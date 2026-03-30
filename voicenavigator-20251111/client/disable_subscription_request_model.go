// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DisableSubscriptionRequest
	GetInstanceId() *string
}

type DisableSubscriptionRequest struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *DisableSubscriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableSubscriptionRequest) SetInstanceId(v string) *DisableSubscriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
