// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type EnterpriseAccountChangeLoginPasswordRequest struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	EncryptPassword *string `json:"EncryptPassword,omitempty" xml:"EncryptPassword,omitempty"`
	OrientedEcId    *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId    *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId    *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnterpriseAccountChangeLoginPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountChangeLoginPasswordRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) SetAppName(v string) *EnterpriseAccountChangeLoginPasswordRequest {
	s.AppName = &v
	return s
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) SetEncryptPassword(v string) *EnterpriseAccountChangeLoginPasswordRequest {
	s.EncryptPassword = &v
	return s
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) SetOrientedEcId(v string) *EnterpriseAccountChangeLoginPasswordRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) SetOrientedLeId(v string) *EnterpriseAccountChangeLoginPasswordRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) SetOrientedNbId(v string) *EnterpriseAccountChangeLoginPasswordRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) SetPk(v string) *EnterpriseAccountChangeLoginPasswordRequest {
	s.Pk = &v
	return s
}

func (s *EnterpriseAccountChangeLoginPasswordRequest) SetRequestId(v string) *EnterpriseAccountChangeLoginPasswordRequest {
	s.RequestId = &v
	return s
}

type EnterpriseAccountChangeLoginPasswordResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Pass      *bool   `json:"Pass,omitempty" xml:"Pass,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountChangeLoginPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountChangeLoginPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountChangeLoginPasswordResponseBody) SetCode(v string) *EnterpriseAccountChangeLoginPasswordResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseAccountChangeLoginPasswordResponseBody) SetMessage(v string) *EnterpriseAccountChangeLoginPasswordResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseAccountChangeLoginPasswordResponseBody) SetPass(v bool) *EnterpriseAccountChangeLoginPasswordResponseBody {
	s.Pass = &v
	return s
}

func (s *EnterpriseAccountChangeLoginPasswordResponseBody) SetRequestId(v string) *EnterpriseAccountChangeLoginPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountChangeLoginPasswordResponseBody) SetSuccess(v bool) *EnterpriseAccountChangeLoginPasswordResponseBody {
	s.Success = &v
	return s
}

type EnterpriseAccountChangeLoginPasswordResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseAccountChangeLoginPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountChangeLoginPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountChangeLoginPasswordResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountChangeLoginPasswordResponse) SetHeaders(v map[string]*string) *EnterpriseAccountChangeLoginPasswordResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseAccountChangeLoginPasswordResponse) SetStatusCode(v int32) *EnterpriseAccountChangeLoginPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseAccountChangeLoginPasswordResponse) SetBody(v *EnterpriseAccountChangeLoginPasswordResponseBody) *EnterpriseAccountChangeLoginPasswordResponse {
	s.Body = v
	return s
}

type EnterpriseAccountChangeSecurityEmailRequest struct {
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// This parameter is required.
	SecurityEmail *string `json:"SecurityEmail,omitempty" xml:"SecurityEmail,omitempty"`
	// This parameter is required.
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s EnterpriseAccountChangeSecurityEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountChangeSecurityEmailRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) SetAppName(v string) *EnterpriseAccountChangeSecurityEmailRequest {
	s.AppName = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) SetOrientedEcId(v string) *EnterpriseAccountChangeSecurityEmailRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) SetOrientedLeId(v string) *EnterpriseAccountChangeSecurityEmailRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) SetOrientedNbId(v string) *EnterpriseAccountChangeSecurityEmailRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) SetPk(v string) *EnterpriseAccountChangeSecurityEmailRequest {
	s.Pk = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) SetRequestId(v string) *EnterpriseAccountChangeSecurityEmailRequest {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) SetSecurityEmail(v string) *EnterpriseAccountChangeSecurityEmailRequest {
	s.SecurityEmail = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) SetVerifyCode(v string) *EnterpriseAccountChangeSecurityEmailRequest {
	s.VerifyCode = &v
	return s
}

type EnterpriseAccountChangeSecurityEmailResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountChangeSecurityEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountChangeSecurityEmailResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountChangeSecurityEmailResponseBody) SetCode(v string) *EnterpriseAccountChangeSecurityEmailResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityEmailResponseBody) SetData(v bool) *EnterpriseAccountChangeSecurityEmailResponseBody {
	s.Data = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityEmailResponseBody) SetMessage(v string) *EnterpriseAccountChangeSecurityEmailResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityEmailResponseBody) SetRequestId(v string) *EnterpriseAccountChangeSecurityEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityEmailResponseBody) SetSuccess(v bool) *EnterpriseAccountChangeSecurityEmailResponseBody {
	s.Success = &v
	return s
}

type EnterpriseAccountChangeSecurityEmailResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseAccountChangeSecurityEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountChangeSecurityEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountChangeSecurityEmailResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountChangeSecurityEmailResponse) SetHeaders(v map[string]*string) *EnterpriseAccountChangeSecurityEmailResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseAccountChangeSecurityEmailResponse) SetStatusCode(v int32) *EnterpriseAccountChangeSecurityEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityEmailResponse) SetBody(v *EnterpriseAccountChangeSecurityEmailResponseBody) *EnterpriseAccountChangeSecurityEmailResponse {
	s.Body = v
	return s
}

type EnterpriseAccountChangeSecurityMobileRequest struct {
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
	// This parameter is required.
	NewMobile    *string `json:"NewMobile,omitempty" xml:"NewMobile,omitempty"`
	OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// This parameter is required.
	VerificationCode *string `json:"VerificationCode,omitempty" xml:"VerificationCode,omitempty"`
}

func (s EnterpriseAccountChangeSecurityMobileRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountChangeSecurityMobileRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetAppName(v string) *EnterpriseAccountChangeSecurityMobileRequest {
	s.AppName = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetEncryptTicket(v string) *EnterpriseAccountChangeSecurityMobileRequest {
	s.EncryptTicket = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetNewMobile(v string) *EnterpriseAccountChangeSecurityMobileRequest {
	s.NewMobile = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetOrientedEcId(v string) *EnterpriseAccountChangeSecurityMobileRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetOrientedLeId(v string) *EnterpriseAccountChangeSecurityMobileRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetOrientedNbId(v string) *EnterpriseAccountChangeSecurityMobileRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetPk(v string) *EnterpriseAccountChangeSecurityMobileRequest {
	s.Pk = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetRequestId(v string) *EnterpriseAccountChangeSecurityMobileRequest {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetVerificationCode(v string) *EnterpriseAccountChangeSecurityMobileRequest {
	s.VerificationCode = &v
	return s
}

type EnterpriseAccountChangeSecurityMobileResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountChangeSecurityMobileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountChangeSecurityMobileResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountChangeSecurityMobileResponseBody) SetCode(v string) *EnterpriseAccountChangeSecurityMobileResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityMobileResponseBody) SetData(v bool) *EnterpriseAccountChangeSecurityMobileResponseBody {
	s.Data = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityMobileResponseBody) SetMessage(v string) *EnterpriseAccountChangeSecurityMobileResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityMobileResponseBody) SetRequestId(v string) *EnterpriseAccountChangeSecurityMobileResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityMobileResponseBody) SetSuccess(v bool) *EnterpriseAccountChangeSecurityMobileResponseBody {
	s.Success = &v
	return s
}

type EnterpriseAccountChangeSecurityMobileResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseAccountChangeSecurityMobileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountChangeSecurityMobileResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountChangeSecurityMobileResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountChangeSecurityMobileResponse) SetHeaders(v map[string]*string) *EnterpriseAccountChangeSecurityMobileResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseAccountChangeSecurityMobileResponse) SetStatusCode(v int32) *EnterpriseAccountChangeSecurityMobileResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseAccountChangeSecurityMobileResponse) SetBody(v *EnterpriseAccountChangeSecurityMobileResponseBody) *EnterpriseAccountChangeSecurityMobileResponse {
	s.Body = v
	return s
}

type EnterpriseAccountQueryAccountGrantedRolesRequest struct {
	AppName          *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	IsOpenApi        *bool   `json:"IsOpenApi,omitempty" xml:"IsOpenApi,omitempty"`
	OrientedEcId     *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId     *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId     *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	Pk               *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	ShowCompleteInfo *bool   `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
}

func (s EnterpriseAccountQueryAccountGrantedRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountQueryAccountGrantedRolesRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) SetAppName(v string) *EnterpriseAccountQueryAccountGrantedRolesRequest {
	s.AppName = &v
	return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) SetIsOpenApi(v bool) *EnterpriseAccountQueryAccountGrantedRolesRequest {
	s.IsOpenApi = &v
	return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) SetOrientedEcId(v string) *EnterpriseAccountQueryAccountGrantedRolesRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) SetOrientedLeId(v string) *EnterpriseAccountQueryAccountGrantedRolesRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) SetOrientedNbId(v string) *EnterpriseAccountQueryAccountGrantedRolesRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) SetPk(v string) *EnterpriseAccountQueryAccountGrantedRolesRequest {
	s.Pk = &v
	return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesRequest) SetShowCompleteInfo(v bool) *EnterpriseAccountQueryAccountGrantedRolesRequest {
	s.ShowCompleteInfo = &v
	return s
}

type EnterpriseAccountQueryAccountGrantedRolesResponseBody struct {
	Code      *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*EnterpriseAccountQueryAccountGrantedRolesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountQueryAccountGrantedRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountQueryAccountGrantedRolesResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBody) SetCode(v string) *EnterpriseAccountQueryAccountGrantedRolesResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBody) SetData(v []*EnterpriseAccountQueryAccountGrantedRolesResponseBodyData) *EnterpriseAccountQueryAccountGrantedRolesResponseBody {
	s.Data = v
	return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBody) SetMessage(v string) *EnterpriseAccountQueryAccountGrantedRolesResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBody) SetRequestId(v string) *EnterpriseAccountQueryAccountGrantedRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBody) SetSuccess(v bool) *EnterpriseAccountQueryAccountGrantedRolesResponseBody {
	s.Success = &v
	return s
}

type EnterpriseAccountQueryAccountGrantedRolesResponseBodyData struct {
	BizRoleCode *string `json:"BizRoleCode,omitempty" xml:"BizRoleCode,omitempty"`
	BizRoleName *string `json:"BizRoleName,omitempty" xml:"BizRoleName,omitempty"`
}

func (s EnterpriseAccountQueryAccountGrantedRolesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountQueryAccountGrantedRolesResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBodyData) SetBizRoleCode(v string) *EnterpriseAccountQueryAccountGrantedRolesResponseBodyData {
	s.BizRoleCode = &v
	return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBodyData) SetBizRoleName(v string) *EnterpriseAccountQueryAccountGrantedRolesResponseBodyData {
	s.BizRoleName = &v
	return s
}

type EnterpriseAccountQueryAccountGrantedRolesResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseAccountQueryAccountGrantedRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountQueryAccountGrantedRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountQueryAccountGrantedRolesResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponse) SetHeaders(v map[string]*string) *EnterpriseAccountQueryAccountGrantedRolesResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponse) SetStatusCode(v int32) *EnterpriseAccountQueryAccountGrantedRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponse) SetBody(v *EnterpriseAccountQueryAccountGrantedRolesResponseBody) *EnterpriseAccountQueryAccountGrantedRolesResponse {
	s.Body = v
	return s
}

type EnterpriseAccountQueryAccountsInfoRequest struct {
	EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OrientedEcId  *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId  *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId  *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	// This parameter is required.
	PksJson *string `json:"PksJson,omitempty" xml:"PksJson,omitempty"`
	// This parameter is required.
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShowCompleteInfo *bool   `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
}

func (s EnterpriseAccountQueryAccountsInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountQueryAccountsInfoRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetEncryptTicket(v string) *EnterpriseAccountQueryAccountsInfoRequest {
	s.EncryptTicket = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetMaxResults(v int32) *EnterpriseAccountQueryAccountsInfoRequest {
	s.MaxResults = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetNextToken(v string) *EnterpriseAccountQueryAccountsInfoRequest {
	s.NextToken = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetOrientedEcId(v string) *EnterpriseAccountQueryAccountsInfoRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetOrientedLeId(v string) *EnterpriseAccountQueryAccountsInfoRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetOrientedNbId(v string) *EnterpriseAccountQueryAccountsInfoRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetPksJson(v string) *EnterpriseAccountQueryAccountsInfoRequest {
	s.PksJson = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetRequestId(v string) *EnterpriseAccountQueryAccountsInfoRequest {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetShowCompleteInfo(v bool) *EnterpriseAccountQueryAccountsInfoRequest {
	s.ShowCompleteInfo = &v
	return s
}

type EnterpriseAccountQueryAccountsInfoResponseBody struct {
	AccountInfoDtoList []*EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList `json:"AccountInfoDtoList,omitempty" xml:"AccountInfoDtoList,omitempty" type:"Repeated"`
	Code               *string                                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	MaxResults         *int32                                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message            *string                                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken          *string                                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId          *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountQueryAccountsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountQueryAccountsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) SetAccountInfoDtoList(v []*EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) *EnterpriseAccountQueryAccountsInfoResponseBody {
	s.AccountInfoDtoList = v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) SetCode(v string) *EnterpriseAccountQueryAccountsInfoResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) SetMaxResults(v int32) *EnterpriseAccountQueryAccountsInfoResponseBody {
	s.MaxResults = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) SetMessage(v string) *EnterpriseAccountQueryAccountsInfoResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) SetNextToken(v string) *EnterpriseAccountQueryAccountsInfoResponseBody {
	s.NextToken = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) SetRequestId(v string) *EnterpriseAccountQueryAccountsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) SetSuccess(v bool) *EnterpriseAccountQueryAccountsInfoResponseBody {
	s.Success = &v
	return s
}

type EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList struct {
	Alias                 *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	BelongToMasterAccount *bool   `json:"BelongToMasterAccount,omitempty" xml:"BelongToMasterAccount,omitempty"`
	Email                 *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EnterpriseEcId        *string `json:"EnterpriseEcId,omitempty" xml:"EnterpriseEcId,omitempty"`
	EnterpriseEntityId    *string `json:"EnterpriseEntityId,omitempty" xml:"EnterpriseEntityId,omitempty"`
	EnterpriseLicenseNo   *string `json:"EnterpriseLicenseNo,omitempty" xml:"EnterpriseLicenseNo,omitempty"`
	EnterpriseName        *string `json:"EnterpriseName,omitempty" xml:"EnterpriseName,omitempty"`
	EnterpriseNbId        *string `json:"EnterpriseNbId,omitempty" xml:"EnterpriseNbId,omitempty"`
	FreezeLogin           *bool   `json:"FreezeLogin,omitempty" xml:"FreezeLogin,omitempty"`
	LoginId               *string `json:"LoginId,omitempty" xml:"LoginId,omitempty"`
	ManageInviteTimeStamp *string `json:"ManageInviteTimeStamp,omitempty" xml:"ManageInviteTimeStamp,omitempty"`
	Pk                    *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	SecurityMobile        *string `json:"SecurityMobile,omitempty" xml:"SecurityMobile,omitempty"`
}

func (s EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetAlias(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
	s.Alias = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetBelongToMasterAccount(v bool) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
	s.BelongToMasterAccount = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetEmail(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
	s.Email = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetEnterpriseEcId(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
	s.EnterpriseEcId = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetEnterpriseEntityId(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
	s.EnterpriseEntityId = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetEnterpriseLicenseNo(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
	s.EnterpriseLicenseNo = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetEnterpriseName(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
	s.EnterpriseName = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetEnterpriseNbId(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
	s.EnterpriseNbId = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetFreezeLogin(v bool) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
	s.FreezeLogin = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetLoginId(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
	s.LoginId = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetManageInviteTimeStamp(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
	s.ManageInviteTimeStamp = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetPk(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
	s.Pk = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetSecurityMobile(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
	s.SecurityMobile = &v
	return s
}

type EnterpriseAccountQueryAccountsInfoResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseAccountQueryAccountsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountQueryAccountsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountQueryAccountsInfoResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountQueryAccountsInfoResponse) SetHeaders(v map[string]*string) *EnterpriseAccountQueryAccountsInfoResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponse) SetStatusCode(v int32) *EnterpriseAccountQueryAccountsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponse) SetBody(v *EnterpriseAccountQueryAccountsInfoResponseBody) *EnterpriseAccountQueryAccountsInfoResponse {
	s.Body = v
	return s
}

type EnterpriseAccountQueryLoginSettingsRequest struct {
	AppName          *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	IsOpenApi        *bool   `json:"IsOpenApi,omitempty" xml:"IsOpenApi,omitempty"`
	OrientedEcId     *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId     *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId     *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	Pk               *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	ShowCompleteInfo *bool   `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
}

func (s EnterpriseAccountQueryLoginSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountQueryLoginSettingsRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) SetAppName(v string) *EnterpriseAccountQueryLoginSettingsRequest {
	s.AppName = &v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) SetIsOpenApi(v bool) *EnterpriseAccountQueryLoginSettingsRequest {
	s.IsOpenApi = &v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) SetOrientedEcId(v string) *EnterpriseAccountQueryLoginSettingsRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) SetOrientedLeId(v string) *EnterpriseAccountQueryLoginSettingsRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) SetOrientedNbId(v string) *EnterpriseAccountQueryLoginSettingsRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) SetPk(v string) *EnterpriseAccountQueryLoginSettingsRequest {
	s.Pk = &v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsRequest) SetShowCompleteInfo(v bool) *EnterpriseAccountQueryLoginSettingsRequest {
	s.ShowCompleteInfo = &v
	return s
}

type EnterpriseAccountQueryLoginSettingsResponseBody struct {
	Code      *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *EnterpriseAccountQueryLoginSettingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountQueryLoginSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountQueryLoginSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBody) SetCode(v string) *EnterpriseAccountQueryLoginSettingsResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBody) SetData(v *EnterpriseAccountQueryLoginSettingsResponseBodyData) *EnterpriseAccountQueryLoginSettingsResponseBody {
	s.Data = v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBody) SetMessage(v string) *EnterpriseAccountQueryLoginSettingsResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBody) SetRequestId(v string) *EnterpriseAccountQueryLoginSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBody) SetSuccess(v bool) *EnterpriseAccountQueryLoginSettingsResponseBody {
	s.Success = &v
	return s
}

type EnterpriseAccountQueryLoginSettingsResponseBodyData struct {
	IpMaskDto                 *EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto      `json:"IpMaskDto,omitempty" xml:"IpMaskDto,omitempty" type:"Struct"`
	MfaBindStatus             *string                                                            `json:"MfaBindStatus,omitempty" xml:"MfaBindStatus,omitempty"`
	RiskControlDto            *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto `json:"RiskControlDto,omitempty" xml:"RiskControlDto,omitempty" type:"Struct"`
	SecurityMobileLoginStatus *string                                                            `json:"SecurityMobileLoginStatus,omitempty" xml:"SecurityMobileLoginStatus,omitempty"`
	SessionExpireTime         *int32                                                             `json:"SessionExpireTime,omitempty" xml:"SessionExpireTime,omitempty"`
}

func (s EnterpriseAccountQueryLoginSettingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountQueryLoginSettingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyData) SetIpMaskDto(v *EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto) *EnterpriseAccountQueryLoginSettingsResponseBodyData {
	s.IpMaskDto = v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyData) SetMfaBindStatus(v string) *EnterpriseAccountQueryLoginSettingsResponseBodyData {
	s.MfaBindStatus = &v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyData) SetRiskControlDto(v *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) *EnterpriseAccountQueryLoginSettingsResponseBodyData {
	s.RiskControlDto = v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyData) SetSecurityMobileLoginStatus(v string) *EnterpriseAccountQueryLoginSettingsResponseBodyData {
	s.SecurityMobileLoginStatus = &v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyData) SetSessionExpireTime(v int32) *EnterpriseAccountQueryLoginSettingsResponseBodyData {
	s.SessionExpireTime = &v
	return s
}

type EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto struct {
	IpMaskEnabledStatus *string   `json:"IpMaskEnabledStatus,omitempty" xml:"IpMaskEnabledStatus,omitempty"`
	IpMasks             []*string `json:"IpMasks,omitempty" xml:"IpMasks,omitempty" type:"Repeated"`
}

func (s EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto) SetIpMaskEnabledStatus(v string) *EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto {
	s.IpMaskEnabledStatus = &v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto) SetIpMasks(v []*string) *EnterpriseAccountQueryLoginSettingsResponseBodyDataIpMaskDto {
	s.IpMasks = v
	return s
}

type EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto struct {
	ProtectLevel       *string `json:"ProtectLevel,omitempty" xml:"ProtectLevel,omitempty"`
	RiskControlEnabled *bool   `json:"RiskControlEnabled,omitempty" xml:"RiskControlEnabled,omitempty"`
	VerifyDetail       *string `json:"VerifyDetail,omitempty" xml:"VerifyDetail,omitempty"`
	VerifyType         *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) SetProtectLevel(v string) *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto {
	s.ProtectLevel = &v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) SetRiskControlEnabled(v bool) *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto {
	s.RiskControlEnabled = &v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) SetVerifyDetail(v string) *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto {
	s.VerifyDetail = &v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto) SetVerifyType(v string) *EnterpriseAccountQueryLoginSettingsResponseBodyDataRiskControlDto {
	s.VerifyType = &v
	return s
}

type EnterpriseAccountQueryLoginSettingsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseAccountQueryLoginSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountQueryLoginSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountQueryLoginSettingsResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountQueryLoginSettingsResponse) SetHeaders(v map[string]*string) *EnterpriseAccountQueryLoginSettingsResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponse) SetStatusCode(v int32) *EnterpriseAccountQueryLoginSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponse) SetBody(v *EnterpriseAccountQueryLoginSettingsResponseBody) *EnterpriseAccountQueryLoginSettingsResponse {
	s.Body = v
	return s
}

type EnterpriseAccountRemoveMfaRequest struct {
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnterpriseAccountRemoveMfaRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountRemoveMfaRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountRemoveMfaRequest) SetAppName(v string) *EnterpriseAccountRemoveMfaRequest {
	s.AppName = &v
	return s
}

func (s *EnterpriseAccountRemoveMfaRequest) SetOrientedEcId(v string) *EnterpriseAccountRemoveMfaRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseAccountRemoveMfaRequest) SetOrientedLeId(v string) *EnterpriseAccountRemoveMfaRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseAccountRemoveMfaRequest) SetOrientedNbId(v string) *EnterpriseAccountRemoveMfaRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseAccountRemoveMfaRequest) SetPk(v string) *EnterpriseAccountRemoveMfaRequest {
	s.Pk = &v
	return s
}

func (s *EnterpriseAccountRemoveMfaRequest) SetRequestId(v string) *EnterpriseAccountRemoveMfaRequest {
	s.RequestId = &v
	return s
}

type EnterpriseAccountRemoveMfaResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountRemoveMfaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountRemoveMfaResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountRemoveMfaResponseBody) SetCode(v string) *EnterpriseAccountRemoveMfaResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseAccountRemoveMfaResponseBody) SetData(v bool) *EnterpriseAccountRemoveMfaResponseBody {
	s.Data = &v
	return s
}

func (s *EnterpriseAccountRemoveMfaResponseBody) SetMessage(v string) *EnterpriseAccountRemoveMfaResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseAccountRemoveMfaResponseBody) SetRequestId(v string) *EnterpriseAccountRemoveMfaResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountRemoveMfaResponseBody) SetSuccess(v bool) *EnterpriseAccountRemoveMfaResponseBody {
	s.Success = &v
	return s
}

type EnterpriseAccountRemoveMfaResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseAccountRemoveMfaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountRemoveMfaResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountRemoveMfaResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountRemoveMfaResponse) SetHeaders(v map[string]*string) *EnterpriseAccountRemoveMfaResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseAccountRemoveMfaResponse) SetStatusCode(v int32) *EnterpriseAccountRemoveMfaResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseAccountRemoveMfaResponse) SetBody(v *EnterpriseAccountRemoveMfaResponseBody) *EnterpriseAccountRemoveMfaResponse {
	s.Body = v
	return s
}

type EnterpriseAccountSeparateEaRequest struct {
	EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
	OrientedEcId  *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId  *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId  *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	Pk            *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s EnterpriseAccountSeparateEaRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountSeparateEaRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountSeparateEaRequest) SetEncryptTicket(v string) *EnterpriseAccountSeparateEaRequest {
	s.EncryptTicket = &v
	return s
}

func (s *EnterpriseAccountSeparateEaRequest) SetOrientedEcId(v string) *EnterpriseAccountSeparateEaRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseAccountSeparateEaRequest) SetOrientedLeId(v string) *EnterpriseAccountSeparateEaRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseAccountSeparateEaRequest) SetOrientedNbId(v string) *EnterpriseAccountSeparateEaRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseAccountSeparateEaRequest) SetPk(v string) *EnterpriseAccountSeparateEaRequest {
	s.Pk = &v
	return s
}

type EnterpriseAccountSeparateEaResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountSeparateEaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountSeparateEaResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountSeparateEaResponseBody) SetCode(v string) *EnterpriseAccountSeparateEaResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseAccountSeparateEaResponseBody) SetMessage(v string) *EnterpriseAccountSeparateEaResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseAccountSeparateEaResponseBody) SetRequestId(v string) *EnterpriseAccountSeparateEaResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountSeparateEaResponseBody) SetSuccess(v bool) *EnterpriseAccountSeparateEaResponseBody {
	s.Success = &v
	return s
}

type EnterpriseAccountSeparateEaResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseAccountSeparateEaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountSeparateEaResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountSeparateEaResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountSeparateEaResponse) SetHeaders(v map[string]*string) *EnterpriseAccountSeparateEaResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseAccountSeparateEaResponse) SetStatusCode(v int32) *EnterpriseAccountSeparateEaResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseAccountSeparateEaResponse) SetBody(v *EnterpriseAccountSeparateEaResponseBody) *EnterpriseAccountSeparateEaResponse {
	s.Body = v
	return s
}

type EnterpriseAccountUpdateAccountAliasRequest struct {
	Alias         *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
	OrientedEcId  *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId  *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId  *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	Pk            *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnterpriseAccountUpdateAccountAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateAccountAliasRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) SetAlias(v string) *EnterpriseAccountUpdateAccountAliasRequest {
	s.Alias = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) SetAppName(v string) *EnterpriseAccountUpdateAccountAliasRequest {
	s.AppName = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) SetEncryptTicket(v string) *EnterpriseAccountUpdateAccountAliasRequest {
	s.EncryptTicket = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) SetOrientedEcId(v string) *EnterpriseAccountUpdateAccountAliasRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) SetOrientedLeId(v string) *EnterpriseAccountUpdateAccountAliasRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) SetOrientedNbId(v string) *EnterpriseAccountUpdateAccountAliasRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) SetPk(v string) *EnterpriseAccountUpdateAccountAliasRequest {
	s.Pk = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountAliasRequest) SetRequestId(v string) *EnterpriseAccountUpdateAccountAliasRequest {
	s.RequestId = &v
	return s
}

type EnterpriseAccountUpdateAccountAliasResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountUpdateAccountAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateAccountAliasResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateAccountAliasResponseBody) SetCode(v string) *EnterpriseAccountUpdateAccountAliasResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountAliasResponseBody) SetData(v bool) *EnterpriseAccountUpdateAccountAliasResponseBody {
	s.Data = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountAliasResponseBody) SetMessage(v string) *EnterpriseAccountUpdateAccountAliasResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountAliasResponseBody) SetRequestId(v string) *EnterpriseAccountUpdateAccountAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountAliasResponseBody) SetSuccess(v bool) *EnterpriseAccountUpdateAccountAliasResponseBody {
	s.Success = &v
	return s
}

type EnterpriseAccountUpdateAccountAliasResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseAccountUpdateAccountAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountUpdateAccountAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateAccountAliasResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateAccountAliasResponse) SetHeaders(v map[string]*string) *EnterpriseAccountUpdateAccountAliasResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseAccountUpdateAccountAliasResponse) SetStatusCode(v int32) *EnterpriseAccountUpdateAccountAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountAliasResponse) SetBody(v *EnterpriseAccountUpdateAccountAliasResponseBody) *EnterpriseAccountUpdateAccountAliasResponse {
	s.Body = v
	return s
}

type EnterpriseAccountUpdateAccountBizRoleGrantRequest struct {
	BizRoleCodeListJson *string `json:"BizRoleCodeListJson,omitempty" xml:"BizRoleCodeListJson,omitempty"`
	EncryptTicket       *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
	OrientedEcId        *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId        *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId        *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	Pk                  *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s EnterpriseAccountUpdateAccountBizRoleGrantRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateAccountBizRoleGrantRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) SetBizRoleCodeListJson(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest {
	s.BizRoleCodeListJson = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) SetEncryptTicket(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest {
	s.EncryptTicket = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) SetOrientedEcId(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) SetOrientedLeId(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) SetOrientedNbId(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantRequest) SetPk(v string) *EnterpriseAccountUpdateAccountBizRoleGrantRequest {
	s.Pk = &v
	return s
}

type EnterpriseAccountUpdateAccountBizRoleGrantResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) SetCode(v string) *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) SetMessage(v string) *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) SetRequestId(v string) *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) SetSuccess(v bool) *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody {
	s.Success = &v
	return s
}

type EnterpriseAccountUpdateAccountBizRoleGrantResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountUpdateAccountBizRoleGrantResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateAccountBizRoleGrantResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponse) SetHeaders(v map[string]*string) *EnterpriseAccountUpdateAccountBizRoleGrantResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponse) SetStatusCode(v int32) *EnterpriseAccountUpdateAccountBizRoleGrantResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponse) SetBody(v *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) *EnterpriseAccountUpdateAccountBizRoleGrantResponse {
	s.Body = v
	return s
}

