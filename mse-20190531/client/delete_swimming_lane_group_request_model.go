// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSwimmingLaneGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteSwimmingLaneGroupRequest
	GetAcceptLanguage() *string
	SetGroupId(v int64) *DeleteSwimmingLaneGroupRequest
	GetGroupId() *int64
	SetName(v string) *DeleteSwimmingLaneGroupRequest
	GetName() *string
	SetNamespace(v string) *DeleteSwimmingLaneGroupRequest
	GetNamespace() *string
}

type DeleteSwimmingLaneGroupRequest struct {
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
	// example:
	//
	// 145
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name.
	//
	// example:
	//
	// example-app
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the Microservices Engine (MSE) namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DeleteSwimmingLaneGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSwimmingLaneGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSwimmingLaneGroupRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteSwimmingLaneGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DeleteSwimmingLaneGroupRequest) GetName() *string {
	return s.Name
}

func (s *DeleteSwimmingLaneGroupRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteSwimmingLaneGroupRequest) SetAcceptLanguage(v string) *DeleteSwimmingLaneGroupRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteSwimmingLaneGroupRequest) SetGroupId(v int64) *DeleteSwimmingLaneGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteSwimmingLaneGroupRequest) SetName(v string) *DeleteSwimmingLaneGroupRequest {
	s.Name = &v
	return s
}

func (s *DeleteSwimmingLaneGroupRequest) SetNamespace(v string) *DeleteSwimmingLaneGroupRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteSwimmingLaneGroupRequest) Validate() error {
	return dara.Validate(s)
}
