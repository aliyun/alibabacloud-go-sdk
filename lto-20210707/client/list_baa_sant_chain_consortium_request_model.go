// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSAntChainConsortiumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListBaaSAntChainConsortiumRequest
	GetRegionId() *string
}

type ListBaaSAntChainConsortiumRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBaaSAntChainConsortiumRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSAntChainConsortiumRequest) GoString() string {
	return s.String()
}

func (s *ListBaaSAntChainConsortiumRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBaaSAntChainConsortiumRequest) SetRegionId(v string) *ListBaaSAntChainConsortiumRequest {
	s.RegionId = &v
	return s
}

func (s *ListBaaSAntChainConsortiumRequest) Validate() error {
	return dara.Validate(s)
}