type EnterpriseAccountUpdateIpMaskRequest struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	IpMasksJson  *string `json:"IpMasksJson,omitempty" xml:"IpMasksJson,omitempty"`
	OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// This parameter is required.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnterpriseAccountUpdateIpMaskRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateIpMaskRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateIpMaskRequest) SetAppName(v string) *EnterpriseAccountUpdateIpMaskRequest {
	s.AppName = &v
	return s
}

func (s *EnterpriseAccountUpdateIpMaskRequest) SetIpMasksJson(v string) *EnterpriseAccountUpdateIpMaskRequest {
	s.IpMasksJson = &v
	return s
}

func (s *EnterpriseAccountUpdateIpMaskRequest) SetOrientedEcId(v string) *EnterpriseAccountUpdateIpMaskRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseAccountUpdateIpMaskRequest) SetOrientedLeId(v string) *EnterpriseAccountUpdateIpMaskRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseAccountUpdateIpMaskRequest) SetOrientedNbId(v string) *EnterpriseAccountUpdateIpMaskRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseAccountUpdateIpMaskRequest) SetPk(v string) *EnterpriseAccountUpdateIpMaskRequest {
	s.Pk = &v
	return s
}

func (s *EnterpriseAccountUpdateIpMaskRequest) SetRequestId(v string) *EnterpriseAccountUpdateIpMaskRequest {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountUpdateIpMaskRequest) SetStatus(v string) *EnterpriseAccountUpdateIpMaskRequest {
	s.Status = &v
	return s
}

type EnterpriseAccountUpdateIpMaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountUpdateIpMaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateIpMaskResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateIpMaskResponseBody) SetCode(v string) *EnterpriseAccountUpdateIpMaskResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseAccountUpdateIpMaskResponseBody) SetData(v bool) *EnterpriseAccountUpdateIpMaskResponseBody {
	s.Data = &v
	return s
}

func (s *EnterpriseAccountUpdateIpMaskResponseBody) SetMessage(v string) *EnterpriseAccountUpdateIpMaskResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseAccountUpdateIpMaskResponseBody) SetRequestId(v string) *EnterpriseAccountUpdateIpMaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountUpdateIpMaskResponseBody) SetSuccess(v bool) *EnterpriseAccountUpdateIpMaskResponseBody {
	s.Success = &v
	return s
}

type EnterpriseAccountUpdateIpMaskResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseAccountUpdateIpMaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountUpdateIpMaskResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateIpMaskResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateIpMaskResponse) SetHeaders(v map[string]*string) *EnterpriseAccountUpdateIpMaskResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseAccountUpdateIpMaskResponse) SetStatusCode(v int32) *EnterpriseAccountUpdateIpMaskResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseAccountUpdateIpMaskResponse) SetBody(v *EnterpriseAccountUpdateIpMaskResponseBody) *EnterpriseAccountUpdateIpMaskResponse {
	s.Body = v
	return s
}

type EnterpriseAccountUpdateOperateRiskControlRequest struct {
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	ProductLevel *string `json:"ProductLevel,omitempty" xml:"ProductLevel,omitempty"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// This parameter is required.
	ValidateType *string `json:"ValidateType,omitempty" xml:"ValidateType,omitempty"`
}

func (s EnterpriseAccountUpdateOperateRiskControlRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateOperateRiskControlRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) SetAppName(v string) *EnterpriseAccountUpdateOperateRiskControlRequest {
	s.AppName = &v
	return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) SetOrientedEcId(v string) *EnterpriseAccountUpdateOperateRiskControlRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) SetOrientedLeId(v string) *EnterpriseAccountUpdateOperateRiskControlRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) SetOrientedNbId(v string) *EnterpriseAccountUpdateOperateRiskControlRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) SetPk(v string) *EnterpriseAccountUpdateOperateRiskControlRequest {
	s.Pk = &v
	return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) SetProductLevel(v string) *EnterpriseAccountUpdateOperateRiskControlRequest {
	s.ProductLevel = &v
	return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) SetRequestId(v string) *EnterpriseAccountUpdateOperateRiskControlRequest {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlRequest) SetValidateType(v string) *EnterpriseAccountUpdateOperateRiskControlRequest {
	s.ValidateType = &v
	return s
}

type EnterpriseAccountUpdateOperateRiskControlResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountUpdateOperateRiskControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateOperateRiskControlResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponseBody) SetCode(v string) *EnterpriseAccountUpdateOperateRiskControlResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponseBody) SetData(v bool) *EnterpriseAccountUpdateOperateRiskControlResponseBody {
	s.Data = &v
	return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponseBody) SetMessage(v string) *EnterpriseAccountUpdateOperateRiskControlResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponseBody) SetRequestId(v string) *EnterpriseAccountUpdateOperateRiskControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponseBody) SetSuccess(v bool) *EnterpriseAccountUpdateOperateRiskControlResponseBody {
	s.Success = &v
	return s
}

type EnterpriseAccountUpdateOperateRiskControlResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseAccountUpdateOperateRiskControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountUpdateOperateRiskControlResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateOperateRiskControlResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponse) SetHeaders(v map[string]*string) *EnterpriseAccountUpdateOperateRiskControlResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponse) SetStatusCode(v int32) *EnterpriseAccountUpdateOperateRiskControlResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponse) SetBody(v *EnterpriseAccountUpdateOperateRiskControlResponseBody) *EnterpriseAccountUpdateOperateRiskControlResponse {
	s.Body = v
	return s
}

type EnterpriseAccountUpdateSecurityMobileLoginStatusRequest struct {
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// This parameter is required.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) SetAppName(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest {
	s.AppName = &v
	return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) SetOrientedEcId(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) SetOrientedLeId(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) SetOrientedNbId(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) SetPk(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest {
	s.Pk = &v
	return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) SetRequestId(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) SetStatus(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest {
	s.Status = &v
	return s
}

type EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Pass      *bool   `json:"Pass,omitempty" xml:"Pass,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) SetCode(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) SetMessage(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) SetPass(v bool) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody {
	s.Pass = &v
	return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) SetRequestId(v string) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) SetSuccess(v bool) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody {
	s.Success = &v
	return s
}

type EnterpriseAccountUpdateSecurityMobileLoginStatusResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountUpdateSecurityMobileLoginStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateSecurityMobileLoginStatusResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse) SetHeaders(v map[string]*string) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse) SetStatusCode(v int32) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse) SetBody(v *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse {
	s.Body = v
	return s
}

type EnterpriseAccountUpdateSessionExpireTimeRequest struct {
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// This parameter is required.
	SessionExpireTime *int32 `json:"SessionExpireTime,omitempty" xml:"SessionExpireTime,omitempty"`
}

func (s EnterpriseAccountUpdateSessionExpireTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateSessionExpireTimeRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) SetAppName(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest {
	s.AppName = &v
	return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) SetOrientedEcId(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) SetOrientedLeId(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) SetOrientedNbId(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) SetPk(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest {
	s.Pk = &v
	return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) SetRequestId(v string) *EnterpriseAccountUpdateSessionExpireTimeRequest {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeRequest) SetSessionExpireTime(v int32) *EnterpriseAccountUpdateSessionExpireTimeRequest {
	s.SessionExpireTime = &v
	return s
}

type EnterpriseAccountUpdateSessionExpireTimeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountUpdateSessionExpireTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateSessionExpireTimeResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponseBody) SetCode(v string) *EnterpriseAccountUpdateSessionExpireTimeResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponseBody) SetData(v bool) *EnterpriseAccountUpdateSessionExpireTimeResponseBody {
	s.Data = &v
	return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponseBody) SetMessage(v string) *EnterpriseAccountUpdateSessionExpireTimeResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponseBody) SetRequestId(v string) *EnterpriseAccountUpdateSessionExpireTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponseBody) SetSuccess(v bool) *EnterpriseAccountUpdateSessionExpireTimeResponseBody {
	s.Success = &v
	return s
}

type EnterpriseAccountUpdateSessionExpireTimeResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseAccountUpdateSessionExpireTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountUpdateSessionExpireTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseAccountUpdateSessionExpireTimeResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponse) SetHeaders(v map[string]*string) *EnterpriseAccountUpdateSessionExpireTimeResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponse) SetStatusCode(v int32) *EnterpriseAccountUpdateSessionExpireTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponse) SetBody(v *EnterpriseAccountUpdateSessionExpireTimeResponseBody) *EnterpriseAccountUpdateSessionExpireTimeResponse {
	s.Body = v
	return s
}

type EnterpriseOrgQueryLoadTreeRequest struct {
	EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
	// example:
	//
	// true
	LoadOrgOnly  *bool   `json:"LoadOrgOnly,omitempty" xml:"LoadOrgOnly,omitempty"`
	OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CF20ED94-D406-512F-9798-4E1F65720BF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnterpriseOrgQueryLoadTreeRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseOrgQueryLoadTreeRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetEncryptTicket(v string) *EnterpriseOrgQueryLoadTreeRequest {
	s.EncryptTicket = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetLoadOrgOnly(v bool) *EnterpriseOrgQueryLoadTreeRequest {
	s.LoadOrgOnly = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetOrientedEcId(v string) *EnterpriseOrgQueryLoadTreeRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetOrientedLeId(v string) *EnterpriseOrgQueryLoadTreeRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetOrientedNbId(v string) *EnterpriseOrgQueryLoadTreeRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeRequest) SetRequestId(v string) *EnterpriseOrgQueryLoadTreeRequest {
	s.RequestId = &v
	return s
}

type EnterpriseOrgQueryLoadTreeResponseBody struct {
	// example:
	//
	// successful
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A93073FC-1E86-58BA-AB83-54DA6BDB4F03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TreeDto *string `json:"TreeDto,omitempty" xml:"TreeDto,omitempty"`
}

func (s EnterpriseOrgQueryLoadTreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseOrgQueryLoadTreeResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) SetCode(v string) *EnterpriseOrgQueryLoadTreeResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) SetMessage(v string) *EnterpriseOrgQueryLoadTreeResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) SetRequestId(v string) *EnterpriseOrgQueryLoadTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) SetSuccess(v bool) *EnterpriseOrgQueryLoadTreeResponseBody {
	s.Success = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeResponseBody) SetTreeDto(v string) *EnterpriseOrgQueryLoadTreeResponseBody {
	s.TreeDto = &v
	return s
}

type EnterpriseOrgQueryLoadTreeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseOrgQueryLoadTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseOrgQueryLoadTreeResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseOrgQueryLoadTreeResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseOrgQueryLoadTreeResponse) SetHeaders(v map[string]*string) *EnterpriseOrgQueryLoadTreeResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeResponse) SetStatusCode(v int32) *EnterpriseOrgQueryLoadTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseOrgQueryLoadTreeResponse) SetBody(v *EnterpriseOrgQueryLoadTreeResponseBody) *EnterpriseOrgQueryLoadTreeResponse {
	s.Body = v
	return s
}

type EnterpriseRegisterAccountRequest struct {
	// example:
	//
	// 资方支付平台
	Alias           *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	EncryptPassword *string `json:"EncryptPassword,omitempty" xml:"EncryptPassword,omitempty"`
	EncryptTicket   *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
	LoginEmail      *string `json:"LoginEmail,omitempty" xml:"LoginEmail,omitempty"`
	// example:
	//
	// 668514d8083f058f78f7199a
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	OrientedEcId   *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId   *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId   *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// A93073FC-1E86-58BA-AB83-54DA6BDB4F03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	ShowCompleteInfo *bool   `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
	SiteNick         *string `json:"SiteNick,omitempty" xml:"SiteNick,omitempty"`
}

