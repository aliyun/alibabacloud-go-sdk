// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSampleConsistencyJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetSampleConsistencyJobRequest
	GetInstanceId() *string
}

type GetSampleConsistencyJobRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetSampleConsistencyJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSampleConsistencyJobRequest) GoString() string {
	return s.String()
}

func (s *GetSampleConsistencyJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSampleConsistencyJobRequest) SetInstanceId(v string) *GetSampleConsistencyJobRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSampleConsistencyJobRequest) Validate() error {
	return dara.Validate(s)
}
