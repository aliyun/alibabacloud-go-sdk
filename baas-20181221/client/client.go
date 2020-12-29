// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AcceptFabricInvitationRequest struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsAccepted *bool   `json:"IsAccepted,omitempty" xml:"IsAccepted,omitempty"`
}

func (s AcceptFabricInvitationRequest) String() string {
	return tea.Prettify(s)
}

func (s AcceptFabricInvitationRequest) GoString() string {
	return s.String()
}

func (s *AcceptFabricInvitationRequest) SetCode(v string) *AcceptFabricInvitationRequest {
	s.Code = &v
	return s
}

func (s *AcceptFabricInvitationRequest) SetIsAccepted(v bool) *AcceptFabricInvitationRequest {
	s.IsAccepted = &v
	return s
}

type AcceptFabricInvitationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AcceptFabricInvitationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AcceptFabricInvitationResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptFabricInvitationResponseBody) SetRequestId(v string) *AcceptFabricInvitationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcceptFabricInvitationResponseBody) SetErrorCode(v int32) *AcceptFabricInvitationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AcceptFabricInvitationResponseBody) SetSuccess(v bool) *AcceptFabricInvitationResponseBody {
	s.Success = &v
	return s
}

type AcceptFabricInvitationResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AcceptFabricInvitationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AcceptFabricInvitationResponse) String() string {
	return tea.Prettify(s)
}

func (s AcceptFabricInvitationResponse) GoString() string {
	return s.String()
}

func (s *AcceptFabricInvitationResponse) SetHeaders(v map[string]*string) *AcceptFabricInvitationResponse {
	s.Headers = v
	return s
}

func (s *AcceptFabricInvitationResponse) SetBody(v *AcceptFabricInvitationResponseBody) *AcceptFabricInvitationResponse {
	s.Body = v
	return s
}

type ApplyAntChainCertificateRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	UploadReq  *string `json:"UploadReq,omitempty" xml:"UploadReq,omitempty"`
}

func (s ApplyAntChainCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyAntChainCertificateRequest) GoString() string {
	return s.String()
}

func (s *ApplyAntChainCertificateRequest) SetAntChainId(v string) *ApplyAntChainCertificateRequest {
	s.AntChainId = &v
	return s
}

func (s *ApplyAntChainCertificateRequest) SetUploadReq(v string) *ApplyAntChainCertificateRequest {
	s.UploadReq = &v
	return s
}

type ApplyAntChainCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ApplyAntChainCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyAntChainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyAntChainCertificateResponseBody) SetRequestId(v string) *ApplyAntChainCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyAntChainCertificateResponseBody) SetResult(v string) *ApplyAntChainCertificateResponseBody {
	s.Result = &v
	return s
}

type ApplyAntChainCertificateResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyAntChainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyAntChainCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyAntChainCertificateResponse) GoString() string {
	return s.String()
}

func (s *ApplyAntChainCertificateResponse) SetHeaders(v map[string]*string) *ApplyAntChainCertificateResponse {
	s.Headers = v
	return s
}

func (s *ApplyAntChainCertificateResponse) SetBody(v *ApplyAntChainCertificateResponseBody) *ApplyAntChainCertificateResponse {
	s.Body = v
	return s
}

type ApplyAntChainCertificateWithKeyAutoCreationRequest struct {
	AntChainId           *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	ConsortiumId         *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	CountryName          *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	StateOrProvinceName  *string `json:"StateOrProvinceName,omitempty" xml:"StateOrProvinceName,omitempty"`
	LocalityName         *string `json:"LocalityName,omitempty" xml:"LocalityName,omitempty"`
	OrganizationName     *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	OrganizationUnitName *string `json:"OrganizationUnitName,omitempty" xml:"OrganizationUnitName,omitempty"`
	CommonName           *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
}

func (s ApplyAntChainCertificateWithKeyAutoCreationRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyAntChainCertificateWithKeyAutoCreationRequest) GoString() string {
	return s.String()
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationRequest) SetAntChainId(v string) *ApplyAntChainCertificateWithKeyAutoCreationRequest {
	s.AntChainId = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationRequest) SetConsortiumId(v string) *ApplyAntChainCertificateWithKeyAutoCreationRequest {
	s.ConsortiumId = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationRequest) SetPassword(v string) *ApplyAntChainCertificateWithKeyAutoCreationRequest {
	s.Password = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationRequest) SetCountryName(v string) *ApplyAntChainCertificateWithKeyAutoCreationRequest {
	s.CountryName = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationRequest) SetStateOrProvinceName(v string) *ApplyAntChainCertificateWithKeyAutoCreationRequest {
	s.StateOrProvinceName = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationRequest) SetLocalityName(v string) *ApplyAntChainCertificateWithKeyAutoCreationRequest {
	s.LocalityName = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationRequest) SetOrganizationName(v string) *ApplyAntChainCertificateWithKeyAutoCreationRequest {
	s.OrganizationName = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationRequest) SetOrganizationUnitName(v string) *ApplyAntChainCertificateWithKeyAutoCreationRequest {
	s.OrganizationUnitName = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationRequest) SetCommonName(v string) *ApplyAntChainCertificateWithKeyAutoCreationRequest {
	s.CommonName = &v
	return s
}

type ApplyAntChainCertificateWithKeyAutoCreationResponseBody struct {
	RequestId *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ApplyAntChainCertificateWithKeyAutoCreationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyAntChainCertificateWithKeyAutoCreationResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponseBody) SetRequestId(v string) *ApplyAntChainCertificateWithKeyAutoCreationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponseBody) SetResult(v *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResult) *ApplyAntChainCertificateWithKeyAutoCreationResponseBody {
	s.Result = v
	return s
}

type ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResult struct {
	PrivateKey   *string                                                                    `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	DownloadPath *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath `json:"DownloadPath,omitempty" xml:"DownloadPath,omitempty" type:"Struct"`
}

func (s ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResult) SetPrivateKey(v string) *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResult {
	s.PrivateKey = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResult) SetDownloadPath(v *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath) *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResult {
	s.DownloadPath = v
	return s
}

type ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath struct {
	CaCrtUrl     *string `json:"CaCrtUrl,omitempty" xml:"CaCrtUrl,omitempty"`
	SdkUrl       *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
	ClientCrtUrl *string `json:"ClientCrtUrl,omitempty" xml:"ClientCrtUrl,omitempty"`
	TrustCaUrl   *string `json:"TrustCaUrl,omitempty" xml:"TrustCaUrl,omitempty"`
}

func (s ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath) String() string {
	return tea.Prettify(s)
}

func (s ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath) GoString() string {
	return s.String()
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath) SetCaCrtUrl(v string) *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath {
	s.CaCrtUrl = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath) SetSdkUrl(v string) *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath {
	s.SdkUrl = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath) SetClientCrtUrl(v string) *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath {
	s.ClientCrtUrl = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath) SetTrustCaUrl(v string) *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath {
	s.TrustCaUrl = &v
	return s
}

type ApplyAntChainCertificateWithKeyAutoCreationResponse struct {
	Headers map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyAntChainCertificateWithKeyAutoCreationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyAntChainCertificateWithKeyAutoCreationResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyAntChainCertificateWithKeyAutoCreationResponse) GoString() string {
	return s.String()
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponse) SetHeaders(v map[string]*string) *ApplyAntChainCertificateWithKeyAutoCreationResponse {
	s.Headers = v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponse) SetBody(v *ApplyAntChainCertificateWithKeyAutoCreationResponseBody) *ApplyAntChainCertificateWithKeyAutoCreationResponse {
	s.Body = v
	return s
}

type BatchAddAntChainMiniAppQRCodeAuthorizedUsersRequest struct {
	AntChainId *string                `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	PhoneList  map[string]interface{} `json:"PhoneList,omitempty" xml:"PhoneList,omitempty"`
}

func (s BatchAddAntChainMiniAppQRCodeAuthorizedUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddAntChainMiniAppQRCodeAuthorizedUsersRequest) GoString() string {
	return s.String()
}

func (s *BatchAddAntChainMiniAppQRCodeAuthorizedUsersRequest) SetAntChainId(v string) *BatchAddAntChainMiniAppQRCodeAuthorizedUsersRequest {
	s.AntChainId = &v
	return s
}

func (s *BatchAddAntChainMiniAppQRCodeAuthorizedUsersRequest) SetPhoneList(v map[string]interface{}) *BatchAddAntChainMiniAppQRCodeAuthorizedUsersRequest {
	s.PhoneList = v
	return s
}

type BatchAddAntChainMiniAppQRCodeAuthorizedUsersShrinkRequest struct {
	AntChainId      *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	PhoneListShrink *string `json:"PhoneList,omitempty" xml:"PhoneList,omitempty"`
}

func (s BatchAddAntChainMiniAppQRCodeAuthorizedUsersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddAntChainMiniAppQRCodeAuthorizedUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchAddAntChainMiniAppQRCodeAuthorizedUsersShrinkRequest) SetAntChainId(v string) *BatchAddAntChainMiniAppQRCodeAuthorizedUsersShrinkRequest {
	s.AntChainId = &v
	return s
}

func (s *BatchAddAntChainMiniAppQRCodeAuthorizedUsersShrinkRequest) SetPhoneListShrink(v string) *BatchAddAntChainMiniAppQRCodeAuthorizedUsersShrinkRequest {
	s.PhoneListShrink = &v
	return s
}

type BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponseBody) SetRequestId(v string) *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponseBody) SetResult(v string) *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponseBody {
	s.Result = &v
	return s
}

type BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse struct {
	Headers map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse) GoString() string {
	return s.String()
}

func (s *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse) SetHeaders(v map[string]*string) *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse {
	s.Headers = v
	return s
}

func (s *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse) SetBody(v *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponseBody) *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse {
	s.Body = v
	return s
}

type CheckFabricConsortiumDomainRequest struct {
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
}

func (s CheckFabricConsortiumDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckFabricConsortiumDomainRequest) GoString() string {
	return s.String()
}

func (s *CheckFabricConsortiumDomainRequest) SetDomainCode(v string) *CheckFabricConsortiumDomainRequest {
	s.DomainCode = &v
	return s
}

type CheckFabricConsortiumDomainResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *CheckFabricConsortiumDomainResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CheckFabricConsortiumDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckFabricConsortiumDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CheckFabricConsortiumDomainResponseBody) SetRequestId(v string) *CheckFabricConsortiumDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckFabricConsortiumDomainResponseBody) SetErrorCode(v int32) *CheckFabricConsortiumDomainResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckFabricConsortiumDomainResponseBody) SetSuccess(v bool) *CheckFabricConsortiumDomainResponseBody {
	s.Success = &v
	return s
}

func (s *CheckFabricConsortiumDomainResponseBody) SetResult(v *CheckFabricConsortiumDomainResponseBodyResult) *CheckFabricConsortiumDomainResponseBody {
	s.Result = v
	return s
}

type CheckFabricConsortiumDomainResponseBodyResult struct {
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Valid  *bool   `json:"Valid,omitempty" xml:"Valid,omitempty"`
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
}

func (s CheckFabricConsortiumDomainResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CheckFabricConsortiumDomainResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CheckFabricConsortiumDomainResponseBodyResult) SetDomain(v string) *CheckFabricConsortiumDomainResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *CheckFabricConsortiumDomainResponseBodyResult) SetValid(v bool) *CheckFabricConsortiumDomainResponseBodyResult {
	s.Valid = &v
	return s
}

func (s *CheckFabricConsortiumDomainResponseBodyResult) SetPrompt(v string) *CheckFabricConsortiumDomainResponseBodyResult {
	s.Prompt = &v
	return s
}

type CheckFabricConsortiumDomainResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckFabricConsortiumDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckFabricConsortiumDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckFabricConsortiumDomainResponse) GoString() string {
	return s.String()
}

func (s *CheckFabricConsortiumDomainResponse) SetHeaders(v map[string]*string) *CheckFabricConsortiumDomainResponse {
	s.Headers = v
	return s
}

func (s *CheckFabricConsortiumDomainResponse) SetBody(v *CheckFabricConsortiumDomainResponseBody) *CheckFabricConsortiumDomainResponse {
	s.Body = v
	return s
}

type CheckFabricOrganizationDomainRequest struct {
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s CheckFabricOrganizationDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckFabricOrganizationDomainRequest) GoString() string {
	return s.String()
}

func (s *CheckFabricOrganizationDomainRequest) SetDomainCode(v string) *CheckFabricOrganizationDomainRequest {
	s.DomainCode = &v
	return s
}

func (s *CheckFabricOrganizationDomainRequest) SetDomain(v string) *CheckFabricOrganizationDomainRequest {
	s.Domain = &v
	return s
}

type CheckFabricOrganizationDomainResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *CheckFabricOrganizationDomainResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CheckFabricOrganizationDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckFabricOrganizationDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CheckFabricOrganizationDomainResponseBody) SetRequestId(v string) *CheckFabricOrganizationDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckFabricOrganizationDomainResponseBody) SetErrorCode(v int32) *CheckFabricOrganizationDomainResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckFabricOrganizationDomainResponseBody) SetSuccess(v bool) *CheckFabricOrganizationDomainResponseBody {
	s.Success = &v
	return s
}

func (s *CheckFabricOrganizationDomainResponseBody) SetResult(v *CheckFabricOrganizationDomainResponseBodyResult) *CheckFabricOrganizationDomainResponseBody {
	s.Result = v
	return s
}

type CheckFabricOrganizationDomainResponseBodyResult struct {
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Valid  *bool   `json:"Valid,omitempty" xml:"Valid,omitempty"`
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
}

func (s CheckFabricOrganizationDomainResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CheckFabricOrganizationDomainResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CheckFabricOrganizationDomainResponseBodyResult) SetDomain(v string) *CheckFabricOrganizationDomainResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *CheckFabricOrganizationDomainResponseBodyResult) SetValid(v bool) *CheckFabricOrganizationDomainResponseBodyResult {
	s.Valid = &v
	return s
}

func (s *CheckFabricOrganizationDomainResponseBodyResult) SetPrompt(v string) *CheckFabricOrganizationDomainResponseBodyResult {
	s.Prompt = &v
	return s
}

type CheckFabricOrganizationDomainResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckFabricOrganizationDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckFabricOrganizationDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckFabricOrganizationDomainResponse) GoString() string {
	return s.String()
}

func (s *CheckFabricOrganizationDomainResponse) SetHeaders(v map[string]*string) *CheckFabricOrganizationDomainResponse {
	s.Headers = v
	return s
}

func (s *CheckFabricOrganizationDomainResponse) SetBody(v *CheckFabricOrganizationDomainResponseBody) *CheckFabricOrganizationDomainResponse {
	s.Body = v
	return s
}

type ConfirmFabricConsortiumMemberRequest struct {
	ConsortiumId *string                                             `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	Organization []*ConfirmFabricConsortiumMemberRequestOrganization `json:"Organization,omitempty" xml:"Organization,omitempty" type:"Repeated"`
}

func (s ConfirmFabricConsortiumMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmFabricConsortiumMemberRequest) GoString() string {
	return s.String()
}

func (s *ConfirmFabricConsortiumMemberRequest) SetConsortiumId(v string) *ConfirmFabricConsortiumMemberRequest {
	s.ConsortiumId = &v
	return s
}

func (s *ConfirmFabricConsortiumMemberRequest) SetOrganization(v []*ConfirmFabricConsortiumMemberRequestOrganization) *ConfirmFabricConsortiumMemberRequest {
	s.Organization = v
	return s
}

type ConfirmFabricConsortiumMemberRequestOrganization struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ConfirmFabricConsortiumMemberRequestOrganization) String() string {
	return tea.Prettify(s)
}

func (s ConfirmFabricConsortiumMemberRequestOrganization) GoString() string {
	return s.String()
}

func (s *ConfirmFabricConsortiumMemberRequestOrganization) SetOrganizationId(v string) *ConfirmFabricConsortiumMemberRequestOrganization {
	s.OrganizationId = &v
	return s
}

type ConfirmFabricConsortiumMemberResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ConfirmFabricConsortiumMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmFabricConsortiumMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmFabricConsortiumMemberResponseBody) SetRequestId(v string) *ConfirmFabricConsortiumMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmFabricConsortiumMemberResponseBody) SetErrorCode(v int32) *ConfirmFabricConsortiumMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ConfirmFabricConsortiumMemberResponseBody) SetSuccess(v bool) *ConfirmFabricConsortiumMemberResponseBody {
	s.Success = &v
	return s
}

func (s *ConfirmFabricConsortiumMemberResponseBody) SetResult(v bool) *ConfirmFabricConsortiumMemberResponseBody {
	s.Result = &v
	return s
}

type ConfirmFabricConsortiumMemberResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfirmFabricConsortiumMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfirmFabricConsortiumMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmFabricConsortiumMemberResponse) GoString() string {
	return s.String()
}

func (s *ConfirmFabricConsortiumMemberResponse) SetHeaders(v map[string]*string) *ConfirmFabricConsortiumMemberResponse {
	s.Headers = v
	return s
}

func (s *ConfirmFabricConsortiumMemberResponse) SetBody(v *ConfirmFabricConsortiumMemberResponseBody) *ConfirmFabricConsortiumMemberResponse {
	s.Body = v
	return s
}

type CopyAntChainContractProjectRequest struct {
	ProjectId          *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectVersion     *string `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
}

func (s CopyAntChainContractProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyAntChainContractProjectRequest) GoString() string {
	return s.String()
}

func (s *CopyAntChainContractProjectRequest) SetProjectId(v string) *CopyAntChainContractProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *CopyAntChainContractProjectRequest) SetProjectName(v string) *CopyAntChainContractProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CopyAntChainContractProjectRequest) SetProjectVersion(v string) *CopyAntChainContractProjectRequest {
	s.ProjectVersion = &v
	return s
}

func (s *CopyAntChainContractProjectRequest) SetProjectDescription(v string) *CopyAntChainContractProjectRequest {
	s.ProjectDescription = &v
	return s
}

type CopyAntChainContractProjectResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CopyAntChainContractProjectResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CopyAntChainContractProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyAntChainContractProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CopyAntChainContractProjectResponseBody) SetRequestId(v string) *CopyAntChainContractProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyAntChainContractProjectResponseBody) SetResult(v *CopyAntChainContractProjectResponseBodyResult) *CopyAntChainContractProjectResponseBody {
	s.Result = v
	return s
}

type CopyAntChainContractProjectResponseBodyResult struct {
	UpdateTime         *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ConsortiumId       *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	CreateTime         *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ProjectId          *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectVersion     *string `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
}

func (s CopyAntChainContractProjectResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CopyAntChainContractProjectResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CopyAntChainContractProjectResponseBodyResult) SetUpdateTime(v int64) *CopyAntChainContractProjectResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *CopyAntChainContractProjectResponseBodyResult) SetConsortiumId(v string) *CopyAntChainContractProjectResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *CopyAntChainContractProjectResponseBodyResult) SetCreateTime(v int64) *CopyAntChainContractProjectResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *CopyAntChainContractProjectResponseBodyResult) SetProjectId(v string) *CopyAntChainContractProjectResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *CopyAntChainContractProjectResponseBodyResult) SetProjectName(v string) *CopyAntChainContractProjectResponseBodyResult {
	s.ProjectName = &v
	return s
}

func (s *CopyAntChainContractProjectResponseBodyResult) SetProjectVersion(v string) *CopyAntChainContractProjectResponseBodyResult {
	s.ProjectVersion = &v
	return s
}

func (s *CopyAntChainContractProjectResponseBodyResult) SetProjectDescription(v string) *CopyAntChainContractProjectResponseBodyResult {
	s.ProjectDescription = &v
	return s
}

type CopyAntChainContractProjectResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CopyAntChainContractProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CopyAntChainContractProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyAntChainContractProjectResponse) GoString() string {
	return s.String()
}

func (s *CopyAntChainContractProjectResponse) SetHeaders(v map[string]*string) *CopyAntChainContractProjectResponse {
	s.Headers = v
	return s
}

func (s *CopyAntChainContractProjectResponse) SetBody(v *CopyAntChainContractProjectResponseBody) *CopyAntChainContractProjectResponse {
	s.Body = v
	return s
}

type CreateAntChainAccountRequest struct {
	AntChainId           *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	Account              *string `json:"Account,omitempty" xml:"Account,omitempty"`
	AccountPubKey        *string `json:"AccountPubKey,omitempty" xml:"AccountPubKey,omitempty"`
	AccountRecoverPubKey *string `json:"AccountRecoverPubKey,omitempty" xml:"AccountRecoverPubKey,omitempty"`
}

func (s CreateAntChainAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAntChainAccountRequest) SetAntChainId(v string) *CreateAntChainAccountRequest {
	s.AntChainId = &v
	return s
}

func (s *CreateAntChainAccountRequest) SetAccount(v string) *CreateAntChainAccountRequest {
	s.Account = &v
	return s
}

func (s *CreateAntChainAccountRequest) SetAccountPubKey(v string) *CreateAntChainAccountRequest {
	s.AccountPubKey = &v
	return s
}

func (s *CreateAntChainAccountRequest) SetAccountRecoverPubKey(v string) *CreateAntChainAccountRequest {
	s.AccountRecoverPubKey = &v
	return s
}

type CreateAntChainAccountResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateAntChainAccountResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateAntChainAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAntChainAccountResponseBody) SetRequestId(v string) *CreateAntChainAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAntChainAccountResponseBody) SetResult(v *CreateAntChainAccountResponseBodyResult) *CreateAntChainAccountResponseBody {
	s.Result = v
	return s
}

type CreateAntChainAccountResponseBodyResult struct {
	Account    *string `json:"Account,omitempty" xml:"Account,omitempty"`
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s CreateAntChainAccountResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainAccountResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAntChainAccountResponseBodyResult) SetAccount(v string) *CreateAntChainAccountResponseBodyResult {
	s.Account = &v
	return s
}

func (s *CreateAntChainAccountResponseBodyResult) SetAntChainId(v string) *CreateAntChainAccountResponseBodyResult {
	s.AntChainId = &v
	return s
}

type CreateAntChainAccountResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAntChainAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAntChainAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateAntChainAccountResponse) SetHeaders(v map[string]*string) *CreateAntChainAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateAntChainAccountResponse) SetBody(v *CreateAntChainAccountResponseBody) *CreateAntChainAccountResponse {
	s.Body = v
	return s
}

type CreateAntChainAccountWithKeyPairAutoCreationRequest struct {
	AntChainId      *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	Account         *string `json:"Account,omitempty" xml:"Account,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RecoverPassword *string `json:"RecoverPassword,omitempty" xml:"RecoverPassword,omitempty"`
}

func (s CreateAntChainAccountWithKeyPairAutoCreationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainAccountWithKeyPairAutoCreationRequest) GoString() string {
	return s.String()
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationRequest) SetAntChainId(v string) *CreateAntChainAccountWithKeyPairAutoCreationRequest {
	s.AntChainId = &v
	return s
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationRequest) SetAccount(v string) *CreateAntChainAccountWithKeyPairAutoCreationRequest {
	s.Account = &v
	return s
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationRequest) SetPassword(v string) *CreateAntChainAccountWithKeyPairAutoCreationRequest {
	s.Password = &v
	return s
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationRequest) SetRecoverPassword(v string) *CreateAntChainAccountWithKeyPairAutoCreationRequest {
	s.RecoverPassword = &v
	return s
}

type CreateAntChainAccountWithKeyPairAutoCreationResponseBody struct {
	RequestId *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateAntChainAccountWithKeyPairAutoCreationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainAccountWithKeyPairAutoCreationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationResponseBody) SetRequestId(v string) *CreateAntChainAccountWithKeyPairAutoCreationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationResponseBody) SetResult(v *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult) *CreateAntChainAccountWithKeyPairAutoCreationResponseBody {
	s.Result = v
	return s
}

type CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult struct {
	AccountPublicKey         *string `json:"AccountPublicKey,omitempty" xml:"AccountPublicKey,omitempty"`
	Account                  *string `json:"Account,omitempty" xml:"Account,omitempty"`
	AccountRecoverPrivateKey *string `json:"AccountRecoverPrivateKey,omitempty" xml:"AccountRecoverPrivateKey,omitempty"`
	AccountRecoverPublicKey  *string `json:"AccountRecoverPublicKey,omitempty" xml:"AccountRecoverPublicKey,omitempty"`
	AccountPrivateKey        *string `json:"AccountPrivateKey,omitempty" xml:"AccountPrivateKey,omitempty"`
	AntChainId               *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult) SetAccountPublicKey(v string) *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult {
	s.AccountPublicKey = &v
	return s
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult) SetAccount(v string) *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult {
	s.Account = &v
	return s
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult) SetAccountRecoverPrivateKey(v string) *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult {
	s.AccountRecoverPrivateKey = &v
	return s
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult) SetAccountRecoverPublicKey(v string) *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult {
	s.AccountRecoverPublicKey = &v
	return s
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult) SetAccountPrivateKey(v string) *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult {
	s.AccountPrivateKey = &v
	return s
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult) SetAntChainId(v string) *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult {
	s.AntChainId = &v
	return s
}

type CreateAntChainAccountWithKeyPairAutoCreationResponse struct {
	Headers map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAntChainAccountWithKeyPairAutoCreationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAntChainAccountWithKeyPairAutoCreationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainAccountWithKeyPairAutoCreationResponse) GoString() string {
	return s.String()
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationResponse) SetHeaders(v map[string]*string) *CreateAntChainAccountWithKeyPairAutoCreationResponse {
	s.Headers = v
	return s
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationResponse) SetBody(v *CreateAntChainAccountWithKeyPairAutoCreationResponseBody) *CreateAntChainAccountWithKeyPairAutoCreationResponse {
	s.Body = v
	return s
}

type CreateAntChainConsortiumRequest struct {
	ConsortiumName        *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	ConsortiumDescription *string `json:"ConsortiumDescription,omitempty" xml:"ConsortiumDescription,omitempty"`
}

func (s CreateAntChainConsortiumRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainConsortiumRequest) GoString() string {
	return s.String()
}

func (s *CreateAntChainConsortiumRequest) SetConsortiumName(v string) *CreateAntChainConsortiumRequest {
	s.ConsortiumName = &v
	return s
}

func (s *CreateAntChainConsortiumRequest) SetConsortiumDescription(v string) *CreateAntChainConsortiumRequest {
	s.ConsortiumDescription = &v
	return s
}

type CreateAntChainConsortiumResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateAntChainConsortiumResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateAntChainConsortiumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainConsortiumResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAntChainConsortiumResponseBody) SetRequestId(v string) *CreateAntChainConsortiumResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAntChainConsortiumResponseBody) SetResult(v *CreateAntChainConsortiumResponseBodyResult) *CreateAntChainConsortiumResponseBody {
	s.Result = v
	return s
}

type CreateAntChainConsortiumResponseBodyResult struct {
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
}

func (s CreateAntChainConsortiumResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainConsortiumResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAntChainConsortiumResponseBodyResult) SetConsortiumId(v string) *CreateAntChainConsortiumResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

type CreateAntChainConsortiumResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAntChainConsortiumResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAntChainConsortiumResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainConsortiumResponse) GoString() string {
	return s.String()
}

func (s *CreateAntChainConsortiumResponse) SetHeaders(v map[string]*string) *CreateAntChainConsortiumResponse {
	s.Headers = v
	return s
}

func (s *CreateAntChainConsortiumResponse) SetBody(v *CreateAntChainConsortiumResponseBody) *CreateAntChainConsortiumResponse {
	s.Body = v
	return s
}

type CreateAntChainContractContentRequest struct {
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ParentContentId *string `json:"ParentContentId,omitempty" xml:"ParentContentId,omitempty"`
	ContentName     *string `json:"ContentName,omitempty" xml:"ContentName,omitempty"`
	IsDirectory     *bool   `json:"IsDirectory,omitempty" xml:"IsDirectory,omitempty"`
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s CreateAntChainContractContentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainContractContentRequest) GoString() string {
	return s.String()
}

func (s *CreateAntChainContractContentRequest) SetProjectId(v string) *CreateAntChainContractContentRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateAntChainContractContentRequest) SetParentContentId(v string) *CreateAntChainContractContentRequest {
	s.ParentContentId = &v
	return s
}

func (s *CreateAntChainContractContentRequest) SetContentName(v string) *CreateAntChainContractContentRequest {
	s.ContentName = &v
	return s
}

func (s *CreateAntChainContractContentRequest) SetIsDirectory(v bool) *CreateAntChainContractContentRequest {
	s.IsDirectory = &v
	return s
}

func (s *CreateAntChainContractContentRequest) SetContent(v string) *CreateAntChainContractContentRequest {
	s.Content = &v
	return s
}

type CreateAntChainContractContentResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateAntChainContractContentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateAntChainContractContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainContractContentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAntChainContractContentResponseBody) SetRequestId(v string) *CreateAntChainContractContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAntChainContractContentResponseBody) SetResult(v *CreateAntChainContractContentResponseBodyResult) *CreateAntChainContractContentResponseBody {
	s.Result = v
	return s
}

type CreateAntChainContractContentResponseBodyResult struct {
	ContentName     *string `json:"ContentName,omitempty" xml:"ContentName,omitempty"`
	ParentContentId *string `json:"ParentContentId,omitempty" xml:"ParentContentId,omitempty"`
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	IsDirectory     *bool   `json:"IsDirectory,omitempty" xml:"IsDirectory,omitempty"`
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ContentId       *string `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
}

func (s CreateAntChainContractContentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainContractContentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAntChainContractContentResponseBodyResult) SetContentName(v string) *CreateAntChainContractContentResponseBodyResult {
	s.ContentName = &v
	return s
}

func (s *CreateAntChainContractContentResponseBodyResult) SetParentContentId(v string) *CreateAntChainContractContentResponseBodyResult {
	s.ParentContentId = &v
	return s
}

func (s *CreateAntChainContractContentResponseBodyResult) SetUpdateTime(v string) *CreateAntChainContractContentResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *CreateAntChainContractContentResponseBodyResult) SetCreateTime(v string) *CreateAntChainContractContentResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *CreateAntChainContractContentResponseBodyResult) SetProjectId(v string) *CreateAntChainContractContentResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *CreateAntChainContractContentResponseBodyResult) SetIsDirectory(v bool) *CreateAntChainContractContentResponseBodyResult {
	s.IsDirectory = &v
	return s
}

func (s *CreateAntChainContractContentResponseBodyResult) SetContent(v string) *CreateAntChainContractContentResponseBodyResult {
	s.Content = &v
	return s
}

func (s *CreateAntChainContractContentResponseBodyResult) SetContentId(v string) *CreateAntChainContractContentResponseBodyResult {
	s.ContentId = &v
	return s
}

type CreateAntChainContractContentResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAntChainContractContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAntChainContractContentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainContractContentResponse) GoString() string {
	return s.String()
}

func (s *CreateAntChainContractContentResponse) SetHeaders(v map[string]*string) *CreateAntChainContractContentResponse {
	s.Headers = v
	return s
}

func (s *CreateAntChainContractContentResponse) SetBody(v *CreateAntChainContractContentResponseBody) *CreateAntChainContractContentResponse {
	s.Body = v
	return s
}

type CreateAntChainContractProjectRequest struct {
	ConsortiumId       *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectVersion     *string `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
}

func (s CreateAntChainContractProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainContractProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateAntChainContractProjectRequest) SetConsortiumId(v string) *CreateAntChainContractProjectRequest {
	s.ConsortiumId = &v
	return s
}

func (s *CreateAntChainContractProjectRequest) SetProjectName(v string) *CreateAntChainContractProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateAntChainContractProjectRequest) SetProjectVersion(v string) *CreateAntChainContractProjectRequest {
	s.ProjectVersion = &v
	return s
}

func (s *CreateAntChainContractProjectRequest) SetProjectDescription(v string) *CreateAntChainContractProjectRequest {
	s.ProjectDescription = &v
	return s
}

type CreateAntChainContractProjectResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateAntChainContractProjectResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateAntChainContractProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainContractProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAntChainContractProjectResponseBody) SetRequestId(v string) *CreateAntChainContractProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAntChainContractProjectResponseBody) SetResult(v *CreateAntChainContractProjectResponseBodyResult) *CreateAntChainContractProjectResponseBody {
	s.Result = v
	return s
}

