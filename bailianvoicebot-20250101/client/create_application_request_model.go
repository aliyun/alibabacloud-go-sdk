// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *CreateApplicationRequest
	GetBusinessUnitId() *string
	SetConcurrency(v int32) *CreateApplicationRequest
	GetConcurrency() *int32
	SetDescription(v string) *CreateApplicationRequest
	GetDescription() *string
	SetName(v string) *CreateApplicationRequest
	GetName() *string
	SetNluAccessType(v string) *CreateApplicationRequest
	GetNluAccessType() *string
}

type CreateApplicationRequest struct {
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
	// example:
	//
	// MANAGED
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *CreateApplicationRequest) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *CreateApplicationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationRequest) GetName() *string {
	return s.Name
}

func (s *CreateApplicationRequest) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *CreateApplicationRequest) SetBusinessUnitId(v string) *CreateApplicationRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *CreateApplicationRequest) SetConcurrency(v int32) *CreateApplicationRequest {
	s.Concurrency = &v
	return s
}

func (s *CreateApplicationRequest) SetDescription(v string) *CreateApplicationRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationRequest) SetName(v string) *CreateApplicationRequest {
	s.Name = &v
	return s
}

func (s *CreateApplicationRequest) SetNluAccessType(v string) *CreateApplicationRequest {
	s.NluAccessType = &v
	return s
}

func (s *CreateApplicationRequest) Validate() error {
	return dara.Validate(s)
}
