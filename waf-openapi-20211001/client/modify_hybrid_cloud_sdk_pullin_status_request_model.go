// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudSdkPullinStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyHybridCloudSdkPullinStatusRequest
	GetInstanceId() *string
	SetMid(v string) *ModifyHybridCloudSdkPullinStatusRequest
	GetMid() *string
	SetPullinStatus(v string) *ModifyHybridCloudSdkPullinStatusRequest
	GetPullinStatus() *string
}

type ModifyHybridCloudSdkPullinStatusRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-7pp2le***01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the SDK.
	//
	// This parameter is required.
	//
	// example:
	//
	// b3dbc5153317c79d8ca9f9***ea
	Mid *string `json:"Mid,omitempty" xml:"Mid,omitempty"`
	// The status of traffic redirection. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	PullinStatus *string `json:"PullinStatus,omitempty" xml:"PullinStatus,omitempty"`
}

func (s ModifyHybridCloudSdkPullinStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudSdkPullinStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudSdkPullinStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHybridCloudSdkPullinStatusRequest) GetMid() *string {
	return s.Mid
}

func (s *ModifyHybridCloudSdkPullinStatusRequest) GetPullinStatus() *string {
	return s.PullinStatus
}

func (s *ModifyHybridCloudSdkPullinStatusRequest) SetInstanceId(v string) *ModifyHybridCloudSdkPullinStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHybridCloudSdkPullinStatusRequest) SetMid(v string) *ModifyHybridCloudSdkPullinStatusRequest {
	s.Mid = &v
	return s
}

func (s *ModifyHybridCloudSdkPullinStatusRequest) SetPullinStatus(v string) *ModifyHybridCloudSdkPullinStatusRequest {
	s.PullinStatus = &v
	return s
}

func (s *ModifyHybridCloudSdkPullinStatusRequest) Validate() error {
	return dara.Validate(s)
}
