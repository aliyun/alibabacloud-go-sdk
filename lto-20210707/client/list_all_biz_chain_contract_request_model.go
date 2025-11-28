// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllBizChainContractRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizChainId(v string) *ListAllBizChainContractRequest
	GetBizChainId() *string
	SetRegionId(v string) *ListAllBizChainContractRequest
	GetRegionId() *string
}

type ListAllBizChainContractRequest struct {
	// This parameter is required.
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAllBizChainContractRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllBizChainContractRequest) GoString() string {
	return s.String()
}

func (s *ListAllBizChainContractRequest) GetBizChainId() *string {
	return s.BizChainId
}

func (s *ListAllBizChainContractRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAllBizChainContractRequest) SetBizChainId(v string) *ListAllBizChainContractRequest {
	s.BizChainId = &v
	return s
}

func (s *ListAllBizChainContractRequest) SetRegionId(v string) *ListAllBizChainContractRequest {
	s.RegionId = &v
	return s
}

func (s *ListAllBizChainContractRequest) Validate() error {
	return dara.Validate(s)
}
