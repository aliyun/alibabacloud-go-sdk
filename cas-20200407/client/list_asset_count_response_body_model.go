// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetCountList(v []*ListAssetCountResponseBodyAssetCountList) *ListAssetCountResponseBody
	GetAssetCountList() []*ListAssetCountResponseBodyAssetCountList
	SetCurrentPage(v int64) *ListAssetCountResponseBody
	GetCurrentPage() *int64
	SetRequestId(v string) *ListAssetCountResponseBody
	GetRequestId() *string
	SetShowSize(v int64) *ListAssetCountResponseBody
	GetShowSize() *int64
	SetTotalCount(v int64) *ListAssetCountResponseBody
	GetTotalCount() *int64
}

type ListAssetCountResponseBody struct {
	AssetCountList []*ListAssetCountResponseBodyAssetCountList `json:"AssetCountList,omitempty" xml:"AssetCountList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAssetCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAssetCountResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssetCountResponseBody) GetAssetCountList() []*ListAssetCountResponseBodyAssetCountList {
	return s.AssetCountList
}

func (s *ListAssetCountResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListAssetCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAssetCountResponseBody) GetShowSize() *int64 {
	return s.ShowSize
}

func (s *ListAssetCountResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAssetCountResponseBody) SetAssetCountList(v []*ListAssetCountResponseBodyAssetCountList) *ListAssetCountResponseBody {
	s.AssetCountList = v
	return s
}

func (s *ListAssetCountResponseBody) SetCurrentPage(v int64) *ListAssetCountResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListAssetCountResponseBody) SetRequestId(v string) *ListAssetCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssetCountResponseBody) SetShowSize(v int64) *ListAssetCountResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListAssetCountResponseBody) SetTotalCount(v int64) *ListAssetCountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAssetCountResponseBody) Validate() error {
	if s.AssetCountList != nil {
		for _, item := range s.AssetCountList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAssetCountResponseBodyAssetCountList struct {
	// example:
	//
	// 99
	AliyunAssetCount *int64 `json:"AliyunAssetCount,omitempty" xml:"AliyunAssetCount,omitempty"`
	// example:
	//
	// 99
	CertificateCount *int32 `json:"CertificateCount,omitempty" xml:"CertificateCount,omitempty"`
	// example:
	//
	// 1767680115423
	CountDate *int64 `json:"CountDate,omitempty" xml:"CountDate,omitempty"`
	// example:
	//
	// 99
	DomainAssetCount *int32 `json:"DomainAssetCount,omitempty" xml:"DomainAssetCount,omitempty"`
	// example:
	//
	// 99
	MultiCloudAssetCount *int64 `json:"MultiCloudAssetCount,omitempty" xml:"MultiCloudAssetCount,omitempty"`
	// example:
	//
	// 99
	Points *int64 `json:"Points,omitempty" xml:"Points,omitempty"`
}

func (s ListAssetCountResponseBodyAssetCountList) String() string {
	return dara.Prettify(s)
}

func (s ListAssetCountResponseBodyAssetCountList) GoString() string {
	return s.String()
}

func (s *ListAssetCountResponseBodyAssetCountList) GetAliyunAssetCount() *int64 {
	return s.AliyunAssetCount
}

func (s *ListAssetCountResponseBodyAssetCountList) GetCertificateCount() *int32 {
	return s.CertificateCount
}

func (s *ListAssetCountResponseBodyAssetCountList) GetCountDate() *int64 {
	return s.CountDate
}

func (s *ListAssetCountResponseBodyAssetCountList) GetDomainAssetCount() *int32 {
	return s.DomainAssetCount
}

func (s *ListAssetCountResponseBodyAssetCountList) GetMultiCloudAssetCount() *int64 {
	return s.MultiCloudAssetCount
}

func (s *ListAssetCountResponseBodyAssetCountList) GetPoints() *int64 {
	return s.Points
}

func (s *ListAssetCountResponseBodyAssetCountList) SetAliyunAssetCount(v int64) *ListAssetCountResponseBodyAssetCountList {
	s.AliyunAssetCount = &v
	return s
}

func (s *ListAssetCountResponseBodyAssetCountList) SetCertificateCount(v int32) *ListAssetCountResponseBodyAssetCountList {
	s.CertificateCount = &v
	return s
}

func (s *ListAssetCountResponseBodyAssetCountList) SetCountDate(v int64) *ListAssetCountResponseBodyAssetCountList {
	s.CountDate = &v
	return s
}

func (s *ListAssetCountResponseBodyAssetCountList) SetDomainAssetCount(v int32) *ListAssetCountResponseBodyAssetCountList {
	s.DomainAssetCount = &v
	return s
}

func (s *ListAssetCountResponseBodyAssetCountList) SetMultiCloudAssetCount(v int64) *ListAssetCountResponseBodyAssetCountList {
	s.MultiCloudAssetCount = &v
	return s
}

func (s *ListAssetCountResponseBodyAssetCountList) SetPoints(v int64) *ListAssetCountResponseBodyAssetCountList {
	s.Points = &v
	return s
}

func (s *ListAssetCountResponseBodyAssetCountList) Validate() error {
	return dara.Validate(s)
}
