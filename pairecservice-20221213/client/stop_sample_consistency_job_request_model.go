// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSampleConsistencyJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StopSampleConsistencyJobRequest
	GetInstanceId() *string
}

type StopSampleConsistencyJobRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopSampleConsistencyJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StopSampleConsistencyJobRequest) GoString() string {
	return s.String()
}

func (s *StopSampleConsistencyJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopSampleConsistencyJobRequest) SetInstanceId(v string) *StopSampleConsistencyJobRequest {
	s.InstanceId = &v
	return s
}

func (s *StopSampleConsistencyJobRequest) Validate() error {
	return dara.Validate(s)
}
