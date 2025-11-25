// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeTagKeysResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTagKeysResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeTagKeysResponseBody
	GetRequestId() *string
	SetTagKeys(v []*DescribeTagKeysResponseBodyTagKeys) *DescribeTagKeysResponseBody
	GetTagKeys() []*DescribeTagKeysResponseBodyTagKeys
	SetTotalCount(v int32) *DescribeTagKeysResponseBody
	GetTotalCount() *int32
}

type DescribeTagKeysResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6623EA1F-30FB-5BC8-BEC9-74D55F6F08F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the tag keys.
	TagKeys []*DescribeTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTagKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTagKeysResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTagKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagKeysResponseBody) GetTagKeys() []*DescribeTagKeysResponseBodyTagKeys {
	return s.TagKeys
}

func (s *DescribeTagKeysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTagKeysResponseBody) SetPageNumber(v int32) *DescribeTagKeysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagKeysResponseBody) SetPageSize(v int32) *DescribeTagKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTagKeysResponseBody) SetRequestId(v string) *DescribeTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagKeysResponseBody) SetTagKeys(v []*DescribeTagKeysResponseBodyTagKeys) *DescribeTagKeysResponseBody {
	s.TagKeys = v
	return s
}

func (s *DescribeTagKeysResponseBody) SetTotalCount(v int32) *DescribeTagKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTagKeysResponseBody) Validate() error {
	if s.TagKeys != nil {
		for _, item := range s.TagKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTagKeysResponseBodyTagKeys struct {
	// The number of Anti-DDoS Proxy (Chinese Mainland) instances to which the tag key is added.
	//
	// example:
	//
	// 2
	TagCount *int32 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// The tag key.
	//
	// example:
	//
	// aa1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s DescribeTagKeysResponseBodyTagKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeysResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysResponseBodyTagKeys) GetTagCount() *int32 {
	return s.TagCount
}

func (s *DescribeTagKeysResponseBodyTagKeys) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeTagKeysResponseBodyTagKeys) SetTagCount(v int32) *DescribeTagKeysResponseBodyTagKeys {
	s.TagCount = &v
	return s
}

func (s *DescribeTagKeysResponseBodyTagKeys) SetTagKey(v string) *DescribeTagKeysResponseBodyTagKeys {
	s.TagKey = &v
	return s
}

func (s *DescribeTagKeysResponseBodyTagKeys) Validate() error {
	return dara.Validate(s)
}
