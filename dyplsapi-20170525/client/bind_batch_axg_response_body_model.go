// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindBatchAxgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BindBatchAxgResponseBody
	GetCode() *string
	SetMessage(v string) *BindBatchAxgResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindBatchAxgResponseBody
	GetRequestId() *string
	SetSecretBindList(v *BindBatchAxgResponseBodySecretBindList) *BindBatchAxgResponseBody
	GetSecretBindList() *BindBatchAxgResponseBodySecretBindList
}

type BindBatchAxgResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5DCCA8CD-6C0A-50B4-A496-B1D2AB48A1C3
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretBindList *BindBatchAxgResponseBodySecretBindList `json:"SecretBindList,omitempty" xml:"SecretBindList,omitempty" type:"Struct"`
}

func (s BindBatchAxgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindBatchAxgResponseBody) GoString() string {
	return s.String()
}

func (s *BindBatchAxgResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindBatchAxgResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindBatchAxgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindBatchAxgResponseBody) GetSecretBindList() *BindBatchAxgResponseBodySecretBindList {
	return s.SecretBindList
}

func (s *BindBatchAxgResponseBody) SetCode(v string) *BindBatchAxgResponseBody {
	s.Code = &v
	return s
}

func (s *BindBatchAxgResponseBody) SetMessage(v string) *BindBatchAxgResponseBody {
	s.Message = &v
	return s
}

func (s *BindBatchAxgResponseBody) SetRequestId(v string) *BindBatchAxgResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindBatchAxgResponseBody) SetSecretBindList(v *BindBatchAxgResponseBodySecretBindList) *BindBatchAxgResponseBody {
	s.SecretBindList = v
	return s
}

func (s *BindBatchAxgResponseBody) Validate() error {
	if s.SecretBindList != nil {
		if err := s.SecretBindList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindBatchAxgResponseBodySecretBindList struct {
	SecretBind []*BindBatchAxgResponseBodySecretBindListSecretBind `json:"SecretBind,omitempty" xml:"SecretBind,omitempty" type:"Repeated"`
}

func (s BindBatchAxgResponseBodySecretBindList) String() string {
	return dara.Prettify(s)
}

func (s BindBatchAxgResponseBodySecretBindList) GoString() string {
	return s.String()
}

func (s *BindBatchAxgResponseBodySecretBindList) GetSecretBind() []*BindBatchAxgResponseBodySecretBindListSecretBind {
	return s.SecretBind
}

func (s *BindBatchAxgResponseBodySecretBindList) SetSecretBind(v []*BindBatchAxgResponseBodySecretBindListSecretBind) *BindBatchAxgResponseBodySecretBindList {
	s.SecretBind = v
	return s
}

func (s *BindBatchAxgResponseBodySecretBindList) Validate() error {
	if s.SecretBind != nil {
		for _, item := range s.SecretBind {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BindBatchAxgResponseBodySecretBindListSecretBind struct {
	// example:
	//
	// isv.INVALID_PARAMETERS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 257
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// 1234
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// ringConfig invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 13333333333
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// example:
	//
	// 13333333333
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// example:
	//
	// 1000085060515673
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s BindBatchAxgResponseBodySecretBindListSecretBind) String() string {
	return dara.Prettify(s)
}

func (s BindBatchAxgResponseBodySecretBindListSecretBind) GoString() string {
	return s.String()
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) GetCode() *string {
	return s.Code
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) GetExtension() *string {
	return s.Extension
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) GetGroupId() *string {
	return s.GroupId
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) GetMessage() *string {
	return s.Message
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) GetPhoneNoA() *string {
	return s.PhoneNoA
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) GetSecretNo() *string {
	return s.SecretNo
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) GetSubsId() *string {
	return s.SubsId
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) SetCode(v string) *BindBatchAxgResponseBodySecretBindListSecretBind {
	s.Code = &v
	return s
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) SetExtension(v string) *BindBatchAxgResponseBodySecretBindListSecretBind {
	s.Extension = &v
	return s
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) SetGroupId(v string) *BindBatchAxgResponseBodySecretBindListSecretBind {
	s.GroupId = &v
	return s
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) SetMessage(v string) *BindBatchAxgResponseBodySecretBindListSecretBind {
	s.Message = &v
	return s
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) SetPhoneNoA(v string) *BindBatchAxgResponseBodySecretBindListSecretBind {
	s.PhoneNoA = &v
	return s
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) SetSecretNo(v string) *BindBatchAxgResponseBodySecretBindListSecretBind {
	s.SecretNo = &v
	return s
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) SetSubsId(v string) *BindBatchAxgResponseBodySecretBindListSecretBind {
	s.SubsId = &v
	return s
}

func (s *BindBatchAxgResponseBodySecretBindListSecretBind) Validate() error {
	return dara.Validate(s)
}
