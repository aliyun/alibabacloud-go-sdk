// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllSystemContractRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockChainType(v string) *ListAllSystemContractRequest
	GetBlockChainType() *string
	SetRegionId(v string) *ListAllSystemContractRequest
	GetRegionId() *string
}

type ListAllSystemContractRequest struct {
	// This parameter is required.
	BlockChainType *string `json:"BlockChainType,omitempty" xml:"BlockChainType,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAllSystemContractRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllSystemContractRequest) GoString() string {
	return s.String()
}

func (s *ListAllSystemContractRequest) GetBlockChainType() *string {
	return s.BlockChainType
}

func (s *ListAllSystemContractRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAllSystemContractRequest) SetBlockChainType(v string) *ListAllSystemContractRequest {
	s.BlockChainType = &v
	return s
}

func (s *ListAllSystemContractRequest) SetRegionId(v string) *ListAllSystemContractRequest {
	s.RegionId = &v
	return s
}

func (s *ListAllSystemContractRequest) Validate() error {
	return dara.Validate(s)
}
