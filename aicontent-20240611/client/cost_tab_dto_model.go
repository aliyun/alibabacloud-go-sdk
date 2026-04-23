// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostTabDTO interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *CostTabDTO
	GetKey() *string
	SetLabel(v string) *CostTabDTO
	GetLabel() *string
	SetModelTypes(v []*string) *CostTabDTO
	GetModelTypes() []*string
}

type CostTabDTO struct {
	// example:
	//
	// all
	Key        *string   `json:"key,omitempty" xml:"key,omitempty"`
	Label      *string   `json:"label,omitempty" xml:"label,omitempty"`
	ModelTypes []*string `json:"modelTypes,omitempty" xml:"modelTypes,omitempty" type:"Repeated"`
}

func (s CostTabDTO) String() string {
	return dara.Prettify(s)
}

func (s CostTabDTO) GoString() string {
	return s.String()
}

func (s *CostTabDTO) GetKey() *string {
	return s.Key
}

func (s *CostTabDTO) GetLabel() *string {
	return s.Label
}

func (s *CostTabDTO) GetModelTypes() []*string {
	return s.ModelTypes
}

func (s *CostTabDTO) SetKey(v string) *CostTabDTO {
	s.Key = &v
	return s
}

func (s *CostTabDTO) SetLabel(v string) *CostTabDTO {
	s.Label = &v
	return s
}

func (s *CostTabDTO) SetModelTypes(v []*string) *CostTabDTO {
	s.ModelTypes = v
	return s
}

func (s *CostTabDTO) Validate() error {
	return dara.Validate(s)
}
