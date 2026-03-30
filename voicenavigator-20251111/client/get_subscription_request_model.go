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
	// example:
	//
	// 19ac2375-53e3-477f-abe9-6cd334227981
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
