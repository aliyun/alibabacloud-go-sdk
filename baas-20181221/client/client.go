// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ChaincodePackage struct {
	// example:
	//
	// cmcc-1.5:227c3cd472a8a2fcaa1c63f15812a1894348a985dc723006dd47d9f2848d10ed
	ChaincodePackageId *string `json:"ChaincodePackageId,omitempty" xml:"ChaincodePackageId,omitempty"`
	// example:
	//
	// 15bab12c29edbc095c33545b9867aec9be469c74a6dd5783df6ee74a24334c77
	Checksum *string `json:"Checksum,omitempty" xml:"Checksum,omitempty"`
	// example:
	//
	// 1533025590
	DeleteTime *string `json:"DeleteTime,omitempty" xml:"DeleteTime,omitempty"`
	// example:
	//
	// false
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// example:
	//
	// 1533025590
	InstallTime *string `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	// example:
	//
	// label
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// N02eVKuRJDA2ODtieMUxVw==
	Md5sum *string `json:"Md5sum,omitempty" xml:"Md5sum,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// hi003/7efbe860-1a40-42c9-9e6b-0bc08d7ee3a4/cmcc.tar.gz
	OssURL *string `json:"OssURL,omitempty" xml:"OssURL,omitempty"`
	// example:
	//
	// 26842
	ProviderBid *string `json:"ProviderBid,omitempty" xml:"ProviderBid,omitempty"`
	// example:
	//
	// 1313158505072210
	ProviderUid *string `json:"ProviderUid,omitempty" xml:"ProviderUid,omitempty"`
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// golang
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	// example:
	//
	// 1533025590
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s ChaincodePackage) String() string {
	return tea.Prettify(s)
}

func (s ChaincodePackage) GoString() string {
	return s.String()
}

func (s *ChaincodePackage) SetChaincodePackageId(v string) *ChaincodePackage {
	s.ChaincodePackageId = &v
	return s
}

func (s *ChaincodePackage) SetChecksum(v string) *ChaincodePackage {
	s.Checksum = &v
	return s
}

func (s *ChaincodePackage) SetDeleteTime(v string) *ChaincodePackage {
	s.DeleteTime = &v
	return s
}

func (s *ChaincodePackage) SetDeleted(v bool) *ChaincodePackage {
	s.Deleted = &v
	return s
}

func (s *ChaincodePackage) SetInstallTime(v string) *ChaincodePackage {
	s.InstallTime = &v
	return s
}

func (s *ChaincodePackage) SetLabel(v string) *ChaincodePackage {
	s.Label = &v
	return s
}

func (s *ChaincodePackage) SetMd5sum(v string) *ChaincodePackage {
	s.Md5sum = &v
	return s
}

func (s *ChaincodePackage) SetMessage(v string) *ChaincodePackage {
	s.Message = &v
	return s
}

func (s *ChaincodePackage) SetOrganizationId(v string) *ChaincodePackage {
	s.OrganizationId = &v
	return s
}

func (s *ChaincodePackage) SetOssURL(v string) *ChaincodePackage {
	s.OssURL = &v
	return s
}

func (s *ChaincodePackage) SetProviderBid(v string) *ChaincodePackage {
	s.ProviderBid = &v
	return s
}

func (s *ChaincodePackage) SetProviderUid(v string) *ChaincodePackage {
	s.ProviderUid = &v
	return s
}

func (s *ChaincodePackage) SetState(v string) *ChaincodePackage {
	s.State = &v
	return s
}

func (s *ChaincodePackage) SetType(v int32) *ChaincodePackage {
	s.Type = &v
	return s
}

func (s *ChaincodePackage) SetTypeName(v string) *ChaincodePackage {
	s.TypeName = &v
	return s
}

func (s *ChaincodePackage) SetUploadTime(v string) *ChaincodePackage {
	s.UploadTime = &v
	return s
}

type ChaincodeVO struct {
	// example:
	//
	// a5f04baed775b016e8dc386783f40a2c8b13823cac22bc5fd11846cca4a13a9a
	ChaincodeDefinitionId *string `json:"ChaincodeDefinitionId,omitempty" xml:"ChaincodeDefinitionId,omitempty"`
	// example:
	//
	// code-sacc-jbkuldd3i7ad
	ChaincodeId *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	// example:
	//
	// sacc_v1.0:698d4e73c5f1d0239fbf0313ac79ddfaf44d5a803ecb02a4d1983cca4a706860
	ChaincodePackageId *string `json:"ChaincodePackageId,omitempty" xml:"ChaincodePackageId,omitempty"`
	// example:
	//
	// chan-test-b0wx5suhgyb9p
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// f100
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1533025590
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1533025590
	DeployTime *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	// example:
	//
	// OR ("perf9141MSP.member")
	EndorsePolicy *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	// example:
	//
	// false
	InitRequired *bool `json:"InitRequired,omitempty" xml:"InitRequired,omitempty"`
	// example:
	//
	// args:"key" args:"val"
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// true
	Install *bool `json:"Install,omitempty" xml:"Install,omitempty"`
	// example:
	//
	// false
	Management *bool `json:"Management,omitempty" xml:"Management,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// sacc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// chaincode/sacc
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// example:
	//
	// perf9141
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// example:
	//
	// Pending
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ChaincodeVO) String() string {
	return tea.Prettify(s)
}

func (s ChaincodeVO) GoString() string {
	return s.String()
}

func (s *ChaincodeVO) SetChaincodeDefinitionId(v string) *ChaincodeVO {
	s.ChaincodeDefinitionId = &v
	return s
}

func (s *ChaincodeVO) SetChaincodeId(v string) *ChaincodeVO {
	s.ChaincodeId = &v
	return s
}

func (s *ChaincodeVO) SetChaincodePackageId(v string) *ChaincodeVO {
	s.ChaincodePackageId = &v
	return s
}

func (s *ChaincodeVO) SetChannelId(v string) *ChaincodeVO {
	s.ChannelId = &v
	return s
}

func (s *ChaincodeVO) SetChannelName(v string) *ChaincodeVO {
	s.ChannelName = &v
	return s
}

func (s *ChaincodeVO) SetConsortiumId(v string) *ChaincodeVO {
	s.ConsortiumId = &v
	return s
}

func (s *ChaincodeVO) SetCreateTime(v string) *ChaincodeVO {
	s.CreateTime = &v
	return s
}

func (s *ChaincodeVO) SetDeployTime(v string) *ChaincodeVO {
	s.DeployTime = &v
	return s
}

func (s *ChaincodeVO) SetEndorsePolicy(v string) *ChaincodeVO {
	s.EndorsePolicy = &v
	return s
}

func (s *ChaincodeVO) SetInitRequired(v bool) *ChaincodeVO {
	s.InitRequired = &v
	return s
}

func (s *ChaincodeVO) SetInput(v string) *ChaincodeVO {
	s.Input = &v
	return s
}

func (s *ChaincodeVO) SetInstall(v bool) *ChaincodeVO {
	s.Install = &v
	return s
}

func (s *ChaincodeVO) SetManagement(v bool) *ChaincodeVO {
	s.Management = &v
	return s
}

func (s *ChaincodeVO) SetMessage(v string) *ChaincodeVO {
	s.Message = &v
	return s
}

func (s *ChaincodeVO) SetName(v string) *ChaincodeVO {
	s.Name = &v
	return s
}

func (s *ChaincodeVO) SetPath(v string) *ChaincodeVO {
	s.Path = &v
	return s
}

func (s *ChaincodeVO) SetProviderId(v string) *ChaincodeVO {
	s.ProviderId = &v
	return s
}

func (s *ChaincodeVO) SetProviderName(v string) *ChaincodeVO {
	s.ProviderName = &v
	return s
}

func (s *ChaincodeVO) SetState(v string) *ChaincodeVO {
	s.State = &v
	return s
}

func (s *ChaincodeVO) SetType(v int32) *ChaincodeVO {
	s.Type = &v
	return s
}

func (s *ChaincodeVO) SetVersion(v string) *ChaincodeVO {
	s.Version = &v
	return s
}

type AcceptFabricInvitationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// secrettoken
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	IsAccepted *bool `json:"IsAccepted,omitempty" xml:"IsAccepted,omitempty"`
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 52CC8D5B-79A8-4769-BC7B-8940556ADA49
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AcceptFabricInvitationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AcceptFabricInvitationResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptFabricInvitationResponseBody) SetErrorCode(v int32) *AcceptFabricInvitationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AcceptFabricInvitationResponseBody) SetRequestId(v string) *AcceptFabricInvitationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcceptFabricInvitationResponseBody) SetSuccess(v bool) *AcceptFabricInvitationResponseBody {
	s.Success = &v
	return s
}

type AcceptFabricInvitationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptFabricInvitationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AcceptFabricInvitationResponse) SetStatusCode(v int32) *AcceptFabricInvitationResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptFabricInvitationResponse) SetBody(v *AcceptFabricInvitationResponseBody) *AcceptFabricInvitationResponse {
	s.Body = v
	return s
}

type ApplyAntChainCertificateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bDXK6boZ
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// LS0tLS1...
	UploadReq *string `json:"UploadReq,omitempty" xml:"UploadReq,omitempty"`
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
	// example:
	//
	// D3FB9E67-0E31-4B8B-8895-3660CCE8CA62
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyAntChainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ApplyAntChainCertificateResponse) SetStatusCode(v int32) *ApplyAntChainCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyAntChainCertificateResponse) SetBody(v *ApplyAntChainCertificateResponseBody) *ApplyAntChainCertificateResponse {
	s.Body = v
	return s
}

type ApplyAntChainCertificateWithKeyAutoCreationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1q8B5R9p
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// commonName
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DV80nJXq
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// countryName
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// localityName
	LocalityName *string `json:"LocalityName,omitempty" xml:"LocalityName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// organizationName
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// organizationUnitName
	OrganizationUnitName *string `json:"OrganizationUnitName,omitempty" xml:"OrganizationUnitName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// stateOrProvinceName
	StateOrProvinceName *string `json:"StateOrProvinceName,omitempty" xml:"StateOrProvinceName,omitempty"`
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

func (s *ApplyAntChainCertificateWithKeyAutoCreationRequest) SetCommonName(v string) *ApplyAntChainCertificateWithKeyAutoCreationRequest {
	s.CommonName = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationRequest) SetConsortiumId(v string) *ApplyAntChainCertificateWithKeyAutoCreationRequest {
	s.ConsortiumId = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationRequest) SetCountryName(v string) *ApplyAntChainCertificateWithKeyAutoCreationRequest {
	s.CountryName = &v
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

func (s *ApplyAntChainCertificateWithKeyAutoCreationRequest) SetPassword(v string) *ApplyAntChainCertificateWithKeyAutoCreationRequest {
	s.Password = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationRequest) SetStateOrProvinceName(v string) *ApplyAntChainCertificateWithKeyAutoCreationRequest {
	s.StateOrProvinceName = &v
	return s
}

type ApplyAntChainCertificateWithKeyAutoCreationResponseBody struct {
	// example:
	//
	// 3E1894AF-FEE7-4C6C-8C2A-2B3EE0105B07
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
	DownloadPath *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath `json:"DownloadPath,omitempty" xml:"DownloadPath,omitempty" type:"Struct"`
	// example:
	//
	// -----BEGIN ENCRYPTED PRIVATE KEY-----\\nMII***A=\\n-----END ENCRYPTED PRIVATE KEY-----\\n
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
}

func (s ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResult) SetDownloadPath(v *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath) *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResult {
	s.DownloadPath = v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResult) SetPrivateKey(v string) *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResult {
	s.PrivateKey = &v
	return s
}

type ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath struct {
	// example:
	//
	// http://***ca.crt
	CaCrtUrl *string `json:"CaCrtUrl,omitempty" xml:"CaCrtUrl,omitempty"`
	// example:
	//
	// http://***client.crt
	ClientCrtUrl *string `json:"ClientCrtUrl,omitempty" xml:"ClientCrtUrl,omitempty"`
	// example:
	//
	// http://***
	SdkUrl *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
	// example:
	//
	// http://***trustCa
	TrustCaUrl *string `json:"TrustCaUrl,omitempty" xml:"TrustCaUrl,omitempty"`
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

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath) SetClientCrtUrl(v string) *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath {
	s.ClientCrtUrl = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath) SetSdkUrl(v string) *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath {
	s.SdkUrl = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath) SetTrustCaUrl(v string) *ApplyAntChainCertificateWithKeyAutoCreationResponseBodyResultDownloadPath {
	s.TrustCaUrl = &v
	return s
}

type ApplyAntChainCertificateWithKeyAutoCreationResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyAntChainCertificateWithKeyAutoCreationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponse) SetStatusCode(v int32) *ApplyAntChainCertificateWithKeyAutoCreationResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyAntChainCertificateWithKeyAutoCreationResponse) SetBody(v *ApplyAntChainCertificateWithKeyAutoCreationResponseBody) *ApplyAntChainCertificateWithKeyAutoCreationResponse {
	s.Body = v
	return s
}

type ApproveFabricChaincodeDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// code-sacc-jbkuldd3i7ad
	ChaincodeId *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	// example:
	//
	// sacc_v1.0:698d4e73c5f1d0239fbf0313ac79ddfaf44d5a803ecb02a4d1983cca4a706860
	ChaincodePackageId *string `json:"ChaincodePackageId,omitempty" xml:"ChaincodePackageId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ApproveFabricChaincodeDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveFabricChaincodeDefinitionRequest) GoString() string {
	return s.String()
}

func (s *ApproveFabricChaincodeDefinitionRequest) SetChaincodeId(v string) *ApproveFabricChaincodeDefinitionRequest {
	s.ChaincodeId = &v
	return s
}

func (s *ApproveFabricChaincodeDefinitionRequest) SetChaincodePackageId(v string) *ApproveFabricChaincodeDefinitionRequest {
	s.ChaincodePackageId = &v
	return s
}

func (s *ApproveFabricChaincodeDefinitionRequest) SetLocation(v string) *ApproveFabricChaincodeDefinitionRequest {
	s.Location = &v
	return s
}

func (s *ApproveFabricChaincodeDefinitionRequest) SetOrganizationId(v string) *ApproveFabricChaincodeDefinitionRequest {
	s.OrganizationId = &v
	return s
}

type ApproveFabricChaincodeDefinitionResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ChaincodeVO `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApproveFabricChaincodeDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveFabricChaincodeDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveFabricChaincodeDefinitionResponseBody) SetErrorCode(v int32) *ApproveFabricChaincodeDefinitionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ApproveFabricChaincodeDefinitionResponseBody) SetMessage(v string) *ApproveFabricChaincodeDefinitionResponseBody {
	s.Message = &v
	return s
}

func (s *ApproveFabricChaincodeDefinitionResponseBody) SetRequestId(v string) *ApproveFabricChaincodeDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApproveFabricChaincodeDefinitionResponseBody) SetResult(v *ChaincodeVO) *ApproveFabricChaincodeDefinitionResponseBody {
	s.Result = v
	return s
}

func (s *ApproveFabricChaincodeDefinitionResponseBody) SetSuccess(v bool) *ApproveFabricChaincodeDefinitionResponseBody {
	s.Success = &v
	return s
}

type ApproveFabricChaincodeDefinitionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveFabricChaincodeDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveFabricChaincodeDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveFabricChaincodeDefinitionResponse) GoString() string {
	return s.String()
}

func (s *ApproveFabricChaincodeDefinitionResponse) SetHeaders(v map[string]*string) *ApproveFabricChaincodeDefinitionResponse {
	s.Headers = v
	return s
}

func (s *ApproveFabricChaincodeDefinitionResponse) SetStatusCode(v int32) *ApproveFabricChaincodeDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveFabricChaincodeDefinitionResponse) SetBody(v *ApproveFabricChaincodeDefinitionResponseBody) *ApproveFabricChaincodeDefinitionResponse {
	s.Body = v
	return s
}

type BatchAddAntChainMiniAppQRCodeAuthorizedUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bDXK6boZ
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["1380013****", "1380013****"]
	PhoneList map[string]interface{} `json:"PhoneList,omitempty" xml:"PhoneList,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// bDXK6boZ
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["1380013****", "1380013****"]
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
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse) SetStatusCode(v int32) *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse) SetBody(v *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponseBody) *BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse {
	s.Body = v
	return s
}

type CheckFabricConsortiumDomainRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bank
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 8F80A214-89FC-4348-9B3F-15446B3DC1FA
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CheckFabricConsortiumDomainResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckFabricConsortiumDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckFabricConsortiumDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CheckFabricConsortiumDomainResponseBody) SetErrorCode(v int32) *CheckFabricConsortiumDomainResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckFabricConsortiumDomainResponseBody) SetRequestId(v string) *CheckFabricConsortiumDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckFabricConsortiumDomainResponseBody) SetResult(v *CheckFabricConsortiumDomainResponseBodyResult) *CheckFabricConsortiumDomainResponseBody {
	s.Result = v
	return s
}

func (s *CheckFabricConsortiumDomainResponseBody) SetSuccess(v bool) *CheckFabricConsortiumDomainResponseBody {
	s.Success = &v
	return s
}

type CheckFabricConsortiumDomainResponseBodyResult struct {
	// example:
	//
	// bank
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// OK
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// true
	Valid *bool `json:"Valid,omitempty" xml:"Valid,omitempty"`
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

func (s *CheckFabricConsortiumDomainResponseBodyResult) SetPrompt(v string) *CheckFabricConsortiumDomainResponseBodyResult {
	s.Prompt = &v
	return s
}

func (s *CheckFabricConsortiumDomainResponseBodyResult) SetValid(v bool) *CheckFabricConsortiumDomainResponseBodyResult {
	s.Valid = &v
	return s
}

type CheckFabricConsortiumDomainResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckFabricConsortiumDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CheckFabricConsortiumDomainResponse) SetStatusCode(v int32) *CheckFabricConsortiumDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckFabricConsortiumDomainResponse) SetBody(v *CheckFabricConsortiumDomainResponseBody) *CheckFabricConsortiumDomainResponse {
	s.Body = v
	return s
}

type CheckFabricOrganizationDomainRequest struct {
	// example:
	//
	// bank
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// bank
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
}

func (s CheckFabricOrganizationDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckFabricOrganizationDomainRequest) GoString() string {
	return s.String()
}

func (s *CheckFabricOrganizationDomainRequest) SetDomain(v string) *CheckFabricOrganizationDomainRequest {
	s.Domain = &v
	return s
}

func (s *CheckFabricOrganizationDomainRequest) SetDomainCode(v string) *CheckFabricOrganizationDomainRequest {
	s.DomainCode = &v
	return s
}

type CheckFabricOrganizationDomainResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// C30A14C7-800E-468B-9EB2-C704DA49295E
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CheckFabricOrganizationDomainResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckFabricOrganizationDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckFabricOrganizationDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CheckFabricOrganizationDomainResponseBody) SetErrorCode(v int32) *CheckFabricOrganizationDomainResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckFabricOrganizationDomainResponseBody) SetRequestId(v string) *CheckFabricOrganizationDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckFabricOrganizationDomainResponseBody) SetResult(v *CheckFabricOrganizationDomainResponseBodyResult) *CheckFabricOrganizationDomainResponseBody {
	s.Result = v
	return s
}

func (s *CheckFabricOrganizationDomainResponseBody) SetSuccess(v bool) *CheckFabricOrganizationDomainResponseBody {
	s.Success = &v
	return s
}

type CheckFabricOrganizationDomainResponseBodyResult struct {
	// example:
	//
	// bank
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// OK
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// true
	Valid *bool `json:"Valid,omitempty" xml:"Valid,omitempty"`
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

func (s *CheckFabricOrganizationDomainResponseBodyResult) SetPrompt(v string) *CheckFabricOrganizationDomainResponseBodyResult {
	s.Prompt = &v
	return s
}

func (s *CheckFabricOrganizationDomainResponseBodyResult) SetValid(v bool) *CheckFabricOrganizationDomainResponseBodyResult {
	s.Valid = &v
	return s
}

type CheckFabricOrganizationDomainResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckFabricOrganizationDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CheckFabricOrganizationDomainResponse) SetStatusCode(v int32) *CheckFabricOrganizationDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckFabricOrganizationDomainResponse) SetBody(v *CheckFabricOrganizationDomainResponseBody) *CheckFabricOrganizationDomainResponse {
	s.Body = v
	return s
}

type ConfirmFabricConsortiumMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// consortium-aaaaaa-akpcsjjac2jd
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa1-1oxw31d046jtl
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 50B3ACF3-CE12-433C-A834-9E8C657A4934
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfirmFabricConsortiumMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmFabricConsortiumMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmFabricConsortiumMemberResponseBody) SetErrorCode(v int32) *ConfirmFabricConsortiumMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ConfirmFabricConsortiumMemberResponseBody) SetRequestId(v string) *ConfirmFabricConsortiumMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmFabricConsortiumMemberResponseBody) SetResult(v bool) *ConfirmFabricConsortiumMemberResponseBody {
	s.Result = &v
	return s
}

func (s *ConfirmFabricConsortiumMemberResponseBody) SetSuccess(v bool) *ConfirmFabricConsortiumMemberResponseBody {
	s.Success = &v
	return s
}

type ConfirmFabricConsortiumMemberResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmFabricConsortiumMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ConfirmFabricConsortiumMemberResponse) SetStatusCode(v int32) *ConfirmFabricConsortiumMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmFabricConsortiumMemberResponse) SetBody(v *ConfirmFabricConsortiumMemberResponseBody) *ConfirmFabricConsortiumMemberResponse {
	s.Body = v
	return s
}

type CopyAntChainContractProjectRequest struct {
	// example:
	//
	// copy project description
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2L9VK68g
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// copyproject
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1.0.1
	ProjectVersion *string `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
}

func (s CopyAntChainContractProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyAntChainContractProjectRequest) GoString() string {
	return s.String()
}

func (s *CopyAntChainContractProjectRequest) SetProjectDescription(v string) *CopyAntChainContractProjectRequest {
	s.ProjectDescription = &v
	return s
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

type CopyAntChainContractProjectResponseBody struct {
	// example:
	//
	// 45D67F6F-C723-4AD8-8462-F94EE5FF22E6
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
	// example:
	//
	// DV80nJXq
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1563953475248
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// copy project description
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
	// example:
	//
	// R6XMEdXe
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// copyproject
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// v1.0.1
	ProjectVersion *string `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
	// example:
	//
	// 1563953475248
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CopyAntChainContractProjectResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CopyAntChainContractProjectResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CopyAntChainContractProjectResponseBodyResult) SetConsortiumId(v string) *CopyAntChainContractProjectResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *CopyAntChainContractProjectResponseBodyResult) SetCreateTime(v int64) *CopyAntChainContractProjectResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *CopyAntChainContractProjectResponseBodyResult) SetProjectDescription(v string) *CopyAntChainContractProjectResponseBodyResult {
	s.ProjectDescription = &v
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

func (s *CopyAntChainContractProjectResponseBodyResult) SetUpdateTime(v int64) *CopyAntChainContractProjectResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type CopyAntChainContractProjectResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyAntChainContractProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CopyAntChainContractProjectResponse) SetStatusCode(v int32) *CopyAntChainContractProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyAntChainContractProjectResponse) SetBody(v *CopyAntChainContractProjectResponseBody) *CopyAntChainContractProjectResponse {
	s.Body = v
	return s
}

type CreateAntChainAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AEDC32...
	AccountPubKey *string `json:"AccountPubKey,omitempty" xml:"AccountPubKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AEDC32...
	AccountRecoverPubKey *string `json:"AccountRecoverPubKey,omitempty" xml:"AccountRecoverPubKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bDXK6boZ
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s CreateAntChainAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainAccountRequest) GoString() string {
	return s.String()
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

func (s *CreateAntChainAccountRequest) SetAntChainId(v string) *CreateAntChainAccountRequest {
	s.AntChainId = &v
	return s
}

type CreateAntChainAccountResponseBody struct {
	// example:
	//
	// 0301F6CB-4FA6-4D03-8668-963623B63D0F
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
	// example:
	//
	// test
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// example:
	//
	// bDXK6boZ
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAntChainAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateAntChainAccountResponse) SetStatusCode(v int32) *CreateAntChainAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAntChainAccountResponse) SetBody(v *CreateAntChainAccountResponseBody) *CreateAntChainAccountResponse {
	s.Body = v
	return s
}

type CreateAntChainAccountWithKeyPairAutoCreationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// accountname
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1q8B5R9p
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ************
	RecoverPassword *string `json:"RecoverPassword,omitempty" xml:"RecoverPassword,omitempty"`
}

func (s CreateAntChainAccountWithKeyPairAutoCreationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainAccountWithKeyPairAutoCreationRequest) GoString() string {
	return s.String()
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationRequest) SetAccount(v string) *CreateAntChainAccountWithKeyPairAutoCreationRequest {
	s.Account = &v
	return s
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationRequest) SetAntChainId(v string) *CreateAntChainAccountWithKeyPairAutoCreationRequest {
	s.AntChainId = &v
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
	// example:
	//
	// 5F058BB4-3043-4638-86D4-EED84AD1AE54
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
	// example:
	//
	// accountname
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// example:
	//
	// -----BEGIN ENCRYPTED PRIVATE KEY-----\\nM***8=\\n-----END ENCRYPTED PRIVATE KEY-----\\n
	AccountPrivateKey *string `json:"AccountPrivateKey,omitempty" xml:"AccountPrivateKey,omitempty"`
	// example:
	//
	// b***7
	AccountPublicKey *string `json:"AccountPublicKey,omitempty" xml:"AccountPublicKey,omitempty"`
	// example:
	//
	// -----BEGIN ENCRYPTED PRIVATE KEY-----\\nM***Io=\\n-----END ENCRYPTED PRIVATE KEY-----\\n
	AccountRecoverPrivateKey *string `json:"AccountRecoverPrivateKey,omitempty" xml:"AccountRecoverPrivateKey,omitempty"`
	// example:
	//
	// d***b
	AccountRecoverPublicKey *string `json:"AccountRecoverPublicKey,omitempty" xml:"AccountRecoverPublicKey,omitempty"`
	// example:
	//
	// 1q8B5R9p
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult) SetAccount(v string) *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult {
	s.Account = &v
	return s
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult) SetAccountPrivateKey(v string) *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult {
	s.AccountPrivateKey = &v
	return s
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult) SetAccountPublicKey(v string) *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult {
	s.AccountPublicKey = &v
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

func (s *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult) SetAntChainId(v string) *CreateAntChainAccountWithKeyPairAutoCreationResponseBodyResult {
	s.AntChainId = &v
	return s
}

type CreateAntChainAccountWithKeyPairAutoCreationResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAntChainAccountWithKeyPairAutoCreationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateAntChainAccountWithKeyPairAutoCreationResponse) SetStatusCode(v int32) *CreateAntChainAccountWithKeyPairAutoCreationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAntChainAccountWithKeyPairAutoCreationResponse) SetBody(v *CreateAntChainAccountWithKeyPairAutoCreationResponseBody) *CreateAntChainAccountWithKeyPairAutoCreationResponse {
	s.Body = v
	return s
}

type CreateAntChainConsortiumRequest struct {
	ConsortiumDescription *string `json:"ConsortiumDescription,omitempty" xml:"ConsortiumDescription,omitempty"`
	// This parameter is required.
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
}

func (s CreateAntChainConsortiumRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainConsortiumRequest) GoString() string {
	return s.String()
}

func (s *CreateAntChainConsortiumRequest) SetConsortiumDescription(v string) *CreateAntChainConsortiumRequest {
	s.ConsortiumDescription = &v
	return s
}

func (s *CreateAntChainConsortiumRequest) SetConsortiumName(v string) *CreateAntChainConsortiumRequest {
	s.ConsortiumName = &v
	return s
}

type CreateAntChainConsortiumResponseBody struct {
	// example:
	//
	// 87FAFC6B-E7AE-4626-B7F7-AE0B88E612D5
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
	// example:
	//
	// A28nO1ow
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAntChainConsortiumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateAntChainConsortiumResponse) SetStatusCode(v int32) *CreateAntChainConsortiumResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAntChainConsortiumResponse) SetBody(v *CreateAntChainConsortiumResponseBody) *CreateAntChainConsortiumResponse {
	s.Body = v
	return s
}

type CreateAntChainContractContentRequest struct {
	// example:
	//
	// ""
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// newFile
	ContentName *string `json:"ContentName,omitempty" xml:"ContentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	IsDirectory *bool `json:"IsDirectory,omitempty" xml:"IsDirectory,omitempty"`
	// example:
	//
	// R38DYDop
	ParentContentId *string `json:"ParentContentId,omitempty" xml:"ParentContentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bDXK1b8Z
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateAntChainContractContentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainContractContentRequest) GoString() string {
	return s.String()
}

func (s *CreateAntChainContractContentRequest) SetContent(v string) *CreateAntChainContractContentRequest {
	s.Content = &v
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

func (s *CreateAntChainContractContentRequest) SetParentContentId(v string) *CreateAntChainContractContentRequest {
	s.ParentContentId = &v
	return s
}

func (s *CreateAntChainContractContentRequest) SetProjectId(v string) *CreateAntChainContractContentRequest {
	s.ProjectId = &v
	return s
}

type CreateAntChainContractContentResponseBody struct {
	// example:
	//
	// C407F102-062A-44B3-BBDC-E3B2718F633D
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
	// example:
	//
	// ""
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// R6XMn59e
	ContentId *string `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	// example:
	//
	// newFile
	ContentName *string `json:"ContentName,omitempty" xml:"ContentName,omitempty"`
	// example:
	//
	// 1564024964404
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// false
	IsDirectory *bool `json:"IsDirectory,omitempty" xml:"IsDirectory,omitempty"`
	// example:
	//
	// R38DYDop
	ParentContentId *string `json:"ParentContentId,omitempty" xml:"ParentContentId,omitempty"`
	// example:
	//
	// bDXK1b8Z
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 1564024964404
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateAntChainContractContentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainContractContentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAntChainContractContentResponseBodyResult) SetContent(v string) *CreateAntChainContractContentResponseBodyResult {
	s.Content = &v
	return s
}

func (s *CreateAntChainContractContentResponseBodyResult) SetContentId(v string) *CreateAntChainContractContentResponseBodyResult {
	s.ContentId = &v
	return s
}

func (s *CreateAntChainContractContentResponseBodyResult) SetContentName(v string) *CreateAntChainContractContentResponseBodyResult {
	s.ContentName = &v
	return s
}

func (s *CreateAntChainContractContentResponseBodyResult) SetCreateTime(v string) *CreateAntChainContractContentResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *CreateAntChainContractContentResponseBodyResult) SetIsDirectory(v bool) *CreateAntChainContractContentResponseBodyResult {
	s.IsDirectory = &v
	return s
}

func (s *CreateAntChainContractContentResponseBodyResult) SetParentContentId(v string) *CreateAntChainContractContentResponseBodyResult {
	s.ParentContentId = &v
	return s
}

func (s *CreateAntChainContractContentResponseBodyResult) SetProjectId(v string) *CreateAntChainContractContentResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *CreateAntChainContractContentResponseBodyResult) SetUpdateTime(v string) *CreateAntChainContractContentResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type CreateAntChainContractContentResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAntChainContractContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateAntChainContractContentResponse) SetStatusCode(v int32) *CreateAntChainContractContentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAntChainContractContentResponse) SetBody(v *CreateAntChainContractContentResponseBody) *CreateAntChainContractContentResponse {
	s.Body = v
	return s
}

type CreateAntChainContractProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DV80nJXq
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// project description
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// projectname
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1.0.0
	ProjectVersion *string `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
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

func (s *CreateAntChainContractProjectRequest) SetProjectDescription(v string) *CreateAntChainContractProjectRequest {
	s.ProjectDescription = &v
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

type CreateAntChainContractProjectResponseBody struct {
	// example:
	//
	// 4261A4C3-B8B5-4555-AEBC-944C3EAB23DB
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
	// example:
	//
	// DV80nJXq
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1563951889044
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// project description
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
	// example:
	//
	// R38DAbop
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// projectname
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// v1.0.0
	ProjectVersion *string `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
	// example:
	//
	// 1563951889140
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateAntChainContractProjectResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainContractProjectResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAntChainContractProjectResponseBodyResult) SetConsortiumId(v string) *CreateAntChainContractProjectResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *CreateAntChainContractProjectResponseBodyResult) SetCreateTime(v int64) *CreateAntChainContractProjectResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *CreateAntChainContractProjectResponseBodyResult) SetProjectDescription(v string) *CreateAntChainContractProjectResponseBodyResult {
	s.ProjectDescription = &v
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

func (s *CreateAntChainContractProjectResponseBodyResult) SetUpdateTime(v int64) *CreateAntChainContractProjectResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type CreateAntChainContractProjectResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAntChainContractProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateAntChainContractProjectResponse) SetStatusCode(v int32) *CreateAntChainContractProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAntChainContractProjectResponse) SetBody(v *CreateAntChainContractProjectResponseBody) *CreateAntChainContractProjectResponse {
	s.Body = v
	return s
}

type CreateAntChainKmsAccountNewRequest struct {
	// example:
	//
	// accountname
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s CreateAntChainKmsAccountNewRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainKmsAccountNewRequest) GoString() string {
	return s.String()
}

func (s *CreateAntChainKmsAccountNewRequest) SetAccount(v string) *CreateAntChainKmsAccountNewRequest {
	s.Account = &v
	return s
}

func (s *CreateAntChainKmsAccountNewRequest) SetAntChainId(v string) *CreateAntChainKmsAccountNewRequest {
	s.AntChainId = &v
	return s
}

type CreateAntChainKmsAccountNewResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 063C2175-C755-5C85-9C42-5EB432B67D9B
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateAntChainKmsAccountNewResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAntChainKmsAccountNewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainKmsAccountNewResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAntChainKmsAccountNewResponseBody) SetCode(v string) *CreateAntChainKmsAccountNewResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAntChainKmsAccountNewResponseBody) SetHttpStatusCode(v string) *CreateAntChainKmsAccountNewResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAntChainKmsAccountNewResponseBody) SetMessage(v string) *CreateAntChainKmsAccountNewResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAntChainKmsAccountNewResponseBody) SetRequestId(v string) *CreateAntChainKmsAccountNewResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAntChainKmsAccountNewResponseBody) SetResult(v *CreateAntChainKmsAccountNewResponseBodyResult) *CreateAntChainKmsAccountNewResponseBody {
	s.Result = v
	return s
}

func (s *CreateAntChainKmsAccountNewResponseBody) SetResultCode(v string) *CreateAntChainKmsAccountNewResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateAntChainKmsAccountNewResponseBody) SetResultMessage(v string) *CreateAntChainKmsAccountNewResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateAntChainKmsAccountNewResponseBody) SetSuccess(v bool) *CreateAntChainKmsAccountNewResponseBody {
	s.Success = &v
	return s
}

type CreateAntChainKmsAccountNewResponseBodyResult struct {
	// example:
	//
	// 1234565181228596659400e4-3bd7-4a83-b9af-d12654bb5779
	MyKmsKeyId *string `json:"MyKmsKeyId,omitempty" xml:"MyKmsKeyId,omitempty"`
	// example:
	//
	// d408f5c5c3c118b23646f4059c81fbe5a1b067a12f96f4a7a5e09eae10c81288cdf64aa887af89738b9ec5d16bc124594bf820a4afa24fe0de9d8dcfd1d99500
	PubKey *string `json:"PubKey,omitempty" xml:"PubKey,omitempty"`
}

func (s CreateAntChainKmsAccountNewResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainKmsAccountNewResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAntChainKmsAccountNewResponseBodyResult) SetMyKmsKeyId(v string) *CreateAntChainKmsAccountNewResponseBodyResult {
	s.MyKmsKeyId = &v
	return s
}

func (s *CreateAntChainKmsAccountNewResponseBodyResult) SetPubKey(v string) *CreateAntChainKmsAccountNewResponseBodyResult {
	s.PubKey = &v
	return s
}

type CreateAntChainKmsAccountNewResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAntChainKmsAccountNewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAntChainKmsAccountNewResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAntChainKmsAccountNewResponse) GoString() string {
	return s.String()
}

func (s *CreateAntChainKmsAccountNewResponse) SetHeaders(v map[string]*string) *CreateAntChainKmsAccountNewResponse {
	s.Headers = v
	return s
}

func (s *CreateAntChainKmsAccountNewResponse) SetStatusCode(v int32) *CreateAntChainKmsAccountNewResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAntChainKmsAccountNewResponse) SetBody(v *CreateAntChainKmsAccountNewResponseBody) *CreateAntChainKmsAccountNewResponse {
	s.Body = v
	return s
}

type CreateFabricChaincodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// chan-first-channel-1w55v3u39x2xz
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// consortium-aaaaaa-akpcsjjac2jd
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// OR (\\"aaaaaa1MSP.peer\\")
	EndorsePolicy *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa1-1oxw31d046jtl
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://chaincode.oss-cn-beijing.aliyuncs.com
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aaaaaa1/a0cf8729-7d47-44c9-9ed5-6d2fdfc8dc2c/sacc.cc
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
}

func (s CreateFabricChaincodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChaincodeRequest) GoString() string {
	return s.String()
}

func (s *CreateFabricChaincodeRequest) SetChannelId(v string) *CreateFabricChaincodeRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateFabricChaincodeRequest) SetConsortiumId(v string) *CreateFabricChaincodeRequest {
	s.ConsortiumId = &v
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

func (s *CreateFabricChaincodeRequest) SetOrganizationId(v string) *CreateFabricChaincodeRequest {
	s.OrganizationId = &v
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

type CreateFabricChaincodeResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// CD264CC8-B892-4CDC-BD31-7D179EE6E396
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateFabricChaincodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFabricChaincodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChaincodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFabricChaincodeResponseBody) SetErrorCode(v int32) *CreateFabricChaincodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFabricChaincodeResponseBody) SetRequestId(v string) *CreateFabricChaincodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFabricChaincodeResponseBody) SetResult(v *CreateFabricChaincodeResponseBodyResult) *CreateFabricChaincodeResponseBody {
	s.Result = v
	return s
}

func (s *CreateFabricChaincodeResponseBody) SetSuccess(v bool) *CreateFabricChaincodeResponseBody {
	s.Success = &v
	return s
}

type CreateFabricChaincodeResponseBodyResult struct {
	// example:
	//
	// code-sacc-1pr09q7jmo0np
	ChaincodeId *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	// example:
	//
	// mychaincode
	ChaincodeName *string `json:"ChaincodeName,omitempty" xml:"ChaincodeName,omitempty"`
	// example:
	//
	// 1.0
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" xml:"ChaincodeVersion,omitempty"`
	// example:
	//
	// first-channel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// consortium-aaaaaa-akpcsjjac2jd
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1544766801000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1544766801000
	DeployTime *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	// example:
	//
	// OR (\\"aaaaaa1MSP.peer\\")
	EndorsePolicy *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	// example:
	//
	// string
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// true
	Install *bool `json:"Install,omitempty" xml:"Install,omitempty"`
	// example:
	//
	// string
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// go-sdk-demo/chaincode/src
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// peers-aaaaaa1-1oxw31d046jtl
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// example:
	//
	// peers-aaaaaa1-1oxw31d046jtl
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// example:
	//
	// Instantiatable
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFabricChaincodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChaincodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateFabricChaincodeResponseBodyResult) SetChaincodeId(v string) *CreateFabricChaincodeResponseBodyResult {
	s.ChaincodeId = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetChaincodeName(v string) *CreateFabricChaincodeResponseBodyResult {
	s.ChaincodeName = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetChaincodeVersion(v string) *CreateFabricChaincodeResponseBodyResult {
	s.ChaincodeVersion = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetChannelName(v string) *CreateFabricChaincodeResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetConsortiumId(v string) *CreateFabricChaincodeResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetCreateTime(v string) *CreateFabricChaincodeResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetDeployTime(v string) *CreateFabricChaincodeResponseBodyResult {
	s.DeployTime = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetEndorsePolicy(v string) *CreateFabricChaincodeResponseBodyResult {
	s.EndorsePolicy = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetInput(v string) *CreateFabricChaincodeResponseBodyResult {
	s.Input = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetInstall(v bool) *CreateFabricChaincodeResponseBodyResult {
	s.Install = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetMessage(v string) *CreateFabricChaincodeResponseBodyResult {
	s.Message = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetPath(v string) *CreateFabricChaincodeResponseBodyResult {
	s.Path = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetProviderId(v string) *CreateFabricChaincodeResponseBodyResult {
	s.ProviderId = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetProviderName(v string) *CreateFabricChaincodeResponseBodyResult {
	s.ProviderName = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetState(v string) *CreateFabricChaincodeResponseBodyResult {
	s.State = &v
	return s
}

func (s *CreateFabricChaincodeResponseBodyResult) SetType(v int32) *CreateFabricChaincodeResponseBodyResult {
	s.Type = &v
	return s
}

type CreateFabricChaincodeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFabricChaincodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateFabricChaincodeResponse) SetStatusCode(v int32) *CreateFabricChaincodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFabricChaincodeResponse) SetBody(v *CreateFabricChaincodeResponseBody) *CreateFabricChaincodeResponse {
	s.Body = v
	return s
}

type CreateFabricChaincodePackageRequest struct {
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hi003/7efbe860-1a40-42c9-9e6b-0bc08d7ee3a4/cmcc.tar.gz
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
}

func (s CreateFabricChaincodePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChaincodePackageRequest) GoString() string {
	return s.String()
}

func (s *CreateFabricChaincodePackageRequest) SetLocation(v string) *CreateFabricChaincodePackageRequest {
	s.Location = &v
	return s
}

func (s *CreateFabricChaincodePackageRequest) SetOrganizationId(v string) *CreateFabricChaincodePackageRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateFabricChaincodePackageRequest) SetOssUrl(v string) *CreateFabricChaincodePackageRequest {
	s.OssUrl = &v
	return s
}

type CreateFabricChaincodePackageResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ChaincodePackage `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFabricChaincodePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChaincodePackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFabricChaincodePackageResponseBody) SetErrorCode(v int32) *CreateFabricChaincodePackageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFabricChaincodePackageResponseBody) SetMessage(v string) *CreateFabricChaincodePackageResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFabricChaincodePackageResponseBody) SetRequestId(v string) *CreateFabricChaincodePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFabricChaincodePackageResponseBody) SetResult(v *ChaincodePackage) *CreateFabricChaincodePackageResponseBody {
	s.Result = v
	return s
}

func (s *CreateFabricChaincodePackageResponseBody) SetSuccess(v bool) *CreateFabricChaincodePackageResponseBody {
	s.Success = &v
	return s
}

type CreateFabricChaincodePackageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFabricChaincodePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFabricChaincodePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChaincodePackageResponse) GoString() string {
	return s.String()
}

func (s *CreateFabricChaincodePackageResponse) SetHeaders(v map[string]*string) *CreateFabricChaincodePackageResponse {
	s.Headers = v
	return s
}

func (s *CreateFabricChaincodePackageResponse) SetStatusCode(v int32) *CreateFabricChaincodePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFabricChaincodePackageResponse) SetBody(v *CreateFabricChaincodePackageResponseBody) *CreateFabricChaincodePackageResponse {
	s.Body = v
	return s
}

type CreateFabricChannelRequest struct {
	// example:
	//
	// 2
	BatchTimeout *int32 `json:"BatchTimeout,omitempty" xml:"BatchTimeout,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// channelx
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// consortium-aaaaaa-akpcsjjac2jd
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 50
	MaxMessageCount *int32 `json:"MaxMessageCount,omitempty" xml:"MaxMessageCount,omitempty"`
	// This parameter is required.
	Organization []*CreateFabricChannelRequestOrganization `json:"Organization,omitempty" xml:"Organization,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	PreferredMaxBytes *int32 `json:"PreferredMaxBytes,omitempty" xml:"PreferredMaxBytes,omitempty"`
}

func (s CreateFabricChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateFabricChannelRequest) SetBatchTimeout(v int32) *CreateFabricChannelRequest {
	s.BatchTimeout = &v
	return s
}

func (s *CreateFabricChannelRequest) SetChannelName(v string) *CreateFabricChannelRequest {
	s.ChannelName = &v
	return s
}

func (s *CreateFabricChannelRequest) SetConsortiumId(v string) *CreateFabricChannelRequest {
	s.ConsortiumId = &v
	return s
}

func (s *CreateFabricChannelRequest) SetMaxMessageCount(v int32) *CreateFabricChannelRequest {
	s.MaxMessageCount = &v
	return s
}

func (s *CreateFabricChannelRequest) SetOrganization(v []*CreateFabricChannelRequestOrganization) *CreateFabricChannelRequest {
	s.Organization = v
	return s
}

func (s *CreateFabricChannelRequest) SetPreferredMaxBytes(v int32) *CreateFabricChannelRequest {
	s.PreferredMaxBytes = &v
	return s
}

type CreateFabricChannelRequestOrganization struct {
	// example:
	//
	// peers-aaaaaa1-1oxw31d046jtl
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateFabricChannelResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFabricChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFabricChannelResponseBody) SetErrorCode(v int32) *CreateFabricChannelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFabricChannelResponseBody) SetRequestId(v string) *CreateFabricChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFabricChannelResponseBody) SetResult(v *CreateFabricChannelResponseBodyResult) *CreateFabricChannelResponseBody {
	s.Result = v
	return s
}

func (s *CreateFabricChannelResponseBody) SetSuccess(v bool) *CreateFabricChannelResponseBody {
	s.Success = &v
	return s
}

type CreateFabricChannelResponseBodyResult struct {
	// example:
	//
	// 2
	BatchTimeout *int32 `json:"BatchTimeout,omitempty" xml:"BatchTimeout,omitempty"`
	// example:
	//
	// 2
	BlockCount *int32 `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
	// example:
	//
	// 2
	ChaincodeCount *int32 `json:"ChaincodeCount,omitempty" xml:"ChaincodeCount,omitempty"`
	// example:
	//
	// chan-channelx-1l1hmckuuisxo
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// channelx
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// consortium-aaaaaa-akpcsjjac2jd
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// aaaaaa
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	// example:
	//
	// 1544768139624
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 50
	MaxMessageCount *int32 `json:"MaxMessageCount,omitempty" xml:"MaxMessageCount,omitempty"`
	// example:
	//
	// 1
	MemberCount *int32 `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	// example:
	//
	// 26842
	OwnerBid *string `json:"OwnerBid,omitempty" xml:"OwnerBid,omitempty"`
	// example:
	//
	// uid-
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 1234
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// example:
	//
	// 12
	PreferredMaxBytes *int32 `json:"PreferredMaxBytes,omitempty" xml:"PreferredMaxBytes,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// true
	SupportConfig *bool `json:"SupportConfig,omitempty" xml:"SupportConfig,omitempty"`
	// example:
	//
	// 1544768139624
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

func (s *CreateFabricChannelResponseBodyResult) SetBlockCount(v int32) *CreateFabricChannelResponseBodyResult {
	s.BlockCount = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetChaincodeCount(v int32) *CreateFabricChannelResponseBodyResult {
	s.ChaincodeCount = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetChannelId(v string) *CreateFabricChannelResponseBodyResult {
	s.ChannelId = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetChannelName(v string) *CreateFabricChannelResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetConsortiumId(v string) *CreateFabricChannelResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetConsortiumName(v string) *CreateFabricChannelResponseBodyResult {
	s.ConsortiumName = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetCreateTime(v string) *CreateFabricChannelResponseBodyResult {
	s.CreateTime = &v
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

func (s *CreateFabricChannelResponseBodyResult) SetOwnerBid(v string) *CreateFabricChannelResponseBodyResult {
	s.OwnerBid = &v
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

func (s *CreateFabricChannelResponseBodyResult) SetPreferredMaxBytes(v int32) *CreateFabricChannelResponseBodyResult {
	s.PreferredMaxBytes = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetRequestId(v string) *CreateFabricChannelResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetState(v string) *CreateFabricChannelResponseBodyResult {
	s.State = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetSupportConfig(v bool) *CreateFabricChannelResponseBodyResult {
	s.SupportConfig = &v
	return s
}

func (s *CreateFabricChannelResponseBodyResult) SetUpdateTime(v string) *CreateFabricChannelResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type CreateFabricChannelResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFabricChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateFabricChannelResponse) SetStatusCode(v int32) *CreateFabricChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFabricChannelResponse) SetBody(v *CreateFabricChannelResponseBody) *CreateFabricChannelResponse {
	s.Body = v
	return s
}

type CreateFabricChannelMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// chan-channelx-1l1hmckuuisxo
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa1-1oxw31d046jtl
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFabricChannelMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricChannelMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFabricChannelMemberResponseBody) SetErrorCode(v int32) *CreateFabricChannelMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFabricChannelMemberResponseBody) SetRequestId(v string) *CreateFabricChannelMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFabricChannelMemberResponseBody) SetResult(v bool) *CreateFabricChannelMemberResponseBody {
	s.Result = &v
	return s
}

func (s *CreateFabricChannelMemberResponseBody) SetSuccess(v bool) *CreateFabricChannelMemberResponseBody {
	s.Success = &v
	return s
}

type CreateFabricChannelMemberResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFabricChannelMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateFabricChannelMemberResponse) SetStatusCode(v int32) *CreateFabricChannelMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFabricChannelMemberResponse) SetBody(v *CreateFabricChannelMemberResponseBody) *CreateFabricChannelMemberResponse {
	s.Body = v
	return s
}

type CreateFabricConsortiumRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// all
	ChannelPolicy *string `json:"ChannelPolicy,omitempty" xml:"ChannelPolicy,omitempty"`
	// example:
	//
	// some
	ConsortiumDescription *string `json:"ConsortiumDescription,omitempty" xml:"ConsortiumDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aaaaa
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// helloworld
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// etcdraft
	OrdererType *string `json:"OrdererType,omitempty" xml:"OrdererType,omitempty"`
	// example:
	//
	// 2
	OrderersCount *int32                                       `json:"OrderersCount,omitempty" xml:"OrderersCount,omitempty"`
	Organization  []*CreateFabricConsortiumRequestOrganization `json:"Organization,omitempty" xml:"Organization,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	PaymentDuration *int32 `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// month
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	// example:
	//
	// 2
	PeersCount *int32 `json:"PeersCount,omitempty" xml:"PeersCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// basic
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	// example:
	//
	// random
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateFabricConsortiumRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricConsortiumRequest) GoString() string {
	return s.String()
}

func (s *CreateFabricConsortiumRequest) SetChannelPolicy(v string) *CreateFabricConsortiumRequest {
	s.ChannelPolicy = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetConsortiumDescription(v string) *CreateFabricConsortiumRequest {
	s.ConsortiumDescription = &v
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

func (s *CreateFabricConsortiumRequest) SetLocation(v string) *CreateFabricConsortiumRequest {
	s.Location = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetOrdererType(v string) *CreateFabricConsortiumRequest {
	s.OrdererType = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetOrderersCount(v int32) *CreateFabricConsortiumRequest {
	s.OrderersCount = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetOrganization(v []*CreateFabricConsortiumRequestOrganization) *CreateFabricConsortiumRequest {
	s.Organization = v
	return s
}

func (s *CreateFabricConsortiumRequest) SetPaymentDuration(v int32) *CreateFabricConsortiumRequest {
	s.PaymentDuration = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetPaymentDurationUnit(v string) *CreateFabricConsortiumRequest {
	s.PaymentDurationUnit = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetPeersCount(v int32) *CreateFabricConsortiumRequest {
	s.PeersCount = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetSpecName(v string) *CreateFabricConsortiumRequest {
	s.SpecName = &v
	return s
}

func (s *CreateFabricConsortiumRequest) SetZoneId(v string) *CreateFabricConsortiumRequest {
	s.ZoneId = &v
	return s
}

type CreateFabricConsortiumRequestOrganization struct {
	// example:
	//
	// peers-yidio-1tuigx42b1goc
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateFabricConsortiumResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFabricConsortiumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricConsortiumResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFabricConsortiumResponseBody) SetErrorCode(v int32) *CreateFabricConsortiumResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFabricConsortiumResponseBody) SetRequestId(v string) *CreateFabricConsortiumResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFabricConsortiumResponseBody) SetResult(v *CreateFabricConsortiumResponseBodyResult) *CreateFabricConsortiumResponseBody {
	s.Result = v
	return s
}

func (s *CreateFabricConsortiumResponseBody) SetSuccess(v bool) *CreateFabricConsortiumResponseBody {
	s.Success = &v
	return s
}

type CreateFabricConsortiumResponseBodyResult struct {
	// example:
	//
	// 0
	ChannelCount *int32 `json:"ChannelCount,omitempty" xml:"ChannelCount,omitempty"`
	// example:
	//
	// Any
	ChannelPolicy *string `json:"ChannelPolicy,omitempty" xml:"ChannelPolicy,omitempty"`
	// example:
	//
	// Pending
	ClusterState *string `json:"ClusterState,omitempty" xml:"ClusterState,omitempty"`
	// example:
	//
	// lianmenyumingyi
	CodeName *string `json:"CodeName,omitempty" xml:"CodeName,omitempty"`
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwlulg7
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// myconsortium
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	// example:
	//
	// 1544086901984
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// string
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// hello
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 2
	MemberCount *int32 `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	// example:
	//
	// 1
	OrdererCount *int32 `json:"OrdererCount,omitempty" xml:"OrdererCount,omitempty"`
	// example:
	//
	// Kafka
	OrdererType *string `json:"OrdererType,omitempty" xml:"OrdererType,omitempty"`
	// example:
	//
	// 26842
	OwnerBid *string `json:"OwnerBid,omitempty" xml:"OwnerBid,omitempty"`
	// example:
	//
	// 1019556
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Pending
	ServiceState *string `json:"ServiceState,omitempty" xml:"ServiceState,omitempty"`
	// example:
	//
	// basic
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	// example:
	//
	// cn-hangzhou
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *CreateFabricConsortiumResponseBodyResult) SetChannelPolicy(v string) *CreateFabricConsortiumResponseBodyResult {
	s.ChannelPolicy = &v
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

func (s *CreateFabricConsortiumResponseBodyResult) SetConsortiumId(v string) *CreateFabricConsortiumResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetConsortiumName(v string) *CreateFabricConsortiumResponseBodyResult {
	s.ConsortiumName = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetCreateTime(v string) *CreateFabricConsortiumResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetDescription(v string) *CreateFabricConsortiumResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetDomain(v string) *CreateFabricConsortiumResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetMemberCount(v int32) *CreateFabricConsortiumResponseBodyResult {
	s.MemberCount = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetOrdererCount(v int32) *CreateFabricConsortiumResponseBodyResult {
	s.OrdererCount = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetOrdererType(v string) *CreateFabricConsortiumResponseBodyResult {
	s.OrdererType = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetOwnerBid(v string) *CreateFabricConsortiumResponseBodyResult {
	s.OwnerBid = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetOwnerUid(v int64) *CreateFabricConsortiumResponseBodyResult {
	s.OwnerUid = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetRegionId(v string) *CreateFabricConsortiumResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetServiceState(v string) *CreateFabricConsortiumResponseBodyResult {
	s.ServiceState = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetSpecName(v string) *CreateFabricConsortiumResponseBodyResult {
	s.SpecName = &v
	return s
}

func (s *CreateFabricConsortiumResponseBodyResult) SetZoneId(v string) *CreateFabricConsortiumResponseBodyResult {
	s.ZoneId = &v
	return s
}

type CreateFabricConsortiumResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFabricConsortiumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateFabricConsortiumResponse) SetStatusCode(v int32) *CreateFabricConsortiumResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFabricConsortiumResponse) SetBody(v *CreateFabricConsortiumResponseBody) *CreateFabricConsortiumResponse {
	s.Body = v
	return s
}

type CreateFabricConsortiumMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// consortium-aaaaaa-akpcsjjac2jd
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// This parameter is required.
	Organization []*CreateFabricConsortiumMemberRequestOrganization `json:"Organization,omitempty" xml:"Organization,omitempty" type:"Repeated"`
}

func (s CreateFabricConsortiumMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricConsortiumMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateFabricConsortiumMemberRequest) SetCode(v string) *CreateFabricConsortiumMemberRequest {
	s.Code = &v
	return s
}

func (s *CreateFabricConsortiumMemberRequest) SetConsortiumId(v string) *CreateFabricConsortiumMemberRequest {
	s.ConsortiumId = &v
	return s
}

func (s *CreateFabricConsortiumMemberRequest) SetOrganization(v []*CreateFabricConsortiumMemberRequestOrganization) *CreateFabricConsortiumMemberRequest {
	s.Organization = v
	return s
}

type CreateFabricConsortiumMemberRequestOrganization struct {
	// example:
	//
	// peers-aaaaaa1-1oxw31d046jtl
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFabricConsortiumMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricConsortiumMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFabricConsortiumMemberResponseBody) SetErrorCode(v int32) *CreateFabricConsortiumMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFabricConsortiumMemberResponseBody) SetRequestId(v string) *CreateFabricConsortiumMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFabricConsortiumMemberResponseBody) SetResult(v bool) *CreateFabricConsortiumMemberResponseBody {
	s.Result = &v
	return s
}

func (s *CreateFabricConsortiumMemberResponseBody) SetSuccess(v bool) *CreateFabricConsortiumMemberResponseBody {
	s.Success = &v
	return s
}

type CreateFabricConsortiumMemberResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFabricConsortiumMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateFabricConsortiumMemberResponse) SetStatusCode(v int32) *CreateFabricConsortiumMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFabricConsortiumMemberResponse) SetBody(v *CreateFabricConsortiumMemberResponseBody) *CreateFabricConsortiumMemberResponse {
	s.Body = v
	return s
}

type CreateFabricOrganizationRequest struct {
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// domain
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// name
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12
	PaymentDuration *int32 `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	// example:
	//
	// 2
	PeersCount *int32 `json:"PeersCount,omitempty" xml:"PeersCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// basic
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
}

func (s CreateFabricOrganizationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricOrganizationRequest) GoString() string {
	return s.String()
}

func (s *CreateFabricOrganizationRequest) SetDescription(v string) *CreateFabricOrganizationRequest {
	s.Description = &v
	return s
}

func (s *CreateFabricOrganizationRequest) SetDomain(v string) *CreateFabricOrganizationRequest {
	s.Domain = &v
	return s
}

func (s *CreateFabricOrganizationRequest) SetLocation(v string) *CreateFabricOrganizationRequest {
	s.Location = &v
	return s
}

func (s *CreateFabricOrganizationRequest) SetOrganizationName(v string) *CreateFabricOrganizationRequest {
	s.OrganizationName = &v
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

func (s *CreateFabricOrganizationRequest) SetPeersCount(v int32) *CreateFabricOrganizationRequest {
	s.PeersCount = &v
	return s
}

func (s *CreateFabricOrganizationRequest) SetSpecName(v string) *CreateFabricOrganizationRequest {
	s.SpecName = &v
	return s
}

type CreateFabricOrganizationResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateFabricOrganizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFabricOrganizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFabricOrganizationResponseBody) SetErrorCode(v int32) *CreateFabricOrganizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFabricOrganizationResponseBody) SetRequestId(v string) *CreateFabricOrganizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFabricOrganizationResponseBody) SetResult(v *CreateFabricOrganizationResponseBodyResult) *CreateFabricOrganizationResponseBody {
	s.Result = v
	return s
}

func (s *CreateFabricOrganizationResponseBody) SetSuccess(v bool) *CreateFabricOrganizationResponseBody {
	s.Success = &v
	return s
}

type CreateFabricOrganizationResponseBodyResult struct {
	// example:
	//
	// Pending
	ClusterState *string `json:"ClusterState,omitempty" xml:"ClusterState,omitempty"`
	// example:
	//
	// name
	CodeName *string `json:"CodeName,omitempty" xml:"CodeName,omitempty"`
	// example:
	//
	// 2
	ConsortiumCount *int32 `json:"ConsortiumCount,omitempty" xml:"ConsortiumCount,omitempty"`
	// example:
	//
	// 1544086901984
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// domain
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// desc
	OrganizationDescription *string `json:"OrganizationDescription,omitempty" xml:"OrganizationDescription,omitempty"`
	// example:
	//
	// zuzhiyumingyi-hc5d1bwlulg7
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// zuzhiyumingyi
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	// example:
	//
	// 26842
	OwnerBid *string `json:"OwnerBid,omitempty" xml:"OwnerBid,omitempty"`
	// example:
	//
	// name
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 1019556
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// example:
	//
	// 3
	PeerCount *int32 `json:"PeerCount,omitempty" xml:"PeerCount,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Pending
	ServiceState *string `json:"ServiceState,omitempty" xml:"ServiceState,omitempty"`
	// example:
	//
	// baisc
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	// example:
	//
	// 2
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	// example:
	//
	// zone
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateFabricOrganizationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricOrganizationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateFabricOrganizationResponseBodyResult) SetClusterState(v string) *CreateFabricOrganizationResponseBodyResult {
	s.ClusterState = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetCodeName(v string) *CreateFabricOrganizationResponseBodyResult {
	s.CodeName = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetConsortiumCount(v int32) *CreateFabricOrganizationResponseBodyResult {
	s.ConsortiumCount = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetCreateTime(v string) *CreateFabricOrganizationResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetDomain(v string) *CreateFabricOrganizationResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetOrganizationDescription(v string) *CreateFabricOrganizationResponseBodyResult {
	s.OrganizationDescription = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetOrganizationId(v string) *CreateFabricOrganizationResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetOrganizationName(v string) *CreateFabricOrganizationResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetOwnerBid(v string) *CreateFabricOrganizationResponseBodyResult {
	s.OwnerBid = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetOwnerName(v string) *CreateFabricOrganizationResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetOwnerUid(v int64) *CreateFabricOrganizationResponseBodyResult {
	s.OwnerUid = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetPeerCount(v int32) *CreateFabricOrganizationResponseBodyResult {
	s.PeerCount = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetRegionId(v string) *CreateFabricOrganizationResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetRequestId(v string) *CreateFabricOrganizationResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetServiceState(v string) *CreateFabricOrganizationResponseBodyResult {
	s.ServiceState = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetSpecName(v string) *CreateFabricOrganizationResponseBodyResult {
	s.SpecName = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetUserCount(v int32) *CreateFabricOrganizationResponseBodyResult {
	s.UserCount = &v
	return s
}

func (s *CreateFabricOrganizationResponseBodyResult) SetZoneId(v string) *CreateFabricOrganizationResponseBodyResult {
	s.ZoneId = &v
	return s
}

type CreateFabricOrganizationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFabricOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateFabricOrganizationResponse) SetStatusCode(v int32) *CreateFabricOrganizationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFabricOrganizationResponse) SetBody(v *CreateFabricOrganizationResponseBody) *CreateFabricOrganizationResponse {
	s.Body = v
	return s
}

type CreateFabricOrganizationUserRequest struct {
	// example:
	//
	// foo=foo1,bar=bar1
	Attrs *string `json:"Attrs,omitempty" xml:"Attrs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-yidio-1tuigx42b****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// username
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateFabricOrganizationUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricOrganizationUserRequest) GoString() string {
	return s.String()
}

func (s *CreateFabricOrganizationUserRequest) SetAttrs(v string) *CreateFabricOrganizationUserRequest {
	s.Attrs = &v
	return s
}

func (s *CreateFabricOrganizationUserRequest) SetOrganizationId(v string) *CreateFabricOrganizationUserRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateFabricOrganizationUserRequest) SetPassword(v string) *CreateFabricOrganizationUserRequest {
	s.Password = &v
	return s
}

func (s *CreateFabricOrganizationUserRequest) SetUsername(v string) *CreateFabricOrganizationUserRequest {
	s.Username = &v
	return s
}

type CreateFabricOrganizationUserResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateFabricOrganizationUserResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFabricOrganizationUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricOrganizationUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFabricOrganizationUserResponseBody) SetErrorCode(v int32) *CreateFabricOrganizationUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFabricOrganizationUserResponseBody) SetRequestId(v string) *CreateFabricOrganizationUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFabricOrganizationUserResponseBody) SetResult(v *CreateFabricOrganizationUserResponseBodyResult) *CreateFabricOrganizationUserResponseBody {
	s.Result = v
	return s
}

func (s *CreateFabricOrganizationUserResponseBody) SetSuccess(v bool) *CreateFabricOrganizationUserResponseBody {
	s.Success = &v
	return s
}

type CreateFabricOrganizationUserResponseBodyResult struct {
	// example:
	//
	// 1544086901984
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1544411108000
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// fullname
	Fullname *string `json:"Fullname,omitempty" xml:"Fullname,omitempty"`
	// example:
	//
	// peers-yidio-1tuigx42b****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// username
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateFabricOrganizationUserResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateFabricOrganizationUserResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateFabricOrganizationUserResponseBodyResult) SetCreateTime(v string) *CreateFabricOrganizationUserResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *CreateFabricOrganizationUserResponseBodyResult) SetExpireTime(v string) *CreateFabricOrganizationUserResponseBodyResult {
	s.ExpireTime = &v
	return s
}

func (s *CreateFabricOrganizationUserResponseBodyResult) SetFullname(v string) *CreateFabricOrganizationUserResponseBodyResult {
	s.Fullname = &v
	return s
}

func (s *CreateFabricOrganizationUserResponseBodyResult) SetOrganizationId(v string) *CreateFabricOrganizationUserResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *CreateFabricOrganizationUserResponseBodyResult) SetPassword(v string) *CreateFabricOrganizationUserResponseBodyResult {
	s.Password = &v
	return s
}

func (s *CreateFabricOrganizationUserResponseBodyResult) SetUsername(v string) *CreateFabricOrganizationUserResponseBodyResult {
	s.Username = &v
	return s
}

type CreateFabricOrganizationUserResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFabricOrganizationUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateFabricOrganizationUserResponse) SetStatusCode(v int32) *CreateFabricOrganizationUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFabricOrganizationUserResponse) SetBody(v *CreateFabricOrganizationUserResponseBody) *CreateFabricOrganizationUserResponse {
	s.Body = v
	return s
}

type DeleteAntChainConsortiumRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// q0oWq92P
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
	// example:
	//
	// F3685787-9F2F-4D60-ADD6-07A5179552CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAntChainConsortiumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteAntChainConsortiumResponse) SetStatusCode(v int32) *DeleteAntChainConsortiumResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAntChainConsortiumResponse) SetBody(v *DeleteAntChainConsortiumResponseBody) *DeleteAntChainConsortiumResponse {
	s.Body = v
	return s
}

type DeleteAntChainContractContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// GKX7KZX0
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
	// example:
	//
	// 7E54F588-39DE-446F-9A16-B9BB0B05DF44
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAntChainContractContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteAntChainContractContentResponse) SetStatusCode(v int32) *DeleteAntChainContractContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAntChainContractContentResponse) SetBody(v *DeleteAntChainContractContentResponseBody) *DeleteAntChainContractContentResponse {
	s.Body = v
	return s
}

type DeleteAntChainContractProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// AVXeKq8R
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
	// example:
	//
	// 4FAC5F05-FDF0-4599-9522-0120525F0145
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAntChainContractProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteAntChainContractProjectResponse) SetStatusCode(v int32) *DeleteAntChainContractProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAntChainContractProjectResponse) SetBody(v *DeleteAntChainContractProjectResponseBody) *DeleteAntChainContractProjectResponse {
	s.Body = v
	return s
}

type DeleteAntChainMiniAppQRCodeAuthorizedUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bDXK6boZ
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1380013****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
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
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAntChainMiniAppQRCodeAuthorizedUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteAntChainMiniAppQRCodeAuthorizedUserResponse) SetStatusCode(v int32) *DeleteAntChainMiniAppQRCodeAuthorizedUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAntChainMiniAppQRCodeAuthorizedUserResponse) SetBody(v *DeleteAntChainMiniAppQRCodeAuthorizedUserResponseBody) *DeleteAntChainMiniAppQRCodeAuthorizedUserResponse {
	s.Body = v
	return s
}

type DeleteFabricChaincodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// code-sacc-1pr09q7jm****
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFabricChaincodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFabricChaincodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFabricChaincodeResponseBody) SetErrorCode(v int32) *DeleteFabricChaincodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFabricChaincodeResponseBody) SetRequestId(v string) *DeleteFabricChaincodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFabricChaincodeResponseBody) SetSuccess(v bool) *DeleteFabricChaincodeResponseBody {
	s.Success = &v
	return s
}

type DeleteFabricChaincodeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFabricChaincodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteFabricChaincodeResponse) SetStatusCode(v int32) *DeleteFabricChaincodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFabricChaincodeResponse) SetBody(v *DeleteFabricChaincodeResponseBody) *DeleteFabricChaincodeResponse {
	s.Body = v
	return s
}

type DescribeAntChainAccountsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pYogqb9v
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAntChainAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainAccountsRequest) SetAntChainId(v string) *DescribeAntChainAccountsRequest {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainAccountsRequest) SetPageNumber(v int32) *DescribeAntChainAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainAccountsRequest) SetPageSize(v int32) *DescribeAntChainAccountsRequest {
	s.PageSize = &v
	return s
}

type DescribeAntChainAccountsResponseBody struct {
	// example:
	//
	// 1692940F-87C5-468F-8188-6C95FC05BBA5
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
	Accounts   []*DescribeAntChainAccountsResponseBodyResultAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	Pagination *DescribeAntChainAccountsResponseBodyResultPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s DescribeAntChainAccountsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainAccountsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainAccountsResponseBodyResult) SetAccounts(v []*DescribeAntChainAccountsResponseBodyResultAccounts) *DescribeAntChainAccountsResponseBodyResult {
	s.Accounts = v
	return s
}

func (s *DescribeAntChainAccountsResponseBodyResult) SetPagination(v *DescribeAntChainAccountsResponseBodyResultPagination) *DescribeAntChainAccountsResponseBodyResult {
	s.Pagination = v
	return s
}

type DescribeAntChainAccountsResponseBodyResultAccounts struct {
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// example:
	//
	// de44334f1094cbd37e952ef691a837f36eab2e9afdf6bed9bf523c2dde40ff7b1ff526ec11b510d320f536685f75dd90f67fda8d41a69738a76a77c27b1373fb
	AccountPublicKey *string `json:"AccountPublicKey,omitempty" xml:"AccountPublicKey,omitempty"`
	// example:
	//
	// ee54a9c4511345f2de05889540eedb8dfed7b57ab0647c5894e4333596bc2ecef07a0469331d016ad9978b39367a0a5c217ef92ffffcbd6716db6667353891a9
	AccountRecoveryKey *string `json:"AccountRecoveryKey,omitempty" xml:"AccountRecoveryKey,omitempty"`
	// example:
	//
	// NORMAL
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// example:
	//
	// pYogqb9v
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s DescribeAntChainAccountsResponseBodyResultAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainAccountsResponseBodyResultAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAntChainAccountsResponseBodyResultAccounts) SetAccount(v string) *DescribeAntChainAccountsResponseBodyResultAccounts {
	s.Account = &v
	return s
}

func (s *DescribeAntChainAccountsResponseBodyResultAccounts) SetAccountPublicKey(v string) *DescribeAntChainAccountsResponseBodyResultAccounts {
	s.AccountPublicKey = &v
	return s
}

func (s *DescribeAntChainAccountsResponseBodyResultAccounts) SetAccountRecoveryKey(v string) *DescribeAntChainAccountsResponseBodyResultAccounts {
	s.AccountRecoveryKey = &v
	return s
}

func (s *DescribeAntChainAccountsResponseBodyResultAccounts) SetAccountStatus(v string) *DescribeAntChainAccountsResponseBodyResultAccounts {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAntChainAccountsResponseBodyResultAccounts) SetAntChainId(v string) *DescribeAntChainAccountsResponseBodyResultAccounts {
	s.AntChainId = &v
	return s
}

type DescribeAntChainAccountsResponseBodyResultPagination struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainAccountsResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainAccountsResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainAccountsResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainAccountsResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainAccountsResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainAccountsResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainAccountsResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainAccountsResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainAccountsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainAccountsResponse) SetStatusCode(v int32) *DescribeAntChainAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainAccountsResponse) SetBody(v *DescribeAntChainAccountsResponseBody) *DescribeAntChainAccountsResponse {
	s.Body = v
	return s
}

type DescribeAntChainAccountsV2Request struct {
	// This parameter is required.
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// M8GaMEyX
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAntChainAccountsV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainAccountsV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainAccountsV2Request) SetAntChainId(v string) *DescribeAntChainAccountsV2Request {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainAccountsV2Request) SetConsortiumId(v string) *DescribeAntChainAccountsV2Request {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainAccountsV2Request) SetPageNumber(v int32) *DescribeAntChainAccountsV2Request {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainAccountsV2Request) SetPageSize(v int32) *DescribeAntChainAccountsV2Request {
	s.PageSize = &v
	return s
}

type DescribeAntChainAccountsV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainAccountsV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainAccountsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainAccountsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainAccountsV2ResponseBody) SetCode(v string) *DescribeAntChainAccountsV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainAccountsV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainAccountsV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainAccountsV2ResponseBody) SetMessage(v string) *DescribeAntChainAccountsV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainAccountsV2ResponseBody) SetRequestId(v string) *DescribeAntChainAccountsV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainAccountsV2ResponseBody) SetResult(v *DescribeAntChainAccountsV2ResponseBodyResult) *DescribeAntChainAccountsV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainAccountsV2ResponseBody) SetResultCode(v string) *DescribeAntChainAccountsV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainAccountsV2ResponseBody) SetResultMessage(v string) *DescribeAntChainAccountsV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainAccountsV2ResponseBody) SetSuccess(v bool) *DescribeAntChainAccountsV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainAccountsV2ResponseBodyResult struct {
	Accounts   []*DescribeAntChainAccountsV2ResponseBodyResultAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	Pagination *DescribeAntChainAccountsV2ResponseBodyResultPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s DescribeAntChainAccountsV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainAccountsV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainAccountsV2ResponseBodyResult) SetAccounts(v []*DescribeAntChainAccountsV2ResponseBodyResultAccounts) *DescribeAntChainAccountsV2ResponseBodyResult {
	s.Accounts = v
	return s
}

func (s *DescribeAntChainAccountsV2ResponseBodyResult) SetPagination(v *DescribeAntChainAccountsV2ResponseBodyResultPagination) *DescribeAntChainAccountsV2ResponseBodyResult {
	s.Pagination = v
	return s
}

type DescribeAntChainAccountsV2ResponseBodyResultAccounts struct {
	// example:
	//
	// test
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// example:
	//
	// 2aa43bae103b6840ce8efdfe6f3fe5e52f8d1db0f44ff762df87ba17eb209979a0e22c934b2728c6c1bab864a6da52de60c5da89793bd839650a1a153e876e32
	AccountPublicKey *string `json:"AccountPublicKey,omitempty" xml:"AccountPublicKey,omitempty"`
	// example:
	//
	// 5a36312d78681794258bb33372586c676adf150ad69e67dbfcaae61bba3607705950bc9efe1bf4a17ac24b05b1615a410e48d2a005dca251c6173495bb47ae29
	AccountRecoveryKey *string `json:"AccountRecoveryKey,omitempty" xml:"AccountRecoveryKey,omitempty"`
	// example:
	//
	// NORMAL
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s DescribeAntChainAccountsV2ResponseBodyResultAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainAccountsV2ResponseBodyResultAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAntChainAccountsV2ResponseBodyResultAccounts) SetAccount(v string) *DescribeAntChainAccountsV2ResponseBodyResultAccounts {
	s.Account = &v
	return s
}

func (s *DescribeAntChainAccountsV2ResponseBodyResultAccounts) SetAccountPublicKey(v string) *DescribeAntChainAccountsV2ResponseBodyResultAccounts {
	s.AccountPublicKey = &v
	return s
}

func (s *DescribeAntChainAccountsV2ResponseBodyResultAccounts) SetAccountRecoveryKey(v string) *DescribeAntChainAccountsV2ResponseBodyResultAccounts {
	s.AccountRecoveryKey = &v
	return s
}

func (s *DescribeAntChainAccountsV2ResponseBodyResultAccounts) SetAccountStatus(v string) *DescribeAntChainAccountsV2ResponseBodyResultAccounts {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAntChainAccountsV2ResponseBodyResultAccounts) SetAntChainId(v string) *DescribeAntChainAccountsV2ResponseBodyResultAccounts {
	s.AntChainId = &v
	return s
}

type DescribeAntChainAccountsV2ResponseBodyResultPagination struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainAccountsV2ResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainAccountsV2ResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainAccountsV2ResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainAccountsV2ResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainAccountsV2ResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainAccountsV2ResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainAccountsV2ResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainAccountsV2ResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainAccountsV2Response struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainAccountsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainAccountsV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainAccountsV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainAccountsV2Response) SetHeaders(v map[string]*string) *DescribeAntChainAccountsV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainAccountsV2Response) SetStatusCode(v int32) *DescribeAntChainAccountsV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainAccountsV2Response) SetBody(v *DescribeAntChainAccountsV2ResponseBody) *DescribeAntChainAccountsV2Response {
	s.Body = v
	return s
}

type DescribeAntChainBlockRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pYogqb9v
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
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
	// example:
	//
	// 4DDA2E72-778A-4D79-BE70-448A26A57227
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
	// example:
	//
	// pYogqb9v
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// aeec5963f8deeeae820aaf302f0c925db9fa8d07b9898dac782335f817554e47
	BlockHash *string `json:"BlockHash,omitempty" xml:"BlockHash,omitempty"`
	// example:
	//
	// 1562851727742
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1234
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// e11767c7b9d92563663a76f10c69e8354788001e2c9e6bd9267239a81d3bb523
	PreviousHash *string `json:"PreviousHash,omitempty" xml:"PreviousHash,omitempty"`
	// example:
	//
	// 0000000000000000000000000000000000000000000000000000000000000000
	RootTxHash *string `json:"RootTxHash,omitempty" xml:"RootTxHash,omitempty"`
	// example:
	//
	// []
	TransSummaryList *string `json:"TransSummaryList,omitempty" xml:"TransSummaryList,omitempty"`
	// example:
	//
	// 0
	TransactionSize *int32 `json:"TransactionSize,omitempty" xml:"TransactionSize,omitempty"`
	// example:
	//
	// 2
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAntChainBlockResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainBlockResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainBlockResponseBodyResult) SetAntChainId(v string) *DescribeAntChainBlockResponseBodyResult {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainBlockResponseBodyResult) SetBlockHash(v string) *DescribeAntChainBlockResponseBodyResult {
	s.BlockHash = &v
	return s
}

func (s *DescribeAntChainBlockResponseBodyResult) SetCreateTime(v int64) *DescribeAntChainBlockResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainBlockResponseBodyResult) SetHeight(v int32) *DescribeAntChainBlockResponseBodyResult {
	s.Height = &v
	return s
}

func (s *DescribeAntChainBlockResponseBodyResult) SetPreviousHash(v string) *DescribeAntChainBlockResponseBodyResult {
	s.PreviousHash = &v
	return s
}

func (s *DescribeAntChainBlockResponseBodyResult) SetRootTxHash(v string) *DescribeAntChainBlockResponseBodyResult {
	s.RootTxHash = &v
	return s
}

func (s *DescribeAntChainBlockResponseBodyResult) SetTransSummaryList(v string) *DescribeAntChainBlockResponseBodyResult {
	s.TransSummaryList = &v
	return s
}

func (s *DescribeAntChainBlockResponseBodyResult) SetTransactionSize(v int32) *DescribeAntChainBlockResponseBodyResult {
	s.TransactionSize = &v
	return s
}

func (s *DescribeAntChainBlockResponseBodyResult) SetVersion(v int64) *DescribeAntChainBlockResponseBodyResult {
	s.Version = &v
	return s
}

type DescribeAntChainBlockResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainBlockResponse) SetStatusCode(v int32) *DescribeAntChainBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainBlockResponse) SetBody(v *DescribeAntChainBlockResponseBody) *DescribeAntChainBlockResponse {
	s.Body = v
	return s
}

type DescribeAntChainBlockV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// M8GaMEyX
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s DescribeAntChainBlockV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainBlockV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainBlockV2Request) SetAntChainId(v string) *DescribeAntChainBlockV2Request {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainBlockV2Request) SetConsortiumId(v string) *DescribeAntChainBlockV2Request {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainBlockV2Request) SetHeight(v int64) *DescribeAntChainBlockV2Request {
	s.Height = &v
	return s
}

type DescribeAntChainBlockV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainBlockV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainBlockV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainBlockV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainBlockV2ResponseBody) SetCode(v string) *DescribeAntChainBlockV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainBlockV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBody) SetMessage(v string) *DescribeAntChainBlockV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBody) SetRequestId(v string) *DescribeAntChainBlockV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBody) SetResult(v *DescribeAntChainBlockV2ResponseBodyResult) *DescribeAntChainBlockV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBody) SetResultCode(v string) *DescribeAntChainBlockV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBody) SetResultMessage(v string) *DescribeAntChainBlockV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBody) SetSuccess(v bool) *DescribeAntChainBlockV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainBlockV2ResponseBodyResult struct {
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// 5abf96debc87f1b64dcbaf4fa57fd46f2a06acdb5de0ba91ef9718d81aebafc7
	BlockHash *string `json:"BlockHash,omitempty" xml:"BlockHash,omitempty"`
	// example:
	//
	// 1609223363570
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 254761
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 2444ef0617e0c6845549dead70f118c5a58f03c04575ecb79e283ab5c34b491d
	PreviousHash *string `json:"PreviousHash,omitempty" xml:"PreviousHash,omitempty"`
	// example:
	//
	// 0000000000000000000000000000000000000000000000000000000000000000
	RootTxHash       *string                                                      `json:"RootTxHash,omitempty" xml:"RootTxHash,omitempty"`
	TransSummaryList []*DescribeAntChainBlockV2ResponseBodyResultTransSummaryList `json:"TransSummaryList,omitempty" xml:"TransSummaryList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TransactionSize *int32 `json:"TransactionSize,omitempty" xml:"TransactionSize,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAntChainBlockV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainBlockV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainBlockV2ResponseBodyResult) SetAntChainId(v string) *DescribeAntChainBlockV2ResponseBodyResult {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResult) SetBlockHash(v string) *DescribeAntChainBlockV2ResponseBodyResult {
	s.BlockHash = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResult) SetCreateTime(v int64) *DescribeAntChainBlockV2ResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResult) SetHeight(v int32) *DescribeAntChainBlockV2ResponseBodyResult {
	s.Height = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResult) SetPreviousHash(v string) *DescribeAntChainBlockV2ResponseBodyResult {
	s.PreviousHash = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResult) SetRootTxHash(v string) *DescribeAntChainBlockV2ResponseBodyResult {
	s.RootTxHash = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResult) SetTransSummaryList(v []*DescribeAntChainBlockV2ResponseBodyResultTransSummaryList) *DescribeAntChainBlockV2ResponseBodyResult {
	s.TransSummaryList = v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResult) SetTransactionSize(v int32) *DescribeAntChainBlockV2ResponseBodyResult {
	s.TransactionSize = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResult) SetVersion(v int64) *DescribeAntChainBlockV2ResponseBodyResult {
	s.Version = &v
	return s
}

type DescribeAntChainBlockV2ResponseBodyResultTransSummaryList struct {
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// example:
	//
	// 5abf96debc87f1b64dcbaf4fa57fd46f2a06acdb5de0ba91ef9718d81aebafc7
	BlockHash *string `json:"BlockHash,omitempty" xml:"BlockHash,omitempty"`
	// example:
	//
	// 0
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 1609223363570
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// e7d3e769f3f593dadcb8634cc5b09fc90dd3a61c4a06a79cb0923662fe6fae6b
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 4000000
	GasUsed *int64 `json:"GasUsed,omitempty" xml:"GasUsed,omitempty"`
	// example:
	//
	// 076bba1b726b3bcb958cba6fffc03eaa5cbed59320271dcbc0e05648a18a94f2
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	// example:
	//
	// 254761
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0
	ReferenceCount *int32 `json:"ReferenceCount,omitempty" xml:"ReferenceCount,omitempty"`
	// example:
	//
	// e93372533f323b2f12783aa3a586135cf421486439c2cdcde47411b78f9839ec
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// example:
	//
	// CALL_CONTRACT
	TransTypeV10 *string `json:"TransTypeV10,omitempty" xml:"TransTypeV10,omitempty"`
	// example:
	//
	// ""
	TransTypeV6 *string `json:"TransTypeV6,omitempty" xml:"TransTypeV6,omitempty"`
}

func (s DescribeAntChainBlockV2ResponseBodyResultTransSummaryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainBlockV2ResponseBodyResultTransSummaryList) GoString() string {
	return s.String()
}

func (s *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList) SetAlias(v string) *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList {
	s.Alias = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList) SetBlockHash(v string) *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList {
	s.BlockHash = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList) SetCategory(v int32) *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList {
	s.Category = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList) SetCreateTime(v int64) *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList) SetFrom(v string) *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList {
	s.From = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList) SetGasUsed(v int64) *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList {
	s.GasUsed = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList) SetHash(v string) *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList {
	s.Hash = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList) SetHeight(v int64) *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList {
	s.Height = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList) SetReferenceCount(v int32) *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList {
	s.ReferenceCount = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList) SetTo(v string) *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList {
	s.To = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList) SetTransTypeV10(v string) *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList {
	s.TransTypeV10 = &v
	return s
}

func (s *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList) SetTransTypeV6(v string) *DescribeAntChainBlockV2ResponseBodyResultTransSummaryList {
	s.TransTypeV6 = &v
	return s
}

type DescribeAntChainBlockV2Response struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainBlockV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainBlockV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainBlockV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainBlockV2Response) SetHeaders(v map[string]*string) *DescribeAntChainBlockV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainBlockV2Response) SetStatusCode(v int32) *DescribeAntChainBlockV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainBlockV2Response) SetBody(v *DescribeAntChainBlockV2ResponseBody) *DescribeAntChainBlockV2Response {
	s.Body = v
	return s
}

type DescribeAntChainCertificateApplicationsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1q8B5R9p
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *DescribeAntChainCertificateApplicationsRequest) SetPageNumber(v int32) *DescribeAntChainCertificateApplicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsRequest) SetPageSize(v int32) *DescribeAntChainCertificateApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsRequest) SetStatus(v string) *DescribeAntChainCertificateApplicationsRequest {
	s.Status = &v
	return s
}

type DescribeAntChainCertificateApplicationsResponseBody struct {
	// example:
	//
	// B232A84E-7944-4EB0-83D0-9D409E12E0A8
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
	CertificateApplications []*DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications `json:"CertificateApplications,omitempty" xml:"CertificateApplications,omitempty" type:"Repeated"`
	Pagination              *DescribeAntChainCertificateApplicationsResponseBodyResultPagination                `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s DescribeAntChainCertificateApplicationsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainCertificateApplicationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResult) SetCertificateApplications(v []*DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) *DescribeAntChainCertificateApplicationsResponseBodyResult {
	s.CertificateApplications = v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResult) SetPagination(v *DescribeAntChainCertificateApplicationsResponseBodyResultPagination) *DescribeAntChainCertificateApplicationsResponseBodyResult {
	s.Pagination = v
	return s
}

type DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications struct {
	// example:
	//
	// 1q8B5R9p
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// Bid
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// example:
	//
	// 1563949275000
	Createtime *int64 `json:"Createtime,omitempty" xml:"Createtime,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1563949275000
	Updatetime *int64 `json:"Updatetime,omitempty" xml:"Updatetime,omitempty"`
	// example:
	//
	// uid-128712635330****
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) GoString() string {
	return s.String()
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) SetAntChainId(v string) *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) SetBid(v string) *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications {
	s.Bid = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) SetCreatetime(v int64) *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications {
	s.Createtime = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) SetStatus(v string) *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications {
	s.Status = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) SetUpdatetime(v int64) *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications {
	s.Updatetime = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications) SetUsername(v string) *DescribeAntChainCertificateApplicationsResponseBodyResultCertificateApplications {
	s.Username = &v
	return s
}

type DescribeAntChainCertificateApplicationsResponseBodyResultPagination struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainCertificateApplicationsResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainCertificateApplicationsResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainCertificateApplicationsResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainCertificateApplicationsResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainCertificateApplicationsResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainCertificateApplicationsResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainCertificateApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainCertificateApplicationsResponse) SetStatusCode(v int32) *DescribeAntChainCertificateApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsResponse) SetBody(v *DescribeAntChainCertificateApplicationsResponseBody) *DescribeAntChainCertificateApplicationsResponse {
	s.Body = v
	return s
}

type DescribeAntChainCertificateApplicationsV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// M8GaMEyX
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAntChainCertificateApplicationsV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainCertificateApplicationsV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainCertificateApplicationsV2Request) SetAntChainId(v string) *DescribeAntChainCertificateApplicationsV2Request {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2Request) SetConsortiumId(v string) *DescribeAntChainCertificateApplicationsV2Request {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2Request) SetPageNumber(v int32) *DescribeAntChainCertificateApplicationsV2Request {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2Request) SetPageSize(v int32) *DescribeAntChainCertificateApplicationsV2Request {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2Request) SetStatus(v string) *DescribeAntChainCertificateApplicationsV2Request {
	s.Status = &v
	return s
}

type DescribeAntChainCertificateApplicationsV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainCertificateApplicationsV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainCertificateApplicationsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainCertificateApplicationsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBody) SetCode(v string) *DescribeAntChainCertificateApplicationsV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainCertificateApplicationsV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBody) SetMessage(v string) *DescribeAntChainCertificateApplicationsV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBody) SetRequestId(v string) *DescribeAntChainCertificateApplicationsV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBody) SetResult(v *DescribeAntChainCertificateApplicationsV2ResponseBodyResult) *DescribeAntChainCertificateApplicationsV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBody) SetResultCode(v string) *DescribeAntChainCertificateApplicationsV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBody) SetResultMessage(v string) *DescribeAntChainCertificateApplicationsV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBody) SetSuccess(v bool) *DescribeAntChainCertificateApplicationsV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainCertificateApplicationsV2ResponseBodyResult struct {
	CertificateApplications []*DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications `json:"CertificateApplications,omitempty" xml:"CertificateApplications,omitempty" type:"Repeated"`
	Pagination              *DescribeAntChainCertificateApplicationsV2ResponseBodyResultPagination                `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s DescribeAntChainCertificateApplicationsV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainCertificateApplicationsV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBodyResult) SetCertificateApplications(v []*DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications) *DescribeAntChainCertificateApplicationsV2ResponseBodyResult {
	s.CertificateApplications = v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBodyResult) SetPagination(v *DescribeAntChainCertificateApplicationsV2ResponseBodyResultPagination) *DescribeAntChainCertificateApplicationsV2ResponseBodyResult {
	s.Pagination = v
	return s
}

type DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications struct {
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// example:
	//
	// 1609848404000
	Createtime *int64 `json:"Createtime,omitempty" xml:"Createtime,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1609848404000
	Updatetime *int64 `json:"Updatetime,omitempty" xml:"Updatetime,omitempty"`
	// example:
	//
	// uid-103477475017****
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications) GoString() string {
	return s.String()
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications) SetAntChainId(v string) *DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications) SetBid(v string) *DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications {
	s.Bid = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications) SetCreatetime(v int64) *DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications {
	s.Createtime = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications) SetStatus(v string) *DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications {
	s.Status = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications) SetUpdatetime(v int64) *DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications {
	s.Updatetime = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications) SetUsername(v string) *DescribeAntChainCertificateApplicationsV2ResponseBodyResultCertificateApplications {
	s.Username = &v
	return s
}

type DescribeAntChainCertificateApplicationsV2ResponseBodyResultPagination struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainCertificateApplicationsV2ResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainCertificateApplicationsV2ResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainCertificateApplicationsV2ResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainCertificateApplicationsV2ResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2ResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainCertificateApplicationsV2ResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainCertificateApplicationsV2Response struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainCertificateApplicationsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainCertificateApplicationsV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainCertificateApplicationsV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainCertificateApplicationsV2Response) SetHeaders(v map[string]*string) *DescribeAntChainCertificateApplicationsV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2Response) SetStatusCode(v int32) *DescribeAntChainCertificateApplicationsV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainCertificateApplicationsV2Response) SetBody(v *DescribeAntChainCertificateApplicationsV2ResponseBody) *DescribeAntChainCertificateApplicationsV2Response {
	s.Body = v
	return s
}

type DescribeAntChainConsortiumsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAntChainConsortiumsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainConsortiumsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainConsortiumsRequest) SetPageNumber(v int32) *DescribeAntChainConsortiumsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainConsortiumsRequest) SetPageSize(v int32) *DescribeAntChainConsortiumsRequest {
	s.PageSize = &v
	return s
}

type DescribeAntChainConsortiumsResponseBody struct {
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
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
	// example:
	//
	// 1
	ChainNum              *int64  `json:"ChainNum,omitempty" xml:"ChainNum,omitempty"`
	ConsortiumDescription *string `json:"ConsortiumDescription,omitempty" xml:"ConsortiumDescription,omitempty"`
	// example:
	//
	// DV80nJXq
	ConsortiumId   *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	// example:
	//
	// 1562845453000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2
	MemberNum *int64 `json:"MemberNum,omitempty" xml:"MemberNum,omitempty"`
	// example:
	//
	// SuperAdmin
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) GoString() string {
	return s.String()
}

func (s *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) SetChainNum(v int64) *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums {
	s.ChainNum = &v
	return s
}

func (s *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) SetConsortiumDescription(v string) *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums {
	s.ConsortiumDescription = &v
	return s
}

func (s *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) SetConsortiumId(v string) *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) SetConsortiumName(v string) *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums {
	s.ConsortiumName = &v
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

func (s *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums) SetStatus(v string) *DescribeAntChainConsortiumsResponseBodyResultAntConsortiums {
	s.Status = &v
	return s
}

type DescribeAntChainConsortiumsResponseBodyResultPagination struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainConsortiumsResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainConsortiumsResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainConsortiumsResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainConsortiumsResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainConsortiumsResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainConsortiumsResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainConsortiumsResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainConsortiumsResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainConsortiumsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainConsortiumsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainConsortiumsResponse) SetStatusCode(v int32) *DescribeAntChainConsortiumsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainConsortiumsResponse) SetBody(v *DescribeAntChainConsortiumsResponseBody) *DescribeAntChainConsortiumsResponse {
	s.Body = v
	return s
}

type DescribeAntChainConsortiumsV2Request struct {
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAntChainConsortiumsV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainConsortiumsV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainConsortiumsV2Request) SetPageNumber(v int32) *DescribeAntChainConsortiumsV2Request {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2Request) SetPageSize(v int32) *DescribeAntChainConsortiumsV2Request {
	s.PageSize = &v
	return s
}

type DescribeAntChainConsortiumsV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainConsortiumsV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainConsortiumsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainConsortiumsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainConsortiumsV2ResponseBody) SetCode(v string) *DescribeAntChainConsortiumsV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainConsortiumsV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBody) SetMessage(v string) *DescribeAntChainConsortiumsV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBody) SetRequestId(v string) *DescribeAntChainConsortiumsV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBody) SetResult(v *DescribeAntChainConsortiumsV2ResponseBodyResult) *DescribeAntChainConsortiumsV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBody) SetResultCode(v string) *DescribeAntChainConsortiumsV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBody) SetResultMessage(v string) *DescribeAntChainConsortiumsV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBody) SetSuccess(v bool) *DescribeAntChainConsortiumsV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainConsortiumsV2ResponseBodyResult struct {
	AntConsortiums []*DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums `json:"AntConsortiums,omitempty" xml:"AntConsortiums,omitempty" type:"Repeated"`
	Pagination     *DescribeAntChainConsortiumsV2ResponseBodyResultPagination       `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s DescribeAntChainConsortiumsV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainConsortiumsV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainConsortiumsV2ResponseBodyResult) SetAntConsortiums(v []*DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums) *DescribeAntChainConsortiumsV2ResponseBodyResult {
	s.AntConsortiums = v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBodyResult) SetPagination(v *DescribeAntChainConsortiumsV2ResponseBodyResultPagination) *DescribeAntChainConsortiumsV2ResponseBodyResult {
	s.Pagination = v
	return s
}

type DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums struct {
	// example:
	//
	// 2
	ChainNum *int64 `json:"ChainNum,omitempty" xml:"ChainNum,omitempty"`
	// example:
	//
	// test
	ConsortiumDescription *string `json:"ConsortiumDescription,omitempty" xml:"ConsortiumDescription,omitempty"`
	// example:
	//
	// M8GaMEyX
	ConsortiumId   *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	// example:
	//
	// 1609745002000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// false
	IsEmptyConsortium *bool `json:"IsEmptyConsortium,omitempty" xml:"IsEmptyConsortium,omitempty"`
	// example:
	//
	// 2
	MemberNum *int64 `json:"MemberNum,omitempty" xml:"MemberNum,omitempty"`
	// example:
	//
	// Member
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums) GoString() string {
	return s.String()
}

func (s *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums) SetChainNum(v int64) *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums {
	s.ChainNum = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums) SetConsortiumDescription(v string) *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums {
	s.ConsortiumDescription = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums) SetConsortiumId(v string) *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums) SetConsortiumName(v string) *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums {
	s.ConsortiumName = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums) SetCreateTime(v int64) *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums) SetIsEmptyConsortium(v bool) *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums {
	s.IsEmptyConsortium = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums) SetMemberNum(v int64) *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums {
	s.MemberNum = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums) SetRole(v string) *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums {
	s.Role = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums) SetStatus(v string) *DescribeAntChainConsortiumsV2ResponseBodyResultAntConsortiums {
	s.Status = &v
	return s
}

type DescribeAntChainConsortiumsV2ResponseBodyResultPagination struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainConsortiumsV2ResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainConsortiumsV2ResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainConsortiumsV2ResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainConsortiumsV2ResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainConsortiumsV2ResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2ResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainConsortiumsV2ResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainConsortiumsV2Response struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainConsortiumsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainConsortiumsV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainConsortiumsV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainConsortiumsV2Response) SetHeaders(v map[string]*string) *DescribeAntChainConsortiumsV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainConsortiumsV2Response) SetStatusCode(v int32) *DescribeAntChainConsortiumsV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainConsortiumsV2Response) SetBody(v *DescribeAntChainConsortiumsV2ResponseBody) *DescribeAntChainConsortiumsV2Response {
	s.Body = v
	return s
}

type DescribeAntChainContractProjectContentTreeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2L9VK68g
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
	// example:
	//
	// A0DF307A-A44A-476F-99D8-BA4F205350BD
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
	Children []map[string]interface{} `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// example:
	//
	// description
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
	// example:
	//
	// 2L9VK68g
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// projectname
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 5
	ProjectVersion *string `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
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

func (s *DescribeAntChainContractProjectContentTreeResponseBodyResult) SetProjectDescription(v string) *DescribeAntChainContractProjectContentTreeResponseBodyResult {
	s.ProjectDescription = &v
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

type DescribeAntChainContractProjectContentTreeResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainContractProjectContentTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainContractProjectContentTreeResponse) SetStatusCode(v int32) *DescribeAntChainContractProjectContentTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeResponse) SetBody(v *DescribeAntChainContractProjectContentTreeResponseBody) *DescribeAntChainContractProjectContentTreeResponse {
	s.Body = v
	return s
}

type DescribeAntChainContractProjectContentTreeV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// M8GaMEyX
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2L9VK68g
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeAntChainContractProjectContentTreeV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectContentTreeV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectContentTreeV2Request) SetConsortiumId(v string) *DescribeAntChainContractProjectContentTreeV2Request {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeV2Request) SetProjectId(v string) *DescribeAntChainContractProjectContentTreeV2Request {
	s.ProjectId = &v
	return s
}

type DescribeAntChainContractProjectContentTreeV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// "{"description":"","projectId":"","projectName":"","treeContractContent":[{"content":"","fileName":"","id":"","isDir":false,"parentId\\":""}],"type":"","version":""}"
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainContractProjectContentTreeV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectContentTreeV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectContentTreeV2ResponseBody) SetCode(v string) *DescribeAntChainContractProjectContentTreeV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainContractProjectContentTreeV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeV2ResponseBody) SetMessage(v string) *DescribeAntChainContractProjectContentTreeV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeV2ResponseBody) SetRequestId(v string) *DescribeAntChainContractProjectContentTreeV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeV2ResponseBody) SetResult(v string) *DescribeAntChainContractProjectContentTreeV2ResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeV2ResponseBody) SetResultCode(v string) *DescribeAntChainContractProjectContentTreeV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeV2ResponseBody) SetResultMessage(v string) *DescribeAntChainContractProjectContentTreeV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeV2ResponseBody) SetSuccess(v bool) *DescribeAntChainContractProjectContentTreeV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainContractProjectContentTreeV2Response struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainContractProjectContentTreeV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainContractProjectContentTreeV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectContentTreeV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectContentTreeV2Response) SetHeaders(v map[string]*string) *DescribeAntChainContractProjectContentTreeV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeV2Response) SetStatusCode(v int32) *DescribeAntChainContractProjectContentTreeV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainContractProjectContentTreeV2Response) SetBody(v *DescribeAntChainContractProjectContentTreeV2ResponseBody) *DescribeAntChainContractProjectContentTreeV2Response {
	s.Body = v
	return s
}

type DescribeAntChainContractProjectsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DV80nJXq
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *DescribeAntChainContractProjectsRequest) SetPageNumber(v int32) *DescribeAntChainContractProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainContractProjectsRequest) SetPageSize(v int32) *DescribeAntChainContractProjectsRequest {
	s.PageSize = &v
	return s
}

type DescribeAntChainContractProjectsResponseBody struct {
	// example:
	//
	// 5EDAC16F-EA3E-4017-B223-D76220F985DF
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
	ContractProjects []*DescribeAntChainContractProjectsResponseBodyResultContractProjects `json:"ContractProjects,omitempty" xml:"ContractProjects,omitempty" type:"Repeated"`
	Pagination       *DescribeAntChainContractProjectsResponseBodyResultPagination         `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s DescribeAntChainContractProjectsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectsResponseBodyResult) SetContractProjects(v []*DescribeAntChainContractProjectsResponseBodyResultContractProjects) *DescribeAntChainContractProjectsResponseBodyResult {
	s.ContractProjects = v
	return s
}

func (s *DescribeAntChainContractProjectsResponseBodyResult) SetPagination(v *DescribeAntChainContractProjectsResponseBodyResultPagination) *DescribeAntChainContractProjectsResponseBodyResult {
	s.Pagination = v
	return s
}

type DescribeAntChainContractProjectsResponseBodyResultContractProjects struct {
	// example:
	//
	// DV80nJXq
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1562847564000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// description
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
	// example:
	//
	// 2L9VK68g
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// projectname
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 5
	ProjectVersion *string `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
	// example:
	//
	// 1563954499000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeAntChainContractProjectsResponseBodyResultContractProjects) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectsResponseBodyResultContractProjects) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectsResponseBodyResultContractProjects) SetConsortiumId(v string) *DescribeAntChainContractProjectsResponseBodyResultContractProjects {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainContractProjectsResponseBodyResultContractProjects) SetCreateTime(v int64) *DescribeAntChainContractProjectsResponseBodyResultContractProjects {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainContractProjectsResponseBodyResultContractProjects) SetProjectDescription(v string) *DescribeAntChainContractProjectsResponseBodyResultContractProjects {
	s.ProjectDescription = &v
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

func (s *DescribeAntChainContractProjectsResponseBodyResultContractProjects) SetUpdateTime(v int64) *DescribeAntChainContractProjectsResponseBodyResultContractProjects {
	s.UpdateTime = &v
	return s
}

type DescribeAntChainContractProjectsResponseBodyResultPagination struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainContractProjectsResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectsResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectsResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainContractProjectsResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainContractProjectsResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainContractProjectsResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainContractProjectsResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainContractProjectsResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainContractProjectsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainContractProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainContractProjectsResponse) SetStatusCode(v int32) *DescribeAntChainContractProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainContractProjectsResponse) SetBody(v *DescribeAntChainContractProjectsResponseBody) *DescribeAntChainContractProjectsResponse {
	s.Body = v
	return s
}

type DescribeAntChainContractProjectsV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// M8GaMEyX
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAntChainContractProjectsV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectsV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectsV2Request) SetConsortiumId(v string) *DescribeAntChainContractProjectsV2Request {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2Request) SetPageNumber(v int32) *DescribeAntChainContractProjectsV2Request {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2Request) SetPageSize(v int32) *DescribeAntChainContractProjectsV2Request {
	s.PageSize = &v
	return s
}

type DescribeAntChainContractProjectsV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainContractProjectsV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainContractProjectsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectsV2ResponseBody) SetCode(v string) *DescribeAntChainContractProjectsV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainContractProjectsV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2ResponseBody) SetMessage(v string) *DescribeAntChainContractProjectsV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2ResponseBody) SetRequestId(v string) *DescribeAntChainContractProjectsV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2ResponseBody) SetResult(v *DescribeAntChainContractProjectsV2ResponseBodyResult) *DescribeAntChainContractProjectsV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainContractProjectsV2ResponseBody) SetResultCode(v string) *DescribeAntChainContractProjectsV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2ResponseBody) SetResultMessage(v string) *DescribeAntChainContractProjectsV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2ResponseBody) SetSuccess(v bool) *DescribeAntChainContractProjectsV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainContractProjectsV2ResponseBodyResult struct {
	ContractProjects []*DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects `json:"ContractProjects,omitempty" xml:"ContractProjects,omitempty" type:"Repeated"`
	Pagination       *DescribeAntChainContractProjectsV2ResponseBodyResultPagination         `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s DescribeAntChainContractProjectsV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectsV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectsV2ResponseBodyResult) SetContractProjects(v []*DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects) *DescribeAntChainContractProjectsV2ResponseBodyResult {
	s.ContractProjects = v
	return s
}

func (s *DescribeAntChainContractProjectsV2ResponseBodyResult) SetPagination(v *DescribeAntChainContractProjectsV2ResponseBodyResultPagination) *DescribeAntChainContractProjectsV2ResponseBodyResult {
	s.Pagination = v
	return s
}

type DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects struct {
	// example:
	//
	// M8GaMEyX
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1609848383000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// test
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
	// example:
	//
	// RXwQj6m8
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 1.0.0
	ProjectVersion *string `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
	// example:
	//
	// 1609848383000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects) SetConsortiumId(v string) *DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects) SetCreateTime(v int64) *DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects) SetProjectDescription(v string) *DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects {
	s.ProjectDescription = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects) SetProjectId(v string) *DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects {
	s.ProjectId = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects) SetProjectName(v string) *DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects {
	s.ProjectName = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects) SetProjectVersion(v string) *DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects {
	s.ProjectVersion = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects) SetUpdateTime(v int64) *DescribeAntChainContractProjectsV2ResponseBodyResultContractProjects {
	s.UpdateTime = &v
	return s
}

type DescribeAntChainContractProjectsV2ResponseBodyResultPagination struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainContractProjectsV2ResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectsV2ResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectsV2ResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainContractProjectsV2ResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2ResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainContractProjectsV2ResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2ResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainContractProjectsV2ResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainContractProjectsV2Response struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainContractProjectsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainContractProjectsV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainContractProjectsV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainContractProjectsV2Response) SetHeaders(v map[string]*string) *DescribeAntChainContractProjectsV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainContractProjectsV2Response) SetStatusCode(v int32) *DescribeAntChainContractProjectsV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainContractProjectsV2Response) SetBody(v *DescribeAntChainContractProjectsV2ResponseBody) *DescribeAntChainContractProjectsV2Response {
	s.Body = v
	return s
}

type DescribeAntChainDownloadPathsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1q8B5R9p
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
	// example:
	//
	// E6F487D4-8606-41B5-B289-46D5EFBD3417
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
	// example:
	//
	// http://***ca.crt
	CaCrtUrl *string `json:"CaCrtUrl,omitempty" xml:"CaCrtUrl,omitempty"`
	// example:
	//
	// http://***client.crt
	ClientCrtUrl *string `json:"ClientCrtUrl,omitempty" xml:"ClientCrtUrl,omitempty"`
	// example:
	//
	// http://***
	SdkUrl *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
	// example:
	//
	// http://***trustCa
	TrustCaUrl *string `json:"TrustCaUrl,omitempty" xml:"TrustCaUrl,omitempty"`
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

func (s *DescribeAntChainDownloadPathsResponseBodyResult) SetClientCrtUrl(v string) *DescribeAntChainDownloadPathsResponseBodyResult {
	s.ClientCrtUrl = &v
	return s
}

func (s *DescribeAntChainDownloadPathsResponseBodyResult) SetSdkUrl(v string) *DescribeAntChainDownloadPathsResponseBodyResult {
	s.SdkUrl = &v
	return s
}

func (s *DescribeAntChainDownloadPathsResponseBodyResult) SetTrustCaUrl(v string) *DescribeAntChainDownloadPathsResponseBodyResult {
	s.TrustCaUrl = &v
	return s
}

type DescribeAntChainDownloadPathsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainDownloadPathsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainDownloadPathsResponse) SetStatusCode(v int32) *DescribeAntChainDownloadPathsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainDownloadPathsResponse) SetBody(v *DescribeAntChainDownloadPathsResponseBody) *DescribeAntChainDownloadPathsResponse {
	s.Body = v
	return s
}

type DescribeAntChainDownloadPathsV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// M8GaMEyX
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
}

func (s DescribeAntChainDownloadPathsV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainDownloadPathsV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainDownloadPathsV2Request) SetAntChainId(v string) *DescribeAntChainDownloadPathsV2Request {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainDownloadPathsV2Request) SetConsortiumId(v string) *DescribeAntChainDownloadPathsV2Request {
	s.ConsortiumId = &v
	return s
}

type DescribeAntChainDownloadPathsV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainDownloadPathsV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainDownloadPathsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainDownloadPathsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainDownloadPathsV2ResponseBody) SetCode(v string) *DescribeAntChainDownloadPathsV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainDownloadPathsV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainDownloadPathsV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainDownloadPathsV2ResponseBody) SetMessage(v string) *DescribeAntChainDownloadPathsV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainDownloadPathsV2ResponseBody) SetRequestId(v string) *DescribeAntChainDownloadPathsV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainDownloadPathsV2ResponseBody) SetResult(v *DescribeAntChainDownloadPathsV2ResponseBodyResult) *DescribeAntChainDownloadPathsV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainDownloadPathsV2ResponseBody) SetResultCode(v string) *DescribeAntChainDownloadPathsV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainDownloadPathsV2ResponseBody) SetResultMessage(v string) *DescribeAntChainDownloadPathsV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainDownloadPathsV2ResponseBody) SetSuccess(v bool) *DescribeAntChainDownloadPathsV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainDownloadPathsV2ResponseBodyResult struct {
	// example:
	//
	// http://***ca.crt
	CaCrtUrl *string `json:"CaCrtUrl,omitempty" xml:"CaCrtUrl,omitempty"`
	// example:
	//
	// http://***client.crt
	ClientCrtUrl *string `json:"ClientCrtUrl,omitempty" xml:"ClientCrtUrl,omitempty"`
	// example:
	//
	// http://***
	SdkUrl *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
	// example:
	//
	// http://***trustCa
	TrustCaUrl *string `json:"TrustCaUrl,omitempty" xml:"TrustCaUrl,omitempty"`
}

func (s DescribeAntChainDownloadPathsV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainDownloadPathsV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainDownloadPathsV2ResponseBodyResult) SetCaCrtUrl(v string) *DescribeAntChainDownloadPathsV2ResponseBodyResult {
	s.CaCrtUrl = &v
	return s
}

func (s *DescribeAntChainDownloadPathsV2ResponseBodyResult) SetClientCrtUrl(v string) *DescribeAntChainDownloadPathsV2ResponseBodyResult {
	s.ClientCrtUrl = &v
	return s
}

func (s *DescribeAntChainDownloadPathsV2ResponseBodyResult) SetSdkUrl(v string) *DescribeAntChainDownloadPathsV2ResponseBodyResult {
	s.SdkUrl = &v
	return s
}

func (s *DescribeAntChainDownloadPathsV2ResponseBodyResult) SetTrustCaUrl(v string) *DescribeAntChainDownloadPathsV2ResponseBodyResult {
	s.TrustCaUrl = &v
	return s
}

type DescribeAntChainDownloadPathsV2Response struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainDownloadPathsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainDownloadPathsV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainDownloadPathsV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainDownloadPathsV2Response) SetHeaders(v map[string]*string) *DescribeAntChainDownloadPathsV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainDownloadPathsV2Response) SetStatusCode(v int32) *DescribeAntChainDownloadPathsV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainDownloadPathsV2Response) SetBody(v *DescribeAntChainDownloadPathsV2ResponseBody) *DescribeAntChainDownloadPathsV2Response {
	s.Body = v
	return s
}

type DescribeAntChainInformationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pYogqb9v
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
	// example:
	//
	// 91246411-C7E9-4EFD-9074-3939AEDFC1BB
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
	// example:
	//
	// 0
	AbnormalNodes *int32 `json:"AbnormalNodes,omitempty" xml:"AbnormalNodes,omitempty"`
	// example:
	//
	// pYogqb9v
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// 365236
	BlockHeight *int32 `json:"BlockHeight,omitempty" xml:"BlockHeight,omitempty"`
	// example:
	//
	// 1562847396000
	CreateTime *int64                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	NodeInfos  []*DescribeAntChainInformationResponseBodyResultNodeInfos `json:"NodeInfos,omitempty" xml:"NodeInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 4
	NodeNumber *int32 `json:"NodeNumber,omitempty" xml:"NodeNumber,omitempty"`
	// example:
	//
	// true
	Normal *bool `json:"Normal,omitempty" xml:"Normal,omitempty"`
	// example:
	//
	// 42
	TransactionSum *int32 `json:"TransactionSum,omitempty" xml:"TransactionSum,omitempty"`
	// example:
	//
	// 0.10.2.4.7
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAntChainInformationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainInformationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainInformationResponseBodyResult) SetAbnormalNodes(v int32) *DescribeAntChainInformationResponseBodyResult {
	s.AbnormalNodes = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResult) SetAntChainId(v string) *DescribeAntChainInformationResponseBodyResult {
	s.AntChainId = &v
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

func (s *DescribeAntChainInformationResponseBodyResult) SetNodeInfos(v []*DescribeAntChainInformationResponseBodyResultNodeInfos) *DescribeAntChainInformationResponseBodyResult {
	s.NodeInfos = v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResult) SetNodeNumber(v int32) *DescribeAntChainInformationResponseBodyResult {
	s.NodeNumber = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResult) SetNormal(v bool) *DescribeAntChainInformationResponseBodyResult {
	s.Normal = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResult) SetTransactionSum(v int32) *DescribeAntChainInformationResponseBodyResult {
	s.TransactionSum = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResult) SetVersion(v string) *DescribeAntChainInformationResponseBodyResult {
	s.Version = &v
	return s
}

type DescribeAntChainInformationResponseBodyResultNodeInfos struct {
	// example:
	//
	// 365216
	BlockHeight *int64 `json:"BlockHeight,omitempty" xml:"BlockHeight,omitempty"`
	// example:
	//
	// 121.199.195.8 18130
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0.10
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAntChainInformationResponseBodyResultNodeInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainInformationResponseBodyResultNodeInfos) GoString() string {
	return s.String()
}

func (s *DescribeAntChainInformationResponseBodyResultNodeInfos) SetBlockHeight(v int64) *DescribeAntChainInformationResponseBodyResultNodeInfos {
	s.BlockHeight = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResultNodeInfos) SetNodeName(v string) *DescribeAntChainInformationResponseBodyResultNodeInfos {
	s.NodeName = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResultNodeInfos) SetStatus(v bool) *DescribeAntChainInformationResponseBodyResultNodeInfos {
	s.Status = &v
	return s
}

func (s *DescribeAntChainInformationResponseBodyResultNodeInfos) SetVersion(v string) *DescribeAntChainInformationResponseBodyResultNodeInfos {
	s.Version = &v
	return s
}

type DescribeAntChainInformationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainInformationResponse) SetStatusCode(v int32) *DescribeAntChainInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainInformationResponse) SetBody(v *DescribeAntChainInformationResponseBody) *DescribeAntChainInformationResponse {
	s.Body = v
	return s
}

type DescribeAntChainInformationV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// M8GaMEyX
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
}

func (s DescribeAntChainInformationV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainInformationV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainInformationV2Request) SetAntChainId(v string) *DescribeAntChainInformationV2Request {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainInformationV2Request) SetConsortiumId(v string) *DescribeAntChainInformationV2Request {
	s.ConsortiumId = &v
	return s
}

type DescribeAntChainInformationV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainInformationV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainInformationV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainInformationV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainInformationV2ResponseBody) SetCode(v string) *DescribeAntChainInformationV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainInformationV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBody) SetMessage(v string) *DescribeAntChainInformationV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBody) SetRequestId(v string) *DescribeAntChainInformationV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBody) SetResult(v *DescribeAntChainInformationV2ResponseBodyResult) *DescribeAntChainInformationV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBody) SetResultCode(v string) *DescribeAntChainInformationV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBody) SetResultMessage(v string) *DescribeAntChainInformationV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBody) SetSuccess(v bool) *DescribeAntChainInformationV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainInformationV2ResponseBodyResult struct {
	// example:
	//
	// 0
	AbnormalNodes *int32 `json:"AbnormalNodes,omitempty" xml:"AbnormalNodes,omitempty"`
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// 259988
	BlockHeight *int32 `json:"BlockHeight,omitempty" xml:"BlockHeight,omitempty"`
	// example:
	//
	// 1609221924000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// false
	IsRole    *bool                                                       `json:"IsRole,omitempty" xml:"IsRole,omitempty"`
	NodeInfos []*DescribeAntChainInformationV2ResponseBodyResultNodeInfos `json:"NodeInfos,omitempty" xml:"NodeInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 4
	NodeNumber *int32 `json:"NodeNumber,omitempty" xml:"NodeNumber,omitempty"`
	// example:
	//
	// true
	Normal *bool `json:"Normal,omitempty" xml:"Normal,omitempty"`
	// example:
	//
	// 6
	TransactionSum *int32 `json:"TransactionSum,omitempty" xml:"TransactionSum,omitempty"`
	// example:
	//
	// 0.10
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAntChainInformationV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainInformationV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainInformationV2ResponseBodyResult) SetAbnormalNodes(v int32) *DescribeAntChainInformationV2ResponseBodyResult {
	s.AbnormalNodes = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBodyResult) SetAntChainId(v string) *DescribeAntChainInformationV2ResponseBodyResult {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBodyResult) SetBlockHeight(v int32) *DescribeAntChainInformationV2ResponseBodyResult {
	s.BlockHeight = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBodyResult) SetCreateTime(v int64) *DescribeAntChainInformationV2ResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBodyResult) SetIsRole(v bool) *DescribeAntChainInformationV2ResponseBodyResult {
	s.IsRole = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBodyResult) SetNodeInfos(v []*DescribeAntChainInformationV2ResponseBodyResultNodeInfos) *DescribeAntChainInformationV2ResponseBodyResult {
	s.NodeInfos = v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBodyResult) SetNodeNumber(v int32) *DescribeAntChainInformationV2ResponseBodyResult {
	s.NodeNumber = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBodyResult) SetNormal(v bool) *DescribeAntChainInformationV2ResponseBodyResult {
	s.Normal = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBodyResult) SetTransactionSum(v int32) *DescribeAntChainInformationV2ResponseBodyResult {
	s.TransactionSum = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBodyResult) SetVersion(v string) *DescribeAntChainInformationV2ResponseBodyResult {
	s.Version = &v
	return s
}

type DescribeAntChainInformationV2ResponseBodyResultNodeInfos struct {
	// example:
	//
	// 259988
	BlockHeight *int64 `json:"BlockHeight,omitempty" xml:"BlockHeight,omitempty"`
	// example:
	//
	// 8.136.158.115 18130
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0.10
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAntChainInformationV2ResponseBodyResultNodeInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainInformationV2ResponseBodyResultNodeInfos) GoString() string {
	return s.String()
}

func (s *DescribeAntChainInformationV2ResponseBodyResultNodeInfos) SetBlockHeight(v int64) *DescribeAntChainInformationV2ResponseBodyResultNodeInfos {
	s.BlockHeight = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBodyResultNodeInfos) SetNodeName(v string) *DescribeAntChainInformationV2ResponseBodyResultNodeInfos {
	s.NodeName = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBodyResultNodeInfos) SetStatus(v bool) *DescribeAntChainInformationV2ResponseBodyResultNodeInfos {
	s.Status = &v
	return s
}

func (s *DescribeAntChainInformationV2ResponseBodyResultNodeInfos) SetVersion(v string) *DescribeAntChainInformationV2ResponseBodyResultNodeInfos {
	s.Version = &v
	return s
}

type DescribeAntChainInformationV2Response struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainInformationV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainInformationV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainInformationV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainInformationV2Response) SetHeaders(v map[string]*string) *DescribeAntChainInformationV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainInformationV2Response) SetStatusCode(v int32) *DescribeAntChainInformationV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainInformationV2Response) SetBody(v *DescribeAntChainInformationV2ResponseBody) *DescribeAntChainInformationV2Response {
	s.Body = v
	return s
}

type DescribeAntChainLatestBlocksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pYogqb9v
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
	// example:
	//
	// B52C552F-DDE6-4779-9BA2-9DD455275528
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainLatestBlocksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainLatestBlocksResponse) SetStatusCode(v int32) *DescribeAntChainLatestBlocksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainLatestBlocksResponse) SetBody(v *DescribeAntChainLatestBlocksResponseBody) *DescribeAntChainLatestBlocksResponse {
	s.Body = v
	return s
}

type DescribeAntChainLatestBlocksV2Request struct {
	// This parameter is required.
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// M8GaMEyX
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
}

func (s DescribeAntChainLatestBlocksV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainLatestBlocksV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainLatestBlocksV2Request) SetAntChainId(v string) *DescribeAntChainLatestBlocksV2Request {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2Request) SetConsortiumId(v string) *DescribeAntChainLatestBlocksV2Request {
	s.ConsortiumId = &v
	return s
}

type DescribeAntChainLatestBlocksV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeAntChainLatestBlocksV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainLatestBlocksV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainLatestBlocksV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainLatestBlocksV2ResponseBody) SetCode(v string) *DescribeAntChainLatestBlocksV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainLatestBlocksV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2ResponseBody) SetMessage(v string) *DescribeAntChainLatestBlocksV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2ResponseBody) SetRequestId(v string) *DescribeAntChainLatestBlocksV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2ResponseBody) SetResult(v []*DescribeAntChainLatestBlocksV2ResponseBodyResult) *DescribeAntChainLatestBlocksV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainLatestBlocksV2ResponseBody) SetResultCode(v string) *DescribeAntChainLatestBlocksV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2ResponseBody) SetResultMessage(v string) *DescribeAntChainLatestBlocksV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2ResponseBody) SetSuccess(v bool) *DescribeAntChainLatestBlocksV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainLatestBlocksV2ResponseBodyResult struct {
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// example:
	//
	// b21baa257788d97eb0006672ccc5008f465365e1dec88d0bbd833e150414b3d5
	BizData *string `json:"BizData,omitempty" xml:"BizData,omitempty"`
	// example:
	//
	// b21baa257788d97eb0006672ccc5008f465365e1dec88d0bbd833e150414b3d5
	BlockHash *string `json:"BlockHash,omitempty" xml:"BlockHash,omitempty"`
	// example:
	//
	// 1610002621000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 259808
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// f208834bdc72bd6bb05c5ef1a35abbc8295a16deda9526b7b78c69ec24591b9f
	PreviousHash *string `json:"PreviousHash,omitempty" xml:"PreviousHash,omitempty"`
	// example:
	//
	// 0000000000000000000000000000000000000000000000000000000000000000
	RootTxHash *string `json:"RootTxHash,omitempty" xml:"RootTxHash,omitempty"`
	// example:
	//
	// 1024
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 0
	TransactionSize *int64 `json:"TransactionSize,omitempty" xml:"TransactionSize,omitempty"`
	// example:
	//
	// 33556226
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAntChainLatestBlocksV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainLatestBlocksV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainLatestBlocksV2ResponseBodyResult) SetAlias(v string) *DescribeAntChainLatestBlocksV2ResponseBodyResult {
	s.Alias = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2ResponseBodyResult) SetBizData(v string) *DescribeAntChainLatestBlocksV2ResponseBodyResult {
	s.BizData = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2ResponseBodyResult) SetBlockHash(v string) *DescribeAntChainLatestBlocksV2ResponseBodyResult {
	s.BlockHash = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2ResponseBodyResult) SetCreateTime(v int64) *DescribeAntChainLatestBlocksV2ResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2ResponseBodyResult) SetHeight(v int64) *DescribeAntChainLatestBlocksV2ResponseBodyResult {
	s.Height = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2ResponseBodyResult) SetPreviousHash(v string) *DescribeAntChainLatestBlocksV2ResponseBodyResult {
	s.PreviousHash = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2ResponseBodyResult) SetRootTxHash(v string) *DescribeAntChainLatestBlocksV2ResponseBodyResult {
	s.RootTxHash = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2ResponseBodyResult) SetSize(v int64) *DescribeAntChainLatestBlocksV2ResponseBodyResult {
	s.Size = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2ResponseBodyResult) SetTransactionSize(v int64) *DescribeAntChainLatestBlocksV2ResponseBodyResult {
	s.TransactionSize = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2ResponseBodyResult) SetVersion(v int64) *DescribeAntChainLatestBlocksV2ResponseBodyResult {
	s.Version = &v
	return s
}

type DescribeAntChainLatestBlocksV2Response struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainLatestBlocksV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainLatestBlocksV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainLatestBlocksV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainLatestBlocksV2Response) SetHeaders(v map[string]*string) *DescribeAntChainLatestBlocksV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainLatestBlocksV2Response) SetStatusCode(v int32) *DescribeAntChainLatestBlocksV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainLatestBlocksV2Response) SetBody(v *DescribeAntChainLatestBlocksV2ResponseBody) *DescribeAntChainLatestBlocksV2Response {
	s.Body = v
	return s
}

type DescribeAntChainLatestTransactionDigestsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pYogqb9v
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
	// example:
	//
	// 3348160E-187B-47BD-B1B9-76AD4AFA4558
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
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainLatestTransactionDigestsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainLatestTransactionDigestsResponse) SetStatusCode(v int32) *DescribeAntChainLatestTransactionDigestsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainLatestTransactionDigestsResponse) SetBody(v *DescribeAntChainLatestTransactionDigestsResponseBody) *DescribeAntChainLatestTransactionDigestsResponse {
	s.Body = v
	return s
}

type DescribeAntChainLatestTransactionDigestsV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// M8GaMEyX
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
}

func (s DescribeAntChainLatestTransactionDigestsV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainLatestTransactionDigestsV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainLatestTransactionDigestsV2Request) SetAntChainId(v string) *DescribeAntChainLatestTransactionDigestsV2Request {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainLatestTransactionDigestsV2Request) SetConsortiumId(v string) *DescribeAntChainLatestTransactionDigestsV2Request {
	s.ConsortiumId = &v
	return s
}

type DescribeAntChainLatestTransactionDigestsV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainLatestTransactionDigestsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainLatestTransactionDigestsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainLatestTransactionDigestsV2ResponseBody) SetCode(v string) *DescribeAntChainLatestTransactionDigestsV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainLatestTransactionDigestsV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainLatestTransactionDigestsV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainLatestTransactionDigestsV2ResponseBody) SetMessage(v string) *DescribeAntChainLatestTransactionDigestsV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainLatestTransactionDigestsV2ResponseBody) SetRequestId(v string) *DescribeAntChainLatestTransactionDigestsV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainLatestTransactionDigestsV2ResponseBody) SetResult(v []*string) *DescribeAntChainLatestTransactionDigestsV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainLatestTransactionDigestsV2ResponseBody) SetResultCode(v string) *DescribeAntChainLatestTransactionDigestsV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainLatestTransactionDigestsV2ResponseBody) SetResultMessage(v string) *DescribeAntChainLatestTransactionDigestsV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainLatestTransactionDigestsV2ResponseBody) SetSuccess(v bool) *DescribeAntChainLatestTransactionDigestsV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainLatestTransactionDigestsV2Response struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainLatestTransactionDigestsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainLatestTransactionDigestsV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainLatestTransactionDigestsV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainLatestTransactionDigestsV2Response) SetHeaders(v map[string]*string) *DescribeAntChainLatestTransactionDigestsV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainLatestTransactionDigestsV2Response) SetStatusCode(v int32) *DescribeAntChainLatestTransactionDigestsV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainLatestTransactionDigestsV2Response) SetBody(v *DescribeAntChainLatestTransactionDigestsV2ResponseBody) *DescribeAntChainLatestTransactionDigestsV2Response {
	s.Body = v
	return s
}

type DescribeAntChainMembersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// q0oWq92P
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAntChainMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMembersRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMembersRequest) SetConsortiumId(v string) *DescribeAntChainMembersRequest {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainMembersRequest) SetPageNumber(v int32) *DescribeAntChainMembersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainMembersRequest) SetPageSize(v int32) *DescribeAntChainMembersRequest {
	s.PageSize = &v
	return s
}

type DescribeAntChainMembersResponseBody struct {
	// example:
	//
	// E0DB75A2-28C8-4295-8851-A157FFB3DFED
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
	// example:
	//
	// 1557477823000
	JoinTime *int64 `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// example:
	//
	// uid-1388601577407805
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// uid-1562589998337656
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// example:
	//
	// Member
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// AllianceNotJoined
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAntChainMembersResponseBodyResultMembers) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMembersResponseBodyResultMembers) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMembersResponseBodyResultMembers) SetJoinTime(v int64) *DescribeAntChainMembersResponseBodyResultMembers {
	s.JoinTime = &v
	return s
}

func (s *DescribeAntChainMembersResponseBodyResultMembers) SetMemberId(v string) *DescribeAntChainMembersResponseBodyResultMembers {
	s.MemberId = &v
	return s
}

func (s *DescribeAntChainMembersResponseBodyResultMembers) SetMemberName(v string) *DescribeAntChainMembersResponseBodyResultMembers {
	s.MemberName = &v
	return s
}

func (s *DescribeAntChainMembersResponseBodyResultMembers) SetRole(v string) *DescribeAntChainMembersResponseBodyResultMembers {
	s.Role = &v
	return s
}

func (s *DescribeAntChainMembersResponseBodyResultMembers) SetStatus(v string) *DescribeAntChainMembersResponseBodyResultMembers {
	s.Status = &v
	return s
}

type DescribeAntChainMembersResponseBodyResultPagination struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainMembersResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMembersResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMembersResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainMembersResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainMembersResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainMembersResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainMembersResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainMembersResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainMembersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainMembersResponse) SetStatusCode(v int32) *DescribeAntChainMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainMembersResponse) SetBody(v *DescribeAntChainMembersResponseBody) *DescribeAntChainMembersResponse {
	s.Body = v
	return s
}

type DescribeAntChainMembersV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// M8GaMEyX
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAntChainMembersV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMembersV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMembersV2Request) SetConsortiumId(v string) *DescribeAntChainMembersV2Request {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainMembersV2Request) SetPageNumber(v int32) *DescribeAntChainMembersV2Request {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainMembersV2Request) SetPageSize(v int32) *DescribeAntChainMembersV2Request {
	s.PageSize = &v
	return s
}

type DescribeAntChainMembersV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainMembersV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainMembersV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMembersV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMembersV2ResponseBody) SetCode(v string) *DescribeAntChainMembersV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainMembersV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainMembersV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainMembersV2ResponseBody) SetMessage(v string) *DescribeAntChainMembersV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainMembersV2ResponseBody) SetRequestId(v string) *DescribeAntChainMembersV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainMembersV2ResponseBody) SetResult(v *DescribeAntChainMembersV2ResponseBodyResult) *DescribeAntChainMembersV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainMembersV2ResponseBody) SetResultCode(v string) *DescribeAntChainMembersV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainMembersV2ResponseBody) SetResultMessage(v string) *DescribeAntChainMembersV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainMembersV2ResponseBody) SetSuccess(v bool) *DescribeAntChainMembersV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainMembersV2ResponseBodyResult struct {
	Members    []*DescribeAntChainMembersV2ResponseBodyResultMembers  `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	Pagination *DescribeAntChainMembersV2ResponseBodyResultPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s DescribeAntChainMembersV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMembersV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMembersV2ResponseBodyResult) SetMembers(v []*DescribeAntChainMembersV2ResponseBodyResultMembers) *DescribeAntChainMembersV2ResponseBodyResult {
	s.Members = v
	return s
}

func (s *DescribeAntChainMembersV2ResponseBodyResult) SetPagination(v *DescribeAntChainMembersV2ResponseBodyResultPagination) *DescribeAntChainMembersV2ResponseBodyResult {
	s.Pagination = v
	return s
}

type DescribeAntChainMembersV2ResponseBodyResultMembers struct {
	// example:
	//
	// 1609745002000
	JoinTime *int64 `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// example:
	//
	// 1034774750177934
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// uid-1034774750177934
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// example:
	//
	// Member
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// AllianceJoined
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAntChainMembersV2ResponseBodyResultMembers) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMembersV2ResponseBodyResultMembers) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMembersV2ResponseBodyResultMembers) SetJoinTime(v int64) *DescribeAntChainMembersV2ResponseBodyResultMembers {
	s.JoinTime = &v
	return s
}

func (s *DescribeAntChainMembersV2ResponseBodyResultMembers) SetMemberId(v string) *DescribeAntChainMembersV2ResponseBodyResultMembers {
	s.MemberId = &v
	return s
}

func (s *DescribeAntChainMembersV2ResponseBodyResultMembers) SetMemberName(v string) *DescribeAntChainMembersV2ResponseBodyResultMembers {
	s.MemberName = &v
	return s
}

func (s *DescribeAntChainMembersV2ResponseBodyResultMembers) SetRole(v string) *DescribeAntChainMembersV2ResponseBodyResultMembers {
	s.Role = &v
	return s
}

func (s *DescribeAntChainMembersV2ResponseBodyResultMembers) SetStatus(v string) *DescribeAntChainMembersV2ResponseBodyResultMembers {
	s.Status = &v
	return s
}

type DescribeAntChainMembersV2ResponseBodyResultPagination struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainMembersV2ResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMembersV2ResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMembersV2ResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainMembersV2ResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainMembersV2ResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainMembersV2ResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainMembersV2ResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainMembersV2ResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainMembersV2Response struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainMembersV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainMembersV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMembersV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMembersV2Response) SetHeaders(v map[string]*string) *DescribeAntChainMembersV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainMembersV2Response) SetStatusCode(v int32) *DescribeAntChainMembersV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainMembersV2Response) SetBody(v *DescribeAntChainMembersV2ResponseBody) *DescribeAntChainMembersV2Response {
	s.Body = v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAccessLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bDXK6boZ
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MINI_APP_BROWSER_TRANSACTION
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
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
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
	// example:
	//
	// 100
	AccessAlipayAccountCount *int64 `json:"AccessAlipayAccountCount,omitempty" xml:"AccessAlipayAccountCount,omitempty"`
	// example:
	//
	// 1000
	AccessCount *int64 `json:"AccessCount,omitempty" xml:"AccessCount,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBodyResult) SetAccessAlipayAccountCount(v int64) *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBodyResult {
	s.AccessAlipayAccountCount = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBodyResult) SetAccessCount(v int64) *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBodyResult {
	s.AccessCount = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse) SetStatusCode(v int32) *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse) SetBody(v *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponseBody) *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse {
	s.Body = v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MINI_APP_BROWSER_TRANSACTION
	QRCodeType *string `json:"QRCodeType,omitempty" xml:"QRCodeType,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Request) SetAntChainId(v string) *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Request {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Request) SetQRCodeType(v string) *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Request {
	s.QRCodeType = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody) SetCode(v string) *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody) SetMessage(v string) *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody) SetRequestId(v string) *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody) SetResult(v *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBodyResult) *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody) SetResultCode(v string) *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody) SetResultMessage(v string) *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody) SetSuccess(v bool) *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBodyResult struct {
	// example:
	//
	// 10
	AccessAlipayAccountCount *int64 `json:"AccessAlipayAccountCount,omitempty" xml:"AccessAlipayAccountCount,omitempty"`
	// example:
	//
	// 100
	AccessCount *int64 `json:"AccessCount,omitempty" xml:"AccessCount,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBodyResult) SetAccessAlipayAccountCount(v int64) *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBodyResult {
	s.AccessAlipayAccountCount = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBodyResult) SetAccessCount(v int64) *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBodyResult {
	s.AccessCount = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Response struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Response) SetHeaders(v map[string]*string) *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Response) SetStatusCode(v int32) *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Response) SetBody(v *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2ResponseBody) *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Response {
	s.Body = v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bDXK6boZ
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MINI_APP_BROWSER_TRANSACTION
	QRCodeType *string `json:"QRCodeType,omitempty" xml:"QRCodeType,omitempty"`
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

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest) SetPageNumber(v int32) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest) SetPageSize(v int32) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest) SetQRCodeType(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest {
	s.QRCodeType = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBody struct {
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
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
	// example:
	//
	// bDXK6boZ
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// ALL_USER_AUTHORIZATION
	AuthorizationType  *string                                                                                    `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	AuthorizedUserList []*DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultAuthorizedUserList `json:"AuthorizedUserList,omitempty" xml:"AuthorizedUserList,omitempty" type:"Repeated"`
	Pagination         *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultPagination           `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	// example:
	//
	// MINI_APP_BROWSER_TRANSACTION
	QRCodeType *string `json:"QRCodeType,omitempty" xml:"QRCodeType,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult) SetAntChainId(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult {
	s.AntChainId = &v
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

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult) SetPagination(v *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultPagination) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult {
	s.Pagination = v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult) SetQRCodeType(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResult {
	s.QRCodeType = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultAuthorizedUserList struct {
	// example:
	//
	// 1580958389000
	GmtAuthorized *string `json:"GmtAuthorized,omitempty" xml:"GmtAuthorized,omitempty"`
	// example:
	//
	// 1380013****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
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

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBodyResultPagination struct {
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
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

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse struct {
	Headers    map[string]*string                                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse) SetStatusCode(v int32) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse) SetBody(v *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponseBody) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse {
	s.Body = v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MINI_APP_BROWSER_TRANSACTION
	QRCodeType *string `json:"QRCodeType,omitempty" xml:"QRCodeType,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Request) SetAntChainId(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Request {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Request) SetPageNumber(v int32) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Request {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Request) SetPageSize(v int32) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Request {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Request) SetQRCodeType(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Request {
	s.QRCodeType = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string                                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody) SetCode(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody) SetMessage(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody) SetRequestId(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody) SetResult(v *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResult) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody) SetResultCode(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody) SetResultMessage(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody) SetSuccess(v bool) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResult struct {
	// example:
	//
	// 4505A1E4-0EDD-4A02-A4D7-B49219328D49
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// SPECIFIC_USER_AUTHORIZATION
	AuthorizationType  *string                                                                                      `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	AuthorizedUserList []*DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultAuthorizedUserList `json:"AuthorizedUserList,omitempty" xml:"AuthorizedUserList,omitempty" type:"Repeated"`
	Pagination         *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultPagination           `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
	// example:
	//
	// MINI_APP_BROWSER_TRANSACTION
	QRCodeType *string `json:"QRCodeType,omitempty" xml:"QRCodeType,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResult) SetAntChainId(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResult {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResult) SetAuthorizationType(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResult {
	s.AuthorizationType = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResult) SetAuthorizedUserList(v []*DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultAuthorizedUserList) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResult {
	s.AuthorizedUserList = v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResult) SetPagination(v *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultPagination) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResult {
	s.Pagination = v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResult) SetQRCodeType(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResult {
	s.QRCodeType = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultAuthorizedUserList struct {
	// example:
	//
	// 2021-01-07 10:55:42
	GmtAuthorized *string `json:"GmtAuthorized,omitempty" xml:"GmtAuthorized,omitempty"`
	// example:
	//
	// 1380013****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultAuthorizedUserList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultAuthorizedUserList) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultAuthorizedUserList) SetGmtAuthorized(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultAuthorizedUserList {
	s.GmtAuthorized = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultAuthorizedUserList) SetPhone(v string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultAuthorizedUserList {
	s.Phone = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultPagination struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Response struct {
	Headers    map[string]*string                                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Response) SetHeaders(v map[string]*string) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Response) SetStatusCode(v int32) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Response) SetBody(v *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2ResponseBody) *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Response {
	s.Body = v
	return s
}

type DescribeAntChainMiniAppBrowserTransactionQRCodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bDXK6boZ
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4a1f1371012cf840a9e4c6bde430d5ba600e422e592d153bd2689783cd1f87d0
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
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
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
	// example:
	//
	// bDXK6boZ
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// iVBORw0KGgo......ABJRU5ErkJggg==
	Base64QRCodePNG *string `json:"Base64QRCodePNG,omitempty" xml:"Base64QRCodePNG,omitempty"`
	// example:
	//
	// https://render.antfin.com/p/s/miniapp-web/?type=trans&from=aliyun&bizid=bDXK6boZ&hash=10692388f8b729cbb42d2985eebd3567812def357eb826a7ec810ce97e1f295e
	QRCodeContent *string `json:"QRCodeContent,omitempty" xml:"QRCodeContent,omitempty"`
	// example:
	//
	// 10692388f8b729cbb42d2985eebd3567812def357eb826a7ec810ce97e1f295e
	TransactionHash *string `json:"TransactionHash,omitempty" xml:"TransactionHash,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult) SetAntChainId(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult) SetBase64QRCodePNG(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult {
	s.Base64QRCodePNG = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult) SetQRCodeContent(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult {
	s.QRCodeContent = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult) SetTransactionHash(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBodyResult {
	s.TransactionHash = &v
	return s
}

type DescribeAntChainMiniAppBrowserTransactionQRCodeResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeResponse) SetStatusCode(v int32) *DescribeAntChainMiniAppBrowserTransactionQRCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeResponse) SetBody(v *DescribeAntChainMiniAppBrowserTransactionQRCodeResponseBody) *DescribeAntChainMiniAppBrowserTransactionQRCodeResponse {
	s.Body = v
	return s
}

type DescribeAntChainMiniAppBrowserTransactionQRCodeNewRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bDXK6boZ
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// 9782291a957d28bec9e02f3dcb27d78392815ede7670b70793e0fd9763309cda
	ContractId *string `json:"ContractId,omitempty" xml:"ContractId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4a1f1371012cf840a9e4c6bde430d5ba600e422e592d153bd2689783cd1f87d0
	TransactionHash *string `json:"TransactionHash,omitempty" xml:"TransactionHash,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeNewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeNewRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewRequest) SetAntChainId(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewRequest {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewRequest) SetContractId(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewRequest {
	s.ContractId = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewRequest) SetTransactionHash(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewRequest {
	s.TransactionHash = &v
	return s
}

type DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody) SetCode(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody) SetHttpStatusCode(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody) SetMessage(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody) SetRequestId(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody) SetResult(v *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBodyResult) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody) SetResultCode(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody) SetResultMessage(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody) SetSuccess(v bool) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBodyResult struct {
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// iVBORw0KGgoAAAANSUhEUgAAAPAAAADwCAIAAACxN37FAAAcb0lEQVR42u3dCZAc1XkAYGlPSRZIGIwEtos4tlM4lE05hhBsU
	Base64QRCodePNG *string `json:"Base64QRCodePNG,omitempty" xml:"Base64QRCodePNG,omitempty"`
	// example:
	//
	// https://render.antfin.com/p/s/miniapp-web/?type=trans&from=aliyun&bizid=8bd720bde18c4b37b0f4a1c7834db163&hash=ac73c8fa158436513e0b398632d9a082e04cee3eac6f9fb50087a46d801bdfd1&contractId=
	QRCodeContent *string `json:"QRCodeContent,omitempty" xml:"QRCodeContent,omitempty"`
	// example:
	//
	// ac73c8fa158436513e0b398632d9a082e04cee3eac6f9fb50087a46d801bdfd1
	TransactionHash *string `json:"TransactionHash,omitempty" xml:"TransactionHash,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBodyResult) SetAntChainId(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBodyResult {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBodyResult) SetBase64QRCodePNG(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBodyResult {
	s.Base64QRCodePNG = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBodyResult) SetQRCodeContent(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBodyResult {
	s.QRCodeContent = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBodyResult) SetTransactionHash(v string) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBodyResult {
	s.TransactionHash = &v
	return s
}

type DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponse struct {
	Headers    map[string]*string                                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponse) SetHeaders(v map[string]*string) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponse) SetStatusCode(v int32) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponse) SetBody(v *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponseBody) *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponse {
	s.Body = v
	return s
}

type DescribeAntChainNodesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1q8B5R9p
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
	// example:
	//
	// BA9738FE-F427-44FD-A1D9-EB4EFB43075A
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainNodesResponse) SetStatusCode(v int32) *DescribeAntChainNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainNodesResponse) SetBody(v *DescribeAntChainNodesResponseBody) *DescribeAntChainNodesResponse {
	s.Body = v
	return s
}

type DescribeAntChainNodesV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// M8GaMEyX
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
}

func (s DescribeAntChainNodesV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainNodesV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainNodesV2Request) SetAntChainId(v string) *DescribeAntChainNodesV2Request {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainNodesV2Request) SetConsortiumId(v string) *DescribeAntChainNodesV2Request {
	s.ConsortiumId = &v
	return s
}

type DescribeAntChainNodesV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainNodesV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainNodesV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainNodesV2ResponseBody) SetCode(v string) *DescribeAntChainNodesV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainNodesV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainNodesV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainNodesV2ResponseBody) SetMessage(v string) *DescribeAntChainNodesV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainNodesV2ResponseBody) SetRequestId(v string) *DescribeAntChainNodesV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainNodesV2ResponseBody) SetResult(v []*string) *DescribeAntChainNodesV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainNodesV2ResponseBody) SetResultCode(v string) *DescribeAntChainNodesV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainNodesV2ResponseBody) SetResultMessage(v string) *DescribeAntChainNodesV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainNodesV2ResponseBody) SetSuccess(v bool) *DescribeAntChainNodesV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainNodesV2Response struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainNodesV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainNodesV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainNodesV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainNodesV2Response) SetHeaders(v map[string]*string) *DescribeAntChainNodesV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainNodesV2Response) SetStatusCode(v int32) *DescribeAntChainNodesV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainNodesV2Response) SetBody(v *DescribeAntChainNodesV2ResponseBody) *DescribeAntChainNodesV2Response {
	s.Body = v
	return s
}

type DescribeAntChainQRCodeAuthorizationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bDXK6boZ
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MINI_APP_BROWSER_TRANSACTION
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
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
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
	// example:
	//
	// bDXK6boZ
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// ALL_USER_AUTHORIZATION
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// example:
	//
	// MINI_APP_BROWSER_TRANSACTION
	QRCodeType *string `json:"QRCodeType,omitempty" xml:"QRCodeType,omitempty"`
}

func (s DescribeAntChainQRCodeAuthorizationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainQRCodeAuthorizationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainQRCodeAuthorizationResponseBodyResult) SetAntChainId(v string) *DescribeAntChainQRCodeAuthorizationResponseBodyResult {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationResponseBodyResult) SetAuthorizationType(v string) *DescribeAntChainQRCodeAuthorizationResponseBodyResult {
	s.AuthorizationType = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationResponseBodyResult) SetQRCodeType(v string) *DescribeAntChainQRCodeAuthorizationResponseBodyResult {
	s.QRCodeType = &v
	return s
}

type DescribeAntChainQRCodeAuthorizationResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainQRCodeAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainQRCodeAuthorizationResponse) SetStatusCode(v int32) *DescribeAntChainQRCodeAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationResponse) SetBody(v *DescribeAntChainQRCodeAuthorizationResponseBody) *DescribeAntChainQRCodeAuthorizationResponse {
	s.Body = v
	return s
}

type DescribeAntChainQRCodeAuthorizationV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MINI_APP_BROWSER_TRANSACTION
	QRCodeType *string `json:"QRCodeType,omitempty" xml:"QRCodeType,omitempty"`
}

func (s DescribeAntChainQRCodeAuthorizationV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainQRCodeAuthorizationV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainQRCodeAuthorizationV2Request) SetAntChainId(v string) *DescribeAntChainQRCodeAuthorizationV2Request {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationV2Request) SetQRCodeType(v string) *DescribeAntChainQRCodeAuthorizationV2Request {
	s.QRCodeType = &v
	return s
}

type DescribeAntChainQRCodeAuthorizationV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 980061FC-058D-4298-8598-D9DDB10D0ED4
	RequestId *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainQRCodeAuthorizationV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainQRCodeAuthorizationV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainQRCodeAuthorizationV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainQRCodeAuthorizationV2ResponseBody) SetCode(v string) *DescribeAntChainQRCodeAuthorizationV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainQRCodeAuthorizationV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationV2ResponseBody) SetMessage(v string) *DescribeAntChainQRCodeAuthorizationV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationV2ResponseBody) SetRequestId(v string) *DescribeAntChainQRCodeAuthorizationV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationV2ResponseBody) SetResult(v *DescribeAntChainQRCodeAuthorizationV2ResponseBodyResult) *DescribeAntChainQRCodeAuthorizationV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationV2ResponseBody) SetResultCode(v string) *DescribeAntChainQRCodeAuthorizationV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationV2ResponseBody) SetResultMessage(v string) *DescribeAntChainQRCodeAuthorizationV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationV2ResponseBody) SetSuccess(v bool) *DescribeAntChainQRCodeAuthorizationV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainQRCodeAuthorizationV2ResponseBodyResult struct {
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// ALL_USER_AUTHORIZATION
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// example:
	//
	// MINI_APP_BROWSER_TRANSACTION
	QRCodeType *string `json:"QRCodeType,omitempty" xml:"QRCodeType,omitempty"`
}

func (s DescribeAntChainQRCodeAuthorizationV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainQRCodeAuthorizationV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainQRCodeAuthorizationV2ResponseBodyResult) SetAntChainId(v string) *DescribeAntChainQRCodeAuthorizationV2ResponseBodyResult {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationV2ResponseBodyResult) SetAuthorizationType(v string) *DescribeAntChainQRCodeAuthorizationV2ResponseBodyResult {
	s.AuthorizationType = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationV2ResponseBodyResult) SetQRCodeType(v string) *DescribeAntChainQRCodeAuthorizationV2ResponseBodyResult {
	s.QRCodeType = &v
	return s
}

type DescribeAntChainQRCodeAuthorizationV2Response struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainQRCodeAuthorizationV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainQRCodeAuthorizationV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainQRCodeAuthorizationV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainQRCodeAuthorizationV2Response) SetHeaders(v map[string]*string) *DescribeAntChainQRCodeAuthorizationV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationV2Response) SetStatusCode(v int32) *DescribeAntChainQRCodeAuthorizationV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainQRCodeAuthorizationV2Response) SetBody(v *DescribeAntChainQRCodeAuthorizationV2ResponseBody) *DescribeAntChainQRCodeAuthorizationV2Response {
	s.Body = v
	return s
}

type DescribeAntChainTransactionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pYogqb9v
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b3b0d2db83d3e670449d1e2a39d1d13b7e0e6080b0f9c6945f79eca68d1dd2ca
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
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
	// example:
	//
	// 7DCE6520-E5DB-4859-90AC-BDDB8A5501C6
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
	// example:
	//
	// 1168bc5dd0b78d15446b15ea5a7f7822a7d07c007dd3a50becf98da220fc08f6
	BlockHash *string `json:"BlockHash,omitempty" xml:"BlockHash,omitempty"`
	// example:
	//
	// 368169
	BlockHeight *int64 `json:"BlockHeight,omitempty" xml:"BlockHeight,omitempty"`
	// example:
	//
	// 10
	BlockVersion *string `json:"BlockVersion,omitempty" xml:"BlockVersion,omitempty"`
	// example:
	//
	// 1563954336850
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// b3b0d2db83d3e670449d1e2a39d1d13b7e0e6080b0f9c6945f79eca68d1dd2ca
	Hash        *string                                                   `json:"Hash,omitempty" xml:"Hash,omitempty"`
	Transaction *DescribeAntChainTransactionResponseBodyResultTransaction `json:"Transaction,omitempty" xml:"Transaction,omitempty" type:"Struct"`
}

func (s DescribeAntChainTransactionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionResponseBodyResult) SetBlockHash(v string) *DescribeAntChainTransactionResponseBodyResult {
	s.BlockHash = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResult) SetBlockHeight(v int64) *DescribeAntChainTransactionResponseBodyResult {
	s.BlockHeight = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResult) SetBlockVersion(v string) *DescribeAntChainTransactionResponseBodyResult {
	s.BlockVersion = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResult) SetCreateTime(v int64) *DescribeAntChainTransactionResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResult) SetHash(v string) *DescribeAntChainTransactionResponseBodyResult {
	s.Hash = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResult) SetTransaction(v *DescribeAntChainTransactionResponseBodyResultTransaction) *DescribeAntChainTransactionResponseBodyResult {
	s.Transaction = v
	return s
}

type DescribeAntChainTransactionResponseBodyResultTransaction struct {
	// example:
	//
	// “”
	Data       *string   `json:"Data,omitempty" xml:"Data,omitempty"`
	Extentions []*string `json:"Extentions,omitempty" xml:"Extentions,omitempty" type:"Repeated"`
	// example:
	//
	// e7d3e769f3f593dadcb8634cc5b09fc90dd3a61c4a06a79cb0923662fe6fae6b
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 4000000
	Gas *string `json:"Gas,omitempty" xml:"Gas,omitempty"`
	// example:
	//
	// b3b0d2db83d3e670449d1e2a39d1d13b7e0e6080b0f9c6945f79eca68d1dd2ca
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	// example:
	//
	// 5675407026657953619
	Nonce *string `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
	// Period
	//
	// example:
	//
	// 1563954336850
	Period     *int64    `json:"Period,omitempty" xml:"Period,omitempty"`
	Signatures []*string `json:"Signatures,omitempty" xml:"Signatures,omitempty" type:"Repeated"`
	// example:
	//
	// 1563954336850
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 961085f3c7ef705ad587d0cbe11d7863a5a11af7451f4e9b1edadd74402addf5
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// example:
	//
	// UNFREEZE_ACCOUNT_CONTRACT
	TxType *string `json:"TxType,omitempty" xml:"TxType,omitempty"`
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAntChainTransactionResponseBodyResultTransaction) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionResponseBodyResultTransaction) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetData(v string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Data = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetExtentions(v []*string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Extentions = v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetFrom(v string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.From = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetGas(v string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Gas = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetHash(v string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Hash = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetNonce(v string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Nonce = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetPeriod(v int64) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Period = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetSignatures(v []*string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Signatures = v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetTimestamp(v int64) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Timestamp = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetTo(v string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.To = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetTxType(v string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.TxType = &v
	return s
}

func (s *DescribeAntChainTransactionResponseBodyResultTransaction) SetValue(v string) *DescribeAntChainTransactionResponseBodyResultTransaction {
	s.Value = &v
	return s
}

type DescribeAntChainTransactionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainTransactionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainTransactionResponse) SetStatusCode(v int32) *DescribeAntChainTransactionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainTransactionResponse) SetBody(v *DescribeAntChainTransactionResponseBody) *DescribeAntChainTransactionResponse {
	s.Body = v
	return s
}

type DescribeAntChainTransactionReceiptRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pYogqb9v
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b3b0d2db83d3e670449d1e2a39d1d13b7e0e6080b0f9c6945f79eca68d1dd2ca
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
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
	// example:
	//
	// 9EC8FDC1-ED32-4EC0-AEC2-AFAA906ADBAB
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
	// example:
	//
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 20000
	GasUsed *string   `json:"GasUsed,omitempty" xml:"GasUsed,omitempty"`
	Logs    []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
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

func (s *DescribeAntChainTransactionReceiptResponseBodyResult) SetGasUsed(v string) *DescribeAntChainTransactionReceiptResponseBodyResult {
	s.GasUsed = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptResponseBodyResult) SetLogs(v []*string) *DescribeAntChainTransactionReceiptResponseBodyResult {
	s.Logs = v
	return s
}

func (s *DescribeAntChainTransactionReceiptResponseBodyResult) SetResult(v int64) *DescribeAntChainTransactionReceiptResponseBodyResult {
	s.Result = &v
	return s
}

type DescribeAntChainTransactionReceiptResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainTransactionReceiptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainTransactionReceiptResponse) SetStatusCode(v int32) *DescribeAntChainTransactionReceiptResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptResponse) SetBody(v *DescribeAntChainTransactionReceiptResponseBody) *DescribeAntChainTransactionReceiptResponse {
	s.Body = v
	return s
}

type DescribeAntChainTransactionReceiptV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// M8GaMEyX
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
}

func (s DescribeAntChainTransactionReceiptV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionReceiptV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionReceiptV2Request) SetAntChainId(v string) *DescribeAntChainTransactionReceiptV2Request {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptV2Request) SetConsortiumId(v string) *DescribeAntChainTransactionReceiptV2Request {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptV2Request) SetHash(v string) *DescribeAntChainTransactionReceiptV2Request {
	s.Hash = &v
	return s
}

type DescribeAntChainTransactionReceiptV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainTransactionReceiptV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainTransactionReceiptV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionReceiptV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionReceiptV2ResponseBody) SetCode(v string) *DescribeAntChainTransactionReceiptV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainTransactionReceiptV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptV2ResponseBody) SetMessage(v string) *DescribeAntChainTransactionReceiptV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptV2ResponseBody) SetRequestId(v string) *DescribeAntChainTransactionReceiptV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptV2ResponseBody) SetResult(v *DescribeAntChainTransactionReceiptV2ResponseBodyResult) *DescribeAntChainTransactionReceiptV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainTransactionReceiptV2ResponseBody) SetResultCode(v string) *DescribeAntChainTransactionReceiptV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptV2ResponseBody) SetResultMessage(v string) *DescribeAntChainTransactionReceiptV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptV2ResponseBody) SetSuccess(v bool) *DescribeAntChainTransactionReceiptV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainTransactionReceiptV2ResponseBodyResult struct {
	// example:
	//
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 20000
	GasUsed *string   `json:"GasUsed,omitempty" xml:"GasUsed,omitempty"`
	Logs    []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeAntChainTransactionReceiptV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionReceiptV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionReceiptV2ResponseBodyResult) SetData(v string) *DescribeAntChainTransactionReceiptV2ResponseBodyResult {
	s.Data = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptV2ResponseBodyResult) SetGasUsed(v string) *DescribeAntChainTransactionReceiptV2ResponseBodyResult {
	s.GasUsed = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptV2ResponseBodyResult) SetLogs(v []*string) *DescribeAntChainTransactionReceiptV2ResponseBodyResult {
	s.Logs = v
	return s
}

func (s *DescribeAntChainTransactionReceiptV2ResponseBodyResult) SetResult(v int64) *DescribeAntChainTransactionReceiptV2ResponseBodyResult {
	s.Result = &v
	return s
}

type DescribeAntChainTransactionReceiptV2Response struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainTransactionReceiptV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainTransactionReceiptV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionReceiptV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionReceiptV2Response) SetHeaders(v map[string]*string) *DescribeAntChainTransactionReceiptV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainTransactionReceiptV2Response) SetStatusCode(v int32) *DescribeAntChainTransactionReceiptV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainTransactionReceiptV2Response) SetBody(v *DescribeAntChainTransactionReceiptV2ResponseBody) *DescribeAntChainTransactionReceiptV2Response {
	s.Body = v
	return s
}

type DescribeAntChainTransactionStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pYogqb9v
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1563897600000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1563206400000
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
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

func (s *DescribeAntChainTransactionStatisticsRequest) SetEnd(v int64) *DescribeAntChainTransactionStatisticsRequest {
	s.End = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsRequest) SetStart(v int64) *DescribeAntChainTransactionStatisticsRequest {
	s.Start = &v
	return s
}

type DescribeAntChainTransactionStatisticsResponseBody struct {
	// example:
	//
	// E93884AC-6C21-4FEA-8E3A-7377D33B194F
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
	// example:
	//
	// pYogqb9v
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// 1563294007000
	CreatTime *int64 `json:"CreatTime,omitempty" xml:"CreatTime,omitempty"`
	// example:
	//
	// 2019-07-16
	Dt *string `json:"Dt,omitempty" xml:"Dt,omitempty"`
	// example:
	//
	// 148018
	LastSumBlockHeight *int64 `json:"LastSumBlockHeight,omitempty" xml:"LastSumBlockHeight,omitempty"`
	// example:
	//
	// 0
	TransCount *int64 `json:"TransCount,omitempty" xml:"TransCount,omitempty"`
}

func (s DescribeAntChainTransactionStatisticsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionStatisticsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionStatisticsResponseBodyResult) SetAntChainId(v string) *DescribeAntChainTransactionStatisticsResponseBodyResult {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsResponseBodyResult) SetCreatTime(v int64) *DescribeAntChainTransactionStatisticsResponseBodyResult {
	s.CreatTime = &v
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

func (s *DescribeAntChainTransactionStatisticsResponseBodyResult) SetTransCount(v int64) *DescribeAntChainTransactionStatisticsResponseBodyResult {
	s.TransCount = &v
	return s
}

type DescribeAntChainTransactionStatisticsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainTransactionStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainTransactionStatisticsResponse) SetStatusCode(v int32) *DescribeAntChainTransactionStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsResponse) SetBody(v *DescribeAntChainTransactionStatisticsResponseBody) *DescribeAntChainTransactionStatisticsResponse {
	s.Body = v
	return s
}

type DescribeAntChainTransactionStatisticsV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// M8GaMEyX
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1609776000000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// 1609084800000
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeAntChainTransactionStatisticsV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionStatisticsV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionStatisticsV2Request) SetAntChainId(v string) *DescribeAntChainTransactionStatisticsV2Request {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsV2Request) SetConsortiumId(v string) *DescribeAntChainTransactionStatisticsV2Request {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsV2Request) SetEnd(v int64) *DescribeAntChainTransactionStatisticsV2Request {
	s.End = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsV2Request) SetStart(v int64) *DescribeAntChainTransactionStatisticsV2Request {
	s.Start = &v
	return s
}

type DescribeAntChainTransactionStatisticsV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeAntChainTransactionStatisticsV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainTransactionStatisticsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionStatisticsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionStatisticsV2ResponseBody) SetCode(v string) *DescribeAntChainTransactionStatisticsV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainTransactionStatisticsV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsV2ResponseBody) SetMessage(v string) *DescribeAntChainTransactionStatisticsV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsV2ResponseBody) SetRequestId(v string) *DescribeAntChainTransactionStatisticsV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsV2ResponseBody) SetResult(v []*DescribeAntChainTransactionStatisticsV2ResponseBodyResult) *DescribeAntChainTransactionStatisticsV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainTransactionStatisticsV2ResponseBody) SetResultCode(v string) *DescribeAntChainTransactionStatisticsV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsV2ResponseBody) SetResultMessage(v string) *DescribeAntChainTransactionStatisticsV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsV2ResponseBody) SetSuccess(v bool) *DescribeAntChainTransactionStatisticsV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainTransactionStatisticsV2ResponseBodyResult struct {
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// 1609258222000
	CreatTime *int64 `json:"CreatTime,omitempty" xml:"CreatTime,omitempty"`
	// example:
	//
	// 1609171200000
	Dt *int64 `json:"Dt,omitempty" xml:"Dt,omitempty"`
	// example:
	//
	// 11511
	LastSumBlockHeight *int64 `json:"LastSumBlockHeight,omitempty" xml:"LastSumBlockHeight,omitempty"`
	// example:
	//
	// 4
	TransCount *int64 `json:"TransCount,omitempty" xml:"TransCount,omitempty"`
}

func (s DescribeAntChainTransactionStatisticsV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionStatisticsV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionStatisticsV2ResponseBodyResult) SetAntChainId(v string) *DescribeAntChainTransactionStatisticsV2ResponseBodyResult {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsV2ResponseBodyResult) SetCreatTime(v int64) *DescribeAntChainTransactionStatisticsV2ResponseBodyResult {
	s.CreatTime = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsV2ResponseBodyResult) SetDt(v int64) *DescribeAntChainTransactionStatisticsV2ResponseBodyResult {
	s.Dt = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsV2ResponseBodyResult) SetLastSumBlockHeight(v int64) *DescribeAntChainTransactionStatisticsV2ResponseBodyResult {
	s.LastSumBlockHeight = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsV2ResponseBodyResult) SetTransCount(v int64) *DescribeAntChainTransactionStatisticsV2ResponseBodyResult {
	s.TransCount = &v
	return s
}

type DescribeAntChainTransactionStatisticsV2Response struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainTransactionStatisticsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainTransactionStatisticsV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionStatisticsV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionStatisticsV2Response) SetHeaders(v map[string]*string) *DescribeAntChainTransactionStatisticsV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainTransactionStatisticsV2Response) SetStatusCode(v int32) *DescribeAntChainTransactionStatisticsV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainTransactionStatisticsV2Response) SetBody(v *DescribeAntChainTransactionStatisticsV2ResponseBody) *DescribeAntChainTransactionStatisticsV2Response {
	s.Body = v
	return s
}

type DescribeAntChainTransactionV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// example:
	//
	// M8GaMEyX
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 227072dcd1a5d78098da99ccc8024304b9fb54fc6b65e37ef919d7f2da080061
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
}

func (s DescribeAntChainTransactionV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionV2Request) SetAntChainId(v string) *DescribeAntChainTransactionV2Request {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainTransactionV2Request) SetConsortiumId(v string) *DescribeAntChainTransactionV2Request {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainTransactionV2Request) SetHash(v string) *DescribeAntChainTransactionV2Request {
	s.Hash = &v
	return s
}

type DescribeAntChainTransactionV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainTransactionV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainTransactionV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionV2ResponseBody) SetCode(v string) *DescribeAntChainTransactionV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainTransactionV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBody) SetMessage(v string) *DescribeAntChainTransactionV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBody) SetRequestId(v string) *DescribeAntChainTransactionV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBody) SetResult(v *DescribeAntChainTransactionV2ResponseBodyResult) *DescribeAntChainTransactionV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBody) SetResultCode(v string) *DescribeAntChainTransactionV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBody) SetResultMessage(v string) *DescribeAntChainTransactionV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBody) SetSuccess(v bool) *DescribeAntChainTransactionV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainTransactionV2ResponseBodyResult struct {
	// example:
	//
	// 1168bc5dd0b78d15446b15ea5a7f7822a7d07c007dd3a50becf98da220fc08f6
	BlockHash *string `json:"BlockHash,omitempty" xml:"BlockHash,omitempty"`
	// example:
	//
	// 100
	BlockHeight *int64 `json:"BlockHeight,omitempty" xml:"BlockHeight,omitempty"`
	// example:
	//
	// 10
	BlockVersion *string `json:"BlockVersion,omitempty" xml:"BlockVersion,omitempty"`
	// example:
	//
	// 1563954336850
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// b3b0d2db83d3e670449d1e2a39d1d13b7e0e6080b0f9c6945f79eca68d1dd2ca
	Hash        *string                                                     `json:"Hash,omitempty" xml:"Hash,omitempty"`
	Transaction *DescribeAntChainTransactionV2ResponseBodyResultTransaction `json:"Transaction,omitempty" xml:"Transaction,omitempty" type:"Struct"`
}

func (s DescribeAntChainTransactionV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionV2ResponseBodyResult) SetBlockHash(v string) *DescribeAntChainTransactionV2ResponseBodyResult {
	s.BlockHash = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBodyResult) SetBlockHeight(v int64) *DescribeAntChainTransactionV2ResponseBodyResult {
	s.BlockHeight = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBodyResult) SetBlockVersion(v string) *DescribeAntChainTransactionV2ResponseBodyResult {
	s.BlockVersion = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBodyResult) SetCreateTime(v int64) *DescribeAntChainTransactionV2ResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBodyResult) SetHash(v string) *DescribeAntChainTransactionV2ResponseBodyResult {
	s.Hash = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBodyResult) SetTransaction(v *DescribeAntChainTransactionV2ResponseBodyResultTransaction) *DescribeAntChainTransactionV2ResponseBodyResult {
	s.Transaction = v
	return s
}

type DescribeAntChainTransactionV2ResponseBodyResultTransaction struct {
	// example:
	//
	// ""
	Data       *string   `json:"Data,omitempty" xml:"Data,omitempty"`
	Extentions []*string `json:"Extentions,omitempty" xml:"Extentions,omitempty" type:"Repeated"`
	// example:
	//
	// e7d3e769f3f593dadcb8634cc5b09fc90dd3a61c4a06a79cb0923662fe6fae6b
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 4000000
	Gas *string `json:"Gas,omitempty" xml:"Gas,omitempty"`
	// example:
	//
	// b3b0d2db83d3e670449d1e2a39d1d13b7e0e6080b0f9c6945f79eca68d1dd2ca
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	// example:
	//
	// 5675407026657953619
	Nonce *string `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
	// example:
	//
	// 1563954336850
	Period     *int64    `json:"Period,omitempty" xml:"Period,omitempty"`
	Signatures []*string `json:"Signatures,omitempty" xml:"Signatures,omitempty" type:"Repeated"`
	// example:
	//
	// 1563954336850
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 961085f3c7ef705ad587d0cbe11d7863a5a11af7451f4e9b1edadd74402addf5
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// example:
	//
	// UNFREEZE_ACCOUNT_CONTRACT
	TxType *string `json:"TxType,omitempty" xml:"TxType,omitempty"`
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAntChainTransactionV2ResponseBodyResultTransaction) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionV2ResponseBodyResultTransaction) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionV2ResponseBodyResultTransaction) SetData(v string) *DescribeAntChainTransactionV2ResponseBodyResultTransaction {
	s.Data = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBodyResultTransaction) SetExtentions(v []*string) *DescribeAntChainTransactionV2ResponseBodyResultTransaction {
	s.Extentions = v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBodyResultTransaction) SetFrom(v string) *DescribeAntChainTransactionV2ResponseBodyResultTransaction {
	s.From = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBodyResultTransaction) SetGas(v string) *DescribeAntChainTransactionV2ResponseBodyResultTransaction {
	s.Gas = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBodyResultTransaction) SetHash(v string) *DescribeAntChainTransactionV2ResponseBodyResultTransaction {
	s.Hash = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBodyResultTransaction) SetNonce(v string) *DescribeAntChainTransactionV2ResponseBodyResultTransaction {
	s.Nonce = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBodyResultTransaction) SetPeriod(v int64) *DescribeAntChainTransactionV2ResponseBodyResultTransaction {
	s.Period = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBodyResultTransaction) SetSignatures(v []*string) *DescribeAntChainTransactionV2ResponseBodyResultTransaction {
	s.Signatures = v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBodyResultTransaction) SetTimestamp(v int64) *DescribeAntChainTransactionV2ResponseBodyResultTransaction {
	s.Timestamp = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBodyResultTransaction) SetTo(v string) *DescribeAntChainTransactionV2ResponseBodyResultTransaction {
	s.To = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBodyResultTransaction) SetTxType(v string) *DescribeAntChainTransactionV2ResponseBodyResultTransaction {
	s.TxType = &v
	return s
}

func (s *DescribeAntChainTransactionV2ResponseBodyResultTransaction) SetValue(v string) *DescribeAntChainTransactionV2ResponseBodyResultTransaction {
	s.Value = &v
	return s
}

type DescribeAntChainTransactionV2Response struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainTransactionV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainTransactionV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainTransactionV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainTransactionV2Response) SetHeaders(v map[string]*string) *DescribeAntChainTransactionV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainTransactionV2Response) SetStatusCode(v int32) *DescribeAntChainTransactionV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainTransactionV2Response) SetBody(v *DescribeAntChainTransactionV2ResponseBody) *DescribeAntChainTransactionV2Response {
	s.Body = v
	return s
}

type DescribeAntChainsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// AroN3X2l
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAntChainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntChainsRequest) SetConsortiumId(v string) *DescribeAntChainsRequest {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainsRequest) SetPageNumber(v int32) *DescribeAntChainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainsRequest) SetPageSize(v int32) *DescribeAntChainsRequest {
	s.PageSize = &v
	return s
}

type DescribeAntChainsResponseBody struct {
	// example:
	//
	// 52EC3196-D0B8-4FBE-A8B6-6DDABABE11C9
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
	AntChains []*DescribeAntChainsResponseBodyResultAntChains `json:"AntChains,omitempty" xml:"AntChains,omitempty" type:"Repeated"`
	// example:
	//
	// false
	IsExist    *bool                                          `json:"IsExist,omitempty" xml:"IsExist,omitempty"`
	Pagination *DescribeAntChainsResponseBodyResultPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s DescribeAntChainsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainsResponseBodyResult) SetAntChains(v []*DescribeAntChainsResponseBodyResultAntChains) *DescribeAntChainsResponseBodyResult {
	s.AntChains = v
	return s
}

func (s *DescribeAntChainsResponseBodyResult) SetIsExist(v bool) *DescribeAntChainsResponseBodyResult {
	s.IsExist = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResult) SetPagination(v *DescribeAntChainsResponseBodyResultPagination) *DescribeAntChainsResponseBodyResult {
	s.Pagination = v
	return s
}

type DescribeAntChainsResponseBodyResultAntChains struct {
	// example:
	//
	// pYogqb9v
	AntChainId   *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	AntChainName *string `json:"AntChainName,omitempty" xml:"AntChainName,omitempty"`
	// example:
	//
	// Contract
	ChainType *string `json:"ChainType,omitempty" xml:"ChainType,omitempty"`
	// example:
	//
	// classic
	CipherSuit *string `json:"CipherSuit,omitempty" xml:"CipherSuit,omitempty"`
	// example:
	//
	// 1562847396000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1626019200000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// true
	IsAdmin *bool `json:"IsAdmin,omitempty" xml:"IsAdmin,omitempty"`
	// example:
	//
	// ChainApplied
	MemberStatus *string `json:"MemberStatus,omitempty" xml:"MemberStatus,omitempty"`
	// example:
	//
	// fdmt
	MerkleTreeSuit *string `json:"MerkleTreeSuit,omitempty" xml:"MerkleTreeSuit,omitempty"`
	// example:
	//
	// Running
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// 4
	NodeNum *int32 `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Basic
	ResourceSize *string `json:"ResourceSize,omitempty" xml:"ResourceSize,omitempty"`
	// example:
	//
	// rsa
	TlsAlgo *string `json:"TlsAlgo,omitempty" xml:"TlsAlgo,omitempty"`
	// example:
	//
	// 2.19.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAntChainsResponseBodyResultAntChains) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainsResponseBodyResultAntChains) GoString() string {
	return s.String()
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetAntChainId(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetAntChainName(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.AntChainName = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetChainType(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.ChainType = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetCipherSuit(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.CipherSuit = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetCreateTime(v int64) *DescribeAntChainsResponseBodyResultAntChains {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetExpireTime(v int64) *DescribeAntChainsResponseBodyResultAntChains {
	s.ExpireTime = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetIsAdmin(v bool) *DescribeAntChainsResponseBodyResultAntChains {
	s.IsAdmin = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetMemberStatus(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.MemberStatus = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetMerkleTreeSuit(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.MerkleTreeSuit = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetNetwork(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.Network = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetNodeNum(v int32) *DescribeAntChainsResponseBodyResultAntChains {
	s.NodeNum = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetRegionId(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.RegionId = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultAntChains) SetResourceSize(v string) *DescribeAntChainsResponseBodyResultAntChains {
	s.ResourceSize = &v
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

type DescribeAntChainsResponseBodyResultPagination struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainsResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainsResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainsResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainsResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainsResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainsResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainsResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAntChainsResponse) SetStatusCode(v int32) *DescribeAntChainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainsResponse) SetBody(v *DescribeAntChainsResponseBody) *DescribeAntChainsResponse {
	s.Body = v
	return s
}

type DescribeAntChainsV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// M8GaMEyX
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAntChainsV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainsV2Request) GoString() string {
	return s.String()
}

func (s *DescribeAntChainsV2Request) SetConsortiumId(v string) *DescribeAntChainsV2Request {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeAntChainsV2Request) SetPageNumber(v int32) *DescribeAntChainsV2Request {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainsV2Request) SetPageSize(v int32) *DescribeAntChainsV2Request {
	s.PageSize = &v
	return s
}

type DescribeAntChainsV2ResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DBD6C69E-22B9-419E-B072-7A715F3AA330
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAntChainsV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// OK
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAntChainsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntChainsV2ResponseBody) SetCode(v string) *DescribeAntChainsV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBody) SetHttpStatusCode(v string) *DescribeAntChainsV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBody) SetMessage(v string) *DescribeAntChainsV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBody) SetRequestId(v string) *DescribeAntChainsV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBody) SetResult(v *DescribeAntChainsV2ResponseBodyResult) *DescribeAntChainsV2ResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAntChainsV2ResponseBody) SetResultCode(v string) *DescribeAntChainsV2ResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBody) SetResultMessage(v string) *DescribeAntChainsV2ResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBody) SetSuccess(v bool) *DescribeAntChainsV2ResponseBody {
	s.Success = &v
	return s
}

type DescribeAntChainsV2ResponseBodyResult struct {
	AntChains []*DescribeAntChainsV2ResponseBodyResultAntChains `json:"AntChains,omitempty" xml:"AntChains,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IsExist    *bool                                            `json:"IsExist,omitempty" xml:"IsExist,omitempty"`
	Pagination *DescribeAntChainsV2ResponseBodyResultPagination `json:"Pagination,omitempty" xml:"Pagination,omitempty" type:"Struct"`
}

func (s DescribeAntChainsV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainsV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAntChainsV2ResponseBodyResult) SetAntChains(v []*DescribeAntChainsV2ResponseBodyResultAntChains) *DescribeAntChainsV2ResponseBodyResult {
	s.AntChains = v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResult) SetIsExist(v bool) *DescribeAntChainsV2ResponseBodyResult {
	s.IsExist = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResult) SetPagination(v *DescribeAntChainsV2ResponseBodyResultPagination) *DescribeAntChainsV2ResponseBodyResult {
	s.Pagination = v
	return s
}

type DescribeAntChainsV2ResponseBodyResultAntChains struct {
	// example:
	//
	// 8bd720bde18c4b37b0f4a1c7834db163
	AntChainId   *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	AntChainName *string `json:"AntChainName,omitempty" xml:"AntChainName,omitempty"`
	// example:
	//
	// Contract
	ChainType *string `json:"ChainType,omitempty" xml:"ChainType,omitempty"`
	// example:
	//
	// classic
	CipherSuit *string `json:"CipherSuit,omitempty" xml:"CipherSuit,omitempty"`
	// example:
	//
	// 1609221924000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1672329600000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// ALIYUN202012291405173265824745109
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// false
	IsAdmin *bool `json:"IsAdmin,omitempty" xml:"IsAdmin,omitempty"`
	// example:
	//
	// ChainApplied
	MemberStatus *string `json:"MemberStatus,omitempty" xml:"MemberStatus,omitempty"`
	// example:
	//
	// fdmt
	MerkleTreeSuit *string `json:"MerkleTreeSuit,omitempty" xml:"MerkleTreeSuit,omitempty"`
	// example:
	//
	// true
	MonitorStatus *bool `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	// example:
	//
	// Running
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// 4
	NodeNum *int32 `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Basic
	ResourceSize *string `json:"ResourceSize,omitempty" xml:"ResourceSize,omitempty"`
	// example:
	//
	// CREATE
	RestStatus *string `json:"RestStatus,omitempty" xml:"RestStatus,omitempty"`
	// example:
	//
	// rsa
	TlsAlgo *string `json:"TlsAlgo,omitempty" xml:"TlsAlgo,omitempty"`
	// example:
	//
	// 2.19.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAntChainsV2ResponseBodyResultAntChains) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainsV2ResponseBodyResultAntChains) GoString() string {
	return s.String()
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetAntChainId(v string) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.AntChainId = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetAntChainName(v string) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.AntChainName = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetChainType(v string) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.ChainType = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetCipherSuit(v string) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.CipherSuit = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetCreateTime(v int64) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.CreateTime = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetExpireTime(v int64) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.ExpireTime = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetInstanceId(v string) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.InstanceId = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetIsAdmin(v bool) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.IsAdmin = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetMemberStatus(v string) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.MemberStatus = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetMerkleTreeSuit(v string) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.MerkleTreeSuit = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetMonitorStatus(v bool) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.MonitorStatus = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetNetwork(v string) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.Network = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetNodeNum(v int32) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.NodeNum = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetRegionId(v string) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.RegionId = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetResourceSize(v string) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.ResourceSize = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetRestStatus(v string) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.RestStatus = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetTlsAlgo(v string) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.TlsAlgo = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultAntChains) SetVersion(v string) *DescribeAntChainsV2ResponseBodyResultAntChains {
	s.Version = &v
	return s
}

type DescribeAntChainsV2ResponseBodyResultPagination struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntChainsV2ResponseBodyResultPagination) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainsV2ResponseBodyResultPagination) GoString() string {
	return s.String()
}

func (s *DescribeAntChainsV2ResponseBodyResultPagination) SetPageNumber(v int32) *DescribeAntChainsV2ResponseBodyResultPagination {
	s.PageNumber = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultPagination) SetPageSize(v int32) *DescribeAntChainsV2ResponseBodyResultPagination {
	s.PageSize = &v
	return s
}

func (s *DescribeAntChainsV2ResponseBodyResultPagination) SetTotalCount(v int32) *DescribeAntChainsV2ResponseBodyResultPagination {
	s.TotalCount = &v
	return s
}

type DescribeAntChainsV2Response struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntChainsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntChainsV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntChainsV2Response) GoString() string {
	return s.String()
}

func (s *DescribeAntChainsV2Response) SetHeaders(v map[string]*string) *DescribeAntChainsV2Response {
	s.Headers = v
	return s
}

func (s *DescribeAntChainsV2Response) SetStatusCode(v int32) *DescribeAntChainsV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntChainsV2Response) SetBody(v *DescribeAntChainsV2ResponseBody) *DescribeAntChainsV2Response {
	s.Body = v
	return s
}

type DescribeEthereumDeletableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eth-1ilrhcav4ibdl
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeEthereumDeletableResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEthereumDeletableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEthereumDeletableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEthereumDeletableResponseBody) SetErrorCode(v int32) *DescribeEthereumDeletableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeEthereumDeletableResponseBody) SetRequestId(v string) *DescribeEthereumDeletableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEthereumDeletableResponseBody) SetResult(v *DescribeEthereumDeletableResponseBodyResult) *DescribeEthereumDeletableResponseBody {
	s.Result = v
	return s
}

func (s *DescribeEthereumDeletableResponseBody) SetSuccess(v bool) *DescribeEthereumDeletableResponseBody {
	s.Success = &v
	return s
}

type DescribeEthereumDeletableResponseBodyResult struct {
	// example:
	//
	// true
	Deletable *bool `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
	// example:
	//
	// eth-1ilrhcav4ibdl
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEthereumDeletableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeEthereumDeletableResponse) SetStatusCode(v int32) *DescribeEthereumDeletableResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEthereumDeletableResponse) SetBody(v *DescribeEthereumDeletableResponseBody) *DescribeEthereumDeletableResponse {
	s.Body = v
	return s
}

type DescribeFabricCandidateOrganizationsRequest struct {
	// example:
	//
	// cn-hangzhou
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricCandidateOrganizationsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricCandidateOrganizationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricCandidateOrganizationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricCandidateOrganizationsResponseBody) SetErrorCode(v int32) *DescribeFabricCandidateOrganizationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricCandidateOrganizationsResponseBody) SetRequestId(v string) *DescribeFabricCandidateOrganizationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricCandidateOrganizationsResponseBody) SetResult(v []*DescribeFabricCandidateOrganizationsResponseBodyResult) *DescribeFabricCandidateOrganizationsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricCandidateOrganizationsResponseBody) SetSuccess(v bool) *DescribeFabricCandidateOrganizationsResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricCandidateOrganizationsResponseBodyResult struct {
	// example:
	//
	// Running
	ClusterState *string `json:"ClusterState,omitempty" xml:"ClusterState,omitempty"`
	// example:
	//
	// peers-1oxw31d04****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// org1
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	// example:
	//
	// Running
	ServiceState *string `json:"ServiceState,omitempty" xml:"ServiceState,omitempty"`
}

func (s DescribeFabricCandidateOrganizationsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricCandidateOrganizationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricCandidateOrganizationsResponseBodyResult) SetClusterState(v string) *DescribeFabricCandidateOrganizationsResponseBodyResult {
	s.ClusterState = &v
	return s
}

func (s *DescribeFabricCandidateOrganizationsResponseBodyResult) SetOrganizationId(v string) *DescribeFabricCandidateOrganizationsResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricCandidateOrganizationsResponseBodyResult) SetOrganizationName(v string) *DescribeFabricCandidateOrganizationsResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *DescribeFabricCandidateOrganizationsResponseBodyResult) SetServiceState(v string) *DescribeFabricCandidateOrganizationsResponseBodyResult {
	s.ServiceState = &v
	return s
}

type DescribeFabricCandidateOrganizationsResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricCandidateOrganizationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricCandidateOrganizationsResponse) SetStatusCode(v int32) *DescribeFabricCandidateOrganizationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricCandidateOrganizationsResponse) SetBody(v *DescribeFabricCandidateOrganizationsResponseBody) *DescribeFabricCandidateOrganizationsResponse {
	s.Body = v
	return s
}

type DescribeFabricChaincodeDefinitionTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// code-sacc-jbkuldd3i7ad
	ChaincodeId *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DescribeFabricChaincodeDefinitionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChaincodeDefinitionTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricChaincodeDefinitionTaskRequest) SetChaincodeId(v string) *DescribeFabricChaincodeDefinitionTaskRequest {
	s.ChaincodeId = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskRequest) SetOrganizationId(v string) *DescribeFabricChaincodeDefinitionTaskRequest {
	s.OrganizationId = &v
	return s
}

type DescribeFabricChaincodeDefinitionTaskResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeFabricChaincodeDefinitionTaskResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricChaincodeDefinitionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChaincodeDefinitionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBody) SetErrorCode(v int32) *DescribeFabricChaincodeDefinitionTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBody) SetRequestId(v string) *DescribeFabricChaincodeDefinitionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBody) SetResult(v *DescribeFabricChaincodeDefinitionTaskResponseBodyResult) *DescribeFabricChaincodeDefinitionTaskResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBody) SetSuccess(v bool) *DescribeFabricChaincodeDefinitionTaskResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricChaincodeDefinitionTaskResponseBodyResult struct {
	Approvers []*string `json:"Approvers,omitempty" xml:"Approvers,omitempty" type:"Repeated"`
	// example:
	//
	// f100
	ChannelName *string                                                         `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	Content     *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// example:
	//
	// 1533025590
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// perf9141MSP
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// ChaincodeDefinition
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeFabricChaincodeDefinitionTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChaincodeDefinitionTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResult) SetApprovers(v []*string) *DescribeFabricChaincodeDefinitionTaskResponseBodyResult {
	s.Approvers = v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResult) SetChannelName(v string) *DescribeFabricChaincodeDefinitionTaskResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResult) SetContent(v *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContent) *DescribeFabricChaincodeDefinitionTaskResponseBodyResult {
	s.Content = v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResult) SetCreateTime(v int64) *DescribeFabricChaincodeDefinitionTaskResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResult) SetCreator(v string) *DescribeFabricChaincodeDefinitionTaskResponseBodyResult {
	s.Creator = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResult) SetDescription(v string) *DescribeFabricChaincodeDefinitionTaskResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResult) SetStatus(v string) *DescribeFabricChaincodeDefinitionTaskResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResult) SetTaskId(v string) *DescribeFabricChaincodeDefinitionTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResult) SetType(v string) *DescribeFabricChaincodeDefinitionTaskResponseBodyResult {
	s.Type = &v
	return s
}

type DescribeFabricChaincodeDefinitionTaskResponseBodyResultContent struct {
	ChaincodeDefinition *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition `json:"ChaincodeDefinition,omitempty" xml:"ChaincodeDefinition,omitempty" type:"Struct"`
}

func (s DescribeFabricChaincodeDefinitionTaskResponseBodyResultContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChaincodeDefinitionTaskResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContent) SetChaincodeDefinition(v *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition) *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContent {
	s.ChaincodeDefinition = v
	return s
}

type DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition struct {
	// example:
	//
	// sacc_v1.0:698d4e73c5f1d0239fbf0313ac79ddfaf44d5a803ecb02a4d1983cca4a706860
	ChaincodePackageId *string `json:"ChaincodePackageId,omitempty" xml:"ChaincodePackageId,omitempty"`
	// example:
	//
	// collectionconfig
	CollectionConfig *string `json:"CollectionConfig,omitempty" xml:"CollectionConfig,omitempty"`
	// example:
	//
	// OR ("perf9141MSP.member")
	EndorsementPolicy *string `json:"EndorsementPolicy,omitempty" xml:"EndorsementPolicy,omitempty"`
	// example:
	//
	// true
	InitRequired *bool `json:"InitRequired,omitempty" xml:"InitRequired,omitempty"`
	// example:
	//
	// sacc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Sequence *int64 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// example:
	//
	// 698d4e73c5f1d0239fbf0313ac79ddfaf44d5a803ecb02a4d1983cca4a706860
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// example:
	//
	// 1.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition) GoString() string {
	return s.String()
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition) SetChaincodePackageId(v string) *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition {
	s.ChaincodePackageId = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition) SetCollectionConfig(v string) *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition {
	s.CollectionConfig = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition) SetEndorsementPolicy(v string) *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition {
	s.EndorsementPolicy = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition) SetInitRequired(v bool) *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition {
	s.InitRequired = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition) SetName(v string) *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition {
	s.Name = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition) SetSequence(v int64) *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition {
	s.Sequence = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition) SetUid(v string) *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition {
	s.Uid = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition) SetVersion(v string) *DescribeFabricChaincodeDefinitionTaskResponseBodyResultContentChaincodeDefinition {
	s.Version = &v
	return s
}

type DescribeFabricChaincodeDefinitionTaskResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricChaincodeDefinitionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFabricChaincodeDefinitionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChaincodeDefinitionTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricChaincodeDefinitionTaskResponse) SetHeaders(v map[string]*string) *DescribeFabricChaincodeDefinitionTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponse) SetStatusCode(v int32) *DescribeFabricChaincodeDefinitionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricChaincodeDefinitionTaskResponse) SetBody(v *DescribeFabricChaincodeDefinitionTaskResponseBody) *DescribeFabricChaincodeDefinitionTaskResponse {
	s.Body = v
	return s
}

type DescribeFabricChaincodeUploadPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa1-1oxw31d04****
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeFabricChaincodeUploadPolicyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricChaincodeUploadPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChaincodeUploadPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBody) SetErrorCode(v int32) *DescribeFabricChaincodeUploadPolicyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBody) SetRequestId(v string) *DescribeFabricChaincodeUploadPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBody) SetResult(v *DescribeFabricChaincodeUploadPolicyResponseBodyResult) *DescribeFabricChaincodeUploadPolicyResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBody) SetSuccess(v bool) *DescribeFabricChaincodeUploadPolicyResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricChaincodeUploadPolicyResponseBodyResult struct {
	// example:
	//
	// accessId
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// test3/4bd1de21-2fff-4405-a8e9-f864c4bee24f
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// example:
	//
	// 1533025590
	Expire *int32 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// example:
	//
	// https://chaincode.oss-cn-beijing.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAxOC0wNy0zMVQwODoyNjozMC40NzdaIiwiY29uZGl
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// signature
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s DescribeFabricChaincodeUploadPolicyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChaincodeUploadPolicyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBodyResult) SetAccessId(v string) *DescribeFabricChaincodeUploadPolicyResponseBodyResult {
	s.AccessId = &v
	return s
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBodyResult) SetDir(v string) *DescribeFabricChaincodeUploadPolicyResponseBodyResult {
	s.Dir = &v
	return s
}

func (s *DescribeFabricChaincodeUploadPolicyResponseBodyResult) SetExpire(v int32) *DescribeFabricChaincodeUploadPolicyResponseBodyResult {
	s.Expire = &v
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

func (s *DescribeFabricChaincodeUploadPolicyResponseBodyResult) SetSignature(v string) *DescribeFabricChaincodeUploadPolicyResponseBodyResult {
	s.Signature = &v
	return s
}

type DescribeFabricChaincodeUploadPolicyResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricChaincodeUploadPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricChaincodeUploadPolicyResponse) SetStatusCode(v int32) *DescribeFabricChaincodeUploadPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricChaincodeUploadPolicyResponse) SetBody(v *DescribeFabricChaincodeUploadPolicyResponseBody) *DescribeFabricChaincodeUploadPolicyResponse {
	s.Body = v
	return s
}

type DescribeFabricChannelMembersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// chan-channelx-1l1hmckuu****
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricChannelMembersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricChannelMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChannelMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricChannelMembersResponseBody) SetErrorCode(v int32) *DescribeFabricChannelMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBody) SetRequestId(v string) *DescribeFabricChannelMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBody) SetResult(v []*DescribeFabricChannelMembersResponseBodyResult) *DescribeFabricChannelMembersResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricChannelMembersResponseBody) SetSuccess(v bool) *DescribeFabricChannelMembersResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricChannelMembersResponseBodyResult struct {
	// example:
	//
	// 1544411108000
	AcceptTime *string `json:"AcceptTime,omitempty" xml:"AcceptTime,omitempty"`
	// example:
	//
	// chan-channelx-1l1hmckuuisxo
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1544411108000
	InviteTime *string `json:"InviteTime,omitempty" xml:"InviteTime,omitempty"`
	// example:
	//
	// description
	OrganizationDescription *string `json:"OrganizationDescription,omitempty" xml:"OrganizationDescription,omitempty"`
	// example:
	//
	// hello
	OrganizationDomain *string `json:"OrganizationDomain,omitempty" xml:"OrganizationDomain,omitempty"`
	// example:
	//
	// peers-aaaaaa1-1oxw31d046jtl
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// name
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// true
	WithPeer *bool `json:"WithPeer,omitempty" xml:"WithPeer,omitempty"`
}

func (s DescribeFabricChannelMembersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricChannelMembersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetAcceptTime(v string) *DescribeFabricChannelMembersResponseBodyResult {
	s.AcceptTime = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetChannelId(v string) *DescribeFabricChannelMembersResponseBodyResult {
	s.ChannelId = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetInviteTime(v string) *DescribeFabricChannelMembersResponseBodyResult {
	s.InviteTime = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetOrganizationDescription(v string) *DescribeFabricChannelMembersResponseBodyResult {
	s.OrganizationDescription = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetOrganizationDomain(v string) *DescribeFabricChannelMembersResponseBodyResult {
	s.OrganizationDomain = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetOrganizationId(v string) *DescribeFabricChannelMembersResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetOrganizationName(v string) *DescribeFabricChannelMembersResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetState(v string) *DescribeFabricChannelMembersResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricChannelMembersResponseBodyResult) SetWithPeer(v bool) *DescribeFabricChannelMembersResponseBodyResult {
	s.WithPeer = &v
	return s
}

type DescribeFabricChannelMembersResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricChannelMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricChannelMembersResponse) SetStatusCode(v int32) *DescribeFabricChannelMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricChannelMembersResponse) SetBody(v *DescribeFabricChannelMembersResponseBody) *DescribeFabricChannelMembersResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumAdminStatusRequest struct {
	// example:
	//
	// cn-hangzhou
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricConsortiumAdminStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricConsortiumAdminStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumAdminStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumAdminStatusResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumAdminStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumAdminStatusResponseBody) SetRequestId(v string) *DescribeFabricConsortiumAdminStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumAdminStatusResponseBody) SetResult(v []*DescribeFabricConsortiumAdminStatusResponseBodyResult) *DescribeFabricConsortiumAdminStatusResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricConsortiumAdminStatusResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumAdminStatusResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricConsortiumAdminStatusResponseBodyResult struct {
	// example:
	//
	// true
	ConsortiumAdministrator *bool `json:"ConsortiumAdministrator,omitempty" xml:"ConsortiumAdministrator,omitempty"`
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwlulg7
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
}

func (s DescribeFabricConsortiumAdminStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumAdminStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumAdminStatusResponseBodyResult) SetConsortiumAdministrator(v bool) *DescribeFabricConsortiumAdminStatusResponseBodyResult {
	s.ConsortiumAdministrator = &v
	return s
}

func (s *DescribeFabricConsortiumAdminStatusResponseBodyResult) SetConsortiumId(v string) *DescribeFabricConsortiumAdminStatusResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

type DescribeFabricConsortiumAdminStatusResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricConsortiumAdminStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricConsortiumAdminStatusResponse) SetStatusCode(v int32) *DescribeFabricConsortiumAdminStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricConsortiumAdminStatusResponse) SetBody(v *DescribeFabricConsortiumAdminStatusResponseBody) *DescribeFabricConsortiumAdminStatusResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumChaincodesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwlulg7
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricConsortiumChaincodesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricConsortiumChaincodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumChaincodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumChaincodesResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumChaincodesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBody) SetRequestId(v string) *DescribeFabricConsortiumChaincodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBody) SetResult(v []*DescribeFabricConsortiumChaincodesResponseBodyResult) *DescribeFabricConsortiumChaincodesResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumChaincodesResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricConsortiumChaincodesResponseBodyResult struct {
	// example:
	//
	// code-sacc-1pr09q7jmo0np
	ChaincodeId *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	// example:
	//
	// sacc
	ChaincodeName *string `json:"ChaincodeName,omitempty" xml:"ChaincodeName,omitempty"`
	// example:
	//
	// 1.0
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" xml:"ChaincodeVersion,omitempty"`
	// example:
	//
	// chan-channelx-1l1hmckuuisxo
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// cname
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwlulg7
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1544411108000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1544411108000
	DeployTime *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	// example:
	//
	// OR (\\"aaaaaa1MSP.peer\\")
	EndorsePolicy *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	// example:
	//
	// input
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// true
	Install *bool `json:"Install,omitempty" xml:"Install,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// go-sdk-demo/chaincode/src
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// example:
	//
	// name
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// example:
	//
	// Instantiatable
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeFabricConsortiumChaincodesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumChaincodesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetChaincodeId(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.ChaincodeId = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetChaincodeName(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.ChaincodeName = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetChaincodeVersion(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.ChaincodeVersion = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetChannelId(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.ChannelId = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetChannelName(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetConsortiumId(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetCreateTime(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetDeployTime(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.DeployTime = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetEndorsePolicy(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.EndorsePolicy = &v
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

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetMessage(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.Message = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetPath(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.Path = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetProviderId(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.ProviderId = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetProviderName(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.ProviderName = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetState(v string) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponseBodyResult) SetType(v int32) *DescribeFabricConsortiumChaincodesResponseBodyResult {
	s.Type = &v
	return s
}

type DescribeFabricConsortiumChaincodesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricConsortiumChaincodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricConsortiumChaincodesResponse) SetStatusCode(v int32) *DescribeFabricConsortiumChaincodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricConsortiumChaincodesResponse) SetBody(v *DescribeFabricConsortiumChaincodesResponseBody) *DescribeFabricConsortiumChaincodesResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumChannelsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwlulg7
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricConsortiumChannelsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricConsortiumChannelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumChannelsResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumChannelsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBody) SetRequestId(v string) *DescribeFabricConsortiumChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBody) SetResult(v []*DescribeFabricConsortiumChannelsResponseBodyResult) *DescribeFabricConsortiumChannelsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumChannelsResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricConsortiumChannelsResponseBodyResult struct {
	// example:
	//
	// 2
	BatchTimeout *int32 `json:"BatchTimeout,omitempty" xml:"BatchTimeout,omitempty"`
	// example:
	//
	// 2
	BlockCount *int32 `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
	// example:
	//
	// 2
	ChaincodeCount *int32 `json:"ChaincodeCount,omitempty" xml:"ChaincodeCount,omitempty"`
	// example:
	//
	// chan-channelx-1l1hmckuuisxo
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// mychammel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// 1
	ConsortiumChannelId *int32 `json:"ConsortiumChannelId,omitempty" xml:"ConsortiumChannelId,omitempty"`
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwlulg7
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// name
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	// example:
	//
	// 1544411108000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1544411108000
	DeleteTime *string `json:"DeleteTime,omitempty" xml:"DeleteTime,omitempty"`
	// example:
	//
	// true
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// example:
	//
	// 200
	MaxMessageCount *int32 `json:"MaxMessageCount,omitempty" xml:"MaxMessageCount,omitempty"`
	// example:
	//
	// 3
	MemberCount *int32 `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	// example:
	//
	// 3
	MemberJoinedCount *string `json:"MemberJoinedCount,omitempty" xml:"MemberJoinedCount,omitempty"`
	// example:
	//
	// true
	NeedJoined *bool `json:"NeedJoined,omitempty" xml:"NeedJoined,omitempty"`
	// example:
	//
	// 26842
	OwnerBid *string `json:"OwnerBid,omitempty" xml:"OwnerBid,omitempty"`
	// example:
	//
	// uid-*
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 212
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// example:
	//
	// 102410241
	PreferredMaxBytes *int32 `json:"PreferredMaxBytes,omitempty" xml:"PreferredMaxBytes,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// true
	SupportChannelConfig *bool `json:"SupportChannelConfig,omitempty" xml:"SupportChannelConfig,omitempty"`
	// example:
	//
	// 1544411108000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetBlockCount(v int32) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.BlockCount = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetChaincodeCount(v int32) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.ChaincodeCount = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetChannelId(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.ChannelId = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetChannelName(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetConsortiumChannelId(v int32) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.ConsortiumChannelId = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetConsortiumId(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetConsortiumName(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.ConsortiumName = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetCreateTime(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetDeleteTime(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.DeleteTime = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetDeleted(v bool) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.Deleted = &v
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

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetMemberJoinedCount(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.MemberJoinedCount = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetNeedJoined(v bool) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.NeedJoined = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetOwnerBid(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.OwnerBid = &v
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

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetPreferredMaxBytes(v int32) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.PreferredMaxBytes = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetRequestId(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetState(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetSupportChannelConfig(v bool) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.SupportChannelConfig = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponseBodyResult) SetUpdateTime(v string) *DescribeFabricConsortiumChannelsResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type DescribeFabricConsortiumChannelsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricConsortiumChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricConsortiumChannelsResponse) SetStatusCode(v int32) *DescribeFabricConsortiumChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricConsortiumChannelsResponse) SetBody(v *DescribeFabricConsortiumChannelsResponseBody) *DescribeFabricConsortiumChannelsResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumConfigResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeFabricConsortiumConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricConsortiumConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumConfigResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumConfigResponseBody) SetRequestId(v string) *DescribeFabricConsortiumConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumConfigResponseBody) SetResult(v *DescribeFabricConsortiumConfigResponseBodyResult) *DescribeFabricConsortiumConfigResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricConsortiumConfigResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumConfigResponseBody {
	s.Success = &v
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricConsortiumConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricConsortiumConfigResponse) SetStatusCode(v int32) *DescribeFabricConsortiumConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricConsortiumConfigResponse) SetBody(v *DescribeFabricConsortiumConfigResponseBody) *DescribeFabricConsortiumConfigResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumDeletableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwl****
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeFabricConsortiumDeletableResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricConsortiumDeletableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumDeletableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumDeletableResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumDeletableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBody) SetRequestId(v string) *DescribeFabricConsortiumDeletableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBody) SetResult(v *DescribeFabricConsortiumDeletableResponseBodyResult) *DescribeFabricConsortiumDeletableResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumDeletableResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricConsortiumDeletableResponseBodyResult struct {
	// example:
	//
	// codename
	CodeName *string `json:"CodeName,omitempty" xml:"CodeName,omitempty"`
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwlulg7
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// myconsortium
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	// example:
	//
	// true
	Deletable *bool `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// domain
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// cn-hangzhou
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeFabricConsortiumDeletableResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumDeletableResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumDeletableResponseBodyResult) SetCodeName(v string) *DescribeFabricConsortiumDeletableResponseBodyResult {
	s.CodeName = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBodyResult) SetConsortiumId(v string) *DescribeFabricConsortiumDeletableResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBodyResult) SetConsortiumName(v string) *DescribeFabricConsortiumDeletableResponseBodyResult {
	s.ConsortiumName = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBodyResult) SetDeletable(v bool) *DescribeFabricConsortiumDeletableResponseBodyResult {
	s.Deletable = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBodyResult) SetDescription(v string) *DescribeFabricConsortiumDeletableResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBodyResult) SetDomain(v string) *DescribeFabricConsortiumDeletableResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponseBodyResult) SetRegionId(v string) *DescribeFabricConsortiumDeletableResponseBodyResult {
	s.RegionId = &v
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

type DescribeFabricConsortiumDeletableResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricConsortiumDeletableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricConsortiumDeletableResponse) SetStatusCode(v int32) *DescribeFabricConsortiumDeletableResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricConsortiumDeletableResponse) SetBody(v *DescribeFabricConsortiumDeletableResponseBody) *DescribeFabricConsortiumDeletableResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumMemberApprovalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwlulg7
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricConsortiumMemberApprovalResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricConsortiumMemberApprovalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumMemberApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumMemberApprovalResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBody) SetRequestId(v string) *DescribeFabricConsortiumMemberApprovalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBody) SetResult(v []*DescribeFabricConsortiumMemberApprovalResponseBodyResult) *DescribeFabricConsortiumMemberApprovalResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumMemberApprovalResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricConsortiumMemberApprovalResponseBodyResult struct {
	// example:
	//
	// Any
	ChannelCreatePolicy *string `json:"ChannelCreatePolicy,omitempty" xml:"ChannelCreatePolicy,omitempty"`
	// example:
	//
	// 1544411108000
	ConfirmTime *string `json:"ConfirmTime,omitempty" xml:"ConfirmTime,omitempty"`
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwlulg7
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// name
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	// example:
	//
	// domain
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// orgname
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeFabricConsortiumMemberApprovalResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumMemberApprovalResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBodyResult) SetChannelCreatePolicy(v string) *DescribeFabricConsortiumMemberApprovalResponseBodyResult {
	s.ChannelCreatePolicy = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBodyResult) SetConfirmTime(v string) *DescribeFabricConsortiumMemberApprovalResponseBodyResult {
	s.ConfirmTime = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBodyResult) SetConsortiumId(v string) *DescribeFabricConsortiumMemberApprovalResponseBodyResult {
	s.ConsortiumId = &v
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

func (s *DescribeFabricConsortiumMemberApprovalResponseBodyResult) SetOrganizationName(v string) *DescribeFabricConsortiumMemberApprovalResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponseBodyResult) SetState(v string) *DescribeFabricConsortiumMemberApprovalResponseBodyResult {
	s.State = &v
	return s
}

type DescribeFabricConsortiumMemberApprovalResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricConsortiumMemberApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricConsortiumMemberApprovalResponse) SetStatusCode(v int32) *DescribeFabricConsortiumMemberApprovalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricConsortiumMemberApprovalResponse) SetBody(v *DescribeFabricConsortiumMemberApprovalResponseBody) *DescribeFabricConsortiumMemberApprovalResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumMembersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwl****
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricConsortiumMembersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricConsortiumMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumMembersResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumMembersResponseBody) SetRequestId(v string) *DescribeFabricConsortiumMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumMembersResponseBody) SetResult(v []*DescribeFabricConsortiumMembersResponseBodyResult) *DescribeFabricConsortiumMembersResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricConsortiumMembersResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumMembersResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricConsortiumMembersResponseBodyResult struct {
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwlulg7
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// domain
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1544411108000
	JoinedTime *string `json:"JoinedTime,omitempty" xml:"JoinedTime,omitempty"`
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// name
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
}

func (s DescribeFabricConsortiumMembersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumMembersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumMembersResponseBodyResult) SetConsortiumId(v string) *DescribeFabricConsortiumMembersResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumMembersResponseBodyResult) SetDescription(v string) *DescribeFabricConsortiumMembersResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeFabricConsortiumMembersResponseBodyResult) SetDomain(v string) *DescribeFabricConsortiumMembersResponseBodyResult {
	s.Domain = &v
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

func (s *DescribeFabricConsortiumMembersResponseBodyResult) SetOrganizationName(v string) *DescribeFabricConsortiumMembersResponseBodyResult {
	s.OrganizationName = &v
	return s
}

type DescribeFabricConsortiumMembersResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricConsortiumMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricConsortiumMembersResponse) SetStatusCode(v int32) *DescribeFabricConsortiumMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricConsortiumMembersResponse) SetBody(v *DescribeFabricConsortiumMembersResponseBody) *DescribeFabricConsortiumMembersResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumOrderersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// consortium-aaaaaa-akpcsjjac2jd
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricConsortiumOrderersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricConsortiumOrderersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumOrderersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumOrderersResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumOrderersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponseBody) SetRequestId(v string) *DescribeFabricConsortiumOrderersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponseBody) SetResult(v []*DescribeFabricConsortiumOrderersResponseBodyResult) *DescribeFabricConsortiumOrderersResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumOrderersResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricConsortiumOrderersResponseBodyResult struct {
	// example:
	//
	// 1544411108000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// domain
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// ecs.n1.small
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// order1
	OrdererName *string `json:"OrdererName,omitempty" xml:"OrdererName,omitempty"`
	// example:
	//
	// 7050
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 1544411108000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeFabricConsortiumOrderersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumOrderersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumOrderersResponseBodyResult) SetCreateTime(v string) *DescribeFabricConsortiumOrderersResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponseBodyResult) SetDomain(v string) *DescribeFabricConsortiumOrderersResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponseBodyResult) SetInstanceType(v string) *DescribeFabricConsortiumOrderersResponseBodyResult {
	s.InstanceType = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponseBodyResult) SetOrdererName(v string) *DescribeFabricConsortiumOrderersResponseBodyResult {
	s.OrdererName = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponseBodyResult) SetPort(v int32) *DescribeFabricConsortiumOrderersResponseBodyResult {
	s.Port = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponseBodyResult) SetUpdateTime(v string) *DescribeFabricConsortiumOrderersResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type DescribeFabricConsortiumOrderersResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricConsortiumOrderersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricConsortiumOrderersResponse) SetStatusCode(v int32) *DescribeFabricConsortiumOrderersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricConsortiumOrderersResponse) SetBody(v *DescribeFabricConsortiumOrderersResponseBody) *DescribeFabricConsortiumOrderersResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumSpecsResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricConsortiumSpecsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricConsortiumSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumSpecsResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumSpecsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumSpecsResponseBody) SetRequestId(v string) *DescribeFabricConsortiumSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumSpecsResponseBody) SetResult(v []*DescribeFabricConsortiumSpecsResponseBodyResult) *DescribeFabricConsortiumSpecsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricConsortiumSpecsResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumSpecsResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricConsortiumSpecsResponseBodyResult struct {
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// basic
	SpecName  *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	SpecTitle *string `json:"SpecTitle,omitempty" xml:"SpecTitle,omitempty"`
}

func (s DescribeFabricConsortiumSpecsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumSpecsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumSpecsResponseBodyResult) SetEnable(v bool) *DescribeFabricConsortiumSpecsResponseBodyResult {
	s.Enable = &v
	return s
}

func (s *DescribeFabricConsortiumSpecsResponseBodyResult) SetSpecName(v string) *DescribeFabricConsortiumSpecsResponseBodyResult {
	s.SpecName = &v
	return s
}

func (s *DescribeFabricConsortiumSpecsResponseBodyResult) SetSpecTitle(v string) *DescribeFabricConsortiumSpecsResponseBodyResult {
	s.SpecTitle = &v
	return s
}

type DescribeFabricConsortiumSpecsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricConsortiumSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricConsortiumSpecsResponse) SetStatusCode(v int32) *DescribeFabricConsortiumSpecsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricConsortiumSpecsResponse) SetBody(v *DescribeFabricConsortiumSpecsResponseBody) *DescribeFabricConsortiumSpecsResponse {
	s.Body = v
	return s
}

type DescribeFabricConsortiumsRequest struct {
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwl****
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Location *string                                `json:"Location,omitempty" xml:"Location,omitempty"`
	Tag      []*DescribeFabricConsortiumsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricConsortiumsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricConsortiumsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricConsortiumsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricConsortiumsResponseBody) SetErrorCode(v int32) *DescribeFabricConsortiumsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBody) SetRequestId(v string) *DescribeFabricConsortiumsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBody) SetResult(v []*DescribeFabricConsortiumsResponseBodyResult) *DescribeFabricConsortiumsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricConsortiumsResponseBody) SetSuccess(v bool) *DescribeFabricConsortiumsResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricConsortiumsResponseBodyResult struct {
	// example:
	//
	// 2
	ChannelCount *int32 `json:"ChannelCount,omitempty" xml:"ChannelCount,omitempty"`
	// example:
	//
	// Any
	ChannelPolicy *string `json:"ChannelPolicy,omitempty" xml:"ChannelPolicy,omitempty"`
	// example:
	//
	// code
	CodeName *string `json:"CodeName,omitempty" xml:"CodeName,omitempty"`
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwl****
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// name
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	// example:
	//
	// 1544411108000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// domain
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1544411108000
	ExpiredTime  *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// example:
	//
	// 2
	OrganizationCount *int32 `json:"OrganizationCount,omitempty" xml:"OrganizationCount,omitempty"`
	// example:
	//
	// 26842
	OwnerBid *string `json:"OwnerBid,omitempty" xml:"OwnerBid,omitempty"`
	// example:
	//
	// uid-125566
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 12042815
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 3A929FAC-A82D-4EAB-A0FF-D867426D8B23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// basic
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	// example:
	//
	// Pending
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// true
	SupportChannelConfig *bool                                              `json:"SupportChannelConfig,omitempty" xml:"SupportChannelConfig,omitempty"`
	Tags                 []*DescribeFabricConsortiumsResponseBodyResultTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *DescribeFabricConsortiumsResponseBodyResult) SetChannelPolicy(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.ChannelPolicy = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetCodeName(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.CodeName = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetConsortiumId(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetConsortiumName(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.ConsortiumName = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetCreateTime(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetDomain(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetExpiredTime(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetMajorVersion(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.MajorVersion = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetOrganizationCount(v int32) *DescribeFabricConsortiumsResponseBodyResult {
	s.OrganizationCount = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetOwnerBid(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.OwnerBid = &v
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

func (s *DescribeFabricConsortiumsResponseBodyResult) SetRegionId(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetRequestId(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetSpecName(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.SpecName = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetState(v string) *DescribeFabricConsortiumsResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetSupportChannelConfig(v bool) *DescribeFabricConsortiumsResponseBodyResult {
	s.SupportChannelConfig = &v
	return s
}

func (s *DescribeFabricConsortiumsResponseBodyResult) SetTags(v []*DescribeFabricConsortiumsResponseBodyResultTags) *DescribeFabricConsortiumsResponseBodyResult {
	s.Tags = v
	return s
}

type DescribeFabricConsortiumsResponseBodyResultTags struct {
	// example:
	//
	// app
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// nat_internet
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricConsortiumsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricConsortiumsResponse) SetStatusCode(v int32) *DescribeFabricConsortiumsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricConsortiumsResponse) SetBody(v *DescribeFabricConsortiumsResponseBody) *DescribeFabricConsortiumsResponse {
	s.Body = v
	return s
}

type DescribeFabricExplorerRequest struct {
	// example:
	//
	// {}
	ExBody *string `json:"ExBody,omitempty" xml:"ExBody,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GET
	ExMethod *string `json:"ExMethod,omitempty" xml:"ExMethod,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /api/v1/channel
	ExUrl *string `json:"ExUrl,omitempty" xml:"ExUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-test11-tzwny25vqicq
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DescribeFabricExplorerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricExplorerRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricExplorerRequest) SetExBody(v string) *DescribeFabricExplorerRequest {
	s.ExBody = &v
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

func (s *DescribeFabricExplorerRequest) SetOrganizationId(v string) *DescribeFabricExplorerRequest {
	s.OrganizationId = &v
	return s
}

type DescribeFabricExplorerResponseBody struct {
	// example:
	//
	// ExplorerException
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// An error in the monitoring service has occurred.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 19E91C1D-9BD0-468F-A7DA-915263C6689F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricExplorerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricExplorerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricExplorerResponseBody) SetDynamicCode(v string) *DescribeFabricExplorerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeFabricExplorerResponseBody) SetDynamicMessage(v string) *DescribeFabricExplorerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeFabricExplorerResponseBody) SetErrorCode(v int32) *DescribeFabricExplorerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricExplorerResponseBody) SetRequestId(v string) *DescribeFabricExplorerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricExplorerResponseBody) SetResult(v string) *DescribeFabricExplorerResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeFabricExplorerResponseBody) SetSuccess(v bool) *DescribeFabricExplorerResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricExplorerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricExplorerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricExplorerResponse) SetStatusCode(v int32) *DescribeFabricExplorerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricExplorerResponse) SetBody(v *DescribeFabricExplorerResponseBody) *DescribeFabricExplorerResponse {
	s.Body = v
	return s
}

type DescribeFabricInvitationCodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwl****
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
	// example:
	//
	// InvitationCodeNotFound
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// The invitation code could not found.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 77E25463-56CB-4F1C-8EB6-8ADDFE39EF0F
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeFabricInvitationCodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricInvitationCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricInvitationCodeResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeFabricInvitationCodeResponseBody) SetRequestId(v string) *DescribeFabricInvitationCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBody) SetResult(v *DescribeFabricInvitationCodeResponseBodyResult) *DescribeFabricInvitationCodeResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBody) SetSuccess(v bool) *DescribeFabricInvitationCodeResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricInvitationCodeResponseBodyResult struct {
	// example:
	//
	// true
	Accepted *bool `json:"Accepted,omitempty" xml:"Accepted,omitempty"`
	// example:
	//
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwlulg7
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// abc@126.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 1544411108000
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 2
	InvitationId *int32 `json:"InvitationId,omitempty" xml:"InvitationId,omitempty"`
	// example:
	//
	// 1544411108000
	SendTime *string `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	// example:
	//
	// 27534
	SenderBid *string `json:"SenderBid,omitempty" xml:"SenderBid,omitempty"`
	// example:
	//
	// 26345345
	SenderId *int64 `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// example:
	//
	// uid-35324
	SenderName *string `json:"SenderName,omitempty" xml:"SenderName,omitempty"`
	// example:
	//
	// http://baas.console.aliyun.test/invite?code=92e7ef1934892
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeFabricInvitationCodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricInvitationCodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetAccepted(v bool) *DescribeFabricInvitationCodeResponseBodyResult {
	s.Accepted = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetCode(v string) *DescribeFabricInvitationCodeResponseBodyResult {
	s.Code = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetConsortiumId(v string) *DescribeFabricInvitationCodeResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetEmail(v string) *DescribeFabricInvitationCodeResponseBodyResult {
	s.Email = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetExpireTime(v string) *DescribeFabricInvitationCodeResponseBodyResult {
	s.ExpireTime = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetInvitationId(v int32) *DescribeFabricInvitationCodeResponseBodyResult {
	s.InvitationId = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetSendTime(v string) *DescribeFabricInvitationCodeResponseBodyResult {
	s.SendTime = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetSenderBid(v string) *DescribeFabricInvitationCodeResponseBodyResult {
	s.SenderBid = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetSenderId(v int64) *DescribeFabricInvitationCodeResponseBodyResult {
	s.SenderId = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetSenderName(v string) *DescribeFabricInvitationCodeResponseBodyResult {
	s.SenderName = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponseBodyResult) SetUrl(v string) *DescribeFabricInvitationCodeResponseBodyResult {
	s.Url = &v
	return s
}

type DescribeFabricInvitationCodeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricInvitationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricInvitationCodeResponse) SetStatusCode(v int32) *DescribeFabricInvitationCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricInvitationCodeResponse) SetBody(v *DescribeFabricInvitationCodeResponseBody) *DescribeFabricInvitationCodeResponse {
	s.Body = v
	return s
}

type DescribeFabricInviterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 200
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// C445762C-8909-4BD8-A2BD-BB45BF2441D8
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeFabricInviterResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricInviterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricInviterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricInviterResponseBody) SetErrorCode(v int32) *DescribeFabricInviterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricInviterResponseBody) SetRequestId(v string) *DescribeFabricInviterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricInviterResponseBody) SetResult(v *DescribeFabricInviterResponseBodyResult) *DescribeFabricInviterResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricInviterResponseBody) SetSuccess(v bool) *DescribeFabricInviterResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricInviterResponseBodyResult struct {
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwl****
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// name
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	// example:
	//
	// 1544411108000
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 3524234
	InviterId *int64 `json:"InviterId,omitempty" xml:"InviterId,omitempty"`
	// example:
	//
	// name
	InviterName *string `json:"InviterName,omitempty" xml:"InviterName,omitempty"`
}

func (s DescribeFabricInviterResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricInviterResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricInviterResponseBodyResult) SetConsortiumId(v string) *DescribeFabricInviterResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricInviterResponseBodyResult) SetConsortiumName(v string) *DescribeFabricInviterResponseBodyResult {
	s.ConsortiumName = &v
	return s
}

func (s *DescribeFabricInviterResponseBodyResult) SetExpireTime(v string) *DescribeFabricInviterResponseBodyResult {
	s.ExpireTime = &v
	return s
}

func (s *DescribeFabricInviterResponseBodyResult) SetInviterId(v int64) *DescribeFabricInviterResponseBodyResult {
	s.InviterId = &v
	return s
}

func (s *DescribeFabricInviterResponseBodyResult) SetInviterName(v string) *DescribeFabricInviterResponseBodyResult {
	s.InviterName = &v
	return s
}

type DescribeFabricInviterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricInviterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricInviterResponse) SetStatusCode(v int32) *DescribeFabricInviterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricInviterResponse) SetBody(v *DescribeFabricInviterResponseBody) *DescribeFabricInviterResponse {
	s.Body = v
	return s
}

type DescribeFabricOrdererLogsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// consortium-test-iyuso3g773ki
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 50
	Lines *string `json:"Lines,omitempty" xml:"Lines,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// orderer1
	OrdererName *string `json:"OrdererName,omitempty" xml:"OrdererName,omitempty"`
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

func (s *DescribeFabricOrdererLogsRequest) SetLines(v string) *DescribeFabricOrdererLogsRequest {
	s.Lines = &v
	return s
}

func (s *DescribeFabricOrdererLogsRequest) SetOrdererName(v string) *DescribeFabricOrdererLogsRequest {
	s.OrdererName = &v
	return s
}

type DescribeFabricOrdererLogsResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 508EA949-6349-4C86-B8E7-1D90669382B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Kafka.Retry.Producer.RetryMax=3
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricOrdererLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrdererLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrdererLogsResponseBody) SetErrorCode(v int32) *DescribeFabricOrdererLogsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrdererLogsResponseBody) SetRequestId(v string) *DescribeFabricOrdererLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrdererLogsResponseBody) SetResult(v string) *DescribeFabricOrdererLogsResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeFabricOrdererLogsResponseBody) SetSuccess(v bool) *DescribeFabricOrdererLogsResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricOrdererLogsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricOrdererLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricOrdererLogsResponse) SetStatusCode(v int32) *DescribeFabricOrdererLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricOrdererLogsResponse) SetBody(v *DescribeFabricOrdererLogsResponseBody) *DescribeFabricOrdererLogsResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationRequest struct {
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w****
	OrganizationId *string                                 `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Tag            []*DescribeFabricOrganizationRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeFabricOrganizationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationRequest) SetLocation(v string) *DescribeFabricOrganizationRequest {
	s.Location = &v
	return s
}

func (s *DescribeFabricOrganizationRequest) SetOrganizationId(v string) *DescribeFabricOrganizationRequest {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricOrganizationRequest) SetTag(v []*DescribeFabricOrganizationRequestTag) *DescribeFabricOrganizationRequest {
	s.Tag = v
	return s
}

type DescribeFabricOrganizationRequestTag struct {
	// This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeFabricOrganizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricOrganizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBody) SetRequestId(v string) *DescribeFabricOrganizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBody) SetResult(v *DescribeFabricOrganizationResponseBodyResult) *DescribeFabricOrganizationResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricOrganizationResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricOrganizationResponseBodyResult struct {
	// example:
	//
	// client
	CANAME *string `json:"CANAME,omitempty" xml:"CANAME,omitempty"`
	// example:
	//
	// https://ca1.org1.alibabacloudbaas.com:31154
	CAUrl *string `json:"CAUrl,omitempty" xml:"CAUrl,omitempty"`
	// example:
	//
	// org1
	CodeName *string `json:"CodeName,omitempty" xml:"CodeName,omitempty"`
	// example:
	//
	// 2
	ConsortiumCount *int32 `json:"ConsortiumCount,omitempty" xml:"ConsortiumCount,omitempty"`
	// example:
	//
	// 1533025590
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// org1.alibabacloudbaas.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// org1MSP
	MSP *string `json:"MSP,omitempty" xml:"MSP,omitempty"`
	// example:
	//
	// This is org1
	OrganizationDescription *string `json:"OrganizationDescription,omitempty" xml:"OrganizationDescription,omitempty"`
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// org1
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	// example:
	//
	// 2544
	OwnerBid *string `json:"OwnerBid,omitempty" xml:"OwnerBid,omitempty"`
	// example:
	//
	// uid-23434
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 23434
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// example:
	//
	// 2
	PeerCount *int32 `json:"PeerCount,omitempty" xml:"PeerCount,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// basic
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	// example:
	//
	// Running
	State *string                                             `json:"State,omitempty" xml:"State,omitempty"`
	Tags  []*DescribeFabricOrganizationResponseBodyResultTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	// example:
	//
	// cn-hangzhou-1a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeFabricOrganizationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetCANAME(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.CANAME = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetCAUrl(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.CAUrl = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetCodeName(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.CodeName = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetConsortiumCount(v int32) *DescribeFabricOrganizationResponseBodyResult {
	s.ConsortiumCount = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetCreateTime(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetDomain(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetMSP(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.MSP = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetOrganizationDescription(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.OrganizationDescription = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetOrganizationId(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetOrganizationName(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetOwnerBid(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.OwnerBid = &v
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

func (s *DescribeFabricOrganizationResponseBodyResult) SetPeerCount(v int32) *DescribeFabricOrganizationResponseBodyResult {
	s.PeerCount = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetRegionId(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetRequestId(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetSpecName(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.SpecName = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetState(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetTags(v []*DescribeFabricOrganizationResponseBodyResultTags) *DescribeFabricOrganizationResponseBodyResult {
	s.Tags = v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetUserCount(v int32) *DescribeFabricOrganizationResponseBodyResult {
	s.UserCount = &v
	return s
}

func (s *DescribeFabricOrganizationResponseBodyResult) SetZoneId(v string) *DescribeFabricOrganizationResponseBodyResult {
	s.ZoneId = &v
	return s
}

type DescribeFabricOrganizationResponseBodyResultTags struct {
	// example:
	//
	// BlockChainOwner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// DepartmentA
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricOrganizationResponse) SetStatusCode(v int32) *DescribeFabricOrganizationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricOrganizationResponse) SetBody(v *DescribeFabricOrganizationResponseBody) *DescribeFabricOrganizationResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationChaincodePackageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DescribeFabricOrganizationChaincodePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationChaincodePackageRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationChaincodePackageRequest) SetOrganizationId(v string) *DescribeFabricOrganizationChaincodePackageRequest {
	s.OrganizationId = &v
	return s
}

type DescribeFabricOrganizationChaincodePackageResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ChaincodePackage `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricOrganizationChaincodePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationChaincodePackageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationChaincodePackageResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationChaincodePackageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodePackageResponseBody) SetMessage(v string) *DescribeFabricOrganizationChaincodePackageResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodePackageResponseBody) SetRequestId(v string) *DescribeFabricOrganizationChaincodePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodePackageResponseBody) SetResult(v []*ChaincodePackage) *DescribeFabricOrganizationChaincodePackageResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricOrganizationChaincodePackageResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationChaincodePackageResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricOrganizationChaincodePackageResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricOrganizationChaincodePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFabricOrganizationChaincodePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationChaincodePackageResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationChaincodePackageResponse) SetHeaders(v map[string]*string) *DescribeFabricOrganizationChaincodePackageResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricOrganizationChaincodePackageResponse) SetStatusCode(v int32) *DescribeFabricOrganizationChaincodePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodePackageResponse) SetBody(v *DescribeFabricOrganizationChaincodePackageResponseBody) *DescribeFabricOrganizationChaincodePackageResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationChaincodesRequest struct {
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DescribeFabricOrganizationChaincodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationChaincodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationChaincodesRequest) SetLocation(v string) *DescribeFabricOrganizationChaincodesRequest {
	s.Location = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesRequest) SetOrganizationId(v string) *DescribeFabricOrganizationChaincodesRequest {
	s.OrganizationId = &v
	return s
}

type DescribeFabricOrganizationChaincodesResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricOrganizationChaincodesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricOrganizationChaincodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationChaincodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationChaincodesResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationChaincodesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBody) SetRequestId(v string) *DescribeFabricOrganizationChaincodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBody) SetResult(v []*DescribeFabricOrganizationChaincodesResponseBodyResult) *DescribeFabricOrganizationChaincodesResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationChaincodesResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricOrganizationChaincodesResponseBodyResult struct {
	// example:
	//
	// code-sacc-1pr09q7jmo0np
	ChaincodeId *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	// example:
	//
	// mycc
	ChaincodeName *string `json:"ChaincodeName,omitempty" xml:"ChaincodeName,omitempty"`
	// example:
	//
	// 0.3
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" xml:"ChaincodeVersion,omitempty"`
	// example:
	//
	// chan-first-channel-31hlgpen5k5lig
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// first-channel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwlulg7
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1533025590
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// uid-23425
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 1533025590
	DeployTime *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	// example:
	//
	// OR (&#39;perf9141MSP.member&#39;)
	EndorsePolicy *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	// example:
	//
	// Installed
	Installed *string `json:"Installed,omitempty" xml:"Installed,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// RUNNING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeFabricOrganizationChaincodesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationChaincodesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetChaincodeId(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.ChaincodeId = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetChaincodeName(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.ChaincodeName = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetChaincodeVersion(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.ChaincodeVersion = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetChannelId(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.ChannelId = &v
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

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetCreateTime(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.CreateTime = &v
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

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetEndorsePolicy(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.EndorsePolicy = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetInstalled(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.Installed = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetMessage(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.Message = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponseBodyResult) SetState(v string) *DescribeFabricOrganizationChaincodesResponseBodyResult {
	s.State = &v
	return s
}

type DescribeFabricOrganizationChaincodesResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricOrganizationChaincodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricOrganizationChaincodesResponse) SetStatusCode(v int32) *DescribeFabricOrganizationChaincodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricOrganizationChaincodesResponse) SetBody(v *DescribeFabricOrganizationChaincodesResponseBody) *DescribeFabricOrganizationChaincodesResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationChannelsRequest struct {
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DescribeFabricOrganizationChannelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationChannelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationChannelsRequest) SetLocation(v string) *DescribeFabricOrganizationChannelsRequest {
	s.Location = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsRequest) SetOrganizationId(v string) *DescribeFabricOrganizationChannelsRequest {
	s.OrganizationId = &v
	return s
}

type DescribeFabricOrganizationChannelsResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricOrganizationChannelsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricOrganizationChannelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationChannelsResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationChannelsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBody) SetRequestId(v string) *DescribeFabricOrganizationChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBody) SetResult(v []*DescribeFabricOrganizationChannelsResponseBodyResult) *DescribeFabricOrganizationChannelsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationChannelsResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricOrganizationChannelsResponseBodyResult struct {
	// example:
	//
	// 2
	BatchTimeout *int32 `json:"BatchTimeout,omitempty" xml:"BatchTimeout,omitempty"`
	// example:
	//
	// 3
	BlockCount *int32 `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
	// example:
	//
	// 2
	ChaincodeCount *int32  `json:"ChaincodeCount,omitempty" xml:"ChaincodeCount,omitempty"`
	ChannelId      *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// mychannel
	ChannelName  *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// name
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	// example:
	//
	// 1533025590
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1533025590
	DeleteTime *string `json:"DeleteTime,omitempty" xml:"DeleteTime,omitempty"`
	// example:
	//
	// true
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// example:
	//
	// 12
	MaxMessageCount *int32 `json:"MaxMessageCount,omitempty" xml:"MaxMessageCount,omitempty"`
	// example:
	//
	// 3
	MemberCount *int32  `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	OwnerBid    *string `json:"OwnerBid,omitempty" xml:"OwnerBid,omitempty"`
	OwnerName   *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerUid    *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// example:
	//
	// 24
	PreferredMaxBytes *int32 `json:"PreferredMaxBytes,omitempty" xml:"PreferredMaxBytes,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// true
	SupportChannelConfig *bool `json:"SupportChannelConfig,omitempty" xml:"SupportChannelConfig,omitempty"`
	// example:
	//
	// 1533025590
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeFabricOrganizationChannelsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationChannelsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetBatchTimeout(v int32) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.BatchTimeout = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetBlockCount(v int32) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.BlockCount = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetChaincodeCount(v int32) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.ChaincodeCount = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetChannelId(v string) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.ChannelId = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetChannelName(v string) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetConsortiumId(v string) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetConsortiumName(v string) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.ConsortiumName = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetCreateTime(v string) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetDeleteTime(v string) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.DeleteTime = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetDeleted(v bool) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.Deleted = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetMaxMessageCount(v int32) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.MaxMessageCount = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetMemberCount(v int32) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.MemberCount = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetOwnerBid(v string) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.OwnerBid = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetOwnerName(v string) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetOwnerUid(v int64) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.OwnerUid = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetPreferredMaxBytes(v int32) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.PreferredMaxBytes = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetRequestId(v string) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetState(v string) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetSupportChannelConfig(v bool) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.SupportChannelConfig = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponseBodyResult) SetUpdateTime(v string) *DescribeFabricOrganizationChannelsResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type DescribeFabricOrganizationChannelsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricOrganizationChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFabricOrganizationChannelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationChannelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationChannelsResponse) SetHeaders(v map[string]*string) *DescribeFabricOrganizationChannelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponse) SetStatusCode(v int32) *DescribeFabricOrganizationChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricOrganizationChannelsResponse) SetBody(v *DescribeFabricOrganizationChannelsResponseBody) *DescribeFabricOrganizationChannelsResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationDeletableRequest struct {
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DescribeFabricOrganizationDeletableRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationDeletableRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationDeletableRequest) SetLocation(v string) *DescribeFabricOrganizationDeletableRequest {
	s.Location = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableRequest) SetOrganizationId(v string) *DescribeFabricOrganizationDeletableRequest {
	s.OrganizationId = &v
	return s
}

type DescribeFabricOrganizationDeletableResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeFabricOrganizationDeletableResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricOrganizationDeletableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationDeletableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationDeletableResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationDeletableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBody) SetRequestId(v string) *DescribeFabricOrganizationDeletableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBody) SetResult(v *DescribeFabricOrganizationDeletableResponseBodyResult) *DescribeFabricOrganizationDeletableResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationDeletableResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricOrganizationDeletableResponseBodyResult struct {
	// example:
	//
	// name
	CodeName *string `json:"CodeName,omitempty" xml:"CodeName,omitempty"`
	// example:
	//
	// true
	Deletable *bool `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
	// example:
	//
	// domain
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// desc
	OrganizationDescription *string `json:"OrganizationDescription,omitempty" xml:"OrganizationDescription,omitempty"`
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// orgname
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// zone
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeFabricOrganizationDeletableResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationDeletableResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationDeletableResponseBodyResult) SetCodeName(v string) *DescribeFabricOrganizationDeletableResponseBodyResult {
	s.CodeName = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBodyResult) SetDeletable(v bool) *DescribeFabricOrganizationDeletableResponseBodyResult {
	s.Deletable = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBodyResult) SetDomain(v string) *DescribeFabricOrganizationDeletableResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBodyResult) SetOrganizationDescription(v string) *DescribeFabricOrganizationDeletableResponseBodyResult {
	s.OrganizationDescription = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBodyResult) SetOrganizationId(v string) *DescribeFabricOrganizationDeletableResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBodyResult) SetOrganizationName(v string) *DescribeFabricOrganizationDeletableResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponseBodyResult) SetRegionId(v string) *DescribeFabricOrganizationDeletableResponseBodyResult {
	s.RegionId = &v
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

type DescribeFabricOrganizationDeletableResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricOrganizationDeletableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricOrganizationDeletableResponse) SetStatusCode(v int32) *DescribeFabricOrganizationDeletableResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricOrganizationDeletableResponse) SetBody(v *DescribeFabricOrganizationDeletableResponseBody) *DescribeFabricOrganizationDeletableResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationMembersRequest struct {
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DescribeFabricOrganizationMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationMembersRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationMembersRequest) SetLocation(v string) *DescribeFabricOrganizationMembersRequest {
	s.Location = &v
	return s
}

func (s *DescribeFabricOrganizationMembersRequest) SetOrganizationId(v string) *DescribeFabricOrganizationMembersRequest {
	s.OrganizationId = &v
	return s
}

type DescribeFabricOrganizationMembersResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricOrganizationMembersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricOrganizationMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationMembersResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBody) SetRequestId(v string) *DescribeFabricOrganizationMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBody) SetResult(v []*DescribeFabricOrganizationMembersResponseBodyResult) *DescribeFabricOrganizationMembersResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationMembersResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricOrganizationMembersResponseBodyResult struct {
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwl****
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// name
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// domain
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1533025590
	JoinedTime *string `json:"JoinedTime,omitempty" xml:"JoinedTime,omitempty"`
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// name
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeFabricOrganizationMembersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationMembersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationMembersResponseBodyResult) SetConsortiumId(v string) *DescribeFabricOrganizationMembersResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBodyResult) SetConsortiumName(v string) *DescribeFabricOrganizationMembersResponseBodyResult {
	s.ConsortiumName = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBodyResult) SetDescription(v string) *DescribeFabricOrganizationMembersResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBodyResult) SetDomain(v string) *DescribeFabricOrganizationMembersResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBodyResult) SetJoinedTime(v string) *DescribeFabricOrganizationMembersResponseBodyResult {
	s.JoinedTime = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBodyResult) SetOrganizationId(v string) *DescribeFabricOrganizationMembersResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBodyResult) SetOrganizationName(v string) *DescribeFabricOrganizationMembersResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponseBodyResult) SetState(v string) *DescribeFabricOrganizationMembersResponseBodyResult {
	s.State = &v
	return s
}

type DescribeFabricOrganizationMembersResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricOrganizationMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricOrganizationMembersResponse) SetStatusCode(v int32) *DescribeFabricOrganizationMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricOrganizationMembersResponse) SetBody(v *DescribeFabricOrganizationMembersResponseBody) *DescribeFabricOrganizationMembersResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationPeersRequest struct {
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DescribeFabricOrganizationPeersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationPeersRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationPeersRequest) SetLocation(v string) *DescribeFabricOrganizationPeersRequest {
	s.Location = &v
	return s
}

func (s *DescribeFabricOrganizationPeersRequest) SetOrganizationId(v string) *DescribeFabricOrganizationPeersRequest {
	s.OrganizationId = &v
	return s
}

type DescribeFabricOrganizationPeersResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricOrganizationPeersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricOrganizationPeersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationPeersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationPeersResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationPeersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBody) SetRequestId(v string) *DescribeFabricOrganizationPeersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBody) SetResult(v []*DescribeFabricOrganizationPeersResponseBodyResult) *DescribeFabricOrganizationPeersResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationPeersResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricOrganizationPeersResponseBodyResult struct {
	// example:
	//
	// 1533025590
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// domain
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// basic
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 10.0.0.2
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// example:
	//
	// 10.0.0.2
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// example:
	//
	// false
	IsAnchor *bool `json:"IsAnchor,omitempty" xml:"IsAnchor,omitempty"`
	// example:
	//
	// name
	OrganizationPeerName *string `json:"OrganizationPeerName,omitempty" xml:"OrganizationPeerName,omitempty"`
	// example:
	//
	// 1234
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 1533025590
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeFabricOrganizationPeersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationPeersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetCreateTime(v string) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetDomain(v string) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetInstanceType(v string) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.InstanceType = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetInternetIp(v string) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.InternetIp = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetIntranetIp(v string) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.IntranetIp = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetIsAnchor(v bool) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.IsAnchor = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetOrganizationPeerName(v string) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.OrganizationPeerName = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetPort(v int32) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.Port = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponseBodyResult) SetUpdateTime(v string) *DescribeFabricOrganizationPeersResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type DescribeFabricOrganizationPeersResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricOrganizationPeersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricOrganizationPeersResponse) SetStatusCode(v int32) *DescribeFabricOrganizationPeersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricOrganizationPeersResponse) SetBody(v *DescribeFabricOrganizationPeersResponseBody) *DescribeFabricOrganizationPeersResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationSpecsResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricOrganizationSpecsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricOrganizationSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationSpecsResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationSpecsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationSpecsResponseBody) SetRequestId(v string) *DescribeFabricOrganizationSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationSpecsResponseBody) SetResult(v []*DescribeFabricOrganizationSpecsResponseBodyResult) *DescribeFabricOrganizationSpecsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricOrganizationSpecsResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationSpecsResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricOrganizationSpecsResponseBodyResult struct {
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// basic
	OrganizationSpecsName *string `json:"OrganizationSpecsName,omitempty" xml:"OrganizationSpecsName,omitempty"`
	Title                 *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeFabricOrganizationSpecsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationSpecsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationSpecsResponseBodyResult) SetEnable(v bool) *DescribeFabricOrganizationSpecsResponseBodyResult {
	s.Enable = &v
	return s
}

func (s *DescribeFabricOrganizationSpecsResponseBodyResult) SetOrganizationSpecsName(v string) *DescribeFabricOrganizationSpecsResponseBodyResult {
	s.OrganizationSpecsName = &v
	return s
}

func (s *DescribeFabricOrganizationSpecsResponseBodyResult) SetTitle(v string) *DescribeFabricOrganizationSpecsResponseBodyResult {
	s.Title = &v
	return s
}

type DescribeFabricOrganizationSpecsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricOrganizationSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricOrganizationSpecsResponse) SetStatusCode(v int32) *DescribeFabricOrganizationSpecsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricOrganizationSpecsResponse) SetBody(v *DescribeFabricOrganizationSpecsResponseBody) *DescribeFabricOrganizationSpecsResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationUsersRequest struct {
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DescribeFabricOrganizationUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationUsersRequest) SetLocation(v string) *DescribeFabricOrganizationUsersRequest {
	s.Location = &v
	return s
}

func (s *DescribeFabricOrganizationUsersRequest) SetOrganizationId(v string) *DescribeFabricOrganizationUsersRequest {
	s.OrganizationId = &v
	return s
}

type DescribeFabricOrganizationUsersResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricOrganizationUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricOrganizationUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationUsersResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationUsersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBody) SetRequestId(v string) *DescribeFabricOrganizationUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBody) SetResult(v []*DescribeFabricOrganizationUsersResponseBodyResult) *DescribeFabricOrganizationUsersResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationUsersResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricOrganizationUsersResponseBodyResult struct {
	// example:
	//
	// foo=foo1,bar=bar1
	Attrs *string `json:"Attrs,omitempty" xml:"Attrs,omitempty"`
	// example:
	//
	// 26842
	CallerBid *string `json:"CallerBid,omitempty" xml:"CallerBid,omitempty"`
	// example:
	//
	// 84927598****
	CallerUid *int64 `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	// example:
	//
	// 1533025590
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1533025590
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// name
	FullName *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// username
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeFabricOrganizationUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationUsersResponseBodyResult) SetAttrs(v string) *DescribeFabricOrganizationUsersResponseBodyResult {
	s.Attrs = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBodyResult) SetCallerBid(v string) *DescribeFabricOrganizationUsersResponseBodyResult {
	s.CallerBid = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBodyResult) SetCallerUid(v int64) *DescribeFabricOrganizationUsersResponseBodyResult {
	s.CallerUid = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBodyResult) SetCreateTime(v string) *DescribeFabricOrganizationUsersResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBodyResult) SetExpireTime(v string) *DescribeFabricOrganizationUsersResponseBodyResult {
	s.ExpireTime = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponseBodyResult) SetFullName(v string) *DescribeFabricOrganizationUsersResponseBodyResult {
	s.FullName = &v
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

func (s *DescribeFabricOrganizationUsersResponseBodyResult) SetUsername(v string) *DescribeFabricOrganizationUsersResponseBodyResult {
	s.Username = &v
	return s
}

type DescribeFabricOrganizationUsersResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricOrganizationUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricOrganizationUsersResponse) SetStatusCode(v int32) *DescribeFabricOrganizationUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricOrganizationUsersResponse) SetBody(v *DescribeFabricOrganizationUsersResponseBody) *DescribeFabricOrganizationUsersResponse {
	s.Body = v
	return s
}

type DescribeFabricOrganizationsRequest struct {
	// example:
	//
	// cn-hangzhou
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
	// This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeFabricOrganizationsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricOrganizationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationsResponseBody) SetErrorCode(v int32) *DescribeFabricOrganizationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBody) SetRequestId(v string) *DescribeFabricOrganizationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBody) SetResult(v []*DescribeFabricOrganizationsResponseBodyResult) *DescribeFabricOrganizationsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFabricOrganizationsResponseBody) SetSuccess(v bool) *DescribeFabricOrganizationsResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricOrganizationsResponseBodyResult struct {
	// example:
	//
	// name
	CodeName *string `json:"CodeName,omitempty" xml:"CodeName,omitempty"`
	// example:
	//
	// 2
	ConsortiumCount *int32 `json:"ConsortiumCount,omitempty" xml:"ConsortiumCount,omitempty"`
	// example:
	//
	// 1533025590
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// domain
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// example:
	//
	// desc
	OrganizationDescription *string `json:"OrganizationDescription,omitempty" xml:"OrganizationDescription,omitempty"`
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// name
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	// example:
	//
	// bid
	OwnerBid *string `json:"OwnerBid,omitempty" xml:"OwnerBid,omitempty"`
	// example:
	//
	// name
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 1232
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// example:
	//
	// 3
	PeerCount *int32 `json:"PeerCount,omitempty" xml:"PeerCount,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// basic、
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	// example:
	//
	// Running
	State *string                                              `json:"State,omitempty" xml:"State,omitempty"`
	Tags  []*DescribeFabricOrganizationsResponseBodyResultTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	// example:
	//
	// zone
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeFabricOrganizationsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricOrganizationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetCodeName(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.CodeName = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetConsortiumCount(v int32) *DescribeFabricOrganizationsResponseBodyResult {
	s.ConsortiumCount = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetCreateTime(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetDomain(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetMajorVersion(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.MajorVersion = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetOrganizationDescription(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.OrganizationDescription = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetOrganizationId(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetOrganizationName(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetOwnerBid(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.OwnerBid = &v
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

func (s *DescribeFabricOrganizationsResponseBodyResult) SetPeerCount(v int32) *DescribeFabricOrganizationsResponseBodyResult {
	s.PeerCount = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetRegionId(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetRequestId(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetSpecName(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.SpecName = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetState(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetTags(v []*DescribeFabricOrganizationsResponseBodyResultTags) *DescribeFabricOrganizationsResponseBodyResult {
	s.Tags = v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetUserCount(v int32) *DescribeFabricOrganizationsResponseBodyResult {
	s.UserCount = &v
	return s
}

func (s *DescribeFabricOrganizationsResponseBodyResult) SetZoneId(v string) *DescribeFabricOrganizationsResponseBodyResult {
	s.ZoneId = &v
	return s
}

type DescribeFabricOrganizationsResponseBodyResultTags struct {
	// example:
	//
	// BlockChainOwner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// DepartmentA
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricOrganizationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricOrganizationsResponse) SetStatusCode(v int32) *DescribeFabricOrganizationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricOrganizationsResponse) SetBody(v *DescribeFabricOrganizationsResponseBody) *DescribeFabricOrganizationsResponse {
	s.Body = v
	return s
}

type DescribeFabricPeerLogsRequest struct {
	// example:
	//
	// 50
	Lines *string `json:"Lines,omitempty" xml:"Lines,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-test11-tzwny25vqicq
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peer1
	PeerName *string `json:"PeerName,omitempty" xml:"PeerName,omitempty"`
}

func (s DescribeFabricPeerLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricPeerLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFabricPeerLogsRequest) SetLines(v string) *DescribeFabricPeerLogsRequest {
	s.Lines = &v
	return s
}

func (s *DescribeFabricPeerLogsRequest) SetOrganizationId(v string) *DescribeFabricPeerLogsRequest {
	s.OrganizationId = &v
	return s
}

func (s *DescribeFabricPeerLogsRequest) SetPeerName(v string) *DescribeFabricPeerLogsRequest {
	s.PeerName = &v
	return s
}

type DescribeFabricPeerLogsResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 78A09CAF-4944-4CAE-B25D-F5B6BA27D9F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// xxxx
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFabricPeerLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFabricPeerLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFabricPeerLogsResponseBody) SetErrorCode(v int32) *DescribeFabricPeerLogsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFabricPeerLogsResponseBody) SetRequestId(v string) *DescribeFabricPeerLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFabricPeerLogsResponseBody) SetResult(v string) *DescribeFabricPeerLogsResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeFabricPeerLogsResponseBody) SetSuccess(v bool) *DescribeFabricPeerLogsResponseBody {
	s.Success = &v
	return s
}

type DescribeFabricPeerLogsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFabricPeerLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFabricPeerLogsResponse) SetStatusCode(v int32) *DescribeFabricPeerLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFabricPeerLogsResponse) SetBody(v *DescribeFabricPeerLogsResponseBody) *DescribeFabricPeerLogsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// D1597848-E708-522C-9275-3A202FC8B63A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetErrorCode(v int32) *DescribeRegionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// example:
	//
	// baas.cn-hangzhou.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeRootDomainResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// onaliyunbaas.com
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRootDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRootDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRootDomainResponseBody) SetErrorCode(v int32) *DescribeRootDomainResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeRootDomainResponseBody) SetRequestId(v string) *DescribeRootDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRootDomainResponseBody) SetResult(v string) *DescribeRootDomainResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeRootDomainResponseBody) SetSuccess(v bool) *DescribeRootDomainResponseBody {
	s.Success = &v
	return s
}

type DescribeRootDomainResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRootDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRootDomainResponse) SetStatusCode(v int32) *DescribeRootDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRootDomainResponse) SetBody(v *DescribeRootDomainResponseBody) *DescribeRootDomainResponse {
	s.Body = v
	return s
}

type DescribeTasksResponseBody struct {
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// msg
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeTasksResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeTasksResponseBody) SetRequestId(v string) *DescribeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTasksResponseBody) SetResult(v []*DescribeTasksResponseBodyResult) *DescribeTasksResponseBody {
	s.Result = v
	return s
}

func (s *DescribeTasksResponseBody) SetSuccess(v bool) *DescribeTasksResponseBody {
	s.Success = &v
	return s
}

type DescribeTasksResponseBodyResult struct {
	// example:
	//
	// action
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// false
	Handled *bool `json:"Handled,omitempty" xml:"Handled,omitempty"`
	// example:
	//
	// type
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// example:
	//
	// 1533025590
	RequestTime *int64 `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	// example:
	//
	// 1533025590
	ResponseTime *string `json:"ResponseTime,omitempty" xml:"ResponseTime,omitempty"`
	// example:
	//
	// result
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// sender
	Sender *string `json:"Sender,omitempty" xml:"Sender,omitempty"`
	// example:
	//
	// target
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// example:
	//
	// 33
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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

func (s *DescribeTasksResponseBodyResult) SetHandled(v bool) *DescribeTasksResponseBodyResult {
	s.Handled = &v
	return s
}

func (s *DescribeTasksResponseBodyResult) SetOperationType(v string) *DescribeTasksResponseBodyResult {
	s.OperationType = &v
	return s
}

func (s *DescribeTasksResponseBodyResult) SetRequestTime(v int64) *DescribeTasksResponseBodyResult {
	s.RequestTime = &v
	return s
}

func (s *DescribeTasksResponseBodyResult) SetResponseTime(v string) *DescribeTasksResponseBodyResult {
	s.ResponseTime = &v
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

func (s *DescribeTasksResponseBodyResult) SetTarget(v string) *DescribeTasksResponseBodyResult {
	s.Target = &v
	return s
}

func (s *DescribeTasksResponseBodyResult) SetTaskId(v int32) *DescribeTasksResponseBodyResult {
	s.TaskId = &v
	return s
}

type DescribeTasksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeTasksResponse) SetStatusCode(v int32) *DescribeTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTasksResponse) SetBody(v *DescribeTasksResponseBody) *DescribeTasksResponse {
	s.Body = v
	return s
}

type DownloadFabricOrganizationSDKRequest struct {
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// username
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DownloadFabricOrganizationSDKRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadFabricOrganizationSDKRequest) GoString() string {
	return s.String()
}

func (s *DownloadFabricOrganizationSDKRequest) SetLocation(v string) *DownloadFabricOrganizationSDKRequest {
	s.Location = &v
	return s
}

func (s *DownloadFabricOrganizationSDKRequest) SetOrganizationId(v string) *DownloadFabricOrganizationSDKRequest {
	s.OrganizationId = &v
	return s
}

func (s *DownloadFabricOrganizationSDKRequest) SetUsername(v string) *DownloadFabricOrganizationSDKRequest {
	s.Username = &v
	return s
}

type DownloadFabricOrganizationSDKResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DownloadFabricOrganizationSDKResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DownloadFabricOrganizationSDKResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadFabricOrganizationSDKResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadFabricOrganizationSDKResponseBody) SetErrorCode(v int32) *DownloadFabricOrganizationSDKResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DownloadFabricOrganizationSDKResponseBody) SetRequestId(v string) *DownloadFabricOrganizationSDKResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadFabricOrganizationSDKResponseBody) SetResult(v []*DownloadFabricOrganizationSDKResponseBodyResult) *DownloadFabricOrganizationSDKResponseBody {
	s.Result = v
	return s
}

func (s *DownloadFabricOrganizationSDKResponseBody) SetSuccess(v bool) *DownloadFabricOrganizationSDKResponseBody {
	s.Success = &v
	return s
}

type DownloadFabricOrganizationSDKResponseBodyResult struct {
	// example:
	//
	// base64string
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// dir/file
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DownloadFabricOrganizationSDKResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DownloadFabricOrganizationSDKResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DownloadFabricOrganizationSDKResponseBodyResult) SetContent(v string) *DownloadFabricOrganizationSDKResponseBodyResult {
	s.Content = &v
	return s
}

func (s *DownloadFabricOrganizationSDKResponseBodyResult) SetPath(v string) *DownloadFabricOrganizationSDKResponseBodyResult {
	s.Path = &v
	return s
}

type DownloadFabricOrganizationSDKResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadFabricOrganizationSDKResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DownloadFabricOrganizationSDKResponse) SetStatusCode(v int32) *DownloadFabricOrganizationSDKResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadFabricOrganizationSDKResponse) SetBody(v *DownloadFabricOrganizationSDKResponseBody) *DownloadFabricOrganizationSDKResponse {
	s.Body = v
	return s
}

type FreezeAntChainAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// accountname
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1q8B5R9p
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s FreezeAntChainAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s FreezeAntChainAccountRequest) GoString() string {
	return s.String()
}

func (s *FreezeAntChainAccountRequest) SetAccount(v string) *FreezeAntChainAccountRequest {
	s.Account = &v
	return s
}

func (s *FreezeAntChainAccountRequest) SetAntChainId(v string) *FreezeAntChainAccountRequest {
	s.AntChainId = &v
	return s
}

type FreezeAntChainAccountResponseBody struct {
	// example:
	//
	// 9F862442-F6AD-4B5F-8969-CE23DE1413CB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FreezeAntChainAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *FreezeAntChainAccountResponse) SetStatusCode(v int32) *FreezeAntChainAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *FreezeAntChainAccountResponse) SetBody(v *FreezeAntChainAccountResponseBody) *FreezeAntChainAccountResponse {
	s.Body = v
	return s
}

type InstallFabricChaincodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// chan-channelx-1l1hmckuuisxo
	ChaincodeId *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s InstallFabricChaincodeRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallFabricChaincodeRequest) GoString() string {
	return s.String()
}

func (s *InstallFabricChaincodeRequest) SetChaincodeId(v string) *InstallFabricChaincodeRequest {
	s.ChaincodeId = &v
	return s
}

func (s *InstallFabricChaincodeRequest) SetLocation(v string) *InstallFabricChaincodeRequest {
	s.Location = &v
	return s
}

func (s *InstallFabricChaincodeRequest) SetOrganizationId(v string) *InstallFabricChaincodeRequest {
	s.OrganizationId = &v
	return s
}

type InstallFabricChaincodeResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *InstallFabricChaincodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InstallFabricChaincodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallFabricChaincodeResponseBody) GoString() string {
	return s.String()
}

func (s *InstallFabricChaincodeResponseBody) SetErrorCode(v int32) *InstallFabricChaincodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InstallFabricChaincodeResponseBody) SetRequestId(v string) *InstallFabricChaincodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallFabricChaincodeResponseBody) SetResult(v *InstallFabricChaincodeResponseBodyResult) *InstallFabricChaincodeResponseBody {
	s.Result = v
	return s
}

func (s *InstallFabricChaincodeResponseBody) SetSuccess(v bool) *InstallFabricChaincodeResponseBody {
	s.Success = &v
	return s
}

type InstallFabricChaincodeResponseBodyResult struct {
	// example:
	//
	// cc-198jejf8f8chi8
	ChaincodeId *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	// example:
	//
	// mychannel
	ChaincodeName *string `json:"ChaincodeName,omitempty" xml:"ChaincodeName,omitempty"`
	// example:
	//
	// 1.0
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" xml:"ChaincodeVersion,omitempty"`
	// example:
	//
	// mychannel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwlulg7
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1533025590
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1533025590
	DeployTime *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	// example:
	//
	// OR ("perf9141MSP.member")
	EndorsePolicy *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	// example:
	//
	// input
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// false
	Install *bool `json:"Install,omitempty" xml:"Install,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// github.com/hyperledger/fabric-samples/chaincode/sacc
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// providerid
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// example:
	//
	// prividername
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// example:
	//
	// Pending
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s InstallFabricChaincodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s InstallFabricChaincodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InstallFabricChaincodeResponseBodyResult) SetChaincodeId(v string) *InstallFabricChaincodeResponseBodyResult {
	s.ChaincodeId = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetChaincodeName(v string) *InstallFabricChaincodeResponseBodyResult {
	s.ChaincodeName = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetChaincodeVersion(v string) *InstallFabricChaincodeResponseBodyResult {
	s.ChaincodeVersion = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetChannelName(v string) *InstallFabricChaincodeResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetConsortiumId(v string) *InstallFabricChaincodeResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetCreateTime(v string) *InstallFabricChaincodeResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetDeployTime(v string) *InstallFabricChaincodeResponseBodyResult {
	s.DeployTime = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetEndorsePolicy(v string) *InstallFabricChaincodeResponseBodyResult {
	s.EndorsePolicy = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetInput(v string) *InstallFabricChaincodeResponseBodyResult {
	s.Input = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetInstall(v bool) *InstallFabricChaincodeResponseBodyResult {
	s.Install = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetMessage(v string) *InstallFabricChaincodeResponseBodyResult {
	s.Message = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetPath(v string) *InstallFabricChaincodeResponseBodyResult {
	s.Path = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetProviderId(v string) *InstallFabricChaincodeResponseBodyResult {
	s.ProviderId = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetProviderName(v string) *InstallFabricChaincodeResponseBodyResult {
	s.ProviderName = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetState(v string) *InstallFabricChaincodeResponseBodyResult {
	s.State = &v
	return s
}

func (s *InstallFabricChaincodeResponseBodyResult) SetType(v int32) *InstallFabricChaincodeResponseBodyResult {
	s.Type = &v
	return s
}

type InstallFabricChaincodeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallFabricChaincodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *InstallFabricChaincodeResponse) SetStatusCode(v int32) *InstallFabricChaincodeResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallFabricChaincodeResponse) SetBody(v *InstallFabricChaincodeResponseBody) *InstallFabricChaincodeResponse {
	s.Body = v
	return s
}

type InstallFabricChaincodePackageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// chan-channelx-1l1hmckuuisxo
	ChaincodePackageId *string `json:"ChaincodePackageId,omitempty" xml:"ChaincodePackageId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s InstallFabricChaincodePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallFabricChaincodePackageRequest) GoString() string {
	return s.String()
}

func (s *InstallFabricChaincodePackageRequest) SetChaincodePackageId(v string) *InstallFabricChaincodePackageRequest {
	s.ChaincodePackageId = &v
	return s
}

func (s *InstallFabricChaincodePackageRequest) SetLocation(v string) *InstallFabricChaincodePackageRequest {
	s.Location = &v
	return s
}

func (s *InstallFabricChaincodePackageRequest) SetOrganizationId(v string) *InstallFabricChaincodePackageRequest {
	s.OrganizationId = &v
	return s
}

type InstallFabricChaincodePackageResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ChaincodePackage `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InstallFabricChaincodePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallFabricChaincodePackageResponseBody) GoString() string {
	return s.String()
}

func (s *InstallFabricChaincodePackageResponseBody) SetErrorCode(v int32) *InstallFabricChaincodePackageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InstallFabricChaincodePackageResponseBody) SetMessage(v string) *InstallFabricChaincodePackageResponseBody {
	s.Message = &v
	return s
}

func (s *InstallFabricChaincodePackageResponseBody) SetRequestId(v string) *InstallFabricChaincodePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallFabricChaincodePackageResponseBody) SetResult(v *ChaincodePackage) *InstallFabricChaincodePackageResponseBody {
	s.Result = v
	return s
}

func (s *InstallFabricChaincodePackageResponseBody) SetSuccess(v bool) *InstallFabricChaincodePackageResponseBody {
	s.Success = &v
	return s
}

type InstallFabricChaincodePackageResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallFabricChaincodePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallFabricChaincodePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallFabricChaincodePackageResponse) GoString() string {
	return s.String()
}

func (s *InstallFabricChaincodePackageResponse) SetHeaders(v map[string]*string) *InstallFabricChaincodePackageResponse {
	s.Headers = v
	return s
}

func (s *InstallFabricChaincodePackageResponse) SetStatusCode(v int32) *InstallFabricChaincodePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallFabricChaincodePackageResponse) SetBody(v *InstallFabricChaincodePackageResponseBody) *InstallFabricChaincodePackageResponse {
	s.Body = v
	return s
}

type InstantiateFabricChaincodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-198jejf8f8****
	ChaincodeId *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	// example:
	//
	// [{"requiredPeerCount":0,"name":"collectionName","maxPeerCount":3,"blockToLive":0,"policy":"OR(\\"Org1MSP.peer\\")"}]
	CollectionConfig *string `json:"CollectionConfig,omitempty" xml:"CollectionConfig,omitempty"`
	// example:
	//
	// OR (&#39;perf9141MSP.member&#39;)
	EndorsePolicy *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s InstantiateFabricChaincodeRequest) String() string {
	return tea.Prettify(s)
}

func (s InstantiateFabricChaincodeRequest) GoString() string {
	return s.String()
}

func (s *InstantiateFabricChaincodeRequest) SetChaincodeId(v string) *InstantiateFabricChaincodeRequest {
	s.ChaincodeId = &v
	return s
}

func (s *InstantiateFabricChaincodeRequest) SetCollectionConfig(v string) *InstantiateFabricChaincodeRequest {
	s.CollectionConfig = &v
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

func (s *InstantiateFabricChaincodeRequest) SetOrganizationId(v string) *InstantiateFabricChaincodeRequest {
	s.OrganizationId = &v
	return s
}

type InstantiateFabricChaincodeResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *InstantiateFabricChaincodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InstantiateFabricChaincodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstantiateFabricChaincodeResponseBody) GoString() string {
	return s.String()
}

func (s *InstantiateFabricChaincodeResponseBody) SetErrorCode(v int32) *InstantiateFabricChaincodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBody) SetRequestId(v string) *InstantiateFabricChaincodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBody) SetResult(v *InstantiateFabricChaincodeResponseBodyResult) *InstantiateFabricChaincodeResponseBody {
	s.Result = v
	return s
}

func (s *InstantiateFabricChaincodeResponseBody) SetSuccess(v bool) *InstantiateFabricChaincodeResponseBody {
	s.Success = &v
	return s
}

type InstantiateFabricChaincodeResponseBodyResult struct {
	// example:
	//
	// cc-198jejf8f8chi8
	ChaincodeId *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	// example:
	//
	// channelname
	ChaincodeName *string `json:"ChaincodeName,omitempty" xml:"ChaincodeName,omitempty"`
	// example:
	//
	// 1.0
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" xml:"ChaincodeVersion,omitempty"`
	// example:
	//
	// channelname
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwlulg7
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1533025590
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1533025590
	DeployTime *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	// example:
	//
	// OR (\\"aaaaaa1MSP.peer\\")
	EndorsePolicy *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	// example:
	//
	// input
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// false
	Install *bool `json:"Install,omitempty" xml:"Install,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// go-sdk-demo/chaincode/src
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// id
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// example:
	//
	// name
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s InstantiateFabricChaincodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s InstantiateFabricChaincodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetChaincodeId(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.ChaincodeId = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetChaincodeName(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.ChaincodeName = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetChaincodeVersion(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.ChaincodeVersion = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetChannelName(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetConsortiumId(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetCreateTime(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetDeployTime(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.DeployTime = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetEndorsePolicy(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.EndorsePolicy = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetInput(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.Input = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetInstall(v bool) *InstantiateFabricChaincodeResponseBodyResult {
	s.Install = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetMessage(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.Message = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetPath(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.Path = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetProviderId(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.ProviderId = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetProviderName(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.ProviderName = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetState(v string) *InstantiateFabricChaincodeResponseBodyResult {
	s.State = &v
	return s
}

func (s *InstantiateFabricChaincodeResponseBodyResult) SetType(v int32) *InstantiateFabricChaincodeResponseBodyResult {
	s.Type = &v
	return s
}

type InstantiateFabricChaincodeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstantiateFabricChaincodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *InstantiateFabricChaincodeResponse) SetStatusCode(v int32) *InstantiateFabricChaincodeResponse {
	s.StatusCode = &v
	return s
}

func (s *InstantiateFabricChaincodeResponse) SetBody(v *InstantiateFabricChaincodeResponseBody) *InstantiateFabricChaincodeResponse {
	s.Body = v
	return s
}

type JoinFabricChannelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// chan-channelx-1l1hmckuu****
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// accept
	Do *string `json:"Do,omitempty" xml:"Do,omitempty"`
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*JoinFabricChannelResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s JoinFabricChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinFabricChannelResponseBody) GoString() string {
	return s.String()
}

func (s *JoinFabricChannelResponseBody) SetErrorCode(v int32) *JoinFabricChannelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *JoinFabricChannelResponseBody) SetRequestId(v string) *JoinFabricChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinFabricChannelResponseBody) SetResult(v []*JoinFabricChannelResponseBodyResult) *JoinFabricChannelResponseBody {
	s.Result = v
	return s
}

func (s *JoinFabricChannelResponseBody) SetSuccess(v bool) *JoinFabricChannelResponseBody {
	s.Success = &v
	return s
}

type JoinFabricChannelResponseBodyResult struct {
	// example:
	//
	// 1533025590
	AcceptTime *string `json:"AcceptTime,omitempty" xml:"AcceptTime,omitempty"`
	// example:
	//
	// 1533025590
	ApproveTime *string `json:"ApproveTime,omitempty" xml:"ApproveTime,omitempty"`
	// example:
	//
	// channelid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1533025590
	ConfirmTime *string `json:"ConfirmTime,omitempty" xml:"ConfirmTime,omitempty"`
	// example:
	//
	// 1533025590
	DestroyTime *string `json:"DestroyTime,omitempty" xml:"DestroyTime,omitempty"`
	// example:
	//
	// 1533025590
	InviteTime *string `json:"InviteTime,omitempty" xml:"InviteTime,omitempty"`
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// true
	WithPeer *bool `json:"WithPeer,omitempty" xml:"WithPeer,omitempty"`
}

func (s JoinFabricChannelResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s JoinFabricChannelResponseBodyResult) GoString() string {
	return s.String()
}

func (s *JoinFabricChannelResponseBodyResult) SetAcceptTime(v string) *JoinFabricChannelResponseBodyResult {
	s.AcceptTime = &v
	return s
}

func (s *JoinFabricChannelResponseBodyResult) SetApproveTime(v string) *JoinFabricChannelResponseBodyResult {
	s.ApproveTime = &v
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

func (s *JoinFabricChannelResponseBodyResult) SetDestroyTime(v string) *JoinFabricChannelResponseBodyResult {
	s.DestroyTime = &v
	return s
}

func (s *JoinFabricChannelResponseBodyResult) SetInviteTime(v string) *JoinFabricChannelResponseBodyResult {
	s.InviteTime = &v
	return s
}

func (s *JoinFabricChannelResponseBodyResult) SetOrganizationId(v string) *JoinFabricChannelResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *JoinFabricChannelResponseBodyResult) SetState(v string) *JoinFabricChannelResponseBodyResult {
	s.State = &v
	return s
}

func (s *JoinFabricChannelResponseBodyResult) SetWithPeer(v bool) *JoinFabricChannelResponseBodyResult {
	s.WithPeer = &v
	return s
}

type JoinFabricChannelResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JoinFabricChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *JoinFabricChannelResponse) SetStatusCode(v int32) *JoinFabricChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinFabricChannelResponse) SetBody(v *JoinFabricChannelResponseBody) *JoinFabricChannelResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// example:
	//
	// MWJlMTY4ZmY4MzdmMDQzYmRlMTdjMDMxNDM0MWM4NDI3MTA0N2IzMQo=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// peers-tf000eihjli000shanghai-yqxvj52lruxx
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// organization
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value2
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
	// example:
	//
	// MWJlMTY4ZmY4MzdmMDQzYmRlMTdjMDMxNDM0MWM4NDI3MTA0N2IzMQo=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 21B8738D-254C-4933-9FC1-3C31B2611482
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
	// example:
	//
	// peers-tf000eihjli000shanghai-yqxvj52lruxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// organization
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// key1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// value1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ResetAntChainCertificateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pYogqb9v
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
	// example:
	//
	// 6ED983D3-C131-4104-8BDD-1F37992258E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAntChainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ResetAntChainCertificateResponse) SetStatusCode(v int32) *ResetAntChainCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAntChainCertificateResponse) SetBody(v *ResetAntChainCertificateResponseBody) *ResetAntChainCertificateResponse {
	s.Body = v
	return s
}

type ResetAntChainUserCertificateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bDXK6boZ
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// uid-128712635330****
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	// example:
	//
	// 3EB828D1-1E9D-4EC2-A002-139FE998D674
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAntChainUserCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ResetAntChainUserCertificateResponse) SetStatusCode(v int32) *ResetAntChainUserCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAntChainUserCertificateResponse) SetBody(v *ResetAntChainUserCertificateResponseBody) *ResetAntChainUserCertificateResponse {
	s.Body = v
	return s
}

type ResetFabricOrganizationUserPasswordRequest struct {
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pwd
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// username
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ResetFabricOrganizationUserPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetFabricOrganizationUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetFabricOrganizationUserPasswordRequest) SetLocation(v string) *ResetFabricOrganizationUserPasswordRequest {
	s.Location = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordRequest) SetOrganizationId(v string) *ResetFabricOrganizationUserPasswordRequest {
	s.OrganizationId = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordRequest) SetPassword(v string) *ResetFabricOrganizationUserPasswordRequest {
	s.Password = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordRequest) SetUsername(v string) *ResetFabricOrganizationUserPasswordRequest {
	s.Username = &v
	return s
}

type ResetFabricOrganizationUserPasswordResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ResetFabricOrganizationUserPasswordResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResetFabricOrganizationUserPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetFabricOrganizationUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetFabricOrganizationUserPasswordResponseBody) SetErrorCode(v int32) *ResetFabricOrganizationUserPasswordResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponseBody) SetRequestId(v string) *ResetFabricOrganizationUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponseBody) SetResult(v *ResetFabricOrganizationUserPasswordResponseBodyResult) *ResetFabricOrganizationUserPasswordResponseBody {
	s.Result = v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponseBody) SetSuccess(v bool) *ResetFabricOrganizationUserPasswordResponseBody {
	s.Success = &v
	return s
}

type ResetFabricOrganizationUserPasswordResponseBodyResult struct {
	// example:
	//
	// 1533025590
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1533025590
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// username
	Fullname *string `json:"Fullname,omitempty" xml:"Fullname,omitempty"`
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// username
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ResetFabricOrganizationUserPasswordResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ResetFabricOrganizationUserPasswordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ResetFabricOrganizationUserPasswordResponseBodyResult) SetCreateTime(v string) *ResetFabricOrganizationUserPasswordResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponseBodyResult) SetExpireTime(v string) *ResetFabricOrganizationUserPasswordResponseBodyResult {
	s.ExpireTime = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponseBodyResult) SetFullname(v string) *ResetFabricOrganizationUserPasswordResponseBodyResult {
	s.Fullname = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponseBodyResult) SetOrganizationId(v string) *ResetFabricOrganizationUserPasswordResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponseBodyResult) SetPassword(v string) *ResetFabricOrganizationUserPasswordResponseBodyResult {
	s.Password = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponseBodyResult) SetUsername(v string) *ResetFabricOrganizationUserPasswordResponseBodyResult {
	s.Username = &v
	return s
}

type ResetFabricOrganizationUserPasswordResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetFabricOrganizationUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ResetFabricOrganizationUserPasswordResponse) SetStatusCode(v int32) *ResetFabricOrganizationUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetFabricOrganizationUserPasswordResponse) SetBody(v *ResetFabricOrganizationUserPasswordResponseBody) *ResetFabricOrganizationUserPasswordResponse {
	s.Body = v
	return s
}

type SubmitFabricChaincodeDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sacc_v1.0:698d4e73c5f1d0239fbf0313ac79ddfaf44d5a803ecb02a4d1983cca4a706860
	ChaincodePackageId *string `json:"ChaincodePackageId,omitempty" xml:"ChaincodePackageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" xml:"ChaincodeVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chan-channelx-1l1hmckuuisxo
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// collectionconfig
	CollectionConfig *string `json:"CollectionConfig,omitempty" xml:"CollectionConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OR ("perf9141MSP.member")
	EndorsePolicy *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	// example:
	//
	// false
	InitRequired *bool `json:"InitRequired,omitempty" xml:"InitRequired,omitempty"`
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sacc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s SubmitFabricChaincodeDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitFabricChaincodeDefinitionRequest) GoString() string {
	return s.String()
}

func (s *SubmitFabricChaincodeDefinitionRequest) SetChaincodePackageId(v string) *SubmitFabricChaincodeDefinitionRequest {
	s.ChaincodePackageId = &v
	return s
}

func (s *SubmitFabricChaincodeDefinitionRequest) SetChaincodeVersion(v string) *SubmitFabricChaincodeDefinitionRequest {
	s.ChaincodeVersion = &v
	return s
}

func (s *SubmitFabricChaincodeDefinitionRequest) SetChannelId(v string) *SubmitFabricChaincodeDefinitionRequest {
	s.ChannelId = &v
	return s
}

func (s *SubmitFabricChaincodeDefinitionRequest) SetCollectionConfig(v string) *SubmitFabricChaincodeDefinitionRequest {
	s.CollectionConfig = &v
	return s
}

func (s *SubmitFabricChaincodeDefinitionRequest) SetEndorsePolicy(v string) *SubmitFabricChaincodeDefinitionRequest {
	s.EndorsePolicy = &v
	return s
}

func (s *SubmitFabricChaincodeDefinitionRequest) SetInitRequired(v bool) *SubmitFabricChaincodeDefinitionRequest {
	s.InitRequired = &v
	return s
}

func (s *SubmitFabricChaincodeDefinitionRequest) SetLocation(v string) *SubmitFabricChaincodeDefinitionRequest {
	s.Location = &v
	return s
}

func (s *SubmitFabricChaincodeDefinitionRequest) SetName(v string) *SubmitFabricChaincodeDefinitionRequest {
	s.Name = &v
	return s
}

func (s *SubmitFabricChaincodeDefinitionRequest) SetOrganizationId(v string) *SubmitFabricChaincodeDefinitionRequest {
	s.OrganizationId = &v
	return s
}

type SubmitFabricChaincodeDefinitionResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ChaincodeVO `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitFabricChaincodeDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitFabricChaincodeDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitFabricChaincodeDefinitionResponseBody) SetErrorCode(v int32) *SubmitFabricChaincodeDefinitionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitFabricChaincodeDefinitionResponseBody) SetMessage(v string) *SubmitFabricChaincodeDefinitionResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitFabricChaincodeDefinitionResponseBody) SetRequestId(v string) *SubmitFabricChaincodeDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitFabricChaincodeDefinitionResponseBody) SetResult(v *ChaincodeVO) *SubmitFabricChaincodeDefinitionResponseBody {
	s.Result = v
	return s
}

func (s *SubmitFabricChaincodeDefinitionResponseBody) SetSuccess(v bool) *SubmitFabricChaincodeDefinitionResponseBody {
	s.Success = &v
	return s
}

type SubmitFabricChaincodeDefinitionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitFabricChaincodeDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitFabricChaincodeDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitFabricChaincodeDefinitionResponse) GoString() string {
	return s.String()
}

func (s *SubmitFabricChaincodeDefinitionResponse) SetHeaders(v map[string]*string) *SubmitFabricChaincodeDefinitionResponse {
	s.Headers = v
	return s
}

func (s *SubmitFabricChaincodeDefinitionResponse) SetStatusCode(v int32) *SubmitFabricChaincodeDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitFabricChaincodeDefinitionResponse) SetBody(v *SubmitFabricChaincodeDefinitionResponseBody) *SubmitFabricChaincodeDefinitionResponse {
	s.Body = v
	return s
}

type SynchronizeFabricChaincodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// code-sacc-1pr09q7jm****
	ChaincodeId *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s SynchronizeFabricChaincodeRequest) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeFabricChaincodeRequest) GoString() string {
	return s.String()
}

func (s *SynchronizeFabricChaincodeRequest) SetChaincodeId(v string) *SynchronizeFabricChaincodeRequest {
	s.ChaincodeId = &v
	return s
}

func (s *SynchronizeFabricChaincodeRequest) SetOrganizationId(v string) *SynchronizeFabricChaincodeRequest {
	s.OrganizationId = &v
	return s
}

type SynchronizeFabricChaincodeResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 1890FA4F-067A-4CE9-AC9B-2BD2E58FB5D3
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *SynchronizeFabricChaincodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SynchronizeFabricChaincodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeFabricChaincodeResponseBody) GoString() string {
	return s.String()
}

func (s *SynchronizeFabricChaincodeResponseBody) SetErrorCode(v int32) *SynchronizeFabricChaincodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBody) SetRequestId(v string) *SynchronizeFabricChaincodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBody) SetResult(v *SynchronizeFabricChaincodeResponseBodyResult) *SynchronizeFabricChaincodeResponseBody {
	s.Result = v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBody) SetSuccess(v bool) *SynchronizeFabricChaincodeResponseBody {
	s.Success = &v
	return s
}

type SynchronizeFabricChaincodeResponseBodyResult struct {
	// example:
	//
	// code-sacc-1pr09q7jm****
	ChaincodeId *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	// example:
	//
	// mychaincode
	ChaincodeName *string `json:"ChaincodeName,omitempty" xml:"ChaincodeName,omitempty"`
	// example:
	//
	// 1.0
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" xml:"ChaincodeVersion,omitempty"`
	// example:
	//
	// chan-channelx-1l1hmckuuisxo
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// consortium-lianmenyumingyi-hc5d1bwl****
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1533025590
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1533025590
	DeployTime *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	// example:
	//
	// OR (\\"aaaaaa1MSP.peer\\")
	EndorsePolicy *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	// example:
	//
	// input
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// false
	Install *bool `json:"Install,omitempty" xml:"Install,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// go-sdk-demo/chaincode/src
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// provider
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// example:
	//
	// name
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SynchronizeFabricChaincodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeFabricChaincodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetChaincodeId(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.ChaincodeId = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetChaincodeName(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.ChaincodeName = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetChaincodeVersion(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.ChaincodeVersion = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetChannelName(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetConsortiumId(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetCreateTime(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetDeployTime(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.DeployTime = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetEndorsePolicy(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.EndorsePolicy = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetInput(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.Input = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetInstall(v bool) *SynchronizeFabricChaincodeResponseBodyResult {
	s.Install = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetMessage(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.Message = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetPath(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.Path = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetProviderId(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.ProviderId = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetProviderName(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.ProviderName = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetState(v string) *SynchronizeFabricChaincodeResponseBodyResult {
	s.State = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponseBodyResult) SetType(v int32) *SynchronizeFabricChaincodeResponseBodyResult {
	s.Type = &v
	return s
}

type SynchronizeFabricChaincodeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SynchronizeFabricChaincodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SynchronizeFabricChaincodeResponse) SetStatusCode(v int32) *SynchronizeFabricChaincodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SynchronizeFabricChaincodeResponse) SetBody(v *SynchronizeFabricChaincodeResponseBody) *SynchronizeFabricChaincodeResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// example:
	//
	// peers-tf000eihjli000beijing-bl6rdwi82tget
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// organization
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
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
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 21B8738D-254C-4933-9FC1-3C31B2611482
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetErrorCode(v int32) *TagResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) SetResult(v bool) *TagResourcesResponseBody {
	s.Result = &v
	return s
}

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnfreezeAntChainAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// accountname
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1q8B5R9p
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
}

func (s UnfreezeAntChainAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s UnfreezeAntChainAccountRequest) GoString() string {
	return s.String()
}

func (s *UnfreezeAntChainAccountRequest) SetAccount(v string) *UnfreezeAntChainAccountRequest {
	s.Account = &v
	return s
}

func (s *UnfreezeAntChainAccountRequest) SetAntChainId(v string) *UnfreezeAntChainAccountRequest {
	s.AntChainId = &v
	return s
}

type UnfreezeAntChainAccountResponseBody struct {
	// example:
	//
	// 6ED983D3-C131-4104-8BDD-1F37992258E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnfreezeAntChainAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UnfreezeAntChainAccountResponse) SetStatusCode(v int32) *UnfreezeAntChainAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *UnfreezeAntChainAccountResponse) SetBody(v *UnfreezeAntChainAccountResponseBody) *UnfreezeAntChainAccountResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// example:
	//
	// true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// example:
	//
	// peers-tf000eihjli000shanghai-yqxvj52lruxx
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// organization
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// key1
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 21B8738D-254C-4933-9FC1-3C31B2611482
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetErrorCode(v int32) *UntagResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetResult(v bool) *UntagResourcesResponseBody {
	s.Result = &v
	return s
}

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateAntChainRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pYogqb9v
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
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
	// example:
	//
	// F6A52E53-2CCB-4C3C-A44D-45FAEBC24E47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAntChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateAntChainResponse) SetStatusCode(v int32) *UpdateAntChainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAntChainResponse) SetBody(v *UpdateAntChainResponseBody) *UpdateAntChainResponse {
	s.Body = v
	return s
}

type UpdateAntChainConsortiumRequest struct {
	ConsortiumDescription *string `json:"ConsortiumDescription,omitempty" xml:"ConsortiumDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// q0oWq92P
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// This parameter is required.
	ConsortiumName *string `json:"ConsortiumName,omitempty" xml:"ConsortiumName,omitempty"`
}

func (s UpdateAntChainConsortiumRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainConsortiumRequest) GoString() string {
	return s.String()
}

func (s *UpdateAntChainConsortiumRequest) SetConsortiumDescription(v string) *UpdateAntChainConsortiumRequest {
	s.ConsortiumDescription = &v
	return s
}

func (s *UpdateAntChainConsortiumRequest) SetConsortiumId(v string) *UpdateAntChainConsortiumRequest {
	s.ConsortiumId = &v
	return s
}

func (s *UpdateAntChainConsortiumRequest) SetConsortiumName(v string) *UpdateAntChainConsortiumRequest {
	s.ConsortiumName = &v
	return s
}

type UpdateAntChainConsortiumResponseBody struct {
	// example:
	//
	// 1E55C947-19B4-4AAD-9D22-0B8DA8904B3D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAntChainConsortiumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateAntChainConsortiumResponse) SetStatusCode(v int32) *UpdateAntChainConsortiumResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAntChainConsortiumResponse) SetBody(v *UpdateAntChainConsortiumResponseBody) *UpdateAntChainConsortiumResponse {
	s.Body = v
	return s
}

type UpdateAntChainContractContentRequest struct {
	// example:
	//
	// update content
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// R6XMn59e
	ContentId *string `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	// example:
	//
	// newFile
	ContentName *string `json:"ContentName,omitempty" xml:"ContentName,omitempty"`
	// example:
	//
	// R38DYDop
	ParentContentId *string `json:"ParentContentId,omitempty" xml:"ParentContentId,omitempty"`
}

func (s UpdateAntChainContractContentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainContractContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAntChainContractContentRequest) SetContent(v string) *UpdateAntChainContractContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateAntChainContractContentRequest) SetContentId(v string) *UpdateAntChainContractContentRequest {
	s.ContentId = &v
	return s
}

func (s *UpdateAntChainContractContentRequest) SetContentName(v string) *UpdateAntChainContractContentRequest {
	s.ContentName = &v
	return s
}

func (s *UpdateAntChainContractContentRequest) SetParentContentId(v string) *UpdateAntChainContractContentRequest {
	s.ParentContentId = &v
	return s
}

type UpdateAntChainContractContentResponseBody struct {
	// example:
	//
	// 08CBD579-C430-4F1F-9011-C120FA1E9B97
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAntChainContractContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateAntChainContractContentResponse) SetStatusCode(v int32) *UpdateAntChainContractContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAntChainContractContentResponse) SetBody(v *UpdateAntChainContractContentResponseBody) *UpdateAntChainContractContentResponse {
	s.Body = v
	return s
}

type UpdateAntChainContractProjectRequest struct {
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bDXK1b8Z
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// v1.1.1
	ProjectVersion *string `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
}

func (s UpdateAntChainContractProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainContractProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateAntChainContractProjectRequest) SetProjectDescription(v string) *UpdateAntChainContractProjectRequest {
	s.ProjectDescription = &v
	return s
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

type UpdateAntChainContractProjectResponseBody struct {
	// example:
	//
	// 9EBD381E-A19E-4875-8771-EDA08A0416D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAntChainContractProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateAntChainContractProjectResponse) SetStatusCode(v int32) *UpdateAntChainContractProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAntChainContractProjectResponse) SetBody(v *UpdateAntChainContractProjectResponseBody) *UpdateAntChainContractProjectResponse {
	s.Body = v
	return s
}

type UpdateAntChainMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// AroN3X2l
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1287126353301234
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// updatename
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
}

func (s UpdateAntChainMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAntChainMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateAntChainMemberRequest) SetConsortiumId(v string) *UpdateAntChainMemberRequest {
	s.ConsortiumId = &v
	return s
}

func (s *UpdateAntChainMemberRequest) SetMemberId(v string) *UpdateAntChainMemberRequest {
	s.MemberId = &v
	return s
}

func (s *UpdateAntChainMemberRequest) SetMemberName(v string) *UpdateAntChainMemberRequest {
	s.MemberName = &v
	return s
}

type UpdateAntChainMemberResponseBody struct {
	// example:
	//
	// 5BA242B2-AB0C-4D06-94EB-23CEC4F7B95E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAntChainMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateAntChainMemberResponse) SetStatusCode(v int32) *UpdateAntChainMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAntChainMemberResponse) SetBody(v *UpdateAntChainMemberResponseBody) *UpdateAntChainMemberResponse {
	s.Body = v
	return s
}

type UpdateAntChainQRCodeAuthorizationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bDXK6boZ
	AntChainId *string `json:"AntChainId,omitempty" xml:"AntChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALL_USER_AUTHORIZATION
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MINI_APP_BROWSER_TRANSACTION
	QRCodeType *string `json:"QRCodeType,omitempty" xml:"QRCodeType,omitempty"`
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

func (s *UpdateAntChainQRCodeAuthorizationRequest) SetAuthorizationType(v string) *UpdateAntChainQRCodeAuthorizationRequest {
	s.AuthorizationType = &v
	return s
}

func (s *UpdateAntChainQRCodeAuthorizationRequest) SetQRCodeType(v string) *UpdateAntChainQRCodeAuthorizationRequest {
	s.QRCodeType = &v
	return s
}

type UpdateAntChainQRCodeAuthorizationResponseBody struct {
	// example:
	//
	// D68D66B6-1964-4073-8714-B49F5EF1AEFC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAntChainQRCodeAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateAntChainQRCodeAuthorizationResponse) SetStatusCode(v int32) *UpdateAntChainQRCodeAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAntChainQRCodeAuthorizationResponse) SetBody(v *UpdateAntChainQRCodeAuthorizationResponseBody) *UpdateAntChainQRCodeAuthorizationResponse {
	s.Body = v
	return s
}

type UpgradeFabricChaincodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// code-sacc-fk9z4r97****
	ChaincodeId *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	// example:
	//
	// [{"requiredPeerCount":0,"name":"collectionName","maxPeerCount":3,"blockToLive":0,"policy":"OR(\\"Org1MSP.peer\\")"}]
	CollectionConfig *string `json:"CollectionConfig,omitempty" xml:"CollectionConfig,omitempty"`
	// example:
	//
	// OR ("wholesalerMSP.peer")
	EndorsePolicy *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-wholesaler-15ix77tof****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s UpgradeFabricChaincodeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFabricChaincodeRequest) GoString() string {
	return s.String()
}

func (s *UpgradeFabricChaincodeRequest) SetChaincodeId(v string) *UpgradeFabricChaincodeRequest {
	s.ChaincodeId = &v
	return s
}

func (s *UpgradeFabricChaincodeRequest) SetCollectionConfig(v string) *UpgradeFabricChaincodeRequest {
	s.CollectionConfig = &v
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

func (s *UpgradeFabricChaincodeRequest) SetOrganizationId(v string) *UpgradeFabricChaincodeRequest {
	s.OrganizationId = &v
	return s
}

type UpgradeFabricChaincodeResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// A6CC6C63-2D71-4D0C-AEBE-E95F0127C514
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpgradeFabricChaincodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeFabricChaincodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFabricChaincodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeFabricChaincodeResponseBody) SetErrorCode(v int32) *UpgradeFabricChaincodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBody) SetRequestId(v string) *UpgradeFabricChaincodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBody) SetResult(v *UpgradeFabricChaincodeResponseBodyResult) *UpgradeFabricChaincodeResponseBody {
	s.Result = v
	return s
}

func (s *UpgradeFabricChaincodeResponseBody) SetSuccess(v bool) *UpgradeFabricChaincodeResponseBody {
	s.Success = &v
	return s
}

type UpgradeFabricChaincodeResponseBodyResult struct {
	// example:
	//
	// code-sacc-fk9z4r97****
	ChaincodeId *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	// example:
	//
	// mychaincode
	ChaincodeName *string `json:"ChaincodeName,omitempty" xml:"ChaincodeName,omitempty"`
	// example:
	//
	// 1.0
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" xml:"ChaincodeVersion,omitempty"`
	// example:
	//
	// cc
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// consortium-supplychain-1pxzsp5tb****
	ConsortiumId *string `json:"ConsortiumId,omitempty" xml:"ConsortiumId,omitempty"`
	// example:
	//
	// 1545215617000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1545215617000
	DeployTime *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	// example:
	//
	// OR ("wholesalerMSP.peer")
	EndorsePolicy *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	// example:
	//
	// input
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// false
	Install *bool `json:"Install,omitempty" xml:"Install,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// github.com/sacc
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// 1928592
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// example:
	//
	// uid-1928592
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpgradeFabricChaincodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFabricChaincodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetChaincodeId(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.ChaincodeId = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetChaincodeName(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.ChaincodeName = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetChaincodeVersion(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.ChaincodeVersion = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetChannelName(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.ChannelName = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetConsortiumId(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.ConsortiumId = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetCreateTime(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetDeployTime(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.DeployTime = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetEndorsePolicy(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.EndorsePolicy = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetInput(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.Input = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetInstall(v bool) *UpgradeFabricChaincodeResponseBodyResult {
	s.Install = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetMessage(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.Message = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetPath(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.Path = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetProviderId(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.ProviderId = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetProviderName(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.ProviderName = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetState(v string) *UpgradeFabricChaincodeResponseBodyResult {
	s.State = &v
	return s
}

func (s *UpgradeFabricChaincodeResponseBodyResult) SetType(v int32) *UpgradeFabricChaincodeResponseBodyResult {
	s.Type = &v
	return s
}

type UpgradeFabricChaincodeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeFabricChaincodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpgradeFabricChaincodeResponse) SetStatusCode(v int32) *UpgradeFabricChaincodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeFabricChaincodeResponse) SetBody(v *UpgradeFabricChaincodeResponseBody) *UpgradeFabricChaincodeResponse {
	s.Body = v
	return s
}

type UpgradeFabricChaincodeDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// code-sacc-jbkuldd3i7ad
	ChaincodeId *string `json:"ChaincodeId,omitempty" xml:"ChaincodeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sacc_v1.0:698d4e73c5f1d0239fbf0313ac79ddfaf44d5a803ecb02a4d1983cca4a706860
	ChaincodePackageId *string `json:"ChaincodePackageId,omitempty" xml:"ChaincodePackageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.1
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" xml:"ChaincodeVersion,omitempty"`
	// example:
	//
	// collectionconfig
	CollectionConfig *string `json:"CollectionConfig,omitempty" xml:"CollectionConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OR ("perf9141MSP.member")
	EndorsePolicy *string `json:"EndorsePolicy,omitempty" xml:"EndorsePolicy,omitempty"`
	// example:
	//
	// true
	InitRequired *bool `json:"InitRequired,omitempty" xml:"InitRequired,omitempty"`
	// example:
	//
	// cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// peers-aaaaaa2-1eqnj5o5w9dt3
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s UpgradeFabricChaincodeDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFabricChaincodeDefinitionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeFabricChaincodeDefinitionRequest) SetChaincodeId(v string) *UpgradeFabricChaincodeDefinitionRequest {
	s.ChaincodeId = &v
	return s
}

func (s *UpgradeFabricChaincodeDefinitionRequest) SetChaincodePackageId(v string) *UpgradeFabricChaincodeDefinitionRequest {
	s.ChaincodePackageId = &v
	return s
}

func (s *UpgradeFabricChaincodeDefinitionRequest) SetChaincodeVersion(v string) *UpgradeFabricChaincodeDefinitionRequest {
	s.ChaincodeVersion = &v
	return s
}

func (s *UpgradeFabricChaincodeDefinitionRequest) SetCollectionConfig(v string) *UpgradeFabricChaincodeDefinitionRequest {
	s.CollectionConfig = &v
	return s
}

func (s *UpgradeFabricChaincodeDefinitionRequest) SetEndorsePolicy(v string) *UpgradeFabricChaincodeDefinitionRequest {
	s.EndorsePolicy = &v
	return s
}

func (s *UpgradeFabricChaincodeDefinitionRequest) SetInitRequired(v bool) *UpgradeFabricChaincodeDefinitionRequest {
	s.InitRequired = &v
	return s
}

func (s *UpgradeFabricChaincodeDefinitionRequest) SetLocation(v string) *UpgradeFabricChaincodeDefinitionRequest {
	s.Location = &v
	return s
}

func (s *UpgradeFabricChaincodeDefinitionRequest) SetOrganizationId(v string) *UpgradeFabricChaincodeDefinitionRequest {
	s.OrganizationId = &v
	return s
}

type UpgradeFabricChaincodeDefinitionResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1E55C947-19B4-4AAD-9D22-0B8DA8904B3D
	RequestId *string      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ChaincodeVO `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeFabricChaincodeDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFabricChaincodeDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeFabricChaincodeDefinitionResponseBody) SetErrorCode(v int32) *UpgradeFabricChaincodeDefinitionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpgradeFabricChaincodeDefinitionResponseBody) SetMessage(v string) *UpgradeFabricChaincodeDefinitionResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeFabricChaincodeDefinitionResponseBody) SetRequestId(v string) *UpgradeFabricChaincodeDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeFabricChaincodeDefinitionResponseBody) SetResult(v *ChaincodeVO) *UpgradeFabricChaincodeDefinitionResponseBody {
	s.Result = v
	return s
}

func (s *UpgradeFabricChaincodeDefinitionResponseBody) SetSuccess(v bool) *UpgradeFabricChaincodeDefinitionResponseBody {
	s.Success = &v
	return s
}

type UpgradeFabricChaincodeDefinitionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeFabricChaincodeDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeFabricChaincodeDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFabricChaincodeDefinitionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeFabricChaincodeDefinitionResponse) SetHeaders(v map[string]*string) *UpgradeFabricChaincodeDefinitionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeFabricChaincodeDefinitionResponse) SetStatusCode(v int32) *UpgradeFabricChaincodeDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeFabricChaincodeDefinitionResponse) SetBody(v *UpgradeFabricChaincodeDefinitionResponseBody) *UpgradeFabricChaincodeDefinitionResponse {
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

// @param request - AcceptFabricInvitationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptFabricInvitationResponse
func (client *Client) AcceptFabricInvitationWithOptions(request *AcceptFabricInvitationRequest, runtime *util.RuntimeOptions) (_result *AcceptFabricInvitationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.IsAccepted)) {
		body["IsAccepted"] = request.IsAccepted
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AcceptFabricInvitation"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AcceptFabricInvitationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AcceptFabricInvitationRequest
//
// @return AcceptFabricInvitationResponse
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

// @param request - ApplyAntChainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyAntChainCertificateResponse
func (client *Client) ApplyAntChainCertificateWithOptions(request *ApplyAntChainCertificateRequest, runtime *util.RuntimeOptions) (_result *ApplyAntChainCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.UploadReq)) {
		body["UploadReq"] = request.UploadReq
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyAntChainCertificate"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyAntChainCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ApplyAntChainCertificateRequest
//
// @return ApplyAntChainCertificateResponse
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

// @param request - ApplyAntChainCertificateWithKeyAutoCreationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyAntChainCertificateWithKeyAutoCreationResponse
func (client *Client) ApplyAntChainCertificateWithKeyAutoCreationWithOptions(request *ApplyAntChainCertificateWithKeyAutoCreationRequest, runtime *util.RuntimeOptions) (_result *ApplyAntChainCertificateWithKeyAutoCreationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.CommonName)) {
		body["CommonName"] = request.CommonName
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.CountryName)) {
		body["CountryName"] = request.CountryName
	}

	if !tea.BoolValue(util.IsUnset(request.LocalityName)) {
		body["LocalityName"] = request.LocalityName
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationName)) {
		body["OrganizationName"] = request.OrganizationName
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationUnitName)) {
		body["OrganizationUnitName"] = request.OrganizationUnitName
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.StateOrProvinceName)) {
		body["StateOrProvinceName"] = request.StateOrProvinceName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyAntChainCertificateWithKeyAutoCreation"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyAntChainCertificateWithKeyAutoCreationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ApplyAntChainCertificateWithKeyAutoCreationRequest
//
// @return ApplyAntChainCertificateWithKeyAutoCreationResponse
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

// Summary:
//
// ApproveFabricChaincodeDefinition
//
// @param request - ApproveFabricChaincodeDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApproveFabricChaincodeDefinitionResponse
func (client *Client) ApproveFabricChaincodeDefinitionWithOptions(request *ApproveFabricChaincodeDefinitionRequest, runtime *util.RuntimeOptions) (_result *ApproveFabricChaincodeDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChaincodeId)) {
		body["ChaincodeId"] = request.ChaincodeId
	}

	if !tea.BoolValue(util.IsUnset(request.ChaincodePackageId)) {
		body["ChaincodePackageId"] = request.ChaincodePackageId
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApproveFabricChaincodeDefinition"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApproveFabricChaincodeDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ApproveFabricChaincodeDefinition
//
// @param request - ApproveFabricChaincodeDefinitionRequest
//
// @return ApproveFabricChaincodeDefinitionResponse
func (client *Client) ApproveFabricChaincodeDefinition(request *ApproveFabricChaincodeDefinitionRequest) (_result *ApproveFabricChaincodeDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApproveFabricChaincodeDefinitionResponse{}
	_body, _err := client.ApproveFabricChaincodeDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - BatchAddAntChainMiniAppQRCodeAuthorizedUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneListShrink)) {
		body["PhoneList"] = request.PhoneListShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchAddAntChainMiniAppQRCodeAuthorizedUsers"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - BatchAddAntChainMiniAppQRCodeAuthorizedUsersRequest
//
// @return BatchAddAntChainMiniAppQRCodeAuthorizedUsersResponse
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

// @param request - CheckFabricConsortiumDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckFabricConsortiumDomainResponse
func (client *Client) CheckFabricConsortiumDomainWithOptions(request *CheckFabricConsortiumDomainRequest, runtime *util.RuntimeOptions) (_result *CheckFabricConsortiumDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainCode)) {
		body["DomainCode"] = request.DomainCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckFabricConsortiumDomain"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckFabricConsortiumDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckFabricConsortiumDomainRequest
//
// @return CheckFabricConsortiumDomainResponse
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

// @param request - CheckFabricOrganizationDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckFabricOrganizationDomainResponse
func (client *Client) CheckFabricOrganizationDomainWithOptions(request *CheckFabricOrganizationDomainRequest, runtime *util.RuntimeOptions) (_result *CheckFabricOrganizationDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.DomainCode)) {
		body["DomainCode"] = request.DomainCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckFabricOrganizationDomain"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckFabricOrganizationDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckFabricOrganizationDomainRequest
//
// @return CheckFabricOrganizationDomainResponse
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

// @param request - ConfirmFabricConsortiumMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmFabricConsortiumMemberResponse
func (client *Client) ConfirmFabricConsortiumMemberWithOptions(request *ConfirmFabricConsortiumMemberRequest, runtime *util.RuntimeOptions) (_result *ConfirmFabricConsortiumMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		query["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.Organization)) {
		query["Organization"] = request.Organization
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfirmFabricConsortiumMember"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfirmFabricConsortiumMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ConfirmFabricConsortiumMemberRequest
//
// @return ConfirmFabricConsortiumMemberResponse
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

// @param request - CopyAntChainContractProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyAntChainContractProjectResponse
func (client *Client) CopyAntChainContractProjectWithOptions(request *CopyAntChainContractProjectRequest, runtime *util.RuntimeOptions) (_result *CopyAntChainContractProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectDescription)) {
		body["ProjectDescription"] = request.ProjectDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectVersion)) {
		body["ProjectVersion"] = request.ProjectVersion
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CopyAntChainContractProject"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyAntChainContractProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CopyAntChainContractProjectRequest
//
// @return CopyAntChainContractProjectResponse
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

// @param request - CreateAntChainAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAntChainAccountResponse
func (client *Client) CreateAntChainAccountWithOptions(request *CreateAntChainAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAntChainAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		body["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPubKey)) {
		body["AccountPubKey"] = request.AccountPubKey
	}

	if !tea.BoolValue(util.IsUnset(request.AccountRecoverPubKey)) {
		body["AccountRecoverPubKey"] = request.AccountRecoverPubKey
	}

	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAntChainAccount"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAntChainAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAntChainAccountRequest
//
// @return CreateAntChainAccountResponse
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

// @param request - CreateAntChainAccountWithKeyPairAutoCreationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAntChainAccountWithKeyPairAutoCreationResponse
func (client *Client) CreateAntChainAccountWithKeyPairAutoCreationWithOptions(request *CreateAntChainAccountWithKeyPairAutoCreationRequest, runtime *util.RuntimeOptions) (_result *CreateAntChainAccountWithKeyPairAutoCreationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		body["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RecoverPassword)) {
		body["RecoverPassword"] = request.RecoverPassword
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAntChainAccountWithKeyPairAutoCreation"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAntChainAccountWithKeyPairAutoCreationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAntChainAccountWithKeyPairAutoCreationRequest
//
// @return CreateAntChainAccountWithKeyPairAutoCreationResponse
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

// @param request - CreateAntChainConsortiumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAntChainConsortiumResponse
func (client *Client) CreateAntChainConsortiumWithOptions(request *CreateAntChainConsortiumRequest, runtime *util.RuntimeOptions) (_result *CreateAntChainConsortiumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumDescription)) {
		body["ConsortiumDescription"] = request.ConsortiumDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumName)) {
		body["ConsortiumName"] = request.ConsortiumName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAntChainConsortium"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAntChainConsortiumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAntChainConsortiumRequest
//
// @return CreateAntChainConsortiumResponse
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

// @param request - CreateAntChainContractContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAntChainContractContentResponse
func (client *Client) CreateAntChainContractContentWithOptions(request *CreateAntChainContractContentRequest, runtime *util.RuntimeOptions) (_result *CreateAntChainContractContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentName)) {
		body["ContentName"] = request.ContentName
	}

	if !tea.BoolValue(util.IsUnset(request.IsDirectory)) {
		body["IsDirectory"] = request.IsDirectory
	}

	if !tea.BoolValue(util.IsUnset(request.ParentContentId)) {
		body["ParentContentId"] = request.ParentContentId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAntChainContractContent"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAntChainContractContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAntChainContractContentRequest
//
// @return CreateAntChainContractContentResponse
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

// @param request - CreateAntChainContractProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAntChainContractProjectResponse
func (client *Client) CreateAntChainContractProjectWithOptions(request *CreateAntChainContractProjectRequest, runtime *util.RuntimeOptions) (_result *CreateAntChainContractProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectDescription)) {
		body["ProjectDescription"] = request.ProjectDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectVersion)) {
		body["ProjectVersion"] = request.ProjectVersion
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAntChainContractProject"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAntChainContractProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAntChainContractProjectRequest
//
// @return CreateAntChainContractProjectResponse
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

// Summary:
//
// 创建MyChain托管账户
//
// @param request - CreateAntChainKmsAccountNewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAntChainKmsAccountNewResponse
func (client *Client) CreateAntChainKmsAccountNewWithOptions(request *CreateAntChainKmsAccountNewRequest, runtime *util.RuntimeOptions) (_result *CreateAntChainKmsAccountNewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		body["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAntChainKmsAccountNew"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAntChainKmsAccountNewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建MyChain托管账户
//
// @param request - CreateAntChainKmsAccountNewRequest
//
// @return CreateAntChainKmsAccountNewResponse
func (client *Client) CreateAntChainKmsAccountNew(request *CreateAntChainKmsAccountNewRequest) (_result *CreateAntChainKmsAccountNewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAntChainKmsAccountNewResponse{}
	_body, _err := client.CreateAntChainKmsAccountNewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateFabricChaincodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFabricChaincodeResponse
func (client *Client) CreateFabricChaincodeWithOptions(request *CreateFabricChaincodeRequest, runtime *util.RuntimeOptions) (_result *CreateFabricChaincodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		body["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.EndorsePolicy)) {
		body["EndorsePolicy"] = request.EndorsePolicy
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucket)) {
		body["OssBucket"] = request.OssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.OssUrl)) {
		body["OssUrl"] = request.OssUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFabricChaincode"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFabricChaincodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateFabricChaincodeRequest
//
// @return CreateFabricChaincodeResponse
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

// Summary:
//
// CreateFabricChaincodePackage
//
// @param request - CreateFabricChaincodePackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFabricChaincodePackageResponse
func (client *Client) CreateFabricChaincodePackageWithOptions(request *CreateFabricChaincodePackageRequest, runtime *util.RuntimeOptions) (_result *CreateFabricChaincodePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.OssUrl)) {
		body["OssUrl"] = request.OssUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFabricChaincodePackage"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFabricChaincodePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// CreateFabricChaincodePackage
//
// @param request - CreateFabricChaincodePackageRequest
//
// @return CreateFabricChaincodePackageResponse
func (client *Client) CreateFabricChaincodePackage(request *CreateFabricChaincodePackageRequest) (_result *CreateFabricChaincodePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFabricChaincodePackageResponse{}
	_body, _err := client.CreateFabricChaincodePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateFabricChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFabricChannelResponse
func (client *Client) CreateFabricChannelWithOptions(request *CreateFabricChannelRequest, runtime *util.RuntimeOptions) (_result *CreateFabricChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelName)) {
		query["ChannelName"] = request.ChannelName
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		query["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.Organization)) {
		query["Organization"] = request.Organization
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BatchTimeout)) {
		body["BatchTimeout"] = request.BatchTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.MaxMessageCount)) {
		body["MaxMessageCount"] = request.MaxMessageCount
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredMaxBytes)) {
		body["PreferredMaxBytes"] = request.PreferredMaxBytes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFabricChannel"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFabricChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateFabricChannelRequest
//
// @return CreateFabricChannelResponse
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

// @param request - CreateFabricChannelMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFabricChannelMemberResponse
func (client *Client) CreateFabricChannelMemberWithOptions(request *CreateFabricChannelMemberRequest, runtime *util.RuntimeOptions) (_result *CreateFabricChannelMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.Organization)) {
		query["Organization"] = request.Organization
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFabricChannelMember"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFabricChannelMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateFabricChannelMemberRequest
//
// @return CreateFabricChannelMemberResponse
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

// @param request - CreateFabricConsortiumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFabricConsortiumResponse
func (client *Client) CreateFabricConsortiumWithOptions(request *CreateFabricConsortiumRequest, runtime *util.RuntimeOptions) (_result *CreateFabricConsortiumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelPolicy)) {
		body["ChannelPolicy"] = request.ChannelPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumDescription)) {
		body["ConsortiumDescription"] = request.ConsortiumDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumName)) {
		body["ConsortiumName"] = request.ConsortiumName
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OrdererType)) {
		body["OrdererType"] = request.OrdererType
	}

	if !tea.BoolValue(util.IsUnset(request.OrderersCount)) {
		body["OrderersCount"] = request.OrderersCount
	}

	if !tea.BoolValue(util.IsUnset(request.Organization)) {
		body["Organization"] = request.Organization
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentDuration)) {
		body["PaymentDuration"] = request.PaymentDuration
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentDurationUnit)) {
		body["PaymentDurationUnit"] = request.PaymentDurationUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PeersCount)) {
		body["PeersCount"] = request.PeersCount
	}

	if !tea.BoolValue(util.IsUnset(request.SpecName)) {
		body["SpecName"] = request.SpecName
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFabricConsortium"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFabricConsortiumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateFabricConsortiumRequest
//
// @return CreateFabricConsortiumResponse
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

// @param request - CreateFabricConsortiumMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFabricConsortiumMemberResponse
func (client *Client) CreateFabricConsortiumMemberWithOptions(request *CreateFabricConsortiumMemberRequest, runtime *util.RuntimeOptions) (_result *CreateFabricConsortiumMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		query["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.Organization)) {
		query["Organization"] = request.Organization
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFabricConsortiumMember"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFabricConsortiumMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateFabricConsortiumMemberRequest
//
// @return CreateFabricConsortiumMemberResponse
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

// @param request - CreateFabricOrganizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFabricOrganizationResponse
func (client *Client) CreateFabricOrganizationWithOptions(request *CreateFabricOrganizationRequest, runtime *util.RuntimeOptions) (_result *CreateFabricOrganizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		query["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationName)) {
		query["OrganizationName"] = request.OrganizationName
	}

	if !tea.BoolValue(util.IsUnset(request.SpecName)) {
		query["SpecName"] = request.SpecName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PaymentDuration)) {
		body["PaymentDuration"] = request.PaymentDuration
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentDurationUnit)) {
		body["PaymentDurationUnit"] = request.PaymentDurationUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PeersCount)) {
		body["PeersCount"] = request.PeersCount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFabricOrganization"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFabricOrganizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateFabricOrganizationRequest
//
// @return CreateFabricOrganizationResponse
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

// @param request - CreateFabricOrganizationUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFabricOrganizationUserResponse
func (client *Client) CreateFabricOrganizationUserWithOptions(request *CreateFabricOrganizationUserRequest, runtime *util.RuntimeOptions) (_result *CreateFabricOrganizationUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Attrs)) {
		body["Attrs"] = request.Attrs
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFabricOrganizationUser"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFabricOrganizationUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateFabricOrganizationUserRequest
//
// @return CreateFabricOrganizationUserResponse
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

// @param request - DeleteAntChainConsortiumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAntChainConsortiumResponse
func (client *Client) DeleteAntChainConsortiumWithOptions(request *DeleteAntChainConsortiumRequest, runtime *util.RuntimeOptions) (_result *DeleteAntChainConsortiumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAntChainConsortium"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAntChainConsortiumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAntChainConsortiumRequest
//
// @return DeleteAntChainConsortiumResponse
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

// @param request - DeleteAntChainContractContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAntChainContractContentResponse
func (client *Client) DeleteAntChainContractContentWithOptions(request *DeleteAntChainContractContentRequest, runtime *util.RuntimeOptions) (_result *DeleteAntChainContractContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentId)) {
		body["ContentId"] = request.ContentId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAntChainContractContent"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAntChainContractContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAntChainContractContentRequest
//
// @return DeleteAntChainContractContentResponse
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

// @param request - DeleteAntChainContractProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAntChainContractProjectResponse
func (client *Client) DeleteAntChainContractProjectWithOptions(request *DeleteAntChainContractProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteAntChainContractProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAntChainContractProject"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAntChainContractProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAntChainContractProjectRequest
//
// @return DeleteAntChainContractProjectResponse
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

// @param request - DeleteAntChainMiniAppQRCodeAuthorizedUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAntChainMiniAppQRCodeAuthorizedUserResponse
func (client *Client) DeleteAntChainMiniAppQRCodeAuthorizedUserWithOptions(request *DeleteAntChainMiniAppQRCodeAuthorizedUserRequest, runtime *util.RuntimeOptions) (_result *DeleteAntChainMiniAppQRCodeAuthorizedUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["Phone"] = request.Phone
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAntChainMiniAppQRCodeAuthorizedUser"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAntChainMiniAppQRCodeAuthorizedUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAntChainMiniAppQRCodeAuthorizedUserRequest
//
// @return DeleteAntChainMiniAppQRCodeAuthorizedUserResponse
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

// @param request - DeleteFabricChaincodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFabricChaincodeResponse
func (client *Client) DeleteFabricChaincodeWithOptions(request *DeleteFabricChaincodeRequest, runtime *util.RuntimeOptions) (_result *DeleteFabricChaincodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChaincodeId)) {
		body["ChaincodeId"] = request.ChaincodeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFabricChaincode"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFabricChaincodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteFabricChaincodeRequest
//
// @return DeleteFabricChaincodeResponse
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

// @param request - DescribeAntChainAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainAccountsResponse
func (client *Client) DescribeAntChainAccountsWithOptions(request *DescribeAntChainAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainAccounts"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainAccountsRequest
//
// @return DescribeAntChainAccountsResponse
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

// @param request - DescribeAntChainAccountsV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainAccountsV2Response
func (client *Client) DescribeAntChainAccountsV2WithOptions(request *DescribeAntChainAccountsV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainAccountsV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainAccountsV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainAccountsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainAccountsV2Request
//
// @return DescribeAntChainAccountsV2Response
func (client *Client) DescribeAntChainAccountsV2(request *DescribeAntChainAccountsV2Request) (_result *DescribeAntChainAccountsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainAccountsV2Response{}
	_body, _err := client.DescribeAntChainAccountsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainBlockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainBlockResponse
func (client *Client) DescribeAntChainBlockWithOptions(request *DescribeAntChainBlockRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainBlockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.Height)) {
		body["Height"] = request.Height
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainBlock"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainBlockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainBlockRequest
//
// @return DescribeAntChainBlockResponse
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

// @param request - DescribeAntChainBlockV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainBlockV2Response
func (client *Client) DescribeAntChainBlockV2WithOptions(request *DescribeAntChainBlockV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainBlockV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.Height)) {
		body["Height"] = request.Height
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainBlockV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainBlockV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainBlockV2Request
//
// @return DescribeAntChainBlockV2Response
func (client *Client) DescribeAntChainBlockV2(request *DescribeAntChainBlockV2Request) (_result *DescribeAntChainBlockV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainBlockV2Response{}
	_body, _err := client.DescribeAntChainBlockV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainCertificateApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainCertificateApplicationsResponse
func (client *Client) DescribeAntChainCertificateApplicationsWithOptions(request *DescribeAntChainCertificateApplicationsRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainCertificateApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainCertificateApplications"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainCertificateApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainCertificateApplicationsRequest
//
// @return DescribeAntChainCertificateApplicationsResponse
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

// @param request - DescribeAntChainCertificateApplicationsV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainCertificateApplicationsV2Response
func (client *Client) DescribeAntChainCertificateApplicationsV2WithOptions(request *DescribeAntChainCertificateApplicationsV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainCertificateApplicationsV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainCertificateApplicationsV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainCertificateApplicationsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainCertificateApplicationsV2Request
//
// @return DescribeAntChainCertificateApplicationsV2Response
func (client *Client) DescribeAntChainCertificateApplicationsV2(request *DescribeAntChainCertificateApplicationsV2Request) (_result *DescribeAntChainCertificateApplicationsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainCertificateApplicationsV2Response{}
	_body, _err := client.DescribeAntChainCertificateApplicationsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainConsortiumsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainConsortiumsResponse
func (client *Client) DescribeAntChainConsortiumsWithOptions(request *DescribeAntChainConsortiumsRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainConsortiumsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainConsortiums"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainConsortiumsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainConsortiumsRequest
//
// @return DescribeAntChainConsortiumsResponse
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

// @param request - DescribeAntChainConsortiumsV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainConsortiumsV2Response
func (client *Client) DescribeAntChainConsortiumsV2WithOptions(request *DescribeAntChainConsortiumsV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainConsortiumsV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainConsortiumsV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainConsortiumsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainConsortiumsV2Request
//
// @return DescribeAntChainConsortiumsV2Response
func (client *Client) DescribeAntChainConsortiumsV2(request *DescribeAntChainConsortiumsV2Request) (_result *DescribeAntChainConsortiumsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainConsortiumsV2Response{}
	_body, _err := client.DescribeAntChainConsortiumsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainContractProjectContentTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainContractProjectContentTreeResponse
func (client *Client) DescribeAntChainContractProjectContentTreeWithOptions(request *DescribeAntChainContractProjectContentTreeRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainContractProjectContentTreeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainContractProjectContentTree"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainContractProjectContentTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainContractProjectContentTreeRequest
//
// @return DescribeAntChainContractProjectContentTreeResponse
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

// @param request - DescribeAntChainContractProjectContentTreeV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainContractProjectContentTreeV2Response
func (client *Client) DescribeAntChainContractProjectContentTreeV2WithOptions(request *DescribeAntChainContractProjectContentTreeV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainContractProjectContentTreeV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainContractProjectContentTreeV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainContractProjectContentTreeV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainContractProjectContentTreeV2Request
//
// @return DescribeAntChainContractProjectContentTreeV2Response
func (client *Client) DescribeAntChainContractProjectContentTreeV2(request *DescribeAntChainContractProjectContentTreeV2Request) (_result *DescribeAntChainContractProjectContentTreeV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainContractProjectContentTreeV2Response{}
	_body, _err := client.DescribeAntChainContractProjectContentTreeV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainContractProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainContractProjectsResponse
func (client *Client) DescribeAntChainContractProjectsWithOptions(request *DescribeAntChainContractProjectsRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainContractProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainContractProjects"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainContractProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainContractProjectsRequest
//
// @return DescribeAntChainContractProjectsResponse
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

// @param request - DescribeAntChainContractProjectsV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainContractProjectsV2Response
func (client *Client) DescribeAntChainContractProjectsV2WithOptions(request *DescribeAntChainContractProjectsV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainContractProjectsV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainContractProjectsV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainContractProjectsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainContractProjectsV2Request
//
// @return DescribeAntChainContractProjectsV2Response
func (client *Client) DescribeAntChainContractProjectsV2(request *DescribeAntChainContractProjectsV2Request) (_result *DescribeAntChainContractProjectsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainContractProjectsV2Response{}
	_body, _err := client.DescribeAntChainContractProjectsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainDownloadPathsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainDownloadPathsResponse
func (client *Client) DescribeAntChainDownloadPathsWithOptions(request *DescribeAntChainDownloadPathsRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainDownloadPathsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainDownloadPaths"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainDownloadPathsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainDownloadPathsRequest
//
// @return DescribeAntChainDownloadPathsResponse
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

// @param request - DescribeAntChainDownloadPathsV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainDownloadPathsV2Response
func (client *Client) DescribeAntChainDownloadPathsV2WithOptions(request *DescribeAntChainDownloadPathsV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainDownloadPathsV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainDownloadPathsV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainDownloadPathsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainDownloadPathsV2Request
//
// @return DescribeAntChainDownloadPathsV2Response
func (client *Client) DescribeAntChainDownloadPathsV2(request *DescribeAntChainDownloadPathsV2Request) (_result *DescribeAntChainDownloadPathsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainDownloadPathsV2Response{}
	_body, _err := client.DescribeAntChainDownloadPathsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainInformationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainInformationResponse
func (client *Client) DescribeAntChainInformationWithOptions(request *DescribeAntChainInformationRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainInformation"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainInformationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainInformationRequest
//
// @return DescribeAntChainInformationResponse
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

// @param request - DescribeAntChainInformationV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainInformationV2Response
func (client *Client) DescribeAntChainInformationV2WithOptions(request *DescribeAntChainInformationV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainInformationV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainInformationV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainInformationV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainInformationV2Request
//
// @return DescribeAntChainInformationV2Response
func (client *Client) DescribeAntChainInformationV2(request *DescribeAntChainInformationV2Request) (_result *DescribeAntChainInformationV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainInformationV2Response{}
	_body, _err := client.DescribeAntChainInformationV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainLatestBlocksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainLatestBlocksResponse
func (client *Client) DescribeAntChainLatestBlocksWithOptions(request *DescribeAntChainLatestBlocksRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainLatestBlocksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainLatestBlocks"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainLatestBlocksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainLatestBlocksRequest
//
// @return DescribeAntChainLatestBlocksResponse
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

// @param request - DescribeAntChainLatestBlocksV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainLatestBlocksV2Response
func (client *Client) DescribeAntChainLatestBlocksV2WithOptions(request *DescribeAntChainLatestBlocksV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainLatestBlocksV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainLatestBlocksV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainLatestBlocksV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainLatestBlocksV2Request
//
// @return DescribeAntChainLatestBlocksV2Response
func (client *Client) DescribeAntChainLatestBlocksV2(request *DescribeAntChainLatestBlocksV2Request) (_result *DescribeAntChainLatestBlocksV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainLatestBlocksV2Response{}
	_body, _err := client.DescribeAntChainLatestBlocksV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainLatestTransactionDigestsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainLatestTransactionDigestsResponse
func (client *Client) DescribeAntChainLatestTransactionDigestsWithOptions(request *DescribeAntChainLatestTransactionDigestsRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainLatestTransactionDigestsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainLatestTransactionDigests"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainLatestTransactionDigestsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainLatestTransactionDigestsRequest
//
// @return DescribeAntChainLatestTransactionDigestsResponse
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

// @param request - DescribeAntChainLatestTransactionDigestsV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainLatestTransactionDigestsV2Response
func (client *Client) DescribeAntChainLatestTransactionDigestsV2WithOptions(request *DescribeAntChainLatestTransactionDigestsV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainLatestTransactionDigestsV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainLatestTransactionDigestsV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainLatestTransactionDigestsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainLatestTransactionDigestsV2Request
//
// @return DescribeAntChainLatestTransactionDigestsV2Response
func (client *Client) DescribeAntChainLatestTransactionDigestsV2(request *DescribeAntChainLatestTransactionDigestsV2Request) (_result *DescribeAntChainLatestTransactionDigestsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainLatestTransactionDigestsV2Response{}
	_body, _err := client.DescribeAntChainLatestTransactionDigestsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainMembersResponse
func (client *Client) DescribeAntChainMembersWithOptions(request *DescribeAntChainMembersRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainMembers"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainMembersRequest
//
// @return DescribeAntChainMembersResponse
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

// @param request - DescribeAntChainMembersV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainMembersV2Response
func (client *Client) DescribeAntChainMembersV2WithOptions(request *DescribeAntChainMembersV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainMembersV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainMembersV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainMembersV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainMembersV2Request
//
// @return DescribeAntChainMembersV2Response
func (client *Client) DescribeAntChainMembersV2(request *DescribeAntChainMembersV2Request) (_result *DescribeAntChainMembersV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainMembersV2Response{}
	_body, _err := client.DescribeAntChainMembersV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainMiniAppBrowserQRCodeAccessLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse
func (client *Client) DescribeAntChainMiniAppBrowserQRCodeAccessLogWithOptions(request *DescribeAntChainMiniAppBrowserQRCodeAccessLogRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.QRCodeType)) {
		body["QRCodeType"] = request.QRCodeType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainMiniAppBrowserQRCodeAccessLog"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainMiniAppBrowserQRCodeAccessLogRequest
//
// @return DescribeAntChainMiniAppBrowserQRCodeAccessLogResponse
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

// @param request - DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Response
func (client *Client) DescribeAntChainMiniAppBrowserQRCodeAccessLogV2WithOptions(request *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.QRCodeType)) {
		body["QRCodeType"] = request.QRCodeType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainMiniAppBrowserQRCodeAccessLogV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Request
//
// @return DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Response
func (client *Client) DescribeAntChainMiniAppBrowserQRCodeAccessLogV2(request *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Request) (_result *DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainMiniAppBrowserQRCodeAccessLogV2Response{}
	_body, _err := client.DescribeAntChainMiniAppBrowserQRCodeAccessLogV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse
func (client *Client) DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersWithOptions(request *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QRCodeType)) {
		body["QRCodeType"] = request.QRCodeType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsers"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersRequest
//
// @return DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersResponse
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

// @param request - DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Response
func (client *Client) DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2WithOptions(request *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QRCodeType)) {
		body["QRCodeType"] = request.QRCodeType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Request
//
// @return DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Response
func (client *Client) DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2(request *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Request) (_result *DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2Response{}
	_body, _err := client.DescribeAntChainMiniAppBrowserQRCodeAuthorizedUsersV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainMiniAppBrowserTransactionQRCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainMiniAppBrowserTransactionQRCodeResponse
func (client *Client) DescribeAntChainMiniAppBrowserTransactionQRCodeWithOptions(request *DescribeAntChainMiniAppBrowserTransactionQRCodeRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainMiniAppBrowserTransactionQRCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionHash)) {
		body["TransactionHash"] = request.TransactionHash
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainMiniAppBrowserTransactionQRCode"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainMiniAppBrowserTransactionQRCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainMiniAppBrowserTransactionQRCodeRequest
//
// @return DescribeAntChainMiniAppBrowserTransactionQRCodeResponse
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

// @param request - DescribeAntChainMiniAppBrowserTransactionQRCodeNewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponse
func (client *Client) DescribeAntChainMiniAppBrowserTransactionQRCodeNewWithOptions(request *DescribeAntChainMiniAppBrowserTransactionQRCodeNewRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.ContractId)) {
		body["ContractId"] = request.ContractId
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionHash)) {
		body["TransactionHash"] = request.TransactionHash
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainMiniAppBrowserTransactionQRCodeNew"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainMiniAppBrowserTransactionQRCodeNewRequest
//
// @return DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponse
func (client *Client) DescribeAntChainMiniAppBrowserTransactionQRCodeNew(request *DescribeAntChainMiniAppBrowserTransactionQRCodeNewRequest) (_result *DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainMiniAppBrowserTransactionQRCodeNewResponse{}
	_body, _err := client.DescribeAntChainMiniAppBrowserTransactionQRCodeNewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainNodesResponse
func (client *Client) DescribeAntChainNodesWithOptions(request *DescribeAntChainNodesRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainNodes"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainNodesRequest
//
// @return DescribeAntChainNodesResponse
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

// @param request - DescribeAntChainNodesV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainNodesV2Response
func (client *Client) DescribeAntChainNodesV2WithOptions(request *DescribeAntChainNodesV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainNodesV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainNodesV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainNodesV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainNodesV2Request
//
// @return DescribeAntChainNodesV2Response
func (client *Client) DescribeAntChainNodesV2(request *DescribeAntChainNodesV2Request) (_result *DescribeAntChainNodesV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainNodesV2Response{}
	_body, _err := client.DescribeAntChainNodesV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainQRCodeAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainQRCodeAuthorizationResponse
func (client *Client) DescribeAntChainQRCodeAuthorizationWithOptions(request *DescribeAntChainQRCodeAuthorizationRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainQRCodeAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.QRCodeType)) {
		body["QRCodeType"] = request.QRCodeType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainQRCodeAuthorization"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainQRCodeAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainQRCodeAuthorizationRequest
//
// @return DescribeAntChainQRCodeAuthorizationResponse
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

// @param request - DescribeAntChainQRCodeAuthorizationV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainQRCodeAuthorizationV2Response
func (client *Client) DescribeAntChainQRCodeAuthorizationV2WithOptions(request *DescribeAntChainQRCodeAuthorizationV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainQRCodeAuthorizationV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.QRCodeType)) {
		body["QRCodeType"] = request.QRCodeType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainQRCodeAuthorizationV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainQRCodeAuthorizationV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainQRCodeAuthorizationV2Request
//
// @return DescribeAntChainQRCodeAuthorizationV2Response
func (client *Client) DescribeAntChainQRCodeAuthorizationV2(request *DescribeAntChainQRCodeAuthorizationV2Request) (_result *DescribeAntChainQRCodeAuthorizationV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainQRCodeAuthorizationV2Response{}
	_body, _err := client.DescribeAntChainQRCodeAuthorizationV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainTransactionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainTransactionResponse
func (client *Client) DescribeAntChainTransactionWithOptions(request *DescribeAntChainTransactionRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainTransactionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.Hash)) {
		body["Hash"] = request.Hash
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainTransaction"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainTransactionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainTransactionRequest
//
// @return DescribeAntChainTransactionResponse
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

// @param request - DescribeAntChainTransactionReceiptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainTransactionReceiptResponse
func (client *Client) DescribeAntChainTransactionReceiptWithOptions(request *DescribeAntChainTransactionReceiptRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainTransactionReceiptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.Hash)) {
		body["Hash"] = request.Hash
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainTransactionReceipt"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainTransactionReceiptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainTransactionReceiptRequest
//
// @return DescribeAntChainTransactionReceiptResponse
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

// @param request - DescribeAntChainTransactionReceiptV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainTransactionReceiptV2Response
func (client *Client) DescribeAntChainTransactionReceiptV2WithOptions(request *DescribeAntChainTransactionReceiptV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainTransactionReceiptV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.Hash)) {
		body["Hash"] = request.Hash
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainTransactionReceiptV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainTransactionReceiptV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainTransactionReceiptV2Request
//
// @return DescribeAntChainTransactionReceiptV2Response
func (client *Client) DescribeAntChainTransactionReceiptV2(request *DescribeAntChainTransactionReceiptV2Request) (_result *DescribeAntChainTransactionReceiptV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainTransactionReceiptV2Response{}
	_body, _err := client.DescribeAntChainTransactionReceiptV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainTransactionStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainTransactionStatisticsResponse
func (client *Client) DescribeAntChainTransactionStatisticsWithOptions(request *DescribeAntChainTransactionStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainTransactionStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		body["End"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		body["Start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainTransactionStatistics"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainTransactionStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainTransactionStatisticsRequest
//
// @return DescribeAntChainTransactionStatisticsResponse
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

// @param request - DescribeAntChainTransactionStatisticsV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainTransactionStatisticsV2Response
func (client *Client) DescribeAntChainTransactionStatisticsV2WithOptions(request *DescribeAntChainTransactionStatisticsV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainTransactionStatisticsV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		body["End"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		body["Start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainTransactionStatisticsV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainTransactionStatisticsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainTransactionStatisticsV2Request
//
// @return DescribeAntChainTransactionStatisticsV2Response
func (client *Client) DescribeAntChainTransactionStatisticsV2(request *DescribeAntChainTransactionStatisticsV2Request) (_result *DescribeAntChainTransactionStatisticsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainTransactionStatisticsV2Response{}
	_body, _err := client.DescribeAntChainTransactionStatisticsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainTransactionV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainTransactionV2Response
func (client *Client) DescribeAntChainTransactionV2WithOptions(request *DescribeAntChainTransactionV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainTransactionV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.Hash)) {
		body["Hash"] = request.Hash
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainTransactionV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainTransactionV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainTransactionV2Request
//
// @return DescribeAntChainTransactionV2Response
func (client *Client) DescribeAntChainTransactionV2(request *DescribeAntChainTransactionV2Request) (_result *DescribeAntChainTransactionV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainTransactionV2Response{}
	_body, _err := client.DescribeAntChainTransactionV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAntChainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainsResponse
func (client *Client) DescribeAntChainsWithOptions(request *DescribeAntChainsRequest, runtime *util.RuntimeOptions) (_result *DescribeAntChainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChains"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainsRequest
//
// @return DescribeAntChainsResponse
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

// @param request - DescribeAntChainsV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAntChainsV2Response
func (client *Client) DescribeAntChainsV2WithOptions(request *DescribeAntChainsV2Request, runtime *util.RuntimeOptions) (_result *DescribeAntChainsV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAntChainsV2"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAntChainsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAntChainsV2Request
//
// @return DescribeAntChainsV2Response
func (client *Client) DescribeAntChainsV2(request *DescribeAntChainsV2Request) (_result *DescribeAntChainsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntChainsV2Response{}
	_body, _err := client.DescribeAntChainsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeEthereumDeletableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEthereumDeletableResponse
func (client *Client) DescribeEthereumDeletableWithOptions(request *DescribeEthereumDeletableRequest, runtime *util.RuntimeOptions) (_result *DescribeEthereumDeletableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EthereumId)) {
		body["EthereumId"] = request.EthereumId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEthereumDeletable"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEthereumDeletableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeEthereumDeletableRequest
//
// @return DescribeEthereumDeletableResponse
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

// @param request - DescribeFabricCandidateOrganizationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricCandidateOrganizationsResponse
func (client *Client) DescribeFabricCandidateOrganizationsWithOptions(request *DescribeFabricCandidateOrganizationsRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricCandidateOrganizationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricCandidateOrganizations"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricCandidateOrganizationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricCandidateOrganizationsRequest
//
// @return DescribeFabricCandidateOrganizationsResponse
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

// Summary:
//
// DescribeFabricChaincodeDefinitionTask
//
// @param request - DescribeFabricChaincodeDefinitionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricChaincodeDefinitionTaskResponse
func (client *Client) DescribeFabricChaincodeDefinitionTaskWithOptions(request *DescribeFabricChaincodeDefinitionTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricChaincodeDefinitionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChaincodeId)) {
		body["ChaincodeId"] = request.ChaincodeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricChaincodeDefinitionTask"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricChaincodeDefinitionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// DescribeFabricChaincodeDefinitionTask
//
// @param request - DescribeFabricChaincodeDefinitionTaskRequest
//
// @return DescribeFabricChaincodeDefinitionTaskResponse
func (client *Client) DescribeFabricChaincodeDefinitionTask(request *DescribeFabricChaincodeDefinitionTaskRequest) (_result *DescribeFabricChaincodeDefinitionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricChaincodeDefinitionTaskResponse{}
	_body, _err := client.DescribeFabricChaincodeDefinitionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeFabricChaincodeUploadPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricChaincodeUploadPolicyResponse
func (client *Client) DescribeFabricChaincodeUploadPolicyWithOptions(request *DescribeFabricChaincodeUploadPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricChaincodeUploadPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricChaincodeUploadPolicy"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricChaincodeUploadPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricChaincodeUploadPolicyRequest
//
// @return DescribeFabricChaincodeUploadPolicyResponse
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

// @param request - DescribeFabricChannelMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricChannelMembersResponse
func (client *Client) DescribeFabricChannelMembersWithOptions(request *DescribeFabricChannelMembersRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricChannelMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricChannelMembers"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricChannelMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricChannelMembersRequest
//
// @return DescribeFabricChannelMembersResponse
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

// @param request - DescribeFabricConsortiumAdminStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricConsortiumAdminStatusResponse
func (client *Client) DescribeFabricConsortiumAdminStatusWithOptions(request *DescribeFabricConsortiumAdminStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumAdminStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricConsortiumAdminStatus"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricConsortiumAdminStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricConsortiumAdminStatusRequest
//
// @return DescribeFabricConsortiumAdminStatusResponse
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

// @param request - DescribeFabricConsortiumChaincodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricConsortiumChaincodesResponse
func (client *Client) DescribeFabricConsortiumChaincodesWithOptions(request *DescribeFabricConsortiumChaincodesRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumChaincodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricConsortiumChaincodes"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricConsortiumChaincodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricConsortiumChaincodesRequest
//
// @return DescribeFabricConsortiumChaincodesResponse
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

// @param request - DescribeFabricConsortiumChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricConsortiumChannelsResponse
func (client *Client) DescribeFabricConsortiumChannelsWithOptions(request *DescribeFabricConsortiumChannelsRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumChannelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		query["ConsortiumId"] = request.ConsortiumId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricConsortiumChannels"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricConsortiumChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricConsortiumChannelsRequest
//
// @return DescribeFabricConsortiumChannelsResponse
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

// @param request - DescribeFabricConsortiumConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricConsortiumConfigResponse
func (client *Client) DescribeFabricConsortiumConfigWithOptions(runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricConsortiumConfig"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricConsortiumConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return DescribeFabricConsortiumConfigResponse
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

// @param request - DescribeFabricConsortiumDeletableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricConsortiumDeletableResponse
func (client *Client) DescribeFabricConsortiumDeletableWithOptions(request *DescribeFabricConsortiumDeletableRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumDeletableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		query["ConsortiumId"] = request.ConsortiumId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricConsortiumDeletable"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricConsortiumDeletableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricConsortiumDeletableRequest
//
// @return DescribeFabricConsortiumDeletableResponse
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

// @param request - DescribeFabricConsortiumMemberApprovalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricConsortiumMemberApprovalResponse
func (client *Client) DescribeFabricConsortiumMemberApprovalWithOptions(request *DescribeFabricConsortiumMemberApprovalRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumMemberApprovalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		query["ConsortiumId"] = request.ConsortiumId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricConsortiumMemberApproval"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricConsortiumMemberApprovalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricConsortiumMemberApprovalRequest
//
// @return DescribeFabricConsortiumMemberApprovalResponse
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

// @param request - DescribeFabricConsortiumMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricConsortiumMembersResponse
func (client *Client) DescribeFabricConsortiumMembersWithOptions(request *DescribeFabricConsortiumMembersRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricConsortiumMembers"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricConsortiumMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricConsortiumMembersRequest
//
// @return DescribeFabricConsortiumMembersResponse
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

// @param request - DescribeFabricConsortiumOrderersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricConsortiumOrderersResponse
func (client *Client) DescribeFabricConsortiumOrderersWithOptions(request *DescribeFabricConsortiumOrderersRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumOrderersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricConsortiumOrderers"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricConsortiumOrderersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricConsortiumOrderersRequest
//
// @return DescribeFabricConsortiumOrderersResponse
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

// @param request - DescribeFabricConsortiumSpecsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricConsortiumSpecsResponse
func (client *Client) DescribeFabricConsortiumSpecsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumSpecsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricConsortiumSpecs"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricConsortiumSpecsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return DescribeFabricConsortiumSpecsResponse
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

// @param request - DescribeFabricConsortiumsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricConsortiumsResponse
func (client *Client) DescribeFabricConsortiumsWithOptions(request *DescribeFabricConsortiumsRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricConsortiumsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		query["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricConsortiums"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricConsortiumsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricConsortiumsRequest
//
// @return DescribeFabricConsortiumsResponse
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

// @param request - DescribeFabricExplorerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricExplorerResponse
func (client *Client) DescribeFabricExplorerWithOptions(request *DescribeFabricExplorerRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricExplorerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExBody)) {
		query["ExBody"] = request.ExBody
	}

	if !tea.BoolValue(util.IsUnset(request.ExMethod)) {
		query["ExMethod"] = request.ExMethod
	}

	if !tea.BoolValue(util.IsUnset(request.ExUrl)) {
		query["ExUrl"] = request.ExUrl
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricExplorer"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricExplorerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricExplorerRequest
//
// @return DescribeFabricExplorerResponse
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

// @param request - DescribeFabricInvitationCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricInvitationCodeResponse
func (client *Client) DescribeFabricInvitationCodeWithOptions(request *DescribeFabricInvitationCodeRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricInvitationCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricInvitationCode"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricInvitationCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricInvitationCodeRequest
//
// @return DescribeFabricInvitationCodeResponse
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

// @param request - DescribeFabricInviterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricInviterResponse
func (client *Client) DescribeFabricInviterWithOptions(request *DescribeFabricInviterRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricInviterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["Code"] = request.Code
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricInviter"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricInviterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricInviterRequest
//
// @return DescribeFabricInviterResponse
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

// @param request - DescribeFabricOrdererLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricOrdererLogsResponse
func (client *Client) DescribeFabricOrdererLogsWithOptions(request *DescribeFabricOrdererLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrdererLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		query["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.Lines)) {
		query["Lines"] = request.Lines
	}

	if !tea.BoolValue(util.IsUnset(request.OrdererName)) {
		query["OrdererName"] = request.OrdererName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricOrdererLogs"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricOrdererLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricOrdererLogsRequest
//
// @return DescribeFabricOrdererLogsResponse
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

// @param request - DescribeFabricOrganizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricOrganizationResponse
func (client *Client) DescribeFabricOrganizationWithOptions(request *DescribeFabricOrganizationRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricOrganization"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricOrganizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricOrganizationRequest
//
// @return DescribeFabricOrganizationResponse
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

// Summary:
//
// DescribeFabricOrganizationChaincodePackage
//
// @param request - DescribeFabricOrganizationChaincodePackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricOrganizationChaincodePackageResponse
func (client *Client) DescribeFabricOrganizationChaincodePackageWithOptions(request *DescribeFabricOrganizationChaincodePackageRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationChaincodePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricOrganizationChaincodePackage"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricOrganizationChaincodePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// DescribeFabricOrganizationChaincodePackage
//
// @param request - DescribeFabricOrganizationChaincodePackageRequest
//
// @return DescribeFabricOrganizationChaincodePackageResponse
func (client *Client) DescribeFabricOrganizationChaincodePackage(request *DescribeFabricOrganizationChaincodePackageRequest) (_result *DescribeFabricOrganizationChaincodePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricOrganizationChaincodePackageResponse{}
	_body, _err := client.DescribeFabricOrganizationChaincodePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeFabricOrganizationChaincodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricOrganizationChaincodesResponse
func (client *Client) DescribeFabricOrganizationChaincodesWithOptions(request *DescribeFabricOrganizationChaincodesRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationChaincodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricOrganizationChaincodes"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricOrganizationChaincodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricOrganizationChaincodesRequest
//
// @return DescribeFabricOrganizationChaincodesResponse
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

// @param request - DescribeFabricOrganizationChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricOrganizationChannelsResponse
func (client *Client) DescribeFabricOrganizationChannelsWithOptions(request *DescribeFabricOrganizationChannelsRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationChannelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricOrganizationChannels"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricOrganizationChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricOrganizationChannelsRequest
//
// @return DescribeFabricOrganizationChannelsResponse
func (client *Client) DescribeFabricOrganizationChannels(request *DescribeFabricOrganizationChannelsRequest) (_result *DescribeFabricOrganizationChannelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFabricOrganizationChannelsResponse{}
	_body, _err := client.DescribeFabricOrganizationChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeFabricOrganizationDeletableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricOrganizationDeletableResponse
func (client *Client) DescribeFabricOrganizationDeletableWithOptions(request *DescribeFabricOrganizationDeletableRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationDeletableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricOrganizationDeletable"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricOrganizationDeletableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricOrganizationDeletableRequest
//
// @return DescribeFabricOrganizationDeletableResponse
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

// @param request - DescribeFabricOrganizationMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricOrganizationMembersResponse
func (client *Client) DescribeFabricOrganizationMembersWithOptions(request *DescribeFabricOrganizationMembersRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricOrganizationMembers"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricOrganizationMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricOrganizationMembersRequest
//
// @return DescribeFabricOrganizationMembersResponse
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

// @param request - DescribeFabricOrganizationPeersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricOrganizationPeersResponse
func (client *Client) DescribeFabricOrganizationPeersWithOptions(request *DescribeFabricOrganizationPeersRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationPeersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricOrganizationPeers"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricOrganizationPeersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricOrganizationPeersRequest
//
// @return DescribeFabricOrganizationPeersResponse
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

// @param request - DescribeFabricOrganizationSpecsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricOrganizationSpecsResponse
func (client *Client) DescribeFabricOrganizationSpecsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationSpecsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricOrganizationSpecs"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricOrganizationSpecsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return DescribeFabricOrganizationSpecsResponse
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

// @param request - DescribeFabricOrganizationUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricOrganizationUsersResponse
func (client *Client) DescribeFabricOrganizationUsersWithOptions(request *DescribeFabricOrganizationUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricOrganizationUsers"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricOrganizationUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricOrganizationUsersRequest
//
// @return DescribeFabricOrganizationUsersResponse
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

// @param request - DescribeFabricOrganizationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricOrganizationsResponse
func (client *Client) DescribeFabricOrganizationsWithOptions(request *DescribeFabricOrganizationsRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricOrganizationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricOrganizations"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricOrganizationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricOrganizationsRequest
//
// @return DescribeFabricOrganizationsResponse
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

// @param request - DescribeFabricPeerLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFabricPeerLogsResponse
func (client *Client) DescribeFabricPeerLogsWithOptions(request *DescribeFabricPeerLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeFabricPeerLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lines)) {
		query["Lines"] = request.Lines
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.PeerName)) {
		query["PeerName"] = request.PeerName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFabricPeerLogs"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFabricPeerLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeFabricPeerLogsRequest
//
// @return DescribeFabricPeerLogsResponse
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

// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
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

// @param request - DescribeRootDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRootDomainResponse
func (client *Client) DescribeRootDomainWithOptions(runtime *util.RuntimeOptions) (_result *DescribeRootDomainResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeRootDomain"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRootDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return DescribeRootDomainResponse
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

// @param request - DescribeTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTasksResponse
func (client *Client) DescribeTasksWithOptions(runtime *util.RuntimeOptions) (_result *DescribeTasksResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeTasks"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return DescribeTasksResponse
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

// @param request - DownloadFabricOrganizationSDKRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadFabricOrganizationSDKResponse
func (client *Client) DownloadFabricOrganizationSDKWithOptions(request *DownloadFabricOrganizationSDKRequest, runtime *util.RuntimeOptions) (_result *DownloadFabricOrganizationSDKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DownloadFabricOrganizationSDK"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DownloadFabricOrganizationSDKResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DownloadFabricOrganizationSDKRequest
//
// @return DownloadFabricOrganizationSDKResponse
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

// @param request - FreezeAntChainAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FreezeAntChainAccountResponse
func (client *Client) FreezeAntChainAccountWithOptions(request *FreezeAntChainAccountRequest, runtime *util.RuntimeOptions) (_result *FreezeAntChainAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		body["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FreezeAntChainAccount"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FreezeAntChainAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - FreezeAntChainAccountRequest
//
// @return FreezeAntChainAccountResponse
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

// @param request - InstallFabricChaincodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallFabricChaincodeResponse
func (client *Client) InstallFabricChaincodeWithOptions(request *InstallFabricChaincodeRequest, runtime *util.RuntimeOptions) (_result *InstallFabricChaincodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChaincodeId)) {
		body["ChaincodeId"] = request.ChaincodeId
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallFabricChaincode"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallFabricChaincodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - InstallFabricChaincodeRequest
//
// @return InstallFabricChaincodeResponse
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

// Summary:
//
// InstallFabricChaincodePackage
//
// @param request - InstallFabricChaincodePackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallFabricChaincodePackageResponse
func (client *Client) InstallFabricChaincodePackageWithOptions(request *InstallFabricChaincodePackageRequest, runtime *util.RuntimeOptions) (_result *InstallFabricChaincodePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChaincodePackageId)) {
		body["ChaincodePackageId"] = request.ChaincodePackageId
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallFabricChaincodePackage"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallFabricChaincodePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// InstallFabricChaincodePackage
//
// @param request - InstallFabricChaincodePackageRequest
//
// @return InstallFabricChaincodePackageResponse
func (client *Client) InstallFabricChaincodePackage(request *InstallFabricChaincodePackageRequest) (_result *InstallFabricChaincodePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallFabricChaincodePackageResponse{}
	_body, _err := client.InstallFabricChaincodePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - InstantiateFabricChaincodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstantiateFabricChaincodeResponse
func (client *Client) InstantiateFabricChaincodeWithOptions(request *InstantiateFabricChaincodeRequest, runtime *util.RuntimeOptions) (_result *InstantiateFabricChaincodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChaincodeId)) {
		body["ChaincodeId"] = request.ChaincodeId
	}

	if !tea.BoolValue(util.IsUnset(request.CollectionConfig)) {
		body["CollectionConfig"] = request.CollectionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.EndorsePolicy)) {
		body["EndorsePolicy"] = request.EndorsePolicy
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InstantiateFabricChaincode"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InstantiateFabricChaincodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - InstantiateFabricChaincodeRequest
//
// @return InstantiateFabricChaincodeResponse
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

// @param request - JoinFabricChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return JoinFabricChannelResponse
func (client *Client) JoinFabricChannelWithOptions(request *JoinFabricChannelRequest, runtime *util.RuntimeOptions) (_result *JoinFabricChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.Do)) {
		query["Do"] = request.Do
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("JoinFabricChannel"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &JoinFabricChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - JoinFabricChannelRequest
//
// @return JoinFabricChannelResponse
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

// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
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

// @param request - ResetAntChainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAntChainCertificateResponse
func (client *Client) ResetAntChainCertificateWithOptions(request *ResetAntChainCertificateRequest, runtime *util.RuntimeOptions) (_result *ResetAntChainCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAntChainCertificate"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetAntChainCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResetAntChainCertificateRequest
//
// @return ResetAntChainCertificateResponse
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

// @param request - ResetAntChainUserCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAntChainUserCertificateResponse
func (client *Client) ResetAntChainUserCertificateWithOptions(request *ResetAntChainUserCertificateRequest, runtime *util.RuntimeOptions) (_result *ResetAntChainUserCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAntChainUserCertificate"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetAntChainUserCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResetAntChainUserCertificateRequest
//
// @return ResetAntChainUserCertificateResponse
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

// @param request - ResetFabricOrganizationUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetFabricOrganizationUserPasswordResponse
func (client *Client) ResetFabricOrganizationUserPasswordWithOptions(request *ResetFabricOrganizationUserPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetFabricOrganizationUserPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetFabricOrganizationUserPassword"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetFabricOrganizationUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResetFabricOrganizationUserPasswordRequest
//
// @return ResetFabricOrganizationUserPasswordResponse
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

// Summary:
//
// SubmitFabricChaincodeDefinition
//
// @param request - SubmitFabricChaincodeDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitFabricChaincodeDefinitionResponse
func (client *Client) SubmitFabricChaincodeDefinitionWithOptions(request *SubmitFabricChaincodeDefinitionRequest, runtime *util.RuntimeOptions) (_result *SubmitFabricChaincodeDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChaincodePackageId)) {
		body["ChaincodePackageId"] = request.ChaincodePackageId
	}

	if !tea.BoolValue(util.IsUnset(request.ChaincodeVersion)) {
		body["ChaincodeVersion"] = request.ChaincodeVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		body["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CollectionConfig)) {
		body["CollectionConfig"] = request.CollectionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.EndorsePolicy)) {
		body["EndorsePolicy"] = request.EndorsePolicy
	}

	if !tea.BoolValue(util.IsUnset(request.InitRequired)) {
		body["InitRequired"] = request.InitRequired
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitFabricChaincodeDefinition"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitFabricChaincodeDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// SubmitFabricChaincodeDefinition
//
// @param request - SubmitFabricChaincodeDefinitionRequest
//
// @return SubmitFabricChaincodeDefinitionResponse
func (client *Client) SubmitFabricChaincodeDefinition(request *SubmitFabricChaincodeDefinitionRequest) (_result *SubmitFabricChaincodeDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitFabricChaincodeDefinitionResponse{}
	_body, _err := client.SubmitFabricChaincodeDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SynchronizeFabricChaincodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SynchronizeFabricChaincodeResponse
func (client *Client) SynchronizeFabricChaincodeWithOptions(request *SynchronizeFabricChaincodeRequest, runtime *util.RuntimeOptions) (_result *SynchronizeFabricChaincodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChaincodeId)) {
		body["ChaincodeId"] = request.ChaincodeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SynchronizeFabricChaincode"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SynchronizeFabricChaincodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SynchronizeFabricChaincodeRequest
//
// @return SynchronizeFabricChaincodeResponse
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

// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
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

// @param request - UnfreezeAntChainAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnfreezeAntChainAccountResponse
func (client *Client) UnfreezeAntChainAccountWithOptions(request *UnfreezeAntChainAccountRequest, runtime *util.RuntimeOptions) (_result *UnfreezeAntChainAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		body["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnfreezeAntChainAccount"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnfreezeAntChainAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UnfreezeAntChainAccountRequest
//
// @return UnfreezeAntChainAccountResponse
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

// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
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

// @param request - UpdateAntChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAntChainResponse
func (client *Client) UpdateAntChainWithOptions(request *UpdateAntChainRequest, runtime *util.RuntimeOptions) (_result *UpdateAntChainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.AntChainName)) {
		body["AntChainName"] = request.AntChainName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAntChain"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAntChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAntChainRequest
//
// @return UpdateAntChainResponse
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

// @param request - UpdateAntChainConsortiumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAntChainConsortiumResponse
func (client *Client) UpdateAntChainConsortiumWithOptions(request *UpdateAntChainConsortiumRequest, runtime *util.RuntimeOptions) (_result *UpdateAntChainConsortiumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumDescription)) {
		body["ConsortiumDescription"] = request.ConsortiumDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsortiumName)) {
		body["ConsortiumName"] = request.ConsortiumName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAntChainConsortium"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAntChainConsortiumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAntChainConsortiumRequest
//
// @return UpdateAntChainConsortiumResponse
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

// @param request - UpdateAntChainContractContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAntChainContractContentResponse
func (client *Client) UpdateAntChainContractContentWithOptions(request *UpdateAntChainContractContentRequest, runtime *util.RuntimeOptions) (_result *UpdateAntChainContractContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentId)) {
		body["ContentId"] = request.ContentId
	}

	if !tea.BoolValue(util.IsUnset(request.ContentName)) {
		body["ContentName"] = request.ContentName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentContentId)) {
		body["ParentContentId"] = request.ParentContentId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAntChainContractContent"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAntChainContractContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAntChainContractContentRequest
//
// @return UpdateAntChainContractContentResponse
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

// @param request - UpdateAntChainContractProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAntChainContractProjectResponse
func (client *Client) UpdateAntChainContractProjectWithOptions(request *UpdateAntChainContractProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateAntChainContractProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectDescription)) {
		body["ProjectDescription"] = request.ProjectDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectVersion)) {
		body["ProjectVersion"] = request.ProjectVersion
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAntChainContractProject"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAntChainContractProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAntChainContractProjectRequest
//
// @return UpdateAntChainContractProjectResponse
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

// @param request - UpdateAntChainMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAntChainMemberResponse
func (client *Client) UpdateAntChainMemberWithOptions(request *UpdateAntChainMemberRequest, runtime *util.RuntimeOptions) (_result *UpdateAntChainMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsortiumId)) {
		body["ConsortiumId"] = request.ConsortiumId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		body["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberName)) {
		body["MemberName"] = request.MemberName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAntChainMember"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAntChainMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAntChainMemberRequest
//
// @return UpdateAntChainMemberResponse
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

// @param request - UpdateAntChainQRCodeAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAntChainQRCodeAuthorizationResponse
func (client *Client) UpdateAntChainQRCodeAuthorizationWithOptions(request *UpdateAntChainQRCodeAuthorizationRequest, runtime *util.RuntimeOptions) (_result *UpdateAntChainQRCodeAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AntChainId)) {
		body["AntChainId"] = request.AntChainId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationType)) {
		body["AuthorizationType"] = request.AuthorizationType
	}

	if !tea.BoolValue(util.IsUnset(request.QRCodeType)) {
		body["QRCodeType"] = request.QRCodeType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAntChainQRCodeAuthorization"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAntChainQRCodeAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAntChainQRCodeAuthorizationRequest
//
// @return UpdateAntChainQRCodeAuthorizationResponse
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

// Description:
//
// ***
//
// @param request - UpgradeFabricChaincodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeFabricChaincodeResponse
func (client *Client) UpgradeFabricChaincodeWithOptions(request *UpgradeFabricChaincodeRequest, runtime *util.RuntimeOptions) (_result *UpgradeFabricChaincodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChaincodeId)) {
		body["ChaincodeId"] = request.ChaincodeId
	}

	if !tea.BoolValue(util.IsUnset(request.CollectionConfig)) {
		body["CollectionConfig"] = request.CollectionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.EndorsePolicy)) {
		body["EndorsePolicy"] = request.EndorsePolicy
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeFabricChaincode"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeFabricChaincodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// ***
//
// @param request - UpgradeFabricChaincodeRequest
//
// @return UpgradeFabricChaincodeResponse
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

// Summary:
//
// UpgradeFabricChaincodeDefinition
//
// @param request - UpgradeFabricChaincodeDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeFabricChaincodeDefinitionResponse
func (client *Client) UpgradeFabricChaincodeDefinitionWithOptions(request *UpgradeFabricChaincodeDefinitionRequest, runtime *util.RuntimeOptions) (_result *UpgradeFabricChaincodeDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChaincodeId)) {
		body["ChaincodeId"] = request.ChaincodeId
	}

	if !tea.BoolValue(util.IsUnset(request.ChaincodePackageId)) {
		body["ChaincodePackageId"] = request.ChaincodePackageId
	}

	if !tea.BoolValue(util.IsUnset(request.ChaincodeVersion)) {
		body["ChaincodeVersion"] = request.ChaincodeVersion
	}

	if !tea.BoolValue(util.IsUnset(request.CollectionConfig)) {
		body["CollectionConfig"] = request.CollectionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.EndorsePolicy)) {
		body["EndorsePolicy"] = request.EndorsePolicy
	}

	if !tea.BoolValue(util.IsUnset(request.InitRequired)) {
		body["InitRequired"] = request.InitRequired
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		body["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeFabricChaincodeDefinition"),
		Version:     tea.String("2018-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeFabricChaincodeDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// UpgradeFabricChaincodeDefinition
//
// @param request - UpgradeFabricChaincodeDefinitionRequest
//
// @return UpgradeFabricChaincodeDefinitionResponse
func (client *Client) UpgradeFabricChaincodeDefinition(request *UpgradeFabricChaincodeDefinitionRequest) (_result *UpgradeFabricChaincodeDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeFabricChaincodeDefinitionResponse{}
	_body, _err := client.UpgradeFabricChaincodeDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
