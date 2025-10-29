// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstimatedLogstashRestartTimeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EstimatedLogstashRestartTimeResponseBody
  GetRequestId() *string 
  SetResult(v *EstimatedLogstashRestartTimeResponseBodyResult) *EstimatedLogstashRestartTimeResponseBody
  GetResult() *EstimatedLogstashRestartTimeResponseBodyResult 
}

type EstimatedLogstashRestartTimeResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The returned result.
  Result *EstimatedLogstashRestartTimeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EstimatedLogstashRestartTimeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EstimatedLogstashRestartTimeResponseBody) GoString() string {
  return s.String()
}

func (s *EstimatedLogstashRestartTimeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EstimatedLogstashRestartTimeResponseBody) GetResult() *EstimatedLogstashRestartTimeResponseBodyResult  {
  return s.Result
}

func (s *EstimatedLogstashRestartTimeResponseBody) SetRequestId(v string) *EstimatedLogstashRestartTimeResponseBody {
  s.RequestId = &v
  return s
}

func (s *EstimatedLogstashRestartTimeResponseBody) SetResult(v *EstimatedLogstashRestartTimeResponseBodyResult) *EstimatedLogstashRestartTimeResponseBody {
  s.Result = v
  return s
}

func (s *EstimatedLogstashRestartTimeResponseBody) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EstimatedLogstashRestartTimeResponseBodyResult struct {
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
  // 600
  Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s EstimatedLogstashRestartTimeResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s EstimatedLogstashRestartTimeResponseBodyResult) GoString() string {
  return s.String()
}

func (s *EstimatedLogstashRestartTimeResponseBodyResult) GetUnit() *string  {
  return s.Unit
}

func (s *EstimatedLogstashRestartTimeResponseBodyResult) GetValue() *int64  {
  return s.Value
}

func (s *EstimatedLogstashRestartTimeResponseBodyResult) SetUnit(v string) *EstimatedLogstashRestartTimeResponseBodyResult {
  s.Unit = &v
  return s
}

func (s *EstimatedLogstashRestartTimeResponseBodyResult) SetValue(v int64) *EstimatedLogstashRestartTimeResponseBodyResult {
  s.Value = &v
  return s
}

func (s *EstimatedLogstashRestartTimeResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

