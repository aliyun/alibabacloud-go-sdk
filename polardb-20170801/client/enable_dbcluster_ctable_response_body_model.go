// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBClusterCTableResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableDBClusterCTableResponseBody
  GetRequestId() *string 
}

type EnableDBClusterCTableResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 2921D843-433A-5FB3-A03B-4EC093B219F8
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableDBClusterCTableResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableDBClusterCTableResponseBody) GoString() string {
  return s.String()
}

func (s *EnableDBClusterCTableResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableDBClusterCTableResponseBody) SetRequestId(v string) *EnableDBClusterCTableResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableDBClusterCTableResponseBody) Validate() error {
  return dara.Validate(s)
}

