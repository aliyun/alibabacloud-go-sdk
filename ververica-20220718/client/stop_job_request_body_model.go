// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopJobRequestBody interface {
	dara.Model
	String() string
	GoString() string
	SetStopStrategy(v string) *StopJobRequestBody
	GetStopStrategy() *string
}

type StopJobRequestBody struct {
	// This parameter is required.
	//
	// example:
	//
	// NONE
	StopStrategy *string `json:"stopStrategy,omitempty" xml:"stopStrategy,omitempty"`
}

func (s StopJobRequestBody) String() string {
	return dara.Prettify(s)
}

func (s StopJobRequestBody) GoString() string {
	return s.String()
}

func (s *StopJobRequestBody) GetStopStrategy() *string {
	return s.StopStrategy
}

func (s *StopJobRequestBody) SetStopStrategy(v string) *StopJobRequestBody {
	s.StopStrategy = &v
	return s
}

func (s *StopJobRequestBody) Validate() error {
	return dara.Validate(s)
}
