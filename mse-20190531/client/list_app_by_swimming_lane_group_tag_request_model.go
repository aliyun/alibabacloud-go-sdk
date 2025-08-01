// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppBySwimmingLaneGroupTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListAppBySwimmingLaneGroupTagRequest
	GetAcceptLanguage() *string
	SetGroupId(v int64) *ListAppBySwimmingLaneGroupTagRequest
	GetGroupId() *int64
	SetNamespace(v string) *ListAppBySwimmingLaneGroupTagRequest
	GetNamespace() *string
	SetTag(v string) *ListAppBySwimmingLaneGroupTagRequest
	GetTag() *string
}

type ListAppBySwimmingLaneGroupTagRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
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
	// 119
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the Microservices Engine (MSE) namespace that you want to query.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// gray
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListAppBySwimmingLaneGroupTagRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppBySwimmingLaneGroupTagRequest) GoString() string {
	return s.String()
}

func (s *ListAppBySwimmingLaneGroupTagRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListAppBySwimmingLaneGroupTagRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListAppBySwimmingLaneGroupTagRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListAppBySwimmingLaneGroupTagRequest) GetTag() *string {
	return s.Tag
}

func (s *ListAppBySwimmingLaneGroupTagRequest) SetAcceptLanguage(v string) *ListAppBySwimmingLaneGroupTagRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagRequest) SetGroupId(v int64) *ListAppBySwimmingLaneGroupTagRequest {
	s.GroupId = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagRequest) SetNamespace(v string) *ListAppBySwimmingLaneGroupTagRequest {
	s.Namespace = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagRequest) SetTag(v string) *ListAppBySwimmingLaneGroupTagRequest {
	s.Tag = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagRequest) Validate() error {
	return dara.Validate(s)
}
