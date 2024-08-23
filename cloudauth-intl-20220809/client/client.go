// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CardOcrRequest struct {
	// example:
	//
	// 00000006
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// F
	IdFaceQuality      *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// example:
	//
	// https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	// example:
	//
	// dso9322***dsjsd22
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// T
	Ocr *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	// example:
	//
	// ID_OCR_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// F
	Spoof *string `json:"Spoof,omitempty" xml:"Spoof,omitempty"`
}

func (s CardOcrRequest) String() string {
	return tea.Prettify(s)
}

func (s CardOcrRequest) GoString() string {
	return s.String()
}

func (s *CardOcrRequest) SetDocType(v string) *CardOcrRequest {
	s.DocType = &v
	return s
}

func (s *CardOcrRequest) SetIdFaceQuality(v string) *CardOcrRequest {
	s.IdFaceQuality = &v
	return s
}

func (s *CardOcrRequest) SetIdOcrPictureBase64(v string) *CardOcrRequest {
	s.IdOcrPictureBase64 = &v
	return s
}

func (s *CardOcrRequest) SetIdOcrPictureUrl(v string) *CardOcrRequest {
	s.IdOcrPictureUrl = &v
	return s
}

func (s *CardOcrRequest) SetMerchantBizId(v string) *CardOcrRequest {
	s.MerchantBizId = &v
	return s
}

func (s *CardOcrRequest) SetMerchantUserId(v string) *CardOcrRequest {
	s.MerchantUserId = &v
	return s
}

func (s *CardOcrRequest) SetOcr(v string) *CardOcrRequest {
	s.Ocr = &v
	return s
}

func (s *CardOcrRequest) SetProductCode(v string) *CardOcrRequest {
	s.ProductCode = &v
	return s
}

func (s *CardOcrRequest) SetSpoof(v string) *CardOcrRequest {
	s.Spoof = &v
	return s
}

type CardOcrResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB356FE-BB6A-5DCC-B4C5-E8051787EBA1
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CardOcrResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CardOcrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CardOcrResponseBody) GoString() string {
	return s.String()
}

func (s *CardOcrResponseBody) SetCode(v string) *CardOcrResponseBody {
	s.Code = &v
	return s
}

func (s *CardOcrResponseBody) SetMessage(v string) *CardOcrResponseBody {
	s.Message = &v
	return s
}

func (s *CardOcrResponseBody) SetRequestId(v string) *CardOcrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CardOcrResponseBody) SetResult(v *CardOcrResponseBodyResult) *CardOcrResponseBody {
	s.Result = v
	return s
}

type CardOcrResponseBodyResult struct {
	ExtCardInfo *string `json:"ExtCardInfo,omitempty" xml:"ExtCardInfo,omitempty"`
	ExtIdInfo   *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s CardOcrResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CardOcrResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CardOcrResponseBodyResult) SetExtCardInfo(v string) *CardOcrResponseBodyResult {
	s.ExtCardInfo = &v
	return s
}

func (s *CardOcrResponseBodyResult) SetExtIdInfo(v string) *CardOcrResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *CardOcrResponseBodyResult) SetPassed(v string) *CardOcrResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *CardOcrResponseBodyResult) SetSubCode(v string) *CardOcrResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *CardOcrResponseBodyResult) SetTransactionId(v string) *CardOcrResponseBodyResult {
	s.TransactionId = &v
	return s
}

type CardOcrResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CardOcrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CardOcrResponse) String() string {
	return tea.Prettify(s)
}

func (s CardOcrResponse) GoString() string {
	return s.String()
}

func (s *CardOcrResponse) SetHeaders(v map[string]*string) *CardOcrResponse {
	s.Headers = v
	return s
}

func (s *CardOcrResponse) SetStatusCode(v int32) *CardOcrResponse {
	s.StatusCode = &v
	return s
}

func (s *CardOcrResponse) SetBody(v *CardOcrResponseBody) *CardOcrResponse {
	s.Body = v
	return s
}

type CheckResultRequest struct {
	// example:
	//
	// ***
	ExtraImageControlList *string `json:"ExtraImageControlList,omitempty" xml:"ExtraImageControlList,omitempty"`
	// example:
	//
	// N
	IsReturnImage *string `json:"IsReturnImage,omitempty" xml:"IsReturnImage,omitempty"`
	// example:
	//
	// djs20d***9-dsskc
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// Y
	ReturnFiveCategorySpoofResult *string `json:"ReturnFiveCategorySpoofResult,omitempty" xml:"ReturnFiveCategorySpoofResult,omitempty"`
	// example:
	//
	// 4ab0b***cbde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s CheckResultRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckResultRequest) GoString() string {
	return s.String()
}

func (s *CheckResultRequest) SetExtraImageControlList(v string) *CheckResultRequest {
	s.ExtraImageControlList = &v
	return s
}

func (s *CheckResultRequest) SetIsReturnImage(v string) *CheckResultRequest {
	s.IsReturnImage = &v
	return s
}

func (s *CheckResultRequest) SetMerchantBizId(v string) *CheckResultRequest {
	s.MerchantBizId = &v
	return s
}

func (s *CheckResultRequest) SetReturnFiveCategorySpoofResult(v string) *CheckResultRequest {
	s.ReturnFiveCategorySpoofResult = &v
	return s
}

func (s *CheckResultRequest) SetTransactionId(v string) *CheckResultRequest {
	s.TransactionId = &v
	return s
}

type CheckResultResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CheckResultResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CheckResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *CheckResultResponseBody) SetCode(v string) *CheckResultResponseBody {
	s.Code = &v
	return s
}

func (s *CheckResultResponseBody) SetMessage(v string) *CheckResultResponseBody {
	s.Message = &v
	return s
}

func (s *CheckResultResponseBody) SetRequestId(v string) *CheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckResultResponseBody) SetResult(v *CheckResultResponseBodyResult) *CheckResultResponseBody {
	s.Result = v
	return s
}

type CheckResultResponseBodyResult struct {
	// example:
	//
	// **
	EkycResult *string `json:"EkycResult,omitempty" xml:"EkycResult,omitempty"`
	// example:
	//
	// **
	ExtBasicInfo *string `json:"ExtBasicInfo,omitempty" xml:"ExtBasicInfo,omitempty"`
	// example:
	//
	// **
	ExtFaceInfo *string `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty"`
	// example:
	//
	// **
	ExtIdInfo *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	// example:
	//
	// **
	ExtRiskInfo *string `json:"ExtRiskInfo,omitempty" xml:"ExtRiskInfo,omitempty"`
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// ***
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s CheckResultResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CheckResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CheckResultResponseBodyResult) SetEkycResult(v string) *CheckResultResponseBodyResult {
	s.EkycResult = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtBasicInfo(v string) *CheckResultResponseBodyResult {
	s.ExtBasicInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtFaceInfo(v string) *CheckResultResponseBodyResult {
	s.ExtFaceInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtIdInfo(v string) *CheckResultResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtRiskInfo(v string) *CheckResultResponseBodyResult {
	s.ExtRiskInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetPassed(v string) *CheckResultResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetSubCode(v string) *CheckResultResponseBodyResult {
	s.SubCode = &v
	return s
}

type CheckResultResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckResultResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckResultResponse) GoString() string {
	return s.String()
}

func (s *CheckResultResponse) SetHeaders(v map[string]*string) *CheckResultResponse {
	s.Headers = v
	return s
}

func (s *CheckResultResponse) SetStatusCode(v int32) *CheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckResultResponse) SetBody(v *CheckResultResponseBody) *CheckResultResponse {
	s.Body = v
	return s
}

type DeleteVerifyResultRequest struct {
	// example:
	//
	// Y / N
	DeleteAfterQuery *string `json:"DeleteAfterQuery,omitempty" xml:"DeleteAfterQuery,omitempty"`
	// example:
	//
	// Img / Text / All
	DeleteType *string `json:"DeleteType,omitempty" xml:"DeleteType,omitempty"`
	// example:
	//
	// 4ab0b***cbde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s DeleteVerifyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVerifyResultRequest) GoString() string {
	return s.String()
}

func (s *DeleteVerifyResultRequest) SetDeleteAfterQuery(v string) *DeleteVerifyResultRequest {
	s.DeleteAfterQuery = &v
	return s
}

func (s *DeleteVerifyResultRequest) SetDeleteType(v string) *DeleteVerifyResultRequest {
	s.DeleteType = &v
	return s
}

func (s *DeleteVerifyResultRequest) SetTransactionId(v string) *DeleteVerifyResultRequest {
	s.TransactionId = &v
	return s
}

type DeleteVerifyResultResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeleteVerifyResultResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteVerifyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVerifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVerifyResultResponseBody) SetCode(v string) *DeleteVerifyResultResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVerifyResultResponseBody) SetMessage(v string) *DeleteVerifyResultResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVerifyResultResponseBody) SetRequestId(v string) *DeleteVerifyResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVerifyResultResponseBody) SetResult(v *DeleteVerifyResultResponseBodyResult) *DeleteVerifyResultResponseBody {
	s.Result = v
	return s
}

type DeleteVerifyResultResponseBodyResult struct {
	// example:
	//
	// Y/N
	DeleteResult *string `json:"DeleteResult,omitempty" xml:"DeleteResult,omitempty"`
	// example:
	//
	// 4ab0b***cbde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s DeleteVerifyResultResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteVerifyResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteVerifyResultResponseBodyResult) SetDeleteResult(v string) *DeleteVerifyResultResponseBodyResult {
	s.DeleteResult = &v
	return s
}

func (s *DeleteVerifyResultResponseBodyResult) SetTransactionId(v string) *DeleteVerifyResultResponseBodyResult {
	s.TransactionId = &v
	return s
}

type DeleteVerifyResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVerifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVerifyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVerifyResultResponse) GoString() string {
	return s.String()
}

func (s *DeleteVerifyResultResponse) SetHeaders(v map[string]*string) *DeleteVerifyResultResponse {
	s.Headers = v
	return s
}

func (s *DeleteVerifyResultResponse) SetStatusCode(v int32) *DeleteVerifyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVerifyResultResponse) SetBody(v *DeleteVerifyResultResponseBody) *DeleteVerifyResultResponse {
	s.Body = v
	return s
}

type DescribeAddressLabelsRequest struct {
	// example:
	//
	// 0xabds8292***dskkds
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// ETH
	Coin *string `json:"Coin,omitempty" xml:"Coin,omitempty"`
	// example:
	//
	// dso9322***dsjsd22
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
}

func (s DescribeAddressLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressLabelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddressLabelsRequest) SetAddress(v string) *DescribeAddressLabelsRequest {
	s.Address = &v
	return s
}

func (s *DescribeAddressLabelsRequest) SetCoin(v string) *DescribeAddressLabelsRequest {
	s.Coin = &v
	return s
}

func (s *DescribeAddressLabelsRequest) SetMerchantBizId(v string) *DescribeAddressLabelsRequest {
	s.MerchantBizId = &v
	return s
}

type DescribeAddressLabelsResponseBody struct {
	// example:
	//
	// Success
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeAddressLabelsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAddressLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAddressLabelsResponseBody) SetCode(v string) *DescribeAddressLabelsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAddressLabelsResponseBody) SetData(v *DescribeAddressLabelsResponseBodyData) *DescribeAddressLabelsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAddressLabelsResponseBody) SetMessage(v string) *DescribeAddressLabelsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAddressLabelsResponseBody) SetRequestId(v string) *DescribeAddressLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAddressLabelsResponseBody) SetSuccess(v bool) *DescribeAddressLabelsResponseBody {
	s.Success = &v
	return s
}

type DescribeAddressLabelsResponseBodyData struct {
	LabelList []*string `json:"LabelList,omitempty" xml:"LabelList,omitempty" type:"Repeated"`
}

func (s DescribeAddressLabelsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressLabelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAddressLabelsResponseBodyData) SetLabelList(v []*string) *DescribeAddressLabelsResponseBodyData {
	s.LabelList = v
	return s
}

type DescribeAddressLabelsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAddressLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAddressLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressLabelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAddressLabelsResponse) SetHeaders(v map[string]*string) *DescribeAddressLabelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAddressLabelsResponse) SetStatusCode(v int32) *DescribeAddressLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAddressLabelsResponse) SetBody(v *DescribeAddressLabelsResponseBody) *DescribeAddressLabelsResponse {
	s.Body = v
	return s
}

type DescribeAddressOverviewRequest struct {
	// example:
	//
	// 0xabds***djskjds
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// ETH
	Coin *string `json:"Coin,omitempty" xml:"Coin,omitempty"`
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
}

func (s DescribeAddressOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddressOverviewRequest) SetAddress(v string) *DescribeAddressOverviewRequest {
	s.Address = &v
	return s
}

func (s *DescribeAddressOverviewRequest) SetCoin(v string) *DescribeAddressOverviewRequest {
	s.Coin = &v
	return s
}

func (s *DescribeAddressOverviewRequest) SetMerchantBizId(v string) *DescribeAddressOverviewRequest {
	s.MerchantBizId = &v
	return s
}

