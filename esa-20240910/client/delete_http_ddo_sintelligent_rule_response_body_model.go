// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpDDoSIntelligentRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordName(v string) *DeleteHttpDDoSIntelligentRuleResponseBody
	GetRecordName() *string
	SetRequestId(v string) *DeleteHttpDDoSIntelligentRuleResponseBody
	GetRequestId() *string
	SetRuleId(v int64) *DeleteHttpDDoSIntelligentRuleResponseBody
	GetRuleId() *int64
	SetSiteId(v int64) *DeleteHttpDDoSIntelligentRuleResponseBody
	GetSiteId() *int64
}

type DeleteHttpDDoSIntelligentRuleResponseBody struct {
	// example:
	//
	// www.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20757864
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteHttpDDoSIntelligentRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpDDoSIntelligentRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHttpDDoSIntelligentRuleResponseBody) GetRecordName() *string {
	return s.RecordName
}

func (s *DeleteHttpDDoSIntelligentRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHttpDDoSIntelligentRuleResponseBody) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DeleteHttpDDoSIntelligentRuleResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteHttpDDoSIntelligentRuleResponseBody) SetRecordName(v string) *DeleteHttpDDoSIntelligentRuleResponseBody {
	s.RecordName = &v
	return s
}

func (s *DeleteHttpDDoSIntelligentRuleResponseBody) SetRequestId(v string) *DeleteHttpDDoSIntelligentRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHttpDDoSIntelligentRuleResponseBody) SetRuleId(v int64) *DeleteHttpDDoSIntelligentRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *DeleteHttpDDoSIntelligentRuleResponseBody) SetSiteId(v int64) *DeleteHttpDDoSIntelligentRuleResponseBody {
	s.SiteId = &v
	return s
}

func (s *DeleteHttpDDoSIntelligentRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
