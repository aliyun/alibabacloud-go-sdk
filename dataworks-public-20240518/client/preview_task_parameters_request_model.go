// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewTaskParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizdate(v int64) *PreviewTaskParametersRequest
	GetBizdate() *int64
	SetCron(v string) *PreviewTaskParametersRequest
	GetCron() *string
	SetParameters(v string) *PreviewTaskParametersRequest
	GetParameters() *string
	SetPreviewCount(v int32) *PreviewTaskParametersRequest
	GetPreviewCount() *int32
	SetProjectEnv(v string) *PreviewTaskParametersRequest
	GetProjectEnv() *string
	SetProjectId(v int64) *PreviewTaskParametersRequest
	GetProjectId() *int64
	SetStartTime(v string) *PreviewTaskParametersRequest
	GetStartTime() *string
}

type PreviewTaskParametersRequest struct {
	// The business date. The value is a 13-digit UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1792425600000
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The scheduling CRON expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The list of script parameters. The length cannot exceed 65,535 characters. Separate multiple parameters with spaces, such as para1=$bizdate para2=$[hh24]. An empty string is allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// para1=$bizdate para2=$[hh24]
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The number of scheduling cycles to preview. Valid values: 1 to 10. Default value: 3.
	//
	// example:
	//
	// 3
	PreviewCount *int32 `json:"PreviewCount,omitempty" xml:"PreviewCount,omitempty"`
	// The project environment. Valid values:
	//
	// - Prod: production
	//
	// - Dev: development
	//
	// Default value: Prod.
	//
	// example:
	//
	// Prod
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The workspace ID. The workspace must belong to the current tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The preview start time for the first scheduling day. The format is HH:mm:ss. Default value: 00:00:00. This parameter takes effect only for non-daily schedules and is used to filter out cycles that occur before this time on the first scheduling day.
	//
	// example:
	//
	// 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PreviewTaskParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewTaskParametersRequest) GoString() string {
	return s.String()
}

func (s *PreviewTaskParametersRequest) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *PreviewTaskParametersRequest) GetCron() *string {
	return s.Cron
}

func (s *PreviewTaskParametersRequest) GetParameters() *string {
	return s.Parameters
}

func (s *PreviewTaskParametersRequest) GetPreviewCount() *int32 {
	return s.PreviewCount
}

func (s *PreviewTaskParametersRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *PreviewTaskParametersRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *PreviewTaskParametersRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *PreviewTaskParametersRequest) SetBizdate(v int64) *PreviewTaskParametersRequest {
	s.Bizdate = &v
	return s
}

func (s *PreviewTaskParametersRequest) SetCron(v string) *PreviewTaskParametersRequest {
	s.Cron = &v
	return s
}

func (s *PreviewTaskParametersRequest) SetParameters(v string) *PreviewTaskParametersRequest {
	s.Parameters = &v
	return s
}

func (s *PreviewTaskParametersRequest) SetPreviewCount(v int32) *PreviewTaskParametersRequest {
	s.PreviewCount = &v
	return s
}

func (s *PreviewTaskParametersRequest) SetProjectEnv(v string) *PreviewTaskParametersRequest {
	s.ProjectEnv = &v
	return s
}

func (s *PreviewTaskParametersRequest) SetProjectId(v int64) *PreviewTaskParametersRequest {
	s.ProjectId = &v
	return s
}

func (s *PreviewTaskParametersRequest) SetStartTime(v string) *PreviewTaskParametersRequest {
	s.StartTime = &v
	return s
}

func (s *PreviewTaskParametersRequest) Validate() error {
	return dara.Validate(s)
}
