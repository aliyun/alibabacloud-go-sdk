// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseContactQueryDetailRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseContactQueryDetailRequest
  GetAppName() *string 
  SetContactId(v int64) *EnterpriseContactQueryDetailRequest
  GetContactId() *int64 
  SetOrientedEcId(v string) *EnterpriseContactQueryDetailRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseContactQueryDetailRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseContactQueryDetailRequest
  GetOrientedNbId() *string 
}

type EnterpriseContactQueryDetailRequest struct {
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

func (s EnterpriseContactQueryDetailRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactQueryDetailRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseContactQueryDetailRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseContactQueryDetailRequest) GetContactId() *int64  {
  return s.ContactId
}

func (s *EnterpriseContactQueryDetailRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseContactQueryDetailRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseContactQueryDetailRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseContactQueryDetailRequest) SetAppName(v string) *EnterpriseContactQueryDetailRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseContactQueryDetailRequest) SetContactId(v int64) *EnterpriseContactQueryDetailRequest {
  s.ContactId = &v
  return s
}

func (s *EnterpriseContactQueryDetailRequest) SetOrientedEcId(v string) *EnterpriseContactQueryDetailRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseContactQueryDetailRequest) SetOrientedLeId(v string) *EnterpriseContactQueryDetailRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseContactQueryDetailRequest) SetOrientedNbId(v string) *EnterpriseContactQueryDetailRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseContactQueryDetailRequest) Validate() error {
  return dara.Validate(s)
}

