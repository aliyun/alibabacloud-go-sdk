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
	SetRequestId(v string) *ListCustomDomainOutput
	GetRequestId() *string
}

type ListCustomDomainOutput struct {
	CustomDomains []*CustomDomain `json:"customDomains,omitempty" xml:"customDomains,omitempty" type:"Repeated"`
	NextToken     *string         `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId     *string         `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

func (s *ListCustomDomainOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomDomainOutput) SetCustomDomains(v []*CustomDomain) *ListCustomDomainOutput {
	s.CustomDomains = v
	return s
}

func (s *ListCustomDomainOutput) SetNextToken(v string) *ListCustomDomainOutput {
	s.NextToken = &v
	return s
}

func (s *ListCustomDomainOutput) SetRequestId(v string) *ListCustomDomainOutput {
	s.RequestId = &v
	return s
}

func (s *ListCustomDomainOutput) Validate() error {
	return dara.Validate(s)
}
