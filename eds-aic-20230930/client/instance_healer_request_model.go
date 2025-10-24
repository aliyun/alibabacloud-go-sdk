// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceHealerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdList(v []*string) *InstanceHealerRequest
	GetInstanceIdList() []*string
	SetStrategy(v string) *InstanceHealerRequest
	GetStrategy() *string
	SetTimeout(v int64) *InstanceHealerRequest
	GetTimeout() *int64
}

type InstanceHealerRequest struct {
	// This parameter is required.
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// Clean
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// example:
	//
	// 30
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s InstanceHealerRequest) String() string {
	return dara.Prettify(s)
}

func (s InstanceHealerRequest) GoString() string {
	return s.String()
}

func (s *InstanceHealerRequest) GetInstanceIdList() []*string {
	return s.InstanceIdList
}

func (s *InstanceHealerRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *InstanceHealerRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *InstanceHealerRequest) SetInstanceIdList(v []*string) *InstanceHealerRequest {
	s.InstanceIdList = v
	return s
}

func (s *InstanceHealerRequest) SetStrategy(v string) *InstanceHealerRequest {
	s.Strategy = &v
	return s
}

func (s *InstanceHealerRequest) SetTimeout(v int64) *InstanceHealerRequest {
	s.Timeout = &v
	return s
}

func (s *InstanceHealerRequest) Validate() error {
	return dara.Validate(s)
}
