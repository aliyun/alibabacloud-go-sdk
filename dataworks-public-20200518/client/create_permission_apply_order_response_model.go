// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePermissionApplyOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePermissionApplyOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePermissionApplyOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreatePermissionApplyOrderResponseBody) *CreatePermissionApplyOrderResponse
	GetBody() *CreatePermissionApplyOrderResponseBody
}

type CreatePermissionApplyOrderResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePermissionApplyOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePermissionApplyOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePermissionApplyOrderResponse) GoString() string {
	return s.String()
}

func (s *CreatePermissionApplyOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePermissionApplyOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePermissionApplyOrderResponse) GetBody() *CreatePermissionApplyOrderResponseBody {
	return s.Body
}

func (s *CreatePermissionApplyOrderResponse) SetHeaders(v map[string]*string) *CreatePermissionApplyOrderResponse {
	s.Headers = v
	return s
}

func (s *CreatePermissionApplyOrderResponse) SetStatusCode(v int32) *CreatePermissionApplyOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePermissionApplyOrderResponse) SetBody(v *CreatePermissionApplyOrderResponseBody) *CreatePermissionApplyOrderResponse {
	s.Body = v
	return s
}

func (s *CreatePermissionApplyOrderResponse) Validate() error {
	return dara.Validate(s)
}
