// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyApproveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyApproveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyApproveResponse
	GetStatusCode() *int32
	SetBody(v *ApplyApproveResponseBody) *ApplyApproveResponse
	GetBody() *ApplyApproveResponseBody
}

type ApplyApproveResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyApproveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyApproveResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyApproveResponse) GoString() string {
	return s.String()
}

func (s *ApplyApproveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyApproveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyApproveResponse) GetBody() *ApplyApproveResponseBody {
	return s.Body
}

func (s *ApplyApproveResponse) SetHeaders(v map[string]*string) *ApplyApproveResponse {
	s.Headers = v
	return s
}

func (s *ApplyApproveResponse) SetStatusCode(v int32) *ApplyApproveResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyApproveResponse) SetBody(v *ApplyApproveResponseBody) *ApplyApproveResponse {
	s.Body = v
	return s
}

func (s *ApplyApproveResponse) Validate() error {
	return dara.Validate(s)
}
