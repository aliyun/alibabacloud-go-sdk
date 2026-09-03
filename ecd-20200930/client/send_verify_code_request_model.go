// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendVerifyCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtraInfo(v string) *SendVerifyCodeRequest
	GetExtraInfo() *string
	SetRegionId(v string) *SendVerifyCodeRequest
	GetRegionId() *string
	SetVerifyCodeAction(v string) *SendVerifyCodeRequest
	GetVerifyCodeAction() *string
}

type SendVerifyCodeRequest struct {
	// The information required to send the verification code, in JSON format. When verifying a CEN instance, provide the CEN instance ID and the Alibaba Cloud account ID to which the CEN instance belongs.
	//
	// - CenId: the CEN instance ID.
	//
	// - CenOwnerId: the Alibaba Cloud account ID to which the CEN instance belongs.
	//
	// > If the specified CenId belongs to the current Alibaba Cloud account, this parameter is not required. If the specified CenId belongs to a different Alibaba Cloud account, specify the Alibaba Cloud account ID of the owner.
	//
	// example:
	//
	// {"cenOwnerId": 1234567890******,"cenId": "cen-3weq30r6t0s7t4****"}
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// The region ID. Call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The action associated with the verification code.
	//
	// This parameter is required.
	//
	// example:
	//
	// eds_cenID_securityverification
	VerifyCodeAction *string `json:"VerifyCodeAction,omitempty" xml:"VerifyCodeAction,omitempty"`
}

func (s SendVerifyCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s SendVerifyCodeRequest) GoString() string {
	return s.String()
}

func (s *SendVerifyCodeRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *SendVerifyCodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SendVerifyCodeRequest) GetVerifyCodeAction() *string {
	return s.VerifyCodeAction
}

func (s *SendVerifyCodeRequest) SetExtraInfo(v string) *SendVerifyCodeRequest {
	s.ExtraInfo = &v
	return s
}

func (s *SendVerifyCodeRequest) SetRegionId(v string) *SendVerifyCodeRequest {
	s.RegionId = &v
	return s
}

func (s *SendVerifyCodeRequest) SetVerifyCodeAction(v string) *SendVerifyCodeRequest {
	s.VerifyCodeAction = &v
	return s
}

func (s *SendVerifyCodeRequest) Validate() error {
	return dara.Validate(s)
}
