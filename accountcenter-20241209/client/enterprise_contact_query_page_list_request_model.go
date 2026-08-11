// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseContactQueryPageListRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseContactQueryPageListRequest
  GetAppName() *string 
  SetOrientedEcId(v string) *EnterpriseContactQueryPageListRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseContactQueryPageListRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseContactQueryPageListRequest
  GetOrientedNbId() *string 
  SetPageNo(v int32) *EnterpriseContactQueryPageListRequest
  GetPageNo() *int32 
  SetPageSize(v int32) *EnterpriseContactQueryPageListRequest
  GetPageSize() *int32 
  SetPrivateContact(v bool) *EnterpriseContactQueryPageListRequest
  GetPrivateContact() *bool 
  SetQuery(v string) *EnterpriseContactQueryPageListRequest
  GetQuery() *string 
  SetSharedContact(v bool) *EnterpriseContactQueryPageListRequest
  GetSharedContact() *bool 
  SetShowCompleteInfo(v bool) *EnterpriseContactQueryPageListRequest
  GetShowCompleteInfo() *bool 
}

type EnterpriseContactQueryPageListRequest struct {
  // The client application name.
  // 
  // example:
  // 
  // xxx
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  // The entity ID of the cross-enterprise management object.
  // 
  // example:
  // 
  // null
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  // The currently switched enterprise.
  // 
  // example:
  // 
  // null
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  // The marketplace ID of the cross-enterprise management object.
  // 
  // example:
  // 
  // null
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  // Page number, default 1.
  // 
  // example:
  // 
  // 1
  PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
  // Page size, default 10.
  // 
  // example:
  // 
  // 10
  PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
  // The flag for querying private contacts. The default value of this parameter is set to false for this API.
  // 
  // example:
  // 
  // false
  PrivateContact *bool `json:"PrivateContact,omitempty" xml:"PrivateContact,omitempty"`
  // Name, email, or mobile number.
  // 
  // example:
  // 
  // 1xxxxxxxxxx
  Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
  // Whether it is a shared contact. The default value of this parameter is set to true for this API.
  // 
  // example:
  // 
  // true
  SharedContact *bool `json:"SharedContact,omitempty" xml:"SharedContact,omitempty"`
  // Whether to display complete information.
  // 
  // example:
  // 
  // false
  ShowCompleteInfo *bool `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
}

func (s EnterpriseContactQueryPageListRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactQueryPageListRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseContactQueryPageListRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseContactQueryPageListRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseContactQueryPageListRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseContactQueryPageListRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseContactQueryPageListRequest) GetPageNo() *int32  {
  return s.PageNo
}

func (s *EnterpriseContactQueryPageListRequest) GetPageSize() *int32  {
  return s.PageSize
}

func (s *EnterpriseContactQueryPageListRequest) GetPrivateContact() *bool  {
  return s.PrivateContact
}

func (s *EnterpriseContactQueryPageListRequest) GetQuery() *string  {
  return s.Query
}

func (s *EnterpriseContactQueryPageListRequest) GetSharedContact() *bool  {
  return s.SharedContact
}

func (s *EnterpriseContactQueryPageListRequest) GetShowCompleteInfo() *bool  {
  return s.ShowCompleteInfo
}

func (s *EnterpriseContactQueryPageListRequest) SetAppName(v string) *EnterpriseContactQueryPageListRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseContactQueryPageListRequest) SetOrientedEcId(v string) *EnterpriseContactQueryPageListRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseContactQueryPageListRequest) SetOrientedLeId(v string) *EnterpriseContactQueryPageListRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseContactQueryPageListRequest) SetOrientedNbId(v string) *EnterpriseContactQueryPageListRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseContactQueryPageListRequest) SetPageNo(v int32) *EnterpriseContactQueryPageListRequest {
  s.PageNo = &v
  return s
}

func (s *EnterpriseContactQueryPageListRequest) SetPageSize(v int32) *EnterpriseContactQueryPageListRequest {
  s.PageSize = &v
  return s
}

func (s *EnterpriseContactQueryPageListRequest) SetPrivateContact(v bool) *EnterpriseContactQueryPageListRequest {
  s.PrivateContact = &v
  return s
}

func (s *EnterpriseContactQueryPageListRequest) SetQuery(v string) *EnterpriseContactQueryPageListRequest {
  s.Query = &v
  return s
}

func (s *EnterpriseContactQueryPageListRequest) SetSharedContact(v bool) *EnterpriseContactQueryPageListRequest {
  s.SharedContact = &v
  return s
}

func (s *EnterpriseContactQueryPageListRequest) SetShowCompleteInfo(v bool) *EnterpriseContactQueryPageListRequest {
  s.ShowCompleteInfo = &v
  return s
}

func (s *EnterpriseContactQueryPageListRequest) Validate() error {
  return dara.Validate(s)
}

