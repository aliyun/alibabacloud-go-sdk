// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRCVClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListRCVClustersRequest
	GetRegionId() *string
}

type ListRCVClustersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRCVClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRCVClustersRequest) GoString() string {
	return s.String()
}

func (s *ListRCVClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRCVClustersRequest) SetRegionId(v string) *ListRCVClustersRequest {
	s.RegionId = &v
	return s
}

func (s *ListRCVClustersRequest) Validate() error {
	return dara.Validate(s)
}
