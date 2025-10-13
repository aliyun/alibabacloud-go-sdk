// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebApplicationRevisionsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListWebApplicationRevisionsOutput
	GetNextToken() *string
	SetRevisions(v []*Revision) *ListWebApplicationRevisionsOutput
	GetRevisions() []*Revision
}

type ListWebApplicationRevisionsOutput struct {
	NextToken *string     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Revisions []*Revision `json:"Revisions,omitempty" xml:"Revisions,omitempty" type:"Repeated"`
}

func (s ListWebApplicationRevisionsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListWebApplicationRevisionsOutput) GoString() string {
	return s.String()
}

func (s *ListWebApplicationRevisionsOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWebApplicationRevisionsOutput) GetRevisions() []*Revision {
	return s.Revisions
}

func (s *ListWebApplicationRevisionsOutput) SetNextToken(v string) *ListWebApplicationRevisionsOutput {
	s.NextToken = &v
	return s
}

func (s *ListWebApplicationRevisionsOutput) SetRevisions(v []*Revision) *ListWebApplicationRevisionsOutput {
	s.Revisions = v
	return s
}

func (s *ListWebApplicationRevisionsOutput) Validate() error {
	if s.Revisions != nil {
		for _, item := range s.Revisions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
