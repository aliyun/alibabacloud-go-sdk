// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*UpdateInstanceLabelsRequestLabels) *UpdateInstanceLabelsRequest
	GetLabels() []*UpdateInstanceLabelsRequestLabels
}

type UpdateInstanceLabelsRequest struct {
	// The tags that you want to update.
	//
	// This parameter is required.
	Labels []*UpdateInstanceLabelsRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s UpdateInstanceLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceLabelsRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceLabelsRequest) GetLabels() []*UpdateInstanceLabelsRequestLabels {
	return s.Labels
}

func (s *UpdateInstanceLabelsRequest) SetLabels(v []*UpdateInstanceLabelsRequestLabels) *UpdateInstanceLabelsRequest {
	s.Labels = v
	return s
}

func (s *UpdateInstanceLabelsRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceLabelsRequestLabels struct {
	// The key of the custom tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// customLabelKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the custom tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// labelValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateInstanceLabelsRequestLabels) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceLabelsRequestLabels) GoString() string {
	return s.String()
}

func (s *UpdateInstanceLabelsRequestLabels) GetKey() *string {
	return s.Key
}

func (s *UpdateInstanceLabelsRequestLabels) GetValue() *string {
	return s.Value
}

func (s *UpdateInstanceLabelsRequestLabels) SetKey(v string) *UpdateInstanceLabelsRequestLabels {
	s.Key = &v
	return s
}

func (s *UpdateInstanceLabelsRequestLabels) SetValue(v string) *UpdateInstanceLabelsRequestLabels {
	s.Value = &v
	return s
}

func (s *UpdateInstanceLabelsRequestLabels) Validate() error {
	return dara.Validate(s)
}
