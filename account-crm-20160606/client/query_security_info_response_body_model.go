// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecurityInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountSecurityInfoDto(v *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) *QuerySecurityInfoResponseBody
	GetAccountSecurityInfoDto() *QuerySecurityInfoResponseBodyAccountSecurityInfoDto
	SetCode(v string) *QuerySecurityInfoResponseBody
	GetCode() *string
	SetMessage(v string) *QuerySecurityInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySecurityInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySecurityInfoResponseBody
	GetSuccess() *bool
}

type QuerySecurityInfoResponseBody struct {
	AccountSecurityInfoDto *QuerySecurityInfoResponseBodyAccountSecurityInfoDto `json:"AccountSecurityInfoDto,omitempty" xml:"AccountSecurityInfoDto,omitempty" type:"Struct"`
	Code                   *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message                *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId              *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySecurityInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySecurityInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySecurityInfoResponseBody) GetAccountSecurityInfoDto() *QuerySecurityInfoResponseBodyAccountSecurityInfoDto {
	return s.AccountSecurityInfoDto
}

func (s *QuerySecurityInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySecurityInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySecurityInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySecurityInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySecurityInfoResponseBody) SetAccountSecurityInfoDto(v *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) *QuerySecurityInfoResponseBody {
	s.AccountSecurityInfoDto = v
	return s
}

func (s *QuerySecurityInfoResponseBody) SetCode(v string) *QuerySecurityInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySecurityInfoResponseBody) SetMessage(v string) *QuerySecurityInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySecurityInfoResponseBody) SetRequestId(v string) *QuerySecurityInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySecurityInfoResponseBody) SetSuccess(v bool) *QuerySecurityInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySecurityInfoResponseBody) Validate() error {
	if s.AccountSecurityInfoDto != nil {
		if err := s.AccountSecurityInfoDto.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySecurityInfoResponseBodyAccountSecurityInfoDto struct {
	AliyunId        *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NationalityCode *string `json:"NationalityCode,omitempty" xml:"NationalityCode,omitempty"`
	Pk              *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	ProfileType     *string `json:"ProfileType,omitempty" xml:"ProfileType,omitempty"`
	SecurityEmail   *string `json:"SecurityEmail,omitempty" xml:"SecurityEmail,omitempty"`
	SecurityMobile  *string `json:"SecurityMobile,omitempty" xml:"SecurityMobile,omitempty"`
}

func (s QuerySecurityInfoResponseBodyAccountSecurityInfoDto) String() string {
	return dara.Prettify(s)
}

func (s QuerySecurityInfoResponseBodyAccountSecurityInfoDto) GoString() string {
	return s.String()
}

func (s *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) GetAliyunId() *string {
	return s.AliyunId
}

func (s *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) GetName() *string {
	return s.Name
}

func (s *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) GetNationalityCode() *string {
	return s.NationalityCode
}

func (s *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) GetPk() *string {
	return s.Pk
}

func (s *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) GetProfileType() *string {
	return s.ProfileType
}

func (s *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) GetSecurityEmail() *string {
	return s.SecurityEmail
}

func (s *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) GetSecurityMobile() *string {
	return s.SecurityMobile
}

func (s *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) SetAliyunId(v string) *QuerySecurityInfoResponseBodyAccountSecurityInfoDto {
	s.AliyunId = &v
	return s
}

func (s *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) SetName(v string) *QuerySecurityInfoResponseBodyAccountSecurityInfoDto {
	s.Name = &v
	return s
}

func (s *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) SetNationalityCode(v string) *QuerySecurityInfoResponseBodyAccountSecurityInfoDto {
	s.NationalityCode = &v
	return s
}

func (s *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) SetPk(v string) *QuerySecurityInfoResponseBodyAccountSecurityInfoDto {
	s.Pk = &v
	return s
}

func (s *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) SetProfileType(v string) *QuerySecurityInfoResponseBodyAccountSecurityInfoDto {
	s.ProfileType = &v
	return s
}

func (s *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) SetSecurityEmail(v string) *QuerySecurityInfoResponseBodyAccountSecurityInfoDto {
	s.SecurityEmail = &v
	return s
}

func (s *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) SetSecurityMobile(v string) *QuerySecurityInfoResponseBodyAccountSecurityInfoDto {
	s.SecurityMobile = &v
	return s
}

func (s *QuerySecurityInfoResponseBodyAccountSecurityInfoDto) Validate() error {
	return dara.Validate(s)
}
