// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBClusterDynamoDBResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableDBClusterDynamoDBResponseBody
  GetRequestId() *string 
}

type EnableDBClusterDynamoDBResponseBody struct {
  // Id of the request
  // 
  // example:
  // 
  // CD3FA5F3-FAF3-44CA-AFFF-BAF869******
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableDBClusterDynamoDBResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableDBClusterDynamoDBResponseBody) GoString() string {
  return s.String()
}

func (s *EnableDBClusterDynamoDBResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableDBClusterDynamoDBResponseBody) SetRequestId(v string) *EnableDBClusterDynamoDBResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableDBClusterDynamoDBResponseBody) Validate() error {
  return dara.Validate(s)
}

