// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebApplicationsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListWebApplicationsOutput
	GetNextToken() *string
	SetWebApplicationWithInstanceCount(v []*WebApplicationWithInstanceCount) *ListWebApplicationsOutput
	GetWebApplicationWithInstanceCount() []*WebApplicationWithInstanceCount
}

type ListWebApplicationsOutput struct {
	NextToken                       *string                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	WebApplicationWithInstanceCount []*WebApplicationWithInstanceCount `json:"WebApplicationWithInstanceCount,omitempty" xml:"WebApplicationWithInstanceCount,omitempty" type:"Repeated"`
}

func (s ListWebApplicationsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListWebApplicationsOutput) GoString() string {
	return s.String()
}

func (s *ListWebApplicationsOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWebApplicationsOutput) GetWebApplicationWithInstanceCount() []*WebApplicationWithInstanceCount {
	return s.WebApplicationWithInstanceCount
}

func (s *ListWebApplicationsOutput) SetNextToken(v string) *ListWebApplicationsOutput {
	s.NextToken = &v
	return s
}

func (s *ListWebApplicationsOutput) SetWebApplicationWithInstanceCount(v []*WebApplicationWithInstanceCount) *ListWebApplicationsOutput {
	s.WebApplicationWithInstanceCount = v
	return s
}

func (s *ListWebApplicationsOutput) Validate() error {
	if s.WebApplicationWithInstanceCount != nil {
		for _, item := range s.WebApplicationWithInstanceCount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
