// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsDLQMessageGetByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsDLQMessageGetByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsDLQMessageGetByIdResponse
	GetStatusCode() *int32
	SetBody(v *OnsDLQMessageGetByIdResponseBody) *OnsDLQMessageGetByIdResponse
	GetBody() *OnsDLQMessageGetByIdResponseBody
}

type OnsDLQMessageGetByIdResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsDLQMessageGetByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsDLQMessageGetByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessageGetByIdResponse) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageGetByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsDLQMessageGetByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsDLQMessageGetByIdResponse) GetBody() *OnsDLQMessageGetByIdResponseBody {
	return s.Body
}

func (s *OnsDLQMessageGetByIdResponse) SetHeaders(v map[string]*string) *OnsDLQMessageGetByIdResponse {
	s.Headers = v
	return s
}

func (s *OnsDLQMessageGetByIdResponse) SetStatusCode(v int32) *OnsDLQMessageGetByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponse) SetBody(v *OnsDLQMessageGetByIdResponseBody) *OnsDLQMessageGetByIdResponse {
	s.Body = v
	return s
}

func (s *OnsDLQMessageGetByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
