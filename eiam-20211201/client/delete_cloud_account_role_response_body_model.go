// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudAccountRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCloudAccountRoleResponseBody
	GetRequestId() *string
}

type DeleteCloudAccountRoleResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCloudAccountRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudAccountRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudAccountRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudAccountRoleResponseBody) SetRequestId(v string) *DeleteCloudAccountRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudAccountRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
