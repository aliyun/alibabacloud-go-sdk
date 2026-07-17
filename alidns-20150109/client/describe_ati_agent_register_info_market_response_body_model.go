// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAtiAgentRegisterInfoMarketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) *DescribeAtiAgentRegisterInfoMarketResponseBody
	GetAccessDeniedDetail() *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail
	SetAgentDescription(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody
	GetAgentDescription() *string
	SetAgentDisplayName(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody
	GetAgentDisplayName() *string
	SetAgentHost(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody
	GetAgentHost() *string
	SetAgentId(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody
	GetAgentId() *string
	SetAgentRegisterInfoId(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody
	GetAgentRegisterInfoId() *string
	SetAgentVersion(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody
	GetAgentVersion() *string
	SetCategories(v *DescribeAtiAgentRegisterInfoMarketResponseBodyCategories) *DescribeAtiAgentRegisterInfoMarketResponseBody
	GetCategories() *DescribeAtiAgentRegisterInfoMarketResponseBodyCategories
	SetEndpoints(v *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpoints) *DescribeAtiAgentRegisterInfoMarketResponseBody
	GetEndpoints() *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpoints
	SetMaxResults(v int32) *DescribeAtiAgentRegisterInfoMarketResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody
	GetStatus() *string
	SetTrustCardContent(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody
	GetTrustCardContent() *string
	SetTrustLevel(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody
	GetTrustLevel() *string
}

type DescribeAtiAgentRegisterInfoMarketResponseBody struct {
	AccessDeniedDetail *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 支付服务
	AgentDescription *string `json:"AgentDescription,omitempty" xml:"AgentDescription,omitempty"`
	// example:
	//
	// 测试Agent
	AgentDisplayName *string `json:"AgentDisplayName,omitempty" xml:"AgentDisplayName,omitempty"`
	// example:
	//
	// www.example.com
	AgentHost *string `json:"AgentHost,omitempty" xml:"AgentHost,omitempty"`
	// example:
	//
	// gsc01629925@5e0964fd-951c-4e45-b518-d09d4d2db8ca
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 2074753647748672512
	AgentRegisterInfoId *string `json:"AgentRegisterInfoId,omitempty" xml:"AgentRegisterInfoId,omitempty"`
	// example:
	//
	// 3.9.3
	AgentVersion *string                                                   `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	Categories   *DescribeAtiAgentRegisterInfoMarketResponseBodyCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Struct"`
	Endpoints    *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpoints  `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	// example:
	//
	// 500
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 4698691
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 0629502C-6224-5DC9-A8ED-2ED73A2E3931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Disable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 忽略
	TrustCardContent *string `json:"TrustCardContent,omitempty" xml:"TrustCardContent,omitempty"`
	// example:
	//
	// 基础认证
	TrustLevel *string `json:"TrustLevel,omitempty" xml:"TrustLevel,omitempty"`
}

func (s DescribeAtiAgentRegisterInfoMarketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiAgentRegisterInfoMarketResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) GetAccessDeniedDetail() *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) GetAgentDescription() *string {
	return s.AgentDescription
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) GetAgentDisplayName() *string {
	return s.AgentDisplayName
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) GetAgentHost() *string {
	return s.AgentHost
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) GetAgentRegisterInfoId() *string {
	return s.AgentRegisterInfoId
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) GetCategories() *DescribeAtiAgentRegisterInfoMarketResponseBodyCategories {
	return s.Categories
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) GetEndpoints() *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpoints {
	return s.Endpoints
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) GetTrustCardContent() *string {
	return s.TrustCardContent
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) GetTrustLevel() *string {
	return s.TrustLevel
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) SetAccessDeniedDetail(v *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) *DescribeAtiAgentRegisterInfoMarketResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) SetAgentDescription(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody {
	s.AgentDescription = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) SetAgentDisplayName(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody {
	s.AgentDisplayName = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) SetAgentHost(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody {
	s.AgentHost = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) SetAgentId(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody {
	s.AgentId = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) SetAgentRegisterInfoId(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody {
	s.AgentRegisterInfoId = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) SetAgentVersion(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody {
	s.AgentVersion = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) SetCategories(v *DescribeAtiAgentRegisterInfoMarketResponseBodyCategories) *DescribeAtiAgentRegisterInfoMarketResponseBody {
	s.Categories = v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) SetEndpoints(v *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpoints) *DescribeAtiAgentRegisterInfoMarketResponseBody {
	s.Endpoints = v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) SetMaxResults(v int32) *DescribeAtiAgentRegisterInfoMarketResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) SetNextToken(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) SetRequestId(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) SetStatus(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) SetTrustCardContent(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody {
	s.TrustCardContent = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) SetTrustLevel(v string) *DescribeAtiAgentRegisterInfoMarketResponseBody {
	s.TrustLevel = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Categories != nil {
		if err := s.Categories.Validate(); err != nil {
			return err
		}
	}
	if s.Endpoints != nil {
		if err := s.Endpoints.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// AddRspDomainServerHoldStatusForGatewayOte
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// 2015555733387XXXX
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// 1046973331XXXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// AQFohtp4aIbaeEXXXXQxNjFDLUIzMzgtNTXXXX05NkFCLUI2RkY5XXXXzAzQQ==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeAtiAgentRegisterInfoMarketResponseBodyCategories struct {
	Category []*string `json:"category,omitempty" xml:"category,omitempty" type:"Repeated"`
}

func (s DescribeAtiAgentRegisterInfoMarketResponseBodyCategories) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiAgentRegisterInfoMarketResponseBodyCategories) GoString() string {
	return s.String()
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyCategories) GetCategory() []*string {
	return s.Category
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyCategories) SetCategory(v []*string) *DescribeAtiAgentRegisterInfoMarketResponseBodyCategories {
	s.Category = v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyCategories) Validate() error {
	return dara.Validate(s)
}

type DescribeAtiAgentRegisterInfoMarketResponseBodyEndpoints struct {
	Endpoint []*DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint `json:"Endpoint,omitempty" xml:"Endpoint,omitempty" type:"Repeated"`
}

func (s DescribeAtiAgentRegisterInfoMarketResponseBodyEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiAgentRegisterInfoMarketResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpoints) GetEndpoint() []*DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint {
	return s.Endpoint
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpoints) SetEndpoint(v []*DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint) *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpoints {
	s.Endpoint = v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpoints) Validate() error {
	if s.Endpoint != nil {
		for _, item := range s.Endpoint {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint struct {
	AgentUrl    *string                                                                    `json:"AgentUrl,omitempty" xml:"AgentUrl,omitempty"`
	MetadataUrl *string                                                                    `json:"MetadataUrl,omitempty" xml:"MetadataUrl,omitempty"`
	Protocol    *string                                                                    `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Transports  *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpointTransports `json:"Transports,omitempty" xml:"Transports,omitempty" type:"Struct"`
}

func (s DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint) GetAgentUrl() *string {
	return s.AgentUrl
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint) GetMetadataUrl() *string {
	return s.MetadataUrl
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint) GetTransports() *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpointTransports {
	return s.Transports
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint) SetAgentUrl(v string) *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint {
	s.AgentUrl = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint) SetMetadataUrl(v string) *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint {
	s.MetadataUrl = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint) SetProtocol(v string) *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint {
	s.Protocol = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint) SetTransports(v *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpointTransports) *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint {
	s.Transports = v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpoint) Validate() error {
	if s.Transports != nil {
		if err := s.Transports.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpointTransports struct {
	Transport []*string `json:"Transport,omitempty" xml:"Transport,omitempty" type:"Repeated"`
}

func (s DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpointTransports) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpointTransports) GoString() string {
	return s.String()
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpointTransports) GetTransport() []*string {
	return s.Transport
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpointTransports) SetTransport(v []*string) *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpointTransports {
	s.Transport = v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponseBodyEndpointsEndpointTransports) Validate() error {
	return dara.Validate(s)
}
