// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDefenseResourceResponseBody
	GetRequestId() *string
	SetResource(v *DescribeDefenseResourceResponseBodyResource) *DescribeDefenseResourceResponseBody
	GetResource() *DescribeDefenseResourceResponseBodyResource
}

type DescribeDefenseResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 439AADF2-368C-5E98-B14E-3086****0573
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The protected object.
	Resource *DescribeDefenseResourceResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
}

func (s DescribeDefenseResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseResourceResponseBody) GetResource() *DescribeDefenseResourceResponseBodyResource {
	return s.Resource
}

func (s *DescribeDefenseResourceResponseBody) SetRequestId(v string) *DescribeDefenseResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseResourceResponseBody) SetResource(v *DescribeDefenseResourceResponseBodyResource) *DescribeDefenseResourceResponseBody {
	s.Resource = v
	return s
}

func (s *DescribeDefenseResourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDefenseResourceResponseBodyResource struct {
	// The status of the tracking cookie.
	//
	// 	- **0**: disabled.
	//
	// 	- **1**: enabled.
	//
	// example:
	//
	// 0
	AcwCookieStatus *int32 `json:"AcwCookieStatus,omitempty" xml:"AcwCookieStatus,omitempty"`
	// The status of the secure attribute of the tracking cookie.
	//
	// 	- **0**: disabled.
	//
	// 	- **1**: enabled.
	//
	// example:
	//
	// 0
	AcwSecureStatus *int32 `json:"AcwSecureStatus,omitempty" xml:"AcwSecureStatus,omitempty"`
	// The status of the secure attribute of the slider CAPTCHA cookie.
	//
	// 	- **0**: disabled.
	//
	// 	- **1**: enabled.
	//
	// example:
	//
	// 0
	AcwV3SecureStatus *int32 `json:"AcwV3SecureStatus,omitempty" xml:"AcwV3SecureStatus,omitempty"`
	// The custom header fields.
	//
	// >  If the value of XffStatus is 1, the first IP address in the specified header field is used as the originating IP address of the client to prevent X-Forwarded-For (XFF) forgery. If you specify multiple header fields, WAF reads the values of the header fields in sequence until the originating IP address is obtained. If the originating IP address cannot be obtained, the first IP address in the XFF header field is used as the originating IP address of the client.
	CustomHeaders []*string `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty" type:"Repeated"`
	// The description of the protected object.
	//
	// example:
	//
	// This is Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The details of the protected object. Different key-value pairs indicate different attributes of the protected object.
	//
	// example:
	//
	// {
	//
	// "product": "waf",
	//
	//  "domain": "demo.aliyundoc****.com"
	//
	// }
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The time when the protected object was created. Unit: milliseconds.
	//
	// example:
	//
	// 1607493144000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the protected object was modified. Unit: milliseconds.
	//
	// example:
	//
	// 1691720010000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The user ID (UID) of the Alibaba Cloud account to which the protected object belongs.
	//
	// example:
	//
	// 170457******9107
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The pattern used for the protected object.
	//
	// example:
	//
	// domain
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The name of the cloud service.
	//
	// example:
	//
	// alb
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The name of the protected object.
	//
	// example:
	//
	// alb-rencs***
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The name of the protected object group to which the protected object belongs.
	//
	// example:
	//
	// example_resource_group
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfmoiy****p2oq
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The origin of the protected object. Valid values:
	//
	// 	- **custom**
	//
	// 	- **access**
	//
	// example:
	//
	// custom
	ResourceOrigin *string `json:"ResourceOrigin,omitempty" xml:"ResourceOrigin,omitempty"`
	// The response header.
	ResponseHeaders []*DescribeDefenseResourceResponseBodyResourceResponseHeaders `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty" type:"Repeated"`
	// Indicates whether a Layer 7 proxy is deployed in front of WAF, such as Anti-DDoS Proxy and Alibaba Cloud CDN. Valid values:
	//
	// 	- **0**: No Layer 7 proxy is deployed.
	//
	// 	- **1**: A Layer 7 proxy is deployed.
	//
	// example:
	//
	// 0
	XffStatus *int32 `json:"XffStatus,omitempty" xml:"XffStatus,omitempty"`
}

func (s DescribeDefenseResourceResponseBodyResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceResponseBodyResource) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceResponseBodyResource) GetAcwCookieStatus() *int32 {
	return s.AcwCookieStatus
}

func (s *DescribeDefenseResourceResponseBodyResource) GetAcwSecureStatus() *int32 {
	return s.AcwSecureStatus
}

