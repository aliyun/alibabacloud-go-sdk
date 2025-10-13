// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrainingJobLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*UpdateTrainingJobLabelsRequestLabels) *UpdateTrainingJobLabelsRequest
	GetLabels() []*UpdateTrainingJobLabelsRequestLabels
}

type UpdateTrainingJobLabelsRequest struct {
	Labels []*UpdateTrainingJobLabelsRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s UpdateTrainingJobLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrainingJobLabelsRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrainingJobLabelsRequest) GetLabels() []*UpdateTrainingJobLabelsRequestLabels {
	return s.Labels
}

func (s *UpdateTrainingJobLabelsRequest) SetLabels(v []*UpdateTrainingJobLabelsRequestLabels) *UpdateTrainingJobLabelsRequest {
	s.Labels = v
	return s
}

func (s *UpdateTrainingJobLabelsRequest) Validate() error {
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

type UpdateTrainingJobLabelsRequestLabels struct {
	// example:
	//
	// RootModelID
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// model-ad8cv770kl
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTrainingJobLabelsRequestLabels) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrainingJobLabelsRequestLabels) GoString() string {
	return s.String()
}

func (s *UpdateTrainingJobLabelsRequestLabels) GetKey() *string {
	return s.Key
}

func (s *UpdateTrainingJobLabelsRequestLabels) GetValue() *string {
	return s.Value
}

func (s *UpdateTrainingJobLabelsRequestLabels) SetKey(v string) *UpdateTrainingJobLabelsRequestLabels {
	s.Key = &v
	return s
}

func (s *UpdateTrainingJobLabelsRequestLabels) SetValue(v string) *UpdateTrainingJobLabelsRequestLabels {
	s.Value = &v
	return s
}

func (s *UpdateTrainingJobLabelsRequestLabels) Validate() error {
	return dara.Validate(s)
}
