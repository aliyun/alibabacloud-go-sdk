// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeARMServerInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAICSpecs(v []*string) *DescribeARMServerInstancesRequest
	GetAICSpecs() []*string
	SetDescribeAICInstances(v bool) *DescribeARMServerInstancesRequest
	GetDescribeAICInstances() *bool
	SetEnsRegionIds(v []*string) *DescribeARMServerInstancesRequest
	GetEnsRegionIds() []*string
	SetMaxDate(v string) *DescribeARMServerInstancesRequest
	GetMaxDate() *string
	SetMinDate(v string) *DescribeARMServerInstancesRequest
	GetMinDate() *string
	SetName(v string) *DescribeARMServerInstancesRequest
	GetName() *string
	SetNamespace(v string) *DescribeARMServerInstancesRequest
	GetNamespace() *string
	SetOrderByParams(v string) *DescribeARMServerInstancesRequest
	GetOrderByParams() *string
	SetPageNumber(v int32) *DescribeARMServerInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeARMServerInstancesRequest
	GetPageSize() *int32
	SetServerIds(v []*string) *DescribeARMServerInstancesRequest
	GetServerIds() []*string
	SetServerSpecs(v []*string) *DescribeARMServerInstancesRequest
	GetServerSpecs() []*string
	SetStates(v []*string) *DescribeARMServerInstancesRequest
	GetStates() []*string
	SetTags(v []*DescribeARMServerInstancesRequestTags) *DescribeARMServerInstancesRequest
	GetTags() []*DescribeARMServerInstancesRequestTags
}

type DescribeARMServerInstancesRequest struct {
	// The list of container specifications.
	AICSpecs []*string `json:"AICSpecs,omitempty" xml:"AICSpecs,omitempty" type:"Repeated"`
	// Specifies whether the query results include container information. Valid values:
	//
	// - **true**: Included.
	//
	// - **false**: Not included.
	//
	// example:
	//
	// true
	DescribeAICInstances *bool `json:"DescribeAICInstances,omitempty" xml:"DescribeAICInstances,omitempty"`
	// The list of ENS edge zone IDs.
	EnsRegionIds []*string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty" type:"Repeated"`
	// The end time of the query. Format: "2006-01-02". By default, no time limit is applied.
	//
	// example:
	//
	// 2006-01-03
	MaxDate *string `json:"MaxDate,omitempty" xml:"MaxDate,omitempty"`
	// The start time of the query. Format: "2006-01-02". By default, no time limit is applied.
	//
	// example:
	//
	// 2006-01-02
	MinDate *string `json:"MinDate,omitempty" xml:"MinDate,omitempty"`
	// The server name.
	//
	// example:
	//
	// Server-Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The sorting method for query results. You can sort results by one or more of the following fields: ServerIdSort (server ID), ServerNameSort (server name), ExpireTimeSort (expiration time), CreationTimeSort (creation time), and EnsRegionIdSort (edge zone ID).
	//
	// Valid values: asc (ascending order) and desc (descending order).
	//
	// example:
	//
	// {"ServerIdSort":"desc","ServerNameSort":"desc","ExpireTimeSort":"asc"}
	OrderByParams *string `json:"OrderByParams,omitempty" xml:"OrderByParams,omitempty"`
	// The page number of the query list. Minimum value: **1**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. Maximum value: **100**.
	//
	// Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of server instance IDs.
	ServerIds []*string `json:"ServerIds,omitempty" xml:"ServerIds,omitempty" type:"Repeated"`
	// The list of server specifications.
	ServerSpecs []*string `json:"ServerSpecs,omitempty" xml:"ServerSpecs,omitempty" type:"Repeated"`
	// The list of operation states.
	States []*string `json:"States,omitempty" xml:"States,omitempty" type:"Repeated"`
	// The list of tags.
	Tags []*DescribeARMServerInstancesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeARMServerInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeARMServerInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeARMServerInstancesRequest) GetAICSpecs() []*string {
	return s.AICSpecs
}