type CreateAntChainContractProjectResponseBodyResult struct {
	UpdateTime         *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ConsortiumId       *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	CreateTime         *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ProjectId          *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectVersion     *string `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
}

func (s CreateAntChainContractProjectResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainContractProjectResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAntChainContractProjectResponseBodyResult) SetUpdateTime(v int64) *CreateAntChainContractProjectResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *CreateAntChainContractProjectResponseBodyResult) SetConsortiumId(v string) *CreateAntChainContractProjectResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *CreateAntChainContractProjectResponseBodyResult) SetCreateTime(v int64) *CreateAntChainContractProjectResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *CreateAntChainContractProjectResponseBodyResult) SetProjectId(v string) *CreateAntChainContractProjectResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *CreateAntChainContractProjectResponseBodyResult) SetProjectName(v string) *CreateAntChainContractProjectResponseBodyResult {
	s.ProjectName = &v
	return s
}

func (s *CreateAntChainContractProjectResponseBodyResult) SetProjectVersion(v string) *CreateAntChainContractProjectResponseBodyResult {
	s.ProjectVersion = &v
	return s
}

func (s *CreateAntChainContractProjectResponseBodyResult) SetProjectDescription(v string) *CreateAntChainContractProjectResponseBodyResult {
	s.ProjectDescription = &v
	return s
}

type CreateAntChainContractProjectResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAntChainContractProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAntChainContractProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainContractProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateAntChainContractProjectResponse) SetHeaders(v map[string]*string) *CreateAntChainContractProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateAntChainContractProjectResponse) SetBody(v *CreateAntChainContractProjectResponseBody) *CreateAntChainContractProjectResponse {
	s.Body = v
	return s
}

type CreateFabricChaincodeRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	ChannelId      *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ConsortiumId   *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	OssBucket      *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssUrl         *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	EndorsePolicy  *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	Location       *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s CreateFabricChaincodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChaincodeRequest) GoString() string {
	return s.String()
}

func (s *CreateFabricChaincodeRequest) SetOrganizationId(v string) *CreateFabricChaincodeRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateFabricChaincodeRequest) SetChannelId(v string) *CreateFabricChaincodeRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateFabricChaincodeRequest) SetConsortiumId(v string) *CreateFabricChaincodeRequest {
	s.ConsortiumId = &v
	return s
}

func (s *CreateFabricChaincodeRequest) SetOssBucket(v string) *CreateFabricChaincodeRequest {
	s.OssBucket = &v
	return s
}

func (s *CreateFabricChaincodeRequest) SetOssUrl(v string) *CreateFabricChaincodeRequest {
	s.OssUrl = &v
	return s
}

func (s *CreateFabricChaincodeRequest) SetEndorsePolicy(v string) *CreateFabricChaincodeRequest {
	s.EndorsePolicy = &v
	return s
}

func (s *CreateFabricChaincodeRequest) SetLocation(v string) *CreateFabricChaincodeRequest {
	s.Location = &v
	return s
}

type CreateFabricChaincodeResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *CreateFabricChaincodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateFabricChaincodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChaincodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFabricChaincodeResponseBody) SetRequestId(v string) *CreateFabricChaincodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFabricChaincodeResponseBody) SetErrorCode(v int32) *CreateFabricChaincodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFabricChaincodeResponseBody) SetSuccess(v bool) *CreateFabricChaincodeResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFabricChaincodeResponseBody) SetResult(v *CreateFabricChaincodeResponseBodyResult) *CreateFabricChaincodeResponseBody {
	s.Result = v
	return s
}

type CreateFabricChaincodeResponseBodyResult struct {
	Type             *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	EndorsePolicy    *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChaincodeId      *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ProviderName     *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	ChaincodeName    *string `json:"ChaincodeName,omitempty" xml:"ChaincodeName,omitempty"`
	Install          *bool   `json:"Install,omitempty" xml:"Install,omitempty"`
	Input            *string `json:"Input,omitempty" xml:"Input,omitempty"`
	ProviderId       *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	DeployTime       *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" xml:"ChaincodeVersion,omitempty"`
	ConsortiumId     *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ChannelName      *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	Path             *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateFabricChaincodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChaincodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateFabricChaincodeResponseBodyResult) SetType(v int32) *CreateFabricChaincodeResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetEndorsePolicy(v string) *CreateFabricChaincodeResponseBodyResult {
	s.EndorsePolicy = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetState(v string) *CreateFabricChaincodeResponseBodyResult {
	s.State = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetCreateTime(v string) *CreateFabricChaincodeResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetChaincodeId(v string) *CreateFabricChaincodeResponseBodyResult {
	s.ChaincodeId = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetMessage(v string) *CreateFabricChaincodeResponseBodyResult {
	s.Message = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetProviderName(v string) *CreateFabricChaincodeResponseBodyResult {
	s.ProviderName = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetChaincodeName(v string) *CreateFabricChaincodeResponseBodyResult {
	s.ChaincodeName = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetInstall(v bool) *CreateFabricChaincodeResponseBodyResult {
	s.Install = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetInput(v string) *CreateFabricChaincodeResponseBodyResult {
	s.Input = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetProviderId(v string) *CreateFabricChaincodeResponseBodyResult {
	s.ProviderId = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetDeployTime(v string) *CreateFabricChaincodeResponseBodyResult {
	s.DeployTime = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetChaincodeVersion(v string) *CreateFabricChaincodeResponseBodyResult {
	s.ChaincodeVersion = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetConsortiumId(v string) *CreateFabricChaincodeResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetChannelName(v string) *CreateFabricChaincodeResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetPath(v string) *CreateFabricChaincodeResponseBodyResult {
	s.Path = &v
	return s
}

type CreateFabricChaincodeResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFabricChaincodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFabricChaincodeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChaincodeResponse) GoString() string {
	return s.String()
}

func (s *CreateFabricChaincodeResponse) SetHeaders(v map[string]*string) *CreateFabricChaincodeResponse {
	s.Headers = v
	return s
}

func (s *CreateFabricChaincodeResponse) SetBody(v *CreateFabricChaincodeResponseBody) *CreateFabricChaincodeResponse {
	s.Body = v
	return s
}

type CreateFabricChannelRequest struct {
	ConsortiumId      *string                                   `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ChannelName       *string                                   `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	BatchTimeout      *int32                                    `json:"BatchTimeout,omitempty" xml:"BatchTimeout,omitempty"`
	MaxMessageCount   *int32                                    `json:"MaxMessageCount,omitempty" xml:"MaxMessageCount,omitempty"`
	PreferredMaxBytes *int32                                    `json:"PreferredMaxBytes,omitempty" xml:"PreferredMaxBytes,omitempty"`
	Organization      []*CreateFabricChannelRequestOrganization `json:"Organization,omitempty" xml:"Organization,omitempty" type:"Repeated"`
}

func (s CreateFabricChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateFabricChannelRequest) SetConsortiumId(v string) *CreateFabricChannelRequest {
	s.ConsortiumId = &v
	return s
}

func (s *CreateFabricChannelRequest) SetChannelName(v string) *CreateFabricChannelRequest {
	s.ChannelName = &v
	return s
}

func (s *CreateFabricChannelRequest) SetBatchTimeout(v int32) *CreateFabricChannelRequest {
	s.BatchTimeout = &v
	return s
}

func (s *CreateFabricChannelRequest) SetMaxMessageCount(v int32) *CreateFabricChannelRequest {
	s.MaxMessageCount = &v
	return s
}

func (s *CreateFabricChannelRequest) SetPreferredMaxBytes(v int32) *CreateFabricChannelRequest {
	s.PreferredMaxBytes = &v
	return s
}

func (s *CreateFabricChannelRequest) SetOrganization(v []*CreateFabricChannelRequestOrganization) *CreateFabricChannelRequest {
	s.Organization = v
	return s
}

type CreateFabricChannelRequestOrganization struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateFabricChannelRequestOrganization) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChannelRequestOrganization) GoString() string {
	return s.String()
}

func (s *CreateFabricChannelRequestOrganization) SetId(v string) *CreateFabricChannelRequestOrganization {
	s.Id = &v
	return s
}

type CreateFabricChannelResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *CreateFabricChannelResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateFabricChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFabricChannelResponseBody) SetRequestId(v string) *CreateFabricChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFabricChannelResponseBody) SetErrorCode(v int32) *CreateFabricChannelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFabricChannelResponseBody) SetSuccess(v bool) *CreateFabricChannelResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFabricChannelResponseBody) SetResult(v *CreateFabricChannelResponseBodyResult) *CreateFabricChannelResponseBody {
	s.Result = v
	return s
}

type CreateFabricChannelResponseBodyResult struct {
	BatchTimeout      *int32  `json:"BatchTimeout,omitempty" xml:"BatchTimeout,omitempty"`
	UpdateTime        *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ChaincodeCount    *int32  `json:"ChaincodeCount,omitempty" xml:"ChaincodeCount,omitempty"`
	State             *string `json:"State,omitempty" xml:"State,omitempty"`
	PreferredMaxBytes *int32  `json:"PreferredMaxBytes,omitempty" xml:"PreferredMaxBytes,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	OwnerName         *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerUid          *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	OwnerBid          *string `json:"OwnerBid,omitempty" xml:"OwnerBid,omitempty"`
	MaxMessageCount   *int32  `json:"MaxMessageCount,omitempty" xml:"MaxMessageCount,omitempty"`
	MemberCount       *int32  `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConsortiumId      *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ChannelName       *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	SupportConfig     *bool   `json:"SupportConfig,omitempty" xml:"SupportConfig,omitempty"`
	ChannelId         *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ConsortiumName    *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	BlockCount        *int32  `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
}

func (s CreateFabricChannelResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChannelResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateFabricChannelResponseBodyResult) SetBatchTimeout(v int32) *CreateFabricChannelResponseBodyResult {
	s.BatchTimeout = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetUpdateTime(v string) *CreateFabricChannelResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetChaincodeCount(v int32) *CreateFabricChannelResponseBodyResult {
	s.ChaincodeCount = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetState(v string) *CreateFabricChannelResponseBodyResult {
	s.State = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetPreferredMaxBytes(v int32) *CreateFabricChannelResponseBodyResult {
	s.PreferredMaxBytes = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetCreateTime(v string) *CreateFabricChannelResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetOwnerName(v string) *CreateFabricChannelResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetOwnerUid(v int64) *CreateFabricChannelResponseBodyResult {
	s.OwnerUid = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetOwnerBid(v string) *CreateFabricChannelResponseBodyResult {
	s.OwnerBid = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetMaxMessageCount(v int32) *CreateFabricChannelResponseBodyResult {
	s.MaxMessageCount = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetMemberCount(v int32) *CreateFabricChannelResponseBodyResult {
	s.MemberCount = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetRequestId(v string) *CreateFabricChannelResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetConsortiumId(v string) *CreateFabricChannelResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetChannelName(v string) *CreateFabricChannelResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetSupportConfig(v bool) *CreateFabricChannelResponseBodyResult {
	s.SupportConfig = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetChannelId(v string) *CreateFabricChannelResponseBodyResult {
	s.ChannelId = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetConsortiumName(v string) *CreateFabricChannelResponseBodyResult {
	s.ConsortiumName = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetBlockCount(v int32) *CreateFabricChannelResponseBodyResult {
	s.BlockCount = &v
	return s
}

type CreateFabricChannelResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFabricChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFabricChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChannelResponse) GoString() string {
	return s.String()
}

func (s *CreateFabricChannelResponse) SetHeaders(v map[string]*string) *CreateFabricChannelResponse {
	s.Headers = v
	return s
}

func (s *CreateFabricChannelResponse) SetBody(v *CreateFabricChannelResponseBody) *CreateFabricChannelResponse {
	s.Body = v
	return s
}

type CreateFabricChannelMemberRequest struct {
	ChannelId    *string                                         `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Organization []*CreateFabricChannelMemberRequestOrganization `json:"Organization,omitempty" xml:"Organization,omitempty" type:"Repeated"`
}

func (s CreateFabricChannelMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChannelMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateFabricChannelMemberRequest) SetChannelId(v string) *CreateFabricChannelMemberRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateFabricChannelMemberRequest) SetOrganization(v []*CreateFabricChannelMemberRequestOrganization) *CreateFabricChannelMemberRequest {
	s.Organization = v
	return s
}

type CreateFabricChannelMemberRequestOrganization struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateFabricChannelMemberRequestOrganization) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChannelMemberRequestOrganization) GoString() string {
	return s.String()
}

func (s *CreateFabricChannelMemberRequestOrganization) SetOrganizationId(v string) *CreateFabricChannelMemberRequestOrganization {
	s.OrganizationId = &v
	return s
}

type CreateFabricChannelMemberResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateFabricChannelMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChannelMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFabricChannelMemberResponseBody) SetRequestId(v string) *CreateFabricChannelMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFabricChannelMemberResponseBody) SetErrorCode(v int32) *CreateFabricChannelMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFabricChannelMemberResponseBody) SetSuccess(v bool) *CreateFabricChannelMemberResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFabricChannelMemberResponseBody) SetResult(v bool) *CreateFabricChannelMemberResponseBody {
	s.Result = &v
	return s
}

type CreateFabricChannelMemberResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFabricChannelMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFabricChannelMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChannelMemberResponse) GoString() string {
	return s.String()
}

func (s *CreateFabricChannelMemberResponse) SetHeaders(v map[string]*string) *CreateFabricChannelMemberResponse {
	s.Headers = v
	return s
}

func (s *CreateFabricChannelMemberResponse) SetBody(v *CreateFabricChannelMemberResponseBody) *CreateFabricChannelMemberResponse {
	s.Body = v
	return s
}

type CreateFabricConsortiumRequest struct {
	Location              *string                                      `json:"Location,omitempty" xml:"Location,omitempty"`
	OrdererType           *string                                      `json:"OrdererType,omitempty" xml:"OrdererType,omitempty"`
	ZoneId                *string                                      `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ConsortiumName        *string                                      `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	Domain                *string                                      `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ConsortiumDescription *string                                      `json:"ConsortiumDescription,omitempty" xml:"ConsortiumDescription,omitempty"`
	ChannelPolicy         *string                                      `json:"ChannelPolicy,omitempty" xml:"ChannelPolicy,omitempty"`
	SpecName              *string                                      `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	OrderersCount         *int32                                       `json:"OrderersCount,omitempty" xml:"OrderersCount,omitempty"`
	PeersCount            *int32                                       `json:"PeersCount,omitempty" xml:"PeersCount,omitempty"`
	PaymentDurationUnit   *string                                      `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	PaymentDuration       *int32                                       `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	Organization          []*CreateFabricConsortiumRequestOrganization `json:"Organization,omitempty" xml:"Organization,omitempty" type:"Repeated"`
}

func (s CreateFabricConsortiumRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricConsortiumRequest) GoString() string {
	return s.String()
}

func (s *CreateFabricConsortiumRequest) SetLocation(v string) *CreateFabricConsortiumRequest {
	s.Location = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetOrdererType(v string) *CreateFabricConsortiumRequest {
	s.OrdererType = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetZoneId(v string) *CreateFabricConsortiumRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetConsortiumName(v string) *CreateFabricConsortiumRequest {
	s.ConsortiumName = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetDomain(v string) *CreateFabricConsortiumRequest {
	s.Domain = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetConsortiumDescription(v string) *CreateFabricConsortiumRequest {
	s.ConsortiumDescription = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetChannelPolicy(v string) *CreateFabricConsortiumRequest {
	s.ChannelPolicy = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetSpecName(v string) *CreateFabricConsortiumRequest {
	s.SpecName = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetOrderersCount(v int32) *CreateFabricConsortiumRequest {
	s.OrderersCount = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetPeersCount(v int32) *CreateFabricConsortiumRequest {
	s.PeersCount = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetPaymentDurationUnit(v string) *CreateFabricConsortiumRequest {
	s.PaymentDurationUnit = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetPaymentDuration(v int32) *CreateFabricConsortiumRequest {
	s.PaymentDuration = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetOrganization(v []*CreateFabricConsortiumRequestOrganization) *CreateFabricConsortiumRequest {
	s.Organization = v
	return s
}

type CreateFabricConsortiumRequestOrganization struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateFabricConsortiumRequestOrganization) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricConsortiumRequestOrganization) GoString() string {
	return s.String()
}

func (s *CreateFabricConsortiumRequestOrganization) SetOrganizationId(v string) *CreateFabricConsortiumRequestOrganization {
	s.OrganizationId = &v
	return s
}

type CreateFabricConsortiumResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *CreateFabricConsortiumResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateFabricConsortiumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricConsortiumResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFabricConsortiumResponseBody) SetRequestId(v string) *CreateFabricConsortiumResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFabricConsortiumResponseBody) SetErrorCode(v int32) *CreateFabricConsortiumResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFabricConsortiumResponseBody) SetSuccess(v bool) *CreateFabricConsortiumResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFabricConsortiumResponseBody) SetResult(v *CreateFabricConsortiumResponseBodyResult) *CreateFabricConsortiumResponseBody {
	s.Result = v
	return s
}

type CreateFabricConsortiumResponseBodyResult struct {
	ChannelCount   *int32  `json:"ChannelCount,omitempty" xml:"ChannelCount,omitempty"`
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SpecName       *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	OrdererCount   *int32  `json:"OrdererCount,omitempty" xml:"OrdererCount,omitempty"`
	ServiceState   *string `json:"ServiceState,omitempty" xml:"ServiceState,omitempty"`
	OwnerUid       *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	ClusterState   *string `json:"ClusterState,omitempty" xml:"ClusterState,omitempty"`
	CodeName       *string `json:"CodeName,omitempty" xml:"CodeName,omitempty"`
	OwnerBid       *string `json:"OwnerBid,omitempty" xml:"OwnerBid,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MemberCount    *int32  `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	ChannelPolicy  *string `json:"ChannelPolicy,omitempty" xml:"ChannelPolicy,omitempty"`
	OrdererType    *string `json:"OrdererType,omitempty" xml:"OrdererType,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ConsortiumId   *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ZoneId         *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
}

func (s CreateFabricConsortiumResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricConsortiumResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateFabricConsortiumResponseBodyResult) SetChannelCount(v int32) *CreateFabricConsortiumResponseBodyResult {
	s.ChannelCount = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetDomain(v string) *CreateFabricConsortiumResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetCreateTime(v string) *CreateFabricConsortiumResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetSpecName(v string) *CreateFabricConsortiumResponseBodyResult {
	s.SpecName = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetOrdererCount(v int32) *CreateFabricConsortiumResponseBodyResult {
	s.OrdererCount = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetServiceState(v string) *CreateFabricConsortiumResponseBodyResult {
	s.ServiceState = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetOwnerUid(v int64) *CreateFabricConsortiumResponseBodyResult {
	s.OwnerUid = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetClusterState(v string) *CreateFabricConsortiumResponseBodyResult {
	s.ClusterState = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetCodeName(v string) *CreateFabricConsortiumResponseBodyResult {
	s.CodeName = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetOwnerBid(v string) *CreateFabricConsortiumResponseBodyResult {
	s.OwnerBid = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetRegionId(v string) *CreateFabricConsortiumResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetMemberCount(v int32) *CreateFabricConsortiumResponseBodyResult {
	s.MemberCount = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetChannelPolicy(v string) *CreateFabricConsortiumResponseBodyResult {
	s.ChannelPolicy = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetOrdererType(v string) *CreateFabricConsortiumResponseBodyResult {
	s.OrdererType = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetDescription(v string) *CreateFabricConsortiumResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetConsortiumId(v string) *CreateFabricConsortiumResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetZoneId(v string) *CreateFabricConsortiumResponseBodyResult {
	s.ZoneId = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetConsortiumName(v string) *CreateFabricConsortiumResponseBodyResult {
	s.ConsortiumName = &v
	return s
}

type CreateFabricConsortiumResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFabricConsortiumResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFabricConsortiumResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricConsortiumResponse) GoString() string {
	return s.String()
}

func (s *CreateFabricConsortiumResponse) SetHeaders(v map[string]*string) *CreateFabricConsortiumResponse {
	s.Headers = v
	return s
}

func (s *CreateFabricConsortiumResponse) SetBody(v *CreateFabricConsortiumResponseBody) *CreateFabricConsortiumResponse {
	s.Body = v
	return s
}

type CreateFabricConsortiumMemberRequest struct {
	ConsortiumId *string                                            `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	Code         *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Organization []*CreateFabricConsortiumMemberRequestOrganization `json:"Organization,omitempty" xml:"Organization,omitempty" type:"Repeated"`
}

func (s CreateFabricConsortiumMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricConsortiumMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateFabricConsortiumMemberRequest) SetConsortiumId(v string) *CreateFabricConsortiumMemberRequest {
	s.ConsortiumId = &v
	return s
}

func (s *CreateFabricConsortiumMemberRequest) SetCode(v string) *CreateFabricConsortiumMemberRequest {
	s.Code = &v
	return s
}

func (s *CreateFabricConsortiumMemberRequest) SetOrganization(v []*CreateFabricConsortiumMemberRequestOrganization) *CreateFabricConsortiumMemberRequest {
	s.Organization = v
	return s
}

type CreateFabricConsortiumMemberRequestOrganization struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateFabricConsortiumMemberRequestOrganization) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricConsortiumMemberRequestOrganization) GoString() string {
	return s.String()
}

func (s *CreateFabricConsortiumMemberRequestOrganization) SetOrganizationId(v string) *CreateFabricConsortiumMemberRequestOrganization {
	s.OrganizationId = &v
	return s
}

type CreateFabricConsortiumMemberResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateFabricConsortiumMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricConsortiumMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFabricConsortiumMemberResponseBody) SetRequestId(v string) *CreateFabricConsortiumMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFabricConsortiumMemberResponseBody) SetErrorCode(v int32) *CreateFabricConsortiumMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFabricConsortiumMemberResponseBody) SetSuccess(v bool) *CreateFabricConsortiumMemberResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFabricConsortiumMemberResponseBody) SetResult(v bool) *CreateFabricConsortiumMemberResponseBody {
	s.Result = &v
	return s
}

type CreateFabricConsortiumMemberResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFabricConsortiumMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFabricConsortiumMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricConsortiumMemberResponse) GoString() string {
	return s.String()
}

func (s *CreateFabricConsortiumMemberResponse) SetHeaders(v map[string]*string) *CreateFabricConsortiumMemberResponse {
	s.Headers = v
	return s
}

func (s *CreateFabricConsortiumMemberResponse) SetBody(v *CreateFabricConsortiumMemberResponseBody) *CreateFabricConsortiumMemberResponse {
	s.Body = v
	return s
}

type CreateFabricOrganizationRequest struct {
	OrganizationName    *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	Location            *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Domain              *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SpecName            *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	PeersCount          *int32  `json:"PeersCount,omitempty" xml:"PeersCount,omitempty"`
	PaymentDuration     *int32  `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
}

func (s CreateFabricOrganizationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricOrganizationRequest) GoString() string {
	return s.String()
}

func (s *CreateFabricOrganizationRequest) SetOrganizationName(v string) *CreateFabricOrganizationRequest {
	s.OrganizationName = &v
	return s
}

func (s *CreateFabricOrganizationRequest) SetLocation(v string) *CreateFabricOrganizationRequest {
	s.Location = &v
	return s
}

func (s *CreateFabricOrganizationRequest) SetDomain(v string) *CreateFabricOrganizationRequest {
	s.Domain = &v
	return s
}

func (s *CreateFabricOrganizationRequest) SetDescription(v string) *CreateFabricOrganizationRequest {
	s.Description = &v
	return s
}

func (s *CreateFabricOrganizationRequest) SetSpecName(v string) *CreateFabricOrganizationRequest {
	s.SpecName = &v
	return s
}

func (s *CreateFabricOrganizationRequest) SetPeersCount(v int32) *CreateFabricOrganizationRequest {
	s.PeersCount = &v
	return s
}

func (s *CreateFabricOrganizationRequest) SetPaymentDuration(v int32) *CreateFabricOrganizationRequest {
	s.PaymentDuration = &v
	return s
}

func (s *CreateFabricOrganizationRequest) SetPaymentDurationUnit(v string) *CreateFabricOrganizationRequest {
	s.PaymentDurationUnit = &v
	return s
}

type CreateFabricOrganizationResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *CreateFabricOrganizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateFabricOrganizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFabricOrganizationResponseBody) SetRequestId(v string) *CreateFabricOrganizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFabricOrganizationResponseBody) SetErrorCode(v int32) *CreateFabricOrganizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFabricOrganizationResponseBody) SetSuccess(v bool) *CreateFabricOrganizationResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFabricOrganizationResponseBody) SetResult(v *CreateFabricOrganizationResponseBodyResult) *CreateFabricOrganizationResponseBody {
	s.Result = v
	return s
}

type CreateFabricOrganizationResponseBodyResult struct {
	Domain                  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	PeerCount               *int32  `json:"PeerCount,omitempty" xml:"PeerCount,omitempty"`
	CreateTime              *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ConsortiumCount         *int32  `json:"ConsortiumCount,omitempty" xml:"ConsortiumCount,omitempty"`
	SpecName                *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	OwnerName               *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	ServiceState            *string `json:"ServiceState,omitempty" xml:"ServiceState,omitempty"`
	OwnerUid                *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	ClusterState            *string `json:"ClusterState,omitempty" xml:"ClusterState,omitempty"`
	CodeName                *string `json:"CodeName,omitempty" xml:"CodeName,omitempty"`
	OwnerBid                *string `json:"OwnerBid,omitempty" xml:"OwnerBid,omitempty"`
	OrganizationDescription *string `json:"OrganizationDescription,omitempty" xml:"OrganizationDescription,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OrganizationId          *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ZoneId                  *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	UserCount               *int32  `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	OrganizationName        *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
}

func (s CreateFabricOrganizationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricOrganizationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateFabricOrganizationResponseBodyResult) SetDomain(v string) *CreateFabricOrganizationResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetPeerCount(v int32) *CreateFabricOrganizationResponseBodyResult {
	s.PeerCount = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetCreateTime(v string) *CreateFabricOrganizationResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetConsortiumCount(v int32) *CreateFabricOrganizationResponseBodyResult {
	s.ConsortiumCount = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetSpecName(v string) *CreateFabricOrganizationResponseBodyResult {
	s.SpecName = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetOwnerName(v string) *CreateFabricOrganizationResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetServiceState(v string) *CreateFabricOrganizationResponseBodyResult {
	s.ServiceState = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetOwnerUid(v int64) *CreateFabricOrganizationResponseBodyResult {
	s.OwnerUid = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetClusterState(v string) *CreateFabricOrganizationResponseBodyResult {
	s.ClusterState = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetCodeName(v string) *CreateFabricOrganizationResponseBodyResult {
	s.CodeName = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetOwnerBid(v string) *CreateFabricOrganizationResponseBodyResult {
	s.OwnerBid = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetOrganizationDescription(v string) *CreateFabricOrganizationResponseBodyResult {
	s.OrganizationDescription = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetRegionId(v string) *CreateFabricOrganizationResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetOrganizationId(v string) *CreateFabricOrganizationResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetRequestId(v string) *CreateFabricOrganizationResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetZoneId(v string) *CreateFabricOrganizationResponseBodyResult {
	s.ZoneId = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetUserCount(v int32) *CreateFabricOrganizationResponseBodyResult {
	s.UserCount = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetOrganizationName(v string) *CreateFabricOrganizationResponseBodyResult {
	s.OrganizationName = &v
	return s
}

type CreateFabricOrganizationResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFabricOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFabricOrganizationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricOrganizationResponse) GoString() string {
	return s.String()
}

func (s *CreateFabricOrganizationResponse) SetHeaders(v map[string]*string) *CreateFabricOrganizationResponse {
	s.Headers = v
	return s
}

func (s *CreateFabricOrganizationResponse) SetBody(v *CreateFabricOrganizationResponseBody) *CreateFabricOrganizationResponse {
	s.Body = v
	return s
}

type CreateFabricOrganizationUserRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Username       *string `json:"Username,omitempty" xml:"Username,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Attrs          *string `json:"Attrs,omitempty" xml:"Attrs,omitempty"`
}

func (s CreateFabricOrganizationUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricOrganizationUserRequest) GoString() string {
	return s.String()
}

func (s *CreateFabricOrganizationUserRequest) SetOrganizationId(v string) *CreateFabricOrganizationUserRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateFabricOrganizationUserRequest) SetUsername(v string) *CreateFabricOrganizationUserRequest {
	s.Username = &v
	return s
}

func (s *CreateFabricOrganizationUserRequest) SetPassword(v string) *CreateFabricOrganizationUserRequest {
	s.Password = &v
	return s
}

func (s *CreateFabricOrganizationUserRequest) SetAttrs(v string) *CreateFabricOrganizationUserRequest {
	s.Attrs = &v
	return s
}

type CreateFabricOrganizationUserResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *CreateFabricOrganizationUserResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateFabricOrganizationUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricOrganizationUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFabricOrganizationUserResponseBody) SetRequestId(v string) *CreateFabricOrganizationUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFabricOrganizationUserResponseBody) SetErrorCode(v int32) *CreateFabricOrganizationUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFabricOrganizationUserResponseBody) SetSuccess(v bool) *CreateFabricOrganizationUserResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFabricOrganizationUserResponseBody) SetResult(v *CreateFabricOrganizationUserResponseBodyResult) *CreateFabricOrganizationUserResponseBody {
	s.Result = v
	return s
}

type CreateFabricOrganizationUserResponseBodyResult struct {
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ExpireTime     *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Username       *string `json:"Username,omitempty" xml:"Username,omitempty"`
	Fullname       *string `json:"Fullname,omitempty" xml:"Fullname,omitempty"`
}

func (s CreateFabricOrganizationUserResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricOrganizationUserResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateFabricOrganizationUserResponseBodyResult) SetPassword(v string) *CreateFabricOrganizationUserResponseBodyResult {
	s.Password = &v
	return s
}

func (s *CreateFabricOrganizationUserResponseBodyResult) SetExpireTime(v string) *CreateFabricOrganizationUserResponseBodyResult {
	s.ExpireTime = &v
	return s
}

func (s *CreateFabricOrganizationUserResponseBodyResult) SetCreateTime(v string) *CreateFabricOrganizationUserResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *CreateFabricOrganizationUserResponseBodyResult) SetOrganizationId(v string) *CreateFabricOrganizationUserResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *CreateFabricOrganizationUserResponseBodyResult) SetUsername(v string) *CreateFabricOrganizationUserResponseBodyResult {
	s.Username = &v
	return s
}

func (s *CreateFabricOrganizationUserResponseBodyResult) SetFullname(v string) *CreateFabricOrganizationUserResponseBodyResult {
	s.Fullname = &v
	return s
}

type CreateFabricOrganizationUserResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFabricOrganizationUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFabricOrganizationUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricOrganizationUserResponse) GoString() string {
	return s.String()
}

func (s *CreateFabricOrganizationUserResponse) SetHeaders(v map[string]*string) *CreateFabricOrganizationUserResponse {
	s.Headers = v
	return s
}

func (s *CreateFabricOrganizationUserResponse) SetBody(v *CreateFabricOrganizationUserResponseBody) *CreateFabricOrganizationUserResponse {
	s.Body = v
	return s
}

type DeleteAntChainConsortiumRequest struct {
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
}

func (s DeleteAntChainConsortiumRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntChainConsortiumRequest) GoString() string {
	return s.String()
}

func (s *DeleteAntChainConsortiumRequest) SetConsortiumId(v string) *DeleteAntChainConsortiumRequest {
	s.ConsortiumId = &v
	return s
}

type DeleteAntChainConsortiumResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteAntChainConsortiumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntChainConsortiumResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAntChainConsortiumResponseBody) SetRequestId(v string) *DeleteAntChainConsortiumResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAntChainConsortiumResponseBody) SetResult(v string) *DeleteAntChainConsortiumResponseBody {
	s.Result = &v
	return s
}

type DeleteAntChainConsortiumResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAntChainConsortiumResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAntChainConsortiumResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntChainConsortiumResponse) GoString() string {
	return s.String()
}

func (s *DeleteAntChainConsortiumResponse) SetHeaders(v map[string]*string) *DeleteAntChainConsortiumResponse {
	s.Headers = v
	return s
}

func (s *DeleteAntChainConsortiumResponse) SetBody(v *DeleteAntChainConsortiumResponseBody) *DeleteAntChainConsortiumResponse {
	s.Body = v
	return s
}

type DeleteAntChainContractContentRequest struct {
	ContentId *string `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
}

func (s DeleteAntChainContractContentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntChainContractContentRequest) GoString() string {
	return s.String()
}

func (s *DeleteAntChainContractContentRequest) SetContentId(v string) *DeleteAntChainContractContentRequest {
	s.ContentId = &v
	return s
}

type DeleteAntChainContractContentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteAntChainContractContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntChainContractContentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAntChainContractContentResponseBody) SetRequestId(v string) *DeleteAntChainContractContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAntChainContractContentResponseBody) SetResult(v string) *DeleteAntChainContractContentResponseBody {
	s.Result = &v
	return s
}

type DeleteAntChainContractContentResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAntChainContractContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAntChainContractContentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntChainContractContentResponse) GoString() string {
	return s.String()
}

func (s *DeleteAntChainContractContentResponse) SetHeaders(v map[string]*string) *DeleteAntChainContractContentResponse {
	s.Headers = v
	return s
}

func (s *DeleteAntChainContractContentResponse) SetBody(v *DeleteAntChainContractContentResponseBody) *DeleteAntChainContractContentResponse {
	s.Body = v
	return s
}

type DeleteAntChainContractProjectRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteAntChainContractProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntChainContractProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteAntChainContractProjectRequest) SetProjectId(v string) *DeleteAntChainContractProjectRequest {
	s.ProjectId = &v
	return s
}

type DeleteAntChainContractProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteAntChainContractProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntChainContractProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAntChainContractProjectResponseBody) SetRequestId(v string) *DeleteAntChainContractProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAntChainContractProjectResponseBody) SetResult(v string) *DeleteAntChainContractProjectResponseBody {
	s.Result = &v
	return s
}

type DeleteAntChainContractProjectResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAntChainContractProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAntChainContractProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntChainContractProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteAntChainContractProjectResponse) SetHeaders(v map[string]*string) *DeleteAntChainContractProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteAntChainContractProjectResponse) SetBody(v *DeleteAntChainContractProjectResponseBody) *DeleteAntChainContractProjectResponse {
	s.Body = v
	return s
}

type DeleteAntChainMiniAppQRCodeAuthorizedUserRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	Phone      *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s DeleteAntChainMiniAppQRCodeAuthorizedUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntChainMiniAppQRCodeAuthorizedUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteAntChainMiniAppQRCodeAuthorizedUserRequest) SetAntChainId(v string) *DeleteAntChainMiniAppQRCodeAuthorizedUserRequest {
	s.AntChainId = &v
	return s
}

