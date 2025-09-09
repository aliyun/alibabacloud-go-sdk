// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSqlFlashbackMatchSwitchResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableSqlFlashbackMatchSwitchResponseBody
  GetRequestId() *string 
  SetResult(v bool) *EnableSqlFlashbackMatchSwitchResponseBody
  GetResult() *bool 
  SetSuccess(v bool) *EnableSqlFlashbackMatchSwitchResponseBody
  GetSuccess() *bool 
}

type EnableSqlFlashbackMatchSwitchResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 463A5F0F-12AD-4544-A902-B2B983******
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether SqlFlashbackMatchSwitch is enabled or not.
  // 
  // example:
  // 
  // true
  Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
  // Indicates whether the request was sent successfully or not.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableSqlFlashbackMatchSwitchResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSqlFlashbackMatchSwitchResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSqlFlashbackMatchSwitchResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSqlFlashbackMatchSwitchResponseBody) GetResult() *bool  {
  return s.Result
}

func (s *EnableSqlFlashbackMatchSwitchResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableSqlFlashbackMatchSwitchResponseBody) SetRequestId(v string) *EnableSqlFlashbackMatchSwitchResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSqlFlashbackMatchSwitchResponseBody) SetResult(v bool) *EnableSqlFlashbackMatchSwitchResponseBody {
  s.Result = &v
  return s
}

func (s *EnableSqlFlashbackMatchSwitchResponseBody) SetSuccess(v bool) *EnableSqlFlashbackMatchSwitchResponseBody {
  s.Success = &v
  return s
}

func (s *EnableSqlFlashbackMatchSwitchResponseBody) Validate() error {
  return dara.Validate(s)
}

