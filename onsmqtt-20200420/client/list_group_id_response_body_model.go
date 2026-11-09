// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListGroupIdResponseBodyData) *ListGroupIdResponseBody
	GetData() []*ListGroupIdResponseBodyData
	SetRequestId(v string) *ListGroupIdResponseBody
	GetRequestId() *string
}

type ListGroupIdResponseBody struct {
	// The details of the returned group IDs.
	Data []*ListGroupIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The unique ID for the request. This is a common parameter.
	//
	// example:
	//
	// 95996EEB-D894-44FA-A87C-940F5CD9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGroupIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupIdResponseBody) GetData() []*ListGroupIdResponseBodyData {
	return s.Data
}

func (s *ListGroupIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupIdResponseBody) SetData(v []*ListGroupIdResponseBodyData) *ListGroupIdResponseBody {
	s.Data = v
	return s
}

func (s *ListGroupIdResponseBody) SetRequestId(v string) *ListGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupIdResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupIdResponseBodyData struct {
	// The creation timestamp of the group ID.
	//
	// example:
	//
	// 1564577317000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The group ID.
	//
	// example:
	//
	// GID_test1
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Indicates whether the instance has an independent namespace. Valid values:
	//
	// - **true**: The instance has an independent namespace. In this case, resource names must be unique within the instance, but they can be the same as resource names in other instances.
	//
	// - **false**: The instance does not have an independent namespace. In this case, resource names must be globally unique across all instances.
	//
	// example:
	//
	// true
	IndependentNaming *bool `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	// The ID of the instance to which the group ID belongs.
	//
	// example:
	//
	// post-cn-45910tj****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The tags attached to the group ID.
	Tags []*ListGroupIdResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timestamp of the last update for the group ID.
	//
	// example:
	//
	// 1564577317000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListGroupIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGroupIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGroupIdResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListGroupIdResponseBodyData) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupIdResponseBodyData) GetIndependentNaming() *bool {
	return s.IndependentNaming
}

func (s *ListGroupIdResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGroupIdResponseBodyData) GetTags() []*ListGroupIdResponseBodyDataTags {
	return s.Tags
}

func (s *ListGroupIdResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListGroupIdResponseBodyData) SetCreateTime(v int64) *ListGroupIdResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListGroupIdResponseBodyData) SetGroupId(v string) *ListGroupIdResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *ListGroupIdResponseBodyData) SetIndependentNaming(v bool) *ListGroupIdResponseBodyData {
	s.IndependentNaming = &v
	return s
}

func (s *ListGroupIdResponseBodyData) SetInstanceId(v string) *ListGroupIdResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListGroupIdResponseBodyData) SetTags(v []*ListGroupIdResponseBodyDataTags) *ListGroupIdResponseBodyData {
	s.Tags = v
	return s
}

func (s *ListGroupIdResponseBodyData) SetUpdateTime(v int64) *ListGroupIdResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListGroupIdResponseBodyData) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupIdResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListGroupIdResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s ListGroupIdResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *ListGroupIdResponseBodyDataTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListGroupIdResponseBodyDataTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListGroupIdResponseBodyDataTags) SetTagKey(v string) *ListGroupIdResponseBodyDataTags {
	s.TagKey = &v
	return s
}

func (s *ListGroupIdResponseBodyDataTags) SetTagValue(v string) *ListGroupIdResponseBodyDataTags {
	s.TagValue = &v
	return s
}

func (s *ListGroupIdResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
