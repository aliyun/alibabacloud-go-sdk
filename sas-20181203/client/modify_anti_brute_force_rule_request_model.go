// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAntiBruteForceRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultRule(v bool) *ModifyAntiBruteForceRuleRequest
	GetDefaultRule() *bool
	SetFailCount(v int32) *ModifyAntiBruteForceRuleRequest
	GetFailCount() *int32
	SetForbiddenTime(v int32) *ModifyAntiBruteForceRuleRequest
	GetForbiddenTime() *int32
	SetId(v int64) *ModifyAntiBruteForceRuleRequest
	GetId() *int64
	SetName(v string) *ModifyAntiBruteForceRuleRequest
	GetName() *string
	SetProtocolType(v *ModifyAntiBruteForceRuleRequestProtocolType) *ModifyAntiBruteForceRuleRequest
	GetProtocolType() *ModifyAntiBruteForceRuleRequestProtocolType
	SetResourceOwnerId(v int64) *ModifyAntiBruteForceRuleRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *ModifyAntiBruteForceRuleRequest
	GetSourceIp() *string
	SetSpan(v int32) *ModifyAntiBruteForceRuleRequest
	GetSpan() *int32
	SetUuidList(v []*string) *ModifyAntiBruteForceRuleRequest
	GetUuidList() []*string
}

type ModifyAntiBruteForceRuleRequest struct {
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
	ProtocolType    *ModifyAntiBruteForceRuleRequestProtocolType `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty" type:"Struct"`
	ResourceOwnerId *int64                                       `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s ModifyAntiBruteForceRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAntiBruteForceRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyAntiBruteForceRuleRequest) GetDefaultRule() *bool {
	return s.DefaultRule
}

func (s *ModifyAntiBruteForceRuleRequest) GetFailCount() *int32 {
	return s.FailCount
}

func (s *ModifyAntiBruteForceRuleRequest) GetForbiddenTime() *int32 {
	return s.ForbiddenTime
}

func (s *ModifyAntiBruteForceRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyAntiBruteForceRuleRequest) GetName() *string {
	return s.Name
}

func (s *ModifyAntiBruteForceRuleRequest) GetProtocolType() *ModifyAntiBruteForceRuleRequestProtocolType {
	return s.ProtocolType
}

func (s *ModifyAntiBruteForceRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAntiBruteForceRuleRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyAntiBruteForceRuleRequest) GetSpan() *int32 {
	return s.Span
}

func (s *ModifyAntiBruteForceRuleRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *ModifyAntiBruteForceRuleRequest) SetDefaultRule(v bool) *ModifyAntiBruteForceRuleRequest {
	s.DefaultRule = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetFailCount(v int32) *ModifyAntiBruteForceRuleRequest {
	s.FailCount = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetForbiddenTime(v int32) *ModifyAntiBruteForceRuleRequest {
	s.ForbiddenTime = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetId(v int64) *ModifyAntiBruteForceRuleRequest {
	s.Id = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetName(v string) *ModifyAntiBruteForceRuleRequest {
	s.Name = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetProtocolType(v *ModifyAntiBruteForceRuleRequestProtocolType) *ModifyAntiBruteForceRuleRequest {
	s.ProtocolType = v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetResourceOwnerId(v int64) *ModifyAntiBruteForceRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetSourceIp(v string) *ModifyAntiBruteForceRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetSpan(v int32) *ModifyAntiBruteForceRuleRequest {
	s.Span = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetUuidList(v []*string) *ModifyAntiBruteForceRuleRequest {
	s.UuidList = v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) Validate() error {
	if s.ProtocolType != nil {
		if err := s.ProtocolType.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAntiBruteForceRuleRequestProtocolType struct {
	// Specifies whether to enable RDP interception. This is enabled by default. Valid values:
	//
	//   - **on**: enabled
	//
	//   - **off**: disabled.
	//
	// example:
	//
	// on
	Rdp *string `json:"Rdp,omitempty" xml:"Rdp,omitempty"`
	// Specifies whether to enable SqlServer interception. This is disabled by default. Valid values:
	//
	//   - **on**: enabled
	//
	//   - **off**: disabled.
	//
	// example:
	//
	// off
	SqlServer *string `json:"SqlServer,omitempty" xml:"SqlServer,omitempty"`
	// Specifies whether to enable SSH interception. This is enabled by default. Valid values:
	//
	//   - **on**: enabled
	//
	//   - **off**: disabled.
	//
	// example:
	//
	// on
	Ssh *string `json:"Ssh,omitempty" xml:"Ssh,omitempty"`
}

func (s ModifyAntiBruteForceRuleRequestProtocolType) String() string {
	return dara.Prettify(s)
}

func (s ModifyAntiBruteForceRuleRequestProtocolType) GoString() string {
	return s.String()
}

func (s *ModifyAntiBruteForceRuleRequestProtocolType) GetRdp() *string {
	return s.Rdp
}

func (s *ModifyAntiBruteForceRuleRequestProtocolType) GetSqlServer() *string {
	return s.SqlServer
}

func (s *ModifyAntiBruteForceRuleRequestProtocolType) GetSsh() *string {
	return s.Ssh
}

func (s *ModifyAntiBruteForceRuleRequestProtocolType) SetRdp(v string) *ModifyAntiBruteForceRuleRequestProtocolType {
	s.Rdp = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequestProtocolType) SetSqlServer(v string) *ModifyAntiBruteForceRuleRequestProtocolType {
	s.SqlServer = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequestProtocolType) SetSsh(v string) *ModifyAntiBruteForceRuleRequestProtocolType {
	s.Ssh = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequestProtocolType) Validate() error {
	return dara.Validate(s)
}