func (s *DeleteAntChainMiniAppQRCodeAuthorizedUserRequest) SetPhone(v string) *DeleteAntChainMiniAppQRCodeAuthorizedUserRequest {
	s.Phone = &v
	return s
}

type DeleteAntChainMiniAppQRCodeAuthorizedUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteAntChainMiniAppQRCodeAuthorizedUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntChainMiniAppQRCodeAuthorizedUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAntChainMiniAppQRCodeAuthorizedUserResponseBody) SetRequestId(v string) *DeleteAntChainMiniAppQRCodeAuthorizedUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAntChainMiniAppQRCodeAuthorizedUserResponseBody) SetResult(v string) *DeleteAntChainMiniAppQRCodeAuthorizedUserResponseBody {
	s.Result = &v
	return s
}

type DeleteAntChainMiniAppQRCodeAuthorizedUserResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAntChainMiniAppQRCodeAuthorizedUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAntChainMiniAppQRCodeAuthorizedUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntChainMiniAppQRCodeAuthorizedUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteAntChainMiniAppQRCodeAuthorizedUserResponse) SetHeaders(v map[string]*string) *DeleteAntChainMiniAppQRCodeAuthorizedUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteAntChainMiniAppQRCodeAuthorizedUserResponse) SetBody(v *DeleteAntChainMiniAppQRCodeAuthorizedUserResponseBody) *DeleteAntChainMiniAppQRCodeAuthorizedUserResponse {
	s.Body = v
	return s
}

type DeleteFabricChaincodeRequest struct {
	ChaincodeId *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
}

func (s DeleteFabricChaincodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFabricChaincodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteFabricChaincodeRequest) SetChaincodeId(v string) *DeleteFabricChaincodeRequest {
	s.ChaincodeId = &v
	return s
}

type DeleteFabricChaincodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFabricChaincodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFabricChaincodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFabricChaincodeResponseBody) SetRequestId(v string) *DeleteFabricChaincodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFabricChaincodeResponseBody) SetErrorCode(v int32) *DeleteFabricChaincodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFabricChaincodeResponseBody) SetSuccess(v bool) *DeleteFabricChaincodeResponseBody {
	s.Success = &v
	return s
}

type DeleteFabricChaincodeResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFabricChaincodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFabricChaincodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFabricChaincodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteFabricChaincodeResponse) SetHeaders(v map[string]*string) *DeleteFabricChaincodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteFabricChaincodeResponse) SetBody(v *DeleteFabricChaincodeResponseBody) *DeleteFabricChaincodeResponse {
	s.Body = v
	return s
}

type DescribeAntChainAccountsRequest struct {
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s DescribeAntChainAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainAccountsRequest) SetPageSize(v int32) *DescribeAntChainAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainAccountsRequest) SetPageNumber(v int32) *DescribeAntChainAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainAccountsRequest) SetAntChainId(v string) *DescribeAntChainAccountsRequest {
	s.AntChainId = &v
	return s
}

type DescribeAntChainAccountsResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainAccountsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAntChainAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainAccountsResponseBody) SetRequestId(v string) *DescribeAntChainAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainAccountsResponseBody) SetResult(v *DescribeAntChainAccountsResponseBodyResult) *DescribeAntChainAccountsResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainAccountsResponseBodyResult struct {
	Pagination *DescribeAntChainAccountsResponseBodyResultPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	Accounts   []*DescribeAntChainAccountsResponseBodyResultAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
}

func (s DescribeAntChainAccountsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainAccountsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainAccountsResponseBodyResult) SetPagination(v *DescribeAntChainAccountsResponseBodyResultPagination) *DescribeAntChainAccountsResponseBodyResult {
	s.Pagination = v
	return s
}

func (s *DescribeAntChainAccountsResponseBodyResult) SetAccounts(v []*DescribeAntChainAccountsResponseBodyResultAccounts) *DescribeAntChainAccountsResponseBodyResult {
	s.Accounts = v
	return s
}

type DescribeAntChainAccountsResponseBodyResultPagination struct {
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainAccountsResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainAccountsResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainAccountsResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainAccountsResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainAccountsResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainAccountsResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainAccountsResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainAccountsResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainAccountsResponseBodyResultAccounts struct {
	AccountPublicKey   *string `json:"AccountPublicKey,omitempty" xml:"AccountPublicKey,omitempty"`
	Account            *string `json:"Account,omitempty" xml:"Account,omitempty"`
	AccountStatus      *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	AccountRecoveryKey *string `json:"AccountRecoveryKey,omitempty" xml:"AccountRecoveryKey,omitempty"`
	AntChainId         *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s DescribeAntChainAccountsResponseBodyResultAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainAccountsResponseBodyResultAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAntChainAccountsResponseBodyResultAccounts) SetAccountPublicKey(v string) *DescribeAntChainAccountsResponseBodyResultAccounts {
	s.AccountPublicKey = &v
	return s
}

func (s *DescribeAntChainAccountsResponseBodyResultAccounts) SetAccount(v string) *DescribeAntChainAccountsResponseBodyResultAccounts {
	s.Account = &v
	return s
}

func (s *DescribeAntChainAccountsResponseBodyResultAccounts) SetAccountStatus(v string) *DescribeAntChainAccountsResponseBodyResultAccounts {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAntChainAccountsResponseBodyResultAccounts) SetAccountRecoveryKey(v string) *DescribeAntChainAccountsResponseBodyResultAccounts {
	s.AccountRecoveryKey = &v
	return s
}

func (s *DescribeAntChainAccountsResponseBodyResultAccounts) SetAntChainId(v string) *DescribeAntChainAccountsResponseBodyResultAccounts {
	s.AntChainId = &v
	return s
}

type DescribeAntChainAccountsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainAccountsResponse) SetHeaders(v map[string]*string) *DescribeAntChainAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainAccountsResponse) SetBody(v *DescribeAntChainAccountsResponseBody) *DescribeAntChainAccountsResponse {
	s.Body = v
	return s
}

type DescribeAntChainBlockRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	Height     *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s DescribeAntChainBlockRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainBlockRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainBlockRequest) SetAntChainId(v string) *DescribeAntChainBlockRequest {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainBlockRequest) SetHeight(v int64) *DescribeAntChainBlockRequest {
	s.Height = &v
	return s
}

type DescribeAntChainBlockResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainBlockResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAntChainBlockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainBlockResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainBlockResponseBody) SetRequestId(v string) *DescribeAntChainBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainBlockResponseBody) SetResult(v *DescribeAntChainBlockResponseBodyResult) *DescribeAntChainBlockResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainBlockResponseBodyResult struct {
	PreviousHash     *string `json:"PreviousHash,omitempty" xml:"PreviousHash,omitempty"`
	Version          *int64  `json:"Version,omitempty" xml:"Version,omitempty"`
	TransactionSize  *int32  `json:"TransactionSize,omitempty" xml:"TransactionSize,omitempty"`
	CreateTime       *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RootTxHash       *string `json:"RootTxHash,omitempty" xml:"RootTxHash,omitempty"`
	Height           *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	BlockHash        *string `json:"BlockHash,omitempty" xml:"BlockHash,omitempty"`
	AntChainId       *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	TransSummaryList *string `json:"TransSummaryList,omitempty" xml:"TransSummaryList,omitempty"`
}

func (s DescribeAntChainBlockResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainBlockResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainBlockResponseBodyResult) SetPreviousHash(v string) *DescribeAntChainBlockResponseBodyResult {
	s.PreviousHash = &v
	return s
}

func (s *DescribeAntChainBlockResponseBodyResult) SetVersion(v int64) *DescribeAntChainBlockResponseBodyResult {
	s.Version = &v
	return s
}

func (s *DescribeAntChainBlockResponseBodyResult) SetTransactionSize(v int32) *DescribeAntChainBlockResponseBodyResult {
	s.TransactionSize = &v
	return s
}

func (s *DescribeAntChainBlockResponseBodyResult) SetCreateTime(v int64) *DescribeAntChainBlockResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainBlockResponseBodyResult) SetRootTxHash(v string) *DescribeAntChainBlockResponseBodyResult {
	s.RootTxHash = &v
	return s
}

func (s *DescribeAntChainBlockResponseBodyResult) SetHeight(v int32) *DescribeAntChainBlockResponseBodyResult {
	s.Height = &v
	return s
}

func (s *DescribeAntChainBlockResponseBodyResult) SetBlockHash(v string) *DescribeAntChainBlockResponseBodyResult {
	s.BlockHash = &v
	return s
}

func (s *DescribeAntChainBlockResponseBodyResult) SetAntChainId(v string) *DescribeAntChainBlockResponseBodyResult {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainBlockResponseBodyResult) SetTransSummaryList(v string) *DescribeAntChainBlockResponseBodyResult {
	s.TransSummaryList = &v
	return s
}

type DescribeAntChainBlockResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainBlockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainBlockResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainBlockResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainBlockResponse) SetHeaders(v map[string]*string) *DescribeAntChainBlockResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainBlockResponse) SetBody(v *DescribeAntChainBlockResponseBody) *DescribeAntChainBlockResponse {
	s.Body = v
	return s
}

type DescribeAntChainCertificateApplicationsRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeAntChainCertificateApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainCertificateApplicationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainCertificateApplicationsRequest) SetAntChainId(v string) *DescribeAntChainCertificateApplicationsRequest {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsRequest) SetStatus(v string) *DescribeAntChainCertificateApplicationsRequest {
	s.Status = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsRequest) SetPageSize(v int32) *DescribeAntChainCertificateApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsRequest) SetPageNumber(v int32) *DescribeAntChainCertificateApplicationsRequest {
	s.PageNumber = &v
	return s
}

type DescribeAntChainCertificateApplicationsResponseBody struct {
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainCertificateApplicationsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAntChainCertificateApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainCertificateApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainCertificateApplicationsResponseBody) SetRequestId(v string) *DescribeAntChainCertificateApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBody) SetResult(v *DescribeAntChainCertificateApplicationsResponseBodyResult) *DescribeAntChainCertificateApplicationsResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainCertificateApplicationsResponseBodyResult struct {
	Pagination              *DescribeAntChainCertificateApplicationsResponseBodyResultPagination                `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	CertificateApplications []*DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications `json:"CertificateApplications,omitempty" xml:"CertificateApplications,omitempty" type:"Repeated"`
}

func (s DescribeAntChainCertificateApplicationsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainCertificateApplicationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResult) SetPagination(v *DescribeAntChainCertificateApplicationsResponseBodyResultPagination) *DescribeAntChainCertificateApplicationsResponseBodyResult {
	s.Pagination = v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResult) SetCertificateApplications(v []*DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) *DescribeAntChainCertificateApplicationsResponseBodyResult {
	s.CertificateApplications = v
	return s
}

type DescribeAntChainCertificateApplicationsResponseBodyResultPagination struct {
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainCertificateApplicationsResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainCertificateApplicationsResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainCertificateApplicationsResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainCertificateApplicationsResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainCertificateApplicationsResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Updatetime *int64  `json:"Updatetime,omitempty" xml:"Updatetime,omitempty"`
	Createtime *int64  `json:"Createtime,omitempty" xml:"Createtime,omitempty"`
	Bid        *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) GoString() string {
	return s.String()
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) SetStatus(v string) *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications {
	s.Status = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) SetUpdatetime(v int64) *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications {
	s.Updatetime = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) SetCreatetime(v int64) *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications {
	s.Createtime = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) SetBid(v string) *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications {
	s.Bid = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) SetAntChainId(v string) *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) SetUsername(v string) *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications {
	s.Username = &v
	return s
}

type DescribeAntChainCertificateApplicationsResponse struct {
	Headers map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainCertificateApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainCertificateApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainCertificateApplicationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainCertificateApplicationsResponse) SetHeaders(v map[string]*string) *DescribeAntChainCertificateApplicationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponse) SetBody(v *DescribeAntChainCertificateApplicationsResponseBody) *DescribeAntChainCertificateApplicationsResponse {
	s.Body = v
	return s
}

type DescribeAntChainConsortiumsRequest struct {
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeAntChainConsortiumsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainConsortiumsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainConsortiumsRequest) SetPageSize(v int32) *DescribeAntChainConsortiumsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainConsortiumsRequest) SetPageNumber(v int32) *DescribeAntChainConsortiumsRequest {
	s.PageNumber = &v
	return s
}

type DescribeAntChainConsortiumsResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainConsortiumsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAntChainConsortiumsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainConsortiumsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainConsortiumsResponseBody) SetRequestId(v string) *DescribeAntChainConsortiumsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainConsortiumsResponseBody) SetResult(v *DescribeAntChainConsortiumsResponseBodyResult) *DescribeAntChainConsortiumsResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainConsortiumsResponseBodyResult struct {
	AntConsortiums []*DescribeAntChainConsortiumsResponseBodyResultAntConsortiums `json:"AntConsortiums,omitempty" xml:"AntConsortiums,omitempty" type:"Repeated"`
	Pagination     *DescribeAntChainConsortiumsResponseBodyResultPagination       `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s DescribeAntChainConsortiumsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainConsortiumsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainConsortiumsResponseBodyResult) SetAntConsortiums(v []*DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) *DescribeAntChainConsortiumsResponseBodyResult {
	s.AntConsortiums = v
	return s
}

func (s *DescribeAntChainConsortiumsResponseBodyResult) SetPagination(v *DescribeAntChainConsortiumsResponseBodyResultPagination) *DescribeAntChainConsortiumsResponseBodyResult {
	s.Pagination = v
	return s
}

type DescribeAntChainConsortiumsResponseBodyResultAntConsortiums struct {
	ConsortiumDescription *string `json:"ConsortiumDescription,omitempty" xml:"ConsortiumDescription,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ConsortiumId          *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MemberNum             *int64  `json:"MemberNum,omitempty" xml:"MemberNum,omitempty"`
	Role                  *string `json:"Role,omitempty" xml:"Role,omitempty"`
	ConsortiumName        *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	ChainNum              *int64  `json:"ChainNum,omitempty" xml:"ChainNum,omitempty"`
}

func (s DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) GoString() string {
	return s.String()
}

func (s *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) SetConsortiumDescription(v string) *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums {
	s.ConsortiumDescription = &v
	return s
}

func (s *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) SetStatus(v string) *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums {
	s.Status = &v
	return s
}

func (s *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) SetConsortiumId(v string) *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) SetCreateTime(v int64) *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) SetMemberNum(v int64) *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums {
	s.MemberNum = &v
	return s
}

func (s *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) SetRole(v string) *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums {
	s.Role = &v
	return s
}

func (s *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) SetConsortiumName(v string) *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums {
	s.ConsortiumName = &v
	return s
}

func (s *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) SetChainNum(v int64) *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums {
	s.ChainNum = &v
	return s
}

type DescribeAntChainConsortiumsResponseBodyResultPagination struct {
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainConsortiumsResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainConsortiumsResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainConsortiumsResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainConsortiumsResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainConsortiumsResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainConsortiumsResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainConsortiumsResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainConsortiumsResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainConsortiumsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainConsortiumsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainConsortiumsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainConsortiumsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainConsortiumsResponse) SetHeaders(v map[string]*string) *DescribeAntChainConsortiumsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainConsortiumsResponse) SetBody(v *DescribeAntChainConsortiumsResponseBody) *DescribeAntChainConsortiumsResponse {
	s.Body = v
	return s
}

type DescribeAntChainContractProjectContentTreeRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeAntChainContractProjectContentTreeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectContentTreeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectContentTreeRequest) SetProjectId(v string) *DescribeAntChainContractProjectContentTreeRequest {
	s.ProjectId = &v
	return s
}

type DescribeAntChainContractProjectContentTreeResponseBody struct {
	RequestId *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainContractProjectContentTreeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAntChainContractProjectContentTreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectContentTreeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectContentTreeResponseBody) SetRequestId(v string) *DescribeAntChainContractProjectContentTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeResponseBody) SetResult(v *DescribeAntChainContractProjectContentTreeResponseBodyResult) *DescribeAntChainContractProjectContentTreeResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainContractProjectContentTreeResponseBodyResult struct {
	Children           []map[string]interface{} `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	ProjectId          *string                  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName        *string                  `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectVersion     *string                  `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
	ProjectDescription *string                  `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
}

func (s DescribeAntChainContractProjectContentTreeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectContentTreeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectContentTreeResponseBodyResult) SetChildren(v []map[string]interface{}) *DescribeAntChainContractProjectContentTreeResponseBodyResult {
	s.Children = v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeResponseBodyResult) SetProjectId(v string) *DescribeAntChainContractProjectContentTreeResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeResponseBodyResult) SetProjectName(v string) *DescribeAntChainContractProjectContentTreeResponseBodyResult {
	s.ProjectName = &v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeResponseBodyResult) SetProjectVersion(v string) *DescribeAntChainContractProjectContentTreeResponseBodyResult {
	s.ProjectVersion = &v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeResponseBodyResult) SetProjectDescription(v string) *DescribeAntChainContractProjectContentTreeResponseBodyResult {
	s.ProjectDescription = &v
	return s
}

type DescribeAntChainContractProjectContentTreeResponse struct {
	Headers map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainContractProjectContentTreeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainContractProjectContentTreeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectContentTreeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectContentTreeResponse) SetHeaders(v map[string]*string) *DescribeAntChainContractProjectContentTreeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeResponse) SetBody(v *DescribeAntChainContractProjectContentTreeResponseBody) *DescribeAntChainContractProjectContentTreeResponse {
	s.Body = v
	return s
}

type DescribeAntChainContractProjectsRequest struct {
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeAntChainContractProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectsRequest) SetConsortiumId(v string) *DescribeAntChainContractProjectsRequest {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainContractProjectsRequest) SetPageSize(v int32) *DescribeAntChainContractProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainContractProjectsRequest) SetPageNumber(v int32) *DescribeAntChainContractProjectsRequest {
	s.PageNumber = &v
	return s
}

type DescribeAntChainContractProjectsResponseBody struct {
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainContractProjectsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAntChainContractProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectsResponseBody) SetRequestId(v string) *DescribeAntChainContractProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainContractProjectsResponseBody) SetResult(v *DescribeAntChainContractProjectsResponseBodyResult) *DescribeAntChainContractProjectsResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainContractProjectsResponseBodyResult struct {
	Pagination       *DescribeAntChainContractProjectsResponseBodyResultPagination         `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	ContractProjects []*DescribeAntChainContractProjectsResponseBodyResultContractProjects `json:"ContractProjects,omitempty" xml:"ContractProjects,omitempty" type:"Repeated"`
}

func (s DescribeAntChainContractProjectsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectsResponseBodyResult) SetPagination(v *DescribeAntChainContractProjectsResponseBodyResultPagination) *DescribeAntChainContractProjectsResponseBodyResult {
	s.Pagination = v
	return s
}

func (s *DescribeAntChainContractProjectsResponseBodyResult) SetContractProjects(v []*DescribeAntChainContractProjectsResponseBodyResultContractProjects) *DescribeAntChainContractProjectsResponseBodyResult {
	s.ContractProjects = v
	return s
}

type DescribeAntChainContractProjectsResponseBodyResultPagination struct {
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainContractProjectsResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectsResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectsResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainContractProjectsResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainContractProjectsResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainContractProjectsResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainContractProjectsResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainContractProjectsResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainContractProjectsResponseBodyResultContractProjects struct {
	UpdateTime         *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ConsortiumId       *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	CreateTime         *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ProjectId          *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectVersion     *string `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
}

func (s DescribeAntChainContractProjectsResponseBodyResultContractProjects) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectsResponseBodyResultContractProjects) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectsResponseBodyResultContractProjects) SetUpdateTime(v int64) *DescribeAntChainContractProjectsResponseBodyResultContractProjects {
	s.UpdateTime = &v
	return s
}

func (s *DescribeAntChainContractProjectsResponseBodyResultContractProjects) SetConsortiumId(v string) *DescribeAntChainContractProjectsResponseBodyResultContractProjects {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainContractProjectsResponseBodyResultContractProjects) SetCreateTime(v int64) *DescribeAntChainContractProjectsResponseBodyResultContractProjects {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainContractProjectsResponseBodyResultContractProjects) SetProjectId(v string) *DescribeAntChainContractProjectsResponseBodyResultContractProjects {
	s.ProjectId = &v
	return s
}

func (s *DescribeAntChainContractProjectsResponseBodyResultContractProjects) SetProjectName(v string) *DescribeAntChainContractProjectsResponseBodyResultContractProjects {
	s.ProjectName = &v
	return s
}

func (s *DescribeAntChainContractProjectsResponseBodyResultContractProjects) SetProjectVersion(v string) *DescribeAntChainContractProjectsResponseBodyResultContractProjects {
	s.ProjectVersion = &v
	return s
}

func (s *DescribeAntChainContractProjectsResponseBodyResultContractProjects) SetProjectDescription(v string) *DescribeAntChainContractProjectsResponseBodyResultContractProjects {
	s.ProjectDescription = &v
	return s
}

type DescribeAntChainContractProjectsResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainContractProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainContractProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectsResponse) SetHeaders(v map[string]*string) *DescribeAntChainContractProjectsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainContractProjectsResponse) SetBody(v *DescribeAntChainContractProjectsResponseBody) *DescribeAntChainContractProjectsResponse {
	s.Body = v
	return s
}

type DescribeAntChainDownloadPathsRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s DescribeAntChainDownloadPathsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainDownloadPathsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainDownloadPathsRequest) SetAntChainId(v string) *DescribeAntChainDownloadPathsRequest {
	s.AntChainId = &v
	return s
}

type DescribeAntChainDownloadPathsResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainDownloadPathsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAntChainDownloadPathsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainDownloadPathsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainDownloadPathsResponseBody) SetRequestId(v string) *DescribeAntChainDownloadPathsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainDownloadPathsResponseBody) SetResult(v *DescribeAntChainDownloadPathsResponseBodyResult) *DescribeAntChainDownloadPathsResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainDownloadPathsResponseBodyResult struct {
	CaCrtUrl     *string `json:"CaCrtUrl,omitempty" xml:"CaCrtUrl,omitempty"`
	SdkUrl       *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
	ClientCrtUrl *string `json:"ClientCrtUrl,omitempty" xml:"ClientCrtUrl,omitempty"`
	TrustCaUrl   *string `json:"TrustCaUrl,omitempty" xml:"TrustCaUrl,omitempty"`
}

func (s DescribeAntChainDownloadPathsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainDownloadPathsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainDownloadPathsResponseBodyResult) SetCaCrtUrl(v string) *DescribeAntChainDownloadPathsResponseBodyResult {
	s.CaCrtUrl = &v
	return s
}

func (s *DescribeAntChainDownloadPathsResponseBodyResult) SetSdkUrl(v string) *DescribeAntChainDownloadPathsResponseBodyResult {
	s.SdkUrl = &v
	return s
}

func (s *DescribeAntChainDownloadPathsResponseBodyResult) SetClientCrtUrl(v string) *DescribeAntChainDownloadPathsResponseBodyResult {
	s.ClientCrtUrl = &v
	return s
}

func (s *DescribeAntChainDownloadPathsResponseBodyResult) SetTrustCaUrl(v string) *DescribeAntChainDownloadPathsResponseBodyResult {
	s.TrustCaUrl = &v
	return s
}

type DescribeAntChainDownloadPathsResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainDownloadPathsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainDownloadPathsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainDownloadPathsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainDownloadPathsResponse) SetHeaders(v map[string]*string) *DescribeAntChainDownloadPathsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainDownloadPathsResponse) SetBody(v *DescribeAntChainDownloadPathsResponseBody) *DescribeAntChainDownloadPathsResponse {
	s.Body = v
	return s
}

type DescribeAntChainInformationRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s DescribeAntChainInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainInformationRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainInformationRequest) SetAntChainId(v string) *DescribeAntChainInformationRequest {
	s.AntChainId = &v
	return s
}

type DescribeAntChainInformationResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainInformationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAntChainInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainInformationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainInformationResponseBody) SetRequestId(v string) *DescribeAntChainInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainInformationResponseBody) SetResult(v *DescribeAntChainInformationResponseBodyResult) *DescribeAntChainInformationResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainInformationResponseBodyResult struct {
	TransactionSum *int32                                                    `json:"TransactionSum,omitempty" xml:"TransactionSum,omitempty"`
	Version        *string                                                   `json:"Version,omitempty" xml:"Version,omitempty"`
	BlockHeight    *int32                                                    `json:"BlockHeight,omitempty" xml:"BlockHeight,omitempty"`
	CreateTime     *int64                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	NodeNumber     *int32                                                    `json:"NodeNumber,omitempty" xml:"NodeNumber,omitempty"`
	AbnormalNodes  *int32                                                    `json:"AbnormalNodes,omitempty" xml:"AbnormalNodes,omitempty"`
	NodeInfos      []*DescribeAntChainInformationResponseBodyResultNodeInfos `json:"NodeInfos,omitempty" xml:"NodeInfos,omitempty" type:"Repeated"`
	AntChainId     *string                                                   `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	Normal         *bool                                                     `json:"Normal,omitempty" xml:"Normal,omitempty"`
}

func (s DescribeAntChainInformationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainInformationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainInformationResponseBodyResult) SetTransactionSum(v int32) *DescribeAntChainInformationResponseBodyResult {
	s.TransactionSum = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResult) SetVersion(v string) *DescribeAntChainInformationResponseBodyResult {
	s.Version = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResult) SetBlockHeight(v int32) *DescribeAntChainInformationResponseBodyResult {
	s.BlockHeight = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResult) SetCreateTime(v int64) *DescribeAntChainInformationResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResult) SetNodeNumber(v int32) *DescribeAntChainInformationResponseBodyResult {
	s.NodeNumber = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResult) SetAbnormalNodes(v int32) *DescribeAntChainInformationResponseBodyResult {
	s.AbnormalNodes = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResult) SetNodeInfos(v []*DescribeAntChainInformationResponseBodyResultNodeInfos) *DescribeAntChainInformationResponseBodyResult {
	s.NodeInfos = v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResult) SetAntChainId(v string) *DescribeAntChainInformationResponseBodyResult {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResult) SetNormal(v bool) *DescribeAntChainInformationResponseBodyResult {
	s.Normal = &v
	return s
}

type DescribeAntChainInformationResponseBodyResultNodeInfos struct {
	Status      *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	NodeName    *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	BlockHeight *int64  `json:"BlockHeight,omitempty" xml:"BlockHeight,omitempty"`
}

func (s DescribeAntChainInformationResponseBodyResultNodeInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainInformationResponseBodyResultNodeInfos) GoString() string {
	return s.String()
}

func (s *DescribeAntChainInformationResponseBodyResultNodeInfos) SetStatus(v bool) *DescribeAntChainInformationResponseBodyResultNodeInfos {
	s.Status = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResultNodeInfos) SetNodeName(v string) *DescribeAntChainInformationResponseBodyResultNodeInfos {
	s.NodeName = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResultNodeInfos) SetVersion(v string) *DescribeAntChainInformationResponseBodyResultNodeInfos {
	s.Version = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResultNodeInfos) SetBlockHeight(v int64) *DescribeAntChainInformationResponseBodyResultNodeInfos {
	s.BlockHeight = &v
	return s
}

type DescribeAntChainInformationResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainInformationResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainInformationResponse) SetHeaders(v map[string]*string) *DescribeAntChainInformationResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainInformationResponse) SetBody(v *DescribeAntChainInformationResponseBody) *DescribeAntChainInformationResponse {
	s.Body = v
	return s
}

type DescribeAntChainLatestBlocksRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s DescribeAntChainLatestBlocksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainLatestBlocksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainLatestBlocksRequest) SetAntChainId(v string) *DescribeAntChainLatestBlocksRequest {
	s.AntChainId = &v
	return s
}

type DescribeAntChainLatestBlocksResponseBody struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeAntChainLatestBlocksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainLatestBlocksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainLatestBlocksResponseBody) SetRequestId(v string) *DescribeAntChainLatestBlocksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainLatestBlocksResponseBody) SetResult(v []map[string]interface{}) *DescribeAntChainLatestBlocksResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainLatestBlocksResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainLatestBlocksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainLatestBlocksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainLatestBlocksResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainLatestBlocksResponse) SetHeaders(v map[string]*string) *DescribeAntChainLatestBlocksResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainLatestBlocksResponse) SetBody(v *DescribeAntChainLatestBlocksResponseBody) *DescribeAntChainLatestBlocksResponse {
	s.Body = v
	return s
}

type DescribeAntChainLatestTransactionDigestsRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s DescribeAntChainLatestTransactionDigestsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainLatestTransactionDigestsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainLatestTransactionDigestsRequest) SetAntChainId(v string) *DescribeAntChainLatestTransactionDigestsRequest {
	s.AntChainId = &v
	return s
}

type DescribeAntChainLatestTransactionDigestsResponseBody struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeAntChainLatestTransactionDigestsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainLatestTransactionDigestsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainLatestTransactionDigestsResponseBody) SetRequestId(v string) *DescribeAntChainLatestTransactionDigestsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainLatestTransactionDigestsResponseBody) SetResult(v []map[string]interface{}) *DescribeAntChainLatestTransactionDigestsResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainLatestTransactionDigestsResponse struct {
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainLatestTransactionDigestsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainLatestTransactionDigestsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainLatestTransactionDigestsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainLatestTransactionDigestsResponse) SetHeaders(v map[string]*string) *DescribeAntChainLatestTransactionDigestsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainLatestTransactionDigestsResponse) SetBody(v *DescribeAntChainLatestTransactionDigestsResponseBody) *DescribeAntChainLatestTransactionDigestsResponse {
	s.Body = v
	return s
}

type DescribeAntChainMembersRequest struct {
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
}

func (s DescribeAntChainMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMembersRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMembersRequest) SetPageSize(v int32) *DescribeAntChainMembersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainMembersRequest) SetPageNumber(v int32) *DescribeAntChainMembersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainMembersRequest) SetConsortiumId(v string) *DescribeAntChainMembersRequest {
	s.ConsortiumId = &v
	return s
}

type DescribeAntChainMembersResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainMembersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAntChainMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMembersResponseBody) SetRequestId(v string) *DescribeAntChainMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainMembersResponseBody) SetResult(v *DescribeAntChainMembersResponseBodyResult) *DescribeAntChainMembersResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainMembersResponseBodyResult struct {
	Members    []*DescribeAntChainMembersResponseBodyResultMembers  `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	Pagination *DescribeAntChainMembersResponseBodyResultPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s DescribeAntChainMembersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMembersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMembersResponseBodyResult) SetMembers(v []*DescribeAntChainMembersResponseBodyResultMembers) *DescribeAntChainMembersResponseBodyResult {
	s.Members = v
	return s
}

func (s *DescribeAntChainMembersResponseBodyResult) SetPagination(v *DescribeAntChainMembersResponseBodyResultPagination) *DescribeAntChainMembersResponseBodyResult {
	s.Pagination = v
	return s
}

type DescribeAntChainMembersResponseBodyResultMembers struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	JoinTime   *int64  `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	Role       *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeAntChainMembersResponseBodyResultMembers) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMembersResponseBodyResultMembers) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMembersResponseBodyResultMembers) SetStatus(v string) *DescribeAntChainMembersResponseBodyResultMembers {
	s.Status = &v
	return s
}

func (s *DescribeAntChainMembersResponseBodyResultMembers) SetMemberName(v string) *DescribeAntChainMembersResponseBodyResultMembers {
	s.MemberName = &v
	return s
}

func (s *DescribeAntChainMembersResponseBodyResultMembers) SetJoinTime(v int64) *DescribeAntChainMembersResponseBodyResultMembers {
	s.JoinTime = &v
	return s
}

func (s *DescribeAntChainMembersResponseBodyResultMembers) SetMemberId(v string) *DescribeAntChainMembersResponseBodyResultMembers {
	s.MemberId = &v
	return s
}

