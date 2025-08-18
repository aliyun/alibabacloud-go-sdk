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
	// The mode of smart HTTP DDoS protection. Valid values:
	//
	// 	- **observe**: alert.
	//
	// 	- **defense**: block.
	//
	// This parameter is required.
	//
	// example:
	//
	// defense
	AiMode *string `json:"AiMode,omitempty" xml:"AiMode,omitempty"`
	// The level of smart HTTP DDoS protection. Valid values:
	//
	// 	- **level0**: very loose.
	//
	// 	- **level30**: loose.
	//
	// 	- **level60**: normal.
	//
	// 	- **level90**: strict.
	//
	// This parameter is required.
	//
	// example:
	//
	// level60
	AiTemplate *string `json:"AiTemplate,omitempty" xml:"AiTemplate,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
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
