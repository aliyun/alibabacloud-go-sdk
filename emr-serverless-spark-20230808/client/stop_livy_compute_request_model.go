// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLivyComputeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *StopLivyComputeRequest
	GetRegionId() *string
}

type StopLivyComputeRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s StopLivyComputeRequest) String() string {
	return dara.Prettify(s)
}

func (s StopLivyComputeRequest) GoString() string {
	return s.String()
}

func (s *StopLivyComputeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopLivyComputeRequest) SetRegionId(v string) *StopLivyComputeRequest {
	s.RegionId = &v
	return s
}

func (s *StopLivyComputeRequest) Validate() error {
	return dara.Validate(s)
}
