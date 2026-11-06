// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBClusterDocumentDBResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableDBClusterDocumentDBResponseBody
  GetRequestId() *string 
}

type EnableDBClusterDocumentDBResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 2921D843-433A-5FB3-A03B-4EC093B219F8
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableDBClusterDocumentDBResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableDBClusterDocumentDBResponseBody) GoString() string {
  return s.String()
}

func (s *EnableDBClusterDocumentDBResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableDBClusterDocumentDBResponseBody) SetRequestId(v string) *EnableDBClusterDocumentDBResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableDBClusterDocumentDBResponseBody) Validate() error {
  return dara.Validate(s)
}

