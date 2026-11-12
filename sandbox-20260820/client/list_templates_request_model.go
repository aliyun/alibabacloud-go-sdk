// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTemplatesRequest
	GetNextToken() *string
	SetResourceTags(v []*ListTemplatesRequestResourceTags) *ListTemplatesRequest
	GetResourceTags() []*ListTemplatesRequestResourceTags
	SetTeamID(v string) *ListTemplatesRequest
	GetTeamID() *string
}

type ListTemplatesRequest struct {
	// The maximum number of entries to return per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Set this parameter to the value of nextToken returned in the response of the previous page.
	//
	// example:
	//
	// cae5f900-8b1d-4c0e-9c2a-1a2b3c4d5e6f
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The resource tags.
	ResourceTags []*ListTemplatesRequestResourceTags `json:"resourceTags" xml:"resourceTags" type:"Repeated"`
	// The team ID.
	//
	// example:
	//
	// 13b721e6-8cc8-5df2-af13-80316f7508af
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplatesRequest) GetResourceTags() []*ListTemplatesRequestResourceTags {
	return s.ResourceTags
}

func (s *ListTemplatesRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *ListTemplatesRequest) SetMaxResults(v int32) *ListTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesRequest) SetNextToken(v string) *ListTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplatesRequest) SetResourceTags(v []*ListTemplatesRequestResourceTags) *ListTemplatesRequest {
	s.ResourceTags = v
	return s
}

func (s *ListTemplatesRequest) SetTeamID(v string) *ListTemplatesRequest {
	s.TeamID = &v
	return s
}

func (s *ListTemplatesRequest) Validate() error {
	if s.ResourceTags != nil {
		for _, item := range s.ResourceTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTemplatesRequestResourceTags struct {
	// The tag key.
	//
	// example:
	//
	// project
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// demo
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListTemplatesRequestResourceTags) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesRequestResourceTags) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequestResourceTags) GetKey() *string {
	return s.Key
}

func (s *ListTemplatesRequestResourceTags) GetValue() *string {
	return s.Value
}

func (s *ListTemplatesRequestResourceTags) SetKey(v string) *ListTemplatesRequestResourceTags {
	s.Key = &v
	return s
}

func (s *ListTemplatesRequestResourceTags) SetValue(v string) *ListTemplatesRequestResourceTags {
	s.Value = &v
	return s
}

func (s *ListTemplatesRequestResourceTags) Validate() error {
	return dara.Validate(s)
}
