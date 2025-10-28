// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*Label) *CreateModelLabelsRequest
	GetLabels() []*Label
}

type CreateModelLabelsRequest struct {
	// The tags.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s CreateModelLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelLabelsRequest) GoString() string {
	return s.String()
}

func (s *CreateModelLabelsRequest) GetLabels() []*Label {
	return s.Labels
}

func (s *CreateModelLabelsRequest) SetLabels(v []*Label) *CreateModelLabelsRequest {
	s.Labels = v
	return s
}

func (s *CreateModelLabelsRequest) Validate() error {
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
