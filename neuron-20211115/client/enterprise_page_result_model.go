// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterprisePageResult interface {
  dara.Model
  String() string
  GoString() string
  SetData(v []*Enterprise) *EnterprisePageResult
  GetData() []*Enterprise 
  SetRequestId(v string) *EnterprisePageResult
  GetRequestId() *string 
  SetTotal(v int64) *EnterprisePageResult
  GetTotal() *int64 
}

type EnterprisePageResult struct {
  Data []*Enterprise `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // 24
  Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s EnterprisePageResult) String() string {
  return dara.Prettify(s)
}

func (s EnterprisePageResult) GoString() string {
  return s.String()
}

func (s *EnterprisePageResult) GetData() []*Enterprise  {
  return s.Data
}

func (s *EnterprisePageResult) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterprisePageResult) GetTotal() *int64  {
  return s.Total
}

func (s *EnterprisePageResult) SetData(v []*Enterprise) *EnterprisePageResult {
  s.Data = v
  return s
}

func (s *EnterprisePageResult) SetRequestId(v string) *EnterprisePageResult {
  s.RequestId = &v
  return s
}

func (s *EnterprisePageResult) SetTotal(v int64) *EnterprisePageResult {
  s.Total = &v
  return s
}

func (s *EnterprisePageResult) Validate() error {
  if s.Data != nil {
    for _, item := range s.Data {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

