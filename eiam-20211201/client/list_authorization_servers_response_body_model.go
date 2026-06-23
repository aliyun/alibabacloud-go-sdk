// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationServers(v []*ListAuthorizationServersResponseBodyAuthorizationServers) *ListAuthorizationServersResponseBody
	GetAuthorizationServers() []*ListAuthorizationServersResponseBodyAuthorizationServers
	SetMaxResults(v int32) *ListAuthorizationServersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAuthorizationServersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAuthorizationServersResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListAuthorizationServersResponseBody
	GetTotalCount() *int64
}

type ListAuthorizationServersResponseBody struct {
	// The list of authorization servers.
	AuthorizationServers []*ListAuthorizationServersResponseBodyAuthorizationServers `json:"AuthorizationServers,omitempty" xml:"AuthorizationServers,omitempty" type:"Repeated"`
	// The number of entries per page for a paged query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token returned in this call.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of authorization servers.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAuthorizationServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizationServersResponseBody) GetAuthorizationServers() []*ListAuthorizationServersResponseBodyAuthorizationServers {
	return s.AuthorizationServers
}

func (s *ListAuthorizationServersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAuthorizationServersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuthorizationServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizationServersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAuthorizationServersResponseBody) SetAuthorizationServers(v []*ListAuthorizationServersResponseBodyAuthorizationServers) *ListAuthorizationServersResponseBody {
	s.AuthorizationServers = v
	return s
}

func (s *ListAuthorizationServersResponseBody) SetMaxResults(v int32) *ListAuthorizationServersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizationServersResponseBody) SetNextToken(v string) *ListAuthorizationServersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizationServersResponseBody) SetRequestId(v string) *ListAuthorizationServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizationServersResponseBody) SetTotalCount(v int64) *ListAuthorizationServersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuthorizationServersResponseBody) Validate() error {
	if s.AuthorizationServers != nil {
		for _, item := range s.AuthorizationServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuthorizationServersResponseBodyAuthorizationServers struct {
	// The unique identifier of the authorization server.
	//
	// example:
	//
	// iauths_system
	AuthorizationServerId *string `json:"AuthorizationServerId,omitempty" xml:"AuthorizationServerId,omitempty"`
	// The name of the authorization server.
	//
	// example:
	//
	// my_authorization_server
	AuthorizationServerName *string `json:"AuthorizationServerName,omitempty" xml:"AuthorizationServerName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creation type.
	//
	// example:
	//
	// user_custom
	CreationType *string `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	// The description of the authorization server.
	//
	// example:
	//
	// description of authorization server
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The currently active issuer address.
	//
	// example:
	//
	// https://xxxx.aliyunidaas.com/api/v2/iauths_system/oauth2/token
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The domain name used by the issuer.
	//
	// example:
	//
	// xxxx.aliyunidaas.com
	IssuerDomain *string `json:"IssuerDomain,omitempty" xml:"IssuerDomain,omitempty"`
	// The issuer mode.
	//
	// example:
	//
	// static
	IssuerMode *string `json:"IssuerMode,omitempty" xml:"IssuerMode,omitempty"`
	// The status.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListAuthorizationServersResponseBodyAuthorizationServers) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationServersResponseBodyAuthorizationServers) GoString() string {
	return s.String()
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) GetAuthorizationServerId() *string {
	return s.AuthorizationServerId
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) GetAuthorizationServerName() *string {
	return s.AuthorizationServerName
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) GetCreationType() *string {
	return s.CreationType
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) GetDescription() *string {
	return s.Description
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) GetIssuer() *string {
	return s.Issuer
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) GetIssuerDomain() *string {
	return s.IssuerDomain
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) GetIssuerMode() *string {
	return s.IssuerMode
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) GetStatus() *string {
	return s.Status
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) SetAuthorizationServerId(v string) *ListAuthorizationServersResponseBodyAuthorizationServers {
	s.AuthorizationServerId = &v
	return s
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) SetAuthorizationServerName(v string) *ListAuthorizationServersResponseBodyAuthorizationServers {
	s.AuthorizationServerName = &v
	return s
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) SetCreateTime(v int64) *ListAuthorizationServersResponseBodyAuthorizationServers {
	s.CreateTime = &v
	return s
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) SetCreationType(v string) *ListAuthorizationServersResponseBodyAuthorizationServers {
	s.CreationType = &v
	return s
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) SetDescription(v string) *ListAuthorizationServersResponseBodyAuthorizationServers {
	s.Description = &v
	return s
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) SetInstanceId(v string) *ListAuthorizationServersResponseBodyAuthorizationServers {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) SetIssuer(v string) *ListAuthorizationServersResponseBodyAuthorizationServers {
	s.Issuer = &v
	return s
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) SetIssuerDomain(v string) *ListAuthorizationServersResponseBodyAuthorizationServers {
	s.IssuerDomain = &v
	return s
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) SetIssuerMode(v string) *ListAuthorizationServersResponseBodyAuthorizationServers {
	s.IssuerMode = &v
	return s
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) SetStatus(v string) *ListAuthorizationServersResponseBodyAuthorizationServers {
	s.Status = &v
	return s
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) SetUpdateTime(v int64) *ListAuthorizationServersResponseBodyAuthorizationServers {
	s.UpdateTime = &v
	return s
}

func (s *ListAuthorizationServersResponseBodyAuthorizationServers) Validate() error {
	return dara.Validate(s)
}
