// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPermissionApplyOrderDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPermissionApplyOrderDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPermissionApplyOrderDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetPermissionApplyOrderDetailResponseBody) *GetPermissionApplyOrderDetailResponse
	GetBody() *GetPermissionApplyOrderDetailResponseBody
}

type GetPermissionApplyOrderDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPermissionApplyOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPermissionApplyOrderDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionApplyOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetPermissionApplyOrderDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPermissionApplyOrderDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPermissionApplyOrderDetailResponse) GetBody() *GetPermissionApplyOrderDetailResponseBody {
	return s.Body
}

func (s *GetPermissionApplyOrderDetailResponse) SetHeaders(v map[string]*string) *GetPermissionApplyOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetPermissionApplyOrderDetailResponse) SetStatusCode(v int32) *GetPermissionApplyOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponse) SetBody(v *GetPermissionApplyOrderDetailResponseBody) *GetPermissionApplyOrderDetailResponse {
	s.Body = v
	return s
}

func (s *GetPermissionApplyOrderDetailResponse) Validate() error {
	return dara.Validate(s)
}