type DescribeAddressOverviewResponseBody struct {
	// example:
	//
	// Success
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeAddressOverviewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAddressOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAddressOverviewResponseBody) SetCode(v string) *DescribeAddressOverviewResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAddressOverviewResponseBody) SetData(v *DescribeAddressOverviewResponseBodyData) *DescribeAddressOverviewResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAddressOverviewResponseBody) SetMessage(v string) *DescribeAddressOverviewResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAddressOverviewResponseBody) SetRequestId(v string) *DescribeAddressOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAddressOverviewResponseBody) SetSuccess(v bool) *DescribeAddressOverviewResponseBody {
	s.Success = &v
	return s
}

type DescribeAddressOverviewResponseBodyData struct {
	// example:
	//
	// 54.4216
	Balance *float32 `json:"Balance,omitempty" xml:"Balance,omitempty"`
	// example:
	//
	// 1441800674
	FirstSeen *int64 `json:"FirstSeen,omitempty" xml:"FirstSeen,omitempty"`
	// example:
	//
	// 1682056727
	LastSeen *int64 `json:"LastSeen,omitempty" xml:"LastSeen,omitempty"`
	// example:
	//
	// 1470
	ReceivedTxsCount *int32 `json:"ReceivedTxsCount,omitempty" xml:"ReceivedTxsCount,omitempty"`
	// example:
	//
	// 227
	SpentTxsCount *int32 `json:"SpentTxsCount,omitempty" xml:"SpentTxsCount,omitempty"`
	// example:
	//
	// 916263.8
	TotalReceived *float32 `json:"TotalReceived,omitempty" xml:"TotalReceived,omitempty"`
	// example:
	//
	// 916205.4
	TotalSpent *float32 `json:"TotalSpent,omitempty" xml:"TotalSpent,omitempty"`
	// example:
	//
	// 1697
	TxsCount *int64 `json:"TxsCount,omitempty" xml:"TxsCount,omitempty"`
}

func (s DescribeAddressOverviewResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAddressOverviewResponseBodyData) SetBalance(v float32) *DescribeAddressOverviewResponseBodyData {
	s.Balance = &v
	return s
}

func (s *DescribeAddressOverviewResponseBodyData) SetFirstSeen(v int64) *DescribeAddressOverviewResponseBodyData {
	s.FirstSeen = &v
	return s
}

func (s *DescribeAddressOverviewResponseBodyData) SetLastSeen(v int64) *DescribeAddressOverviewResponseBodyData {
	s.LastSeen = &v
	return s
}

func (s *DescribeAddressOverviewResponseBodyData) SetReceivedTxsCount(v int32) *DescribeAddressOverviewResponseBodyData {
	s.ReceivedTxsCount = &v
	return s
}

func (s *DescribeAddressOverviewResponseBodyData) SetSpentTxsCount(v int32) *DescribeAddressOverviewResponseBodyData {
	s.SpentTxsCount = &v
	return s
}

func (s *DescribeAddressOverviewResponseBodyData) SetTotalReceived(v float32) *DescribeAddressOverviewResponseBodyData {
	s.TotalReceived = &v
	return s
}

func (s *DescribeAddressOverviewResponseBodyData) SetTotalSpent(v float32) *DescribeAddressOverviewResponseBodyData {
	s.TotalSpent = &v
	return s
}

func (s *DescribeAddressOverviewResponseBodyData) SetTxsCount(v int64) *DescribeAddressOverviewResponseBodyData {
	s.TxsCount = &v
	return s
}

type DescribeAddressOverviewResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAddressOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAddressOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeAddressOverviewResponse) SetHeaders(v map[string]*string) *DescribeAddressOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeAddressOverviewResponse) SetStatusCode(v int32) *DescribeAddressOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAddressOverviewResponse) SetBody(v *DescribeAddressOverviewResponseBody) *DescribeAddressOverviewResponse {
	s.Body = v
	return s
}

type DescribeMaliciousAddressRequest struct {
	// example:
	//
	// BTC
	Coin *string `json:"Coin,omitempty" xml:"Coin,omitempty"`
	// example:
	//
	// 2017-08-20
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// djs20dsjk2-dsjd29-dsskc
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 2017-07-20
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeMaliciousAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMaliciousAddressRequest) GoString() string {
	return s.String()
}

func (s *DescribeMaliciousAddressRequest) SetCoin(v string) *DescribeMaliciousAddressRequest {
	s.Coin = &v
	return s
}

func (s *DescribeMaliciousAddressRequest) SetEnd(v string) *DescribeMaliciousAddressRequest {
	s.End = &v
	return s
}

func (s *DescribeMaliciousAddressRequest) SetMerchantBizId(v string) *DescribeMaliciousAddressRequest {
	s.MerchantBizId = &v
	return s
}

func (s *DescribeMaliciousAddressRequest) SetStart(v string) *DescribeMaliciousAddressRequest {
	s.Start = &v
	return s
}

type DescribeMaliciousAddressResponseBody struct {
	// example:
	//
	// Success
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*DescribeMaliciousAddressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMaliciousAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMaliciousAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMaliciousAddressResponseBody) SetCode(v string) *DescribeMaliciousAddressResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMaliciousAddressResponseBody) SetData(v []*DescribeMaliciousAddressResponseBodyData) *DescribeMaliciousAddressResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMaliciousAddressResponseBody) SetMessage(v string) *DescribeMaliciousAddressResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMaliciousAddressResponseBody) SetRequestId(v string) *DescribeMaliciousAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMaliciousAddressResponseBody) SetSuccess(v bool) *DescribeMaliciousAddressResponseBody {
	s.Success = &v
	return s
}

type DescribeMaliciousAddressResponseBodyData struct {
	// example:
	//
	// 2017-07-18 00:00:00
	AddTime *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	// example:
	//
	// f6d239ff***df816
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// ETH
	Coin *string `json:"Coin,omitempty" xml:"Coin,omitempty"`
	// example:
	//
	// FAKE_Coindash_2
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Tag    *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DescribeMaliciousAddressResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMaliciousAddressResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMaliciousAddressResponseBodyData) SetAddTime(v string) *DescribeMaliciousAddressResponseBodyData {
	s.AddTime = &v
	return s
}

func (s *DescribeMaliciousAddressResponseBodyData) SetAddress(v string) *DescribeMaliciousAddressResponseBodyData {
	s.Address = &v
	return s
}

func (s *DescribeMaliciousAddressResponseBodyData) SetCoin(v string) *DescribeMaliciousAddressResponseBodyData {
	s.Coin = &v
	return s
}

func (s *DescribeMaliciousAddressResponseBodyData) SetDetail(v string) *DescribeMaliciousAddressResponseBodyData {
	s.Detail = &v
	return s
}

func (s *DescribeMaliciousAddressResponseBodyData) SetTag(v string) *DescribeMaliciousAddressResponseBodyData {
	s.Tag = &v
	return s
}

type DescribeMaliciousAddressResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMaliciousAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMaliciousAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMaliciousAddressResponse) GoString() string {
	return s.String()
}

func (s *DescribeMaliciousAddressResponse) SetHeaders(v map[string]*string) *DescribeMaliciousAddressResponse {
	s.Headers = v
	return s
}

func (s *DescribeMaliciousAddressResponse) SetStatusCode(v int32) *DescribeMaliciousAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMaliciousAddressResponse) SetBody(v *DescribeMaliciousAddressResponseBody) *DescribeMaliciousAddressResponse {
	s.Body = v
	return s
}

type DescribeRiskScoreRequest struct {
	// example:
	//
	// 0xabds***djskjds
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// ETH
	Coin *string `json:"Coin,omitempty" xml:"Coin,omitempty"`
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
}

func (s DescribeRiskScoreRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskScoreRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskScoreRequest) SetAddress(v string) *DescribeRiskScoreRequest {
	s.Address = &v
	return s
}

func (s *DescribeRiskScoreRequest) SetCoin(v string) *DescribeRiskScoreRequest {
	s.Coin = &v
	return s
}

func (s *DescribeRiskScoreRequest) SetMerchantBizId(v string) *DescribeRiskScoreRequest {
	s.MerchantBizId = &v
	return s
}

type DescribeRiskScoreResponseBody struct {
	// example:
	//
	// Success
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeRiskScoreResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRiskScoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskScoreResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskScoreResponseBody) SetCode(v string) *DescribeRiskScoreResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRiskScoreResponseBody) SetData(v *DescribeRiskScoreResponseBodyData) *DescribeRiskScoreResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRiskScoreResponseBody) SetMessage(v string) *DescribeRiskScoreResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRiskScoreResponseBody) SetRequestId(v string) *DescribeRiskScoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskScoreResponseBody) SetSuccess(v bool) *DescribeRiskScoreResponseBody {
	s.Success = &v
	return s
}

type DescribeRiskScoreResponseBodyData struct {
	DetailList []*string `json:"DetailList,omitempty" xml:"DetailList,omitempty" type:"Repeated"`
	// example:
	//
	// MMFinance Exploiter
	HackingEvent *string `json:"HackingEvent,omitempty" xml:"HackingEvent,omitempty"`
	// example:
	//
	// Severe
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// 100
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s DescribeRiskScoreResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskScoreResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRiskScoreResponseBodyData) SetDetailList(v []*string) *DescribeRiskScoreResponseBodyData {
	s.DetailList = v
	return s
}

func (s *DescribeRiskScoreResponseBodyData) SetHackingEvent(v string) *DescribeRiskScoreResponseBodyData {
	s.HackingEvent = &v
	return s
}

func (s *DescribeRiskScoreResponseBodyData) SetRiskLevel(v string) *DescribeRiskScoreResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *DescribeRiskScoreResponseBodyData) SetScore(v int32) *DescribeRiskScoreResponseBodyData {
	s.Score = &v
	return s
}

type DescribeRiskScoreResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskScoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskScoreResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskScoreResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskScoreResponse) SetHeaders(v map[string]*string) *DescribeRiskScoreResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskScoreResponse) SetStatusCode(v int32) *DescribeRiskScoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskScoreResponse) SetBody(v *DescribeRiskScoreResponseBody) *DescribeRiskScoreResponse {
	s.Body = v
	return s
}

type DescribeTransactionsListRequest struct {
	// example:
	//
	// 0xabds***djskjds
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// ETH
	Coin *string `json:"Coin,omitempty" xml:"Coin,omitempty"`
	// example:
	//
	// 1682235649
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 1682232649
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// example:
	//
	// in
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTransactionsListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransactionsListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTransactionsListRequest) SetAddress(v string) *DescribeTransactionsListRequest {
	s.Address = &v
	return s
}

func (s *DescribeTransactionsListRequest) SetCoin(v string) *DescribeTransactionsListRequest {
	s.Coin = &v
	return s
}

func (s *DescribeTransactionsListRequest) SetEndTimestamp(v int64) *DescribeTransactionsListRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeTransactionsListRequest) SetMerchantBizId(v string) *DescribeTransactionsListRequest {
	s.MerchantBizId = &v
	return s
}

func (s *DescribeTransactionsListRequest) SetPage(v int64) *DescribeTransactionsListRequest {
	s.Page = &v
	return s
}

func (s *DescribeTransactionsListRequest) SetStartTimestamp(v int64) *DescribeTransactionsListRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeTransactionsListRequest) SetType(v string) *DescribeTransactionsListRequest {
	s.Type = &v
	return s
}

type DescribeTransactionsListResponseBody struct {
	// example:
	//
	// Success
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeTransactionsListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeTransactionsListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransactionsListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTransactionsListResponseBody) SetCode(v string) *DescribeTransactionsListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTransactionsListResponseBody) SetData(v *DescribeTransactionsListResponseBodyData) *DescribeTransactionsListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTransactionsListResponseBody) SetMessage(v string) *DescribeTransactionsListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTransactionsListResponseBody) SetRequestId(v string) *DescribeTransactionsListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTransactionsListResponseBody) SetSuccess(v bool) *DescribeTransactionsListResponseBody {
	s.Success = &v
	return s
}

type DescribeTransactionsListResponseBodyData struct {
	In  []*DescribeTransactionsListResponseBodyDataIn  `json:"In,omitempty" xml:"In,omitempty" type:"Repeated"`
	Out []*DescribeTransactionsListResponseBodyDataOut `json:"Out,omitempty" xml:"Out,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 6
	TotalPages *int64 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	// example:
	//
	// 300
	TransactionsOnPage *int64 `json:"TransactionsOnPage,omitempty" xml:"TransactionsOnPage,omitempty"`
}

func (s DescribeTransactionsListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransactionsListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTransactionsListResponseBodyData) SetIn(v []*DescribeTransactionsListResponseBodyDataIn) *DescribeTransactionsListResponseBodyData {
	s.In = v
	return s
}

func (s *DescribeTransactionsListResponseBodyData) SetOut(v []*DescribeTransactionsListResponseBodyDataOut) *DescribeTransactionsListResponseBodyData {
	s.Out = v
	return s
}

func (s *DescribeTransactionsListResponseBodyData) SetPage(v int64) *DescribeTransactionsListResponseBodyData {
	s.Page = &v
	return s
}

func (s *DescribeTransactionsListResponseBodyData) SetTotalPages(v int64) *DescribeTransactionsListResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *DescribeTransactionsListResponseBodyData) SetTransactionsOnPage(v int64) *DescribeTransactionsListResponseBodyData {
	s.TransactionsOnPage = &v
	return s
}

type DescribeTransactionsListResponseBodyDataIn struct {
	// example:
	//
	// 0xd90e2***b90ad053324f31b
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// 0.0845
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// bi***ce
	Label      *string   `json:"Label,omitempty" xml:"Label,omitempty"`
	TxHashList []*string `json:"TxHashList,omitempty" xml:"TxHashList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTransactionsListResponseBodyDataIn) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransactionsListResponseBodyDataIn) GoString() string {
	return s.String()
}

