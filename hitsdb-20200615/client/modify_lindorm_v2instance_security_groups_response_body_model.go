// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLindormV2InstanceSecurityGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLindormV2InstanceSecurityGroupsResponseBody
	GetRequestId() *string
}

type ModifyLindormV2InstanceSecurityGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLindormV2InstanceSecurityGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLindormV2InstanceSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2InstanceSecurityGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLindormV2InstanceSecurityGroupsResponseBody) SetRequestId(v string) *ModifyLindormV2InstanceSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLindormV2InstanceSecurityGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
