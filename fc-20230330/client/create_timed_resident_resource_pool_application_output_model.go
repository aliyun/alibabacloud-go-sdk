// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTimedResidentResourcePoolApplicationOutput interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationStatus(v string) *CreateTimedResidentResourcePoolApplicationOutput
	GetApplicationStatus() *string
	SetTimedPoolId(v string) *CreateTimedResidentResourcePoolApplicationOutput
	GetTimedPoolId() *string
}

type CreateTimedResidentResourcePoolApplicationOutput struct {
	ApplicationStatus *string `json:"applicationStatus,omitempty" xml:"applicationStatus,omitempty"`
	TimedPoolId       *string `json:"timedPoolId,omitempty" xml:"timedPoolId,omitempty"`
}

func (s CreateTimedResidentResourcePoolApplicationOutput) String() string {
	return dara.Prettify(s)
}

func (s CreateTimedResidentResourcePoolApplicationOutput) GoString() string {
	return s.String()
}

func (s *CreateTimedResidentResourcePoolApplicationOutput) GetApplicationStatus() *string {
	return s.ApplicationStatus
}

func (s *CreateTimedResidentResourcePoolApplicationOutput) GetTimedPoolId() *string {
	return s.TimedPoolId
}

func (s *CreateTimedResidentResourcePoolApplicationOutput) SetApplicationStatus(v string) *CreateTimedResidentResourcePoolApplicationOutput {
	s.ApplicationStatus = &v
	return s
}

func (s *CreateTimedResidentResourcePoolApplicationOutput) SetTimedPoolId(v string) *CreateTimedResidentResourcePoolApplicationOutput {
	s.TimedPoolId = &v
	return s
}

func (s *CreateTimedResidentResourcePoolApplicationOutput) Validate() error {
	return dara.Validate(s)
}
