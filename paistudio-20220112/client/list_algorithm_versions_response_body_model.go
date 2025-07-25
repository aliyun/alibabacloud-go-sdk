// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlgorithmVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmVersions(v []*ListAlgorithmVersionsResponseBodyAlgorithmVersions) *ListAlgorithmVersionsResponseBody
	GetAlgorithmVersions() []*ListAlgorithmVersionsResponseBodyAlgorithmVersions
	SetRequestId(v string) *ListAlgorithmVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListAlgorithmVersionsResponseBody
	GetTotalCount() *int64
}

type ListAlgorithmVersionsResponseBody struct {
	AlgorithmVersions []*ListAlgorithmVersionsResponseBodyAlgorithmVersions `json:"AlgorithmVersions,omitempty" xml:"AlgorithmVersions,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 23
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlgorithmVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlgorithmVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlgorithmVersionsResponseBody) GetAlgorithmVersions() []*ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	return s.AlgorithmVersions
}

func (s *ListAlgorithmVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlgorithmVersionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAlgorithmVersionsResponseBody) SetAlgorithmVersions(v []*ListAlgorithmVersionsResponseBodyAlgorithmVersions) *ListAlgorithmVersionsResponseBody {
	s.AlgorithmVersions = v
	return s
}

func (s *ListAlgorithmVersionsResponseBody) SetRequestId(v string) *ListAlgorithmVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBody) SetTotalCount(v int64) *ListAlgorithmVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAlgorithmVersionsResponseBodyAlgorithmVersions struct {
	// example:
	//
	// algo-sidjc8134hv
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// llm_train
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// example:
	//
	// v0.1.0
	AlgorithmVersion *string `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	// example:
	//
	// 2024-01-19T02:00:26Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2024-01-22T02:00:59Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 123456789
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAlgorithmVersionsResponseBodyAlgorithmVersions) String() string {
	return dara.Prettify(s)
}

func (s ListAlgorithmVersionsResponseBodyAlgorithmVersions) GoString() string {
	return s.String()
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) GetAlgorithmId() *string {
	return s.AlgorithmId
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) GetAlgorithmName() *string {
	return s.AlgorithmName
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) GetAlgorithmProvider() *string {
	return s.AlgorithmProvider
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) GetAlgorithmVersion() *string {
	return s.AlgorithmVersion
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) GetUserId() *string {
	return s.UserId
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) SetAlgorithmId(v string) *ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	s.AlgorithmId = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) SetAlgorithmName(v string) *ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	s.AlgorithmName = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) SetAlgorithmProvider(v string) *ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	s.AlgorithmProvider = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) SetAlgorithmVersion(v string) *ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	s.AlgorithmVersion = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) SetGmtCreateTime(v string) *ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	s.GmtCreateTime = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) SetGmtModifiedTime(v string) *ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) SetTenantId(v string) *ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	s.TenantId = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) SetUserId(v string) *ListAlgorithmVersionsResponseBodyAlgorithmVersions {
	s.UserId = &v
	return s
}

func (s *ListAlgorithmVersionsResponseBodyAlgorithmVersions) Validate() error {
	return dara.Validate(s)
}
