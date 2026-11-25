// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseResourceXffShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcwCookieStatus(v int32) *ModifyDefenseResourceXffShrinkRequest
	GetAcwCookieStatus() *int32
	SetAcwSecureStatus(v int32) *ModifyDefenseResourceXffShrinkRequest
	GetAcwSecureStatus() *int32
	SetAcwV3SecureStatus(v int32) *ModifyDefenseResourceXffShrinkRequest
	GetAcwV3SecureStatus() *int32
	SetCookieConfigShrink(v string) *ModifyDefenseResourceXffShrinkRequest
	GetCookieConfigShrink() *string
	SetCustomHeaders(v []*string) *ModifyDefenseResourceXffShrinkRequest
	GetCustomHeaders() []*string
	SetInstanceId(v string) *ModifyDefenseResourceXffShrinkRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyDefenseResourceXffShrinkRequest
	GetRegionId() *string
	SetResource(v string) *ModifyDefenseResourceXffShrinkRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *ModifyDefenseResourceXffShrinkRequest
	GetResourceManagerResourceGroupId() *string
	SetResponseHeaders(v []*ModifyDefenseResourceXffShrinkRequestResponseHeaders) *ModifyDefenseResourceXffShrinkRequest
	GetResponseHeaders() []*ModifyDefenseResourceXffShrinkRequestResponseHeaders
	SetXffStatus(v int32) *ModifyDefenseResourceXffShrinkRequest
	GetXffStatus() *int32
}

type ModifyDefenseResourceXffShrinkRequest struct {
	// The status of the tracking cookie switch. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1*	- (default): Enabled.
	//
	// example:
	//
	// 0
	AcwCookieStatus *int32 `json:"AcwCookieStatus,omitempty" xml:"AcwCookieStatus,omitempty"`
	// The status of the Secure attribute for the tracking cookie. Valid values:
	//
	// - **0*	- (default): Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 0
	AcwSecureStatus *int32 `json:"AcwSecureStatus,omitempty" xml:"AcwSecureStatus,omitempty"`
	// The status of the Secure attribute for the slider cookie. Valid values:
	//
	// - **0*	- (default): Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 0
	AcwV3SecureStatus *int32 `json:"AcwV3SecureStatus,omitempty" xml:"AcwV3SecureStatus,omitempty"`
	// The aggregated configuration of cookie attributes.
	CookieConfigShrink *string `json:"CookieConfig,omitempty" xml:"CookieConfig,omitempty"`
	// The list of specified header fields.
	//
	// > The first IP address in the specified header field is used as the source IP address of the client to prevent X-Forwarded-For (XFF) spoofing. If multiple headers are specified, the system attempts to obtain the source IP address in sequence. If the first header does not contain an IP address, the system checks the second header, and so on. If no IP address is found in any of the specified headers, the first IP address in the X-Forwarded-For header is used.
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
	// The region where the WAF instance resides. Valid values:
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
	// > The protected object must be added to WAF. You can call the [DescribeDefenseResources](https://help.aliyun.com/document_detail/461612.html) operation to query the name of the protected object.
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
	ResponseHeaders []*ModifyDefenseResourceXffShrinkRequestResponseHeaders `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty" type:"Repeated"`
	// Specifies whether a Layer 7 proxy, such as Anti-DDoS or CDN, is enabled before WAF. Valid values:
	//
	// - **0*	- (default): The proxy is disabled.
	//
	// - **1**: The proxy is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	XffStatus *int32 `json:"XffStatus,omitempty" xml:"XffStatus,omitempty"`
}

func (s ModifyDefenseResourceXffShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseResourceXffShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceXffShrinkRequest) GetAcwCookieStatus() *int32 {
	return s.AcwCookieStatus
}

func (s *ModifyDefenseResourceXffShrinkRequest) GetAcwSecureStatus() *int32 {
	return s.AcwSecureStatus
}

func (s *ModifyDefenseResourceXffShrinkRequest) GetAcwV3SecureStatus() *int32 {
	return s.AcwV3SecureStatus
}

