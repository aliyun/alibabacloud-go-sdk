// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomDomainOutput interface {
	dara.Model
	String() string
	GoString() string
	SetCustomDomains(v []*CustomDomain) *ListCustomDomainOutput
	GetCustomDomains() []*CustomDomain
	SetNextToken(v string) *ListCustomDomainOutput
	GetNextToken() *string
}

type ListCustomDomainOutput struct {
	// The custom domain names.
	CustomDomains []*CustomDomain `json:"customDomains" xml:"customDomains" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// next_domain_name
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListCustomDomainOutput) String() string {
	return dara.Prettify(s)
}

func (s ListCustomDomainOutput) GoString() string {
	return s.String()
}

func (s *ListCustomDomainOutput) GetCustomDomains() []*CustomDomain {
	return s.CustomDomains
}

func (s *ListCustomDomainOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomDomainOutput) SetCustomDomains(v []*CustomDomain) *ListCustomDomainOutput {
	s.CustomDomains = v
	return s
}

func (s *ListCustomDomainOutput) SetNextToken(v string) *ListCustomDomainOutput {
	s.NextToken = &v
	return s
}

func (s *ListCustomDomainOutput) Validate() error {
	if s.CustomDomains != nil {
		for _, item := range s.CustomDomains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
