// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateABTestExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ABTestExperiment) *CreateABTestExperimentRequest
	GetBody() *ABTestExperiment
	SetDryRun(v bool) *CreateABTestExperimentRequest
	GetDryRun() *bool
}

type CreateABTestExperimentRequest struct {
	// The request body.
	Body *ABTestExperiment `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to perform a dry run. The default value is false. Valid values:
	//
	// Valid values:
	//
	// - **true**: The system checks the validity of the request parameters.
	//
	// - **false**: The system checks the validity of the request parameters and creates the experiment.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateABTestExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateABTestExperimentRequest) GoString() string {
	return s.String()
}

func (s *CreateABTestExperimentRequest) GetBody() *ABTestExperiment {
	return s.Body
}

func (s *CreateABTestExperimentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateABTestExperimentRequest) SetBody(v *ABTestExperiment) *CreateABTestExperimentRequest {
	s.Body = v
	return s
}

func (s *CreateABTestExperimentRequest) SetDryRun(v bool) *CreateABTestExperimentRequest {
	s.DryRun = &v
	return s
}

func (s *CreateABTestExperimentRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
