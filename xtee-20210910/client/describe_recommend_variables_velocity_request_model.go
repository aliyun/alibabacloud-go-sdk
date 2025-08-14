// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendVariablesVelocityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRecommendVariablesVelocityRequest
	GetLang() *string
	SetRegId(v string) *DescribeRecommendVariablesVelocityRequest
	GetRegId() *string
	SetTaskId(v int64) *DescribeRecommendVariablesVelocityRequest
	GetTaskId() *int64
	SetVariableIdsStr(v string) *DescribeRecommendVariablesVelocityRequest
	GetVariableIdsStr() *string
}

type DescribeRecommendVariablesVelocityRequest struct {
	// Set the language type for request and response, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region Code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Task ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 6770764
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// Variable IDs
	//
	// example:
	//
	// [\\"232\\",\\"233\\"]
	VariableIdsStr *string `json:"variableIdsStr,omitempty" xml:"variableIdsStr,omitempty"`
}

func (s DescribeRecommendVariablesVelocityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendVariablesVelocityRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecommendVariablesVelocityRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRecommendVariablesVelocityRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRecommendVariablesVelocityRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeRecommendVariablesVelocityRequest) GetVariableIdsStr() *string {
	return s.VariableIdsStr
}

func (s *DescribeRecommendVariablesVelocityRequest) SetLang(v string) *DescribeRecommendVariablesVelocityRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRecommendVariablesVelocityRequest) SetRegId(v string) *DescribeRecommendVariablesVelocityRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRecommendVariablesVelocityRequest) SetTaskId(v int64) *DescribeRecommendVariablesVelocityRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeRecommendVariablesVelocityRequest) SetVariableIdsStr(v string) *DescribeRecommendVariablesVelocityRequest {
	s.VariableIdsStr = &v
	return s
}

func (s *DescribeRecommendVariablesVelocityRequest) Validate() error {
	return dara.Validate(s)
}
