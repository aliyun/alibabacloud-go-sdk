// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABTestExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ABTestExperiment) *UpdateABTestExperimentRequest
	GetBody() *ABTestExperiment
	SetDryRun(v bool) *UpdateABTestExperimentRequest
	GetDryRun() *bool
}

type UpdateABTestExperimentRequest struct {
	// The request body. For more information, see [ABTestExperiment](https://help.aliyun.com/document_detail/173617.html).
	Body *ABTestExperiment `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - true: Performs a dry run. The system checks if your AccessKey is valid, if the Resource Access Management (RAM) user is authorized, and if all required parameters are specified. The request is not sent.
	//
	// - false (default): Sends the request to update the experiment.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateABTestExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateABTestExperimentRequest) GoString() string {
	return s.String()
}

func (s *UpdateABTestExperimentRequest) GetBody() *ABTestExperiment {
	return s.Body
}

func (s *UpdateABTestExperimentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateABTestExperimentRequest) SetBody(v *ABTestExperiment) *UpdateABTestExperimentRequest {
	s.Body = v
	return s
}

func (s *UpdateABTestExperimentRequest) SetDryRun(v bool) *UpdateABTestExperimentRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateABTestExperimentRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
