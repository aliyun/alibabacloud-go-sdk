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

type CompanyFourElementsVerificationRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	EpCertName           *string `json:"EpCertName,omitempty" xml:"EpCertName,omitempty"`
	EpCertNo             *string `json:"EpCertNo,omitempty" xml:"EpCertNo,omitempty"`
	LegalPersonCertName  *string `json:"LegalPersonCertName,omitempty" xml:"LegalPersonCertName,omitempty"`
	LegalPersonCertNo    *string `json:"LegalPersonCertNo,omitempty" xml:"LegalPersonCertNo,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CompanyFourElementsVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s CompanyFourElementsVerificationRequest) GoString() string {
	return s.String()
}

func (s *CompanyFourElementsVerificationRequest) SetAuthCode(v string) *CompanyFourElementsVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *CompanyFourElementsVerificationRequest) SetEpCertName(v string) *CompanyFourElementsVerificationRequest {
	s.EpCertName = &v
	return s
}

func (s *CompanyFourElementsVerificationRequest) SetEpCertNo(v string) *CompanyFourElementsVerificationRequest {
	s.EpCertNo = &v
	return s
}

func (s *CompanyFourElementsVerificationRequest) SetLegalPersonCertName(v string) *CompanyFourElementsVerificationRequest {
	s.LegalPersonCertName = &v
	return s
}

func (s *CompanyFourElementsVerificationRequest) SetLegalPersonCertNo(v string) *CompanyFourElementsVerificationRequest {
	s.LegalPersonCertNo = &v
	return s
}

func (s *CompanyFourElementsVerificationRequest) SetOwnerId(v int64) *CompanyFourElementsVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *CompanyFourElementsVerificationRequest) SetResourceOwnerAccount(v string) *CompanyFourElementsVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CompanyFourElementsVerificationRequest) SetResourceOwnerId(v int64) *CompanyFourElementsVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

type CompanyFourElementsVerificationResponseBody struct {
	AccessDeniedDetail *string                                          `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *CompanyFourElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CompanyFourElementsVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompanyFourElementsVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *CompanyFourElementsVerificationResponseBody) SetAccessDeniedDetail(v string) *CompanyFourElementsVerificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CompanyFourElementsVerificationResponseBody) SetCode(v string) *CompanyFourElementsVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *CompanyFourElementsVerificationResponseBody) SetData(v *CompanyFourElementsVerificationResponseBodyData) *CompanyFourElementsVerificationResponseBody {
	s.Data = v
	return s
}

func (s *CompanyFourElementsVerificationResponseBody) SetMessage(v string) *CompanyFourElementsVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *CompanyFourElementsVerificationResponseBody) SetRequestId(v string) *CompanyFourElementsVerificationResponseBody {
	s.RequestId = &v
	return s
}

type CompanyFourElementsVerificationResponseBodyData struct {
	DetailInfo   map[string]interface{} `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty"`
	ReasonCode   *int64                 `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	VerifyResult *string                `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s CompanyFourElementsVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CompanyFourElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompanyFourElementsVerificationResponseBodyData) SetDetailInfo(v map[string]interface{}) *CompanyFourElementsVerificationResponseBodyData {
	s.DetailInfo = v
	return s
}

func (s *CompanyFourElementsVerificationResponseBodyData) SetReasonCode(v int64) *CompanyFourElementsVerificationResponseBodyData {
	s.ReasonCode = &v
	return s
}

func (s *CompanyFourElementsVerificationResponseBodyData) SetVerifyResult(v string) *CompanyFourElementsVerificationResponseBodyData {
	s.VerifyResult = &v
	return s
}

type CompanyFourElementsVerificationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CompanyFourElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompanyFourElementsVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s CompanyFourElementsVerificationResponse) GoString() string {
	return s.String()
}

func (s *CompanyFourElementsVerificationResponse) SetHeaders(v map[string]*string) *CompanyFourElementsVerificationResponse {
	s.Headers = v
	return s
}

func (s *CompanyFourElementsVerificationResponse) SetStatusCode(v int32) *CompanyFourElementsVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CompanyFourElementsVerificationResponse) SetBody(v *CompanyFourElementsVerificationResponseBody) *CompanyFourElementsVerificationResponse {
	s.Body = v
	return s
}

type CompanyThreeElementsVerificationRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	EpCertName           *string `json:"EpCertName,omitempty" xml:"EpCertName,omitempty"`
	EpCertNo             *string `json:"EpCertNo,omitempty" xml:"EpCertNo,omitempty"`
	LegalPersonCertName  *string `json:"LegalPersonCertName,omitempty" xml:"LegalPersonCertName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CompanyThreeElementsVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s CompanyThreeElementsVerificationRequest) GoString() string {
	return s.String()
}

func (s *CompanyThreeElementsVerificationRequest) SetAuthCode(v string) *CompanyThreeElementsVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *CompanyThreeElementsVerificationRequest) SetEpCertName(v string) *CompanyThreeElementsVerificationRequest {
	s.EpCertName = &v
	return s
}

func (s *CompanyThreeElementsVerificationRequest) SetEpCertNo(v string) *CompanyThreeElementsVerificationRequest {
	s.EpCertNo = &v
	return s
}

func (s *CompanyThreeElementsVerificationRequest) SetLegalPersonCertName(v string) *CompanyThreeElementsVerificationRequest {
	s.LegalPersonCertName = &v
	return s
}

func (s *CompanyThreeElementsVerificationRequest) SetOwnerId(v int64) *CompanyThreeElementsVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *CompanyThreeElementsVerificationRequest) SetResourceOwnerAccount(v string) *CompanyThreeElementsVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CompanyThreeElementsVerificationRequest) SetResourceOwnerId(v int64) *CompanyThreeElementsVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

type CompanyThreeElementsVerificationResponseBody struct {
	AccessDeniedDetail *string                                           `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *CompanyThreeElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CompanyThreeElementsVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompanyThreeElementsVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *CompanyThreeElementsVerificationResponseBody) SetAccessDeniedDetail(v string) *CompanyThreeElementsVerificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBody) SetCode(v string) *CompanyThreeElementsVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBody) SetData(v *CompanyThreeElementsVerificationResponseBodyData) *CompanyThreeElementsVerificationResponseBody {
	s.Data = v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBody) SetMessage(v string) *CompanyThreeElementsVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBody) SetRequestId(v string) *CompanyThreeElementsVerificationResponseBody {
	s.RequestId = &v
	return s
}

type CompanyThreeElementsVerificationResponseBodyData struct {
	DetailInfo   map[string]interface{} `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty"`
	ReasonCode   *int64                 `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	VerifyResult *string                `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s CompanyThreeElementsVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CompanyThreeElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompanyThreeElementsVerificationResponseBodyData) SetDetailInfo(v map[string]interface{}) *CompanyThreeElementsVerificationResponseBodyData {
	s.DetailInfo = v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBodyData) SetReasonCode(v int64) *CompanyThreeElementsVerificationResponseBodyData {
	s.ReasonCode = &v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBodyData) SetVerifyResult(v string) *CompanyThreeElementsVerificationResponseBodyData {
	s.VerifyResult = &v
	return s
}

type CompanyThreeElementsVerificationResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CompanyThreeElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompanyThreeElementsVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s CompanyThreeElementsVerificationResponse) GoString() string {
	return s.String()
}

func (s *CompanyThreeElementsVerificationResponse) SetHeaders(v map[string]*string) *CompanyThreeElementsVerificationResponse {
	s.Headers = v
	return s
}

func (s *CompanyThreeElementsVerificationResponse) SetStatusCode(v int32) *CompanyThreeElementsVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CompanyThreeElementsVerificationResponse) SetBody(v *CompanyThreeElementsVerificationResponseBody) *CompanyThreeElementsVerificationResponse {
	s.Body = v
	return s
}

type CompanyTwoElementsVerificationRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	EpCertName           *string `json:"EpCertName,omitempty" xml:"EpCertName,omitempty"`
	EpCertNo             *string `json:"EpCertNo,omitempty" xml:"EpCertNo,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CompanyTwoElementsVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s CompanyTwoElementsVerificationRequest) GoString() string {
	return s.String()
}

func (s *CompanyTwoElementsVerificationRequest) SetAuthCode(v string) *CompanyTwoElementsVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *CompanyTwoElementsVerificationRequest) SetEpCertName(v string) *CompanyTwoElementsVerificationRequest {
	s.EpCertName = &v
	return s
}

func (s *CompanyTwoElementsVerificationRequest) SetEpCertNo(v string) *CompanyTwoElementsVerificationRequest {
	s.EpCertNo = &v
	return s
}

func (s *CompanyTwoElementsVerificationRequest) SetOwnerId(v int64) *CompanyTwoElementsVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *CompanyTwoElementsVerificationRequest) SetResourceOwnerAccount(v string) *CompanyTwoElementsVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CompanyTwoElementsVerificationRequest) SetResourceOwnerId(v int64) *CompanyTwoElementsVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

type CompanyTwoElementsVerificationResponseBody struct {
	AccessDeniedDetail *string                                         `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *CompanyTwoElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CompanyTwoElementsVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompanyTwoElementsVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *CompanyTwoElementsVerificationResponseBody) SetAccessDeniedDetail(v string) *CompanyTwoElementsVerificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBody) SetCode(v string) *CompanyTwoElementsVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBody) SetData(v *CompanyTwoElementsVerificationResponseBodyData) *CompanyTwoElementsVerificationResponseBody {
	s.Data = v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBody) SetMessage(v string) *CompanyTwoElementsVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBody) SetRequestId(v string) *CompanyTwoElementsVerificationResponseBody {
	s.RequestId = &v
	return s
}

type CompanyTwoElementsVerificationResponseBodyData struct {
	DetailInfo   map[string]interface{} `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty"`
	ReasonCode   *string                `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	VerifyResult *string                `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s CompanyTwoElementsVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CompanyTwoElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompanyTwoElementsVerificationResponseBodyData) SetDetailInfo(v map[string]interface{}) *CompanyTwoElementsVerificationResponseBodyData {
	s.DetailInfo = v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBodyData) SetReasonCode(v string) *CompanyTwoElementsVerificationResponseBodyData {
	s.ReasonCode = &v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBodyData) SetVerifyResult(v string) *CompanyTwoElementsVerificationResponseBodyData {
	s.VerifyResult = &v
	return s
}

type CompanyTwoElementsVerificationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CompanyTwoElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompanyTwoElementsVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s CompanyTwoElementsVerificationResponse) GoString() string {
	return s.String()
}

func (s *CompanyTwoElementsVerificationResponse) SetHeaders(v map[string]*string) *CompanyTwoElementsVerificationResponse {
	s.Headers = v
	return s
}

