// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEmptySlsLogstoreResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EmptySlsLogstoreResponseBody
  GetRequestId() *string 
}

type EmptySlsLogstoreResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // CF33B4C3-196E-4015-AADD-5CAD00057B80
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EmptySlsLogstoreResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EmptySlsLogstoreResponseBody) GoString() string {
  return s.String()
}

func (s *EmptySlsLogstoreResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EmptySlsLogstoreResponseBody) SetRequestId(v string) *EmptySlsLogstoreResponseBody {
  s.RequestId = &v
  return s
}

func (s *EmptySlsLogstoreResponseBody) Validate() error {
  return dara.Validate(s)
}

