// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsLogstoreInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogStore(v string) *DescribeSlsLogstoreInfoResponseBody
	GetLogStore() *string
	SetProject(v string) *DescribeSlsLogstoreInfoResponseBody
	GetProject() *string
	SetQuota(v int64) *DescribeSlsLogstoreInfoResponseBody
	GetQuota() *int64
	SetRequestId(v string) *DescribeSlsLogstoreInfoResponseBody
	GetRequestId() *string
	SetTtl(v int32) *DescribeSlsLogstoreInfoResponseBody
	GetTtl() *int32
	SetUsed(v int64) *DescribeSlsLogstoreInfoResponseBody
	GetUsed() *int64
}

type DescribeSlsLogstoreInfoResponseBody struct {
	// example:
	//
	// ddoscoo-logstore
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// example:
	//
	// ddoscoo-project-xxxx-cn-hangzhou
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// example:
	//
	// 5497558138880
	Quota *int64 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Ttl       *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s DescribeSlsLogstoreInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsLogstoreInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogstoreInfoResponseBody) GetLogStore() *string {
	return s.LogStore
}

func (s *DescribeSlsLogstoreInfoResponseBody) GetProject() *string {
	return s.Project
}

func (s *DescribeSlsLogstoreInfoResponseBody) GetQuota() *int64 {
	return s.Quota
}

func (s *DescribeSlsLogstoreInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlsLogstoreInfoResponseBody) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeSlsLogstoreInfoResponseBody) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetLogStore(v string) *DescribeSlsLogstoreInfoResponseBody {
	s.LogStore = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetProject(v string) *DescribeSlsLogstoreInfoResponseBody {
	s.Project = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetQuota(v int64) *DescribeSlsLogstoreInfoResponseBody {
	s.Quota = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetRequestId(v string) *DescribeSlsLogstoreInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetTtl(v int32) *DescribeSlsLogstoreInfoResponseBody {
	s.Ttl = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetUsed(v int64) *DescribeSlsLogstoreInfoResponseBody {
	s.Used = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
