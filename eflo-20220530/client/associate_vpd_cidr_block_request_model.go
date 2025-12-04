// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateVpdCidrBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *AssociateVpdCidrBlockRequest
	GetRegionId() *string
	SetSecondaryCidrBlock(v string) *AssociateVpdCidrBlockRequest
	GetSecondaryCidrBlock() *string
	SetVpdId(v string) *AssociateVpdCidrBlockRequest
	GetVpdId() *string
}

type AssociateVpdCidrBlockRequest struct {
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
	// 172.16.0.0/12
	SecondaryCidrBlock *string `json:"SecondaryCidrBlock,omitempty" xml:"SecondaryCidrBlock,omitempty"`
	// The ID of the Lingjun CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-omqutbff
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s AssociateVpdCidrBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateVpdCidrBlockRequest) GoString() string {
	return s.String()
}

func (s *AssociateVpdCidrBlockRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateVpdCidrBlockRequest) GetSecondaryCidrBlock() *string {
	return s.SecondaryCidrBlock
}

func (s *AssociateVpdCidrBlockRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *AssociateVpdCidrBlockRequest) SetRegionId(v string) *AssociateVpdCidrBlockRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateVpdCidrBlockRequest) SetSecondaryCidrBlock(v string) *AssociateVpdCidrBlockRequest {
	s.SecondaryCidrBlock = &v
	return s
}

func (s *AssociateVpdCidrBlockRequest) SetVpdId(v string) *AssociateVpdCidrBlockRequest {
	s.VpdId = &v
	return s
}

func (s *AssociateVpdCidrBlockRequest) Validate() error {
	return dara.Validate(s)
}
