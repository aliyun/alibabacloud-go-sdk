// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOAuthAuthorizationSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationTokenId(v string) *GetOAuthAuthorizationSessionResponseBody
	GetAuthenticationTokenId() *string
	SetAuthorizationUrl(v string) *GetOAuthAuthorizationSessionResponseBody
	GetAuthorizationUrl() *string
	SetConsumerId(v string) *GetOAuthAuthorizationSessionResponseBody
	GetConsumerId() *string
	SetConsumerType(v string) *GetOAuthAuthorizationSessionResponseBody
	GetConsumerType() *string
	SetCreatorId(v string) *GetOAuthAuthorizationSessionResponseBody
	GetCreatorId() *string
	SetCreatorType(v string) *GetOAuthAuthorizationSessionResponseBody
	GetCreatorType() *string
	SetCredentialProviderIdentifier(v string) *GetOAuthAuthorizationSessionResponseBody
	GetCredentialProviderIdentifier() *string
	SetErrorCode(v string) *GetOAuthAuthorizationSessionResponseBody
	GetErrorCode() *string
	SetErrorDescription(v string) *GetOAuthAuthorizationSessionResponseBody
	GetErrorDescription() *string
	SetExpirationTime(v int64) *GetOAuthAuthorizationSessionResponseBody
	GetExpirationTime() *int64
	SetInstanceId(v string) *GetOAuthAuthorizationSessionResponseBody
	GetInstanceId() *string
	SetSessionId(v string) *GetOAuthAuthorizationSessionResponseBody
	GetSessionId() *string
	SetSessionStatus(v string) *GetOAuthAuthorizationSessionResponseBody
	GetSessionStatus() *string
	SetSessionUri(v string) *GetOAuthAuthorizationSessionResponseBody
	GetSessionUri() *string
}

