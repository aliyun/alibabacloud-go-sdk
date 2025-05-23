// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CertNoThreeElementVerificationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// QASDW@#**
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3***************0
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// iVBOFMKODOFNDFP123DFSMOO...
	CertPicture *string `json:"CertPicture,omitempty" xml:"CertPicture,omitempty"`
	// example:
	//
	// 示例值示例值
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CertNoThreeElementVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s CertNoThreeElementVerificationRequest) GoString() string {
	return s.String()
}

func (s *CertNoThreeElementVerificationRequest) SetAuthCode(v string) *CertNoThreeElementVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *CertNoThreeElementVerificationRequest) SetCertName(v string) *CertNoThreeElementVerificationRequest {
	s.CertName = &v
	return s
}

func (s *CertNoThreeElementVerificationRequest) SetCertNo(v string) *CertNoThreeElementVerificationRequest {
	s.CertNo = &v
	return s
}

func (s *CertNoThreeElementVerificationRequest) SetCertPicture(v string) *CertNoThreeElementVerificationRequest {
	s.CertPicture = &v
	return s
}

func (s *CertNoThreeElementVerificationRequest) SetMask(v string) *CertNoThreeElementVerificationRequest {
	s.Mask = &v
	return s
}

func (s *CertNoThreeElementVerificationRequest) SetOwnerId(v int64) *CertNoThreeElementVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *CertNoThreeElementVerificationRequest) SetResourceOwnerAccount(v string) *CertNoThreeElementVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CertNoThreeElementVerificationRequest) SetResourceOwnerId(v int64) *CertNoThreeElementVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

type CertNoThreeElementVerificationResponseBody struct {
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CertNoThreeElementVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD******177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CertNoThreeElementVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CertNoThreeElementVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *CertNoThreeElementVerificationResponseBody) SetAccessDeniedDetail(v string) *CertNoThreeElementVerificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CertNoThreeElementVerificationResponseBody) SetCode(v string) *CertNoThreeElementVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *CertNoThreeElementVerificationResponseBody) SetData(v *CertNoThreeElementVerificationResponseBodyData) *CertNoThreeElementVerificationResponseBody {
	s.Data = v
	return s
}

func (s *CertNoThreeElementVerificationResponseBody) SetMessage(v string) *CertNoThreeElementVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *CertNoThreeElementVerificationResponseBody) SetRequestId(v string) *CertNoThreeElementVerificationResponseBody {
	s.RequestId = &v
	return s
}

type CertNoThreeElementVerificationResponseBodyData struct {
	// example:
	//
	// 1
	IsConsistent *string `json:"IsConsistent,omitempty" xml:"IsConsistent,omitempty"`
}

func (s CertNoThreeElementVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CertNoThreeElementVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CertNoThreeElementVerificationResponseBodyData) SetIsConsistent(v string) *CertNoThreeElementVerificationResponseBodyData {
	s.IsConsistent = &v
	return s
}

type CertNoThreeElementVerificationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CertNoThreeElementVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CertNoThreeElementVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s CertNoThreeElementVerificationResponse) GoString() string {
	return s.String()
}

func (s *CertNoThreeElementVerificationResponse) SetHeaders(v map[string]*string) *CertNoThreeElementVerificationResponse {
	s.Headers = v
	return s
}

func (s *CertNoThreeElementVerificationResponse) SetStatusCode(v int32) *CertNoThreeElementVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CertNoThreeElementVerificationResponse) SetBody(v *CertNoThreeElementVerificationResponseBody) *CertNoThreeElementVerificationResponse {
	s.Body = v
	return s
}

type CertNoTwoElementVerificationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// QASDW@#**
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 331021200001010000
	CertNo               *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CertNoTwoElementVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s CertNoTwoElementVerificationRequest) GoString() string {
	return s.String()
}

func (s *CertNoTwoElementVerificationRequest) SetAuthCode(v string) *CertNoTwoElementVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *CertNoTwoElementVerificationRequest) SetCertName(v string) *CertNoTwoElementVerificationRequest {
	s.CertName = &v
	return s
}

func (s *CertNoTwoElementVerificationRequest) SetCertNo(v string) *CertNoTwoElementVerificationRequest {
	s.CertNo = &v
	return s
}

func (s *CertNoTwoElementVerificationRequest) SetOwnerId(v int64) *CertNoTwoElementVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *CertNoTwoElementVerificationRequest) SetResourceOwnerAccount(v string) *CertNoTwoElementVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CertNoTwoElementVerificationRequest) SetResourceOwnerId(v int64) *CertNoTwoElementVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

type CertNoTwoElementVerificationResponseBody struct {
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CertNoTwoElementVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CertNoTwoElementVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CertNoTwoElementVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *CertNoTwoElementVerificationResponseBody) SetAccessDeniedDetail(v string) *CertNoTwoElementVerificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CertNoTwoElementVerificationResponseBody) SetCode(v string) *CertNoTwoElementVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *CertNoTwoElementVerificationResponseBody) SetData(v *CertNoTwoElementVerificationResponseBodyData) *CertNoTwoElementVerificationResponseBody {
	s.Data = v
	return s
}

func (s *CertNoTwoElementVerificationResponseBody) SetMessage(v string) *CertNoTwoElementVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *CertNoTwoElementVerificationResponseBody) SetRequestId(v string) *CertNoTwoElementVerificationResponseBody {
	s.RequestId = &v
	return s
}

type CertNoTwoElementVerificationResponseBodyData struct {
	// example:
	//
	// 1
	IsConsistent *string `json:"IsConsistent,omitempty" xml:"IsConsistent,omitempty"`
}

func (s CertNoTwoElementVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CertNoTwoElementVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CertNoTwoElementVerificationResponseBodyData) SetIsConsistent(v string) *CertNoTwoElementVerificationResponseBodyData {
	s.IsConsistent = &v
	return s
}

type CertNoTwoElementVerificationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CertNoTwoElementVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CertNoTwoElementVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s CertNoTwoElementVerificationResponse) GoString() string {
	return s.String()
}

func (s *CertNoTwoElementVerificationResponse) SetHeaders(v map[string]*string) *CertNoTwoElementVerificationResponse {
	s.Headers = v
	return s
}

func (s *CertNoTwoElementVerificationResponse) SetStatusCode(v int32) *CertNoTwoElementVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CertNoTwoElementVerificationResponse) SetBody(v *CertNoTwoElementVerificationResponseBody) *CertNoTwoElementVerificationResponse {
	s.Body = v
	return s
}

type CompanyFourElementsVerificationRequest struct {
	// The authorization code.
	//
	// >  On the [My Applications](https://dytns.console.aliyun.com/analysis/apply) page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/overview?spm=a2c4g.608385.0.0.79847f8b3awqUC), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The enterprise name.
	//
	// example:
	//
	// 示例值示例值
	EpCertName *string `json:"EpCertName,omitempty" xml:"EpCertName,omitempty"`
	// The business license number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9242032*******J627
	EpCertNo *string `json:"EpCertNo,omitempty" xml:"EpCertNo,omitempty"`
	// The name of the legal representative.
	//
	// >  If an enterprise has multiple legal representatives, separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	LegalPersonCertName *string `json:"LegalPersonCertName,omitempty" xml:"LegalPersonCertName,omitempty"`
	// The ID card number of the legal representative.
	//
	// >  If an enterprise has multiple legal representatives, separate the ID card numbers with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 331021********0011
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
	// The details about the access denial.
	//
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *CompanyFourElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID. It is a common parameter and can be used to troubleshoot issues.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The information about the enterprise.
	DetailInfo *CompanyFourElementsVerificationResponseBodyDataDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Struct"`
	// The fields to be verified.
	InconsistentData []*string `json:"InconsistentData,omitempty" xml:"InconsistentData,omitempty" type:"Repeated"`
	// The code of the verification result. Valid values:
	//
	// 	- 0: The four elements belong to the same enterprise.
	//
	// 	- 1: The four elements belong to the same enterprise, but the business status of the enterprise is abnormal.
	//
	// 	- 2: The legal representative information cannot match the enterprise information.
	//
	// 	- 3: The four elements do not belong to the same enterprise.
	//
	// 	- 4: No information about the enterprise is found.
	//
	// 	- 5: No information about the legal representative is found.
	//
	// example:
	//
	// 0
	ReasonCode *int64 `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The verification result. Valid values:
	//
	// 	- true: The four elements belong to the same enterprise and the business status of the enterprise is Active.
	//
	// 	- false: The four elements do not belong to the same enterprise.
	//
	// example:
	//
	// true
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s CompanyFourElementsVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CompanyFourElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompanyFourElementsVerificationResponseBodyData) SetDetailInfo(v *CompanyFourElementsVerificationResponseBodyDataDetailInfo) *CompanyFourElementsVerificationResponseBodyData {
	s.DetailInfo = v
	return s
}

func (s *CompanyFourElementsVerificationResponseBodyData) SetInconsistentData(v []*string) *CompanyFourElementsVerificationResponseBodyData {
	s.InconsistentData = v
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

type CompanyFourElementsVerificationResponseBodyDataDetailInfo struct {
	// The business status of the enterprise.
	//
	// example:
	//
	// Active
	EnterpriseStatus *string `json:"EnterpriseStatus,omitempty" xml:"EnterpriseStatus,omitempty"`
	// The business term of the enterprise.
	//
	// example:
	//
	// 2023-05-25/2053-05-24
	OpenTime *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
}

func (s CompanyFourElementsVerificationResponseBodyDataDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s CompanyFourElementsVerificationResponseBodyDataDetailInfo) GoString() string {
	return s.String()
}

func (s *CompanyFourElementsVerificationResponseBodyDataDetailInfo) SetEnterpriseStatus(v string) *CompanyFourElementsVerificationResponseBodyDataDetailInfo {
	s.EnterpriseStatus = &v
	return s
}

func (s *CompanyFourElementsVerificationResponseBodyDataDetailInfo) SetOpenTime(v string) *CompanyFourElementsVerificationResponseBodyDataDetailInfo {
	s.OpenTime = &v
	return s
}

type CompanyFourElementsVerificationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompanyFourElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The authorization code.
	//
	// >  On the [My Applications](https://dytns.console.aliyun.com/analysis/apply) page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/overview?spm=a2c4g.608385.0.0.79847f8b3awqUC), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The enterprise name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	EpCertName *string `json:"EpCertName,omitempty" xml:"EpCertName,omitempty"`
	// The business license number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9242032*******J627
	EpCertNo *string `json:"EpCertNo,omitempty" xml:"EpCertNo,omitempty"`
	// The name of the legal representative.
	//
	// >  If an enterprise has multiple legal representatives, separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
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
	// The details about the access denial.
	//
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *CompanyThreeElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID. It is a common parameter and can be used to troubleshoot issues.
	//
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD135850177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The information about the enterprise.
	DetailInfo *CompanyThreeElementsVerificationResponseBodyDataDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Struct"`
	// The fields to be verified.
	InconsistentData []*string `json:"InconsistentData,omitempty" xml:"InconsistentData,omitempty" type:"Repeated"`
	// The code of the verification result. Valid values:
	//
	// 	- 0: The three elements belong to the same enterprise.
	//
	// 	- 1: The three elements belong to the same enterprise, and the business status of the enterprise is abnormal.
	//
	// 	- 2: The legal representative information cannot match the enterprise information.
	//
	// 	- 3: The three elements do not belong to the same enterprise.
	//
	// 	- 4: No information about the enterprise is found.
	//
	// 	- 5: No information about the legal representative is found.
	//
	// example:
	//
	// 0
	ReasonCode *int64 `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The verification result. Valid values:
	//
	// 	- true: The three elements belong to the same enterprise and the business status of the enterprise is Active.
	//
	// 	- false: The three elements do not belong to the same enterprise.
	//
	// example:
	//
	// true
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s CompanyThreeElementsVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CompanyThreeElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompanyThreeElementsVerificationResponseBodyData) SetDetailInfo(v *CompanyThreeElementsVerificationResponseBodyDataDetailInfo) *CompanyThreeElementsVerificationResponseBodyData {
	s.DetailInfo = v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBodyData) SetInconsistentData(v []*string) *CompanyThreeElementsVerificationResponseBodyData {
	s.InconsistentData = v
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

type CompanyThreeElementsVerificationResponseBodyDataDetailInfo struct {
	// The business status of the enterprise.
	//
	// example:
	//
	// Active
	EnterpriseStatus *string `json:"EnterpriseStatus,omitempty" xml:"EnterpriseStatus,omitempty"`
	// The business term of the enterprise.
	//
	// example:
	//
	// 2023-05-25/2053-05-24
	OpenTime *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
}

func (s CompanyThreeElementsVerificationResponseBodyDataDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s CompanyThreeElementsVerificationResponseBodyDataDetailInfo) GoString() string {
	return s.String()
}

func (s *CompanyThreeElementsVerificationResponseBodyDataDetailInfo) SetEnterpriseStatus(v string) *CompanyThreeElementsVerificationResponseBodyDataDetailInfo {
	s.EnterpriseStatus = &v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBodyDataDetailInfo) SetOpenTime(v string) *CompanyThreeElementsVerificationResponseBodyDataDetailInfo {
	s.OpenTime = &v
	return s
}

type CompanyThreeElementsVerificationResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompanyThreeElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The authorization code.
	//
	// >  On the [My Applications](https://dytns.console.aliyun.com/analysis/apply) page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/overview?spm=a2c4g.608385.0.0.79847f8b3awqUC), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The enterprise name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	EpCertName *string `json:"EpCertName,omitempty" xml:"EpCertName,omitempty"`
	// The business license number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9242032*******J627
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
	// The details about the access denial.
	//
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *CompanyTwoElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID. It is a common parameter and can be used to troubleshoot issues.
	//
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD135850177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The information about the enterprise.
	DetailInfo *CompanyTwoElementsVerificationResponseBodyDataDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Struct"`
	// The fields to be verified.
	InconsistentData []*string `json:"InconsistentData,omitempty" xml:"InconsistentData,omitempty" type:"Repeated"`
	// The code of the verification result. Valid values:
	//
	// 	- 0: The two elements belong to the same enterprise.
	//
	// 	- 1: The two elements belong to the same enterprise, but the business status of the enterprise is abnormal.
	//
	// 	- 3: The two elements do not belong to the same enterprise.
	//
	// 	- 4: No information about the enterprise is found.
	//
	// example:
	//
	// 0
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The verification result. Valid values:
	//
	// 	- true: The two elements belong to the same enterprise and the business status of the enterprise is Active.
	//
	// 	- false: The two elements do not belong to the same enterprise.
	//
	// example:
	//
	// true
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s CompanyTwoElementsVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CompanyTwoElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompanyTwoElementsVerificationResponseBodyData) SetDetailInfo(v *CompanyTwoElementsVerificationResponseBodyDataDetailInfo) *CompanyTwoElementsVerificationResponseBodyData {
	s.DetailInfo = v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBodyData) SetInconsistentData(v []*string) *CompanyTwoElementsVerificationResponseBodyData {
	s.InconsistentData = v
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

type CompanyTwoElementsVerificationResponseBodyDataDetailInfo struct {
	// The business status of the enterprise.
	//
	// example:
	//
	// Active
	EnterpriseStatus *string `json:"EnterpriseStatus,omitempty" xml:"EnterpriseStatus,omitempty"`
	// The business term of the enterprise.
	//
	// example:
	//
	// 2023-05-25/2053-05-24
	OpenTime *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
}

func (s CompanyTwoElementsVerificationResponseBodyDataDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s CompanyTwoElementsVerificationResponseBodyDataDetailInfo) GoString() string {
	return s.String()
}

func (s *CompanyTwoElementsVerificationResponseBodyDataDetailInfo) SetEnterpriseStatus(v string) *CompanyTwoElementsVerificationResponseBodyDataDetailInfo {
	s.EnterpriseStatus = &v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBodyDataDetailInfo) SetOpenTime(v string) *CompanyTwoElementsVerificationResponseBodyDataDetailInfo {
	s.OpenTime = &v
	return s
}

type CompanyTwoElementsVerificationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompanyTwoElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be queried.
	//
	// >  You can query only one phone number a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 189****1234
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Valid values:
	//
	// 	- **NORMAL**: The phone number is not encrypted.
	//
	// 	- **MD5**
	//
	// 	- **SHA256**
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
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
	// The HTTP status code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **InvalidPhoneNumber.Check**: The phone number is invalid.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *DescribeEmptyNumberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID. It is a common parameter and can be used to troubleshoot issues.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The specified phone number.
	//
	// example:
	//
	// 189****1234
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The returned status for the queried phone number. Valid values:
	//
	// 	- **EMPTY**: The queried phone number is a nonexistent number.
	//
	// 	- **NORMAL**: The queried phone number is valid.
	//
	// 	- **SUSPECT_EMPTY**: The queried phone number is suspected to be a nonexistent number.
	//
	// 	- **UNKNOWN**: The queried phone number is unknown.
	//
	// example:
	//
	// EMPTY
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEmptyNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeMobileOperatorAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeMobileOperatorAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMobileOperatorAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeMobileOperatorAttributeRequest) SetAuthCode(v string) *DescribeMobileOperatorAttributeRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribeMobileOperatorAttributeRequest) SetInputNumber(v string) *DescribeMobileOperatorAttributeRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribeMobileOperatorAttributeRequest) SetMask(v string) *DescribeMobileOperatorAttributeRequest {
	s.Mask = &v
	return s
}

