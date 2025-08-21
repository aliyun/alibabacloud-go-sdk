// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketInfos(v []*ListBucketsResponseBodyBucketInfos) *ListBucketsResponseBody
	GetBucketInfos() []*ListBucketsResponseBodyBucketInfos
	SetRequestId(v string) *ListBucketsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListBucketsResponseBody
	GetTotalCount() *int64
}

type ListBucketsResponseBody struct {
	// The list of bucket information.
	BucketInfos []*ListBucketsResponseBodyBucketInfos `json:"BucketInfos,omitempty" xml:"BucketInfos,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 435769C7-AA6F-4DC5-B3DB-A3DC0DE7E853
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of buckets that match the conditions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBucketsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBucketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBucketsResponseBody) GetBucketInfos() []*ListBucketsResponseBodyBucketInfos {
	return s.BucketInfos
}

func (s *ListBucketsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBucketsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListBucketsResponseBody) SetBucketInfos(v []*ListBucketsResponseBodyBucketInfos) *ListBucketsResponseBody {
	s.BucketInfos = v
	return s
}

func (s *ListBucketsResponseBody) SetRequestId(v string) *ListBucketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBucketsResponseBody) SetTotalCount(v int64) *ListBucketsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBucketsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListBucketsResponseBodyBucketInfos struct {
	// The access control list (ACL) of the bucket.
	//
	// 	- **public-read-write**
	//
	// 	- **public-read**
	//
	// 	- **private*	- (default)
	//
	// example:
	//
	// private
	BucketAcl *string `json:"BucketAcl,omitempty" xml:"BucketAcl,omitempty"`
	// The name of the bucket.
	//
	// example:
	//
	// test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The remarks.
	//
	// example:
	//
	// numb
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the bucket was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-10-12T05:45:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the region where the node is located.
	//
	// example:
	//
	// cn-dalian-unicom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The type of the single-node storage. Set the value to sink.
	//
	// example:
	//
	// sink
	LogicalBucketType *string `json:"LogicalBucketType,omitempty" xml:"LogicalBucketType,omitempty"`
	// The time when the bucket was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-10-12T06:45:00Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s ListBucketsResponseBodyBucketInfos) String() string {
	return dara.Prettify(s)
}

func (s ListBucketsResponseBodyBucketInfos) GoString() string {
	return s.String()
}

func (s *ListBucketsResponseBodyBucketInfos) GetBucketAcl() *string {
	return s.BucketAcl
}

func (s *ListBucketsResponseBodyBucketInfos) GetBucketName() *string {
	return s.BucketName
}

func (s *ListBucketsResponseBodyBucketInfos) GetComment() *string {
	return s.Comment
}

func (s *ListBucketsResponseBodyBucketInfos) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListBucketsResponseBodyBucketInfos) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *ListBucketsResponseBodyBucketInfos) GetLogicalBucketType() *string {
	return s.LogicalBucketType
}

func (s *ListBucketsResponseBodyBucketInfos) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListBucketsResponseBodyBucketInfos) SetBucketAcl(v string) *ListBucketsResponseBodyBucketInfos {
	s.BucketAcl = &v
	return s
}

func (s *ListBucketsResponseBodyBucketInfos) SetBucketName(v string) *ListBucketsResponseBodyBucketInfos {
	s.BucketName = &v
	return s
}

func (s *ListBucketsResponseBodyBucketInfos) SetComment(v string) *ListBucketsResponseBodyBucketInfos {
	s.Comment = &v
	return s
}

func (s *ListBucketsResponseBodyBucketInfos) SetCreateTime(v string) *ListBucketsResponseBodyBucketInfos {
	s.CreateTime = &v
	return s
}

func (s *ListBucketsResponseBodyBucketInfos) SetEnsRegionId(v string) *ListBucketsResponseBodyBucketInfos {
	s.EnsRegionId = &v
	return s
}

func (s *ListBucketsResponseBodyBucketInfos) SetLogicalBucketType(v string) *ListBucketsResponseBodyBucketInfos {
	s.LogicalBucketType = &v
	return s
}

func (s *ListBucketsResponseBodyBucketInfos) SetModifyTime(v string) *ListBucketsResponseBodyBucketInfos {
	s.ModifyTime = &v
	return s
}

func (s *ListBucketsResponseBodyBucketInfos) Validate() error {
	return dara.Validate(s)
}
