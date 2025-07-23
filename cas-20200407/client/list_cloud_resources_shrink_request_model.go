// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdsShrink(v string) *ListCloudResourcesShrinkRequest
	GetCertIdsShrink() *string
	SetCloudName(v string) *ListCloudResourcesShrinkRequest
	GetCloudName() *string
	SetCloudProduct(v string) *ListCloudResourcesShrinkRequest
	GetCloudProduct() *string
	SetCurrentPage(v int32) *ListCloudResourcesShrinkRequest
	GetCurrentPage() *int32
	SetKeyword(v string) *ListCloudResourcesShrinkRequest
	GetKeyword() *string
	SetSecretId(v string) *ListCloudResourcesShrinkRequest
	GetSecretId() *string
	SetShowSize(v int32) *ListCloudResourcesShrinkRequest
	GetShowSize() *int32
}

type ListCloudResourcesShrinkRequest struct {
	// The certificate IDs.
	CertIdsShrink *string `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
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

func (s ListCloudResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCloudResourcesShrinkRequest) GetCertIdsShrink() *string {
	return s.CertIdsShrink
}

func (s *ListCloudResourcesShrinkRequest) GetCloudName() *string {
	return s.CloudName
}

func (s *ListCloudResourcesShrinkRequest) GetCloudProduct() *string {
	return s.CloudProduct
}

func (s *ListCloudResourcesShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCloudResourcesShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCloudResourcesShrinkRequest) GetSecretId() *string {
	return s.SecretId
}

func (s *ListCloudResourcesShrinkRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListCloudResourcesShrinkRequest) SetCertIdsShrink(v string) *ListCloudResourcesShrinkRequest {
	s.CertIdsShrink = &v
	return s
}

func (s *ListCloudResourcesShrinkRequest) SetCloudName(v string) *ListCloudResourcesShrinkRequest {
	s.CloudName = &v
	return s
}

func (s *ListCloudResourcesShrinkRequest) SetCloudProduct(v string) *ListCloudResourcesShrinkRequest {
	s.CloudProduct = &v
	return s
}

func (s *ListCloudResourcesShrinkRequest) SetCurrentPage(v int32) *ListCloudResourcesShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudResourcesShrinkRequest) SetKeyword(v string) *ListCloudResourcesShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListCloudResourcesShrinkRequest) SetSecretId(v string) *ListCloudResourcesShrinkRequest {
	s.SecretId = &v
	return s
}

func (s *ListCloudResourcesShrinkRequest) SetShowSize(v int32) *ListCloudResourcesShrinkRequest {
	s.ShowSize = &v
	return s
}

func (s *ListCloudResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
