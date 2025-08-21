// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListInstanceRequest
	GetAgentKey() *string
	SetName(v string) *ListInstanceRequest
	GetName() *string
	SetPageNumber(v int64) *ListInstanceRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListInstanceRequest
	GetPageSize() *int64
	SetRobotType(v string) *ListInstanceRequest
	GetRobotType() *string
}

type ListInstanceRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// scenario_im
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// scenario_im
	RobotType *string `json:"RobotType,omitempty" xml:"RobotType,omitempty"`
}

func (s ListInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListInstanceRequest) GetName() *string {
	return s.Name
}

func (s *ListInstanceRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListInstanceRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListInstanceRequest) GetRobotType() *string {
	return s.RobotType
}

func (s *ListInstanceRequest) SetAgentKey(v string) *ListInstanceRequest {
	s.AgentKey = &v
	return s
}

func (s *ListInstanceRequest) SetName(v string) *ListInstanceRequest {
	s.Name = &v
	return s
}

func (s *ListInstanceRequest) SetPageNumber(v int64) *ListInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceRequest) SetPageSize(v int64) *ListInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceRequest) SetRobotType(v string) *ListInstanceRequest {
	s.RobotType = &v
	return s
}

func (s *ListInstanceRequest) Validate() error {
	return dara.Validate(s)
}
