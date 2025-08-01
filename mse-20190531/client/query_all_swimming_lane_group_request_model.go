// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAllSwimmingLaneGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QueryAllSwimmingLaneGroupRequest
	GetAcceptLanguage() *string
	SetNamespace(v string) *QueryAllSwimmingLaneGroupRequest
	GetNamespace() *string
}

type QueryAllSwimmingLaneGroupRequest struct {
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
	// The name of the Microservices Engine (MSE) namespace that you want to query.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s QueryAllSwimmingLaneGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAllSwimmingLaneGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryAllSwimmingLaneGroupRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QueryAllSwimmingLaneGroupRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *QueryAllSwimmingLaneGroupRequest) SetAcceptLanguage(v string) *QueryAllSwimmingLaneGroupRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupRequest) SetNamespace(v string) *QueryAllSwimmingLaneGroupRequest {
	s.Namespace = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupRequest) Validate() error {
	return dara.Validate(s)
}
