// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndToEndRealTimeDialogResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EndToEndRealTimeDialogResponseBody
  GetRequestId() *string 
}

type EndToEndRealTimeDialogResponseBody struct {
  // example:
  // 
  // 1C98B466-D6E0-5252-A60B-F345CBB33DDB
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s EndToEndRealTimeDialogResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EndToEndRealTimeDialogResponseBody) GoString() string {
  return s.String()
}

func (s *EndToEndRealTimeDialogResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EndToEndRealTimeDialogResponseBody) SetRequestId(v string) *EndToEndRealTimeDialogResponseBody {
  s.RequestId = &v
  return s
}

func (s *EndToEndRealTimeDialogResponseBody) Validate() error {
  return dara.Validate(s)
}

