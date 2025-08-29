// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteABMetricGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteABMetricGroupRequest
	GetInstanceId() *string
}

type DeleteABMetricGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteABMetricGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteABMetricGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteABMetricGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteABMetricGroupRequest) SetInstanceId(v string) *DeleteABMetricGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteABMetricGroupRequest) Validate() error {
	return dara.Validate(s)
}