func (s *DescribeARMServerInstancesRequest) GetDescribeAICInstances() *bool {
	return s.DescribeAICInstances
}

func (s *DescribeARMServerInstancesRequest) GetEnsRegionIds() []*string {
	return s.EnsRegionIds
}

func (s *DescribeARMServerInstancesRequest) GetMaxDate() *string {
	return s.MaxDate
}

func (s *DescribeARMServerInstancesRequest) GetMinDate() *string {
	return s.MinDate
}

func (s *DescribeARMServerInstancesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeARMServerInstancesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeARMServerInstancesRequest) GetOrderByParams() *string {
	return s.OrderByParams
}

func (s *DescribeARMServerInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeARMServerInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeARMServerInstancesRequest) GetServerIds() []*string {
	return s.ServerIds
}

func (s *DescribeARMServerInstancesRequest) GetServerSpecs() []*string {
	return s.ServerSpecs
}

func (s *DescribeARMServerInstancesRequest) GetStates() []*string {
	return s.States
}

func (s *DescribeARMServerInstancesRequest) GetTags() []*DescribeARMServerInstancesRequestTags {
	return s.Tags
}

func (s *DescribeARMServerInstancesRequest) SetAICSpecs(v []*string) *DescribeARMServerInstancesRequest {
	s.AICSpecs = v
	return s
}

func (s *DescribeARMServerInstancesRequest) SetDescribeAICInstances(v bool) *DescribeARMServerInstancesRequest {
	s.DescribeAICInstances = &v
	return s
}

func (s *DescribeARMServerInstancesRequest) SetEnsRegionIds(v []*string) *DescribeARMServerInstancesRequest {
	s.EnsRegionIds = v
	return s
}

func (s *DescribeARMServerInstancesRequest) SetMaxDate(v string) *DescribeARMServerInstancesRequest {
	s.MaxDate = &v
	return s
}

func (s *DescribeARMServerInstancesRequest) SetMinDate(v string) *DescribeARMServerInstancesRequest {
	s.MinDate = &v
	return s
}

func (s *DescribeARMServerInstancesRequest) SetName(v string) *DescribeARMServerInstancesRequest {
	s.Name = &v
	return s
}

func (s *DescribeARMServerInstancesRequest) SetNamespace(v string) *DescribeARMServerInstancesRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeARMServerInstancesRequest) SetOrderByParams(v string) *DescribeARMServerInstancesRequest {
	s.OrderByParams = &v
	return s
}

func (s *DescribeARMServerInstancesRequest) SetPageNumber(v int32) *DescribeARMServerInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeARMServerInstancesRequest) SetPageSize(v int32) *DescribeARMServerInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeARMServerInstancesRequest) SetServerIds(v []*string) *DescribeARMServerInstancesRequest {
	s.ServerIds = v
	return s
}

func (s *DescribeARMServerInstancesRequest) SetServerSpecs(v []*string) *DescribeARMServerInstancesRequest {
	s.ServerSpecs = v
	return s
}

func (s *DescribeARMServerInstancesRequest) SetStates(v []*string) *DescribeARMServerInstancesRequest {
	s.States = v
	return s
}

func (s *DescribeARMServerInstancesRequest) SetTags(v []*DescribeARMServerInstancesRequestTags) *DescribeARMServerInstancesRequest {
	s.Tags = v
	return s
}

func (s *DescribeARMServerInstancesRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeARMServerInstancesRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeARMServerInstancesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeARMServerInstancesRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeARMServerInstancesRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeARMServerInstancesRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeARMServerInstancesRequestTags) SetKey(v string) *DescribeARMServerInstancesRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeARMServerInstancesRequestTags) SetValue(v string) *DescribeARMServerInstancesRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeARMServerInstancesRequestTags) Validate() error {
	return dara.Validate(s)
}
