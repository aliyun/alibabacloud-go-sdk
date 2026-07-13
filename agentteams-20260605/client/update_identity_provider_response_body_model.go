// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateIdentityProviderResponseBody
	GetCode() *string
	SetData(v *UpdateIdentityProviderResponseBodyData) *UpdateIdentityProviderResponseBody
	GetData() *UpdateIdentityProviderResponseBodyData
	SetHttpStatusCode(v int32) *UpdateIdentityProviderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateIdentityProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateIdentityProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateIdentityProviderResponseBody
	GetSuccess() *bool
}

type UpdateIdentityProviderResponseBody struct {
	Code           *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *UpdateIdentityProviderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateIdentityProviderResponseBody) GetData() *UpdateIdentityProviderResponseBodyData {
	return s.Data
}

func (s *UpdateIdentityProviderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateIdentityProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIdentityProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateIdentityProviderResponseBody) SetCode(v string) *UpdateIdentityProviderResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateIdentityProviderResponseBody) SetData(v *UpdateIdentityProviderResponseBodyData) *UpdateIdentityProviderResponseBody {
	s.Data = v
	return s
}

func (s *UpdateIdentityProviderResponseBody) SetHttpStatusCode(v int32) *UpdateIdentityProviderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateIdentityProviderResponseBody) SetMessage(v string) *UpdateIdentityProviderResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateIdentityProviderResponseBody) SetRequestId(v string) *UpdateIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIdentityProviderResponseBody) SetSuccess(v bool) *UpdateIdentityProviderResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateIdentityProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIdentityProviderResponseBodyData struct {
	BindingId            *int64  `json:"BindingId,omitempty" xml:"BindingId,omitempty"`
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LoginEnabled         *bool   `json:"LoginEnabled,omitempty" xml:"LoginEnabled,omitempty"`
	SyncEnabled          *bool   `json:"SyncEnabled,omitempty" xml:"SyncEnabled,omitempty"`
}

func (s UpdateIdentityProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderResponseBodyData) GetBindingId() *int64 {
	return s.BindingId
}

func (s *UpdateIdentityProviderResponseBodyData) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *UpdateIdentityProviderResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateIdentityProviderResponseBodyData) GetLoginEnabled() *bool {
	return s.LoginEnabled
}

func (s *UpdateIdentityProviderResponseBodyData) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *UpdateIdentityProviderResponseBodyData) SetBindingId(v int64) *UpdateIdentityProviderResponseBodyData {
	s.BindingId = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyData) SetIdentityProviderType(v string) *UpdateIdentityProviderResponseBodyData {
	s.IdentityProviderType = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyData) SetInstanceId(v string) *UpdateIdentityProviderResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyData) SetLoginEnabled(v bool) *UpdateIdentityProviderResponseBodyData {
	s.LoginEnabled = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyData) SetSyncEnabled(v bool) *UpdateIdentityProviderResponseBodyData {
	s.SyncEnabled = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
