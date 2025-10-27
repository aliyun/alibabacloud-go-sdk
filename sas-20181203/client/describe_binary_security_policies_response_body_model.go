// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBinarySecurityPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBinarySecurityPolicies(v []*DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies) *DescribeBinarySecurityPoliciesResponseBody
	GetBinarySecurityPolicies() []*DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies
	SetPageInfo(v *DescribeBinarySecurityPoliciesResponseBodyPageInfo) *DescribeBinarySecurityPoliciesResponseBody
	GetPageInfo() *DescribeBinarySecurityPoliciesResponseBodyPageInfo
	SetRequestId(v string) *DescribeBinarySecurityPoliciesResponseBody
	GetRequestId() *string
}

type DescribeBinarySecurityPoliciesResponseBody struct {
	// The information about security policies.
	BinarySecurityPolicies []*DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies `json:"BinarySecurityPolicies,omitempty" xml:"BinarySecurityPolicies,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeBinarySecurityPoliciesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1EE7B150-D67E-53FD-A52D-3E8E669A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBinarySecurityPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinarySecurityPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBinarySecurityPoliciesResponseBody) GetBinarySecurityPolicies() []*DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies {
	return s.BinarySecurityPolicies
}

func (s *DescribeBinarySecurityPoliciesResponseBody) GetPageInfo() *DescribeBinarySecurityPoliciesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeBinarySecurityPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBinarySecurityPoliciesResponseBody) SetBinarySecurityPolicies(v []*DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies) *DescribeBinarySecurityPoliciesResponseBody {
	s.BinarySecurityPolicies = v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponseBody) SetPageInfo(v *DescribeBinarySecurityPoliciesResponseBodyPageInfo) *DescribeBinarySecurityPoliciesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponseBody) SetRequestId(v string) *DescribeBinarySecurityPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponseBody) Validate() error {
	if s.BinarySecurityPolicies != nil {
		for _, item := range s.BinarySecurityPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies struct {
	// The information about clusters.
	Clusters []*DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPoliciesClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The name of the policy.
	//
	// example:
	//
	// logtail
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The content of the policy. The value is in the JSON format. A key supports the following values:
	//
	// 	- **policyMode**: the type of the policy. Default value: requireAttestor.
	//
	// 	- **requiredAttestors**: the required witnesses.
	//
	// example:
	//
	// {\\"PolicyMode\\":\\"requireAttestor\\",\\"RequiredAttestors\\":[\\"test-xcs-04-11-hhht\\"]}
	Policy map[string]interface{} `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the policy. Valid values:
	//
	// 	- **enabled**
	//
	// 	- **disabled**
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies) GoString() string {
	return s.String()
}

func (s *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies) GetClusters() []*DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPoliciesClusters {
	return s.Clusters
}

func (s *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies) GetName() *string {
	return s.Name
}

func (s *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies) GetPolicy() map[string]interface{} {
	return s.Policy
}

func (s *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies) GetRemark() *string {
	return s.Remark
}

func (s *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies) GetStatus() *string {
	return s.Status
}

func (s *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies) SetClusters(v []*DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPoliciesClusters) *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies {
	s.Clusters = v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies) SetName(v string) *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies {
	s.Name = &v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies) SetPolicy(v map[string]interface{}) *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies {
	s.Policy = v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies) SetRemark(v string) *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies {
	s.Remark = &v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies) SetStatus(v string) *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies {
	s.Status = &v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPolicies) Validate() error {
	if s.Clusters != nil {
		for _, item := range s.Clusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPoliciesClusters struct {
	// The ID of the cluster.
	//
	// example:
	//
	// c316702acdf5f45e1a9dc7fc52f21****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The namespaces.
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
}

func (s DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPoliciesClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPoliciesClusters) GoString() string {
	return s.String()
}

func (s *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPoliciesClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPoliciesClusters) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPoliciesClusters) SetClusterId(v string) *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPoliciesClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPoliciesClusters) SetNamespaces(v []*string) *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPoliciesClusters {
	s.Namespaces = v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponseBodyBinarySecurityPoliciesClusters) Validate() error {
	return dara.Validate(s)
}

type DescribeBinarySecurityPoliciesResponseBodyPageInfo struct {
	// The number of entries returned on each page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 218
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBinarySecurityPoliciesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinarySecurityPoliciesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeBinarySecurityPoliciesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeBinarySecurityPoliciesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeBinarySecurityPoliciesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBinarySecurityPoliciesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBinarySecurityPoliciesResponseBodyPageInfo) SetCount(v int32) *DescribeBinarySecurityPoliciesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeBinarySecurityPoliciesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponseBodyPageInfo) SetPageSize(v int32) *DescribeBinarySecurityPoliciesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeBinarySecurityPoliciesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
