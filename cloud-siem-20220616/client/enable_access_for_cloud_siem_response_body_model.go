// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAccessForCloudSiemResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v bool) *EnableAccessForCloudSiemResponseBody
  GetData() *bool 
  SetRequestId(v string) *EnableAccessForCloudSiemResponseBody
  GetRequestId() *string 
}

type EnableAccessForCloudSiemResponseBody struct {
  // The data returned.
  // 
  // example:
  // 
  // true
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 6276D891-*****-55B2-87B9-74D413F7****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAccessForCloudSiemResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableAccessForCloudSiemResponseBody) GoString() string {
  return s.String()
}

func (s *EnableAccessForCloudSiemResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnableAccessForCloudSiemResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableAccessForCloudSiemResponseBody) SetData(v bool) *EnableAccessForCloudSiemResponseBody {
  s.Data = &v
  return s
}

func (s *EnableAccessForCloudSiemResponseBody) SetRequestId(v string) *EnableAccessForCloudSiemResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableAccessForCloudSiemResponseBody) Validate() error {
  return dara.Validate(s)
}

