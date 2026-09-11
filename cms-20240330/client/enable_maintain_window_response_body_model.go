// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMaintainWindowResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMaintainWindowId(v string) *EnableMaintainWindowResponseBody
  GetMaintainWindowId() *string 
  SetRequestId(v string) *EnableMaintainWindowResponseBody
  GetRequestId() *string 
}

type EnableMaintainWindowResponseBody struct {
  // example:
  // 
  // 123-12-312-31-23123
  MaintainWindowId *string `json:"maintainWindowId,omitempty" xml:"maintainWindowId,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 0CEC5375-C554-562B-A65F-9A629907C1F0
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s EnableMaintainWindowResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableMaintainWindowResponseBody) GoString() string {
  return s.String()
}

func (s *EnableMaintainWindowResponseBody) GetMaintainWindowId() *string  {
  return s.MaintainWindowId
}

func (s *EnableMaintainWindowResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableMaintainWindowResponseBody) SetMaintainWindowId(v string) *EnableMaintainWindowResponseBody {
  s.MaintainWindowId = &v
  return s
}

func (s *EnableMaintainWindowResponseBody) SetRequestId(v string) *EnableMaintainWindowResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableMaintainWindowResponseBody) Validate() error {
  return dara.Validate(s)
}

