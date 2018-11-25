package types

import "github.com/mailru/easyjson"

// все типы, в формате которых нужно пересылать сообщения с клиентом.

// первый уровень
//easyjson:json
type Event struct {
	Method    string              `json:"method,required"`
	Parameter easyjson.RawMessage `json:"parameter,required"`
}

// "send"
type Message struct {
	From      *string  `json:"from"`          // логин пользователя
	To        *string  `json:"to"`            // логин пользователя
	Text      string   `json:"text,required"` // сообщение пользователя
	Reply     *Message `json:"reply"`         // пересылаемое сообщение, во вторую очередь
	Time      string   `json:"time"`          // в формате iso-8601
	Id        uint     `json:"id"`
	IsHistory bool     `json:"is_history,omitempty"` // falsteб если посылает горутина получившая, true если из базы.
}

//easyjson:json
type Messages []Message

// "history"
//easyjson:json
type HistoryRequest struct {
	From  *string `json:"from"`  // логин пользователя
	After string  `json:"after"` // id последнего сообщения, которое не надо пересылать.
}