// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageCachesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageCacheName(v string) *ListImageCachesRequest
	GetImageCacheName() *string
	SetMaxResults(v int32) *ListImageCachesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListImageCachesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListImageCachesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListImageCachesRequest
	GetResourceGroupId() *string
	SetStatus(v string) *ListImageCachesRequest
	GetStatus() *string
	SetTags(v []*ListImageCachesRequestTags) *ListImageCachesRequest
	GetTags() []*ListImageCachesRequestTags
}

type ListImageCachesRequest struct {
	// example:
	//
	// my-imc
	ImageCacheName *string `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdDWBF*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aekzh43v*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// Ready
	Status *string                       `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*ListImageCachesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListImageCachesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImageCachesRequest) GoString() string {
	return s.String()
}

func (s *ListImageCachesRequest) GetImageCacheName() *string {
	return s.ImageCacheName
}

func (s *ListImageCachesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListImageCachesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListImageCachesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListImageCachesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListImageCachesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListImageCachesRequest) GetTags() []*ListImageCachesRequestTags {
	return s.Tags
}

func (s *ListImageCachesRequest) SetImageCacheName(v string) *ListImageCachesRequest {
	s.ImageCacheName = &v
	return s
}

func (s *ListImageCachesRequest) SetMaxResults(v int32) *ListImageCachesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListImageCachesRequest) SetNextToken(v string) *ListImageCachesRequest {
	s.NextToken = &v
	return s
}

func (s *ListImageCachesRequest) SetRegionId(v string) *ListImageCachesRequest {
	s.RegionId = &v
	return s
}

func (s *ListImageCachesRequest) SetResourceGroupId(v string) *ListImageCachesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListImageCachesRequest) SetStatus(v string) *ListImageCachesRequest {
	s.Status = &v
	return s
}

func (s *ListImageCachesRequest) SetTags(v []*ListImageCachesRequestTags) *ListImageCachesRequest {
	s.Tags = v
	return s
}

func (s *ListImageCachesRequest) Validate() error {
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

type ListImageCachesRequestTags struct {
	// example:
	//
	// imc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListImageCachesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListImageCachesRequestTags) GoString() string {
	return s.String()
}

func (s *ListImageCachesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListImageCachesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListImageCachesRequestTags) SetKey(v string) *ListImageCachesRequestTags {
	s.Key = &v
	return s
}

func (s *ListImageCachesRequestTags) SetValue(v string) *ListImageCachesRequestTags {
	s.Value = &v
	return s
}

func (s *ListImageCachesRequestTags) Validate() error {
	return dara.Validate(s)
}
