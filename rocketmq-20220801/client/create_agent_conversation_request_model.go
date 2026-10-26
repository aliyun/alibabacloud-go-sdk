// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateAgentConversationRequest
	GetClientToken() *string
	SetInitialMessage(v *CreateAgentConversationRequestInitialMessage) *CreateAgentConversationRequest
	GetInitialMessage() *CreateAgentConversationRequestInitialMessage
	SetTitle(v string) *CreateAgentConversationRequest
	GetTitle() *string
}

type CreateAgentConversationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// conversation-request-001
	ClientToken    *string                                       `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	InitialMessage *CreateAgentConversationRequestInitialMessage `json:"initialMessage,omitempty" xml:"initialMessage,omitempty" type:"Struct"`
	// example:
	//
	// 消息堆积排查
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateAgentConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentConversationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAgentConversationRequest) GetInitialMessage() *CreateAgentConversationRequestInitialMessage {
	return s.InitialMessage
}

func (s *CreateAgentConversationRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateAgentConversationRequest) SetClientToken(v string) *CreateAgentConversationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAgentConversationRequest) SetInitialMessage(v *CreateAgentConversationRequestInitialMessage) *CreateAgentConversationRequest {
	s.InitialMessage = v
	return s
}

func (s *CreateAgentConversationRequest) SetTitle(v string) *CreateAgentConversationRequest {
	s.Title = &v
	return s
}

func (s *CreateAgentConversationRequest) Validate() error {
	if s.InitialMessage != nil {
		if err := s.InitialMessage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentConversationRequestInitialMessage struct {
	// example:
	//
	// 告警发生于 10:00
	ClientContext *string                                               `json:"clientContext,omitempty" xml:"clientContext,omitempty"`
	Images        []*CreateAgentConversationRequestInitialMessageImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 请排查实例消息堆积原因
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// ticket_request
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// example:
	//
	// ticket-request-example-001
	TicketRequestId *string `json:"ticketRequestId,omitempty" xml:"ticketRequestId,omitempty"`
}

func (s CreateAgentConversationRequestInitialMessage) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentConversationRequestInitialMessage) GoString() string {
	return s.String()
}

func (s *CreateAgentConversationRequestInitialMessage) GetClientContext() *string {
	return s.ClientContext
}

func (s *CreateAgentConversationRequestInitialMessage) GetImages() []*CreateAgentConversationRequestInitialMessageImages {
	return s.Images
}

func (s *CreateAgentConversationRequestInitialMessage) GetMessage() *string {
	return s.Message
}

func (s *CreateAgentConversationRequestInitialMessage) GetMessageType() *string {
	return s.MessageType
}

func (s *CreateAgentConversationRequestInitialMessage) GetTicketRequestId() *string {
	return s.TicketRequestId
}

func (s *CreateAgentConversationRequestInitialMessage) SetClientContext(v string) *CreateAgentConversationRequestInitialMessage {
	s.ClientContext = &v
	return s
}

func (s *CreateAgentConversationRequestInitialMessage) SetImages(v []*CreateAgentConversationRequestInitialMessageImages) *CreateAgentConversationRequestInitialMessage {
	s.Images = v
	return s
}

func (s *CreateAgentConversationRequestInitialMessage) SetMessage(v string) *CreateAgentConversationRequestInitialMessage {
	s.Message = &v
	return s
}

func (s *CreateAgentConversationRequestInitialMessage) SetMessageType(v string) *CreateAgentConversationRequestInitialMessage {
	s.MessageType = &v
	return s
}

func (s *CreateAgentConversationRequestInitialMessage) SetTicketRequestId(v string) *CreateAgentConversationRequestInitialMessage {
	s.TicketRequestId = &v
	return s
}

func (s *CreateAgentConversationRequestInitialMessage) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAgentConversationRequestInitialMessageImages struct {
	// This parameter is required.
	//
	// example:
	//
	// iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAC0lEQVR4nGNgAAIAAAUAAXpeqz8AAAAASUVORK5CYII=
	DataBase64 *string `json:"dataBase64,omitempty" xml:"dataBase64,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// image/png
	MediaType *string `json:"mediaType,omitempty" xml:"mediaType,omitempty"`
}

func (s CreateAgentConversationRequestInitialMessageImages) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentConversationRequestInitialMessageImages) GoString() string {
	return s.String()
}

func (s *CreateAgentConversationRequestInitialMessageImages) GetDataBase64() *string {
	return s.DataBase64
}

func (s *CreateAgentConversationRequestInitialMessageImages) GetMediaType() *string {
	return s.MediaType
}

func (s *CreateAgentConversationRequestInitialMessageImages) SetDataBase64(v string) *CreateAgentConversationRequestInitialMessageImages {
	s.DataBase64 = &v
	return s
}

func (s *CreateAgentConversationRequestInitialMessageImages) SetMediaType(v string) *CreateAgentConversationRequestInitialMessageImages {
	s.MediaType = &v
	return s
}

func (s *CreateAgentConversationRequestInitialMessageImages) Validate() error {
	return dara.Validate(s)
}
