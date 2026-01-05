// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTagsResponseBody
	GetRequestId() *string
	SetTags(v []*ListTagsResponseBodyTags) *ListTagsResponseBody
	GetTags() []*ListTagsResponseBodyTags
	SetTotalCount(v int32) *ListTagsResponseBody
	GetTotalCount() *int32
}

type ListTagsResponseBody struct {
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*ListTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagsResponseBody) GetTags() []*ListTagsResponseBodyTags {
	return s.Tags
}

func (s *ListTagsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTagsResponseBody) SetRequestId(v string) *ListTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagsResponseBody) SetTags(v []*ListTagsResponseBodyTags) *ListTagsResponseBody {
	s.Tags = v
	return s
}

func (s *ListTagsResponseBody) SetTotalCount(v int32) *ListTagsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagsResponseBody) Validate() error {
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

type ListTagsResponseBodyTags struct {
	// example:
	//
	// 2014-10-02T15:01:23Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2014-10-02T15:01:23Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// dataset_version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 1557702098194904
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// inst-my1tk3jggooi5uwks
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 276065346797410278
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// v0.1.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagsResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTags) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListTagsResponseBodyTags) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListTagsResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *ListTagsResponseBodyTags) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListTagsResponseBodyTags) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTagsResponseBodyTags) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagsResponseBodyTags) GetUserId() *string {
	return s.UserId
}

func (s *ListTagsResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *ListTagsResponseBodyTags) SetGmtCreateTime(v string) *ListTagsResponseBodyTags {
	s.GmtCreateTime = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetGmtModifiedTime(v string) *ListTagsResponseBodyTags {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetKey(v string) *ListTagsResponseBodyTags {
	s.Key = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetOwnerId(v string) *ListTagsResponseBodyTags {
	s.OwnerId = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetResourceId(v string) *ListTagsResponseBodyTags {
	s.ResourceId = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetResourceType(v string) *ListTagsResponseBodyTags {
	s.ResourceType = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetUserId(v string) *ListTagsResponseBodyTags {
	s.UserId = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetValue(v string) *ListTagsResponseBodyTags {
	s.Value = &v
	return s
}

func (s *ListTagsResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
