// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*UpdateServiceVersionRequestLabels) *UpdateServiceVersionRequest
	GetLabels() []*UpdateServiceVersionRequestLabels
}

type UpdateServiceVersionRequest struct {
	// This parameter is required.
	Labels []*UpdateServiceVersionRequestLabels `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
}

func (s UpdateServiceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceVersionRequest) GetLabels() []*UpdateServiceVersionRequestLabels {
	return s.Labels
}

func (s *UpdateServiceVersionRequest) SetLabels(v []*UpdateServiceVersionRequestLabels) *UpdateServiceVersionRequest {
	s.Labels = v
	return s
}

func (s *UpdateServiceVersionRequest) Validate() error {
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

type UpdateServiceVersionRequestLabels struct {
	// This parameter is required.
	//
	// example:
	//
	// topology.kubernetes.io/zone
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// cn-hangzhou-k
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateServiceVersionRequestLabels) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceVersionRequestLabels) GoString() string {
	return s.String()
}

func (s *UpdateServiceVersionRequestLabels) GetKey() *string {
	return s.Key
}

func (s *UpdateServiceVersionRequestLabels) GetValue() *string {
	return s.Value
}

func (s *UpdateServiceVersionRequestLabels) SetKey(v string) *UpdateServiceVersionRequestLabels {
	s.Key = &v
	return s
}

func (s *UpdateServiceVersionRequestLabels) SetValue(v string) *UpdateServiceVersionRequestLabels {
	s.Value = &v
	return s
}

func (s *UpdateServiceVersionRequestLabels) Validate() error {
	return dara.Validate(s)
}
