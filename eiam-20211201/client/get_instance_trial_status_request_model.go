// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceTrialStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceTrialStatusRequest
	GetInstanceId() *string
}

type GetInstanceTrialStatusRequest struct {
	// IDaaS EIAM的实例id
	//
	// This parameter is required.
	//
	// example:
	//
	// eiam-111ccc1111
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceTrialStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceTrialStatusRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceTrialStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceTrialStatusRequest) SetInstanceId(v string) *GetInstanceTrialStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceTrialStatusRequest) Validate() error {
	return dara.Validate(s)
}
