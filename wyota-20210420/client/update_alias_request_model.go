// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *UpdateAliasRequest
	GetAlias() *string
	SetSerialNo(v string) *UpdateAliasRequest
	GetSerialNo() *string
	SetUuid(v string) *UpdateAliasRequest
	GetUuid() *string
}

type UpdateAliasRequest struct {
	Alias    *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s UpdateAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAliasRequest) GoString() string {
	return s.String()
}

func (s *UpdateAliasRequest) GetAlias() *string {
	return s.Alias
}

func (s *UpdateAliasRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *UpdateAliasRequest) GetUuid() *string {
	return s.Uuid
}

func (s *UpdateAliasRequest) SetAlias(v string) *UpdateAliasRequest {
	s.Alias = &v
	return s
}

func (s *UpdateAliasRequest) SetSerialNo(v string) *UpdateAliasRequest {
	s.SerialNo = &v
	return s
}

func (s *UpdateAliasRequest) SetUuid(v string) *UpdateAliasRequest {
	s.Uuid = &v
	return s
}

func (s *UpdateAliasRequest) Validate() error {
	return dara.Validate(s)
}
