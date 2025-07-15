// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateAppKeyRequest
	GetAppId() *string
	SetAppKey(v string) *CreateAppKeyRequest
	GetAppKey() *string
	SetAppSecret(v string) *CreateAppKeyRequest
	GetAppSecret() *string
}

type CreateAppKeyRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 111053351
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application AppKey.
	//
	// example:
	//
	// 204203237
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The application AppSecret.
	//
	// example:
	//
	// 6f0a4ad7918a4b41a57fc087d5b066d0
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
}

func (s CreateAppKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateAppKeyRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppKeyRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *CreateAppKeyRequest) GetAppSecret() *string {
	return s.AppSecret
}

func (s *CreateAppKeyRequest) SetAppId(v string) *CreateAppKeyRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppKeyRequest) SetAppKey(v string) *CreateAppKeyRequest {
	s.AppKey = &v
	return s
}

func (s *CreateAppKeyRequest) SetAppSecret(v string) *CreateAppKeyRequest {
	s.AppSecret = &v
	return s
}

func (s *CreateAppKeyRequest) Validate() error {
	return dara.Validate(s)
}
