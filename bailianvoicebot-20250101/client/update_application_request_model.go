// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateApplicationRequest
	GetApplicationId() *string
	SetBusinessUnitId(v string) *UpdateApplicationRequest
	GetBusinessUnitId() *string
	SetConcurrency(v int32) *UpdateApplicationRequest
	GetConcurrency() *int32
	SetDescription(v string) *UpdateApplicationRequest
	GetDescription() *string
	SetName(v string) *UpdateApplicationRequest
	GetName() *string
}

type UpdateApplicationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a395011f-a247-400f-bc69-28796749fd52
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	// example:
	//
	// 10
	Concurrency *int32  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *UpdateApplicationRequest) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *UpdateApplicationRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationRequest) GetName() *string {
	return s.Name
}

func (s *UpdateApplicationRequest) SetApplicationId(v string) *UpdateApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationRequest) SetBusinessUnitId(v string) *UpdateApplicationRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *UpdateApplicationRequest) SetConcurrency(v int32) *UpdateApplicationRequest {
	s.Concurrency = &v
	return s
}

func (s *UpdateApplicationRequest) SetDescription(v string) *UpdateApplicationRequest {
	s.Description = &v
	return s
}

func (s *UpdateApplicationRequest) SetName(v string) *UpdateApplicationRequest {
	s.Name = &v
	return s
}

func (s *UpdateApplicationRequest) Validate() error {
	return dara.Validate(s)
}
