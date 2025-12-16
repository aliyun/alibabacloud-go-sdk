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
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. No endpoint is created. The system checks whether your AccessKey is valid, whether Resource Access Management (RAM) users are authorized, and whether the required parameters are set.
	//
	// 	- false (default): creates an endpoint immediately.
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
