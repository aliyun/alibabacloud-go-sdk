// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeDownStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *ListNodeDownStreamRequest
	GetEnv() *string
	SetListQuery(v *ListNodeDownStreamRequestListQuery) *ListNodeDownStreamRequest
	GetListQuery() *ListNodeDownStreamRequestListQuery
	SetOpTenantId(v int64) *ListNodeDownStreamRequest
	GetOpTenantId() *int64
}

type ListNodeDownStreamRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	ListQuery *ListNodeDownStreamRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListNodeDownStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDownStreamRequest) GoString() string {
	return s.String()
}

func (s *ListNodeDownStreamRequest) GetEnv() *string {
	return s.Env
}

func (s *ListNodeDownStreamRequest) GetListQuery() *ListNodeDownStreamRequestListQuery {
	return s.ListQuery
}

func (s *ListNodeDownStreamRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListNodeDownStreamRequest) SetEnv(v string) *ListNodeDownStreamRequest {
	s.Env = &v
	return s
}

func (s *ListNodeDownStreamRequest) SetListQuery(v *ListNodeDownStreamRequestListQuery) *ListNodeDownStreamRequest {
	s.ListQuery = v
	return s
}

func (s *ListNodeDownStreamRequest) SetOpTenantId(v int64) *ListNodeDownStreamRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListNodeDownStreamRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodeDownStreamRequestListQuery struct {
	// example:
	//
	// 1
	DownStreamDepth *int32                                          `json:"DownStreamDepth,omitempty" xml:"DownStreamDepth,omitempty"`
	FilterList      []*ListNodeDownStreamRequestListQueryFilterList `json:"FilterList,omitempty" xml:"FilterList,omitempty" type:"Repeated"`
	// This parameter is required.
	NodeIdList []*ListNodeDownStreamRequestListQueryNodeIdList `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 123011
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListNodeDownStreamRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDownStreamRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListNodeDownStreamRequestListQuery) GetDownStreamDepth() *int32 {
	return s.DownStreamDepth
}

func (s *ListNodeDownStreamRequestListQuery) GetFilterList() []*ListNodeDownStreamRequestListQueryFilterList {
	return s.FilterList
}

func (s *ListNodeDownStreamRequestListQuery) GetNodeIdList() []*ListNodeDownStreamRequestListQueryNodeIdList {
	return s.NodeIdList
}

func (s *ListNodeDownStreamRequestListQuery) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListNodeDownStreamRequestListQuery) SetDownStreamDepth(v int32) *ListNodeDownStreamRequestListQuery {
	s.DownStreamDepth = &v
	return s
}

func (s *ListNodeDownStreamRequestListQuery) SetFilterList(v []*ListNodeDownStreamRequestListQueryFilterList) *ListNodeDownStreamRequestListQuery {
	s.FilterList = v
	return s
}

func (s *ListNodeDownStreamRequestListQuery) SetNodeIdList(v []*ListNodeDownStreamRequestListQueryNodeIdList) *ListNodeDownStreamRequestListQuery {
	s.NodeIdList = v
	return s
}

func (s *ListNodeDownStreamRequestListQuery) SetProjectId(v int64) *ListNodeDownStreamRequestListQuery {
	s.ProjectId = &v
	return s
}

func (s *ListNodeDownStreamRequestListQuery) Validate() error {
	if s.FilterList != nil {
		for _, item := range s.FilterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NodeIdList != nil {
		for _, item := range s.NodeIdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNodeDownStreamRequestListQueryFilterList struct {
	// example:
	//
	// false
	Exclude *bool `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// example:
	//
	// PROJECT
	Key       *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s ListNodeDownStreamRequestListQueryFilterList) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDownStreamRequestListQueryFilterList) GoString() string {
	return s.String()
}

func (s *ListNodeDownStreamRequestListQueryFilterList) GetExclude() *bool {
	return s.Exclude
}

func (s *ListNodeDownStreamRequestListQueryFilterList) GetKey() *string {
	return s.Key
}

func (s *ListNodeDownStreamRequestListQueryFilterList) GetValueList() []*string {
	return s.ValueList
}

func (s *ListNodeDownStreamRequestListQueryFilterList) SetExclude(v bool) *ListNodeDownStreamRequestListQueryFilterList {
	s.Exclude = &v
	return s
}

func (s *ListNodeDownStreamRequestListQueryFilterList) SetKey(v string) *ListNodeDownStreamRequestListQueryFilterList {
	s.Key = &v
	return s
}

func (s *ListNodeDownStreamRequestListQueryFilterList) SetValueList(v []*string) *ListNodeDownStreamRequestListQueryFilterList {
	s.ValueList = v
	return s
}

func (s *ListNodeDownStreamRequestListQueryFilterList) Validate() error {
	return dara.Validate(s)
}

type ListNodeDownStreamRequestListQueryNodeIdList struct {
	// example:
	//
	// 112
	FieldIdList []*string `json:"FieldIdList,omitempty" xml:"FieldIdList,omitempty" type:"Repeated"`
	// example:
	//
	// n_23431
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListNodeDownStreamRequestListQueryNodeIdList) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDownStreamRequestListQueryNodeIdList) GoString() string {
	return s.String()
}

func (s *ListNodeDownStreamRequestListQueryNodeIdList) GetFieldIdList() []*string {
	return s.FieldIdList
}

func (s *ListNodeDownStreamRequestListQueryNodeIdList) GetId() *string {
	return s.Id
}

func (s *ListNodeDownStreamRequestListQueryNodeIdList) SetFieldIdList(v []*string) *ListNodeDownStreamRequestListQueryNodeIdList {
	s.FieldIdList = v
	return s
}

func (s *ListNodeDownStreamRequestListQueryNodeIdList) SetId(v string) *ListNodeDownStreamRequestListQueryNodeIdList {
	s.Id = &v
	return s
}

func (s *ListNodeDownStreamRequestListQueryNodeIdList) Validate() error {
	return dara.Validate(s)
}
