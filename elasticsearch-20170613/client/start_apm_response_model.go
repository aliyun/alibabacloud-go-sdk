// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartApmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartApmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartApmResponse
	GetStatusCode() *int32
	SetBody(v *StartApmResponseBody) *StartApmResponse
	GetBody() *StartApmResponseBody
}

type StartApmResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartApmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartApmResponse) String() string {
	return dara.Prettify(s)
}

func (s StartApmResponse) GoString() string {
	return s.String()
}

func (s *StartApmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartApmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartApmResponse) GetBody() *StartApmResponseBody {
	return s.Body
}

func (s *StartApmResponse) SetHeaders(v map[string]*string) *StartApmResponse {
	s.Headers = v
	return s
}

func (s *StartApmResponse) SetStatusCode(v int32) *StartApmResponse {
	s.StatusCode = &v
	return s
}

func (s *StartApmResponse) SetBody(v *StartApmResponseBody) *StartApmResponse {
	s.Body = v
	return s
}

func (s *StartApmResponse) Validate() error {
	return dara.Validate(s)
}
