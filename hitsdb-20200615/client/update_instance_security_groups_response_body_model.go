// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceSecurityGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateInstanceSecurityGroupsResponseBody
	GetRequestId() *string
}

type UpdateInstanceSecurityGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceSecurityGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceSecurityGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceSecurityGroupsResponseBody) SetRequestId(v string) *UpdateInstanceSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceSecurityGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
