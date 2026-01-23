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
	FilterQuery *GetTableLineagesRequestFilterQuery `json:"FilterQuery,omitempty" xml:"FilterQuery,omitempty" type:"Struct"`
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
	NeedDownstream     *bool `json:"NeedDownstream,omitempty" xml:"NeedDownstream,omitempty"`
	NeedNotExistObject *bool `json:"NeedNotExistObject,omitempty" xml:"NeedNotExistObject,omitempty"`
	NeedUpstream       *bool `json:"NeedUpstream,omitempty" xml:"NeedUpstream,omitempty"`
	// example:
	//
	// dev
	NodeEnv    *string   `json:"NodeEnv,omitempty" xml:"NodeEnv,omitempty"`
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