func (s *DescribeAntChainMembersResponseBodyResultMembers) SetRole(v string) *DescribeAntChainMembersResponseBodyResultMembers {
	s.Role = &v
	return s
}

type DescribeAntChainMembersResponseBodyResultPagination struct {
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainMembersResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMembersResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMembersResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainMembersResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainMembersResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainMembersResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainMembersResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainMembersResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainMembersResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMembersResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMembersResponse) SetHeaders(v map[string]*string) *DescribeAntChainMembersResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainMembersResponse) SetBody(v *DescribeAntChainMembersResponseBody) *DescribeAntChainMembersResponse {
	s.Body = v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAccessLogRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	QRCodeType *string `json:"QRCodeType,omitempty" xml:"QRCodeType,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogRequest) SetAntChainId(v string) *DescribeAntChainMiniAppBrowserQRCodeAccessLogRequest {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogRequest) SetQRCodeType(v string) *DescribeAntChainMiniAppBrowserQRCodeAccessLogRequest {
	s.QRCodeType = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBody struct {
	RequestId *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBody) SetRequestId(v string) *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBody) SetResult(v *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBodyResult) *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBodyResult struct {
	AccessCount              *int64 `json:"AccessCount,omitempty" xml:"AccessCount,omitempty"`
	AccessAlipayAccountCount *int64 `json:"AccessAlipayAccountCount,omitempty" xml:"AccessAlipayAccountCount,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBodyResult) SetAccessCount(v int64) *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBodyResult {
	s.AccessCount = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBodyResult) SetAccessAlipayAccountCount(v int64) *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBodyResult {
	s.AccessAlipayAccountCount = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse struct {
	Headers map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse) SetHeaders(v map[string]*string) *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse) SetBody(v *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBody) *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse {
	s.Body = v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	QRCodeType *string `json:"QRCodeType,omitempty" xml:"QRCodeType,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest) SetAntChainId(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest) SetQRCodeType(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest {
	s.QRCodeType = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest) SetPageSize(v int32) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest) SetPageNumber(v int32) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest {
	s.PageNumber = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBody struct {
	RequestId *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBody) SetRequestId(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBody) SetResult(v *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult struct {
	QRCodeType         *string                                                                                    `json:"QRCodeType,omitempty" xml:"QRCodeType,omitempty"`
	Pagination         *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultPagination           `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	AuthorizationType  *string                                                                                    `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	AuthorizedUserList []*DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultAuthorizedUserList `json:"AuthorizedUserList,omitempty" xml:"AuthorizedUserList,omitempty" type:"Repeated"`
	AntChainId         *string                                                                                    `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult) SetQRCodeType(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult {
	s.QRCodeType = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult) SetPagination(v *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultPagination) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult {
	s.Pagination = v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult) SetAuthorizationType(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult {
	s.AuthorizationType = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult) SetAuthorizedUserList(v []*DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultAuthorizedUserList) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult {
	s.AuthorizedUserList = v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult) SetAntChainId(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult {
	s.AntChainId = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultPagination struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultAuthorizedUserList struct {
	GmtAuthorized *string `json:"GmtAuthorized,omitempty" xml:"GmtAuthorized,omitempty"`
	Phone         *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultAuthorizedUserList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultAuthorizedUserList) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultAuthorizedUserList) SetGmtAuthorized(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultAuthorizedUserList {
	s.GmtAuthorized = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultAuthorizedUserList) SetPhone(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultAuthorizedUserList {
	s.Phone = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse struct {
	Headers map[string]*string                                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse) SetHeaders(v map[string]*string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse) SetBody(v *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBody) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse {
	s.Body = v
	return s
}

type DescribeAntChainMiniAppBrowserTransactionQRCodeRequest struct {
	AntChainId      *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	TransactionHash *string `json:"TransactionHash,omitempty" xml:"TransactionHash,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeRequest) SetAntChainId(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeRequest {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeRequest) SetTransactionHash(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeRequest {
	s.TransactionHash = &v
	return s
}

type DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBody struct {
	RequestId *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBody) SetRequestId(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBody) SetResult(v *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult) *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult struct {
	Base64QRCodePNG *string `json:"Base64QRCodePNG,omitempty" xml:"Base64QRCodePNG,omitempty"`
	TransactionHash *string `json:"TransactionHash,omitempty" xml:"TransactionHash,omitempty"`
	QRCodeContent   *string `json:"QRCodeContent,omitempty" xml:"QRCodeContent,omitempty"`
	AntChainId      *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult) SetBase64QRCodePNG(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult {
	s.Base64QRCodePNG = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult) SetTransactionHash(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult {
	s.TransactionHash = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult) SetQRCodeContent(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult {
	s.QRCodeContent = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult) SetAntChainId(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult {
	s.AntChainId = &v
	return s
}

type DescribeAntChainMiniAppBrowserTransactionQRCodeResponse struct {
	Headers map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeResponse) SetHeaders(v map[string]*string) *DescribeAntChainMiniAppBrowserTransactionQRCodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeResponse) SetBody(v *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBody) *DescribeAntChainMiniAppBrowserTransactionQRCodeResponse {
	s.Body = v
	return s
}

type DescribeAntChainNodesRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s DescribeAntChainNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainNodesRequest) SetAntChainId(v string) *DescribeAntChainNodesRequest {
	s.AntChainId = &v
	return s
}

type DescribeAntChainNodesResponseBody struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeAntChainNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainNodesResponseBody) SetRequestId(v string) *DescribeAntChainNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainNodesResponseBody) SetResult(v []map[string]interface{}) *DescribeAntChainNodesResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainNodesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainNodesResponse) SetHeaders(v map[string]*string) *DescribeAntChainNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainNodesResponse) SetBody(v *DescribeAntChainNodesResponseBody) *DescribeAntChainNodesResponse {
	s.Body = v
	return s
}

type DescribeAntChainQRCodeAuthorizationRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	QRCodeType *string `json:"QRCodeType,omitempty" xml:"QRCodeType,omitempty"`
}

func (s DescribeAntChainQRCodeAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainQRCodeAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainQRCodeAuthorizationRequest) SetAntChainId(v string) *DescribeAntChainQRCodeAuthorizationRequest {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationRequest) SetQRCodeType(v string) *DescribeAntChainQRCodeAuthorizationRequest {
	s.QRCodeType = &v
	return s
}

type DescribeAntChainQRCodeAuthorizationResponseBody struct {
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainQRCodeAuthorizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAntChainQRCodeAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainQRCodeAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainQRCodeAuthorizationResponseBody) SetRequestId(v string) *DescribeAntChainQRCodeAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationResponseBody) SetResult(v *DescribeAntChainQRCodeAuthorizationResponseBodyResult) *DescribeAntChainQRCodeAuthorizationResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainQRCodeAuthorizationResponseBodyResult struct {
	QRCodeType        *string `json:"QRCodeType,omitempty" xml:"QRCodeType,omitempty"`
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	AntChainId        *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s DescribeAntChainQRCodeAuthorizationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainQRCodeAuthorizationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainQRCodeAuthorizationResponseBodyResult) SetQRCodeType(v string) *DescribeAntChainQRCodeAuthorizationResponseBodyResult {
	s.QRCodeType = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationResponseBodyResult) SetAuthorizationType(v string) *DescribeAntChainQRCodeAuthorizationResponseBodyResult {
	s.AuthorizationType = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationResponseBodyResult) SetAntChainId(v string) *DescribeAntChainQRCodeAuthorizationResponseBodyResult {
	s.AntChainId = &v
	return s
}

type DescribeAntChainQRCodeAuthorizationResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainQRCodeAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainQRCodeAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainQRCodeAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainQRCodeAuthorizationResponse) SetHeaders(v map[string]*string) *DescribeAntChainQRCodeAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationResponse) SetBody(v *DescribeAntChainQRCodeAuthorizationResponseBody) *DescribeAntChainQRCodeAuthorizationResponse {
	s.Body = v
	return s
}

type DescribeAntChainsRequest struct {
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
}

func (s DescribeAntChainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainsRequest) SetPageSize(v int32) *DescribeAntChainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainsRequest) SetPageNumber(v int32) *DescribeAntChainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainsRequest) SetConsortiumId(v string) *DescribeAntChainsRequest {
	s.ConsortiumId = &v
	return s
}

type DescribeAntChainsResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAntChainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainsResponseBody) SetRequestId(v string) *DescribeAntChainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainsResponseBody) SetResult(v *DescribeAntChainsResponseBodyResult) *DescribeAntChainsResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainsResponseBodyResult struct {
	Pagination *DescribeAntChainsResponseBodyResultPagination  `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	AntChains  []*DescribeAntChainsResponseBodyResultAntChains `json:"AntChains,omitempty" xml:"AntChains,omitempty" type:"Repeated"`
	IsExist    *bool                                           `json:"IsExist,omitempty" xml:"IsExist,omitempty"`
}

func (s DescribeAntChainsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainsResponseBodyResult) SetPagination(v *DescribeAntChainsResponseBodyResultPagination) *DescribeAntChainsResponseBodyResult {
	s.Pagination = v
	return s
}

func (s *DescribeAntChainsResponseBodyResult) SetAntChains(v []*DescribeAntChainsResponseBodyResultAntChains) *DescribeAntChainsResponseBodyResult {
	s.AntChains = v
	return s
}

func (s *DescribeAntChainsResponseBodyResult) SetIsExist(v bool) *DescribeAntChainsResponseBodyResult {
	s.IsExist = &v
	return s
}

type DescribeAntChainsResponseBodyResultPagination struct {
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainsResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainsResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainsResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainsResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainsResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainsResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainsResponseBodyResultAntChains struct {
	ExpireTime     *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	CreateTime     *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChainType      *string `json:"ChainType,omitempty" xml:"ChainType,omitempty"`
	IsAdmin        *bool   `json:"IsAdmin,omitempty" xml:"IsAdmin,omitempty"`
	MerkleTreeSuit *string `json:"MerkleTreeSuit,omitempty" xml:"MerkleTreeSuit,omitempty"`
	MemberStatus   *string `json:"MemberStatus,omitempty" xml:"MemberStatus,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AntChainName   *string `json:"AntChainName,omitempty" xml:"AntChainName,omitempty"`
	Network        *string `json:"Network,omitempty" xml:"Network,omitempty"`
	TlsAlgo        *string `json:"TlsAlgo,omitempty" xml:"TlsAlgo,omitempty"`
	Version        *string `json:"Version,omitempty" xml:"Version,omitempty"`
	CipherSuit     *string `json:"CipherSuit,omitempty" xml:"CipherSuit,omitempty"`
	ResourceSize   *string `json:"ResourceSize,omitempty" xml:"ResourceSize,omitempty"`
	NodeNum        *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	AntChainId     *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s DescribeAntChainsResponseBodyResultAntChains) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainsResponseBodyResultAntChains) GoString() string {
	return s.String()
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetExpireTime(v int64) *DescribeAntChainsResponseBodyResultAntChains {
	s.ExpireTime = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetCreateTime(v int64) *DescribeAntChainsResponseBodyResultAntChains {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetChainType(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.ChainType = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetIsAdmin(v bool) *DescribeAntChainsResponseBodyResultAntChains {
	s.IsAdmin = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetMerkleTreeSuit(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.MerkleTreeSuit = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetMemberStatus(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.MemberStatus = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetRegionId(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.RegionId = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetAntChainName(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.AntChainName = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetNetwork(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.Network = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetTlsAlgo(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.TlsAlgo = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetVersion(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.Version = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetCipherSuit(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.CipherSuit = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetResourceSize(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.ResourceSize = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetNodeNum(v int32) *DescribeAntChainsResponseBodyResultAntChains {
	s.NodeNum = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetAntChainId(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.AntChainId = &v
	return s
}

type DescribeAntChainsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainsResponse) SetHeaders(v map[string]*string) *DescribeAntChainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainsResponse) SetBody(v *DescribeAntChainsResponseBody) *DescribeAntChainsResponse {
	s.Body = v
	return s
}

type DescribeAntChainTransactionRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	Hash       *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
}

func (s DescribeAntChainTransactionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionRequest) SetAntChainId(v string) *DescribeAntChainTransactionRequest {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainTransactionRequest) SetHash(v string) *DescribeAntChainTransactionRequest {
	s.Hash = &v
	return s
}

type DescribeAntChainTransactionResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainTransactionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAntChainTransactionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionResponseBody) SetRequestId(v string) *DescribeAntChainTransactionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBody) SetResult(v *DescribeAntChainTransactionResponseBodyResult) *DescribeAntChainTransactionResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainTransactionResponseBodyResult struct {
	Hash         *string                                                   `json:"Hash,omitempty" xml:"Hash,omitempty"`
	BlockVersion *string                                                   `json:"BlockVersion,omitempty" xml:"BlockVersion,omitempty"`
	BlockHeight  *int64                                                    `json:"BlockHeight,omitempty" xml:"BlockHeight,omitempty"`
	CreateTime   *int64                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	BlockHash    *string                                                   `json:"BlockHash,omitempty" xml:"BlockHash,omitempty"`
	Transaction  *DescribeAntChainTransactionResponseBodyResultTransaction `json:"Transaction,omitempty" xml:"Transaction,omitempty" type:"Struct"`
}

func (s DescribeAntChainTransactionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionResponseBodyResult) SetHash(v string) *DescribeAntChainTransactionResponseBodyResult {
	s.Hash = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResult) SetBlockVersion(v string) *DescribeAntChainTransactionResponseBodyResult {
	s.BlockVersion = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResult) SetBlockHeight(v int64) *DescribeAntChainTransactionResponseBodyResult {
	s.BlockHeight = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResult) SetCreateTime(v int64) *DescribeAntChainTransactionResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResult) SetBlockHash(v string) *DescribeAntChainTransactionResponseBodyResult {
	s.BlockHash = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResult) SetTransaction(v *DescribeAntChainTransactionResponseBodyResultTransaction) *DescribeAntChainTransactionResponseBodyResult {
	s.Transaction = v
	return s
}

type DescribeAntChainTransactionResponseBodyResultTransaction struct {
	Hash       *string   `json:"Hash,omitempty" xml:"Hash,omitempty"`
	From       *string   `json:"From,omitempty" xml:"From,omitempty"`
	Data       *string   `json:"Data,omitempty" xml:"Data,omitempty"`
	Nonce      *string   `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
	To         *string   `json:"To,omitempty" xml:"To,omitempty"`
	Timestamp  *int64    `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Gas        *string   `json:"Gas,omitempty" xml:"Gas,omitempty"`
	Period     *int64    `json:"Period,omitempty" xml:"Period,omitempty"`
	Value      *string   `json:"Value,omitempty" xml:"Value,omitempty"`
	Extentions []*string `json:"Extentions,omitempty" xml:"Extentions,omitempty" type:"Repeated"`
	TxType     *string   `json:"TxType,omitempty" xml:"TxType,omitempty"`
	Signatures []*string `json:"Signatures,omitempty" xml:"Signatures,omitempty" type:"Repeated"`
}

func (s DescribeAntChainTransactionResponseBodyResultTransaction) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionResponseBodyResultTransaction) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetHash(v string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Hash = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetFrom(v string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.From = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetData(v string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Data = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetNonce(v string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Nonce = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetTo(v string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.To = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetTimestamp(v int64) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Timestamp = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetGas(v string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Gas = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetPeriod(v int64) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Period = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetValue(v string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Value = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetExtentions(v []*string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Extentions = v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetTxType(v string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.TxType = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetSignatures(v []*string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Signatures = v
	return s
}

type DescribeAntChainTransactionResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainTransactionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainTransactionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionResponse) SetHeaders(v map[string]*string) *DescribeAntChainTransactionResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainTransactionResponse) SetBody(v *DescribeAntChainTransactionResponseBody) *DescribeAntChainTransactionResponse {
	s.Body = v
	return s
}

type DescribeAntChainTransactionReceiptRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	Hash       *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
}

func (s DescribeAntChainTransactionReceiptRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionReceiptRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionReceiptRequest) SetAntChainId(v string) *DescribeAntChainTransactionReceiptRequest {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptRequest) SetHash(v string) *DescribeAntChainTransactionReceiptRequest {
	s.Hash = &v
	return s
}

type DescribeAntChainTransactionReceiptResponseBody struct {
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainTransactionReceiptResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAntChainTransactionReceiptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionReceiptResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionReceiptResponseBody) SetRequestId(v string) *DescribeAntChainTransactionReceiptResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptResponseBody) SetResult(v *DescribeAntChainTransactionReceiptResponseBodyResult) *DescribeAntChainTransactionReceiptResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainTransactionReceiptResponseBodyResult struct {
	Data    *string   `json:"Data,omitempty" xml:"Data,omitempty"`
	Result  *int64    `json:"Result,omitempty" xml:"Result,omitempty"`
	Logs    []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	GasUsed *string   `json:"GasUsed,omitempty" xml:"GasUsed,omitempty"`
}

func (s DescribeAntChainTransactionReceiptResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionReceiptResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionReceiptResponseBodyResult) SetData(v string) *DescribeAntChainTransactionReceiptResponseBodyResult {
	s.Data = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptResponseBodyResult) SetResult(v int64) *DescribeAntChainTransactionReceiptResponseBodyResult {
	s.Result = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptResponseBodyResult) SetLogs(v []*string) *DescribeAntChainTransactionReceiptResponseBodyResult {
	s.Logs = v
	return s
}

func (s *DescribeAntChainTransactionReceiptResponseBodyResult) SetGasUsed(v string) *DescribeAntChainTransactionReceiptResponseBodyResult {
	s.GasUsed = &v
	return s
}

type DescribeAntChainTransactionReceiptResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainTransactionReceiptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainTransactionReceiptResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionReceiptResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionReceiptResponse) SetHeaders(v map[string]*string) *DescribeAntChainTransactionReceiptResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainTransactionReceiptResponse) SetBody(v *DescribeAntChainTransactionReceiptResponseBody) *DescribeAntChainTransactionReceiptResponse {
	s.Body = v
	return s
}

type DescribeAntChainTransactionStatisticsRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	Start      *int64  `json:"Start,omitempty" xml:"Start,omitempty"`
	End        *int64  `json:"End,omitempty" xml:"End,omitempty"`
}

func (s DescribeAntChainTransactionStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionStatisticsRequest) SetAntChainId(v string) *DescribeAntChainTransactionStatisticsRequest {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsRequest) SetStart(v int64) *DescribeAntChainTransactionStatisticsRequest {
	s.Start = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsRequest) SetEnd(v int64) *DescribeAntChainTransactionStatisticsRequest {
	s.End = &v
	return s
}

type DescribeAntChainTransactionStatisticsResponseBody struct {
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeAntChainTransactionStatisticsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeAntChainTransactionStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionStatisticsResponseBody) SetRequestId(v string) *DescribeAntChainTransactionStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsResponseBody) SetResult(v []*DescribeAntChainTransactionStatisticsResponseBodyResult) *DescribeAntChainTransactionStatisticsResponseBody {
	s.Result = v
	return s
}

type DescribeAntChainTransactionStatisticsResponseBodyResult struct {
	CreatTime          *int64  `json:"CreatTime,omitempty" xml:"CreatTime,omitempty"`
	TransCount         *int64  `json:"TransCount,omitempty" xml:"TransCount,omitempty"`
	Dt                 *string `json:"Dt,omitempty" xml:"Dt,omitempty"`
	LastSumBlockHeight *int64  `json:"LastSumBlockHeight,omitempty" xml:"LastSumBlockHeight,omitempty"`
	AntChainId         *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s DescribeAntChainTransactionStatisticsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionStatisticsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionStatisticsResponseBodyResult) SetCreatTime(v int64) *DescribeAntChainTransactionStatisticsResponseBodyResult {
	s.CreatTime = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsResponseBodyResult) SetTransCount(v int64) *DescribeAntChainTransactionStatisticsResponseBodyResult {
	s.TransCount = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsResponseBodyResult) SetDt(v string) *DescribeAntChainTransactionStatisticsResponseBodyResult {
	s.Dt = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsResponseBodyResult) SetLastSumBlockHeight(v int64) *DescribeAntChainTransactionStatisticsResponseBodyResult {
	s.LastSumBlockHeight = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsResponseBodyResult) SetAntChainId(v string) *DescribeAntChainTransactionStatisticsResponseBodyResult {
	s.AntChainId = &v
	return s
}

type DescribeAntChainTransactionStatisticsResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntChainTransactionStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntChainTransactionStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionStatisticsResponse) SetHeaders(v map[string]*string) *DescribeAntChainTransactionStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainTransactionStatisticsResponse) SetBody(v *DescribeAntChainTransactionStatisticsResponseBody) *DescribeAntChainTransactionStatisticsResponse {
	s.Body = v
	return s
}

type DescribeEthereumDeletableRequest struct {
	EthereumId *string `json:"EthereumId,omitempty" xml:"EthereumId,omitempty"`
}

func (s DescribeEthereumDeletableRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEthereumDeletableRequest) GoString() string {
	return s.String()
}

func (s *DescribeEthereumDeletableRequest) SetEthereumId(v string) *DescribeEthereumDeletableRequest {
	s.EthereumId = &v
	return s
}

type DescribeEthereumDeletableResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *DescribeEthereumDeletableResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeEthereumDeletableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEthereumDeletableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEthereumDeletableResponseBody) SetRequestId(v string) *DescribeEthereumDeletableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEthereumDeletableResponseBody) SetErrorCode(v int32) *DescribeEthereumDeletableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeEthereumDeletableResponseBody) SetSuccess(v bool) *DescribeEthereumDeletableResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEthereumDeletableResponseBody) SetResult(v *DescribeEthereumDeletableResponseBodyResult) *DescribeEthereumDeletableResponseBody {
	s.Result = v
	return s
}

type DescribeEthereumDeletableResponseBodyResult struct {
	Deletable  *bool   `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
	EthereumId *string `json:"EthereumId,omitempty" xml:"EthereumId,omitempty"`
}

func (s DescribeEthereumDeletableResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeEthereumDeletableResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeEthereumDeletableResponseBodyResult) SetDeletable(v bool) *DescribeEthereumDeletableResponseBodyResult {
	s.Deletable = &v
	return s
}

func (s *DescribeEthereumDeletableResponseBodyResult) SetEthereumId(v string) *DescribeEthereumDeletableResponseBodyResult {
	s.EthereumId = &v
	return s
}

type DescribeEthereumDeletableResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEthereumDeletableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEthereumDeletableResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEthereumDeletableResponse) GoString() string {
	return s.String()
}

func (s *DescribeEthereumDeletableResponse) SetHeaders(v map[string]*string) *DescribeEthereumDeletableResponse {
	s.Headers = v
	return s
}

func (s *DescribeEthereumDeletableResponse) SetBody(v *DescribeEthereumDeletableResponseBody) *DescribeEthereumDeletableResponse {
	s.Body = v
	return s
}

type DescribeFabricCandidateOrganizationsRequest struct {
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s DescribeFabricCandidateOrganizationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricCandidateOrganizationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricCandidateOrganizationsRequest) SetLocation(v string) *DescribeFabricCandidateOrganizationsRequest {
	s.Location = &v
	return s
}

type DescribeFabricCandidateOrganizationsResponseBody struct {
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DescribeFabricCandidateOrganizationsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeFabricCandidateOrganizationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricCandidateOrganizationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricCandidateOrganizationsResponseBody) SetRequestId(v string) *DescribeFabricCandidateOrganizationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricCandidateOrganizationsResponseBody) SetErrorCode(v int32) *DescribeFabricCandidateOrganizationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricCandidateOrganizationsResponseBody) SetSuccess(v bool) *DescribeFabricCandidateOrganizationsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricCandidateOrganizationsResponseBody) SetResult(v []*DescribeFabricCandidateOrganizationsResponseBodyResult) *DescribeFabricCandidateOrganizationsResponseBody {
	s.Result = v
	return s
}

type DescribeFabricCandidateOrganizationsResponseBodyResult struct {
	ServiceState     *string `json:"ServiceState,omitempty" xml:"ServiceState,omitempty"`
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	ClusterState     *string `json:"ClusterState,omitempty" xml:"ClusterState,omitempty"`
	OrganizationId   *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DescribeFabricCandidateOrganizationsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricCandidateOrganizationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricCandidateOrganizationsResponseBodyResult) SetServiceState(v string) *DescribeFabricCandidateOrganizationsResponseBodyResult {
	s.ServiceState = &v
	return s
}

func (s *DescribeFabricCandidateOrganizationsResponseBodyResult) SetOrganizationName(v string) *DescribeFabricCandidateOrganizationsResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *DescribeFabricCandidateOrganizationsResponseBodyResult) SetClusterState(v string) *DescribeFabricCandidateOrganizationsResponseBodyResult {
	s.ClusterState = &v
	return s
}

func (s *DescribeFabricCandidateOrganizationsResponseBodyResult) SetOrganizationId(v string) *DescribeFabricCandidateOrganizationsResponseBodyResult {
	s.OrganizationId = &v
	return s
}

type DescribeFabricCandidateOrganizationsResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricCandidateOrganizationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricCandidateOrganizationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricCandidateOrganizationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricCandidateOrganizationsResponse) SetHeaders(v map[string]*string) *DescribeFabricCandidateOrganizationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricCandidateOrganizationsResponse) SetBody(v *DescribeFabricCandidateOrganizationsResponseBody) *DescribeFabricCandidateOrganizationsResponse {
	s.Body = v
	return s
}

type DescribeFabricChaincodeUploadPolicyRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DescribeFabricChaincodeUploadPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChaincodeUploadPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricChaincodeUploadPolicyRequest) SetOrganizationId(v string) *DescribeFabricChaincodeUploadPolicyRequest {
	s.OrganizationId = &v
	return s
}

type DescribeFabricChaincodeUploadPolicyResponseBody struct {
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *DescribeFabricChaincodeUploadPolicyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeFabricChaincodeUploadPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChaincodeUploadPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBody) SetRequestId(v string) *DescribeFabricChaincodeUploadPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBody) SetErrorCode(v int32) *DescribeFabricChaincodeUploadPolicyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBody) SetSuccess(v bool) *DescribeFabricChaincodeUploadPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBody) SetResult(v *DescribeFabricChaincodeUploadPolicyResponseBodyResult) *DescribeFabricChaincodeUploadPolicyResponseBody {
	s.Result = v
	return s
}

type DescribeFabricChaincodeUploadPolicyResponseBodyResult struct {
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Expire    *int32  `json:"Expire,omitempty" xml:"Expire,omitempty"`
}

func (s DescribeFabricChaincodeUploadPolicyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChaincodeUploadPolicyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBodyResult) SetSignature(v string) *DescribeFabricChaincodeUploadPolicyResponseBodyResult {
	s.Signature = &v
	return s
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBodyResult) SetHost(v string) *DescribeFabricChaincodeUploadPolicyResponseBodyResult {
	s.Host = &v
	return s
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBodyResult) SetPolicy(v string) *DescribeFabricChaincodeUploadPolicyResponseBodyResult {
	s.Policy = &v
	return s
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBodyResult) SetDir(v string) *DescribeFabricChaincodeUploadPolicyResponseBodyResult {
	s.Dir = &v
	return s
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBodyResult) SetAccessId(v string) *DescribeFabricChaincodeUploadPolicyResponseBodyResult {
	s.AccessId = &v
	return s
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBodyResult) SetExpire(v int32) *DescribeFabricChaincodeUploadPolicyResponseBodyResult {
	s.Expire = &v
	return s
}

type DescribeFabricChaincodeUploadPolicyResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricChaincodeUploadPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricChaincodeUploadPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChaincodeUploadPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricChaincodeUploadPolicyResponse) SetHeaders(v map[string]*string) *DescribeFabricChaincodeUploadPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricChaincodeUploadPolicyResponse) SetBody(v *DescribeFabricChaincodeUploadPolicyResponseBody) *DescribeFabricChaincodeUploadPolicyResponse {
	s.Body = v
	return s
}

type DescribeFabricChannelMembersRequest struct {
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s DescribeFabricChannelMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChannelMembersRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricChannelMembersRequest) SetChannelId(v string) *DescribeFabricChannelMembersRequest {
	s.ChannelId = &v
	return s
}

type DescribeFabricChannelMembersResponseBody struct {
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DescribeFabricChannelMembersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeFabricChannelMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChannelMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricChannelMembersResponseBody) SetRequestId(v string) *DescribeFabricChannelMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBody) SetErrorCode(v int32) *DescribeFabricChannelMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBody) SetSuccess(v bool) *DescribeFabricChannelMembersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBody) SetResult(v []*DescribeFabricChannelMembersResponseBodyResult) *DescribeFabricChannelMembersResponseBody {
	s.Result = v
	return s
}

type DescribeFabricChannelMembersResponseBodyResult struct {
	WithPeer                *bool   `json:"WithPeer,omitempty" xml:"WithPeer,omitempty"`
	AcceptTime              *string `json:"AcceptTime,omitempty" xml:"AcceptTime,omitempty"`
	OrganizationDomain      *string `json:"OrganizationDomain,omitempty" xml:"OrganizationDomain,omitempty"`
	State                   *string `json:"State,omitempty" xml:"State,omitempty"`
	InviteTime              *string `json:"InviteTime,omitempty" xml:"InviteTime,omitempty"`
	ChannelId               *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OrganizationName        *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	OrganizationDescription *string `json:"OrganizationDescription,omitempty" xml:"OrganizationDescription,omitempty"`
	OrganizationId          *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DescribeFabricChannelMembersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChannelMembersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetWithPeer(v bool) *DescribeFabricChannelMembersResponseBodyResult {
	s.WithPeer = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetAcceptTime(v string) *DescribeFabricChannelMembersResponseBodyResult {
	s.AcceptTime = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetOrganizationDomain(v string) *DescribeFabricChannelMembersResponseBodyResult {
	s.OrganizationDomain = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetState(v string) *DescribeFabricChannelMembersResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetInviteTime(v string) *DescribeFabricChannelMembersResponseBodyResult {
	s.InviteTime = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetChannelId(v string) *DescribeFabricChannelMembersResponseBodyResult {
	s.ChannelId = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetOrganizationName(v string) *DescribeFabricChannelMembersResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetOrganizationDescription(v string) *DescribeFabricChannelMembersResponseBodyResult {
	s.OrganizationDescription = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetOrganizationId(v string) *DescribeFabricChannelMembersResponseBodyResult {
	s.OrganizationId = &v
	return s
}

type DescribeFabricChannelMembersResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricChannelMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricChannelMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChannelMembersResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricChannelMembersResponse) SetHeaders(v map[string]*string) *DescribeFabricChannelMembersResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricChannelMembersResponse) SetBody(v *DescribeFabricChannelMembersResponseBody) *DescribeFabricChannelMembersResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumAdminStatusRequest struct {
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s DescribeFabricConsortiumAdminStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumAdminStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumAdminStatusRequest) SetLocation(v string) *DescribeFabricConsortiumAdminStatusRequest {
	s.Location = &v
	return s
}

type DescribeFabricConsortiumAdminStatusResponseBody struct {
	RequestId *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DescribeFabricConsortiumAdminStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeFabricConsortiumAdminStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumAdminStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumAdminStatusResponseBody) SetRequestId(v string) *DescribeFabricConsortiumAdminStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumAdminStatusResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumAdminStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumAdminStatusResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumAdminStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricConsortiumAdminStatusResponseBody) SetResult(v []*DescribeFabricConsortiumAdminStatusResponseBodyResult) *DescribeFabricConsortiumAdminStatusResponseBody {
	s.Result = v
	return s
}

type DescribeFabricConsortiumAdminStatusResponseBodyResult struct {
	ConsortiumId            *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ConsortiumAdministrator *bool   `json:"ConsortiumAdministrator,omitempty" xml:"ConsortiumAdministrator,omitempty"`
}

func (s DescribeFabricConsortiumAdminStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumAdminStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumAdminStatusResponseBodyResult) SetConsortiumId(v string) *DescribeFabricConsortiumAdminStatusResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumAdminStatusResponseBodyResult) SetConsortiumAdministrator(v bool) *DescribeFabricConsortiumAdminStatusResponseBodyResult {
	s.ConsortiumAdministrator = &v
	return s
}

type DescribeFabricConsortiumAdminStatusResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricConsortiumAdminStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricConsortiumAdminStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumAdminStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumAdminStatusResponse) SetHeaders(v map[string]*string) *DescribeFabricConsortiumAdminStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricConsortiumAdminStatusResponse) SetBody(v *DescribeFabricConsortiumAdminStatusResponseBody) *DescribeFabricConsortiumAdminStatusResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumChaincodesRequest struct {
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s DescribeFabricConsortiumChaincodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumChaincodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumChaincodesRequest) SetConsortiumId(v string) *DescribeFabricConsortiumChaincodesRequest {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesRequest) SetLocation(v string) *DescribeFabricConsortiumChaincodesRequest {
	s.Location = &v
	return s
}

type DescribeFabricConsortiumChaincodesResponseBody struct {
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DescribeFabricConsortiumChaincodesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeFabricConsortiumChaincodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumChaincodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumChaincodesResponseBody) SetRequestId(v string) *DescribeFabricConsortiumChaincodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumChaincodesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumChaincodesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBody) SetResult(v []*DescribeFabricConsortiumChaincodesResponseBodyResult) *DescribeFabricConsortiumChaincodesResponseBody {
	s.Result = v
	return s
}