func (s *CompanyTwoElementsVerificationResponse) SetStatusCode(v int32) *CompanyTwoElementsVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CompanyTwoElementsVerificationResponse) SetBody(v *CompanyTwoElementsVerificationResponseBody) *CompanyTwoElementsVerificationResponse {
	s.Body = v
	return s
}

type DescribeEmptyNumberRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeEmptyNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmptyNumberRequest) GoString() string {
	return s.String()
}

func (s *DescribeEmptyNumberRequest) SetAuthCode(v string) *DescribeEmptyNumberRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribeEmptyNumberRequest) SetInputNumber(v string) *DescribeEmptyNumberRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribeEmptyNumberRequest) SetMask(v string) *DescribeEmptyNumberRequest {
	s.Mask = &v
	return s
}

func (s *DescribeEmptyNumberRequest) SetOwnerId(v int64) *DescribeEmptyNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEmptyNumberRequest) SetResourceOwnerAccount(v string) *DescribeEmptyNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEmptyNumberRequest) SetResourceOwnerId(v int64) *DescribeEmptyNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeEmptyNumberResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeEmptyNumberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEmptyNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmptyNumberResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEmptyNumberResponseBody) SetCode(v string) *DescribeEmptyNumberResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEmptyNumberResponseBody) SetData(v *DescribeEmptyNumberResponseBodyData) *DescribeEmptyNumberResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEmptyNumberResponseBody) SetMessage(v string) *DescribeEmptyNumberResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEmptyNumberResponseBody) SetRequestId(v string) *DescribeEmptyNumberResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEmptyNumberResponseBodyData struct {
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEmptyNumberResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmptyNumberResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEmptyNumberResponseBodyData) SetNumber(v string) *DescribeEmptyNumberResponseBodyData {
	s.Number = &v
	return s
}

func (s *DescribeEmptyNumberResponseBodyData) SetStatus(v string) *DescribeEmptyNumberResponseBodyData {
	s.Status = &v
	return s
}

type DescribeEmptyNumberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEmptyNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEmptyNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmptyNumberResponse) GoString() string {
	return s.String()
}

func (s *DescribeEmptyNumberResponse) SetHeaders(v map[string]*string) *DescribeEmptyNumberResponse {
	s.Headers = v
	return s
}

