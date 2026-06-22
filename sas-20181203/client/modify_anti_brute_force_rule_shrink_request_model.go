// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAntiBruteForceRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultRule(v bool) *ModifyAntiBruteForceRuleShrinkRequest
	GetDefaultRule() *bool
	SetFailCount(v int32) *ModifyAntiBruteForceRuleShrinkRequest
	GetFailCount() *int32
	SetForbiddenTime(v int32) *ModifyAntiBruteForceRuleShrinkRequest
	GetForbiddenTime() *int32
	SetId(v int64) *ModifyAntiBruteForceRuleShrinkRequest
	GetId() *int64
	SetName(v string) *ModifyAntiBruteForceRuleShrinkRequest
	GetName() *string
	SetProtocolTypeShrink(v string) *ModifyAntiBruteForceRuleShrinkRequest
	GetProtocolTypeShrink() *string
	SetResourceOwnerId(v int64) *ModifyAntiBruteForceRuleShrinkRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *ModifyAntiBruteForceRuleShrinkRequest
	GetSourceIp() *string
	SetSpan(v int32) *ModifyAntiBruteForceRuleShrinkRequest
	GetSpan() *int32
	SetUuidList(v []*string) *ModifyAntiBruteForceRuleShrinkRequest
	GetUuidList() []*string
}

type ModifyAntiBruteForceRuleShrinkRequest struct {
	// Specifies whether the defense rule against brute-force attacks is set as the default policy in Settings. Valid values:
	//
	// - **true**: The rule is set as the default policy.
	//
	// - **false**: The rule is not set as the default policy.
	//
	// example:
	//
	// true
	DefaultRule *bool `json:"DefaultRule,omitempty" xml:"DefaultRule,omitempty"`
	// The threshold for the number of logon failures. Valid values:
	//
	// - **2**: 2 times
	//
	// - **3**: 3 times
	//
	// - **4**: 4 times
	//
	// - **5**: 5 times
	//
	// - **10**: 10 times
	//
	// - **50**: 50 times
	//
	// - **80**: 80 times
	//
	// - **100**: 100 times.
	//
	// example:
	//
	// 10
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The duration for which logon is prohibited, in minutes. Valid values:
	//
	// - **5**: 5 minutes
	//
	// - **15**: 15 minutes
	//
	// - **30**: 30 minutes
	//
	// - **60**: 1 hour
	//
	// - **120**: 2 hours
	//
	// - **360**: 6 hours
	//
	// - **720**: 12 hours
	//
	// - **1440**: 24 hours
	//
	// - **10080**: 7 days
	//
	// - **52560000**: permanent (100 years).
	//
	// example:
	//
	// 5
	ForbiddenTime *int32 `json:"ForbiddenTime,omitempty" xml:"ForbiddenTime,omitempty"`
	// The ID of the defense rule against brute-force attacks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65778
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the defense rule against brute-force attacks.
	//
	// example:
	//
	// TestRule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The protocol types that the defense rule against brute-force attacks supports for interception.
	ProtocolTypeShrink *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	ResourceOwnerId    *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The threshold for the period of time during which logon failures are counted, in minutes. Valid values:
	//
	// - **1**: 1 minute
	//
	// - **2**: 2 minutes
	//
	// - **5**: 5 minutes
	//
	// - **10**: 10 minutes
	//
	// - **15**: 15 minutes.
	//
	// example:
	//
	// 1
	Span *int32 `json:"Span,omitempty" xml:"Span,omitempty"`
	// The list of servers to which the defense rule against brute-force attacks applies.
	//
	// example:
	//
	// uuid-13213-dasda
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s ModifyAntiBruteForceRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAntiBruteForceRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) GetDefaultRule() *bool {
	return s.DefaultRule
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) GetFailCount() *int32 {
	return s.FailCount
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) GetForbiddenTime() *int32 {
	return s.ForbiddenTime
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) GetProtocolTypeShrink() *string {
	return s.ProtocolTypeShrink
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) GetSpan() *int32 {
	return s.Span
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) SetDefaultRule(v bool) *ModifyAntiBruteForceRuleShrinkRequest {
	s.DefaultRule = &v
	return s
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) SetFailCount(v int32) *ModifyAntiBruteForceRuleShrinkRequest {
	s.FailCount = &v
	return s
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) SetForbiddenTime(v int32) *ModifyAntiBruteForceRuleShrinkRequest {
	s.ForbiddenTime = &v
	return s
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) SetId(v int64) *ModifyAntiBruteForceRuleShrinkRequest {
	s.Id = &v
	return s
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) SetName(v string) *ModifyAntiBruteForceRuleShrinkRequest {
	s.Name = &v
	return s
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) SetProtocolTypeShrink(v string) *ModifyAntiBruteForceRuleShrinkRequest {
	s.ProtocolTypeShrink = &v
	return s
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) SetResourceOwnerId(v int64) *ModifyAntiBruteForceRuleShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) SetSourceIp(v string) *ModifyAntiBruteForceRuleShrinkRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) SetSpan(v int32) *ModifyAntiBruteForceRuleShrinkRequest {
	s.Span = &v
	return s
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) SetUuidList(v []*string) *ModifyAntiBruteForceRuleShrinkRequest {
	s.UuidList = v
	return s
}

func (s *ModifyAntiBruteForceRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
