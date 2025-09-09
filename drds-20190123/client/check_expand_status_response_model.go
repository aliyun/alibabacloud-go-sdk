// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckExpandStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckExpandStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckExpandStatusResponse
	GetStatusCode() *int32
	SetBody(v *CheckExpandStatusResponseBody) *CheckExpandStatusResponse
	GetBody() *CheckExpandStatusResponseBody
}

type CheckExpandStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckExpandStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckExpandStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckExpandStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckExpandStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckExpandStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckExpandStatusResponse) GetBody() *CheckExpandStatusResponseBody {
	return s.Body
}

func (s *CheckExpandStatusResponse) SetHeaders(v map[string]*string) *CheckExpandStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckExpandStatusResponse) SetStatusCode(v int32) *CheckExpandStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckExpandStatusResponse) SetBody(v *CheckExpandStatusResponseBody) *CheckExpandStatusResponse {
	s.Body = v
	return s
}

func (s *CheckExpandStatusResponse) Validate() error {
	return dara.Validate(s)
}
