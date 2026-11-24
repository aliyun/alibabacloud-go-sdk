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
	// The product code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total count.
	//
	// example:
	//
	// 100
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The maximum number of entries per page for a paged query. Maximum value: 50. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token returned by this call.
	//
	// example:
	//
	// 3v3mzZN1QdVsTPNiT0OkD36LC9I+AJHU9z2oXBmJJOyy4nQl7MIUZUYG6fdbYBk+
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6EF60BEC-0242-43AF-BB20-270359FB54A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the successful flow node.
	Result []*DescribeInvoiceForIsvResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	//
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
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInvoiceForIsvResponseBodyResult struct {
	// The Alibaba Cloud user ID of the customer.
	//
	// example:
	//
	// 102277855749****
	AliyunPk *string `json:"AliyunPk,omitempty" xml:"AliyunPk,omitempty"`
	// The service provider invoice remarks.
	//
	// example:
	//
	// None
	CheckNotice *string `json:"CheckNotice,omitempty" xml:"CheckNotice,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-01-01 00:00:00
	CreateTimeStr *string `json:"CreateTimeStr,omitempty" xml:"CreateTimeStr,omitempty"`
	// The order, bill, or contract number.
	EvaluateList []*DescribeInvoiceForIsvResponseBodyResultEvaluateList `json:"EvaluateList,omitempty" xml:"EvaluateList,omitempty" type:"Repeated"`
	// The ID of the invoice application. This corresponds to the InvoiceId parameter when you call the ModifyInvoiceForIsv operation.
	//
	// example:
	//
	// 4072040****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the invoice object.
	//
	// example:
	//
	// 4072040****
	InvoiceId *string `json:"InvoiceId,omitempty" xml:"InvoiceId,omitempty"`
	// **[Deprecated]*	- The invoice sub-collection.
	InvoiceList []*DescribeInvoiceForIsvResponseBodyResultInvoiceList `json:"InvoiceList,omitempty" xml:"InvoiceList,omitempty" type:"Repeated"`
	// The relationship between the invoice title and the customer real-name information. Valid values: `REAL_NAME_MATCH` (The real-name information is consistent), `PERSONAL_APPLY_ENTERPRISE_TITLE` (An individual applies for an enterprise title), `ENTERPRISE_MULTI_TITLE` (Multiple titles for an enterprise), `MASTER_ACCOUNT_OPERATION` (Operated by the main account), and `OTHER` (Other). If not provided, the value is null.
	//
	// example:
	//
	// REAL_NAME_MATCH
	InvoiceTitleRelation *string `json:"InvoiceTitleRelation,omitempty" xml:"InvoiceTitleRelation,omitempty"`
	// The invoice issuance type. Valid values:
	//
	// - 0: Paper.
	//
	// - 1: Electronic.
	//
	// example:
	//
	// 1
	MaterialType *string `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
	// The time when the order status was modified. The time is displayed in GMT.
	//
	// example:
	//
	// 2025-03-04T09:43:18+08:00
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The time when the invoice was modified.
	//
	// example:
	//
	// 2025-01-01 00:00:00
	ModifiedTimeStr *string `json:"ModifiedTimeStr,omitempty" xml:"ModifiedTimeStr,omitempty"`
	// The total invoice amount. Unit: CNY.
	//
	// example:
	//
	// 99.99
	Price *string `json:"Price,omitempty" xml:"Price,omitempty"`
	// The user information of the invoice applicant.
	ReceiptUserInfoDto *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto `json:"ReceiptUserInfoDto,omitempty" xml:"ReceiptUserInfoDto,omitempty" type:"Struct"`
	// The invoice status. Valid values: 0: Pending. 1: Invoiced.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The invoice title.
	//
	// example:
	//
	// Alibaba Cloud Computing Ltd
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The invoice type. Valid values:
	//
	// - 2: Digital electronic general invoice.
	//
	// - 3: Digital electronic special invoice.
	//
	// example:
	//
	// 2
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The address information.
	UserAddressDto *DescribeInvoiceForIsvResponseBodyResultUserAddressDto `json:"UserAddressDto,omitempty" xml:"UserAddressDto,omitempty" type:"Struct"`
	// The user invoice remarks.
	//
	// example:
	//
	// None
	UserNotice *string `json:"UserNotice,omitempty" xml:"UserNotice,omitempty"`
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

func (s *DescribeInvoiceForIsvResponseBodyResult) GetCheckNotice() *string {
	return s.CheckNotice
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

func (s *DescribeInvoiceForIsvResponseBodyResult) GetInvoiceTitleRelation() *string {
	return s.InvoiceTitleRelation
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

func (s *DescribeInvoiceForIsvResponseBodyResult) GetUserNotice() *string {
	return s.UserNotice
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetAliyunPk(v string) *DescribeInvoiceForIsvResponseBodyResult {
	s.AliyunPk = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) SetCheckNotice(v string) *DescribeInvoiceForIsvResponseBodyResult {
	s.CheckNotice = &v
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

func (s *DescribeInvoiceForIsvResponseBodyResult) SetInvoiceTitleRelation(v string) *DescribeInvoiceForIsvResponseBodyResult {
	s.InvoiceTitleRelation = &v
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

func (s *DescribeInvoiceForIsvResponseBodyResult) SetUserNotice(v string) *DescribeInvoiceForIsvResponseBodyResult {
	s.UserNotice = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResult) Validate() error {
	if s.EvaluateList != nil {
		for _, item := range s.EvaluateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InvoiceList != nil {
		for _, item := range s.InvoiceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ReceiptUserInfoDto != nil {
		if err := s.ReceiptUserInfoDto.Validate(); err != nil {
			return err
		}
	}
	if s.UserAddressDto != nil {
		if err := s.UserAddressDto.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInvoiceForIsvResponseBodyResultEvaluateList struct {
	// Specifies whether proxy invoicing is used.
	//
	// example:
	//
	// false
	Agent *bool `json:"Agent,omitempty" xml:"Agent,omitempty"`
	// The amount. Unit: CNY.
	//
	// example:
	//
	// 1
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The time when the transaction occurred.
	//
	// example:
	//
	// 2025-01-01 00:00:00
	BizTimeStr *string `json:"BizTimeStr,omitempty" xml:"BizTimeStr,omitempty"`
	// The ID of the invoicing object.
	//
	// example:
	//
	// 9540765
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of the order or bill. Valid values:
	//
	// - 0: Order.
	//
	// - 1: Bill.
	//
	// example:
	//
	// 0
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 2024091610072000****
	OutBizId *string `json:"OutBizId,omitempty" xml:"OutBizId,omitempty"`
	// The commodity code.
	//
	// example:
	//
	// cmapi00067***
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// **[Deprecated]*	- The product name.
	//
	// example:
	//
	// Image recognition OCR***
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The account ID.
	//
	// example:
	//
	// example@alibaba-inc.com
	RealAliyunId *string `json:"RealAliyunId,omitempty" xml:"RealAliyunId,omitempty"`
	// The account to which the order belongs.
	//
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

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) GetRealAliyunId() *string {
	return s.RealAliyunId
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

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) SetProductCode(v string) *DescribeInvoiceForIsvResponseBodyResultEvaluateList {
	s.ProductCode = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) SetProductName(v string) *DescribeInvoiceForIsvResponseBodyResultEvaluateList {
	s.ProductName = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultEvaluateList) SetRealAliyunId(v string) *DescribeInvoiceForIsvResponseBodyResultEvaluateList {
	s.RealAliyunId = &v
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
	// The ID of the invoice object.
	//
	// example:
	//
	// 4072040****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The applied invoicing amount. Unit: CNY.
	//
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
	// The bank account.
	//
	// example:
	//
	// 000000
	BankAccount *string `json:"BankAccount,omitempty" xml:"BankAccount,omitempty"`
	// The name of the bank.
	//
	// example:
	//
	// 111111
	BankName *string `json:"BankName,omitempty" xml:"BankName,omitempty"`
	// The registered address.
	//
	// example:
	//
	// Hangzhou ****
	RegAddress *string `json:"RegAddress,omitempty" xml:"RegAddress,omitempty"`
	// The registered landline number.
	//
	// example:
	//
	// 0571-****
	RegPhone *string `json:"RegPhone,omitempty" xml:"RegPhone,omitempty"`
	// The unified social credit code of the enterprise.
	//
	// example:
	//
	// 91420100MA49L5****
	TaxNumber *string `json:"TaxNumber,omitempty" xml:"TaxNumber,omitempty"`
}

func (s DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) GoString() string {
	return s.String()
}

func (s *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) GetBankAccount() *string {
	return s.BankAccount
}

func (s *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) GetBankName() *string {
	return s.BankName
}

func (s *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) GetRegAddress() *string {
	return s.RegAddress
}

func (s *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) GetRegPhone() *string {
	return s.RegPhone
}

func (s *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) GetTaxNumber() *string {
	return s.TaxNumber
}

func (s *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) SetBankAccount(v string) *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto {
	s.BankAccount = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) SetBankName(v string) *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto {
	s.BankName = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) SetRegAddress(v string) *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto {
	s.RegAddress = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) SetRegPhone(v string) *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto {
	s.RegPhone = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) SetTaxNumber(v string) *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto {
	s.TaxNumber = &v
	return s
}

func (s *DescribeInvoiceForIsvResponseBodyResultReceiptUserInfoDto) Validate() error {
	return dara.Validate(s)
}

type DescribeInvoiceForIsvResponseBodyResultUserAddressDto struct {
	// The recipient.
	//
	// example:
	//
	// Zhang San
	Addressee *string `json:"Addressee,omitempty" xml:"Addressee,omitempty"`
	// The Alibaba Cloud user ID of the customer.
	//
	// example:
	//
	// 174452687724****
	AliyunPk *string `json:"AliyunPk,omitempty" xml:"AliyunPk,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// ******
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The delivery address.
	//
	// example:
	//
	// Hangzhou ****
	DeliveryAddress *string `json:"DeliveryAddress,omitempty" xml:"DeliveryAddress,omitempty"`
	// The email address.
	//
	// example:
	//
	// demo@example.com
	Emails *string `json:"Emails,omitempty" xml:"Emails,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 130********
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The postal code.
	//
	// example:
	//
	// ******
	PostalCode *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
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
