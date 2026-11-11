// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateCloudAccountPrivilegeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeactivateCloudAccountPrivilegeResponseBody
	GetRequestId() *string
}

type DeactivateCloudAccountPrivilegeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeactivateCloudAccountPrivilegeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeactivateCloudAccountPrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *DeactivateCloudAccountPrivilegeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeactivateCloudAccountPrivilegeResponseBody) SetRequestId(v string) *DeactivateCloudAccountPrivilegeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeactivateCloudAccountPrivilegeResponseBody) Validate() error {
	return dara.Validate(s)
}
