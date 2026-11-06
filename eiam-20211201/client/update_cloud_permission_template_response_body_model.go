// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudPermissionTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudPermissionTemplateResponseBody
	GetRequestId() *string
}

type UpdateCloudPermissionTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCloudPermissionTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudPermissionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudPermissionTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudPermissionTemplateResponseBody) SetRequestId(v string) *UpdateCloudPermissionTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudPermissionTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
