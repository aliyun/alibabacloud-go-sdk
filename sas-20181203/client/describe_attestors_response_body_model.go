// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttestorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttestors(v []*DescribeAttestorsResponseBodyAttestors) *DescribeAttestorsResponseBody
	GetAttestors() []*DescribeAttestorsResponseBodyAttestors
	SetPageInfo(v *DescribeAttestorsResponseBodyPageInfo) *DescribeAttestorsResponseBody
	GetPageInfo() *DescribeAttestorsResponseBodyPageInfo
	SetRequestId(v string) *DescribeAttestorsResponseBody
	GetRequestId() *string
}

type DescribeAttestorsResponseBody struct {
	// The witnesses.
	Attestors []*DescribeAttestorsResponseBodyAttestors `json:"Attestors,omitempty" xml:"Attestors,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeAttestorsResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7BC55C8F-226E-5AF5-9A2C-2EC43864****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAttestorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttestorsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAttestorsResponseBody) GetAttestors() []*DescribeAttestorsResponseBodyAttestors {
	return s.Attestors
}

func (s *DescribeAttestorsResponseBody) GetPageInfo() *DescribeAttestorsResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeAttestorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAttestorsResponseBody) SetAttestors(v []*DescribeAttestorsResponseBodyAttestors) *DescribeAttestorsResponseBody {
	s.Attestors = v
	return s
}

func (s *DescribeAttestorsResponseBody) SetPageInfo(v *DescribeAttestorsResponseBodyPageInfo) *DescribeAttestorsResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeAttestorsResponseBody) SetRequestId(v string) *DescribeAttestorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAttestorsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAttestorsResponseBodyAttestors struct {
	// The ID of the KMS key.
	//
	// example:
	//
	// 2e81355b-f8e7-4090-8082-a8f8124a****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The region ID of the KMS key.
	//
	// example:
	//
	// cn-hangzhou
	KeyRegionId *string `json:"KeyRegionId,omitempty" xml:"KeyRegionId,omitempty"`
	// The version ID of the Key Management Service (KMS) key.
	//
	// example:
	//
	// 8d7c9c91-57ce-4cf4-a959-1e700e13****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The name of the witness.
	//
	// example:
	//
	// attestor-123
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description.
	//
	// example:
	//
	// attestor
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s DescribeAttestorsResponseBodyAttestors) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttestorsResponseBodyAttestors) GoString() string {
	return s.String()
}

func (s *DescribeAttestorsResponseBodyAttestors) GetKeyId() *string {
	return s.KeyId
}

func (s *DescribeAttestorsResponseBodyAttestors) GetKeyRegionId() *string {
	return s.KeyRegionId
}

func (s *DescribeAttestorsResponseBodyAttestors) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *DescribeAttestorsResponseBodyAttestors) GetName() *string {
	return s.Name
}

func (s *DescribeAttestorsResponseBodyAttestors) GetRemark() *string {
	return s.Remark
}

func (s *DescribeAttestorsResponseBodyAttestors) SetKeyId(v string) *DescribeAttestorsResponseBodyAttestors {
	s.KeyId = &v
	return s
}

func (s *DescribeAttestorsResponseBodyAttestors) SetKeyRegionId(v string) *DescribeAttestorsResponseBodyAttestors {
	s.KeyRegionId = &v
	return s
}

func (s *DescribeAttestorsResponseBodyAttestors) SetKeyVersionId(v string) *DescribeAttestorsResponseBodyAttestors {
	s.KeyVersionId = &v
	return s
}

func (s *DescribeAttestorsResponseBodyAttestors) SetName(v string) *DescribeAttestorsResponseBodyAttestors {
	s.Name = &v
	return s
}

func (s *DescribeAttestorsResponseBodyAttestors) SetRemark(v string) *DescribeAttestorsResponseBodyAttestors {
	s.Remark = &v
	return s
}

func (s *DescribeAttestorsResponseBodyAttestors) Validate() error {
	return dara.Validate(s)
}

type DescribeAttestorsResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 122
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAttestorsResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttestorsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAttestorsResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeAttestorsResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAttestorsResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAttestorsResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAttestorsResponseBodyPageInfo) SetCount(v int32) *DescribeAttestorsResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeAttestorsResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeAttestorsResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAttestorsResponseBodyPageInfo) SetPageSize(v int32) *DescribeAttestorsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeAttestorsResponseBodyPageInfo) SetTotalCount(v int32) *DescribeAttestorsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeAttestorsResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
