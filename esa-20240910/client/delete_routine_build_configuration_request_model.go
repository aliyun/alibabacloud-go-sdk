// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineBuildConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoutineName(v string) *DeleteRoutineBuildConfigurationRequest
	GetRoutineName() *string
}

type DeleteRoutineBuildConfigurationRequest struct {
	// The ER name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-routine
	RoutineName *string `json:"RoutineName,omitempty" xml:"RoutineName,omitempty"`
}

func (s DeleteRoutineBuildConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineBuildConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoutineBuildConfigurationRequest) GetRoutineName() *string {
	return s.RoutineName
}

func (s *DeleteRoutineBuildConfigurationRequest) SetRoutineName(v string) *DeleteRoutineBuildConfigurationRequest {
	s.RoutineName = &v
	return s
}

func (s *DeleteRoutineBuildConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
