// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarClawCronJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DeletePolarClawCronJobRequest
	GetApplicationId() *string
	SetJobId(v string) *DeletePolarClawCronJobRequest
	GetJobId() *string
	SetRestart(v bool) *DeletePolarClawCronJobRequest
	GetRestart() *bool
}

type DeletePolarClawCronJobRequest struct {
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

func (s DeletePolarClawCronJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarClawCronJobRequest) GoString() string {
	return s.String()
}

func (s *DeletePolarClawCronJobRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeletePolarClawCronJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *DeletePolarClawCronJobRequest) GetRestart() *bool {
	return s.Restart
}

func (s *DeletePolarClawCronJobRequest) SetApplicationId(v string) *DeletePolarClawCronJobRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeletePolarClawCronJobRequest) SetJobId(v string) *DeletePolarClawCronJobRequest {
	s.JobId = &v
	return s
}

func (s *DeletePolarClawCronJobRequest) SetRestart(v bool) *DeletePolarClawCronJobRequest {
	s.Restart = &v
	return s
}

func (s *DeletePolarClawCronJobRequest) Validate() error {
	return dara.Validate(s)
}
