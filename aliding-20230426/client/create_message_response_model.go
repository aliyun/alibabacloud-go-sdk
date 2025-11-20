// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMessageResponse
	GetStatusCode() *int32
	SetBody(v *CreateMessageResponseBody) *CreateMessageResponse
	GetBody() *CreateMessageResponseBody
}

type CreateMessageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageResponse) GoString() string {
	return s.String()
}

func (s *CreateMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMessageResponse) GetBody() *CreateMessageResponseBody {
	return s.Body
}

func (s *CreateMessageResponse) SetHeaders(v map[string]*string) *CreateMessageResponse {
	s.Headers = v
	return s
}

func (s *CreateMessageResponse) SetStatusCode(v int32) *CreateMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMessageResponse) SetBody(v *CreateMessageResponseBody) *CreateMessageResponse {
	s.Body = v
	return s
}

func (s *CreateMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