func (s *DescribeTransactionsListResponseBodyDataIn) SetAddress(v string) *DescribeTransactionsListResponseBodyDataIn {
	s.Address = &v
	return s
}

func (s *DescribeTransactionsListResponseBodyDataIn) SetAmount(v float32) *DescribeTransactionsListResponseBodyDataIn {
	s.Amount = &v
	return s
}

func (s *DescribeTransactionsListResponseBodyDataIn) SetLabel(v string) *DescribeTransactionsListResponseBodyDataIn {
	s.Label = &v
	return s
}

func (s *DescribeTransactionsListResponseBodyDataIn) SetTxHashList(v []*string) *DescribeTransactionsListResponseBodyDataIn {
	s.TxHashList = v
	return s
}

func (s *DescribeTransactionsListResponseBodyDataIn) SetType(v int32) *DescribeTransactionsListResponseBodyDataIn {
	s.Type = &v
	return s
}

type DescribeTransactionsListResponseBodyDataOut struct {
	// example:
	//
	// 0xd90e2***b90ad053324f31b
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// 743
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// Tor***uter
	Label      *string   `json:"Label,omitempty" xml:"Label,omitempty"`
	TxHashList []*string `json:"TxHashList,omitempty" xml:"TxHashList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTransactionsListResponseBodyDataOut) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransactionsListResponseBodyDataOut) GoString() string {
	return s.String()
}

func (s *DescribeTransactionsListResponseBodyDataOut) SetAddress(v string) *DescribeTransactionsListResponseBodyDataOut {
	s.Address = &v
	return s
}

func (s *DescribeTransactionsListResponseBodyDataOut) SetAmount(v float32) *DescribeTransactionsListResponseBodyDataOut {
	s.Amount = &v
	return s
}

func (s *DescribeTransactionsListResponseBodyDataOut) SetLabel(v string) *DescribeTransactionsListResponseBodyDataOut {
	s.Label = &v
	return s
}

func (s *DescribeTransactionsListResponseBodyDataOut) SetTxHashList(v []*string) *DescribeTransactionsListResponseBodyDataOut {
	s.TxHashList = v
	return s
}

func (s *DescribeTransactionsListResponseBodyDataOut) SetType(v int32) *DescribeTransactionsListResponseBodyDataOut {
	s.Type = &v
	return s
}

type DescribeTransactionsListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTransactionsListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTransactionsListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransactionsListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTransactionsListResponse) SetHeaders(v map[string]*string) *DescribeTransactionsListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTransactionsListResponse) SetStatusCode(v int32) *DescribeTransactionsListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTransactionsListResponse) SetBody(v *DescribeTransactionsListResponseBody) *DescribeTransactionsListResponse {
	s.Body = v
	return s
}

type DescribeWeb3AddressLabelsRequest struct {
	// The address hash.
	//
	// example:
	//
	// 0000980145045a5c5acad3d2df0cf3b2afxxxxxx
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// This is the short name of blockchain。
	//
	// [ ETH, MATIC, BNB ]
	//
	// example:
	//
	// ETH
	ChainShortName *string `json:"ChainShortName,omitempty" xml:"ChainShortName,omitempty"`
	// A unique business ID for tracing purpose. For example，the sequence ID from the merchant\\"s business-related database.
	//
	// example:
	//
	// dso932dsjsd22
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
}

func (s DescribeWeb3AddressLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3AddressLabelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWeb3AddressLabelsRequest) SetAddress(v string) *DescribeWeb3AddressLabelsRequest {
	s.Address = &v
	return s
}

func (s *DescribeWeb3AddressLabelsRequest) SetChainShortName(v string) *DescribeWeb3AddressLabelsRequest {
	s.ChainShortName = &v
	return s
}

func (s *DescribeWeb3AddressLabelsRequest) SetMerchantBizId(v string) *DescribeWeb3AddressLabelsRequest {
	s.MerchantBizId = &v
	return s
}

type DescribeWeb3AddressLabelsResponseBody struct {
	// Return code. For the full list of codes, see Codes and Messages.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data *DescribeWeb3AddressLabelsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response detailed message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request, which can be used to locate issues.
	//
	// example:
	//
	// 7F971622-38C0-5F56-B2EC-315367979B4F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWeb3AddressLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3AddressLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWeb3AddressLabelsResponseBody) SetCode(v string) *DescribeWeb3AddressLabelsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBody) SetData(v *DescribeWeb3AddressLabelsResponseBodyData) *DescribeWeb3AddressLabelsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBody) SetHttpStatusCode(v int32) *DescribeWeb3AddressLabelsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBody) SetMessage(v string) *DescribeWeb3AddressLabelsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBody) SetRequestId(v string) *DescribeWeb3AddressLabelsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeWeb3AddressLabelsResponseBodyData struct {
	// address
	//
	// example:
	//
	// 2341980145045A5c5acad3d2Df0cF3B2Afxxxxxx
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// amount of native currency
	//
	// example:
	//
	// 0
	Balance *string `json:"Balance,omitempty" xml:"Balance,omitempty"`
	// native currency of the chain
	//
	// example:
	//
	// ETH
	BalanceSymbol *string `json:"BalanceSymbol,omitempty" xml:"BalanceSymbol,omitempty"`
	// ChainNameEnumstring name of blockchain
	//
	// example:
	//
	// Ethereum Mainnet
	ChainName *string `json:"ChainName,omitempty" xml:"ChainName,omitempty"`
	// ChainShortName
	//
	// example:
	//
	// eth
	ChainShortName *string `json:"ChainShortName,omitempty" xml:"ChainShortName,omitempty"`
	// 0: EOA; 1: Contract
	//
	// example:
	//
	// 0
	ContractAddress *string `json:"ContractAddress,omitempty" xml:"ContractAddress,omitempty"`
	// the address of deployer if the current address is a contract, null if the current address is an EOA
	//
	// example:
	//
	// dAC17F958D2ee523a2206206994597C13Dxxxxxx
	CreateContractAddress *string `json:"CreateContractAddress,omitempty" xml:"CreateContractAddress,omitempty"`
	// contract creation hash if the current address is a contract, null if the current address is an EOA
	//
	// example:
	//
	// dAC17F958D2ee523a2206206994597C13Dxxxxxx
	CreateContractTransactionHash *string `json:"CreateContractTransactionHash,omitempty" xml:"CreateContractTransactionHash,omitempty"`
	// customized assessment detail
	//
	// example:
	//
	// ""
	CustomRiskAssessment *string `json:"CustomRiskAssessment,omitempty" xml:"CustomRiskAssessment,omitempty"`
	// the first transaction hash sent by the address
	//
	// example:
	//
	// 20230304
	FirstTransactionTime *string `json:"FirstTransactionTime,omitempty" xml:"FirstTransactionTime,omitempty"`
	// 0: Not validator; 1: validator
	//
	// example:
	//
	// 0
	IsProducerAddress *string `json:"IsProducerAddress,omitempty" xml:"IsProducerAddress,omitempty"`
	// the latest transaction hash sent by the address
	//
	// example:
	//
	// 20230304
	LastTransactionTime *string `json:"LastTransactionTime,omitempty" xml:"LastTransactionTime,omitempty"`
	// the amount of native currency received in 180 days
	//
	// example:
	//
	// 0
	ReceiveAmount *string `json:"ReceiveAmount,omitempty" xml:"ReceiveAmount,omitempty"`
	// the amount of native currency sent in 180 days
	//
	// example:
	//
	// 0
	SendAmount *string `json:"SendAmount,omitempty" xml:"SendAmount,omitempty"`
	// tag
	//
	// example:
	//
	// contracts:Tether: Tether_USD,token_standard:erc20
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// if the address is an erc20 token, returns the token name
	//
	// example:
	//
	// Tether USD
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// the number of erc20 tokens involved with current address in 180 days
	//
	// example:
	//
	// 1
	TokenAmount *int32 `json:"TokenAmount,omitempty" xml:"TokenAmount,omitempty"`
	// address of erc20 tokens involved with current address in 180 days
	//
	// example:
	//
	// ["{"ERC721":[]"]}"]
	TokenList *string `json:"TokenList,omitempty" xml:"TokenList,omitempty"`
	// the number of transactions
	//
	// example:
	//
	// 2
	TransactionCount *int32 `json:"TransactionCount,omitempty" xml:"TransactionCount,omitempty"`
}

func (s DescribeWeb3AddressLabelsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3AddressLabelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetAddress(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.Address = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetBalance(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.Balance = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetBalanceSymbol(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.BalanceSymbol = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetChainName(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.ChainName = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetChainShortName(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.ChainShortName = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetContractAddress(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.ContractAddress = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetCreateContractAddress(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.CreateContractAddress = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetCreateContractTransactionHash(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.CreateContractTransactionHash = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetCustomRiskAssessment(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.CustomRiskAssessment = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetFirstTransactionTime(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.FirstTransactionTime = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetIsProducerAddress(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.IsProducerAddress = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetLastTransactionTime(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.LastTransactionTime = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetReceiveAmount(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.ReceiveAmount = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetSendAmount(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.SendAmount = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetTag(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.Tag = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetToken(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.Token = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetTokenAmount(v int32) *DescribeWeb3AddressLabelsResponseBodyData {
	s.TokenAmount = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetTokenList(v string) *DescribeWeb3AddressLabelsResponseBodyData {
	s.TokenList = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponseBodyData) SetTransactionCount(v int32) *DescribeWeb3AddressLabelsResponseBodyData {
	s.TransactionCount = &v
	return s
}

type DescribeWeb3AddressLabelsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWeb3AddressLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWeb3AddressLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3AddressLabelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWeb3AddressLabelsResponse) SetHeaders(v map[string]*string) *DescribeWeb3AddressLabelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWeb3AddressLabelsResponse) SetStatusCode(v int32) *DescribeWeb3AddressLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWeb3AddressLabelsResponse) SetBody(v *DescribeWeb3AddressLabelsResponseBody) *DescribeWeb3AddressLabelsResponse {
	s.Body = v
	return s
}

type DescribeWeb3RiskScoreRequest struct {
	// This is the short name of blockchain。
	//
	// [ ETH, MATIC, BNB ]
	//
	// example:
	//
	// ETH
	ChainShortName *string `json:"ChainShortName,omitempty" xml:"ChainShortName,omitempty"`
	// minimum: 1
	//
	// maximum: 100
	//
	// the maximum depth for risk analysis. For UTXO-based blockchains, default and maximum is enforced at 100.
	//
	// For account-based blockchains, default and maximum is enforced at 6
	//
	// example:
	//
	// 2
	Depth *int32 `json:"Depth,omitempty" xml:"Depth,omitempty"`
	// A unique business ID for tracing purpose. For example，the sequence ID from the merchant\\"s business-related database.
	//
	// example:
	//
	// e0c34a353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// For TRANSACTION objects, you need to provide the transaction hash。
	//
	// For ADDRESS objects, you need to provide the address or reference address hash。
	//
	// example:
	//
	// 0000980145045a5c5acad3d2df0cf3b2afxxxxxx
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The object of the analysis.
	//
	// [ TRANSACTION, ADDRESS ]
	//
	// example:
	//
	// TRANSACTION
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s DescribeWeb3RiskScoreRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3RiskScoreRequest) GoString() string {
	return s.String()
}

func (s *DescribeWeb3RiskScoreRequest) SetChainShortName(v string) *DescribeWeb3RiskScoreRequest {
	s.ChainShortName = &v
	return s
}

func (s *DescribeWeb3RiskScoreRequest) SetDepth(v int32) *DescribeWeb3RiskScoreRequest {
	s.Depth = &v
	return s
}

func (s *DescribeWeb3RiskScoreRequest) SetMerchantBizId(v string) *DescribeWeb3RiskScoreRequest {
	s.MerchantBizId = &v
	return s
}

func (s *DescribeWeb3RiskScoreRequest) SetObjectId(v string) *DescribeWeb3RiskScoreRequest {
	s.ObjectId = &v
	return s
}

func (s *DescribeWeb3RiskScoreRequest) SetObjectType(v string) *DescribeWeb3RiskScoreRequest {
	s.ObjectType = &v
	return s
}

type DescribeWeb3RiskScoreResponseBody struct {
	// Return code. For the full list of codes, see Codes and Messages.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data *DescribeWeb3RiskScoreResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response detailed message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request, which can be used to locate issues.
	//
	// example:
	//
	// 5E63B760-0ECB-5C07-8503-A65C27876968
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWeb3RiskScoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3RiskScoreResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWeb3RiskScoreResponseBody) SetCode(v string) *DescribeWeb3RiskScoreResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeWeb3RiskScoreResponseBody) SetData(v *DescribeWeb3RiskScoreResponseBodyData) *DescribeWeb3RiskScoreResponseBody {
	s.Data = v
	return s
}

func (s *DescribeWeb3RiskScoreResponseBody) SetHttpStatusCode(v int32) *DescribeWeb3RiskScoreResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeWeb3RiskScoreResponseBody) SetMessage(v string) *DescribeWeb3RiskScoreResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeWeb3RiskScoreResponseBody) SetRequestId(v string) *DescribeWeb3RiskScoreResponseBody {
	s.RequestId = &v
	return s
}

type DescribeWeb3RiskScoreResponseBodyData struct {
	// risk results
	RiskResults []*DescribeWeb3RiskScoreResponseBodyDataRiskResults `json:"RiskResults,omitempty" xml:"RiskResults,omitempty" type:"Repeated"`
	// Risk score
	//
	// example:
	//
	// 80
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s DescribeWeb3RiskScoreResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3RiskScoreResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeWeb3RiskScoreResponseBodyData) SetRiskResults(v []*DescribeWeb3RiskScoreResponseBodyDataRiskResults) *DescribeWeb3RiskScoreResponseBodyData {
	s.RiskResults = v
	return s
}

func (s *DescribeWeb3RiskScoreResponseBodyData) SetScore(v string) *DescribeWeb3RiskScoreResponseBodyData {
	s.Score = &v
	return s
}

type DescribeWeb3RiskScoreResponseBodyDataRiskResults struct {
	// description
	//
	// example:
	//
	// incoming address risk critical xxxxxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// [ CRITICAL, HIGH, MEDIUM, LOW, NO ]
	//
	// 100: Critical
	//
	// 67-99: High risk
	//
	// 34-66: Medium risk
	//
	// 1-33: Low risk
	//
	// 0: No risk
	//
	// example:
	//
	// HIGH
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// Address
	//
	// Transaction
	//
	// example:
	//
	// Address
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeWeb3RiskScoreResponseBodyDataRiskResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3RiskScoreResponseBodyDataRiskResults) GoString() string {
	return s.String()
}

func (s *DescribeWeb3RiskScoreResponseBodyDataRiskResults) SetDescription(v string) *DescribeWeb3RiskScoreResponseBodyDataRiskResults {
	s.Description = &v
	return s
}

func (s *DescribeWeb3RiskScoreResponseBodyDataRiskResults) SetSeverity(v string) *DescribeWeb3RiskScoreResponseBodyDataRiskResults {
	s.Severity = &v
	return s
}

func (s *DescribeWeb3RiskScoreResponseBodyDataRiskResults) SetType(v string) *DescribeWeb3RiskScoreResponseBodyDataRiskResults {
	s.Type = &v
	return s
}

type DescribeWeb3RiskScoreResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWeb3RiskScoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWeb3RiskScoreResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3RiskScoreResponse) GoString() string {
	return s.String()
}

func (s *DescribeWeb3RiskScoreResponse) SetHeaders(v map[string]*string) *DescribeWeb3RiskScoreResponse {
	s.Headers = v
	return s
}

func (s *DescribeWeb3RiskScoreResponse) SetStatusCode(v int32) *DescribeWeb3RiskScoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWeb3RiskScoreResponse) SetBody(v *DescribeWeb3RiskScoreResponseBody) *DescribeWeb3RiskScoreResponse {
	s.Body = v
	return s
}

type DescribeWeb3TransactionLabelsRequest struct {
	// This is the short name of blockchain。
	//
	// [ ETH, MATIC, BNB ]
	//
	// example:
	//
	// ETH
	ChainShortName *string `json:"ChainShortName,omitempty" xml:"ChainShortName,omitempty"`
	// A unique business ID for tracing purpose. For example，the sequence ID from the merchant\\"s business-related database.
	//
	// example:
	//
	// e0c34a353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The Transaction hash.
	//
	// example:
	//
	// c92880148d4896d8a2093a891a8f08916fe141fba474ede4101f81fb2bxxxxxx
	Transaction *string `json:"Transaction,omitempty" xml:"Transaction,omitempty"`
}

func (s DescribeWeb3TransactionLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3TransactionLabelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWeb3TransactionLabelsRequest) SetChainShortName(v string) *DescribeWeb3TransactionLabelsRequest {
	s.ChainShortName = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsRequest) SetMerchantBizId(v string) *DescribeWeb3TransactionLabelsRequest {
	s.MerchantBizId = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsRequest) SetTransaction(v string) *DescribeWeb3TransactionLabelsRequest {
	s.Transaction = &v
	return s
}

type DescribeWeb3TransactionLabelsResponseBody struct {
	// Return code. For the full list of codes, see Codes and Messages.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data *DescribeWeb3TransactionLabelsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response detailed message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request, which can be used to locate issues.
	//
	// example:
	//
	// 7A0D192A-CC0C-5DE5-A3B6-A14CF45508F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWeb3TransactionLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3TransactionLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWeb3TransactionLabelsResponseBody) SetCode(v string) *DescribeWeb3TransactionLabelsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBody) SetData(v *DescribeWeb3TransactionLabelsResponseBodyData) *DescribeWeb3TransactionLabelsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBody) SetHttpStatusCode(v int32) *DescribeWeb3TransactionLabelsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBody) SetMessage(v string) *DescribeWeb3TransactionLabelsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBody) SetRequestId(v string) *DescribeWeb3TransactionLabelsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeWeb3TransactionLabelsResponseBodyData struct {
	// the amount of native currency
	//
	// example:
	//
	// 27
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// chainName
	//
	// example:
	//
	// ETH
	ChainName *string `json:"ChainName,omitempty" xml:"ChainName,omitempty"`
	// short name of blockchain
	//
	// example:
	//
	// eth
	ChainShortName *string `json:"ChainShortName,omitempty" xml:"ChainShortName,omitempty"`
	// contract details
	ContractDetails []*DescribeWeb3TransactionLabelsResponseBodyDataContractDetails `json:"ContractDetails,omitempty" xml:"ContractDetails,omitempty" type:"Repeated"`
	// error log
	//
	// example:
	//
	// “”
	ErrorLog *string `json:"ErrorLog,omitempty" xml:"ErrorLog,omitempty"`
	// gasLimit
	//
	// example:
	//
	// 1232
	GasLimit *int32 `json:"GasLimit,omitempty" xml:"GasLimit,omitempty"`
	// gasPrice
	//
	// example:
	//
	// 5034
	GasPrice *string `json:"GasPrice,omitempty" xml:"GasPrice,omitempty"`
	// gasUsed
	//
	// example:
	//
	// 234
	GasUsed *int32 `json:"GasUsed,omitempty" xml:"GasUsed,omitempty"`
	// height
	//
	// example:
	//
	// 17087552
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// the position of the transaction in the block
	//
	// example:
	//
	// 94
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// input data
	//
	// example:
	//
	// a9059cbb000000000000000000000000c7019579cb9bdb9204e61c7179ba2f88f9d2990b000000000xxxxxx
	InputData *string `json:"InputData,omitempty" xml:"InputData,omitempty"`
	// input details
	InputDetails []*DescribeWeb3TransactionLabelsResponseBodyDataInputDetails `json:"InputDetails,omitempty" xml:"InputDetails,omitempty" type:"Repeated"`
	// the method name of contract call. For external transaction method: [\\"CALL\\",\\"CALLCODE\\",\\"DELEGATECALL\\",\\"STATICCALL\\"]; for internal transaction method: the first 4 bytes of the hash of the method name
	//
	// example:
	//
	// a9059cbb
	MethodId *string `json:"MethodId,omitempty" xml:"MethodId,omitempty"`
	// nonce
	//
	// example:
	//
	// 6242724
	Nonce *string `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
	// output details
	OutputDetails []*DescribeWeb3TransactionLabelsResponseBodyDataOutputDetails `json:"OutputDetails,omitempty" xml:"OutputDetails,omitempty" type:"Repeated"`
	// the transaction state. 1: success 0: fail
	//
	// example:
	//
	// 1
	State *int32 `json:"State,omitempty" xml:"State,omitempty"`
	// token transfer details
	TokenTransferDetails []*DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails `json:"TokenTransferDetails,omitempty" xml:"TokenTransferDetails,omitempty" type:"Repeated"`
	// the symbol of native currency
	//
	// example:
	//
	// ETH
	TransactionSymbol *string `json:"TransactionSymbol,omitempty" xml:"TransactionSymbol,omitempty"`
	// the block timestamp
	//
	// example:
	//
	// 1681991807
	TransactionTime *string `json:"TransactionTime,omitempty" xml:"TransactionTime,omitempty"`
	// Integer	0: legacy; 1: eip 2930; 2: eip 1559
	//
	// example:
	//
	// 1
	TransactionType *string `json:"TransactionType,omitempty" xml:"TransactionType,omitempty"`
	// the transaction fee in eth
	//
	// example:
	//
	// 0.002321489548255059
	Txfee *string `json:"Txfee,omitempty" xml:"Txfee,omitempty"`
	// Txid
	//
	// example:
	//
	// c92880148d4896d8a2093a891a8f08916fe141fba474ede410xxxxxx
	Txid *string `json:"Txid,omitempty" xml:"Txid,omitempty"`
}

