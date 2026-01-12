// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSourcePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSourcePolicyResponseBody
	GetCode() *string
	SetData(v *CreateSourcePolicyResponseBodyData) *CreateSourcePolicyResponseBody
	GetData() *CreateSourcePolicyResponseBodyData
	SetErrorName(v string) *CreateSourcePolicyResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *CreateSourcePolicyResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *CreateSourcePolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSourcePolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSourcePolicyResponseBody
	GetSuccess() *bool
}

type CreateSourcePolicyResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateSourcePolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorName *string                             `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                              `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSourcePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSourcePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSourcePolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSourcePolicyResponseBody) GetData() *CreateSourcePolicyResponseBodyData {
	return s.Data
}

func (s *CreateSourcePolicyResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *CreateSourcePolicyResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CreateSourcePolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSourcePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSourcePolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSourcePolicyResponseBody) SetCode(v string) *CreateSourcePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSourcePolicyResponseBody) SetData(v *CreateSourcePolicyResponseBodyData) *CreateSourcePolicyResponseBody {
	s.Data = v
	return s
}

func (s *CreateSourcePolicyResponseBody) SetErrorName(v string) *CreateSourcePolicyResponseBody {
	s.ErrorName = &v
	return s
}

func (s *CreateSourcePolicyResponseBody) SetHttpCode(v int32) *CreateSourcePolicyResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateSourcePolicyResponseBody) SetMessage(v string) *CreateSourcePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSourcePolicyResponseBody) SetRequestId(v string) *CreateSourcePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSourcePolicyResponseBody) SetSuccess(v bool) *CreateSourcePolicyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSourcePolicyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSourcePolicyResponseBodyData struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s CreateSourcePolicyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSourcePolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSourcePolicyResponseBodyData) GetAccessId() *string {
	return s.AccessId
}

func (s *CreateSourcePolicyResponseBodyData) GetDir() *string {
	return s.Dir
}

func (s *CreateSourcePolicyResponseBodyData) GetExpire() *string {
	return s.Expire
}

func (s *CreateSourcePolicyResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *CreateSourcePolicyResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *CreateSourcePolicyResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *CreateSourcePolicyResponseBodyData) SetAccessId(v string) *CreateSourcePolicyResponseBodyData {
	s.AccessId = &v
	return s
}

func (s *CreateSourcePolicyResponseBodyData) SetDir(v string) *CreateSourcePolicyResponseBodyData {
	s.Dir = &v
	return s
}

func (s *CreateSourcePolicyResponseBodyData) SetExpire(v string) *CreateSourcePolicyResponseBodyData {
	s.Expire = &v
	return s
}

func (s *CreateSourcePolicyResponseBodyData) SetHost(v string) *CreateSourcePolicyResponseBodyData {
	s.Host = &v
	return s
}

func (s *CreateSourcePolicyResponseBodyData) SetPolicy(v string) *CreateSourcePolicyResponseBodyData {
	s.Policy = &v
	return s
}

func (s *CreateSourcePolicyResponseBodyData) SetSignature(v string) *CreateSourcePolicyResponseBodyData {
	s.Signature = &v
	return s
}

func (s *CreateSourcePolicyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
