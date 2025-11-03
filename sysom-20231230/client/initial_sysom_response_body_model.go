// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitialSysomResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InitialSysomResponseBody
	GetRequestId() *string
	SetCode(v string) *InitialSysomResponseBody
	GetCode() *string
	SetData(v *InitialSysomResponseBodyData) *InitialSysomResponseBody
	GetData() *InitialSysomResponseBodyData
	SetMessage(v string) *InitialSysomResponseBody
	GetMessage() *string
}

type InitialSysomResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data      *InitialSysomResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                       `json:"message,omitempty" xml:"message,omitempty"`
}

func (s InitialSysomResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitialSysomResponseBody) GoString() string {
	return s.String()
}

func (s *InitialSysomResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitialSysomResponseBody) GetCode() *string {
	return s.Code
}

func (s *InitialSysomResponseBody) GetData() *InitialSysomResponseBodyData {
	return s.Data
}

func (s *InitialSysomResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitialSysomResponseBody) SetRequestId(v string) *InitialSysomResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitialSysomResponseBody) SetCode(v string) *InitialSysomResponseBody {
	s.Code = &v
	return s
}

func (s *InitialSysomResponseBody) SetData(v *InitialSysomResponseBodyData) *InitialSysomResponseBody {
	s.Data = v
	return s
}

func (s *InitialSysomResponseBody) SetMessage(v string) *InitialSysomResponseBody {
	s.Message = &v
	return s
}

func (s *InitialSysomResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitialSysomResponseBodyData struct {
	RoleExist *bool `json:"role_exist,omitempty" xml:"role_exist,omitempty"`
}

func (s InitialSysomResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InitialSysomResponseBodyData) GoString() string {
	return s.String()
}

func (s *InitialSysomResponseBodyData) GetRoleExist() *bool {
	return s.RoleExist
}

func (s *InitialSysomResponseBodyData) SetRoleExist(v bool) *InitialSysomResponseBodyData {
	s.RoleExist = &v
	return s
}

func (s *InitialSysomResponseBodyData) Validate() error {
	return dara.Validate(s)
}