func (s DescribeWeb3TransactionLabelsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3TransactionLabelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetAmount(v string) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.Amount = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetChainName(v string) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.ChainName = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetChainShortName(v string) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.ChainShortName = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetContractDetails(v []*DescribeWeb3TransactionLabelsResponseBodyDataContractDetails) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.ContractDetails = v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetErrorLog(v string) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.ErrorLog = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetGasLimit(v int32) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.GasLimit = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetGasPrice(v string) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.GasPrice = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetGasUsed(v int32) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.GasUsed = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetHeight(v int32) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.Height = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetIndex(v int32) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.Index = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetInputData(v string) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.InputData = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetInputDetails(v []*DescribeWeb3TransactionLabelsResponseBodyDataInputDetails) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.InputDetails = v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetMethodId(v string) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.MethodId = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetNonce(v string) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.Nonce = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetOutputDetails(v []*DescribeWeb3TransactionLabelsResponseBodyDataOutputDetails) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.OutputDetails = v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetState(v int32) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.State = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetTokenTransferDetails(v []*DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.TokenTransferDetails = v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetTransactionSymbol(v string) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.TransactionSymbol = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetTransactionTime(v string) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.TransactionTime = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetTransactionType(v string) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.TransactionType = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetTxfee(v string) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.Txfee = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyData) SetTxid(v string) *DescribeWeb3TransactionLabelsResponseBodyData {
	s.Txid = &v
	return s
}

type DescribeWeb3TransactionLabelsResponseBodyDataContractDetails struct {
	// the value of internal transaction
	//
	// example:
	//
	// 945.5
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// the sender of internal transaction
	//
	// example:
	//
	// C7019579cB9bdb9204e61C7179ba2F88F9dxxxxxx
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// the gaslimit of internal transaction
	//
	// example:
	//
	// 20712
	GasLimit *int32 `json:"GasLimit,omitempty" xml:"GasLimit,omitempty"`
	// the call layer of internal transaction
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// the receiver of internal transaction
	//
	// example:
	//
	// C7019579cB9bdb9204e61C7179ba2F88F9dxxxxxx
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s DescribeWeb3TransactionLabelsResponseBodyDataContractDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3TransactionLabelsResponseBodyDataContractDetails) GoString() string {
	return s.String()
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataContractDetails) SetAmount(v string) *DescribeWeb3TransactionLabelsResponseBodyDataContractDetails {
	s.Amount = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataContractDetails) SetFrom(v string) *DescribeWeb3TransactionLabelsResponseBodyDataContractDetails {
	s.From = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataContractDetails) SetGasLimit(v int32) *DescribeWeb3TransactionLabelsResponseBodyDataContractDetails {
	s.GasLimit = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataContractDetails) SetIndex(v int32) *DescribeWeb3TransactionLabelsResponseBodyDataContractDetails {
	s.Index = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataContractDetails) SetTo(v string) *DescribeWeb3TransactionLabelsResponseBodyDataContractDetails {
	s.To = &v
	return s
}

type DescribeWeb3TransactionLabelsResponseBodyDataInputDetails struct {
	// example: 15. the amount of transation sent by the address
	//
	// example:
	//
	// 3234
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// the address hash
	//
	// example:
	//
	// 21a31Ee1afC51d94C2eFcCAa2xxxxxx
	InputHash *string `json:"InputHash,omitempty" xml:"InputHash,omitempty"`
	// example: true. is it a contract
	//
	// example:
	//
	// true
	IsContract *string `json:"IsContract,omitempty" xml:"IsContract,omitempty"`
	// example: Dex . the tag of the address
	//
	// example:
	//
	// contracts:Tether: Tether_USD,token_standard:xxxxxx
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DescribeWeb3TransactionLabelsResponseBodyDataInputDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3TransactionLabelsResponseBodyDataInputDetails) GoString() string {
	return s.String()
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataInputDetails) SetAmount(v int32) *DescribeWeb3TransactionLabelsResponseBodyDataInputDetails {
	s.Amount = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataInputDetails) SetInputHash(v string) *DescribeWeb3TransactionLabelsResponseBodyDataInputDetails {
	s.InputHash = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataInputDetails) SetIsContract(v string) *DescribeWeb3TransactionLabelsResponseBodyDataInputDetails {
	s.IsContract = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataInputDetails) SetTag(v string) *DescribeWeb3TransactionLabelsResponseBodyDataInputDetails {
	s.Tag = &v
	return s
}

type DescribeWeb3TransactionLabelsResponseBodyDataOutputDetails struct {
	// example: 15. the amount of transation sent by the address
	//
	// example:
	//
	// 1500
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// the address hash
	//
	// example:
	//
	// 21a31Ee1afC51d94C2eFcCAa2xxxxxx
	InputHash *string `json:"InputHash,omitempty" xml:"InputHash,omitempty"`
	// example: true. is it a contract
	//
	// example:
	//
	// true
	IsContract *string `json:"IsContract,omitempty" xml:"IsContract,omitempty"`
	// example: Dex. the tag of the address
	//
	// example:
	//
	// contracts:Tether: Tether_USD,token_standard:xxxxxx
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DescribeWeb3TransactionLabelsResponseBodyDataOutputDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3TransactionLabelsResponseBodyDataOutputDetails) GoString() string {
	return s.String()
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataOutputDetails) SetAmount(v int32) *DescribeWeb3TransactionLabelsResponseBodyDataOutputDetails {
	s.Amount = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataOutputDetails) SetInputHash(v string) *DescribeWeb3TransactionLabelsResponseBodyDataOutputDetails {
	s.InputHash = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataOutputDetails) SetIsContract(v string) *DescribeWeb3TransactionLabelsResponseBodyDataOutputDetails {
	s.IsContract = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataOutputDetails) SetTag(v string) *DescribeWeb3TransactionLabelsResponseBodyDataOutputDetails {
	s.Tag = &v
	return s
}

type DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails struct {
	// the token amount
	//
	// example:
	//
	// 945.5
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// the sender of the token
	//
	// example:
	//
	// 21a31Ee1afC51d94C2eFcCAa2092aD1028xxxxxx
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// the call layer of to token transfer
	//
	// example:
	//
	// 0
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// the token symbol
	//
	// example:
	//
	// USDT
	Symbol *string `json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	// the receiver of the token
	//
	// example:
	//
	// C7019579cB9bdb9204e61C7179ba2F88F9dxxxxxx
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// the token name
	//
	// example:
	//
	// Tether USD
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// the token address
	//
	// example:
	//
	// dAC17F958D2ee523a2206206994597C13Dxxxxxx
	TokenContractAddress *string `json:"TokenContractAddress,omitempty" xml:"TokenContractAddress,omitempty"`
	// NFT ID, if the token is erc721
	//
	// example:
	//
	// -1
	TokenId *string `json:"TokenId,omitempty" xml:"TokenId,omitempty"`
}

func (s DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails) GoString() string {
	return s.String()
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails) SetAmount(v string) *DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails {
	s.Amount = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails) SetFrom(v string) *DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails {
	s.From = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails) SetIndex(v int32) *DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails {
	s.Index = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails) SetSymbol(v string) *DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails {
	s.Symbol = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails) SetTo(v string) *DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails {
	s.To = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails) SetToken(v string) *DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails {
	s.Token = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails) SetTokenContractAddress(v string) *DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails {
	s.TokenContractAddress = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails) SetTokenId(v string) *DescribeWeb3TransactionLabelsResponseBodyDataTokenTransferDetails {
	s.TokenId = &v
	return s
}

type DescribeWeb3TransactionLabelsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWeb3TransactionLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWeb3TransactionLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWeb3TransactionLabelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWeb3TransactionLabelsResponse) SetHeaders(v map[string]*string) *DescribeWeb3TransactionLabelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponse) SetStatusCode(v int32) *DescribeWeb3TransactionLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWeb3TransactionLabelsResponse) SetBody(v *DescribeWeb3TransactionLabelsResponseBody) *DescribeWeb3TransactionLabelsResponse {
	s.Body = v
	return s
}

type DocOcrRequest struct {
	// example:
	//
	// 00000006
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// F
	IdFaceQuality      *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// example:
	//
	// https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	IdThreshold     *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// example:
	//
	// dso9322***dsjsd22
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// T
	Ocr         *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// F
	Spoof *string `json:"Spoof,omitempty" xml:"Spoof,omitempty"`
}

func (s DocOcrRequest) String() string {
	return tea.Prettify(s)
}

func (s DocOcrRequest) GoString() string {
	return s.String()
}

func (s *DocOcrRequest) SetDocType(v string) *DocOcrRequest {
	s.DocType = &v
	return s
}

func (s *DocOcrRequest) SetIdFaceQuality(v string) *DocOcrRequest {
	s.IdFaceQuality = &v
	return s
}

func (s *DocOcrRequest) SetIdOcrPictureBase64(v string) *DocOcrRequest {
	s.IdOcrPictureBase64 = &v
	return s
}

func (s *DocOcrRequest) SetIdOcrPictureUrl(v string) *DocOcrRequest {
	s.IdOcrPictureUrl = &v
	return s
}

func (s *DocOcrRequest) SetIdThreshold(v string) *DocOcrRequest {
	s.IdThreshold = &v
	return s
}

func (s *DocOcrRequest) SetMerchantBizId(v string) *DocOcrRequest {
	s.MerchantBizId = &v
	return s
}

func (s *DocOcrRequest) SetMerchantUserId(v string) *DocOcrRequest {
	s.MerchantUserId = &v
	return s
}

func (s *DocOcrRequest) SetOcr(v string) *DocOcrRequest {
	s.Ocr = &v
	return s
}

func (s *DocOcrRequest) SetProductCode(v string) *DocOcrRequest {
	s.ProductCode = &v
	return s
}

func (s *DocOcrRequest) SetSpoof(v string) *DocOcrRequest {
	s.Spoof = &v
	return s
}

type DocOcrResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 86C40EC3-5940-5F47-995C-BFE90B70E540
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DocOcrResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DocOcrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DocOcrResponseBody) GoString() string {
	return s.String()
}

func (s *DocOcrResponseBody) SetCode(v string) *DocOcrResponseBody {
	s.Code = &v
	return s
}

func (s *DocOcrResponseBody) SetMessage(v string) *DocOcrResponseBody {
	s.Message = &v
	return s
}

func (s *DocOcrResponseBody) SetRequestId(v string) *DocOcrResponseBody {
	s.RequestId = &v
	return s
}

func (s *DocOcrResponseBody) SetResult(v *DocOcrResponseBodyResult) *DocOcrResponseBody {
	s.Result = v
	return s
}

type DocOcrResponseBodyResult struct {
	ExtIdInfo *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s DocOcrResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DocOcrResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DocOcrResponseBodyResult) SetExtIdInfo(v string) *DocOcrResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *DocOcrResponseBodyResult) SetPassed(v string) *DocOcrResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *DocOcrResponseBodyResult) SetSubCode(v string) *DocOcrResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *DocOcrResponseBodyResult) SetTransactionId(v string) *DocOcrResponseBodyResult {
	s.TransactionId = &v
	return s
}

type DocOcrResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocOcrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocOcrResponse) String() string {
	return tea.Prettify(s)
}

func (s DocOcrResponse) GoString() string {
	return s.String()
}

func (s *DocOcrResponse) SetHeaders(v map[string]*string) *DocOcrResponse {
	s.Headers = v
	return s
}

func (s *DocOcrResponse) SetStatusCode(v int32) *DocOcrResponse {
	s.StatusCode = &v
	return s
}

func (s *DocOcrResponse) SetBody(v *DocOcrResponseBody) *DocOcrResponse {
	s.Body = v
	return s
}

type EkycVerifyRequest struct {
	// example:
	//
	// T
	Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	// example:
	//
	// F
	Crop    *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// example:
	//
	// 410***************
	DocNo *string `json:"DocNo,omitempty" xml:"DocNo,omitempty"`
	// example:
	//
	// 00000001
	DocType           *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// example:
	//
	// https://digital-face-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	FacePictureUrl     *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// example:
	//
	// https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	IdThreshold     *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 123456
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// eKYC_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s EkycVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s EkycVerifyRequest) GoString() string {
	return s.String()
}

func (s *EkycVerifyRequest) SetAuthorize(v string) *EkycVerifyRequest {
	s.Authorize = &v
	return s
}

func (s *EkycVerifyRequest) SetCrop(v string) *EkycVerifyRequest {
	s.Crop = &v
	return s
}

func (s *EkycVerifyRequest) SetDocName(v string) *EkycVerifyRequest {
	s.DocName = &v
	return s
}

func (s *EkycVerifyRequest) SetDocNo(v string) *EkycVerifyRequest {
	s.DocNo = &v
	return s
}

func (s *EkycVerifyRequest) SetDocType(v string) *EkycVerifyRequest {
	s.DocType = &v
	return s
}

func (s *EkycVerifyRequest) SetFacePictureBase64(v string) *EkycVerifyRequest {
	s.FacePictureBase64 = &v
	return s
}

func (s *EkycVerifyRequest) SetFacePictureUrl(v string) *EkycVerifyRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *EkycVerifyRequest) SetIdOcrPictureBase64(v string) *EkycVerifyRequest {
	s.IdOcrPictureBase64 = &v
	return s
}

func (s *EkycVerifyRequest) SetIdOcrPictureUrl(v string) *EkycVerifyRequest {
	s.IdOcrPictureUrl = &v
	return s
}

func (s *EkycVerifyRequest) SetIdThreshold(v string) *EkycVerifyRequest {
	s.IdThreshold = &v
	return s
}

func (s *EkycVerifyRequest) SetMerchantBizId(v string) *EkycVerifyRequest {
	s.MerchantBizId = &v
	return s
}

func (s *EkycVerifyRequest) SetMerchantUserId(v string) *EkycVerifyRequest {
	s.MerchantUserId = &v
	return s
}

func (s *EkycVerifyRequest) SetProductCode(v string) *EkycVerifyRequest {
	s.ProductCode = &v
	return s
}

type EkycVerifyResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB356FE-BB6A-5DCC-B4C5-E8051787EBA1
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *EkycVerifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EkycVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EkycVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *EkycVerifyResponseBody) SetCode(v string) *EkycVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *EkycVerifyResponseBody) SetMessage(v string) *EkycVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *EkycVerifyResponseBody) SetRequestId(v string) *EkycVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *EkycVerifyResponseBody) SetResult(v *EkycVerifyResponseBodyResult) *EkycVerifyResponseBody {
	s.Result = v
	return s
}

type EkycVerifyResponseBodyResult struct {
	// example:
	//
	// {
	//
	// "faceAttack": "N",
	//
	// "faceComparisonScore": 52.57,
	//
	// "facePassed": "N",
	//
	// "authorityComparisonScore": 80.39
	//
	// }
	ExtFaceInfo *string `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty"`
	ExtIdInfo   *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 205
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// example:
	//
	// 4ab0b***cbde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s EkycVerifyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EkycVerifyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EkycVerifyResponseBodyResult) SetExtFaceInfo(v string) *EkycVerifyResponseBodyResult {
	s.ExtFaceInfo = &v
	return s
}

func (s *EkycVerifyResponseBodyResult) SetExtIdInfo(v string) *EkycVerifyResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *EkycVerifyResponseBodyResult) SetPassed(v string) *EkycVerifyResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *EkycVerifyResponseBodyResult) SetSubCode(v string) *EkycVerifyResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *EkycVerifyResponseBodyResult) SetTransactionId(v string) *EkycVerifyResponseBodyResult {
	s.TransactionId = &v
	return s
}

type EkycVerifyResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EkycVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EkycVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s EkycVerifyResponse) GoString() string {
	return s.String()
}

func (s *EkycVerifyResponse) SetHeaders(v map[string]*string) *EkycVerifyResponse {
	s.Headers = v
	return s
}

func (s *EkycVerifyResponse) SetStatusCode(v int32) *EkycVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *EkycVerifyResponse) SetBody(v *EkycVerifyResponseBody) *EkycVerifyResponse {
	s.Body = v
	return s
}

type FaceCompareRequest struct {
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	MerchantBizId     *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	SourceFacePicture *string `json:"SourceFacePicture,omitempty" xml:"SourceFacePicture,omitempty"`
	// example:
	//
	// https://***face1.jpeg
	SourceFacePictureUrl *string `json:"SourceFacePictureUrl,omitempty" xml:"SourceFacePictureUrl,omitempty"`
	TargetFacePicture    *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
	// example:
	//
	// https://***face2.jpeg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
}

func (s FaceCompareRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareRequest) GoString() string {
	return s.String()
}

func (s *FaceCompareRequest) SetMerchantBizId(v string) *FaceCompareRequest {
	s.MerchantBizId = &v
	return s
}

func (s *FaceCompareRequest) SetSourceFacePicture(v string) *FaceCompareRequest {
	s.SourceFacePicture = &v
	return s
}

func (s *FaceCompareRequest) SetSourceFacePictureUrl(v string) *FaceCompareRequest {
	s.SourceFacePictureUrl = &v
	return s
}

func (s *FaceCompareRequest) SetTargetFacePicture(v string) *FaceCompareRequest {
	s.TargetFacePicture = &v
	return s
}

func (s *FaceCompareRequest) SetTargetFacePictureUrl(v string) *FaceCompareRequest {
	s.TargetFacePictureUrl = &v
	return s
}

type FaceCompareResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB356FE-BB6A-5DCC-B4C5-E8051787EBA1
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *FaceCompareResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s FaceCompareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareResponseBody) GoString() string {
	return s.String()
}

