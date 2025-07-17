// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateStandardGroupRequest
	GetDescription() *string
	SetGroupId(v int64) *UpdateStandardGroupRequest
	GetGroupId() *int64
	SetGroupName(v string) *UpdateStandardGroupRequest
	GetGroupName() *string
	SetTid(v int64) *UpdateStandardGroupRequest
	GetTid() *int64
}

type UpdateStandardGroupRequest struct {
	// The description of the security rule set.
	//
	// This parameter is required.
	//
	// example:
	//
	// Production Environment test rules
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The security rule set ID. You can call the [ListStandardGroups](https://help.aliyun.com/document_detail/465940.html) operation to obtain the ID of the security rule set.
	//
	// This parameter is required.
	//
	// example:
	//
	// 242***
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the security rule set.
	//
	// This parameter is required.
	//
	// example:
	//
	// poc_test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The tenant ID.
	//
	// >  To view the tenant ID, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 23****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateStandardGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateStandardGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateStandardGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdateStandardGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateStandardGroupRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateStandardGroupRequest) SetDescription(v string) *UpdateStandardGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateStandardGroupRequest) SetGroupId(v int64) *UpdateStandardGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateStandardGroupRequest) SetGroupName(v string) *UpdateStandardGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateStandardGroupRequest) SetTid(v int64) *UpdateStandardGroupRequest {
	s.Tid = &v
	return s
}

func (s *UpdateStandardGroupRequest) Validate() error {
	return dara.Validate(s)
}
