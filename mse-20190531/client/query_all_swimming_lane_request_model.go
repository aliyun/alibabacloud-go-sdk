// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAllSwimmingLaneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QueryAllSwimmingLaneRequest
	GetAcceptLanguage() *string
	SetGroupId(v int64) *QueryAllSwimmingLaneRequest
	GetGroupId() *int64
	SetNamespace(v string) *QueryAllSwimmingLaneRequest
	GetNamespace() *string
}

type QueryAllSwimmingLaneRequest struct {
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
	// 186
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the MSE namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s QueryAllSwimmingLaneRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAllSwimmingLaneRequest) GoString() string {
	return s.String()
}

func (s *QueryAllSwimmingLaneRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QueryAllSwimmingLaneRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *QueryAllSwimmingLaneRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *QueryAllSwimmingLaneRequest) SetAcceptLanguage(v string) *QueryAllSwimmingLaneRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QueryAllSwimmingLaneRequest) SetGroupId(v int64) *QueryAllSwimmingLaneRequest {
	s.GroupId = &v
	return s
}

func (s *QueryAllSwimmingLaneRequest) SetNamespace(v string) *QueryAllSwimmingLaneRequest {
	s.Namespace = &v
	return s
}

func (s *QueryAllSwimmingLaneRequest) Validate() error {
	return dara.Validate(s)
}
