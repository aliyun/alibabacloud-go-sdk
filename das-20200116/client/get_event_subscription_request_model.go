// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetEventSubscriptionRequest
	GetInstanceId() *string
}

type GetEventSubscriptionRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetEventSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEventSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *GetEventSubscriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEventSubscriptionRequest) SetInstanceId(v string) *GetEventSubscriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetEventSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
