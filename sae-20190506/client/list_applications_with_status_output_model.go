// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsWithStatusOutput interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ApplicationWithStatus) *ListApplicationsWithStatusOutput
	GetApplications() []*ApplicationWithStatus
	SetNextToken(v string) *ListApplicationsWithStatusOutput
	GetNextToken() *string
	SetRequestId(v string) *ListApplicationsWithStatusOutput
	GetRequestId() *string
}

type ListApplicationsWithStatusOutput struct {
	Applications []*ApplicationWithStatus `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	NextToken    *string                  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId    *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListApplicationsWithStatusOutput) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsWithStatusOutput) GoString() string {
	return s.String()
}

func (s *ListApplicationsWithStatusOutput) GetApplications() []*ApplicationWithStatus {
	return s.Applications
}

func (s *ListApplicationsWithStatusOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationsWithStatusOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsWithStatusOutput) SetApplications(v []*ApplicationWithStatus) *ListApplicationsWithStatusOutput {
	s.Applications = v
	return s
}

func (s *ListApplicationsWithStatusOutput) SetNextToken(v string) *ListApplicationsWithStatusOutput {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsWithStatusOutput) SetRequestId(v string) *ListApplicationsWithStatusOutput {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsWithStatusOutput) Validate() error {
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
