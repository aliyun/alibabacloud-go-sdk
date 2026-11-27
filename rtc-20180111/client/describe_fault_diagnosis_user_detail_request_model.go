// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaultDiagnosisUserDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeFaultDiagnosisUserDetailRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeFaultDiagnosisUserDetailRequest
	GetChannelId() *string
	SetCreatedTs(v int64) *DescribeFaultDiagnosisUserDetailRequest
	GetCreatedTs() *int64
	SetFaultType(v string) *DescribeFaultDiagnosisUserDetailRequest
	GetFaultType() *string
	SetQueryCallUserInfo(v bool) *DescribeFaultDiagnosisUserDetailRequest
	GetQueryCallUserInfo() *bool
	SetUserId(v string) *DescribeFaultDiagnosisUserDetailRequest
	GetUserId() *string
}

type DescribeFaultDiagnosisUserDetailRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 频道ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 311
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// 频道创建时间，使用UNIX时间戳表示，单位：秒。
	//
	// This parameter is required.
	//
	// example:
	//
	// 1615892596
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// 异常类型。取值：
	//
	// - **JOIN_SLOW**：进频道慢 。
	//
	// - **AUDIO_STUCK**：音频卡顿。
	//
	// - **VIDEO_STUCK**：视频卡顿。
	//
	// - **VIDEO_VAGUE**：视频模糊。
	//
	// - **HIGH_DELAY**：通话延迟高 。
	//
	// - **FIRST_FRAME_SLOW**：接收首屏慢。
	//
	// This parameter is required.
	//
	// example:
	//
	// JOIN_SLOW
	FaultType *string `json:"FaultType,omitempty" xml:"FaultType,omitempty"`
	// 是否查询通话用户信息，为空默认是false。
	//
	// example:
	//
	// true
	QueryCallUserInfo *bool `json:"QueryCallUserInfo,omitempty" xml:"QueryCallUserInfo,omitempty"`
	// 用户ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c906531af5f9****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeFaultDiagnosisUserDetailRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeFaultDiagnosisUserDetailRequest) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeFaultDiagnosisUserDetailRequest) GetFaultType() *string {
	return s.FaultType
}

func (s *DescribeFaultDiagnosisUserDetailRequest) GetQueryCallUserInfo() *bool {
	return s.QueryCallUserInfo
}

func (s *DescribeFaultDiagnosisUserDetailRequest) GetUserId() *string {
	return s.UserId
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetAppId(v string) *DescribeFaultDiagnosisUserDetailRequest {
	s.AppId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetChannelId(v string) *DescribeFaultDiagnosisUserDetailRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetCreatedTs(v int64) *DescribeFaultDiagnosisUserDetailRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetFaultType(v string) *DescribeFaultDiagnosisUserDetailRequest {
	s.FaultType = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetQueryCallUserInfo(v bool) *DescribeFaultDiagnosisUserDetailRequest {
	s.QueryCallUserInfo = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetUserId(v string) *DescribeFaultDiagnosisUserDetailRequest {
	s.UserId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailRequest) Validate() error {
	return dara.Validate(s)
}
