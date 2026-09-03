// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *VerifyCenRequest
	GetCenId() *string
	SetCenOwnerId(v int64) *VerifyCenRequest
	GetCenOwnerId() *int64
	SetCidrBlock(v string) *VerifyCenRequest
	GetCidrBlock() *string
	SetRegionId(v string) *VerifyCenRequest
	GetRegionId() *string
	SetVerifyCode(v string) *VerifyCenRequest
	GetVerifyCode() *string
}

type VerifyCenRequest struct {
	// The CEN instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-3gwy16dojz1m65****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The Alibaba Cloud account ID to which the CEN instance belongs.
	//
	// - If the specified CenId belongs to your Alibaba Cloud account, you do not need to configure this parameter.
	//
	// - If the specified CenId belongs to another Alibaba Cloud account, specify the ID of that Alibaba Cloud account.
	//
	// example:
	//
	// 102681951715****
	CenOwnerId *int64 `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The IPv4 CIDR block of the office network.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/16
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The verification code. If the specified CenId belongs to another Alibaba Cloud account, you must first call [SendVerifyCode](https://help.aliyun.com/document_detail/436847.html) to obtain the verification code.
	//
	// example:
	//
	// 12****
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s VerifyCenRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyCenRequest) GoString() string {
	return s.String()
}

func (s *VerifyCenRequest) GetCenId() *string {
	return s.CenId
}

func (s *VerifyCenRequest) GetCenOwnerId() *int64 {
	return s.CenOwnerId
}

func (s *VerifyCenRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *VerifyCenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *VerifyCenRequest) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *VerifyCenRequest) SetCenId(v string) *VerifyCenRequest {
	s.CenId = &v
	return s
}

func (s *VerifyCenRequest) SetCenOwnerId(v int64) *VerifyCenRequest {
	s.CenOwnerId = &v
	return s
}

func (s *VerifyCenRequest) SetCidrBlock(v string) *VerifyCenRequest {
	s.CidrBlock = &v
	return s
}

func (s *VerifyCenRequest) SetRegionId(v string) *VerifyCenRequest {
	s.RegionId = &v
	return s
}

func (s *VerifyCenRequest) SetVerifyCode(v string) *VerifyCenRequest {
	s.VerifyCode = &v
	return s
}

func (s *VerifyCenRequest) Validate() error {
	return dara.Validate(s)
}
