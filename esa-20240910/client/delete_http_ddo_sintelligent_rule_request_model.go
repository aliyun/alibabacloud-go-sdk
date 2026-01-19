// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpDDoSIntelligentRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordName(v string) *DeleteHttpDDoSIntelligentRuleRequest
	GetRecordName() *string
	SetRuleId(v int64) *DeleteHttpDDoSIntelligentRuleRequest
	GetRuleId() *int64
	SetSiteId(v int64) *DeleteHttpDDoSIntelligentRuleRequest
	GetSiteId() *int64
}

type DeleteHttpDDoSIntelligentRuleRequest struct {
	// Record name.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// Rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20928021
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteHttpDDoSIntelligentRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpDDoSIntelligentRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteHttpDDoSIntelligentRuleRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *DeleteHttpDDoSIntelligentRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DeleteHttpDDoSIntelligentRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteHttpDDoSIntelligentRuleRequest) SetRecordName(v string) *DeleteHttpDDoSIntelligentRuleRequest {
	s.RecordName = &v
	return s
}

func (s *DeleteHttpDDoSIntelligentRuleRequest) SetRuleId(v int64) *DeleteHttpDDoSIntelligentRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteHttpDDoSIntelligentRuleRequest) SetSiteId(v int64) *DeleteHttpDDoSIntelligentRuleRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteHttpDDoSIntelligentRuleRequest) Validate() error {
	return dara.Validate(s)
}