type GetOAuthAuthorizationSessionResponseBody struct {
	// The authentication token ID.
	//
	// example:
	//
	// atntkn_01l6lot7o4e4r77oelp6qtuxxxxx
	AuthenticationTokenId *string `json:"authenticationTokenId,omitempty" xml:"authenticationTokenId,omitempty"`
	// The user authorization URL.
	//
	// example:
	//
	// https://login.dingtalk.com/oauth2/auth?client_id=...
	AuthorizationUrl *string `json:"authorizationUrl,omitempty" xml:"authorizationUrl,omitempty"`
	// The authentication token consumer ID.
	//
	// example:
	//
	// app_ngtkgrrxxxxktg5eao6z4xxxxx
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// The authentication token consumer type.
	//
	// example:
	//
	// application
	ConsumerType *string `json:"consumerType,omitempty" xml:"consumerType,omitempty"`
	// The authentication token creator ID.
	//
	// example:
	//
	// app_ngtkgrrxxxxktg5eao6z4xxxxx
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The authentication token creator type.
	//
	// example:
	//
	// application
	CreatorType *string `json:"creatorType,omitempty" xml:"creatorType,omitempty"`
	// The credential provider business identifier.
	//
	// example:
	//
	// test_example_identifier
	CredentialProviderIdentifier *string `json:"credentialProviderIdentifier,omitempty" xml:"credentialProviderIdentifier,omitempty"`
	// The error code.
	//
	// example:
	//
	// access_denied
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error description.
	//
	// example:
	//
	// The user denied the authorization request
	ErrorDescription *string `json:"errorDescription,omitempty" xml:"errorDescription,omitempty"`
	// The authentication token expiration time. UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1704153600000
	ExpirationTime *int64 `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The authorization session ID.
	//
	// example:
	//
	// atpoas_01l6losojlojbbv01adsq56xxxxx
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The authorization session status.
	//
	// example:
	//
	// pending
	SessionStatus *string `json:"sessionStatus,omitempty" xml:"sessionStatus,omitempty"`
	// The authorization session URI.
	//
	// example:
	//
	// urn:ietf:params:oauth:request_uri:atpoas_01l6ljnvrpc5niakl3gj3amxxxxxx
	SessionUri *string `json:"sessionUri,omitempty" xml:"sessionUri,omitempty"`
}

func (s GetOAuthAuthorizationSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOAuthAuthorizationSessionResponseBody) GoString() string {
	return s.String()
}

func (s *GetOAuthAuthorizationSessionResponseBody) GetAuthenticationTokenId() *string {
	return s.AuthenticationTokenId
}

func (s *GetOAuthAuthorizationSessionResponseBody) GetAuthorizationUrl() *string {
	return s.AuthorizationUrl
}

func (s *GetOAuthAuthorizationSessionResponseBody) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *GetOAuthAuthorizationSessionResponseBody) GetConsumerType() *string {
	return s.ConsumerType
}

func (s *GetOAuthAuthorizationSessionResponseBody) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetOAuthAuthorizationSessionResponseBody) GetCreatorType() *string {
	return s.CreatorType
}

func (s *GetOAuthAuthorizationSessionResponseBody) GetCredentialProviderIdentifier() *string {
	return s.CredentialProviderIdentifier
}

func (s *GetOAuthAuthorizationSessionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetOAuthAuthorizationSessionResponseBody) GetErrorDescription() *string {
	return s.ErrorDescription
}

func (s *GetOAuthAuthorizationSessionResponseBody) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *GetOAuthAuthorizationSessionResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetOAuthAuthorizationSessionResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *GetOAuthAuthorizationSessionResponseBody) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *GetOAuthAuthorizationSessionResponseBody) GetSessionUri() *string {
	return s.SessionUri
}

func (s *GetOAuthAuthorizationSessionResponseBody) SetAuthenticationTokenId(v string) *GetOAuthAuthorizationSessionResponseBody {
	s.AuthenticationTokenId = &v
	return s
}

func (s *GetOAuthAuthorizationSessionResponseBody) SetAuthorizationUrl(v string) *GetOAuthAuthorizationSessionResponseBody {
	s.AuthorizationUrl = &v
	return s
}

func (s *GetOAuthAuthorizationSessionResponseBody) SetConsumerId(v string) *GetOAuthAuthorizationSessionResponseBody {
	s.ConsumerId = &v
	return s
}

func (s *GetOAuthAuthorizationSessionResponseBody) SetConsumerType(v string) *GetOAuthAuthorizationSessionResponseBody {
	s.ConsumerType = &v
	return s
}

func (s *GetOAuthAuthorizationSessionResponseBody) SetCreatorId(v string) *GetOAuthAuthorizationSessionResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetOAuthAuthorizationSessionResponseBody) SetCreatorType(v string) *GetOAuthAuthorizationSessionResponseBody {
	s.CreatorType = &v
	return s
}

func (s *GetOAuthAuthorizationSessionResponseBody) SetCredentialProviderIdentifier(v string) *GetOAuthAuthorizationSessionResponseBody {
	s.CredentialProviderIdentifier = &v
	return s
}

func (s *GetOAuthAuthorizationSessionResponseBody) SetErrorCode(v string) *GetOAuthAuthorizationSessionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOAuthAuthorizationSessionResponseBody) SetErrorDescription(v string) *GetOAuthAuthorizationSessionResponseBody {
	s.ErrorDescription = &v
	return s
}

func (s *GetOAuthAuthorizationSessionResponseBody) SetExpirationTime(v int64) *GetOAuthAuthorizationSessionResponseBody {
	s.ExpirationTime = &v
	return s
}

func (s *GetOAuthAuthorizationSessionResponseBody) SetInstanceId(v string) *GetOAuthAuthorizationSessionResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetOAuthAuthorizationSessionResponseBody) SetSessionId(v string) *GetOAuthAuthorizationSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *GetOAuthAuthorizationSessionResponseBody) SetSessionStatus(v string) *GetOAuthAuthorizationSessionResponseBody {
	s.SessionStatus = &v
	return s
}

func (s *GetOAuthAuthorizationSessionResponseBody) SetSessionUri(v string) *GetOAuthAuthorizationSessionResponseBody {
	s.SessionUri = &v
	return s
}

func (s *GetOAuthAuthorizationSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
