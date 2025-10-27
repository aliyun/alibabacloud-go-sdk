// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAntiBruteForceRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultRule(v bool) *CreateAntiBruteForceRuleRequest
	GetDefaultRule() *bool
	SetFailCount(v int32) *CreateAntiBruteForceRuleRequest
	GetFailCount() *int32
	SetForbiddenTime(v int32) *CreateAntiBruteForceRuleRequest
	GetForbiddenTime() *int32
	SetName(v string) *CreateAntiBruteForceRuleRequest
	GetName() *string
	SetProtocolType(v *CreateAntiBruteForceRuleRequestProtocolType) *CreateAntiBruteForceRuleRequest
	GetProtocolType() *CreateAntiBruteForceRuleRequestProtocolType
	SetResourceOwnerId(v int64) *CreateAntiBruteForceRuleRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *CreateAntiBruteForceRuleRequest
	GetSourceIp() *string
	SetSpan(v int32) *CreateAntiBruteForceRuleRequest
	GetSpan() *int32
	SetUuidList(v []*string) *CreateAntiBruteForceRuleRequest
	GetUuidList() []*string
}

type CreateAntiBruteForceRuleRequest struct {
	// Specifies whether to set the defense rule as the default rule. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	DefaultRule *bool `json:"DefaultRule,omitempty" xml:"DefaultRule,omitempty"`
	// The maximum number of failed logon attempts from an account. Valid values: 2, 3, 4, 5, 10, 50, 80, and 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The period of time during which logons from an account are not allowed. Unit: minutes. Valid values:
	//
	// 	- **5**
	//
	// 	- **15**
	//
	// 	- **30**
	//
	// 	- **60**
	//
	// 	- **120**
	//
	// 	- **360**
	//
	// 	- **720**
	//
	// 	- **1440**
	//
	// 	- **10080**
	//
	// 	- **52560000**: permanent
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	ForbiddenTime *int32 `json:"ForbiddenTime,omitempty" xml:"ForbiddenTime,omitempty"`
	// The name of the defense rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestAntiBruteForceRule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The types of protocols supported for interception by the brute force attack rule creation.
	ProtocolType    *CreateAntiBruteForceRuleRequestProtocolType `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty" type:"Struct"`
	ResourceOwnerId *int64                                       `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The maximum period of time during which failed logon attempts from an account can occur. Unit: minutes. Valid values:
	//
	// 	- **1**
	//
	// 	- **2**
	//
	// 	- **5**
	//
	// 	- **10**
	//
	// 	- **15**
	//
	// >  To configure a defense rule, you must specify the Span, FailCount, and ForbiddenTime parameters. If the number of failed logon attempts from an account within the minutes specified by Span exceeds the value specified by FailCount, the account cannot be used for logons within the minutes specified by ForbiddenTime.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Span *int32 `json:"Span,omitempty" xml:"Span,omitempty"`
	// The UUIDs of the servers to which you want to apply the defense rule.
	//
	// This parameter is required.
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s CreateAntiBruteForceRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAntiBruteForceRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAntiBruteForceRuleRequest) GetDefaultRule() *bool {
	return s.DefaultRule
}

func (s *CreateAntiBruteForceRuleRequest) GetFailCount() *int32 {
	return s.FailCount
}

func (s *CreateAntiBruteForceRuleRequest) GetForbiddenTime() *int32 {
	return s.ForbiddenTime
}

func (s *CreateAntiBruteForceRuleRequest) GetName() *string {
	return s.Name
}

func (s *CreateAntiBruteForceRuleRequest) GetProtocolType() *CreateAntiBruteForceRuleRequestProtocolType {
	return s.ProtocolType
}

func (s *CreateAntiBruteForceRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAntiBruteForceRuleRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateAntiBruteForceRuleRequest) GetSpan() *int32 {
	return s.Span
}

func (s *CreateAntiBruteForceRuleRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *CreateAntiBruteForceRuleRequest) SetDefaultRule(v bool) *CreateAntiBruteForceRuleRequest {
	s.DefaultRule = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) SetFailCount(v int32) *CreateAntiBruteForceRuleRequest {
	s.FailCount = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) SetForbiddenTime(v int32) *CreateAntiBruteForceRuleRequest {
	s.ForbiddenTime = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) SetName(v string) *CreateAntiBruteForceRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) SetProtocolType(v *CreateAntiBruteForceRuleRequestProtocolType) *CreateAntiBruteForceRuleRequest {
	s.ProtocolType = v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) SetResourceOwnerId(v int64) *CreateAntiBruteForceRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) SetSourceIp(v string) *CreateAntiBruteForceRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) SetSpan(v int32) *CreateAntiBruteForceRuleRequest {
	s.Span = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) SetUuidList(v []*string) *CreateAntiBruteForceRuleRequest {
	s.UuidList = v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) Validate() error {
	if s.ProtocolType != nil {
		if err := s.ProtocolType.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAntiBruteForceRuleRequestProtocolType struct {
	// Whether to enable RDP interception, default is on. Values:
	//
	// - **on**: Enable
	//
	// - **off**: Disable
	//
	// example:
	//
	// on
	Rdp *string `json:"Rdp,omitempty" xml:"Rdp,omitempty"`
	// Whether to enable the SqlServer interception method, default is off. Values:
	//
	// - **on**: Enable
	//
	// - **off**: Disable
	//
	// example:
	//
	// off
	SqlServer *string `json:"SqlServer,omitempty" xml:"SqlServer,omitempty"`
	// Whether to enable SSH interception, default is on. Values:
	//
	// - **on**: Enable
	//
	// - **off**: Disable
	//
	// example:
	//
	// on
	Ssh *string `json:"Ssh,omitempty" xml:"Ssh,omitempty"`
}

func (s CreateAntiBruteForceRuleRequestProtocolType) String() string {
	return dara.Prettify(s)
}

func (s CreateAntiBruteForceRuleRequestProtocolType) GoString() string {
	return s.String()
}

func (s *CreateAntiBruteForceRuleRequestProtocolType) GetRdp() *string {
	return s.Rdp
}

func (s *CreateAntiBruteForceRuleRequestProtocolType) GetSqlServer() *string {
	return s.SqlServer
}

func (s *CreateAntiBruteForceRuleRequestProtocolType) GetSsh() *string {
	return s.Ssh
}

func (s *CreateAntiBruteForceRuleRequestProtocolType) SetRdp(v string) *CreateAntiBruteForceRuleRequestProtocolType {
	s.Rdp = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequestProtocolType) SetSqlServer(v string) *CreateAntiBruteForceRuleRequestProtocolType {
	s.SqlServer = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequestProtocolType) SetSsh(v string) *CreateAntiBruteForceRuleRequestProtocolType {
	s.Ssh = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequestProtocolType) Validate() error {
	return dara.Validate(s)
}
