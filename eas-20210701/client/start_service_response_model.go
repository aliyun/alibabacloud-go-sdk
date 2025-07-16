// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartServiceResponse
	GetStatusCode() *int32
	SetBody(v *StartServiceResponseBody) *StartServiceResponse
	GetBody() *StartServiceResponseBody
}

type StartServiceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartServiceResponse) GoString() string {
	return s.String()
}

func (s *StartServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartServiceResponse) GetBody() *StartServiceResponseBody {
	return s.Body
}

func (s *StartServiceResponse) SetHeaders(v map[string]*string) *StartServiceResponse {
	s.Headers = v
	return s
}

func (s *StartServiceResponse) SetStatusCode(v int32) *StartServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartServiceResponse) SetBody(v *StartServiceResponseBody) *StartServiceResponse {
	s.Body = v
	return s
}

func (s *StartServiceResponse) Validate() error {
	return dara.Validate(s)
}
