// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLastAccelerationEngineJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCubeId(v string) *QueryLastAccelerationEngineJobRequest
	GetCubeId() *string
}

type QueryLastAccelerationEngineJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// d23d30c0-****-6c92bf668356
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
}

func (s QueryLastAccelerationEngineJobRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLastAccelerationEngineJobRequest) GoString() string {
	return s.String()
}

func (s *QueryLastAccelerationEngineJobRequest) GetCubeId() *string {
	return s.CubeId
}

func (s *QueryLastAccelerationEngineJobRequest) SetCubeId(v string) *QueryLastAccelerationEngineJobRequest {
	s.CubeId = &v
	return s
}

func (s *QueryLastAccelerationEngineJobRequest) Validate() error {
	return dara.Validate(s)
}
