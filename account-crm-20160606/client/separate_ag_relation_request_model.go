// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSeparateAgRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMpk(v string) *SeparateAgRelationRequest
	GetMpk() *string
	SetPk(v string) *SeparateAgRelationRequest
	GetPk() *string
}

type SeparateAgRelationRequest struct {
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s SeparateAgRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s SeparateAgRelationRequest) GoString() string {
	return s.String()
}

func (s *SeparateAgRelationRequest) GetMpk() *string {
	return s.Mpk
}

func (s *SeparateAgRelationRequest) GetPk() *string {
	return s.Pk
}

func (s *SeparateAgRelationRequest) SetMpk(v string) *SeparateAgRelationRequest {
	s.Mpk = &v
	return s
}

func (s *SeparateAgRelationRequest) SetPk(v string) *SeparateAgRelationRequest {
	s.Pk = &v
	return s
}

func (s *SeparateAgRelationRequest) Validate() error {
	return dara.Validate(s)
}
