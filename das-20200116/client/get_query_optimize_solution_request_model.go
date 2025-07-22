// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeSolutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *GetQueryOptimizeSolutionRequest
	GetEngine() *string
	SetInstanceId(v string) *GetQueryOptimizeSolutionRequest
	GetInstanceId() *string
	SetRuleIds(v string) *GetQueryOptimizeSolutionRequest
	GetRuleIds() *string
	SetSqlId(v string) *GetQueryOptimizeSolutionRequest
	GetSqlId() *string
}

type GetQueryOptimizeSolutionRequest struct {
	// The database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **PolarDBMySQL**
	//
	// 	- **PostgreSQL**
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The instance ID. You can call the [GetQueryOptimizeDataStats](https://help.aliyun.com/document_detail/405261.html) operation to query the instance ID.
	//
	// example:
	//
	// rm-bp1o3z6beqpej****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The tag ID. For more information, see [Query governance](https://help.aliyun.com/document_detail/290038.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// LARGE_ROWS_EXAMINED
	RuleIds *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
	// The SQL template ID. You can call the [GetQueryOptimizeDataStats](https://help.aliyun.com/document_detail/405261.html) operation to query the SQL template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 05fecf7e7b3efd123c4d5197035f****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
}

func (s GetQueryOptimizeSolutionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeSolutionRequest) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeSolutionRequest) GetEngine() *string {
	return s.Engine
}

func (s *GetQueryOptimizeSolutionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQueryOptimizeSolutionRequest) GetRuleIds() *string {
	return s.RuleIds
}

func (s *GetQueryOptimizeSolutionRequest) GetSqlId() *string {
	return s.SqlId
}

func (s *GetQueryOptimizeSolutionRequest) SetEngine(v string) *GetQueryOptimizeSolutionRequest {
	s.Engine = &v
	return s
}

func (s *GetQueryOptimizeSolutionRequest) SetInstanceId(v string) *GetQueryOptimizeSolutionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQueryOptimizeSolutionRequest) SetRuleIds(v string) *GetQueryOptimizeSolutionRequest {
	s.RuleIds = &v
	return s
}

func (s *GetQueryOptimizeSolutionRequest) SetSqlId(v string) *GetQueryOptimizeSolutionRequest {
	s.SqlId = &v
	return s
}

func (s *GetQueryOptimizeSolutionRequest) Validate() error {
	return dara.Validate(s)
}