func (s EnterpriseRegisterAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRegisterAccountRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseRegisterAccountRequest) SetAlias(v string) *EnterpriseRegisterAccountRequest {
	s.Alias = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetEncryptPassword(v string) *EnterpriseRegisterAccountRequest {
	s.EncryptPassword = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetEncryptTicket(v string) *EnterpriseRegisterAccountRequest {
	s.EncryptTicket = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetLoginEmail(v string) *EnterpriseRegisterAccountRequest {
	s.LoginEmail = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetOrganizationId(v string) *EnterpriseRegisterAccountRequest {
	s.OrganizationId = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetOrientedEcId(v string) *EnterpriseRegisterAccountRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetOrientedLeId(v string) *EnterpriseRegisterAccountRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetOrientedNbId(v string) *EnterpriseRegisterAccountRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetRequestId(v string) *EnterpriseRegisterAccountRequest {
	s.RequestId = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetShowCompleteInfo(v bool) *EnterpriseRegisterAccountRequest {
	s.ShowCompleteInfo = &v
	return s
}

func (s *EnterpriseRegisterAccountRequest) SetSiteNick(v string) *EnterpriseRegisterAccountRequest {
	s.SiteNick = &v
	return s
}

type EnterpriseRegisterAccountResponseBody struct {
	AccountInfo *EnterpriseRegisterAccountResponseBodyAccountInfo `json:"AccountInfo,omitempty" xml:"AccountInfo,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// The operation is not allowed. Channel state (RELEASED) does not meet expectations (ANSWERED).
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BDFCF081-7BCD-52D5-9D82-0F58D96EFD92
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseRegisterAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRegisterAccountResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseRegisterAccountResponseBody) SetAccountInfo(v *EnterpriseRegisterAccountResponseBodyAccountInfo) *EnterpriseRegisterAccountResponseBody {
	s.AccountInfo = v
	return s
}

func (s *EnterpriseRegisterAccountResponseBody) SetCode(v string) *EnterpriseRegisterAccountResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseRegisterAccountResponseBody) SetMessage(v string) *EnterpriseRegisterAccountResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseRegisterAccountResponseBody) SetRequestId(v string) *EnterpriseRegisterAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseRegisterAccountResponseBody) SetSuccess(v bool) *EnterpriseRegisterAccountResponseBody {
	s.Success = &v
	return s
}

type EnterpriseRegisterAccountResponseBodyAccountInfo struct {
	EnterpriseLicenseNo *string `json:"EnterpriseLicenseNo,omitempty" xml:"EnterpriseLicenseNo,omitempty"`
	// example:
	//
	// 海南屿可网络科技有限公司
	EnterpriseName *string `json:"EnterpriseName,omitempty" xml:"EnterpriseName,omitempty"`
	// example:
	//
	// 195529
	LoginId *string `json:"LoginId,omitempty" xml:"LoginId,omitempty"`
	// example:
	//
	// 5190216604405754
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s EnterpriseRegisterAccountResponseBodyAccountInfo) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRegisterAccountResponseBodyAccountInfo) GoString() string {
	return s.String()
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) SetEnterpriseLicenseNo(v string) *EnterpriseRegisterAccountResponseBodyAccountInfo {
	s.EnterpriseLicenseNo = &v
	return s
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) SetEnterpriseName(v string) *EnterpriseRegisterAccountResponseBodyAccountInfo {
	s.EnterpriseName = &v
	return s
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) SetLoginId(v string) *EnterpriseRegisterAccountResponseBodyAccountInfo {
	s.LoginId = &v
	return s
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) SetPk(v string) *EnterpriseRegisterAccountResponseBodyAccountInfo {
	s.Pk = &v
	return s
}

type EnterpriseRegisterAccountResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseRegisterAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseRegisterAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRegisterAccountResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseRegisterAccountResponse) SetHeaders(v map[string]*string) *EnterpriseRegisterAccountResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseRegisterAccountResponse) SetStatusCode(v int32) *EnterpriseRegisterAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseRegisterAccountResponse) SetBody(v *EnterpriseRegisterAccountResponseBody) *EnterpriseRegisterAccountResponse {
	s.Body = v
	return s
}

type EnterpriseRoleCreateBizRoleRequest struct {
	BizPermissionCodeListJson *string `json:"BizPermissionCodeListJson,omitempty" xml:"BizPermissionCodeListJson,omitempty"`
	BizRoleDesc               *string `json:"BizRoleDesc,omitempty" xml:"BizRoleDesc,omitempty"`
	BizRoleName               *string `json:"BizRoleName,omitempty" xml:"BizRoleName,omitempty"`
	EncryptTicket             *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
	OrientedEcId              *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId              *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId              *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	ResourceType              *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s EnterpriseRoleCreateBizRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleCreateBizRoleRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleCreateBizRoleRequest) SetBizPermissionCodeListJson(v string) *EnterpriseRoleCreateBizRoleRequest {
	s.BizPermissionCodeListJson = &v
	return s
}

func (s *EnterpriseRoleCreateBizRoleRequest) SetBizRoleDesc(v string) *EnterpriseRoleCreateBizRoleRequest {
	s.BizRoleDesc = &v
	return s
}

func (s *EnterpriseRoleCreateBizRoleRequest) SetBizRoleName(v string) *EnterpriseRoleCreateBizRoleRequest {
	s.BizRoleName = &v
	return s
}

func (s *EnterpriseRoleCreateBizRoleRequest) SetEncryptTicket(v string) *EnterpriseRoleCreateBizRoleRequest {
	s.EncryptTicket = &v
	return s
}

func (s *EnterpriseRoleCreateBizRoleRequest) SetOrientedEcId(v string) *EnterpriseRoleCreateBizRoleRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseRoleCreateBizRoleRequest) SetOrientedLeId(v string) *EnterpriseRoleCreateBizRoleRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseRoleCreateBizRoleRequest) SetOrientedNbId(v string) *EnterpriseRoleCreateBizRoleRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseRoleCreateBizRoleRequest) SetResourceType(v string) *EnterpriseRoleCreateBizRoleRequest {
	s.ResourceType = &v
	return s
}

type EnterpriseRoleCreateBizRoleResponseBody struct {
	BizRoleCode *string `json:"BizRoleCode,omitempty" xml:"BizRoleCode,omitempty"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseRoleCreateBizRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleCreateBizRoleResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleCreateBizRoleResponseBody) SetBizRoleCode(v string) *EnterpriseRoleCreateBizRoleResponseBody {
	s.BizRoleCode = &v
	return s
}

func (s *EnterpriseRoleCreateBizRoleResponseBody) SetCode(v string) *EnterpriseRoleCreateBizRoleResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseRoleCreateBizRoleResponseBody) SetMessage(v string) *EnterpriseRoleCreateBizRoleResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseRoleCreateBizRoleResponseBody) SetRequestId(v string) *EnterpriseRoleCreateBizRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseRoleCreateBizRoleResponseBody) SetSuccess(v bool) *EnterpriseRoleCreateBizRoleResponseBody {
	s.Success = &v
	return s
}

type EnterpriseRoleCreateBizRoleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseRoleCreateBizRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseRoleCreateBizRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleCreateBizRoleResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleCreateBizRoleResponse) SetHeaders(v map[string]*string) *EnterpriseRoleCreateBizRoleResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseRoleCreateBizRoleResponse) SetStatusCode(v int32) *EnterpriseRoleCreateBizRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseRoleCreateBizRoleResponse) SetBody(v *EnterpriseRoleCreateBizRoleResponseBody) *EnterpriseRoleCreateBizRoleResponse {
	s.Body = v
	return s
}

type EnterpriseRoleDeleteBizRoleRequest struct {
	BizRoleCode   *string `json:"BizRoleCode,omitempty" xml:"BizRoleCode,omitempty"`
	EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
	OrientedEcId  *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId  *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId  *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
}

func (s EnterpriseRoleDeleteBizRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleDeleteBizRoleRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleDeleteBizRoleRequest) SetBizRoleCode(v string) *EnterpriseRoleDeleteBizRoleRequest {
	s.BizRoleCode = &v
	return s
}

func (s *EnterpriseRoleDeleteBizRoleRequest) SetEncryptTicket(v string) *EnterpriseRoleDeleteBizRoleRequest {
	s.EncryptTicket = &v
	return s
}

func (s *EnterpriseRoleDeleteBizRoleRequest) SetOrientedEcId(v string) *EnterpriseRoleDeleteBizRoleRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseRoleDeleteBizRoleRequest) SetOrientedLeId(v string) *EnterpriseRoleDeleteBizRoleRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseRoleDeleteBizRoleRequest) SetOrientedNbId(v string) *EnterpriseRoleDeleteBizRoleRequest {
	s.OrientedNbId = &v
	return s
}

type EnterpriseRoleDeleteBizRoleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseRoleDeleteBizRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleDeleteBizRoleResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleDeleteBizRoleResponseBody) SetCode(v string) *EnterpriseRoleDeleteBizRoleResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseRoleDeleteBizRoleResponseBody) SetMessage(v string) *EnterpriseRoleDeleteBizRoleResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseRoleDeleteBizRoleResponseBody) SetRequestId(v string) *EnterpriseRoleDeleteBizRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseRoleDeleteBizRoleResponseBody) SetSuccess(v bool) *EnterpriseRoleDeleteBizRoleResponseBody {
	s.Success = &v
	return s
}

type EnterpriseRoleDeleteBizRoleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseRoleDeleteBizRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseRoleDeleteBizRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleDeleteBizRoleResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleDeleteBizRoleResponse) SetHeaders(v map[string]*string) *EnterpriseRoleDeleteBizRoleResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseRoleDeleteBizRoleResponse) SetStatusCode(v int32) *EnterpriseRoleDeleteBizRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseRoleDeleteBizRoleResponse) SetBody(v *EnterpriseRoleDeleteBizRoleResponseBody) *EnterpriseRoleDeleteBizRoleResponse {
	s.Body = v
	return s
}

type EnterpriseRoleQueryAccountForRoleGrantByPageRequest struct {
	BizRoleCode      *string `json:"BizRoleCode,omitempty" xml:"BizRoleCode,omitempty"`
	EncryptTicket    *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
	MaxResults       *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OrgId            *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrientedEcId     *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId     *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId     *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	PageNo           *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query            *string `json:"Query,omitempty" xml:"Query,omitempty"`
	ShowCompleteInfo *bool   `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
}

func (s EnterpriseRoleQueryAccountForRoleGrantByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleQueryAccountForRoleGrantByPageRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetBizRoleCode(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
	s.BizRoleCode = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetEncryptTicket(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
	s.EncryptTicket = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetMaxResults(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
	s.MaxResults = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetNextToken(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
	s.NextToken = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetOrgId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
	s.OrgId = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetOrientedEcId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetOrientedLeId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetOrientedNbId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetPageNo(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
	s.PageNo = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetPageSize(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
	s.PageSize = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetQuery(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
	s.Query = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetShowCompleteInfo(v bool) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
	s.ShowCompleteInfo = &v
	return s
}

type EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody struct {
	Accounts   []*EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	Code       *string                                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	MaxResults *int32                                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message    *string                                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken  *string                                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageNo     *int32                                                              `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32                                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                                              `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetAccounts(v []*EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
	s.Accounts = v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetCode(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetMaxResults(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
	s.MaxResults = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetMessage(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetNextToken(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
	s.NextToken = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetPageNo(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
	s.PageNo = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetPageSize(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetRequestId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetSuccess(v bool) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
	s.Success = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetTotalCount(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetTotalPage(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
	s.TotalPage = &v
	return s
}

type EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts struct {
	Alias    *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	AliyunId *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	Granted  *bool   `json:"Granted,omitempty" xml:"Granted,omitempty"`
	Pk       *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) SetAlias(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts {
	s.Alias = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) SetAliyunId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts {
	s.AliyunId = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) SetGranted(v bool) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts {
	s.Granted = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) SetPk(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts {
	s.Pk = &v
	return s
}

type EnterpriseRoleQueryAccountForRoleGrantByPageResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseRoleQueryAccountForRoleGrantByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleQueryAccountForRoleGrantByPageResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponse) SetHeaders(v map[string]*string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponse) SetStatusCode(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponse) SetBody(v *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) *EnterpriseRoleQueryAccountForRoleGrantByPageResponse {
	s.Body = v
	return s
}

type EnterpriseRoleQueryBizRoleByPageRequest struct {
	EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OrientedEcId  *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId  *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId  *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	PageNo        *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query         *string `json:"Query,omitempty" xml:"Query,omitempty"`
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SrcType       *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleByPageRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetEncryptTicket(v string) *EnterpriseRoleQueryBizRoleByPageRequest {
	s.EncryptTicket = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetMaxResults(v int32) *EnterpriseRoleQueryBizRoleByPageRequest {
	s.MaxResults = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetNextToken(v string) *EnterpriseRoleQueryBizRoleByPageRequest {
	s.NextToken = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetOrientedEcId(v string) *EnterpriseRoleQueryBizRoleByPageRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetOrientedLeId(v string) *EnterpriseRoleQueryBizRoleByPageRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetOrientedNbId(v string) *EnterpriseRoleQueryBizRoleByPageRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetPageNo(v int32) *EnterpriseRoleQueryBizRoleByPageRequest {
	s.PageNo = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetPageSize(v int32) *EnterpriseRoleQueryBizRoleByPageRequest {
	s.PageSize = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetQuery(v string) *EnterpriseRoleQueryBizRoleByPageRequest {
	s.Query = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetResourceType(v string) *EnterpriseRoleQueryBizRoleByPageRequest {
	s.ResourceType = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetSrcType(v string) *EnterpriseRoleQueryBizRoleByPageRequest {
	s.SrcType = &v
	return s
}

type EnterpriseRoleQueryBizRoleByPageResponseBody struct {
	BizRoles   []*EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles `json:"BizRoles,omitempty" xml:"BizRoles,omitempty" type:"Repeated"`
	Code       *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	MaxResults *int32                                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message    *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken  *string                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageNo     *int32                                                  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                                  `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleByPageResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetBizRoles(v []*EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) *EnterpriseRoleQueryBizRoleByPageResponseBody {
	s.BizRoles = v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetCode(v string) *EnterpriseRoleQueryBizRoleByPageResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetMaxResults(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBody {
	s.MaxResults = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetMessage(v string) *EnterpriseRoleQueryBizRoleByPageResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetNextToken(v string) *EnterpriseRoleQueryBizRoleByPageResponseBody {
	s.NextToken = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetPageNo(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBody {
	s.PageNo = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetPageSize(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetRequestId(v string) *EnterpriseRoleQueryBizRoleByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetSuccess(v bool) *EnterpriseRoleQueryBizRoleByPageResponseBody {
	s.Success = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetTotalCount(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetTotalPage(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBody {
	s.TotalPage = &v
	return s
}

type EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles struct {
	BizPermissionCount    *int32    `json:"BizPermissionCount,omitempty" xml:"BizPermissionCount,omitempty"`
	BizPermissionNameList []*string `json:"BizPermissionNameList,omitempty" xml:"BizPermissionNameList,omitempty" type:"Repeated"`
	BizRoleCode           *string   `json:"BizRoleCode,omitempty" xml:"BizRoleCode,omitempty"`
	BizRoleDesc           *string   `json:"BizRoleDesc,omitempty" xml:"BizRoleDesc,omitempty"`
	BizRoleName           *string   `json:"BizRoleName,omitempty" xml:"BizRoleName,omitempty"`
	GrantedPkCount        *int32    `json:"GrantedPkCount,omitempty" xml:"GrantedPkCount,omitempty"`
	ResourceType          *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SrcType               *string   `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) SetBizPermissionCount(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles {
	s.BizPermissionCount = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) SetBizPermissionNameList(v []*string) *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles {
	s.BizPermissionNameList = v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) SetBizRoleCode(v string) *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles {
	s.BizRoleCode = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) SetBizRoleDesc(v string) *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles {
	s.BizRoleDesc = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) SetBizRoleName(v string) *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles {
	s.BizRoleName = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) SetGrantedPkCount(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles {
	s.GrantedPkCount = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) SetResourceType(v string) *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles {
	s.ResourceType = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) SetSrcType(v string) *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles {
	s.SrcType = &v
	return s
}

type EnterpriseRoleQueryBizRoleByPageResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseRoleQueryBizRoleByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleByPageResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleQueryBizRoleByPageResponse) SetHeaders(v map[string]*string) *EnterpriseRoleQueryBizRoleByPageResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponse) SetStatusCode(v int32) *EnterpriseRoleQueryBizRoleByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponse) SetBody(v *EnterpriseRoleQueryBizRoleByPageResponseBody) *EnterpriseRoleQueryBizRoleByPageResponse {
	s.Body = v
	return s
}

type EnterpriseRoleQueryBizRoleDetailRequest struct {
	BizRoleCode   *string `json:"BizRoleCode,omitempty" xml:"BizRoleCode,omitempty"`
	EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
	OrientedEcId  *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId  *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId  *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleDetailRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleQueryBizRoleDetailRequest) SetBizRoleCode(v string) *EnterpriseRoleQueryBizRoleDetailRequest {
	s.BizRoleCode = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailRequest) SetEncryptTicket(v string) *EnterpriseRoleQueryBizRoleDetailRequest {
	s.EncryptTicket = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailRequest) SetOrientedEcId(v string) *EnterpriseRoleQueryBizRoleDetailRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailRequest) SetOrientedLeId(v string) *EnterpriseRoleQueryBizRoleDetailRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailRequest) SetOrientedNbId(v string) *EnterpriseRoleQueryBizRoleDetailRequest {
	s.OrientedNbId = &v
	return s
}

type EnterpriseRoleQueryBizRoleDetailResponseBody struct {
	BizRoleDetail *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail `json:"BizRoleDetail,omitempty" xml:"BizRoleDetail,omitempty" type:"Struct"`
	Code          *string                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleDetailResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBody) SetBizRoleDetail(v *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) *EnterpriseRoleQueryBizRoleDetailResponseBody {
	s.BizRoleDetail = v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBody) SetCode(v string) *EnterpriseRoleQueryBizRoleDetailResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBody) SetMessage(v string) *EnterpriseRoleQueryBizRoleDetailResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBody) SetRequestId(v string) *EnterpriseRoleQueryBizRoleDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBody) SetSuccess(v bool) *EnterpriseRoleQueryBizRoleDetailResponseBody {
	s.Success = &v
	return s
}

type EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail struct {
	BizPermissions []*EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions `json:"BizPermissions,omitempty" xml:"BizPermissions,omitempty" type:"Repeated"`
	BizRoleCode    *string                                                                    `json:"BizRoleCode,omitempty" xml:"BizRoleCode,omitempty"`
	BizRoleDesc    *string                                                                    `json:"BizRoleDesc,omitempty" xml:"BizRoleDesc,omitempty"`
	BizRoleName    *string                                                                    `json:"BizRoleName,omitempty" xml:"BizRoleName,omitempty"`
	ResourceType   *string                                                                    `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SrcType        *string                                                                    `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) SetBizPermissions(v []*EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail {
	s.BizPermissions = v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) SetBizRoleCode(v string) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail {
	s.BizRoleCode = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) SetBizRoleDesc(v string) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail {
	s.BizRoleDesc = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) SetBizRoleName(v string) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail {
	s.BizRoleName = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) SetResourceType(v string) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail {
	s.ResourceType = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) SetSrcType(v string) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail {
	s.SrcType = &v
	return s
}

type EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions struct {
	BizPermissionCode *string `json:"BizPermissionCode,omitempty" xml:"BizPermissionCode,omitempty"`
	BizPermissionDesc *string `json:"BizPermissionDesc,omitempty" xml:"BizPermissionDesc,omitempty"`
	BizPermissionName *string `json:"BizPermissionName,omitempty" xml:"BizPermissionName,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions) SetBizPermissionCode(v string) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions {
	s.BizPermissionCode = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions) SetBizPermissionDesc(v string) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions {
	s.BizPermissionDesc = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions) SetBizPermissionName(v string) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions {
	s.BizPermissionName = &v
	return s
}

type EnterpriseRoleQueryBizRoleDetailResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseRoleQueryBizRoleDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleDetailResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleQueryBizRoleDetailResponse) SetHeaders(v map[string]*string) *EnterpriseRoleQueryBizRoleDetailResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponse) SetStatusCode(v int32) *EnterpriseRoleQueryBizRoleDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponse) SetBody(v *EnterpriseRoleQueryBizRoleDetailResponseBody) *EnterpriseRoleQueryBizRoleDetailResponse {
	s.Body = v
	return s
}

type EnterpriseRoleUpdateBizRoleRequest struct {
	BizPermissionCodeListJson *string `json:"BizPermissionCodeListJson,omitempty" xml:"BizPermissionCodeListJson,omitempty"`
	BizRoleCode               *string `json:"BizRoleCode,omitempty" xml:"BizRoleCode,omitempty"`
	BizRoleDesc               *string `json:"BizRoleDesc,omitempty" xml:"BizRoleDesc,omitempty"`
	BizRoleName               *string `json:"BizRoleName,omitempty" xml:"BizRoleName,omitempty"`
	EncryptTicket             *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
	OrientedEcId              *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId              *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId              *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
}

func (s EnterpriseRoleUpdateBizRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleUpdateBizRoleRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleUpdateBizRoleRequest) SetBizPermissionCodeListJson(v string) *EnterpriseRoleUpdateBizRoleRequest {
	s.BizPermissionCodeListJson = &v
	return s
}

func (s *EnterpriseRoleUpdateBizRoleRequest) SetBizRoleCode(v string) *EnterpriseRoleUpdateBizRoleRequest {
	s.BizRoleCode = &v
	return s
}

func (s *EnterpriseRoleUpdateBizRoleRequest) SetBizRoleDesc(v string) *EnterpriseRoleUpdateBizRoleRequest {
	s.BizRoleDesc = &v
	return s
}

func (s *EnterpriseRoleUpdateBizRoleRequest) SetBizRoleName(v string) *EnterpriseRoleUpdateBizRoleRequest {
	s.BizRoleName = &v
	return s
}

func (s *EnterpriseRoleUpdateBizRoleRequest) SetEncryptTicket(v string) *EnterpriseRoleUpdateBizRoleRequest {
	s.EncryptTicket = &v
	return s
}

func (s *EnterpriseRoleUpdateBizRoleRequest) SetOrientedEcId(v string) *EnterpriseRoleUpdateBizRoleRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseRoleUpdateBizRoleRequest) SetOrientedLeId(v string) *EnterpriseRoleUpdateBizRoleRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseRoleUpdateBizRoleRequest) SetOrientedNbId(v string) *EnterpriseRoleUpdateBizRoleRequest {
	s.OrientedNbId = &v
	return s
}

type EnterpriseRoleUpdateBizRoleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseRoleUpdateBizRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleUpdateBizRoleResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleUpdateBizRoleResponseBody) SetCode(v string) *EnterpriseRoleUpdateBizRoleResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseRoleUpdateBizRoleResponseBody) SetMessage(v string) *EnterpriseRoleUpdateBizRoleResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseRoleUpdateBizRoleResponseBody) SetRequestId(v string) *EnterpriseRoleUpdateBizRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseRoleUpdateBizRoleResponseBody) SetSuccess(v bool) *EnterpriseRoleUpdateBizRoleResponseBody {
	s.Success = &v
	return s
}

type EnterpriseRoleUpdateBizRoleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseRoleUpdateBizRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseRoleUpdateBizRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseRoleUpdateBizRoleResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseRoleUpdateBizRoleResponse) SetHeaders(v map[string]*string) *EnterpriseRoleUpdateBizRoleResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseRoleUpdateBizRoleResponse) SetStatusCode(v int32) *EnterpriseRoleUpdateBizRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseRoleUpdateBizRoleResponse) SetBody(v *EnterpriseRoleUpdateBizRoleResponseBody) *EnterpriseRoleUpdateBizRoleResponse {
	s.Body = v
	return s
}

