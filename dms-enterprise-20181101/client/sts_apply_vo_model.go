// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStsApplyVO interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunId(v string) *StsApplyVO
	GetAliyunId() *string
	SetDuration(v int64) *StsApplyVO
	GetDuration() *int64
}

type StsApplyVO struct {
	AliyunId *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	Duration *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s StsApplyVO) String() string {
	return dara.Prettify(s)
}

func (s StsApplyVO) GoString() string {
	return s.String()
}

func (s *StsApplyVO) GetAliyunId() *string {
	return s.AliyunId
}

func (s *StsApplyVO) GetDuration() *int64 {
	return s.Duration
}

func (s *StsApplyVO) SetAliyunId(v string) *StsApplyVO {
	s.AliyunId = &v
	return s
}

func (s *StsApplyVO) SetDuration(v int64) *StsApplyVO {
	s.Duration = &v
	return s
}

func (s *StsApplyVO) Validate() error {
	return dara.Validate(s)
}