type DescribeFabricConsortiumChaincodesResponseBodyResult struct {
	Type             *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	EndorsePolicy    *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChaincodeId      *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ProviderName     *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	ChaincodeName    *string `json:"ChaincodeName,omitempty" xml:"ChaincodeName,omitempty"`
	Input            *string `json:"Input,omitempty" xml:"Input,omitempty"`
	Install          *bool   `json:"Install,omitempty" xml:"Install,omitempty"`
	ProviderId       *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	DeployTime       *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" xml:"ChaincodeVersion,omitempty"`
	ConsortiumId     *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ChannelName      *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	ChannelId        *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Path             *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DescribeFabricConsortiumChaincodesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumChaincodesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetType(v int32) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.Type = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetEndorsePolicy(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.EndorsePolicy = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetState(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetCreateTime(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetChaincodeId(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.ChaincodeId = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetMessage(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.Message = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetProviderName(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.ProviderName = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetChaincodeName(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.ChaincodeName = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetInput(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.Input = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetInstall(v bool) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.Install = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetProviderId(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.ProviderId = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetDeployTime(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.DeployTime = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetChaincodeVersion(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.ChaincodeVersion = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetConsortiumId(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetChannelName(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetChannelId(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.ChannelId = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetPath(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.Path = &v
	return s
}

type DescribeFabricConsortiumChaincodesResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricConsortiumChaincodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricConsortiumChaincodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumChaincodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumChaincodesResponse) SetHeaders(v map[string]*string) *DescribeFabricConsortiumChaincodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponse) SetBody(v *DescribeFabricConsortiumChaincodesResponseBody) *DescribeFabricConsortiumChaincodesResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumChannelsRequest struct {
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s DescribeFabricConsortiumChannelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumChannelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumChannelsRequest) SetConsortiumId(v string) *DescribeFabricConsortiumChannelsRequest {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsRequest) SetLocation(v string) *DescribeFabricConsortiumChannelsRequest {
	s.Location = &v
	return s
}

type DescribeFabricConsortiumChannelsResponseBody struct {
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DescribeFabricConsortiumChannelsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeFabricConsortiumChannelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumChannelsResponseBody) SetRequestId(v string) *DescribeFabricConsortiumChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumChannelsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumChannelsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBody) SetResult(v []*DescribeFabricConsortiumChannelsResponseBodyResult) *DescribeFabricConsortiumChannelsResponseBody {
	s.Result = v
	return s
}

type DescribeFabricConsortiumChannelsResponseBodyResult struct {
	BatchTimeout         *int32  `json:"BatchTimeout,omitempty" xml:"BatchTimeout,omitempty"`
	UpdateTime           *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ChaincodeCount       *int32  `json:"ChaincodeCount,omitempty" xml:"ChaincodeCount,omitempty"`
	State                *string `json:"State,omitempty" xml:"State,omitempty"`
	MemberJoinedCount    *string `json:"MemberJoinedCount,omitempty" xml:"MemberJoinedCount,omitempty"`
	PreferredMaxBytes    *int32  `json:"PreferredMaxBytes,omitempty" xml:"PreferredMaxBytes,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SupportChannelConfig *bool   `json:"SupportChannelConfig,omitempty" xml:"SupportChannelConfig,omitempty"`
	OwnerName            *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerUid             *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	OwnerBid             *string `json:"OwnerBid,omitempty" xml:"OwnerBid,omitempty"`
	MaxMessageCount      *int32  `json:"MaxMessageCount,omitempty" xml:"MaxMessageCount,omitempty"`
	MemberCount          *int32  `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	NeedJoined           *bool   `json:"NeedJoined,omitempty" xml:"NeedJoined,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConsortiumId         *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ChannelName          *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	DeleteTime           *string `json:"DeleteTime,omitempty" xml:"DeleteTime,omitempty"`
	ChannelId            *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ConsortiumChannelId  *int32  `json:"ConsortiumChannelId,omitempty" xml:"ConsortiumChannelId,omitempty"`
	Deleted              *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ConsortiumName       *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	BlockCount           *int32  `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
}

func (s DescribeFabricConsortiumChannelsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumChannelsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetBatchTimeout(v int32) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.BatchTimeout = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetUpdateTime(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetChaincodeCount(v int32) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.ChaincodeCount = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetState(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetMemberJoinedCount(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.MemberJoinedCount = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetPreferredMaxBytes(v int32) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.PreferredMaxBytes = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetCreateTime(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetSupportChannelConfig(v bool) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.SupportChannelConfig = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetOwnerName(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetOwnerUid(v int64) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.OwnerUid = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetOwnerBid(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.OwnerBid = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetMaxMessageCount(v int32) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.MaxMessageCount = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetMemberCount(v int32) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.MemberCount = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetNeedJoined(v bool) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.NeedJoined = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetRequestId(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetConsortiumId(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetChannelName(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetDeleteTime(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.DeleteTime = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetChannelId(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.ChannelId = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetConsortiumChannelId(v int32) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.ConsortiumChannelId = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetDeleted(v bool) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.Deleted = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetConsortiumName(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.ConsortiumName = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetBlockCount(v int32) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.BlockCount = &v
	return s
}

type DescribeFabricConsortiumChannelsResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricConsortiumChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricConsortiumChannelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumChannelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumChannelsResponse) SetHeaders(v map[string]*string) *DescribeFabricConsortiumChannelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponse) SetBody(v *DescribeFabricConsortiumChannelsResponseBody) *DescribeFabricConsortiumChannelsResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumConfigResponseBody struct {
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *DescribeFabricConsortiumConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeFabricConsortiumConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumConfigResponseBody) SetRequestId(v string) *DescribeFabricConsortiumConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumConfigResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumConfigResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricConsortiumConfigResponseBody) SetResult(v *DescribeFabricConsortiumConfigResponseBodyResult) *DescribeFabricConsortiumConfigResponseBody {
	s.Result = v
	return s
}

type DescribeFabricConsortiumConfigResponseBodyResult struct {
	ChannelPolicy []*string `json:"ChannelPolicy,omitempty" xml:"ChannelPolicy,omitempty" type:"Repeated"`
	OrdererType   []*string `json:"OrdererType,omitempty" xml:"OrdererType,omitempty" type:"Repeated"`
}

func (s DescribeFabricConsortiumConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumConfigResponseBodyResult) SetChannelPolicy(v []*string) *DescribeFabricConsortiumConfigResponseBodyResult {
	s.ChannelPolicy = v
	return s
}

func (s *DescribeFabricConsortiumConfigResponseBodyResult) SetOrdererType(v []*string) *DescribeFabricConsortiumConfigResponseBodyResult {
	s.OrdererType = v
	return s
}

type DescribeFabricConsortiumConfigResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricConsortiumConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricConsortiumConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumConfigResponse) SetHeaders(v map[string]*string) *DescribeFabricConsortiumConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricConsortiumConfigResponse) SetBody(v *DescribeFabricConsortiumConfigResponseBody) *DescribeFabricConsortiumConfigResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumDeletableRequest struct {
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s DescribeFabricConsortiumDeletableRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumDeletableRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumDeletableRequest) SetConsortiumId(v string) *DescribeFabricConsortiumDeletableRequest {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableRequest) SetLocation(v string) *DescribeFabricConsortiumDeletableRequest {
	s.Location = &v
	return s
}

type DescribeFabricConsortiumDeletableResponseBody struct {
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *DescribeFabricConsortiumDeletableResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeFabricConsortiumDeletableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumDeletableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumDeletableResponseBody) SetRequestId(v string) *DescribeFabricConsortiumDeletableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumDeletableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumDeletableResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBody) SetResult(v *DescribeFabricConsortiumDeletableResponseBodyResult) *DescribeFabricConsortiumDeletableResponseBody {
	s.Result = v
	return s
}

type DescribeFabricConsortiumDeletableResponseBodyResult struct {
	Deletable      *bool   `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ConsortiumId   *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	State          *string `json:"State,omitempty" xml:"State,omitempty"`
	ZoneId         *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	CodeName       *string `json:"CodeName,omitempty" xml:"CodeName,omitempty"`
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFabricConsortiumDeletableResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumDeletableResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumDeletableResponseBodyResult) SetDeletable(v bool) *DescribeFabricConsortiumDeletableResponseBodyResult {
	s.Deletable = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBodyResult) SetDomain(v string) *DescribeFabricConsortiumDeletableResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBodyResult) SetDescription(v string) *DescribeFabricConsortiumDeletableResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBodyResult) SetConsortiumId(v string) *DescribeFabricConsortiumDeletableResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBodyResult) SetState(v string) *DescribeFabricConsortiumDeletableResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBodyResult) SetZoneId(v string) *DescribeFabricConsortiumDeletableResponseBodyResult {
	s.ZoneId = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBodyResult) SetCodeName(v string) *DescribeFabricConsortiumDeletableResponseBodyResult {
	s.CodeName = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBodyResult) SetConsortiumName(v string) *DescribeFabricConsortiumDeletableResponseBodyResult {
	s.ConsortiumName = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBodyResult) SetRegionId(v string) *DescribeFabricConsortiumDeletableResponseBodyResult {
	s.RegionId = &v
	return s
}

type DescribeFabricConsortiumDeletableResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricConsortiumDeletableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricConsortiumDeletableResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumDeletableResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumDeletableResponse) SetHeaders(v map[string]*string) *DescribeFabricConsortiumDeletableResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponse) SetBody(v *DescribeFabricConsortiumDeletableResponseBody) *DescribeFabricConsortiumDeletableResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumMemberApprovalRequest struct {
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s DescribeFabricConsortiumMemberApprovalRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumMemberApprovalRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumMemberApprovalRequest) SetConsortiumId(v string) *DescribeFabricConsortiumMemberApprovalRequest {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalRequest) SetLocation(v string) *DescribeFabricConsortiumMemberApprovalRequest {
	s.Location = &v
	return s
}

type DescribeFabricConsortiumMemberApprovalResponseBody struct {
	RequestId *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DescribeFabricConsortiumMemberApprovalResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeFabricConsortiumMemberApprovalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumMemberApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBody) SetRequestId(v string) *DescribeFabricConsortiumMemberApprovalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumMemberApprovalResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumMemberApprovalResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBody) SetResult(v []*DescribeFabricConsortiumMemberApprovalResponseBodyResult) *DescribeFabricConsortiumMemberApprovalResponseBody {
	s.Result = v
	return s
}

type DescribeFabricConsortiumMemberApprovalResponseBodyResult struct {
	ConsortiumId        *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	State               *string `json:"State,omitempty" xml:"State,omitempty"`
	ChannelCreatePolicy *string `json:"ChannelCreatePolicy,omitempty" xml:"ChannelCreatePolicy,omitempty"`
	ConfirmTime         *string `json:"ConfirmTime,omitempty" xml:"ConfirmTime,omitempty"`
	OrganizationName    *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	ConsortiumName      *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	DomainName          *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OrganizationId      *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DescribeFabricConsortiumMemberApprovalResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumMemberApprovalResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBodyResult) SetConsortiumId(v string) *DescribeFabricConsortiumMemberApprovalResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBodyResult) SetState(v string) *DescribeFabricConsortiumMemberApprovalResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBodyResult) SetChannelCreatePolicy(v string) *DescribeFabricConsortiumMemberApprovalResponseBodyResult {
	s.ChannelCreatePolicy = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBodyResult) SetConfirmTime(v string) *DescribeFabricConsortiumMemberApprovalResponseBodyResult {
	s.ConfirmTime = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBodyResult) SetOrganizationName(v string) *DescribeFabricConsortiumMemberApprovalResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBodyResult) SetConsortiumName(v string) *DescribeFabricConsortiumMemberApprovalResponseBodyResult {
	s.ConsortiumName = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBodyResult) SetDomainName(v string) *DescribeFabricConsortiumMemberApprovalResponseBodyResult {
	s.DomainName = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBodyResult) SetOrganizationId(v string) *DescribeFabricConsortiumMemberApprovalResponseBodyResult {
	s.OrganizationId = &v
	return s
}

type DescribeFabricConsortiumMemberApprovalResponse struct {
	Headers map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricConsortiumMemberApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricConsortiumMemberApprovalResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumMemberApprovalResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumMemberApprovalResponse) SetHeaders(v map[string]*string) *DescribeFabricConsortiumMemberApprovalResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponse) SetBody(v *DescribeFabricConsortiumMemberApprovalResponseBody) *DescribeFabricConsortiumMemberApprovalResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumMembersRequest struct {
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s DescribeFabricConsortiumMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumMembersRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumMembersRequest) SetConsortiumId(v string) *DescribeFabricConsortiumMembersRequest {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumMembersRequest) SetLocation(v string) *DescribeFabricConsortiumMembersRequest {
	s.Location = &v
	return s
}

type DescribeFabricConsortiumMembersResponseBody struct {
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DescribeFabricConsortiumMembersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeFabricConsortiumMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumMembersResponseBody) SetRequestId(v string) *DescribeFabricConsortiumMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumMembersResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumMembersResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumMembersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricConsortiumMembersResponseBody) SetResult(v []*DescribeFabricConsortiumMembersResponseBodyResult) *DescribeFabricConsortiumMembersResponseBody {
	s.Result = v
	return s
}

type DescribeFabricConsortiumMembersResponseBodyResult struct {
	Domain           *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ConsortiumId     *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	JoinedTime       *string `json:"JoinedTime,omitempty" xml:"JoinedTime,omitempty"`
	OrganizationId   *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DescribeFabricConsortiumMembersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumMembersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumMembersResponseBodyResult) SetDomain(v string) *DescribeFabricConsortiumMembersResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricConsortiumMembersResponseBodyResult) SetDescription(v string) *DescribeFabricConsortiumMembersResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeFabricConsortiumMembersResponseBodyResult) SetConsortiumId(v string) *DescribeFabricConsortiumMembersResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumMembersResponseBodyResult) SetOrganizationName(v string) *DescribeFabricConsortiumMembersResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *DescribeFabricConsortiumMembersResponseBodyResult) SetJoinedTime(v string) *DescribeFabricConsortiumMembersResponseBodyResult {
	s.JoinedTime = &v
	return s
}

func (s *DescribeFabricConsortiumMembersResponseBodyResult) SetOrganizationId(v string) *DescribeFabricConsortiumMembersResponseBodyResult {
	s.OrganizationId = &v
	return s
}

type DescribeFabricConsortiumMembersResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricConsortiumMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricConsortiumMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumMembersResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumMembersResponse) SetHeaders(v map[string]*string) *DescribeFabricConsortiumMembersResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricConsortiumMembersResponse) SetBody(v *DescribeFabricConsortiumMembersResponseBody) *DescribeFabricConsortiumMembersResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumOrderersRequest struct {
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s DescribeFabricConsortiumOrderersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumOrderersRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumOrderersRequest) SetConsortiumId(v string) *DescribeFabricConsortiumOrderersRequest {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersRequest) SetLocation(v string) *DescribeFabricConsortiumOrderersRequest {
	s.Location = &v
	return s
}

type DescribeFabricConsortiumOrderersResponseBody struct {
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DescribeFabricConsortiumOrderersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeFabricConsortiumOrderersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumOrderersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumOrderersResponseBody) SetRequestId(v string) *DescribeFabricConsortiumOrderersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumOrderersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumOrderersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponseBody) SetResult(v []*DescribeFabricConsortiumOrderersResponseBodyResult) *DescribeFabricConsortiumOrderersResponseBody {
	s.Result = v
	return s
}

type DescribeFabricConsortiumOrderersResponseBodyResult struct {
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	UpdateTime   *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	OrdererName  *string `json:"OrdererName,omitempty" xml:"OrdererName,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Port         *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeFabricConsortiumOrderersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumOrderersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumOrderersResponseBodyResult) SetDomain(v string) *DescribeFabricConsortiumOrderersResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponseBodyResult) SetUpdateTime(v string) *DescribeFabricConsortiumOrderersResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponseBodyResult) SetOrdererName(v string) *DescribeFabricConsortiumOrderersResponseBodyResult {
	s.OrdererName = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponseBodyResult) SetCreateTime(v string) *DescribeFabricConsortiumOrderersResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponseBodyResult) SetPort(v int32) *DescribeFabricConsortiumOrderersResponseBodyResult {
	s.Port = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponseBodyResult) SetInstanceType(v string) *DescribeFabricConsortiumOrderersResponseBodyResult {
	s.InstanceType = &v
	return s
}

type DescribeFabricConsortiumOrderersResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricConsortiumOrderersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricConsortiumOrderersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumOrderersResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumOrderersResponse) SetHeaders(v map[string]*string) *DescribeFabricConsortiumOrderersResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponse) SetBody(v *DescribeFabricConsortiumOrderersResponseBody) *DescribeFabricConsortiumOrderersResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumsRequest struct {
	ConsortiumId *string                                `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	Location     *string                                `json:"Location,omitempty" xml:"Location,omitempty"`
	Tag          []*DescribeFabricConsortiumsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeFabricConsortiumsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumsRequest) SetConsortiumId(v string) *DescribeFabricConsortiumsRequest {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumsRequest) SetLocation(v string) *DescribeFabricConsortiumsRequest {
	s.Location = &v
	return s
}

func (s *DescribeFabricConsortiumsRequest) SetTag(v []*DescribeFabricConsortiumsRequestTag) *DescribeFabricConsortiumsRequest {
	s.Tag = v
	return s
}

type DescribeFabricConsortiumsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFabricConsortiumsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumsRequestTag) SetKey(v string) *DescribeFabricConsortiumsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeFabricConsortiumsRequestTag) SetValue(v string) *DescribeFabricConsortiumsRequestTag {
	s.Value = &v
	return s
}

type DescribeFabricConsortiumsResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DescribeFabricConsortiumsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeFabricConsortiumsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumsResponseBody) SetRequestId(v string) *DescribeFabricConsortiumsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBody) SetResult(v []*DescribeFabricConsortiumsResponseBodyResult) *DescribeFabricConsortiumsResponseBody {
	s.Result = v
	return s
}

type DescribeFabricConsortiumsResponseBodyResult struct {
	ChannelCount         *int32                                             `json:"ChannelCount,omitempty" xml:"ChannelCount,omitempty"`
	Domain               *string                                            `json:"Domain,omitempty" xml:"Domain,omitempty"`
	State                *string                                            `json:"State,omitempty" xml:"State,omitempty"`
	CreateTime           *string                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Tags                 []*DescribeFabricConsortiumsResponseBodyResultTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	SpecName             *string                                            `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	SupportChannelConfig *bool                                              `json:"SupportChannelConfig,omitempty" xml:"SupportChannelConfig,omitempty"`
	OwnerName            *string                                            `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerUid             *int64                                             `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	CodeName             *string                                            `json:"CodeName,omitempty" xml:"CodeName,omitempty"`
	OwnerBid             *string                                            `json:"OwnerBid,omitempty" xml:"OwnerBid,omitempty"`
	RegionId             *string                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ChannelPolicy        *string                                            `json:"ChannelPolicy,omitempty" xml:"ChannelPolicy,omitempty"`
	RequestId            *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConsortiumId         *string                                            `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ExpiredTime          *string                                            `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	OrganizationCount    *int32                                             `json:"OrganizationCount,omitempty" xml:"OrganizationCount,omitempty"`
	ConsortiumName       *string                                            `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
}

func (s DescribeFabricConsortiumsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetChannelCount(v int32) *DescribeFabricConsortiumsResponseBodyResult {
	s.ChannelCount = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetDomain(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetState(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetCreateTime(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetTags(v []*DescribeFabricConsortiumsResponseBodyResultTags) *DescribeFabricConsortiumsResponseBodyResult {
	s.Tags = v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetSpecName(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.SpecName = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetSupportChannelConfig(v bool) *DescribeFabricConsortiumsResponseBodyResult {
	s.SupportChannelConfig = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetOwnerName(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetOwnerUid(v int64) *DescribeFabricConsortiumsResponseBodyResult {
	s.OwnerUid = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetCodeName(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.CodeName = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetOwnerBid(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.OwnerBid = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetRegionId(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetChannelPolicy(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.ChannelPolicy = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetRequestId(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetConsortiumId(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetExpiredTime(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetOrganizationCount(v int32) *DescribeFabricConsortiumsResponseBodyResult {
	s.OrganizationCount = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetConsortiumName(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.ConsortiumName = &v
	return s
}

type DescribeFabricConsortiumsResponseBodyResultTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFabricConsortiumsResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumsResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumsResponseBodyResultTags) SetKey(v string) *DescribeFabricConsortiumsResponseBodyResultTags {
	s.Key = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResultTags) SetValue(v string) *DescribeFabricConsortiumsResponseBodyResultTags {
	s.Value = &v
	return s
}

type DescribeFabricConsortiumsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricConsortiumsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricConsortiumsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumsResponse) SetHeaders(v map[string]*string) *DescribeFabricConsortiumsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricConsortiumsResponse) SetBody(v *DescribeFabricConsortiumsResponseBody) *DescribeFabricConsortiumsResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumSpecsResponseBody struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DescribeFabricConsortiumSpecsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeFabricConsortiumSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumSpecsResponseBody) SetRequestId(v string) *DescribeFabricConsortiumSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumSpecsResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumSpecsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumSpecsResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumSpecsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricConsortiumSpecsResponseBody) SetResult(v []*DescribeFabricConsortiumSpecsResponseBodyResult) *DescribeFabricConsortiumSpecsResponseBody {
	s.Result = v
	return s
}

type DescribeFabricConsortiumSpecsResponseBodyResult struct {
	SpecName  *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	SpecTitle *string `json:"SpecTitle,omitempty" xml:"SpecTitle,omitempty"`
	Enable    *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeFabricConsortiumSpecsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumSpecsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumSpecsResponseBodyResult) SetSpecName(v string) *DescribeFabricConsortiumSpecsResponseBodyResult {
	s.SpecName = &v
	return s
}

func (s *DescribeFabricConsortiumSpecsResponseBodyResult) SetSpecTitle(v string) *DescribeFabricConsortiumSpecsResponseBodyResult {
	s.SpecTitle = &v
	return s
}

func (s *DescribeFabricConsortiumSpecsResponseBodyResult) SetEnable(v bool) *DescribeFabricConsortiumSpecsResponseBodyResult {
	s.Enable = &v
	return s
}

type DescribeFabricConsortiumSpecsResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricConsortiumSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricConsortiumSpecsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumSpecsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumSpecsResponse) SetHeaders(v map[string]*string) *DescribeFabricConsortiumSpecsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricConsortiumSpecsResponse) SetBody(v *DescribeFabricConsortiumSpecsResponseBody) *DescribeFabricConsortiumSpecsResponse {
	s.Body = v
	return s
}

type DescribeFabricExplorerRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	ExMethod       *string `json:"ExMethod,omitempty" xml:"ExMethod,omitempty"`
	ExUrl          *string `json:"ExUrl,omitempty" xml:"ExUrl,omitempty"`
	ExBody         *string `json:"ExBody,omitempty" xml:"ExBody,omitempty"`
}

func (s DescribeFabricExplorerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricExplorerRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricExplorerRequest) SetOrganizationId(v string) *DescribeFabricExplorerRequest {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricExplorerRequest) SetExMethod(v string) *DescribeFabricExplorerRequest {
	s.ExMethod = &v
	return s
}

func (s *DescribeFabricExplorerRequest) SetExUrl(v string) *DescribeFabricExplorerRequest {
	s.ExUrl = &v
	return s
}

func (s *DescribeFabricExplorerRequest) SetExBody(v string) *DescribeFabricExplorerRequest {
	s.ExBody = &v
	return s
}

type DescribeFabricExplorerResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	ErrorCode      *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Result         *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeFabricExplorerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricExplorerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricExplorerResponseBody) SetRequestId(v string) *DescribeFabricExplorerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricExplorerResponseBody) SetDynamicCode(v string) *DescribeFabricExplorerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeFabricExplorerResponseBody) SetErrorCode(v int32) *DescribeFabricExplorerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricExplorerResponseBody) SetDynamicMessage(v string) *DescribeFabricExplorerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeFabricExplorerResponseBody) SetSuccess(v bool) *DescribeFabricExplorerResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricExplorerResponseBody) SetResult(v string) *DescribeFabricExplorerResponseBody {
	s.Result = &v
	return s
}

type DescribeFabricExplorerResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricExplorerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricExplorerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricExplorerResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricExplorerResponse) SetHeaders(v map[string]*string) *DescribeFabricExplorerResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricExplorerResponse) SetBody(v *DescribeFabricExplorerResponseBody) *DescribeFabricExplorerResponse {
	s.Body = v
	return s
}

type DescribeFabricInvitationCodeRequest struct {
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
}

func (s DescribeFabricInvitationCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricInvitationCodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricInvitationCodeRequest) SetConsortiumId(v string) *DescribeFabricInvitationCodeRequest {
	s.ConsortiumId = &v
	return s
}

type DescribeFabricInvitationCodeResponseBody struct {
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DynamicCode    *string                                         `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                         `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *int32                                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success        *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Result         *DescribeFabricInvitationCodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeFabricInvitationCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricInvitationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricInvitationCodeResponseBody) SetRequestId(v string) *DescribeFabricInvitationCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBody) SetDynamicCode(v string) *DescribeFabricInvitationCodeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBody) SetDynamicMessage(v string) *DescribeFabricInvitationCodeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBody) SetErrorCode(v int32) *DescribeFabricInvitationCodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBody) SetSuccess(v bool) *DescribeFabricInvitationCodeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBody) SetResult(v *DescribeFabricInvitationCodeResponseBodyResult) *DescribeFabricInvitationCodeResponseBody {
	s.Result = v
	return s
}

type DescribeFabricInvitationCodeResponseBodyResult struct {
	SenderId     *int64  `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	SenderBid    *string `json:"SenderBid,omitempty" xml:"SenderBid,omitempty"`
	ExpireTime   *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	InvitationId *int32  `json:"InvitationId,omitempty" xml:"InvitationId,omitempty"`
	SenderName   *string `json:"SenderName,omitempty" xml:"SenderName,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Accepted     *bool   `json:"Accepted,omitempty" xml:"Accepted,omitempty"`
	SendTime     *string `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
}

func (s DescribeFabricInvitationCodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricInvitationCodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetSenderId(v int64) *DescribeFabricInvitationCodeResponseBodyResult {
	s.SenderId = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetEmail(v string) *DescribeFabricInvitationCodeResponseBodyResult {
	s.Email = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetSenderBid(v string) *DescribeFabricInvitationCodeResponseBodyResult {
	s.SenderBid = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetExpireTime(v string) *DescribeFabricInvitationCodeResponseBodyResult {
	s.ExpireTime = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetConsortiumId(v string) *DescribeFabricInvitationCodeResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetInvitationId(v int32) *DescribeFabricInvitationCodeResponseBodyResult {
	s.InvitationId = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetSenderName(v string) *DescribeFabricInvitationCodeResponseBodyResult {
	s.SenderName = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetCode(v string) *DescribeFabricInvitationCodeResponseBodyResult {
	s.Code = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetUrl(v string) *DescribeFabricInvitationCodeResponseBodyResult {
	s.Url = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetAccepted(v bool) *DescribeFabricInvitationCodeResponseBodyResult {
	s.Accepted = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetSendTime(v string) *DescribeFabricInvitationCodeResponseBodyResult {
	s.SendTime = &v
	return s
}

type DescribeFabricInvitationCodeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricInvitationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricInvitationCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricInvitationCodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricInvitationCodeResponse) SetHeaders(v map[string]*string) *DescribeFabricInvitationCodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricInvitationCodeResponse) SetBody(v *DescribeFabricInvitationCodeResponseBody) *DescribeFabricInvitationCodeResponse {
	s.Body = v
	return s
}

type DescribeFabricInviterRequest struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeFabricInviterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricInviterRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricInviterRequest) SetCode(v string) *DescribeFabricInviterRequest {
	s.Code = &v
	return s
}

type DescribeFabricInviterResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *DescribeFabricInviterResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeFabricInviterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricInviterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricInviterResponseBody) SetRequestId(v string) *DescribeFabricInviterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricInviterResponseBody) SetErrorCode(v int32) *DescribeFabricInviterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricInviterResponseBody) SetSuccess(v bool) *DescribeFabricInviterResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricInviterResponseBody) SetResult(v *DescribeFabricInviterResponseBodyResult) *DescribeFabricInviterResponseBody {
	s.Result = v
	return s
}

type DescribeFabricInviterResponseBodyResult struct {
	InviterId      *int64  `json:"InviterId,omitempty" xml:"InviterId,omitempty"`
	ExpireTime     *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ConsortiumId   *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	InviterName    *string `json:"InviterName,omitempty" xml:"InviterName,omitempty"`
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
}

func (s DescribeFabricInviterResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricInviterResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricInviterResponseBodyResult) SetInviterId(v int64) *DescribeFabricInviterResponseBodyResult {
	s.InviterId = &v
	return s
}

func (s *DescribeFabricInviterResponseBodyResult) SetExpireTime(v string) *DescribeFabricInviterResponseBodyResult {
	s.ExpireTime = &v
	return s
}

func (s *DescribeFabricInviterResponseBodyResult) SetConsortiumId(v string) *DescribeFabricInviterResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricInviterResponseBodyResult) SetInviterName(v string) *DescribeFabricInviterResponseBodyResult {
	s.InviterName = &v
	return s
}

func (s *DescribeFabricInviterResponseBodyResult) SetConsortiumName(v string) *DescribeFabricInviterResponseBodyResult {
	s.ConsortiumName = &v
	return s
}

type DescribeFabricInviterResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricInviterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricInviterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricInviterResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricInviterResponse) SetHeaders(v map[string]*string) *DescribeFabricInviterResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricInviterResponse) SetBody(v *DescribeFabricInviterResponseBody) *DescribeFabricInviterResponse {
	s.Body = v
	return s
}

type DescribeFabricOrdererLogsRequest struct {
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	OrdererName  *string `json:"OrdererName,omitempty" xml:"OrdererName,omitempty"`
	Lines        *string `json:"Lines,omitempty" xml:"Lines,omitempty"`
}

func (s DescribeFabricOrdererLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrdererLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrdererLogsRequest) SetConsortiumId(v string) *DescribeFabricOrdererLogsRequest {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricOrdererLogsRequest) SetOrdererName(v string) *DescribeFabricOrdererLogsRequest {
	s.OrdererName = &v
	return s
}

func (s *DescribeFabricOrdererLogsRequest) SetLines(v string) *DescribeFabricOrdererLogsRequest {
	s.Lines = &v
	return s
}

type DescribeFabricOrdererLogsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeFabricOrdererLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrdererLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrdererLogsResponseBody) SetRequestId(v string) *DescribeFabricOrdererLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrdererLogsResponseBody) SetErrorCode(v int32) *DescribeFabricOrdererLogsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrdererLogsResponseBody) SetSuccess(v bool) *DescribeFabricOrdererLogsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricOrdererLogsResponseBody) SetResult(v string) *DescribeFabricOrdererLogsResponseBody {
	s.Result = &v
	return s
}

type DescribeFabricOrdererLogsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricOrdererLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricOrdererLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrdererLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrdererLogsResponse) SetHeaders(v map[string]*string) *DescribeFabricOrdererLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricOrdererLogsResponse) SetBody(v *DescribeFabricOrdererLogsResponseBody) *DescribeFabricOrdererLogsResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationRequest struct {
	OrganizationId *string                                 `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Location       *string                                 `json:"Location,omitempty" xml:"Location,omitempty"`
	Tag            []*DescribeFabricOrganizationRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeFabricOrganizationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationRequest) SetOrganizationId(v string) *DescribeFabricOrganizationRequest {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricOrganizationRequest) SetLocation(v string) *DescribeFabricOrganizationRequest {
	s.Location = &v
	return s
}

func (s *DescribeFabricOrganizationRequest) SetTag(v []*DescribeFabricOrganizationRequestTag) *DescribeFabricOrganizationRequest {
	s.Tag = v
	return s
}

type DescribeFabricOrganizationRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFabricOrganizationRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationRequestTag) SetKey(v string) *DescribeFabricOrganizationRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeFabricOrganizationRequestTag) SetValue(v string) *DescribeFabricOrganizationRequestTag {
	s.Value = &v
	return s
}

type DescribeFabricOrganizationResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *DescribeFabricOrganizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeFabricOrganizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationResponseBody) SetRequestId(v string) *DescribeFabricOrganizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBody) SetResult(v *DescribeFabricOrganizationResponseBodyResult) *DescribeFabricOrganizationResponseBody {
	s.Result = v
	return s
}

type DescribeFabricOrganizationResponseBodyResult struct {
	Domain                  *string                                             `json:"Domain,omitempty" xml:"Domain,omitempty"`
	PeerCount               *int32                                              `json:"PeerCount,omitempty" xml:"PeerCount,omitempty"`
	State                   *string                                             `json:"State,omitempty" xml:"State,omitempty"`
	CreateTime              *string                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ConsortiumCount         *int32                                              `json:"ConsortiumCount,omitempty" xml:"ConsortiumCount,omitempty"`
	Tags                    []*DescribeFabricOrganizationResponseBodyResultTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	SpecName                *string                                             `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	OwnerName               *string                                             `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerUid                *int64                                              `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	CodeName                *string                                             `json:"CodeName,omitempty" xml:"CodeName,omitempty"`
	OwnerBid                *string                                             `json:"OwnerBid,omitempty" xml:"OwnerBid,omitempty"`
	OrganizationDescription *string                                             `json:"OrganizationDescription,omitempty" xml:"OrganizationDescription,omitempty"`
	RegionId                *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OrganizationId          *string                                             `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	MSP                     *string                                             `json:"MSP,omitempty" xml:"MSP,omitempty"`
	RequestId               *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CAUrl                   *string                                             `json:"CAUrl,omitempty" xml:"CAUrl,omitempty"`
	CANAME                  *string                                             `json:"CANAME,omitempty" xml:"CANAME,omitempty"`
	ZoneId                  *string                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	UserCount               *int32                                              `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	OrganizationName        *string                                             `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
}

func (s DescribeFabricOrganizationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetDomain(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetPeerCount(v int32) *DescribeFabricOrganizationResponseBodyResult {
	s.PeerCount = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetState(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetCreateTime(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetConsortiumCount(v int32) *DescribeFabricOrganizationResponseBodyResult {
	s.ConsortiumCount = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetTags(v []*DescribeFabricOrganizationResponseBodyResultTags) *DescribeFabricOrganizationResponseBodyResult {
	s.Tags = v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetSpecName(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.SpecName = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetOwnerName(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetOwnerUid(v int64) *DescribeFabricOrganizationResponseBodyResult {
	s.OwnerUid = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetCodeName(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.CodeName = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetOwnerBid(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.OwnerBid = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetOrganizationDescription(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.OrganizationDescription = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetRegionId(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetOrganizationId(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetMSP(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.MSP = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetRequestId(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetCAUrl(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.CAUrl = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetCANAME(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.CANAME = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetZoneId(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.ZoneId = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetUserCount(v int32) *DescribeFabricOrganizationResponseBodyResult {
	s.UserCount = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetOrganizationName(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.OrganizationName = &v
	return s
}

type DescribeFabricOrganizationResponseBodyResultTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFabricOrganizationResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationResponseBodyResultTags) SetKey(v string) *DescribeFabricOrganizationResponseBodyResultTags {
	s.Key = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResultTags) SetValue(v string) *DescribeFabricOrganizationResponseBodyResultTags {
	s.Value = &v
	return s
}

type DescribeFabricOrganizationResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricOrganizationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationResponse) SetHeaders(v map[string]*string) *DescribeFabricOrganizationResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricOrganizationResponse) SetBody(v *DescribeFabricOrganizationResponseBody) *DescribeFabricOrganizationResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationChaincodesRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Location       *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s DescribeFabricOrganizationChaincodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationChaincodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationChaincodesRequest) SetOrganizationId(v string) *DescribeFabricOrganizationChaincodesRequest {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesRequest) SetLocation(v string) *DescribeFabricOrganizationChaincodesRequest {
	s.Location = &v
	return s
}

type DescribeFabricOrganizationChaincodesResponseBody struct {
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DescribeFabricOrganizationChaincodesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeFabricOrganizationChaincodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationChaincodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationChaincodesResponseBody) SetRequestId(v string) *DescribeFabricOrganizationChaincodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationChaincodesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationChaincodesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBody) SetResult(v []*DescribeFabricOrganizationChaincodesResponseBodyResult) *DescribeFabricOrganizationChaincodesResponseBody {
	s.Result = v
	return s
}

type DescribeFabricOrganizationChaincodesResponseBodyResult struct {
	EndorsePolicy    *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChaincodeId      *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ChaincodeName    *string `json:"ChaincodeName,omitempty" xml:"ChaincodeName,omitempty"`
	Installed        *string `json:"Installed,omitempty" xml:"Installed,omitempty"`
	Creator          *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DeployTime       *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" xml:"ChaincodeVersion,omitempty"`
	ChannelName      *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	ConsortiumId     *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ChannelId        *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s DescribeFabricOrganizationChaincodesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationChaincodesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetEndorsePolicy(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.EndorsePolicy = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetState(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetCreateTime(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetChaincodeId(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.ChaincodeId = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetMessage(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.Message = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetChaincodeName(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.ChaincodeName = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetInstalled(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.Installed = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetCreator(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.Creator = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetDeployTime(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.DeployTime = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetChaincodeVersion(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.ChaincodeVersion = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetChannelName(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetConsortiumId(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetChannelId(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.ChannelId = &v
	return s
}

type DescribeFabricOrganizationChaincodesResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricOrganizationChaincodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricOrganizationChaincodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationChaincodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationChaincodesResponse) SetHeaders(v map[string]*string) *DescribeFabricOrganizationChaincodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponse) SetBody(v *DescribeFabricOrganizationChaincodesResponseBody) *DescribeFabricOrganizationChaincodesResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationDeletableRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Location       *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s DescribeFabricOrganizationDeletableRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationDeletableRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationDeletableRequest) SetOrganizationId(v string) *DescribeFabricOrganizationDeletableRequest {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableRequest) SetLocation(v string) *DescribeFabricOrganizationDeletableRequest {
	s.Location = &v
	return s
}

type DescribeFabricOrganizationDeletableResponseBody struct {
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *DescribeFabricOrganizationDeletableResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeFabricOrganizationDeletableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationDeletableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationDeletableResponseBody) SetRequestId(v string) *DescribeFabricOrganizationDeletableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationDeletableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationDeletableResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBody) SetResult(v *DescribeFabricOrganizationDeletableResponseBodyResult) *DescribeFabricOrganizationDeletableResponseBody {
	s.Result = v
	return s
}

type DescribeFabricOrganizationDeletableResponseBodyResult struct {
	Deletable               *bool   `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
	Domain                  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	State                   *string `json:"State,omitempty" xml:"State,omitempty"`
	ZoneId                  *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	OrganizationName        *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	CodeName                *string `json:"CodeName,omitempty" xml:"CodeName,omitempty"`
	OrganizationDescription *string `json:"OrganizationDescription,omitempty" xml:"OrganizationDescription,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OrganizationId          *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DescribeFabricOrganizationDeletableResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationDeletableResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationDeletableResponseBodyResult) SetDeletable(v bool) *DescribeFabricOrganizationDeletableResponseBodyResult {
	s.Deletable = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBodyResult) SetDomain(v string) *DescribeFabricOrganizationDeletableResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBodyResult) SetState(v string) *DescribeFabricOrganizationDeletableResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBodyResult) SetZoneId(v string) *DescribeFabricOrganizationDeletableResponseBodyResult {
	s.ZoneId = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBodyResult) SetOrganizationName(v string) *DescribeFabricOrganizationDeletableResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBodyResult) SetCodeName(v string) *DescribeFabricOrganizationDeletableResponseBodyResult {
	s.CodeName = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBodyResult) SetOrganizationDescription(v string) *DescribeFabricOrganizationDeletableResponseBodyResult {
	s.OrganizationDescription = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBodyResult) SetRegionId(v string) *DescribeFabricOrganizationDeletableResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBodyResult) SetOrganizationId(v string) *DescribeFabricOrganizationDeletableResponseBodyResult {
	s.OrganizationId = &v
	return s
}

type DescribeFabricOrganizationDeletableResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricOrganizationDeletableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricOrganizationDeletableResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationDeletableResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationDeletableResponse) SetHeaders(v map[string]*string) *DescribeFabricOrganizationDeletableResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponse) SetBody(v *DescribeFabricOrganizationDeletableResponseBody) *DescribeFabricOrganizationDeletableResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationMembersRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Location       *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s DescribeFabricOrganizationMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationMembersRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationMembersRequest) SetOrganizationId(v string) *DescribeFabricOrganizationMembersRequest {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricOrganizationMembersRequest) SetLocation(v string) *DescribeFabricOrganizationMembersRequest {
	s.Location = &v
	return s
}

type DescribeFabricOrganizationMembersResponseBody struct {
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DescribeFabricOrganizationMembersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeFabricOrganizationMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationMembersResponseBody) SetRequestId(v string) *DescribeFabricOrganizationMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationMembersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBody) SetResult(v []*DescribeFabricOrganizationMembersResponseBodyResult) *DescribeFabricOrganizationMembersResponseBody {
	s.Result = v
	return s
}

type DescribeFabricOrganizationMembersResponseBodyResult struct {
	Domain           *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ConsortiumId     *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	JoinedTime       *string `json:"JoinedTime,omitempty" xml:"JoinedTime,omitempty"`
	ConsortiumName   *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	OrganizationId   *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DescribeFabricOrganizationMembersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationMembersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationMembersResponseBodyResult) SetDomain(v string) *DescribeFabricOrganizationMembersResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBodyResult) SetConsortiumId(v string) *DescribeFabricOrganizationMembersResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBodyResult) SetDescription(v string) *DescribeFabricOrganizationMembersResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBodyResult) SetState(v string) *DescribeFabricOrganizationMembersResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBodyResult) SetOrganizationName(v string) *DescribeFabricOrganizationMembersResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBodyResult) SetJoinedTime(v string) *DescribeFabricOrganizationMembersResponseBodyResult {
	s.JoinedTime = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBodyResult) SetConsortiumName(v string) *DescribeFabricOrganizationMembersResponseBodyResult {
	s.ConsortiumName = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBodyResult) SetOrganizationId(v string) *DescribeFabricOrganizationMembersResponseBodyResult {
	s.OrganizationId = &v
	return s
}

type DescribeFabricOrganizationMembersResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricOrganizationMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricOrganizationMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationMembersResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationMembersResponse) SetHeaders(v map[string]*string) *DescribeFabricOrganizationMembersResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricOrganizationMembersResponse) SetBody(v *DescribeFabricOrganizationMembersResponseBody) *DescribeFabricOrganizationMembersResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationPeersRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Location       *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s DescribeFabricOrganizationPeersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationPeersRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationPeersRequest) SetOrganizationId(v string) *DescribeFabricOrganizationPeersRequest {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricOrganizationPeersRequest) SetLocation(v string) *DescribeFabricOrganizationPeersRequest {
	s.Location = &v
	return s
}

type DescribeFabricOrganizationPeersResponseBody struct {
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DescribeFabricOrganizationPeersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeFabricOrganizationPeersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationPeersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationPeersResponseBody) SetRequestId(v string) *DescribeFabricOrganizationPeersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationPeersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationPeersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBody) SetResult(v []*DescribeFabricOrganizationPeersResponseBodyResult) *DescribeFabricOrganizationPeersResponseBody {
	s.Result = v
	return s
}

type DescribeFabricOrganizationPeersResponseBodyResult struct {
	UpdateTime           *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Domain               *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InternetIp           *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	IsAnchor             *bool   `json:"IsAnchor,omitempty" xml:"IsAnchor,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Port                 *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	OrganizationPeerName *string `json:"OrganizationPeerName,omitempty" xml:"OrganizationPeerName,omitempty"`
	IntranetIp           *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
}

func (s DescribeFabricOrganizationPeersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationPeersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetUpdateTime(v string) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetDomain(v string) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetInternetIp(v string) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.InternetIp = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetCreateTime(v string) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetIsAnchor(v bool) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.IsAnchor = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetInstanceType(v string) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.InstanceType = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetPort(v int32) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.Port = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetOrganizationPeerName(v string) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.OrganizationPeerName = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetIntranetIp(v string) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.IntranetIp = &v
	return s
}

type DescribeFabricOrganizationPeersResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricOrganizationPeersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricOrganizationPeersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationPeersResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationPeersResponse) SetHeaders(v map[string]*string) *DescribeFabricOrganizationPeersResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricOrganizationPeersResponse) SetBody(v *DescribeFabricOrganizationPeersResponseBody) *DescribeFabricOrganizationPeersResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationsRequest struct {
	Location *string                                  `json:"Location,omitempty" xml:"Location,omitempty"`
	Tag      []*DescribeFabricOrganizationsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeFabricOrganizationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationsRequest) SetLocation(v string) *DescribeFabricOrganizationsRequest {
	s.Location = &v
	return s
}

func (s *DescribeFabricOrganizationsRequest) SetTag(v []*DescribeFabricOrganizationsRequestTag) *DescribeFabricOrganizationsRequest {
	s.Tag = v
	return s
}

type DescribeFabricOrganizationsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFabricOrganizationsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationsRequestTag) SetKey(v string) *DescribeFabricOrganizationsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeFabricOrganizationsRequestTag) SetValue(v string) *DescribeFabricOrganizationsRequestTag {
	s.Value = &v
	return s
}

type DescribeFabricOrganizationsResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DescribeFabricOrganizationsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeFabricOrganizationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationsResponseBody) SetRequestId(v string) *DescribeFabricOrganizationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBody) SetResult(v []*DescribeFabricOrganizationsResponseBodyResult) *DescribeFabricOrganizationsResponseBody {
	s.Result = v
	return s
}

type DescribeFabricOrganizationsResponseBodyResult struct {
	Domain                  *string                                              `json:"Domain,omitempty" xml:"Domain,omitempty"`
	PeerCount               *int32                                               `json:"PeerCount,omitempty" xml:"PeerCount,omitempty"`
	State                   *string                                              `json:"State,omitempty" xml:"State,omitempty"`
	CreateTime              *string                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ConsortiumCount         *int32                                               `json:"ConsortiumCount,omitempty" xml:"ConsortiumCount,omitempty"`
	Tags                    []*DescribeFabricOrganizationsResponseBodyResultTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	SpecName                *string                                              `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	OwnerName               *string                                              `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerUid                *int64                                               `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	CodeName                *string                                              `json:"CodeName,omitempty" xml:"CodeName,omitempty"`
	OwnerBid                *string                                              `json:"OwnerBid,omitempty" xml:"OwnerBid,omitempty"`
	OrganizationDescription *string                                              `json:"OrganizationDescription,omitempty" xml:"OrganizationDescription,omitempty"`
	RegionId                *string                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OrganizationId          *string                                              `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	RequestId               *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ZoneId                  *string                                              `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	UserCount               *int32                                               `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	OrganizationName        *string                                              `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
}

func (s DescribeFabricOrganizationsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetDomain(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetPeerCount(v int32) *DescribeFabricOrganizationsResponseBodyResult {
	s.PeerCount = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetState(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetCreateTime(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetConsortiumCount(v int32) *DescribeFabricOrganizationsResponseBodyResult {
	s.ConsortiumCount = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetTags(v []*DescribeFabricOrganizationsResponseBodyResultTags) *DescribeFabricOrganizationsResponseBodyResult {
	s.Tags = v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetSpecName(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.SpecName = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetOwnerName(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetOwnerUid(v int64) *DescribeFabricOrganizationsResponseBodyResult {
	s.OwnerUid = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetCodeName(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.CodeName = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetOwnerBid(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.OwnerBid = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetOrganizationDescription(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.OrganizationDescription = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetRegionId(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetOrganizationId(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetRequestId(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetZoneId(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.ZoneId = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetUserCount(v int32) *DescribeFabricOrganizationsResponseBodyResult {
	s.UserCount = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetOrganizationName(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.OrganizationName = &v
	return s
}

type DescribeFabricOrganizationsResponseBodyResultTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFabricOrganizationsResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationsResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationsResponseBodyResultTags) SetKey(v string) *DescribeFabricOrganizationsResponseBodyResultTags {
	s.Key = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResultTags) SetValue(v string) *DescribeFabricOrganizationsResponseBodyResultTags {
	s.Value = &v
	return s
}

type DescribeFabricOrganizationsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricOrganizationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricOrganizationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationsResponse) SetHeaders(v map[string]*string) *DescribeFabricOrganizationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricOrganizationsResponse) SetBody(v *DescribeFabricOrganizationsResponseBody) *DescribeFabricOrganizationsResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationSpecsResponseBody struct {
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DescribeFabricOrganizationSpecsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeFabricOrganizationSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationSpecsResponseBody) SetRequestId(v string) *DescribeFabricOrganizationSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationSpecsResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationSpecsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationSpecsResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationSpecsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricOrganizationSpecsResponseBody) SetResult(v []*DescribeFabricOrganizationSpecsResponseBodyResult) *DescribeFabricOrganizationSpecsResponseBody {
	s.Result = v
	return s
}

type DescribeFabricOrganizationSpecsResponseBodyResult struct {
	Title                 *string `json:"Title,omitempty" xml:"Title,omitempty"`
	OrganizationSpecsName *string `json:"OrganizationSpecsName,omitempty" xml:"OrganizationSpecsName,omitempty"`
	Enable                *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeFabricOrganizationSpecsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationSpecsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationSpecsResponseBodyResult) SetTitle(v string) *DescribeFabricOrganizationSpecsResponseBodyResult {
	s.Title = &v
	return s
}

func (s *DescribeFabricOrganizationSpecsResponseBodyResult) SetOrganizationSpecsName(v string) *DescribeFabricOrganizationSpecsResponseBodyResult {
	s.OrganizationSpecsName = &v
	return s
}

func (s *DescribeFabricOrganizationSpecsResponseBodyResult) SetEnable(v bool) *DescribeFabricOrganizationSpecsResponseBodyResult {
	s.Enable = &v
	return s
}

type DescribeFabricOrganizationSpecsResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricOrganizationSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricOrganizationSpecsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationSpecsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationSpecsResponse) SetHeaders(v map[string]*string) *DescribeFabricOrganizationSpecsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricOrganizationSpecsResponse) SetBody(v *DescribeFabricOrganizationSpecsResponseBody) *DescribeFabricOrganizationSpecsResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationUsersRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Location       *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s DescribeFabricOrganizationUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationUsersRequest) SetOrganizationId(v string) *DescribeFabricOrganizationUsersRequest {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricOrganizationUsersRequest) SetLocation(v string) *DescribeFabricOrganizationUsersRequest {
	s.Location = &v
	return s
}

type DescribeFabricOrganizationUsersResponseBody struct {
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DescribeFabricOrganizationUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeFabricOrganizationUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationUsersResponseBody) SetRequestId(v string) *DescribeFabricOrganizationUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationUsersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationUsersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBody) SetResult(v []*DescribeFabricOrganizationUsersResponseBodyResult) *DescribeFabricOrganizationUsersResponseBody {
	s.Result = v
	return s
}

type DescribeFabricOrganizationUsersResponseBodyResult struct {
	ExpireTime     *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Attrs          *string `json:"Attrs,omitempty" xml:"Attrs,omitempty"`
	CallerBid      *string `json:"CallerBid,omitempty" xml:"CallerBid,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FullName       *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	CallerUid      *int64  `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	Username       *string `json:"Username,omitempty" xml:"Username,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFabricOrganizationUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationUsersResponseBodyResult) SetExpireTime(v string) *DescribeFabricOrganizationUsersResponseBodyResult {
	s.ExpireTime = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBodyResult) SetAttrs(v string) *DescribeFabricOrganizationUsersResponseBodyResult {
	s.Attrs = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBodyResult) SetCallerBid(v string) *DescribeFabricOrganizationUsersResponseBodyResult {
	s.CallerBid = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBodyResult) SetCreateTime(v string) *DescribeFabricOrganizationUsersResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBodyResult) SetFullName(v string) *DescribeFabricOrganizationUsersResponseBodyResult {
	s.FullName = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBodyResult) SetCallerUid(v int64) *DescribeFabricOrganizationUsersResponseBodyResult {
	s.CallerUid = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBodyResult) SetUsername(v string) *DescribeFabricOrganizationUsersResponseBodyResult {
	s.Username = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBodyResult) SetOrganizationId(v string) *DescribeFabricOrganizationUsersResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBodyResult) SetRegionId(v string) *DescribeFabricOrganizationUsersResponseBodyResult {
	s.RegionId = &v
	return s
}

type DescribeFabricOrganizationUsersResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricOrganizationUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricOrganizationUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationUsersResponse) SetHeaders(v map[string]*string) *DescribeFabricOrganizationUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricOrganizationUsersResponse) SetBody(v *DescribeFabricOrganizationUsersResponseBody) *DescribeFabricOrganizationUsersResponse {
	s.Body = v
	return s
}

type DescribeFabricPeerLogsRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	PeerName       *string `json:"PeerName,omitempty" xml:"PeerName,omitempty"`
	Lines          *string `json:"Lines,omitempty" xml:"Lines,omitempty"`
}

func (s DescribeFabricPeerLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricPeerLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricPeerLogsRequest) SetOrganizationId(v string) *DescribeFabricPeerLogsRequest {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricPeerLogsRequest) SetPeerName(v string) *DescribeFabricPeerLogsRequest {
	s.PeerName = &v
	return s
}

func (s *DescribeFabricPeerLogsRequest) SetLines(v string) *DescribeFabricPeerLogsRequest {
	s.Lines = &v
	return s
}

type DescribeFabricPeerLogsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeFabricPeerLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricPeerLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricPeerLogsResponseBody) SetRequestId(v string) *DescribeFabricPeerLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricPeerLogsResponseBody) SetErrorCode(v int32) *DescribeFabricPeerLogsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricPeerLogsResponseBody) SetSuccess(v bool) *DescribeFabricPeerLogsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFabricPeerLogsResponseBody) SetResult(v string) *DescribeFabricPeerLogsResponseBody {
	s.Result = &v
	return s
}

type DescribeFabricPeerLogsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFabricPeerLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFabricPeerLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricPeerLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricPeerLogsResponse) SetHeaders(v map[string]*string) *DescribeFabricPeerLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricPeerLogsResponse) SetBody(v *DescribeFabricPeerLogsResponseBody) *DescribeFabricPeerLogsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	ErrorCode *int32                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetErrorCode(v int32) *DescribeRegionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeRootDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeRootDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRootDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRootDomainResponseBody) SetRequestId(v string) *DescribeRootDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRootDomainResponseBody) SetErrorCode(v int32) *DescribeRootDomainResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeRootDomainResponseBody) SetSuccess(v bool) *DescribeRootDomainResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRootDomainResponseBody) SetResult(v string) *DescribeRootDomainResponseBody {
	s.Result = &v
	return s
}

type DescribeRootDomainResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRootDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRootDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRootDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeRootDomainResponse) SetHeaders(v map[string]*string) *DescribeRootDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeRootDomainResponse) SetBody(v *DescribeRootDomainResponseBody) *DescribeRootDomainResponse {
	s.Body = v
	return s
}

type DescribeTasksResponseBody struct {
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DynamicCode    *string                            `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                            `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *int32                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Result         []*DescribeTasksResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBody) SetRequestId(v string) *DescribeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTasksResponseBody) SetDynamicCode(v string) *DescribeTasksResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeTasksResponseBody) SetDynamicMessage(v string) *DescribeTasksResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeTasksResponseBody) SetErrorCode(v int32) *DescribeTasksResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeTasksResponseBody) SetSuccess(v bool) *DescribeTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTasksResponseBody) SetResult(v []*DescribeTasksResponseBodyResult) *DescribeTasksResponseBody {
	s.Result = v
	return s
}

type DescribeTasksResponseBodyResult struct {
	Action        *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Result        *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Sender        *string `json:"Sender,omitempty" xml:"Sender,omitempty"`
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	Handled       *bool   `json:"Handled,omitempty" xml:"Handled,omitempty"`
	ResponseTime  *string `json:"ResponseTime,omitempty" xml:"ResponseTime,omitempty"`
	Target        *string `json:"Target,omitempty" xml:"Target,omitempty"`
	TaskId        *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestTime   *int64  `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
}

func (s DescribeTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyResult) SetAction(v string) *DescribeTasksResponseBodyResult {
	s.Action = &v
	return s
}

func (s *DescribeTasksResponseBodyResult) SetResult(v string) *DescribeTasksResponseBodyResult {
	s.Result = &v
	return s
}

func (s *DescribeTasksResponseBodyResult) SetSender(v string) *DescribeTasksResponseBodyResult {
	s.Sender = &v
	return s
}

func (s *DescribeTasksResponseBodyResult) SetOperationType(v string) *DescribeTasksResponseBodyResult {
	s.OperationType = &v
	return s
}

func (s *DescribeTasksResponseBodyResult) SetHandled(v bool) *DescribeTasksResponseBodyResult {
	s.Handled = &v
	return s
}

func (s *DescribeTasksResponseBodyResult) SetResponseTime(v string) *DescribeTasksResponseBodyResult {
	s.ResponseTime = &v
	return s
}

func (s *DescribeTasksResponseBodyResult) SetTarget(v string) *DescribeTasksResponseBodyResult {
	s.Target = &v
	return s
}

func (s *DescribeTasksResponseBodyResult) SetTaskId(v int32) *DescribeTasksResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *DescribeTasksResponseBodyResult) SetRequestTime(v int64) *DescribeTasksResponseBodyResult {
	s.RequestTime = &v
	return s
}

type DescribeTasksResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponse) SetHeaders(v map[string]*string) *DescribeTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeTasksResponse) SetBody(v *DescribeTasksResponseBody) *DescribeTasksResponse {
	s.Body = v
	return s
}

type DownloadFabricOrganizationSDKRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Username       *string `json:"Username,omitempty" xml:"Username,omitempty"`
	Location       *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s DownloadFabricOrganizationSDKRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadFabricOrganizationSDKRequest) GoString() string {
	return s.String()
}

func (s *DownloadFabricOrganizationSDKRequest) SetOrganizationId(v string) *DownloadFabricOrganizationSDKRequest {
	s.OrganizationId = &v
	return s
}

func (s *DownloadFabricOrganizationSDKRequest) SetUsername(v string) *DownloadFabricOrganizationSDKRequest {
	s.Username = &v
	return s
}

func (s *DownloadFabricOrganizationSDKRequest) SetLocation(v string) *DownloadFabricOrganizationSDKRequest {
	s.Location = &v
	return s
}

type DownloadFabricOrganizationSDKResponseBody struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*DownloadFabricOrganizationSDKResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DownloadFabricOrganizationSDKResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadFabricOrganizationSDKResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadFabricOrganizationSDKResponseBody) SetRequestId(v string) *DownloadFabricOrganizationSDKResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadFabricOrganizationSDKResponseBody) SetErrorCode(v int32) *DownloadFabricOrganizationSDKResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DownloadFabricOrganizationSDKResponseBody) SetSuccess(v bool) *DownloadFabricOrganizationSDKResponseBody {
	s.Success = &v
	return s
}

func (s *DownloadFabricOrganizationSDKResponseBody) SetResult(v []*DownloadFabricOrganizationSDKResponseBodyResult) *DownloadFabricOrganizationSDKResponseBody {
	s.Result = v
	return s
}

type DownloadFabricOrganizationSDKResponseBodyResult struct {
	Path    *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s DownloadFabricOrganizationSDKResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DownloadFabricOrganizationSDKResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DownloadFabricOrganizationSDKResponseBodyResult) SetPath(v string) *DownloadFabricOrganizationSDKResponseBodyResult {
	s.Path = &v
	return s
}

func (s *DownloadFabricOrganizationSDKResponseBodyResult) SetContent(v string) *DownloadFabricOrganizationSDKResponseBodyResult {
	s.Content = &v
	return s
}

type DownloadFabricOrganizationSDKResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DownloadFabricOrganizationSDKResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadFabricOrganizationSDKResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadFabricOrganizationSDKResponse) GoString() string {
	return s.String()
}

func (s *DownloadFabricOrganizationSDKResponse) SetHeaders(v map[string]*string) *DownloadFabricOrganizationSDKResponse {
	s.Headers = v
	return s
}

func (s *DownloadFabricOrganizationSDKResponse) SetBody(v *DownloadFabricOrganizationSDKResponseBody) *DownloadFabricOrganizationSDKResponse {
	s.Body = v
	return s
}

type FreezeAntChainAccountRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	Account    *string `json:"Account,omitempty" xml:"Account,omitempty"`
}

func (s FreezeAntChainAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s FreezeAntChainAccountRequest) GoString() string {
	return s.String()
}

func (s *FreezeAntChainAccountRequest) SetAntChainId(v string) *FreezeAntChainAccountRequest {
	s.AntChainId = &v
	return s
}

func (s *FreezeAntChainAccountRequest) SetAccount(v string) *FreezeAntChainAccountRequest {
	s.Account = &v
	return s
}

type FreezeAntChainAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s FreezeAntChainAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FreezeAntChainAccountResponseBody) GoString() string {
	return s.String()
}

func (s *FreezeAntChainAccountResponseBody) SetRequestId(v string) *FreezeAntChainAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *FreezeAntChainAccountResponseBody) SetResult(v string) *FreezeAntChainAccountResponseBody {
	s.Result = &v
	return s
}

type FreezeAntChainAccountResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FreezeAntChainAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FreezeAntChainAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s FreezeAntChainAccountResponse) GoString() string {
	return s.String()
}

func (s *FreezeAntChainAccountResponse) SetHeaders(v map[string]*string) *FreezeAntChainAccountResponse {
	s.Headers = v
	return s
}

func (s *FreezeAntChainAccountResponse) SetBody(v *FreezeAntChainAccountResponseBody) *FreezeAntChainAccountResponse {
	s.Body = v
	return s
}

type InstallFabricChaincodeRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	ChaincodeId    *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	Location       *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s InstallFabricChaincodeRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallFabricChaincodeRequest) GoString() string {
	return s.String()
}

func (s *InstallFabricChaincodeRequest) SetOrganizationId(v string) *InstallFabricChaincodeRequest {
	s.OrganizationId = &v
	return s
}

func (s *InstallFabricChaincodeRequest) SetChaincodeId(v string) *InstallFabricChaincodeRequest {
	s.ChaincodeId = &v
	return s
}

func (s *InstallFabricChaincodeRequest) SetLocation(v string) *InstallFabricChaincodeRequest {
	s.Location = &v
	return s
}

type InstallFabricChaincodeResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *InstallFabricChaincodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s InstallFabricChaincodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallFabricChaincodeResponseBody) GoString() string {
	return s.String()
}

func (s *InstallFabricChaincodeResponseBody) SetRequestId(v string) *InstallFabricChaincodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallFabricChaincodeResponseBody) SetErrorCode(v int32) *InstallFabricChaincodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InstallFabricChaincodeResponseBody) SetSuccess(v bool) *InstallFabricChaincodeResponseBody {
	s.Success = &v
	return s
}

func (s *InstallFabricChaincodeResponseBody) SetResult(v *InstallFabricChaincodeResponseBodyResult) *InstallFabricChaincodeResponseBody {
	s.Result = v
	return s
}

