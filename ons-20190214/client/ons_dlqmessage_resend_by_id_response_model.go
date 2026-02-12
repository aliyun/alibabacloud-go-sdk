// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsDLQMessageResendByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsDLQMessageResendByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsDLQMessageResendByIdResponse
	GetStatusCode() *int32
	SetBody(v *OnsDLQMessageResendByIdResponseBody) *OnsDLQMessageResendByIdResponse
	GetBody() *OnsDLQMessageResendByIdResponseBody
}

type OnsDLQMessageResendByIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsDLQMessageResendByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsDLQMessageResendByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessageResendByIdResponse) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageResendByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsDLQMessageResendByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsDLQMessageResendByIdResponse) GetBody() *OnsDLQMessageResendByIdResponseBody {
	return s.Body
}

func (s *OnsDLQMessageResendByIdResponse) SetHeaders(v map[string]*string) *OnsDLQMessageResendByIdResponse {
	s.Headers = v
	return s
}

func (s *OnsDLQMessageResendByIdResponse) SetStatusCode(v int32) *OnsDLQMessageResendByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsDLQMessageResendByIdResponse) SetBody(v *OnsDLQMessageResendByIdResponseBody) *OnsDLQMessageResendByIdResponse {
	s.Body = v
	return s
}

func (s *OnsDLQMessageResendByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
