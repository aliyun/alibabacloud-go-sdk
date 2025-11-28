// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSAntChainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaaSAntChainConsortiumId(v string) *ListBaaSAntChainRequest
	GetBaaSAntChainConsortiumId() *string
	SetRegionId(v string) *ListBaaSAntChainRequest
	GetRegionId() *string
}

type ListBaaSAntChainRequest struct {
	// This parameter is required.
	BaaSAntChainConsortiumId *string `json:"BaaSAntChainConsortiumId,omitempty" xml:"BaaSAntChainConsortiumId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBaaSAntChainRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSAntChainRequest) GoString() string {
	return s.String()
}

func (s *ListBaaSAntChainRequest) GetBaaSAntChainConsortiumId() *string {
	return s.BaaSAntChainConsortiumId
}

func (s *ListBaaSAntChainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBaaSAntChainRequest) SetBaaSAntChainConsortiumId(v string) *ListBaaSAntChainRequest {
	s.BaaSAntChainConsortiumId = &v
	return s
}

func (s *ListBaaSAntChainRequest) SetRegionId(v string) *ListBaaSAntChainRequest {
	s.RegionId = &v
	return s
}

func (s *ListBaaSAntChainRequest) Validate() error {
	return dara.Validate(s)
}
