// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnLineagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterQuery(v *GetTableColumnLineagesRequestFilterQuery) *GetTableColumnLineagesRequest
	GetFilterQuery() *GetTableColumnLineagesRequestFilterQuery
	SetOpTenantId(v int64) *GetTableColumnLineagesRequest
	GetOpTenantId() *int64
	SetTableGuid(v string) *GetTableColumnLineagesRequest
	GetTableGuid() *string
}

type GetTableColumnLineagesRequest struct {
	FilterQuery *GetTableColumnLineagesRequestFilterQuery `json:"FilterQuery,omitempty" xml:"FilterQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1121
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s GetTableColumnLineagesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnLineagesRequest) GoString() string {
	return s.String()
}

func (s *GetTableColumnLineagesRequest) GetFilterQuery() *GetTableColumnLineagesRequestFilterQuery {
	return s.FilterQuery
}

func (s *GetTableColumnLineagesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetTableColumnLineagesRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetTableColumnLineagesRequest) SetFilterQuery(v *GetTableColumnLineagesRequestFilterQuery) *GetTableColumnLineagesRequest {
	s.FilterQuery = v
	return s
}

func (s *GetTableColumnLineagesRequest) SetOpTenantId(v int64) *GetTableColumnLineagesRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetTableColumnLineagesRequest) SetTableGuid(v string) *GetTableColumnLineagesRequest {
	s.TableGuid = &v
	return s
}

func (s *GetTableColumnLineagesRequest) Validate() error {
	if s.FilterQuery != nil {
		if err := s.FilterQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTableColumnLineagesRequestFilterQuery struct {
	NeedDownstream     *bool `json:"NeedDownstream,omitempty" xml:"NeedDownstream,omitempty"`
	NeedNotExistObject *bool `json:"NeedNotExistObject,omitempty" xml:"NeedNotExistObject,omitempty"`
	NeedUpstream       *bool `json:"NeedUpstream,omitempty" xml:"NeedUpstream,omitempty"`
	// example:
	//
	// dev
	NodeEnv    *string   `json:"NodeEnv,omitempty" xml:"NodeEnv,omitempty"`
	NodeIdList []*string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
}

func (s GetTableColumnLineagesRequestFilterQuery) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnLineagesRequestFilterQuery) GoString() string {
	return s.String()
}

func (s *GetTableColumnLineagesRequestFilterQuery) GetNeedDownstream() *bool {
	return s.NeedDownstream
}

func (s *GetTableColumnLineagesRequestFilterQuery) GetNeedNotExistObject() *bool {
	return s.NeedNotExistObject
}

func (s *GetTableColumnLineagesRequestFilterQuery) GetNeedUpstream() *bool {
	return s.NeedUpstream
}

func (s *GetTableColumnLineagesRequestFilterQuery) GetNodeEnv() *string {
	return s.NodeEnv
}

func (s *GetTableColumnLineagesRequestFilterQuery) GetNodeIdList() []*string {
	return s.NodeIdList
}

func (s *GetTableColumnLineagesRequestFilterQuery) SetNeedDownstream(v bool) *GetTableColumnLineagesRequestFilterQuery {
	s.NeedDownstream = &v
	return s
}

func (s *GetTableColumnLineagesRequestFilterQuery) SetNeedNotExistObject(v bool) *GetTableColumnLineagesRequestFilterQuery {
	s.NeedNotExistObject = &v
	return s
}

func (s *GetTableColumnLineagesRequestFilterQuery) SetNeedUpstream(v bool) *GetTableColumnLineagesRequestFilterQuery {
	s.NeedUpstream = &v
	return s
}

func (s *GetTableColumnLineagesRequestFilterQuery) SetNodeEnv(v string) *GetTableColumnLineagesRequestFilterQuery {
	s.NodeEnv = &v
	return s
}

func (s *GetTableColumnLineagesRequestFilterQuery) SetNodeIdList(v []*string) *GetTableColumnLineagesRequestFilterQuery {
	s.NodeIdList = v
	return s
}

func (s *GetTableColumnLineagesRequestFilterQuery) Validate() error {
	return dara.Validate(s)
}
