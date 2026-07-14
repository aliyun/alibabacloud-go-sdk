// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableNotifyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableNotifyPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableNotifyPolicyResponseBody
	GetSuccess() *bool
	SetUuid(v string) *DisableNotifyPolicyResponseBody
	GetUuid() *string
}

type DisableNotifyPolicyResponseBody struct {
	// The unique ID of the request. Used for troubleshooting and ticket tracking.
	//
	// example:
	//
	// 70D52620-2609-1802-9788-6BC592C83F03
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the operation was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The uuid of the notification policy that was operated on.
	//
	// example:
	//
	// 04779a183add4f2ca06ab440f16cc580
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s DisableNotifyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableNotifyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DisableNotifyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableNotifyPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableNotifyPolicyResponseBody) GetUuid() *string {
	return s.Uuid
}

func (s *DisableNotifyPolicyResponseBody) SetRequestId(v string) *DisableNotifyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableNotifyPolicyResponseBody) SetSuccess(v bool) *DisableNotifyPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DisableNotifyPolicyResponseBody) SetUuid(v string) *DisableNotifyPolicyResponseBody {
	s.Uuid = &v
	return s
}

func (s *DisableNotifyPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
