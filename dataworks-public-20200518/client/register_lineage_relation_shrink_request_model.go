// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterLineageRelationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLineageRelationRegisterVOShrink(v string) *RegisterLineageRelationShrinkRequest
	GetLineageRelationRegisterVOShrink() *string
}

type RegisterLineageRelationShrinkRequest struct {
	// The structure whose lineage you want to register to DataWorks.
	//
	// This parameter is required.
	LineageRelationRegisterVOShrink *string `json:"LineageRelationRegisterVO,omitempty" xml:"LineageRelationRegisterVO,omitempty"`
}

func (s RegisterLineageRelationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterLineageRelationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RegisterLineageRelationShrinkRequest) GetLineageRelationRegisterVOShrink() *string {
	return s.LineageRelationRegisterVOShrink
}

func (s *RegisterLineageRelationShrinkRequest) SetLineageRelationRegisterVOShrink(v string) *RegisterLineageRelationShrinkRequest {
	s.LineageRelationRegisterVOShrink = &v
	return s
}

func (s *RegisterLineageRelationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
