// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeploymentJobStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuyCount(v int32) *DescribeDeploymentJobStatusResponseBody
	GetBuyCount() *int32
	SetCertCount(v int32) *DescribeDeploymentJobStatusResponseBody
	GetCertCount() *int32
	SetCostCount(v int32) *DescribeDeploymentJobStatusResponseBody
	GetCostCount() *int32
	SetFailedCount(v int32) *DescribeDeploymentJobStatusResponseBody
	GetFailedCount() *int32
	SetMatchWorkerCount(v int32) *DescribeDeploymentJobStatusResponseBody
	GetMatchWorkerCount() *int32
	SetProductWorkerCount(v []*DescribeDeploymentJobStatusResponseBodyProductWorkerCount) *DescribeDeploymentJobStatusResponseBody
	GetProductWorkerCount() []*DescribeDeploymentJobStatusResponseBodyProductWorkerCount
	SetRequestId(v string) *DescribeDeploymentJobStatusResponseBody
	GetRequestId() *string
	SetResourceCount(v int32) *DescribeDeploymentJobStatusResponseBody
	GetResourceCount() *int32
	SetRollbackCount(v int32) *DescribeDeploymentJobStatusResponseBody
	GetRollbackCount() *int32
	SetRollbackFailedCount(v int32) *DescribeDeploymentJobStatusResponseBody
	GetRollbackFailedCount() *int32
	SetRollbackSuccessCount(v int32) *DescribeDeploymentJobStatusResponseBody
	GetRollbackSuccessCount() *int32
	SetSuccessCount(v int32) *DescribeDeploymentJobStatusResponseBody
	GetSuccessCount() *int32
	SetUseCount(v int32) *DescribeDeploymentJobStatusResponseBody
	GetUseCount() *int32
	SetWorkerCount(v int32) *DescribeDeploymentJobStatusResponseBody
	GetWorkerCount() *int32
}

type DescribeDeploymentJobStatusResponseBody struct {
	// The total number of purchased resources.
	//
	// example:
	//
	// 500
	BuyCount *int32 `json:"BuyCount,omitempty" xml:"BuyCount,omitempty"`
	// The number of certificates involved in the deployment task.
	//
	// example:
	//
	// 20
	CertCount *int32 `json:"CertCount,omitempty" xml:"CertCount,omitempty"`
	// The number of resources consumed by worker tasks.
	//
	// example:
	//
	// 20
	CostCount *int32 `json:"CostCount,omitempty" xml:"CostCount,omitempty"`
	// The number of failed worker tasks, excluding rollback tasks.
	//
	// example:
	//
	// 20
	FailedCount *int32 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The total number of worker tasks that match the certificate.
	//
	// example:
	//
	// 20
	MatchWorkerCount *int32 `json:"MatchWorkerCount,omitempty" xml:"MatchWorkerCount,omitempty"`
	// The number of cloud resources to which certificates are deployed in the deployment task.
	ProductWorkerCount []*DescribeDeploymentJobStatusResponseBodyProductWorkerCount `json:"ProductWorkerCount,omitempty" xml:"ProductWorkerCount,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of resources.
	//
	// example:
	//
	// 4127
	ResourceCount *int32 `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	// The number of worker tasks that are rolled backed.
	//
	// example:
	//
	// 20
	RollbackCount *int32 `json:"RollbackCount,omitempty" xml:"RollbackCount,omitempty"`
	// The number of worker tasks that failed to be rolled back.
	//
	// example:
	//
	// 20
	RollbackFailedCount *int32 `json:"RollbackFailedCount,omitempty" xml:"RollbackFailedCount,omitempty"`
	// The number of worker tasks that are successfully rolled back.
	//
	// example:
	//
	// 20
	RollbackSuccessCount *int32 `json:"RollbackSuccessCount,omitempty" xml:"RollbackSuccessCount,omitempty"`
	// The number of successful worker tasks, excluding rollback tasks.
	//
	// example:
	//
	// 20
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The total number of used resources.
	//
	// example:
	//
	// 300
	UseCount *int32 `json:"UseCount,omitempty" xml:"UseCount,omitempty"`
	// The total number of resources to which certificate are deployed in the current cloud service. The value indicates the total number of worker tasks.
	//
	// example:
	//
	// 20
	WorkerCount *int32 `json:"WorkerCount,omitempty" xml:"WorkerCount,omitempty"`
}

