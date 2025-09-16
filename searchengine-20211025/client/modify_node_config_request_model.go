// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *ModifyNodeConfigRequest
	GetActive() *bool
	SetDataDuplicateNumber(v int32) *ModifyNodeConfigRequest
	GetDataDuplicateNumber() *int32
	SetDataFragmentNumber(v int32) *ModifyNodeConfigRequest
	GetDataFragmentNumber() *int32
	SetFlowRatio(v int32) *ModifyNodeConfigRequest
	GetFlowRatio() *int32
	SetMinServicePercent(v int32) *ModifyNodeConfigRequest
	GetMinServicePercent() *int32
	SetPublished(v bool) *ModifyNodeConfigRequest
	GetPublished() *bool
	SetClusterName(v string) *ModifyNodeConfigRequest
	GetClusterName() *string
	SetDataSourceName(v string) *ModifyNodeConfigRequest
	GetDataSourceName() *string
	SetName(v string) *ModifyNodeConfigRequest
	GetName() *string
	SetType(v string) *ModifyNodeConfigRequest
	GetType() *string
}

type ModifyNodeConfigRequest struct {
	// Specifies whether to enable the index.
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The number of data replicas.
	//
	// example:
	//
	// 1
	DataDuplicateNumber *int32 `json:"dataDuplicateNumber,omitempty" xml:"dataDuplicateNumber,omitempty"`
	// The number of data shards.
	//
	// example:
	//
	// 1
	DataFragmentNumber *int32 `json:"dataFragmentNumber,omitempty" xml:"dataFragmentNumber,omitempty"`
	// The traffic percentage.
	//
	// example:
	//
	// -100
	FlowRatio *int32 `json:"flowRatio,omitempty" xml:"flowRatio,omitempty"`
	// The minimum service ratio.
	//
	// example:
	//
	// 10
	MinServicePercent *int32 `json:"minServicePercent,omitempty" xml:"minServicePercent,omitempty"`
	// Specifies whether to mount the cluster.
	//
	// example:
	//
	// true
	Published *bool `json:"published,omitempty" xml:"published,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// vpc_sh_domain_2
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// The name of the data source. Valid values: -search: search for data. -not_search: do not search for data.
	//
	// example:
	//
	// ha-cn-2r42ostoc01_0704
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// The name of the configuration before the modification.
	//
	// This parameter is required.
	//
	// example:
	//
	// ha-cn-zvp2iv9a401_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the algorithm. Valid values:
	//
	// 	- pop: a popularity model.
	//
	// 	- cp: a category prediction model.
	//
	// 	- hot: a top search model.
	//
	// 	- hint: a hint model.
	//
	// 	- suggest: a drop-down suggestions model.
	//
	// This parameter is required.
	//
	// example:
	//
	// " "
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModifyNodeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodeConfigRequest) GetActive() *bool {
	return s.Active
}

func (s *ModifyNodeConfigRequest) GetDataDuplicateNumber() *int32 {
	return s.DataDuplicateNumber
}

func (s *ModifyNodeConfigRequest) GetDataFragmentNumber() *int32 {
	return s.DataFragmentNumber
}

func (s *ModifyNodeConfigRequest) GetFlowRatio() *int32 {
	return s.FlowRatio
}

func (s *ModifyNodeConfigRequest) GetMinServicePercent() *int32 {
	return s.MinServicePercent
}

func (s *ModifyNodeConfigRequest) GetPublished() *bool {
	return s.Published
}

func (s *ModifyNodeConfigRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ModifyNodeConfigRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ModifyNodeConfigRequest) GetName() *string {
	return s.Name
}

func (s *ModifyNodeConfigRequest) GetType() *string {
	return s.Type
}

func (s *ModifyNodeConfigRequest) SetActive(v bool) *ModifyNodeConfigRequest {
	s.Active = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetDataDuplicateNumber(v int32) *ModifyNodeConfigRequest {
	s.DataDuplicateNumber = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetDataFragmentNumber(v int32) *ModifyNodeConfigRequest {
	s.DataFragmentNumber = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetFlowRatio(v int32) *ModifyNodeConfigRequest {
	s.FlowRatio = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetMinServicePercent(v int32) *ModifyNodeConfigRequest {
	s.MinServicePercent = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetPublished(v bool) *ModifyNodeConfigRequest {
	s.Published = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetClusterName(v string) *ModifyNodeConfigRequest {
	s.ClusterName = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetDataSourceName(v string) *ModifyNodeConfigRequest {
	s.DataSourceName = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetName(v string) *ModifyNodeConfigRequest {
	s.Name = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetType(v string) *ModifyNodeConfigRequest {
	s.Type = &v
	return s
}

func (s *ModifyNodeConfigRequest) Validate() error {
	return dara.Validate(s)
}
