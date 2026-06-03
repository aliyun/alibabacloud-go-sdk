// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCertificateResponseBody
	GetCode() *string
	SetData(v []*ListCertificateResponseBodyData) *ListCertificateResponseBody
	GetData() []*ListCertificateResponseBodyData
	SetPageNumber(v int32) *ListCertificateResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCertificateResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCertificateResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCertificateResponseBody
	GetTotalCount() *int32
}

type ListCertificateResponseBody struct {
	Code       *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListCertificateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ListCertificateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCertificateResponseBody) GetData() []*ListCertificateResponseBodyData {
	return s.Data
}

func (s *ListCertificateResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCertificateResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCertificateResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCertificateResponseBody) SetCode(v string) *ListCertificateResponseBody {
	s.Code = &v
	return s
}

func (s *ListCertificateResponseBody) SetData(v []*ListCertificateResponseBodyData) *ListCertificateResponseBody {
	s.Data = v
	return s
}

func (s *ListCertificateResponseBody) SetPageNumber(v int32) *ListCertificateResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCertificateResponseBody) SetPageSize(v int32) *ListCertificateResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCertificateResponseBody) SetRequestId(v string) *ListCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCertificateResponseBody) SetTotalCount(v int32) *ListCertificateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCertificateResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCertificateResponseBodyData struct {
	AuthorizationTime *int64   `json:"AuthorizationTime,omitempty" xml:"AuthorizationTime,omitempty"`
	BlockchainNo      *string  `json:"BlockchainNo,omitempty" xml:"BlockchainNo,omitempty"`
	BusinessTypeList  []*int32 `json:"BusinessTypeList,omitempty" xml:"BusinessTypeList,omitempty" type:"Repeated"`
	CancelTime        *int64   `json:"CancelTime,omitempty" xml:"CancelTime,omitempty"`
	CompanyName       *string  `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	EndDate           *string  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	PhoneNo           *string  `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	SchemeType        *int32   `json:"SchemeType,omitempty" xml:"SchemeType,omitempty"`
	StartDate         *string  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Status            *int32   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCertificateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCertificateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCertificateResponseBodyData) GetAuthorizationTime() *int64 {
	return s.AuthorizationTime
}

func (s *ListCertificateResponseBodyData) GetBlockchainNo() *string {
	return s.BlockchainNo
}

func (s *ListCertificateResponseBodyData) GetBusinessTypeList() []*int32 {
	return s.BusinessTypeList
}

func (s *ListCertificateResponseBodyData) GetCancelTime() *int64 {
	return s.CancelTime
}

func (s *ListCertificateResponseBodyData) GetCompanyName() *string {
	return s.CompanyName
}

func (s *ListCertificateResponseBodyData) GetEndDate() *string {
	return s.EndDate
}

func (s *ListCertificateResponseBodyData) GetPhoneNo() *string {
	return s.PhoneNo
}

func (s *ListCertificateResponseBodyData) GetSchemeType() *int32 {
	return s.SchemeType
}

func (s *ListCertificateResponseBodyData) GetStartDate() *string {
	return s.StartDate
}

func (s *ListCertificateResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListCertificateResponseBodyData) SetAuthorizationTime(v int64) *ListCertificateResponseBodyData {
	s.AuthorizationTime = &v
	return s
}

func (s *ListCertificateResponseBodyData) SetBlockchainNo(v string) *ListCertificateResponseBodyData {
	s.BlockchainNo = &v
	return s
}

func (s *ListCertificateResponseBodyData) SetBusinessTypeList(v []*int32) *ListCertificateResponseBodyData {
	s.BusinessTypeList = v
	return s
}

func (s *ListCertificateResponseBodyData) SetCancelTime(v int64) *ListCertificateResponseBodyData {
	s.CancelTime = &v
	return s
}

func (s *ListCertificateResponseBodyData) SetCompanyName(v string) *ListCertificateResponseBodyData {
	s.CompanyName = &v
	return s
}

func (s *ListCertificateResponseBodyData) SetEndDate(v string) *ListCertificateResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *ListCertificateResponseBodyData) SetPhoneNo(v string) *ListCertificateResponseBodyData {
	s.PhoneNo = &v
	return s
}

func (s *ListCertificateResponseBodyData) SetSchemeType(v int32) *ListCertificateResponseBodyData {
	s.SchemeType = &v
	return s
}

func (s *ListCertificateResponseBodyData) SetStartDate(v string) *ListCertificateResponseBodyData {
	s.StartDate = &v
	return s
}

func (s *ListCertificateResponseBodyData) SetStatus(v int32) *ListCertificateResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListCertificateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
