// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNotificationPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *DeleteNotificationPolicyResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteNotificationPolicyResponseBody
	GetRequestId() *string
}

type DeleteNotificationPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the notification policy.
	//
	// example:
	//
	// A5EC8221-08F2-4C95-9AF1-49FD998C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNotificationPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNotificationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNotificationPolicyResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteNotificationPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNotificationPolicyResponseBody) SetIsSuccess(v bool) *DeleteNotificationPolicyResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteNotificationPolicyResponseBody) SetRequestId(v string) *DeleteNotificationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNotificationPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
