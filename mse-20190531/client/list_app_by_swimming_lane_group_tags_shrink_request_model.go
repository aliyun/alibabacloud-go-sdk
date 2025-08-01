// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppBySwimmingLaneGroupTagsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListAppBySwimmingLaneGroupTagsShrinkRequest
	GetAcceptLanguage() *string
	SetGroupId(v int64) *ListAppBySwimmingLaneGroupTagsShrinkRequest
	GetGroupId() *int64
	SetNamespace(v string) *ListAppBySwimmingLaneGroupTagsShrinkRequest
	GetNamespace() *string
	SetTagsShrink(v string) *ListAppBySwimmingLaneGroupTagsShrinkRequest
	GetTagsShrink() *string
}

type ListAppBySwimmingLaneGroupTagsShrinkRequest struct {
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListAppBySwimmingLaneGroupTagsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppBySwimmingLaneGroupTagsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAppBySwimmingLaneGroupTagsShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListAppBySwimmingLaneGroupTagsShrinkRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListAppBySwimmingLaneGroupTagsShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListAppBySwimmingLaneGroupTagsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListAppBySwimmingLaneGroupTagsShrinkRequest) SetAcceptLanguage(v string) *ListAppBySwimmingLaneGroupTagsShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagsShrinkRequest) SetGroupId(v int64) *ListAppBySwimmingLaneGroupTagsShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagsShrinkRequest) SetNamespace(v string) *ListAppBySwimmingLaneGroupTagsShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagsShrinkRequest) SetTagsShrink(v string) *ListAppBySwimmingLaneGroupTagsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
