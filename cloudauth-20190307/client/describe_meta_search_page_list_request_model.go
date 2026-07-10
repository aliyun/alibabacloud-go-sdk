// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetaSearchPageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApi(v string) *DescribeMetaSearchPageListRequest
	GetApi() *string
	SetBankCard(v string) *DescribeMetaSearchPageListRequest
	GetBankCard() *string
	SetBizCode(v string) *DescribeMetaSearchPageListRequest
	GetBizCode() *string
	SetCurrentPage(v int32) *DescribeMetaSearchPageListRequest
	GetCurrentPage() *int32
	SetEndDate(v int64) *DescribeMetaSearchPageListRequest
	GetEndDate() *int64
	SetIdentifyNum(v string) *DescribeMetaSearchPageListRequest
	GetIdentifyNum() *string
	SetIspName(v string) *DescribeMetaSearchPageListRequest
	GetIspName() *string
	SetMobile(v string) *DescribeMetaSearchPageListRequest
	GetMobile() *string
	SetPageSize(v int32) *DescribeMetaSearchPageListRequest
	GetPageSize() *int32
	SetReqId(v string) *DescribeMetaSearchPageListRequest
	GetReqId() *string
	SetStartDate(v int64) *DescribeMetaSearchPageListRequest
	GetStartDate() *int64
	SetSubCode(v string) *DescribeMetaSearchPageListRequest
	GetSubCode() *string
	SetUserName(v string) *DescribeMetaSearchPageListRequest
	GetUserName() *string
	SetVehicleNum(v string) *DescribeMetaSearchPageListRequest
	GetVehicleNum() *string
}

type DescribeMetaSearchPageListRequest struct {
	// The product API. Valid values:
	//
	// - **ID_CARD_2_META**: ID card two-element verification
	//
	// - **ID_PERIOD**: ID card validity period verification
	//
	// - **MOBILE_ONLINE_LENGTH**: mobile number online duration
	//
	// - **MOBILE_ONLINE_STATUS**: mobile number online status
	//
	// - **MOBILE_3_META_SIMPLE**: mobile number three-element verification (simple edition)
	//
	// - **MOBILE_3_META**: mobile number three-element verification (detailed edition)
	//
	// - **MOBILE_2_META**: mobile number two-element verification
	//
	// - **BANK_CARD_N_META**: bank card verification (detailed edition)
	//
	// - **MOBILE_DETECT**: phone number detection
	//
	// - **VEHICLE_N_META**: vehicle element verification (enhanced edition)
	//
	// - **VEHICLE_PENTA_INFO**: vehicle five-element information recognition
	//
	// - **VEHICLE_LICENSE_INFO**: vehicle information recognition
	//
	// - **VEHICLE_INSURE_DATE**: vehicle insurance date query
	//
	// - **VEHICLE_CHECK**: vehicle element verification.
	//
	// This parameter is required.
	//
	// example:
	//
	// ID_CARD_2_META
	Api *string `json:"Api,omitempty" xml:"Api,omitempty"`
	// The bank card number.
	//
	// example:
	//
	// 610*************1181
	BankCard *string `json:"BankCard,omitempty" xml:"BankCard,omitempty"`
	// The verification status. Valid values:
	//
	// - **1**: Verification passed.
	//
	// - **2**: Verification failed.
	//
	// - **3**: No record found.
	//
	// example:
	//
	// 2
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time of the query. The value is a UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1739926800000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The ID card number.
	//
	// example:
	//
	// 522132197411184XXX
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// The name of the telecommunications service provider. Valid values:
	//
	// - **CMCC**: China Mobile
	//
	// - **CUCC**: China Unicom
	//
	// - **CTCC**: China Telecom.
	//
	// example:
	//
	// CTCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// The mobile phone number.
	//
	// example:
	//
	// 19127612221
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B012DB99-6C10-5740-81E0-B3A8C1C1B9C1
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1760198400000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The result code. For more information, see [official documentation](https://www.alibabacloud.com/help/en/id-verification/information-verification/).
	//
	// example:
	//
	// 205
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// The name.
	//
	// example:
	//
	// 张三
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The license plate number.
	//
	// example:
	//
	// 陕A9****
	VehicleNum *string `json:"VehicleNum,omitempty" xml:"VehicleNum,omitempty"`
}

func (s DescribeMetaSearchPageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaSearchPageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetaSearchPageListRequest) GetApi() *string {
	return s.Api
}

func (s *DescribeMetaSearchPageListRequest) GetBankCard() *string {
	return s.BankCard
}

func (s *DescribeMetaSearchPageListRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *DescribeMetaSearchPageListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeMetaSearchPageListRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeMetaSearchPageListRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *DescribeMetaSearchPageListRequest) GetIspName() *string {
	return s.IspName
}

func (s *DescribeMetaSearchPageListRequest) GetMobile() *string {
	return s.Mobile
}

func (s *DescribeMetaSearchPageListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMetaSearchPageListRequest) GetReqId() *string {
	return s.ReqId
}

func (s *DescribeMetaSearchPageListRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeMetaSearchPageListRequest) GetSubCode() *string {
	return s.SubCode
}

func (s *DescribeMetaSearchPageListRequest) GetUserName() *string {
	return s.UserName
}

func (s *DescribeMetaSearchPageListRequest) GetVehicleNum() *string {
	return s.VehicleNum
}

func (s *DescribeMetaSearchPageListRequest) SetApi(v string) *DescribeMetaSearchPageListRequest {
	s.Api = &v
	return s
}

func (s *DescribeMetaSearchPageListRequest) SetBankCard(v string) *DescribeMetaSearchPageListRequest {
	s.BankCard = &v
	return s
}

func (s *DescribeMetaSearchPageListRequest) SetBizCode(v string) *DescribeMetaSearchPageListRequest {
	s.BizCode = &v
	return s
}

func (s *DescribeMetaSearchPageListRequest) SetCurrentPage(v int32) *DescribeMetaSearchPageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeMetaSearchPageListRequest) SetEndDate(v int64) *DescribeMetaSearchPageListRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeMetaSearchPageListRequest) SetIdentifyNum(v string) *DescribeMetaSearchPageListRequest {
	s.IdentifyNum = &v
	return s
}

func (s *DescribeMetaSearchPageListRequest) SetIspName(v string) *DescribeMetaSearchPageListRequest {
	s.IspName = &v
	return s
}

func (s *DescribeMetaSearchPageListRequest) SetMobile(v string) *DescribeMetaSearchPageListRequest {
	s.Mobile = &v
	return s
}

func (s *DescribeMetaSearchPageListRequest) SetPageSize(v int32) *DescribeMetaSearchPageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMetaSearchPageListRequest) SetReqId(v string) *DescribeMetaSearchPageListRequest {
	s.ReqId = &v
	return s
}

func (s *DescribeMetaSearchPageListRequest) SetStartDate(v int64) *DescribeMetaSearchPageListRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeMetaSearchPageListRequest) SetSubCode(v string) *DescribeMetaSearchPageListRequest {
	s.SubCode = &v
	return s
}

func (s *DescribeMetaSearchPageListRequest) SetUserName(v string) *DescribeMetaSearchPageListRequest {
	s.UserName = &v
	return s
}

func (s *DescribeMetaSearchPageListRequest) SetVehicleNum(v string) *DescribeMetaSearchPageListRequest {
	s.VehicleNum = &v
	return s
}

func (s *DescribeMetaSearchPageListRequest) Validate() error {
	return dara.Validate(s)
}
