// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetExperimentLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*LabelInfo) *SetExperimentLabelsRequest
	GetLabels() []*LabelInfo
}

type SetExperimentLabelsRequest struct {
	// The tags.
	Labels []*LabelInfo `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s SetExperimentLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s SetExperimentLabelsRequest) GoString() string {
	return s.String()
}

func (s *SetExperimentLabelsRequest) GetLabels() []*LabelInfo {
	return s.Labels
}

func (s *SetExperimentLabelsRequest) SetLabels(v []*LabelInfo) *SetExperimentLabelsRequest {
	s.Labels = v
	return s
}

func (s *SetExperimentLabelsRequest) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
