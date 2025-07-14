// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllotDatasetAccelerationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCubeId(v string) *AllotDatasetAccelerationTaskRequest
	GetCubeId() *string
}

type AllotDatasetAccelerationTaskRequest struct {
	// The dataset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
}

func (s AllotDatasetAccelerationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s AllotDatasetAccelerationTaskRequest) GoString() string {
	return s.String()
}

func (s *AllotDatasetAccelerationTaskRequest) GetCubeId() *string {
	return s.CubeId
}

func (s *AllotDatasetAccelerationTaskRequest) SetCubeId(v string) *AllotDatasetAccelerationTaskRequest {
	s.CubeId = &v
	return s
}

func (s *AllotDatasetAccelerationTaskRequest) Validate() error {
	return dara.Validate(s)
}
