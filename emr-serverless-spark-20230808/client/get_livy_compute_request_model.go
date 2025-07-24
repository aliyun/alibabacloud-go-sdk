// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLivyComputeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetLivyComputeRequest
	GetRegionId() *string
}

type GetLivyComputeRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetLivyComputeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLivyComputeRequest) GoString() string {
	return s.String()
}

func (s *GetLivyComputeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLivyComputeRequest) SetRegionId(v string) *GetLivyComputeRequest {
	s.RegionId = &v
	return s
}

func (s *GetLivyComputeRequest) Validate() error {
	return dara.Validate(s)
}
