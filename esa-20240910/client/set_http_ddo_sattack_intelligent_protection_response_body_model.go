// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHttpDDoSAttackIntelligentProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAiMode(v string) *SetHttpDDoSAttackIntelligentProtectionResponseBody
	GetAiMode() *string
	SetAiTemplate(v string) *SetHttpDDoSAttackIntelligentProtectionResponseBody
	GetAiTemplate() *string
	SetRequestId(v string) *SetHttpDDoSAttackIntelligentProtectionResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *SetHttpDDoSAttackIntelligentProtectionResponseBody
	GetSiteId() *int64
}

type SetHttpDDoSAttackIntelligentProtectionResponseBody struct {
	// The HTTP DDoS intelligent protection pattern. Valid values:
	//
	// - **observe**: Warning.
	//
	// - **defense**: Defense.
	//
	// example:
	//
	// defense
	AiMode *string `json:"AiMode,omitempty" xml:"AiMode,omitempty"`
	// The HTTP DDoS intelligent protection level. Valid values:
	//
	// - **level0**: Very loose.
	//
	// - **level30**: Loose.
	//
	// - **level60**: Normal.
	//
	// - **level90**: Strict.
	//
	// example:
	//
	// level60
	AiTemplate *string `json:"AiTemplate,omitempty" xml:"AiTemplate,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 156A6B-677B1A-4297B7-9187B7-2B44792
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the site.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s SetHttpDDoSAttackIntelligentProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetHttpDDoSAttackIntelligentProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponseBody) GetAiMode() *string {
	return s.AiMode
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponseBody) GetAiTemplate() *string {
	return s.AiTemplate
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponseBody) SetAiMode(v string) *SetHttpDDoSAttackIntelligentProtectionResponseBody {
	s.AiMode = &v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponseBody) SetAiTemplate(v string) *SetHttpDDoSAttackIntelligentProtectionResponseBody {
	s.AiTemplate = &v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponseBody) SetRequestId(v string) *SetHttpDDoSAttackIntelligentProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponseBody) SetSiteId(v int64) *SetHttpDDoSAttackIntelligentProtectionResponseBody {
	s.SiteId = &v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
