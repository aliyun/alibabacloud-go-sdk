// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNotifyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNotifyPolicy(v *NotifyPolicy) *UpdateNotifyPolicyResponseBody
	GetNotifyPolicy() *NotifyPolicy
	SetRequestId(v string) *UpdateNotifyPolicyResponseBody
	GetRequestId() *string
}

type UpdateNotifyPolicyResponseBody struct {
	// The details of the notification policy object, including the policy UUID, name, description, and enabling status, as well as sub-entities such as the notification policy (group merging, notification routing, and channels), subscriptions (event filtering, cross-workspace routing, and legacy product event subscriptions), and response plans (escalation, repeated notifications, automatic recovery, and action integration).
	NotifyPolicy *NotifyPolicy `json:"notifyPolicy,omitempty" xml:"notifyPolicy,omitempty"`
	// The unique ID of the request, which is used for troubleshooting and ticket tracking.
	//
	// example:
	//
	// 70D52620-2609-1802-9788-6BC592C83F03
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateNotifyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNotifyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNotifyPolicyResponseBody) GetNotifyPolicy() *NotifyPolicy {
	return s.NotifyPolicy
}

func (s *UpdateNotifyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNotifyPolicyResponseBody) SetNotifyPolicy(v *NotifyPolicy) *UpdateNotifyPolicyResponseBody {
	s.NotifyPolicy = v
	return s
}

func (s *UpdateNotifyPolicyResponseBody) SetRequestId(v string) *UpdateNotifyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNotifyPolicyResponseBody) Validate() error {
	if s.NotifyPolicy != nil {
		if err := s.NotifyPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
