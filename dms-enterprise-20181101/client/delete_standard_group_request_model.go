// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *DeleteStandardGroupRequest
	GetGroupId() *int64
	SetTid(v int64) *DeleteStandardGroupRequest
	GetTid() *int64
}

type DeleteStandardGroupRequest struct {
	// The security rule set ID. You can call the [ListStandardGroups](https://help.aliyun.com/document_detail/465940.html) operation to obtain the ID of the security rule set.
	//
	// This parameter is required.
	//
	// example:
	//
	// 242***
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 23****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DeleteStandardGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteStandardGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DeleteStandardGroupRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteStandardGroupRequest) SetGroupId(v int64) *DeleteStandardGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteStandardGroupRequest) SetTid(v int64) *DeleteStandardGroupRequest {
	s.Tid = &v
	return s
}

func (s *DeleteStandardGroupRequest) Validate() error {
	return dara.Validate(s)
}
