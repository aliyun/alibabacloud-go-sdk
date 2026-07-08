// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudResources(v []*DescribeCloudResourcesResponseBodyCloudResources) *DescribeCloudResourcesResponseBody
	GetCloudResources() []*DescribeCloudResourcesResponseBodyCloudResources
	SetRequestId(v string) *DescribeCloudResourcesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeCloudResourcesResponseBody
	GetTotalCount() *int64
}

type DescribeCloudResourcesResponseBody struct {
	// A list of cloud services.
	CloudResources []*DescribeCloudResourcesResponseBodyCloudResources `json:"CloudResources,omitempty" xml:"CloudResources,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C327F81A-CCE2-5B21-817C-F93E29C5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of cloud services found.
	//
	// example:
	//
	// 121
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourcesResponseBody) GetCloudResources() []*DescribeCloudResourcesResponseBodyCloudResources {
	return s.CloudResources
}

func (s *DescribeCloudResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudResourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeCloudResourcesResponseBody) SetCloudResources(v []*DescribeCloudResourcesResponseBodyCloudResources) *DescribeCloudResourcesResponseBody {
	s.CloudResources = v
	return s
}

func (s *DescribeCloudResourcesResponseBody) SetRequestId(v string) *DescribeCloudResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudResourcesResponseBody) SetTotalCount(v int64) *DescribeCloudResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCloudResourcesResponseBody) Validate() error {
	if s.CloudResources != nil {
		for _, item := range s.CloudResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudResourcesResponseBodyCloudResources struct {
	// The number of HTTP ports that are added to WAF.
	//
	// > This parameter is returned only for ECS and CLB instances.
	//
	// example:
	//
	// 1
	HttpPortCount *int32 `json:"HttpPortCount,omitempty" xml:"HttpPortCount,omitempty"`
	// The number of HTTPS ports that are added to WAF.
	//
	// > This parameter is returned only for ECS and CLB instances.
	//
	// example:
	//
	// 1
	HttpsPortCount *int32 `json:"HttpsPortCount,omitempty" xml:"HttpsPortCount,omitempty"`
	// The UID of the account to which the resource belongs.
	//
	// example:
	//
	// 11769793******
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The domain name of the resource. This parameter is returned for FC or SAE instances.
	//
	// example:
	//
	// test-resource-domain
	ResourceDomain *string `json:"ResourceDomain,omitempty" xml:"ResourceDomain,omitempty"`
	// The name of the function. This parameter is returned for FC instances.
	//
	// example:
	//
	// test-resource-function
	ResourceFunction *string `json:"ResourceFunction,omitempty" xml:"ResourceFunction,omitempty"`
	// Deprecated
	//
	// The ID of the resource instance.
	//
	// example:
	//
	// alb-ffff****
	ResourceInstance *string `json:"ResourceInstance,omitempty" xml:"ResourceInstance,omitempty"`
	// The ID of the instance that is added to WAF.
	//
	// example:
	//
	// lb-uf60ub45fr9b***
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The IP address of the instance that is added to WAF.
	//
	// example:
	//
	// 1.1.1.1
	ResourceInstanceIp *string `json:"ResourceInstanceIp,omitempty" xml:"ResourceInstanceIp,omitempty"`
	// The name of the instance that is added to WAF.
	//
	// example:
	//
	// test-name
	ResourceInstanceName *string `json:"ResourceInstanceName,omitempty" xml:"ResourceInstanceName,omitempty"`
	// Deprecated
	//
	// The name of the instance.
	//
	// example:
	//
	// test-resource-name
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The cloud service to which the resource belongs. Valid values:
	//
	// - **alb**: ALB
	//
	// - **mse**: MSE
	//
	// - **fc**: FC
	//
	// - **sae**: SAE
	//
	// - **ecs**: ECS
	//
	// - **clb4**: CLB that uses TCP
	//
	// - **clb7**: CLB that uses HTTP or HTTPS
	//
	// - **nlb**: NLB
	//
	// example:
	//
	// alb
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
	// The region where the resource resides.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The name of the route. This parameter is returned for MSE instances.
	//
	// example:
	//
	// test-route-name
	ResourceRouteName *string `json:"ResourceRouteName,omitempty" xml:"ResourceRouteName,omitempty"`
	// The name of the service. This parameter is returned for FC instances.
	//
	// example:
	//
	// test-resource-service
	ResourceService *string `json:"ResourceService,omitempty" xml:"ResourceService,omitempty"`
}

func (s DescribeCloudResourcesResponseBodyCloudResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourcesResponseBodyCloudResources) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) GetHttpPortCount() *int32 {
	return s.HttpPortCount
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) GetHttpsPortCount() *int32 {
	return s.HttpsPortCount
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) GetResourceDomain() *string {
	return s.ResourceDomain
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) GetResourceFunction() *string {
	return s.ResourceFunction
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) GetResourceInstance() *string {
	return s.ResourceInstance
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) GetResourceInstanceIp() *string {
	return s.ResourceInstanceIp
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) GetResourceInstanceName() *string {
	return s.ResourceInstanceName
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) GetResourceProduct() *string {
	return s.ResourceProduct
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) GetResourceRouteName() *string {
	return s.ResourceRouteName
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) GetResourceService() *string {
	return s.ResourceService
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetHttpPortCount(v int32) *DescribeCloudResourcesResponseBodyCloudResources {
	s.HttpPortCount = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetHttpsPortCount(v int32) *DescribeCloudResourcesResponseBodyCloudResources {
	s.HttpsPortCount = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetOwnerUserId(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceDomain(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceDomain = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceFunction(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceFunction = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceInstance(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceInstance = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceInstanceId(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceInstanceIp(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceInstanceIp = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceInstanceName(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceInstanceName = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceName(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceName = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceProduct(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceProduct = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceRegionId(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceRouteName(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceRouteName = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceService(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceService = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) Validate() error {
	return dara.Validate(s)
}
