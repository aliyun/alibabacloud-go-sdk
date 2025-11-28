// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSFabricConsortiumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListBaaSFabricConsortiumRequest
	GetRegionId() *string
}

type ListBaaSFabricConsortiumRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBaaSFabricConsortiumRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSFabricConsortiumRequest) GoString() string {
	return s.String()
}

func (s *ListBaaSFabricConsortiumRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBaaSFabricConsortiumRequest) SetRegionId(v string) *ListBaaSFabricConsortiumRequest {
	s.RegionId = &v
	return s
}

func (s *ListBaaSFabricConsortiumRequest) Validate() error {
	return dara.Validate(s)
}