func (s *DescribeMobileOperatorAttributeRequest) SetOwnerId(v int64) *DescribeMobileOperatorAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMobileOperatorAttributeRequest) SetResourceOwnerAccount(v string) *DescribeMobileOperatorAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMobileOperatorAttributeRequest) SetResourceOwnerId(v int64) *DescribeMobileOperatorAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeMobileOperatorAttributeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeMobileOperatorAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMobileOperatorAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMobileOperatorAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMobileOperatorAttributeResponseBody) SetAccessDeniedDetail(v string) *DescribeMobileOperatorAttributeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBody) SetCode(v string) *DescribeMobileOperatorAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBody) SetData(v *DescribeMobileOperatorAttributeResponseBodyData) *DescribeMobileOperatorAttributeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBody) SetMessage(v string) *DescribeMobileOperatorAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBody) SetRequestId(v string) *DescribeMobileOperatorAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMobileOperatorAttributeResponseBodyData struct {
	// example:
	//
	// 示例值示例值示例值
	BasicCarrier *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	// example:
	//
	// 示例值
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// example:
	//
	// 示例值示例值
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// true
	IsNumberPortability *bool `json:"IsNumberPortability,omitempty" xml:"IsNumberPortability,omitempty"`
	// example:
	//
	// 示例值示例值
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// example:
	//
	// 示例值示例值
	RealNumber *string `json:"RealNumber,omitempty" xml:"RealNumber,omitempty"`
	// example:
	//
	// 示例值示例值
	SegmentCarrier *string `json:"SegmentCarrier,omitempty" xml:"SegmentCarrier,omitempty"`
}

func (s DescribeMobileOperatorAttributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMobileOperatorAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) SetBasicCarrier(v string) *DescribeMobileOperatorAttributeResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) SetCarrier(v string) *DescribeMobileOperatorAttributeResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) SetCity(v string) *DescribeMobileOperatorAttributeResponseBodyData {
	s.City = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) SetIsNumberPortability(v bool) *DescribeMobileOperatorAttributeResponseBodyData {
	s.IsNumberPortability = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) SetProvince(v string) *DescribeMobileOperatorAttributeResponseBodyData {
	s.Province = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) SetRealNumber(v string) *DescribeMobileOperatorAttributeResponseBodyData {
	s.RealNumber = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) SetSegmentCarrier(v string) *DescribeMobileOperatorAttributeResponseBodyData {
	s.SegmentCarrier = &v
	return s
}

type DescribeMobileOperatorAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMobileOperatorAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMobileOperatorAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMobileOperatorAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeMobileOperatorAttributeResponse) SetHeaders(v map[string]*string) *DescribeMobileOperatorAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeMobileOperatorAttributeResponse) SetStatusCode(v int32) *DescribeMobileOperatorAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponse) SetBody(v *DescribeMobileOperatorAttributeResponseBody) *DescribeMobileOperatorAttributeResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberAnalysisRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// QASDW@#**
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1310000****
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// example:
	//
	// MD5
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// example:
	//
	// 0
	NumberType *int64 `json:"NumberType,omitempty" xml:"NumberType,omitempty"`
	OwnerId    *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 10
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribePhoneNumberAnalysisResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetAccessDeniedDetail(v string) *DescribePhoneNumberAnalysisResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// example:
	//
	// NO
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 18752785620
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The authorization code.
	//
	// >  On the ****[**Labels**](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click Activate Now, enter the required information, and then submit your application. After your application is approved, you can obtain an authorization code.
	//
	// This parameter is required.
	//
	// example:
	//
	// HwD***nG
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// 187****5620
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The model parameter configuration. This field is required by some labels.
	//
	// example:
	//
	// {"trainingJobId": "17**********48"}
	ModelConfig *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The score threshold for the phone number. Valid values: **0 to 100**.
	//
	// >  The system provided by Alibaba Cloud determines whether to accept the specified score threshold. When the system does not accept the specified score threshold, the value of this field is invalid.
	//
	// example:
	//
	// 96
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. Valid values:
	//
	// 	- OK: The request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *DescribePhoneNumberAnalysisAIResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD******177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberAnalysisAIResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisAIResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisAIResponseBody) SetAccessDeniedDetail(v string) *DescribePhoneNumberAnalysisAIResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The returned code.
	//
	// 	- YES: The specified phone number is valid.
	//
	// 	- NO: The specified phone number is invalid.
	//
	// 	- UNKNOWN: The specified phone number is unknown
	//
	// example:
	//
	// YES
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The specified phone number.
	//
	// example:
	//
	// 187****5620
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberAnalysisAIResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribePhoneNumberAnalysisPaiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	ModelConfig *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 16
	Rate                 *int64  `json:"Rate,omitempty" xml:"Rate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberAnalysisPaiRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisPaiRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisPaiRequest) SetAuthCode(v string) *DescribePhoneNumberAnalysisPaiRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiRequest) SetInputNumber(v string) *DescribePhoneNumberAnalysisPaiRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiRequest) SetModelConfig(v string) *DescribePhoneNumberAnalysisPaiRequest {
	s.ModelConfig = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiRequest) SetOwnerId(v int64) *DescribePhoneNumberAnalysisPaiRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiRequest) SetRate(v int64) *DescribePhoneNumberAnalysisPaiRequest {
	s.Rate = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberAnalysisPaiRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberAnalysisPaiRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribePhoneNumberAnalysisPaiResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberAnalysisPaiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisPaiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisPaiResponseBody) SetAccessDeniedDetail(v string) *DescribePhoneNumberAnalysisPaiResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiResponseBody) SetCode(v string) *DescribePhoneNumberAnalysisPaiResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiResponseBody) SetData(v []*string) *DescribePhoneNumberAnalysisPaiResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiResponseBody) SetMessage(v string) *DescribePhoneNumberAnalysisPaiResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiResponseBody) SetRequestId(v string) *DescribePhoneNumberAnalysisPaiResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberAnalysisPaiResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberAnalysisPaiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhoneNumberAnalysisPaiResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisPaiResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisPaiResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberAnalysisPaiResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiResponse) SetStatusCode(v int32) *DescribePhoneNumberAnalysisPaiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiResponse) SetBody(v *DescribePhoneNumberAnalysisPaiResponseBody) *DescribePhoneNumberAnalysisPaiResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberAnalysisTransparentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// QASDW@#**
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 187****5620
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	NumberType           *string `json:"NumberType,omitempty" xml:"NumberType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberAnalysisTransparentRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisTransparentRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) SetAuthCode(v string) *DescribePhoneNumberAnalysisTransparentRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) SetInputNumber(v string) *DescribePhoneNumberAnalysisTransparentRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) SetIp(v string) *DescribePhoneNumberAnalysisTransparentRequest {
	s.Ip = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) SetNumberType(v string) *DescribePhoneNumberAnalysisTransparentRequest {
	s.NumberType = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) SetOwnerId(v int64) *DescribePhoneNumberAnalysisTransparentRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberAnalysisTransparentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberAnalysisTransparentRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribePhoneNumberAnalysisTransparentResponseBody struct {
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribePhoneNumberAnalysisTransparentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberAnalysisTransparentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisTransparentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBody) SetAccessDeniedDetail(v string) *DescribePhoneNumberAnalysisTransparentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBody) SetCode(v string) *DescribePhoneNumberAnalysisTransparentResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBody) SetData(v *DescribePhoneNumberAnalysisTransparentResponseBodyData) *DescribePhoneNumberAnalysisTransparentResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBody) SetMessage(v string) *DescribePhoneNumberAnalysisTransparentResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBody) SetRequestId(v string) *DescribePhoneNumberAnalysisTransparentResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberAnalysisTransparentResponseBodyData struct {
	// example:
	//
	// 示例值示例值示例值
	DeviceRisk *string `json:"Device_risk,omitempty" xml:"Device_risk,omitempty"`
	// example:
	//
	// 示例值示例值
	IpRisk *string `json:"Ip_risk,omitempty" xml:"Ip_risk,omitempty"`
	// example:
	//
	// 0.6
	Score1 *string `json:"Score1,omitempty" xml:"Score1,omitempty"`
	// example:
	//
	// 0.2
	Score2 *string `json:"Score2,omitempty" xml:"Score2,omitempty"`
	// example:
	//
	// 0.8
	Score3 *string `json:"Score3,omitempty" xml:"Score3,omitempty"`
}

func (s DescribePhoneNumberAnalysisTransparentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisTransparentResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBodyData) SetDeviceRisk(v string) *DescribePhoneNumberAnalysisTransparentResponseBodyData {
	s.DeviceRisk = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBodyData) SetIpRisk(v string) *DescribePhoneNumberAnalysisTransparentResponseBodyData {
	s.IpRisk = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBodyData) SetScore1(v string) *DescribePhoneNumberAnalysisTransparentResponseBodyData {
	s.Score1 = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBodyData) SetScore2(v string) *DescribePhoneNumberAnalysisTransparentResponseBodyData {
	s.Score2 = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponseBodyData) SetScore3(v string) *DescribePhoneNumberAnalysisTransparentResponseBodyData {
	s.Score3 = &v
	return s
}

type DescribePhoneNumberAnalysisTransparentResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberAnalysisTransparentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhoneNumberAnalysisTransparentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisTransparentResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisTransparentResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberAnalysisTransparentResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponse) SetStatusCode(v int32) *DescribePhoneNumberAnalysisTransparentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponse) SetBody(v *DescribePhoneNumberAnalysisTransparentResponseBody) *DescribePhoneNumberAnalysisTransparentResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberAttributeRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
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
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **InvalidParameter**: The specified phone number is invalid or the parameter format is invalid.
	//
	// 	- **PhoneNumberNotfound**: No attribute information can be found for the specified phone number.
	//
	// 	- **isp.UNKNOWN**: An unknown exception occurred.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The attribute information about the phone number.
	PhoneNumberAttribute *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute `json:"PhoneNumberAttribute,omitempty" xml:"PhoneNumberAttribute,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The basic carrier. Valid values:
	//
	// 	- **China Mobile**
	//
	// 	- **China Unicom**
	//
	// 	- **China Telecom**
	BasicCarrier *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	// The actual carrier, including the virtual network operator (VNO). If the phone number involves mobile number portability, the value of this parameter is the carrier after mobile number portability.
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The city where the phone number is registered.
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// Indicates whether the phone number involves mobile number portability. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsNumberPortability *bool `json:"IsNumberPortability,omitempty" xml:"IsNumberPortability,omitempty"`
	// The number segment to which the phone number belongs.
	//
	// example:
	//
	// 139
	NumberSegment *int64 `json:"NumberSegment,omitempty" xml:"NumberSegment,omitempty"`
	// The province where the phone number is registered.
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// QASDW@#**
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The carrier. Valid values:
	//
	// 	- **MOBILE**: China Mobile
	//
	// 	- **UNICOM**: China Unicom
	//
	// 	- **TELECOM**: China Telecom
	//
	// >  Alibaba Cloud automatically determines the carrier based on the carrier who assigns the phone number. Therefore, the value of this field does not affect the query result.
	//
	// example:
	//
	// UNICOM
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The phone number to be queried.
	//
	// 	- If the value of Mask is NORMAL, specify an 11-digit phone number in plaintext.
	//
	// 	- If the value of Mask is MD5, specify a 32-bit string that is encrypted by using MD5.
	//
	// 	- If the value of Mask is SHA256, specify a 64-bit string that is encrypted by using SHA256.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Valid values:
	//
	// 	- **NORMAL**: The phone number is not encrypted.
	//
	// 	- **MD5**
	//
	// 	- **SHA256**
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
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
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **PortabilityNumberNotSupported**: The phone number that is involved in mobile number portability is not supported.
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// >  You are charged if the value of Code is OK and the value of VerifyResult is not -1. For more information, see [Pricing](https://help.aliyun.com/document_detail/154751.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *DescribePhoneNumberOnlineTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The carrier code. Valid values:
	//
	// 	- **CMCC**: China Mobile
	//
	// 	- **CUCC**: China Unicom
	//
	// 	- **CTCC**: China Telecom
	//
	// 	- **CBN**: China Broadnet
	//
	// example:
	//
	// CMCC
	CarrierCode *string `json:"CarrierCode,omitempty" xml:"CarrierCode,omitempty"`
	// The enumerated value of the usage period of a phone number. Valid values:
	//
	// 	- **-1**: No usage period information is available for the phone number.
	//
	// 	- **0**: The phone number status is abnormal. For example, the phone number is a nonexistent number.
	//
	// 	- **1*	- :[0-3) months.
	//
	// 	- **2*	- :[3-6] months.
	//
	// 	- **3*	- :(6-12] months.
	//
	// 	- **4*	- :(12-24] months.
	//
	// 	- **5*	- :(24,+) months.
	//
	// example:
	//
	// 1
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberOnlineTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// example:
	//
	// 示例值
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The phone number to be queried.
	//
	// 	- If the value of Mask is NORMAL, specify an 11-digit phone number in plaintext.
	//
	// 	- If the value of Mask is MD5, specify a 32-bit string that is encrypted by using MD5.
	//
	// 	- If the value of Mask is SHA256, specify a 64-bit string that is encrypted by using SHA256.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****1234
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Valid values:
	//
	// 	- **NORMAL**: The phone number is not encrypted.
	//
	// 	- **MD5**: The phone number is MD5-encrypted.
	//
	// 	- **SHA256**: The phone number is SHA256-encrypted.
	//
	// > Letters in the string must be uppercase.
	//
	// This parameter is required.
	//
	// example:
	//
	// MD5
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	ResultCount *string `json:"ResultCount,omitempty" xml:"ResultCount,omitempty"`
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

func (s *DescribePhoneNumberOperatorAttributeRequest) SetFlowName(v string) *DescribePhoneNumberOperatorAttributeRequest {
	s.FlowName = &v
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

func (s *DescribePhoneNumberOperatorAttributeRequest) SetResultCount(v string) *DescribePhoneNumberOperatorAttributeRequest {
	s.ResultCount = &v
	return s
}

type DescribePhoneNumberOperatorAttributeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **InvalidParameter**: The specified phone number is invalid or the parameter format is invalid.
	//
	// 	- **PhoneNumberNotfound**: No attribute information can be found for the specified phone number.
	//
	// 	- **isp.UNKNOWN**: An unknown exception occurred.
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *DescribePhoneNumberOperatorAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD135850177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) SetAccessDeniedDetail(v string) *DescribePhoneNumberOperatorAttributeResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The basic carrier. Valid values:
	//
	// 	- **China Mobile**
	//
	// 	- **China Unicom**
	//
	// 	- **China Telecom**
	//
	// 	- **China Broadnet**
	//
	// example:
	//
	// China Mobile
	BasicCarrier *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	// The actual carrier, including the virtual network operator (VNO). If the phone number involves mobile number portability, the value of this parameter is the carrier after mobile number portability.
	//
	// example:
	//
	// China Mobile
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The city where the phone number is registered.
	//
	// example:
	//
	// hangzhou
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// Indicates whether the phone number involves mobile number portability. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsNumberPortability *bool `json:"IsNumberPortability,omitempty" xml:"IsNumberPortability,omitempty"`
	// The number segment to which the phone number belongs.
	//
	// example:
	//
	// 139
	NumberSegment *int64 `json:"NumberSegment,omitempty" xml:"NumberSegment,omitempty"`
	// The province where the phone number is registered.
	//
	// example:
	//
	// zhejiang
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
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
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberOperatorAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribePhoneNumberOperatorAttributeAnnualRequest struct {
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	Mask     *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	Number   *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeAnnualRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeAnnualRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeAnnualRequest) SetAuthCode(v string) *DescribePhoneNumberOperatorAttributeAnnualRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualRequest) SetMask(v string) *DescribePhoneNumberOperatorAttributeAnnualRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualRequest) SetNumber(v string) *DescribePhoneNumberOperatorAttributeAnnualRequest {
	s.Number = &v
	return s
}

type DescribePhoneNumberOperatorAttributeAnnualResponseBody struct {
	Code      *string                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeAnnualResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeAnnualResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBody) SetCode(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBody) SetData(v *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) *DescribePhoneNumberOperatorAttributeAnnualResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBody) SetMessage(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBody) SetRequestId(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBody) SetSuccess(v bool) *DescribePhoneNumberOperatorAttributeAnnualResponseBody {
	s.Success = &v
	return s
}

type DescribePhoneNumberOperatorAttributeAnnualResponseBodyData struct {
	BasicCarrier      *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	Carrier           *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	City              *string `json:"City,omitempty" xml:"City,omitempty"`
	Mts               *string `json:"Mts,omitempty" xml:"Mts,omitempty"`
	NumberPortability *bool   `json:"NumberPortability,omitempty" xml:"NumberPortability,omitempty"`
	Province          *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) SetBasicCarrier(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) SetCarrier(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) SetCity(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData {
	s.City = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) SetMts(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData {
	s.Mts = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) SetNumberPortability(v bool) *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData {
	s.NumberPortability = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) SetProvince(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData {
	s.Province = &v
	return s
}

type DescribePhoneNumberOperatorAttributeAnnualResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberOperatorAttributeAnnualResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeAnnualResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeAnnualResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberOperatorAttributeAnnualResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponse) SetStatusCode(v int32) *DescribePhoneNumberOperatorAttributeAnnualResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponse) SetBody(v *DescribePhoneNumberOperatorAttributeAnnualResponseBody) *DescribePhoneNumberOperatorAttributeAnnualResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberOperatorAttributeAnnualUseRequest struct {
	// This parameter is required.
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// This parameter is required.
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeAnnualUseRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeAnnualUseRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseRequest) SetAuthCode(v string) *DescribePhoneNumberOperatorAttributeAnnualUseRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseRequest) SetInputNumber(v string) *DescribePhoneNumberOperatorAttributeAnnualUseRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseRequest) SetMask(v string) *DescribePhoneNumberOperatorAttributeAnnualUseRequest {
	s.Mask = &v
	return s
}

type DescribePhoneNumberOperatorAttributeAnnualUseResponseBody struct {
	AccessDeniedDetail *string                                                        `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) SetAccessDeniedDetail(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) SetCode(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) SetData(v *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) SetMessage(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) SetRequestId(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData struct {
	BasicCarrier        *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	Carrier             *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	City                *string `json:"City,omitempty" xml:"City,omitempty"`
	IsNumberPortability *bool   `json:"IsNumberPortability,omitempty" xml:"IsNumberPortability,omitempty"`
	NumberSegment       *int64  `json:"NumberSegment,omitempty" xml:"NumberSegment,omitempty"`
	Province            *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) SetBasicCarrier(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) SetCarrier(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) SetCity(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData {
	s.City = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) SetIsNumberPortability(v bool) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData {
	s.IsNumberPortability = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) SetNumberSegment(v int64) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData {
	s.NumberSegment = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) SetProvince(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData {
	s.Province = &v
	return s
}

type DescribePhoneNumberOperatorAttributeAnnualUseResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeAnnualUseResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeAnnualUseResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberOperatorAttributeAnnualUseResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponse) SetStatusCode(v int32) *DescribePhoneNumberOperatorAttributeAnnualUseResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponse) SetBody(v *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) *DescribePhoneNumberOperatorAttributeAnnualUseResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberRiskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberRiskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberRiskRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberRiskRequest) SetAuthCode(v string) *DescribePhoneNumberRiskRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberRiskRequest) SetInputNumber(v string) *DescribePhoneNumberRiskRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberRiskRequest) SetMask(v string) *DescribePhoneNumberRiskRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneNumberRiskRequest) SetOwnerId(v int64) *DescribePhoneNumberRiskRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberRiskRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberRiskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberRiskRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberRiskRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribePhoneNumberRiskResponseBody struct {
	AccessDeniedDetail *string                                  `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *DescribePhoneNumberRiskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberRiskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberRiskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberRiskResponseBody) SetAccessDeniedDetail(v string) *DescribePhoneNumberRiskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribePhoneNumberRiskResponseBody) SetCode(v string) *DescribePhoneNumberRiskResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberRiskResponseBody) SetData(v *DescribePhoneNumberRiskResponseBodyData) *DescribePhoneNumberRiskResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberRiskResponseBody) SetMessage(v string) *DescribePhoneNumberRiskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberRiskResponseBody) SetRequestId(v string) *DescribePhoneNumberRiskResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberRiskResponseBodyData struct {
	// example:
	//
	// 示例值
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s DescribePhoneNumberRiskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberRiskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberRiskResponseBodyData) SetVerifyResult(v string) *DescribePhoneNumberRiskResponseBodyData {
	s.VerifyResult = &v
	return s
}

type DescribePhoneNumberRiskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberRiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhoneNumberRiskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberRiskResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberRiskResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberRiskResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberRiskResponse) SetStatusCode(v int32) *DescribePhoneNumberRiskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberRiskResponse) SetBody(v *DescribePhoneNumberRiskResponseBody) *DescribePhoneNumberRiskResponse {
	s.Body = v
	return s
}

type DescribePhoneTwiceTelVerifyRequest struct {
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be queried.
	//
	// 	- If the value of Mask is NORMAL, specify an 11-digit phone number in plaintext.
	//
	// 	- If the value of Mask is MD5, specify a 32-bit string that is encrypted by using MD5.
	//
	// 	- If the value of Mask is SHA256, specify a 64-bit string that is encrypted by using SHA256.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139*******
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Valid values:
	//
	// 	- **NORMAL**: The phone number is not encrypted.
	//
	// 	- **MD5**
	//
	// 	- **SHA256**
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The registration time. Specify the time in the yyyy-MM-dd HH:mm:ss format. This time is the service registration time of the mobile phone user. If the service registration time is later than the time when the phone number is assigned by a carrier, it indicates that the phone number is not a reassigned number. Otherwise, the phone number is a reassigned number.
	//
	// >
	//
	// 	- If a carrier allocates a single number multiple times, the system will determine whether the phone number is a reassigned number based on the time when the carrier last allocated the phone number.
	//
	// 	- The service registration time must be later than 00:00:00 on January 1, 1970.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **PortabilityNumberNotSupported**: The phone number that is involved in mobile number portability is not supported.
	//
	// 	- **RequestNumberNotSupported**: You are not allowed to query phone numbers assigned by China Broadnet (that is, phone numbers start with 192) and phone numbers assigned by virtual network operators (VNOs).
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// >  You are charged for phone number verifications if the value of Code is OK and the value of VerifyResult is not 0. For more information, see [Pricing](https://help.aliyun.com/document_detail/154751.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *DescribePhoneTwiceTelVerifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID. It is a common parameter and can be used to troubleshoot and locate issues.
	//
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD135850177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The carrier. Valid values:
	//
	// 	- **CMCC**: China Mobile
	//
	// 	- **CUCC**: China Unicom
	//
	// 	- **CTCC**: China Telecom
	//
	// >  The returned result indicates the carrier who assigns the phone number. If the phone number involves mobile number portability, the carrier after mobile number portability is returned.
	//
	// example:
	//
	// CMCC
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The result of the request. Valid values:
	//
	// 	- **0**: It is unable to judge whether the phone number is a reassigned number.
	//
	// 	- **1**: The phone number is a reassigned number.
	//
	// 	- **2**: The phone number is not a reassigned number.
	//
	// 	- **3**: The phone number has been canceled.
	//
	// example:
	//
	// 1
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneTwiceTelVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetUAIDApplyTokenSignRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// HwD97InG
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CM
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30300
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// 示例值示例值
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b8b5b3a*******0b9893484fdf412c99
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 955EC1A869852EA8BC66F********D7C6E92017BBD5B001C736EFEAFB775C232
	ParamKey *string `json:"ParamKey,omitempty" xml:"ParamKey,omitempty"`
	// example:
	//
	// B2F0D4CD7A216D16CE2AF4BBC********29A454FDDD991F919106C12CB89ABA8
	ParamStr             *string `json:"ParamStr,omitempty" xml:"ParamStr,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20121227180001165
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetUAIDApplyTokenSignRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUAIDApplyTokenSignRequest) GoString() string {
	return s.String()
}

