// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseContactDeleteRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseContactDeleteRequest
  GetAppName() *string 
  SetContactId(v int64) *EnterpriseContactDeleteRequest
  GetContactId() *int64 
  SetOrientedEcId(v string) *EnterpriseContactDeleteRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseContactDeleteRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseContactDeleteRequest
  GetOrientedNbId() *string 
}

type EnterpriseContactDeleteRequest struct {
  // example:
  // 
  // xxx
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  // example:
  // 
  // xxx
  ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
  // example:
  // 
  // null
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  // example:
  // 
  // null
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  // example:
  // 
  // null
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
}

func (s EnterpriseContactDeleteRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactDeleteRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseContactDeleteRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseContactDeleteRequest) GetContactId() *int64  {
  return s.ContactId
}

func (s *EnterpriseContactDeleteRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseContactDeleteRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseContactDeleteRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseContactDeleteRequest) SetAppName(v string) *EnterpriseContactDeleteRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseContactDeleteRequest) SetContactId(v int64) *EnterpriseContactDeleteRequest {
  s.ContactId = &v
  return s
}

func (s *EnterpriseContactDeleteRequest) SetOrientedEcId(v string) *EnterpriseContactDeleteRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseContactDeleteRequest) SetOrientedLeId(v string) *EnterpriseContactDeleteRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseContactDeleteRequest) SetOrientedNbId(v string) *EnterpriseContactDeleteRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseContactDeleteRequest) Validate() error {
  return dara.Validate(s)
}

