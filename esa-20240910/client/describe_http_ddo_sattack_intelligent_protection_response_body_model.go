// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHttpDDoSAttackIntelligentProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAiMode(v string) *DescribeHttpDDoSAttackIntelligentProtectionResponseBody
	GetAiMode() *string
	SetAiTemplate(v string) *DescribeHttpDDoSAttackIntelligentProtectionResponseBody
	GetAiTemplate() *string
	SetRequestId(v string) *DescribeHttpDDoSAttackIntelligentProtectionResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *DescribeHttpDDoSAttackIntelligentProtectionResponseBody
	GetSiteId() *int64
}

type DescribeHttpDDoSAttackIntelligentProtectionResponseBody struct {
	// The mode of smart HTTP DDoS protection. Valid values:
	//
	// 	- **observe**: alert.
	//
	// 	- **defense**: block.
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
	// example:
	//
	// level60
	AiTemplate *string `json:"AiTemplate,omitempty" xml:"AiTemplate,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-3C82-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The website ID.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DescribeHttpDDoSAttackIntelligentProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSAttackIntelligentProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponseBody) GetAiMode() *string {
	return s.AiMode
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponseBody) GetAiTemplate() *string {
	return s.AiTemplate
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponseBody) SetAiMode(v string) *DescribeHttpDDoSAttackIntelligentProtectionResponseBody {
	s.AiMode = &v
	return s
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponseBody) SetAiTemplate(v string) *DescribeHttpDDoSAttackIntelligentProtectionResponseBody {
	s.AiTemplate = &v
	return s
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponseBody) SetRequestId(v string) *DescribeHttpDDoSAttackIntelligentProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponseBody) SetSiteId(v int64) *DescribeHttpDDoSAttackIntelligentProtectionResponseBody {
	s.SiteId = &v
	return s
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
