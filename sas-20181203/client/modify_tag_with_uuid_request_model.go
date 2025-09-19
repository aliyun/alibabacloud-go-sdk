// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTagWithUuidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMachineTypes(v string) *ModifyTagWithUuidRequest
	GetMachineTypes() *string
	SetTagId(v string) *ModifyTagWithUuidRequest
	GetTagId() *string
	SetTagList(v string) *ModifyTagWithUuidRequest
	GetTagList() *string
	SetTarget(v string) *ModifyTagWithUuidRequest
	GetTarget() *string
	SetUuidList(v string) *ModifyTagWithUuidRequest
	GetUuidList() *string
}

type ModifyTagWithUuidRequest struct {
	// The type of the asset to query. If you do not specify this parameter, the tags of all asset types are queried. Valid values:
	//
	// 	- **ecs**: server
	//
	// 	- **cloud_product**: Alibaba Cloud service
	//
	// example:
	//
	// ecs
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
	// The ID of the tag that you want to manage.
	//
	// >  You can call the [DescribeGroupedTags](~~DescribeGroupedTags~~) operation to query the IDs of tags.
	//
	// example:
	//
	// 3897941
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The names of the tags that you want to manage. Separate multiple tag names with commas (,).
	//
	// >  You can call the [DescribeGroupedTags](~~DescribeGroupedTags~~) operation to query the names of tags.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac,ad
	TagList *string `json:"TagList,omitempty" xml:"TagList,omitempty"`
	// The details of the server for which you want to manage the tag. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **Target**: the UUID of the server that you want to add or remove.
	//
	// 	- **targetType**: the method by which the server is added. Valid values:
	//
	//     	- **uuid**: by server
	//
	//     	- **groupId**: by server group
	//
	// 	- **flag**: the operation that you want to perform on the server. Valid values:
	//
	//     	- **del**: removes the tag from the server.
	//
	//     	- **add**: adds the tag to the server.
	//
	// example:
	//
	// [{"target":"inet-7c676676-06fa-442e-90fb-b802e5d6****","targetType":"uuid","flag":"add"}]
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The UUIDs of the servers.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// example:
	//
	// e2737dc3-78f4-4653-a986-dc5fad4b****,c189f0e3-df22-42d5-a73d-02c05667****
	UuidList *string `json:"UuidList,omitempty" xml:"UuidList,omitempty"`
}

func (s ModifyTagWithUuidRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTagWithUuidRequest) GoString() string {
	return s.String()
}

func (s *ModifyTagWithUuidRequest) GetMachineTypes() *string {
	return s.MachineTypes
}

func (s *ModifyTagWithUuidRequest) GetTagId() *string {
	return s.TagId
}

func (s *ModifyTagWithUuidRequest) GetTagList() *string {
	return s.TagList
}

func (s *ModifyTagWithUuidRequest) GetTarget() *string {
	return s.Target
}

func (s *ModifyTagWithUuidRequest) GetUuidList() *string {
	return s.UuidList
}

func (s *ModifyTagWithUuidRequest) SetMachineTypes(v string) *ModifyTagWithUuidRequest {
	s.MachineTypes = &v
	return s
}

func (s *ModifyTagWithUuidRequest) SetTagId(v string) *ModifyTagWithUuidRequest {
	s.TagId = &v
	return s
}

func (s *ModifyTagWithUuidRequest) SetTagList(v string) *ModifyTagWithUuidRequest {
	s.TagList = &v
	return s
}

func (s *ModifyTagWithUuidRequest) SetTarget(v string) *ModifyTagWithUuidRequest {
	s.Target = &v
	return s
}

func (s *ModifyTagWithUuidRequest) SetUuidList(v string) *ModifyTagWithUuidRequest {
	s.UuidList = &v
	return s
}

func (s *ModifyTagWithUuidRequest) Validate() error {
	return dara.Validate(s)
}
