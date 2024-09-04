// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateEenterpriseDIDRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 01357967-61d1-42a9-8a90-f0dd8a161411
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// F76iBtCdYuE0DHC33a5amzv3ioUfAqjuBxgek3RwjxBZ2cGP0HDxZy0k8Hs7yNtZRmM3h3KmRn4liSH4gLOOK4P41l
	OwnerUniqueID *string `json:"OwnerUniqueID,omitempty" xml:"OwnerUniqueID,omitempty"`
}

func (s CreateEenterpriseDIDRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEenterpriseDIDRequest) GoString() string {
	return s.String()
}

func (s *CreateEenterpriseDIDRequest) SetClientToken(v string) *CreateEenterpriseDIDRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEenterpriseDIDRequest) SetOwnerUniqueID(v string) *CreateEenterpriseDIDRequest {
	s.OwnerUniqueID = &v
	return s
}

type CreateEenterpriseDIDResponseBody struct {
	// example:
	//
	// "did:mychain:xxx"
	DID *string `json:"DID,omitempty" xml:"DID,omitempty"`
	// example:
	//
	// "7CEDB9B0-E68A-4E67-A258-EEE342695921"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// "OK"
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// ""
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEenterpriseDIDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEenterpriseDIDResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEenterpriseDIDResponseBody) SetDID(v string) *CreateEenterpriseDIDResponseBody {
	s.DID = &v
	return s
}

func (s *CreateEenterpriseDIDResponseBody) SetRequestId(v string) *CreateEenterpriseDIDResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEenterpriseDIDResponseBody) SetResultCode(v string) *CreateEenterpriseDIDResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateEenterpriseDIDResponseBody) SetResultMessage(v string) *CreateEenterpriseDIDResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateEenterpriseDIDResponseBody) SetSuccess(v bool) *CreateEenterpriseDIDResponseBody {
	s.Success = &v
	return s
}

type CreateEenterpriseDIDResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEenterpriseDIDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEenterpriseDIDResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEenterpriseDIDResponse) GoString() string {
	return s.String()
}

func (s *CreateEenterpriseDIDResponse) SetHeaders(v map[string]*string) *CreateEenterpriseDIDResponse {
	s.Headers = v
	return s
}

func (s *CreateEenterpriseDIDResponse) SetStatusCode(v int32) *CreateEenterpriseDIDResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEenterpriseDIDResponse) SetBody(v *CreateEenterpriseDIDResponseBody) *CreateEenterpriseDIDResponse {
	s.Body = v
	return s
}

type CreatePersonalDIDRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 01357967-61d1-42a9-8a90-f0dd8a161411
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 73BUce7y51BlPFxHucfjdOLC9cdWNnPqc7pyXXJe05c2twjraDPwDm1KhmD6v1d7tUyxjYoNhXLdX18zzp5rsogU87
	OwnerUniqueID *string `json:"OwnerUniqueID,omitempty" xml:"OwnerUniqueID,omitempty"`
}

func (s CreatePersonalDIDRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePersonalDIDRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalDIDRequest) SetClientToken(v string) *CreatePersonalDIDRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePersonalDIDRequest) SetOwnerUniqueID(v string) *CreatePersonalDIDRequest {
	s.OwnerUniqueID = &v
	return s
}

type CreatePersonalDIDResponseBody struct {
	// example:
	//
	// "did:mychain:xxx"
	DID *string `json:"DID,omitempty" xml:"DID,omitempty"`
	// example:
	//
	// "7C171509-B966-4AD2-B654-7BE14F1F3AA6"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// "OK"
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// ""
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePersonalDIDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePersonalDIDResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalDIDResponseBody) SetDID(v string) *CreatePersonalDIDResponseBody {
	s.DID = &v
	return s
}

func (s *CreatePersonalDIDResponseBody) SetRequestId(v string) *CreatePersonalDIDResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalDIDResponseBody) SetResultCode(v string) *CreatePersonalDIDResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreatePersonalDIDResponseBody) SetResultMessage(v string) *CreatePersonalDIDResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreatePersonalDIDResponseBody) SetSuccess(v bool) *CreatePersonalDIDResponseBody {
	s.Success = &v
	return s
}

type CreatePersonalDIDResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalDIDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalDIDResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePersonalDIDResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalDIDResponse) SetHeaders(v map[string]*string) *CreatePersonalDIDResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalDIDResponse) SetStatusCode(v int32) *CreatePersonalDIDResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalDIDResponse) SetBody(v *CreatePersonalDIDResponseBody) *CreatePersonalDIDResponse {
	s.Body = v
	return s
}

type CreateTenantDIDRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 01357967-61d1-42a9-8a90-f0dd8a161411
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateTenantDIDRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantDIDRequest) GoString() string {
	return s.String()
}

func (s *CreateTenantDIDRequest) SetClientToken(v string) *CreateTenantDIDRequest {
	s.ClientToken = &v
	return s
}

type CreateTenantDIDResponseBody struct {
	// example:
	//
	// "did:mychain:xxx"
	DID *string `json:"DID,omitempty" xml:"DID,omitempty"`
	// example:
	//
	// "757DB186-A865-4F65-935D-7D990E0CE451"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// "OK"
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// ""
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTenantDIDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantDIDResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTenantDIDResponseBody) SetDID(v string) *CreateTenantDIDResponseBody {
	s.DID = &v
	return s
}

func (s *CreateTenantDIDResponseBody) SetRequestId(v string) *CreateTenantDIDResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTenantDIDResponseBody) SetResultCode(v string) *CreateTenantDIDResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateTenantDIDResponseBody) SetResultMessage(v string) *CreateTenantDIDResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateTenantDIDResponseBody) SetSuccess(v bool) *CreateTenantDIDResponseBody {
	s.Success = &v
	return s
}

type CreateTenantDIDResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTenantDIDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTenantDIDResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantDIDResponse) GoString() string {
	return s.String()
}

func (s *CreateTenantDIDResponse) SetHeaders(v map[string]*string) *CreateTenantDIDResponse {
	s.Headers = v
	return s
}

func (s *CreateTenantDIDResponse) SetStatusCode(v int32) *CreateTenantDIDResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTenantDIDResponse) SetBody(v *CreateTenantDIDResponseBody) *CreateTenantDIDResponse {
	s.Body = v
	return s
}

type GetDIDRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// did:mychain:xxx
	DID *string `json:"DID,omitempty" xml:"DID,omitempty"`
}

func (s GetDIDRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDIDRequest) GoString() string {
	return s.String()
}

func (s *GetDIDRequest) SetDID(v string) *GetDIDRequest {
	s.DID = &v
	return s
}

type GetDIDResponseBody struct {
	// example:
	//
	// "{\\"creator\\":\\"did:mychain:xxx\\",\\"created\\":\\"2020-05-22T13:59:49+0800\\",\\"service\\":[],\\"index\\":[],\\"id\\":\\"did:mychain:xxx\\",\\"publicKey\\":[{\\"controller\\":\\"did:mychain:xxx\\",\\"id\\":\\"keys-1\\",\\"publicKey\\":\\"xxx\\",\\"type\\":\\"Secp256k1VerificationKey2018\\"}],\\"type\\":\\"Corporate\\",\\"@context\\":\\"https://w3id.org/did/v1\\",\\"updated\\":\\"2020-05-22T13:59:49+0800\\",\\"version\\":0,\\"authentication\\":[\\"keys-1\\"]}"
	Doc *string `json:"Doc,omitempty" xml:"Doc,omitempty"`
	// example:
	//
	// "2C93E421-AD9D-4ABE-B519-6E1ACD18934C"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// "OK"
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// ”“
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDIDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDIDResponseBody) GoString() string {
	return s.String()
}

func (s *GetDIDResponseBody) SetDoc(v string) *GetDIDResponseBody {
	s.Doc = &v
	return s
}