type EnterpriseTodoDealAccountTodoRequest struct {
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TodoId       *string `json:"TodoId,omitempty" xml:"TodoId,omitempty"`
}

func (s EnterpriseTodoDealAccountTodoRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoDealAccountTodoRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoDealAccountTodoRequest) SetAppName(v string) *EnterpriseTodoDealAccountTodoRequest {
	s.AppName = &v
	return s
}

func (s *EnterpriseTodoDealAccountTodoRequest) SetOrientedEcId(v string) *EnterpriseTodoDealAccountTodoRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseTodoDealAccountTodoRequest) SetOrientedLeId(v string) *EnterpriseTodoDealAccountTodoRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseTodoDealAccountTodoRequest) SetOrientedNbId(v string) *EnterpriseTodoDealAccountTodoRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseTodoDealAccountTodoRequest) SetRemark(v string) *EnterpriseTodoDealAccountTodoRequest {
	s.Remark = &v
	return s
}

func (s *EnterpriseTodoDealAccountTodoRequest) SetStatus(v string) *EnterpriseTodoDealAccountTodoRequest {
	s.Status = &v
	return s
}

func (s *EnterpriseTodoDealAccountTodoRequest) SetTodoId(v string) *EnterpriseTodoDealAccountTodoRequest {
	s.TodoId = &v
	return s
}

type EnterpriseTodoDealAccountTodoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseTodoDealAccountTodoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoDealAccountTodoResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoDealAccountTodoResponseBody) SetCode(v string) *EnterpriseTodoDealAccountTodoResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseTodoDealAccountTodoResponseBody) SetData(v bool) *EnterpriseTodoDealAccountTodoResponseBody {
	s.Data = &v
	return s
}

func (s *EnterpriseTodoDealAccountTodoResponseBody) SetMessage(v string) *EnterpriseTodoDealAccountTodoResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseTodoDealAccountTodoResponseBody) SetRequestId(v string) *EnterpriseTodoDealAccountTodoResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseTodoDealAccountTodoResponseBody) SetSuccess(v bool) *EnterpriseTodoDealAccountTodoResponseBody {
	s.Success = &v
	return s
}

type EnterpriseTodoDealAccountTodoResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseTodoDealAccountTodoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseTodoDealAccountTodoResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoDealAccountTodoResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoDealAccountTodoResponse) SetHeaders(v map[string]*string) *EnterpriseTodoDealAccountTodoResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseTodoDealAccountTodoResponse) SetStatusCode(v int32) *EnterpriseTodoDealAccountTodoResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseTodoDealAccountTodoResponse) SetBody(v *EnterpriseTodoDealAccountTodoResponseBody) *EnterpriseTodoDealAccountTodoResponse {
	s.Body = v
	return s
}

