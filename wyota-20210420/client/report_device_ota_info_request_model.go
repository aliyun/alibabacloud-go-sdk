// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportDeviceOtaInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseVersion(v string) *ReportDeviceOtaInfoRequest
	GetBaseVersion() *string
	SetDeviceId(v string) *ReportDeviceOtaInfoRequest
	GetDeviceId() *string
	SetModel(v string) *ReportDeviceOtaInfoRequest
	GetModel() *string
	SetNote(v string) *ReportDeviceOtaInfoRequest
	GetNote() *string
	SetStatus(v int32) *ReportDeviceOtaInfoRequest
	GetStatus() *int32
	SetTargetVersion(v string) *ReportDeviceOtaInfoRequest
	GetTargetVersion() *string
}

type ReportDeviceOtaInfoRequest struct {
	BaseVersion   *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Model         *string `json:"Model,omitempty" xml:"Model,omitempty"`
	Note          *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
}

func (s ReportDeviceOtaInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportDeviceOtaInfoRequest) GoString() string {
	return s.String()
}

func (s *ReportDeviceOtaInfoRequest) GetBaseVersion() *string {
	return s.BaseVersion
}

func (s *ReportDeviceOtaInfoRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ReportDeviceOtaInfoRequest) GetModel() *string {
	return s.Model
}

func (s *ReportDeviceOtaInfoRequest) GetNote() *string {
	return s.Note
}

func (s *ReportDeviceOtaInfoRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ReportDeviceOtaInfoRequest) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *ReportDeviceOtaInfoRequest) SetBaseVersion(v string) *ReportDeviceOtaInfoRequest {
	s.BaseVersion = &v
	return s
}

func (s *ReportDeviceOtaInfoRequest) SetDeviceId(v string) *ReportDeviceOtaInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *ReportDeviceOtaInfoRequest) SetModel(v string) *ReportDeviceOtaInfoRequest {
	s.Model = &v
	return s
}

func (s *ReportDeviceOtaInfoRequest) SetNote(v string) *ReportDeviceOtaInfoRequest {
	s.Note = &v
	return s
}

func (s *ReportDeviceOtaInfoRequest) SetStatus(v int32) *ReportDeviceOtaInfoRequest {
	s.Status = &v
	return s
}

func (s *ReportDeviceOtaInfoRequest) SetTargetVersion(v string) *ReportDeviceOtaInfoRequest {
	s.TargetVersion = &v
	return s
}

func (s *ReportDeviceOtaInfoRequest) Validate() error {
	return dara.Validate(s)
}
