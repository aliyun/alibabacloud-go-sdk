// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkerResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListWorkerResourceResponseBody
	GetCurrentPage() *int32
	SetData(v []*ListWorkerResourceResponseBodyData) *ListWorkerResourceResponseBody
	GetData() []*ListWorkerResourceResponseBodyData
	SetRequestId(v string) *ListWorkerResourceResponseBody
	GetRequestId() *string
	SetShowSize(v int32) *ListWorkerResourceResponseBody
	GetShowSize() *int32
	SetTotal(v int64) *ListWorkerResourceResponseBody
	GetTotal() *int64
}

type ListWorkerResourceResponseBody struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The response parameters.
	Data []*ListWorkerResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 3E50D480-9765-5CFD-9650-9BACCECA5135
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries per page. Default value: **50**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 8
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListWorkerResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkerResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkerResourceResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListWorkerResourceResponseBody) GetData() []*ListWorkerResourceResponseBodyData {
	return s.Data
}

func (s *ListWorkerResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkerResourceResponseBody) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListWorkerResourceResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListWorkerResourceResponseBody) SetCurrentPage(v int32) *ListWorkerResourceResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListWorkerResourceResponseBody) SetData(v []*ListWorkerResourceResponseBodyData) *ListWorkerResourceResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkerResourceResponseBody) SetRequestId(v string) *ListWorkerResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkerResourceResponseBody) SetShowSize(v int32) *ListWorkerResourceResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListWorkerResourceResponseBody) SetTotal(v int64) *ListWorkerResourceResponseBody {
	s.Total = &v
	return s
}

func (s *ListWorkerResourceResponseBody) Validate() error {
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

type ListWorkerResourceResponseBodyData struct {
	// The domain name bound to the certificate in the worker task.
	//
	// example:
	//
	// www.example.com
	CertDomain *string `json:"CertDomain,omitempty" xml:"CertDomain,omitempty"`
	// The ID of the certificate in the worker task.
	//
	// example:
	//
	// 12073663
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The instance ID of the certificate in the worker task.
	//
	// example:
	//
	// lsn-jzk2h0xz5dmynuphb8@1883
	CertInstanceId *string `json:"CertInstanceId,omitempty" xml:"CertInstanceId,omitempty"`
	// The name of the certificate in the worker task.
	//
	// example:
	//
	// testCertName
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The cloud service provider to which the cloud resource in the worker task belongs.
	//
	// >  This parameter is not returned if you deploy certificates to Alibaba Cloud services.
	//
	// example:
	//
	// aliyun
	CloudName *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	// The cloud service to which the cloud resource in the worker task belongs. Valid values:
	//
	// 	- **CDN**: Alibaba Cloud CDN (CDN). This value is supported only at the China site (aliyun.com).
	//
	// 	- **SLB**: Classic Load Balancer (CLB). This value is supported only at the China site (aliyun.com).
	//
	// 	- **DCDN**: Dynamic Content Delivery Network (DCDN). This value is supported only at the China site (aliyun.com).
	//
	// 	- **DDOS**: Anti-DDoS. This value is supported only at the China site (aliyun.com).
	//
	// 	- **LIVE**: ApsaraVideo Live. This value is supported only at the China site (aliyun.com).
	//
	// 	- **webHosting**: Cloud Web Hosting. This value is supported only at the China site (aliyun.com).
	//
	// 	- **VOD**: ApsaraVideo VOD. This value is supported only at the China site (aliyun.com).
	//
	// 	- **CR**: Container Registry. This value is supported only at the China site (aliyun.com).
	//
	// 	- **ALB**: Application Load Balancer (ALB).
	//
	// 	- **APIGateway**: API Gateway.
	//
	// 	- **FC**: Function Compute.
	//
	// 	- **GA**: Global Accelerator (GA).
	//
	// 	- **MSE**: Microservices Engine (MSE).
	//
	// 	- **NLB**: Network Load Balancer (NLB).
	//
	// 	- **OSS**: Object Storage Service (OSS).
	//
	// 	- **SAE**: Serverless App Engine (SAE).
	//
	// 	- **TencentCDN**: Tencent Cloud Content Delivery Network (CDN).
	//
	// 	- **WAF**: Web Application Firewall (WAF).
	//
	// example:
	//
	// SLB
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	// The original region ID of the cloud resource in the worker task. The value is the region ID defined by the cloud service provider. This parameter is required only when you deploy certificates to services of multiple clouds.
	//
	// example:
	//
	// cn-hangzhou
	CloudRegion *string `json:"CloudRegion,omitempty" xml:"CloudRegion,omitempty"`
	// Indicates whether the cloud resource in the worker task is the default resource. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// 0
	DefaultResource *bool `json:"DefaultResource,omitempty" xml:"DefaultResource,omitempty"`
	// The time when the worker task was created. The time is a timestamp in seconds.
	//
	// example:
	//
	// 1680228896000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the worker task was last modified. The time is a timestamp in seconds.
	//
	// example:
	//
	// 1681956830000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the worker task.
	//
	// example:
	//
	// 22487
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the cloud resource in the worker task.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// cas-cn-0pp118nuu40b
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the deployment task to which the worker task belongs.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The listener ID of the cloud resource in the worker task.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// lsn-lhl8y7s1e1ngc3m3zz@81
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the namespace in SAE. This parameter is returned only if you deploy certificates to SAE.
	//
	// example:
	//
	// 32fed52a-4bf7-44f6-955f-e82ada68ef18
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The order ID of the worker task, which is the same as the order ID of the certificate.
	//
	// >  If the CertId parameter is returned, this parameter is not returned.
	//
	// example:
	//
	// 9349278
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The listening port of the cloud resource in the worker task.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// 4431
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the cloud resource in the worker task.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the certificate that was originally bound to the cloud resource in the worker task.
	//
	// example:
	//
	// 123
	ResourceCertId *int64 `json:"ResourceCertId,omitempty" xml:"ResourceCertId,omitempty"`
	// The domain name that was originally bound to the cloud resource in the worker task.
	//
	// example:
	//
	// www.example.com
	ResourceDomain *string `json:"ResourceDomain,omitempty" xml:"ResourceDomain,omitempty"`
	// The ID of the cloud resource in the worker task.
	//
	// example:
	//
	// 1286999
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The status of the worker task. Valid values:
	//
	// 	- **editing**
	//
	// 	- **pending**
	//
	// 	- **scheduling**
	//
	// 	- **processing**
	//
	// 	- **error**
	//
	// 	- **success**
	//
	// 	- **rollback**
	//
	// 	- **rollback_success**
	//
	// 	- **rollback_error**
	//
	// example:
	//
	// editing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the Alibaba Cloud account to which the worker task belongs.
	//
	// example:
	//
	// 1666884372152785
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListWorkerResourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWorkerResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkerResourceResponseBodyData) GetCertDomain() *string {
	return s.CertDomain
}

func (s *ListWorkerResourceResponseBodyData) GetCertId() *int64 {
	return s.CertId
}

func (s *ListWorkerResourceResponseBodyData) GetCertInstanceId() *string {
	return s.CertInstanceId
}

func (s *ListWorkerResourceResponseBodyData) GetCertName() *string {
	return s.CertName
}

func (s *ListWorkerResourceResponseBodyData) GetCloudName() *string {
	return s.CloudName
}

func (s *ListWorkerResourceResponseBodyData) GetCloudProduct() *string {
	return s.CloudProduct
}

func (s *ListWorkerResourceResponseBodyData) GetCloudRegion() *string {
	return s.CloudRegion
}

func (s *ListWorkerResourceResponseBodyData) GetDefaultResource() *bool {
	return s.DefaultResource
}

func (s *ListWorkerResourceResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListWorkerResourceResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListWorkerResourceResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListWorkerResourceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWorkerResourceResponseBodyData) GetJobId() *int64 {
	return s.JobId
}

func (s *ListWorkerResourceResponseBodyData) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListWorkerResourceResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListWorkerResourceResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListWorkerResourceResponseBodyData) GetPort() *int32 {
	return s.Port
}

