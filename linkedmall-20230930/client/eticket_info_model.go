// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEticketInfo interface {
  dara.Model
  String() string
  GoString() string
  SetAvailableNum(v int64) *EticketInfo
  GetAvailableNum() *int64 
  SetCode(v string) *EticketInfo
  GetCode() *string 
  SetCodeStatus(v int64) *EticketInfo
  GetCodeStatus() *int64 
  SetEndTime(v string) *EticketInfo
  GetEndTime() *string 
  SetLockNum(v int64) *EticketInfo
  GetLockNum() *int64 
  SetQrcodeUrl(v string) *EticketInfo
  GetQrcodeUrl() *string 
  SetStartTime(v string) *EticketInfo
  GetStartTime() *string 
  SetUseTime(v string) *EticketInfo
  GetUseTime() *string 
  SetUsedNum(v int64) *EticketInfo
  GetUsedNum() *int64 
}

type EticketInfo struct {
  // example:
  // 
  // 0
  AvailableNum *int64 `json:"availableNum,omitempty" xml:"availableNum,omitempty"`
  // example:
  // 
  // taobao******tpg
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  // example:
  // 
  // -1
  CodeStatus *int64 `json:"codeStatus,omitempty" xml:"codeStatus,omitempty"`
  // example:
  // 
  // 2026-08-02T23:59:59.000+08:00
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // example:
  // 
  // 0
  LockNum *int64 `json:"lockNum,omitempty" xml:"lockNum,omitempty"`
  // example:
  // 
  // http://qrcode.alicdn.com/img.jpg
  QrcodeUrl *string `json:"qrcodeUrl,omitempty" xml:"qrcodeUrl,omitempty"`
  // example:
  // 
  // 2026-02-04T00:00:00.000+08:00
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // example:
  // 
  // 2026-02-04T15:07:59.000+08:00
  UseTime *string `json:"useTime,omitempty" xml:"useTime,omitempty"`
  // example:
  // 
  // 1
  UsedNum *int64 `json:"usedNum,omitempty" xml:"usedNum,omitempty"`
}

func (s EticketInfo) String() string {
  return dara.Prettify(s)
}

func (s EticketInfo) GoString() string {
  return s.String()
}

func (s *EticketInfo) GetAvailableNum() *int64  {
  return s.AvailableNum
}

func (s *EticketInfo) GetCode() *string  {
  return s.Code
}

func (s *EticketInfo) GetCodeStatus() *int64  {
  return s.CodeStatus
}

func (s *EticketInfo) GetEndTime() *string  {
  return s.EndTime
}

func (s *EticketInfo) GetLockNum() *int64  {
  return s.LockNum
}

func (s *EticketInfo) GetQrcodeUrl() *string  {
  return s.QrcodeUrl
}

func (s *EticketInfo) GetStartTime() *string  {
  return s.StartTime
}

func (s *EticketInfo) GetUseTime() *string  {
  return s.UseTime
}

func (s *EticketInfo) GetUsedNum() *int64  {
  return s.UsedNum
}

func (s *EticketInfo) SetAvailableNum(v int64) *EticketInfo {
  s.AvailableNum = &v
  return s
}

func (s *EticketInfo) SetCode(v string) *EticketInfo {
  s.Code = &v
  return s
}

func (s *EticketInfo) SetCodeStatus(v int64) *EticketInfo {
  s.CodeStatus = &v
  return s
}

func (s *EticketInfo) SetEndTime(v string) *EticketInfo {
  s.EndTime = &v
  return s
}

func (s *EticketInfo) SetLockNum(v int64) *EticketInfo {
  s.LockNum = &v
  return s
}

func (s *EticketInfo) SetQrcodeUrl(v string) *EticketInfo {
  s.QrcodeUrl = &v
  return s
}

func (s *EticketInfo) SetStartTime(v string) *EticketInfo {
  s.StartTime = &v
  return s
}

func (s *EticketInfo) SetUseTime(v string) *EticketInfo {
  s.UseTime = &v
  return s
}

func (s *EticketInfo) SetUsedNum(v int64) *EticketInfo {
  s.UsedNum = &v
  return s
}

func (s *EticketInfo) Validate() error {
  return dara.Validate(s)
}

