// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReConfigApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReConfigApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReConfigApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ReConfigApplicationResponseBody) *ReConfigApplicationResponse
	GetBody() *ReConfigApplicationResponseBody
}

type ReConfigApplicationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReConfigApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReConfigApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ReConfigApplicationResponse) GoString() string {
	return s.String()
}

func (s *ReConfigApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReConfigApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReConfigApplicationResponse) GetBody() *ReConfigApplicationResponseBody {
	return s.Body
}

func (s *ReConfigApplicationResponse) SetHeaders(v map[string]*string) *ReConfigApplicationResponse {
	s.Headers = v
	return s
}

func (s *ReConfigApplicationResponse) SetStatusCode(v int32) *ReConfigApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ReConfigApplicationResponse) SetBody(v *ReConfigApplicationResponseBody) *ReConfigApplicationResponse {
	s.Body = v
	return s
}

func (s *ReConfigApplicationResponse) Validate() error {
	return dara.Validate(s)
}
