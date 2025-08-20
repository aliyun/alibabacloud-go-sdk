// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStackGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteStackGroupRequest
	GetRegionId() *string
	SetStackGroupName(v string) *DeleteStackGroupRequest
	GetStackGroupName() *string
}

type DeleteStackGroupRequest struct {
	// The ID of the region to which the stack group belongs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stack group. The name must be unique in a region.
	//
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). The name must start with a digit or a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyStackGroup
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
}

func (s DeleteStackGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStackGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteStackGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteStackGroupRequest) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *DeleteStackGroupRequest) SetRegionId(v string) *DeleteStackGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStackGroupRequest) SetStackGroupName(v string) *DeleteStackGroupRequest {
	s.StackGroupName = &v
	return s
}

func (s *DeleteStackGroupRequest) Validate() error {
	return dara.Validate(s)
}
