// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCustomBlockRecordResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableCustomBlockRecordResponseBody
  GetRequestId() *string 
}

type EnableCustomBlockRecordResponseBody struct {
  // The ID of the request, which is used to locate and troubleshoot issues.
  // 
  // example:
  // 
  // F02D5F26-70B9-53BD-9CDF-A316FD11****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCustomBlockRecordResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableCustomBlockRecordResponseBody) GoString() string {
  return s.String()
}

func (s *EnableCustomBlockRecordResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableCustomBlockRecordResponseBody) SetRequestId(v string) *EnableCustomBlockRecordResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableCustomBlockRecordResponseBody) Validate() error {
  return dara.Validate(s)
}

