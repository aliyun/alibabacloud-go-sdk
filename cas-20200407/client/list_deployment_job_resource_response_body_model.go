// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentJobResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDeploymentJobResourceResponseBodyData) *ListDeploymentJobResourceResponseBody
	GetData() []*ListDeploymentJobResourceResponseBodyData
	SetRequestId(v string) *ListDeploymentJobResourceResponseBody
	GetRequestId() *string
}

type ListDeploymentJobResourceResponseBody struct {
	// The response parameters.
	Data []*ListDeploymentJobResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeploymentJobResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentJobResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobResourceResponseBody) GetData() []*ListDeploymentJobResourceResponseBodyData {
	return s.Data
}

func (s *ListDeploymentJobResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeploymentJobResourceResponseBody) SetData(v []*ListDeploymentJobResourceResponseBodyData) *ListDeploymentJobResourceResponseBody {
	s.Data = v
	return s
}

func (s *ListDeploymentJobResourceResponseBody) SetRequestId(v string) *ListDeploymentJobResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBody) Validate() error {
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

type ListDeploymentJobResourceResponseBodyData struct {
	// The end date of the certificate bound to the cloud resource. The value is a timestamp in seconds.
	//
	// example:
	//
	// 1681956830000
	CertEndTime *string `json:"CertEndTime,omitempty" xml:"CertEndTime,omitempty"`
	// The ID of the certificate bound to the cloud resource.
	//
	// example:
	//
	// 11599949
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The name of the certificate bound to the cloud resource.
	//
	// example:
	//
	// sc-SSL
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The start date of the certificate bound to the cloud resource. The value is a timestamp in seconds.
	//
	// example:
	//
	// 1681956830000
	CertStartTime *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	// The AccessKey ID used to access cloud resources.
	//
	// >  This parameter is required only when you deploy certificates to services of multiple clouds.
	//
	// example:
	//
	// 1234
	CloudAccessId *string `json:"CloudAccessId,omitempty" xml:"CloudAccessId,omitempty"`
	// The cloud service provider of the cloud resource. Valid values:
	//
	// 	- **aliyun**: Alibaba Cloud
	//
	// 	- **Tencent**: Tencent Cloud
	//
	// example:
	//
	// aliyun
	CloudName *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	// The cloud service. Valid values:
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
	// The region ID of the cloud service provider to which the cloud resource belongs.
	//
	// example:
	//
	// cn-hangzhou
	CloudRegion *string `json:"CloudRegion,omitempty" xml:"CloudRegion,omitempty"`
	// Indicates whether the cloud resource is the default resource. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// 0
	DefaultResource *int32 `json:"DefaultResource,omitempty" xml:"DefaultResource,omitempty"`
	// The domain name bound to the cloud resource.
	//
	// example:
	//
	// aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether HTTPS is enabled for the cloud resource. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	EnableHttps *int32 `json:"EnableHttps,omitempty" xml:"EnableHttps,omitempty"`
	// The time when the cloud resource was created. The time is a timestamp in seconds.
	//
	// example:
	//
	// 1673423339000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the cloud resource was last modified. The time is in the timestamp format.
	//
	// example:
	//
	// 1681956830000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the cloud resource.
	//
	// example:
	//
	// 20979
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID of the cloud resource.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// cas-cn-m7r1qocw91at
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The listener ID of the cloud resource.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// lsn-vwdff0q20poq5xazb9@443
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The listening port of the cloud resource.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// 8047
	ListenerPort *string `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The region ID of the cloud resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The other metadata related to the cloud resource.
	//
	// example:
	//
	// {\\"camera_model\\":\\"GIFSHOW [1267087617][OnePlus
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the cloud resource.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether an Alibaba Cloud SSL certificate is used. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// >  This parameter is required only when you deploy certificates to services of multiple clouds.
	//
	// example:
	//
	// 1
	UseSsl *int32 `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1666884372152785
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListDeploymentJobResourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentJobResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobResourceResponseBodyData) GetCertEndTime() *string {
	return s.CertEndTime
}

func (s *ListDeploymentJobResourceResponseBodyData) GetCertId() *int64 {
	return s.CertId
}

func (s *ListDeploymentJobResourceResponseBodyData) GetCertName() *string {
	return s.CertName
}

func (s *ListDeploymentJobResourceResponseBodyData) GetCertStartTime() *string {
	return s.CertStartTime
}

func (s *ListDeploymentJobResourceResponseBodyData) GetCloudAccessId() *string {
	return s.CloudAccessId
}

func (s *ListDeploymentJobResourceResponseBodyData) GetCloudName() *string {
	return s.CloudName
}

func (s *ListDeploymentJobResourceResponseBodyData) GetCloudProduct() *string {
	return s.CloudProduct
}

func (s *ListDeploymentJobResourceResponseBodyData) GetCloudRegion() *string {
	return s.CloudRegion
}

func (s *ListDeploymentJobResourceResponseBodyData) GetDefaultResource() *int32 {
	return s.DefaultResource
}

func (s *ListDeploymentJobResourceResponseBodyData) GetDomain() *string {
	return s.Domain
}

func (s *ListDeploymentJobResourceResponseBodyData) GetEnableHttps() *int32 {
	return s.EnableHttps
}

func (s *ListDeploymentJobResourceResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListDeploymentJobResourceResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDeploymentJobResourceResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListDeploymentJobResourceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDeploymentJobResourceResponseBodyData) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListDeploymentJobResourceResponseBodyData) GetListenerPort() *string {
	return s.ListenerPort
}

