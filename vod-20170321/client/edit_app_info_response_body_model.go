// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditAppInfoResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *EditAppInfoResponseBodyData) *EditAppInfoResponseBody
  GetData() *EditAppInfoResponseBodyData 
  SetRequestId(v string) *EditAppInfoResponseBody
  GetRequestId() *string 
}

type EditAppInfoResponseBody struct {
  Data *EditAppInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditAppInfoResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditAppInfoResponseBody) GoString() string {
  return s.String()
}

func (s *EditAppInfoResponseBody) GetData() *EditAppInfoResponseBodyData  {
  return s.Data
}

func (s *EditAppInfoResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditAppInfoResponseBody) SetData(v *EditAppInfoResponseBodyData) *EditAppInfoResponseBody {
  s.Data = v
  return s
}

func (s *EditAppInfoResponseBody) SetRequestId(v string) *EditAppInfoResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditAppInfoResponseBody) Validate() error {
  return dara.Validate(s)
}

type EditAppInfoResponseBodyData struct {
  ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
}

func (s EditAppInfoResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EditAppInfoResponseBodyData) GoString() string {
  return s.String()
}

func (s *EditAppInfoResponseBodyData) GetItemId() *string  {
  return s.ItemId
}

func (s *EditAppInfoResponseBodyData) SetItemId(v string) *EditAppInfoResponseBodyData {
  s.ItemId = &v
  return s
}

func (s *EditAppInfoResponseBodyData) Validate() error {
  return dara.Validate(s)
}

