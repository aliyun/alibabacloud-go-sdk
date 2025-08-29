// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteABMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteABMetricRequest
	GetInstanceId() *string
}

type DeleteABMetricRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteABMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteABMetricRequest) GoString() string {
	return s.String()
}

func (s *DeleteABMetricRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteABMetricRequest) SetInstanceId(v string) *DeleteABMetricRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteABMetricRequest) Validate() error {
	return dara.Validate(s)
}
