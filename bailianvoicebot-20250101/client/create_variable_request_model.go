// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *CreateVariableRequest
	GetBusinessUnitId() *string
	SetDescription(v string) *CreateVariableRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateVariableRequest
	GetDisplayName() *string
	SetName(v string) *CreateVariableRequest
	GetName() *string
}

type CreateVariableRequest struct {
	// example:
	//
	// llm-baployoyopf22m2r
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// userType
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVariableRequest) GoString() string {
	return s.String()
}

func (s *CreateVariableRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *CreateVariableRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVariableRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateVariableRequest) GetName() *string {
	return s.Name
}

func (s *CreateVariableRequest) SetBusinessUnitId(v string) *CreateVariableRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *CreateVariableRequest) SetDescription(v string) *CreateVariableRequest {
	s.Description = &v
	return s
}

func (s *CreateVariableRequest) SetDisplayName(v string) *CreateVariableRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateVariableRequest) SetName(v string) *CreateVariableRequest {
	s.Name = &v
	return s
}

func (s *CreateVariableRequest) Validate() error {
	return dara.Validate(s)
}
