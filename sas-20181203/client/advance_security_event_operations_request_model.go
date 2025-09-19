// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAdvanceSecurityEventOperationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventName(v string) *AdvanceSecurityEventOperationsRequest
	GetEventName() *string
	SetEventType(v string) *AdvanceSecurityEventOperationsRequest
	GetEventType() *string
	SetResourceOwnerId(v int64) *AdvanceSecurityEventOperationsRequest
	GetResourceOwnerId() *int64
	SetRuleId(v int32) *AdvanceSecurityEventOperationsRequest
	GetRuleId() *int32
}

type AdvanceSecurityEventOperationsRequest struct {
	// The alert name.
	//
	// example:
	//
	// Execution of malicious script code
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The alert event type. Valid values:
	//
	// 	- Suspicious process
	//
	// 	- Webshell
	//
	// 	- Unusual logon
	//
	// 	- Exception
	//
	// 	- Sensitive file tampering
	//
	// 	- Malicious process (cloud threat detection)
	//
	// 	- Unusual network connection
	//
	// 	- Abnormal account
	//
	// 	- Application intrusion event
	//
	// 	- Cloud threat detection
	//
	// 	- Precision defense
	//
	// 	- Application whitelist
	//
	// 	- Persistent webshell
	//
	// 	- Web application threat detection
	//
	// 	- Malicious script
	//
	// 	- Threat intelligence
	//
	// 	- Malicious network activity
	//
	// 	- Cluster exception
	//
	// 	- Webshell (on-premises threat detection)
	//
	// 	- Vulnerability exploitation
	//
	// 	- Malicious process (on-premises threat detection)
	//
	// 	- Trusted exception
	//
	// 	- Others
	//
	// For more information about alert types, see [Alerts](https://help.aliyun.com/document_detail/68388.html).
	//
	// example:
	//
	// Malicious script
	EventType       *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 123
	RuleId *int32 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s AdvanceSecurityEventOperationsRequest) String() string {
	return dara.Prettify(s)
}

func (s AdvanceSecurityEventOperationsRequest) GoString() string {
	return s.String()
}

func (s *AdvanceSecurityEventOperationsRequest) GetEventName() *string {
	return s.EventName
}

func (s *AdvanceSecurityEventOperationsRequest) GetEventType() *string {
	return s.EventType
}

func (s *AdvanceSecurityEventOperationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AdvanceSecurityEventOperationsRequest) GetRuleId() *int32 {
	return s.RuleId
}

func (s *AdvanceSecurityEventOperationsRequest) SetEventName(v string) *AdvanceSecurityEventOperationsRequest {
	s.EventName = &v
	return s
}

func (s *AdvanceSecurityEventOperationsRequest) SetEventType(v string) *AdvanceSecurityEventOperationsRequest {
	s.EventType = &v
	return s
}

func (s *AdvanceSecurityEventOperationsRequest) SetResourceOwnerId(v int64) *AdvanceSecurityEventOperationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AdvanceSecurityEventOperationsRequest) SetRuleId(v int32) *AdvanceSecurityEventOperationsRequest {
	s.RuleId = &v
	return s
}

func (s *AdvanceSecurityEventOperationsRequest) Validate() error {
	return dara.Validate(s)
}
