// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCloudPermissionTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableCloudPermissionTemplateResponseBody
	GetRequestId() *string
}

type DisableCloudPermissionTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableCloudPermissionTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableCloudPermissionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DisableCloudPermissionTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableCloudPermissionTemplateResponseBody) SetRequestId(v string) *DisableCloudPermissionTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableCloudPermissionTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
