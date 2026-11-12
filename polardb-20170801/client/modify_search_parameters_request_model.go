// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySearchParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifySearchParametersRequest
	GetDBClusterId() *string
	SetParameters(v string) *ModifySearchParametersRequest
	GetParameters() *string
}

type ModifySearchParametersRequest struct {
	// The ID of the PolarDB primary cluster, which is used to locate the Search cluster associated with the current account. Before you call this operation, you must enable the Search feature in the console. You can check the status by using the SearchClusterStatus field returned by DescribeDBClusterAttribute. Usage notes: Prerequisites: The cluster must have the Search feature enabled in the console. You can verify the status by using the SearchClusterStatus field returned by DescribeDBClusterAttribute.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-xxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// A non-empty JSON object string where the keys are parameter names and all values must be strings, such as 10% for percentages or 48mb for capacities. You can only modify parameters that are visible and modifiable in the template of the current version. Dynamic parameters take effect and are persisted immediately. Static parameters take effect after a restart triggered by the task flow. A success response only indicates that the task is accepted.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"indices.memory.index_buffer_size":"10%","indices.memory.min_index_buffer_size":"48mb"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s ModifySearchParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySearchParametersRequest) GoString() string {
	return s.String()
}

func (s *ModifySearchParametersRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifySearchParametersRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifySearchParametersRequest) SetDBClusterId(v string) *ModifySearchParametersRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifySearchParametersRequest) SetParameters(v string) *ModifySearchParametersRequest {
	s.Parameters = &v
	return s
}

func (s *ModifySearchParametersRequest) Validate() error {
	return dara.Validate(s)
}
