// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSpeed(v string) *ListPipelinesRequest
	GetSpeed() *string
}

type ListPipelinesRequest struct {
	// The type of the MPS queue.
	//
	// Valid values:
	//
	// 	- Boost: MPS queue with transcoding speed boosted.
	//
	// 	- Standard: standard MPS queue.
	//
	// 	- NarrowBandHDV2: MPS queue that supports Narrowband HD 2.0.
	//
	// example:
	//
	// Standard
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
}

func (s ListPipelinesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesRequest) GoString() string {
	return s.String()
}

func (s *ListPipelinesRequest) GetSpeed() *string {
	return s.Speed
}

func (s *ListPipelinesRequest) SetSpeed(v string) *ListPipelinesRequest {
	s.Speed = &v
	return s
}

func (s *ListPipelinesRequest) Validate() error {
	return dara.Validate(s)
}
