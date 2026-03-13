// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaggedResourcesOutput interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTaggedResourcesOutput
	GetNextToken() *string
	SetResources(v []*Resource) *ListTaggedResourcesOutput
	GetResources() []*Resource
}

type ListTaggedResourcesOutput struct {
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The tagged resources.
	Resources []*Resource `json:"resources" xml:"resources" type:"Repeated"`
}

func (s ListTaggedResourcesOutput) String() string {
	return dara.Prettify(s)
}

func (s ListTaggedResourcesOutput) GoString() string {
	return s.String()
}

func (s *ListTaggedResourcesOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTaggedResourcesOutput) GetResources() []*Resource {
	return s.Resources
}

func (s *ListTaggedResourcesOutput) SetNextToken(v string) *ListTaggedResourcesOutput {
	s.NextToken = &v
	return s
}

func (s *ListTaggedResourcesOutput) SetResources(v []*Resource) *ListTaggedResourcesOutput {
	s.Resources = v
	return s
}

func (s *ListTaggedResourcesOutput) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
