// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSampleConsistencyJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteSampleConsistencyJobRequest
	GetInstanceId() *string
}

type DeleteSampleConsistencyJobRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteSampleConsistencyJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSampleConsistencyJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteSampleConsistencyJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteSampleConsistencyJobRequest) SetInstanceId(v string) *DeleteSampleConsistencyJobRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSampleConsistencyJobRequest) Validate() error {
	return dara.Validate(s)
}