func (s *GetUAIDApplyTokenSignRequest) SetAuthCode(v string) *GetUAIDApplyTokenSignRequest {
	s.AuthCode = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetCarrier(v string) *GetUAIDApplyTokenSignRequest {
	s.Carrier = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetClientType(v string) *GetUAIDApplyTokenSignRequest {
	s.ClientType = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetFormat(v string) *GetUAIDApplyTokenSignRequest {
	s.Format = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetOutId(v string) *GetUAIDApplyTokenSignRequest {
	s.OutId = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetOwnerId(v int64) *GetUAIDApplyTokenSignRequest {
	s.OwnerId = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetParamKey(v string) *GetUAIDApplyTokenSignRequest {
	s.ParamKey = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetParamStr(v string) *GetUAIDApplyTokenSignRequest {
	s.ParamStr = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetResourceOwnerAccount(v string) *GetUAIDApplyTokenSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetResourceOwnerId(v int64) *GetUAIDApplyTokenSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetTime(v string) *GetUAIDApplyTokenSignRequest {
	s.Time = &v
	return s
}

type GetUAIDApplyTokenSignResponseBody struct {
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetUAIDApplyTokenSignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD******177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUAIDApplyTokenSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUAIDApplyTokenSignResponseBody) GoString() string {
	return s.String()
}

func (s *GetUAIDApplyTokenSignResponseBody) SetAccessDeniedDetail(v string) *GetUAIDApplyTokenSignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetUAIDApplyTokenSignResponseBody) SetCode(v string) *GetUAIDApplyTokenSignResponseBody {
	s.Code = &v
	return s
}

func (s *GetUAIDApplyTokenSignResponseBody) SetData(v *GetUAIDApplyTokenSignResponseBodyData) *GetUAIDApplyTokenSignResponseBody {
	s.Data = v
	return s
}

func (s *GetUAIDApplyTokenSignResponseBody) SetMessage(v string) *GetUAIDApplyTokenSignResponseBody {
	s.Message = &v
	return s
}

func (s *GetUAIDApplyTokenSignResponseBody) SetRequestId(v string) *GetUAIDApplyTokenSignResponseBody {
	s.RequestId = &v
	return s
}

type GetUAIDApplyTokenSignResponseBodyData struct {
	// example:
	//
	// CM
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// example:
	//
	// b8b5b3a*******0b9893484fdf412c99
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// example:
	//
	// D2E204D74EEB373E468632********23F592C4C9
	Sign *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
}

func (s GetUAIDApplyTokenSignResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUAIDApplyTokenSignResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUAIDApplyTokenSignResponseBodyData) SetCarrier(v string) *GetUAIDApplyTokenSignResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *GetUAIDApplyTokenSignResponseBodyData) SetOutId(v string) *GetUAIDApplyTokenSignResponseBodyData {
	s.OutId = &v
	return s
}

func (s *GetUAIDApplyTokenSignResponseBodyData) SetSign(v string) *GetUAIDApplyTokenSignResponseBodyData {
	s.Sign = &v
	return s
}

type GetUAIDApplyTokenSignResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUAIDApplyTokenSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUAIDApplyTokenSignResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUAIDApplyTokenSignResponse) GoString() string {
	return s.String()
}

func (s *GetUAIDApplyTokenSignResponse) SetHeaders(v map[string]*string) *GetUAIDApplyTokenSignResponse {
	s.Headers = v
	return s
}

func (s *GetUAIDApplyTokenSignResponse) SetStatusCode(v int32) *GetUAIDApplyTokenSignResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUAIDApplyTokenSignResponse) SetBody(v *GetUAIDApplyTokenSignResponseBody) *GetUAIDApplyTokenSignResponse {
	s.Body = v
	return s
}

type GetUAIDConversionSignRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// example:
	//
	// 示例值
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 示例值示例值
	ParamKey *string `json:"ParamKey,omitempty" xml:"ParamKey,omitempty"`
	// example:
	//
	// 示例值示例值
	ParamStr             *string `json:"ParamStr,omitempty" xml:"ParamStr,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 示例值示例值
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetUAIDConversionSignRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUAIDConversionSignRequest) GoString() string {
	return s.String()
}

func (s *GetUAIDConversionSignRequest) SetAuthCode(v string) *GetUAIDConversionSignRequest {
	s.AuthCode = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetCarrier(v string) *GetUAIDConversionSignRequest {
	s.Carrier = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetClientType(v string) *GetUAIDConversionSignRequest {
	s.ClientType = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetFormat(v string) *GetUAIDConversionSignRequest {
	s.Format = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetOutId(v string) *GetUAIDConversionSignRequest {
	s.OutId = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetOwnerId(v int64) *GetUAIDConversionSignRequest {
	s.OwnerId = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetParamKey(v string) *GetUAIDConversionSignRequest {
	s.ParamKey = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetParamStr(v string) *GetUAIDConversionSignRequest {
	s.ParamStr = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetResourceOwnerAccount(v string) *GetUAIDConversionSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetResourceOwnerId(v int64) *GetUAIDConversionSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetTime(v string) *GetUAIDConversionSignRequest {
	s.Time = &v
	return s
}

type GetUAIDConversionSignResponseBody struct {
	AccessDeniedDetail *string                                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *GetUAIDConversionSignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUAIDConversionSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUAIDConversionSignResponseBody) GoString() string {
	return s.String()
}

func (s *GetUAIDConversionSignResponseBody) SetAccessDeniedDetail(v string) *GetUAIDConversionSignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetUAIDConversionSignResponseBody) SetCode(v string) *GetUAIDConversionSignResponseBody {
	s.Code = &v
	return s
}

func (s *GetUAIDConversionSignResponseBody) SetData(v *GetUAIDConversionSignResponseBodyData) *GetUAIDConversionSignResponseBody {
	s.Data = v
	return s
}

func (s *GetUAIDConversionSignResponseBody) SetMessage(v string) *GetUAIDConversionSignResponseBody {
	s.Message = &v
	return s
}

func (s *GetUAIDConversionSignResponseBody) SetRequestId(v string) *GetUAIDConversionSignResponseBody {
	s.RequestId = &v
	return s
}

type GetUAIDConversionSignResponseBodyData struct {
	// example:
	//
	// 示例值示例值示例值
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// example:
	//
	// 示例值示例值
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// example:
	//
	// 示例值
	Sign *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
}

func (s GetUAIDConversionSignResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUAIDConversionSignResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUAIDConversionSignResponseBodyData) SetCarrier(v string) *GetUAIDConversionSignResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *GetUAIDConversionSignResponseBodyData) SetOutId(v string) *GetUAIDConversionSignResponseBodyData {
	s.OutId = &v
	return s
}

func (s *GetUAIDConversionSignResponseBodyData) SetSign(v string) *GetUAIDConversionSignResponseBodyData {
	s.Sign = &v
	return s
}

type GetUAIDConversionSignResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUAIDConversionSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUAIDConversionSignResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUAIDConversionSignResponse) GoString() string {
	return s.String()
}

func (s *GetUAIDConversionSignResponse) SetHeaders(v map[string]*string) *GetUAIDConversionSignResponse {
	s.Headers = v
	return s
}

func (s *GetUAIDConversionSignResponse) SetStatusCode(v int32) *GetUAIDConversionSignResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUAIDConversionSignResponse) SetBody(v *GetUAIDConversionSignResponseBody) *GetUAIDConversionSignResponse {
	s.Body = v
	return s
}

type InvalidPhoneNumberFilterRequest struct {
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// QASDW@#**
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number.
	//
	// >  Only the NORMAL encryption method is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
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
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **MobileNumberIllegal**: The phone number is invalid.
	//
	// 	- **EncyrptTypeIllegal**: The encryption type is invalid.
	//
	// 	- **MobileNumberTypeNotMatch**: The phone number does not match the encryption type.
	//
	// 	- **CarrierIllegal**: The carrier type is invalid.
	//
	// 	- **AuthCodeNotExist**: The authorization code does not exist.
	//
	// 	- **PortabilityNumberNotSupported**: Mobile number portability is not supported.
	//
	// 	- **Unknown**: An unknown exception occurred.
	//
	// 	- **AuthCodeAndApiNotMatch**: A system exception occurred.
	//
	// 	- **AuthCodeAndApiNotMatch**: The authorization code does not match the API operation.
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details about the returned entries.
	Data []*InvalidPhoneNumberFilterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The returned filter results.
	//
	// 	- **YES**: the valid phone number. The mappings are returned.
	//
	// 	- **NO**: the invalid phone number. No mappings are returned.
	//
	// example:
	//
	// YES
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The encrypted phone number.
	//
	// example:
	//
	// 1400513****
	EncryptedNumber *string `json:"EncryptedNumber,omitempty" xml:"EncryptedNumber,omitempty"`
	// The time when the phone number expires.
	//
	// example:
	//
	// 2022-05-27 16:05:23
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The original phone number.
	//
	// example:
	//
	// 1390000****
	OriginalNumber *string `json:"OriginalNumber,omitempty" xml:"OriginalNumber,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvalidPhoneNumberFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// This parameter is required.
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PhoneNumberConvertServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// QASDW@#**
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be queried.
	//
	// >  You can query only one phone number at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Set the value to **NORMAL**.
	//
	// >  Only the NORMAL encryption method is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
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

func (s *PhoneNumberEncryptRequest) SetOutId(v string) *PhoneNumberEncryptRequest {
	s.OutId = &v
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
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details about the returned entries.
	Data []*PhoneNumberEncryptResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The encrypted phone number.
	//
	// example:
	//
	// 1400513****
	EncryptedNumber *string `json:"EncryptedNumber,omitempty" xml:"EncryptedNumber,omitempty"`
	// The time when the phone number expires.
	//
	// example:
	//
	// 2022-05-27 16:05:23
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The original phone number.
	//
	// example:
	//
	// 1390000****
	OriginalNumber *string `json:"OriginalNumber,omitempty" xml:"OriginalNumber,omitempty"`
	OutId          *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
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

func (s *PhoneNumberEncryptResponseBodyData) SetOutId(v string) *PhoneNumberEncryptResponseBodyData {
	s.OutId = &v
	return s
}

type PhoneNumberEncryptResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PhoneNumberEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be queried.
	//
	// 	- If the value of Mask is NORMAL, specify an 11-digit phone number in plaintext.
	//
	// 	- If the value of Mask is MD5, specify a 32-bit string that is encrypted by using MD5.
	//
	// 	- If the value of Mask is SHA256, specify a 64-bit string that is encrypted by using SHA256.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****1234
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Valid values:
	//
	// 	- **NORMAL**: The phone number is not encrypted.
	//
	// 	- **MD5**
	//
	// 	- **SHA256**
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
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
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **OperatorLimit**: The carrier prohibits the query of the phone number.
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *PhoneNumberStatusForAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID. It is a common parameter and can be used to troubleshoot issues.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The basic carrier who assings the phone number. If the queried phone number involves mobile number portability, the carrier after mobile number portability is returned. Valid values:
	//
	// 	- **CMCC**: China Mobile
	//
	// 	- **CUCC**: China Unicom
	//
	// 	- **CTCC**: China Telecom
	//
	// >  You are not allowed to query the phone numbers assigned by China Broadnet.
	//
	// example:
	//
	// CMCC
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The returned status for the queried phone number. Valid values:
	//
	// 	- **NORMAL**: The queried phone number is valid.
	//
	// 	- **SHUTDOWN**: The queried phone number is suspended.
	//
	// 	- **POWER_OFF**: The queried phone number cannot be connected.
	//
	// 	- **NOT_EXIST**: The queried phone number is a nonexistent number.
	//
	// 	- **DEFECT**: The queried phone number is invalid.
	//
	// 	- **UNKNOWN**: The queried phone number is unknown.
	//
	// >  Due to system adjustment of the carrier, the BUSY and POWER_OFF states cannot be returned for the numbers assigned by China Telecom. [For more information, see the official announcements](https://help.aliyun.com/document_detail/2489709.html).
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PhoneNumberStatusForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be queried.
	//
	// 	- If the value of Mask is NORMAL, the value of this field is an 11-digit phone number.
	//
	// 	- If the value of Mask is MD5, the value of this field is a 32-bit encrypted string.
	//
	// 	- If the value of Mask is SHA256, the value of this field is a 64-bit encrypted string.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
	//
	// example:
	//
	// 139****1234
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Valid values:
	//
	// 	- **NORMAL**: The phone number is not encrypted.
	//
	// 	- **MD5**
	//
	// 	- **SHA256**
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
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
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **OperatorLimit**: The carrier prohibits the query of the phone number.
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// >  For a list of error codes, see [Service error codes](https://next.api.aliyun.com/document/Dytnsapi/2020-02-17/errorCode).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *PhoneNumberStatusForPublicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID. It is a common parameter and can be used to troubleshoot issues.
	//
	// example:
	//
	// CC3BB6D2-****-****-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The basic carrier who assigns the phone number. If the queried phone number involves mobile number portability, the carrier after mobile number portability is returned.
	//
	// Valid values:
	//
	// 	- **CMCC**: China Mobile
	//
	// 	- **CUCC**: China Unicom
	//
	// 	- **CTCC**: China Telecom
	//
	// >  You are not allowed to query the phone numbers assigned by China Broadnet.
	//
	// example:
	//
	// CMCC
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The returned status for the queried phone number. Valid values:
	//
	// 	- **NORMAL**: The queried phone number can be reached.
	//
	// 	- **SHUTDOWN**: The queried phone number is suspended.
	//
	// 	- **POWER_OFF**: The phone is powered off.
	//
	// 	- **NOT_EXIST**: The queried phone number is a nonexistent number.
	//
	// 	- **SUSPECTED_POWER_OFF**: The phone is suspected to be powered off.
	//
	// 	- **BUSY**: The queried phone number is busy.
	//
	// 	- **UNKNOWN**: The queried phone number is unknown.
	//
	// >  Due to system adjustment of the carrier, the BUSY and POWER_OFF states cannot be returned for the numbers assigned by China Telecom. [For more information, see the official announcements](https://help.aliyun.com/document_detail/2489709.html).
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PhoneNumberStatusForPublicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be queried.
	//
	// 	- If the value of Mask is NORMAL, the value of this field is an 11-digit phone number.
	//
	// 	- If the value of Mask is MD5, the value of this field is a 32-bit encrypted string.
	//
	// 	- If the value of Mask is SHA256, the value of this field is a 64-bit encrypted string.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 189****8999
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Valid values:
	//
	// 	- **NORMAL**: The phone number is not encrypted.
	//
	// 	- **MD5**
	//
	// 	- **SHA256**
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
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
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **OperatorLimit**: The carrier prohibits the query of the phone number.
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *PhoneNumberStatusForRealResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID. It is a common parameter and can be used to troubleshoot issues.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The basic carrier who assigns the phone number. If the queried phone number involves mobile number portability, the carrier after mobile number portability is returned. Valid values:
	//
	// 	- **CMCC**: China Mobile
	//
	// 	- **CUCC**: China Unicom
	//
	// 	- **CTCC**: China Telecom
	//
	// >  You are not allowed to query the phone numbers assigned by China Broadnet.
	//
	// example:
	//
	// CMCC
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The returned status for the queried phone number. Valid values:
	//
	// 	- **NORMAL**: The queried phone number can be reached.
	//
	// 	- **SHUTDOWN**: The queried phone number is suspended.
	//
	// 	- **POWER_OFF**: The phone is powered off.
	//
	// 	- **NOT_EXIST**: The queried phone number is a nonexistent number.
	//
	// 	- **BUSY**: The queried phone number is busy.
	//
	// 	- **SUSPECTED_POWER_OFF**: The phone is suspected to be powered off.
	//
	// 	- **DEFECT**: The queried phone number is invalid.
	//
	// 	- **UNKNOWN**: The queried phone number is unknown.
	//
	// >  Due to system adjustment of the carrier, the BUSY and POWER_OFF states cannot be returned for the numbers assigned by China Telecom. [For more information, see the official announcements](https://help.aliyun.com/document_detail/2489709.html).
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PhoneNumberStatusForRealResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// zf08***pi6
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be queried.
	//
	// 	- If the value of Mask is NORMAL, specify an 11-digit phone number in plaintext.
	//
	// 	- If the value of Mask is MD5, specify a 32-bit string that is encrypted by using MD5.
	//
	// 	- If the value of Mask is SHA256, specify a 64-bit string that is encrypted by using SHA256.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 181****1234
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Valid values:
	//
	// 	- **NORMAL**: plaintext
	//
	// 	- **MD5**
	//
	// 	- **SHA256**
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
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
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **OperatorLimit**: The carrier prohibits the query of the phone number.
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *PhoneNumberStatusForSmsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID. It is a common parameter and can be used to troubleshoot issues.
	//
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD135850177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The basic carrier who assigns the phone number. If the queried phone number involves mobile number portability, the carrier after mobile number portability is returned. Valid values:
	//
	// 	- **CMCC**: China Mobile
	//
	// 	- **CUCC**: China Unicom
	//
	// 	- **CTCC**: China Telecom
	//
	// >  You are not allowed to query the phone numbers assigned by China Broadnet.
	//
	// example:
	//
	// CMCC
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The returned status for the queried phone number. Valid values:
	//
	// 	- **NORMAL**: The queried phone number can be reached.
	//
	// 	- **SHUTDOWN**: The queried phone number is suspended.
	//
	// 	- **POWER_OFF**: The phone is powered off.
	//
	// 	- **NOT_EXIST**: The queried phone number is a nonexistent number.
	//
	// 	- **DEFECT**: The queried phone number is invalid.
	//
	// 	- **UNKNOWN**: The queried phone number is unknown.
	//
	// >  Due to system adjustment of the carrier, the BUSY, SUSPECTED_POWER_OFF, and POWER_OFF states cannot be returned for the numbers assigned by China Telecom. [For more information, see the official announcements](https://help.aliyun.com/document_detail/2489709.html).
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PhoneNumberStatusForSmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type PhoneNumberStatusForVoiceRequest struct {
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be queried.
	//
	// 	- If the value of Mask is NORMAL, specify an 11-digit phone number in plaintext.
	//
	// 	- If the value of Mask is MD5, specify a 32-bit string that is encrypted by using MD5.
	//
	// 	- If the value of Mask is SHA256, specify a 64-bit string that is encrypted by using SHA256.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****1234
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method. Valid values:
	//
	// 	- **NORMAL**: plaintext
	//
	// 	- **MD5**
	//
	// 	- **SHA256**
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
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
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **OperatorLimit**: The carrier prohibits the query of the phone number.
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *PhoneNumberStatusForVoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID. It is a common parameter and can be used to troubleshoot issues.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The basic carrier who assigns the phone number. If the queried phone number involves mobile number portability, the carrier after mobile number portability is returned. Valid values:
	//
	// 	- **CMCC**: China Mobile
	//
	// 	- **CUCC**: China Unicom
	//
	// 	- **CTCC**: China Telecom
	//
	// >  You are not allowed to query the phone numbers assigned by China Broadnet.
	//
	// example:
	//
	// CTCC
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The returned status for the queried phone number. Valid values:
	//
	// 	- **NORMAL**: The queried phone number can be reached.
	//
	// 	- **SHUTDOWN**: The queried phone number is suspended.
	//
	// 	- **POWER_OFF**: The phone is powered off.
	//
	// 	- **NOT_EXIST**: The queried phone number is a nonexistent number.
	//
	// 	- **SUSPECTED_POWER_OFF**: The phone is suspected to be powered off.
	//
	// 	- **DEFECT**: The queried phone number is invalid.
	//
	// 	- **UNKNOWN**: The queried phone number is unknown.
	//
	// >  Due to system adjustment of the carrier, the BUSY, SUSPECTED_POWER_OFF, and POWER_OFF states cannot be returned for the numbers assigned by China Telecom. [For more information, see the official announcements](https://help.aliyun.com/document_detail/2489709.html).
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PhoneNumberStatusForVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The tag ID.
	//
	// example:
	//
	// 22
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
	// The response code. **OK*	- indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6ADDCD31-6BC7-5913-A47F-E29A07E37FEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAvailableAuthCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type QueryPhoneNumberOnlineTimeRequest struct {
	// example:
	//
	// 示例值示例值示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// example:
	//
	// 示例值
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// example:
	//
	// 示例值示例值
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryPhoneNumberOnlineTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneNumberOnlineTimeRequest) GoString() string {
	return s.String()
}

func (s *QueryPhoneNumberOnlineTimeRequest) SetAuthCode(v string) *QueryPhoneNumberOnlineTimeRequest {
	s.AuthCode = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeRequest) SetInputNumber(v string) *QueryPhoneNumberOnlineTimeRequest {
	s.InputNumber = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeRequest) SetMask(v string) *QueryPhoneNumberOnlineTimeRequest {
	s.Mask = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeRequest) SetOwnerId(v int64) *QueryPhoneNumberOnlineTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeRequest) SetResourceOwnerAccount(v string) *QueryPhoneNumberOnlineTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeRequest) SetResourceOwnerId(v int64) *QueryPhoneNumberOnlineTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryPhoneNumberOnlineTimeResponseBody struct {
	AccessDeniedDetail *string                                     `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *QueryPhoneNumberOnlineTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPhoneNumberOnlineTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneNumberOnlineTimeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPhoneNumberOnlineTimeResponseBody) SetAccessDeniedDetail(v string) *QueryPhoneNumberOnlineTimeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponseBody) SetCode(v string) *QueryPhoneNumberOnlineTimeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponseBody) SetData(v *QueryPhoneNumberOnlineTimeResponseBodyData) *QueryPhoneNumberOnlineTimeResponseBody {
	s.Data = v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponseBody) SetMessage(v string) *QueryPhoneNumberOnlineTimeResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponseBody) SetRequestId(v string) *QueryPhoneNumberOnlineTimeResponseBody {
	s.RequestId = &v
	return s
}

type QueryPhoneNumberOnlineTimeResponseBodyData struct {
	// example:
	//
	// 示例值示例值示例值
	CarrierCode *string `json:"CarrierCode,omitempty" xml:"CarrierCode,omitempty"`
	// example:
	//
	// 29
	VerifyResult *int64 `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s QueryPhoneNumberOnlineTimeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneNumberOnlineTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPhoneNumberOnlineTimeResponseBodyData) SetCarrierCode(v string) *QueryPhoneNumberOnlineTimeResponseBodyData {
	s.CarrierCode = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponseBodyData) SetVerifyResult(v int64) *QueryPhoneNumberOnlineTimeResponseBodyData {
	s.VerifyResult = &v
	return s
}

type QueryPhoneNumberOnlineTimeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPhoneNumberOnlineTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPhoneNumberOnlineTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneNumberOnlineTimeResponse) GoString() string {
	return s.String()
}

func (s *QueryPhoneNumberOnlineTimeResponse) SetHeaders(v map[string]*string) *QueryPhoneNumberOnlineTimeResponse {
	s.Headers = v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponse) SetStatusCode(v int32) *QueryPhoneNumberOnlineTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponse) SetBody(v *QueryPhoneNumberOnlineTimeResponseBody) *QueryPhoneNumberOnlineTimeResponse {
	s.Body = v
	return s
}

type QueryPhoneTwiceTelVerifyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryPhoneTwiceTelVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneTwiceTelVerifyRequest) GoString() string {
	return s.String()
}

