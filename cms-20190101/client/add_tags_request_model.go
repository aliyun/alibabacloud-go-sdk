// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupIds(v []*string) *AddTagsRequest
	GetGroupIds() []*string
	SetRegionId(v string) *AddTagsRequest
	GetRegionId() *string
	SetTag(v []*AddTagsRequestTag) *AddTagsRequest
	GetTag() []*AddTagsRequestTag
}

type AddTagsRequest struct {
	// The IDs of the application groups.
	//
	// The value of N can be from 1 to 20.
	//
	// For information about how to obtain application group IDs, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/2513168.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 7301****
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// A list of tags.
	//
	// This parameter is required.
	Tag []*AddTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s AddTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTagsRequest) GoString() string {
	return s.String()
}

func (s *AddTagsRequest) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *AddTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddTagsRequest) GetTag() []*AddTagsRequestTag {
	return s.Tag
}

func (s *AddTagsRequest) SetGroupIds(v []*string) *AddTagsRequest {
	s.GroupIds = v
	return s
}

func (s *AddTagsRequest) SetRegionId(v string) *AddTagsRequest {
	s.RegionId = &v
	return s
}

func (s *AddTagsRequest) SetTag(v []*AddTagsRequestTag) *AddTagsRequest {
	s.Tag = v
	return s
}

func (s *AddTagsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddTagsRequestTag struct {
	// The tag key.
	//
	// The value of N can be from 1 to 3. The tag key must be 1 to 64 characters in length.
	//
	// Enter a new tag key or use an existing tag key from Cloud Monitor. For information about how to obtain tag keys, see [DescribeTagKeyList](https://help.aliyun.com/document_detail/2513189.html).
	//
	// > The tag key cannot start with `aliyun` or `acs:`. The `Tag.N.Key` and `Tag.N.Value` parameters must be specified together.
	//
	// This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The value of N can be from 1 to 3. The tag value must be 1 to 64 characters in length.
	//
	// Enter a new tag value or use an existing tag value from Cloud Monitor. For information about how to obtain tag values, see [DescribeTagValueList](https://help.aliyun.com/document_detail/2513188.html).
	//
	// > The tag value cannot start with `aliyun` or `acs:`. The `Tag.N.Key` and `Tag.N.Value` parameters must be specified together.
	//
	// This parameter is required.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddTagsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AddTagsRequestTag) GoString() string {
	return s.String()
}

func (s *AddTagsRequestTag) GetKey() *string {
	return s.Key
}

func (s *AddTagsRequestTag) GetValue() *string {
	return s.Value
}

func (s *AddTagsRequestTag) SetKey(v string) *AddTagsRequestTag {
	s.Key = &v
	return s
}

func (s *AddTagsRequestTag) SetValue(v string) *AddTagsRequestTag {
	s.Value = &v
	return s
}

func (s *AddTagsRequestTag) Validate() error {
	return dara.Validate(s)
}
