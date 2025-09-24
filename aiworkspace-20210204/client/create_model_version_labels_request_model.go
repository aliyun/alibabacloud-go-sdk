// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelVersionLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*Label) *CreateModelVersionLabelsRequest
	GetLabels() []*Label
}

type CreateModelVersionLabelsRequest struct {
	// The tags.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s CreateModelVersionLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelVersionLabelsRequest) GoString() string {
	return s.String()
}

func (s *CreateModelVersionLabelsRequest) GetLabels() []*Label {
	return s.Labels
}

func (s *CreateModelVersionLabelsRequest) SetLabels(v []*Label) *CreateModelVersionLabelsRequest {
	s.Labels = v
	return s
}

func (s *CreateModelVersionLabelsRequest) Validate() error {
	return dara.Validate(s)
}