func (s *QueryPhoneTwiceTelVerifyRequest) SetAuthCode(v string) *QueryPhoneTwiceTelVerifyRequest {
	s.AuthCode = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyRequest) SetInputNumber(v string) *QueryPhoneTwiceTelVerifyRequest {
	s.InputNumber = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyRequest) SetMask(v string) *QueryPhoneTwiceTelVerifyRequest {
	s.Mask = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyRequest) SetOwnerId(v int64) *QueryPhoneTwiceTelVerifyRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyRequest) SetResourceOwnerAccount(v string) *QueryPhoneTwiceTelVerifyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyRequest) SetResourceOwnerId(v int64) *QueryPhoneTwiceTelVerifyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyRequest) SetStartTime(v string) *QueryPhoneTwiceTelVerifyRequest {
	s.StartTime = &v
	return s
}

type QueryPhoneTwiceTelVerifyResponseBody struct {
	AccessDeniedDetail *string                                   `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *QueryPhoneTwiceTelVerifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPhoneTwiceTelVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneTwiceTelVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPhoneTwiceTelVerifyResponseBody) SetAccessDeniedDetail(v string) *QueryPhoneTwiceTelVerifyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponseBody) SetCode(v string) *QueryPhoneTwiceTelVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponseBody) SetData(v *QueryPhoneTwiceTelVerifyResponseBodyData) *QueryPhoneTwiceTelVerifyResponseBody {
	s.Data = v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponseBody) SetMessage(v string) *QueryPhoneTwiceTelVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponseBody) SetRequestId(v string) *QueryPhoneTwiceTelVerifyResponseBody {
	s.RequestId = &v
	return s
}

type QueryPhoneTwiceTelVerifyResponseBodyData struct {
	// example:
	//
	// 示例值示例值
	CarrierCode *string `json:"CarrierCode,omitempty" xml:"CarrierCode,omitempty"`
	// example:
	//
	// 73
	VerifyResult *int64 `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s QueryPhoneTwiceTelVerifyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneTwiceTelVerifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPhoneTwiceTelVerifyResponseBodyData) SetCarrierCode(v string) *QueryPhoneTwiceTelVerifyResponseBodyData {
	s.CarrierCode = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponseBodyData) SetVerifyResult(v int64) *QueryPhoneTwiceTelVerifyResponseBodyData {
	s.VerifyResult = &v
	return s
}

type QueryPhoneTwiceTelVerifyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPhoneTwiceTelVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPhoneTwiceTelVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneTwiceTelVerifyResponse) GoString() string {
	return s.String()
}

func (s *QueryPhoneTwiceTelVerifyResponse) SetHeaders(v map[string]*string) *QueryPhoneTwiceTelVerifyResponse {
	s.Headers = v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponse) SetStatusCode(v int32) *QueryPhoneTwiceTelVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponse) SetBody(v *QueryPhoneTwiceTelVerifyResponseBody) *QueryPhoneTwiceTelVerifyResponse {
	s.Body = v
	return s
}

type QueryTagApplyRuleRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag ID.
	//
	// example:
	//
	// 61
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
	// The response code. **OK*	- indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryTagApplyRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CC3BB6D2-****-****-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The requirements for application materials.
	//
	// example:
	//
	// Requirements for application materials
	ApplyMaterialDesc *string `json:"ApplyMaterialDesc,omitempty" xml:"ApplyMaterialDesc,omitempty"`
	// Indicates whether the application is automatically approved.
	//
	// example:
	//
	// 16
	AutoAudit *int64 `json:"AutoAudit,omitempty" xml:"AutoAudit,omitempty"`
	// The URL for the billing documentation.
	//
	// example:
	//
	// aliyundoc.com
	ChargingStandardLink *string `json:"ChargingStandardLink,omitempty" xml:"ChargingStandardLink,omitempty"`
	// indicates whether encrypted queries are supported.
	//
	// example:
	//
	// 99
	EncryptedQuery *int64 `json:"EncryptedQuery,omitempty" xml:"EncryptedQuery,omitempty"`
	// Indicates whether application materials are required.
	//
	// example:
	//
	// 46
	NeedApplyMaterial *int64 `json:"NeedApplyMaterial,omitempty" xml:"NeedApplyMaterial,omitempty"`
	// The URL for the service agreement.
	//
	// example:
	//
	// example.aliyundoc.com
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTagApplyRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The industry ID.
	//
	// example:
	//
	// 58
	IndustryId           *int64  `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The scene ID.
	//
	// example:
	//
	// 83
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The tag ID.
	//
	// example:
	//
	// 71
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
	// The response code. **OK*	- indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*QueryTagInfoBySelectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1C3B8084-3A7D-570B-BC84-BF945A9CF65E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The list of available authorization codes.
	AuthCodeList   []*string `json:"AuthCodeList,omitempty" xml:"AuthCodeList,omitempty" type:"Repeated"`
	ComplexityType *string   `json:"ComplexityType,omitempty" xml:"ComplexityType,omitempty"`
	// The URL for the API demo.
	//
	// example:
	//
	// https://help.aliyun.com/document_detail/388997.html?spm=a2c4g.2573870.0.0.3aa921cbOrtqJz
	DemoAddress *string `json:"DemoAddress,omitempty" xml:"DemoAddress,omitempty"`
	// The URL for the API documentation.
	//
	// example:
	//
	// https://help.aliyun.com/document_detail/388997.html?spm=a2c4g.2573870.0.0.3aa921cbOrtqJz
	DocAddress *string `json:"DocAddress,omitempty" xml:"DocAddress,omitempty"`
	// The URL for the definitions of the enumerated values.
	//
	// example:
	//
	// example.aliyundoc.com
	EnumDefinitionAddress *string `json:"EnumDefinitionAddress,omitempty" xml:"EnumDefinitionAddress,omitempty"`
	// The flow name.
	//
	// example:
	//
	// process name
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The industry ID.
	//
	// example:
	//
	// 83
	IndustryId *int64 `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	// The industry name.
	//
	// example:
	//
	// logistics
	IndustryName *string `json:"IndustryName,omitempty" xml:"IndustryName,omitempty"`
	// The list of tag parameters.
	ParamList           []*QueryTagInfoBySelectionResponseBodyDataParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	RichTextDescription *string                                             `json:"RichTextDescription,omitempty" xml:"RichTextDescription,omitempty"`
	// The scene ID.
	//
	// example:
	//
	// 41
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The scene name.
	//
	// example:
	//
	// General scenario
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The tag ID.
	//
	// example:
	//
	// 31
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The tag name.
	//
	// example:
	//
	// Number ownership
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

