// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsLogStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogStoreName(v string) *DescribeSlsLogStoreResponseBody
	GetLogStoreName() *string
	SetProjectName(v string) *DescribeSlsLogStoreResponseBody
	GetProjectName() *string
	SetQuota(v int64) *DescribeSlsLogStoreResponseBody
	GetQuota() *int64
	SetRequestId(v string) *DescribeSlsLogStoreResponseBody
	GetRequestId() *string
	SetTtl(v int32) *DescribeSlsLogStoreResponseBody
	GetTtl() *int32
	SetUsed(v int64) *DescribeSlsLogStoreResponseBody
	GetUsed() *int64
}

type DescribeSlsLogStoreResponseBody struct {
	// The name of the Logstore.
	//
	// example:
	//
	// wafng-logstore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// wafng-project-14316572********-cn-hangzhou
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The capacity of the Logstore. Unit: bytes.
	//
	// example:
	//
	// 3298534883328
	Quota *int64 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CAC0A24B-486A-5E12-9894-BE860E5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The storage duration of the Logstore. Unit: days.
	//
	// example:
	//
	// 180
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The used capacity of the Logstore. Unit: bytes.
	//
	// example:
	//
	// 35471136
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s DescribeSlsLogStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsLogStoreResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogStoreResponseBody) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeSlsLogStoreResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeSlsLogStoreResponseBody) GetQuota() *int64 {
	return s.Quota
}

func (s *DescribeSlsLogStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlsLogStoreResponseBody) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeSlsLogStoreResponseBody) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeSlsLogStoreResponseBody) SetLogStoreName(v string) *DescribeSlsLogStoreResponseBody {
	s.LogStoreName = &v
	return s
}

func (s *DescribeSlsLogStoreResponseBody) SetProjectName(v string) *DescribeSlsLogStoreResponseBody {
	s.ProjectName = &v
	return s
}

func (s *DescribeSlsLogStoreResponseBody) SetQuota(v int64) *DescribeSlsLogStoreResponseBody {
	s.Quota = &v
	return s
}

func (s *DescribeSlsLogStoreResponseBody) SetRequestId(v string) *DescribeSlsLogStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsLogStoreResponseBody) SetTtl(v int32) *DescribeSlsLogStoreResponseBody {
	s.Ttl = &v
	return s
}

func (s *DescribeSlsLogStoreResponseBody) SetUsed(v int64) *DescribeSlsLogStoreResponseBody {
	s.Used = &v
	return s
}

func (s *DescribeSlsLogStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
