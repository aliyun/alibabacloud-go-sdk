// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotUpdateJobParams interface {
	dara.Model
	String() string
	GoString() string
	SetRescaleJobParam(v *RescaleJobParam) *HotUpdateJobParams
	GetRescaleJobParam() *RescaleJobParam
	SetUpdateJobConfigParam(v *UpdateJobConfigParam) *HotUpdateJobParams
	GetUpdateJobConfigParam() *UpdateJobConfigParam
}

type HotUpdateJobParams struct {
	// 资源动态更新。
	RescaleJobParam *RescaleJobParam `json:"rescaleJobParam,omitempty" xml:"rescaleJobParam,omitempty"`
	// 运行参数动态更新。
	UpdateJobConfigParam *UpdateJobConfigParam `json:"updateJobConfigParam,omitempty" xml:"updateJobConfigParam,omitempty"`
}

func (s HotUpdateJobParams) String() string {
	return dara.Prettify(s)
}

func (s HotUpdateJobParams) GoString() string {
	return s.String()
}

func (s *HotUpdateJobParams) GetRescaleJobParam() *RescaleJobParam {
	return s.RescaleJobParam
}

func (s *HotUpdateJobParams) GetUpdateJobConfigParam() *UpdateJobConfigParam {
	return s.UpdateJobConfigParam
}

func (s *HotUpdateJobParams) SetRescaleJobParam(v *RescaleJobParam) *HotUpdateJobParams {
	s.RescaleJobParam = v
	return s
}

func (s *HotUpdateJobParams) SetUpdateJobConfigParam(v *UpdateJobConfigParam) *HotUpdateJobParams {
	s.UpdateJobConfigParam = v
	return s
}

func (s *HotUpdateJobParams) Validate() error {
	if s.RescaleJobParam != nil {
		if err := s.RescaleJobParam.Validate(); err != nil {
			return err
		}
	}
	if s.UpdateJobConfigParam != nil {
		if err := s.UpdateJobConfigParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}
