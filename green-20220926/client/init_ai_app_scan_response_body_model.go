// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitAiAppScanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthInfo(v *InitAiAppScanResponseBodyAuthInfo) *InitAiAppScanResponseBody
	GetAuthInfo() *InitAiAppScanResponseBodyAuthInfo
	SetAuthInfoConfig(v map[string]*AuthInfoConfigValue) *InitAiAppScanResponseBody
	GetAuthInfoConfig() map[string]*AuthInfoConfigValue
	SetAuthStatus(v string) *InitAiAppScanResponseBody
	GetAuthStatus() *string
	SetOpenStatus(v string) *InitAiAppScanResponseBody
	GetOpenStatus() *string
	SetReadyStatus(v string) *InitAiAppScanResponseBody
	GetReadyStatus() *string
	SetRequestId(v string) *InitAiAppScanResponseBody
	GetRequestId() *string
}

type InitAiAppScanResponseBody struct {
	// The access entry information.
	AuthInfo *InitAiAppScanResponseBodyAuthInfo `json:"AuthInfo,omitempty" xml:"AuthInfo,omitempty" type:"Struct"`
	// The access information.
	AuthInfoConfig map[string]*AuthInfoConfigValue `json:"AuthInfoConfig,omitempty" xml:"AuthInfoConfig,omitempty"`
	// The authorization status.
	//
	// example:
	//
	// enabled
	AuthStatus *string `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// The service activation status.
	//
	// example:
	//
	// enabled
	OpenStatus *string `json:"OpenStatus,omitempty" xml:"OpenStatus,omitempty"`
	// The ready status.
	//
	// example:
	//
	// enabled
	ReadyStatus *string `json:"ReadyStatus,omitempty" xml:"ReadyStatus,omitempty"`
	// The ID assigned by the backend to uniquely identify a request. Used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitAiAppScanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitAiAppScanResponseBody) GoString() string {
	return s.String()
}

func (s *InitAiAppScanResponseBody) GetAuthInfo() *InitAiAppScanResponseBodyAuthInfo {
	return s.AuthInfo
}

func (s *InitAiAppScanResponseBody) GetAuthInfoConfig() map[string]*AuthInfoConfigValue {
	return s.AuthInfoConfig
}

func (s *InitAiAppScanResponseBody) GetAuthStatus() *string {
	return s.AuthStatus
}

func (s *InitAiAppScanResponseBody) GetOpenStatus() *string {
	return s.OpenStatus
}

func (s *InitAiAppScanResponseBody) GetReadyStatus() *string {
	return s.ReadyStatus
}

func (s *InitAiAppScanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitAiAppScanResponseBody) SetAuthInfo(v *InitAiAppScanResponseBodyAuthInfo) *InitAiAppScanResponseBody {
	s.AuthInfo = v
	return s
}

func (s *InitAiAppScanResponseBody) SetAuthInfoConfig(v map[string]*AuthInfoConfigValue) *InitAiAppScanResponseBody {
	s.AuthInfoConfig = v
	return s
}

func (s *InitAiAppScanResponseBody) SetAuthStatus(v string) *InitAiAppScanResponseBody {
	s.AuthStatus = &v
	return s
}

func (s *InitAiAppScanResponseBody) SetOpenStatus(v string) *InitAiAppScanResponseBody {
	s.OpenStatus = &v
	return s
}

func (s *InitAiAppScanResponseBody) SetReadyStatus(v string) *InitAiAppScanResponseBody {
	s.ReadyStatus = &v
	return s
}

func (s *InitAiAppScanResponseBody) SetRequestId(v string) *InitAiAppScanResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitAiAppScanResponseBody) Validate() error {
	if s.AuthInfo != nil {
		if err := s.AuthInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitAiAppScanResponseBodyAuthInfo struct {
	// The credential.
	//
	// example:
	//
	// token-xxx
	AuthToken *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
	// The private domain name.
	//
	// example:
	//
	// https://xxx
	PrivateDomain *string `json:"PrivateDomain,omitempty" xml:"PrivateDomain,omitempty"`
	// The project space.
	//
	// example:
	//
	// proj-xxx
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The public domain name.
	//
	// example:
	//
	// https://xxx
	PublicDomain *string `json:"PublicDomain,omitempty" xml:"PublicDomain,omitempty"`
}

func (s InitAiAppScanResponseBodyAuthInfo) String() string {
	return dara.Prettify(s)
}

func (s InitAiAppScanResponseBodyAuthInfo) GoString() string {
	return s.String()
}

func (s *InitAiAppScanResponseBodyAuthInfo) GetAuthToken() *string {
	return s.AuthToken
}

func (s *InitAiAppScanResponseBodyAuthInfo) GetPrivateDomain() *string {
	return s.PrivateDomain
}

func (s *InitAiAppScanResponseBodyAuthInfo) GetProject() *string {
	return s.Project
}

func (s *InitAiAppScanResponseBodyAuthInfo) GetPublicDomain() *string {
	return s.PublicDomain
}

func (s *InitAiAppScanResponseBodyAuthInfo) SetAuthToken(v string) *InitAiAppScanResponseBodyAuthInfo {
	s.AuthToken = &v
	return s
}

func (s *InitAiAppScanResponseBodyAuthInfo) SetPrivateDomain(v string) *InitAiAppScanResponseBodyAuthInfo {
	s.PrivateDomain = &v
	return s
}

func (s *InitAiAppScanResponseBodyAuthInfo) SetProject(v string) *InitAiAppScanResponseBodyAuthInfo {
	s.Project = &v
	return s
}

func (s *InitAiAppScanResponseBodyAuthInfo) SetPublicDomain(v string) *InitAiAppScanResponseBodyAuthInfo {
	s.PublicDomain = &v
	return s
}

func (s *InitAiAppScanResponseBodyAuthInfo) Validate() error {
	return dara.Validate(s)
}
