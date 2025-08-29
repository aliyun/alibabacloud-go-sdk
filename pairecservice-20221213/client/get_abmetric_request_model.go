// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetABMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetABMetricRequest
	GetInstanceId() *string
}

type GetABMetricRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetABMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s GetABMetricRequest) GoString() string {
	return s.String()
}

func (s *GetABMetricRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetABMetricRequest) SetInstanceId(v string) *GetABMetricRequest {
	s.InstanceId = &v
	return s
}

func (s *GetABMetricRequest) Validate() error {
	return dara.Validate(s)
}
