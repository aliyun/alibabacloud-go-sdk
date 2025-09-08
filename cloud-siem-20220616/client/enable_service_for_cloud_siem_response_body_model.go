// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableServiceForCloudSiemResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v bool) *EnableServiceForCloudSiemResponseBody
  GetData() *bool 
  SetRequestId(v string) *EnableServiceForCloudSiemResponseBody
  GetRequestId() *string 
}

type EnableServiceForCloudSiemResponseBody struct {
  // Indicates whether the threat analysis feature is authorized to access the resource directory. Valid values:
  // 
  // 	- true
  // 
  // 	- false
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

func (s EnableServiceForCloudSiemResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableServiceForCloudSiemResponseBody) GoString() string {
  return s.String()
}

func (s *EnableServiceForCloudSiemResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnableServiceForCloudSiemResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableServiceForCloudSiemResponseBody) SetData(v bool) *EnableServiceForCloudSiemResponseBody {
  s.Data = &v
  return s
}

func (s *EnableServiceForCloudSiemResponseBody) SetRequestId(v string) *EnableServiceForCloudSiemResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableServiceForCloudSiemResponseBody) Validate() error {
  return dara.Validate(s)
}

