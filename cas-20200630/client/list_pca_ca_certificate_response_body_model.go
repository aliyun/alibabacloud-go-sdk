// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPcaCaCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ListPcaCaCertificateResponseBodyList) *ListPcaCaCertificateResponseBody
	GetList() []*ListPcaCaCertificateResponseBodyList
	SetMaxResults(v int32) *ListPcaCaCertificateResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListPcaCaCertificateResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPcaCaCertificateResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListPcaCaCertificateResponseBody
	GetTotalCount() *int64
}

type ListPcaCaCertificateResponseBody struct {
	List []*ListPcaCaCertificateResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 12345678-1234-1234-1234-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPcaCaCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPcaCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ListPcaCaCertificateResponseBody) GetList() []*ListPcaCaCertificateResponseBodyList {
	return s.List
}

func (s *ListPcaCaCertificateResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPcaCaCertificateResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPcaCaCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPcaCaCertificateResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListPcaCaCertificateResponseBody) SetList(v []*ListPcaCaCertificateResponseBodyList) *ListPcaCaCertificateResponseBody {
	s.List = v
	return s
}

func (s *ListPcaCaCertificateResponseBody) SetMaxResults(v int32) *ListPcaCaCertificateResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPcaCaCertificateResponseBody) SetNextToken(v string) *ListPcaCaCertificateResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPcaCaCertificateResponseBody) SetRequestId(v string) *ListPcaCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPcaCaCertificateResponseBody) SetTotalCount(v int64) *ListPcaCaCertificateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPcaCaCertificateResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPcaCaCertificateResponseBodyList struct {
	// example:
	//
	// 1ef78be5-******-b5ef0f0eba3d
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// example:
	//
	// Example Co., Ltd.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// example:
	//
	// 1ef78be5-******-b5ef0f0eba3d
	IssuerIdentifier *string `json:"IssuerIdentifier,omitempty" xml:"IssuerIdentifier,omitempty"`
	// example:
	//
	// cas_deposit-cn-******
	PrivateCaInstanceId *string `json:"PrivateCaInstanceId,omitempty" xml:"PrivateCaInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	PrivateCaRegionId *string `json:"PrivateCaRegionId,omitempty" xml:"PrivateCaRegionId,omitempty"`
	// example:
	//
	// ISSUE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 166********
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListPcaCaCertificateResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListPcaCaCertificateResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListPcaCaCertificateResponseBodyList) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *ListPcaCaCertificateResponseBodyList) GetCommonName() *string {
	return s.CommonName
}

func (s *ListPcaCaCertificateResponseBodyList) GetIssuerIdentifier() *string {
	return s.IssuerIdentifier
}

func (s *ListPcaCaCertificateResponseBodyList) GetPrivateCaInstanceId() *string {
	return s.PrivateCaInstanceId
}

func (s *ListPcaCaCertificateResponseBodyList) GetPrivateCaRegionId() *string {
	return s.PrivateCaRegionId
}

func (s *ListPcaCaCertificateResponseBodyList) GetStatus() *string {
	return s.Status
}

func (s *ListPcaCaCertificateResponseBodyList) GetUserId() *string {
	return s.UserId
}

func (s *ListPcaCaCertificateResponseBodyList) SetCertIdentifier(v string) *ListPcaCaCertificateResponseBodyList {
	s.CertIdentifier = &v
	return s
}

func (s *ListPcaCaCertificateResponseBodyList) SetCommonName(v string) *ListPcaCaCertificateResponseBodyList {
	s.CommonName = &v
	return s
}

func (s *ListPcaCaCertificateResponseBodyList) SetIssuerIdentifier(v string) *ListPcaCaCertificateResponseBodyList {
	s.IssuerIdentifier = &v
	return s
}

func (s *ListPcaCaCertificateResponseBodyList) SetPrivateCaInstanceId(v string) *ListPcaCaCertificateResponseBodyList {
	s.PrivateCaInstanceId = &v
	return s
}

func (s *ListPcaCaCertificateResponseBodyList) SetPrivateCaRegionId(v string) *ListPcaCaCertificateResponseBodyList {
	s.PrivateCaRegionId = &v
	return s
}

func (s *ListPcaCaCertificateResponseBodyList) SetStatus(v string) *ListPcaCaCertificateResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListPcaCaCertificateResponseBodyList) SetUserId(v string) *ListPcaCaCertificateResponseBodyList {
	s.UserId = &v
	return s
}

func (s *ListPcaCaCertificateResponseBodyList) Validate() error {
	return dara.Validate(s)
}
