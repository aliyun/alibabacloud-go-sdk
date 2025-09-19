// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAntiBruteForceRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultRule(v bool) *CreateAntiBruteForceRuleShrinkRequest
	GetDefaultRule() *bool
	SetFailCount(v int32) *CreateAntiBruteForceRuleShrinkRequest
	GetFailCount() *int32
	SetForbiddenTime(v int32) *CreateAntiBruteForceRuleShrinkRequest
	GetForbiddenTime() *int32
	SetName(v string) *CreateAntiBruteForceRuleShrinkRequest
	GetName() *string
	SetProtocolTypeShrink(v string) *CreateAntiBruteForceRuleShrinkRequest
	GetProtocolTypeShrink() *string
	SetResourceOwnerId(v int64) *CreateAntiBruteForceRuleShrinkRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *CreateAntiBruteForceRuleShrinkRequest
	GetSourceIp() *string
	SetSpan(v int32) *CreateAntiBruteForceRuleShrinkRequest
	GetSpan() *int32
	SetUuidList(v []*string) *CreateAntiBruteForceRuleShrinkRequest
	GetUuidList() []*string
}

type CreateAntiBruteForceRuleShrinkRequest struct {
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
	ProtocolTypeShrink *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	ResourceOwnerId    *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s CreateAntiBruteForceRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAntiBruteForceRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAntiBruteForceRuleShrinkRequest) GetDefaultRule() *bool {
	return s.DefaultRule
}

func (s *CreateAntiBruteForceRuleShrinkRequest) GetFailCount() *int32 {
	return s.FailCount
}

func (s *CreateAntiBruteForceRuleShrinkRequest) GetForbiddenTime() *int32 {
	return s.ForbiddenTime
}

func (s *CreateAntiBruteForceRuleShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateAntiBruteForceRuleShrinkRequest) GetProtocolTypeShrink() *string {
	return s.ProtocolTypeShrink
}

func (s *CreateAntiBruteForceRuleShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAntiBruteForceRuleShrinkRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateAntiBruteForceRuleShrinkRequest) GetSpan() *int32 {
	return s.Span
}

func (s *CreateAntiBruteForceRuleShrinkRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *CreateAntiBruteForceRuleShrinkRequest) SetDefaultRule(v bool) *CreateAntiBruteForceRuleShrinkRequest {
	s.DefaultRule = &v
	return s
}

func (s *CreateAntiBruteForceRuleShrinkRequest) SetFailCount(v int32) *CreateAntiBruteForceRuleShrinkRequest {
	s.FailCount = &v
	return s
}

func (s *CreateAntiBruteForceRuleShrinkRequest) SetForbiddenTime(v int32) *CreateAntiBruteForceRuleShrinkRequest {
	s.ForbiddenTime = &v
	return s
}

func (s *CreateAntiBruteForceRuleShrinkRequest) SetName(v string) *CreateAntiBruteForceRuleShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateAntiBruteForceRuleShrinkRequest) SetProtocolTypeShrink(v string) *CreateAntiBruteForceRuleShrinkRequest {
	s.ProtocolTypeShrink = &v
	return s
}

func (s *CreateAntiBruteForceRuleShrinkRequest) SetResourceOwnerId(v int64) *CreateAntiBruteForceRuleShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAntiBruteForceRuleShrinkRequest) SetSourceIp(v string) *CreateAntiBruteForceRuleShrinkRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateAntiBruteForceRuleShrinkRequest) SetSpan(v int32) *CreateAntiBruteForceRuleShrinkRequest {
	s.Span = &v
	return s
}

func (s *CreateAntiBruteForceRuleShrinkRequest) SetUuidList(v []*string) *CreateAntiBruteForceRuleShrinkRequest {
	s.UuidList = v
	return s
}

func (s *CreateAntiBruteForceRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
