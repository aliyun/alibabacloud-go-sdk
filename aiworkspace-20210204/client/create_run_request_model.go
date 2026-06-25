// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExperimentId(v string) *CreateRunRequest
	GetExperimentId() *string
	SetLabels(v []*Label) *CreateRunRequest
	GetLabels() []*Label
	SetName(v string) *CreateRunRequest
	GetName() *string
	SetParams(v []*RunParam) *CreateRunRequest
	GetParams() []*RunParam
	SetSourceId(v string) *CreateRunRequest
	GetSourceId() *string
	SetSourceType(v string) *CreateRunRequest
	GetSourceType() *string
}

type CreateRunRequest struct {
	// The ID of the experiment associated with the run.
	//
	// This parameter is required.
	//
	// example:
	//
	// exp-6thbb5xrbmp*****
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// The list of labels for the run.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The name of the run. The naming convention is as follows:
	//
	// - Starts with a lowercase or uppercase letter.
	//
	// - Can contain lowercase letters, uppercase letters, digits, underscores (_), and hyphens (-).
	//
	// - The length must be 1 to 63 characters.
	//
	// If this parameter is left empty, the server-generated random ID (RunID) is used as the name.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of parameters for the run.
	Params []*RunParam `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The ID of the PAI workload associated with the run.
	//
	// example:
	//
	// job-jdnhf***fnrimv
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type of the PAI workload associated with the run. Options include TrainingService, DLC, or empty. This parameter is optional. The default value is empty.
	//
	// example:
	//
	// DLC
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s CreateRunRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRunRequest) GoString() string {
	return s.String()
}

func (s *CreateRunRequest) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *CreateRunRequest) GetLabels() []*Label {
	return s.Labels
}

func (s *CreateRunRequest) GetName() *string {
	return s.Name
}

func (s *CreateRunRequest) GetParams() []*RunParam {
	return s.Params
}

func (s *CreateRunRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateRunRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateRunRequest) SetExperimentId(v string) *CreateRunRequest {
	s.ExperimentId = &v
	return s
}

func (s *CreateRunRequest) SetLabels(v []*Label) *CreateRunRequest {
	s.Labels = v
	return s
}

func (s *CreateRunRequest) SetName(v string) *CreateRunRequest {
	s.Name = &v
	return s
}

func (s *CreateRunRequest) SetParams(v []*RunParam) *CreateRunRequest {
	s.Params = v
	return s
}

func (s *CreateRunRequest) SetSourceId(v string) *CreateRunRequest {
	s.SourceId = &v
	return s
}

func (s *CreateRunRequest) SetSourceType(v string) *CreateRunRequest {
	s.SourceType = &v
	return s
}

func (s *CreateRunRequest) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Params != nil {
		for _, item := range s.Params {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
