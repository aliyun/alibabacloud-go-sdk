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
	SetResourceDirectoryAccountId(v int64) *AdvanceSecurityEventOperationsRequest
	GetResourceDirectoryAccountId() *int64
	SetResourceOwnerId(v int64) *AdvanceSecurityEventOperationsRequest
	GetResourceOwnerId() *int64
	SetRuleId(v int32) *AdvanceSecurityEventOperationsRequest
	GetRuleId() *int32
}

type AdvanceSecurityEventOperationsRequest struct {
	// The alert name. The EventName and EventType parameters must be specified together. If only one of them is specified, the API returns a 400 error.
	//
	// example:
	//
	// Malicious script code execution
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The type of the alerting event. Valid values:
	//
	// - Abnormal process behavior
	//
	// - Web shell
	//
	// - Unusual logon
	//
	// - Abnormal event
	//
	// - Sensitive file tampering
	//
	// - Malicious process (cloud scan)
	//
	// - Suspicious network connection
	//
	// - Abnormal account
	//
	// - Application intrusion event
	//
	// - Cloud service threat detection
	//
	// - Precise defense
	//
	// - Application whitelist
	//
	// - Persistent backdoor
	//
	// - Web application threat detection
	//
	// - Malicious script
	//
	// - Threat intelligence
	//
	// - Malicious network behavior
	//
	// - Container cluster exception
	//
	// - Web shell (local scan)
	//
	// - Vulnerability exploits
	//
	// - Malicious process (local scan)
	//
	// - Trusted exception
	//
	// - Other
	//
	// For more information about alert types, see [Security alert check items](https://help.aliyun.com/document_detail/68388.html).
	//
	// The EventName and EventType parameters must be specified together. If only one of them is specified, the API returns a 400 error.
	//
	// example:
	//
	// Malicious script
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The member account ID in the resource directory (Alibaba Cloud account).
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	ResourceOwnerId            *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *AdvanceSecurityEventOperationsRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
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

func (s *AdvanceSecurityEventOperationsRequest) SetResourceDirectoryAccountId(v int64) *AdvanceSecurityEventOperationsRequest {
	s.ResourceDirectoryAccountId = &v
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
