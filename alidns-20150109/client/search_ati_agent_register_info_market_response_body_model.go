// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAtiAgentRegisterInfoMarketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) *SearchAtiAgentRegisterInfoMarketResponseBody
	GetAccessDeniedDetail() *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail
	SetAgents(v *SearchAtiAgentRegisterInfoMarketResponseBodyAgents) *SearchAtiAgentRegisterInfoMarketResponseBody
	GetAgents() *SearchAtiAgentRegisterInfoMarketResponseBodyAgents
	SetMaxResults(v int32) *SearchAtiAgentRegisterInfoMarketResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *SearchAtiAgentRegisterInfoMarketResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *SearchAtiAgentRegisterInfoMarketResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchAtiAgentRegisterInfoMarketResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *SearchAtiAgentRegisterInfoMarketResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *SearchAtiAgentRegisterInfoMarketResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *SearchAtiAgentRegisterInfoMarketResponseBody
	GetTotalPages() *int32
}

type SearchAtiAgentRegisterInfoMarketResponseBody struct {
	// The details of the access denial. This field is returned only when RAM authentication fails.
	AccessDeniedDetail *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Agents             *SearchAtiAgentRegisterInfoMarketResponseBodyAgents             `json:"Agents,omitempty" xml:"Agents,omitempty" type:"Struct"`
	// The number of entries per batch query.
	//
	// example:
	//
	// 500
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query.
	//
	// example:
	//
	// 4698691
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The current page number. Minimum value: **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for the paged query. Settings: maximum value: **100**. Default value: **20**. This parameter controls paging behavior.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 11
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s SearchAtiAgentRegisterInfoMarketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchAtiAgentRegisterInfoMarketResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) GetAccessDeniedDetail() *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) GetAgents() *SearchAtiAgentRegisterInfoMarketResponseBodyAgents {
	return s.Agents
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) SetAccessDeniedDetail(v *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) *SearchAtiAgentRegisterInfoMarketResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) SetAgents(v *SearchAtiAgentRegisterInfoMarketResponseBodyAgents) *SearchAtiAgentRegisterInfoMarketResponseBody {
	s.Agents = v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) SetMaxResults(v int32) *SearchAtiAgentRegisterInfoMarketResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) SetNextToken(v string) *SearchAtiAgentRegisterInfoMarketResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) SetPageNumber(v int32) *SearchAtiAgentRegisterInfoMarketResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) SetPageSize(v int32) *SearchAtiAgentRegisterInfoMarketResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) SetRequestId(v string) *SearchAtiAgentRegisterInfoMarketResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) SetTotalItems(v int32) *SearchAtiAgentRegisterInfoMarketResponseBody {
	s.TotalItems = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) SetTotalPages(v int32) *SearchAtiAgentRegisterInfoMarketResponseBody {
	s.TotalPages = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Agents != nil {
		if err := s.Agents.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that was attempted.
	//
	// example:
	//
	// CreateUser
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authorization principal.
	//
	// example:
	//
	// 2015555733387XXXX
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authorization principal.
	//
	// example:
	//
	// 10469733312XXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The identity type.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encrypted diagnostic message.
	//
	// example:
	//
	// AQFohtp4aIbaeEXXXXQxNjFDLUIzMzgtNTXXXX05NkFCLUI2RkY5XXXXzAzQQ==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The reason for the authentication failure. Valid values:
	//
	// - ExplicitDeny: explicit deny.
	//
	// - ImplicitDeny: implicit deny.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) SetAuthAction(v string) *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) SetPolicyType(v string) *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type SearchAtiAgentRegisterInfoMarketResponseBodyAgents struct {
	Agent []*SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Repeated"`
}

func (s SearchAtiAgentRegisterInfoMarketResponseBodyAgents) String() string {
	return dara.Prettify(s)
}

func (s SearchAtiAgentRegisterInfoMarketResponseBodyAgents) GoString() string {
	return s.String()
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgents) GetAgent() []*SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent {
	return s.Agent
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgents) SetAgent(v []*SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) *SearchAtiAgentRegisterInfoMarketResponseBodyAgents {
	s.Agent = v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgents) Validate() error {
	if s.Agent != nil {
		for _, item := range s.Agent {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent struct {
	AgentDescription *string                                                           `json:"AgentDescription,omitempty" xml:"AgentDescription,omitempty"`
	AgentDisplayName *string                                                           `json:"AgentDisplayName,omitempty" xml:"AgentDisplayName,omitempty"`
	AgentHost        *string                                                           `json:"AgentHost,omitempty" xml:"AgentHost,omitempty"`
	AgentId          *string                                                           `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentVersion     *string                                                           `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	CreateTimestamp  *int64                                                            `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Protocols        *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgentProtocols `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Struct"`
	Status           *string                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	TrustCardUrl     *string                                                           `json:"TrustCardUrl,omitempty" xml:"TrustCardUrl,omitempty"`
	TrustLevel       *string                                                           `json:"TrustLevel,omitempty" xml:"TrustLevel,omitempty"`
	UpdateTimestamp  *int64                                                            `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) String() string {
	return dara.Prettify(s)
}

func (s SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) GoString() string {
	return s.String()
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) GetAgentDescription() *string {
	return s.AgentDescription
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) GetAgentDisplayName() *string {
	return s.AgentDisplayName
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) GetAgentHost() *string {
	return s.AgentHost
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) GetAgentId() *string {
	return s.AgentId
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) GetProtocols() *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgentProtocols {
	return s.Protocols
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) GetStatus() *string {
	return s.Status
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) GetTrustCardUrl() *string {
	return s.TrustCardUrl
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) GetTrustLevel() *string {
	return s.TrustLevel
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) SetAgentDescription(v string) *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent {
	s.AgentDescription = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) SetAgentDisplayName(v string) *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent {
	s.AgentDisplayName = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) SetAgentHost(v string) *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent {
	s.AgentHost = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) SetAgentId(v string) *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent {
	s.AgentId = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) SetAgentVersion(v string) *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent {
	s.AgentVersion = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) SetCreateTimestamp(v int64) *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent {
	s.CreateTimestamp = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) SetProtocols(v *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgentProtocols) *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent {
	s.Protocols = v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) SetStatus(v string) *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent {
	s.Status = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) SetTrustCardUrl(v string) *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent {
	s.TrustCardUrl = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) SetTrustLevel(v string) *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent {
	s.TrustLevel = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) SetUpdateTimestamp(v int64) *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent {
	s.UpdateTimestamp = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgent) Validate() error {
	if s.Protocols != nil {
		if err := s.Protocols.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgentProtocols struct {
	Protocol []*string `json:"Protocol,omitempty" xml:"Protocol,omitempty" type:"Repeated"`
}

func (s SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgentProtocols) String() string {
	return dara.Prettify(s)
}

func (s SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgentProtocols) GoString() string {
	return s.String()
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgentProtocols) GetProtocol() []*string {
	return s.Protocol
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgentProtocols) SetProtocol(v []*string) *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgentProtocols {
	s.Protocol = v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponseBodyAgentsAgentProtocols) Validate() error {
	return dara.Validate(s)
}
