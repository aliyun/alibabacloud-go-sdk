// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppBySwimmingLaneGroupTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListAppBySwimmingLaneGroupTagsRequest
	GetAcceptLanguage() *string
	SetGroupId(v int64) *ListAppBySwimmingLaneGroupTagsRequest
	GetGroupId() *int64
	SetNamespace(v string) *ListAppBySwimmingLaneGroupTagsRequest
	GetNamespace() *string
	SetTags(v []*string) *ListAppBySwimmingLaneGroupTagsRequest
	GetTags() []*string
}

type ListAppBySwimmingLaneGroupTagsRequest struct {
	// The language in which you want to display the results. Valid values: zh and en. zh indicates Chinese, which is the default value. en indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the lane group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the MSE namespace that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The tag based on which you want to list applications.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListAppBySwimmingLaneGroupTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppBySwimmingLaneGroupTagsRequest) GoString() string {
	return s.String()
}

func (s *ListAppBySwimmingLaneGroupTagsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListAppBySwimmingLaneGroupTagsRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListAppBySwimmingLaneGroupTagsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListAppBySwimmingLaneGroupTagsRequest) GetTags() []*string {
	return s.Tags
}

func (s *ListAppBySwimmingLaneGroupTagsRequest) SetAcceptLanguage(v string) *ListAppBySwimmingLaneGroupTagsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagsRequest) SetGroupId(v int64) *ListAppBySwimmingLaneGroupTagsRequest {
	s.GroupId = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagsRequest) SetNamespace(v string) *ListAppBySwimmingLaneGroupTagsRequest {
	s.Namespace = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagsRequest) SetTags(v []*string) *ListAppBySwimmingLaneGroupTagsRequest {
	s.Tags = v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagsRequest) Validate() error {
	return dara.Validate(s)
}
