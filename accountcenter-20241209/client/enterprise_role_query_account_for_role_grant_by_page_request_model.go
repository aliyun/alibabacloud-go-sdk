// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleQueryAccountForRoleGrantByPageRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBizRoleCode(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest
  GetBizRoleCode() *string 
  SetEncryptTicket(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest
  GetEncryptTicket() *string 
  SetMaxResults(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest
  GetMaxResults() *int32 
  SetNextToken(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest
  GetNextToken() *string 
  SetOrgId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest
  GetOrgId() *string 
  SetOrientedEcId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest
  GetOrientedNbId() *string 
  SetPageNo(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest
  GetPageNo() *int32 
  SetPageSize(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest
  GetPageSize() *int32 
  SetQuery(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest
  GetQuery() *string 
  SetShowCompleteInfo(v bool) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest
  GetShowCompleteInfo() *bool 
}

type EnterpriseRoleQueryAccountForRoleGrantByPageRequest struct {
  BizRoleCode *string `json:"BizRoleCode,omitempty" xml:"BizRoleCode,omitempty"`
  EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
  OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
  PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
  Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
  ShowCompleteInfo *bool `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
}

func (s EnterpriseRoleQueryAccountForRoleGrantByPageRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleQueryAccountForRoleGrantByPageRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) GetBizRoleCode() *string  {
  return s.BizRoleCode
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) GetEncryptTicket() *string  {
  return s.EncryptTicket
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) GetOrgId() *string  {
  return s.OrgId
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) GetPageNo() *int32  {
  return s.PageNo
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) GetPageSize() *int32  {
  return s.PageSize
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) GetQuery() *string  {
  return s.Query
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) GetShowCompleteInfo() *bool  {
  return s.ShowCompleteInfo
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetBizRoleCode(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
  s.BizRoleCode = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetEncryptTicket(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
  s.EncryptTicket = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetMaxResults(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
  s.MaxResults = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetNextToken(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
  s.NextToken = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetOrgId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
  s.OrgId = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetOrientedEcId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetOrientedLeId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetOrientedNbId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetPageNo(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
  s.PageNo = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetPageSize(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
  s.PageSize = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetQuery(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
  s.Query = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) SetShowCompleteInfo(v bool) *EnterpriseRoleQueryAccountForRoleGrantByPageRequest {
  s.ShowCompleteInfo = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) Validate() error {
  return dara.Validate(s)
}

