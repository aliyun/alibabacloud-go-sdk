// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudAccountRoleDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudAccountRoleDescriptionResponseBody
	GetRequestId() *string
}

type UpdateCloudAccountRoleDescriptionResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCloudAccountRoleDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudAccountRoleDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudAccountRoleDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudAccountRoleDescriptionResponseBody) SetRequestId(v string) *UpdateCloudAccountRoleDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudAccountRoleDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