func (s *QueryTagInfoBySelectionResponseBodyData) SetComplexityType(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.ComplexityType = &v
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

func (s *QueryTagInfoBySelectionResponseBodyData) SetRichTextDescription(v string) *QueryTagInfoBySelectionResponseBodyData {
	s.RichTextDescription = &v
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
	// The English name of the parameter.
	//
	// example:
	//
	// preame
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The input hint.
	//
	// example:
	//
	// none
	Hint *string `json:"Hint,omitempty" xml:"Hint,omitempty"`
	// Indicates whether the parameter is required.
	//
	// example:
	//
	// false
	Must *bool `json:"Must,omitempty" xml:"Must,omitempty"`
	// The Chinese name of the parameter.
	//
	// example:
	//
	// none
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type. The code that corresponds to EnumUIWidgetTypes.
	//
	// example:
	//
	// aqzx
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The definitions of the enumerated values such as Code or Desc.
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
	// The English name.
	//
	// example:
	//
	// Aliyun
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The Chinese name.
	//
	// example:
	//
	// 阿里云
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTagInfoBySelectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 18
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 66
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
	// The response code. **OK*	- indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryTagListPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A4475657-BB7E-585D-9E09-37934F096103
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The page number.
	//
	// example:
	//
	// 11
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 24
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The retruned data.
	Records []*QueryTagListPageResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The total number of returned entries.
	//
	// example:
	//
	// 32
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 91
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
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
	// The API operation that is called by the frontend.
	//
	// example:
	//
	// TwoElementsVerification
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// Code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The URL for the API documentation.
	//
	// example:
	//
	// https://help.aliyun.com/document_detail/388997.html?spm=a2c4g.388997.0.0.cf804cc7DX4vlP
	DocAddress *string `json:"DocAddress,omitempty" xml:"DocAddress,omitempty"`
	// The tag ID.
	//
	// example:
	//
	// 75
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The industry ID.
	//
	// example:
	//
	// 2
	IndustryId *int64 `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	// The industry name.
	//
	// example:
	//
	// Test
	IndustryName *string `json:"IndustryName,omitempty" xml:"IndustryName,omitempty"`
	// The tag description.
	//
	// example:
	//
	// for autotest new
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// Indicates whether the number is activated.
	//
	// example:
	//
	// 45
	IsOpen *int64 `json:"IsOpen,omitempty" xml:"IsOpen,omitempty"`
	// The tag name.
	//
	// example:
	//
	// Aliyun
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 	- 0: The number is hidden.
	//
	// 	- 1: The number is public.
	//
	// example:
	//
	// 1
	SaleStatusStr *string `json:"SaleStatusStr,omitempty" xml:"SaleStatusStr,omitempty"`
	// The scene ID.
	//
	// example:
	//
	// 13
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The scene name.
	//
	// example:
	//
	// check
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTagListPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The beginning of the time range to query.
	//
	// example:
	//
	// 20230308
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end of the time range to query.
	//
	// example:
	//
	// 20230406
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag ID.
	//
	// example:
	//
	// 14
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
	// The response code. **OK*	- indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*QueryUsageStatisticsByTagIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D45CC751-34DF-5797-81FB-9A2ED6DC024B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The authorization code.
	//
	// example:
	//
	// g61I8UV5zd
	AuthorizationCode *string `json:"AuthorizationCode,omitempty" xml:"AuthorizationCode,omitempty"`
	// The numbers for which the query failed.
	//
	// example:
	//
	// 71
	FailTotal *int64 `json:"FailTotal,omitempty" xml:"FailTotal,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 20230312
	GmtDateStr *string `json:"GmtDateStr,omitempty" xml:"GmtDateStr,omitempty"`
	// The ID of the authorization code usage record.
	//
	// example:
	//
	// 17
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The industry name.
	//
	// example:
	//
	// Home security
	IndustryName *string `json:"IndustryName,omitempty" xml:"IndustryName,omitempty"`
	// The customer product ID (PID).
	//
	// example:
	//
	// 89
	PartnerId *int64 `json:"PartnerId,omitempty" xml:"PartnerId,omitempty"`
	// The scene name.
	//
	// example:
	//
	// Return visit
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The numbers for which the query succeeded.
	//
	// example:
	//
	// 93
	SuccessTotal *int64 `json:"SuccessTotal,omitempty" xml:"SuccessTotal,omitempty"`
	// The tag name.
	//
	// example:
	//
	// 69
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The tag name.
	//
	// example:
	//
	// Alibaba Cloud Query
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The total quantity of numbers that are involved in the query.
	//
	// example:
	//
	// 41
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUsageStatisticsByTagIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// QASDW@#**
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The ID card number to be verified.
	//
	// 	- If the value of Mask is NORMAL, specify a value in plaintext for this field.
	//
	// 	- If the value of Mask is MD5, specify a MD5-encrypted value for this field.
	//
	// 	- If the value of Mask is SHA256, specify a SHA256-encrypted value for this field.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 83d8040d3cb2181e04****dc6ff5566d4493876a4a5da782887446356b0a787e
	CertCode *string `json:"CertCode,omitempty" xml:"CertCode,omitempty"`
	// The phone number to be verified.
	//
	// 	- If the value of Mask is NORMAL, specify a value in plaintext for this field.
	//
	// 	- If the value of Mask is MD5, specify a MD5-encrypted value for this field.
	//
	// 	- If the value of Mask is SHA256, specify a SHA256-encrypted value for this field.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method. Valid values:
	//
	// 	- **NORMAL**: The phone number is not encrypted.
	//
	// 	- **MD5**
	//
	// 	- **SHA256**
	//
	// This parameter is required.
	//
	// example:
	//
	// MD5
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The name to be verified.
	//
	// 	- If the value of Mask is NORMAL, specify a value in plaintext for this field.
	//
	// 	- If the value of Mask is MD5, specify a MD5-encrypted value for this field.
	//
	// 	- If the value of Mask is SHA256, specify a SHA256-encrypted value for this field.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
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
	// The response code.
	//
	// 	- **OK**: The request is successful.
	//
	// 	- For more information, see Error codes in this documentation.
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *ThreeElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The basic carrier. Valid values:
	//
	// 	- **China Mobile**
	//
	// 	- **China Unicom**
	//
	// 	- **China Telecom**
	//
	// example:
	//
	// China Mobile
	BasicCarrier *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	// Indicates whether the specified name, phone number, and ID card number belong to the same user. Valid values:
	//
	// 	- **1**: The specified name, phone number, and ID card number belong to the same user.
	//
	// 	- **0**: The specified name, phone number, and ID card number do not belong to the same user.
	//
	// 	- **2**: The specified name, phone number, and ID card number cannot be found.
	//
	// **Note*	- The phone number registration data of a user is usually updated one or three days after registration. The registration data can be queried only after the update. The following table shows the verification results under different phone number states.
	//
	// |Carrier/Phone number state|Out-of-service|Nonexistent|Canceled|
	//
	// |---|---|---|---|
	//
	// |China Mobile|Verifications can be carried out normally.|The specified name, phone number, and ID card number cannot be found.|The specified name, phone number, and ID card number cannot be found.|
	//
	// |China Unicom|Verifications can be carried out normally.|The specified name, phone number, and ID card number do not belong to the same user.|The specified name, phone number, and ID card number do not belong to the same user.|
	//
	// |China Telecom|Verifications can be carried out normally.|The specified name, phone number, and ID card number cannot be found.|The specified name, phone number, and ID card number cannot be found.|
	//
	// example:
	//
	// 1
	IsConsistent *int32 `json:"IsConsistent,omitempty" xml:"IsConsistent,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ThreeElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// QASDW@#**
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be verified.
	//
	// 	- If the value of Mask is NORMAL, specify a value in plaintext for this field.
	//
	// 	- If the value of Mask is MD5, specify a MD5-encrypted value for this field.
	//
	// 	- If the value of Mask is SHA256, specify a SHA256-encrypted value for this field.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method. Valid values:
	//
	// 	- **NORMAL**: plaintext
	//
	// 	- **MD5**
	//
	// 	- **SHA256**
	//
	// This parameter is required.
	//
	// example:
	//
	// MD5
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The name to be verified.
	//
	// 	- If the value of Mask is NORMAL, specify a value in plaintext for this field.
	//
	// 	- If the value of Mask is MD5, specify a MD5-encrypted value for this field.
	//
	// 	- If the value of Mask is SHA256, specify a SHA256-encrypted value for this field.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
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
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- For more information, see Error codes in this documentation.
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number or name at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *TwoElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The basic carriers. Valid values:
	//
	// 	- **China Mobile**
	//
	// 	- **China Unicom**
	//
	// 	- **China Telecom**
	//
	// >  You are not allowed to verify numbers assigned by China Broadnet.
	//
	// example:
	//
	// China Mobile
	BasicCarrier *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	// Indicates whether the specified name and phone number belong to the same user. Valid values:
	//
	// 	- **1**: The specified name and phone number belong to the same user.
	//
	// 	- **0**: The specified name and phone number do not belong to the same user.
	//
	// 	- **2**: The specified name and phone number cannot be found.
	//
	// The phone number registration data of a user is usually updated one or three days after registration. The registration data can be queried only after the update. The following table shows the verification results under different phone number states.
	//
	// |Carrier/Phone number state|Out-of-service|Nonexistent|Canceled|
	//
	// |---|---|---|---|
	//
	// |China Mobile|Verifications can be carried out normally.|The specified name and phone number cannot be found.|The specified name and phone number cannot be found.|
	//
	// |China Unicom|Verifications can be carried out normally.|The specified name and phone number do not belong to the same user.|The specified name and phone number do not belong to the same user.|
	//
	// |China Telecom|Verifications can be carried out normally.|The specified name and phone number cannot be found.|The specified name and phone number cannot be found.|
	//
	// example:
	//
	// 1
	IsConsistent *int32 `json:"IsConsistent,omitempty" xml:"IsConsistent,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TwoElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UAIDCollectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// example:
	//
	// 示例值
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 示例值示例值
	Province             *string `json:"Province,omitempty" xml:"Province,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// 示例值示例值
	UserGrantId *string `json:"UserGrantId,omitempty" xml:"UserGrantId,omitempty"`
}

func (s UAIDCollectionRequest) String() string {
	return tea.Prettify(s)
}

func (s UAIDCollectionRequest) GoString() string {
	return s.String()
}

func (s *UAIDCollectionRequest) SetAuthCode(v string) *UAIDCollectionRequest {
	s.AuthCode = &v
	return s
}

func (s *UAIDCollectionRequest) SetCarrier(v string) *UAIDCollectionRequest {
	s.Carrier = &v
	return s
}

func (s *UAIDCollectionRequest) SetIp(v string) *UAIDCollectionRequest {
	s.Ip = &v
	return s
}

func (s *UAIDCollectionRequest) SetOutId(v string) *UAIDCollectionRequest {
	s.OutId = &v
	return s
}

func (s *UAIDCollectionRequest) SetOwnerId(v int64) *UAIDCollectionRequest {
	s.OwnerId = &v
	return s
}

func (s *UAIDCollectionRequest) SetProvince(v string) *UAIDCollectionRequest {
	s.Province = &v
	return s
}

func (s *UAIDCollectionRequest) SetResourceOwnerAccount(v string) *UAIDCollectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UAIDCollectionRequest) SetResourceOwnerId(v int64) *UAIDCollectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UAIDCollectionRequest) SetToken(v string) *UAIDCollectionRequest {
	s.Token = &v
	return s
}

func (s *UAIDCollectionRequest) SetUserGrantId(v string) *UAIDCollectionRequest {
	s.UserGrantId = &v
	return s
}

type UAIDCollectionResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *UAIDCollectionResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UAIDCollectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UAIDCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *UAIDCollectionResponseBody) SetAccessDeniedDetail(v string) *UAIDCollectionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UAIDCollectionResponseBody) SetCode(v string) *UAIDCollectionResponseBody {
	s.Code = &v
	return s
}

func (s *UAIDCollectionResponseBody) SetMessage(v string) *UAIDCollectionResponseBody {
	s.Message = &v
	return s
}

func (s *UAIDCollectionResponseBody) SetModel(v *UAIDCollectionResponseBodyModel) *UAIDCollectionResponseBody {
	s.Model = v
	return s
}

func (s *UAIDCollectionResponseBody) SetRequestId(v string) *UAIDCollectionResponseBody {
	s.RequestId = &v
	return s
}

type UAIDCollectionResponseBodyModel struct {
	// example:
	//
	// 示例值
	Uaid *string `json:"Uaid,omitempty" xml:"Uaid,omitempty"`
}

func (s UAIDCollectionResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s UAIDCollectionResponseBodyModel) GoString() string {
	return s.String()
}

func (s *UAIDCollectionResponseBodyModel) SetUaid(v string) *UAIDCollectionResponseBodyModel {
	s.Uaid = &v
	return s
}

type UAIDCollectionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UAIDCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UAIDCollectionResponse) String() string {
	return tea.Prettify(s)
}

func (s UAIDCollectionResponse) GoString() string {
	return s.String()
}

func (s *UAIDCollectionResponse) SetHeaders(v map[string]*string) *UAIDCollectionResponse {
	s.Headers = v
	return s
}

func (s *UAIDCollectionResponse) SetStatusCode(v int32) *UAIDCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UAIDCollectionResponse) SetBody(v *UAIDCollectionResponseBody) *UAIDCollectionResponse {
	s.Body = v
	return s
}

type UAIDConversionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	UaidList *string `json:"UaidList,omitempty" xml:"UaidList,omitempty"`
}

func (s UAIDConversionRequest) String() string {
	return tea.Prettify(s)
}

func (s UAIDConversionRequest) GoString() string {
	return s.String()
}

func (s *UAIDConversionRequest) SetAuthCode(v string) *UAIDConversionRequest {
	s.AuthCode = &v
	return s
}

func (s *UAIDConversionRequest) SetCarrier(v string) *UAIDConversionRequest {
	s.Carrier = &v
	return s
}

func (s *UAIDConversionRequest) SetOutId(v string) *UAIDConversionRequest {
	s.OutId = &v
	return s
}

func (s *UAIDConversionRequest) SetOwnerId(v int64) *UAIDConversionRequest {
	s.OwnerId = &v
	return s
}

func (s *UAIDConversionRequest) SetResourceOwnerAccount(v string) *UAIDConversionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UAIDConversionRequest) SetResourceOwnerId(v int64) *UAIDConversionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UAIDConversionRequest) SetUaidList(v string) *UAIDConversionRequest {
	s.UaidList = &v
	return s
}

type UAIDConversionResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *UAIDConversionResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UAIDConversionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UAIDConversionResponseBody) GoString() string {
	return s.String()
}

func (s *UAIDConversionResponseBody) SetAccessDeniedDetail(v string) *UAIDConversionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UAIDConversionResponseBody) SetCode(v string) *UAIDConversionResponseBody {
	s.Code = &v
	return s
}

func (s *UAIDConversionResponseBody) SetMessage(v string) *UAIDConversionResponseBody {
	s.Message = &v
	return s
}

func (s *UAIDConversionResponseBody) SetModel(v *UAIDConversionResponseBodyModel) *UAIDConversionResponseBody {
	s.Model = v
	return s
}

func (s *UAIDConversionResponseBody) SetRequestId(v string) *UAIDConversionResponseBody {
	s.RequestId = &v
	return s
}

type UAIDConversionResponseBodyModel struct {
	// example:
	//
	// 示例值示例值示例值
	PhoneList *string `json:"PhoneList,omitempty" xml:"PhoneList,omitempty"`
}

func (s UAIDConversionResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s UAIDConversionResponseBodyModel) GoString() string {
	return s.String()
}

func (s *UAIDConversionResponseBodyModel) SetPhoneList(v string) *UAIDConversionResponseBodyModel {
	s.PhoneList = &v
	return s
}

type UAIDConversionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UAIDConversionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UAIDConversionResponse) String() string {
	return tea.Prettify(s)
}

func (s UAIDConversionResponse) GoString() string {
	return s.String()
}

func (s *UAIDConversionResponse) SetHeaders(v map[string]*string) *UAIDConversionResponse {
	s.Headers = v
	return s
}

func (s *UAIDConversionResponse) SetStatusCode(v int32) *UAIDConversionResponse {
	s.StatusCode = &v
	return s
}

func (s *UAIDConversionResponse) SetBody(v *UAIDConversionResponseBody) *UAIDConversionResponse {
	s.Body = v
	return s
}

type UAIDVerificationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// HwD97InG
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CM
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// example:
	//
	// 示例值
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// b8b5b3a*******0b9893484fdf412c99
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 示例值示例值
	Province             *string `json:"Province,omitempty" xml:"Province,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MIGfMA0********3DQEBAQUAA4GNADCB
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// Md****a3Em
	UserGrantId *string `json:"UserGrantId,omitempty" xml:"UserGrantId,omitempty"`
}

func (s UAIDVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s UAIDVerificationRequest) GoString() string {
	return s.String()
}

func (s *UAIDVerificationRequest) SetAuthCode(v string) *UAIDVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *UAIDVerificationRequest) SetCarrier(v string) *UAIDVerificationRequest {
	s.Carrier = &v
	return s
}

func (s *UAIDVerificationRequest) SetIp(v string) *UAIDVerificationRequest {
	s.Ip = &v
	return s
}

func (s *UAIDVerificationRequest) SetOutId(v string) *UAIDVerificationRequest {
	s.OutId = &v
	return s
}

func (s *UAIDVerificationRequest) SetOwnerId(v int64) *UAIDVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *UAIDVerificationRequest) SetProvince(v string) *UAIDVerificationRequest {
	s.Province = &v
	return s
}

func (s *UAIDVerificationRequest) SetResourceOwnerAccount(v string) *UAIDVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UAIDVerificationRequest) SetResourceOwnerId(v int64) *UAIDVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UAIDVerificationRequest) SetToken(v string) *UAIDVerificationRequest {
	s.Token = &v
	return s
}

func (s *UAIDVerificationRequest) SetUserGrantId(v string) *UAIDVerificationRequest {
	s.UserGrantId = &v
	return s
}

type UAIDVerificationResponseBody struct {
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UAIDVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD******177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UAIDVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UAIDVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *UAIDVerificationResponseBody) SetAccessDeniedDetail(v string) *UAIDVerificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UAIDVerificationResponseBody) SetCode(v string) *UAIDVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *UAIDVerificationResponseBody) SetData(v *UAIDVerificationResponseBodyData) *UAIDVerificationResponseBody {
	s.Data = v
	return s
}

func (s *UAIDVerificationResponseBody) SetMessage(v string) *UAIDVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *UAIDVerificationResponseBody) SetRequestId(v string) *UAIDVerificationResponseBody {
	s.RequestId = &v
	return s
}

type UAIDVerificationResponseBodyData struct {
	// example:
	//
	// B1E0C1********9F757AF52A035
	Uaid *string `json:"Uaid,omitempty" xml:"Uaid,omitempty"`
}

func (s UAIDVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UAIDVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *UAIDVerificationResponseBodyData) SetUaid(v string) *UAIDVerificationResponseBodyData {
	s.Uaid = &v
	return s
}

type UAIDVerificationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UAIDVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UAIDVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s UAIDVerificationResponse) GoString() string {
	return s.String()
}

func (s *UAIDVerificationResponse) SetHeaders(v map[string]*string) *UAIDVerificationResponse {
	s.Headers = v
	return s
}

