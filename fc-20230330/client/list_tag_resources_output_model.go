// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesOutput interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagResourcesOutput
	GetNextToken() *string
	SetRequestId(v string) *ListTagResourcesOutput
	GetRequestId() *string
	SetTagResources(v []*TagResource) *ListTagResourcesOutput
	GetTagResources() []*TagResource
}

type ListTagResourcesOutput struct {
	// example:
	//
	// next_token
	NextToken    *string        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*TagResource `json:"TagResources" xml:"TagResources" type:"Repeated"`
}

func (s ListTagResourcesOutput) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesOutput) GoString() string {
	return s.String()
}

func (s *ListTagResourcesOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagResourcesOutput) GetTagResources() []*TagResource {
	return s.TagResources
}

func (s *ListTagResourcesOutput) SetNextToken(v string) *ListTagResourcesOutput {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesOutput) SetRequestId(v string) *ListTagResourcesOutput {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesOutput) SetTagResources(v []*TagResource) *ListTagResourcesOutput {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesOutput) Validate() error {
	return dara.Validate(s)
}
