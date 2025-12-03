// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopEaiJupyterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StopEaiJupyterRequest
	GetInstanceId() *string
	SetRegionId(v string) *StopEaiJupyterRequest
	GetRegionId() *string
}

type StopEaiJupyterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eais-hze3x2gv9wimdj0k****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopEaiJupyterRequest) String() string {
	return dara.Prettify(s)
}

func (s StopEaiJupyterRequest) GoString() string {
	return s.String()
}

func (s *StopEaiJupyterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopEaiJupyterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopEaiJupyterRequest) SetInstanceId(v string) *StopEaiJupyterRequest {
	s.InstanceId = &v
	return s
}

func (s *StopEaiJupyterRequest) SetRegionId(v string) *StopEaiJupyterRequest {
	s.RegionId = &v
	return s
}

func (s *StopEaiJupyterRequest) Validate() error {
	return dara.Validate(s)
}
