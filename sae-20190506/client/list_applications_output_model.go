// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*Application) *ListApplicationsOutput
	GetApplications() []*Application
	SetNextToken(v string) *ListApplicationsOutput
	GetNextToken() *string
	SetRequestId(v string) *ListApplicationsOutput
	GetRequestId() *string
}

type ListApplicationsOutput struct {
	Applications []*Application `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	NextToken    *string        `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId    *string        `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListApplicationsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsOutput) GoString() string {
	return s.String()
}

func (s *ListApplicationsOutput) GetApplications() []*Application {
	return s.Applications
}

func (s *ListApplicationsOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationsOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsOutput) SetApplications(v []*Application) *ListApplicationsOutput {
	s.Applications = v
	return s
}

func (s *ListApplicationsOutput) SetNextToken(v string) *ListApplicationsOutput {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsOutput) SetRequestId(v string) *ListApplicationsOutput {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsOutput) Validate() error {
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