type InstallFabricChaincodeResponseBodyResult struct {
	Type             *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	EndorsePolicy    *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChaincodeId      *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ProviderName     *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	ChaincodeName    *string `json:"ChaincodeName,omitempty" xml:"ChaincodeName,omitempty"`
	Install          *bool   `json:"Install,omitempty" xml:"Install,omitempty"`
	Input            *string `json:"Input,omitempty" xml:"Input,omitempty"`
	ProviderId       *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	DeployTime       *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" xml:"ChaincodeVersion,omitempty"`
	ConsortiumId     *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ChannelName      *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	Path             *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s InstallFabricChaincodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s InstallFabricChaincodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InstallFabricChaincodeResponseBodyResult) SetType(v int32) *InstallFabricChaincodeResponseBodyResult {
	s.Type = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetEndorsePolicy(v string) *InstallFabricChaincodeResponseBodyResult {
	s.EndorsePolicy = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetState(v string) *InstallFabricChaincodeResponseBodyResult {
	s.State = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetCreateTime(v string) *InstallFabricChaincodeResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetChaincodeId(v string) *InstallFabricChaincodeResponseBodyResult {
	s.ChaincodeId = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetMessage(v string) *InstallFabricChaincodeResponseBodyResult {
	s.Message = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetProviderName(v string) *InstallFabricChaincodeResponseBodyResult {
	s.ProviderName = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetChaincodeName(v string) *InstallFabricChaincodeResponseBodyResult {
	s.ChaincodeName = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetInstall(v bool) *InstallFabricChaincodeResponseBodyResult {
	s.Install = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetInput(v string) *InstallFabricChaincodeResponseBodyResult {
	s.Input = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetProviderId(v string) *InstallFabricChaincodeResponseBodyResult {
	s.ProviderId = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetDeployTime(v string) *InstallFabricChaincodeResponseBodyResult {
	s.DeployTime = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetChaincodeVersion(v string) *InstallFabricChaincodeResponseBodyResult {
	s.ChaincodeVersion = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetConsortiumId(v string) *InstallFabricChaincodeResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetChannelName(v string) *InstallFabricChaincodeResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetPath(v string) *InstallFabricChaincodeResponseBodyResult {
	s.Path = &v
	return s
}

type InstallFabricChaincodeResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InstallFabricChaincodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallFabricChaincodeResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallFabricChaincodeResponse) GoString() string {
	return s.String()
}

func (s *InstallFabricChaincodeResponse) SetHeaders(v map[string]*string) *InstallFabricChaincodeResponse {
	s.Headers = v
	return s
}

func (s *InstallFabricChaincodeResponse) SetBody(v *InstallFabricChaincodeResponseBody) *InstallFabricChaincodeResponse {
	s.Body = v
	return s
}

type InstantiateFabricChaincodeRequest struct {
	OrganizationId   *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	ChaincodeId      *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	EndorsePolicy    *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	Location         *string `json:"Location,omitempty" xml:"Location,omitempty"`
	CollectionConfig *string `json:"CollectionConfig,omitempty" xml:"CollectionConfig,omitempty"`
}

func (s InstantiateFabricChaincodeRequest) String() string {
	return tea.Prettify(s)
}

func (s InstantiateFabricChaincodeRequest) GoString() string {
	return s.String()
}

func (s *InstantiateFabricChaincodeRequest) SetOrganizationId(v string) *InstantiateFabricChaincodeRequest {
	s.OrganizationId = &v
	return s
}

func (s *InstantiateFabricChaincodeRequest) SetChaincodeId(v string) *InstantiateFabricChaincodeRequest {
	s.ChaincodeId = &v
	return s
}

func (s *InstantiateFabricChaincodeRequest) SetEndorsePolicy(v string) *InstantiateFabricChaincodeRequest {
	s.EndorsePolicy = &v
	return s
}

func (s *InstantiateFabricChaincodeRequest) SetLocation(v string) *InstantiateFabricChaincodeRequest {
	s.Location = &v
	return s
}

func (s *InstantiateFabricChaincodeRequest) SetCollectionConfig(v string) *InstantiateFabricChaincodeRequest {
	s.CollectionConfig = &v
	return s
}

type InstantiateFabricChaincodeResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *InstantiateFabricChaincodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s InstantiateFabricChaincodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstantiateFabricChaincodeResponseBody) GoString() string {
	return s.String()
}

func (s *InstantiateFabricChaincodeResponseBody) SetRequestId(v string) *InstantiateFabricChaincodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBody) SetErrorCode(v int32) *InstantiateFabricChaincodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBody) SetSuccess(v bool) *InstantiateFabricChaincodeResponseBody {
	s.Success = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBody) SetResult(v *InstantiateFabricChaincodeResponseBodyResult) *InstantiateFabricChaincodeResponseBody {
	s.Result = v
	return s
}

type InstantiateFabricChaincodeResponseBodyResult struct {
	Type             *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	EndorsePolicy    *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChaincodeId      *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ProviderName     *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	ChaincodeName    *string `json:"ChaincodeName,omitempty" xml:"ChaincodeName,omitempty"`
	Install          *bool   `json:"Install,omitempty" xml:"Install,omitempty"`
	Input            *string `json:"Input,omitempty" xml:"Input,omitempty"`
	ProviderId       *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	DeployTime       *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" xml:"ChaincodeVersion,omitempty"`
	ConsortiumId     *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ChannelName      *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	Path             *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s InstantiateFabricChaincodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s InstantiateFabricChaincodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetType(v int32) *InstantiateFabricChaincodeResponseBodyResult {
	s.Type = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetEndorsePolicy(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.EndorsePolicy = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetState(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.State = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetCreateTime(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetChaincodeId(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.ChaincodeId = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetMessage(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.Message = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetProviderName(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.ProviderName = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetChaincodeName(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.ChaincodeName = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetInstall(v bool) *InstantiateFabricChaincodeResponseBodyResult {
	s.Install = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetInput(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.Input = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetProviderId(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.ProviderId = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetDeployTime(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.DeployTime = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetChaincodeVersion(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.ChaincodeVersion = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetConsortiumId(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetChannelName(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetPath(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.Path = &v
	return s
}

type InstantiateFabricChaincodeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InstantiateFabricChaincodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstantiateFabricChaincodeResponse) String() string {
	return tea.Prettify(s)
}

func (s InstantiateFabricChaincodeResponse) GoString() string {
	return s.String()
}

func (s *InstantiateFabricChaincodeResponse) SetHeaders(v map[string]*string) *InstantiateFabricChaincodeResponse {
	s.Headers = v
	return s
}

func (s *InstantiateFabricChaincodeResponse) SetBody(v *InstantiateFabricChaincodeResponseBody) *InstantiateFabricChaincodeResponse {
	s.Body = v
	return s
}

type JoinFabricChannelRequest struct {
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Do        *string `json:"Do,omitempty" xml:"Do,omitempty"`
	Location  *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s JoinFabricChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinFabricChannelRequest) GoString() string {
	return s.String()
}

func (s *JoinFabricChannelRequest) SetChannelId(v string) *JoinFabricChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *JoinFabricChannelRequest) SetDo(v string) *JoinFabricChannelRequest {
	s.Do = &v
	return s
}

func (s *JoinFabricChannelRequest) SetLocation(v string) *JoinFabricChannelRequest {
	s.Location = &v
	return s
}

type JoinFabricChannelResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    []*JoinFabricChannelResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s JoinFabricChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinFabricChannelResponseBody) GoString() string {
	return s.String()
}

func (s *JoinFabricChannelResponseBody) SetRequestId(v string) *JoinFabricChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinFabricChannelResponseBody) SetErrorCode(v int32) *JoinFabricChannelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *JoinFabricChannelResponseBody) SetSuccess(v bool) *JoinFabricChannelResponseBody {
	s.Success = &v
	return s
}

func (s *JoinFabricChannelResponseBody) SetResult(v []*JoinFabricChannelResponseBodyResult) *JoinFabricChannelResponseBody {
	s.Result = v
	return s
}

type JoinFabricChannelResponseBodyResult struct {
	WithPeer       *bool   `json:"WithPeer,omitempty" xml:"WithPeer,omitempty"`
	AcceptTime     *string `json:"AcceptTime,omitempty" xml:"AcceptTime,omitempty"`
	State          *string `json:"State,omitempty" xml:"State,omitempty"`
	DestroyTime    *string `json:"DestroyTime,omitempty" xml:"DestroyTime,omitempty"`
	InviteTime     *string `json:"InviteTime,omitempty" xml:"InviteTime,omitempty"`
	ChannelId      *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ConfirmTime    *string `json:"ConfirmTime,omitempty" xml:"ConfirmTime,omitempty"`
	ApproveTime    *string `json:"ApproveTime,omitempty" xml:"ApproveTime,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s JoinFabricChannelResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s JoinFabricChannelResponseBodyResult) GoString() string {
	return s.String()
}

func (s *JoinFabricChannelResponseBodyResult) SetWithPeer(v bool) *JoinFabricChannelResponseBodyResult {
	s.WithPeer = &v
	return s
}

func (s *JoinFabricChannelResponseBodyResult) SetAcceptTime(v string) *JoinFabricChannelResponseBodyResult {
	s.AcceptTime = &v
	return s
}

func (s *JoinFabricChannelResponseBodyResult) SetState(v string) *JoinFabricChannelResponseBodyResult {
	s.State = &v
	return s
}

func (s *JoinFabricChannelResponseBodyResult) SetDestroyTime(v string) *JoinFabricChannelResponseBodyResult {
	s.DestroyTime = &v
	return s
}

func (s *JoinFabricChannelResponseBodyResult) SetInviteTime(v string) *JoinFabricChannelResponseBodyResult {
	s.InviteTime = &v
	return s
}

func (s *JoinFabricChannelResponseBodyResult) SetChannelId(v string) *JoinFabricChannelResponseBodyResult {
	s.ChannelId = &v
	return s
}

func (s *JoinFabricChannelResponseBodyResult) SetConfirmTime(v string) *JoinFabricChannelResponseBodyResult {
	s.ConfirmTime = &v
	return s
}

func (s *JoinFabricChannelResponseBodyResult) SetApproveTime(v string) *JoinFabricChannelResponseBodyResult {
	s.ApproveTime = &v
	return s
}

func (s *JoinFabricChannelResponseBodyResult) SetOrganizationId(v string) *JoinFabricChannelResponseBodyResult {
	s.OrganizationId = &v
	return s
}

type JoinFabricChannelResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *JoinFabricChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JoinFabricChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinFabricChannelResponse) GoString() string {
	return s.String()
}

func (s *JoinFabricChannelResponse) SetHeaders(v map[string]*string) *JoinFabricChannelResponse {
	s.Headers = v
	return s
}

func (s *JoinFabricChannelResponse) SetBody(v *JoinFabricChannelResponseBody) *JoinFabricChannelResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ResetAntChainCertificateRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s ResetAntChainCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAntChainCertificateRequest) GoString() string {
	return s.String()
}

func (s *ResetAntChainCertificateRequest) SetAntChainId(v string) *ResetAntChainCertificateRequest {
	s.AntChainId = &v
	return s
}

type ResetAntChainCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ResetAntChainCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAntChainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAntChainCertificateResponseBody) SetRequestId(v string) *ResetAntChainCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAntChainCertificateResponseBody) SetResult(v string) *ResetAntChainCertificateResponseBody {
	s.Result = &v
	return s
}

type ResetAntChainCertificateResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetAntChainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetAntChainCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAntChainCertificateResponse) GoString() string {
	return s.String()
}

func (s *ResetAntChainCertificateResponse) SetHeaders(v map[string]*string) *ResetAntChainCertificateResponse {
	s.Headers = v
	return s
}

func (s *ResetAntChainCertificateResponse) SetBody(v *ResetAntChainCertificateResponseBody) *ResetAntChainCertificateResponse {
	s.Body = v
	return s
}

type ResetAntChainUserCertificateRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ResetAntChainUserCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAntChainUserCertificateRequest) GoString() string {
	return s.String()
}

func (s *ResetAntChainUserCertificateRequest) SetAntChainId(v string) *ResetAntChainUserCertificateRequest {
	s.AntChainId = &v
	return s
}

func (s *ResetAntChainUserCertificateRequest) SetUsername(v string) *ResetAntChainUserCertificateRequest {
	s.Username = &v
	return s
}

type ResetAntChainUserCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ResetAntChainUserCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAntChainUserCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAntChainUserCertificateResponseBody) SetRequestId(v string) *ResetAntChainUserCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAntChainUserCertificateResponseBody) SetResult(v string) *ResetAntChainUserCertificateResponseBody {
	s.Result = &v
	return s
}

type ResetAntChainUserCertificateResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetAntChainUserCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetAntChainUserCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAntChainUserCertificateResponse) GoString() string {
	return s.String()
}

func (s *ResetAntChainUserCertificateResponse) SetHeaders(v map[string]*string) *ResetAntChainUserCertificateResponse {
	s.Headers = v
	return s
}

func (s *ResetAntChainUserCertificateResponse) SetBody(v *ResetAntChainUserCertificateResponseBody) *ResetAntChainUserCertificateResponse {
	s.Body = v
	return s
}

type ResetFabricOrganizationUserPasswordRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Username       *string `json:"Username,omitempty" xml:"Username,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Location       *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s ResetFabricOrganizationUserPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetFabricOrganizationUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetFabricOrganizationUserPasswordRequest) SetOrganizationId(v string) *ResetFabricOrganizationUserPasswordRequest {
	s.OrganizationId = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordRequest) SetUsername(v string) *ResetFabricOrganizationUserPasswordRequest {
	s.Username = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordRequest) SetPassword(v string) *ResetFabricOrganizationUserPasswordRequest {
	s.Password = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordRequest) SetLocation(v string) *ResetFabricOrganizationUserPasswordRequest {
	s.Location = &v
	return s
}

type ResetFabricOrganizationUserPasswordResponseBody struct {
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *ResetFabricOrganizationUserPasswordResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ResetFabricOrganizationUserPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetFabricOrganizationUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetFabricOrganizationUserPasswordResponseBody) SetRequestId(v string) *ResetFabricOrganizationUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponseBody) SetErrorCode(v int32) *ResetFabricOrganizationUserPasswordResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponseBody) SetSuccess(v bool) *ResetFabricOrganizationUserPasswordResponseBody {
	s.Success = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponseBody) SetResult(v *ResetFabricOrganizationUserPasswordResponseBodyResult) *ResetFabricOrganizationUserPasswordResponseBody {
	s.Result = v
	return s
}

type ResetFabricOrganizationUserPasswordResponseBodyResult struct {
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ExpireTime     *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Username       *string `json:"Username,omitempty" xml:"Username,omitempty"`
	Fullname       *string `json:"Fullname,omitempty" xml:"Fullname,omitempty"`
}

func (s ResetFabricOrganizationUserPasswordResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ResetFabricOrganizationUserPasswordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ResetFabricOrganizationUserPasswordResponseBodyResult) SetPassword(v string) *ResetFabricOrganizationUserPasswordResponseBodyResult {
	s.Password = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponseBodyResult) SetExpireTime(v string) *ResetFabricOrganizationUserPasswordResponseBodyResult {
	s.ExpireTime = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponseBodyResult) SetCreateTime(v string) *ResetFabricOrganizationUserPasswordResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponseBodyResult) SetOrganizationId(v string) *ResetFabricOrganizationUserPasswordResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponseBodyResult) SetUsername(v string) *ResetFabricOrganizationUserPasswordResponseBodyResult {
	s.Username = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponseBodyResult) SetFullname(v string) *ResetFabricOrganizationUserPasswordResponseBodyResult {
	s.Fullname = &v
	return s
}

type ResetFabricOrganizationUserPasswordResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetFabricOrganizationUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetFabricOrganizationUserPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetFabricOrganizationUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetFabricOrganizationUserPasswordResponse) SetHeaders(v map[string]*string) *ResetFabricOrganizationUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponse) SetBody(v *ResetFabricOrganizationUserPasswordResponseBody) *ResetFabricOrganizationUserPasswordResponse {
	s.Body = v
	return s
}

type SynchronizeFabricChaincodeRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	ChaincodeId    *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
}

func (s SynchronizeFabricChaincodeRequest) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeFabricChaincodeRequest) GoString() string {
	return s.String()
}

func (s *SynchronizeFabricChaincodeRequest) SetOrganizationId(v string) *SynchronizeFabricChaincodeRequest {
	s.OrganizationId = &v
	return s
}

func (s *SynchronizeFabricChaincodeRequest) SetChaincodeId(v string) *SynchronizeFabricChaincodeRequest {
	s.ChaincodeId = &v
	return s
}

type SynchronizeFabricChaincodeResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *SynchronizeFabricChaincodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s SynchronizeFabricChaincodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeFabricChaincodeResponseBody) GoString() string {
	return s.String()
}

func (s *SynchronizeFabricChaincodeResponseBody) SetRequestId(v string) *SynchronizeFabricChaincodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBody) SetErrorCode(v int32) *SynchronizeFabricChaincodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBody) SetSuccess(v bool) *SynchronizeFabricChaincodeResponseBody {
	s.Success = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBody) SetResult(v *SynchronizeFabricChaincodeResponseBodyResult) *SynchronizeFabricChaincodeResponseBody {
	s.Result = v
	return s
}

type SynchronizeFabricChaincodeResponseBodyResult struct {
	Type             *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	EndorsePolicy    *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChaincodeId      *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ProviderName     *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	ChaincodeName    *string `json:"ChaincodeName,omitempty" xml:"ChaincodeName,omitempty"`
	Install          *bool   `json:"Install,omitempty" xml:"Install,omitempty"`
	Input            *string `json:"Input,omitempty" xml:"Input,omitempty"`
	ProviderId       *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	DeployTime       *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" xml:"ChaincodeVersion,omitempty"`
	ConsortiumId     *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ChannelName      *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	Path             *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s SynchronizeFabricChaincodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeFabricChaincodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetType(v int32) *SynchronizeFabricChaincodeResponseBodyResult {
	s.Type = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetEndorsePolicy(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.EndorsePolicy = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetState(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.State = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetCreateTime(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetChaincodeId(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.ChaincodeId = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetMessage(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.Message = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetProviderName(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.ProviderName = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetChaincodeName(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.ChaincodeName = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetInstall(v bool) *SynchronizeFabricChaincodeResponseBodyResult {
	s.Install = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetInput(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.Input = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetProviderId(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.ProviderId = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetDeployTime(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.DeployTime = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetChaincodeVersion(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.ChaincodeVersion = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetConsortiumId(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetChannelName(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetPath(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.Path = &v
	return s
}

type SynchronizeFabricChaincodeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SynchronizeFabricChaincodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SynchronizeFabricChaincodeResponse) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeFabricChaincodeResponse) GoString() string {
	return s.String()
}

func (s *SynchronizeFabricChaincodeResponse) SetHeaders(v map[string]*string) *SynchronizeFabricChaincodeResponse {
	s.Headers = v
	return s
}

func (s *SynchronizeFabricChaincodeResponse) SetBody(v *SynchronizeFabricChaincodeResponseBody) *SynchronizeFabricChaincodeResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) SetErrorCode(v int32) *TagResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *TagResourcesResponseBody) SetResult(v bool) *TagResourcesResponseBody {
	s.Result = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnfreezeAntChainAccountRequest struct {
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	Account    *string `json:"Account,omitempty" xml:"Account,omitempty"`
}

func (s UnfreezeAntChainAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s UnfreezeAntChainAccountRequest) GoString() string {
	return s.String()
}

func (s *UnfreezeAntChainAccountRequest) SetAntChainId(v string) *UnfreezeAntChainAccountRequest {
	s.AntChainId = &v
	return s
}

func (s *UnfreezeAntChainAccountRequest) SetAccount(v string) *UnfreezeAntChainAccountRequest {
	s.Account = &v
	return s
}

type UnfreezeAntChainAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UnfreezeAntChainAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnfreezeAntChainAccountResponseBody) GoString() string {
	return s.String()
}

func (s *UnfreezeAntChainAccountResponseBody) SetRequestId(v string) *UnfreezeAntChainAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnfreezeAntChainAccountResponseBody) SetResult(v string) *UnfreezeAntChainAccountResponseBody {
	s.Result = &v
	return s
}

type UnfreezeAntChainAccountResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnfreezeAntChainAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnfreezeAntChainAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s UnfreezeAntChainAccountResponse) GoString() string {
	return s.String()
}

func (s *UnfreezeAntChainAccountResponse) SetHeaders(v map[string]*string) *UnfreezeAntChainAccountResponse {
	s.Headers = v
	return s
}

func (s *UnfreezeAntChainAccountResponse) SetBody(v *UnfreezeAntChainAccountResponseBody) *UnfreezeAntChainAccountResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetErrorCode(v int32) *UntagResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *UntagResourcesResponseBody) SetResult(v bool) *UntagResourcesResponseBody {
	s.Result = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateAntChainRequest struct {
	AntChainId   *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	AntChainName *string `json:"AntChainName,omitempty" xml:"AntChainName,omitempty"`
}

func (s UpdateAntChainRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainRequest) GoString() string {
	return s.String()
}

func (s *UpdateAntChainRequest) SetAntChainId(v string) *UpdateAntChainRequest {
	s.AntChainId = &v
	return s
}

func (s *UpdateAntChainRequest) SetAntChainName(v string) *UpdateAntChainRequest {
	s.AntChainName = &v
	return s
}

type UpdateAntChainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateAntChainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAntChainResponseBody) SetRequestId(v string) *UpdateAntChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAntChainResponseBody) SetResult(v string) *UpdateAntChainResponseBody {
	s.Result = &v
	return s
}

type UpdateAntChainResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAntChainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAntChainResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainResponse) GoString() string {
	return s.String()
}

func (s *UpdateAntChainResponse) SetHeaders(v map[string]*string) *UpdateAntChainResponse {
	s.Headers = v
	return s
}

func (s *UpdateAntChainResponse) SetBody(v *UpdateAntChainResponseBody) *UpdateAntChainResponse {
	s.Body = v
	return s
}

type UpdateAntChainConsortiumRequest struct {
	ConsortiumId          *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ConsortiumName        *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	ConsortiumDescription *string `json:"ConsortiumDescription,omitempty" xml:"ConsortiumDescription,omitempty"`
}

func (s UpdateAntChainConsortiumRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainConsortiumRequest) GoString() string {
	return s.String()
}

func (s *UpdateAntChainConsortiumRequest) SetConsortiumId(v string) *UpdateAntChainConsortiumRequest {
	s.ConsortiumId = &v
	return s
}

func (s *UpdateAntChainConsortiumRequest) SetConsortiumName(v string) *UpdateAntChainConsortiumRequest {
	s.ConsortiumName = &v
	return s
}

func (s *UpdateAntChainConsortiumRequest) SetConsortiumDescription(v string) *UpdateAntChainConsortiumRequest {
	s.ConsortiumDescription = &v
	return s
}

type UpdateAntChainConsortiumResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateAntChainConsortiumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainConsortiumResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAntChainConsortiumResponseBody) SetRequestId(v string) *UpdateAntChainConsortiumResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAntChainConsortiumResponseBody) SetResult(v string) *UpdateAntChainConsortiumResponseBody {
	s.Result = &v
	return s
}

type UpdateAntChainConsortiumResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAntChainConsortiumResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAntChainConsortiumResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainConsortiumResponse) GoString() string {
	return s.String()
}

func (s *UpdateAntChainConsortiumResponse) SetHeaders(v map[string]*string) *UpdateAntChainConsortiumResponse {
	s.Headers = v
	return s
}

func (s *UpdateAntChainConsortiumResponse) SetBody(v *UpdateAntChainConsortiumResponseBody) *UpdateAntChainConsortiumResponse {
	s.Body = v
	return s
}

type UpdateAntChainContractContentRequest struct {
	ContentId       *string `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	ParentContentId *string `json:"ParentContentId,omitempty" xml:"ParentContentId,omitempty"`
	ContentName     *string `json:"ContentName,omitempty" xml:"ContentName,omitempty"`
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s UpdateAntChainContractContentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainContractContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAntChainContractContentRequest) SetContentId(v string) *UpdateAntChainContractContentRequest {
	s.ContentId = &v
	return s
}

func (s *UpdateAntChainContractContentRequest) SetParentContentId(v string) *UpdateAntChainContractContentRequest {
	s.ParentContentId = &v
	return s
}

func (s *UpdateAntChainContractContentRequest) SetContentName(v string) *UpdateAntChainContractContentRequest {
	s.ContentName = &v
	return s
}

func (s *UpdateAntChainContractContentRequest) SetContent(v string) *UpdateAntChainContractContentRequest {
	s.Content = &v
	return s
}

type UpdateAntChainContractContentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateAntChainContractContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainContractContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAntChainContractContentResponseBody) SetRequestId(v string) *UpdateAntChainContractContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAntChainContractContentResponseBody) SetResult(v string) *UpdateAntChainContractContentResponseBody {
	s.Result = &v
	return s
}

type UpdateAntChainContractContentResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAntChainContractContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAntChainContractContentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainContractContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateAntChainContractContentResponse) SetHeaders(v map[string]*string) *UpdateAntChainContractContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateAntChainContractContentResponse) SetBody(v *UpdateAntChainContractContentResponseBody) *UpdateAntChainContractContentResponse {
	s.Body = v
	return s
}

type UpdateAntChainContractProjectRequest struct {
	ProjectId          *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectVersion     *string `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
}

func (s UpdateAntChainContractProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainContractProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateAntChainContractProjectRequest) SetProjectId(v string) *UpdateAntChainContractProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateAntChainContractProjectRequest) SetProjectName(v string) *UpdateAntChainContractProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateAntChainContractProjectRequest) SetProjectVersion(v string) *UpdateAntChainContractProjectRequest {
	s.ProjectVersion = &v
	return s
}

func (s *UpdateAntChainContractProjectRequest) SetProjectDescription(v string) *UpdateAntChainContractProjectRequest {
	s.ProjectDescription = &v
	return s
}

type UpdateAntChainContractProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateAntChainContractProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainContractProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAntChainContractProjectResponseBody) SetRequestId(v string) *UpdateAntChainContractProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAntChainContractProjectResponseBody) SetResult(v string) *UpdateAntChainContractProjectResponseBody {
	s.Result = &v
	return s
}

type UpdateAntChainContractProjectResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAntChainContractProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAntChainContractProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainContractProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateAntChainContractProjectResponse) SetHeaders(v map[string]*string) *UpdateAntChainContractProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateAntChainContractProjectResponse) SetBody(v *UpdateAntChainContractProjectResponseBody) *UpdateAntChainContractProjectResponse {
	s.Body = v
	return s
}

type UpdateAntChainMemberRequest struct {
	MemberName   *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	MemberId     *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s UpdateAntChainMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateAntChainMemberRequest) SetMemberName(v string) *UpdateAntChainMemberRequest {
	s.MemberName = &v
	return s
}

func (s *UpdateAntChainMemberRequest) SetConsortiumId(v string) *UpdateAntChainMemberRequest {
	s.ConsortiumId = &v
	return s
}

func (s *UpdateAntChainMemberRequest) SetMemberId(v string) *UpdateAntChainMemberRequest {
	s.MemberId = &v
	return s
}

type UpdateAntChainMemberResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateAntChainMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAntChainMemberResponseBody) SetRequestId(v string) *UpdateAntChainMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAntChainMemberResponseBody) SetResult(v string) *UpdateAntChainMemberResponseBody {
	s.Result = &v
	return s
}

type UpdateAntChainMemberResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAntChainMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAntChainMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateAntChainMemberResponse) SetHeaders(v map[string]*string) *UpdateAntChainMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateAntChainMemberResponse) SetBody(v *UpdateAntChainMemberResponseBody) *UpdateAntChainMemberResponse {
	s.Body = v
	return s
}

type UpdateAntChainQRCodeAuthorizationRequest struct {
	AntChainId        *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	QRCodeType        *string `json:"QRCodeType,omitempty" xml:"QRCodeType,omitempty"`
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
}

func (s UpdateAntChainQRCodeAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainQRCodeAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *UpdateAntChainQRCodeAuthorizationRequest) SetAntChainId(v string) *UpdateAntChainQRCodeAuthorizationRequest {
	s.AntChainId = &v
	return s
}

func (s *UpdateAntChainQRCodeAuthorizationRequest) SetQRCodeType(v string) *UpdateAntChainQRCodeAuthorizationRequest {
	s.QRCodeType = &v
	return s
}

func (s *UpdateAntChainQRCodeAuthorizationRequest) SetAuthorizationType(v string) *UpdateAntChainQRCodeAuthorizationRequest {
	s.AuthorizationType = &v
	return s
}

type UpdateAntChainQRCodeAuthorizationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateAntChainQRCodeAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainQRCodeAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAntChainQRCodeAuthorizationResponseBody) SetRequestId(v string) *UpdateAntChainQRCodeAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAntChainQRCodeAuthorizationResponseBody) SetResult(v string) *UpdateAntChainQRCodeAuthorizationResponseBody {
	s.Result = &v
	return s
}

type UpdateAntChainQRCodeAuthorizationResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAntChainQRCodeAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAntChainQRCodeAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainQRCodeAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *UpdateAntChainQRCodeAuthorizationResponse) SetHeaders(v map[string]*string) *UpdateAntChainQRCodeAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *UpdateAntChainQRCodeAuthorizationResponse) SetBody(v *UpdateAntChainQRCodeAuthorizationResponseBody) *UpdateAntChainQRCodeAuthorizationResponse {
	s.Body = v
	return s
}

type UpgradeFabricChaincodeRequest struct {
	OrganizationId   *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	ChaincodeId      *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	EndorsePolicy    *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	Location         *string `json:"Location,omitempty" xml:"Location,omitempty"`
	CollectionConfig *string `json:"CollectionConfig,omitempty" xml:"CollectionConfig,omitempty"`
}

func (s UpgradeFabricChaincodeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFabricChaincodeRequest) GoString() string {
	return s.String()
}

func (s *UpgradeFabricChaincodeRequest) SetOrganizationId(v string) *UpgradeFabricChaincodeRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpgradeFabricChaincodeRequest) SetChaincodeId(v string) *UpgradeFabricChaincodeRequest {
	s.ChaincodeId = &v
	return s
}

func (s *UpgradeFabricChaincodeRequest) SetEndorsePolicy(v string) *UpgradeFabricChaincodeRequest {
	s.EndorsePolicy = &v
	return s
}

func (s *UpgradeFabricChaincodeRequest) SetLocation(v string) *UpgradeFabricChaincodeRequest {
	s.Location = &v
	return s
}

func (s *UpgradeFabricChaincodeRequest) SetCollectionConfig(v string) *UpgradeFabricChaincodeRequest {
	s.CollectionConfig = &v
	return s
}

type UpgradeFabricChaincodeResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *int32                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *UpgradeFabricChaincodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpgradeFabricChaincodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFabricChaincodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeFabricChaincodeResponseBody) SetRequestId(v string) *UpgradeFabricChaincodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBody) SetErrorCode(v int32) *UpgradeFabricChaincodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBody) SetSuccess(v bool) *UpgradeFabricChaincodeResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBody) SetResult(v *UpgradeFabricChaincodeResponseBodyResult) *UpgradeFabricChaincodeResponseBody {
	s.Result = v
	return s
}