func (s *DescribeEmptyNumberResponse) SetStatusCode(v int32) *DescribeEmptyNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEmptyNumberResponse) SetBody(v *DescribeEmptyNumberResponseBody) *DescribeEmptyNumberResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberAnalysisRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	NumberType           *int64  `json:"NumberType,omitempty" xml:"NumberType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Rate                 *int64  `json:"Rate,omitempty" xml:"Rate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisRequest) SetAuthCode(v string) *DescribePhoneNumberAnalysisRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetInputNumber(v string) *DescribePhoneNumberAnalysisRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetMask(v string) *DescribePhoneNumberAnalysisRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetNumberType(v int64) *DescribePhoneNumberAnalysisRequest {
	s.NumberType = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetOwnerId(v int64) *DescribePhoneNumberAnalysisRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetRate(v int64) *DescribePhoneNumberAnalysisRequest {
	s.Rate = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberAnalysisRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberAnalysisRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribePhoneNumberAnalysisResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribePhoneNumberAnalysisResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetCode(v string) *DescribePhoneNumberAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetData(v *DescribePhoneNumberAnalysisResponseBodyData) *DescribePhoneNumberAnalysisResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetMessage(v string) *DescribePhoneNumberAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetRequestId(v string) *DescribePhoneNumberAnalysisResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberAnalysisResponseBodyData struct {
	List []*DescribePhoneNumberAnalysisResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribePhoneNumberAnalysisResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisResponseBodyData) SetList(v []*DescribePhoneNumberAnalysisResponseBodyDataList) *DescribePhoneNumberAnalysisResponseBodyData {
	s.List = v
	return s
}

type DescribePhoneNumberAnalysisResponseBodyDataList struct {
	Code   *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s DescribePhoneNumberAnalysisResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisResponseBodyDataList) SetCode(v string) *DescribePhoneNumberAnalysisResponseBodyDataList {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBodyDataList) SetNumber(v string) *DescribePhoneNumberAnalysisResponseBodyDataList {
	s.Number = &v
	return s
}

type DescribePhoneNumberAnalysisResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePhoneNumberAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneNumberAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberAnalysisResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberAnalysisResponse) SetStatusCode(v int32) *DescribePhoneNumberAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponse) SetBody(v *DescribePhoneNumberAnalysisResponseBody) *DescribePhoneNumberAnalysisResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberAnalysisAIRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	ModelConfig          *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Rate                 *int64  `json:"Rate,omitempty" xml:"Rate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberAnalysisAIRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisAIRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisAIRequest) SetAuthCode(v string) *DescribePhoneNumberAnalysisAIRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIRequest) SetInputNumber(v string) *DescribePhoneNumberAnalysisAIRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIRequest) SetModelConfig(v string) *DescribePhoneNumberAnalysisAIRequest {
	s.ModelConfig = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIRequest) SetOwnerId(v int64) *DescribePhoneNumberAnalysisAIRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIRequest) SetRate(v int64) *DescribePhoneNumberAnalysisAIRequest {
	s.Rate = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberAnalysisAIRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberAnalysisAIRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribePhoneNumberAnalysisAIResponseBody struct {
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribePhoneNumberAnalysisAIResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberAnalysisAIResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisAIResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisAIResponseBody) SetCode(v string) *DescribePhoneNumberAnalysisAIResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIResponseBody) SetData(v *DescribePhoneNumberAnalysisAIResponseBodyData) *DescribePhoneNumberAnalysisAIResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberAnalysisAIResponseBody) SetMessage(v string) *DescribePhoneNumberAnalysisAIResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIResponseBody) SetRequestId(v string) *DescribePhoneNumberAnalysisAIResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberAnalysisAIResponseBodyData struct {
	Code   *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s DescribePhoneNumberAnalysisAIResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisAIResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisAIResponseBodyData) SetCode(v string) *DescribePhoneNumberAnalysisAIResponseBodyData {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIResponseBodyData) SetNumber(v string) *DescribePhoneNumberAnalysisAIResponseBodyData {
	s.Number = &v
	return s
}

type DescribePhoneNumberAnalysisAIResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePhoneNumberAnalysisAIResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneNumberAnalysisAIResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisAIResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisAIResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberAnalysisAIResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberAnalysisAIResponse) SetStatusCode(v int32) *DescribePhoneNumberAnalysisAIResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIResponse) SetBody(v *DescribePhoneNumberAnalysisAIResponseBody) *DescribePhoneNumberAnalysisAIResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberAttributeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAttributeRequest) SetOwnerId(v int64) *DescribePhoneNumberAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberAttributeRequest) SetPhoneNumber(v string) *DescribePhoneNumberAttributeRequest {
	s.PhoneNumber = &v
	return s
}

func (s *DescribePhoneNumberAttributeRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberAttributeRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribePhoneNumberAttributeResponseBody struct {
	Code                 *string                                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message              *string                                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	PhoneNumberAttribute *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute `json:"PhoneNumberAttribute,omitempty" xml:"PhoneNumberAttribute,omitempty" type:"Struct"`
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAttributeResponseBody) SetCode(v string) *DescribePhoneNumberAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBody) SetMessage(v string) *DescribePhoneNumberAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBody) SetPhoneNumberAttribute(v *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) *DescribePhoneNumberAttributeResponseBody {
	s.PhoneNumberAttribute = v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBody) SetRequestId(v string) *DescribePhoneNumberAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute struct {
	BasicCarrier        *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	Carrier             *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	City                *string `json:"City,omitempty" xml:"City,omitempty"`
	IsNumberPortability *bool   `json:"IsNumberPortability,omitempty" xml:"IsNumberPortability,omitempty"`
	NumberSegment       *int64  `json:"NumberSegment,omitempty" xml:"NumberSegment,omitempty"`
	Province            *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetBasicCarrier(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.BasicCarrier = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetCarrier(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetCity(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.City = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetIsNumberPortability(v bool) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.IsNumberPortability = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetNumberSegment(v int64) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.NumberSegment = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetProvince(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.Province = &v
	return s
}

type DescribePhoneNumberAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePhoneNumberAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneNumberAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAttributeResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberAttributeResponse) SetStatusCode(v int32) *DescribePhoneNumberAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponse) SetBody(v *DescribePhoneNumberAttributeResponseBody) *DescribePhoneNumberAttributeResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberOnlineTimeRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	Carrier              *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberOnlineTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOnlineTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetAuthCode(v string) *DescribePhoneNumberOnlineTimeRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetCarrier(v string) *DescribePhoneNumberOnlineTimeRequest {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetInputNumber(v string) *DescribePhoneNumberOnlineTimeRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetMask(v string) *DescribePhoneNumberOnlineTimeRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetOwnerId(v int64) *DescribePhoneNumberOnlineTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberOnlineTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberOnlineTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribePhoneNumberOnlineTimeResponseBody struct {
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribePhoneNumberOnlineTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberOnlineTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOnlineTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) SetCode(v string) *DescribePhoneNumberOnlineTimeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) SetData(v *DescribePhoneNumberOnlineTimeResponseBodyData) *DescribePhoneNumberOnlineTimeResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) SetMessage(v string) *DescribePhoneNumberOnlineTimeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) SetRequestId(v string) *DescribePhoneNumberOnlineTimeResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberOnlineTimeResponseBodyData struct {
	CarrierCode  *string `json:"CarrierCode,omitempty" xml:"CarrierCode,omitempty"`
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s DescribePhoneNumberOnlineTimeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOnlineTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOnlineTimeResponseBodyData) SetCarrierCode(v string) *DescribePhoneNumberOnlineTimeResponseBodyData {
	s.CarrierCode = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponseBodyData) SetVerifyResult(v string) *DescribePhoneNumberOnlineTimeResponseBodyData {
	s.VerifyResult = &v
	return s
}

type DescribePhoneNumberOnlineTimeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePhoneNumberOnlineTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneNumberOnlineTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOnlineTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOnlineTimeResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberOnlineTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponse) SetStatusCode(v int32) *DescribePhoneNumberOnlineTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponse) SetBody(v *DescribePhoneNumberOnlineTimeResponseBody) *DescribePhoneNumberOnlineTimeResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberOperatorAttributeRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetAuthCode(v string) *DescribePhoneNumberOperatorAttributeRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetInputNumber(v string) *DescribePhoneNumberOperatorAttributeRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetMask(v string) *DescribePhoneNumberOperatorAttributeRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetOwnerId(v int64) *DescribePhoneNumberOperatorAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberOperatorAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberOperatorAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribePhoneNumberOperatorAttributeResponseBody struct {
	Code      *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribePhoneNumberOperatorAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) SetCode(v string) *DescribePhoneNumberOperatorAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) SetData(v *DescribePhoneNumberOperatorAttributeResponseBodyData) *DescribePhoneNumberOperatorAttributeResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) SetMessage(v string) *DescribePhoneNumberOperatorAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) SetRequestId(v string) *DescribePhoneNumberOperatorAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberOperatorAttributeResponseBodyData struct {
	BasicCarrier        *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	Carrier             *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	City                *string `json:"City,omitempty" xml:"City,omitempty"`
	IsNumberPortability *bool   `json:"IsNumberPortability,omitempty" xml:"IsNumberPortability,omitempty"`
	NumberSegment       *int64  `json:"NumberSegment,omitempty" xml:"NumberSegment,omitempty"`
	Province            *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetBasicCarrier(v string) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetCarrier(v string) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetCity(v string) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.City = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetIsNumberPortability(v bool) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.IsNumberPortability = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetNumberSegment(v int64) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.NumberSegment = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetProvince(v string) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.Province = &v
	return s
}

type DescribePhoneNumberOperatorAttributeResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePhoneNumberOperatorAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneNumberOperatorAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberOperatorAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponse) SetStatusCode(v int32) *DescribePhoneNumberOperatorAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponse) SetBody(v *DescribePhoneNumberOperatorAttributeResponseBody) *DescribePhoneNumberOperatorAttributeResponse {
	s.Body = v
	return s
}

type DescribePhoneTwiceTelVerifyRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePhoneTwiceTelVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneTwiceTelVerifyRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetAuthCode(v string) *DescribePhoneTwiceTelVerifyRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetInputNumber(v string) *DescribePhoneTwiceTelVerifyRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetMask(v string) *DescribePhoneTwiceTelVerifyRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetOwnerId(v int64) *DescribePhoneTwiceTelVerifyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetResourceOwnerAccount(v string) *DescribePhoneTwiceTelVerifyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetResourceOwnerId(v int64) *DescribePhoneTwiceTelVerifyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetStartTime(v string) *DescribePhoneTwiceTelVerifyRequest {
	s.StartTime = &v
	return s
}

type DescribePhoneTwiceTelVerifyResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribePhoneTwiceTelVerifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneTwiceTelVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneTwiceTelVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) SetCode(v string) *DescribePhoneTwiceTelVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) SetData(v *DescribePhoneTwiceTelVerifyResponseBodyData) *DescribePhoneTwiceTelVerifyResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) SetMessage(v string) *DescribePhoneTwiceTelVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) SetRequestId(v string) *DescribePhoneTwiceTelVerifyResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneTwiceTelVerifyResponseBodyData struct {
	Carrier      *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s DescribePhoneTwiceTelVerifyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneTwiceTelVerifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneTwiceTelVerifyResponseBodyData) SetCarrier(v string) *DescribePhoneTwiceTelVerifyResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponseBodyData) SetVerifyResult(v string) *DescribePhoneTwiceTelVerifyResponseBodyData {
	s.VerifyResult = &v
	return s
}

type DescribePhoneTwiceTelVerifyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePhoneTwiceTelVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneTwiceTelVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneTwiceTelVerifyResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneTwiceTelVerifyResponse) SetHeaders(v map[string]*string) *DescribePhoneTwiceTelVerifyResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponse) SetStatusCode(v int32) *DescribePhoneTwiceTelVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponse) SetBody(v *DescribePhoneTwiceTelVerifyResponseBody) *DescribePhoneTwiceTelVerifyResponse {
	s.Body = v
	return s
}

type InvalidPhoneNumberFilterRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s InvalidPhoneNumberFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s InvalidPhoneNumberFilterRequest) GoString() string {
	return s.String()
}

func (s *InvalidPhoneNumberFilterRequest) SetAuthCode(v string) *InvalidPhoneNumberFilterRequest {
	s.AuthCode = &v
	return s
}

func (s *InvalidPhoneNumberFilterRequest) SetInputNumber(v string) *InvalidPhoneNumberFilterRequest {
	s.InputNumber = &v
	return s
}

func (s *InvalidPhoneNumberFilterRequest) SetMask(v string) *InvalidPhoneNumberFilterRequest {
	s.Mask = &v
	return s
}

func (s *InvalidPhoneNumberFilterRequest) SetOwnerId(v int64) *InvalidPhoneNumberFilterRequest {
	s.OwnerId = &v
	return s
}

func (s *InvalidPhoneNumberFilterRequest) SetResourceOwnerAccount(v string) *InvalidPhoneNumberFilterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *InvalidPhoneNumberFilterRequest) SetResourceOwnerId(v int64) *InvalidPhoneNumberFilterRequest {
	s.ResourceOwnerId = &v
	return s
}

type InvalidPhoneNumberFilterResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*InvalidPhoneNumberFilterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InvalidPhoneNumberFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvalidPhoneNumberFilterResponseBody) GoString() string {
	return s.String()
}

func (s *InvalidPhoneNumberFilterResponseBody) SetCode(v string) *InvalidPhoneNumberFilterResponseBody {
	s.Code = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBody) SetData(v []*InvalidPhoneNumberFilterResponseBodyData) *InvalidPhoneNumberFilterResponseBody {
	s.Data = v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBody) SetMessage(v string) *InvalidPhoneNumberFilterResponseBody {
	s.Message = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBody) SetRequestId(v string) *InvalidPhoneNumberFilterResponseBody {
	s.RequestId = &v
	return s
}

type InvalidPhoneNumberFilterResponseBodyData struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	EncryptedNumber *string `json:"EncryptedNumber,omitempty" xml:"EncryptedNumber,omitempty"`
	ExpireTime      *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	OriginalNumber  *string `json:"OriginalNumber,omitempty" xml:"OriginalNumber,omitempty"`
}

func (s InvalidPhoneNumberFilterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InvalidPhoneNumberFilterResponseBodyData) GoString() string {
	return s.String()
}

func (s *InvalidPhoneNumberFilterResponseBodyData) SetCode(v string) *InvalidPhoneNumberFilterResponseBodyData {
	s.Code = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBodyData) SetEncryptedNumber(v string) *InvalidPhoneNumberFilterResponseBodyData {
	s.EncryptedNumber = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBodyData) SetExpireTime(v string) *InvalidPhoneNumberFilterResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponseBodyData) SetOriginalNumber(v string) *InvalidPhoneNumberFilterResponseBodyData {
	s.OriginalNumber = &v
	return s
}

type InvalidPhoneNumberFilterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InvalidPhoneNumberFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvalidPhoneNumberFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s InvalidPhoneNumberFilterResponse) GoString() string {
	return s.String()
}

func (s *InvalidPhoneNumberFilterResponse) SetHeaders(v map[string]*string) *InvalidPhoneNumberFilterResponse {
	s.Headers = v
	return s
}

func (s *InvalidPhoneNumberFilterResponse) SetStatusCode(v int32) *InvalidPhoneNumberFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponse) SetBody(v *InvalidPhoneNumberFilterResponseBody) *InvalidPhoneNumberFilterResponse {
	s.Body = v
	return s
}

type PhoneNumberConvertServiceRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberConvertServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberConvertServiceRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberConvertServiceRequest) SetAuthCode(v string) *PhoneNumberConvertServiceRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberConvertServiceRequest) SetInputNumber(v string) *PhoneNumberConvertServiceRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberConvertServiceRequest) SetMask(v string) *PhoneNumberConvertServiceRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberConvertServiceRequest) SetOwnerId(v int64) *PhoneNumberConvertServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberConvertServiceRequest) SetResourceOwnerAccount(v string) *PhoneNumberConvertServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberConvertServiceRequest) SetResourceOwnerId(v int64) *PhoneNumberConvertServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

type PhoneNumberConvertServiceResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*PhoneNumberConvertServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberConvertServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberConvertServiceResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberConvertServiceResponseBody) SetCode(v string) *PhoneNumberConvertServiceResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberConvertServiceResponseBody) SetData(v []*PhoneNumberConvertServiceResponseBodyData) *PhoneNumberConvertServiceResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberConvertServiceResponseBody) SetMessage(v string) *PhoneNumberConvertServiceResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberConvertServiceResponseBody) SetRequestId(v string) *PhoneNumberConvertServiceResponseBody {
	s.RequestId = &v
	return s
}

type PhoneNumberConvertServiceResponseBodyData struct {
	ConverResult *bool   `json:"ConverResult,omitempty" xml:"ConverResult,omitempty"`
	Number       *string `json:"Number,omitempty" xml:"Number,omitempty"`
	NumberMd5    *string `json:"NumberMd5,omitempty" xml:"NumberMd5,omitempty"`
	NumberSha256 *string `json:"NumberSha256,omitempty" xml:"NumberSha256,omitempty"`
}

func (s PhoneNumberConvertServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberConvertServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberConvertServiceResponseBodyData) SetConverResult(v bool) *PhoneNumberConvertServiceResponseBodyData {
	s.ConverResult = &v
	return s
}

func (s *PhoneNumberConvertServiceResponseBodyData) SetNumber(v string) *PhoneNumberConvertServiceResponseBodyData {
	s.Number = &v
	return s
}

func (s *PhoneNumberConvertServiceResponseBodyData) SetNumberMd5(v string) *PhoneNumberConvertServiceResponseBodyData {
	s.NumberMd5 = &v
	return s
}

func (s *PhoneNumberConvertServiceResponseBodyData) SetNumberSha256(v string) *PhoneNumberConvertServiceResponseBodyData {
	s.NumberSha256 = &v
	return s
}

type PhoneNumberConvertServiceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PhoneNumberConvertServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PhoneNumberConvertServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberConvertServiceResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberConvertServiceResponse) SetHeaders(v map[string]*string) *PhoneNumberConvertServiceResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberConvertServiceResponse) SetStatusCode(v int32) *PhoneNumberConvertServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberConvertServiceResponse) SetBody(v *PhoneNumberConvertServiceResponseBody) *PhoneNumberConvertServiceResponse {
	s.Body = v
	return s
}

type PhoneNumberEncryptRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberEncryptRequest) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberEncryptRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberEncryptRequest) SetAuthCode(v string) *PhoneNumberEncryptRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberEncryptRequest) SetInputNumber(v string) *PhoneNumberEncryptRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberEncryptRequest) SetMask(v string) *PhoneNumberEncryptRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberEncryptRequest) SetOwnerId(v int64) *PhoneNumberEncryptRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberEncryptRequest) SetResourceOwnerAccount(v string) *PhoneNumberEncryptRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberEncryptRequest) SetResourceOwnerId(v int64) *PhoneNumberEncryptRequest {
	s.ResourceOwnerId = &v
	return s
}

type PhoneNumberEncryptResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*PhoneNumberEncryptResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberEncryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberEncryptResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberEncryptResponseBody) SetCode(v string) *PhoneNumberEncryptResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberEncryptResponseBody) SetData(v []*PhoneNumberEncryptResponseBodyData) *PhoneNumberEncryptResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberEncryptResponseBody) SetMessage(v string) *PhoneNumberEncryptResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberEncryptResponseBody) SetRequestId(v string) *PhoneNumberEncryptResponseBody {
	s.RequestId = &v
	return s
}

type PhoneNumberEncryptResponseBodyData struct {
	EncryptedNumber *string `json:"EncryptedNumber,omitempty" xml:"EncryptedNumber,omitempty"`
	ExpireTime      *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	OriginalNumber  *string `json:"OriginalNumber,omitempty" xml:"OriginalNumber,omitempty"`
}

func (s PhoneNumberEncryptResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberEncryptResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberEncryptResponseBodyData) SetEncryptedNumber(v string) *PhoneNumberEncryptResponseBodyData {
	s.EncryptedNumber = &v
	return s
}

func (s *PhoneNumberEncryptResponseBodyData) SetExpireTime(v string) *PhoneNumberEncryptResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *PhoneNumberEncryptResponseBodyData) SetOriginalNumber(v string) *PhoneNumberEncryptResponseBodyData {
	s.OriginalNumber = &v
	return s
}

type PhoneNumberEncryptResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PhoneNumberEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PhoneNumberEncryptResponse) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberEncryptResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberEncryptResponse) SetHeaders(v map[string]*string) *PhoneNumberEncryptResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberEncryptResponse) SetStatusCode(v int32) *PhoneNumberEncryptResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberEncryptResponse) SetBody(v *PhoneNumberEncryptResponseBody) *PhoneNumberEncryptResponse {
	s.Body = v
	return s
}

type PhoneNumberStatusForAccountRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberStatusForAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForAccountRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForAccountRequest) SetAuthCode(v string) *PhoneNumberStatusForAccountRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberStatusForAccountRequest) SetInputNumber(v string) *PhoneNumberStatusForAccountRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberStatusForAccountRequest) SetMask(v string) *PhoneNumberStatusForAccountRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberStatusForAccountRequest) SetOwnerId(v int64) *PhoneNumberStatusForAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberStatusForAccountRequest) SetResourceOwnerAccount(v string) *PhoneNumberStatusForAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberStatusForAccountRequest) SetResourceOwnerId(v int64) *PhoneNumberStatusForAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

type PhoneNumberStatusForAccountResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PhoneNumberStatusForAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberStatusForAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForAccountResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForAccountResponseBody) SetCode(v string) *PhoneNumberStatusForAccountResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberStatusForAccountResponseBody) SetData(v *PhoneNumberStatusForAccountResponseBodyData) *PhoneNumberStatusForAccountResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberStatusForAccountResponseBody) SetMessage(v string) *PhoneNumberStatusForAccountResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberStatusForAccountResponseBody) SetRequestId(v string) *PhoneNumberStatusForAccountResponseBody {
	s.RequestId = &v
	return s
}

type PhoneNumberStatusForAccountResponseBodyData struct {
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PhoneNumberStatusForAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForAccountResponseBodyData) SetCarrier(v string) *PhoneNumberStatusForAccountResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *PhoneNumberStatusForAccountResponseBodyData) SetStatus(v string) *PhoneNumberStatusForAccountResponseBodyData {
	s.Status = &v
	return s
}

type PhoneNumberStatusForAccountResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PhoneNumberStatusForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PhoneNumberStatusForAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForAccountResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForAccountResponse) SetHeaders(v map[string]*string) *PhoneNumberStatusForAccountResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberStatusForAccountResponse) SetStatusCode(v int32) *PhoneNumberStatusForAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberStatusForAccountResponse) SetBody(v *PhoneNumberStatusForAccountResponseBody) *PhoneNumberStatusForAccountResponse {
	s.Body = v
	return s
}

type PhoneNumberStatusForPublicRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberStatusForPublicRequest) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForPublicRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForPublicRequest) SetAuthCode(v string) *PhoneNumberStatusForPublicRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberStatusForPublicRequest) SetInputNumber(v string) *PhoneNumberStatusForPublicRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberStatusForPublicRequest) SetMask(v string) *PhoneNumberStatusForPublicRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberStatusForPublicRequest) SetOwnerId(v int64) *PhoneNumberStatusForPublicRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberStatusForPublicRequest) SetResourceOwnerAccount(v string) *PhoneNumberStatusForPublicRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberStatusForPublicRequest) SetResourceOwnerId(v int64) *PhoneNumberStatusForPublicRequest {
	s.ResourceOwnerId = &v
	return s
}

type PhoneNumberStatusForPublicResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PhoneNumberStatusForPublicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberStatusForPublicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForPublicResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForPublicResponseBody) SetCode(v string) *PhoneNumberStatusForPublicResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberStatusForPublicResponseBody) SetData(v *PhoneNumberStatusForPublicResponseBodyData) *PhoneNumberStatusForPublicResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberStatusForPublicResponseBody) SetMessage(v string) *PhoneNumberStatusForPublicResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberStatusForPublicResponseBody) SetRequestId(v string) *PhoneNumberStatusForPublicResponseBody {
	s.RequestId = &v
	return s
}

type PhoneNumberStatusForPublicResponseBodyData struct {
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PhoneNumberStatusForPublicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForPublicResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForPublicResponseBodyData) SetCarrier(v string) *PhoneNumberStatusForPublicResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *PhoneNumberStatusForPublicResponseBodyData) SetStatus(v string) *PhoneNumberStatusForPublicResponseBodyData {
	s.Status = &v
	return s
}

type PhoneNumberStatusForPublicResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PhoneNumberStatusForPublicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PhoneNumberStatusForPublicResponse) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForPublicResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForPublicResponse) SetHeaders(v map[string]*string) *PhoneNumberStatusForPublicResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberStatusForPublicResponse) SetStatusCode(v int32) *PhoneNumberStatusForPublicResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberStatusForPublicResponse) SetBody(v *PhoneNumberStatusForPublicResponseBody) *PhoneNumberStatusForPublicResponse {
	s.Body = v
	return s
}

type PhoneNumberStatusForRealRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberStatusForRealRequest) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForRealRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForRealRequest) SetAuthCode(v string) *PhoneNumberStatusForRealRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberStatusForRealRequest) SetInputNumber(v string) *PhoneNumberStatusForRealRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberStatusForRealRequest) SetMask(v string) *PhoneNumberStatusForRealRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberStatusForRealRequest) SetOwnerId(v int64) *PhoneNumberStatusForRealRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberStatusForRealRequest) SetResourceOwnerAccount(v string) *PhoneNumberStatusForRealRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberStatusForRealRequest) SetResourceOwnerId(v int64) *PhoneNumberStatusForRealRequest {
	s.ResourceOwnerId = &v
	return s
}

type PhoneNumberStatusForRealResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PhoneNumberStatusForRealResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberStatusForRealResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForRealResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForRealResponseBody) SetCode(v string) *PhoneNumberStatusForRealResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberStatusForRealResponseBody) SetData(v *PhoneNumberStatusForRealResponseBodyData) *PhoneNumberStatusForRealResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberStatusForRealResponseBody) SetMessage(v string) *PhoneNumberStatusForRealResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberStatusForRealResponseBody) SetRequestId(v string) *PhoneNumberStatusForRealResponseBody {
	s.RequestId = &v
	return s
}

type PhoneNumberStatusForRealResponseBodyData struct {
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PhoneNumberStatusForRealResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForRealResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForRealResponseBodyData) SetCarrier(v string) *PhoneNumberStatusForRealResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *PhoneNumberStatusForRealResponseBodyData) SetStatus(v string) *PhoneNumberStatusForRealResponseBodyData {
	s.Status = &v
	return s
}

type PhoneNumberStatusForRealResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PhoneNumberStatusForRealResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PhoneNumberStatusForRealResponse) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForRealResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForRealResponse) SetHeaders(v map[string]*string) *PhoneNumberStatusForRealResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberStatusForRealResponse) SetStatusCode(v int32) *PhoneNumberStatusForRealResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberStatusForRealResponse) SetBody(v *PhoneNumberStatusForRealResponseBody) *PhoneNumberStatusForRealResponse {
	s.Body = v
	return s
}

type PhoneNumberStatusForSmsRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberStatusForSmsRequest) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForSmsRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForSmsRequest) SetAuthCode(v string) *PhoneNumberStatusForSmsRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberStatusForSmsRequest) SetInputNumber(v string) *PhoneNumberStatusForSmsRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberStatusForSmsRequest) SetMask(v string) *PhoneNumberStatusForSmsRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberStatusForSmsRequest) SetOwnerId(v int64) *PhoneNumberStatusForSmsRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberStatusForSmsRequest) SetResourceOwnerAccount(v string) *PhoneNumberStatusForSmsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberStatusForSmsRequest) SetResourceOwnerId(v int64) *PhoneNumberStatusForSmsRequest {
	s.ResourceOwnerId = &v
	return s
}

type PhoneNumberStatusForSmsResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PhoneNumberStatusForSmsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberStatusForSmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForSmsResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForSmsResponseBody) SetCode(v string) *PhoneNumberStatusForSmsResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberStatusForSmsResponseBody) SetData(v *PhoneNumberStatusForSmsResponseBodyData) *PhoneNumberStatusForSmsResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberStatusForSmsResponseBody) SetMessage(v string) *PhoneNumberStatusForSmsResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberStatusForSmsResponseBody) SetRequestId(v string) *PhoneNumberStatusForSmsResponseBody {
	s.RequestId = &v
	return s
}

type PhoneNumberStatusForSmsResponseBodyData struct {
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PhoneNumberStatusForSmsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForSmsResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForSmsResponseBodyData) SetCarrier(v string) *PhoneNumberStatusForSmsResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *PhoneNumberStatusForSmsResponseBodyData) SetStatus(v string) *PhoneNumberStatusForSmsResponseBodyData {
	s.Status = &v
	return s
}

type PhoneNumberStatusForSmsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PhoneNumberStatusForSmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PhoneNumberStatusForSmsResponse) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForSmsResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForSmsResponse) SetHeaders(v map[string]*string) *PhoneNumberStatusForSmsResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberStatusForSmsResponse) SetStatusCode(v int32) *PhoneNumberStatusForSmsResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberStatusForSmsResponse) SetBody(v *PhoneNumberStatusForSmsResponseBody) *PhoneNumberStatusForSmsResponse {
	s.Body = v
	return s
}

type PhoneNumberStatusForVirtualRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberStatusForVirtualRequest) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForVirtualRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVirtualRequest) SetAuthCode(v string) *PhoneNumberStatusForVirtualRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberStatusForVirtualRequest) SetInputNumber(v string) *PhoneNumberStatusForVirtualRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberStatusForVirtualRequest) SetMask(v string) *PhoneNumberStatusForVirtualRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberStatusForVirtualRequest) SetOwnerId(v int64) *PhoneNumberStatusForVirtualRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberStatusForVirtualRequest) SetResourceOwnerAccount(v string) *PhoneNumberStatusForVirtualRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberStatusForVirtualRequest) SetResourceOwnerId(v int64) *PhoneNumberStatusForVirtualRequest {
	s.ResourceOwnerId = &v
	return s
}

type PhoneNumberStatusForVirtualResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PhoneNumberStatusForVirtualResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberStatusForVirtualResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForVirtualResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVirtualResponseBody) SetCode(v string) *PhoneNumberStatusForVirtualResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberStatusForVirtualResponseBody) SetData(v *PhoneNumberStatusForVirtualResponseBodyData) *PhoneNumberStatusForVirtualResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberStatusForVirtualResponseBody) SetMessage(v string) *PhoneNumberStatusForVirtualResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberStatusForVirtualResponseBody) SetRequestId(v string) *PhoneNumberStatusForVirtualResponseBody {
	s.RequestId = &v
	return s
}

type PhoneNumberStatusForVirtualResponseBodyData struct {
	IsPrivacyNumber *bool `json:"IsPrivacyNumber,omitempty" xml:"IsPrivacyNumber,omitempty"`
}

func (s PhoneNumberStatusForVirtualResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForVirtualResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVirtualResponseBodyData) SetIsPrivacyNumber(v bool) *PhoneNumberStatusForVirtualResponseBodyData {
	s.IsPrivacyNumber = &v
	return s
}

type PhoneNumberStatusForVirtualResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PhoneNumberStatusForVirtualResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PhoneNumberStatusForVirtualResponse) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForVirtualResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVirtualResponse) SetHeaders(v map[string]*string) *PhoneNumberStatusForVirtualResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberStatusForVirtualResponse) SetStatusCode(v int32) *PhoneNumberStatusForVirtualResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberStatusForVirtualResponse) SetBody(v *PhoneNumberStatusForVirtualResponseBody) *PhoneNumberStatusForVirtualResponse {
	s.Body = v
	return s
}

type PhoneNumberStatusForVoiceRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberStatusForVoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForVoiceRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVoiceRequest) SetAuthCode(v string) *PhoneNumberStatusForVoiceRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberStatusForVoiceRequest) SetInputNumber(v string) *PhoneNumberStatusForVoiceRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberStatusForVoiceRequest) SetMask(v string) *PhoneNumberStatusForVoiceRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberStatusForVoiceRequest) SetOwnerId(v int64) *PhoneNumberStatusForVoiceRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberStatusForVoiceRequest) SetResourceOwnerAccount(v string) *PhoneNumberStatusForVoiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberStatusForVoiceRequest) SetResourceOwnerId(v int64) *PhoneNumberStatusForVoiceRequest {
	s.ResourceOwnerId = &v
	return s
}

type PhoneNumberStatusForVoiceResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PhoneNumberStatusForVoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberStatusForVoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVoiceResponseBody) SetCode(v string) *PhoneNumberStatusForVoiceResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberStatusForVoiceResponseBody) SetData(v *PhoneNumberStatusForVoiceResponseBodyData) *PhoneNumberStatusForVoiceResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberStatusForVoiceResponseBody) SetMessage(v string) *PhoneNumberStatusForVoiceResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberStatusForVoiceResponseBody) SetRequestId(v string) *PhoneNumberStatusForVoiceResponseBody {
	s.RequestId = &v
	return s
}

type PhoneNumberStatusForVoiceResponseBodyData struct {
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PhoneNumberStatusForVoiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForVoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVoiceResponseBodyData) SetCarrier(v string) *PhoneNumberStatusForVoiceResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *PhoneNumberStatusForVoiceResponseBodyData) SetStatus(v string) *PhoneNumberStatusForVoiceResponseBodyData {
	s.Status = &v
	return s
}

type PhoneNumberStatusForVoiceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PhoneNumberStatusForVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PhoneNumberStatusForVoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s PhoneNumberStatusForVoiceResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVoiceResponse) SetHeaders(v map[string]*string) *PhoneNumberStatusForVoiceResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberStatusForVoiceResponse) SetStatusCode(v int32) *PhoneNumberStatusForVoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberStatusForVoiceResponse) SetBody(v *PhoneNumberStatusForVoiceResponseBody) *PhoneNumberStatusForVoiceResponse {
	s.Body = v
	return s
}

type QueryAvailableAuthCodeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 标签id
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s QueryAvailableAuthCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAvailableAuthCodeRequest) GoString() string {
	return s.String()
}

func (s *QueryAvailableAuthCodeRequest) SetOwnerId(v int64) *QueryAvailableAuthCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAvailableAuthCodeRequest) SetResourceOwnerAccount(v string) *QueryAvailableAuthCodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryAvailableAuthCodeRequest) SetResourceOwnerId(v int64) *QueryAvailableAuthCodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryAvailableAuthCodeRequest) SetTagId(v int64) *QueryAvailableAuthCodeRequest {
	s.TagId = &v
	return s
}

type QueryAvailableAuthCodeResponseBody struct {
	Code      *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAvailableAuthCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAvailableAuthCodeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAvailableAuthCodeResponseBody) SetCode(v string) *QueryAvailableAuthCodeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAvailableAuthCodeResponseBody) SetData(v []*string) *QueryAvailableAuthCodeResponseBody {
	s.Data = v
	return s
}

func (s *QueryAvailableAuthCodeResponseBody) SetMessage(v string) *QueryAvailableAuthCodeResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAvailableAuthCodeResponseBody) SetRequestId(v string) *QueryAvailableAuthCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAvailableAuthCodeResponseBody) SetSuccess(v bool) *QueryAvailableAuthCodeResponseBody {
	s.Success = &v
	return s
}

type QueryAvailableAuthCodeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAvailableAuthCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAvailableAuthCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAvailableAuthCodeResponse) GoString() string {
	return s.String()
}

func (s *QueryAvailableAuthCodeResponse) SetHeaders(v map[string]*string) *QueryAvailableAuthCodeResponse {
	s.Headers = v
	return s
}

func (s *QueryAvailableAuthCodeResponse) SetStatusCode(v int32) *QueryAvailableAuthCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAvailableAuthCodeResponse) SetBody(v *QueryAvailableAuthCodeResponseBody) *QueryAvailableAuthCodeResponse {
	s.Body = v
	return s
}

type QueryTagApplyRuleRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 标签id
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s QueryTagApplyRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTagApplyRuleRequest) GoString() string {
	return s.String()
}

func (s *QueryTagApplyRuleRequest) SetOwnerId(v int64) *QueryTagApplyRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTagApplyRuleRequest) SetResourceOwnerAccount(v string) *QueryTagApplyRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTagApplyRuleRequest) SetResourceOwnerId(v int64) *QueryTagApplyRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryTagApplyRuleRequest) SetTagId(v int64) *QueryTagApplyRuleRequest {
	s.TagId = &v
	return s
}

type QueryTagApplyRuleResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryTagApplyRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTagApplyRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTagApplyRuleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTagApplyRuleResponseBody) SetCode(v string) *QueryTagApplyRuleResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTagApplyRuleResponseBody) SetData(v *QueryTagApplyRuleResponseBodyData) *QueryTagApplyRuleResponseBody {
	s.Data = v
	return s
}

func (s *QueryTagApplyRuleResponseBody) SetMessage(v string) *QueryTagApplyRuleResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTagApplyRuleResponseBody) SetRequestId(v string) *QueryTagApplyRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTagApplyRuleResponseBody) SetSuccess(v bool) *QueryTagApplyRuleResponseBody {
	s.Success = &v
	return s
}

type QueryTagApplyRuleResponseBodyData struct {
	// 申请材料要求
	ApplyMaterialDesc *string `json:"ApplyMaterialDesc,omitempty" xml:"ApplyMaterialDesc,omitempty"`
	// 是否自动审批
	AutoAudit *int64 `json:"AutoAudit,omitempty" xml:"AutoAudit,omitempty"`
	// 计费标准说明链接
	ChargingStandardLink *string `json:"ChargingStandardLink,omitempty" xml:"ChargingStandardLink,omitempty"`
	// 是否支持加密查询
	EncryptedQuery *int64 `json:"EncryptedQuery,omitempty" xml:"EncryptedQuery,omitempty"`
	// 是否需要提供申请材料
	NeedApplyMaterial *int64 `json:"NeedApplyMaterial,omitempty" xml:"NeedApplyMaterial,omitempty"`
	// 服务协议链接
	SlaLink *string `json:"SlaLink,omitempty" xml:"SlaLink,omitempty"`
}

func (s QueryTagApplyRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTagApplyRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTagApplyRuleResponseBodyData) SetApplyMaterialDesc(v string) *QueryTagApplyRuleResponseBodyData {
	s.ApplyMaterialDesc = &v
	return s
}

func (s *QueryTagApplyRuleResponseBodyData) SetAutoAudit(v int64) *QueryTagApplyRuleResponseBodyData {
	s.AutoAudit = &v
	return s
}

func (s *QueryTagApplyRuleResponseBodyData) SetChargingStandardLink(v string) *QueryTagApplyRuleResponseBodyData {
	s.ChargingStandardLink = &v
	return s
}

func (s *QueryTagApplyRuleResponseBodyData) SetEncryptedQuery(v int64) *QueryTagApplyRuleResponseBodyData {
	s.EncryptedQuery = &v
	return s
}

func (s *QueryTagApplyRuleResponseBodyData) SetNeedApplyMaterial(v int64) *QueryTagApplyRuleResponseBodyData {
	s.NeedApplyMaterial = &v
	return s
}

func (s *QueryTagApplyRuleResponseBodyData) SetSlaLink(v string) *QueryTagApplyRuleResponseBodyData {
	s.SlaLink = &v
	return s
}

type QueryTagApplyRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTagApplyRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTagApplyRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTagApplyRuleResponse) GoString() string {
	return s.String()
}

func (s *QueryTagApplyRuleResponse) SetHeaders(v map[string]*string) *QueryTagApplyRuleResponse {
	s.Headers = v
	return s
}

func (s *QueryTagApplyRuleResponse) SetStatusCode(v int32) *QueryTagApplyRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTagApplyRuleResponse) SetBody(v *QueryTagApplyRuleResponseBody) *QueryTagApplyRuleResponse {
	s.Body = v
	return s
}

type QueryTagInfoBySelectionRequest struct {
	// 行业id
	IndustryId           *int64  `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 场景id
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// 标签id
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s QueryTagInfoBySelectionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTagInfoBySelectionRequest) GoString() string {
	return s.String()
}

