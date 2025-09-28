// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGateVerifyBillingPublicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryGateVerifyBillingPublicResponseBody
	GetCode() *string
	SetData(v *QueryGateVerifyBillingPublicResponseBodyData) *QueryGateVerifyBillingPublicResponseBody
	GetData() *QueryGateVerifyBillingPublicResponseBodyData
	SetMessage(v string) *QueryGateVerifyBillingPublicResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryGateVerifyBillingPublicResponseBody
	GetRequestId() *string
}

type QueryGateVerifyBillingPublicResponseBody struct {
	// The response code. Valid values:
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The billing information about each verification service.
	Data *QueryGateVerifyBillingPublicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryGateVerifyBillingPublicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyBillingPublicResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingPublicResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryGateVerifyBillingPublicResponseBody) GetData() *QueryGateVerifyBillingPublicResponseBodyData {
	return s.Data
}

func (s *QueryGateVerifyBillingPublicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryGateVerifyBillingPublicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryGateVerifyBillingPublicResponseBody) SetCode(v string) *QueryGateVerifyBillingPublicResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBody) SetData(v *QueryGateVerifyBillingPublicResponseBodyData) *QueryGateVerifyBillingPublicResponseBody {
	s.Data = v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBody) SetMessage(v string) *QueryGateVerifyBillingPublicResponseBody {
	s.Message = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBody) SetRequestId(v string) *QueryGateVerifyBillingPublicResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryGateVerifyBillingPublicResponseBodyData struct {
	// The fees generated for all verification services. Unitrogen: CNY.
	//
	// example:
	//
	// 1234
	AmountSum *string `json:"AmountSum,omitempty" xml:"AmountSum,omitempty"`
	// The details of fees.
	SceneBillingList []*QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList `json:"SceneBillingList,omitempty" xml:"SceneBillingList,omitempty" type:"Repeated"`
}

func (s QueryGateVerifyBillingPublicResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyBillingPublicResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingPublicResponseBodyData) GetAmountSum() *string {
	return s.AmountSum
}

func (s *QueryGateVerifyBillingPublicResponseBodyData) GetSceneBillingList() []*QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	return s.SceneBillingList
}

func (s *QueryGateVerifyBillingPublicResponseBodyData) SetAmountSum(v string) *QueryGateVerifyBillingPublicResponseBodyData {
	s.AmountSum = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyData) SetSceneBillingList(v []*QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) *QueryGateVerifyBillingPublicResponseBodyData {
	s.SceneBillingList = v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList struct {
	// The billable items.
	//
	// example:
	//
	// 74
	Add *string `json:"Add,omitempty" xml:"Add,omitempty"`
	// The fees generated for the verification service. Unitrogen: CNY.
	//
	// example:
	//
	// 1.48
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The application name.
	//
	// example:
	//
	// Aliyun
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The verification method.
	//
	// example:
	//
	// Verification of local phone number
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// The service code.
	//
	// example:
	//
	// FC100000038194004
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The service name.
	//
	// example:
	//
	// Alibaba Cloud Communications
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The unit price. Unit: CNY.
	//
	// example:
	//
	// 0.02
	SinglePrice *string `json:"SinglePrice,omitempty" xml:"SinglePrice,omitempty"`
}

func (s QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) GetAdd() *string {
	return s.Add
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) GetAmount() *string {
	return s.Amount
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) GetAppName() *string {
	return s.AppName
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) GetItemName() *string {
	return s.ItemName
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) GetSceneCode() *string {
	return s.SceneCode
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) GetSceneName() *string {
	return s.SceneName
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) GetSinglePrice() *string {
	return s.SinglePrice
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetAdd(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.Add = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetAmount(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.Amount = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetAppName(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.AppName = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetItemName(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.ItemName = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetSceneCode(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.SceneCode = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetSceneName(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.SceneName = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetSinglePrice(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.SinglePrice = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) Validate() error {
	return dara.Validate(s)
}
