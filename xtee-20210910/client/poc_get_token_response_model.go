// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPocGetTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PocGetTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PocGetTokenResponse
	GetStatusCode() *int32
	SetBody(v *PocGetTokenResponseBody) *PocGetTokenResponse
	GetBody() *PocGetTokenResponseBody
}

type PocGetTokenResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PocGetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PocGetTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s PocGetTokenResponse) GoString() string {
	return s.String()
}

func (s *PocGetTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PocGetTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PocGetTokenResponse) GetBody() *PocGetTokenResponseBody {
	return s.Body
}

func (s *PocGetTokenResponse) SetHeaders(v map[string]*string) *PocGetTokenResponse {
	s.Headers = v
	return s
}

func (s *PocGetTokenResponse) SetStatusCode(v int32) *PocGetTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *PocGetTokenResponse) SetBody(v *PocGetTokenResponseBody) *PocGetTokenResponse {
	s.Body = v
	return s
}

func (s *PocGetTokenResponse) Validate() error {
	return dara.Validate(s)
}
