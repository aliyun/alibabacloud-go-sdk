// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisablePolarClawCronJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DisablePolarClawCronJobRequest
	GetApplicationId() *string
	SetJobId(v string) *DisablePolarClawCronJobRequest
	GetJobId() *string
	SetRestart(v bool) *DisablePolarClawCronJobRequest
	GetRestart() *bool
}

type DisablePolarClawCronJobRequest struct {
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
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s DisablePolarClawCronJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DisablePolarClawCronJobRequest) GoString() string {
	return s.String()
}

func (s *DisablePolarClawCronJobRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DisablePolarClawCronJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *DisablePolarClawCronJobRequest) GetRestart() *bool {
	return s.Restart
}

func (s *DisablePolarClawCronJobRequest) SetApplicationId(v string) *DisablePolarClawCronJobRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisablePolarClawCronJobRequest) SetJobId(v string) *DisablePolarClawCronJobRequest {
	s.JobId = &v
	return s
}

func (s *DisablePolarClawCronJobRequest) SetRestart(v bool) *DisablePolarClawCronJobRequest {
	s.Restart = &v
	return s
}

func (s *DisablePolarClawCronJobRequest) Validate() error {
	return dara.Validate(s)
}
