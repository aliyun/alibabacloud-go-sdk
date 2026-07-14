// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotifyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNotifyPolicy(v *NotifyPolicy) *GetNotifyPolicyResponseBody
	GetNotifyPolicy() *NotifyPolicy
	SetRequestId(v string) *GetNotifyPolicyResponseBody
	GetRequestId() *string
}

type GetNotifyPolicyResponseBody struct {
	// The notification policy object details, including the policy UUID, name, description, enabled status, and sub-entities such as notification policies (noise reduction, notification routing, and channels), subscriptions (event filtering, cross-workspace routing, and legacy product event subscriptions), and response plans (escalation, repeated notifications, automatic recovery, and action integration).
	NotifyPolicy *NotifyPolicy `json:"notifyPolicy,omitempty" xml:"notifyPolicy,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-A01D6CC3F4B8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetNotifyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNotifyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetNotifyPolicyResponseBody) GetNotifyPolicy() *NotifyPolicy {
	return s.NotifyPolicy
}

func (s *GetNotifyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNotifyPolicyResponseBody) SetNotifyPolicy(v *NotifyPolicy) *GetNotifyPolicyResponseBody {
	s.NotifyPolicy = v
	return s
}

func (s *GetNotifyPolicyResponseBody) SetRequestId(v string) *GetNotifyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNotifyPolicyResponseBody) Validate() error {
	if s.NotifyPolicy != nil {
		if err := s.NotifyPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