func (s *QueryTagInfoBySelectionRequest) SetIndustryId(v int64) *QueryTagInfoBySelectionRequest {
	s.IndustryId = &v
	return s
}

func (s *QueryTagInfoBySelectionRequest) SetOwnerId(v int64) *QueryTagInfoBySelectionRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTagInfoBySelectionRequest) SetResourceOwnerAccount(v string) *QueryTagInfoBySelectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTagInfoBySelectionRequest) SetResourceOwnerId(v int64) *QueryTagInfoBySelectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryTagInfoBySelectionRequest) SetSceneId(v int64) *QueryTagInfoBySelectionRequest {
	s.SceneId = &v
	return s
}

func (s *QueryTagInfoBySelectionRequest) SetTagId(v int64) *QueryTagInfoBySelectionRequest {
	s.TagId = &v
	return s
}

type QueryTagInfoBySelectionResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*QueryTagInfoBySelectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTagInfoBySelectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTagInfoBySelectionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTagInfoBySelectionResponseBody) SetCode(v string) *QueryTagInfoBySelectionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBody) SetData(v []*QueryTagInfoBySelectionResponseBodyData) *QueryTagInfoBySelectionResponseBody {
	s.Data = v
	return s
}

func (s *QueryTagInfoBySelectionResponseBody) SetMessage(v string) *QueryTagInfoBySelectionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBody) SetRequestId(v string) *QueryTagInfoBySelectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBody) SetSuccess(v bool) *QueryTagInfoBySelectionResponseBody {
	s.Success = &v
	return s
}

