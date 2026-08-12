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
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) operation to query instance ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-7pp2le***01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The machine identifier (MID). You can call the [DescribeHybridCloudSdkServers](https://help.aliyun.com/document_detail/2982006.html) operation to query the hybrid cloud SDK list and obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// b3dbc5153317c79d8ca9f9***ea
	Mid *string `json:"Mid,omitempty" xml:"Mid,omitempty"`
	// The traffic redirection status. Valid values:
	//
	// - **on**: enabled.
	//
	// - **off**: disabled.
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
