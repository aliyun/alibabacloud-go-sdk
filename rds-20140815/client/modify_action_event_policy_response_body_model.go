// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActionEventPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnableEventLog(v string) *ModifyActionEventPolicyResponseBody
	GetEnableEventLog() *string
	SetRegionId(v string) *ModifyActionEventPolicyResponseBody
	GetRegionId() *string
	SetRequestId(v string) *ModifyActionEventPolicyResponseBody
	GetRequestId() *string
}

type ModifyActionEventPolicyResponseBody struct {
	// Indicates whether the event history feature is enabled.
	//
	// example:
	//
	// True
	EnableEventLog *string `json:"EnableEventLog,omitempty" xml:"EnableEventLog,omitempty"`
	// The ID of the region for which the event history feature is enabled or disabled.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BAC0952C-0EB3-4DE7-A567-B83269BFE43F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyActionEventPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyActionEventPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyActionEventPolicyResponseBody) GetEnableEventLog() *string {
	return s.EnableEventLog
}

func (s *ModifyActionEventPolicyResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyActionEventPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyActionEventPolicyResponseBody) SetEnableEventLog(v string) *ModifyActionEventPolicyResponseBody {
	s.EnableEventLog = &v
	return s
}

func (s *ModifyActionEventPolicyResponseBody) SetRegionId(v string) *ModifyActionEventPolicyResponseBody {
	s.RegionId = &v
	return s
}

func (s *ModifyActionEventPolicyResponseBody) SetRequestId(v string) *ModifyActionEventPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyActionEventPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
