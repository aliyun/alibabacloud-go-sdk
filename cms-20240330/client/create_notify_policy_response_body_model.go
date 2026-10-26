// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNotifyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNotifyPolicy(v *NotifyPolicy) *CreateNotifyPolicyResponseBody
	GetNotifyPolicy() *NotifyPolicy
	SetRequestId(v string) *CreateNotifyPolicyResponseBody
	GetRequestId() *string
}

type CreateNotifyPolicyResponseBody struct {
	// The details of the notification policy object, including the policy UUID, name, description, and enabled status, as well as sub-entities such as the notification strategy (grouping and merging, notification routing, and channels), subscription (event filtering, cross-workspace routing, and legacy product event subscription), and response plan (escalation, repeated notifications, automatic recovery, and action integration).
	NotifyPolicy *NotifyPolicy `json:"notifyPolicy,omitempty" xml:"notifyPolicy,omitempty"`
	// The unique ID of the request. This ID is used for troubleshooting and ticket tracking.
	//
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-A01D6CC3F4B8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateNotifyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNotifyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNotifyPolicyResponseBody) GetNotifyPolicy() *NotifyPolicy {
	return s.NotifyPolicy
}

func (s *CreateNotifyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNotifyPolicyResponseBody) SetNotifyPolicy(v *NotifyPolicy) *CreateNotifyPolicyResponseBody {
	s.NotifyPolicy = v
	return s
}

func (s *CreateNotifyPolicyResponseBody) SetRequestId(v string) *CreateNotifyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNotifyPolicyResponseBody) Validate() error {
	if s.NotifyPolicy != nil {
		if err := s.NotifyPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
