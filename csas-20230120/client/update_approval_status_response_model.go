// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApprovalStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApprovalStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApprovalStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApprovalStatusResponseBody) *UpdateApprovalStatusResponse
	GetBody() *UpdateApprovalStatusResponseBody
}

type UpdateApprovalStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApprovalStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApprovalStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateApprovalStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApprovalStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApprovalStatusResponse) GetBody() *UpdateApprovalStatusResponseBody {
	return s.Body
}

func (s *UpdateApprovalStatusResponse) SetHeaders(v map[string]*string) *UpdateApprovalStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateApprovalStatusResponse) SetStatusCode(v int32) *UpdateApprovalStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApprovalStatusResponse) SetBody(v *UpdateApprovalStatusResponseBody) *UpdateApprovalStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateApprovalStatusResponse) Validate() error {
	return dara.Validate(s)
}
