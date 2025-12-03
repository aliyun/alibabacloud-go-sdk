// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMessagesResponse
	GetStatusCode() *int32
	SetBody(v *GetMessagesResponseBody) *GetMessagesResponse
	GetBody() *GetMessagesResponseBody
}

type GetMessagesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMessagesResponse) GoString() string {
	return s.String()
}

func (s *GetMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMessagesResponse) GetBody() *GetMessagesResponseBody {
	return s.Body
}

func (s *GetMessagesResponse) SetHeaders(v map[string]*string) *GetMessagesResponse {
	s.Headers = v
	return s
}

func (s *GetMessagesResponse) SetStatusCode(v int32) *GetMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessagesResponse) SetBody(v *GetMessagesResponseBody) *GetMessagesResponse {
	s.Body = v
	return s
}

func (s *GetMessagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