func (s *DescribeDefenseResourceResponseBodyResource) GetAcwV3SecureStatus() *int32 {
	return s.AcwV3SecureStatus
}

func (s *DescribeDefenseResourceResponseBodyResource) GetCustomHeaders() []*string {
	return s.CustomHeaders
}

func (s *DescribeDefenseResourceResponseBodyResource) GetDescription() *string {
	return s.Description
}

func (s *DescribeDefenseResourceResponseBodyResource) GetDetail() map[string]interface{} {
	return s.Detail
}

func (s *DescribeDefenseResourceResponseBodyResource) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeDefenseResourceResponseBodyResource) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeDefenseResourceResponseBodyResource) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *DescribeDefenseResourceResponseBodyResource) GetPattern() *string {
	return s.Pattern
}

func (s *DescribeDefenseResourceResponseBodyResource) GetProduct() *string {
	return s.Product
}

func (s *DescribeDefenseResourceResponseBodyResource) GetResource() *string {
	return s.Resource
}

func (s *DescribeDefenseResourceResponseBodyResource) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *DescribeDefenseResourceResponseBodyResource) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseResourceResponseBodyResource) GetResourceOrigin() *string {
	return s.ResourceOrigin
}

func (s *DescribeDefenseResourceResponseBodyResource) GetResponseHeaders() []*DescribeDefenseResourceResponseBodyResourceResponseHeaders {
	return s.ResponseHeaders
}

func (s *DescribeDefenseResourceResponseBodyResource) GetXffStatus() *int32 {
	return s.XffStatus
}

func (s *DescribeDefenseResourceResponseBodyResource) SetAcwCookieStatus(v int32) *DescribeDefenseResourceResponseBodyResource {
	s.AcwCookieStatus = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetAcwSecureStatus(v int32) *DescribeDefenseResourceResponseBodyResource {
	s.AcwSecureStatus = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetAcwV3SecureStatus(v int32) *DescribeDefenseResourceResponseBodyResource {
	s.AcwV3SecureStatus = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetCustomHeaders(v []*string) *DescribeDefenseResourceResponseBodyResource {
	s.CustomHeaders = v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetDescription(v string) *DescribeDefenseResourceResponseBodyResource {
	s.Description = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetDetail(v map[string]interface{}) *DescribeDefenseResourceResponseBodyResource {
	s.Detail = v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetGmtCreate(v int64) *DescribeDefenseResourceResponseBodyResource {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetGmtModified(v int64) *DescribeDefenseResourceResponseBodyResource {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetOwnerUserId(v string) *DescribeDefenseResourceResponseBodyResource {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetPattern(v string) *DescribeDefenseResourceResponseBodyResource {
	s.Pattern = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetProduct(v string) *DescribeDefenseResourceResponseBodyResource {
	s.Product = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetResource(v string) *DescribeDefenseResourceResponseBodyResource {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetResourceGroup(v string) *DescribeDefenseResourceResponseBodyResource {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceResponseBodyResource {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetResourceOrigin(v string) *DescribeDefenseResourceResponseBodyResource {
	s.ResourceOrigin = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetResponseHeaders(v []*DescribeDefenseResourceResponseBodyResourceResponseHeaders) *DescribeDefenseResourceResponseBodyResource {
	s.ResponseHeaders = v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetXffStatus(v int32) *DescribeDefenseResourceResponseBodyResource {
	s.XffStatus = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) Validate() error {
	return dara.Validate(s)
}

type DescribeDefenseResourceResponseBodyResourceResponseHeaders struct {
	// Specifies the key for a custom response header.
	//
	// example:
	//
	// Header-Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Specifies the value for a custom response header.
	//
	// example:
	//
	// Header-Value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDefenseResourceResponseBodyResourceResponseHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceResponseBodyResourceResponseHeaders) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceResponseBodyResourceResponseHeaders) GetKey() *string {
	return s.Key
}

func (s *DescribeDefenseResourceResponseBodyResourceResponseHeaders) GetValue() *string {
	return s.Value
}

func (s *DescribeDefenseResourceResponseBodyResourceResponseHeaders) SetKey(v string) *DescribeDefenseResourceResponseBodyResourceResponseHeaders {
	s.Key = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResourceResponseHeaders) SetValue(v string) *DescribeDefenseResourceResponseBodyResourceResponseHeaders {
	s.Value = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResourceResponseHeaders) Validate() error {
	return dara.Validate(s)
}
