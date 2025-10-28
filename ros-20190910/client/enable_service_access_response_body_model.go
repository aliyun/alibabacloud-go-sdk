// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableServiceAccessResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableServiceAccessResponseBody
  GetRequestId() *string 
  SetServiceAccessInfo(v *EnableServiceAccessResponseBodyServiceAccessInfo) *EnableServiceAccessResponseBody
  GetServiceAccessInfo() *EnableServiceAccessResponseBodyServiceAccessInfo 
}

type EnableServiceAccessResponseBody struct {
  // example:
  // 
  // 23045A5D-720E-5D11-A752-E1568F725C93
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  ServiceAccessInfo *EnableServiceAccessResponseBodyServiceAccessInfo `json:"ServiceAccessInfo,omitempty" xml:"ServiceAccessInfo,omitempty" type:"Struct"`
}

func (s EnableServiceAccessResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableServiceAccessResponseBody) GoString() string {
  return s.String()
}

func (s *EnableServiceAccessResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableServiceAccessResponseBody) GetServiceAccessInfo() *EnableServiceAccessResponseBodyServiceAccessInfo  {
  return s.ServiceAccessInfo
}

func (s *EnableServiceAccessResponseBody) SetRequestId(v string) *EnableServiceAccessResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableServiceAccessResponseBody) SetServiceAccessInfo(v *EnableServiceAccessResponseBodyServiceAccessInfo) *EnableServiceAccessResponseBody {
  s.ServiceAccessInfo = v
  return s
}

func (s *EnableServiceAccessResponseBody) Validate() error {
  if s.ServiceAccessInfo != nil {
    if err := s.ServiceAccessInfo.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnableServiceAccessResponseBodyServiceAccessInfo struct {
  // example:
  // 
  // ENABLED
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnableServiceAccessResponseBodyServiceAccessInfo) String() string {
  return dara.Prettify(s)
}

func (s EnableServiceAccessResponseBodyServiceAccessInfo) GoString() string {
  return s.String()
}

func (s *EnableServiceAccessResponseBodyServiceAccessInfo) GetStatus() *string  {
  return s.Status
}

func (s *EnableServiceAccessResponseBodyServiceAccessInfo) SetStatus(v string) *EnableServiceAccessResponseBodyServiceAccessInfo {
  s.Status = &v
  return s
}

func (s *EnableServiceAccessResponseBodyServiceAccessInfo) Validate() error {
  return dara.Validate(s)
}

