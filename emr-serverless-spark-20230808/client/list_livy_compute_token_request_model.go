// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivyComputeTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListLivyComputeTokenRequest
	GetRegionId() *string
}

type ListLivyComputeTokenRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListLivyComputeTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLivyComputeTokenRequest) GoString() string {
	return s.String()
}

func (s *ListLivyComputeTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLivyComputeTokenRequest) SetRegionId(v string) *ListLivyComputeTokenRequest {
	s.RegionId = &v
	return s
}

func (s *ListLivyComputeTokenRequest) Validate() error {
	return dara.Validate(s)
}
