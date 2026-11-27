// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaultDiagnosisOverallDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeFaultDiagnosisOverallDataRequest
	GetAppId() *string
	SetEndTs(v int64) *DescribeFaultDiagnosisOverallDataRequest
	GetEndTs() *int64
	SetStartTs(v int64) *DescribeFaultDiagnosisOverallDataRequest
	GetStartTs() *int64
	SetStatDim(v string) *DescribeFaultDiagnosisOverallDataRequest
	GetStatDim() *string
}

type DescribeFaultDiagnosisOverallDataRequest struct {
	// The app ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4eah****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time of the query. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1615910399
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// The statistical dimension. Valid values:
	//
	// - **JOIN_SLOW_USER**: the number of users who joined the channel slowly
	//
	// - **AUDIO_STUCK_USER**: the number of users who experienced audio stuttering
	//
	// - **VIDEO_STUCK_USER**: the number of users who experienced video stuttering
	//
	// - **VIDEO_VAGUE_USER**: the number of users who experienced blurry video
	//
	// - **HIGH_DELAY_USER**: the number of users who experienced high call latency
	//
	// - **FIRST_FRAME_SLOW_USER**: the number of users who experienced slow first screen rendering
	//
	// This parameter is required.
	//
	// example:
	//
	// JOIN_SLOW_USER
	StatDim *string `json:"StatDim,omitempty" xml:"StatDim,omitempty"`
}

func (s DescribeFaultDiagnosisOverallDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeFaultDiagnosisOverallDataRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribeFaultDiagnosisOverallDataRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeFaultDiagnosisOverallDataRequest) GetStatDim() *string {
	return s.StatDim
}

func (s *DescribeFaultDiagnosisOverallDataRequest) SetAppId(v string) *DescribeFaultDiagnosisOverallDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataRequest) SetEndTs(v int64) *DescribeFaultDiagnosisOverallDataRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataRequest) SetStartTs(v int64) *DescribeFaultDiagnosisOverallDataRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataRequest) SetStatDim(v string) *DescribeFaultDiagnosisOverallDataRequest {
	s.StatDim = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataRequest) Validate() error {
	return dara.Validate(s)
}
