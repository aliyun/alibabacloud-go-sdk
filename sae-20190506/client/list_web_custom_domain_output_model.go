// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebCustomDomainOutput interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListWebCustomDomainOutput
	GetNextToken() *string
	SetWebCustomDomains(v []*WebCustomDomain) *ListWebCustomDomainOutput
	GetWebCustomDomains() []*WebCustomDomain
}

type ListWebCustomDomainOutput struct {
	NextToken        *string            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	WebCustomDomains []*WebCustomDomain `json:"WebCustomDomains,omitempty" xml:"WebCustomDomains,omitempty" type:"Repeated"`
}

func (s ListWebCustomDomainOutput) String() string {
	return dara.Prettify(s)
}

func (s ListWebCustomDomainOutput) GoString() string {
	return s.String()
}

func (s *ListWebCustomDomainOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWebCustomDomainOutput) GetWebCustomDomains() []*WebCustomDomain {
	return s.WebCustomDomains
}

func (s *ListWebCustomDomainOutput) SetNextToken(v string) *ListWebCustomDomainOutput {
	s.NextToken = &v
	return s
}

func (s *ListWebCustomDomainOutput) SetWebCustomDomains(v []*WebCustomDomain) *ListWebCustomDomainOutput {
	s.WebCustomDomains = v
	return s
}

func (s *ListWebCustomDomainOutput) Validate() error {
	if s.WebCustomDomains != nil {
		for _, item := range s.WebCustomDomains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
