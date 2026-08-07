// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetSubscriptionRequest
	GetInstanceId() *string
}

type GetSubscriptionRequest struct {
	// 实例ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *GetSubscriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSubscriptionRequest) SetInstanceId(v string) *GetSubscriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
