// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*Label) *CreateDatasetLabelsRequest
	GetLabels() []*Label
}

type CreateDatasetLabelsRequest struct {
	// The tags.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s CreateDatasetLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetLabelsRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetLabelsRequest) GetLabels() []*Label {
	return s.Labels
}

func (s *CreateDatasetLabelsRequest) SetLabels(v []*Label) *CreateDatasetLabelsRequest {
	s.Labels = v
	return s
}

func (s *CreateDatasetLabelsRequest) Validate() error {
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
