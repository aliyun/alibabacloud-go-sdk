// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetABMetricGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetABMetricGroupRequest
	GetInstanceId() *string
}

type GetABMetricGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetABMetricGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetABMetricGroupRequest) GoString() string {
	return s.String()
}

func (s *GetABMetricGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetABMetricGroupRequest) SetInstanceId(v string) *GetABMetricGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *GetABMetricGroupRequest) Validate() error {
	return dara.Validate(s)
}
