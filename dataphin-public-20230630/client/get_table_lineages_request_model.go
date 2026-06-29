// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableLineagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterQuery(v *GetTableLineagesRequestFilterQuery) *GetTableLineagesRequest
	GetFilterQuery() *GetTableLineagesRequestFilterQuery
	SetOpTenantId(v int64) *GetTableLineagesRequest
	GetOpTenantId() *int64
	SetTableGuid(v string) *GetTableLineagesRequest
	GetTableGuid() *string
}

type GetTableLineagesRequest struct {
	// The filter conditions.
	FilterQuery *GetTableLineagesRequestFilterQuery `json:"FilterQuery,omitempty" xml:"FilterQuery,omitempty" type:"Struct"`
	// Tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The GUID of the table, which is the unique identifier of each asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1121
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s GetTableLineagesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableLineagesRequest) GoString() string {
	return s.String()
}

func (s *GetTableLineagesRequest) GetFilterQuery() *GetTableLineagesRequestFilterQuery {
	return s.FilterQuery
}

func (s *GetTableLineagesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetTableLineagesRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetTableLineagesRequest) SetFilterQuery(v *GetTableLineagesRequestFilterQuery) *GetTableLineagesRequest {
	s.FilterQuery = v
	return s
}

func (s *GetTableLineagesRequest) SetOpTenantId(v int64) *GetTableLineagesRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetTableLineagesRequest) SetTableGuid(v string) *GetTableLineagesRequest {
	s.TableGuid = &v
	return s
}

func (s *GetTableLineagesRequest) Validate() error {
	if s.FilterQuery != nil {
		if err := s.FilterQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTableLineagesRequestFilterQuery struct {
	// Specifies whether to query downstream lineage. Default value: false.
	NeedDownstream *bool `json:"NeedDownstream,omitempty" xml:"NeedDownstream,omitempty"`
	// Specifies whether to return tables that do not exist in the asset list. Default value: false.
	NeedNotExistObject *bool `json:"NeedNotExistObject,omitempty" xml:"NeedNotExistObject,omitempty"`
	// Specifies whether to query upstream lineage. Default value: false.
	NeedUpstream *bool `json:"NeedUpstream,omitempty" xml:"NeedUpstream,omitempty"`
	// The environment to which the task belongs. This parameter is used for filtering. Valid values: dev and prod.
	//
	// example:
	//
	// dev
	NodeEnv *string `json:"NodeEnv,omitempty" xml:"NodeEnv,omitempty"`
	// The list of task IDs used for filtering.
	NodeIdList []*string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
}

func (s GetTableLineagesRequestFilterQuery) String() string {
	return dara.Prettify(s)
}

func (s GetTableLineagesRequestFilterQuery) GoString() string {
	return s.String()
}

func (s *GetTableLineagesRequestFilterQuery) GetNeedDownstream() *bool {
	return s.NeedDownstream
}

func (s *GetTableLineagesRequestFilterQuery) GetNeedNotExistObject() *bool {
	return s.NeedNotExistObject
}

func (s *GetTableLineagesRequestFilterQuery) GetNeedUpstream() *bool {
	return s.NeedUpstream
}

func (s *GetTableLineagesRequestFilterQuery) GetNodeEnv() *string {
	return s.NodeEnv
}

func (s *GetTableLineagesRequestFilterQuery) GetNodeIdList() []*string {
	return s.NodeIdList
}

func (s *GetTableLineagesRequestFilterQuery) SetNeedDownstream(v bool) *GetTableLineagesRequestFilterQuery {
	s.NeedDownstream = &v
	return s
}

func (s *GetTableLineagesRequestFilterQuery) SetNeedNotExistObject(v bool) *GetTableLineagesRequestFilterQuery {
	s.NeedNotExistObject = &v
	return s
}

func (s *GetTableLineagesRequestFilterQuery) SetNeedUpstream(v bool) *GetTableLineagesRequestFilterQuery {
	s.NeedUpstream = &v
	return s
}

func (s *GetTableLineagesRequestFilterQuery) SetNodeEnv(v string) *GetTableLineagesRequestFilterQuery {
	s.NodeEnv = &v
	return s
}

func (s *GetTableLineagesRequestFilterQuery) SetNodeIdList(v []*string) *GetTableLineagesRequestFilterQuery {
	s.NodeIdList = v
	return s
}

func (s *GetTableLineagesRequestFilterQuery) Validate() error {
	return dara.Validate(s)
}