func (s *UAIDVerificationResponse) SetStatusCode(v int32) *UAIDVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *UAIDVerificationResponse) SetBody(v *UAIDVerificationResponseBody) *UAIDVerificationResponse {
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

// Summary:
//
// 身份证三要素
//
// @param request - CertNoThreeElementVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CertNoThreeElementVerificationResponse
func (client *Client) CertNoThreeElementVerificationWithOptions(request *CertNoThreeElementVerificationRequest, runtime *util.RuntimeOptions) (_result *CertNoThreeElementVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		query["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.CertNo)) {
		query["CertNo"] = request.CertNo
	}

	if !tea.BoolValue(util.IsUnset(request.CertPicture)) {
		query["CertPicture"] = request.CertPicture
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
		Action:      tea.String("CertNoThreeElementVerification"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CertNoThreeElementVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 身份证三要素
//
// @param request - CertNoThreeElementVerificationRequest
//
// @return CertNoThreeElementVerificationResponse
func (client *Client) CertNoThreeElementVerification(request *CertNoThreeElementVerificationRequest) (_result *CertNoThreeElementVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CertNoThreeElementVerificationResponse{}
	_body, _err := client.CertNoThreeElementVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 身份证二要素认证
//
// @param request - CertNoTwoElementVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CertNoTwoElementVerificationResponse
func (client *Client) CertNoTwoElementVerificationWithOptions(request *CertNoTwoElementVerificationRequest, runtime *util.RuntimeOptions) (_result *CertNoTwoElementVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		query["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.CertNo)) {
		query["CertNo"] = request.CertNo
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
		Action:      tea.String("CertNoTwoElementVerification"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CertNoTwoElementVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 身份证二要素认证
//
// @param request - CertNoTwoElementVerificationRequest
//
// @return CertNoTwoElementVerificationResponse
func (client *Client) CertNoTwoElementVerification(request *CertNoTwoElementVerificationRequest) (_result *CertNoTwoElementVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CertNoTwoElementVerificationResponse{}
	_body, _err := client.CertNoTwoElementVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies whether the enterprise name, the business license number, and the name and ID card of the legal representative belong to the same enterprise. The verification is successful only when the preceding four elements belong to the same enterprise and the business status of the enterprise is Active.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the billing of services related to four-element verification for enterprises. For more information, see [Billing](https://help.aliyun.com/document_detail/154751.html?spm=a2c4g.154007.0.0.3edd7eb6E90YT4).
//
//		- You are charged only if the value of VerifyResult is true or false and the value of ReasonCode is 0, 1, or 2.
//
//		- Before you call this operation, perform the following operations: Log on to the [Cell Phone Number Service console](https://account.aliyun.com/login/login.htm?oauth_callback=https%3A%2F%2Fdytns.console.aliyun.com%2Foverview%3Fspm%3Da2c4g.608385.0.0.79847f8b3awqUC\\&lang=zh). On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// @param request - CompanyFourElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompanyFourElementsVerificationResponse
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

// Summary:
//
// Verifies whether the enterprise name, the business license number, and the name and ID card of the legal representative belong to the same enterprise. The verification is successful only when the preceding four elements belong to the same enterprise and the business status of the enterprise is Active.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the billing of services related to four-element verification for enterprises. For more information, see [Billing](https://help.aliyun.com/document_detail/154751.html?spm=a2c4g.154007.0.0.3edd7eb6E90YT4).
//
//		- You are charged only if the value of VerifyResult is true or false and the value of ReasonCode is 0, 1, or 2.
//
//		- Before you call this operation, perform the following operations: Log on to the [Cell Phone Number Service console](https://account.aliyun.com/login/login.htm?oauth_callback=https%3A%2F%2Fdytns.console.aliyun.com%2Foverview%3Fspm%3Da2c4g.608385.0.0.79847f8b3awqUC\\&lang=zh). On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// @param request - CompanyFourElementsVerificationRequest
//
// @return CompanyFourElementsVerificationResponse
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

// Summary:
//
// Verifies whether the enterprise name, the business license number, and the name of the legal representative belong to the same enterprise. The verification is successful only when the three elements belong to the same enterprise and the business status of the enterprise is Active.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the billing of services related to three-element verification for enterprises. For more information, see [Billing](https://help.aliyun.com/document_detail/154751.html?spm=a2c4g.154007.0.0.3edd7eb6E90YT4).
//
//		- You are charged only if the value of VerifyResult is true or false and the value of ReasonCode is 0, 1, or 2.
//
//		- Before you call this operation, perform the following operations: Log on to the [Cell Phone Number Service console](https://account.aliyun.com/login/login.htm?oauth_callback=https%3A%2F%2Fdytns.console.aliyun.com%2Foverview%3Fspm%3Da2c4g.608385.0.0.79847f8b3awqUC\\&lang=zh). On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// @param request - CompanyThreeElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompanyThreeElementsVerificationResponse
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

// Summary:
//
// Verifies whether the enterprise name, the business license number, and the name of the legal representative belong to the same enterprise. The verification is successful only when the three elements belong to the same enterprise and the business status of the enterprise is Active.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the billing of services related to three-element verification for enterprises. For more information, see [Billing](https://help.aliyun.com/document_detail/154751.html?spm=a2c4g.154007.0.0.3edd7eb6E90YT4).
//
//		- You are charged only if the value of VerifyResult is true or false and the value of ReasonCode is 0, 1, or 2.
//
//		- Before you call this operation, perform the following operations: Log on to the [Cell Phone Number Service console](https://account.aliyun.com/login/login.htm?oauth_callback=https%3A%2F%2Fdytns.console.aliyun.com%2Foverview%3Fspm%3Da2c4g.608385.0.0.79847f8b3awqUC\\&lang=zh). On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// @param request - CompanyThreeElementsVerificationRequest
//
// @return CompanyThreeElementsVerificationResponse
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

// Summary:
//
// Verifies whether the enterprise name and the business license number belong to the same enterprise. The verification is successful only when the two elements belong to the same enterprise and the business status of the enterprise is Active.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the billing of services related to two-element verification for enterprises. For more information, see [Billing](https://help.aliyun.com/document_detail/154751.html?spm=a2c4g.154007.0.0.3edd7eb6E90YT4).
//
//		- You are charged only if the value of VerifyResult is true or false and the value of ReasonCode is 0 or 1.
//
//		- Before you call this operation, perform the following operations: Log on to the [Cell Phone Number Service console](https://account.aliyun.com/login/login.htm?oauth_callback=https%3A%2F%2Fdytns.console.aliyun.com%2Foverview%3Fspm%3Da2c4g.608385.0.0.79847f8b3awqUC\\&lang=zh). On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// @param request - CompanyTwoElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompanyTwoElementsVerificationResponse
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

// Summary:
//
// Verifies whether the enterprise name and the business license number belong to the same enterprise. The verification is successful only when the two elements belong to the same enterprise and the business status of the enterprise is Active.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the billing of services related to two-element verification for enterprises. For more information, see [Billing](https://help.aliyun.com/document_detail/154751.html?spm=a2c4g.154007.0.0.3edd7eb6E90YT4).
//
//		- You are charged only if the value of VerifyResult is true or false and the value of ReasonCode is 0 or 1.
//
//		- Before you call this operation, perform the following operations: Log on to the [Cell Phone Number Service console](https://account.aliyun.com/login/login.htm?oauth_callback=https%3A%2F%2Fdytns.console.aliyun.com%2Foverview%3Fspm%3Da2c4g.608385.0.0.79847f8b3awqUC\\&lang=zh). On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// @param request - CompanyTwoElementsVerificationRequest
//
// @return CompanyTwoElementsVerificationResponse
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

// Summary:
//
// Predicts whether a phone number is a nonexistent number by using AI algorithms.
//
// Description:
//
//	  You can call this operation to verify whether a phone number is a nonexistent number. When you call this operation to verify a number, the system charges you CNY 0.01 per verification based on the number of verifications. **Before you call this operation, make sure that you are familiar with the billing of Cell Phone Number Service.**
//
//		- You are charged only if the value of Code is OK and the value of Status is not UNKNOWN.
//
//		- The prediction is not strictly accurate because Cell Phone Number Service predicts the nonexistent number probability by using AI algorithms. The accuracy rate of the prediction and the recall rate of empty numbers are about 95%. **Pay attention to this point when you call this operation**.
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// ### [](#)Authorization information
//
// By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// @param request - DescribeEmptyNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEmptyNumberResponse
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

// Summary:
//
// Predicts whether a phone number is a nonexistent number by using AI algorithms.
//
// Description:
//
//	  You can call this operation to verify whether a phone number is a nonexistent number. When you call this operation to verify a number, the system charges you CNY 0.01 per verification based on the number of verifications. **Before you call this operation, make sure that you are familiar with the billing of Cell Phone Number Service.**
//
//		- You are charged only if the value of Code is OK and the value of Status is not UNKNOWN.
//
//		- The prediction is not strictly accurate because Cell Phone Number Service predicts the nonexistent number probability by using AI algorithms. The accuracy rate of the prediction and the recall rate of empty numbers are about 95%. **Pay attention to this point when you call this operation**.
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// ### [](#)Authorization information
//
// By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// @param request - DescribeEmptyNumberRequest
//
// @return DescribeEmptyNumberResponse
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

// Summary:
//
// 号码归属地查询v2
//
// @param request - DescribeMobileOperatorAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMobileOperatorAttributeResponse
func (client *Client) DescribeMobileOperatorAttributeWithOptions(request *DescribeMobileOperatorAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeMobileOperatorAttributeResponse, _err error) {
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
		Action:      tea.String("DescribeMobileOperatorAttribute"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMobileOperatorAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 号码归属地查询v2
//
// @param request - DescribeMobileOperatorAttributeRequest
//
// @return DescribeMobileOperatorAttributeResponse
func (client *Client) DescribeMobileOperatorAttribute(request *DescribeMobileOperatorAttributeRequest) (_result *DescribeMobileOperatorAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMobileOperatorAttributeResponse{}
	_body, _err := client.DescribeMobileOperatorAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 号码分析实时查询蚂蚁
//
// @param request - DescribePhoneNumberAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberAnalysisResponse
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

// Summary:
//
// 号码分析实时查询蚂蚁
//
// @param request - DescribePhoneNumberAnalysisRequest
//
// @return DescribePhoneNumberAnalysisResponse
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

// Summary:
//
// Obtains the analysis results of a phone number.
//
// Description:
//
// Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the Labels page, find the label that you want to use, click Activate Now, enter the required information, and then submit your application. After your application is approved, you can use the label. Before you call this operation, make sure that you are familiar with the billing of Cell Phone Number Service.
//
// @param request - DescribePhoneNumberAnalysisAIRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberAnalysisAIResponse
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

// Summary:
//
// Obtains the analysis results of a phone number.
//
// Description:
//
// Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the Labels page, find the label that you want to use, click Activate Now, enter the required information, and then submit your application. After your application is approved, you can use the label. Before you call this operation, make sure that you are familiar with the billing of Cell Phone Number Service.
//
// @param request - DescribePhoneNumberAnalysisAIRequest
//
// @return DescribePhoneNumberAnalysisAIResponse
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

// Summary:
//
// 号码分析服务pai供应商批量查询接口
//
// @param request - DescribePhoneNumberAnalysisPaiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberAnalysisPaiResponse
func (client *Client) DescribePhoneNumberAnalysisPaiWithOptions(request *DescribePhoneNumberAnalysisPaiRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberAnalysisPaiResponse, _err error) {
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
		Action:      tea.String("DescribePhoneNumberAnalysisPai"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberAnalysisPaiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 号码分析服务pai供应商批量查询接口
//
// @param request - DescribePhoneNumberAnalysisPaiRequest
//
// @return DescribePhoneNumberAnalysisPaiResponse
func (client *Client) DescribePhoneNumberAnalysisPai(request *DescribePhoneNumberAnalysisPaiRequest) (_result *DescribePhoneNumberAnalysisPaiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberAnalysisPaiResponse{}
	_body, _err := client.DescribePhoneNumberAnalysisPaiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 泛行业人群筛选
//
// @param request - DescribePhoneNumberAnalysisTransparentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberAnalysisTransparentResponse
func (client *Client) DescribePhoneNumberAnalysisTransparentWithOptions(request *DescribePhoneNumberAnalysisTransparentRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberAnalysisTransparentResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.NumberType)) {
		query["NumberType"] = request.NumberType
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
		Action:      tea.String("DescribePhoneNumberAnalysisTransparent"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberAnalysisTransparentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 泛行业人群筛选
//
// @param request - DescribePhoneNumberAnalysisTransparentRequest
//
// @return DescribePhoneNumberAnalysisTransparentResponse
func (client *Client) DescribePhoneNumberAnalysisTransparent(request *DescribePhoneNumberAnalysisTransparentRequest) (_result *DescribePhoneNumberAnalysisTransparentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberAnalysisTransparentResponse{}
	_body, _err := client.DescribePhoneNumberAnalysisTransparentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribePhoneNumberAttribute is deprecated, please use Dytnsapi::2020-02-17::DescribePhoneNumberOperatorAttribute instead.
//
// Summary:
//
// Queries the carrier, registration location, and mobile number portability information of a phone number.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 2,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePhoneNumberAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberAttributeResponse
// Deprecated
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

// Deprecated: OpenAPI DescribePhoneNumberAttribute is deprecated, please use Dytnsapi::2020-02-17::DescribePhoneNumberOperatorAttribute instead.
//
// Summary:
//
// Queries the carrier, registration location, and mobile number portability information of a phone number.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 2,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePhoneNumberAttributeRequest
//
// @return DescribePhoneNumberAttributeResponse
// Deprecated
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

// Summary:
//
// Queries the usage period of a phone number of a user.
//
// Description:
//
//	  Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
//		- Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 200 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePhoneNumberOnlineTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberOnlineTimeResponse
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

// Summary:
//
// Queries the usage period of a phone number of a user.
//
// Description:
//
//	  Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
//		- Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 200 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePhoneNumberOnlineTimeRequest
//
// @return DescribePhoneNumberOnlineTimeResponse
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

// Summary:
//
// Queries the attribute information about a phone number, including the registration province, registration city, basic carrier (such as China Mobile, China Unicom, China Telecom, or China Broadnet), reseller of mobile communications services (such as Alibaba Communications), mobile number portability, and the number segment to which the phone number belongs.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154008.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- You can call this operation to obtain the carrier, registration location, and mobile number portability information about a phone number. You can query phone numbers in **plaintext*	- and phone numbers that are encrypted by using **MD5*	- and **SHA256**.
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// @param request - DescribePhoneNumberOperatorAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberOperatorAttributeResponse
func (client *Client) DescribePhoneNumberOperatorAttributeWithOptions(request *DescribePhoneNumberOperatorAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberOperatorAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.FlowName)) {
		query["FlowName"] = request.FlowName
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

	if !tea.BoolValue(util.IsUnset(request.ResultCount)) {
		query["ResultCount"] = request.ResultCount
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

// Summary:
//
// Queries the attribute information about a phone number, including the registration province, registration city, basic carrier (such as China Mobile, China Unicom, China Telecom, or China Broadnet), reseller of mobile communications services (such as Alibaba Communications), mobile number portability, and the number segment to which the phone number belongs.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154008.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- You can call this operation to obtain the carrier, registration location, and mobile number portability information about a phone number. You can query phone numbers in **plaintext*	- and phone numbers that are encrypted by using **MD5*	- and **SHA256**.
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// @param request - DescribePhoneNumberOperatorAttributeRequest
//
// @return DescribePhoneNumberOperatorAttributeResponse
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

// Summary:
//
// 号码归属服务(包年包月客户专用)
//
// @param request - DescribePhoneNumberOperatorAttributeAnnualRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberOperatorAttributeAnnualResponse
func (client *Client) DescribePhoneNumberOperatorAttributeAnnualWithOptions(request *DescribePhoneNumberOperatorAttributeAnnualRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberOperatorAttributeAnnualResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		query["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.Number)) {
		query["Number"] = request.Number
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberOperatorAttributeAnnual"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberOperatorAttributeAnnualResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 号码归属服务(包年包月客户专用)
//
// @param request - DescribePhoneNumberOperatorAttributeAnnualRequest
//
// @return DescribePhoneNumberOperatorAttributeAnnualResponse
func (client *Client) DescribePhoneNumberOperatorAttributeAnnual(request *DescribePhoneNumberOperatorAttributeAnnualRequest) (_result *DescribePhoneNumberOperatorAttributeAnnualResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberOperatorAttributeAnnualResponse{}
	_body, _err := client.DescribePhoneNumberOperatorAttributeAnnualWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 号码归属服务(包年包月客户专用)
//
// @param request - DescribePhoneNumberOperatorAttributeAnnualUseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberOperatorAttributeAnnualUseResponse
func (client *Client) DescribePhoneNumberOperatorAttributeAnnualUseWithOptions(request *DescribePhoneNumberOperatorAttributeAnnualUseRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberOperatorAttributeAnnualUseResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberOperatorAttributeAnnualUse"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberOperatorAttributeAnnualUseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 号码归属服务(包年包月客户专用)
//
// @param request - DescribePhoneNumberOperatorAttributeAnnualUseRequest
//
// @return DescribePhoneNumberOperatorAttributeAnnualUseResponse
func (client *Client) DescribePhoneNumberOperatorAttributeAnnualUse(request *DescribePhoneNumberOperatorAttributeAnnualUseRequest) (_result *DescribePhoneNumberOperatorAttributeAnnualUseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberOperatorAttributeAnnualUseResponse{}
	_body, _err := client.DescribePhoneNumberOperatorAttributeAnnualUseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 风险用户评分
//
// @param request - DescribePhoneNumberRiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberRiskResponse
func (client *Client) DescribePhoneNumberRiskWithOptions(request *DescribePhoneNumberRiskRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberRiskResponse, _err error) {
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
		Action:      tea.String("DescribePhoneNumberRisk"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberRiskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 风险用户评分
//
// @param request - DescribePhoneNumberRiskRequest
//
// @return DescribePhoneNumberRiskResponse
func (client *Client) DescribePhoneNumberRisk(request *DescribePhoneNumberRiskRequest) (_result *DescribePhoneNumberRiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberRiskResponse{}
	_body, _err := client.DescribePhoneNumberRiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies whether a phone number is a reassigned phone number by calling this operation.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- You are charged for phone number verifications only if the value of Code is OK and the value of VerifyResult is not 0.
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ## [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// ## [](#)Authorization information
//
// By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// @param request - DescribePhoneTwiceTelVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneTwiceTelVerifyResponse
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

// Summary:
//
// Verifies whether a phone number is a reassigned phone number by calling this operation.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- You are charged for phone number verifications only if the value of Code is OK and the value of VerifyResult is not 0.
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ## [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// ## [](#)Authorization information
//
// By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// @param request - DescribePhoneTwiceTelVerifyRequest
//
// @return DescribePhoneTwiceTelVerifyResponse
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

// Summary:
//
// 获取UAID申请Token所需的签名字段
//
// @param request - GetUAIDApplyTokenSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUAIDApplyTokenSignResponse
func (client *Client) GetUAIDApplyTokenSignWithOptions(request *GetUAIDApplyTokenSignRequest, runtime *util.RuntimeOptions) (_result *GetUAIDApplyTokenSignResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.Format)) {
		query["Format"] = request.Format
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ParamKey)) {
		query["ParamKey"] = request.ParamKey
	}

	if !tea.BoolValue(util.IsUnset(request.ParamStr)) {
		query["ParamStr"] = request.ParamStr
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Time)) {
		query["Time"] = request.Time
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUAIDApplyTokenSign"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUAIDApplyTokenSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取UAID申请Token所需的签名字段
//
// @param request - GetUAIDApplyTokenSignRequest
//
// @return GetUAIDApplyTokenSignResponse
func (client *Client) GetUAIDApplyTokenSign(request *GetUAIDApplyTokenSignRequest) (_result *GetUAIDApplyTokenSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUAIDApplyTokenSignResponse{}
	_body, _err := client.GetUAIDApplyTokenSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取号码采集服务申请Token所需的签名字段
//
// @param request - GetUAIDConversionSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUAIDConversionSignResponse
func (client *Client) GetUAIDConversionSignWithOptions(request *GetUAIDConversionSignRequest, runtime *util.RuntimeOptions) (_result *GetUAIDConversionSignResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.Format)) {
		query["Format"] = request.Format
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ParamKey)) {
		query["ParamKey"] = request.ParamKey
	}

	if !tea.BoolValue(util.IsUnset(request.ParamStr)) {
		query["ParamStr"] = request.ParamStr
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Time)) {
		query["Time"] = request.Time
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUAIDConversionSign"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUAIDConversionSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取号码采集服务申请Token所需的签名字段
//
// @param request - GetUAIDConversionSignRequest
//
// @return GetUAIDConversionSignResponse
func (client *Client) GetUAIDConversionSign(request *GetUAIDConversionSignRequest) (_result *GetUAIDConversionSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUAIDConversionSignResponse{}
	_body, _err := client.GetUAIDConversionSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Filters invalid phone numbers.
//
// Description:
//
// Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - InvalidPhoneNumberFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvalidPhoneNumberFilterResponse
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

// Summary:
//
// Filters invalid phone numbers.
//
// Description:
//
// Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - InvalidPhoneNumberFilterRequest
//
// @return InvalidPhoneNumberFilterResponse
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

// @param request - PhoneNumberConvertServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberConvertServiceResponse
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

// @param request - PhoneNumberConvertServiceRequest
//
// @return PhoneNumberConvertServiceResponse
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

// Summary:
//
// Encrypts the original phone number into a virtual number that starts with 140. Cell Phone Number Service integrates the communications services provided by Alibaba Cloud. This allows you to initiate a call by using a virtual number that starts with 140.
//
// Description:
//
// Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberEncryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberEncryptResponse
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

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
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

// Summary:
//
// Encrypts the original phone number into a virtual number that starts with 140. Cell Phone Number Service integrates the communications services provided by Alibaba Cloud. This allows you to initiate a call by using a virtual number that starts with 140.
//
// Description:
//
// Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberEncryptRequest
//
// @return PhoneNumberEncryptResponse
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

// Summary:
//
// Queries the real-time service state of a phone number. The state includes NORMAL, SHUTDOWN, and NOT_EXIST. You can choose an encryption method for your phone number query, including plaintext, MD5, and SHA256.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 300 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberStatusForAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberStatusForAccountResponse
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

// Summary:
//
// Queries the real-time service state of a phone number. The state includes NORMAL, SHUTDOWN, and NOT_EXIST. You can choose an encryption method for your phone number query, including plaintext, MD5, and SHA256.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 300 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberStatusForAccountRequest
//
// @return PhoneNumberStatusForAccountResponse
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

// Summary:
//
// Queries the real-time service state of a phone number. The state includes NORMAL, SHUTDOWN, and NOT_EXIST. You can choose an encryption method for your phone number query, including plaintext, MD5, and SHA256.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 300 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberStatusForPublicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberStatusForPublicResponse
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

// Summary:
//
// Queries the real-time service state of a phone number. The state includes NORMAL, SHUTDOWN, and NOT_EXIST. You can choose an encryption method for your phone number query, including plaintext, MD5, and SHA256.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 300 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberStatusForPublicRequest
//
// @return PhoneNumberStatusForPublicResponse
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

// Summary:
//
// Queries the real-time service state of a phone number. The state includes NORMAL, SHUTDOWN, and NOT_EXIST. You can choose an encryption method for your phone number query, including plaintext, MD5, and SHA256.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 300 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberStatusForRealRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberStatusForRealResponse
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

// Summary:
//
// Queries the real-time service state of a phone number. The state includes NORMAL, SHUTDOWN, and NOT_EXIST. You can choose an encryption method for your phone number query, including plaintext, MD5, and SHA256.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 300 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberStatusForRealRequest
//
// @return PhoneNumberStatusForRealResponse
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

// Summary:
//
// Queries the real-time service state of a phone number. The state includes NORMAL, SHUTDOWN, and NOT_EXIST. You can choose an encryption method for your phone number query, including plaintext, MD5, and SHA256.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 300 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberStatusForSmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberStatusForSmsResponse
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

// Summary:
//
// Queries the real-time service state of a phone number. The state includes NORMAL, SHUTDOWN, and NOT_EXIST. You can choose an encryption method for your phone number query, including plaintext, MD5, and SHA256.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 300 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberStatusForSmsRequest
//
// @return PhoneNumberStatusForSmsResponse
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

// Summary:
//
// Queries the real-time service state of a phone number. The state includes NORMAL, SHUTDOWN, and NOT_EXIST. You can choose an encryption method for your phone number query, including plaintext, MD5, and SHA256.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 300 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberStatusForVoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberStatusForVoiceResponse
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

// Summary:
//
// Queries the real-time service state of a phone number. The state includes NORMAL, SHUTDOWN, and NOT_EXIST. You can choose an encryption method for your phone number query, including plaintext, MD5, and SHA256.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 300 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberStatusForVoiceRequest
//
// @return PhoneNumberStatusForVoiceResponse
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

// Summary:
//
// Queries available authorization codes.
//
// @param request - QueryAvailableAuthCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAvailableAuthCodeResponse
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

// Summary:
//
// Queries available authorization codes.
//
// @param request - QueryAvailableAuthCodeRequest
//
// @return QueryAvailableAuthCodeResponse
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

// Summary:
//
// 在网时长专用接口
//
// @param request - QueryPhoneNumberOnlineTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPhoneNumberOnlineTimeResponse
func (client *Client) QueryPhoneNumberOnlineTimeWithOptions(request *QueryPhoneNumberOnlineTimeRequest, runtime *util.RuntimeOptions) (_result *QueryPhoneNumberOnlineTimeResponse, _err error) {
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
		Action:      tea.String("QueryPhoneNumberOnlineTime"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPhoneNumberOnlineTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 在网时长专用接口
//
// @param request - QueryPhoneNumberOnlineTimeRequest
//
// @return QueryPhoneNumberOnlineTimeResponse
func (client *Client) QueryPhoneNumberOnlineTime(request *QueryPhoneNumberOnlineTimeRequest) (_result *QueryPhoneNumberOnlineTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPhoneNumberOnlineTimeResponse{}
	_body, _err := client.QueryPhoneNumberOnlineTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 二次号携号转网号码查询
//
// @param request - QueryPhoneTwiceTelVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPhoneTwiceTelVerifyResponse
func (client *Client) QueryPhoneTwiceTelVerifyWithOptions(request *QueryPhoneTwiceTelVerifyRequest, runtime *util.RuntimeOptions) (_result *QueryPhoneTwiceTelVerifyResponse, _err error) {
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
		Action:      tea.String("QueryPhoneTwiceTelVerify"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPhoneTwiceTelVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 二次号携号转网号码查询
//
// @param request - QueryPhoneTwiceTelVerifyRequest
//
// @return QueryPhoneTwiceTelVerifyResponse
func (client *Client) QueryPhoneTwiceTelVerify(request *QueryPhoneTwiceTelVerifyRequest) (_result *QueryPhoneTwiceTelVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPhoneTwiceTelVerifyResponse{}
	_body, _err := client.QueryPhoneTwiceTelVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tag application rules.
//
// @param request - QueryTagApplyRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTagApplyRuleResponse
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

// Summary:
//
// Queries tag application rules.
//
// @param request - QueryTagApplyRuleRequest
//
// @return QueryTagApplyRuleResponse
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

// Summary:
//
// Queries information about tags.
//
// @param request - QueryTagInfoBySelectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTagInfoBySelectionResponse
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

// Summary:
//
// Queries information about tags.
//
// @param request - QueryTagInfoBySelectionRequest
//
// @return QueryTagInfoBySelectionResponse
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

// Summary:
//
// Queries a list of tags by page.
//
// @param request - QueryTagListPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTagListPageResponse
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

// Summary:
//
// Queries a list of tags by page.
//
// @param request - QueryTagListPageRequest
//
// @return QueryTagListPageResponse
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

// Summary:
//
// Queries the usage statistics based on tag IDs.
//
// @param request - QueryUsageStatisticsByTagIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUsageStatisticsByTagIdResponse
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

// Summary:
//
// Queries the usage statistics based on tag IDs.
//
// @param request - QueryUsageStatisticsByTagIdRequest
//
// @return QueryUsageStatisticsByTagIdResponse
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

// Summary:
//
// Verifies whether the name, phone number, and ID card number entered by a user belong to the same user.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
//		- You are charged only if the value of Code is OK and the value of IsConsistent is not 2.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 200 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ThreeElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ThreeElementsVerificationResponse
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

// Summary:
//
// Verifies whether the name, phone number, and ID card number entered by a user belong to the same user.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
//		- You are charged only if the value of Code is OK and the value of IsConsistent is not 2.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 200 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ThreeElementsVerificationRequest
//
// @return ThreeElementsVerificationResponse
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

// Summary:
//
// Verifies whether the name and phone number entered by a user belong to the same user.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
//		- You are charged only if the value of Code is OK and the value of IsConsistent is not 2.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 200 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - TwoElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TwoElementsVerificationResponse
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

// Summary:
//
// Verifies whether the name and phone number entered by a user belong to the same user.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
//		- You are charged only if the value of Code is OK and the value of IsConsistent is not 2.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 200 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - TwoElementsVerificationRequest
//
// @return TwoElementsVerificationResponse
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

// Summary:
//
// # UAID采集
//
// @param request - UAIDCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UAIDCollectionResponse
func (client *Client) UAIDCollectionWithOptions(request *UAIDCollectionRequest, runtime *util.RuntimeOptions) (_result *UAIDCollectionResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		query["Province"] = request.Province
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.UserGrantId)) {
		query["UserGrantId"] = request.UserGrantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UAIDCollection"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UAIDCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # UAID采集
//
// @param request - UAIDCollectionRequest
//
// @return UAIDCollectionResponse
func (client *Client) UAIDCollection(request *UAIDCollectionRequest) (_result *UAIDCollectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UAIDCollectionResponse{}
	_body, _err := client.UAIDCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// uaid号码转换服务
//
// @param request - UAIDConversionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UAIDConversionResponse
func (client *Client) UAIDConversionWithOptions(request *UAIDConversionRequest, runtime *util.RuntimeOptions) (_result *UAIDConversionResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
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

	if !tea.BoolValue(util.IsUnset(request.UaidList)) {
		query["UaidList"] = request.UaidList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UAIDConversion"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UAIDConversionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// uaid号码转换服务
//
// @param request - UAIDConversionRequest
//
// @return UAIDConversionResponse
func (client *Client) UAIDConversion(request *UAIDConversionRequest) (_result *UAIDConversionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UAIDConversionResponse{}
	_body, _err := client.UAIDConversionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取UAID
//
// @param request - UAIDVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UAIDVerificationResponse
func (client *Client) UAIDVerificationWithOptions(request *UAIDVerificationRequest, runtime *util.RuntimeOptions) (_result *UAIDVerificationResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		query["Province"] = request.Province
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.UserGrantId)) {
		query["UserGrantId"] = request.UserGrantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UAIDVerification"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UAIDVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取UAID
//
// @param request - UAIDVerificationRequest
//
// @return UAIDVerificationResponse
func (client *Client) UAIDVerification(request *UAIDVerificationRequest) (_result *UAIDVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UAIDVerificationResponse{}
	_body, _err := client.UAIDVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
