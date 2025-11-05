// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditZeroCreditShutdownRequest interface {
  dara.Model
  String() string
  GoString() string
  SetShutdownPolicy(v string) *EditZeroCreditShutdownRequest
  GetShutdownPolicy() *string 
  SetUid(v int64) *EditZeroCreditShutdownRequest
  GetUid() *int64 
}

type EditZeroCreditShutdownRequest struct {
  // UID
  // 
  // example:
  // 
  // Shutdown Policy</br>
  // 
  // - immediatelyStop, The instances of the specified End User\\"s account will be shutdown immediately once EU triggered the Shutdown Policy.</br>
  // 
  // - delayStop, The instances of the specified End User\\"s account will be shutdown later, even EU have triggered the Shutdown Policy.</br>
  // 
  // - noStop, The instances of the specified End User\\"s account will not be shutdown, after EU have triggered the Shutdown Policy.</br>
  ShutdownPolicy *string `json:"ShutdownPolicy,omitempty" xml:"ShutdownPolicy,omitempty"`
  // No Change History
  // 
  // example:
  // 
  // 1263644979775567
  Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s EditZeroCreditShutdownRequest) String() string {
  return dara.Prettify(s)
}

func (s EditZeroCreditShutdownRequest) GoString() string {
  return s.String()
}

func (s *EditZeroCreditShutdownRequest) GetShutdownPolicy() *string  {
  return s.ShutdownPolicy
}

func (s *EditZeroCreditShutdownRequest) GetUid() *int64  {
  return s.Uid
}

func (s *EditZeroCreditShutdownRequest) SetShutdownPolicy(v string) *EditZeroCreditShutdownRequest {
  s.ShutdownPolicy = &v
  return s
}

func (s *EditZeroCreditShutdownRequest) SetUid(v int64) *EditZeroCreditShutdownRequest {
  s.Uid = &v
  return s
}

func (s *EditZeroCreditShutdownRequest) Validate() error {
  return dara.Validate(s)
}

