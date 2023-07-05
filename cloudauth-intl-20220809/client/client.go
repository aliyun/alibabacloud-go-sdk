// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CardOcrRequest struct {
	DocType            *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	IdFaceQuality      *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	IdOcrPictureUrl    *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	MerchantBizId      *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	MerchantUserId     *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	Ocr                *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	ProductCode        *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Spoof              *string `json:"Spoof,omitempty" xml:"Spoof,omitempty"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	ExtCardInfo   *string `json:"ExtCardInfo,omitempty" xml:"ExtCardInfo,omitempty"`
	ExtIdInfo     *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	Passed        *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	SubCode       *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CardOcrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ExtraImageControlList         *string `json:"ExtraImageControlList,omitempty" xml:"ExtraImageControlList,omitempty"`
	IsReturnImage                 *string `json:"IsReturnImage,omitempty" xml:"IsReturnImage,omitempty"`
	MerchantBizId                 *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	ReturnFiveCategorySpoofResult *string `json:"ReturnFiveCategorySpoofResult,omitempty" xml:"ReturnFiveCategorySpoofResult,omitempty"`
	TransactionId                 *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	EkycResult   *string `json:"EkycResult,omitempty" xml:"EkycResult,omitempty"`
	ExtBasicInfo *string `json:"ExtBasicInfo,omitempty" xml:"ExtBasicInfo,omitempty"`
	ExtFaceInfo  *string `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty"`
	ExtIdInfo    *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	ExtRiskInfo  *string `json:"ExtRiskInfo,omitempty" xml:"ExtRiskInfo,omitempty"`
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeletePictureRequest struct {
	DeletePicAfterQuery *string `json:"DeletePicAfterQuery,omitempty" xml:"DeletePicAfterQuery,omitempty"`
	TransactionId       *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s DeletePictureRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePictureRequest) GoString() string {
	return s.String()
}

func (s *DeletePictureRequest) SetDeletePicAfterQuery(v string) *DeletePictureRequest {
	s.DeletePicAfterQuery = &v
	return s
}

func (s *DeletePictureRequest) SetTransactionId(v string) *DeletePictureRequest {
	s.TransactionId = &v
	return s
}

type DeletePictureResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeletePictureResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeletePictureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePictureResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePictureResponseBody) SetCode(v string) *DeletePictureResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePictureResponseBody) SetMessage(v string) *DeletePictureResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePictureResponseBody) SetRequestId(v string) *DeletePictureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePictureResponseBody) SetResult(v *DeletePictureResponseBodyResult) *DeletePictureResponseBody {
	s.Result = v
	return s
}

type DeletePictureResponseBodyResult struct {
	DeleteResult  *string `json:"DeleteResult,omitempty" xml:"DeleteResult,omitempty"`
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s DeletePictureResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeletePictureResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeletePictureResponseBodyResult) SetDeleteResult(v string) *DeletePictureResponseBodyResult {
	s.DeleteResult = &v
	return s
}

func (s *DeletePictureResponseBodyResult) SetTransactionId(v string) *DeletePictureResponseBodyResult {
	s.TransactionId = &v
	return s
}

type DeletePictureResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePictureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePictureResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePictureResponse) GoString() string {
	return s.String()
}

func (s *DeletePictureResponse) SetHeaders(v map[string]*string) *DeletePictureResponse {
	s.Headers = v
	return s
}

func (s *DeletePictureResponse) SetStatusCode(v int32) *DeletePictureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePictureResponse) SetBody(v *DeletePictureResponseBody) *DeletePictureResponse {
	s.Body = v
	return s
}