func (s *GetDIDResponseBody) SetRequestId(v string) *GetDIDResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDIDResponseBody) SetResultCode(v string) *GetDIDResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetDIDResponseBody) SetResultMessage(v string) *GetDIDResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *GetDIDResponseBody) SetSuccess(v bool) *GetDIDResponseBody {
	s.Success = &v
	return s
}

type GetDIDResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDIDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDIDResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDIDResponse) GoString() string {
	return s.String()
}

func (s *GetDIDResponse) SetHeaders(v map[string]*string) *GetDIDResponse {
	s.Headers = v
	return s
}

func (s *GetDIDResponse) SetStatusCode(v int32) *GetDIDResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDIDResponse) SetBody(v *GetDIDResponseBody) *GetDIDResponse {
	s.Body = v
	return s
}

type IssueNormalVerifiableVCRequest struct {
	// This parameter is required.
	BareClaimStructBody []*IssueNormalVerifiableVCRequestBareClaimStructBody `json:"BareClaimStructBody,omitempty" xml:"BareClaimStructBody,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 01357967-61d1-42a9-8a90-f0dd8a161411
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -1
	Expiration *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "did:mychain:xxx"
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "did:mychain:xxx"
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s IssueNormalVerifiableVCRequest) String() string {
	return tea.Prettify(s)
}

func (s IssueNormalVerifiableVCRequest) GoString() string {
	return s.String()
}

func (s *IssueNormalVerifiableVCRequest) SetBareClaimStructBody(v []*IssueNormalVerifiableVCRequestBareClaimStructBody) *IssueNormalVerifiableVCRequest {
	s.BareClaimStructBody = v
	return s
}

func (s *IssueNormalVerifiableVCRequest) SetClientToken(v string) *IssueNormalVerifiableVCRequest {
	s.ClientToken = &v
	return s
}

func (s *IssueNormalVerifiableVCRequest) SetExpiration(v int64) *IssueNormalVerifiableVCRequest {
	s.Expiration = &v
	return s
}

func (s *IssueNormalVerifiableVCRequest) SetIssuer(v string) *IssueNormalVerifiableVCRequest {
	s.Issuer = &v
	return s
}

func (s *IssueNormalVerifiableVCRequest) SetSubject(v string) *IssueNormalVerifiableVCRequest {
	s.Subject = &v
	return s
}

type IssueNormalVerifiableVCRequestBareClaimStructBody struct {
	// This parameter is required.
	//
	// example:
	//
	// "test"
	Claim *string `json:"Claim,omitempty" xml:"Claim,omitempty"`
	// example:
	//
	// ""
	ClaimType *string `json:"ClaimType,omitempty" xml:"ClaimType,omitempty"`
}

func (s IssueNormalVerifiableVCRequestBareClaimStructBody) String() string {
	return tea.Prettify(s)
}

func (s IssueNormalVerifiableVCRequestBareClaimStructBody) GoString() string {
	return s.String()
}

func (s *IssueNormalVerifiableVCRequestBareClaimStructBody) SetClaim(v string) *IssueNormalVerifiableVCRequestBareClaimStructBody {
	s.Claim = &v
	return s
}

func (s *IssueNormalVerifiableVCRequestBareClaimStructBody) SetClaimType(v string) *IssueNormalVerifiableVCRequestBareClaimStructBody {
	s.ClaimType = &v
	return s
}

type IssueNormalVerifiableVCResponseBody struct {
	// example:
	//
	// "4D1E29A7-17D6-48AD-B5AF-F44FAB68D87D"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// "OK"
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// ""
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// "{\\"proof\\":{\\"type\\":\\"ecdsa\\",\\"verificationMethod\\":\\"did:mychain:xxx#keys-1\\",\\"signatureValue\\":\\"xxx\\"},\\"content\\":{\\"issuanceDate\\":1590127960785,\\"subject\\":\\"did:mychain:xxx\\",\\"expire\\":-1,\\"claim\\":\\"test\\",\\"id\\":\\"vc:mychain:xxx\\",\\"type\\":[\\"VerifiableCredential\\"],\\"version\\":\\"0.7.0\\",\\"@context\\":\\"https://www.w3.org/2018/credentials/v1\\",\\"issuer\\":\\"did:mychain:xxx\\",\\"status\\":{\\"id\\":\\"vc:mychain:xxx\\",\\"type\\":\\"BlockChainStatusList\\"}}}"
	VerifiableClaimContent *string `json:"VerifiableClaimContent,omitempty" xml:"VerifiableClaimContent,omitempty"`
	// example:
	//
	// "vc:mychain:xxx"
	VerifiableClaimId *string `json:"VerifiableClaimId,omitempty" xml:"VerifiableClaimId,omitempty"`
}

func (s IssueNormalVerifiableVCResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IssueNormalVerifiableVCResponseBody) GoString() string {
	return s.String()
}

func (s *IssueNormalVerifiableVCResponseBody) SetRequestId(v string) *IssueNormalVerifiableVCResponseBody {
	s.RequestId = &v
	return s
}

func (s *IssueNormalVerifiableVCResponseBody) SetResultCode(v string) *IssueNormalVerifiableVCResponseBody {
	s.ResultCode = &v
	return s
}

func (s *IssueNormalVerifiableVCResponseBody) SetResultMessage(v string) *IssueNormalVerifiableVCResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *IssueNormalVerifiableVCResponseBody) SetSuccess(v bool) *IssueNormalVerifiableVCResponseBody {
	s.Success = &v
	return s
}

func (s *IssueNormalVerifiableVCResponseBody) SetVerifiableClaimContent(v string) *IssueNormalVerifiableVCResponseBody {
	s.VerifiableClaimContent = &v
	return s
}

func (s *IssueNormalVerifiableVCResponseBody) SetVerifiableClaimId(v string) *IssueNormalVerifiableVCResponseBody {
	s.VerifiableClaimId = &v
	return s
}

type IssueNormalVerifiableVCResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IssueNormalVerifiableVCResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IssueNormalVerifiableVCResponse) String() string {
	return tea.Prettify(s)
}

func (s IssueNormalVerifiableVCResponse) GoString() string {
	return s.String()
}

func (s *IssueNormalVerifiableVCResponse) SetHeaders(v map[string]*string) *IssueNormalVerifiableVCResponse {
	s.Headers = v
	return s
}

func (s *IssueNormalVerifiableVCResponse) SetStatusCode(v int32) *IssueNormalVerifiableVCResponse {
	s.StatusCode = &v
	return s
}

func (s *IssueNormalVerifiableVCResponse) SetBody(v *IssueNormalVerifiableVCResponseBody) *IssueNormalVerifiableVCResponse {
	s.Body = v
	return s
}

type UpdateVCRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// "did:mychain:xxx"
	IssuerDid *string `json:"IssuerDid,omitempty" xml:"IssuerDid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vc:mychain:xxx
	VCId *string `json:"VCId,omitempty" xml:"VCId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "1"
	VCStatus *string `json:"VCStatus,omitempty" xml:"VCStatus,omitempty"`
}

func (s UpdateVCRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVCRequest) GoString() string {
	return s.String()
}

func (s *UpdateVCRequest) SetIssuerDid(v string) *UpdateVCRequest {
	s.IssuerDid = &v
	return s
}

func (s *UpdateVCRequest) SetVCId(v string) *UpdateVCRequest {
	s.VCId = &v
	return s
}

func (s *UpdateVCRequest) SetVCStatus(v string) *UpdateVCRequest {
	s.VCStatus = &v
	return s
}

type UpdateVCResponseBody struct {
	// example:
	//
	// "1265B5EA-704A-4DCA-83F9-29C4D3B69549"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// "OK"
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// ""
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateVCResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVCResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVCResponseBody) SetRequestId(v string) *UpdateVCResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVCResponseBody) SetResultCode(v string) *UpdateVCResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UpdateVCResponseBody) SetResultMessage(v string) *UpdateVCResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *UpdateVCResponseBody) SetSuccess(v bool) *UpdateVCResponseBody {
	s.Success = &v
	return s
}

type UpdateVCResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVCResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVCResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVCResponse) GoString() string {
	return s.String()
}

func (s *UpdateVCResponse) SetHeaders(v map[string]*string) *UpdateVCResponse {
	s.Headers = v
	return s
}

func (s *UpdateVCResponse) SetStatusCode(v int32) *UpdateVCResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVCResponse) SetBody(v *UpdateVCResponseBody) *UpdateVCResponse {
	s.Body = v
	return s
}

type VerifyVerifiableClaimRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"proof": {"type": "ecdsa","verificationMethod": "did:mychain:xxx#keys-1","signatureValue": "xxx"},"content": {"issuanceDate": 1589964299367,"subject": "did:mychain:xxx","expire": -1,"claim": "test01","id": "vc:mychain:xxx","type": ["VerifiableCredential"],"version": "0.7.0","@context": "https://www.w3.org/2018/credentials/v1","issuer": "did:mychain:xxx","status": {"id": "vc:mychain:xxx","type": "BlockChainStatusList"}}}
	VCContent *string `json:"VCContent,omitempty" xml:"VCContent,omitempty"`
}

func (s VerifyVerifiableClaimRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyVerifiableClaimRequest) GoString() string {
	return s.String()
}

func (s *VerifyVerifiableClaimRequest) SetVCContent(v string) *VerifyVerifiableClaimRequest {
	s.VCContent = &v
	return s
}

type VerifyVerifiableClaimResponseBody struct {
	// example:
	//
	// "EEA284E9-B779-4E62-99EA-E2E8E801A745"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// "OK"
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// ""
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VerifyVerifiableClaimResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyVerifiableClaimResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyVerifiableClaimResponseBody) SetRequestId(v string) *VerifyVerifiableClaimResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyVerifiableClaimResponseBody) SetResultCode(v string) *VerifyVerifiableClaimResponseBody {
	s.ResultCode = &v
	return s
}

func (s *VerifyVerifiableClaimResponseBody) SetResultMessage(v string) *VerifyVerifiableClaimResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *VerifyVerifiableClaimResponseBody) SetSuccess(v bool) *VerifyVerifiableClaimResponseBody {
	s.Success = &v
	return s
}

type VerifyVerifiableClaimResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyVerifiableClaimResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyVerifiableClaimResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyVerifiableClaimResponse) GoString() string {
	return s.String()
}

func (s *VerifyVerifiableClaimResponse) SetHeaders(v map[string]*string) *VerifyVerifiableClaimResponse {
	s.Headers = v
	return s
}

func (s *VerifyVerifiableClaimResponse) SetStatusCode(v int32) *VerifyVerifiableClaimResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyVerifiableClaimResponse) SetBody(v *VerifyVerifiableClaimResponseBody) *VerifyVerifiableClaimResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("baasdis"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - CreateEenterpriseDIDRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEenterpriseDIDResponse
func (client *Client) CreateEenterpriseDIDWithOptions(request *CreateEenterpriseDIDRequest, runtime *util.RuntimeOptions) (_result *CreateEenterpriseDIDResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUniqueID)) {
		body["OwnerUniqueID"] = request.OwnerUniqueID
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEenterpriseDID"),
		Version:     tea.String("2020-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEenterpriseDIDResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateEenterpriseDIDRequest
//
// @return CreateEenterpriseDIDResponse
func (client *Client) CreateEenterpriseDID(request *CreateEenterpriseDIDRequest) (_result *CreateEenterpriseDIDResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEenterpriseDIDResponse{}
	_body, _err := client.CreateEenterpriseDIDWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreatePersonalDIDRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalDIDResponse
func (client *Client) CreatePersonalDIDWithOptions(request *CreatePersonalDIDRequest, runtime *util.RuntimeOptions) (_result *CreatePersonalDIDResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUniqueID)) {
		body["OwnerUniqueID"] = request.OwnerUniqueID
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePersonalDID"),
		Version:     tea.String("2020-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePersonalDIDResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreatePersonalDIDRequest
//
// @return CreatePersonalDIDResponse
func (client *Client) CreatePersonalDID(request *CreatePersonalDIDRequest) (_result *CreatePersonalDIDResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePersonalDIDResponse{}
	_body, _err := client.CreatePersonalDIDWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateTenantDIDRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTenantDIDResponse
func (client *Client) CreateTenantDIDWithOptions(request *CreateTenantDIDRequest, runtime *util.RuntimeOptions) (_result *CreateTenantDIDResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTenantDID"),
		Version:     tea.String("2020-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTenantDIDResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateTenantDIDRequest
//
// @return CreateTenantDIDResponse
func (client *Client) CreateTenantDID(request *CreateTenantDIDRequest) (_result *CreateTenantDIDResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTenantDIDResponse{}
	_body, _err := client.CreateTenantDIDWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetDIDRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDIDResponse
func (client *Client) GetDIDWithOptions(request *GetDIDRequest, runtime *util.RuntimeOptions) (_result *GetDIDResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DID)) {
		body["DID"] = request.DID
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDID"),
		Version:     tea.String("2020-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDIDResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetDIDRequest
//
// @return GetDIDResponse
func (client *Client) GetDID(request *GetDIDRequest) (_result *GetDIDResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDIDResponse{}
	_body, _err := client.GetDIDWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - IssueNormalVerifiableVCRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IssueNormalVerifiableVCResponse
func (client *Client) IssueNormalVerifiableVCWithOptions(request *IssueNormalVerifiableVCRequest, runtime *util.RuntimeOptions) (_result *IssueNormalVerifiableVCResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BareClaimStructBody)) {
		body["BareClaimStructBody"] = request.BareClaimStructBody
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Expiration)) {
		body["Expiration"] = request.Expiration
	}

	if !tea.BoolValue(util.IsUnset(request.Issuer)) {
		body["Issuer"] = request.Issuer
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["Subject"] = request.Subject
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("IssueNormalVerifiableVC"),
		Version:     tea.String("2020-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IssueNormalVerifiableVCResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - IssueNormalVerifiableVCRequest
//
// @return IssueNormalVerifiableVCResponse
func (client *Client) IssueNormalVerifiableVC(request *IssueNormalVerifiableVCRequest) (_result *IssueNormalVerifiableVCResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IssueNormalVerifiableVCResponse{}
	_body, _err := client.IssueNormalVerifiableVCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateVCRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVCResponse
func (client *Client) UpdateVCWithOptions(request *UpdateVCRequest, runtime *util.RuntimeOptions) (_result *UpdateVCResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IssuerDid)) {
		body["IssuerDid"] = request.IssuerDid
	}

	if !tea.BoolValue(util.IsUnset(request.VCId)) {
		body["VCId"] = request.VCId
	}

	if !tea.BoolValue(util.IsUnset(request.VCStatus)) {
		body["VCStatus"] = request.VCStatus
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVC"),
		Version:     tea.String("2020-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVCResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateVCRequest
//
// @return UpdateVCResponse
func (client *Client) UpdateVC(request *UpdateVCRequest) (_result *UpdateVCResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVCResponse{}
	_body, _err := client.UpdateVCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - VerifyVerifiableClaimRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyVerifiableClaimResponse
func (client *Client) VerifyVerifiableClaimWithOptions(request *VerifyVerifiableClaimRequest, runtime *util.RuntimeOptions) (_result *VerifyVerifiableClaimResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VCContent)) {
		body["VCContent"] = request.VCContent
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyVerifiableClaim"),
		Version:     tea.String("2020-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyVerifiableClaimResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - VerifyVerifiableClaimRequest
//
// @return VerifyVerifiableClaimResponse
func (client *Client) VerifyVerifiableClaim(request *VerifyVerifiableClaimRequest) (_result *VerifyVerifiableClaimResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyVerifiableClaimResponse{}
	_body, _err := client.VerifyVerifiableClaimWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
