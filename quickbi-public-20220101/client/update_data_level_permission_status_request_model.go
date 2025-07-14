// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLevelPermissionStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCubeId(v string) *UpdateDataLevelPermissionStatusRequest
	GetCubeId() *string
	SetIsOpen(v int32) *UpdateDataLevelPermissionStatusRequest
	GetIsOpen() *int32
	SetRuleType(v string) *UpdateDataLevelPermissionStatusRequest
	GetRuleType() *string
}

type UpdateDataLevelPermissionStatusRequest struct {
	// The ID of the training dataset that you want to remove from the specified custom linguistic model.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	IsOpen *int32 `json:"IsOpen,omitempty" xml:"IsOpen,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ROW_LEVEL
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s UpdateDataLevelPermissionStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLevelPermissionStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataLevelPermissionStatusRequest) GetCubeId() *string {
	return s.CubeId
}

func (s *UpdateDataLevelPermissionStatusRequest) GetIsOpen() *int32 {
	return s.IsOpen
}

func (s *UpdateDataLevelPermissionStatusRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *UpdateDataLevelPermissionStatusRequest) SetCubeId(v string) *UpdateDataLevelPermissionStatusRequest {
	s.CubeId = &v
	return s
}

func (s *UpdateDataLevelPermissionStatusRequest) SetIsOpen(v int32) *UpdateDataLevelPermissionStatusRequest {
	s.IsOpen = &v
	return s
}

func (s *UpdateDataLevelPermissionStatusRequest) SetRuleType(v string) *UpdateDataLevelPermissionStatusRequest {
	s.RuleType = &v
	return s
}

func (s *UpdateDataLevelPermissionStatusRequest) Validate() error {
	return dara.Validate(s)
}
