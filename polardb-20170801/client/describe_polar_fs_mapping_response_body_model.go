// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultAccessKeyId(v string) *DescribePolarFsMappingResponseBody
	GetDefaultAccessKeyId() *string
	SetPageNumber(v string) *DescribePolarFsMappingResponseBody
	GetPageNumber() *string
	SetPageRecordCount(v string) *DescribePolarFsMappingResponseBody
	GetPageRecordCount() *string
	SetPageSize(v string) *DescribePolarFsMappingResponseBody
	GetPageSize() *string
	SetPathMappingItems(v []*DescribePolarFsMappingResponseBodyPathMappingItems) *DescribePolarFsMappingResponseBody
	GetPathMappingItems() []*DescribePolarFsMappingResponseBodyPathMappingItems
	SetRequestId(v string) *DescribePolarFsMappingResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v string) *DescribePolarFsMappingResponseBody
	GetTotalRecordCount() *string
}

type DescribePolarFsMappingResponseBody struct {
	// The default AccessKey ID at the instance level.
	//
	// example:
	//
	// xxx
	DefaultAccessKeyId *string `json:"DefaultAccessKeyId,omitempty" xml:"DefaultAccessKeyId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records on the current page.
	//
	// example:
	//
	// 1
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of path mappings.
	PathMappingItems []*DescribePolarFsMappingResponseBodyPathMappingItems `json:"PathMappingItems,omitempty" xml:"PathMappingItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribePolarFsMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsMappingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarFsMappingResponseBody) GetDefaultAccessKeyId() *string {
	return s.DefaultAccessKeyId
}

func (s *DescribePolarFsMappingResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribePolarFsMappingResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *DescribePolarFsMappingResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribePolarFsMappingResponseBody) GetPathMappingItems() []*DescribePolarFsMappingResponseBodyPathMappingItems {
	return s.PathMappingItems
}

func (s *DescribePolarFsMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarFsMappingResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribePolarFsMappingResponseBody) SetDefaultAccessKeyId(v string) *DescribePolarFsMappingResponseBody {
	s.DefaultAccessKeyId = &v
	return s
}

func (s *DescribePolarFsMappingResponseBody) SetPageNumber(v string) *DescribePolarFsMappingResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarFsMappingResponseBody) SetPageRecordCount(v string) *DescribePolarFsMappingResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribePolarFsMappingResponseBody) SetPageSize(v string) *DescribePolarFsMappingResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePolarFsMappingResponseBody) SetPathMappingItems(v []*DescribePolarFsMappingResponseBodyPathMappingItems) *DescribePolarFsMappingResponseBody {
	s.PathMappingItems = v
	return s
}

func (s *DescribePolarFsMappingResponseBody) SetRequestId(v string) *DescribePolarFsMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarFsMappingResponseBody) SetTotalRecordCount(v string) *DescribePolarFsMappingResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribePolarFsMappingResponseBody) Validate() error {
	if s.PathMappingItems != nil {
		for _, item := range s.PathMappingItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarFsMappingResponseBodyPathMappingItems struct {
	// The storage bucket.
	//
	// example:
	//
	// pfs-xxx.oss-[regionId]-internal.aliyuncs.com
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The AccessKey ID of the storage bucket.
	//
	// example:
	//
	// xxx
	BucketAccessKeyId *string `json:"BucketAccessKeyId,omitempty" xml:"BucketAccessKeyId,omitempty"`
	// The mapping path.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DescribePolarFsMappingResponseBodyPathMappingItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsMappingResponseBodyPathMappingItems) GoString() string {
	return s.String()
}

func (s *DescribePolarFsMappingResponseBodyPathMappingItems) GetBucket() *string {
	return s.Bucket
}

func (s *DescribePolarFsMappingResponseBodyPathMappingItems) GetBucketAccessKeyId() *string {
	return s.BucketAccessKeyId
}

func (s *DescribePolarFsMappingResponseBodyPathMappingItems) GetPath() *string {
	return s.Path
}

func (s *DescribePolarFsMappingResponseBodyPathMappingItems) SetBucket(v string) *DescribePolarFsMappingResponseBodyPathMappingItems {
	s.Bucket = &v
	return s
}

func (s *DescribePolarFsMappingResponseBodyPathMappingItems) SetBucketAccessKeyId(v string) *DescribePolarFsMappingResponseBodyPathMappingItems {
	s.BucketAccessKeyId = &v
	return s
}

func (s *DescribePolarFsMappingResponseBodyPathMappingItems) SetPath(v string) *DescribePolarFsMappingResponseBodyPathMappingItems {
	s.Path = &v
	return s
}

func (s *DescribePolarFsMappingResponseBodyPathMappingItems) Validate() error {
	return dara.Validate(s)
}
