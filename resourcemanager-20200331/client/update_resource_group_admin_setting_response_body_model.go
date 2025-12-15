// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceGroupAdminSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateResourceGroupAdminSettingResponseBody
	GetRequestId() *string
}

type UpdateResourceGroupAdminSettingResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceGroupAdminSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupAdminSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupAdminSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceGroupAdminSettingResponseBody) SetRequestId(v string) *UpdateResourceGroupAdminSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceGroupAdminSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
