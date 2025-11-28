// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBizChainDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizChainId(v string) *ListBizChainDataRequest
	GetBizChainId() *string
	SetEndTime(v int64) *ListBizChainDataRequest
	GetEndTime() *int64
	SetIoTDataDID(v string) *ListBizChainDataRequest
	GetIoTDataDID() *string
	SetMemberId(v string) *ListBizChainDataRequest
	GetMemberId() *string
	SetNum(v int32) *ListBizChainDataRequest
	GetNum() *int32
	SetRegionId(v string) *ListBizChainDataRequest
	GetRegionId() *string
	SetSize(v int32) *ListBizChainDataRequest
	GetSize() *int32
	SetStartTime(v int64) *ListBizChainDataRequest
	GetStartTime() *int64
}

type ListBizChainDataRequest struct {
	// This parameter is required.
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	// This parameter is required.
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IoTDataDID *string `json:"IoTDataDID,omitempty" xml:"IoTDataDID,omitempty"`
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter is required.
	Num      *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListBizChainDataRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBizChainDataRequest) GoString() string {
	return s.String()
}

func (s *ListBizChainDataRequest) GetBizChainId() *string {
	return s.BizChainId
}

func (s *ListBizChainDataRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListBizChainDataRequest) GetIoTDataDID() *string {
	return s.IoTDataDID
}

func (s *ListBizChainDataRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *ListBizChainDataRequest) GetNum() *int32 {
	return s.Num
}

func (s *ListBizChainDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBizChainDataRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListBizChainDataRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListBizChainDataRequest) SetBizChainId(v string) *ListBizChainDataRequest {
	s.BizChainId = &v
	return s
}

func (s *ListBizChainDataRequest) SetEndTime(v int64) *ListBizChainDataRequest {
	s.EndTime = &v
	return s
}

func (s *ListBizChainDataRequest) SetIoTDataDID(v string) *ListBizChainDataRequest {
	s.IoTDataDID = &v
	return s
}

func (s *ListBizChainDataRequest) SetMemberId(v string) *ListBizChainDataRequest {
	s.MemberId = &v
	return s
}

func (s *ListBizChainDataRequest) SetNum(v int32) *ListBizChainDataRequest {
	s.Num = &v
	return s
}

func (s *ListBizChainDataRequest) SetRegionId(v string) *ListBizChainDataRequest {
	s.RegionId = &v
	return s
}

func (s *ListBizChainDataRequest) SetSize(v int32) *ListBizChainDataRequest {
	s.Size = &v
	return s
}

func (s *ListBizChainDataRequest) SetStartTime(v int64) *ListBizChainDataRequest {
	s.StartTime = &v
	return s
}

func (s *ListBizChainDataRequest) Validate() error {
	return dara.Validate(s)
}
