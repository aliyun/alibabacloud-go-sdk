// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelChatResponse
	GetStatusCode() *int32
	SetBody(v *CancelChatResponseBody) *CancelChatResponse
	GetBody() *CancelChatResponseBody
}

type CancelChatResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelChatResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelChatResponse) GoString() string {
	return s.String()
}

func (s *CancelChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelChatResponse) GetBody() *CancelChatResponseBody {
	return s.Body
}

func (s *CancelChatResponse) SetHeaders(v map[string]*string) *CancelChatResponse {
	s.Headers = v
	return s
}

func (s *CancelChatResponse) SetStatusCode(v int32) *CancelChatResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelChatResponse) SetBody(v *CancelChatResponseBody) *CancelChatResponse {
	s.Body = v
	return s
}

func (s *CancelChatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
