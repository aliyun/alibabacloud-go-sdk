// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineBuildConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoutineName(v string) *GetRoutineBuildConfigurationRequest
	GetRoutineName() *string
}

type GetRoutineBuildConfigurationRequest struct {
	// The ER name.
	//
	// example:
	//
	// test-routine
	RoutineName *string `json:"RoutineName,omitempty" xml:"RoutineName,omitempty"`
}

func (s GetRoutineBuildConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineBuildConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetRoutineBuildConfigurationRequest) GetRoutineName() *string {
	return s.RoutineName
}

func (s *GetRoutineBuildConfigurationRequest) SetRoutineName(v string) *GetRoutineBuildConfigurationRequest {
	s.RoutineName = &v
	return s
}

func (s *GetRoutineBuildConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
