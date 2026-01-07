// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleQueryBizRoleByPageRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEncryptTicket(v string) *EnterpriseRoleQueryBizRoleByPageRequest
  GetEncryptTicket() *string 
  SetMaxResults(v int32) *EnterpriseRoleQueryBizRoleByPageRequest
  GetMaxResults() *int32 
  SetNextToken(v string) *EnterpriseRoleQueryBizRoleByPageRequest
  GetNextToken() *string 
  SetOrientedEcId(v string) *EnterpriseRoleQueryBizRoleByPageRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseRoleQueryBizRoleByPageRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseRoleQueryBizRoleByPageRequest
  GetOrientedNbId() *string 
  SetPageNo(v int32) *EnterpriseRoleQueryBizRoleByPageRequest
  GetPageNo() *int32 
  SetPageSize(v int32) *EnterpriseRoleQueryBizRoleByPageRequest
  GetPageSize() *int32 
  SetQuery(v string) *EnterpriseRoleQueryBizRoleByPageRequest
  GetQuery() *string 
  SetResourceType(v string) *EnterpriseRoleQueryBizRoleByPageRequest
  GetResourceType() *string 
  SetSrcType(v string) *EnterpriseRoleQueryBizRoleByPageRequest
  GetSrcType() *string 
}

type EnterpriseRoleQueryBizRoleByPageRequest struct {
  EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
  PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
  Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
  ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
  SrcType *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleByPageRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleByPageRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) GetEncryptTicket() *string  {
  return s.EncryptTicket
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) GetPageNo() *int32  {
  return s.PageNo
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) GetPageSize() *int32  {
  return s.PageSize
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) GetQuery() *string  {
  return s.Query
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) GetResourceType() *string  {
  return s.ResourceType
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) GetSrcType() *string  {
  return s.SrcType
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetEncryptTicket(v string) *EnterpriseRoleQueryBizRoleByPageRequest {
  s.EncryptTicket = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetMaxResults(v int32) *EnterpriseRoleQueryBizRoleByPageRequest {
  s.MaxResults = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetNextToken(v string) *EnterpriseRoleQueryBizRoleByPageRequest {
  s.NextToken = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetOrientedEcId(v string) *EnterpriseRoleQueryBizRoleByPageRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetOrientedLeId(v string) *EnterpriseRoleQueryBizRoleByPageRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetOrientedNbId(v string) *EnterpriseRoleQueryBizRoleByPageRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetPageNo(v int32) *EnterpriseRoleQueryBizRoleByPageRequest {
  s.PageNo = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetPageSize(v int32) *EnterpriseRoleQueryBizRoleByPageRequest {
  s.PageSize = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetQuery(v string) *EnterpriseRoleQueryBizRoleByPageRequest {
  s.Query = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetResourceType(v string) *EnterpriseRoleQueryBizRoleByPageRequest {
  s.ResourceType = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) SetSrcType(v string) *EnterpriseRoleQueryBizRoleByPageRequest {
  s.SrcType = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageRequest) Validate() error {
  return dara.Validate(s)
}

