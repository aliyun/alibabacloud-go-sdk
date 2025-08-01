// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTagsBySwimmingLaneGroupIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetTagsBySwimmingLaneGroupIdRequest
	GetAcceptLanguage() *string
	SetGroupId(v int64) *GetTagsBySwimmingLaneGroupIdRequest
	GetGroupId() *int64
	SetNamespace(v string) *GetTagsBySwimmingLaneGroupIdRequest
	GetNamespace() *string
}

type GetTagsBySwimmingLaneGroupIdRequest struct {
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
	// 154
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the Microservices Engine (MSE) namespace that you want to query.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetTagsBySwimmingLaneGroupIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTagsBySwimmingLaneGroupIdRequest) GoString() string {
	return s.String()
}

func (s *GetTagsBySwimmingLaneGroupIdRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetTagsBySwimmingLaneGroupIdRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *GetTagsBySwimmingLaneGroupIdRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetTagsBySwimmingLaneGroupIdRequest) SetAcceptLanguage(v string) *GetTagsBySwimmingLaneGroupIdRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetTagsBySwimmingLaneGroupIdRequest) SetGroupId(v int64) *GetTagsBySwimmingLaneGroupIdRequest {
	s.GroupId = &v
	return s
}

func (s *GetTagsBySwimmingLaneGroupIdRequest) SetNamespace(v string) *GetTagsBySwimmingLaneGroupIdRequest {
	s.Namespace = &v
	return s
}

func (s *GetTagsBySwimmingLaneGroupIdRequest) Validate() error {
	return dara.Validate(s)
}
