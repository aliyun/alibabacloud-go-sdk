// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHttpDDoSAttackIntelligentProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAiMode(v string) *SetHttpDDoSAttackIntelligentProtectionRequest
	GetAiMode() *string
	SetAiTemplate(v string) *SetHttpDDoSAttackIntelligentProtectionRequest
	GetAiTemplate() *string
	SetSiteId(v int64) *SetHttpDDoSAttackIntelligentProtectionRequest
	GetSiteId() *int64
}

type SetHttpDDoSAttackIntelligentProtectionRequest struct {
	// This parameter is required.
	AiMode *string `json:"AiMode,omitempty" xml:"AiMode,omitempty"`
	// This parameter is required.
	AiTemplate *string `json:"AiTemplate,omitempty" xml:"AiTemplate,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s SetHttpDDoSAttackIntelligentProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetHttpDDoSAttackIntelligentProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackIntelligentProtectionRequest) GetAiMode() *string {
	return s.AiMode
}

func (s *SetHttpDDoSAttackIntelligentProtectionRequest) GetAiTemplate() *string {
	return s.AiTemplate
}

func (s *SetHttpDDoSAttackIntelligentProtectionRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetHttpDDoSAttackIntelligentProtectionRequest) SetAiMode(v string) *SetHttpDDoSAttackIntelligentProtectionRequest {
	s.AiMode = &v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionRequest) SetAiTemplate(v string) *SetHttpDDoSAttackIntelligentProtectionRequest {
	s.AiTemplate = &v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionRequest) SetSiteId(v int64) *SetHttpDDoSAttackIntelligentProtectionRequest {
	s.SiteId = &v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionRequest) Validate() error {
	return dara.Validate(s)
}
