// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSwimmingLaneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteSwimmingLaneRequest
	GetAcceptLanguage() *string
	SetLaneId(v int64) *DeleteSwimmingLaneRequest
	GetLaneId() *int64
	SetNamespace(v string) *DeleteSwimmingLaneRequest
	GetNamespace() *string
}

type DeleteSwimmingLaneRequest struct {
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
	// 229
	LaneId *int64 `json:"LaneId,omitempty" xml:"LaneId,omitempty"`
	// The name of the Microservices Engine (MSE) namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DeleteSwimmingLaneRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSwimmingLaneRequest) GoString() string {
	return s.String()
}

func (s *DeleteSwimmingLaneRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteSwimmingLaneRequest) GetLaneId() *int64 {
	return s.LaneId
}

func (s *DeleteSwimmingLaneRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteSwimmingLaneRequest) SetAcceptLanguage(v string) *DeleteSwimmingLaneRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteSwimmingLaneRequest) SetLaneId(v int64) *DeleteSwimmingLaneRequest {
	s.LaneId = &v
	return s
}

func (s *DeleteSwimmingLaneRequest) SetNamespace(v string) *DeleteSwimmingLaneRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteSwimmingLaneRequest) Validate() error {
	return dara.Validate(s)
}
