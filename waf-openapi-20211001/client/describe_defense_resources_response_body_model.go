// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDefenseResourcesResponseBody
	GetRequestId() *string
	SetResources(v []*DescribeDefenseResourcesResponseBodyResources) *DescribeDefenseResourcesResponseBody
	GetResources() []*DescribeDefenseResourcesResponseBodyResources
	SetTotalCount(v int64) *DescribeDefenseResourcesResponseBody
	GetTotalCount() *int64
}

type DescribeDefenseResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 618F2626-DB27-5187-8C6C-4E61A491DF29
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The protected objects.
	Resources []*DescribeDefenseResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 73
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseResourcesResponseBody) GetResources() []*DescribeDefenseResourcesResponseBodyResources {
	return s.Resources
}

func (s *DescribeDefenseResourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDefenseResourcesResponseBody) SetRequestId(v string) *DescribeDefenseResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBody) SetResources(v []*DescribeDefenseResourcesResponseBodyResources) *DescribeDefenseResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeDefenseResourcesResponseBody) SetTotalCount(v int64) *DescribeDefenseResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBody) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDefenseResourcesResponseBodyResources struct {
	// The status of the tracking cookie.
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled. This is the default value.
	//
	// example:
	//
	// 1
	AcwCookieStatus *int32 `json:"AcwCookieStatus,omitempty" xml:"AcwCookieStatus,omitempty"`
	// The status of the secure attribute of the tracking cookie.
	//
	// 	- **0**: disabled. This is the default value.
	//
	// 	- **1**: enabled.
	//
	// example:
	//
	// 0
	AcwSecureStatus *int32 `json:"AcwSecureStatus,omitempty" xml:"AcwSecureStatus,omitempty"`
	// The status of the secure attribute of the slider CAPTCHA cookie.
	//
	// 	- **0**: disabled. This is the default value.
	//
	// 	- **1**: enabled.
	//
	// example:
	//
	// 0
	AcwV3SecureStatus *int32 `json:"AcwV3SecureStatus,omitempty" xml:"AcwV3SecureStatus,omitempty"`
	// The custom header fields that are used to identify the originating IP addresses of clients. If the value of XffStatus is 1 and CustomHeaders is left empty, the first IP addresses in the XFF header fields are used as the originating IP addresses of clients.
	CustomHeaders []*string `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty" type:"Repeated"`
	// The description of the protected object.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The description of the protected object. Different key-value pairs in a map indicate different properties of the protected object.
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The creation time of the protected object. Unit: seconds.
	//
	// example:
	//
	// 1652149203187
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The most recent modification time of the protected object. Unit: seconds.
	//
	// example:
	//
	// 1665633032000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The Alibaba Cloud account to which the protected object belongs. You can specify this parameter to query protected objects that belong to a specific Alibaba Cloud account. Exact match is supported.
	//
	// example:
	//
	// 135*********46
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The protection pattern.
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
	// test
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The origin of the protected object.
	//
	// example:
	//
	// custom
	ResourceOrigin *string `json:"ResourceOrigin,omitempty" xml:"ResourceOrigin,omitempty"`
	// The response header.
	ResponseHeaders []*DescribeDefenseResourcesResponseBodyResourcesResponseHeaders `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty" type:"Repeated"`
	// Indicates whether the X-Forwarded-For (XFF) header is used.
	//
	// example:
	//
	// 1
	XffStatus *int32 `json:"XffStatus,omitempty" xml:"XffStatus,omitempty"`
}

func (s DescribeDefenseResourcesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetAcwCookieStatus() *int32 {
	return s.AcwCookieStatus
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetAcwSecureStatus() *int32 {
	return s.AcwSecureStatus
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetAcwV3SecureStatus() *int32 {
	return s.AcwV3SecureStatus
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetCustomHeaders() []*string {
	return s.CustomHeaders
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetDescription() *string {
	return s.Description
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetDetail() map[string]interface{} {
	return s.Detail
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetPattern() *string {
	return s.Pattern
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetProduct() *string {
	return s.Product
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetResource() *string {
	return s.Resource
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetResourceOrigin() *string {
	return s.ResourceOrigin
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetResponseHeaders() []*DescribeDefenseResourcesResponseBodyResourcesResponseHeaders {
	return s.ResponseHeaders
}

func (s *DescribeDefenseResourcesResponseBodyResources) GetXffStatus() *int32 {
	return s.XffStatus
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetAcwCookieStatus(v int32) *DescribeDefenseResourcesResponseBodyResources {
	s.AcwCookieStatus = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetAcwSecureStatus(v int32) *DescribeDefenseResourcesResponseBodyResources {
	s.AcwSecureStatus = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetAcwV3SecureStatus(v int32) *DescribeDefenseResourcesResponseBodyResources {
	s.AcwV3SecureStatus = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetCustomHeaders(v []*string) *DescribeDefenseResourcesResponseBodyResources {
	s.CustomHeaders = v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetDescription(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.Description = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetDetail(v map[string]interface{}) *DescribeDefenseResourcesResponseBodyResources {
	s.Detail = v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetGmtCreate(v int64) *DescribeDefenseResourcesResponseBodyResources {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetGmtModified(v int64) *DescribeDefenseResourcesResponseBodyResources {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetOwnerUserId(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetPattern(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.Pattern = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetProduct(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.Product = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetResource(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetResourceGroup(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetResourceOrigin(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.ResourceOrigin = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetResponseHeaders(v []*DescribeDefenseResourcesResponseBodyResourcesResponseHeaders) *DescribeDefenseResourcesResponseBodyResources {
	s.ResponseHeaders = v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetXffStatus(v int32) *DescribeDefenseResourcesResponseBodyResources {
	s.XffStatus = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) Validate() error {
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

type DescribeDefenseResourcesResponseBodyResourcesResponseHeaders struct {
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

func (s DescribeDefenseResourcesResponseBodyResourcesResponseHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourcesResponseBodyResourcesResponseHeaders) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourcesResponseBodyResourcesResponseHeaders) GetKey() *string {
	return s.Key
}

func (s *DescribeDefenseResourcesResponseBodyResourcesResponseHeaders) GetValue() *string {
	return s.Value
}

func (s *DescribeDefenseResourcesResponseBodyResourcesResponseHeaders) SetKey(v string) *DescribeDefenseResourcesResponseBodyResourcesResponseHeaders {
	s.Key = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResourcesResponseHeaders) SetValue(v string) *DescribeDefenseResourcesResponseBodyResourcesResponseHeaders {
	s.Value = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResourcesResponseHeaders) Validate() error {
	return dara.Validate(s)
}