func (s *ListWorkerResourceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListWorkerResourceResponseBodyData) GetResourceCertId() *int64 {
	return s.ResourceCertId
}

func (s *ListWorkerResourceResponseBodyData) GetResourceDomain() *string {
	return s.ResourceDomain
}

func (s *ListWorkerResourceResponseBodyData) GetResourceId() *int64 {
	return s.ResourceId
}

func (s *ListWorkerResourceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListWorkerResourceResponseBodyData) GetUserId() *int64 {
	return s.UserId
}

func (s *ListWorkerResourceResponseBodyData) SetCertDomain(v string) *ListWorkerResourceResponseBodyData {
	s.CertDomain = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetCertId(v int64) *ListWorkerResourceResponseBodyData {
	s.CertId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetCertInstanceId(v string) *ListWorkerResourceResponseBodyData {
	s.CertInstanceId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetCertName(v string) *ListWorkerResourceResponseBodyData {
	s.CertName = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetCloudName(v string) *ListWorkerResourceResponseBodyData {
	s.CloudName = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetCloudProduct(v string) *ListWorkerResourceResponseBodyData {
	s.CloudProduct = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetCloudRegion(v string) *ListWorkerResourceResponseBodyData {
	s.CloudRegion = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetDefaultResource(v bool) *ListWorkerResourceResponseBodyData {
	s.DefaultResource = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetGmtCreate(v string) *ListWorkerResourceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetGmtModified(v string) *ListWorkerResourceResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetId(v int64) *ListWorkerResourceResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetInstanceId(v string) *ListWorkerResourceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetJobId(v int64) *ListWorkerResourceResponseBodyData {
	s.JobId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetListenerId(v string) *ListWorkerResourceResponseBodyData {
	s.ListenerId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetNamespaceId(v string) *ListWorkerResourceResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetOrderId(v int64) *ListWorkerResourceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetPort(v int32) *ListWorkerResourceResponseBodyData {
	s.Port = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetRegionId(v string) *ListWorkerResourceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetResourceCertId(v int64) *ListWorkerResourceResponseBodyData {
	s.ResourceCertId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetResourceDomain(v string) *ListWorkerResourceResponseBodyData {
	s.ResourceDomain = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetResourceId(v int64) *ListWorkerResourceResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetStatus(v string) *ListWorkerResourceResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetUserId(v int64) *ListWorkerResourceResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