type QueryTagInfoBySelectionResponseBodyData struct {
	// 可用的授权码列表
	AuthCodeList []*string `json:"AuthCodeList,omitempty" xml:"AuthCodeList,omitempty" type:"Repeated"`
	// API demo链接
	DemoAddress *string `json:"DemoAddress,omitempty" xml:"DemoAddress,omitempty"`
	// API文档链接
	DocAddress *string `json:"DocAddress,omitempty" xml:"DocAddress,omitempty"`
	// 枚举值定义链接
	EnumDefinitionAddress *string `json:"EnumDefinitionAddress,omitempty" xml:"EnumDefinitionAddress,omitempty"`
	// 流程名称
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// 行业id
	IndustryId *int64 `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	// 行业名称
	IndustryName *string `json:"IndustryName,omitempty" xml:"IndustryName,omitempty"`
	// 标签参数列表
	ParamList []*QueryTagInfoBySelectionResponseBodyDataParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// 场景id
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// 场景名称
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// 标签id
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// 标签名称
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s QueryTagInfoBySelectionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTagInfoBySelectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetAuthCodeList(v []*string) *QueryTagInfoBySelectionResponseBodyData {
	s.AuthCodeList = v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetDemoAddress(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.DemoAddress = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetDocAddress(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.DocAddress = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetEnumDefinitionAddress(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.EnumDefinitionAddress = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetFlowName(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.FlowName = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetIndustryId(v int64) *QueryTagInfoBySelectionResponseBodyData {
	s.IndustryId = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetIndustryName(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.IndustryName = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetParamList(v []*QueryTagInfoBySelectionResponseBodyDataParamList) *QueryTagInfoBySelectionResponseBodyData {
	s.ParamList = v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetSceneId(v int64) *QueryTagInfoBySelectionResponseBodyData {
	s.SceneId = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetSceneName(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.SceneName = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetTagId(v int64) *QueryTagInfoBySelectionResponseBodyData {
	s.TagId = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyData) SetTagName(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.TagName = &v
	return s
}

type QueryTagInfoBySelectionResponseBodyDataParamList struct {
	// 参数英文名
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 输入提示
	Hint *string `json:"Hint,omitempty" xml:"Hint,omitempty"`
	// 是否必填
	Must *bool `json:"Must,omitempty" xml:"Must,omitempty"`
	// 参数中文名
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 类型EnumUIWidgetTypes对应的code
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 枚举值定义，code:desc
	ValueDict []*QueryTagInfoBySelectionResponseBodyDataParamListValueDict `json:"ValueDict,omitempty" xml:"ValueDict,omitempty" type:"Repeated"`
}

func (s QueryTagInfoBySelectionResponseBodyDataParamList) String() string {
	return tea.Prettify(s)
}

func (s QueryTagInfoBySelectionResponseBodyDataParamList) GoString() string {
	return s.String()
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) SetCode(v string) *QueryTagInfoBySelectionResponseBodyDataParamList {
	s.Code = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) SetHint(v string) *QueryTagInfoBySelectionResponseBodyDataParamList {
	s.Hint = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) SetMust(v bool) *QueryTagInfoBySelectionResponseBodyDataParamList {
	s.Must = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) SetName(v string) *QueryTagInfoBySelectionResponseBodyDataParamList {
	s.Name = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) SetType(v string) *QueryTagInfoBySelectionResponseBodyDataParamList {
	s.Type = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamList) SetValueDict(v []*QueryTagInfoBySelectionResponseBodyDataParamListValueDict) *QueryTagInfoBySelectionResponseBodyDataParamList {
	s.ValueDict = v
	return s
}

type QueryTagInfoBySelectionResponseBodyDataParamListValueDict struct {
	// 英文名
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 中文名
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s QueryTagInfoBySelectionResponseBodyDataParamListValueDict) String() string {
	return tea.Prettify(s)
}

func (s QueryTagInfoBySelectionResponseBodyDataParamListValueDict) GoString() string {
	return s.String()
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamListValueDict) SetCode(v string) *QueryTagInfoBySelectionResponseBodyDataParamListValueDict {
	s.Code = &v
	return s
}

func (s *QueryTagInfoBySelectionResponseBodyDataParamListValueDict) SetDesc(v string) *QueryTagInfoBySelectionResponseBodyDataParamListValueDict {
	s.Desc = &v
	return s
}

type QueryTagInfoBySelectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTagInfoBySelectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTagInfoBySelectionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTagInfoBySelectionResponse) GoString() string {
	return s.String()
}

func (s *QueryTagInfoBySelectionResponse) SetHeaders(v map[string]*string) *QueryTagInfoBySelectionResponse {
	s.Headers = v
	return s
}

func (s *QueryTagInfoBySelectionResponse) SetStatusCode(v int32) *QueryTagInfoBySelectionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTagInfoBySelectionResponse) SetBody(v *QueryTagInfoBySelectionResponseBody) *QueryTagInfoBySelectionResponse {
	s.Body = v
	return s
}

type QueryTagListPageRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryTagListPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTagListPageRequest) GoString() string {
	return s.String()
}

func (s *QueryTagListPageRequest) SetOwnerId(v int64) *QueryTagListPageRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTagListPageRequest) SetPageNo(v int64) *QueryTagListPageRequest {
	s.PageNo = &v
	return s
}

func (s *QueryTagListPageRequest) SetPageSize(v int64) *QueryTagListPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTagListPageRequest) SetResourceOwnerAccount(v string) *QueryTagListPageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTagListPageRequest) SetResourceOwnerId(v int64) *QueryTagListPageRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryTagListPageResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryTagListPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTagListPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTagListPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTagListPageResponseBody) SetCode(v string) *QueryTagListPageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTagListPageResponseBody) SetData(v *QueryTagListPageResponseBodyData) *QueryTagListPageResponseBody {
	s.Data = v
	return s
}

func (s *QueryTagListPageResponseBody) SetMessage(v string) *QueryTagListPageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTagListPageResponseBody) SetRequestId(v string) *QueryTagListPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTagListPageResponseBody) SetSuccess(v bool) *QueryTagListPageResponseBody {
	s.Success = &v
	return s
}

type QueryTagListPageResponseBodyData struct {
	PageNo     *int64                                     `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int64                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records    []*QueryTagListPageResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	TotalCount *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int64                                     `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s QueryTagListPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTagListPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTagListPageResponseBodyData) SetPageNo(v int64) *QueryTagListPageResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryTagListPageResponseBodyData) SetPageSize(v int64) *QueryTagListPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryTagListPageResponseBodyData) SetRecords(v []*QueryTagListPageResponseBodyDataRecords) *QueryTagListPageResponseBodyData {
	s.Records = v
	return s
}

func (s *QueryTagListPageResponseBodyData) SetTotalCount(v int64) *QueryTagListPageResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryTagListPageResponseBodyData) SetTotalPage(v int64) *QueryTagListPageResponseBodyData {
	s.TotalPage = &v
	return s
}

type QueryTagListPageResponseBodyDataRecords struct {
	// 前端调用的api名称
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// API文档链接
	DocAddress *string `json:"DocAddress,omitempty" xml:"DocAddress,omitempty"`
	// 标签 id
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 行业id
	IndustryId *int64 `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	// 行业名称
	IndustryName *string `json:"IndustryName,omitempty" xml:"IndustryName,omitempty"`
	// 标签介绍
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// 是否已经申请开通
	IsOpen *int64 `json:"IsOpen,omitempty" xml:"IsOpen,omitempty"`
	// 标签名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 0 隐藏 1 公开
	SaleStatusStr *string `json:"SaleStatusStr,omitempty" xml:"SaleStatusStr,omitempty"`
	// 场景id
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// 场景名称
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s QueryTagListPageResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s QueryTagListPageResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *QueryTagListPageResponseBodyDataRecords) SetApiName(v string) *QueryTagListPageResponseBodyDataRecords {
	s.ApiName = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetCode(v string) *QueryTagListPageResponseBodyDataRecords {
	s.Code = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetDocAddress(v string) *QueryTagListPageResponseBodyDataRecords {
	s.DocAddress = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetId(v int64) *QueryTagListPageResponseBodyDataRecords {
	s.Id = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetIndustryId(v int64) *QueryTagListPageResponseBodyDataRecords {
	s.IndustryId = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetIndustryName(v string) *QueryTagListPageResponseBodyDataRecords {
	s.IndustryName = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetIntroduction(v string) *QueryTagListPageResponseBodyDataRecords {
	s.Introduction = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetIsOpen(v int64) *QueryTagListPageResponseBodyDataRecords {
	s.IsOpen = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetName(v string) *QueryTagListPageResponseBodyDataRecords {
	s.Name = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetSaleStatusStr(v string) *QueryTagListPageResponseBodyDataRecords {
	s.SaleStatusStr = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetSceneId(v int64) *QueryTagListPageResponseBodyDataRecords {
	s.SceneId = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetSceneName(v string) *QueryTagListPageResponseBodyDataRecords {
	s.SceneName = &v
	return s
}

type QueryTagListPageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTagListPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTagListPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTagListPageResponse) GoString() string {
	return s.String()
}

func (s *QueryTagListPageResponse) SetHeaders(v map[string]*string) *QueryTagListPageResponse {
	s.Headers = v
	return s
}

func (s *QueryTagListPageResponse) SetStatusCode(v int32) *QueryTagListPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTagListPageResponse) SetBody(v *QueryTagListPageResponseBody) *QueryTagListPageResponse {
	s.Body = v
	return s
}

type QueryUsageStatisticsByTagIdRequest struct {
	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// 结束时间
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 结束时间
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s QueryUsageStatisticsByTagIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageStatisticsByTagIdRequest) GoString() string {
	return s.String()
}

func (s *QueryUsageStatisticsByTagIdRequest) SetBeginTime(v string) *QueryUsageStatisticsByTagIdRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdRequest) SetEndTime(v string) *QueryUsageStatisticsByTagIdRequest {
	s.EndTime = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdRequest) SetOwnerId(v int64) *QueryUsageStatisticsByTagIdRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdRequest) SetPageNo(v int64) *QueryUsageStatisticsByTagIdRequest {
	s.PageNo = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdRequest) SetPageSize(v int64) *QueryUsageStatisticsByTagIdRequest {
	s.PageSize = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdRequest) SetResourceOwnerAccount(v string) *QueryUsageStatisticsByTagIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdRequest) SetResourceOwnerId(v int64) *QueryUsageStatisticsByTagIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdRequest) SetTagId(v int64) *QueryUsageStatisticsByTagIdRequest {
	s.TagId = &v
	return s
}

type QueryUsageStatisticsByTagIdResponseBody struct {
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*QueryUsageStatisticsByTagIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUsageStatisticsByTagIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageStatisticsByTagIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUsageStatisticsByTagIdResponseBody) SetCode(v string) *QueryUsageStatisticsByTagIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBody) SetData(v []*QueryUsageStatisticsByTagIdResponseBodyData) *QueryUsageStatisticsByTagIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBody) SetMessage(v string) *QueryUsageStatisticsByTagIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBody) SetRequestId(v string) *QueryUsageStatisticsByTagIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBody) SetSuccess(v bool) *QueryUsageStatisticsByTagIdResponseBody {
	s.Success = &v
	return s
}

type QueryUsageStatisticsByTagIdResponseBodyData struct {
	// 授权码
	AuthorizationCode *string `json:"AuthorizationCode,omitempty" xml:"AuthorizationCode,omitempty"`
	// 查询失败号码数
	FailTotal *int64 `json:"FailTotal,omitempty" xml:"FailTotal,omitempty"`
	// 创建时间
	GmtDateStr *string `json:"GmtDateStr,omitempty" xml:"GmtDateStr,omitempty"`
	// 授权码使用记录 id
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 行业名称
	IndustryName *string `json:"IndustryName,omitempty" xml:"IndustryName,omitempty"`
	// 客户 pid
	PartnerId *int64 `json:"PartnerId,omitempty" xml:"PartnerId,omitempty"`
	// 场景名称
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// 查询成功号码数
	SuccessTotal *int64 `json:"SuccessTotal,omitempty" xml:"SuccessTotal,omitempty"`
	// 标签名称
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// 标签名称
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// 查询总号码数
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryUsageStatisticsByTagIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageStatisticsByTagIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetAuthorizationCode(v string) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.AuthorizationCode = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetFailTotal(v int64) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.FailTotal = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetGmtDateStr(v string) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.GmtDateStr = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetId(v int64) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetIndustryName(v string) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.IndustryName = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetPartnerId(v int64) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.PartnerId = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetSceneName(v string) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.SceneName = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetSuccessTotal(v int64) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.SuccessTotal = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetTagId(v int64) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.TagId = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetTagName(v string) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.TagName = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponseBodyData) SetTotal(v int64) *QueryUsageStatisticsByTagIdResponseBodyData {
	s.Total = &v
	return s
}

type QueryUsageStatisticsByTagIdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUsageStatisticsByTagIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUsageStatisticsByTagIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageStatisticsByTagIdResponse) GoString() string {
	return s.String()
}

