// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApprovePermissionApplyOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApprovePermissionApplyOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApprovePermissionApplyOrderResponse
	GetStatusCode() *int32
	SetBody(v *ApprovePermissionApplyOrderResponseBody) *ApprovePermissionApplyOrderResponse
	GetBody() *ApprovePermissionApplyOrderResponseBody
}

type ApprovePermissionApplyOrderResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApprovePermissionApplyOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApprovePermissionApplyOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s ApprovePermissionApplyOrderResponse) GoString() string {
	return s.String()
}

func (s *ApprovePermissionApplyOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApprovePermissionApplyOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApprovePermissionApplyOrderResponse) GetBody() *ApprovePermissionApplyOrderResponseBody {
	return s.Body
}

func (s *ApprovePermissionApplyOrderResponse) SetHeaders(v map[string]*string) *ApprovePermissionApplyOrderResponse {
	s.Headers = v
	return s
}

func (s *ApprovePermissionApplyOrderResponse) SetStatusCode(v int32) *ApprovePermissionApplyOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ApprovePermissionApplyOrderResponse) SetBody(v *ApprovePermissionApplyOrderResponseBody) *ApprovePermissionApplyOrderResponse {
	s.Body = v
	return s
}

func (s *ApprovePermissionApplyOrderResponse) Validate() error {
	return dara.Validate(s)
}
