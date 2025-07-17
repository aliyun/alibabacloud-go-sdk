// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNotificationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteNotificationPolicyRequest
	GetId() *int64
}

type DeleteNotificationPolicyRequest struct {
	// The ID of the notification policy.
	//
	// For more information about how to obtain the ID of a notification policy, see [ListNotificationPolicies](https://help.aliyun.com/document_detail/2612375.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteNotificationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNotificationPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteNotificationPolicyRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteNotificationPolicyRequest) SetId(v int64) *DeleteNotificationPolicyRequest {
	s.Id = &v
	return s
}

func (s *DeleteNotificationPolicyRequest) Validate() error {
	return dara.Validate(s)
}
