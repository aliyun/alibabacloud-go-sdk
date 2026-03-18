// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListPolicyRequest
	GetName() *string
	SetPageNo(v int64) *ListPolicyRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListPolicyRequest
	GetPageSize() *int64
	SetProductType(v string) *ListPolicyRequest
	GetProductType() *string
	SetType(v string) *ListPolicyRequest
	GetType() *string
}

type ListPolicyRequest struct {
	// The name of the policy.
	//
	// example:
	//
	// test**
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The service type. Valid values:
	//
	// 	- **ecs**: Elastic Compute Service (ECS).
	//
	// 	- **slb**: Server Load Balancer (SLB).
	//
	// 	- **eip**: Elastic IP Address (EIP).
	//
	// 	- **gf-eip**: EIP with Anti-DDoS (Enhanced) enabled.
	//
	// >  This parameter is available only if Type is set to `default`.
	//
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **default**: the default mitigation policy.
	//
	// 	- **l3**: IP-specific mitigation policies.
	//
	// 	- **l4**: port-specific mitigation policies.
	//
	// example:
	//
	// l3
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyRequest) GetName() *string {
	return s.Name
}

func (s *ListPolicyRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListPolicyRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPolicyRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListPolicyRequest) GetType() *string {
	return s.Type
}

func (s *ListPolicyRequest) SetName(v string) *ListPolicyRequest {
	s.Name = &v
	return s
}

func (s *ListPolicyRequest) SetPageNo(v int64) *ListPolicyRequest {
	s.PageNo = &v
	return s
}

func (s *ListPolicyRequest) SetPageSize(v int64) *ListPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *ListPolicyRequest) SetProductType(v string) *ListPolicyRequest {
	s.ProductType = &v
	return s
}

func (s *ListPolicyRequest) SetType(v string) *ListPolicyRequest {
	s.Type = &v
	return s
}

func (s *ListPolicyRequest) Validate() error {
	return dara.Validate(s)
}
