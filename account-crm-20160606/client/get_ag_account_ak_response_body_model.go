// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgAccountAkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountAkDto(v *GetAgAccountAkResponseBodyAccountAkDto) *GetAgAccountAkResponseBody
	GetAccountAkDto() *GetAgAccountAkResponseBodyAccountAkDto
	SetCode(v string) *GetAgAccountAkResponseBody
	GetCode() *string
	SetMessage(v string) *GetAgAccountAkResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgAccountAkResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetAgAccountAkResponseBody
	GetSuccess() *string
}

type GetAgAccountAkResponseBody struct {
	AccountAkDto *GetAgAccountAkResponseBodyAccountAkDto `json:"AccountAkDto,omitempty" xml:"AccountAkDto,omitempty" type:"Struct"`
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *string                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgAccountAkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgAccountAkResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgAccountAkResponseBody) GetAccountAkDto() *GetAgAccountAkResponseBodyAccountAkDto {
	return s.AccountAkDto
}

func (s *GetAgAccountAkResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgAccountAkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgAccountAkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgAccountAkResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetAgAccountAkResponseBody) SetAccountAkDto(v *GetAgAccountAkResponseBodyAccountAkDto) *GetAgAccountAkResponseBody {
	s.AccountAkDto = v
	return s
}

func (s *GetAgAccountAkResponseBody) SetCode(v string) *GetAgAccountAkResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgAccountAkResponseBody) SetMessage(v string) *GetAgAccountAkResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgAccountAkResponseBody) SetRequestId(v string) *GetAgAccountAkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgAccountAkResponseBody) SetSuccess(v string) *GetAgAccountAkResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgAccountAkResponseBody) Validate() error {
	if s.AccountAkDto != nil {
		if err := s.AccountAkDto.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgAccountAkResponseBodyAccountAkDto struct {
	Ak     *string `json:"Ak,omitempty" xml:"Ak,omitempty"`
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
}

func (s GetAgAccountAkResponseBodyAccountAkDto) String() string {
	return dara.Prettify(s)
}

func (s GetAgAccountAkResponseBodyAccountAkDto) GoString() string {
	return s.String()
}

func (s *GetAgAccountAkResponseBodyAccountAkDto) GetAk() *string {
	return s.Ak
}

func (s *GetAgAccountAkResponseBodyAccountAkDto) GetSecret() *string {
	return s.Secret
}

func (s *GetAgAccountAkResponseBodyAccountAkDto) SetAk(v string) *GetAgAccountAkResponseBodyAccountAkDto {
	s.Ak = &v
	return s
}

func (s *GetAgAccountAkResponseBodyAccountAkDto) SetSecret(v string) *GetAgAccountAkResponseBodyAccountAkDto {
	s.Secret = &v
	return s
}

func (s *GetAgAccountAkResponseBodyAccountAkDto) Validate() error {
	return dara.Validate(s)
}