func (s *ListDeploymentJobResourceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDeploymentJobResourceResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *ListDeploymentJobResourceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListDeploymentJobResourceResponseBodyData) GetUseSsl() *int32 {
	return s.UseSsl
}

func (s *ListDeploymentJobResourceResponseBodyData) GetUserId() *int64 {
	return s.UserId
}

func (s *ListDeploymentJobResourceResponseBodyData) SetCertEndTime(v string) *ListDeploymentJobResourceResponseBodyData {
	s.CertEndTime = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetCertId(v int64) *ListDeploymentJobResourceResponseBodyData {
	s.CertId = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetCertName(v string) *ListDeploymentJobResourceResponseBodyData {
	s.CertName = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetCertStartTime(v string) *ListDeploymentJobResourceResponseBodyData {
	s.CertStartTime = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetCloudAccessId(v string) *ListDeploymentJobResourceResponseBodyData {
	s.CloudAccessId = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetCloudName(v string) *ListDeploymentJobResourceResponseBodyData {
	s.CloudName = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetCloudProduct(v string) *ListDeploymentJobResourceResponseBodyData {
	s.CloudProduct = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetCloudRegion(v string) *ListDeploymentJobResourceResponseBodyData {
	s.CloudRegion = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetDefaultResource(v int32) *ListDeploymentJobResourceResponseBodyData {
	s.DefaultResource = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetDomain(v string) *ListDeploymentJobResourceResponseBodyData {
	s.Domain = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetEnableHttps(v int32) *ListDeploymentJobResourceResponseBodyData {
	s.EnableHttps = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetGmtCreate(v string) *ListDeploymentJobResourceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetGmtModified(v string) *ListDeploymentJobResourceResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetId(v int64) *ListDeploymentJobResourceResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetInstanceId(v string) *ListDeploymentJobResourceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetListenerId(v string) *ListDeploymentJobResourceResponseBodyData {
	s.ListenerId = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetListenerPort(v string) *ListDeploymentJobResourceResponseBodyData {
	s.ListenerPort = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetRegionId(v string) *ListDeploymentJobResourceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetRemark(v string) *ListDeploymentJobResourceResponseBodyData {
	s.Remark = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetStatus(v string) *ListDeploymentJobResourceResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetUseSsl(v int32) *ListDeploymentJobResourceResponseBodyData {
	s.UseSsl = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetUserId(v int64) *ListDeploymentJobResourceResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
