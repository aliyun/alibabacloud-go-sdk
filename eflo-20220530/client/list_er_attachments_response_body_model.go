// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListErAttachmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListErAttachmentsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListErAttachmentsResponseBody
	GetCode() *int32
	SetContent(v *ListErAttachmentsResponseBodyContent) *ListErAttachmentsResponseBody
	GetContent() *ListErAttachmentsResponseBodyContent
	SetMessage(v string) *ListErAttachmentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListErAttachmentsResponseBody
	GetRequestId() *string
}

type ListErAttachmentsResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Content *ListErAttachmentsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is displayed.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3D9D6E7B-365B-5200-BFA6-9B79E269058C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListErAttachmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListErAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListErAttachmentsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListErAttachmentsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListErAttachmentsResponseBody) GetContent() *ListErAttachmentsResponseBodyContent {
	return s.Content
}

func (s *ListErAttachmentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListErAttachmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListErAttachmentsResponseBody) SetAccessDeniedDetail(v string) *ListErAttachmentsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListErAttachmentsResponseBody) SetCode(v int32) *ListErAttachmentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListErAttachmentsResponseBody) SetContent(v *ListErAttachmentsResponseBodyContent) *ListErAttachmentsResponseBody {
	s.Content = v
	return s
}

func (s *ListErAttachmentsResponseBody) SetMessage(v string) *ListErAttachmentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListErAttachmentsResponseBody) SetRequestId(v string) *ListErAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListErAttachmentsResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListErAttachmentsResponseBodyContent struct {
	// The list of Lingjun HUB network instances.
	Data []*ListErAttachmentsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListErAttachmentsResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListErAttachmentsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListErAttachmentsResponseBodyContent) GetData() []*ListErAttachmentsResponseBodyContentData {
	return s.Data
}

func (s *ListErAttachmentsResponseBodyContent) GetTotal() *int64 {
	return s.Total
}

func (s *ListErAttachmentsResponseBodyContent) SetData(v []*ListErAttachmentsResponseBodyContentData) *ListErAttachmentsResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListErAttachmentsResponseBodyContent) SetTotal(v int64) *ListErAttachmentsResponseBodyContent {
	s.Total = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContent) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListErAttachmentsResponseBodyContentData struct {
	// Whether to cross accounts. Valid values:
	//
	// 	- **true**: The network instance is a cross-account resource.
	//
	// 	- **false**: The current network instance is a resource of the current account.
	//
	// example:
	//
	// false
	Across *bool `json:"Across,omitempty" xml:"Across,omitempty"`
	// Whether to automatically receive all routes from all instances under this Lingjun HUB. Valid values:
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
	// 1669734207000
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
	// vcc-cn-209300qha01
	ErAttachmentName *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	// The ID of the Lingjun HUB instance.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 1640187007000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the network instance. Valid values: **VPD*	- and **VCC**.
	//
	// For more information, see [What is Lingjun?](https://help.aliyun.com/document_detail/444430.html)
	//
	// You can query **Lingjun CIDR blocks*	- and **Lingjun connections*	- by [ListVpds](https://help.aliyun.com/document_detail/2331077.html) and [ListVccs](https://help.aliyun.com/document_detail/2399526.html) respectively.
	//
	// example:
	//
	// vcc-cn-209300qha02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// vcc-wulanchabu-main
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The database type. Valid values:
	//
	// 	- **VPD**: indicates the Lingjun CIDR block.
	//
	// 	- **VCC**: indicates a Lingjun connection.
	//
	// example:
	//
	// VCC
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The returned message.
	//
	// example:
	//
	// test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Lingjun HUB region information.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aekzlki4ehfse4y
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant to which the resource belongs.
	//
	// example:
	//
	// 1111156667137893
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

func (s ListErAttachmentsResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s ListErAttachmentsResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListErAttachmentsResponseBodyContentData) GetAcross() *bool {
	return s.Across
}

func (s *ListErAttachmentsResponseBodyContentData) GetAutoReceiveAllRoute() *bool {
	return s.AutoReceiveAllRoute
}

func (s *ListErAttachmentsResponseBodyContentData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListErAttachmentsResponseBodyContentData) GetErAttachmentId() *string {
	return s.ErAttachmentId
}

func (s *ListErAttachmentsResponseBodyContentData) GetErAttachmentName() *string {
	return s.ErAttachmentName
}

func (s *ListErAttachmentsResponseBodyContentData) GetErId() *string {
	return s.ErId
}

func (s *ListErAttachmentsResponseBodyContentData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListErAttachmentsResponseBodyContentData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListErAttachmentsResponseBodyContentData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListErAttachmentsResponseBodyContentData) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListErAttachmentsResponseBodyContentData) GetMessage() *string {
	return s.Message
}

func (s *ListErAttachmentsResponseBodyContentData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListErAttachmentsResponseBodyContentData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListErAttachmentsResponseBodyContentData) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *ListErAttachmentsResponseBodyContentData) GetStatus() *string {
	return s.Status
}

func (s *ListErAttachmentsResponseBodyContentData) GetTenantId() *string {
	return s.TenantId
}

func (s *ListErAttachmentsResponseBodyContentData) SetAcross(v bool) *ListErAttachmentsResponseBodyContentData {
	s.Across = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetAutoReceiveAllRoute(v bool) *ListErAttachmentsResponseBodyContentData {
	s.AutoReceiveAllRoute = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetCreateTime(v string) *ListErAttachmentsResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetErAttachmentId(v string) *ListErAttachmentsResponseBodyContentData {
	s.ErAttachmentId = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetErAttachmentName(v string) *ListErAttachmentsResponseBodyContentData {
	s.ErAttachmentName = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetErId(v string) *ListErAttachmentsResponseBodyContentData {
	s.ErId = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetGmtModified(v string) *ListErAttachmentsResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetInstanceId(v string) *ListErAttachmentsResponseBodyContentData {
	s.InstanceId = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetInstanceName(v string) *ListErAttachmentsResponseBodyContentData {
	s.InstanceName = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetInstanceType(v string) *ListErAttachmentsResponseBodyContentData {
	s.InstanceType = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetMessage(v string) *ListErAttachmentsResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetRegionId(v string) *ListErAttachmentsResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetResourceGroupId(v string) *ListErAttachmentsResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetResourceTenantId(v string) *ListErAttachmentsResponseBodyContentData {
	s.ResourceTenantId = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetStatus(v string) *ListErAttachmentsResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetTenantId(v string) *ListErAttachmentsResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) Validate() error {
	return dara.Validate(s)
}
