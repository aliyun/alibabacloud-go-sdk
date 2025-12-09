// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListTagResourcesResponseBodyData) *ListTagResourcesResponseBody
	GetData() *ListTagResourcesResponseBodyData
	SetRequestId(v string) *ListTagResourcesResponseBody
	GetRequestId() *string
}

type ListTagResourcesResponseBody struct {
	Data *ListTagResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF5189484043
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) GetData() *ListTagResourcesResponseBodyData {
	return s.Data
}

func (s *ListTagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagResourcesResponseBody) SetData(v *ListTagResourcesResponseBodyData) *ListTagResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTagResourcesResponseBodyData struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 24262
	NextToken    *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TagResources *ListTagResourcesResponseBodyDataTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagResourcesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesResponseBodyData) GetTagResources() *ListTagResourcesResponseBodyDataTagResources {
	return s.TagResources
}

func (s *ListTagResourcesResponseBodyData) SetMaxResults(v int32) *ListTagResourcesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesResponseBodyData) SetNextToken(v string) *ListTagResourcesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBodyData) SetTagResources(v *ListTagResourcesResponseBodyDataTagResources) *ListTagResourcesResponseBodyData {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBodyData) Validate() error {
	if s.TagResources != nil {
		if err := s.TagResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTagResourcesResponseBodyDataTagResources struct {
	TagResources []*ListTagResourcesResponseBodyDataTagResourcesTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyDataTagResources) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyDataTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyDataTagResources) GetTagResources() []*ListTagResourcesResponseBodyDataTagResourcesTagResources {
	return s.TagResources
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetTagResources(v []*ListTagResourcesResponseBodyDataTagResourcesTagResources) *ListTagResourcesResponseBodyDataTagResources {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) Validate() error {
	if s.TagResources != nil {
		for _, item := range s.TagResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagResourcesResponseBodyDataTagResourcesTagResources struct {
	// example:
	//
	// cas-upload-xgjcng
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// ALIYUN::CAS::PCACERTIFICATE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// PVDCDC
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyDataTagResourcesTagResources) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyDataTagResourcesTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyDataTagResourcesTagResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTagResourcesResponseBodyDataTagResourcesTagResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesResponseBodyDataTagResourcesTagResources) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagResourcesResponseBodyDataTagResourcesTagResources) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTagResourcesResponseBodyDataTagResourcesTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyDataTagResourcesTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResourcesTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyDataTagResourcesTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResourcesTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyDataTagResourcesTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResourcesTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyDataTagResourcesTagResources {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResourcesTagResources) Validate() error {
	return dara.Validate(s)
}
