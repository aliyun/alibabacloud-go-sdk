// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditProhibitedDevicesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v bool) *EditProhibitedDevicesResponseBody
  GetData() *bool 
  SetRequestId(v string) *EditProhibitedDevicesResponseBody
  GetRequestId() *string 
}

type EditProhibitedDevicesResponseBody struct {
  // example:
  // 
  // true
  Data *bool `json:"data,omitempty" xml:"data,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 9bc20a5a-b26b-4c28-922a-7cd10b61f96f
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s EditProhibitedDevicesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditProhibitedDevicesResponseBody) GoString() string {
  return s.String()
}

func (s *EditProhibitedDevicesResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EditProhibitedDevicesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditProhibitedDevicesResponseBody) SetData(v bool) *EditProhibitedDevicesResponseBody {
  s.Data = &v
  return s
}

func (s *EditProhibitedDevicesResponseBody) SetRequestId(v string) *EditProhibitedDevicesResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditProhibitedDevicesResponseBody) Validate() error {
  return dara.Validate(s)
}

