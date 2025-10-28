// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLayerVersionInput interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v *InputCodeLocation) *CreateLayerVersionInput
	GetCode() *InputCodeLocation
	SetCompatibleRuntime(v []*string) *CreateLayerVersionInput
	GetCompatibleRuntime() []*string
	SetDescription(v string) *CreateLayerVersionInput
	GetDescription() *string
	SetLicense(v string) *CreateLayerVersionInput
	GetLicense() *string
}

type CreateLayerVersionInput struct {
	Code              *InputCodeLocation `json:"code,omitempty" xml:"code,omitempty"`
	CompatibleRuntime []*string          `json:"compatibleRuntime" xml:"compatibleRuntime" type:"Repeated"`
	// example:
	//
	// my first layer
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// Apache
	License *string `json:"license,omitempty" xml:"license,omitempty"`
}

func (s CreateLayerVersionInput) String() string {
	return dara.Prettify(s)
}

func (s CreateLayerVersionInput) GoString() string {
	return s.String()
}

func (s *CreateLayerVersionInput) GetCode() *InputCodeLocation {
	return s.Code
}

func (s *CreateLayerVersionInput) GetCompatibleRuntime() []*string {
	return s.CompatibleRuntime
}

func (s *CreateLayerVersionInput) GetDescription() *string {
	return s.Description
}

func (s *CreateLayerVersionInput) GetLicense() *string {
	return s.License
}

func (s *CreateLayerVersionInput) SetCode(v *InputCodeLocation) *CreateLayerVersionInput {
	s.Code = v
	return s
}

func (s *CreateLayerVersionInput) SetCompatibleRuntime(v []*string) *CreateLayerVersionInput {
	s.CompatibleRuntime = v
	return s
}

func (s *CreateLayerVersionInput) SetDescription(v string) *CreateLayerVersionInput {
	s.Description = &v
	return s
}

func (s *CreateLayerVersionInput) SetLicense(v string) *CreateLayerVersionInput {
	s.License = &v
	return s
}

func (s *CreateLayerVersionInput) Validate() error {
	if s.Code != nil {
		if err := s.Code.Validate(); err != nil {
			return err
		}
	}
	return nil
}
