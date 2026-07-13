// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BindIdentityProviderResponseBody
	GetCode() *string
	SetData(v *BindIdentityProviderResponseBodyData) *BindIdentityProviderResponseBody
	GetData() *BindIdentityProviderResponseBodyData
	SetHttpStatusCode(v int32) *BindIdentityProviderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BindIdentityProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindIdentityProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindIdentityProviderResponseBody
	GetSuccess() *bool
}

type BindIdentityProviderResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *BindIdentityProviderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *BindIdentityProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindIdentityProviderResponseBody) GetData() *BindIdentityProviderResponseBodyData {
	return s.Data
}

func (s *BindIdentityProviderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BindIdentityProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindIdentityProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindIdentityProviderResponseBody) SetCode(v string) *BindIdentityProviderResponseBody {
	s.Code = &v
	return s
}

func (s *BindIdentityProviderResponseBody) SetData(v *BindIdentityProviderResponseBodyData) *BindIdentityProviderResponseBody {
	s.Data = v
	return s
}

func (s *BindIdentityProviderResponseBody) SetHttpStatusCode(v int32) *BindIdentityProviderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BindIdentityProviderResponseBody) SetMessage(v string) *BindIdentityProviderResponseBody {
	s.Message = &v
	return s
}

func (s *BindIdentityProviderResponseBody) SetRequestId(v string) *BindIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindIdentityProviderResponseBody) SetSuccess(v bool) *BindIdentityProviderResponseBody {
	s.Success = &v
	return s
}

func (s *BindIdentityProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindIdentityProviderResponseBodyData struct {
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s BindIdentityProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindIdentityProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindIdentityProviderResponseBodyData) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *BindIdentityProviderResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BindIdentityProviderResponseBodyData) SetIdentityProviderType(v string) *BindIdentityProviderResponseBodyData {
	s.IdentityProviderType = &v
	return s
}

func (s *BindIdentityProviderResponseBodyData) SetInstanceId(v string) *BindIdentityProviderResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *BindIdentityProviderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
