// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNotifyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNotifyPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteNotifyPolicyResponseBody
	GetSuccess() *bool
	SetUuid(v string) *DeleteNotifyPolicyResponseBody
	GetUuid() *string
}

type DeleteNotifyPolicyResponseBody struct {
	// The unique ID of the request. Used for troubleshooting and ticket tracking.
	//
	// example:
	//
	// A1234567-1234-1234-1234-123456789012
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the operation is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The uuid of the notification policy that was operated on.
	//
	// example:
	//
	// np-12345678-1234-1234-1234-123456789012
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s DeleteNotifyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNotifyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNotifyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNotifyPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteNotifyPolicyResponseBody) GetUuid() *string {
	return s.Uuid
}

func (s *DeleteNotifyPolicyResponseBody) SetRequestId(v string) *DeleteNotifyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNotifyPolicyResponseBody) SetSuccess(v bool) *DeleteNotifyPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteNotifyPolicyResponseBody) SetUuid(v string) *DeleteNotifyPolicyResponseBody {
	s.Uuid = &v
	return s
}

func (s *DeleteNotifyPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
