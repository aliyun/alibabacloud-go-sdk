// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareSampleConsistencyJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CompareSampleConsistencyJobRequest
	GetInstanceId() *string
}

type CompareSampleConsistencyJobRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CompareSampleConsistencyJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CompareSampleConsistencyJobRequest) GoString() string {
	return s.String()
}

func (s *CompareSampleConsistencyJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CompareSampleConsistencyJobRequest) SetInstanceId(v string) *CompareSampleConsistencyJobRequest {
	s.InstanceId = &v
	return s
}

func (s *CompareSampleConsistencyJobRequest) Validate() error {
	return dara.Validate(s)
}
