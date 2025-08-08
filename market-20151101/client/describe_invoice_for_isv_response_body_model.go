// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvoiceForIsvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeInvoiceForIsvResponseBody
	GetCode() *string
	SetCount(v string) *DescribeInvoiceForIsvResponseBody
	GetCount() *string
	SetMaxResults(v int32) *DescribeInvoiceForIsvResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInvoiceForIsvResponseBody
	GetNextToken() *string
	SetPageNumber(v string) *DescribeInvoiceForIsvResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeInvoiceForIsvResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeInvoiceForIsvResponseBody
	GetRequestId() *string
	SetResult(v []*DescribeInvoiceForIsvResponseBodyResult) *DescribeInvoiceForIsvResponseBody
	GetResult() []*DescribeInvoiceForIsvResponseBodyResult
	SetSuccess(v bool) *DescribeInvoiceForIsvResponseBody
	GetSuccess() *bool
}

type DescribeInvoiceForIsvResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 100
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 3v3mzZN1QdVsTPNiT0OkD36LC9I+AJHU9z2oXBmJJOyy4nQl7MIUZUYG6fdbYBk+
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 6EF60BEC-0242-43AF-BB20-270359FB54A7
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeInvoiceForIsvResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInvoiceForIsvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvoiceForIsvResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvoiceForIsvResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeInvoiceForIsvResponseBody) GetCount() *string {
	return s.Count
}

func (s *DescribeInvoiceForIsvResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInvoiceForIsvResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInvoiceForIsvResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeInvoiceForIsvResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInvoiceForIsvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInvoiceForIsvResponseBody) GetResult() []*DescribeInvoiceForIsvResponseBodyResult {
	return s.Result
}

func (s *DescribeInvoiceForIsvResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInvoiceForIsvResponseBody) SetCode(v string) *DescribeInvoiceForIsvResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBody) SetCount(v string) *DescribeInvoiceForIsvResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBody) SetMaxResults(v int32) *DescribeInvoiceForIsvResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBody) SetNextToken(v string) *DescribeInvoiceForIsvResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBody) SetPageNumber(v string) *DescribeInvoiceForIsvResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBody) SetPageSize(v string) *DescribeInvoiceForIsvResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBody) SetRequestId(v string) *DescribeInvoiceForIsvResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBody) SetResult(v []*DescribeInvoiceForIsvResponseBodyResult) *DescribeInvoiceForIsvResponseBody {
	s.Result = v
	return s
}

