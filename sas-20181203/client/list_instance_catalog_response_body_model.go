// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceCatalogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListInstanceCatalogResponseBody
	GetRequestId() *string
	SetVendors(v []*ListInstanceCatalogResponseBodyVendors) *ListInstanceCatalogResponseBody
	GetVendors() []*ListInstanceCatalogResponseBodyVendors
}

type ListInstanceCatalogResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0D42A83F-CE33-5F54-A5AE-05DA39F59E1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The asset types by service provider.
	Vendors []*ListInstanceCatalogResponseBodyVendors `json:"Vendors,omitempty" xml:"Vendors,omitempty" type:"Repeated"`
}

func (s ListInstanceCatalogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceCatalogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceCatalogResponseBody) GetVendors() []*ListInstanceCatalogResponseBodyVendors {
	return s.Vendors
}

func (s *ListInstanceCatalogResponseBody) SetRequestId(v string) *ListInstanceCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceCatalogResponseBody) SetVendors(v []*ListInstanceCatalogResponseBodyVendors) *ListInstanceCatalogResponseBody {
	s.Vendors = v
	return s
}

func (s *ListInstanceCatalogResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstanceCatalogResponseBodyVendors struct {
	// The asset types.
	InstanceTypes []*ListInstanceCatalogResponseBodyVendorsInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The name of the service provider.
	//
	// example:
	//
	// ALIYUN
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the service provider type. Valid values:
	//
	// 	- **0**: an asset provided by Alibaba Cloud
	//
	// 	- **1**: an asset outside Alibaba Cloud
	//
	// 	- **2**: an asset in a data center
	//
	// 	- **3**, **4**, **5**, and **7**: an asset from a third-party cloud service provider
	//
	// 	- **8**: a lightweight cloud asset
	//
	// example:
	//
	// 0
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstanceCatalogResponseBodyVendors) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceCatalogResponseBodyVendors) GoString() string {
	return s.String()
}

func (s *ListInstanceCatalogResponseBodyVendors) GetInstanceTypes() []*ListInstanceCatalogResponseBodyVendorsInstanceTypes {
	return s.InstanceTypes
}

func (s *ListInstanceCatalogResponseBodyVendors) GetName() *string {
	return s.Name
}

func (s *ListInstanceCatalogResponseBodyVendors) GetValue() *int32 {
	return s.Value
}

func (s *ListInstanceCatalogResponseBodyVendors) SetInstanceTypes(v []*ListInstanceCatalogResponseBodyVendorsInstanceTypes) *ListInstanceCatalogResponseBodyVendors {
	s.InstanceTypes = v
	return s
}

func (s *ListInstanceCatalogResponseBodyVendors) SetName(v string) *ListInstanceCatalogResponseBodyVendors {
	s.Name = &v
	return s
}

func (s *ListInstanceCatalogResponseBodyVendors) SetValue(v int32) *ListInstanceCatalogResponseBodyVendors {
	s.Value = &v
	return s
}

func (s *ListInstanceCatalogResponseBodyVendors) Validate() error {
	return dara.Validate(s)
}

