// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgAccountType(v string) *UpdateAgServiceStatusRequest
	GetAgAccountType() *string
	SetMpk(v string) *UpdateAgServiceStatusRequest
	GetMpk() *string
	SetStatus(v string) *UpdateAgServiceStatusRequest
	GetStatus() *string
}

type UpdateAgServiceStatusRequest struct {
	// This parameter is required.
	AgAccountType *string `json:"AgAccountType,omitempty" xml:"AgAccountType,omitempty"`
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateAgServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgServiceStatusRequest) GetAgAccountType() *string {
	return s.AgAccountType
}

func (s *UpdateAgServiceStatusRequest) GetMpk() *string {
	return s.Mpk
}

func (s *UpdateAgServiceStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateAgServiceStatusRequest) SetAgAccountType(v string) *UpdateAgServiceStatusRequest {
	s.AgAccountType = &v
	return s
}

func (s *UpdateAgServiceStatusRequest) SetMpk(v string) *UpdateAgServiceStatusRequest {
	s.Mpk = &v
	return s
}

func (s *UpdateAgServiceStatusRequest) SetStatus(v string) *UpdateAgServiceStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateAgServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
