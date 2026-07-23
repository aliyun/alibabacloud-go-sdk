// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineBuildConfigurationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoutineNames(v string) *ListRoutineBuildConfigurationsRequest
	GetRoutineNames() *string
}

type ListRoutineBuildConfigurationsRequest struct {
	// The list of ER routine names, separated by commas.
	//
	// This parameter is required.
	//
	// example:
	//
	// rwa-test,demo
	RoutineNames *string `json:"RoutineNames,omitempty" xml:"RoutineNames,omitempty"`
}

func (s ListRoutineBuildConfigurationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineBuildConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListRoutineBuildConfigurationsRequest) GetRoutineNames() *string {
	return s.RoutineNames
}

func (s *ListRoutineBuildConfigurationsRequest) SetRoutineNames(v string) *ListRoutineBuildConfigurationsRequest {
	s.RoutineNames = &v
	return s
}

func (s *ListRoutineBuildConfigurationsRequest) Validate() error {
	return dara.Validate(s)
}