type ListInstanceCatalogResponseBodyVendorsInstanceTypes struct {
	// The asset subtypes.
	InstanceSubTypes []*ListInstanceCatalogResponseBodyVendorsInstanceTypesInstanceSubTypes `json:"InstanceSubTypes,omitempty" xml:"InstanceSubTypes,omitempty" type:"Repeated"`
	// The name of the asset type.
	//
	// example:
	//
	// ECS
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the asset type. Valid values:
	//
	// 	- **0**: Elastic Compute Service (ECS)
	//
	// 	- **1**: Server Load Balancer (SLB)
	//
	// 	- **3**: ApsaraDB RDS
	//
	// 	- **4**: ApsaraDB for MongoDB (MongoDB)
	//
	// 	- **5**: Tair (Redis OSS-compatible)
	//
	// 	- **6**: Container Registry
	//
	// 	- **8**: Container Service for Kubernetes (ACK)
	//
	// 	- **9**: Virtual Private Cloud (VPC)
	//
	// 	- **11**: ActionTrail
	//
	// 	- **12**: Alibaba Cloud CDN (CDN)
	//
	// 	- **13**: Certificate Management Service (formerly SSL Certificates Service)
	//
	// 	- **14**: Alibaba Cloud DevOps
	//
	// 	- **15**: Resource Access Management (RAM)
	//
	// 	- **16**: Anti-DDoS
	//
	// 	- **17**: Web Application Firewall (WAF)
	//
	// 	- **18**: Object Storage Service (OSS)
	//
	// 	- **19**: PolarDB
	//
	// 	- **20**: ApsaraDB RDS for PostgreSQL
	//
	// 	- **21**: Microservices Engine (MSE)
	//
	// 	- **22**: Apsara File Storage NAS (NAS)
	//
	// 	- **23**: Data Security Center (DSC)
	//
	// 	- **24**: Elastic IP Address (EIP)
	//
	// 	- **25**: Identity as a Service (IDaaS) - Enterprise Identity Access Management (EIAM)
	//
	// 	- **26**: PolarDB for Xscale (PolarDB-X)
	//
	// 	- **27**: Elasticsearch
	//
	// example:
	//
	// 15
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstanceCatalogResponseBodyVendorsInstanceTypes) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceCatalogResponseBodyVendorsInstanceTypes) GoString() string {
	return s.String()
}

func (s *ListInstanceCatalogResponseBodyVendorsInstanceTypes) GetInstanceSubTypes() []*ListInstanceCatalogResponseBodyVendorsInstanceTypesInstanceSubTypes {
	return s.InstanceSubTypes
}

func (s *ListInstanceCatalogResponseBodyVendorsInstanceTypes) GetName() *string {
	return s.Name
}

func (s *ListInstanceCatalogResponseBodyVendorsInstanceTypes) GetValue() *int32 {
	return s.Value
}

func (s *ListInstanceCatalogResponseBodyVendorsInstanceTypes) SetInstanceSubTypes(v []*ListInstanceCatalogResponseBodyVendorsInstanceTypesInstanceSubTypes) *ListInstanceCatalogResponseBodyVendorsInstanceTypes {
	s.InstanceSubTypes = v
	return s
}

func (s *ListInstanceCatalogResponseBodyVendorsInstanceTypes) SetName(v string) *ListInstanceCatalogResponseBodyVendorsInstanceTypes {
	s.Name = &v
	return s
}

func (s *ListInstanceCatalogResponseBodyVendorsInstanceTypes) SetValue(v int32) *ListInstanceCatalogResponseBodyVendorsInstanceTypes {
	s.Value = &v
	return s
}

func (s *ListInstanceCatalogResponseBodyVendorsInstanceTypes) Validate() error {
	return dara.Validate(s)
}

type ListInstanceCatalogResponseBodyVendorsInstanceTypesInstanceSubTypes struct {
	// The name of the asset subtype.
	//
	// example:
	//
	// SECURITY_GROUP
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the asset subtype.
	//
	// example:
	//
	// 1
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstanceCatalogResponseBodyVendorsInstanceTypesInstanceSubTypes) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceCatalogResponseBodyVendorsInstanceTypesInstanceSubTypes) GoString() string {
	return s.String()
}

func (s *ListInstanceCatalogResponseBodyVendorsInstanceTypesInstanceSubTypes) GetName() *string {
	return s.Name
}

func (s *ListInstanceCatalogResponseBodyVendorsInstanceTypesInstanceSubTypes) GetValue() *int32 {
	return s.Value
}

func (s *ListInstanceCatalogResponseBodyVendorsInstanceTypesInstanceSubTypes) SetName(v string) *ListInstanceCatalogResponseBodyVendorsInstanceTypesInstanceSubTypes {
	s.Name = &v
	return s
}

func (s *ListInstanceCatalogResponseBodyVendorsInstanceTypesInstanceSubTypes) SetValue(v int32) *ListInstanceCatalogResponseBodyVendorsInstanceTypesInstanceSubTypes {
	s.Value = &v
	return s
}

func (s *ListInstanceCatalogResponseBodyVendorsInstanceTypesInstanceSubTypes) Validate() error {
	return dara.Validate(s)
}