type UpgradeFabricChaincodeResponseBodyResult struct {
	Type             *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	EndorsePolicy    *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChaincodeId      *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ProviderName     *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	ChaincodeName    *string `json:"ChaincodeName,omitempty" xml:"ChaincodeName,omitempty"`
	Install          *bool   `json:"Install,omitempty" xml:"Install,omitempty"`
	Input            *string `json:"Input,omitempty" xml:"Input,omitempty"`
	ProviderId       *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	DeployTime       *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" xml:"ChaincodeVersion,omitempty"`
	ConsortiumId     *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ChannelName      *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	Path             *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s UpgradeFabricChaincodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFabricChaincodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetType(v int32) *UpgradeFabricChaincodeResponseBodyResult {
	s.Type = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetEndorsePolicy(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.EndorsePolicy = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetState(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.State = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetCreateTime(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetChaincodeId(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.ChaincodeId = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetMessage(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.Message = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetProviderName(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.ProviderName = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetChaincodeName(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.ChaincodeName = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetInstall(v bool) *UpgradeFabricChaincodeResponseBodyResult {
	s.Install = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetInput(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.Input = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetProviderId(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.ProviderId = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetDeployTime(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.DeployTime = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetChaincodeVersion(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.ChaincodeVersion = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetConsortiumId(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetChannelName(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetPath(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.Path = &v
	return s
}

type UpgradeFabricChaincodeResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeFabricChaincodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeFabricChaincodeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFabricChaincodeResponse) GoString() string {
	return s.String()
}

func (s *UpgradeFabricChaincodeResponse) SetHeaders(v map[string]*string) *UpgradeFabricChaincodeResponse {
	s.Headers = v
	return s
}

func (s *UpgradeFabricChaincodeResponse) SetBody(v *UpgradeFabricChaincodeResponseBody) *UpgradeFabricChaincodeResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":     tea.String("baas.aliyuncs.com"),
		"cn-beijing":     tea.String("baas.aliyuncs.com"),
		"cn-zhangjiakou": tea.String("baas.aliyuncs.com"),
		"cn-huhehaote":   tea.String("baas.aliyuncs.com"),
		"cn-shanghai":    tea.String("baas.aliyuncs.com"),
		"cn-shenzhen":    tea.String("baas.aliyuncs.com"),
		"cn-hongkong":    tea.String("baas.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2": tea.String("baas.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-1": tea.String("baas.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":      tea.String("baas.ap-southeast-1.aliyuncs.com"),
		"us-west-1":      tea.String("baas.ap-southeast-1.aliyuncs.com"),
		"us-east-1":      tea.String("baas.ap-southeast-1.aliyuncs.com"),
		"eu-central-1":   tea.String("baas.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":     tea.String("baas.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("baas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AcceptFabricInvitationWithOptions(request *AcceptFabricInvitationRequest, runtime *util.RuntimeOptions) (_result *AcceptFabricInvitationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AcceptFabricInvitationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AcceptFabricInvitation"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AcceptFabricInvitation(request *AcceptFabricInvitationRequest) (_result *AcceptFabricInvitationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AcceptFabricInvitationResponse{}
	_body, _err := client.AcceptFabricInvitationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyAntChainCertificateWithOptions(request *ApplyAntChainCertificateRequest, runtime *util.RuntimeOptions) (_result *ApplyAntChainCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApplyAntChainCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApplyAntChainCertificate"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyAntChainCertificate(request *ApplyAntChainCertificateRequest) (_result *ApplyAntChainCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyAntChainCertificateResponse{}
	_body, _err := client.ApplyAntChainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyAntChainCertificateWithKeyAutoCreationWithOptions(request *ApplyAntChainCertificateWithKeyAutoCreationRequest, runtime *util.RuntimeOptions) (_result *ApplyAntChainCertificateWithKeyAutoCreationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApplyAntChainCertificateWithKeyAutoCreationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApplyAntChainCertificateWithKeyAutoCreation"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyAntChainCertificateWithKeyAutoCreation(request *ApplyAntChainCertificateWithKeyAutoCreationRequest) (_result *ApplyAntChainCertificateWithKeyAutoCreationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyAntChainCertificateWithKeyAutoCreationResponse{}
	_body, _err := client.ApplyAntChainCertificateWithKeyAutoCreationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchAddAntChainMiniAppQRCodeAuthorizedUsersWithOptions(tmpReq *BatchAddAntChainMiniAppQRCodeAuthorizedUsersRequest, runtime *util.RuntimeOptions) (_result *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchAddAntChainMiniAppQRCodeAuthorizedUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PhoneList)) {
		request.PhoneListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PhoneList, tea.String("PhoneList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchAddAntChainMiniAppQRCodeAuthorizedUsers"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchAddAntChainMiniAppQRCodeAuthorizedUsers(request *BatchAddAntChainMiniAppQRCodeAuthorizedUsersRequest) (_result *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse{}
	_body, _err := client.BatchAddAntChainMiniAppQRCodeAuthorizedUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckFabricConsortiumDomainWithOptions(request *CheckFabricConsortiumDomainRequest, runtime *util.RuntimeOptions) (_result *CheckFabricConsortiumDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckFabricConsortiumDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckFabricConsortiumDomain"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckFabricConsortiumDomain(request *CheckFabricConsortiumDomainRequest) (_result *CheckFabricConsortiumDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckFabricConsortiumDomainResponse{}
	_body, _err := client.CheckFabricConsortiumDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckFabricOrganizationDomainWithOptions(request *CheckFabricOrganizationDomainRequest, runtime *util.RuntimeOptions) (_result *CheckFabricOrganizationDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckFabricOrganizationDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckFabricOrganizationDomain"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckFabricOrganizationDomain(request *CheckFabricOrganizationDomainRequest) (_result *CheckFabricOrganizationDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckFabricOrganizationDomainResponse{}
	_body, _err := client.CheckFabricOrganizationDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfirmFabricConsortiumMemberWithOptions(request *ConfirmFabricConsortiumMemberRequest, runtime *util.RuntimeOptions) (_result *ConfirmFabricConsortiumMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfirmFabricConsortiumMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfirmFabricConsortiumMember"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfirmFabricConsortiumMember(request *ConfirmFabricConsortiumMemberRequest) (_result *ConfirmFabricConsortiumMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfirmFabricConsortiumMemberResponse{}
	_body, _err := client.ConfirmFabricConsortiumMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CopyAntChainContractProjectWithOptions(request *CopyAntChainContractProjectRequest, runtime *util.RuntimeOptions) (_result *CopyAntChainContractProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CopyAntChainContractProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CopyAntChainContractProject"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CopyAntChainContractProject(request *CopyAntChainContractProjectRequest) (_result *CopyAntChainContractProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CopyAntChainContractProjectResponse{}
	_body, _err := client.CopyAntChainContractProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAntChainAccountWithOptions(request *CreateAntChainAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAntChainAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAntChainAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAntChainAccount"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAntChainAccount(request *CreateAntChainAccountRequest) (_result *CreateAntChainAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAntChainAccountResponse{}
	_body, _err := client.CreateAntChainAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAntChainAccountWithKeyPairAutoCreationWithOptions(request *CreateAntChainAccountWithKeyPairAutoCreationRequest, runtime *util.RuntimeOptions) (_result *CreateAntChainAccountWithKeyPairAutoCreationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAntChainAccountWithKeyPairAutoCreationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAntChainAccountWithKeyPairAutoCreation"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAntChainAccountWithKeyPairAutoCreation(request *CreateAntChainAccountWithKeyPairAutoCreationRequest) (_result *CreateAntChainAccountWithKeyPairAutoCreationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAntChainAccountWithKeyPairAutoCreationResponse{}
	_body, _err := client.CreateAntChainAccountWithKeyPairAutoCreationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAntChainConsortiumWithOptions(request *CreateAntChainConsortiumRequest, runtime *util.RuntimeOptions) (_result *CreateAntChainConsortiumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAntChainConsortiumResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAntChainConsortium"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAntChainConsortium(request *CreateAntChainConsortiumRequest) (_result *CreateAntChainConsortiumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAntChainConsortiumResponse{}
	_body, _err := client.CreateAntChainConsortiumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAntChainContractContentWithOptions(request *CreateAntChainContractContentRequest, runtime *util.RuntimeOptions) (_result *CreateAntChainContractContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAntChainContractContentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAntChainContractContent"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAntChainContractContent(request *CreateAntChainContractContentRequest) (_result *CreateAntChainContractContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAntChainContractContentResponse{}
	_body, _err := client.CreateAntChainContractContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAntChainContractProjectWithOptions(request *CreateAntChainContractProjectRequest, runtime *util.RuntimeOptions) (_result *CreateAntChainContractProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAntChainContractProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAntChainContractProject"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAntChainContractProject(request *CreateAntChainContractProjectRequest) (_result *CreateAntChainContractProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAntChainContractProjectResponse{}
	_body, _err := client.CreateAntChainContractProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFabricChaincodeWithOptions(request *CreateFabricChaincodeRequest, runtime *util.RuntimeOptions) (_result *CreateFabricChaincodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFabricChaincodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFabricChaincode"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFabricChaincode(request *CreateFabricChaincodeRequest) (_result *CreateFabricChaincodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFabricChaincodeResponse{}
	_body, _err := client.CreateFabricChaincodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFabricChannelWithOptions(request *CreateFabricChannelRequest, runtime *util.RuntimeOptions) (_result *CreateFabricChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFabricChannelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFabricChannel"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFabricChannel(request *CreateFabricChannelRequest) (_result *CreateFabricChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFabricChannelResponse{}
	_body, _err := client.CreateFabricChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFabricChannelMemberWithOptions(request *CreateFabricChannelMemberRequest, runtime *util.RuntimeOptions) (_result *CreateFabricChannelMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFabricChannelMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFabricChannelMember"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFabricChannelMember(request *CreateFabricChannelMemberRequest) (_result *CreateFabricChannelMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFabricChannelMemberResponse{}
	_body, _err := client.CreateFabricChannelMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFabricConsortiumWithOptions(request *CreateFabricConsortiumRequest, runtime *util.RuntimeOptions) (_result *CreateFabricConsortiumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFabricConsortiumResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFabricConsortium"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFabricConsortium(request *CreateFabricConsortiumRequest) (_result *CreateFabricConsortiumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFabricConsortiumResponse{}
	_body, _err := client.CreateFabricConsortiumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFabricConsortiumMemberWithOptions(request *CreateFabricConsortiumMemberRequest, runtime *util.RuntimeOptions) (_result *CreateFabricConsortiumMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFabricConsortiumMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFabricConsortiumMember"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFabricConsortiumMember(request *CreateFabricConsortiumMemberRequest) (_result *CreateFabricConsortiumMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFabricConsortiumMemberResponse{}
	_body, _err := client.CreateFabricConsortiumMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFabricOrganizationWithOptions(request *CreateFabricOrganizationRequest, runtime *util.RuntimeOptions) (_result *CreateFabricOrganizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFabricOrganizationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFabricOrganization"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFabricOrganization(request *CreateFabricOrganizationRequest) (_result *CreateFabricOrganizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFabricOrganizationResponse{}
	_body, _err := client.CreateFabricOrganizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFabricOrganizationUserWithOptions(request *CreateFabricOrganizationUserRequest, runtime *util.RuntimeOptions) (_result *CreateFabricOrganizationUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFabricOrganizationUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFabricOrganizationUser"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFabricOrganizationUser(request *CreateFabricOrganizationUserRequest) (_result *CreateFabricOrganizationUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFabricOrganizationUserResponse{}
	_body, _err := client.CreateFabricOrganizationUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAntChainConsortiumWithOptions(request *DeleteAntChainConsortiumRequest, runtime *util.RuntimeOptions) (_result *DeleteAntChainConsortiumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAntChainConsortiumResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAntChainConsortium"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAntChainConsortium(request *DeleteAntChainConsortiumRequest) (_result *DeleteAntChainConsortiumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAntChainConsortiumResponse{}
	_body, _err := client.DeleteAntChainConsortiumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAntChainContractContentWithOptions(request *DeleteAntChainContractContentRequest, runtime *util.RuntimeOptions) (_result *DeleteAntChainContractContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAntChainContractContentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAntChainContractContent"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAntChainContractContent(request *DeleteAntChainContractContentRequest) (_result *DeleteAntChainContractContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAntChainContractContentResponse{}
	_body, _err := client.DeleteAntChainContractContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAntChainContractProjectWithOptions(request *DeleteAntChainContractProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteAntChainContractProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAntChainContractProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAntChainContractProject"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAntChainContractProject(request *DeleteAntChainContractProjectRequest) (_result *DeleteAntChainContractProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAntChainContractProjectResponse{}
	_body, _err := client.DeleteAntChainContractProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAntChainMiniAppQRCodeAuthorizedUserWithOptions(request *DeleteAntChainMiniAppQRCodeAuthorizedUserRequest, runtime *util.RuntimeOptions) (_result *DeleteAntChainMiniAppQRCodeAuthorizedUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAntChainMiniAppQRCodeAuthorizedUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAntChainMiniAppQRCodeAuthorizedUser"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAntChainMiniAppQRCodeAuthorizedUser(request *DeleteAntChainMiniAppQRCodeAuthorizedUserRequest) (_result *DeleteAntChainMiniAppQRCodeAuthorizedUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAntChainMiniAppQRCodeAuthorizedUserResponse{}
	_body, _err := client.DeleteAntChainMiniAppQRCodeAuthorizedUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFabricChaincodeWithOptions(request *DeleteFabricChaincodeRequest, runtime *util.RuntimeOptions) (_result *DeleteFabricChaincodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFabricChaincodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFabricChaincode"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFabricChaincode(request *DeleteFabricChaincodeRequest) (_result *DeleteFabricChaincodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFabricChaincodeResponse{}
	_body, _err := client.DeleteFabricChaincodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainAccountsWithOptions(request *DescribeAntChainAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainAccountsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainAccounts"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainAccounts(request *DescribeAntChainAccountsRequest) (_result *DescribeAntChainAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainAccountsResponse{}
	_body, _err := client.DescribeAntChainAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainBlockWithOptions(request *DescribeAntChainBlockRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainBlockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainBlockResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainBlock"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainBlock(request *DescribeAntChainBlockRequest) (_result *DescribeAntChainBlockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainBlockResponse{}
	_body, _err := client.DescribeAntChainBlockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainCertificateApplicationsWithOptions(request *DescribeAntChainCertificateApplicationsRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainCertificateApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainCertificateApplicationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainCertificateApplications"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainCertificateApplications(request *DescribeAntChainCertificateApplicationsRequest) (_result *DescribeAntChainCertificateApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainCertificateApplicationsResponse{}
	_body, _err := client.DescribeAntChainCertificateApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainConsortiumsWithOptions(request *DescribeAntChainConsortiumsRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainConsortiumsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainConsortiumsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainConsortiums"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainConsortiums(request *DescribeAntChainConsortiumsRequest) (_result *DescribeAntChainConsortiumsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainConsortiumsResponse{}
	_body, _err := client.DescribeAntChainConsortiumsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainContractProjectContentTreeWithOptions(request *DescribeAntChainContractProjectContentTreeRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainContractProjectContentTreeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainContractProjectContentTreeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainContractProjectContentTree"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainContractProjectContentTree(request *DescribeAntChainContractProjectContentTreeRequest) (_result *DescribeAntChainContractProjectContentTreeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainContractProjectContentTreeResponse{}
	_body, _err := client.DescribeAntChainContractProjectContentTreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainContractProjectsWithOptions(request *DescribeAntChainContractProjectsRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainContractProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainContractProjectsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainContractProjects"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainContractProjects(request *DescribeAntChainContractProjectsRequest) (_result *DescribeAntChainContractProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainContractProjectsResponse{}
	_body, _err := client.DescribeAntChainContractProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainDownloadPathsWithOptions(request *DescribeAntChainDownloadPathsRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainDownloadPathsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainDownloadPathsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainDownloadPaths"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainDownloadPaths(request *DescribeAntChainDownloadPathsRequest) (_result *DescribeAntChainDownloadPathsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainDownloadPathsResponse{}
	_body, _err := client.DescribeAntChainDownloadPathsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainInformationWithOptions(request *DescribeAntChainInformationRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainInformationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainInformation"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainInformation(request *DescribeAntChainInformationRequest) (_result *DescribeAntChainInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainInformationResponse{}
	_body, _err := client.DescribeAntChainInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainLatestBlocksWithOptions(request *DescribeAntChainLatestBlocksRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainLatestBlocksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainLatestBlocksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainLatestBlocks"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainLatestBlocks(request *DescribeAntChainLatestBlocksRequest) (_result *DescribeAntChainLatestBlocksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainLatestBlocksResponse{}
	_body, _err := client.DescribeAntChainLatestBlocksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainLatestTransactionDigestsWithOptions(request *DescribeAntChainLatestTransactionDigestsRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainLatestTransactionDigestsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainLatestTransactionDigestsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainLatestTransactionDigests"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainLatestTransactionDigests(request *DescribeAntChainLatestTransactionDigestsRequest) (_result *DescribeAntChainLatestTransactionDigestsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainLatestTransactionDigestsResponse{}
	_body, _err := client.DescribeAntChainLatestTransactionDigestsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainMembersWithOptions(request *DescribeAntChainMembersRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainMembers"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainMembers(request *DescribeAntChainMembersRequest) (_result *DescribeAntChainMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainMembersResponse{}
	_body, _err := client.DescribeAntChainMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainMiniAppBrowserQRCodeAccessLogWithOptions(request *DescribeAntChainMiniAppBrowserQRCodeAccessLogRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainMiniAppBrowserQRCodeAccessLog"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainMiniAppBrowserQRCodeAccessLog(request *DescribeAntChainMiniAppBrowserQRCodeAccessLogRequest) (_result *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse{}
	_body, _err := client.DescribeAntChainMiniAppBrowserQRCodeAccessLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersWithOptions(request *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsers"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsers(request *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest) (_result *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse{}
	_body, _err := client.DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainMiniAppBrowserTransactionQRCodeWithOptions(request *DescribeAntChainMiniAppBrowserTransactionQRCodeRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainMiniAppBrowserTransactionQRCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainMiniAppBrowserTransactionQRCodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainMiniAppBrowserTransactionQRCode"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainMiniAppBrowserTransactionQRCode(request *DescribeAntChainMiniAppBrowserTransactionQRCodeRequest) (_result *DescribeAntChainMiniAppBrowserTransactionQRCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainMiniAppBrowserTransactionQRCodeResponse{}
	_body, _err := client.DescribeAntChainMiniAppBrowserTransactionQRCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainNodesWithOptions(request *DescribeAntChainNodesRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainNodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainNodes"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainNodes(request *DescribeAntChainNodesRequest) (_result *DescribeAntChainNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainNodesResponse{}
	_body, _err := client.DescribeAntChainNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainQRCodeAuthorizationWithOptions(request *DescribeAntChainQRCodeAuthorizationRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainQRCodeAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainQRCodeAuthorizationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainQRCodeAuthorization"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainQRCodeAuthorization(request *DescribeAntChainQRCodeAuthorizationRequest) (_result *DescribeAntChainQRCodeAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainQRCodeAuthorizationResponse{}
	_body, _err := client.DescribeAntChainQRCodeAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainsWithOptions(request *DescribeAntChainsRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChains"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChains(request *DescribeAntChainsRequest) (_result *DescribeAntChainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainsResponse{}
	_body, _err := client.DescribeAntChainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainTransactionWithOptions(request *DescribeAntChainTransactionRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainTransactionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainTransactionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainTransaction"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainTransaction(request *DescribeAntChainTransactionRequest) (_result *DescribeAntChainTransactionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainTransactionResponse{}
	_body, _err := client.DescribeAntChainTransactionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainTransactionReceiptWithOptions(request *DescribeAntChainTransactionReceiptRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainTransactionReceiptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainTransactionReceiptResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainTransactionReceipt"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainTransactionReceipt(request *DescribeAntChainTransactionReceiptRequest) (_result *DescribeAntChainTransactionReceiptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainTransactionReceiptResponse{}
	_body, _err := client.DescribeAntChainTransactionReceiptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntChainTransactionStatisticsWithOptions(request *DescribeAntChainTransactionStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainTransactionStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntChainTransactionStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntChainTransactionStatistics"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntChainTransactionStatistics(request *DescribeAntChainTransactionStatisticsRequest) (_result *DescribeAntChainTransactionStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainTransactionStatisticsResponse{}
	_body, _err := client.DescribeAntChainTransactionStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEthereumDeletableWithOptions(request *DescribeEthereumDeletableRequest, runtime *util.RuntimeOptions) (_result *DescribeEthereumDeletableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEthereumDeletableResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEthereumDeletable"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEthereumDeletable(request *DescribeEthereumDeletableRequest) (_result *DescribeEthereumDeletableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEthereumDeletableResponse{}
	_body, _err := client.DescribeEthereumDeletableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricCandidateOrganizationsWithOptions(request *DescribeFabricCandidateOrganizationsRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricCandidateOrganizationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricCandidateOrganizationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricCandidateOrganizations"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricCandidateOrganizations(request *DescribeFabricCandidateOrganizationsRequest) (_result *DescribeFabricCandidateOrganizationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricCandidateOrganizationsResponse{}
	_body, _err := client.DescribeFabricCandidateOrganizationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricChaincodeUploadPolicyWithOptions(request *DescribeFabricChaincodeUploadPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricChaincodeUploadPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricChaincodeUploadPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricChaincodeUploadPolicy"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricChaincodeUploadPolicy(request *DescribeFabricChaincodeUploadPolicyRequest) (_result *DescribeFabricChaincodeUploadPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricChaincodeUploadPolicyResponse{}
	_body, _err := client.DescribeFabricChaincodeUploadPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricChannelMembersWithOptions(request *DescribeFabricChannelMembersRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricChannelMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricChannelMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricChannelMembers"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricChannelMembers(request *DescribeFabricChannelMembersRequest) (_result *DescribeFabricChannelMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricChannelMembersResponse{}
	_body, _err := client.DescribeFabricChannelMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumAdminStatusWithOptions(request *DescribeFabricConsortiumAdminStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumAdminStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricConsortiumAdminStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricConsortiumAdminStatus"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumAdminStatus(request *DescribeFabricConsortiumAdminStatusRequest) (_result *DescribeFabricConsortiumAdminStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricConsortiumAdminStatusResponse{}
	_body, _err := client.DescribeFabricConsortiumAdminStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumChaincodesWithOptions(request *DescribeFabricConsortiumChaincodesRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumChaincodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricConsortiumChaincodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricConsortiumChaincodes"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumChaincodes(request *DescribeFabricConsortiumChaincodesRequest) (_result *DescribeFabricConsortiumChaincodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricConsortiumChaincodesResponse{}
	_body, _err := client.DescribeFabricConsortiumChaincodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumChannelsWithOptions(request *DescribeFabricConsortiumChannelsRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumChannelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricConsortiumChannelsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricConsortiumChannels"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumChannels(request *DescribeFabricConsortiumChannelsRequest) (_result *DescribeFabricConsortiumChannelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricConsortiumChannelsResponse{}
	_body, _err := client.DescribeFabricConsortiumChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumConfigWithOptions(runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeFabricConsortiumConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricConsortiumConfig"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumConfig() (_result *DescribeFabricConsortiumConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricConsortiumConfigResponse{}
	_body, _err := client.DescribeFabricConsortiumConfigWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumDeletableWithOptions(request *DescribeFabricConsortiumDeletableRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumDeletableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricConsortiumDeletableResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricConsortiumDeletable"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumDeletable(request *DescribeFabricConsortiumDeletableRequest) (_result *DescribeFabricConsortiumDeletableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricConsortiumDeletableResponse{}
	_body, _err := client.DescribeFabricConsortiumDeletableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumMemberApprovalWithOptions(request *DescribeFabricConsortiumMemberApprovalRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumMemberApprovalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricConsortiumMemberApprovalResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricConsortiumMemberApproval"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumMemberApproval(request *DescribeFabricConsortiumMemberApprovalRequest) (_result *DescribeFabricConsortiumMemberApprovalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricConsortiumMemberApprovalResponse{}
	_body, _err := client.DescribeFabricConsortiumMemberApprovalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumMembersWithOptions(request *DescribeFabricConsortiumMembersRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricConsortiumMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricConsortiumMembers"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumMembers(request *DescribeFabricConsortiumMembersRequest) (_result *DescribeFabricConsortiumMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricConsortiumMembersResponse{}
	_body, _err := client.DescribeFabricConsortiumMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumOrderersWithOptions(request *DescribeFabricConsortiumOrderersRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumOrderersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricConsortiumOrderersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricConsortiumOrderers"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumOrderers(request *DescribeFabricConsortiumOrderersRequest) (_result *DescribeFabricConsortiumOrderersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricConsortiumOrderersResponse{}
	_body, _err := client.DescribeFabricConsortiumOrderersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumsWithOptions(request *DescribeFabricConsortiumsRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricConsortiumsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricConsortiums"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricConsortiums(request *DescribeFabricConsortiumsRequest) (_result *DescribeFabricConsortiumsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricConsortiumsResponse{}
	_body, _err := client.DescribeFabricConsortiumsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumSpecsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumSpecsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeFabricConsortiumSpecsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricConsortiumSpecs"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricConsortiumSpecs() (_result *DescribeFabricConsortiumSpecsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricConsortiumSpecsResponse{}
	_body, _err := client.DescribeFabricConsortiumSpecsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricExplorerWithOptions(request *DescribeFabricExplorerRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricExplorerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricExplorerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricExplorer"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricExplorer(request *DescribeFabricExplorerRequest) (_result *DescribeFabricExplorerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricExplorerResponse{}
	_body, _err := client.DescribeFabricExplorerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricInvitationCodeWithOptions(request *DescribeFabricInvitationCodeRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricInvitationCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricInvitationCodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricInvitationCode"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricInvitationCode(request *DescribeFabricInvitationCodeRequest) (_result *DescribeFabricInvitationCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricInvitationCodeResponse{}
	_body, _err := client.DescribeFabricInvitationCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricInviterWithOptions(request *DescribeFabricInviterRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricInviterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricInviterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricInviter"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricInviter(request *DescribeFabricInviterRequest) (_result *DescribeFabricInviterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricInviterResponse{}
	_body, _err := client.DescribeFabricInviterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricOrdererLogsWithOptions(request *DescribeFabricOrdererLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrdererLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricOrdererLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricOrdererLogs"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricOrdererLogs(request *DescribeFabricOrdererLogsRequest) (_result *DescribeFabricOrdererLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricOrdererLogsResponse{}
	_body, _err := client.DescribeFabricOrdererLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricOrganizationWithOptions(request *DescribeFabricOrganizationRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricOrganizationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricOrganization"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricOrganization(request *DescribeFabricOrganizationRequest) (_result *DescribeFabricOrganizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricOrganizationResponse{}
	_body, _err := client.DescribeFabricOrganizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricOrganizationChaincodesWithOptions(request *DescribeFabricOrganizationChaincodesRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationChaincodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricOrganizationChaincodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricOrganizationChaincodes"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricOrganizationChaincodes(request *DescribeFabricOrganizationChaincodesRequest) (_result *DescribeFabricOrganizationChaincodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricOrganizationChaincodesResponse{}
	_body, _err := client.DescribeFabricOrganizationChaincodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricOrganizationDeletableWithOptions(request *DescribeFabricOrganizationDeletableRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationDeletableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricOrganizationDeletableResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricOrganizationDeletable"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricOrganizationDeletable(request *DescribeFabricOrganizationDeletableRequest) (_result *DescribeFabricOrganizationDeletableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricOrganizationDeletableResponse{}
	_body, _err := client.DescribeFabricOrganizationDeletableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricOrganizationMembersWithOptions(request *DescribeFabricOrganizationMembersRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricOrganizationMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricOrganizationMembers"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricOrganizationMembers(request *DescribeFabricOrganizationMembersRequest) (_result *DescribeFabricOrganizationMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricOrganizationMembersResponse{}
	_body, _err := client.DescribeFabricOrganizationMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricOrganizationPeersWithOptions(request *DescribeFabricOrganizationPeersRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationPeersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricOrganizationPeersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricOrganizationPeers"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricOrganizationPeers(request *DescribeFabricOrganizationPeersRequest) (_result *DescribeFabricOrganizationPeersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricOrganizationPeersResponse{}
	_body, _err := client.DescribeFabricOrganizationPeersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricOrganizationsWithOptions(request *DescribeFabricOrganizationsRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricOrganizationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricOrganizations"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricOrganizations(request *DescribeFabricOrganizationsRequest) (_result *DescribeFabricOrganizationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricOrganizationsResponse{}
	_body, _err := client.DescribeFabricOrganizationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricOrganizationSpecsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationSpecsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeFabricOrganizationSpecsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricOrganizationSpecs"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricOrganizationSpecs() (_result *DescribeFabricOrganizationSpecsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricOrganizationSpecsResponse{}
	_body, _err := client.DescribeFabricOrganizationSpecsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricOrganizationUsersWithOptions(request *DescribeFabricOrganizationUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricOrganizationUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricOrganizationUsers"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricOrganizationUsers(request *DescribeFabricOrganizationUsersRequest) (_result *DescribeFabricOrganizationUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricOrganizationUsersResponse{}
	_body, _err := client.DescribeFabricOrganizationUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFabricPeerLogsWithOptions(request *DescribeFabricPeerLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricPeerLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFabricPeerLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFabricPeerLogs"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFabricPeerLogs(request *DescribeFabricPeerLogsRequest) (_result *DescribeFabricPeerLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricPeerLogsResponse{}
	_body, _err := client.DescribeFabricPeerLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRootDomainWithOptions(runtime *util.RuntimeOptions) (_result *DescribeRootDomainResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeRootDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRootDomain"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRootDomain() (_result *DescribeRootDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRootDomainResponse{}
	_body, _err := client.DescribeRootDomainWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTasksWithOptions(runtime *util.RuntimeOptions) (_result *DescribeTasksResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTasks"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTasks() (_result *DescribeTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTasksResponse{}
	_body, _err := client.DescribeTasksWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadFabricOrganizationSDKWithOptions(request *DownloadFabricOrganizationSDKRequest, runtime *util.RuntimeOptions) (_result *DownloadFabricOrganizationSDKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DownloadFabricOrganizationSDKResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DownloadFabricOrganizationSDK"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadFabricOrganizationSDK(request *DownloadFabricOrganizationSDKRequest) (_result *DownloadFabricOrganizationSDKResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadFabricOrganizationSDKResponse{}
	_body, _err := client.DownloadFabricOrganizationSDKWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FreezeAntChainAccountWithOptions(request *FreezeAntChainAccountRequest, runtime *util.RuntimeOptions) (_result *FreezeAntChainAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FreezeAntChainAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FreezeAntChainAccount"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FreezeAntChainAccount(request *FreezeAntChainAccountRequest) (_result *FreezeAntChainAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FreezeAntChainAccountResponse{}
	_body, _err := client.FreezeAntChainAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallFabricChaincodeWithOptions(request *InstallFabricChaincodeRequest, runtime *util.RuntimeOptions) (_result *InstallFabricChaincodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InstallFabricChaincodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InstallFabricChaincode"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallFabricChaincode(request *InstallFabricChaincodeRequest) (_result *InstallFabricChaincodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallFabricChaincodeResponse{}
	_body, _err := client.InstallFabricChaincodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstantiateFabricChaincodeWithOptions(request *InstantiateFabricChaincodeRequest, runtime *util.RuntimeOptions) (_result *InstantiateFabricChaincodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InstantiateFabricChaincodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InstantiateFabricChaincode"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstantiateFabricChaincode(request *InstantiateFabricChaincodeRequest) (_result *InstantiateFabricChaincodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstantiateFabricChaincodeResponse{}
	_body, _err := client.InstantiateFabricChaincodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinFabricChannelWithOptions(request *JoinFabricChannelRequest, runtime *util.RuntimeOptions) (_result *JoinFabricChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &JoinFabricChannelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("JoinFabricChannel"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinFabricChannel(request *JoinFabricChannelRequest) (_result *JoinFabricChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinFabricChannelResponse{}
	_body, _err := client.JoinFabricChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetAntChainCertificateWithOptions(request *ResetAntChainCertificateRequest, runtime *util.RuntimeOptions) (_result *ResetAntChainCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetAntChainCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetAntChainCertificate"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetAntChainCertificate(request *ResetAntChainCertificateRequest) (_result *ResetAntChainCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAntChainCertificateResponse{}
	_body, _err := client.ResetAntChainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetAntChainUserCertificateWithOptions(request *ResetAntChainUserCertificateRequest, runtime *util.RuntimeOptions) (_result *ResetAntChainUserCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetAntChainUserCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetAntChainUserCertificate"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetAntChainUserCertificate(request *ResetAntChainUserCertificateRequest) (_result *ResetAntChainUserCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAntChainUserCertificateResponse{}
	_body, _err := client.ResetAntChainUserCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetFabricOrganizationUserPasswordWithOptions(request *ResetFabricOrganizationUserPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetFabricOrganizationUserPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetFabricOrganizationUserPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetFabricOrganizationUserPassword"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetFabricOrganizationUserPassword(request *ResetFabricOrganizationUserPasswordRequest) (_result *ResetFabricOrganizationUserPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetFabricOrganizationUserPasswordResponse{}
	_body, _err := client.ResetFabricOrganizationUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SynchronizeFabricChaincodeWithOptions(request *SynchronizeFabricChaincodeRequest, runtime *util.RuntimeOptions) (_result *SynchronizeFabricChaincodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SynchronizeFabricChaincodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SynchronizeFabricChaincode"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SynchronizeFabricChaincode(request *SynchronizeFabricChaincodeRequest) (_result *SynchronizeFabricChaincodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SynchronizeFabricChaincodeResponse{}
	_body, _err := client.SynchronizeFabricChaincodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnfreezeAntChainAccountWithOptions(request *UnfreezeAntChainAccountRequest, runtime *util.RuntimeOptions) (_result *UnfreezeAntChainAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnfreezeAntChainAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnfreezeAntChainAccount"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnfreezeAntChainAccount(request *UnfreezeAntChainAccountRequest) (_result *UnfreezeAntChainAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnfreezeAntChainAccountResponse{}
	_body, _err := client.UnfreezeAntChainAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAntChainWithOptions(request *UpdateAntChainRequest, runtime *util.RuntimeOptions) (_result *UpdateAntChainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAntChainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAntChain"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAntChain(request *UpdateAntChainRequest) (_result *UpdateAntChainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAntChainResponse{}
	_body, _err := client.UpdateAntChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAntChainConsortiumWithOptions(request *UpdateAntChainConsortiumRequest, runtime *util.RuntimeOptions) (_result *UpdateAntChainConsortiumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAntChainConsortiumResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAntChainConsortium"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAntChainConsortium(request *UpdateAntChainConsortiumRequest) (_result *UpdateAntChainConsortiumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAntChainConsortiumResponse{}
	_body, _err := client.UpdateAntChainConsortiumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAntChainContractContentWithOptions(request *UpdateAntChainContractContentRequest, runtime *util.RuntimeOptions) (_result *UpdateAntChainContractContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAntChainContractContentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAntChainContractContent"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAntChainContractContent(request *UpdateAntChainContractContentRequest) (_result *UpdateAntChainContractContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAntChainContractContentResponse{}
	_body, _err := client.UpdateAntChainContractContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAntChainContractProjectWithOptions(request *UpdateAntChainContractProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateAntChainContractProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAntChainContractProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAntChainContractProject"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAntChainContractProject(request *UpdateAntChainContractProjectRequest) (_result *UpdateAntChainContractProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAntChainContractProjectResponse{}
	_body, _err := client.UpdateAntChainContractProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAntChainMemberWithOptions(request *UpdateAntChainMemberRequest, runtime *util.RuntimeOptions) (_result *UpdateAntChainMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAntChainMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAntChainMember"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAntChainMember(request *UpdateAntChainMemberRequest) (_result *UpdateAntChainMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAntChainMemberResponse{}
	_body, _err := client.UpdateAntChainMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAntChainQRCodeAuthorizationWithOptions(request *UpdateAntChainQRCodeAuthorizationRequest, runtime *util.RuntimeOptions) (_result *UpdateAntChainQRCodeAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAntChainQRCodeAuthorizationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAntChainQRCodeAuthorization"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAntChainQRCodeAuthorization(request *UpdateAntChainQRCodeAuthorizationRequest) (_result *UpdateAntChainQRCodeAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAntChainQRCodeAuthorizationResponse{}
	_body, _err := client.UpdateAntChainQRCodeAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeFabricChaincodeWithOptions(request *UpgradeFabricChaincodeRequest, runtime *util.RuntimeOptions) (_result *UpgradeFabricChaincodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeFabricChaincodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeFabricChaincode"), tea.String("2018-12-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeFabricChaincode(request *UpgradeFabricChaincodeRequest) (_result *UpgradeFabricChaincodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeFabricChaincodeResponse{}
	_body, _err := client.UpgradeFabricChaincodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
