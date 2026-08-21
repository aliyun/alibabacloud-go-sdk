// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseResourceXffRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcwCookieStatus(v int32) *ModifyDefenseResourceXffRequest
	GetAcwCookieStatus() *int32
	SetAcwSecureStatus(v int32) *ModifyDefenseResourceXffRequest
	GetAcwSecureStatus() *int32
	SetAcwV3SecureStatus(v int32) *ModifyDefenseResourceXffRequest
	GetAcwV3SecureStatus() *int32
	SetCustomHeaders(v []*string) *ModifyDefenseResourceXffRequest
	GetCustomHeaders() []*string
	SetInstanceId(v string) *ModifyDefenseResourceXffRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyDefenseResourceXffRequest
	GetRegionId() *string
	SetResource(v string) *ModifyDefenseResourceXffRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *ModifyDefenseResourceXffRequest
	GetResourceManagerResourceGroupId() *string
	SetResponseHeaders(v []*ModifyDefenseResourceXffRequestResponseHeaders) *ModifyDefenseResourceXffRequest
	GetResponseHeaders() []*ModifyDefenseResourceXffRequestResponseHeaders
	SetXffStatus(v int32) *ModifyDefenseResourceXffRequest
	GetXffStatus() *int32
}

type ModifyDefenseResourceXffRequest struct {
	// The status of the tracking cookie switch.
	//
	// - **0**: disabled.
	//
	// - **1 (default)**: enabled.
	//
	// example:
	//
	// 0
	AcwCookieStatus *int32 `json:"AcwCookieStatus,omitempty" xml:"AcwCookieStatus,omitempty"`
	// The status of the secure attribute of the tracking cookie.
	//
	// - **0 (default)**: disabled.
	//
	// - **1**: enabled.
	//
	// example:
	//
	// 0
	AcwSecureStatus *int32 `json:"AcwSecureStatus,omitempty" xml:"AcwSecureStatus,omitempty"`
	// The status of the secure attribute of the slider cookie.
	//
	// - **0 (default)**: disabled.
	//
	// - **1**: enabled.
	//
	// example:
	//
	// 0
	AcwV3SecureStatus *int32 `json:"AcwV3SecureStatus,omitempty" xml:"AcwV3SecureStatus,omitempty"`
	// The list of specified header fields.
	//
	// > The first IP address in the specified header field is used as the client source IP address to prevent XFF spoofing. If multiple headers are specified, the system attempts to obtain the source IP address from the headers in order. If the first header does not contain an IP address, the system tries the second header, and so on. If none of the specified headers contain an IP address, the first IP address in the X-Forwarded-For header is used.
	CustomHeaders []*string `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty" type:"Repeated"`
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-wwo****ek07
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the protected object.
	//
	// > The protected object must have been added to WAF. You can call the [DescribeDefenseResources](https://help.aliyun.com/document_detail/461612.html) operation to query the name of the protected object.
	//
	// This parameter is required.
	//
	// example:
	//
	// alb-4pxu81fgagx3h6y****-alb
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm2ki****miwq
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The response header parameters.
	ResponseHeaders []*ModifyDefenseResourceXffRequestResponseHeaders `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty" type:"Repeated"`
	// Specifies whether a Layer 7 proxy (Anti-DDoS Pro, CDN, or similar) is deployed in front of WAF. Valid values:
	//
	// - **0 (default)**: No Layer 7 proxy is deployed.
	//
	// - **1**: A Layer 7 proxy is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	XffStatus *int32 `json:"XffStatus,omitempty" xml:"XffStatus,omitempty"`
}

func (s ModifyDefenseResourceXffRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseResourceXffRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceXffRequest) GetAcwCookieStatus() *int32 {
	return s.AcwCookieStatus
}

func (s *ModifyDefenseResourceXffRequest) GetAcwSecureStatus() *int32 {
	return s.AcwSecureStatus
}

func (s *ModifyDefenseResourceXffRequest) GetAcwV3SecureStatus() *int32 {
	return s.AcwV3SecureStatus
}

func (s *ModifyDefenseResourceXffRequest) GetCustomHeaders() []*string {
	return s.CustomHeaders
}

func (s *ModifyDefenseResourceXffRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDefenseResourceXffRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDefenseResourceXffRequest) GetResource() *string {
	return s.Resource
}

func (s *ModifyDefenseResourceXffRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyDefenseResourceXffRequest) GetResponseHeaders() []*ModifyDefenseResourceXffRequestResponseHeaders {
	return s.ResponseHeaders
}

func (s *ModifyDefenseResourceXffRequest) GetXffStatus() *int32 {
	return s.XffStatus
}

func (s *ModifyDefenseResourceXffRequest) SetAcwCookieStatus(v int32) *ModifyDefenseResourceXffRequest {
	s.AcwCookieStatus = &v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetAcwSecureStatus(v int32) *ModifyDefenseResourceXffRequest {
	s.AcwSecureStatus = &v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetAcwV3SecureStatus(v int32) *ModifyDefenseResourceXffRequest {
	s.AcwV3SecureStatus = &v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetCustomHeaders(v []*string) *ModifyDefenseResourceXffRequest {
	s.CustomHeaders = v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetInstanceId(v string) *ModifyDefenseResourceXffRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetRegionId(v string) *ModifyDefenseResourceXffRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetResource(v string) *ModifyDefenseResourceXffRequest {
	s.Resource = &v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseResourceXffRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetResponseHeaders(v []*ModifyDefenseResourceXffRequestResponseHeaders) *ModifyDefenseResourceXffRequest {
	s.ResponseHeaders = v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetXffStatus(v int32) *ModifyDefenseResourceXffRequest {
	s.XffStatus = &v
	return s
}

func (s *ModifyDefenseResourceXffRequest) Validate() error {
	if s.ResponseHeaders != nil {
		for _, item := range s.ResponseHeaders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDefenseResourceXffRequestResponseHeaders struct {
	// The key of the custom response header.
	//
	// example:
	//
	// Header-Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the custom response header.
	//
	// example:
	//
	// Header-Value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyDefenseResourceXffRequestResponseHeaders) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseResourceXffRequestResponseHeaders) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceXffRequestResponseHeaders) GetKey() *string {
	return s.Key
}

func (s *ModifyDefenseResourceXffRequestResponseHeaders) GetValue() *string {
	return s.Value
}

func (s *ModifyDefenseResourceXffRequestResponseHeaders) SetKey(v string) *ModifyDefenseResourceXffRequestResponseHeaders {
	s.Key = &v
	return s
}

func (s *ModifyDefenseResourceXffRequestResponseHeaders) SetValue(v string) *ModifyDefenseResourceXffRequestResponseHeaders {
	s.Value = &v
	return s
}

func (s *ModifyDefenseResourceXffRequestResponseHeaders) Validate() error {
	return dara.Validate(s)
}
