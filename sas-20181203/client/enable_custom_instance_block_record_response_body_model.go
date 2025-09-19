// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCustomInstanceBlockRecordResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableCustomInstanceBlockRecordResponseBody
  GetRequestId() *string 
}

type EnableCustomInstanceBlockRecordResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 571B2642-BF51-5BDD-906B-D2340DB9****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCustomInstanceBlockRecordResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableCustomInstanceBlockRecordResponseBody) GoString() string {
  return s.String()
}

func (s *EnableCustomInstanceBlockRecordResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableCustomInstanceBlockRecordResponseBody) SetRequestId(v string) *EnableCustomInstanceBlockRecordResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableCustomInstanceBlockRecordResponseBody) Validate() error {
  return dara.Validate(s)
}

