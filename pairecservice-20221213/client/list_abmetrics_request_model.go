// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListABMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListABMetricsRequest
	GetInstanceId() *string
	SetName(v string) *ListABMetricsRequest
	GetName() *string
	SetPageNumber(v int32) *ListABMetricsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListABMetricsRequest
	GetPageSize() *int32
	SetRealtime(v bool) *ListABMetricsRequest
	GetRealtime() *bool
	SetSceneId(v string) *ListABMetricsRequest
	GetSceneId() *string
	SetTableMetaId(v string) *ListABMetricsRequest
	GetTableMetaId() *string
	SetType(v string) *ListABMetricsRequest
	GetType() *string
}

type ListABMetricsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// home
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// false
	Realtime *bool `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 1
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// example:
	//
	// Single
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListABMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListABMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListABMetricsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListABMetricsRequest) GetName() *string {
	return s.Name
}

func (s *ListABMetricsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListABMetricsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListABMetricsRequest) GetRealtime() *bool {
	return s.Realtime
}

func (s *ListABMetricsRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *ListABMetricsRequest) GetTableMetaId() *string {
	return s.TableMetaId
}

func (s *ListABMetricsRequest) GetType() *string {
	return s.Type
}

func (s *ListABMetricsRequest) SetInstanceId(v string) *ListABMetricsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListABMetricsRequest) SetName(v string) *ListABMetricsRequest {
	s.Name = &v
	return s
}

func (s *ListABMetricsRequest) SetPageNumber(v int32) *ListABMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListABMetricsRequest) SetPageSize(v int32) *ListABMetricsRequest {
	s.PageSize = &v
	return s
}

func (s *ListABMetricsRequest) SetRealtime(v bool) *ListABMetricsRequest {
	s.Realtime = &v
	return s
}

func (s *ListABMetricsRequest) SetSceneId(v string) *ListABMetricsRequest {
	s.SceneId = &v
	return s
}

func (s *ListABMetricsRequest) SetTableMetaId(v string) *ListABMetricsRequest {
	s.TableMetaId = &v
	return s
}

func (s *ListABMetricsRequest) SetType(v string) *ListABMetricsRequest {
	s.Type = &v
	return s
}

func (s *ListABMetricsRequest) Validate() error {
	return dara.Validate(s)
}