type DescribeAddressLabelsRequest struct {
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Coin          *string `json:"Coin,omitempty" xml:"Coin,omitempty"`
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
	Code    *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *DescribeAddressLabelsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAddressLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Coin          *string `json:"Coin,omitempty" xml:"Coin,omitempty"`
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
	Code    *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *DescribeAddressOverviewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Balance          *float32 `json:"Balance,omitempty" xml:"Balance,omitempty"`
	FirstSeen        *int64   `json:"FirstSeen,omitempty" xml:"FirstSeen,omitempty"`
	LastSeen         *int64   `json:"LastSeen,omitempty" xml:"LastSeen,omitempty"`
	ReceivedTxsCount *int32   `json:"ReceivedTxsCount,omitempty" xml:"ReceivedTxsCount,omitempty"`
	SpentTxsCount    *int32   `json:"SpentTxsCount,omitempty" xml:"SpentTxsCount,omitempty"`
	TotalReceived    *float32 `json:"TotalReceived,omitempty" xml:"TotalReceived,omitempty"`
	TotalSpent       *float32 `json:"TotalSpent,omitempty" xml:"TotalSpent,omitempty"`
	TxsCount         *int64   `json:"TxsCount,omitempty" xml:"TxsCount,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAddressOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Coin          *string `json:"Coin,omitempty" xml:"Coin,omitempty"`
	End           *string `json:"End,omitempty" xml:"End,omitempty"`
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	Start         *string `json:"Start,omitempty" xml:"Start,omitempty"`
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
	Code    *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    []*DescribeMaliciousAddressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	AddTime *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Coin    *string `json:"Coin,omitempty" xml:"Coin,omitempty"`
	Detail  *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Tag     *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMaliciousAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Coin          *string `json:"Coin,omitempty" xml:"Coin,omitempty"`
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
	Code    *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *DescribeRiskScoreResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	DetailList   []*string `json:"DetailList,omitempty" xml:"DetailList,omitempty" type:"Repeated"`
	HackingEvent *string   `json:"HackingEvent,omitempty" xml:"HackingEvent,omitempty"`
	RiskLevel    *string   `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Score        *int32    `json:"Score,omitempty" xml:"Score,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRiskScoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Address        *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Coin           *string `json:"Coin,omitempty" xml:"Coin,omitempty"`
	EndTimestamp   *int64  `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	MerchantBizId  *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	Page           *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	StartTimestamp *int64  `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Code    *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *DescribeTransactionsListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	In                 []*DescribeTransactionsListResponseBodyDataIn  `json:"In,omitempty" xml:"In,omitempty" type:"Repeated"`
	Out                []*DescribeTransactionsListResponseBodyDataOut `json:"Out,omitempty" xml:"Out,omitempty" type:"Repeated"`
	Page               *int64                                         `json:"Page,omitempty" xml:"Page,omitempty"`
	TotalPages         *int64                                         `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TransactionsOnPage *int64                                         `json:"TransactionsOnPage,omitempty" xml:"TransactionsOnPage,omitempty"`
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
	Address    *string   `json:"Address,omitempty" xml:"Address,omitempty"`
	Amount     *float32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	Label      *string   `json:"Label,omitempty" xml:"Label,omitempty"`
	TxHashList []*string `json:"TxHashList,omitempty" xml:"TxHashList,omitempty" type:"Repeated"`
	Type       *int32    `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Address    *string   `json:"Address,omitempty" xml:"Address,omitempty"`
	Amount     *float32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	Label      *string   `json:"Label,omitempty" xml:"Label,omitempty"`
	TxHashList []*string `json:"TxHashList,omitempty" xml:"TxHashList,omitempty" type:"Repeated"`
	Type       *int32    `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTransactionsListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DocOcrRequest struct {
	DocType            *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	IdFaceQuality      *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	IdOcrPictureUrl    *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	MerchantBizId      *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	MerchantUserId     *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	Ocr                *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	ProductCode        *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Spoof              *string `json:"Spoof,omitempty" xml:"Spoof,omitempty"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	ExtIdInfo     *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	Passed        *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	SubCode       *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
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
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DocOcrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type FaceCompareRequest struct {
	MerchantBizId        *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	SourceFacePicture    *string `json:"SourceFacePicture,omitempty" xml:"SourceFacePicture,omitempty"`
	SourceFacePictureUrl *string `json:"SourceFacePictureUrl,omitempty" xml:"SourceFacePictureUrl,omitempty"`
	TargetFacePicture    *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	FaceComparisonScore *float64 `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	Passed              *string  `json:"Passed,omitempty" xml:"Passed,omitempty"`
	TransactionId       *string  `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FaceCompareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Crop              *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	FacePictureUrl    *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	FaceQuality       *string `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	MerchantBizId     *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	MerchantUserId    *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	Occlusion         *string `json:"Occlusion,omitempty" xml:"Occlusion,omitempty"`
	ProductCode       *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
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
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
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
	ExtFaceInfo   *FaceLivenessResponseBodyResultExtFaceInfo `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty" type:"Struct"`
	Passed        *string                                    `json:"Passed,omitempty" xml:"Passed,omitempty"`
	SubCode       *string                                    `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	TransactionId *string                                    `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
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
	FaceAttack       *string  `json:"FaceAttack,omitempty" xml:"FaceAttack,omitempty"`
	FaceQualityScore *float64 `json:"FaceQualityScore,omitempty" xml:"FaceQualityScore,omitempty"`
	OcclusionResult  *string  `json:"OcclusionResult,omitempty" xml:"OcclusionResult,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FaceLivenessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type InitializeRequest struct {
	Authorize         *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	Crop              *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DocType           *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	FacePictureUrl    *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	FlowType          *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	IdFaceQuality     *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	IdSpoof           *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
	MerchantBizId     *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	MerchantUserId    *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	MetaInfo          *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	// OCR。
	Ocr           *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	OperationMode *string `json:"OperationMode,omitempty" xml:"OperationMode,omitempty"`
	Pages         *string `json:"Pages,omitempty" xml:"Pages,omitempty"`
	ProductCode   *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductConfig *string `json:"ProductConfig,omitempty" xml:"ProductConfig,omitempty"`
	ReturnUrl     *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	SceneCode     *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	ServiceLevel  *string `json:"ServiceLevel,omitempty" xml:"ServiceLevel,omitempty"`
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

func (s *InitializeRequest) SetCrop(v string) *InitializeRequest {
	s.Crop = &v
	return s
}

func (s *InitializeRequest) SetDocType(v string) *InitializeRequest {
	s.DocType = &v
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

func (s *InitializeRequest) SetFlowType(v string) *InitializeRequest {
	s.FlowType = &v
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

func (s *InitializeRequest) SetOcr(v string) *InitializeRequest {
	s.Ocr = &v
	return s
}

func (s *InitializeRequest) SetOperationMode(v string) *InitializeRequest {
	s.OperationMode = &v
	return s
}

func (s *InitializeRequest) SetPages(v string) *InitializeRequest {
	s.Pages = &v
	return s
}

func (s *InitializeRequest) SetProductCode(v string) *InitializeRequest {
	s.ProductCode = &v
	return s
}

func (s *InitializeRequest) SetProductConfig(v string) *InitializeRequest {
	s.ProductConfig = &v
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

func (s *InitializeRequest) SetServiceLevel(v string) *InitializeRequest {
	s.ServiceLevel = &v
	return s
}

type InitializeResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	ClientCfg      *string `json:"ClientCfg,omitempty" xml:"ClientCfg,omitempty"`
	TransactionId  *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InitializeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

/**
 * @deprecated : CardOcr is deprecated, please use Cloudauth-intl::2022-08-09::DocOcr instead.
 *
 * @param request CardOcrRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CardOcrResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.IdOcrPictureBase64)) {
		query["IdOcrPictureBase64"] = request.IdOcrPictureBase64
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
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

/**
 * @deprecated : CardOcr is deprecated, please use Cloudauth-intl::2022-08-09::DocOcr instead.
 *
 * @param request CardOcrRequest
 * @return CardOcrResponse
 */
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

func (client *Client) DeletePictureWithOptions(request *DeletePictureRequest, runtime *util.RuntimeOptions) (_result *DeletePictureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeletePicAfterQuery)) {
		query["DeletePicAfterQuery"] = request.DeletePicAfterQuery
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionId)) {
		query["TransactionId"] = request.TransactionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePicture"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePictureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePicture(request *DeletePictureRequest) (_result *DeletePictureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePictureResponse{}
	_body, _err := client.DeletePictureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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

	if !tea.BoolValue(util.IsUnset(request.IdOcrPictureBase64)) {
		query["IdOcrPictureBase64"] = request.IdOcrPictureBase64
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
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

func (client *Client) FaceCompareWithOptions(request *FaceCompareRequest, runtime *util.RuntimeOptions) (_result *FaceCompareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceFacePicture)) {
		query["SourceFacePicture"] = request.SourceFacePicture
	}

	if !tea.BoolValue(util.IsUnset(request.SourceFacePictureUrl)) {
		query["SourceFacePictureUrl"] = request.SourceFacePictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFacePicture)) {
		query["TargetFacePicture"] = request.TargetFacePicture
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFacePictureUrl)) {
		query["TargetFacePictureUrl"] = request.TargetFacePictureUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
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

func (client *Client) FaceLivenessWithOptions(request *FaceLivenessRequest, runtime *util.RuntimeOptions) (_result *FaceLivenessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		query["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.FacePictureBase64)) {
		query["FacePictureBase64"] = request.FacePictureBase64
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
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

func (client *Client) InitializeWithOptions(request *InitializeRequest, runtime *util.RuntimeOptions) (_result *InitializeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		query["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.FacePictureBase64)) {
		query["FacePictureBase64"] = request.FacePictureBase64
	}

	if !tea.BoolValue(util.IsUnset(request.FacePictureUrl)) {
		query["FacePictureUrl"] = request.FacePictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FlowType)) {
		query["FlowType"] = request.FlowType
	}

	if !tea.BoolValue(util.IsUnset(request.IdFaceQuality)) {
		query["IdFaceQuality"] = request.IdFaceQuality
	}

	if !tea.BoolValue(util.IsUnset(request.IdSpoof)) {
		query["IdSpoof"] = request.IdSpoof
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

	if !tea.BoolValue(util.IsUnset(request.Ocr)) {
		query["Ocr"] = request.Ocr
	}

	if !tea.BoolValue(util.IsUnset(request.OperationMode)) {
		query["OperationMode"] = request.OperationMode
	}

	if !tea.BoolValue(util.IsUnset(request.Pages)) {
		query["Pages"] = request.Pages
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductConfig)) {
		query["ProductConfig"] = request.ProductConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnUrl)) {
		query["ReturnUrl"] = request.ReturnUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceLevel)) {
		query["ServiceLevel"] = request.ServiceLevel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
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
