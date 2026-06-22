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
	// The new property information of the server group after modification. The following parameters are described:
	//
	// - **groupFlag**: The type of the server group. Valid values: **0*	- (default group) | **1*	- (other group).
	//
	// - **groupId**: The ID of the server group.
	//
	// - **groupIndex**: The sorting number of the server group. Sorted in ascending order.
	//
	// - **groupName**: The name of the server group. Set this parameter to the new name of the server group. The new name must be different from the original name.
	//
	// > Call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to obtain the values of groupFlag and groupId. The values of groupFlag and groupId cannot be modified. Only the value of groupName can be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"groupFlag":1,"groupId":8436682,"groupIndex":1,"groupName":"example"}]
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
