// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPoolInfo(v *GetPoolResponseBodyPoolInfo) *GetPoolResponseBody
	GetPoolInfo() *GetPoolResponseBodyPoolInfo
	SetRequestId(v string) *GetPoolResponseBody
	GetRequestId() *string
}

type GetPoolResponseBody struct {
	PoolInfo *GetPoolResponseBodyPoolInfo `json:"PoolInfo,omitempty" xml:"PoolInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPoolResponseBody) GoString() string {
	return s.String()
}

func (s *GetPoolResponseBody) GetPoolInfo() *GetPoolResponseBodyPoolInfo {
	return s.PoolInfo
}

func (s *GetPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPoolResponseBody) SetPoolInfo(v *GetPoolResponseBodyPoolInfo) *GetPoolResponseBody {
	s.PoolInfo = v
	return s
}

func (s *GetPoolResponseBody) SetRequestId(v string) *GetPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPoolResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPoolResponseBodyPoolInfo struct {
	// example:
	//
	// 2024-12-01 20:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1
	ExectorUsage *int32 `json:"ExectorUsage,omitempty" xml:"ExectorUsage,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// 2000
	MaxExectorNum *int32 `json:"MaxExectorNum,omitempty" xml:"MaxExectorNum,omitempty"`
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Fails to **	- pool: ***.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// Working
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-12-01 20:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetPoolResponseBodyPoolInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPoolResponseBodyPoolInfo) GoString() string {
	return s.String()
}

func (s *GetPoolResponseBodyPoolInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPoolResponseBodyPoolInfo) GetExectorUsage() *int32 {
	return s.ExectorUsage
}

func (s *GetPoolResponseBodyPoolInfo) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *GetPoolResponseBodyPoolInfo) GetMaxExectorNum() *int32 {
	return s.MaxExectorNum
}

func (s *GetPoolResponseBodyPoolInfo) GetPoolName() *string {
	return s.PoolName
}

func (s *GetPoolResponseBodyPoolInfo) GetPriority() *int32 {
	return s.Priority
}

func (s *GetPoolResponseBodyPoolInfo) GetReason() *string {
	return s.Reason
}

func (s *GetPoolResponseBodyPoolInfo) GetStatus() *string {
	return s.Status
}

func (s *GetPoolResponseBodyPoolInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetPoolResponseBodyPoolInfo) SetCreateTime(v string) *GetPoolResponseBodyPoolInfo {
	s.CreateTime = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetExectorUsage(v int32) *GetPoolResponseBodyPoolInfo {
	s.ExectorUsage = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetIsDefault(v bool) *GetPoolResponseBodyPoolInfo {
	s.IsDefault = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetMaxExectorNum(v int32) *GetPoolResponseBodyPoolInfo {
	s.MaxExectorNum = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetPoolName(v string) *GetPoolResponseBodyPoolInfo {
	s.PoolName = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetPriority(v int32) *GetPoolResponseBodyPoolInfo {
	s.Priority = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetReason(v string) *GetPoolResponseBodyPoolInfo {
	s.Reason = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetStatus(v string) *GetPoolResponseBodyPoolInfo {
	s.Status = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetUpdateTime(v string) *GetPoolResponseBodyPoolInfo {
	s.UpdateTime = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) Validate() error {
	return dara.Validate(s)
}
