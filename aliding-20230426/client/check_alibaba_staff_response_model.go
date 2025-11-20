// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAlibabaStaffResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAlibabaStaffResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAlibabaStaffResponse
	GetStatusCode() *int32
	SetBody(v *CheckAlibabaStaffResponseBody) *CheckAlibabaStaffResponse
	GetBody() *CheckAlibabaStaffResponseBody
}

type CheckAlibabaStaffResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAlibabaStaffResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAlibabaStaffResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAlibabaStaffResponse) GoString() string {
	return s.String()
}

func (s *CheckAlibabaStaffResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAlibabaStaffResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAlibabaStaffResponse) GetBody() *CheckAlibabaStaffResponseBody {
	return s.Body
}

func (s *CheckAlibabaStaffResponse) SetHeaders(v map[string]*string) *CheckAlibabaStaffResponse {
	s.Headers = v
	return s
}

func (s *CheckAlibabaStaffResponse) SetStatusCode(v int32) *CheckAlibabaStaffResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAlibabaStaffResponse) SetBody(v *CheckAlibabaStaffResponseBody) *CheckAlibabaStaffResponse {
	s.Body = v
	return s
}

func (s *CheckAlibabaStaffResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
