// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientUserStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateClientUserStatusRequest
	GetId() *string
	SetStatus(v string) *UpdateClientUserStatusRequest
	GetStatus() *string
}

type UpdateClientUserStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1495
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateClientUserStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientUserStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateClientUserStatusRequest) GetId() *string {
	return s.Id
}

func (s *UpdateClientUserStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateClientUserStatusRequest) SetId(v string) *UpdateClientUserStatusRequest {
	s.Id = &v
	return s
}

func (s *UpdateClientUserStatusRequest) SetStatus(v string) *UpdateClientUserStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateClientUserStatusRequest) Validate() error {
	return dara.Validate(s)
}
