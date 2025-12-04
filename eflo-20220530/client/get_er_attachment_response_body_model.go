// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetErAttachmentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetErAttachmentResponseBody
	GetCode() *int32
	SetContent(v *GetErAttachmentResponseBodyContent) *GetErAttachmentResponseBody
	GetContent() *GetErAttachmentResponseBodyContent
	SetMessage(v string) *GetErAttachmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetErAttachmentResponseBody
	GetRequestId() *string
}

type GetErAttachmentResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *GetErAttachmentResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is displayed.)
	//
	// example:
	//
	// You don\\"t have the permission to do this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7F0D9440-1F97-5613-87CD-D3047172A93C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetErAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetErAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetErAttachmentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetErAttachmentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetErAttachmentResponseBody) GetContent() *GetErAttachmentResponseBodyContent {
	return s.Content
}

func (s *GetErAttachmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetErAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetErAttachmentResponseBody) SetAccessDeniedDetail(v string) *GetErAttachmentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetErAttachmentResponseBody) SetCode(v int32) *GetErAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetErAttachmentResponseBody) SetContent(v *GetErAttachmentResponseBodyContent) *GetErAttachmentResponseBody {
	s.Content = v
	return s
}

func (s *GetErAttachmentResponseBody) SetMessage(v string) *GetErAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetErAttachmentResponseBody) SetRequestId(v string) *GetErAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetErAttachmentResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetErAttachmentResponseBodyContent struct {
	// Whether cross-account. Valid values:
	//
	// 	- **true**: The network instance is a cross-account resource.
	//
	// 	- **false**: The current network instance is a resource of the current account.
	//
	// example:
	//
	// fasle
	Across *bool `json:"Across,omitempty" xml:"Across,omitempty"`
	// Indicates whether to automatically receive all routes from all instances under the Lingjun HUB. Valid values:
	//
	// 	- **true**: received automatically.
	//
	// 	- **false**: Not received.
	//
	// example:
	//
	// true
	AutoReceiveAllRoute *bool `json:"AutoReceiveAllRoute,omitempty" xml:"AutoReceiveAllRoute,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1648085472000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Lingjun HUB network instance.
	//
	// example:
	//
	// er-attachment-i1ioibyf
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	// The name of the Lingjun HUB network instance.
	//
	// example:
	//
	// vpd-lxnsj2cx
	ErAttachmentName *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	// The ID of the Lingjun HUB instance.
	//
	// example:
	//
	// er-aueyxxsy
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 1648085472000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the network instance. Valid values: **VPD*	- and **VCC**.
	//
	// For more information, see [What is Lingjun?](https://help.aliyun.com/document_detail/444430.html)
	//
	// You can query **Lingjun CIDR blocks*	- and **Lingjun connections*	- by [ListVpds](https://help.aliyun.com/document_detail/2331077.html) and [ListVccs](https://help.aliyun.com/document_detail/2399526.html?) respectively.
	//
	// example:
	//
	// vpd-lxnsj2cx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// vpd-wulanchabu-main
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The database type. Valid values:
	//
	// 	- **VPD**: indicates the Lingjun CIDR block.
	//
	// 	- **VCC**: indicates a Lingjun connection.
	//
	// example:
	//
	// VPD
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The returned message.
	//
	// example:
	//
	// test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aekzb3n5lgk2ieq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant to which the resource belongs.
	//
	// example:
	//
	// 1620939556166277
	ResourceTenantId *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	// The status of the cache reserve instance. Valid values:
	//
	// 	- **Available**: Normal.
	//
	// 	- **Not Available**: Not available.
	//
	// 	- **Executing**: The task is being executed.
	//
	// 	- **Deleting**: The account is being deleted
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetErAttachmentResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetErAttachmentResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetErAttachmentResponseBodyContent) GetAcross() *bool {
	return s.Across
}

func (s *GetErAttachmentResponseBodyContent) GetAutoReceiveAllRoute() *bool {
	return s.AutoReceiveAllRoute
}

func (s *GetErAttachmentResponseBodyContent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetErAttachmentResponseBodyContent) GetErAttachmentId() *string {
	return s.ErAttachmentId
}

func (s *GetErAttachmentResponseBodyContent) GetErAttachmentName() *string {
	return s.ErAttachmentName
}

func (s *GetErAttachmentResponseBodyContent) GetErId() *string {
	return s.ErId
}

func (s *GetErAttachmentResponseBodyContent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetErAttachmentResponseBodyContent) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetErAttachmentResponseBodyContent) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetErAttachmentResponseBodyContent) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetErAttachmentResponseBodyContent) GetMessage() *string {
	return s.Message
}

func (s *GetErAttachmentResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetErAttachmentResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetErAttachmentResponseBodyContent) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *GetErAttachmentResponseBodyContent) GetStatus() *string {
	return s.Status
}

func (s *GetErAttachmentResponseBodyContent) GetTenantId() *string {
	return s.TenantId
}

func (s *GetErAttachmentResponseBodyContent) SetAcross(v bool) *GetErAttachmentResponseBodyContent {
	s.Across = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetAutoReceiveAllRoute(v bool) *GetErAttachmentResponseBodyContent {
	s.AutoReceiveAllRoute = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetCreateTime(v string) *GetErAttachmentResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetErAttachmentId(v string) *GetErAttachmentResponseBodyContent {
	s.ErAttachmentId = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetErAttachmentName(v string) *GetErAttachmentResponseBodyContent {
	s.ErAttachmentName = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetErId(v string) *GetErAttachmentResponseBodyContent {
	s.ErId = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetGmtModified(v string) *GetErAttachmentResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetInstanceId(v string) *GetErAttachmentResponseBodyContent {
	s.InstanceId = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetInstanceName(v string) *GetErAttachmentResponseBodyContent {
	s.InstanceName = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetInstanceType(v string) *GetErAttachmentResponseBodyContent {
	s.InstanceType = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetMessage(v string) *GetErAttachmentResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetRegionId(v string) *GetErAttachmentResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetResourceGroupId(v string) *GetErAttachmentResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetResourceTenantId(v string) *GetErAttachmentResponseBodyContent {
	s.ResourceTenantId = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetStatus(v string) *GetErAttachmentResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetTenantId(v string) *GetErAttachmentResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
