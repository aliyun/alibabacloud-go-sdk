// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodUserUsageDetailDataExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimension(v string) *CreateVodUserUsageDetailDataExportTaskRequest
	GetDimension() *string
	SetDomainNames(v string) *CreateVodUserUsageDetailDataExportTaskRequest
	GetDomainNames() *string
	SetEndTime(v string) *CreateVodUserUsageDetailDataExportTaskRequest
	GetEndTime() *string
	SetGroup(v string) *CreateVodUserUsageDetailDataExportTaskRequest
	GetGroup() *string
	SetLanguage(v string) *CreateVodUserUsageDetailDataExportTaskRequest
	GetLanguage() *string
	SetOwnerId(v int64) *CreateVodUserUsageDetailDataExportTaskRequest
	GetOwnerId() *int64
	SetStartTime(v string) *CreateVodUserUsageDetailDataExportTaskRequest
	GetStartTime() *string
	SetTaskName(v string) *CreateVodUserUsageDetailDataExportTaskRequest
	GetTaskName() *string
}

type CreateVodUserUsageDetailDataExportTaskRequest struct {
	// This parameter is required.
	Dimension   *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// This parameter is required.
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Group    *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskName  *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateVodUserUsageDetailDataExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVodUserUsageDetailDataExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) GetDimension() *string {
	return s.Dimension
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) GetGroup() *string {
	return s.Group
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) SetDimension(v string) *CreateVodUserUsageDetailDataExportTaskRequest {
	s.Dimension = &v
	return s
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) SetDomainNames(v string) *CreateVodUserUsageDetailDataExportTaskRequest {
	s.DomainNames = &v
	return s
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) SetEndTime(v string) *CreateVodUserUsageDetailDataExportTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) SetGroup(v string) *CreateVodUserUsageDetailDataExportTaskRequest {
	s.Group = &v
	return s
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) SetLanguage(v string) *CreateVodUserUsageDetailDataExportTaskRequest {
	s.Language = &v
	return s
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) SetOwnerId(v int64) *CreateVodUserUsageDetailDataExportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) SetStartTime(v string) *CreateVodUserUsageDetailDataExportTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) SetTaskName(v string) *CreateVodUserUsageDetailDataExportTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateVodUserUsageDetailDataExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
