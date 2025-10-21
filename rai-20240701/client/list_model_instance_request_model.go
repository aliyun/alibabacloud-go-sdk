// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEasServiceName(v string) *ListModelInstanceRequest
	GetEasServiceName() *string
	SetIsSidecarPolicy(v int32) *ListModelInstanceRequest
	GetIsSidecarPolicy() *int32
	SetIsSupportContentSafe(v int32) *ListModelInstanceRequest
	GetIsSupportContentSafe() *int32
	SetIsSupportPromptAttack(v int32) *ListModelInstanceRequest
	GetIsSupportPromptAttack() *int32
	SetIsSupportSensitiveTopic(v int32) *ListModelInstanceRequest
	GetIsSupportSensitiveTopic() *int32
	SetModelSource(v string) *ListModelInstanceRequest
	GetModelSource() *string
	SetOrder(v string) *ListModelInstanceRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListModelInstanceRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelInstanceRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListModelInstanceRequest
	GetRegionId() *string
	SetSortBy(v string) *ListModelInstanceRequest
	GetSortBy() *string
	SetWorkspaceId(v int64) *ListModelInstanceRequest
	GetWorkspaceId() *int64
}

type ListModelInstanceRequest struct {
	// example:
	//
	// rai_content_detection_model
	EasServiceName *string `json:"EasServiceName,omitempty" xml:"EasServiceName,omitempty"`
	// example:
	//
	// 1
	IsSidecarPolicy *int32 `json:"IsSidecarPolicy,omitempty" xml:"IsSidecarPolicy,omitempty"`
	// example:
	//
	// True
	IsSupportContentSafe *int32 `json:"IsSupportContentSafe,omitempty" xml:"IsSupportContentSafe,omitempty"`
	// example:
	//
	// False
	IsSupportPromptAttack *int32 `json:"IsSupportPromptAttack,omitempty" xml:"IsSupportPromptAttack,omitempty"`
	// example:
	//
	// True
	IsSupportSensitiveTopic *int32 `json:"IsSupportSensitiveTopic,omitempty" xml:"IsSupportSensitiveTopic,omitempty"`
	// example:
	//
	// 1
	ModelSource *string `json:"ModelSource,omitempty" xml:"ModelSource,omitempty"`
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// GmtModified
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 643168
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListModelInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListModelInstanceRequest) GetEasServiceName() *string {
	return s.EasServiceName
}

func (s *ListModelInstanceRequest) GetIsSidecarPolicy() *int32 {
	return s.IsSidecarPolicy
}

func (s *ListModelInstanceRequest) GetIsSupportContentSafe() *int32 {
	return s.IsSupportContentSafe
}

func (s *ListModelInstanceRequest) GetIsSupportPromptAttack() *int32 {
	return s.IsSupportPromptAttack
}

func (s *ListModelInstanceRequest) GetIsSupportSensitiveTopic() *int32 {
	return s.IsSupportSensitiveTopic
}

func (s *ListModelInstanceRequest) GetModelSource() *string {
	return s.ModelSource
}

func (s *ListModelInstanceRequest) GetOrder() *string {
	return s.Order
}

func (s *ListModelInstanceRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelInstanceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListModelInstanceRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListModelInstanceRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListModelInstanceRequest) SetEasServiceName(v string) *ListModelInstanceRequest {
	s.EasServiceName = &v
	return s
}

func (s *ListModelInstanceRequest) SetIsSidecarPolicy(v int32) *ListModelInstanceRequest {
	s.IsSidecarPolicy = &v
	return s
}

func (s *ListModelInstanceRequest) SetIsSupportContentSafe(v int32) *ListModelInstanceRequest {
	s.IsSupportContentSafe = &v
	return s
}

func (s *ListModelInstanceRequest) SetIsSupportPromptAttack(v int32) *ListModelInstanceRequest {
	s.IsSupportPromptAttack = &v
	return s
}

func (s *ListModelInstanceRequest) SetIsSupportSensitiveTopic(v int32) *ListModelInstanceRequest {
	s.IsSupportSensitiveTopic = &v
	return s
}

func (s *ListModelInstanceRequest) SetModelSource(v string) *ListModelInstanceRequest {
	s.ModelSource = &v
	return s
}

func (s *ListModelInstanceRequest) SetOrder(v string) *ListModelInstanceRequest {
	s.Order = &v
	return s
}

func (s *ListModelInstanceRequest) SetPageNumber(v int32) *ListModelInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelInstanceRequest) SetPageSize(v int32) *ListModelInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelInstanceRequest) SetRegionId(v string) *ListModelInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ListModelInstanceRequest) SetSortBy(v string) *ListModelInstanceRequest {
	s.SortBy = &v
	return s
}

func (s *ListModelInstanceRequest) SetWorkspaceId(v int64) *ListModelInstanceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListModelInstanceRequest) Validate() error {
	return dara.Validate(s)
}
