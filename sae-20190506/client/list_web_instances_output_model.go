// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebInstancesOutput interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentError(v string) *ListWebInstancesOutput
	GetCurrentError() *string
	SetWebInstances(v []*WebInstanceInfo) *ListWebInstancesOutput
	GetWebInstances() []*WebInstanceInfo
	SetWebVersionStatus(v map[string]*WebVersionStatus) *ListWebInstancesOutput
	GetWebVersionStatus() map[string]*WebVersionStatus
}

type ListWebInstancesOutput struct {
	CurrentError     *string                      `json:"CurrentError,omitempty" xml:"CurrentError,omitempty"`
	WebInstances     []*WebInstanceInfo           `json:"WebInstances,omitempty" xml:"WebInstances,omitempty" type:"Repeated"`
	WebVersionStatus map[string]*WebVersionStatus `json:"WebVersionStatus,omitempty" xml:"WebVersionStatus,omitempty"`
}

func (s ListWebInstancesOutput) String() string {
	return dara.Prettify(s)
}

func (s ListWebInstancesOutput) GoString() string {
	return s.String()
}

func (s *ListWebInstancesOutput) GetCurrentError() *string {
	return s.CurrentError
}

func (s *ListWebInstancesOutput) GetWebInstances() []*WebInstanceInfo {
	return s.WebInstances
}

func (s *ListWebInstancesOutput) GetWebVersionStatus() map[string]*WebVersionStatus {
	return s.WebVersionStatus
}

func (s *ListWebInstancesOutput) SetCurrentError(v string) *ListWebInstancesOutput {
	s.CurrentError = &v
	return s
}

func (s *ListWebInstancesOutput) SetWebInstances(v []*WebInstanceInfo) *ListWebInstancesOutput {
	s.WebInstances = v
	return s
}

func (s *ListWebInstancesOutput) SetWebVersionStatus(v map[string]*WebVersionStatus) *ListWebInstancesOutput {
	s.WebVersionStatus = v
	return s
}

func (s *ListWebInstancesOutput) Validate() error {
	return dara.Validate(s)
}