func (s *ModifyDefenseResourceXffShrinkRequest) GetCookieConfigShrink() *string {
	return s.CookieConfigShrink
}

func (s *ModifyDefenseResourceXffShrinkRequest) GetCustomHeaders() []*string {
	return s.CustomHeaders
}

func (s *ModifyDefenseResourceXffShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDefenseResourceXffShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDefenseResourceXffShrinkRequest) GetResource() *string {
	return s.Resource
}

func (s *ModifyDefenseResourceXffShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyDefenseResourceXffShrinkRequest) GetResponseHeaders() []*ModifyDefenseResourceXffShrinkRequestResponseHeaders {
	return s.ResponseHeaders
}

func (s *ModifyDefenseResourceXffShrinkRequest) GetXffStatus() *int32 {
	return s.XffStatus
}

func (s *ModifyDefenseResourceXffShrinkRequest) SetAcwCookieStatus(v int32) *ModifyDefenseResourceXffShrinkRequest {
	s.AcwCookieStatus = &v
	return s
}

func (s *ModifyDefenseResourceXffShrinkRequest) SetAcwSecureStatus(v int32) *ModifyDefenseResourceXffShrinkRequest {
	s.AcwSecureStatus = &v
	return s
}

func (s *ModifyDefenseResourceXffShrinkRequest) SetAcwV3SecureStatus(v int32) *ModifyDefenseResourceXffShrinkRequest {
	s.AcwV3SecureStatus = &v
	return s
}

func (s *ModifyDefenseResourceXffShrinkRequest) SetCookieConfigShrink(v string) *ModifyDefenseResourceXffShrinkRequest {
	s.CookieConfigShrink = &v
	return s
}

func (s *ModifyDefenseResourceXffShrinkRequest) SetCustomHeaders(v []*string) *ModifyDefenseResourceXffShrinkRequest {
	s.CustomHeaders = v
	return s
}

func (s *ModifyDefenseResourceXffShrinkRequest) SetInstanceId(v string) *ModifyDefenseResourceXffShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseResourceXffShrinkRequest) SetRegionId(v string) *ModifyDefenseResourceXffShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseResourceXffShrinkRequest) SetResource(v string) *ModifyDefenseResourceXffShrinkRequest {
	s.Resource = &v
	return s
}

func (s *ModifyDefenseResourceXffShrinkRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseResourceXffShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefenseResourceXffShrinkRequest) SetResponseHeaders(v []*ModifyDefenseResourceXffShrinkRequestResponseHeaders) *ModifyDefenseResourceXffShrinkRequest {
	s.ResponseHeaders = v
	return s
}

func (s *ModifyDefenseResourceXffShrinkRequest) SetXffStatus(v int32) *ModifyDefenseResourceXffShrinkRequest {
	s.XffStatus = &v
	return s
}

func (s *ModifyDefenseResourceXffShrinkRequest) Validate() error {
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

type ModifyDefenseResourceXffShrinkRequestResponseHeaders struct {
	// The key of the specified custom response header.
	//
	// example:
	//
	// Header-Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the specified custom response header.
	//
	// example:
	//
	// Header-Value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyDefenseResourceXffShrinkRequestResponseHeaders) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseResourceXffShrinkRequestResponseHeaders) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceXffShrinkRequestResponseHeaders) GetKey() *string {
	return s.Key
}

func (s *ModifyDefenseResourceXffShrinkRequestResponseHeaders) GetValue() *string {
	return s.Value
}

func (s *ModifyDefenseResourceXffShrinkRequestResponseHeaders) SetKey(v string) *ModifyDefenseResourceXffShrinkRequestResponseHeaders {
	s.Key = &v
	return s
}

func (s *ModifyDefenseResourceXffShrinkRequestResponseHeaders) SetValue(v string) *ModifyDefenseResourceXffShrinkRequestResponseHeaders {
	s.Value = &v
	return s
}

func (s *ModifyDefenseResourceXffShrinkRequestResponseHeaders) Validate() error {
	return dara.Validate(s)
}
