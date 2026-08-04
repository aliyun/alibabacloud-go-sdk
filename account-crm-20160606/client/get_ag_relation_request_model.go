// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgAccountType(v string) *GetAgRelationRequest
	GetAgAccountType() *string
	SetPk(v string) *GetAgRelationRequest
	GetPk() *string
}

type GetAgRelationRequest struct {
	AgAccountType *string `json:"AgAccountType,omitempty" xml:"AgAccountType,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s GetAgRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgRelationRequest) GoString() string {
	return s.String()
}

func (s *GetAgRelationRequest) GetAgAccountType() *string {
	return s.AgAccountType
}

func (s *GetAgRelationRequest) GetPk() *string {
	return s.Pk
}

func (s *GetAgRelationRequest) SetAgAccountType(v string) *GetAgRelationRequest {
	s.AgAccountType = &v
	return s
}

func (s *GetAgRelationRequest) SetPk(v string) *GetAgRelationRequest {
	s.Pk = &v
	return s
}

func (s *GetAgRelationRequest) Validate() error {
	return dara.Validate(s)
}
