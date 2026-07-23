// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShoppingAssistantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ShoppingAssistantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ShoppingAssistantResponse
	GetStatusCode() *int32
	SetId(v string) *ShoppingAssistantResponse
	GetId() *string
	SetEvent(v string) *ShoppingAssistantResponse
	GetEvent() *string
	SetBody(v *ShoppingAssistantResponseBody) *ShoppingAssistantResponse
	GetBody() *ShoppingAssistantResponseBody
}

type ShoppingAssistantResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                        `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                        `json:"event,omitempty" xml:"event,omitempty"`
	Body       *ShoppingAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ShoppingAssistantResponse) String() string {
	return dara.Prettify(s)
}

func (s ShoppingAssistantResponse) GoString() string {
	return s.String()
}

func (s *ShoppingAssistantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ShoppingAssistantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ShoppingAssistantResponse) GetId() *string {
	return s.Id
}

func (s *ShoppingAssistantResponse) GetEvent() *string {
	return s.Event
}

func (s *ShoppingAssistantResponse) GetBody() *ShoppingAssistantResponseBody {
	return s.Body
}

func (s *ShoppingAssistantResponse) SetHeaders(v map[string]*string) *ShoppingAssistantResponse {
	s.Headers = v
	return s
}

func (s *ShoppingAssistantResponse) SetStatusCode(v int32) *ShoppingAssistantResponse {
	s.StatusCode = &v
	return s
}

func (s *ShoppingAssistantResponse) SetId(v string) *ShoppingAssistantResponse {
	s.Id = &v
	return s
}

func (s *ShoppingAssistantResponse) SetEvent(v string) *ShoppingAssistantResponse {
	s.Event = &v
	return s
}

func (s *ShoppingAssistantResponse) SetBody(v *ShoppingAssistantResponseBody) *ShoppingAssistantResponse {
	s.Body = v
	return s
}

func (s *ShoppingAssistantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