func (s *QueryUsageStatisticsByTagIdResponse) SetHeaders(v map[string]*string) *QueryUsageStatisticsByTagIdResponse {
	s.Headers = v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponse) SetStatusCode(v int32) *QueryUsageStatisticsByTagIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdResponse) SetBody(v *QueryUsageStatisticsByTagIdResponseBody) *QueryUsageStatisticsByTagIdResponse {
	s.Body = v
	return s
}

type ThreeElementsVerificationRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	CertCode             *string `json:"CertCode,omitempty" xml:"CertCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ThreeElementsVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s ThreeElementsVerificationRequest) GoString() string {
	return s.String()
}

func (s *ThreeElementsVerificationRequest) SetAuthCode(v string) *ThreeElementsVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetCertCode(v string) *ThreeElementsVerificationRequest {
	s.CertCode = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetInputNumber(v string) *ThreeElementsVerificationRequest {
	s.InputNumber = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetMask(v string) *ThreeElementsVerificationRequest {
	s.Mask = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetName(v string) *ThreeElementsVerificationRequest {
	s.Name = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetOwnerId(v int64) *ThreeElementsVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetResourceOwnerAccount(v string) *ThreeElementsVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetResourceOwnerId(v int64) *ThreeElementsVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

type ThreeElementsVerificationResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ThreeElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ThreeElementsVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ThreeElementsVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *ThreeElementsVerificationResponseBody) SetCode(v string) *ThreeElementsVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *ThreeElementsVerificationResponseBody) SetData(v *ThreeElementsVerificationResponseBodyData) *ThreeElementsVerificationResponseBody {
	s.Data = v
	return s
}

func (s *ThreeElementsVerificationResponseBody) SetMessage(v string) *ThreeElementsVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *ThreeElementsVerificationResponseBody) SetRequestId(v string) *ThreeElementsVerificationResponseBody {
	s.RequestId = &v
	return s
}

type ThreeElementsVerificationResponseBodyData struct {
	BasicCarrier *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	IsConsistent *int32  `json:"IsConsistent,omitempty" xml:"IsConsistent,omitempty"`
}

func (s ThreeElementsVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ThreeElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ThreeElementsVerificationResponseBodyData) SetBasicCarrier(v string) *ThreeElementsVerificationResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *ThreeElementsVerificationResponseBodyData) SetIsConsistent(v int32) *ThreeElementsVerificationResponseBodyData {
	s.IsConsistent = &v
	return s
}

type ThreeElementsVerificationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ThreeElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ThreeElementsVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s ThreeElementsVerificationResponse) GoString() string {
	return s.String()
}

func (s *ThreeElementsVerificationResponse) SetHeaders(v map[string]*string) *ThreeElementsVerificationResponse {
	s.Headers = v
	return s
}

func (s *ThreeElementsVerificationResponse) SetStatusCode(v int32) *ThreeElementsVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *ThreeElementsVerificationResponse) SetBody(v *ThreeElementsVerificationResponseBody) *ThreeElementsVerificationResponse {
	s.Body = v
	return s
}

type TwoElementsVerificationRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s TwoElementsVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s TwoElementsVerificationRequest) GoString() string {
	return s.String()
}

func (s *TwoElementsVerificationRequest) SetAuthCode(v string) *TwoElementsVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetInputNumber(v string) *TwoElementsVerificationRequest {
	s.InputNumber = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetMask(v string) *TwoElementsVerificationRequest {
	s.Mask = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetName(v string) *TwoElementsVerificationRequest {
	s.Name = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetOwnerId(v int64) *TwoElementsVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetResourceOwnerAccount(v string) *TwoElementsVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetResourceOwnerId(v int64) *TwoElementsVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

type TwoElementsVerificationResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TwoElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TwoElementsVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TwoElementsVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *TwoElementsVerificationResponseBody) SetCode(v string) *TwoElementsVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *TwoElementsVerificationResponseBody) SetData(v *TwoElementsVerificationResponseBodyData) *TwoElementsVerificationResponseBody {
	s.Data = v
	return s
}

func (s *TwoElementsVerificationResponseBody) SetMessage(v string) *TwoElementsVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *TwoElementsVerificationResponseBody) SetRequestId(v string) *TwoElementsVerificationResponseBody {
	s.RequestId = &v
	return s
}

type TwoElementsVerificationResponseBodyData struct {
	BasicCarrier *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	IsConsistent *int32  `json:"IsConsistent,omitempty" xml:"IsConsistent,omitempty"`
}

func (s TwoElementsVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TwoElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *TwoElementsVerificationResponseBodyData) SetBasicCarrier(v string) *TwoElementsVerificationResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *TwoElementsVerificationResponseBodyData) SetIsConsistent(v int32) *TwoElementsVerificationResponseBodyData {
	s.IsConsistent = &v
	return s
}

type TwoElementsVerificationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TwoElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TwoElementsVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s TwoElementsVerificationResponse) GoString() string {
	return s.String()
}

func (s *TwoElementsVerificationResponse) SetHeaders(v map[string]*string) *TwoElementsVerificationResponse {
	s.Headers = v
	return s
}

func (s *TwoElementsVerificationResponse) SetStatusCode(v int32) *TwoElementsVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *TwoElementsVerificationResponse) SetBody(v *TwoElementsVerificationResponseBody) *TwoElementsVerificationResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dytnsapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CompanyFourElementsVerificationWithOptions(request *CompanyFourElementsVerificationRequest, runtime *util.RuntimeOptions) (_result *CompanyFourElementsVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.EpCertName)) {
		query["EpCertName"] = request.EpCertName
	}

	if !tea.BoolValue(util.IsUnset(request.EpCertNo)) {
		query["EpCertNo"] = request.EpCertNo
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonCertName)) {
		query["LegalPersonCertName"] = request.LegalPersonCertName
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonCertNo)) {
		query["LegalPersonCertNo"] = request.LegalPersonCertNo
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CompanyFourElementsVerification"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CompanyFourElementsVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompanyFourElementsVerification(request *CompanyFourElementsVerificationRequest) (_result *CompanyFourElementsVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompanyFourElementsVerificationResponse{}
	_body, _err := client.CompanyFourElementsVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompanyThreeElementsVerificationWithOptions(request *CompanyThreeElementsVerificationRequest, runtime *util.RuntimeOptions) (_result *CompanyThreeElementsVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.EpCertName)) {
		query["EpCertName"] = request.EpCertName
	}

	if !tea.BoolValue(util.IsUnset(request.EpCertNo)) {
		query["EpCertNo"] = request.EpCertNo
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonCertName)) {
		query["LegalPersonCertName"] = request.LegalPersonCertName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CompanyThreeElementsVerification"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CompanyThreeElementsVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompanyThreeElementsVerification(request *CompanyThreeElementsVerificationRequest) (_result *CompanyThreeElementsVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompanyThreeElementsVerificationResponse{}
	_body, _err := client.CompanyThreeElementsVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompanyTwoElementsVerificationWithOptions(request *CompanyTwoElementsVerificationRequest, runtime *util.RuntimeOptions) (_result *CompanyTwoElementsVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.EpCertName)) {
		query["EpCertName"] = request.EpCertName
	}

	if !tea.BoolValue(util.IsUnset(request.EpCertNo)) {
		query["EpCertNo"] = request.EpCertNo
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CompanyTwoElementsVerification"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CompanyTwoElementsVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompanyTwoElementsVerification(request *CompanyTwoElementsVerificationRequest) (_result *CompanyTwoElementsVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompanyTwoElementsVerificationResponse{}
	_body, _err := client.CompanyTwoElementsVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEmptyNumberWithOptions(request *DescribeEmptyNumberRequest, runtime *util.RuntimeOptions) (_result *DescribeEmptyNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEmptyNumber"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEmptyNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEmptyNumber(request *DescribeEmptyNumberRequest) (_result *DescribeEmptyNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEmptyNumberResponse{}
	_body, _err := client.DescribeEmptyNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneNumberAnalysisWithOptions(request *DescribePhoneNumberAnalysisRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.NumberType)) {
		query["NumberType"] = request.NumberType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Rate)) {
		query["Rate"] = request.Rate
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberAnalysis"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneNumberAnalysis(request *DescribePhoneNumberAnalysisRequest) (_result *DescribePhoneNumberAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberAnalysisResponse{}
	_body, _err := client.DescribePhoneNumberAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneNumberAnalysisAIWithOptions(request *DescribePhoneNumberAnalysisAIRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberAnalysisAIResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ModelConfig)) {
		query["ModelConfig"] = request.ModelConfig
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Rate)) {
		query["Rate"] = request.Rate
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberAnalysisAI"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberAnalysisAIResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneNumberAnalysisAI(request *DescribePhoneNumberAnalysisAIRequest) (_result *DescribePhoneNumberAnalysisAIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberAnalysisAIResponse{}
	_body, _err := client.DescribePhoneNumberAnalysisAIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneNumberAttributeWithOptions(request *DescribePhoneNumberAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberAttribute"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneNumberAttribute(request *DescribePhoneNumberAttributeRequest) (_result *DescribePhoneNumberAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberAttributeResponse{}
	_body, _err := client.DescribePhoneNumberAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneNumberOnlineTimeWithOptions(request *DescribePhoneNumberOnlineTimeRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberOnlineTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.Carrier)) {
		query["Carrier"] = request.Carrier
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberOnlineTime"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberOnlineTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneNumberOnlineTime(request *DescribePhoneNumberOnlineTimeRequest) (_result *DescribePhoneNumberOnlineTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberOnlineTimeResponse{}
	_body, _err := client.DescribePhoneNumberOnlineTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneNumberOperatorAttributeWithOptions(request *DescribePhoneNumberOperatorAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberOperatorAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberOperatorAttribute"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberOperatorAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneNumberOperatorAttribute(request *DescribePhoneNumberOperatorAttributeRequest) (_result *DescribePhoneNumberOperatorAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberOperatorAttributeResponse{}
	_body, _err := client.DescribePhoneNumberOperatorAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneTwiceTelVerifyWithOptions(request *DescribePhoneTwiceTelVerifyRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneTwiceTelVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneTwiceTelVerify"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneTwiceTelVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneTwiceTelVerify(request *DescribePhoneTwiceTelVerifyRequest) (_result *DescribePhoneTwiceTelVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneTwiceTelVerifyResponse{}
	_body, _err := client.DescribePhoneTwiceTelVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvalidPhoneNumberFilterWithOptions(request *InvalidPhoneNumberFilterRequest, runtime *util.RuntimeOptions) (_result *InvalidPhoneNumberFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InvalidPhoneNumberFilter"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InvalidPhoneNumberFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvalidPhoneNumberFilter(request *InvalidPhoneNumberFilterRequest) (_result *InvalidPhoneNumberFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InvalidPhoneNumberFilterResponse{}
	_body, _err := client.InvalidPhoneNumberFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PhoneNumberConvertServiceWithOptions(request *PhoneNumberConvertServiceRequest, runtime *util.RuntimeOptions) (_result *PhoneNumberConvertServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PhoneNumberConvertService"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PhoneNumberConvertServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PhoneNumberConvertService(request *PhoneNumberConvertServiceRequest) (_result *PhoneNumberConvertServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PhoneNumberConvertServiceResponse{}
	_body, _err := client.PhoneNumberConvertServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PhoneNumberEncryptWithOptions(request *PhoneNumberEncryptRequest, runtime *util.RuntimeOptions) (_result *PhoneNumberEncryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PhoneNumberEncrypt"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PhoneNumberEncryptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PhoneNumberEncrypt(request *PhoneNumberEncryptRequest) (_result *PhoneNumberEncryptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PhoneNumberEncryptResponse{}
	_body, _err := client.PhoneNumberEncryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PhoneNumberStatusForAccountWithOptions(request *PhoneNumberStatusForAccountRequest, runtime *util.RuntimeOptions) (_result *PhoneNumberStatusForAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PhoneNumberStatusForAccount"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PhoneNumberStatusForAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PhoneNumberStatusForAccount(request *PhoneNumberStatusForAccountRequest) (_result *PhoneNumberStatusForAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PhoneNumberStatusForAccountResponse{}
	_body, _err := client.PhoneNumberStatusForAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PhoneNumberStatusForPublicWithOptions(request *PhoneNumberStatusForPublicRequest, runtime *util.RuntimeOptions) (_result *PhoneNumberStatusForPublicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PhoneNumberStatusForPublic"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PhoneNumberStatusForPublicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PhoneNumberStatusForPublic(request *PhoneNumberStatusForPublicRequest) (_result *PhoneNumberStatusForPublicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PhoneNumberStatusForPublicResponse{}
	_body, _err := client.PhoneNumberStatusForPublicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PhoneNumberStatusForRealWithOptions(request *PhoneNumberStatusForRealRequest, runtime *util.RuntimeOptions) (_result *PhoneNumberStatusForRealResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PhoneNumberStatusForReal"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PhoneNumberStatusForRealResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PhoneNumberStatusForReal(request *PhoneNumberStatusForRealRequest) (_result *PhoneNumberStatusForRealResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PhoneNumberStatusForRealResponse{}
	_body, _err := client.PhoneNumberStatusForRealWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PhoneNumberStatusForSmsWithOptions(request *PhoneNumberStatusForSmsRequest, runtime *util.RuntimeOptions) (_result *PhoneNumberStatusForSmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PhoneNumberStatusForSms"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PhoneNumberStatusForSmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PhoneNumberStatusForSms(request *PhoneNumberStatusForSmsRequest) (_result *PhoneNumberStatusForSmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PhoneNumberStatusForSmsResponse{}
	_body, _err := client.PhoneNumberStatusForSmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PhoneNumberStatusForVirtualWithOptions(request *PhoneNumberStatusForVirtualRequest, runtime *util.RuntimeOptions) (_result *PhoneNumberStatusForVirtualResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PhoneNumberStatusForVirtual"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PhoneNumberStatusForVirtualResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PhoneNumberStatusForVirtual(request *PhoneNumberStatusForVirtualRequest) (_result *PhoneNumberStatusForVirtualResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PhoneNumberStatusForVirtualResponse{}
	_body, _err := client.PhoneNumberStatusForVirtualWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PhoneNumberStatusForVoiceWithOptions(request *PhoneNumberStatusForVoiceRequest, runtime *util.RuntimeOptions) (_result *PhoneNumberStatusForVoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PhoneNumberStatusForVoice"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PhoneNumberStatusForVoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PhoneNumberStatusForVoice(request *PhoneNumberStatusForVoiceRequest) (_result *PhoneNumberStatusForVoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PhoneNumberStatusForVoiceResponse{}
	_body, _err := client.PhoneNumberStatusForVoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAvailableAuthCodeWithOptions(request *QueryAvailableAuthCodeRequest, runtime *util.RuntimeOptions) (_result *QueryAvailableAuthCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TagId)) {
		query["TagId"] = request.TagId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAvailableAuthCode"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAvailableAuthCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAvailableAuthCode(request *QueryAvailableAuthCodeRequest) (_result *QueryAvailableAuthCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAvailableAuthCodeResponse{}
	_body, _err := client.QueryAvailableAuthCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTagApplyRuleWithOptions(request *QueryTagApplyRuleRequest, runtime *util.RuntimeOptions) (_result *QueryTagApplyRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TagId)) {
		query["TagId"] = request.TagId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTagApplyRule"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTagApplyRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTagApplyRule(request *QueryTagApplyRuleRequest) (_result *QueryTagApplyRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTagApplyRuleResponse{}
	_body, _err := client.QueryTagApplyRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTagInfoBySelectionWithOptions(request *QueryTagInfoBySelectionRequest, runtime *util.RuntimeOptions) (_result *QueryTagInfoBySelectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IndustryId)) {
		query["IndustryId"] = request.IndustryId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.TagId)) {
		query["TagId"] = request.TagId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTagInfoBySelection"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTagInfoBySelectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTagInfoBySelection(request *QueryTagInfoBySelectionRequest) (_result *QueryTagInfoBySelectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTagInfoBySelectionResponse{}
	_body, _err := client.QueryTagInfoBySelectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTagListPageWithOptions(request *QueryTagListPageRequest, runtime *util.RuntimeOptions) (_result *QueryTagListPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTagListPage"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTagListPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTagListPage(request *QueryTagListPageRequest) (_result *QueryTagListPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTagListPageResponse{}
	_body, _err := client.QueryTagListPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUsageStatisticsByTagIdWithOptions(request *QueryUsageStatisticsByTagIdRequest, runtime *util.RuntimeOptions) (_result *QueryUsageStatisticsByTagIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TagId)) {
		query["TagId"] = request.TagId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryUsageStatisticsByTagId"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUsageStatisticsByTagIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUsageStatisticsByTagId(request *QueryUsageStatisticsByTagIdRequest) (_result *QueryUsageStatisticsByTagIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUsageStatisticsByTagIdResponse{}
	_body, _err := client.QueryUsageStatisticsByTagIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ThreeElementsVerificationWithOptions(request *ThreeElementsVerificationRequest, runtime *util.RuntimeOptions) (_result *ThreeElementsVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.CertCode)) {
		query["CertCode"] = request.CertCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ThreeElementsVerification"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ThreeElementsVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ThreeElementsVerification(request *ThreeElementsVerificationRequest) (_result *ThreeElementsVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ThreeElementsVerificationResponse{}
	_body, _err := client.ThreeElementsVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TwoElementsVerificationWithOptions(request *TwoElementsVerificationRequest, runtime *util.RuntimeOptions) (_result *TwoElementsVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.InputNumber)) {
		query["InputNumber"] = request.InputNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TwoElementsVerification"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TwoElementsVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TwoElementsVerification(request *TwoElementsVerificationRequest) (_result *TwoElementsVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TwoElementsVerificationResponse{}
	_body, _err := client.TwoElementsVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
