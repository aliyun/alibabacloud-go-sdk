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
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-3gwy16dojz1m65****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The Alibaba Cloud account to which the CEN instance belongs.
	//
	// 	- If you own the CEN instance, you can skip this parameter.
	//
	// 	- If you do not own the CEN instance, you must specify the ID of the account that owns the CEN instance.
	//
	// example:
	//
	// 102681951715****
	CenOwnerId *int64 `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The IPv4 CIDR block of the associated office network.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47.100.XX.XX
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The verification code. If you do not own the CEN instance, you must call the [SendVerifyCode](https://help.aliyun.com/document_detail/436847.html) operation to obtain a verification code.
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
