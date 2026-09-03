// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *AttachCenRequest
	GetCenId() *string
	SetCenOwnerId(v int64) *AttachCenRequest
	GetCenOwnerId() *int64
	SetOfficeSiteId(v string) *AttachCenRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *AttachCenRequest
	GetRegionId() *string
	SetVerifyCode(v string) *AttachCenRequest
	GetVerifyCode() *string
}

type AttachCenRequest struct {
	// The CEN instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-3gwy16dojz1m65****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The Alibaba Cloud account ID of the account to which the CEN instance belongs.
	//
	// - If the specified CenId belongs to the current Alibaba Cloud account, you do not need to configure this parameter.
	//
	// - If the specified CenId belongs to a different Alibaba Cloud account, specify the Alibaba Cloud account ID of that account.
	//
	// example:
	//
	// 102681951715****
	CenOwnerId *int64 `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The ID of the office network.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The verification code. If the specified CenId belongs to a different Alibaba Cloud account, call [SendVerifyCode](https://help.aliyun.com/document_detail/436847.html) to obtain the verification code first.
	//
	// example:
	//
	// 12****
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s AttachCenRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachCenRequest) GoString() string {
	return s.String()
}

func (s *AttachCenRequest) GetCenId() *string {
	return s.CenId
}

func (s *AttachCenRequest) GetCenOwnerId() *int64 {
	return s.CenOwnerId
}

func (s *AttachCenRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *AttachCenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachCenRequest) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *AttachCenRequest) SetCenId(v string) *AttachCenRequest {
	s.CenId = &v
	return s
}

func (s *AttachCenRequest) SetCenOwnerId(v int64) *AttachCenRequest {
	s.CenOwnerId = &v
	return s
}

func (s *AttachCenRequest) SetOfficeSiteId(v string) *AttachCenRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *AttachCenRequest) SetRegionId(v string) *AttachCenRequest {
	s.RegionId = &v
	return s
}

func (s *AttachCenRequest) SetVerifyCode(v string) *AttachCenRequest {
	s.VerifyCode = &v
	return s
}

func (s *AttachCenRequest) Validate() error {
	return dara.Validate(s)
}
