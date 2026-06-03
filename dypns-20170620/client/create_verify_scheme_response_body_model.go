// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVerifySchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVerifySchemeResponseBody
	GetCode() *string
	SetGateVerifySchemeDTO(v *CreateVerifySchemeResponseBodyGateVerifySchemeDTO) *CreateVerifySchemeResponseBody
	GetGateVerifySchemeDTO() *CreateVerifySchemeResponseBodyGateVerifySchemeDTO
	SetHttpStatusCode(v int64) *CreateVerifySchemeResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *CreateVerifySchemeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVerifySchemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateVerifySchemeResponseBody
	GetSuccess() *bool
}

type CreateVerifySchemeResponseBody struct {
	// example:
	//
	// OK
	Code                *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	GateVerifySchemeDTO *CreateVerifySchemeResponseBodyGateVerifySchemeDTO `json:"GateVerifySchemeDTO,omitempty" xml:"GateVerifySchemeDTO,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2E7CA885-8D97-C497-A02C-2D31214D3285
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateVerifySchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVerifySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVerifySchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVerifySchemeResponseBody) GetGateVerifySchemeDTO() *CreateVerifySchemeResponseBodyGateVerifySchemeDTO {
	return s.GateVerifySchemeDTO
}

func (s *CreateVerifySchemeResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *CreateVerifySchemeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVerifySchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVerifySchemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateVerifySchemeResponseBody) SetCode(v string) *CreateVerifySchemeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVerifySchemeResponseBody) SetGateVerifySchemeDTO(v *CreateVerifySchemeResponseBodyGateVerifySchemeDTO) *CreateVerifySchemeResponseBody {
	s.GateVerifySchemeDTO = v
	return s
}

func (s *CreateVerifySchemeResponseBody) SetHttpStatusCode(v int64) *CreateVerifySchemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateVerifySchemeResponseBody) SetMessage(v string) *CreateVerifySchemeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVerifySchemeResponseBody) SetRequestId(v string) *CreateVerifySchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVerifySchemeResponseBody) SetSuccess(v bool) *CreateVerifySchemeResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVerifySchemeResponseBody) Validate() error {
	if s.GateVerifySchemeDTO != nil {
		if err := s.GateVerifySchemeDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVerifySchemeResponseBodyGateVerifySchemeDTO struct {
	// example:
	//
	// FC2**********008
	SchemeCode *string `json:"SchemeCode,omitempty" xml:"SchemeCode,omitempty"`
}

func (s CreateVerifySchemeResponseBodyGateVerifySchemeDTO) String() string {
	return dara.Prettify(s)
}

func (s CreateVerifySchemeResponseBodyGateVerifySchemeDTO) GoString() string {
	return s.String()
}

func (s *CreateVerifySchemeResponseBodyGateVerifySchemeDTO) GetSchemeCode() *string {
	return s.SchemeCode
}

func (s *CreateVerifySchemeResponseBodyGateVerifySchemeDTO) SetSchemeCode(v string) *CreateVerifySchemeResponseBodyGateVerifySchemeDTO {
	s.SchemeCode = &v
	return s
}

func (s *CreateVerifySchemeResponseBodyGateVerifySchemeDTO) Validate() error {
	return dara.Validate(s)
}
