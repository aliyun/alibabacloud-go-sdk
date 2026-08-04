// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgRelationDto(v *GetAgRelationResponseBodyAgRelationDto) *GetAgRelationResponseBody
	GetAgRelationDto() *GetAgRelationResponseBodyAgRelationDto
	SetCode(v string) *GetAgRelationResponseBody
	GetCode() *string
	SetMessage(v string) *GetAgRelationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgRelationResponseBody
	GetSuccess() *bool
}

type GetAgRelationResponseBody struct {
	AgRelationDto *GetAgRelationResponseBodyAgRelationDto `json:"AgRelationDto,omitempty" xml:"AgRelationDto,omitempty" type:"Struct"`
	Code          *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgRelationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgRelationResponseBody) GetAgRelationDto() *GetAgRelationResponseBodyAgRelationDto {
	return s.AgRelationDto
}

func (s *GetAgRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgRelationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgRelationResponseBody) SetAgRelationDto(v *GetAgRelationResponseBodyAgRelationDto) *GetAgRelationResponseBody {
	s.AgRelationDto = v
	return s
}

func (s *GetAgRelationResponseBody) SetCode(v string) *GetAgRelationResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgRelationResponseBody) SetMessage(v string) *GetAgRelationResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgRelationResponseBody) SetRequestId(v string) *GetAgRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgRelationResponseBody) SetSuccess(v bool) *GetAgRelationResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgRelationResponseBody) Validate() error {
	if s.AgRelationDto != nil {
		if err := s.AgRelationDto.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgRelationResponseBodyAgRelationDto struct {
	Mpk  *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	Pk   *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAgRelationResponseBodyAgRelationDto) String() string {
	return dara.Prettify(s)
}

func (s GetAgRelationResponseBodyAgRelationDto) GoString() string {
	return s.String()
}

func (s *GetAgRelationResponseBodyAgRelationDto) GetMpk() *string {
	return s.Mpk
}

func (s *GetAgRelationResponseBodyAgRelationDto) GetPk() *string {
	return s.Pk
}

func (s *GetAgRelationResponseBodyAgRelationDto) GetType() *string {
	return s.Type
}

func (s *GetAgRelationResponseBodyAgRelationDto) SetMpk(v string) *GetAgRelationResponseBodyAgRelationDto {
	s.Mpk = &v
	return s
}

func (s *GetAgRelationResponseBodyAgRelationDto) SetPk(v string) *GetAgRelationResponseBodyAgRelationDto {
	s.Pk = &v
	return s
}

func (s *GetAgRelationResponseBodyAgRelationDto) SetType(v string) *GetAgRelationResponseBodyAgRelationDto {
	s.Type = &v
	return s
}

func (s *GetAgRelationResponseBodyAgRelationDto) Validate() error {
	return dara.Validate(s)
}
