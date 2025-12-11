// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*CreateServiceVersionRequestLabels) *CreateServiceVersionRequest
	GetLabels() []*CreateServiceVersionRequestLabels
	SetName(v string) *CreateServiceVersionRequest
	GetName() *string
}

type CreateServiceVersionRequest struct {
	// The service tags.
	//
	// This parameter is required.
	Labels []*CreateServiceVersionRequestLabels `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The version name.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateServiceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceVersionRequest) GetLabels() []*CreateServiceVersionRequestLabels {
	return s.Labels
}

func (s *CreateServiceVersionRequest) GetName() *string {
	return s.Name
}

func (s *CreateServiceVersionRequest) SetLabels(v []*CreateServiceVersionRequestLabels) *CreateServiceVersionRequest {
	s.Labels = v
	return s
}

func (s *CreateServiceVersionRequest) SetName(v string) *CreateServiceVersionRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceVersionRequest) Validate() error {
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

type CreateServiceVersionRequestLabels struct {
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// topology.kubernetes.io/zone
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// cn-hangzhou-j
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateServiceVersionRequestLabels) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceVersionRequestLabels) GoString() string {
	return s.String()
}

func (s *CreateServiceVersionRequestLabels) GetKey() *string {
	return s.Key
}

func (s *CreateServiceVersionRequestLabels) GetValue() *string {
	return s.Value
}

func (s *CreateServiceVersionRequestLabels) SetKey(v string) *CreateServiceVersionRequestLabels {
	s.Key = &v
	return s
}

func (s *CreateServiceVersionRequestLabels) SetValue(v string) *CreateServiceVersionRequestLabels {
	s.Value = &v
	return s
}

func (s *CreateServiceVersionRequestLabels) Validate() error {
	return dara.Validate(s)
}
