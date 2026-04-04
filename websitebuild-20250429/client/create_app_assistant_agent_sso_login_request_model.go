// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppAssistantAgentSsoLoginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *CreateAppAssistantAgentSsoLoginRequest
	GetBizId() *string
	SetPlatformType(v string) *CreateAppAssistantAgentSsoLoginRequest
	GetPlatformType() *string
	SetTargetUrl(v string) *CreateAppAssistantAgentSsoLoginRequest
	GetTargetUrl() *string
}

type CreateAppAssistantAgentSsoLoginRequest struct {
	// example:
	//
	// WD20250821161210000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// all
	PlatformType *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	// example:
	//
	// http://172.16.70.16:9410/metrics
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
}

func (s CreateAppAssistantAgentSsoLoginRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAssistantAgentSsoLoginRequest) GoString() string {
	return s.String()
}

func (s *CreateAppAssistantAgentSsoLoginRequest) GetBizId() *string {
	return s.BizId
}

func (s *CreateAppAssistantAgentSsoLoginRequest) GetPlatformType() *string {
	return s.PlatformType
}

func (s *CreateAppAssistantAgentSsoLoginRequest) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *CreateAppAssistantAgentSsoLoginRequest) SetBizId(v string) *CreateAppAssistantAgentSsoLoginRequest {
	s.BizId = &v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginRequest) SetPlatformType(v string) *CreateAppAssistantAgentSsoLoginRequest {
	s.PlatformType = &v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginRequest) SetTargetUrl(v string) *CreateAppAssistantAgentSsoLoginRequest {
	s.TargetUrl = &v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginRequest) Validate() error {
	return dara.Validate(s)
}
