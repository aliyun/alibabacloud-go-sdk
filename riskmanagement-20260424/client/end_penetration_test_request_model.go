// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndPenetrationTestRequest interface {
  dara.Model
  String() string
  GoString() string
  SetId(v int64) *EndPenetrationTestRequest
  GetId() *int64 
}

type EndPenetrationTestRequest struct {
  // The unique identity sequence of the application record.
  // 
  // > The field type is Long. Precision loss may occur during the procedure to serialize or deserialize. Make sure that the value does not exceed 9007199254740991.
  // 
  // example:
  // 
  // 2398
  Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s EndPenetrationTestRequest) String() string {
  return dara.Prettify(s)
}

func (s EndPenetrationTestRequest) GoString() string {
  return s.String()
}

func (s *EndPenetrationTestRequest) GetId() *int64  {
  return s.Id
}

func (s *EndPenetrationTestRequest) SetId(v int64) *EndPenetrationTestRequest {
  s.Id = &v
  return s
}

func (s *EndPenetrationTestRequest) Validate() error {
  return dara.Validate(s)
}

