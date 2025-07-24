// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLivyComputeTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetLivyComputeTokenRequest
	GetRegionId() *string
}

type GetLivyComputeTokenRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetLivyComputeTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLivyComputeTokenRequest) GoString() string {
	return s.String()
}

func (s *GetLivyComputeTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLivyComputeTokenRequest) SetRegionId(v string) *GetLivyComputeTokenRequest {
	s.RegionId = &v
	return s
}

func (s *GetLivyComputeTokenRequest) Validate() error {
	return dara.Validate(s)
}
