// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDrdsDbNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDrdsDbNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDrdsDbNameResponse
	GetStatusCode() *int32
	SetBody(v *CheckDrdsDbNameResponseBody) *CheckDrdsDbNameResponse
	GetBody() *CheckDrdsDbNameResponseBody
}

type CheckDrdsDbNameResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDrdsDbNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDrdsDbNameResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDrdsDbNameResponse) GoString() string {
	return s.String()
}

func (s *CheckDrdsDbNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDrdsDbNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDrdsDbNameResponse) GetBody() *CheckDrdsDbNameResponseBody {
	return s.Body
}

func (s *CheckDrdsDbNameResponse) SetHeaders(v map[string]*string) *CheckDrdsDbNameResponse {
	s.Headers = v
	return s
}

func (s *CheckDrdsDbNameResponse) SetStatusCode(v int32) *CheckDrdsDbNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDrdsDbNameResponse) SetBody(v *CheckDrdsDbNameResponseBody) *CheckDrdsDbNameResponse {
	s.Body = v
	return s
}

func (s *CheckDrdsDbNameResponse) Validate() error {
	return dara.Validate(s)
}