func (s *DescribeInvoiceForIsvResponseBody) SetSuccess(v bool) *DescribeInvoiceForIsvResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInvoiceForIsvResponseBodyResult struct {
	// example:
	//
	// 102277855749****
	AliyunPk      *string                                                `json:"AliyunPk,omitempty" xml:"AliyunPk,omitempty"`
	CreateTimeStr *string                                                `json:"CreateTimeStr,omitempty" xml:"CreateTimeStr,omitempty"`
	EvaluateList  []*DescribeInvoiceForIsvResponseBodyResultEvaluateList `json:"EvaluateList,omitempty" xml:"EvaluateList,omitempty" type:"Repeated"`
	// example:
	//
	// 58050005
	Id          *string                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	InvoiceId   *string                                               `json:"InvoiceId,omitempty" xml:"InvoiceId,omitempty"`
	InvoiceList []*DescribeInvoiceForIsvResponseBodyResultInvoiceList `json:"InvoiceList,omitempty" xml:"InvoiceList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	MaterialType *string `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
	// example:
	//
	// 2025-03-04T09:43:18+08:00
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 2025-01-01 00:00:00
	ModifiedTimeStr *string `json:"ModifiedTimeStr,omitempty" xml:"ModifiedTimeStr,omitempty"`
	// example:
	//
	// 99.99
	Price              *string                                                    `json:"Price,omitempty" xml:"Price,omitempty"`
	ReceiptUserInfoDto *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto `json:"ReceiptUserInfoDto,omitempty" xml:"ReceiptUserInfoDto,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2
	Type           *string                                                `json:"Type,omitempty" xml:"Type,omitempty"`
	UserAddressDto *DescribeInvoiceForIsvResponseBodyResultUserAddressDto `json:"UserAddressDto,omitempty" xml:"UserAddressDto,omitempty" type:"Struct"`
}

func (s DescribeInvoiceForIsvResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvoiceForIsvResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeInvoiceForIsvResponseBodyResult) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *DescribeInvoiceForIsvResponseBodyResult) GetCreateTimeStr() *string {
	return s.CreateTimeStr
}

func (s *DescribeInvoiceForIsvResponseBodyResult) GetEvaluateList() []*DescribeInvoiceForIsvResponseBodyResultEvaluateList {
	return s.EvaluateList
}

func (s *DescribeInvoiceForIsvResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *DescribeInvoiceForIsvResponseBodyResult) GetInvoiceId() *string {
	return s.InvoiceId
}

func (s *DescribeInvoiceForIsvResponseBodyResult) GetInvoiceList() []*DescribeInvoiceForIsvResponseBodyResultInvoiceList {
	return s.InvoiceList
}

func (s *DescribeInvoiceForIsvResponseBodyResult) GetMaterialType() *string {
	return s.MaterialType
}

func (s *DescribeInvoiceForIsvResponseBodyResult) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeInvoiceForIsvResponseBodyResult) GetModifiedTimeStr() *string {
	return s.ModifiedTimeStr
}

func (s *DescribeInvoiceForIsvResponseBodyResult) GetPrice() *string {
	return s.Price
}

func (s *DescribeInvoiceForIsvResponseBodyResult) GetReceiptUserInfoDto() *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto {
	return s.ReceiptUserInfoDto
}

func (s *DescribeInvoiceForIsvResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *DescribeInvoiceForIsvResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *DescribeInvoiceForIsvResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *DescribeInvoiceForIsvResponseBodyResult) GetUserAddressDto() *DescribeInvoiceForIsvResponseBodyResultUserAddressDto {
	return s.UserAddressDto
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetAliyunPk(v string) *DescribeInvoiceForIsvResponseBodyResult {
	s.AliyunPk = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetCreateTimeStr(v string) *DescribeInvoiceForIsvResponseBodyResult {
	s.CreateTimeStr = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetEvaluateList(v []*DescribeInvoiceForIsvResponseBodyResultEvaluateList) *DescribeInvoiceForIsvResponseBodyResult {
	s.EvaluateList = v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetId(v string) *DescribeInvoiceForIsvResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetInvoiceId(v string) *DescribeInvoiceForIsvResponseBodyResult {
	s.InvoiceId = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetInvoiceList(v []*DescribeInvoiceForIsvResponseBodyResultInvoiceList) *DescribeInvoiceForIsvResponseBodyResult {
	s.InvoiceList = v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetMaterialType(v string) *DescribeInvoiceForIsvResponseBodyResult {
	s.MaterialType = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetModifiedTime(v string) *DescribeInvoiceForIsvResponseBodyResult {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetModifiedTimeStr(v string) *DescribeInvoiceForIsvResponseBodyResult {
	s.ModifiedTimeStr = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetPrice(v string) *DescribeInvoiceForIsvResponseBodyResult {
	s.Price = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetReceiptUserInfoDto(v *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) *DescribeInvoiceForIsvResponseBodyResult {
	s.ReceiptUserInfoDto = v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetStatus(v string) *DescribeInvoiceForIsvResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetTitle(v string) *DescribeInvoiceForIsvResponseBodyResult {
	s.Title = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetType(v string) *DescribeInvoiceForIsvResponseBodyResult {
	s.Type = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetUserAddressDto(v *DescribeInvoiceForIsvResponseBodyResultUserAddressDto) *DescribeInvoiceForIsvResponseBodyResult {
	s.UserAddressDto = v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type DescribeInvoiceForIsvResponseBodyResultEvaluateList struct {
	// example:
	//
	// false
	Agent *bool `json:"Agent,omitempty" xml:"Agent,omitempty"`
	// example:
	//
	// 1
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// 2025-01-01 00:00:00
	BizTimeStr *string `json:"BizTimeStr,omitempty" xml:"BizTimeStr,omitempty"`
	// example:
	//
	// 9540765
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 2024091610072000****
	OutBizId *string `json:"OutBizId,omitempty" xml:"OutBizId,omitempty"`
	// example:
	//
	// 102277855749****
	RealAliyunPk *string `json:"RealAliyunPk,omitempty" xml:"RealAliyunPk,omitempty"`
}

func (s DescribeInvoiceForIsvResponseBodyResultEvaluateList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvoiceForIsvResponseBodyResultEvaluateList) GoString() string {
	return s.String()
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) GetAgent() *bool {
	return s.Agent
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) GetAmount() *string {
	return s.Amount
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) GetBizTimeStr() *string {
	return s.BizTimeStr
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) GetId() *string {
	return s.Id
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) GetOutBizId() *string {
	return s.OutBizId
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) GetRealAliyunPk() *string {
	return s.RealAliyunPk
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) SetAgent(v bool) *DescribeInvoiceForIsvResponseBodyResultEvaluateList {
	s.Agent = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) SetAmount(v string) *DescribeInvoiceForIsvResponseBodyResultEvaluateList {
	s.Amount = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) SetBizTimeStr(v string) *DescribeInvoiceForIsvResponseBodyResultEvaluateList {
	s.BizTimeStr = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) SetId(v string) *DescribeInvoiceForIsvResponseBodyResultEvaluateList {
	s.Id = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) SetOrderType(v string) *DescribeInvoiceForIsvResponseBodyResultEvaluateList {
	s.OrderType = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) SetOutBizId(v string) *DescribeInvoiceForIsvResponseBodyResultEvaluateList {
	s.OutBizId = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) SetRealAliyunPk(v string) *DescribeInvoiceForIsvResponseBodyResultEvaluateList {
	s.RealAliyunPk = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) Validate() error {
	return dara.Validate(s)
}

type DescribeInvoiceForIsvResponseBodyResultInvoiceList struct {
	// example:
	//
	// 50000018
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 99.99
	InvoiceAmount *string `json:"InvoiceAmount,omitempty" xml:"InvoiceAmount,omitempty"`
}

func (s DescribeInvoiceForIsvResponseBodyResultInvoiceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvoiceForIsvResponseBodyResultInvoiceList) GoString() string {
	return s.String()
}

func (s *DescribeInvoiceForIsvResponseBodyResultInvoiceList) GetId() *string {
	return s.Id
}

func (s *DescribeInvoiceForIsvResponseBodyResultInvoiceList) GetInvoiceAmount() *string {
	return s.InvoiceAmount
}

func (s *DescribeInvoiceForIsvResponseBodyResultInvoiceList) SetId(v string) *DescribeInvoiceForIsvResponseBodyResultInvoiceList {
	s.Id = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultInvoiceList) SetInvoiceAmount(v string) *DescribeInvoiceForIsvResponseBodyResultInvoiceList {
	s.InvoiceAmount = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultInvoiceList) Validate() error {
	return dara.Validate(s)
}

type DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto struct {
	TaxNumber *string `json:"TaxNumber,omitempty" xml:"TaxNumber,omitempty"`
}

func (s DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) GoString() string {
	return s.String()
}

func (s *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) GetTaxNumber() *string {
	return s.TaxNumber
}

func (s *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) SetTaxNumber(v string) *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto {
	s.TaxNumber = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) Validate() error {
	return dara.Validate(s)
}

type DescribeInvoiceForIsvResponseBodyResultUserAddressDto struct {
	Addressee       *string `json:"Addressee,omitempty" xml:"Addressee,omitempty"`
	AliyunPk        *string `json:"AliyunPk,omitempty" xml:"AliyunPk,omitempty"`
	BizType         *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	DeliveryAddress *string `json:"DeliveryAddress,omitempty" xml:"DeliveryAddress,omitempty"`
	Emails          *string `json:"Emails,omitempty" xml:"Emails,omitempty"`
	Phone           *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	PostalCode      *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
}

func (s DescribeInvoiceForIsvResponseBodyResultUserAddressDto) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvoiceForIsvResponseBodyResultUserAddressDto) GoString() string {
	return s.String()
}

func (s *DescribeInvoiceForIsvResponseBodyResultUserAddressDto) GetAddressee() *string {
	return s.Addressee
}

func (s *DescribeInvoiceForIsvResponseBodyResultUserAddressDto) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *DescribeInvoiceForIsvResponseBodyResultUserAddressDto) GetBizType() *string {
	return s.BizType
}

func (s *DescribeInvoiceForIsvResponseBodyResultUserAddressDto) GetDeliveryAddress() *string {
	return s.DeliveryAddress
}

func (s *DescribeInvoiceForIsvResponseBodyResultUserAddressDto) GetEmails() *string {
	return s.Emails
}

func (s *DescribeInvoiceForIsvResponseBodyResultUserAddressDto) GetPhone() *string {
	return s.Phone
}

func (s *DescribeInvoiceForIsvResponseBodyResultUserAddressDto) GetPostalCode() *string {
	return s.PostalCode
}

func (s *DescribeInvoiceForIsvResponseBodyResultUserAddressDto) SetAddressee(v string) *DescribeInvoiceForIsvResponseBodyResultUserAddressDto {
	s.Addressee = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultUserAddressDto) SetAliyunPk(v string) *DescribeInvoiceForIsvResponseBodyResultUserAddressDto {
	s.AliyunPk = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultUserAddressDto) SetBizType(v string) *DescribeInvoiceForIsvResponseBodyResultUserAddressDto {
	s.BizType = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultUserAddressDto) SetDeliveryAddress(v string) *DescribeInvoiceForIsvResponseBodyResultUserAddressDto {
	s.DeliveryAddress = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultUserAddressDto) SetEmails(v string) *DescribeInvoiceForIsvResponseBodyResultUserAddressDto {
	s.Emails = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultUserAddressDto) SetPhone(v string) *DescribeInvoiceForIsvResponseBodyResultUserAddressDto {
	s.Phone = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultUserAddressDto) SetPostalCode(v string) *DescribeInvoiceForIsvResponseBodyResultUserAddressDto {
	s.PostalCode = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultUserAddressDto) Validate() error {
	return dara.Validate(s)
}
