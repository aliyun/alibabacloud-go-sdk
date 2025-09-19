// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBruteForceRecordResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableBruteForceRecordResponseBody
  GetRequestId() *string 
}

type EnableBruteForceRecordResponseBody struct {
  // The ID of the request, which is used to locate and troubleshoot issues.
  // 
  // example:
  // 
  // DFAECA37-4660-5EB6-9A18-8FDF56B3****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableBruteForceRecordResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableBruteForceRecordResponseBody) GoString() string {
  return s.String()
}

func (s *EnableBruteForceRecordResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableBruteForceRecordResponseBody) SetRequestId(v string) *EnableBruteForceRecordResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableBruteForceRecordResponseBody) Validate() error {
  return dara.Validate(s)
}

