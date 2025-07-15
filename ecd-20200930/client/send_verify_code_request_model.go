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
	// The information that is required to send the verification code, in JSON format. When you verify the CEN instance of another Alibaba Cloud account, you must provide the ID of the CEN instance and the ID of the Alibaba Cloud account to which the instance belongs.
	//
	// 	- CenId: the ID of the CEN instance.
	//
	// 	- CenOwnerId: the ID of the Alibaba Cloud account to which the CEN instance belongs.
	//
	// >  If you own the CEN instance, skip this parameter. If you do not own the CEN instance, specify the ID of the Alibaba Cloud account that owns the CEN instance.
	//
	// example:
	//
	// {"cenOwnerId": 1234567890******,"cenId": "cen-3weq30r6t0s7t4****"}
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The action that you want to perform by using the verification code.
	//
	// Valid value:
	//
	// 	- eds_cenID_securityverification: Use the verification code to verify the CEN instance.
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
