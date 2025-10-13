// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterRegionIds(v string) *ListPrometheusInstancesRequest
	GetFilterRegionIds() *string
	SetMaxResults(v int32) *ListPrometheusInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPrometheusInstancesRequest
	GetNextToken() *string
	SetPrometheusInstanceIds(v string) *ListPrometheusInstancesRequest
	GetPrometheusInstanceIds() *string
	SetPrometheusInstanceName(v string) *ListPrometheusInstancesRequest
	GetPrometheusInstanceName() *string
	SetResourceGroupId(v string) *ListPrometheusInstancesRequest
	GetResourceGroupId() *string
	SetResourceType(v string) *ListPrometheusInstancesRequest
	GetResourceType() *string
	SetTag(v []*ListPrometheusInstancesRequestTag) *ListPrometheusInstancesRequest
	GetTag() []*ListPrometheusInstancesRequestTag
	SetVersion(v string) *ListPrometheusInstancesRequest
	GetVersion() *string
}

type ListPrometheusInstancesRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// cn-shenzhen
	FilterRegionIds *string `json:"filterRegionIds,omitempty" xml:"filterRegionIds,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// xxxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// rw-00001,rw-00002,rw-00003
	PrometheusInstanceIds *string `json:"prometheusInstanceIds,omitempty" xml:"prometheusInstanceIds,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// test
	PrometheusInstanceName *string `json:"prometheusInstanceName,omitempty" xml:"prometheusInstanceName,omitempty"`
	// example:
	//
	// rg-aek2bhocin5e2na
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// Prometheus
	ResourceType *string                              `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Tag          []*ListPrometheusInstancesRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
	// if can be null:
	// true
	//
	// example:
	//
	// V2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListPrometheusInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstancesRequest) GetFilterRegionIds() *string {
	return s.FilterRegionIds
}

func (s *ListPrometheusInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPrometheusInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPrometheusInstancesRequest) GetPrometheusInstanceIds() *string {
	return s.PrometheusInstanceIds
}

func (s *ListPrometheusInstancesRequest) GetPrometheusInstanceName() *string {
	return s.PrometheusInstanceName
}

func (s *ListPrometheusInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPrometheusInstancesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListPrometheusInstancesRequest) GetTag() []*ListPrometheusInstancesRequestTag {
	return s.Tag
}

func (s *ListPrometheusInstancesRequest) GetVersion() *string {
	return s.Version
}

func (s *ListPrometheusInstancesRequest) SetFilterRegionIds(v string) *ListPrometheusInstancesRequest {
	s.FilterRegionIds = &v
	return s
}

func (s *ListPrometheusInstancesRequest) SetMaxResults(v int32) *ListPrometheusInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPrometheusInstancesRequest) SetNextToken(v string) *ListPrometheusInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPrometheusInstancesRequest) SetPrometheusInstanceIds(v string) *ListPrometheusInstancesRequest {
	s.PrometheusInstanceIds = &v
	return s
}

func (s *ListPrometheusInstancesRequest) SetPrometheusInstanceName(v string) *ListPrometheusInstancesRequest {
	s.PrometheusInstanceName = &v
	return s
}

func (s *ListPrometheusInstancesRequest) SetResourceGroupId(v string) *ListPrometheusInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPrometheusInstancesRequest) SetResourceType(v string) *ListPrometheusInstancesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListPrometheusInstancesRequest) SetTag(v []*ListPrometheusInstancesRequestTag) *ListPrometheusInstancesRequest {
	s.Tag = v
	return s
}

func (s *ListPrometheusInstancesRequest) SetVersion(v string) *ListPrometheusInstancesRequest {
	s.Version = &v
	return s
}

func (s *ListPrometheusInstancesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrometheusInstancesRequestTag struct {
	// example:
	//
	// testKey
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// testValue
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListPrometheusInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListPrometheusInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListPrometheusInstancesRequestTag) SetKey(v string) *ListPrometheusInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListPrometheusInstancesRequestTag) SetValue(v string) *ListPrometheusInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *ListPrometheusInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
