// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstimatedRestartTimeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EstimatedRestartTimeResponseBody
  GetRequestId() *string 
  SetResult(v *EstimatedRestartTimeResponseBodyResult) *EstimatedRestartTimeResponseBody
  GetResult() *EstimatedRestartTimeResponseBodyResult 
}

type EstimatedRestartTimeResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The returned result.
  Result *EstimatedRestartTimeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EstimatedRestartTimeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EstimatedRestartTimeResponseBody) GoString() string {
  return s.String()
}

func (s *EstimatedRestartTimeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EstimatedRestartTimeResponseBody) GetResult() *EstimatedRestartTimeResponseBodyResult  {
  return s.Result
}

func (s *EstimatedRestartTimeResponseBody) SetRequestId(v string) *EstimatedRestartTimeResponseBody {
  s.RequestId = &v
  return s
}

func (s *EstimatedRestartTimeResponseBody) SetResult(v *EstimatedRestartTimeResponseBodyResult) *EstimatedRestartTimeResponseBody {
  s.Result = v
  return s
}

func (s *EstimatedRestartTimeResponseBody) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EstimatedRestartTimeResponseBodyResult struct {
  // The unit.
  // 
  // example:
  // 
  // second
  Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
  // The estimated restart time.
  // 
  // example:
  // 
  // 50
  Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s EstimatedRestartTimeResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s EstimatedRestartTimeResponseBodyResult) GoString() string {
  return s.String()
}

func (s *EstimatedRestartTimeResponseBodyResult) GetUnit() *string  {
  return s.Unit
}

func (s *EstimatedRestartTimeResponseBodyResult) GetValue() *int64  {
  return s.Value
}

func (s *EstimatedRestartTimeResponseBodyResult) SetUnit(v string) *EstimatedRestartTimeResponseBodyResult {
  s.Unit = &v
  return s
}

func (s *EstimatedRestartTimeResponseBodyResult) SetValue(v int64) *EstimatedRestartTimeResponseBodyResult {
  s.Value = &v
  return s
}

func (s *EstimatedRestartTimeResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

