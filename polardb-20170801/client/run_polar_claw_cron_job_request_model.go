// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPolarClawCronJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RunPolarClawCronJobRequest
	GetApplicationId() *string
	SetJobId(v string) *RunPolarClawCronJobRequest
	GetJobId() *string
	SetMode(v string) *RunPolarClawCronJobRequest
	GetMode() *string
	SetRestart(v bool) *RunPolarClawCronJobRequest
	GetRestart() *bool
}

type RunPolarClawCronJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0ee00f56-f467-4d41-858c-ca4ede2c770e
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// force
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s RunPolarClawCronJobRequest) String() string {
	return dara.Prettify(s)
}

func (s RunPolarClawCronJobRequest) GoString() string {
	return s.String()
}

func (s *RunPolarClawCronJobRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RunPolarClawCronJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *RunPolarClawCronJobRequest) GetMode() *string {
	return s.Mode
}

func (s *RunPolarClawCronJobRequest) GetRestart() *bool {
	return s.Restart
}

func (s *RunPolarClawCronJobRequest) SetApplicationId(v string) *RunPolarClawCronJobRequest {
	s.ApplicationId = &v
	return s
}

func (s *RunPolarClawCronJobRequest) SetJobId(v string) *RunPolarClawCronJobRequest {
	s.JobId = &v
	return s
}

func (s *RunPolarClawCronJobRequest) SetMode(v string) *RunPolarClawCronJobRequest {
	s.Mode = &v
	return s
}

func (s *RunPolarClawCronJobRequest) SetRestart(v bool) *RunPolarClawCronJobRequest {
	s.Restart = &v
	return s
}

func (s *RunPolarClawCronJobRequest) Validate() error {
	return dara.Validate(s)
}
