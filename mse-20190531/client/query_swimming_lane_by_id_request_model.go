// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySwimmingLaneByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QuerySwimmingLaneByIdRequest
	GetAcceptLanguage() *string
	SetLaneId(v int64) *QuerySwimmingLaneByIdRequest
	GetLaneId() *int64
	SetNamespace(v string) *QuerySwimmingLaneByIdRequest
	GetNamespace() *string
}

type QuerySwimmingLaneByIdRequest struct {
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
	// The ID of the lane.
	//
	// This parameter is required.
	//
	// example:
	//
	// 250
	LaneId *int64 `json:"LaneId,omitempty" xml:"LaneId,omitempty"`
	// The name of the Microservices Engine (MSE) namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s QuerySwimmingLaneByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySwimmingLaneByIdRequest) GoString() string {
	return s.String()
}

func (s *QuerySwimmingLaneByIdRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QuerySwimmingLaneByIdRequest) GetLaneId() *int64 {
	return s.LaneId
}

func (s *QuerySwimmingLaneByIdRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *QuerySwimmingLaneByIdRequest) SetAcceptLanguage(v string) *QuerySwimmingLaneByIdRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QuerySwimmingLaneByIdRequest) SetLaneId(v int64) *QuerySwimmingLaneByIdRequest {
	s.LaneId = &v
	return s
}

func (s *QuerySwimmingLaneByIdRequest) SetNamespace(v string) *QuerySwimmingLaneByIdRequest {
	s.Namespace = &v
	return s
}

func (s *QuerySwimmingLaneByIdRequest) Validate() error {
	return dara.Validate(s)
}
