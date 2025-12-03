// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEaiJupyterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StartEaiJupyterRequest
	GetInstanceId() *string
	SetRegionId(v string) *StartEaiJupyterRequest
	GetRegionId() *string
}

type StartEaiJupyterRequest struct {
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

func (s StartEaiJupyterRequest) String() string {
	return dara.Prettify(s)
}

func (s StartEaiJupyterRequest) GoString() string {
	return s.String()
}

func (s *StartEaiJupyterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartEaiJupyterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartEaiJupyterRequest) SetInstanceId(v string) *StartEaiJupyterRequest {
	s.InstanceId = &v
	return s
}

func (s *StartEaiJupyterRequest) SetRegionId(v string) *StartEaiJupyterRequest {
	s.RegionId = &v
	return s
}

func (s *StartEaiJupyterRequest) Validate() error {
	return dara.Validate(s)
}
