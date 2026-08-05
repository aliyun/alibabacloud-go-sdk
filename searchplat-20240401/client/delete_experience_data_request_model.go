// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperienceDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DeleteExperienceDataRequest
	GetDryRun() *bool
}

type DeleteExperienceDataRequest struct {
	// Whether this is a dry run request
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s DeleteExperienceDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperienceDataRequest) GoString() string {
	return s.String()
}

func (s *DeleteExperienceDataRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteExperienceDataRequest) SetDryRun(v bool) *DeleteExperienceDataRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteExperienceDataRequest) Validate() error {
	return dara.Validate(s)
}