func (s DescribeDeploymentJobStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentJobStatusResponseBody) GetBuyCount() *int32 {
	return s.BuyCount
}

func (s *DescribeDeploymentJobStatusResponseBody) GetCertCount() *int32 {
	return s.CertCount
}

func (s *DescribeDeploymentJobStatusResponseBody) GetCostCount() *int32 {
	return s.CostCount
}

func (s *DescribeDeploymentJobStatusResponseBody) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *DescribeDeploymentJobStatusResponseBody) GetMatchWorkerCount() *int32 {
	return s.MatchWorkerCount
}

func (s *DescribeDeploymentJobStatusResponseBody) GetProductWorkerCount() []*DescribeDeploymentJobStatusResponseBodyProductWorkerCount {
	return s.ProductWorkerCount
}

func (s *DescribeDeploymentJobStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeploymentJobStatusResponseBody) GetResourceCount() *int32 {
	return s.ResourceCount
}

func (s *DescribeDeploymentJobStatusResponseBody) GetRollbackCount() *int32 {
	return s.RollbackCount
}

func (s *DescribeDeploymentJobStatusResponseBody) GetRollbackFailedCount() *int32 {
	return s.RollbackFailedCount
}

func (s *DescribeDeploymentJobStatusResponseBody) GetRollbackSuccessCount() *int32 {
	return s.RollbackSuccessCount
}

func (s *DescribeDeploymentJobStatusResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *DescribeDeploymentJobStatusResponseBody) GetUseCount() *int32 {
	return s.UseCount
}

func (s *DescribeDeploymentJobStatusResponseBody) GetWorkerCount() *int32 {
	return s.WorkerCount
}

func (s *DescribeDeploymentJobStatusResponseBody) SetBuyCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.BuyCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetCertCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.CertCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetCostCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.CostCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetFailedCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.FailedCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetMatchWorkerCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.MatchWorkerCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetProductWorkerCount(v []*DescribeDeploymentJobStatusResponseBodyProductWorkerCount) *DescribeDeploymentJobStatusResponseBody {
	s.ProductWorkerCount = v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetRequestId(v string) *DescribeDeploymentJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetResourceCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.ResourceCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetRollbackCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.RollbackCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetRollbackFailedCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.RollbackFailedCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetRollbackSuccessCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.RollbackSuccessCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetSuccessCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetUseCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.UseCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetWorkerCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.WorkerCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) Validate() error {
	if s.ProductWorkerCount != nil {
		for _, item := range s.ProductWorkerCount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDeploymentJobStatusResponseBodyProductWorkerCount struct {
	// The total number of resources of a cloud service in the deployment task.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the cloud service. Valid values:
	//
	// 	- **SLB**: Classic Load Balancer (CLB). This value is supported only at the China site (aliyun.com).
	//
	// 	- **LIVE**: ApsaraVideo Live. This value is supported only at the China site (aliyun.com).
	//
	// 	- **webHosting**: Cloud Web Hosting. This value is supported only at the China site (aliyun.com).
	//
	// 	- **VOD**: ApsaraVideo VOD. This value is supported only at the China site (aliyun.com).
	//
	// 	- **CR**: Container Registry. This value is supported only at the China site (aliyun.com).
	//
	// 	- **DCDN**: Dynamic Content Delivery Network (DCDN).
	//
	// 	- **DDOS**: Anti-DDoS.
	//
	// 	- **CDN**: Alibaba Cloud CDN (CDN).
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
	// NLB
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s DescribeDeploymentJobStatusResponseBodyProductWorkerCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentJobStatusResponseBodyProductWorkerCount) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentJobStatusResponseBodyProductWorkerCount) GetCount() *int32 {
	return s.Count
}

func (s *DescribeDeploymentJobStatusResponseBodyProductWorkerCount) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeDeploymentJobStatusResponseBodyProductWorkerCount) SetCount(v int32) *DescribeDeploymentJobStatusResponseBodyProductWorkerCount {
	s.Count = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBodyProductWorkerCount) SetProductName(v string) *DescribeDeploymentJobStatusResponseBodyProductWorkerCount {
	s.ProductName = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBodyProductWorkerCount) Validate() error {
	return dara.Validate(s)
}
