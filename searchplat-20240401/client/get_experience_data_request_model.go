// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperienceDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *GetExperienceDataRequest
	GetDryRun() *bool
}

type GetExperienceDataRequest struct {
	// Specifies whether to validate the request parameters without performing the actual operation. Default value: false.
	//
	// Valid values:
	//
	// - **true**
	//
	// - **false**.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s GetExperienceDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExperienceDataRequest) GoString() string {
	return s.String()
}

func (s *GetExperienceDataRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *GetExperienceDataRequest) SetDryRun(v bool) *GetExperienceDataRequest {
	s.DryRun = &v
	return s
}

func (s *GetExperienceDataRequest) Validate() error {
	return dara.Validate(s)
}
