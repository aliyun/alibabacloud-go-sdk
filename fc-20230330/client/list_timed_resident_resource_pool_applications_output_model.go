// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTimedResidentResourcePoolApplicationsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*TimedResidentResourcePoolApplication) *ListTimedResidentResourcePoolApplicationsOutput
	GetApplications() []*TimedResidentResourcePoolApplication
	SetNextToken(v string) *ListTimedResidentResourcePoolApplicationsOutput
	GetNextToken() *string
}

type ListTimedResidentResourcePoolApplicationsOutput struct {
	Applications []*TimedResidentResourcePoolApplication `json:"applications" xml:"applications" type:"Repeated"`
	NextToken    *string                                 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListTimedResidentResourcePoolApplicationsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListTimedResidentResourcePoolApplicationsOutput) GoString() string {
	return s.String()
}

func (s *ListTimedResidentResourcePoolApplicationsOutput) GetApplications() []*TimedResidentResourcePoolApplication {
	return s.Applications
}

func (s *ListTimedResidentResourcePoolApplicationsOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTimedResidentResourcePoolApplicationsOutput) SetApplications(v []*TimedResidentResourcePoolApplication) *ListTimedResidentResourcePoolApplicationsOutput {
	s.Applications = v
	return s
}

func (s *ListTimedResidentResourcePoolApplicationsOutput) SetNextToken(v string) *ListTimedResidentResourcePoolApplicationsOutput {
	s.NextToken = &v
	return s
}

func (s *ListTimedResidentResourcePoolApplicationsOutput) Validate() error {
	if s.Applications != nil {
		for _, item := range s.Applications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