func (s *FaceCompareResponseBody) SetCode(v string) *FaceCompareResponseBody {
	s.Code = &v
	return s
}

func (s *FaceCompareResponseBody) SetMessage(v string) *FaceCompareResponseBody {
	s.Message = &v
	return s
}

func (s *FaceCompareResponseBody) SetRequestId(v string) *FaceCompareResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceCompareResponseBody) SetResult(v *FaceCompareResponseBodyResult) *FaceCompareResponseBody {
	s.Result = v
	return s
}

type FaceCompareResponseBodyResult struct {
	// example:
	//
	// 98
	FaceComparisonScore *float64 `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s FaceCompareResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FaceCompareResponseBodyResult) SetFaceComparisonScore(v float64) *FaceCompareResponseBodyResult {
	s.FaceComparisonScore = &v
	return s
}

func (s *FaceCompareResponseBodyResult) SetPassed(v string) *FaceCompareResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *FaceCompareResponseBodyResult) SetTransactionId(v string) *FaceCompareResponseBodyResult {
	s.TransactionId = &v
	return s
}

type FaceCompareResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceCompareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceCompareResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareResponse) GoString() string {
	return s.String()
}

func (s *FaceCompareResponse) SetHeaders(v map[string]*string) *FaceCompareResponse {
	s.Headers = v
	return s
}

func (s *FaceCompareResponse) SetStatusCode(v int32) *FaceCompareResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceCompareResponse) SetBody(v *FaceCompareResponseBody) *FaceCompareResponse {
	s.Body = v
	return s
}

type FaceLivenessRequest struct {
	// example:
	//
	// T
	Crop              *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// example:
	//
	// https://digital-face-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	// example:
	//
	// T
	FaceQuality *string `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// T
	Occlusion *string `json:"Occlusion,omitempty" xml:"Occlusion,omitempty"`
	// example:
	//
	// FACE_LIVENESS_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s FaceLivenessRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceLivenessRequest) GoString() string {
	return s.String()
}

func (s *FaceLivenessRequest) SetCrop(v string) *FaceLivenessRequest {
	s.Crop = &v
	return s
}

func (s *FaceLivenessRequest) SetFacePictureBase64(v string) *FaceLivenessRequest {
	s.FacePictureBase64 = &v
	return s
}

func (s *FaceLivenessRequest) SetFacePictureUrl(v string) *FaceLivenessRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *FaceLivenessRequest) SetFaceQuality(v string) *FaceLivenessRequest {
	s.FaceQuality = &v
	return s
}

func (s *FaceLivenessRequest) SetMerchantBizId(v string) *FaceLivenessRequest {
	s.MerchantBizId = &v
	return s
}

func (s *FaceLivenessRequest) SetMerchantUserId(v string) *FaceLivenessRequest {
	s.MerchantUserId = &v
	return s
}

func (s *FaceLivenessRequest) SetOcclusion(v string) *FaceLivenessRequest {
	s.Occlusion = &v
	return s
}

func (s *FaceLivenessRequest) SetProductCode(v string) *FaceLivenessRequest {
	s.ProductCode = &v
	return s
}

type FaceLivenessResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 42EA58CA-5DF4-55D5-82C4-5E7A40DA62BA
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *FaceLivenessResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s FaceLivenessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceLivenessResponseBody) GoString() string {
	return s.String()
}

func (s *FaceLivenessResponseBody) SetCode(v string) *FaceLivenessResponseBody {
	s.Code = &v
	return s
}

func (s *FaceLivenessResponseBody) SetMessage(v string) *FaceLivenessResponseBody {
	s.Message = &v
	return s
}

func (s *FaceLivenessResponseBody) SetRequestId(v string) *FaceLivenessResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceLivenessResponseBody) SetResult(v *FaceLivenessResponseBodyResult) *FaceLivenessResponseBody {
	s.Result = v
	return s
}

type FaceLivenessResponseBodyResult struct {
	ExtFaceInfo *FaceLivenessResponseBodyResultExtFaceInfo `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty" type:"Struct"`
	// example:
	//
	// N
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 205
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s FaceLivenessResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s FaceLivenessResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FaceLivenessResponseBodyResult) SetExtFaceInfo(v *FaceLivenessResponseBodyResultExtFaceInfo) *FaceLivenessResponseBodyResult {
	s.ExtFaceInfo = v
	return s
}

func (s *FaceLivenessResponseBodyResult) SetPassed(v string) *FaceLivenessResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *FaceLivenessResponseBodyResult) SetSubCode(v string) *FaceLivenessResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *FaceLivenessResponseBodyResult) SetTransactionId(v string) *FaceLivenessResponseBodyResult {
	s.TransactionId = &v
	return s
}

type FaceLivenessResponseBodyResultExtFaceInfo struct {
	// example:
	//
	// Y
	FaceAttack *string `json:"FaceAttack,omitempty" xml:"FaceAttack,omitempty"`
	// example:
	//
	// 87.19
	FaceQualityScore *float64 `json:"FaceQualityScore,omitempty" xml:"FaceQualityScore,omitempty"`
	// example:
	//
	// Y
	OcclusionResult *string `json:"OcclusionResult,omitempty" xml:"OcclusionResult,omitempty"`
}

func (s FaceLivenessResponseBodyResultExtFaceInfo) String() string {
	return tea.Prettify(s)
}

func (s FaceLivenessResponseBodyResultExtFaceInfo) GoString() string {
	return s.String()
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetFaceAttack(v string) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.FaceAttack = &v
	return s
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetFaceQualityScore(v float64) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.FaceQualityScore = &v
	return s
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetOcclusionResult(v string) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.OcclusionResult = &v
	return s
}

type FaceLivenessResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceLivenessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceLivenessResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceLivenessResponse) GoString() string {
	return s.String()
}

func (s *FaceLivenessResponse) SetHeaders(v map[string]*string) *FaceLivenessResponse {
	s.Headers = v
	return s
}

func (s *FaceLivenessResponse) SetStatusCode(v int32) *FaceLivenessResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceLivenessResponse) SetBody(v *FaceLivenessResponseBody) *FaceLivenessResponse {
	s.Body = v
	return s
}

