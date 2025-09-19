// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGroupPropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ModifyGroupPropertyRequest
	GetData() *string
}

type ModifyGroupPropertyRequest struct {
	// The new attributes of the server group. You can specify the following parameters to configure the attributes:
	//
	// 	- **groupFlag**: the type of the server group. Valid values: 0 and 1. The value **0*	- specifies the Default server group. The value **1*	- specifies other server groups.
	//
	// 	- **groupId**: the ID of the server group.
	//
	// 	- **groupIndex**: no meaning. You can leave this parameter empty.
	//
	// 	- **groupName**: the name of the server group. The value is the new name of the server group. The new name cannot be the same as the original name of the server group.
	//
	// >  You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to obtain the values of the groupFlag and groupId parameters. You cannot change the value of the groupFlag or groupId parameter. You can change only the value of the groupName parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"groupFlag":1,"groupId":8436682,"groupIndex":,"groupName":"example"}]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ModifyGroupPropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGroupPropertyRequest) GoString() string {
	return s.String()
}

func (s *ModifyGroupPropertyRequest) GetData() *string {
	return s.Data
}

func (s *ModifyGroupPropertyRequest) SetData(v string) *ModifyGroupPropertyRequest {
	s.Data = &v
	return s
}

func (s *ModifyGroupPropertyRequest) Validate() error {
	return dara.Validate(s)
}
