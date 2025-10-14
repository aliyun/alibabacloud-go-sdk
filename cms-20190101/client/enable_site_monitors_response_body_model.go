// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSiteMonitorsResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableSiteMonitorsResponseBody
  GetCode() *string 
  SetData(v *EnableSiteMonitorsResponseBodyData) *EnableSiteMonitorsResponseBody
  GetData() *EnableSiteMonitorsResponseBodyData 
  SetMessage(v string) *EnableSiteMonitorsResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableSiteMonitorsResponseBody
  GetRequestId() *string 
  SetSuccess(v string) *EnableSiteMonitorsResponseBody
  GetSuccess() *string 
}

type EnableSiteMonitorsResponseBody struct {
  // The responses code.
  // 
  // >  The status code 200 indicates that the request was successful.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The number of detection points that are affected by the site monitoring tasks.
  Data *EnableSiteMonitorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // The returned message.
  // 
  // example:
  // 
  // successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 3fcd12e7-d387-42ee-b77e-661c775bb17f
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- true
  // 
  // 	- false
  // 
  // example:
  // 
  // true
  Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableSiteMonitorsResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSiteMonitorsResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSiteMonitorsResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableSiteMonitorsResponseBody) GetData() *EnableSiteMonitorsResponseBodyData  {
  return s.Data
}

func (s *EnableSiteMonitorsResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableSiteMonitorsResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSiteMonitorsResponseBody) GetSuccess() *string  {
  return s.Success
}

func (s *EnableSiteMonitorsResponseBody) SetCode(v string) *EnableSiteMonitorsResponseBody {
  s.Code = &v
  return s
}

func (s *EnableSiteMonitorsResponseBody) SetData(v *EnableSiteMonitorsResponseBodyData) *EnableSiteMonitorsResponseBody {
  s.Data = v
  return s
}

func (s *EnableSiteMonitorsResponseBody) SetMessage(v string) *EnableSiteMonitorsResponseBody {
  s.Message = &v
  return s
}

func (s *EnableSiteMonitorsResponseBody) SetRequestId(v string) *EnableSiteMonitorsResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSiteMonitorsResponseBody) SetSuccess(v string) *EnableSiteMonitorsResponseBody {
  s.Success = &v
  return s
}

func (s *EnableSiteMonitorsResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnableSiteMonitorsResponseBodyData struct {
  // The number of detection points.
  // 
  // example:
  // 
  // 0
  Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
}

func (s EnableSiteMonitorsResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnableSiteMonitorsResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnableSiteMonitorsResponseBodyData) GetCount() *int32  {
  return s.Count
}

func (s *EnableSiteMonitorsResponseBodyData) SetCount(v int32) *EnableSiteMonitorsResponseBodyData {
  s.Count = &v
  return s
}

func (s *EnableSiteMonitorsResponseBodyData) Validate() error {
  return dara.Validate(s)
}