type EnterpriseTodoQueryAccountTodoListRequest struct {
	AppName          *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	MaxResults       *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OperatePk        *string `json:"OperatePk,omitempty" xml:"OperatePk,omitempty"`
	OrientedEcId     *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId     *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId     *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	Page             *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ShowCompleteInfo *bool   `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TodoType         *string `json:"TodoType,omitempty" xml:"TodoType,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetAppName(v string) *EnterpriseTodoQueryAccountTodoListRequest {
	s.AppName = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetMaxResults(v int32) *EnterpriseTodoQueryAccountTodoListRequest {
	s.MaxResults = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetNextToken(v string) *EnterpriseTodoQueryAccountTodoListRequest {
	s.NextToken = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetOperatePk(v string) *EnterpriseTodoQueryAccountTodoListRequest {
	s.OperatePk = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetOrientedEcId(v string) *EnterpriseTodoQueryAccountTodoListRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetOrientedLeId(v string) *EnterpriseTodoQueryAccountTodoListRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetOrientedNbId(v string) *EnterpriseTodoQueryAccountTodoListRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetPage(v int32) *EnterpriseTodoQueryAccountTodoListRequest {
	s.Page = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetPageSize(v int32) *EnterpriseTodoQueryAccountTodoListRequest {
	s.PageSize = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetShowCompleteInfo(v bool) *EnterpriseTodoQueryAccountTodoListRequest {
	s.ShowCompleteInfo = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetStatus(v string) *EnterpriseTodoQueryAccountTodoListRequest {
	s.Status = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetTodoType(v string) *EnterpriseTodoQueryAccountTodoListRequest {
	s.TodoType = &v
	return s
}

type EnterpriseTodoQueryAccountTodoListResponseBody struct {
	Code       *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *EnterpriseTodoQueryAccountTodoListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	MaxResults *int32                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message    *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken  *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) SetCode(v string) *EnterpriseTodoQueryAccountTodoListResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) SetData(v *EnterpriseTodoQueryAccountTodoListResponseBodyData) *EnterpriseTodoQueryAccountTodoListResponseBody {
	s.Data = v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) SetMaxResults(v int32) *EnterpriseTodoQueryAccountTodoListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) SetMessage(v string) *EnterpriseTodoQueryAccountTodoListResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) SetNextToken(v string) *EnterpriseTodoQueryAccountTodoListResponseBody {
	s.NextToken = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) SetRequestId(v string) *EnterpriseTodoQueryAccountTodoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) SetSuccess(v bool) *EnterpriseTodoQueryAccountTodoListResponseBody {
	s.Success = &v
	return s
}

type EnterpriseTodoQueryAccountTodoListResponseBodyData struct {
	Count    *int32                                                        `json:"Count,omitempty" xml:"Count,omitempty"`
	TodoList []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList `json:"TodoList,omitempty" xml:"TodoList,omitempty" type:"Repeated"`
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyData) SetCount(v int32) *EnterpriseTodoQueryAccountTodoListResponseBodyData {
	s.Count = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyData) SetTodoList(v []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) *EnterpriseTodoQueryAccountTodoListResponseBodyData {
	s.TodoList = v
	return s
}

type EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList struct {
	AliyunId          *string                                                                  `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	ApplicantAliyunId *string                                                                  `json:"ApplicantAliyunId,omitempty" xml:"ApplicantAliyunId,omitempty"`
	ApplicantPk       *string                                                                  `json:"ApplicantPk,omitempty" xml:"ApplicantPk,omitempty"`
	ApplyRemark       *string                                                                  `json:"ApplyRemark,omitempty" xml:"ApplyRemark,omitempty"`
	ApplyTime         *int64                                                                   `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	AuditorAliyunId   *string                                                                  `json:"AuditorAliyunId,omitempty" xml:"AuditorAliyunId,omitempty"`
	AuditorPk         *string                                                                  `json:"AuditorPk,omitempty" xml:"AuditorPk,omitempty"`
	AuditorStatus     *string                                                                  `json:"AuditorStatus,omitempty" xml:"AuditorStatus,omitempty"`
	CanceledTime      *int64                                                                   `json:"CanceledTime,omitempty" xml:"CanceledTime,omitempty"`
	Closed            *bool                                                                    `json:"Closed,omitempty" xml:"Closed,omitempty"`
	CurrAuditor       *bool                                                                    `json:"CurrAuditor,omitempty" xml:"CurrAuditor,omitempty"`
	FromLe            *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe        `json:"FromLe,omitempty" xml:"FromLe,omitempty" type:"Struct"`
	Pk                *string                                                                  `json:"Pk,omitempty" xml:"Pk,omitempty"`
	ProcessList       []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList `json:"ProcessList,omitempty" xml:"ProcessList,omitempty" type:"Repeated"`
	Status            *string                                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	TimeoutTime       *int64                                                                   `json:"TimeoutTime,omitempty" xml:"TimeoutTime,omitempty"`
	ToLe              *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe          `json:"ToLe,omitempty" xml:"ToLe,omitempty" type:"Struct"`
	ToLeAudit         *bool                                                                    `json:"ToLeAudit,omitempty" xml:"ToLeAudit,omitempty"`
	TodoId            *string                                                                  `json:"TodoId,omitempty" xml:"TodoId,omitempty"`
	TodoType          *string                                                                  `json:"TodoType,omitempty" xml:"TodoType,omitempty"`
	TodoView          *string                                                                  `json:"TodoView,omitempty" xml:"TodoView,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetAliyunId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.AliyunId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetApplicantAliyunId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.ApplicantAliyunId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetApplicantPk(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.ApplicantPk = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetApplyRemark(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.ApplyRemark = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetApplyTime(v int64) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.ApplyTime = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetAuditorAliyunId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.AuditorAliyunId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetAuditorPk(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.AuditorPk = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetAuditorStatus(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.AuditorStatus = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetCanceledTime(v int64) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.CanceledTime = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetClosed(v bool) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.Closed = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetCurrAuditor(v bool) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.CurrAuditor = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetFromLe(v *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.FromLe = v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetPk(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.Pk = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetProcessList(v []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.ProcessList = v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetStatus(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.Status = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetTimeoutTime(v int64) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.TimeoutTime = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetToLe(v *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.ToLe = v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetToLeAudit(v bool) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.ToLeAudit = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetTodoId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.TodoId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetTodoType(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.TodoType = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetTodoView(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
	s.TodoView = &v
	return s
}

type EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe struct {
	EcId                   *string                                                                        `json:"EcId,omitempty" xml:"EcId,omitempty"`
	LeId                   *string                                                                        `json:"LeId,omitempty" xml:"LeId,omitempty"`
	LicenseNumber          *string                                                                        `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty"`
	ManageableAccountCount *int64                                                                         `json:"ManageableAccountCount,omitempty" xml:"ManageableAccountCount,omitempty"`
	ManagedAccountCount    *int64                                                                         `json:"ManagedAccountCount,omitempty" xml:"ManagedAccountCount,omitempty"`
	ManagerList            []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList `json:"ManagerList,omitempty" xml:"ManagerList,omitempty" type:"Repeated"`
	NbId                   *string                                                                        `json:"NbId,omitempty" xml:"NbId,omitempty"`
	SubjectName            *string                                                                        `json:"SubjectName,omitempty" xml:"SubjectName,omitempty"`
	SubjectType            *string                                                                        `json:"SubjectType,omitempty" xml:"SubjectType,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetEcId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
	s.EcId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetLeId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
	s.LeId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetLicenseNumber(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
	s.LicenseNumber = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetManageableAccountCount(v int64) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
	s.ManageableAccountCount = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetManagedAccountCount(v int64) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
	s.ManagedAccountCount = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetManagerList(v []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
	s.ManagerList = v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetNbId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
	s.NbId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetSubjectName(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
	s.SubjectName = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetSubjectType(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
	s.SubjectType = &v
	return s
}

type EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList struct {
	AliyunId  *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	Pk        *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	PkEncoded *string `json:"PkEncoded,omitempty" xml:"PkEncoded,omitempty"`
	Role      *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) SetAliyunId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList {
	s.AliyunId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) SetPk(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList {
	s.Pk = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) SetPkEncoded(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList {
	s.PkEncoded = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) SetRole(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList {
	s.Role = &v
	return s
}

type EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList struct {
	AliyunId  *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	AuditTime *int64  `json:"AuditTime,omitempty" xml:"AuditTime,omitempty"`
	EcId      *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
	LeId      *string `json:"LeId,omitempty" xml:"LeId,omitempty"`
	NbId      *string `json:"NbId,omitempty" xml:"NbId,omitempty"`
	Pk        *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) SetAliyunId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList {
	s.AliyunId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) SetAuditTime(v int64) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList {
	s.AuditTime = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) SetEcId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList {
	s.EcId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) SetLeId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList {
	s.LeId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) SetNbId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList {
	s.NbId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) SetPk(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList {
	s.Pk = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) SetRemark(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList {
	s.Remark = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) SetStatus(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList {
	s.Status = &v
	return s
}

type EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe struct {
	EcId                   *string                                                                      `json:"EcId,omitempty" xml:"EcId,omitempty"`
	LeId                   *string                                                                      `json:"LeId,omitempty" xml:"LeId,omitempty"`
	LicenseNumber          *string                                                                      `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty"`
	ManageableAccountCount *int64                                                                       `json:"ManageableAccountCount,omitempty" xml:"ManageableAccountCount,omitempty"`
	ManagedAccountCount    *int64                                                                       `json:"ManagedAccountCount,omitempty" xml:"ManagedAccountCount,omitempty"`
	ManagerList            []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList `json:"ManagerList,omitempty" xml:"ManagerList,omitempty" type:"Repeated"`
	NbId                   *string                                                                      `json:"NbId,omitempty" xml:"NbId,omitempty"`
	SubjectName            *string                                                                      `json:"SubjectName,omitempty" xml:"SubjectName,omitempty"`
	SubjectType            *string                                                                      `json:"SubjectType,omitempty" xml:"SubjectType,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetEcId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
	s.EcId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetLeId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
	s.LeId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetLicenseNumber(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
	s.LicenseNumber = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetManageableAccountCount(v int64) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
	s.ManageableAccountCount = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetManagedAccountCount(v int64) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
	s.ManagedAccountCount = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetManagerList(v []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
	s.ManagerList = v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetNbId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
	s.NbId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetSubjectName(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
	s.SubjectName = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetSubjectType(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
	s.SubjectType = &v
	return s
}

type EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList struct {
	AliyunId  *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	Pk        *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	PkEncoded *string `json:"PkEncoded,omitempty" xml:"PkEncoded,omitempty"`
	Role      *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) SetAliyunId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList {
	s.AliyunId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) SetPk(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList {
	s.Pk = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) SetPkEncoded(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList {
	s.PkEncoded = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) SetRole(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList {
	s.Role = &v
	return s
}

type EnterpriseTodoQueryAccountTodoListResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseTodoQueryAccountTodoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponse) SetHeaders(v map[string]*string) *EnterpriseTodoQueryAccountTodoListResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponse) SetStatusCode(v int32) *EnterpriseTodoQueryAccountTodoListResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponse) SetBody(v *EnterpriseTodoQueryAccountTodoListResponseBody) *EnterpriseTodoQueryAccountTodoListResponse {
	s.Body = v
	return s
}

type EnterpriseTodoQueryAccountTodoListByApplicantRequest struct {
	AppName          *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	MaxResults       *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OperatePk        *string `json:"OperatePk,omitempty" xml:"OperatePk,omitempty"`
	OrientedEcId     *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	OrientedLeId     *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	OrientedNbId     *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	Page             *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ShowCompleteInfo *bool   `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TodoType         *string `json:"TodoType,omitempty" xml:"TodoType,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetAppName(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
	s.AppName = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetMaxResults(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
	s.MaxResults = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetNextToken(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
	s.NextToken = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetOperatePk(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
	s.OperatePk = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetOrientedEcId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
	s.OrientedEcId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetOrientedLeId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
	s.OrientedLeId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetOrientedNbId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
	s.OrientedNbId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetPage(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
	s.Page = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetPageSize(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
	s.PageSize = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetShowCompleteInfo(v bool) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
	s.ShowCompleteInfo = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetStatus(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
	s.Status = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetTodoType(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
	s.TodoType = &v
	return s
}

type EnterpriseTodoQueryAccountTodoListByApplicantResponseBody struct {
	Code       *string                                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	MaxResults *int32                                                         `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message    *string                                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken  *string                                                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) SetCode(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) SetData(v *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody {
	s.Data = v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) SetMaxResults(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody {
	s.MaxResults = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) SetMessage(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) SetNextToken(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody {
	s.NextToken = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) SetRequestId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) SetSuccess(v bool) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody {
	s.Success = &v
	return s
}

type EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData struct {
	Count    *int32                                                                   `json:"Count,omitempty" xml:"Count,omitempty"`
	TodoList []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList `json:"TodoList,omitempty" xml:"TodoList,omitempty" type:"Repeated"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData) SetCount(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData {
	s.Count = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData) SetTodoList(v []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData {
	s.TodoList = v
	return s
}

type EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList struct {
	AliyunId          *string                                                                             `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	ApplicantAliyunId *string                                                                             `json:"ApplicantAliyunId,omitempty" xml:"ApplicantAliyunId,omitempty"`
	ApplicantPk       *string                                                                             `json:"ApplicantPk,omitempty" xml:"ApplicantPk,omitempty"`
	ApplyRemark       *string                                                                             `json:"ApplyRemark,omitempty" xml:"ApplyRemark,omitempty"`
	ApplyTime         *int64                                                                              `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	AuditorAliyunId   *string                                                                             `json:"AuditorAliyunId,omitempty" xml:"AuditorAliyunId,omitempty"`
	AuditorPk         *string                                                                             `json:"AuditorPk,omitempty" xml:"AuditorPk,omitempty"`
	AuditorStatus     *string                                                                             `json:"AuditorStatus,omitempty" xml:"AuditorStatus,omitempty"`
	CanceledTime      *int64                                                                              `json:"CanceledTime,omitempty" xml:"CanceledTime,omitempty"`
	Closed            *bool                                                                               `json:"Closed,omitempty" xml:"Closed,omitempty"`
	CurrAuditor       *bool                                                                               `json:"CurrAuditor,omitempty" xml:"CurrAuditor,omitempty"`
	FromLe            *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe        `json:"FromLe,omitempty" xml:"FromLe,omitempty" type:"Struct"`
	Pk                *string                                                                             `json:"Pk,omitempty" xml:"Pk,omitempty"`
	ProcessList       []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList `json:"ProcessList,omitempty" xml:"ProcessList,omitempty" type:"Repeated"`
	Status            *string                                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	TimeoutTime       *int64                                                                              `json:"TimeoutTime,omitempty" xml:"TimeoutTime,omitempty"`
	ToLe              *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe          `json:"ToLe,omitempty" xml:"ToLe,omitempty" type:"Struct"`
	ToLeAudit         *bool                                                                               `json:"ToLeAudit,omitempty" xml:"ToLeAudit,omitempty"`
	TodoId            *string                                                                             `json:"TodoId,omitempty" xml:"TodoId,omitempty"`
	TodoType          *string                                                                             `json:"TodoType,omitempty" xml:"TodoType,omitempty"`
	TodoView          *string                                                                             `json:"TodoView,omitempty" xml:"TodoView,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetAliyunId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.AliyunId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetApplicantAliyunId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.ApplicantAliyunId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetApplicantPk(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.ApplicantPk = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetApplyRemark(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.ApplyRemark = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetApplyTime(v int64) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.ApplyTime = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetAuditorAliyunId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.AuditorAliyunId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetAuditorPk(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.AuditorPk = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetAuditorStatus(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.AuditorStatus = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetCanceledTime(v int64) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.CanceledTime = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetClosed(v bool) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.Closed = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetCurrAuditor(v bool) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.CurrAuditor = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetFromLe(v *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.FromLe = v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetPk(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.Pk = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetProcessList(v []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.ProcessList = v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetStatus(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.Status = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetTimeoutTime(v int64) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.TimeoutTime = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetToLe(v *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.ToLe = v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetToLeAudit(v bool) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.ToLeAudit = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetTodoId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.TodoId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetTodoType(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.TodoType = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetTodoView(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
	s.TodoView = &v
	return s
}

type EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe struct {
	EcId                   *string                                                                                   `json:"EcId,omitempty" xml:"EcId,omitempty"`
	LeId                   *string                                                                                   `json:"LeId,omitempty" xml:"LeId,omitempty"`
	LicenseNumber          *string                                                                                   `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty"`
	ManageableAccountCount *int64                                                                                    `json:"ManageableAccountCount,omitempty" xml:"ManageableAccountCount,omitempty"`
	ManagedAccountCount    *int64                                                                                    `json:"ManagedAccountCount,omitempty" xml:"ManagedAccountCount,omitempty"`
	ManagerList            []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList `json:"ManagerList,omitempty" xml:"ManagerList,omitempty" type:"Repeated"`
	NbId                   *string                                                                                   `json:"NbId,omitempty" xml:"NbId,omitempty"`
	SubjectName            *string                                                                                   `json:"SubjectName,omitempty" xml:"SubjectName,omitempty"`
	SubjectType            *string                                                                                   `json:"SubjectType,omitempty" xml:"SubjectType,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetEcId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
	s.EcId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetLeId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
	s.LeId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetLicenseNumber(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
	s.LicenseNumber = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetManageableAccountCount(v int64) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
	s.ManageableAccountCount = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetManagedAccountCount(v int64) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
	s.ManagedAccountCount = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetManagerList(v []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
	s.ManagerList = v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetNbId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
	s.NbId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetSubjectName(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
	s.SubjectName = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetSubjectType(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
	s.SubjectType = &v
	return s
}

type EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList struct {
	AliyunId  *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	Pk        *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	PkEncoded *string `json:"PkEncoded,omitempty" xml:"PkEncoded,omitempty"`
	Role      *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) SetAliyunId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList {
	s.AliyunId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) SetPk(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList {
	s.Pk = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) SetPkEncoded(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList {
	s.PkEncoded = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) SetRole(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList {
	s.Role = &v
	return s
}

type EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList struct {
	AliyunId  *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	AuditTime *int64  `json:"AuditTime,omitempty" xml:"AuditTime,omitempty"`
	EcId      *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
	LeId      *string `json:"LeId,omitempty" xml:"LeId,omitempty"`
	NbId      *string `json:"NbId,omitempty" xml:"NbId,omitempty"`
	Pk        *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) SetAliyunId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList {
	s.AliyunId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) SetAuditTime(v int64) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList {
	s.AuditTime = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) SetEcId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList {
	s.EcId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) SetLeId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList {
	s.LeId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) SetNbId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList {
	s.NbId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) SetPk(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList {
	s.Pk = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) SetRemark(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList {
	s.Remark = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) SetStatus(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList {
	s.Status = &v
	return s
}

type EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe struct {
	EcId                   *string                                                                                 `json:"EcId,omitempty" xml:"EcId,omitempty"`
	LeId                   *string                                                                                 `json:"LeId,omitempty" xml:"LeId,omitempty"`
	LicenseNumber          *string                                                                                 `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty"`
	ManageableAccountCount *int64                                                                                  `json:"ManageableAccountCount,omitempty" xml:"ManageableAccountCount,omitempty"`
	ManagedAccountCount    *int64                                                                                  `json:"ManagedAccountCount,omitempty" xml:"ManagedAccountCount,omitempty"`
	ManagerList            []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList `json:"ManagerList,omitempty" xml:"ManagerList,omitempty" type:"Repeated"`
	NbId                   *string                                                                                 `json:"NbId,omitempty" xml:"NbId,omitempty"`
	SubjectName            *string                                                                                 `json:"SubjectName,omitempty" xml:"SubjectName,omitempty"`
	SubjectType            *string                                                                                 `json:"SubjectType,omitempty" xml:"SubjectType,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetEcId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
	s.EcId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetLeId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
	s.LeId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetLicenseNumber(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
	s.LicenseNumber = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetManageableAccountCount(v int64) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
	s.ManageableAccountCount = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetManagedAccountCount(v int64) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
	s.ManagedAccountCount = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetManagerList(v []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
	s.ManagerList = v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetNbId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
	s.NbId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetSubjectName(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
	s.SubjectName = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetSubjectType(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
	s.SubjectType = &v
	return s
}

type EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList struct {
	AliyunId  *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	Pk        *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	PkEncoded *string `json:"PkEncoded,omitempty" xml:"PkEncoded,omitempty"`
	Role      *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) SetAliyunId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList {
	s.AliyunId = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) SetPk(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList {
	s.Pk = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) SetPkEncoded(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList {
	s.PkEncoded = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) SetRole(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList {
	s.Role = &v
	return s
}

type EnterpriseTodoQueryAccountTodoListByApplicantResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponse) SetHeaders(v map[string]*string) *EnterpriseTodoQueryAccountTodoListByApplicantResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponse) SetStatusCode(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponse) SetBody(v *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) *EnterpriseTodoQueryAccountTodoListByApplicantResponse {
	s.Body = v
	return s
}

type EnterpriseUninvitedAdminInviteJoinEnterpriseRequest struct {
	EcId          *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
	EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
	InviteePk     *string `json:"InviteePk,omitempty" xml:"InviteePk,omitempty"`
	LeId          *string `json:"LeId,omitempty" xml:"LeId,omitempty"`
	NbId          *string `json:"NbId,omitempty" xml:"NbId,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) GoString() string {
	return s.String()
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) SetEcId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest {
	s.EcId = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) SetEncryptTicket(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest {
	s.EncryptTicket = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) SetInviteePk(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest {
	s.InviteePk = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) SetLeId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest {
	s.LeId = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) SetNbId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest {
	s.NbId = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) SetRemark(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest {
	s.Remark = &v
	return s
}

type EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody struct {
	Code      *string                                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) GoString() string {
	return s.String()
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) SetCode(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody {
	s.Code = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) SetData(v []*EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody {
	s.Data = v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) SetMessage(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody {
	s.Message = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) SetRequestId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) SetSuccess(v bool) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody {
	s.Success = &v
	return s
}

type EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData struct {
	ApplicantAliyunId *string `json:"ApplicantAliyunId,omitempty" xml:"ApplicantAliyunId,omitempty"`
	ApplicantPk       *string `json:"ApplicantPk,omitempty" xml:"ApplicantPk,omitempty"`
	ApplyRemark       *string `json:"ApplyRemark,omitempty" xml:"ApplyRemark,omitempty"`
	ApplyTime         *int64  `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	AuditorAliyunId   *string `json:"AuditorAliyunId,omitempty" xml:"AuditorAliyunId,omitempty"`
	AuditorPk         *string `json:"AuditorPk,omitempty" xml:"AuditorPk,omitempty"`
	EcId              *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
	LeId              *string `json:"LeId,omitempty" xml:"LeId,omitempty"`
	LeLicenseNo       *string `json:"LeLicenseNo,omitempty" xml:"LeLicenseNo,omitempty"`
	LeName            *string `json:"LeName,omitempty" xml:"LeName,omitempty"`
	Message           *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NbId              *string `json:"NbId,omitempty" xml:"NbId,omitempty"`
	OperatedAliyunId  *string `json:"OperatedAliyunId,omitempty" xml:"OperatedAliyunId,omitempty"`
	OperatedPk        *string `json:"OperatedPk,omitempty" xml:"OperatedPk,omitempty"`
	Success           *string `json:"Success,omitempty" xml:"Success,omitempty"`
	TimeoutTime       *int64  `json:"TimeoutTime,omitempty" xml:"TimeoutTime,omitempty"`
	TodoId            *string `json:"TodoId,omitempty" xml:"TodoId,omitempty"`
	TodoType          *string `json:"TodoType,omitempty" xml:"TodoType,omitempty"`
}

func (s EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetApplicantAliyunId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.ApplicantAliyunId = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetApplicantPk(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.ApplicantPk = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetApplyRemark(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.ApplyRemark = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetApplyTime(v int64) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.ApplyTime = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetAuditorAliyunId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.AuditorAliyunId = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetAuditorPk(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.AuditorPk = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetEcId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.EcId = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetLeId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.LeId = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetLeLicenseNo(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.LeLicenseNo = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetLeName(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.LeName = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetMessage(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.Message = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetNbId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.NbId = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetOperatedAliyunId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.OperatedAliyunId = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetOperatedPk(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.OperatedPk = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetSuccess(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.Success = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetTimeoutTime(v int64) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.TimeoutTime = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetTodoId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.TodoId = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetTodoType(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
	s.TodoType = &v
	return s
}

type EnterpriseUninvitedAdminInviteJoinEnterpriseResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseUninvitedAdminInviteJoinEnterpriseResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterpriseUninvitedAdminInviteJoinEnterpriseResponse) GoString() string {
	return s.String()
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse) SetHeaders(v map[string]*string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse {
	s.Headers = v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse) SetStatusCode(v int32) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse) SetBody(v *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("accountcenter"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 修改登录密码
//
// @param request - EnterpriseAccountChangeLoginPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountChangeLoginPasswordResponse
func (client *Client) EnterpriseAccountChangeLoginPasswordWithOptions(request *EnterpriseAccountChangeLoginPasswordRequest, runtime *util.RuntimeOptions) (_result *EnterpriseAccountChangeLoginPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptPassword)) {
		query["EncryptPassword"] = request.EncryptPassword
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseAccountChangeLoginPassword"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseAccountChangeLoginPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改登录密码
//
// @param request - EnterpriseAccountChangeLoginPasswordRequest
//
// @return EnterpriseAccountChangeLoginPasswordResponse
func (client *Client) EnterpriseAccountChangeLoginPassword(request *EnterpriseAccountChangeLoginPasswordRequest) (_result *EnterpriseAccountChangeLoginPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseAccountChangeLoginPasswordResponse{}
	_body, _err := client.EnterpriseAccountChangeLoginPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改安全邮箱
//
// @param request - EnterpriseAccountChangeSecurityEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountChangeSecurityEmailResponse
func (client *Client) EnterpriseAccountChangeSecurityEmailWithOptions(request *EnterpriseAccountChangeSecurityEmailRequest, runtime *util.RuntimeOptions) (_result *EnterpriseAccountChangeSecurityEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityEmail)) {
		query["SecurityEmail"] = request.SecurityEmail
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyCode)) {
		query["VerifyCode"] = request.VerifyCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseAccountChangeSecurityEmail"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseAccountChangeSecurityEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改安全邮箱
//
// @param request - EnterpriseAccountChangeSecurityEmailRequest
//
// @return EnterpriseAccountChangeSecurityEmailResponse
func (client *Client) EnterpriseAccountChangeSecurityEmail(request *EnterpriseAccountChangeSecurityEmailRequest) (_result *EnterpriseAccountChangeSecurityEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseAccountChangeSecurityEmailResponse{}
	_body, _err := client.EnterpriseAccountChangeSecurityEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改成员账号安全手机号
//
// @param request - EnterpriseAccountChangeSecurityMobileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountChangeSecurityMobileResponse
func (client *Client) EnterpriseAccountChangeSecurityMobileWithOptions(request *EnterpriseAccountChangeSecurityMobileRequest, runtime *util.RuntimeOptions) (_result *EnterpriseAccountChangeSecurityMobileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptTicket)) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !tea.BoolValue(util.IsUnset(request.NewMobile)) {
		query["NewMobile"] = request.NewMobile
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.VerificationCode)) {
		query["VerificationCode"] = request.VerificationCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseAccountChangeSecurityMobile"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseAccountChangeSecurityMobileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改成员账号安全手机号
//
// @param request - EnterpriseAccountChangeSecurityMobileRequest
//
// @return EnterpriseAccountChangeSecurityMobileResponse
func (client *Client) EnterpriseAccountChangeSecurityMobile(request *EnterpriseAccountChangeSecurityMobileRequest) (_result *EnterpriseAccountChangeSecurityMobileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseAccountChangeSecurityMobileResponse{}
	_body, _err := client.EnterpriseAccountChangeSecurityMobileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询纳管账号授权角色
//
// @param request - EnterpriseAccountQueryAccountGrantedRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountQueryAccountGrantedRolesResponse
func (client *Client) EnterpriseAccountQueryAccountGrantedRolesWithOptions(request *EnterpriseAccountQueryAccountGrantedRolesRequest, runtime *util.RuntimeOptions) (_result *EnterpriseAccountQueryAccountGrantedRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.IsOpenApi)) {
		body["IsOpenApi"] = request.IsOpenApi
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		body["Pk"] = request.Pk
	}

	if !tea.BoolValue(util.IsUnset(request.ShowCompleteInfo)) {
		body["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseAccountQueryAccountGrantedRoles"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseAccountQueryAccountGrantedRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询纳管账号授权角色
//
// @param request - EnterpriseAccountQueryAccountGrantedRolesRequest
//
// @return EnterpriseAccountQueryAccountGrantedRolesResponse
func (client *Client) EnterpriseAccountQueryAccountGrantedRoles(request *EnterpriseAccountQueryAccountGrantedRolesRequest) (_result *EnterpriseAccountQueryAccountGrantedRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseAccountQueryAccountGrantedRolesResponse{}
	_body, _err := client.EnterpriseAccountQueryAccountGrantedRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询纳管账号信息
//
// @param request - EnterpriseAccountQueryAccountsInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountQueryAccountsInfoResponse
func (client *Client) EnterpriseAccountQueryAccountsInfoWithOptions(request *EnterpriseAccountQueryAccountsInfoRequest, runtime *util.RuntimeOptions) (_result *EnterpriseAccountQueryAccountsInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptTicket)) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PksJson)) {
		query["PksJson"] = request.PksJson
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowCompleteInfo)) {
		body["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseAccountQueryAccountsInfo"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseAccountQueryAccountsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询纳管账号信息
//
// @param request - EnterpriseAccountQueryAccountsInfoRequest
//
// @return EnterpriseAccountQueryAccountsInfoResponse
func (client *Client) EnterpriseAccountQueryAccountsInfo(request *EnterpriseAccountQueryAccountsInfoRequest) (_result *EnterpriseAccountQueryAccountsInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseAccountQueryAccountsInfoResponse{}
	_body, _err := client.EnterpriseAccountQueryAccountsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询纳管账号登录设置
//
// @param request - EnterpriseAccountQueryLoginSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountQueryLoginSettingsResponse
func (client *Client) EnterpriseAccountQueryLoginSettingsWithOptions(request *EnterpriseAccountQueryLoginSettingsRequest, runtime *util.RuntimeOptions) (_result *EnterpriseAccountQueryLoginSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.IsOpenApi)) {
		body["IsOpenApi"] = request.IsOpenApi
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		body["Pk"] = request.Pk
	}

	if !tea.BoolValue(util.IsUnset(request.ShowCompleteInfo)) {
		body["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseAccountQueryLoginSettings"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseAccountQueryLoginSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询纳管账号登录设置
//
// @param request - EnterpriseAccountQueryLoginSettingsRequest
//
// @return EnterpriseAccountQueryLoginSettingsResponse
func (client *Client) EnterpriseAccountQueryLoginSettings(request *EnterpriseAccountQueryLoginSettingsRequest) (_result *EnterpriseAccountQueryLoginSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseAccountQueryLoginSettingsResponse{}
	_body, _err := client.EnterpriseAccountQueryLoginSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移除mfa
//
// @param request - EnterpriseAccountRemoveMfaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountRemoveMfaResponse
func (client *Client) EnterpriseAccountRemoveMfaWithOptions(request *EnterpriseAccountRemoveMfaRequest, runtime *util.RuntimeOptions) (_result *EnterpriseAccountRemoveMfaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseAccountRemoveMfa"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseAccountRemoveMfaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除mfa
//
// @param request - EnterpriseAccountRemoveMfaRequest
//
// @return EnterpriseAccountRemoveMfaResponse
func (client *Client) EnterpriseAccountRemoveMfa(request *EnterpriseAccountRemoveMfaRequest) (_result *EnterpriseAccountRemoveMfaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseAccountRemoveMfaResponse{}
	_body, _err := client.EnterpriseAccountRemoveMfaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 脱离ea
//
// @param request - EnterpriseAccountSeparateEaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountSeparateEaResponse
func (client *Client) EnterpriseAccountSeparateEaWithOptions(request *EnterpriseAccountSeparateEaRequest, runtime *util.RuntimeOptions) (_result *EnterpriseAccountSeparateEaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptTicket)) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseAccountSeparateEa"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseAccountSeparateEaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 脱离ea
//
// @param request - EnterpriseAccountSeparateEaRequest
//
// @return EnterpriseAccountSeparateEaResponse
func (client *Client) EnterpriseAccountSeparateEa(request *EnterpriseAccountSeparateEaRequest) (_result *EnterpriseAccountSeparateEaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseAccountSeparateEaResponse{}
	_body, _err := client.EnterpriseAccountSeparateEaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新账号企业多账号中的别名
//
// @param request - EnterpriseAccountUpdateAccountAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateAccountAliasResponse
func (client *Client) EnterpriseAccountUpdateAccountAliasWithOptions(request *EnterpriseAccountUpdateAccountAliasRequest, runtime *util.RuntimeOptions) (_result *EnterpriseAccountUpdateAccountAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Alias)) {
		query["Alias"] = request.Alias
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptTicket)) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseAccountUpdateAccountAlias"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseAccountUpdateAccountAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新账号企业多账号中的别名
//
// @param request - EnterpriseAccountUpdateAccountAliasRequest
//
// @return EnterpriseAccountUpdateAccountAliasResponse
func (client *Client) EnterpriseAccountUpdateAccountAlias(request *EnterpriseAccountUpdateAccountAliasRequest) (_result *EnterpriseAccountUpdateAccountAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseAccountUpdateAccountAliasResponse{}
	_body, _err := client.EnterpriseAccountUpdateAccountAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新账号授权
//
// @param request - EnterpriseAccountUpdateAccountBizRoleGrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateAccountBizRoleGrantResponse
func (client *Client) EnterpriseAccountUpdateAccountBizRoleGrantWithOptions(request *EnterpriseAccountUpdateAccountBizRoleGrantRequest, runtime *util.RuntimeOptions) (_result *EnterpriseAccountUpdateAccountBizRoleGrantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRoleCodeListJson)) {
		query["BizRoleCodeListJson"] = request.BizRoleCodeListJson
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptTicket)) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseAccountUpdateAccountBizRoleGrant"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseAccountUpdateAccountBizRoleGrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新账号授权
//
// @param request - EnterpriseAccountUpdateAccountBizRoleGrantRequest
//
// @return EnterpriseAccountUpdateAccountBizRoleGrantResponse
func (client *Client) EnterpriseAccountUpdateAccountBizRoleGrant(request *EnterpriseAccountUpdateAccountBizRoleGrantRequest) (_result *EnterpriseAccountUpdateAccountBizRoleGrantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseAccountUpdateAccountBizRoleGrantResponse{}
	_body, _err := client.EnterpriseAccountUpdateAccountBizRoleGrantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置Ip掩码
//
// @param request - EnterpriseAccountUpdateIpMaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateIpMaskResponse
func (client *Client) EnterpriseAccountUpdateIpMaskWithOptions(request *EnterpriseAccountUpdateIpMaskRequest, runtime *util.RuntimeOptions) (_result *EnterpriseAccountUpdateIpMaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpMasksJson)) {
		query["IpMasksJson"] = request.IpMasksJson
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseAccountUpdateIpMask"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseAccountUpdateIpMaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置Ip掩码
//
// @param request - EnterpriseAccountUpdateIpMaskRequest
//
// @return EnterpriseAccountUpdateIpMaskResponse
func (client *Client) EnterpriseAccountUpdateIpMask(request *EnterpriseAccountUpdateIpMaskRequest) (_result *EnterpriseAccountUpdateIpMaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseAccountUpdateIpMaskResponse{}
	_body, _err := client.EnterpriseAccountUpdateIpMaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新操作风控
//
// @param request - EnterpriseAccountUpdateOperateRiskControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateOperateRiskControlResponse
func (client *Client) EnterpriseAccountUpdateOperateRiskControlWithOptions(request *EnterpriseAccountUpdateOperateRiskControlRequest, runtime *util.RuntimeOptions) (_result *EnterpriseAccountUpdateOperateRiskControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	if !tea.BoolValue(util.IsUnset(request.ProductLevel)) {
		query["ProductLevel"] = request.ProductLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ValidateType)) {
		query["ValidateType"] = request.ValidateType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseAccountUpdateOperateRiskControl"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseAccountUpdateOperateRiskControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新操作风控
//
// @param request - EnterpriseAccountUpdateOperateRiskControlRequest
//
// @return EnterpriseAccountUpdateOperateRiskControlResponse
func (client *Client) EnterpriseAccountUpdateOperateRiskControl(request *EnterpriseAccountUpdateOperateRiskControlRequest) (_result *EnterpriseAccountUpdateOperateRiskControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseAccountUpdateOperateRiskControlResponse{}
	_body, _err := client.EnterpriseAccountUpdateOperateRiskControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改安全手机启用状态
//
// @param request - EnterpriseAccountUpdateSecurityMobileLoginStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateSecurityMobileLoginStatusResponse
func (client *Client) EnterpriseAccountUpdateSecurityMobileLoginStatusWithOptions(request *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest, runtime *util.RuntimeOptions) (_result *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseAccountUpdateSecurityMobileLoginStatus"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseAccountUpdateSecurityMobileLoginStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改安全手机启用状态
//
// @param request - EnterpriseAccountUpdateSecurityMobileLoginStatusRequest
//
// @return EnterpriseAccountUpdateSecurityMobileLoginStatusResponse
func (client *Client) EnterpriseAccountUpdateSecurityMobileLoginStatus(request *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) (_result *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseAccountUpdateSecurityMobileLoginStatusResponse{}
	_body, _err := client.EnterpriseAccountUpdateSecurityMobileLoginStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新过期时间
//
// @param request - EnterpriseAccountUpdateSessionExpireTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateSessionExpireTimeResponse
func (client *Client) EnterpriseAccountUpdateSessionExpireTimeWithOptions(request *EnterpriseAccountUpdateSessionExpireTimeRequest, runtime *util.RuntimeOptions) (_result *EnterpriseAccountUpdateSessionExpireTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionExpireTime)) {
		query["SessionExpireTime"] = request.SessionExpireTime
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseAccountUpdateSessionExpireTime"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseAccountUpdateSessionExpireTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新过期时间
//
// @param request - EnterpriseAccountUpdateSessionExpireTimeRequest
//
// @return EnterpriseAccountUpdateSessionExpireTimeResponse
func (client *Client) EnterpriseAccountUpdateSessionExpireTime(request *EnterpriseAccountUpdateSessionExpireTimeRequest) (_result *EnterpriseAccountUpdateSessionExpireTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseAccountUpdateSessionExpireTimeResponse{}
	_body, _err := client.EnterpriseAccountUpdateSessionExpireTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 组织目录树查询
//
// @param request - EnterpriseOrgQueryLoadTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseOrgQueryLoadTreeResponse
func (client *Client) EnterpriseOrgQueryLoadTreeWithOptions(request *EnterpriseOrgQueryLoadTreeRequest, runtime *util.RuntimeOptions) (_result *EnterpriseOrgQueryLoadTreeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptTicket)) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !tea.BoolValue(util.IsUnset(request.LoadOrgOnly)) {
		query["LoadOrgOnly"] = request.LoadOrgOnly
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseOrgQueryLoadTree"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseOrgQueryLoadTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 组织目录树查询
//
// @param request - EnterpriseOrgQueryLoadTreeRequest
//
// @return EnterpriseOrgQueryLoadTreeResponse
func (client *Client) EnterpriseOrgQueryLoadTree(request *EnterpriseOrgQueryLoadTreeRequest) (_result *EnterpriseOrgQueryLoadTreeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseOrgQueryLoadTreeResponse{}
	_body, _err := client.EnterpriseOrgQueryLoadTreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建成员账号
//
// @param request - EnterpriseRegisterAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRegisterAccountResponse
func (client *Client) EnterpriseRegisterAccountWithOptions(request *EnterpriseRegisterAccountRequest, runtime *util.RuntimeOptions) (_result *EnterpriseRegisterAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Alias)) {
		query["Alias"] = request.Alias
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptPassword)) {
		query["EncryptPassword"] = request.EncryptPassword
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptTicket)) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !tea.BoolValue(util.IsUnset(request.LoginEmail)) {
		query["LoginEmail"] = request.LoginEmail
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowCompleteInfo)) {
		query["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SiteNick)) {
		query["SiteNick"] = request.SiteNick
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseRegisterAccount"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseRegisterAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建成员账号
//
// @param request - EnterpriseRegisterAccountRequest
//
// @return EnterpriseRegisterAccountResponse
func (client *Client) EnterpriseRegisterAccount(request *EnterpriseRegisterAccountRequest) (_result *EnterpriseRegisterAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseRegisterAccountResponse{}
	_body, _err := client.EnterpriseRegisterAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建业务角色
//
// @param request - EnterpriseRoleCreateBizRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleCreateBizRoleResponse
func (client *Client) EnterpriseRoleCreateBizRoleWithOptions(request *EnterpriseRoleCreateBizRoleRequest, runtime *util.RuntimeOptions) (_result *EnterpriseRoleCreateBizRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizPermissionCodeListJson)) {
		query["BizPermissionCodeListJson"] = request.BizPermissionCodeListJson
	}

	if !tea.BoolValue(util.IsUnset(request.BizRoleDesc)) {
		query["BizRoleDesc"] = request.BizRoleDesc
	}

	if !tea.BoolValue(util.IsUnset(request.BizRoleName)) {
		query["BizRoleName"] = request.BizRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptTicket)) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseRoleCreateBizRole"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseRoleCreateBizRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建业务角色
//
// @param request - EnterpriseRoleCreateBizRoleRequest
//
// @return EnterpriseRoleCreateBizRoleResponse
func (client *Client) EnterpriseRoleCreateBizRole(request *EnterpriseRoleCreateBizRoleRequest) (_result *EnterpriseRoleCreateBizRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseRoleCreateBizRoleResponse{}
	_body, _err := client.EnterpriseRoleCreateBizRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除业务角色
//
// @param request - EnterpriseRoleDeleteBizRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleDeleteBizRoleResponse
func (client *Client) EnterpriseRoleDeleteBizRoleWithOptions(request *EnterpriseRoleDeleteBizRoleRequest, runtime *util.RuntimeOptions) (_result *EnterpriseRoleDeleteBizRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRoleCode)) {
		query["BizRoleCode"] = request.BizRoleCode
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptTicket)) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseRoleDeleteBizRole"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseRoleDeleteBizRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除业务角色
//
// @param request - EnterpriseRoleDeleteBizRoleRequest
//
// @return EnterpriseRoleDeleteBizRoleResponse
func (client *Client) EnterpriseRoleDeleteBizRole(request *EnterpriseRoleDeleteBizRoleRequest) (_result *EnterpriseRoleDeleteBizRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseRoleDeleteBizRoleResponse{}
	_body, _err := client.EnterpriseRoleDeleteBizRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 角色授权场景下分页查询账号
//
// @param request - EnterpriseRoleQueryAccountForRoleGrantByPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleQueryAccountForRoleGrantByPageResponse
func (client *Client) EnterpriseRoleQueryAccountForRoleGrantByPageWithOptions(request *EnterpriseRoleQueryAccountForRoleGrantByPageRequest, runtime *util.RuntimeOptions) (_result *EnterpriseRoleQueryAccountForRoleGrantByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRoleCode)) {
		query["BizRoleCode"] = request.BizRoleCode
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptTicket)) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.ShowCompleteInfo)) {
		query["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseRoleQueryAccountForRoleGrantByPage"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseRoleQueryAccountForRoleGrantByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 角色授权场景下分页查询账号
//
// @param request - EnterpriseRoleQueryAccountForRoleGrantByPageRequest
//
// @return EnterpriseRoleQueryAccountForRoleGrantByPageResponse
func (client *Client) EnterpriseRoleQueryAccountForRoleGrantByPage(request *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) (_result *EnterpriseRoleQueryAccountForRoleGrantByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseRoleQueryAccountForRoleGrantByPageResponse{}
	_body, _err := client.EnterpriseRoleQueryAccountForRoleGrantByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询业务角色
//
// @param request - EnterpriseRoleQueryBizRoleByPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleQueryBizRoleByPageResponse
func (client *Client) EnterpriseRoleQueryBizRoleByPageWithOptions(request *EnterpriseRoleQueryBizRoleByPageRequest, runtime *util.RuntimeOptions) (_result *EnterpriseRoleQueryBizRoleByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptTicket)) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SrcType)) {
		query["SrcType"] = request.SrcType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseRoleQueryBizRoleByPage"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseRoleQueryBizRoleByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询业务角色
//
// @param request - EnterpriseRoleQueryBizRoleByPageRequest
//
// @return EnterpriseRoleQueryBizRoleByPageResponse
func (client *Client) EnterpriseRoleQueryBizRoleByPage(request *EnterpriseRoleQueryBizRoleByPageRequest) (_result *EnterpriseRoleQueryBizRoleByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseRoleQueryBizRoleByPageResponse{}
	_body, _err := client.EnterpriseRoleQueryBizRoleByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询业务角色详情
//
// @param request - EnterpriseRoleQueryBizRoleDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleQueryBizRoleDetailResponse
func (client *Client) EnterpriseRoleQueryBizRoleDetailWithOptions(request *EnterpriseRoleQueryBizRoleDetailRequest, runtime *util.RuntimeOptions) (_result *EnterpriseRoleQueryBizRoleDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRoleCode)) {
		query["BizRoleCode"] = request.BizRoleCode
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptTicket)) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseRoleQueryBizRoleDetail"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseRoleQueryBizRoleDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询业务角色详情
//
// @param request - EnterpriseRoleQueryBizRoleDetailRequest
//
// @return EnterpriseRoleQueryBizRoleDetailResponse
func (client *Client) EnterpriseRoleQueryBizRoleDetail(request *EnterpriseRoleQueryBizRoleDetailRequest) (_result *EnterpriseRoleQueryBizRoleDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseRoleQueryBizRoleDetailResponse{}
	_body, _err := client.EnterpriseRoleQueryBizRoleDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新业务角色
//
// @param request - EnterpriseRoleUpdateBizRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleUpdateBizRoleResponse
func (client *Client) EnterpriseRoleUpdateBizRoleWithOptions(request *EnterpriseRoleUpdateBizRoleRequest, runtime *util.RuntimeOptions) (_result *EnterpriseRoleUpdateBizRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizPermissionCodeListJson)) {
		query["BizPermissionCodeListJson"] = request.BizPermissionCodeListJson
	}

	if !tea.BoolValue(util.IsUnset(request.BizRoleCode)) {
		query["BizRoleCode"] = request.BizRoleCode
	}

	if !tea.BoolValue(util.IsUnset(request.BizRoleDesc)) {
		query["BizRoleDesc"] = request.BizRoleDesc
	}

	if !tea.BoolValue(util.IsUnset(request.BizRoleName)) {
		query["BizRoleName"] = request.BizRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptTicket)) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseRoleUpdateBizRole"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseRoleUpdateBizRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新业务角色
//
// @param request - EnterpriseRoleUpdateBizRoleRequest
//
// @return EnterpriseRoleUpdateBizRoleResponse
func (client *Client) EnterpriseRoleUpdateBizRole(request *EnterpriseRoleUpdateBizRoleRequest) (_result *EnterpriseRoleUpdateBizRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseRoleUpdateBizRoleResponse{}
	_body, _err := client.EnterpriseRoleUpdateBizRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 处理待办项
//
// @param request - EnterpriseTodoDealAccountTodoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseTodoDealAccountTodoResponse
func (client *Client) EnterpriseTodoDealAccountTodoWithOptions(request *EnterpriseTodoDealAccountTodoRequest, runtime *util.RuntimeOptions) (_result *EnterpriseTodoDealAccountTodoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TodoId)) {
		body["TodoId"] = request.TodoId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseTodoDealAccountTodo"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseTodoDealAccountTodoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 处理待办项
//
// @param request - EnterpriseTodoDealAccountTodoRequest
//
// @return EnterpriseTodoDealAccountTodoResponse
func (client *Client) EnterpriseTodoDealAccountTodo(request *EnterpriseTodoDealAccountTodoRequest) (_result *EnterpriseTodoDealAccountTodoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseTodoDealAccountTodoResponse{}
	_body, _err := client.EnterpriseTodoDealAccountTodoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询当前登录用户处理的待办项列表
//
// @param request - EnterpriseTodoQueryAccountTodoListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseTodoQueryAccountTodoListResponse
func (client *Client) EnterpriseTodoQueryAccountTodoListWithOptions(request *EnterpriseTodoQueryAccountTodoListRequest, runtime *util.RuntimeOptions) (_result *EnterpriseTodoQueryAccountTodoListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.OperatePk)) {
		body["OperatePk"] = request.OperatePk
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ShowCompleteInfo)) {
		body["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TodoType)) {
		body["TodoType"] = request.TodoType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseTodoQueryAccountTodoList"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseTodoQueryAccountTodoListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询当前登录用户处理的待办项列表
//
// @param request - EnterpriseTodoQueryAccountTodoListRequest
//
// @return EnterpriseTodoQueryAccountTodoListResponse
func (client *Client) EnterpriseTodoQueryAccountTodoList(request *EnterpriseTodoQueryAccountTodoListRequest) (_result *EnterpriseTodoQueryAccountTodoListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseTodoQueryAccountTodoListResponse{}
	_body, _err := client.EnterpriseTodoQueryAccountTodoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询当前登录用户发起的待办项列表
//
// @param request - EnterpriseTodoQueryAccountTodoListByApplicantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseTodoQueryAccountTodoListByApplicantResponse
func (client *Client) EnterpriseTodoQueryAccountTodoListByApplicantWithOptions(request *EnterpriseTodoQueryAccountTodoListByApplicantRequest, runtime *util.RuntimeOptions) (_result *EnterpriseTodoQueryAccountTodoListByApplicantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.OperatePk)) {
		body["OperatePk"] = request.OperatePk
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedEcId)) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedLeId)) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrientedNbId)) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ShowCompleteInfo)) {
		body["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TodoType)) {
		body["TodoType"] = request.TodoType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseTodoQueryAccountTodoListByApplicant"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseTodoQueryAccountTodoListByApplicantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询当前登录用户发起的待办项列表
//
// @param request - EnterpriseTodoQueryAccountTodoListByApplicantRequest
//
// @return EnterpriseTodoQueryAccountTodoListByApplicantResponse
func (client *Client) EnterpriseTodoQueryAccountTodoListByApplicant(request *EnterpriseTodoQueryAccountTodoListByApplicantRequest) (_result *EnterpriseTodoQueryAccountTodoListByApplicantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseTodoQueryAccountTodoListByApplicantResponse{}
	_body, _err := client.EnterpriseTodoQueryAccountTodoListByApplicantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 管理员邀请纳管
//
// @param request - EnterpriseUninvitedAdminInviteJoinEnterpriseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseUninvitedAdminInviteJoinEnterpriseResponse
func (client *Client) EnterpriseUninvitedAdminInviteJoinEnterpriseWithOptions(request *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest, runtime *util.RuntimeOptions) (_result *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcId)) {
		query["EcId"] = request.EcId
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptTicket)) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !tea.BoolValue(util.IsUnset(request.InviteePk)) {
		query["InviteePk"] = request.InviteePk
	}

	if !tea.BoolValue(util.IsUnset(request.LeId)) {
		query["LeId"] = request.LeId
	}

	if !tea.BoolValue(util.IsUnset(request.NbId)) {
		query["NbId"] = request.NbId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterpriseUninvitedAdminInviteJoinEnterprise"),
		Version:     tea.String("2024-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterpriseUninvitedAdminInviteJoinEnterpriseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 管理员邀请纳管
//
// @param request - EnterpriseUninvitedAdminInviteJoinEnterpriseRequest
//
// @return EnterpriseUninvitedAdminInviteJoinEnterpriseResponse
func (client *Client) EnterpriseUninvitedAdminInviteJoinEnterprise(request *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) (_result *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterpriseUninvitedAdminInviteJoinEnterpriseResponse{}
	_body, _err := client.EnterpriseUninvitedAdminInviteJoinEnterpriseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
