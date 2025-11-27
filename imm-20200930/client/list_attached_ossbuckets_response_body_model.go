// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAttachedOSSBucketsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttachedOSSBuckets(v []*ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) *ListAttachedOSSBucketsResponseBody
	GetAttachedOSSBuckets() []*ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets
	SetNextToken(v string) *ListAttachedOSSBucketsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAttachedOSSBucketsResponseBody
	GetRequestId() *string
}

type ListAttachedOSSBucketsResponseBody struct {
	// List of bound OSS Buckets.
	AttachedOSSBuckets []*ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets `json:"AttachedOSSBuckets,omitempty" xml:"AttachedOSSBuckets,omitempty" type:"Repeated"`
	// Pagination token. When the total number of tasks in the list exceeds the set MaxResults, this token is used for pagination. This parameter has a value only when not all matching task lists are returned.
	//
	// Use this value as NextToken in the next call to return the subsequent task list.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpwZw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 9847E7D0-A9A3-0053-84C6-BA16FF******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAttachedOSSBucketsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAttachedOSSBucketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAttachedOSSBucketsResponseBody) GetAttachedOSSBuckets() []*ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets {
	return s.AttachedOSSBuckets
}

func (s *ListAttachedOSSBucketsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAttachedOSSBucketsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAttachedOSSBucketsResponseBody) SetAttachedOSSBuckets(v []*ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) *ListAttachedOSSBucketsResponseBody {
	s.AttachedOSSBuckets = v
	return s
}

func (s *ListAttachedOSSBucketsResponseBody) SetNextToken(v string) *ListAttachedOSSBucketsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAttachedOSSBucketsResponseBody) SetRequestId(v string) *ListAttachedOSSBucketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAttachedOSSBucketsResponseBody) Validate() error {
	if s.AttachedOSSBuckets != nil {
		for _, item := range s.AttachedOSSBuckets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets struct {
	// Timestamp of the project creation time, formatted as RFC3339Nano.
	//
	// example:
	//
	// 2021-06-29T14:50:13.011643661+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Description
	//
	// example:
	//
	// test bucket
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// OSS Bucket name.
	//
	// example:
	//
	// test-bucket
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	// User ID.
	//
	// example:
	//
	// 1023***********
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Project name.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Timestamp of the project modification time, formatted as RFC3339Nano.
	//
	// example:
	//
	// 2021-06-29T14:50:13.011643661+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) String() string {
	return dara.Prettify(s)
}

func (s ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) GoString() string {
	return s.String()
}

func (s *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) GetDescription() *string {
	return s.Description
}

func (s *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) SetCreateTime(v string) *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets {
	s.CreateTime = &v
	return s
}

func (s *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) SetDescription(v string) *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets {
	s.Description = &v
	return s
}

func (s *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) SetOSSBucket(v string) *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets {
	s.OSSBucket = &v
	return s
}

func (s *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) SetOwnerId(v string) *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets {
	s.OwnerId = &v
	return s
}

func (s *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) SetProjectName(v string) *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets {
	s.ProjectName = &v
	return s
}

func (s *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) SetUpdateTime(v string) *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets {
	s.UpdateTime = &v
	return s
}

func (s *ListAttachedOSSBucketsResponseBodyAttachedOSSBuckets) Validate() error {
	return dara.Validate(s)
}