type FraudResultCallBackRequest struct {
	// example:
	//
	// shs2b27333914876c01de4cb22f5841f
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	ExtParams *string `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	// example:
	//
	// PASS
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// production
	VerifyDeployEnv *string `json:"VerifyDeployEnv,omitempty" xml:"VerifyDeployEnv,omitempty"`
}

func (s FraudResultCallBackRequest) String() string {
	return tea.Prettify(s)
}

func (s FraudResultCallBackRequest) GoString() string {
	return s.String()
}

func (s *FraudResultCallBackRequest) SetCertifyId(v string) *FraudResultCallBackRequest {
	s.CertifyId = &v
	return s
}

func (s *FraudResultCallBackRequest) SetExtParams(v string) *FraudResultCallBackRequest {
	s.ExtParams = &v
	return s
}

func (s *FraudResultCallBackRequest) SetResultCode(v string) *FraudResultCallBackRequest {
	s.ResultCode = &v
	return s
}

func (s *FraudResultCallBackRequest) SetVerifyDeployEnv(v string) *FraudResultCallBackRequest {
	s.VerifyDeployEnv = &v
	return s
}

type FraudResultCallBackResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FraudResultCallBackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FraudResultCallBackResponseBody) GoString() string {
	return s.String()
}

func (s *FraudResultCallBackResponseBody) SetCode(v string) *FraudResultCallBackResponseBody {
	s.Code = &v
	return s
}

func (s *FraudResultCallBackResponseBody) SetMessage(v string) *FraudResultCallBackResponseBody {
	s.Message = &v
	return s
}

func (s *FraudResultCallBackResponseBody) SetRequestId(v string) *FraudResultCallBackResponseBody {
	s.RequestId = &v
	return s
}

func (s *FraudResultCallBackResponseBody) SetSuccess(v bool) *FraudResultCallBackResponseBody {
	s.Success = &v
	return s
}

type FraudResultCallBackResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FraudResultCallBackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FraudResultCallBackResponse) String() string {
	return tea.Prettify(s)
}

func (s FraudResultCallBackResponse) GoString() string {
	return s.String()
}

func (s *FraudResultCallBackResponse) SetHeaders(v map[string]*string) *FraudResultCallBackResponse {
	s.Headers = v
	return s
}

func (s *FraudResultCallBackResponse) SetStatusCode(v int32) *FraudResultCallBackResponse {
	s.StatusCode = &v
	return s
}

func (s *FraudResultCallBackResponse) SetBody(v *FraudResultCallBackResponseBody) *FraudResultCallBackResponse {
	s.Body = v
	return s
}

type Id2MetaVerifyIntlRequest struct {
	// example:
	//
	// 429001********8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// example:
	//
	// ID_2META
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Id2MetaVerifyIntlRequest) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaVerifyIntlRequest) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyIntlRequest) SetIdentifyNum(v string) *Id2MetaVerifyIntlRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Id2MetaVerifyIntlRequest) SetParamType(v string) *Id2MetaVerifyIntlRequest {
	s.ParamType = &v
	return s
}

func (s *Id2MetaVerifyIntlRequest) SetProductCode(v string) *Id2MetaVerifyIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *Id2MetaVerifyIntlRequest) SetUserName(v string) *Id2MetaVerifyIntlRequest {
	s.UserName = &v
	return s
}

type Id2MetaVerifyIntlResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EFA11401-C961-5E89-A2D3-BF9883E5CC3D
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *Id2MetaVerifyIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s Id2MetaVerifyIntlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaVerifyIntlResponseBody) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyIntlResponseBody) SetCode(v string) *Id2MetaVerifyIntlResponseBody {
	s.Code = &v
	return s
}

func (s *Id2MetaVerifyIntlResponseBody) SetMessage(v string) *Id2MetaVerifyIntlResponseBody {
	s.Message = &v
	return s
}

func (s *Id2MetaVerifyIntlResponseBody) SetRequestId(v string) *Id2MetaVerifyIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *Id2MetaVerifyIntlResponseBody) SetResult(v *Id2MetaVerifyIntlResponseBodyResult) *Id2MetaVerifyIntlResponseBody {
	s.Result = v
	return s
}

type Id2MetaVerifyIntlResponseBodyResult struct {
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
}

func (s Id2MetaVerifyIntlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaVerifyIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyIntlResponseBodyResult) SetBizCode(v string) *Id2MetaVerifyIntlResponseBodyResult {
	s.BizCode = &v
	return s
}

type Id2MetaVerifyIntlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Id2MetaVerifyIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Id2MetaVerifyIntlResponse) String() string {
	return tea.Prettify(s)
}

func (s Id2MetaVerifyIntlResponse) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyIntlResponse) SetHeaders(v map[string]*string) *Id2MetaVerifyIntlResponse {
	s.Headers = v
	return s
}

func (s *Id2MetaVerifyIntlResponse) SetStatusCode(v int32) *Id2MetaVerifyIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *Id2MetaVerifyIntlResponse) SetBody(v *Id2MetaVerifyIntlResponseBody) *Id2MetaVerifyIntlResponse {
	s.Body = v
	return s
}

type InitializeRequest struct {
	Authorize     *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	CallbackToken *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	CallbackUrl   *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// example:
	//
	// *
	Crop        *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DocScanMode *string `json:"DocScanMode,omitempty" xml:"DocScanMode,omitempty"`
	// example:
	//
	// 01000000
	DocType           *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	DocVideo          *string `json:"DocVideo,omitempty" xml:"DocVideo,omitempty"`
	ExperienceCode    *string `json:"ExperienceCode,omitempty" xml:"ExperienceCode,omitempty"`
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// example:
	//
	// ***
	FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	// example:
	//
	// *
	IdFaceQuality *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	// example:
	//
	// Y
	IdSpoof        *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
	IdThreshold    *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	LanguageConfig *string `json:"LanguageConfig,omitempty" xml:"LanguageConfig,omitempty"`
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 1221****6543
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// {\\"bioMetaInfo\\":\\"4.1.0:2916352,0\\",\\"deviceType\\":\\"web\\",\\"ua\\":\\"Mozilla/5.0 (Macintosh
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	Model    *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// OCR。
	//
	// example:
	//
	// *
	Ocr *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	// example:
	//
	// eKYC
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductFlow *string `json:"ProductFlow,omitempty" xml:"ProductFlow,omitempty"`
	// example:
	//
	// http*****
	ReturnUrl *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	// example:
	//
	// PAY**
	SceneCode     *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	ShowAlbumIcon *string `json:"ShowAlbumIcon,omitempty" xml:"ShowAlbumIcon,omitempty"`
	ShowOcrResult *string `json:"ShowOcrResult,omitempty" xml:"ShowOcrResult,omitempty"`
	StyleConfig   *string `json:"StyleConfig,omitempty" xml:"StyleConfig,omitempty"`
}

func (s InitializeRequest) String() string {
	return tea.Prettify(s)
}

func (s InitializeRequest) GoString() string {
	return s.String()
}

func (s *InitializeRequest) SetAuthorize(v string) *InitializeRequest {
	s.Authorize = &v
	return s
}

func (s *InitializeRequest) SetCallbackToken(v string) *InitializeRequest {
	s.CallbackToken = &v
	return s
}

func (s *InitializeRequest) SetCallbackUrl(v string) *InitializeRequest {
	s.CallbackUrl = &v
	return s
}

func (s *InitializeRequest) SetCrop(v string) *InitializeRequest {
	s.Crop = &v
	return s
}

func (s *InitializeRequest) SetDocScanMode(v string) *InitializeRequest {
	s.DocScanMode = &v
	return s
}

func (s *InitializeRequest) SetDocType(v string) *InitializeRequest {
	s.DocType = &v
	return s
}

func (s *InitializeRequest) SetDocVideo(v string) *InitializeRequest {
	s.DocVideo = &v
	return s
}

func (s *InitializeRequest) SetExperienceCode(v string) *InitializeRequest {
	s.ExperienceCode = &v
	return s
}

func (s *InitializeRequest) SetFacePictureBase64(v string) *InitializeRequest {
	s.FacePictureBase64 = &v
	return s
}

func (s *InitializeRequest) SetFacePictureUrl(v string) *InitializeRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *InitializeRequest) SetIdFaceQuality(v string) *InitializeRequest {
	s.IdFaceQuality = &v
	return s
}

func (s *InitializeRequest) SetIdSpoof(v string) *InitializeRequest {
	s.IdSpoof = &v
	return s
}

func (s *InitializeRequest) SetIdThreshold(v string) *InitializeRequest {
	s.IdThreshold = &v
	return s
}

func (s *InitializeRequest) SetLanguageConfig(v string) *InitializeRequest {
	s.LanguageConfig = &v
	return s
}

func (s *InitializeRequest) SetMerchantBizId(v string) *InitializeRequest {
	s.MerchantBizId = &v
	return s
}

func (s *InitializeRequest) SetMerchantUserId(v string) *InitializeRequest {
	s.MerchantUserId = &v
	return s
}

func (s *InitializeRequest) SetMetaInfo(v string) *InitializeRequest {
	s.MetaInfo = &v
	return s
}

func (s *InitializeRequest) SetModel(v string) *InitializeRequest {
	s.Model = &v
	return s
}

func (s *InitializeRequest) SetOcr(v string) *InitializeRequest {
	s.Ocr = &v
	return s
}

func (s *InitializeRequest) SetProductCode(v string) *InitializeRequest {
	s.ProductCode = &v
	return s
}

func (s *InitializeRequest) SetProductFlow(v string) *InitializeRequest {
	s.ProductFlow = &v
	return s
}

func (s *InitializeRequest) SetReturnUrl(v string) *InitializeRequest {
	s.ReturnUrl = &v
	return s
}

func (s *InitializeRequest) SetSceneCode(v string) *InitializeRequest {
	s.SceneCode = &v
	return s
}

func (s *InitializeRequest) SetSecurityLevel(v string) *InitializeRequest {
	s.SecurityLevel = &v
	return s
}

func (s *InitializeRequest) SetShowAlbumIcon(v string) *InitializeRequest {
	s.ShowAlbumIcon = &v
	return s
}

func (s *InitializeRequest) SetShowOcrResult(v string) *InitializeRequest {
	s.ShowOcrResult = &v
	return s
}

func (s *InitializeRequest) SetStyleConfig(v string) *InitializeRequest {
	s.StyleConfig = &v
	return s
}

type InitializeResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *InitializeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s InitializeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitializeResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeResponseBody) SetCode(v string) *InitializeResponseBody {
	s.Code = &v
	return s
}

func (s *InitializeResponseBody) SetMessage(v string) *InitializeResponseBody {
	s.Message = &v
	return s
}

func (s *InitializeResponseBody) SetRequestId(v string) *InitializeResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeResponseBody) SetResult(v *InitializeResponseBodyResult) *InitializeResponseBody {
	s.Result = v
	return s
}

type InitializeResponseBodyResult struct {
	// example:
	//
	// ***
	ClientCfg *string `json:"ClientCfg,omitempty" xml:"ClientCfg,omitempty"`
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
	// example:
	//
	// http****
	TransactionUrl *string `json:"TransactionUrl,omitempty" xml:"TransactionUrl,omitempty"`
}

func (s InitializeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s InitializeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InitializeResponseBodyResult) SetClientCfg(v string) *InitializeResponseBodyResult {
	s.ClientCfg = &v
	return s
}

func (s *InitializeResponseBodyResult) SetTransactionId(v string) *InitializeResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *InitializeResponseBodyResult) SetTransactionUrl(v string) *InitializeResponseBodyResult {
	s.TransactionUrl = &v
	return s
}

type InitializeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeResponse) String() string {
	return tea.Prettify(s)
}

func (s InitializeResponse) GoString() string {
	return s.String()
}

func (s *InitializeResponse) SetHeaders(v map[string]*string) *InitializeResponse {
	s.Headers = v
	return s
}

func (s *InitializeResponse) SetStatusCode(v int32) *InitializeResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeResponse) SetBody(v *InitializeResponseBody) *InitializeResponse {
	s.Body = v
	return s
}

type Mobile3MetaVerifyIntlRequest struct {
	// example:
	//
	// 429001********8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// example:
	//
	// 186****1234
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// example:
	//
	// MOBILE_3META
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Mobile3MetaVerifyIntlRequest) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaVerifyIntlRequest) GoString() string {
	return s.String()
}

func (s *Mobile3MetaVerifyIntlRequest) SetIdentifyNum(v string) *Mobile3MetaVerifyIntlRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Mobile3MetaVerifyIntlRequest) SetMobile(v string) *Mobile3MetaVerifyIntlRequest {
	s.Mobile = &v
	return s
}

func (s *Mobile3MetaVerifyIntlRequest) SetParamType(v string) *Mobile3MetaVerifyIntlRequest {
	s.ParamType = &v
	return s
}

func (s *Mobile3MetaVerifyIntlRequest) SetProductCode(v string) *Mobile3MetaVerifyIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *Mobile3MetaVerifyIntlRequest) SetUserName(v string) *Mobile3MetaVerifyIntlRequest {
	s.UserName = &v
	return s
}

type Mobile3MetaVerifyIntlResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D241532C-4EE9-5A2A-A5A5-C1FD98CE2EDD
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *Mobile3MetaVerifyIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s Mobile3MetaVerifyIntlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaVerifyIntlResponseBody) GoString() string {
	return s.String()
}

func (s *Mobile3MetaVerifyIntlResponseBody) SetCode(v string) *Mobile3MetaVerifyIntlResponseBody {
	s.Code = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBody) SetMessage(v string) *Mobile3MetaVerifyIntlResponseBody {
	s.Message = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBody) SetRequestId(v string) *Mobile3MetaVerifyIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBody) SetResult(v *Mobile3MetaVerifyIntlResponseBodyResult) *Mobile3MetaVerifyIntlResponseBody {
	s.Result = v
	return s
}

type Mobile3MetaVerifyIntlResponseBodyResult struct {
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s Mobile3MetaVerifyIntlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaVerifyIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *Mobile3MetaVerifyIntlResponseBodyResult) SetBizCode(v string) *Mobile3MetaVerifyIntlResponseBodyResult {
	s.BizCode = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBodyResult) SetIspName(v string) *Mobile3MetaVerifyIntlResponseBodyResult {
	s.IspName = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponseBodyResult) SetSubCode(v string) *Mobile3MetaVerifyIntlResponseBodyResult {
	s.SubCode = &v
	return s
}

type Mobile3MetaVerifyIntlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Mobile3MetaVerifyIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Mobile3MetaVerifyIntlResponse) String() string {
	return tea.Prettify(s)
}

func (s Mobile3MetaVerifyIntlResponse) GoString() string {
	return s.String()
}

func (s *Mobile3MetaVerifyIntlResponse) SetHeaders(v map[string]*string) *Mobile3MetaVerifyIntlResponse {
	s.Headers = v
	return s
}

func (s *Mobile3MetaVerifyIntlResponse) SetStatusCode(v int32) *Mobile3MetaVerifyIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponse) SetBody(v *Mobile3MetaVerifyIntlResponseBody) *Mobile3MetaVerifyIntlResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudauth-intl"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CardOcr is deprecated, please use Cloudauth-intl::2022-08-09::DocOcr instead.
//
// Summary:
//
// 证件OCR识别纯服务端接口
//
// @param request - CardOcrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CardOcrResponse
// Deprecated
func (client *Client) CardOcrWithOptions(request *CardOcrRequest, runtime *util.RuntimeOptions) (_result *CardOcrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		query["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.IdFaceQuality)) {
		query["IdFaceQuality"] = request.IdFaceQuality
	}

	if !tea.BoolValue(util.IsUnset(request.IdOcrPictureUrl)) {
		query["IdOcrPictureUrl"] = request.IdOcrPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Ocr)) {
		query["Ocr"] = request.Ocr
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.Spoof)) {
		query["Spoof"] = request.Spoof
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdOcrPictureBase64)) {
		body["IdOcrPictureBase64"] = request.IdOcrPictureBase64
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CardOcr"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CardOcrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CardOcr is deprecated, please use Cloudauth-intl::2022-08-09::DocOcr instead.
//
// Summary:
//
// 证件OCR识别纯服务端接口
//
// @param request - CardOcrRequest
//
// @return CardOcrResponse
// Deprecated
func (client *Client) CardOcr(request *CardOcrRequest) (_result *CardOcrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CardOcrResponse{}
	_body, _err := client.CardOcrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 结果查询
//
// @param request - CheckResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckResultResponse
func (client *Client) CheckResultWithOptions(request *CheckResultRequest, runtime *util.RuntimeOptions) (_result *CheckResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtraImageControlList)) {
		query["ExtraImageControlList"] = request.ExtraImageControlList
	}

	if !tea.BoolValue(util.IsUnset(request.IsReturnImage)) {
		query["IsReturnImage"] = request.IsReturnImage
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnFiveCategorySpoofResult)) {
		query["ReturnFiveCategorySpoofResult"] = request.ReturnFiveCategorySpoofResult
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionId)) {
		query["TransactionId"] = request.TransactionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckResult"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 结果查询
//
// @param request - CheckResultRequest
//
// @return CheckResultResponse
func (client *Client) CheckResult(request *CheckResultRequest) (_result *CheckResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckResultResponse{}
	_body, _err := client.CheckResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除用户认证记录结果
//
// @param request - DeleteVerifyResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVerifyResultResponse
func (client *Client) DeleteVerifyResultWithOptions(request *DeleteVerifyResultRequest, runtime *util.RuntimeOptions) (_result *DeleteVerifyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeleteAfterQuery)) {
		query["DeleteAfterQuery"] = request.DeleteAfterQuery
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteType)) {
		query["DeleteType"] = request.DeleteType
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionId)) {
		query["TransactionId"] = request.TransactionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVerifyResult"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVerifyResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除用户认证记录结果
//
// @param request - DeleteVerifyResultRequest
//
// @return DeleteVerifyResultResponse
func (client *Client) DeleteVerifyResult(request *DeleteVerifyResultRequest) (_result *DeleteVerifyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVerifyResultResponse{}
	_body, _err := client.DeleteVerifyResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询地址标签接口
//
// @param request - DescribeAddressLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAddressLabelsResponse
func (client *Client) DescribeAddressLabelsWithOptions(request *DescribeAddressLabelsRequest, runtime *util.RuntimeOptions) (_result *DescribeAddressLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.Coin)) {
		query["Coin"] = request.Coin
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAddressLabels"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAddressLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询地址标签接口
//
// @param request - DescribeAddressLabelsRequest
//
// @return DescribeAddressLabelsResponse
func (client *Client) DescribeAddressLabels(request *DescribeAddressLabelsRequest) (_result *DescribeAddressLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAddressLabelsResponse{}
	_body, _err := client.DescribeAddressLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询地址详情接口
//
// @param request - DescribeAddressOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAddressOverviewResponse
func (client *Client) DescribeAddressOverviewWithOptions(request *DescribeAddressOverviewRequest, runtime *util.RuntimeOptions) (_result *DescribeAddressOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.Coin)) {
		query["Coin"] = request.Coin
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAddressOverview"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAddressOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询地址详情接口
//
// @param request - DescribeAddressOverviewRequest
//
// @return DescribeAddressOverviewResponse
func (client *Client) DescribeAddressOverview(request *DescribeAddressOverviewRequest) (_result *DescribeAddressOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAddressOverviewResponse{}
	_body, _err := client.DescribeAddressOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 恶意地址查询
//
// @param request - DescribeMaliciousAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMaliciousAddressResponse
func (client *Client) DescribeMaliciousAddressWithOptions(request *DescribeMaliciousAddressRequest, runtime *util.RuntimeOptions) (_result *DescribeMaliciousAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Coin)) {
		query["Coin"] = request.Coin
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["End"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["Start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMaliciousAddress"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMaliciousAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 恶意地址查询
//
// @param request - DescribeMaliciousAddressRequest
//
// @return DescribeMaliciousAddressResponse
func (client *Client) DescribeMaliciousAddress(request *DescribeMaliciousAddressRequest) (_result *DescribeMaliciousAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMaliciousAddressResponse{}
	_body, _err := client.DescribeMaliciousAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询地址风险接口
//
// @param request - DescribeRiskScoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRiskScoreResponse
func (client *Client) DescribeRiskScoreWithOptions(request *DescribeRiskScoreRequest, runtime *util.RuntimeOptions) (_result *DescribeRiskScoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.Coin)) {
		query["Coin"] = request.Coin
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRiskScore"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRiskScoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询地址风险接口
//
// @param request - DescribeRiskScoreRequest
//
// @return DescribeRiskScoreResponse
func (client *Client) DescribeRiskScore(request *DescribeRiskScoreRequest) (_result *DescribeRiskScoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRiskScoreResponse{}
	_body, _err := client.DescribeRiskScoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询交易接口
//
// @param request - DescribeTransactionsListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTransactionsListResponse
func (client *Client) DescribeTransactionsListWithOptions(request *DescribeTransactionsListRequest, runtime *util.RuntimeOptions) (_result *DescribeTransactionsListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.Coin)) {
		query["Coin"] = request.Coin
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTransactionsList"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTransactionsListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询交易接口
//
// @param request - DescribeTransactionsListRequest
//
// @return DescribeTransactionsListResponse
func (client *Client) DescribeTransactionsList(request *DescribeTransactionsListRequest) (_result *DescribeTransactionsListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTransactionsListResponse{}
	_body, _err := client.DescribeTransactionsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the Information of address.
//
// @param request - DescribeWeb3AddressLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWeb3AddressLabelsResponse
func (client *Client) DescribeWeb3AddressLabelsWithOptions(request *DescribeWeb3AddressLabelsRequest, runtime *util.RuntimeOptions) (_result *DescribeWeb3AddressLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.ChainShortName)) {
		query["ChainShortName"] = request.ChainShortName
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWeb3AddressLabels"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWeb3AddressLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the Information of address.
//
// @param request - DescribeWeb3AddressLabelsRequest
//
// @return DescribeWeb3AddressLabelsResponse
func (client *Client) DescribeWeb3AddressLabels(request *DescribeWeb3AddressLabelsRequest) (_result *DescribeWeb3AddressLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWeb3AddressLabelsResponse{}
	_body, _err := client.DescribeWeb3AddressLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query risk score, risk detail list for a given address
//
// @param request - DescribeWeb3RiskScoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWeb3RiskScoreResponse
func (client *Client) DescribeWeb3RiskScoreWithOptions(request *DescribeWeb3RiskScoreRequest, runtime *util.RuntimeOptions) (_result *DescribeWeb3RiskScoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChainShortName)) {
		query["ChainShortName"] = request.ChainShortName
	}

	if !tea.BoolValue(util.IsUnset(request.Depth)) {
		query["Depth"] = request.Depth
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectId)) {
		query["ObjectId"] = request.ObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		query["ObjectType"] = request.ObjectType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWeb3RiskScore"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWeb3RiskScoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query risk score, risk detail list for a given address
//
// @param request - DescribeWeb3RiskScoreRequest
//
// @return DescribeWeb3RiskScoreResponse
func (client *Client) DescribeWeb3RiskScore(request *DescribeWeb3RiskScoreRequest) (_result *DescribeWeb3RiskScoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWeb3RiskScoreResponse{}
	_body, _err := client.DescribeWeb3RiskScoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the Information of transaction.
//
// @param request - DescribeWeb3TransactionLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWeb3TransactionLabelsResponse
func (client *Client) DescribeWeb3TransactionLabelsWithOptions(request *DescribeWeb3TransactionLabelsRequest, runtime *util.RuntimeOptions) (_result *DescribeWeb3TransactionLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChainShortName)) {
		query["ChainShortName"] = request.ChainShortName
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.Transaction)) {
		query["Transaction"] = request.Transaction
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWeb3TransactionLabels"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWeb3TransactionLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the Information of transaction.
//
// @param request - DescribeWeb3TransactionLabelsRequest
//
// @return DescribeWeb3TransactionLabelsResponse
func (client *Client) DescribeWeb3TransactionLabels(request *DescribeWeb3TransactionLabelsRequest) (_result *DescribeWeb3TransactionLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWeb3TransactionLabelsResponse{}
	_body, _err := client.DescribeWeb3TransactionLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 卡证ocr纯服务端
//
// @param request - DocOcrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DocOcrResponse
func (client *Client) DocOcrWithOptions(request *DocOcrRequest, runtime *util.RuntimeOptions) (_result *DocOcrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		query["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.IdFaceQuality)) {
		query["IdFaceQuality"] = request.IdFaceQuality
	}

	if !tea.BoolValue(util.IsUnset(request.IdOcrPictureUrl)) {
		query["IdOcrPictureUrl"] = request.IdOcrPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdThreshold)) {
		query["IdThreshold"] = request.IdThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Ocr)) {
		query["Ocr"] = request.Ocr
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.Spoof)) {
		query["Spoof"] = request.Spoof
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdOcrPictureBase64)) {
		body["IdOcrPictureBase64"] = request.IdOcrPictureBase64
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DocOcr"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DocOcrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 卡证ocr纯服务端
//
// @param request - DocOcrRequest
//
// @return DocOcrResponse
func (client *Client) DocOcr(request *DocOcrRequest) (_result *DocOcrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DocOcrResponse{}
	_body, _err := client.DocOcrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// ekyc纯服务端接口
//
// @param request - EkycVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EkycVerifyResponse
func (client *Client) EkycVerifyWithOptions(request *EkycVerifyRequest, runtime *util.RuntimeOptions) (_result *EkycVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Authorize)) {
		query["Authorize"] = request.Authorize
	}

	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		query["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.DocName)) {
		query["DocName"] = request.DocName
	}

	if !tea.BoolValue(util.IsUnset(request.DocNo)) {
		query["DocNo"] = request.DocNo
	}

	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		query["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.FacePictureUrl)) {
		query["FacePictureUrl"] = request.FacePictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdOcrPictureUrl)) {
		query["IdOcrPictureUrl"] = request.IdOcrPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdThreshold)) {
		query["IdThreshold"] = request.IdThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FacePictureBase64)) {
		body["FacePictureBase64"] = request.FacePictureBase64
	}

	if !tea.BoolValue(util.IsUnset(request.IdOcrPictureBase64)) {
		body["IdOcrPictureBase64"] = request.IdOcrPictureBase64
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EkycVerify"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EkycVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ekyc纯服务端接口
//
// @param request - EkycVerifyRequest
//
// @return EkycVerifyResponse
func (client *Client) EkycVerify(request *EkycVerifyRequest) (_result *EkycVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EkycVerifyResponse{}
	_body, _err := client.EkycVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人脸比对
//
// @param request - FaceCompareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FaceCompareResponse
func (client *Client) FaceCompareWithOptions(request *FaceCompareRequest, runtime *util.RuntimeOptions) (_result *FaceCompareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceFacePictureUrl)) {
		query["SourceFacePictureUrl"] = request.SourceFacePictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFacePictureUrl)) {
		query["TargetFacePictureUrl"] = request.TargetFacePictureUrl
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceFacePicture)) {
		body["SourceFacePicture"] = request.SourceFacePicture
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFacePicture)) {
		body["TargetFacePicture"] = request.TargetFacePicture
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FaceCompare"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FaceCompareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人脸比对
//
// @param request - FaceCompareRequest
//
// @return FaceCompareResponse
func (client *Client) FaceCompare(request *FaceCompareRequest) (_result *FaceCompareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceCompareResponse{}
	_body, _err := client.FaceCompareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 静默活体API 纯服务端
//
// @param request - FaceLivenessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FaceLivenessResponse
func (client *Client) FaceLivenessWithOptions(request *FaceLivenessRequest, runtime *util.RuntimeOptions) (_result *FaceLivenessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		query["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.FacePictureUrl)) {
		query["FacePictureUrl"] = request.FacePictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FaceQuality)) {
		query["FaceQuality"] = request.FaceQuality
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Occlusion)) {
		query["Occlusion"] = request.Occlusion
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FacePictureBase64)) {
		body["FacePictureBase64"] = request.FacePictureBase64
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FaceLiveness"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FaceLivenessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 静默活体API 纯服务端
//
// @param request - FaceLivenessRequest
//
// @return FaceLivenessResponse
func (client *Client) FaceLiveness(request *FaceLivenessRequest) (_result *FaceLivenessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceLivenessResponse{}
	_body, _err := client.FaceLivenessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 防伪回调接口
//
// @param request - FraudResultCallBackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FraudResultCallBackResponse
func (client *Client) FraudResultCallBackWithOptions(request *FraudResultCallBackRequest, runtime *util.RuntimeOptions) (_result *FraudResultCallBackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertifyId)) {
		query["CertifyId"] = request.CertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtParams)) {
		query["ExtParams"] = request.ExtParams
	}

	if !tea.BoolValue(util.IsUnset(request.ResultCode)) {
		query["ResultCode"] = request.ResultCode
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyDeployEnv)) {
		query["VerifyDeployEnv"] = request.VerifyDeployEnv
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FraudResultCallBack"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FraudResultCallBackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 防伪回调接口
//
// @param request - FraudResultCallBackRequest
//
// @return FraudResultCallBackResponse
func (client *Client) FraudResultCallBack(request *FraudResultCallBackRequest) (_result *FraudResultCallBackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FraudResultCallBackResponse{}
	_body, _err := client.FraudResultCallBackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 身份二要素国际版接口
//
// @param request - Id2MetaVerifyIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id2MetaVerifyIntlResponse
func (client *Client) Id2MetaVerifyIntlWithOptions(request *Id2MetaVerifyIntlRequest, runtime *util.RuntimeOptions) (_result *Id2MetaVerifyIntlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdentifyNum)) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		query["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Id2MetaVerifyIntl"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &Id2MetaVerifyIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 身份二要素国际版接口
//
// @param request - Id2MetaVerifyIntlRequest
//
// @return Id2MetaVerifyIntlResponse
func (client *Client) Id2MetaVerifyIntl(request *Id2MetaVerifyIntlRequest) (_result *Id2MetaVerifyIntlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &Id2MetaVerifyIntlResponse{}
	_body, _err := client.Id2MetaVerifyIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 认证初始化
//
// @param request - InitializeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitializeResponse
func (client *Client) InitializeWithOptions(request *InitializeRequest, runtime *util.RuntimeOptions) (_result *InitializeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Authorize)) {
		query["Authorize"] = request.Authorize
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackToken)) {
		query["CallbackToken"] = request.CallbackToken
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		query["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.DocScanMode)) {
		query["DocScanMode"] = request.DocScanMode
	}

	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		query["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.DocVideo)) {
		query["DocVideo"] = request.DocVideo
	}

	if !tea.BoolValue(util.IsUnset(request.ExperienceCode)) {
		query["ExperienceCode"] = request.ExperienceCode
	}

	if !tea.BoolValue(util.IsUnset(request.FacePictureUrl)) {
		query["FacePictureUrl"] = request.FacePictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdFaceQuality)) {
		query["IdFaceQuality"] = request.IdFaceQuality
	}

	if !tea.BoolValue(util.IsUnset(request.IdSpoof)) {
		query["IdSpoof"] = request.IdSpoof
	}

	if !tea.BoolValue(util.IsUnset(request.IdThreshold)) {
		query["IdThreshold"] = request.IdThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageConfig)) {
		query["LanguageConfig"] = request.LanguageConfig
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.MetaInfo)) {
		query["MetaInfo"] = request.MetaInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		query["Model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.Ocr)) {
		query["Ocr"] = request.Ocr
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductFlow)) {
		query["ProductFlow"] = request.ProductFlow
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnUrl)) {
		query["ReturnUrl"] = request.ReturnUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityLevel)) {
		query["SecurityLevel"] = request.SecurityLevel
	}

	if !tea.BoolValue(util.IsUnset(request.ShowAlbumIcon)) {
		query["ShowAlbumIcon"] = request.ShowAlbumIcon
	}

	if !tea.BoolValue(util.IsUnset(request.ShowOcrResult)) {
		query["ShowOcrResult"] = request.ShowOcrResult
	}

	if !tea.BoolValue(util.IsUnset(request.StyleConfig)) {
		query["StyleConfig"] = request.StyleConfig
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FacePictureBase64)) {
		body["FacePictureBase64"] = request.FacePictureBase64
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Initialize"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitializeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 认证初始化
//
// @param request - InitializeRequest
//
// @return InitializeResponse
func (client *Client) Initialize(request *InitializeRequest) (_result *InitializeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitializeResponse{}
	_body, _err := client.InitializeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 手机号三要素国际版接口
//
// @param request - Mobile3MetaVerifyIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile3MetaVerifyIntlResponse
func (client *Client) Mobile3MetaVerifyIntlWithOptions(request *Mobile3MetaVerifyIntlRequest, runtime *util.RuntimeOptions) (_result *Mobile3MetaVerifyIntlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdentifyNum)) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		query["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Mobile3MetaVerifyIntl"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &Mobile3MetaVerifyIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 手机号三要素国际版接口
//
// @param request - Mobile3MetaVerifyIntlRequest
//
// @return Mobile3MetaVerifyIntlResponse
func (client *Client) Mobile3MetaVerifyIntl(request *Mobile3MetaVerifyIntlRequest) (_result *Mobile3MetaVerifyIntlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &Mobile3MetaVerifyIntlResponse{}
	_body, _err := client.Mobile3MetaVerifyIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
