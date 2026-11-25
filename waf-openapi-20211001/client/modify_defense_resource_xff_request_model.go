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
	SetCookieConfig(v *ModifyDefenseResourceXffRequestCookieConfig) *ModifyDefenseResourceXffRequest
	GetCookieConfig() *ModifyDefenseResourceXffRequestCookieConfig
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
	CookieConfig *ModifyDefenseResourceXffRequestCookieConfig `json:"CookieConfig,omitempty" xml:"CookieConfig,omitempty" type:"Struct"`
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
	ResponseHeaders []*ModifyDefenseResourceXffRequestResponseHeaders `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty" type:"Repeated"`
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

func (s *ModifyDefenseResourceXffRequest) GetCookieConfig() *ModifyDefenseResourceXffRequestCookieConfig {
	return s.CookieConfig
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

func (s *ModifyDefenseResourceXffRequest) SetCookieConfig(v *ModifyDefenseResourceXffRequestCookieConfig) *ModifyDefenseResourceXffRequest {
	s.CookieConfig = v
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
	if s.CookieConfig != nil {
		if err := s.CookieConfig.Validate(); err != nil {
			return err
		}
	}
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

type ModifyDefenseResourceXffRequestCookieConfig struct {
	// The SameSite policy of the slider cookie. Valid values:
	//
	// - **None**: None.
	//
	// - **Lax**: Lax.
	//
	// - **Strict**: Strict.
	//
	// > An empty string indicates that the SameSite attribute is not configured.
	//
	// If the SameSite policy is set to None, the Secure attribute must be enabled for the cookie.
	//
	// example:
	//
	// None
	AcwScV3SameSite *string `json:"AcwScV3SameSite,omitempty" xml:"AcwScV3SameSite,omitempty"`
	// The SameSite policy of the tracking cookie. Valid values:
	//
	// - **None**: None.
	//
	// - **Lax**: Lax.
	//
	// - **Strict**: Strict.
	//
	// > An empty string indicates that the SameSite attribute is not configured.
	//
	// If the SameSite policy is set to None, the Secure attribute must be enabled for the cookie.
	//
	// example:
	//
	// Lax
	AcwTcSameSite *string `json:"AcwTcSameSite,omitempty" xml:"AcwTcSameSite,omitempty"`
}

func (s ModifyDefenseResourceXffRequestCookieConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseResourceXffRequestCookieConfig) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceXffRequestCookieConfig) GetAcwScV3SameSite() *string {
	return s.AcwScV3SameSite
}

func (s *ModifyDefenseResourceXffRequestCookieConfig) GetAcwTcSameSite() *string {
	return s.AcwTcSameSite
}

func (s *ModifyDefenseResourceXffRequestCookieConfig) SetAcwScV3SameSite(v string) *ModifyDefenseResourceXffRequestCookieConfig {
	s.AcwScV3SameSite = &v
	return s
}

func (s *ModifyDefenseResourceXffRequestCookieConfig) SetAcwTcSameSite(v string) *ModifyDefenseResourceXffRequestCookieConfig {
	s.AcwTcSameSite = &v
	return s
}

func (s *ModifyDefenseResourceXffRequestCookieConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyDefenseResourceXffRequestResponseHeaders struct {
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
