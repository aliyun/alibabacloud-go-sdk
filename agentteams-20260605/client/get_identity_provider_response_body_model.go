// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetIdentityProviderResponseBody
	GetCode() *string
	SetData(v *GetIdentityProviderResponseBodyData) *GetIdentityProviderResponseBody
	GetData() *GetIdentityProviderResponseBodyData
	SetHttpStatusCode(v int32) *GetIdentityProviderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetIdentityProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIdentityProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetIdentityProviderResponseBody
	GetSuccess() *bool
}

type GetIdentityProviderResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetIdentityProviderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIdentityProviderResponseBody) GetData() *GetIdentityProviderResponseBodyData {
	return s.Data
}

func (s *GetIdentityProviderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetIdentityProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentityProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIdentityProviderResponseBody) SetCode(v string) *GetIdentityProviderResponseBody {
	s.Code = &v
	return s
}

func (s *GetIdentityProviderResponseBody) SetData(v *GetIdentityProviderResponseBodyData) *GetIdentityProviderResponseBody {
	s.Data = v
	return s
}

func (s *GetIdentityProviderResponseBody) SetHttpStatusCode(v int32) *GetIdentityProviderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetIdentityProviderResponseBody) SetMessage(v string) *GetIdentityProviderResponseBody {
	s.Message = &v
	return s
}

func (s *GetIdentityProviderResponseBody) SetRequestId(v string) *GetIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentityProviderResponseBody) SetSuccess(v bool) *GetIdentityProviderResponseBody {
	s.Success = &v
	return s
}

func (s *GetIdentityProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderResponseBodyData struct {
	EventSubscriptionCallbackUrl *string            `json:"EventSubscriptionCallbackUrl,omitempty" xml:"EventSubscriptionCallbackUrl,omitempty"`
	IdentityProviderType         *string            `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	IdpMetadata                  map[string]*string `json:"IdpMetadata,omitempty" xml:"IdpMetadata,omitempty"`
	InstanceId                   *string            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LoginCallbackUrl             *string            `json:"LoginCallbackUrl,omitempty" xml:"LoginCallbackUrl,omitempty"`
	LoginEnabled                 *bool              `json:"LoginEnabled,omitempty" xml:"LoginEnabled,omitempty"`
	Status                       *string            `json:"Status,omitempty" xml:"Status,omitempty"`
	SyncEnabled                  *bool              `json:"SyncEnabled,omitempty" xml:"SyncEnabled,omitempty"`
}

func (s GetIdentityProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyData) GetEventSubscriptionCallbackUrl() *string {
	return s.EventSubscriptionCallbackUrl
}

func (s *GetIdentityProviderResponseBodyData) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *GetIdentityProviderResponseBodyData) GetIdpMetadata() map[string]*string {
	return s.IdpMetadata
}

func (s *GetIdentityProviderResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetIdentityProviderResponseBodyData) GetLoginCallbackUrl() *string {
	return s.LoginCallbackUrl
}

func (s *GetIdentityProviderResponseBodyData) GetLoginEnabled() *bool {
	return s.LoginEnabled
}

func (s *GetIdentityProviderResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetIdentityProviderResponseBodyData) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *GetIdentityProviderResponseBodyData) SetEventSubscriptionCallbackUrl(v string) *GetIdentityProviderResponseBodyData {
	s.EventSubscriptionCallbackUrl = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetIdentityProviderType(v string) *GetIdentityProviderResponseBodyData {
	s.IdentityProviderType = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetIdpMetadata(v map[string]*string) *GetIdentityProviderResponseBodyData {
	s.IdpMetadata = v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetInstanceId(v string) *GetIdentityProviderResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetLoginCallbackUrl(v string) *GetIdentityProviderResponseBodyData {
	s.LoginCallbackUrl = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetLoginEnabled(v bool) *GetIdentityProviderResponseBodyData {
	s.LoginEnabled = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetStatus(v string) *GetIdentityProviderResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetSyncEnabled(v bool) *GetIdentityProviderResponseBodyData {
	s.SyncEnabled = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
