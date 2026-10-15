// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgRelationDto(v *CreateAgAccountResponseBodyAgRelationDto) *CreateAgAccountResponseBody
	GetAgRelationDto() *CreateAgAccountResponseBodyAgRelationDto
	SetCode(v string) *CreateAgAccountResponseBody
	GetCode() *string
	SetMessage(v string) *CreateAgAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAgAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAgAccountResponseBody
	GetSuccess() *bool
}

type CreateAgAccountResponseBody struct {
	// The level 2 relationship between the master account (MPK) and the member account (PK).
	AgRelationDto *CreateAgAccountResponseBodyAgRelationDto `json:"AgRelationDto,omitempty" xml:"AgRelationDto,omitempty" type:"Struct"`
	// The status code.
	//
	// example:
	//
	// LOGIN_EMAIL_HAS_BEEN_USED
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// loginEmail=685741089H@chinaunicom.cn,has used
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EAE08A27-386C-579E-966D-8853EC3C5D0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAgAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgAccountResponseBody) GetAgRelationDto() *CreateAgAccountResponseBodyAgRelationDto {
	return s.AgRelationDto
}

func (s *CreateAgAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAgAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAgAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAgAccountResponseBody) SetAgRelationDto(v *CreateAgAccountResponseBodyAgRelationDto) *CreateAgAccountResponseBody {
	s.AgRelationDto = v
	return s
}

func (s *CreateAgAccountResponseBody) SetCode(v string) *CreateAgAccountResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAgAccountResponseBody) SetMessage(v string) *CreateAgAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAgAccountResponseBody) SetRequestId(v string) *CreateAgAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgAccountResponseBody) SetSuccess(v bool) *CreateAgAccountResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAgAccountResponseBody) Validate() error {
	if s.AgRelationDto != nil {
		if err := s.AgRelationDto.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgAccountResponseBodyAgRelationDto struct {
	// The UID of the master account (MPK).
	//
	// example:
	//
	// 1785287436011964
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// The UID of the member account (PK) that is managed by the master account (MPK).
	//
	// example:
	//
	// 1728240534507590
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// The name of the RAM role.
	//
	// example:
	//
	// admin-role
	RamAdminRoleName *string `json:"RamAdminRoleName,omitempty" xml:"RamAdminRoleName,omitempty"`
	// The type of the account linking relationship.
	//
	// example:
	//
	// FINACE_CLOUD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAgAccountResponseBodyAgRelationDto) String() string {
	return dara.Prettify(s)
}

func (s CreateAgAccountResponseBodyAgRelationDto) GoString() string {
	return s.String()
}

func (s *CreateAgAccountResponseBodyAgRelationDto) GetMpk() *string {
	return s.Mpk
}

func (s *CreateAgAccountResponseBodyAgRelationDto) GetPk() *string {
	return s.Pk
}

func (s *CreateAgAccountResponseBodyAgRelationDto) GetRamAdminRoleName() *string {
	return s.RamAdminRoleName
}

func (s *CreateAgAccountResponseBodyAgRelationDto) GetType() *string {
	return s.Type
}

func (s *CreateAgAccountResponseBodyAgRelationDto) SetMpk(v string) *CreateAgAccountResponseBodyAgRelationDto {
	s.Mpk = &v
	return s
}

func (s *CreateAgAccountResponseBodyAgRelationDto) SetPk(v string) *CreateAgAccountResponseBodyAgRelationDto {
	s.Pk = &v
	return s
}

func (s *CreateAgAccountResponseBodyAgRelationDto) SetRamAdminRoleName(v string) *CreateAgAccountResponseBodyAgRelationDto {
	s.RamAdminRoleName = &v
	return s
}

func (s *CreateAgAccountResponseBodyAgRelationDto) SetType(v string) *CreateAgAccountResponseBodyAgRelationDto {
	s.Type = &v
	return s
}

func (s *CreateAgAccountResponseBodyAgRelationDto) Validate() error {
	return dara.Validate(s)
}
