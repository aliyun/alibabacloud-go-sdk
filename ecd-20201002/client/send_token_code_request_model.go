// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendTokenCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *SendTokenCodeRequest
	GetClientId() *string
	SetClientOS(v string) *SendTokenCodeRequest
	GetClientOS() *string
	SetClientVersion(v string) *SendTokenCodeRequest
	GetClientVersion() *string
	SetEndUserId(v string) *SendTokenCodeRequest
	GetEndUserId() *string
	SetLoginToken(v string) *SendTokenCodeRequest
	GetLoginToken() *string
	SetOfficeSiteId(v string) *SendTokenCodeRequest
	GetOfficeSiteId() *string
	SetSessionId(v string) *SendTokenCodeRequest
	GetSessionId() *string
	SetTokenCode(v string) *SendTokenCodeRequest
	GetTokenCode() *string
}

type SendTokenCodeRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// f4a0dc8e-1702-4728-9a60-95b27a35****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The operating system on which the client runs.
	//
	// example:
	//
	// Windows_NT 10.0.18363 x64
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The client version. If you use an Alibaba Cloud Workspace client, you can view the client version in the "About" dialog box on the client logon page.
	//
	// example:
	//
	// 2.1.0-R-20210731.151756
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The logon token.
	//
	// example:
	//
	// v28101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-2925105532
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// cd45e873-650d-4d70-acb9-f996187a****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// If two-factor authentication is enabled for clients in the Elastic Desktop Service (EDS) Enterprise console, the system will send a verification code to the user\\"s email address if it detects that the current logged-on user is at risk. This parameter is required if you set `CurrentStage` to `TokenVerify`.
	//
	// example:
	//
	// 63****
	TokenCode *string `json:"TokenCode,omitempty" xml:"TokenCode,omitempty"`
}

func (s SendTokenCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s SendTokenCodeRequest) GoString() string {
	return s.String()
}

func (s *SendTokenCodeRequest) GetClientId() *string {
	return s.ClientId
}

func (s *SendTokenCodeRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *SendTokenCodeRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *SendTokenCodeRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *SendTokenCodeRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *SendTokenCodeRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *SendTokenCodeRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *SendTokenCodeRequest) GetTokenCode() *string {
	return s.TokenCode
}

func (s *SendTokenCodeRequest) SetClientId(v string) *SendTokenCodeRequest {
	s.ClientId = &v
	return s
}

func (s *SendTokenCodeRequest) SetClientOS(v string) *SendTokenCodeRequest {
	s.ClientOS = &v
	return s
}

func (s *SendTokenCodeRequest) SetClientVersion(v string) *SendTokenCodeRequest {
	s.ClientVersion = &v
	return s
}

func (s *SendTokenCodeRequest) SetEndUserId(v string) *SendTokenCodeRequest {
	s.EndUserId = &v
	return s
}

func (s *SendTokenCodeRequest) SetLoginToken(v string) *SendTokenCodeRequest {
	s.LoginToken = &v
	return s
}

func (s *SendTokenCodeRequest) SetOfficeSiteId(v string) *SendTokenCodeRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *SendTokenCodeRequest) SetSessionId(v string) *SendTokenCodeRequest {
	s.SessionId = &v
	return s
}

func (s *SendTokenCodeRequest) SetTokenCode(v string) *SendTokenCodeRequest {
	s.TokenCode = &v
	return s
}

func (s *SendTokenCodeRequest) Validate() error {
	return dara.Validate(s)
}
