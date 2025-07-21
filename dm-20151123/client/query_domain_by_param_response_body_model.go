// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainByParamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *QueryDomainByParamResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryDomainByParamResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryDomainByParamResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryDomainByParamResponseBody
	GetTotalCount() *int32
	SetData(v *QueryDomainByParamResponseBodyData) *QueryDomainByParamResponseBody
	GetData() *QueryDomainByParamResponseBodyData
}

type QueryDomainByParamResponseBody struct {
	// Current page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 8C90CCD3-627C-4F87-AD8C-2F03146071EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// List of domains
	Data *QueryDomainByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryDomainByParamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainByParamResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryDomainByParamResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryDomainByParamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDomainByParamResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryDomainByParamResponseBody) GetData() *QueryDomainByParamResponseBodyData {
	return s.Data
}

func (s *QueryDomainByParamResponseBody) SetPageNumber(v int32) *QueryDomainByParamResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryDomainByParamResponseBody) SetPageSize(v int32) *QueryDomainByParamResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryDomainByParamResponseBody) SetRequestId(v string) *QueryDomainByParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainByParamResponseBody) SetTotalCount(v int32) *QueryDomainByParamResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryDomainByParamResponseBody) SetData(v *QueryDomainByParamResponseBodyData) *QueryDomainByParamResponseBody {
	s.Data = v
	return s
}

func (s *QueryDomainByParamResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryDomainByParamResponseBodyData struct {
	Domain []*QueryDomainByParamResponseBodyDataDomain `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s QueryDomainByParamResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDomainByParamResponseBodyData) GetDomain() []*QueryDomainByParamResponseBodyDataDomain {
	return s.Domain
}

func (s *QueryDomainByParamResponseBodyData) SetDomain(v []*QueryDomainByParamResponseBodyDataDomain) *QueryDomainByParamResponseBodyData {
	s.Domain = v
	return s
}

func (s *QueryDomainByParamResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryDomainByParamResponseBodyDataDomain struct {
	// Track verification
	//
	// example:
	//
	// 0
	CnameAuthStatus *string `json:"CnameAuthStatus,omitempty" xml:"CnameAuthStatus,omitempty"`
	// CName verification status, success: 0; failure: 1
	//
	// example:
	//
	// 0
	ConfirmStatus *string `json:"ConfirmStatus,omitempty" xml:"ConfirmStatus,omitempty"`
	// Creation time
	//
	// example:
	//
	// 2019-09-29T13:28Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Domain ID
	//
	// example:
	//
	// 158923
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// Domain name
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Domain record
	//
	// example:
	//
	// 6bd86901b9fe4618a046
	DomainRecord *string `json:"DomainRecord,omitempty" xml:"DomainRecord,omitempty"`
	// Domain status.
	//
	// - 0: Available, verified
	//
	// - 1: Unavailable, verification failed
	//
	// example:
	//
	// 0
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// ICP filing status.
	//
	// - 1 indicates filed
	//
	// - 0 indicates not filed
	//
	// example:
	//
	// 1
	IcpStatus *string `json:"IcpStatus,omitempty" xml:"IcpStatus,omitempty"`
	// MX authentication status, success: 0, failure: 1.
	//
	// example:
	//
	// 0
	MxAuthStatus *string `json:"MxAuthStatus,omitempty" xml:"MxAuthStatus,omitempty"`
	// SPF authentication status, success: 0, failure: 1.
	//
	// example:
	//
	// 0
	SpfAuthStatus *string `json:"SpfAuthStatus,omitempty" xml:"SpfAuthStatus,omitempty"`
	// Creation time in UTC format.
	//
	// example:
	//
	// 1569734892
	UtcCreateTime *int64 `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
}

func (s QueryDomainByParamResponseBodyDataDomain) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByParamResponseBodyDataDomain) GoString() string {
	return s.String()
}

func (s *QueryDomainByParamResponseBodyDataDomain) GetCnameAuthStatus() *string {
	return s.CnameAuthStatus
}

func (s *QueryDomainByParamResponseBodyDataDomain) GetConfirmStatus() *string {
	return s.ConfirmStatus
}

func (s *QueryDomainByParamResponseBodyDataDomain) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryDomainByParamResponseBodyDataDomain) GetDomainId() *string {
	return s.DomainId
}

func (s *QueryDomainByParamResponseBodyDataDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainByParamResponseBodyDataDomain) GetDomainRecord() *string {
	return s.DomainRecord
}

func (s *QueryDomainByParamResponseBodyDataDomain) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *QueryDomainByParamResponseBodyDataDomain) GetIcpStatus() *string {
	return s.IcpStatus
}

func (s *QueryDomainByParamResponseBodyDataDomain) GetMxAuthStatus() *string {
	return s.MxAuthStatus
}

func (s *QueryDomainByParamResponseBodyDataDomain) GetSpfAuthStatus() *string {
	return s.SpfAuthStatus
}

func (s *QueryDomainByParamResponseBodyDataDomain) GetUtcCreateTime() *int64 {
	return s.UtcCreateTime
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetCnameAuthStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.CnameAuthStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetConfirmStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.ConfirmStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetCreateTime(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.CreateTime = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetDomainId(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.DomainId = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetDomainName(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.DomainName = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetDomainRecord(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.DomainRecord = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetDomainStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.DomainStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetIcpStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.IcpStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetMxAuthStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.MxAuthStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetSpfAuthStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.SpfAuthStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetUtcCreateTime(v int64) *QueryDomainByParamResponseBodyDataDomain {
	s.UtcCreateTime = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) Validate() error {
	return dara.Validate(s)
}
