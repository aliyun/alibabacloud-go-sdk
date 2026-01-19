// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageCachesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageCaches(v []*ListImageCachesResponseBodyImageCaches) *ListImageCachesResponseBody
	GetImageCaches() []*ListImageCachesResponseBodyImageCaches
	SetMaxResults(v int32) *ListImageCachesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListImageCachesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListImageCachesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListImageCachesResponseBody
	GetTotalCount() *int32
}

type ListImageCachesResponseBody struct {
	ImageCaches []*ListImageCachesResponseBodyImageCaches `json:"ImageCaches,omitempty" xml:"ImageCaches,omitempty" type:"Repeated"`
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdDWBF*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 0E234675-3465-4CC3-9D0F-9A864BC*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListImageCachesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImageCachesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageCachesResponseBody) GetImageCaches() []*ListImageCachesResponseBodyImageCaches {
	return s.ImageCaches
}

func (s *ListImageCachesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListImageCachesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListImageCachesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImageCachesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListImageCachesResponseBody) SetImageCaches(v []*ListImageCachesResponseBodyImageCaches) *ListImageCachesResponseBody {
	s.ImageCaches = v
	return s
}

func (s *ListImageCachesResponseBody) SetMaxResults(v int32) *ListImageCachesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListImageCachesResponseBody) SetNextToken(v string) *ListImageCachesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListImageCachesResponseBody) SetRequestId(v string) *ListImageCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageCachesResponseBody) SetTotalCount(v int32) *ListImageCachesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListImageCachesResponseBody) Validate() error {
	if s.ImageCaches != nil {
		for _, item := range s.ImageCaches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListImageCachesResponseBodyImageCaches struct {
	// example:
	//
	// 2025-**-**T07:55:25Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// imc-bp1dj*****
	ImageCacheId *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	// example:
	//
	// my-imc
	ImageCacheName *string   `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	Images         []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-**-**T07:58:25Z
	ReadyTime *string `json:"ReadyTime,omitempty" xml:"ReadyTime,omitempty"`
	// example:
	//
	// rg-aekzh43v*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 8
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// Ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListImageCachesResponseBodyImageCaches) String() string {
	return dara.Prettify(s)
}

func (s ListImageCachesResponseBodyImageCaches) GoString() string {
	return s.String()
}

func (s *ListImageCachesResponseBodyImageCaches) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListImageCachesResponseBodyImageCaches) GetImageCacheId() *string {
	return s.ImageCacheId
}

func (s *ListImageCachesResponseBodyImageCaches) GetImageCacheName() *string {
	return s.ImageCacheName
}

func (s *ListImageCachesResponseBodyImageCaches) GetImages() []*string {
	return s.Images
}

func (s *ListImageCachesResponseBodyImageCaches) GetReadyTime() *string {
	return s.ReadyTime
}

func (s *ListImageCachesResponseBodyImageCaches) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListImageCachesResponseBodyImageCaches) GetSize() *int32 {
	return s.Size
}

func (s *ListImageCachesResponseBodyImageCaches) GetStatus() *string {
	return s.Status
}

func (s *ListImageCachesResponseBodyImageCaches) SetCreateTime(v string) *ListImageCachesResponseBodyImageCaches {
	s.CreateTime = &v
	return s
}

func (s *ListImageCachesResponseBodyImageCaches) SetImageCacheId(v string) *ListImageCachesResponseBodyImageCaches {
	s.ImageCacheId = &v
	return s
}

func (s *ListImageCachesResponseBodyImageCaches) SetImageCacheName(v string) *ListImageCachesResponseBodyImageCaches {
	s.ImageCacheName = &v
	return s
}

func (s *ListImageCachesResponseBodyImageCaches) SetImages(v []*string) *ListImageCachesResponseBodyImageCaches {
	s.Images = v
	return s
}

func (s *ListImageCachesResponseBodyImageCaches) SetReadyTime(v string) *ListImageCachesResponseBodyImageCaches {
	s.ReadyTime = &v
	return s
}

func (s *ListImageCachesResponseBodyImageCaches) SetResourceGroupId(v string) *ListImageCachesResponseBodyImageCaches {
	s.ResourceGroupId = &v
	return s
}

func (s *ListImageCachesResponseBodyImageCaches) SetSize(v int32) *ListImageCachesResponseBodyImageCaches {
	s.Size = &v
	return s
}

func (s *ListImageCachesResponseBodyImageCaches) SetStatus(v string) *ListImageCachesResponseBodyImageCaches {
	s.Status = &v
	return s
}

func (s *ListImageCachesResponseBodyImageCaches) Validate() error {
	return dara.Validate(s)
}
