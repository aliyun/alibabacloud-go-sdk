// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnAssociateVpdCidrBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UnAssociateVpdCidrBlockRequest
	GetRegionId() *string
	SetSecondaryCidrBlock(v string) *UnAssociateVpdCidrBlockRequest
	GetSecondaryCidrBlock() *string
	SetVpdId(v string) *UnAssociateVpdCidrBlockRequest
	GetVpdId() *string
}

type UnAssociateVpdCidrBlockRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The additional CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/16
	SecondaryCidrBlock *string `json:"SecondaryCidrBlock,omitempty" xml:"SecondaryCidrBlock,omitempty"`
	// The ID of the Lingjun CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-aof7dat1
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s UnAssociateVpdCidrBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s UnAssociateVpdCidrBlockRequest) GoString() string {
	return s.String()
}

func (s *UnAssociateVpdCidrBlockRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnAssociateVpdCidrBlockRequest) GetSecondaryCidrBlock() *string {
	return s.SecondaryCidrBlock
}

func (s *UnAssociateVpdCidrBlockRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *UnAssociateVpdCidrBlockRequest) SetRegionId(v string) *UnAssociateVpdCidrBlockRequest {
	s.RegionId = &v
	return s
}

func (s *UnAssociateVpdCidrBlockRequest) SetSecondaryCidrBlock(v string) *UnAssociateVpdCidrBlockRequest {
	s.SecondaryCidrBlock = &v
	return s
}

func (s *UnAssociateVpdCidrBlockRequest) SetVpdId(v string) *UnAssociateVpdCidrBlockRequest {
	s.VpdId = &v
	return s
}

func (s *UnAssociateVpdCidrBlockRequest) Validate() error {
	return dara.Validate(s)
}
