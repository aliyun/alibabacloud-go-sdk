// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckFieldLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckFieldLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckFieldLimitResponse
	GetStatusCode() *int32
	SetBody(v *CheckFieldLimitResponseBody) *CheckFieldLimitResponse
	GetBody() *CheckFieldLimitResponseBody
}

type CheckFieldLimitResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckFieldLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckFieldLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckFieldLimitResponse) GoString() string {
	return s.String()
}

func (s *CheckFieldLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckFieldLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckFieldLimitResponse) GetBody() *CheckFieldLimitResponseBody {
	return s.Body
}

func (s *CheckFieldLimitResponse) SetHeaders(v map[string]*string) *CheckFieldLimitResponse {
	s.Headers = v
	return s
}

func (s *CheckFieldLimitResponse) SetStatusCode(v int32) *CheckFieldLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckFieldLimitResponse) SetBody(v *CheckFieldLimitResponseBody) *CheckFieldLimitResponse {
	s.Body = v
	return s
}

func (s *CheckFieldLimitResponse) Validate() error {
	return dara.Validate(s)
}
