// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditUnfavorableAreaDevicesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v bool) *EditUnfavorableAreaDevicesResponseBody
  GetData() *bool 
  SetRequestId(v string) *EditUnfavorableAreaDevicesResponseBody
  GetRequestId() *string 
}

type EditUnfavorableAreaDevicesResponseBody struct {
  // example:
  // 
  // true
  Data *bool `json:"data,omitempty" xml:"data,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 83A5A7DD-8974-5769-952E-590A97BEA34E
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s EditUnfavorableAreaDevicesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditUnfavorableAreaDevicesResponseBody) GoString() string {
  return s.String()
}

func (s *EditUnfavorableAreaDevicesResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EditUnfavorableAreaDevicesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditUnfavorableAreaDevicesResponseBody) SetData(v bool) *EditUnfavorableAreaDevicesResponseBody {
  s.Data = &v
  return s
}

func (s *EditUnfavorableAreaDevicesResponseBody) SetRequestId(v string) *EditUnfavorableAreaDevicesResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditUnfavorableAreaDevicesResponseBody) Validate() error {
  return dara.Validate(s)
}

