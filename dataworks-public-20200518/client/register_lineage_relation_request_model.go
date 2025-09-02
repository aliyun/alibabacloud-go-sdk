// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterLineageRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLineageRelationRegisterVO(v *LineageRelationRegisterVO) *RegisterLineageRelationRequest
	GetLineageRelationRegisterVO() *LineageRelationRegisterVO
}

type RegisterLineageRelationRequest struct {
	// The structure whose lineage you want to register to DataWorks.
	//
	// This parameter is required.
	LineageRelationRegisterVO *LineageRelationRegisterVO `json:"LineageRelationRegisterVO,omitempty" xml:"LineageRelationRegisterVO,omitempty"`
}

func (s RegisterLineageRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterLineageRelationRequest) GoString() string {
	return s.String()
}

func (s *RegisterLineageRelationRequest) GetLineageRelationRegisterVO() *LineageRelationRegisterVO {
	return s.LineageRelationRegisterVO
}

func (s *RegisterLineageRelationRequest) SetLineageRelationRegisterVO(v *LineageRelationRegisterVO) *RegisterLineageRelationRequest {
	s.LineageRelationRegisterVO = v
	return s
}

func (s *RegisterLineageRelationRequest) Validate() error {
	return dara.Validate(s)
}
