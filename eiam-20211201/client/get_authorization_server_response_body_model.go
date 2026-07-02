// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorizationServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationServer(v *GetAuthorizationServerResponseBodyAuthorizationServer) *GetAuthorizationServerResponseBody
	GetAuthorizationServer() *GetAuthorizationServerResponseBodyAuthorizationServer
	SetRequestId(v string) *GetAuthorizationServerResponseBody
	GetRequestId() *string
}

type GetAuthorizationServerResponseBody struct {
	// The authorization server.
	AuthorizationServer *GetAuthorizationServerResponseBodyAuthorizationServer `json:"AuthorizationServer,omitempty" xml:"AuthorizationServer,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAuthorizationServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationServerResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthorizationServerResponseBody) GetAuthorizationServer() *GetAuthorizationServerResponseBodyAuthorizationServer {
	return s.AuthorizationServer
}

func (s *GetAuthorizationServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuthorizationServerResponseBody) SetAuthorizationServer(v *GetAuthorizationServerResponseBodyAuthorizationServer) *GetAuthorizationServerResponseBody {
	s.AuthorizationServer = v
	return s
}

func (s *GetAuthorizationServerResponseBody) SetRequestId(v string) *GetAuthorizationServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthorizationServerResponseBody) Validate() error {
	if s.AuthorizationServer != nil {
		if err := s.AuthorizationServer.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAuthorizationServerResponseBodyAuthorizationServer struct {
	// The authorization server ID.
	//
	// example:
	//
	// iauths_system
	AuthorizationServerId *string `json:"AuthorizationServerId,omitempty" xml:"AuthorizationServerId,omitempty"`
	// The name of the authorization server.
	//
	// example:
	//
	// System_Default
	AuthorizationServerName *string `json:"AuthorizationServerName,omitempty" xml:"AuthorizationServerName,omitempty"`
	// The time when the authorization server was created.
	//
	// example:
	//
	// 1754620108295
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creation type.
	//
	// example:
	//
	// system_init
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
	// idaas_qzljgbhtwnnsywtdbz7yzy2any
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The issuer of the authorization token.
	//
	// example:
	//
	// https://xxxx.aliyunidaas.com/api/v2/iauths_system/oauth2
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
	// The time when the authorization server was last updated.
	//
	// example:
	//
	// 1781608572164
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// The endpoint configuration of the authorization server.
	ProtocolEndpoint *GetAuthorizationServerResponseBodyAuthorizationServerProtocolEndpoint `json:"ProtocolEndpoint,omitempty" xml:"ProtocolEndpoint,omitempty" type:"Struct"`
	// The status of the authorization server.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAuthorizationServerResponseBodyAuthorizationServer) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationServerResponseBodyAuthorizationServer) GoString() string {
	return s.String()
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) GetAuthorizationServerId() *string {
	return s.AuthorizationServerId
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) GetAuthorizationServerName() *string {
	return s.AuthorizationServerName
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) GetCreationType() *string {
	return s.CreationType
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) GetDescription() *string {
	return s.Description
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) GetIssuer() *string {
	return s.Issuer
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) GetIssuerDomain() *string {
	return s.IssuerDomain
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) GetIssuerMode() *string {
	return s.IssuerMode
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) GetLastUpdateTime() *int64 {
	return s.LastUpdateTime
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) GetProtocolEndpoint() *GetAuthorizationServerResponseBodyAuthorizationServerProtocolEndpoint {
	return s.ProtocolEndpoint
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) GetStatus() *string {
	return s.Status
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) SetAuthorizationServerId(v string) *GetAuthorizationServerResponseBodyAuthorizationServer {
	s.AuthorizationServerId = &v
	return s
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) SetAuthorizationServerName(v string) *GetAuthorizationServerResponseBodyAuthorizationServer {
	s.AuthorizationServerName = &v
	return s
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) SetCreateTime(v int64) *GetAuthorizationServerResponseBodyAuthorizationServer {
	s.CreateTime = &v
	return s
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) SetCreationType(v string) *GetAuthorizationServerResponseBodyAuthorizationServer {
	s.CreationType = &v
	return s
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) SetDescription(v string) *GetAuthorizationServerResponseBodyAuthorizationServer {
	s.Description = &v
	return s
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) SetInstanceId(v string) *GetAuthorizationServerResponseBodyAuthorizationServer {
	s.InstanceId = &v
	return s
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) SetIssuer(v string) *GetAuthorizationServerResponseBodyAuthorizationServer {
	s.Issuer = &v
	return s
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) SetIssuerDomain(v string) *GetAuthorizationServerResponseBodyAuthorizationServer {
	s.IssuerDomain = &v
	return s
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) SetIssuerMode(v string) *GetAuthorizationServerResponseBodyAuthorizationServer {
	s.IssuerMode = &v
	return s
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) SetLastUpdateTime(v int64) *GetAuthorizationServerResponseBodyAuthorizationServer {
	s.LastUpdateTime = &v
	return s
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) SetProtocolEndpoint(v *GetAuthorizationServerResponseBodyAuthorizationServerProtocolEndpoint) *GetAuthorizationServerResponseBodyAuthorizationServer {
	s.ProtocolEndpoint = v
	return s
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) SetStatus(v string) *GetAuthorizationServerResponseBodyAuthorizationServer {
	s.Status = &v
	return s
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServer) Validate() error {
	if s.ProtocolEndpoint != nil {
		if err := s.ProtocolEndpoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAuthorizationServerResponseBodyAuthorizationServerProtocolEndpoint struct {
	// The token endpoint of the authorization server.
	//
	// example:
	//
	// https://xxxx.aliyunidaas.com/api/v2/iauths_system/oauth2/token
	Oauth2TokenEndpoint *string `json:"Oauth2TokenEndpoint,omitempty" xml:"Oauth2TokenEndpoint,omitempty"`
	// The JWKS endpoint of the authorization server.
	//
	// example:
	//
	// https://xxxx.aliyunidaas.com/api/v2/iauths_system/oauth2/jwks
	OidcJwksEndpoint *string `json:"OidcJwksEndpoint,omitempty" xml:"OidcJwksEndpoint,omitempty"`
}

func (s GetAuthorizationServerResponseBodyAuthorizationServerProtocolEndpoint) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationServerResponseBodyAuthorizationServerProtocolEndpoint) GoString() string {
	return s.String()
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServerProtocolEndpoint) GetOauth2TokenEndpoint() *string {
	return s.Oauth2TokenEndpoint
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServerProtocolEndpoint) GetOidcJwksEndpoint() *string {
	return s.OidcJwksEndpoint
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServerProtocolEndpoint) SetOauth2TokenEndpoint(v string) *GetAuthorizationServerResponseBodyAuthorizationServerProtocolEndpoint {
	s.Oauth2TokenEndpoint = &v
	return s
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServerProtocolEndpoint) SetOidcJwksEndpoint(v string) *GetAuthorizationServerResponseBodyAuthorizationServerProtocolEndpoint {
	s.OidcJwksEndpoint = &v
	return s
}

func (s *GetAuthorizationServerResponseBodyAuthorizationServerProtocolEndpoint) Validate() error {
	return dara.Validate(s)
}
