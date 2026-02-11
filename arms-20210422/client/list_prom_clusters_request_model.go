// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListPromClustersRequest
	GetRegionId() *string
}

type ListPromClustersRequest struct {
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPromClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPromClustersRequest) GoString() string {
	return s.String()
}

func (s *ListPromClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPromClustersRequest) SetRegionId(v string) *ListPromClustersRequest {
	s.RegionId = &v
	return s
}

func (s *ListPromClustersRequest) Validate() error {
	return dara.Validate(s)
}
