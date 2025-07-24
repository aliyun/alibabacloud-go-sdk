// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLivyComputeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *StartLivyComputeRequest
	GetRegionId() *string
}

type StartLivyComputeRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s StartLivyComputeRequest) String() string {
	return dara.Prettify(s)
}

func (s StartLivyComputeRequest) GoString() string {
	return s.String()
}

func (s *StartLivyComputeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartLivyComputeRequest) SetRegionId(v string) *StartLivyComputeRequest {
	s.RegionId = &v
	return s
}

func (s *StartLivyComputeRequest) Validate() error {
	return dara.Validate(s)
}
