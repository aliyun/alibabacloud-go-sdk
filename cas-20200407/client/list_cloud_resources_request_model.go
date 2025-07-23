// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertIds(v []*int64) *ListCloudResourcesRequest
	GetCertIds() []*int64
	SetCloudName(v string) *ListCloudResourcesRequest
	GetCloudName() *string
	SetCloudProduct(v string) *ListCloudResourcesRequest
	GetCloudProduct() *string
	SetCurrentPage(v int32) *ListCloudResourcesRequest
	GetCurrentPage() *int32
	SetKeyword(v string) *ListCloudResourcesRequest
	GetKeyword() *string
	SetSecretId(v string) *ListCloudResourcesRequest
	GetSecretId() *string
	SetShowSize(v int32) *ListCloudResourcesRequest
	GetShowSize() *int32
}

type ListCloudResourcesRequest struct {
	// The certificate IDs.
	CertIds []*int64 `json:"CertIds,omitempty" xml:"CertIds,omitempty" type:"Repeated"`
	// The cloud service provider.
	//
	// Valid values:
	//
	// 	- Tencent
	//
	// 	- Huawei
	//
	// 	- Aws
	//
	// 	- aliyun
	//
	// example:
	//
	// Tencent
	CloudName *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	// The cloud service.
	//
	// Valid values when CloudName is set to aliyun:
	//
	// 	- SLB: Classic Load Balancer (CLB). This value is available only on the China site (aliyun.com).
	//
	// 	- LIVE: ApsaraVideo Live. This value is available only on the China site (aliyun.com).
	//
	// 	- webHosting: Cloud Web Hosting. This value is available only on the China site (aliyun.com).
	//
	// 	- VOD: ApsaraVideo VOD. This value is available only on the China site (aliyun.com).
	//
	// 	- CR: Container Registry. This value is available only on the China site (aliyun.com).
	//
	// 	- DCDN: Dynamic Content Delivery Network (DCDN).
	//
	// 	- DDOS: Anti-DDoS.
	//
	// 	- CDN: Alibaba Cloud CDN (CDN).
	//
	// 	- ALB: Application Load Balancer (ALB).
	//
	// 	- APIGateway: API Gateway.
	//
	// 	- FC: Function Compute.
	//
	// 	- GA: Global Accelerator (GA).
	//
	// 	- MSE: Microservices Engine (MSE).
	//
	// 	- NLB: Network Load Balancer (NLB).
	//
	// 	- OSS: Object Storage Service (OSS).
	//
	// 	- SAE: Serverless App Engine (SAE).
	//
	// 	- WAF: Web Application Firewall (WAF).
	//
	// Valid values when CloudName is set to Tencent:
	//
	// 	- TencentCDN: Content Delivery Network (CDN).
	//
	// 	- TencentCLB: CLB.
	//
	// 	- TencentWAF: WAF.
	//
	// Valid value when CloudName is set to Huawei:
	//
	// 	- HuaweiCDN: CDN.
	//
	// Valid values when CloudName is set to Aws:
	//
	// 	- AwsCloudFront: Amazon CloudFront.
	//
	// 	- AwsCLB: CLB.
	//
	// 	- AwsALB: ALB.
	//
	// 	- AwsNLB: NLB.
	//
	// example:
	//
	// SLB
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The keyword of the domain name or instance ID bound to the cloud resource.
	//
	// example:
	//
	// cert-instanceId
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The AccessKey ID that is used to access cloud resources.
	//
	// example:
	//
	// 21
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The number of entries per page. Default value: **20**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListCloudResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudResourcesRequest) GetCertIds() []*int64 {
	return s.CertIds
}

func (s *ListCloudResourcesRequest) GetCloudName() *string {
	return s.CloudName
}

func (s *ListCloudResourcesRequest) GetCloudProduct() *string {
	return s.CloudProduct
}

func (s *ListCloudResourcesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCloudResourcesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCloudResourcesRequest) GetSecretId() *string {
	return s.SecretId
}

func (s *ListCloudResourcesRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListCloudResourcesRequest) SetCertIds(v []*int64) *ListCloudResourcesRequest {
	s.CertIds = v
	return s
}

func (s *ListCloudResourcesRequest) SetCloudName(v string) *ListCloudResourcesRequest {
	s.CloudName = &v
	return s
}

func (s *ListCloudResourcesRequest) SetCloudProduct(v string) *ListCloudResourcesRequest {
	s.CloudProduct = &v
	return s
}

func (s *ListCloudResourcesRequest) SetCurrentPage(v int32) *ListCloudResourcesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudResourcesRequest) SetKeyword(v string) *ListCloudResourcesRequest {
	s.Keyword = &v
	return s
}

func (s *ListCloudResourcesRequest) SetSecretId(v string) *ListCloudResourcesRequest {
	s.SecretId = &v
	return s
}

func (s *ListCloudResourcesRequest) SetShowSize(v int32) *ListCloudResourcesRequest {
	s.ShowSize = &v
	return s
}

func (s *ListCloudResourcesRequest) Validate() error {
	return dara.Validate(s)
}
