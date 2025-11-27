// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelSpecification interface {
	dara.Model
	String() string
	GoString() string
	SetMetaData(v *MetaData) *ModelSpecification
	GetMetaData() *MetaData
	SetSpec(v *Spec) *ModelSpecification
	GetSpec() *Spec
}

type ModelSpecification struct {
	// This parameter is required.
	MetaData *MetaData `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// This parameter is required.
	Spec *Spec `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s ModelSpecification) String() string {
	return dara.Prettify(s)
}

func (s ModelSpecification) GoString() string {
	return s.String()
}

func (s *ModelSpecification) GetMetaData() *MetaData {
	return s.MetaData
}

func (s *ModelSpecification) GetSpec() *Spec {
	return s.Spec
}

func (s *ModelSpecification) SetMetaData(v *MetaData) *ModelSpecification {
	s.MetaData = v
	return s
}

func (s *ModelSpecification) SetSpec(v *Spec) *ModelSpecification {
	s.Spec = v
	return s
}

func (s *ModelSpecification) Validate() error {
	if s.MetaData != nil {
		if err := s.MetaData.Validate(); err != nil {
			return err
		}
	}
	if s.Spec != nil {
		if err := s.Spec.Validate(); err != nil {
			return err
		}
	}
	return nil
}
