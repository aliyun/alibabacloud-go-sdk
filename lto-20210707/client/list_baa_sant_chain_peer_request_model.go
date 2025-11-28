// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSAntChainPeerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaaSAntChainChainId(v string) *ListBaaSAntChainPeerRequest
	GetBaaSAntChainChainId() *string
	SetBaaSAntChainConsortiumId(v string) *ListBaaSAntChainPeerRequest
	GetBaaSAntChainConsortiumId() *string
	SetRegionId(v string) *ListBaaSAntChainPeerRequest
	GetRegionId() *string
}

type ListBaaSAntChainPeerRequest struct {
	// This parameter is required.
	BaaSAntChainChainId *string `json:"BaaSAntChainChainId,omitempty" xml:"BaaSAntChainChainId,omitempty"`
	// This parameter is required.
	BaaSAntChainConsortiumId *string `json:"BaaSAntChainConsortiumId,omitempty" xml:"BaaSAntChainConsortiumId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBaaSAntChainPeerRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSAntChainPeerRequest) GoString() string {
	return s.String()
}

func (s *ListBaaSAntChainPeerRequest) GetBaaSAntChainChainId() *string {
	return s.BaaSAntChainChainId
}

func (s *ListBaaSAntChainPeerRequest) GetBaaSAntChainConsortiumId() *string {
	return s.BaaSAntChainConsortiumId
}

func (s *ListBaaSAntChainPeerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBaaSAntChainPeerRequest) SetBaaSAntChainChainId(v string) *ListBaaSAntChainPeerRequest {
	s.BaaSAntChainChainId = &v
	return s
}

func (s *ListBaaSAntChainPeerRequest) SetBaaSAntChainConsortiumId(v string) *ListBaaSAntChainPeerRequest {
	s.BaaSAntChainConsortiumId = &v
	return s
}

func (s *ListBaaSAntChainPeerRequest) SetRegionId(v string) *ListBaaSAntChainPeerRequest {
	s.RegionId = &v
	return s
}

func (s *ListBaaSAntChainPeerRequest) Validate() error {
	return dara.Validate(s)
}
