// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetVersionLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*Label) *CreateDatasetVersionLabelsRequest
	GetLabels() []*Label
}

type CreateDatasetVersionLabelsRequest struct {
	// The tags.
	//
	// This parameter is required.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s CreateDatasetVersionLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetVersionLabelsRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionLabelsRequest) GetLabels() []*Label {
	return s.Labels
}

func (s *CreateDatasetVersionLabelsRequest) SetLabels(v []*Label) *CreateDatasetVersionLabelsRequest {
	s.Labels = v
	return s
}

func (s *CreateDatasetVersionLabelsRequest) Validate() error {
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
