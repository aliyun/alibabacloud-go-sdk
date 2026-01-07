// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountQueryAccountsInfoRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEncryptTicket(v string) *EnterpriseAccountQueryAccountsInfoRequest
  GetEncryptTicket() *string 
  SetMaxResults(v int32) *EnterpriseAccountQueryAccountsInfoRequest
  GetMaxResults() *int32 
  SetNextToken(v string) *EnterpriseAccountQueryAccountsInfoRequest
  GetNextToken() *string 
  SetOrientedEcId(v string) *EnterpriseAccountQueryAccountsInfoRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseAccountQueryAccountsInfoRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseAccountQueryAccountsInfoRequest
  GetOrientedNbId() *string 
  SetPksJson(v string) *EnterpriseAccountQueryAccountsInfoRequest
  GetPksJson() *string 
  SetRequestId(v string) *EnterpriseAccountQueryAccountsInfoRequest
  GetRequestId() *string 
  SetShowCompleteInfo(v bool) *EnterpriseAccountQueryAccountsInfoRequest
  GetShowCompleteInfo() *bool 
}

type EnterpriseAccountQueryAccountsInfoRequest struct {
  EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  // This parameter is required.
  PksJson *string `json:"PksJson,omitempty" xml:"PksJson,omitempty"`
  // This parameter is required.
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  ShowCompleteInfo *bool `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
}

func (s EnterpriseAccountQueryAccountsInfoRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountQueryAccountsInfoRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) GetEncryptTicket() *string  {
  return s.EncryptTicket
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) GetPksJson() *string  {
  return s.PksJson
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) GetShowCompleteInfo() *bool  {
  return s.ShowCompleteInfo
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetEncryptTicket(v string) *EnterpriseAccountQueryAccountsInfoRequest {
  s.EncryptTicket = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetMaxResults(v int32) *EnterpriseAccountQueryAccountsInfoRequest {
  s.MaxResults = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetNextToken(v string) *EnterpriseAccountQueryAccountsInfoRequest {
  s.NextToken = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetOrientedEcId(v string) *EnterpriseAccountQueryAccountsInfoRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetOrientedLeId(v string) *EnterpriseAccountQueryAccountsInfoRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetOrientedNbId(v string) *EnterpriseAccountQueryAccountsInfoRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetPksJson(v string) *EnterpriseAccountQueryAccountsInfoRequest {
  s.PksJson = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetRequestId(v string) *EnterpriseAccountQueryAccountsInfoRequest {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) SetShowCompleteInfo(v bool) *EnterpriseAccountQueryAccountsInfoRequest {
  s.ShowCompleteInfo = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoRequest) Validate() error {
  return dara.Validate(s)
}

