// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotUpdateJobResult interface {
	dara.Model
	String() string
	GoString() string
	SetHotUpdateParams(v *HotUpdateJobParams) *HotUpdateJobResult
	GetHotUpdateParams() *HotUpdateJobParams
	SetJobHotUpdateId(v string) *HotUpdateJobResult
	GetJobHotUpdateId() *string
	SetJobId(v string) *HotUpdateJobResult
	GetJobId() *string
	SetStatus(v *HotUpdateJobStatus) *HotUpdateJobResult
	GetStatus() *HotUpdateJobStatus
	SetTargetResourceSetting(v *BriefResourceSetting) *HotUpdateJobResult
	GetTargetResourceSetting() *BriefResourceSetting
}

type HotUpdateJobResult struct {
	HotUpdateParams       *HotUpdateJobParams   `json:"hotUpdateParams,omitempty" xml:"hotUpdateParams,omitempty"`
	JobHotUpdateId        *string               `json:"jobHotUpdateId,omitempty" xml:"jobHotUpdateId,omitempty"`
	JobId                 *string               `json:"jobId,omitempty" xml:"jobId,omitempty"`
	Status                *HotUpdateJobStatus   `json:"status,omitempty" xml:"status,omitempty"`
	TargetResourceSetting *BriefResourceSetting `json:"targetResourceSetting,omitempty" xml:"targetResourceSetting,omitempty"`
}

func (s HotUpdateJobResult) String() string {
	return dara.Prettify(s)
}

func (s HotUpdateJobResult) GoString() string {
	return s.String()
}

func (s *HotUpdateJobResult) GetHotUpdateParams() *HotUpdateJobParams {
	return s.HotUpdateParams
}

func (s *HotUpdateJobResult) GetJobHotUpdateId() *string {
	return s.JobHotUpdateId
}

func (s *HotUpdateJobResult) GetJobId() *string {
	return s.JobId
}

func (s *HotUpdateJobResult) GetStatus() *HotUpdateJobStatus {
	return s.Status
}

func (s *HotUpdateJobResult) GetTargetResourceSetting() *BriefResourceSetting {
	return s.TargetResourceSetting
}

func (s *HotUpdateJobResult) SetHotUpdateParams(v *HotUpdateJobParams) *HotUpdateJobResult {
	s.HotUpdateParams = v
	return s
}

func (s *HotUpdateJobResult) SetJobHotUpdateId(v string) *HotUpdateJobResult {
	s.JobHotUpdateId = &v
	return s
}

func (s *HotUpdateJobResult) SetJobId(v string) *HotUpdateJobResult {
	s.JobId = &v
	return s
}

func (s *HotUpdateJobResult) SetStatus(v *HotUpdateJobStatus) *HotUpdateJobResult {
	s.Status = v
	return s
}

func (s *HotUpdateJobResult) SetTargetResourceSetting(v *BriefResourceSetting) *HotUpdateJobResult {
	s.TargetResourceSetting = v
	return s
}

func (s *HotUpdateJobResult) Validate() error {
	if s.HotUpdateParams != nil {
		if err := s.HotUpdateParams.Validate(); err != nil {
			return err
		}
	}
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	if s.TargetResourceSetting != nil {
		if err := s.TargetResourceSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}
