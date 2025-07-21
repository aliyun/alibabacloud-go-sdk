// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTerminalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *AddTerminalRequest
	GetAlias() *string
	SetClientType(v string) *AddTerminalRequest
	GetClientType() *string
	SetMainBizType(v string) *AddTerminalRequest
	GetMainBizType() *string
	SetSerialNumber(v string) *AddTerminalRequest
	GetSerialNumber() *string
	SetTerminalGroupId(v string) *AddTerminalRequest
	GetTerminalGroupId() *string
	SetUuid(v string) *AddTerminalRequest
	GetUuid() *string
}

type AddTerminalRequest struct {
	Alias           *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	ClientType      *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	MainBizType     *string `json:"MainBizType,omitempty" xml:"MainBizType,omitempty"`
	SerialNumber    *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	TerminalGroupId *string `json:"TerminalGroupId,omitempty" xml:"TerminalGroupId,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s AddTerminalRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTerminalRequest) GoString() string {
	return s.String()
}

func (s *AddTerminalRequest) GetAlias() *string {
	return s.Alias
}

func (s *AddTerminalRequest) GetClientType() *string {
	return s.ClientType
}

func (s *AddTerminalRequest) GetMainBizType() *string {
	return s.MainBizType
}

func (s *AddTerminalRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *AddTerminalRequest) GetTerminalGroupId() *string {
	return s.TerminalGroupId
}

func (s *AddTerminalRequest) GetUuid() *string {
	return s.Uuid
}

func (s *AddTerminalRequest) SetAlias(v string) *AddTerminalRequest {
	s.Alias = &v
	return s
}

func (s *AddTerminalRequest) SetClientType(v string) *AddTerminalRequest {
	s.ClientType = &v
	return s
}

func (s *AddTerminalRequest) SetMainBizType(v string) *AddTerminalRequest {
	s.MainBizType = &v
	return s
}

func (s *AddTerminalRequest) SetSerialNumber(v string) *AddTerminalRequest {
	s.SerialNumber = &v
	return s
}

func (s *AddTerminalRequest) SetTerminalGroupId(v string) *AddTerminalRequest {
	s.TerminalGroupId = &v
	return s
}

func (s *AddTerminalRequest) SetUuid(v string) *AddTerminalRequest {
	s.Uuid = &v
	return s
}

func (s *AddTerminalRequest) Validate() error {
	return dara.Validate(s)
}
