// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaultDiagnosisUserListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeFaultDiagnosisUserListRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeFaultDiagnosisUserListRequest
	GetChannelId() *string
	SetEndTs(v int64) *DescribeFaultDiagnosisUserListRequest
	GetEndTs() *int64
	SetFaultTypes(v string) *DescribeFaultDiagnosisUserListRequest
	GetFaultTypes() *string
	SetPageNo(v int32) *DescribeFaultDiagnosisUserListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeFaultDiagnosisUserListRequest
	GetPageSize() *int32
	SetStartTs(v int64) *DescribeFaultDiagnosisUserListRequest
	GetStartTs() *int64
	SetUserId(v string) *DescribeFaultDiagnosisUserListRequest
	GetUserId() *string
}

type DescribeFaultDiagnosisUserListRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 311
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615892596
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// example:
	//
	// JOIN_SLOW,AUDIO_STUCK
	FaultTypes *string `json:"FaultTypes,omitempty" xml:"FaultTypes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615806196
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// example:
	//
	// c906531af5f9****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserListRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserListRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeFaultDiagnosisUserListRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeFaultDiagnosisUserListRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribeFaultDiagnosisUserListRequest) GetFaultTypes() *string {
	return s.FaultTypes
}

func (s *DescribeFaultDiagnosisUserListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeFaultDiagnosisUserListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFaultDiagnosisUserListRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeFaultDiagnosisUserListRequest) GetUserId() *string {
	return s.UserId
}

func (s *DescribeFaultDiagnosisUserListRequest) SetAppId(v string) *DescribeFaultDiagnosisUserListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetChannelId(v string) *DescribeFaultDiagnosisUserListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetEndTs(v int64) *DescribeFaultDiagnosisUserListRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetFaultTypes(v string) *DescribeFaultDiagnosisUserListRequest {
	s.FaultTypes = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetPageNo(v int32) *DescribeFaultDiagnosisUserListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetPageSize(v int32) *DescribeFaultDiagnosisUserListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetStartTs(v int64) *DescribeFaultDiagnosisUserListRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetUserId(v string) *DescribeFaultDiagnosisUserListRequest {
	s.UserId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) Validate() error {
	return dara.Validate(s)
}
