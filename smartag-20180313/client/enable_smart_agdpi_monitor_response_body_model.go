// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSmartAGDpiMonitorResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableSmartAGDpiMonitorResponseBody
  GetRequestId() *string 
}

type EnableSmartAGDpiMonitorResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 64966488-B3E3-41E2-9570-4596117EC12E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableSmartAGDpiMonitorResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSmartAGDpiMonitorResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSmartAGDpiMonitorResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSmartAGDpiMonitorResponseBody) SetRequestId(v string) *EnableSmartAGDpiMonitorResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSmartAGDpiMonitorResponseBody) Validate() error {
  return dara.Validate(s)
}

