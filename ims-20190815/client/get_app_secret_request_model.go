// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAppSecretRequest
	GetAppId() *string
	SetAppSecretId(v string) *GetAppSecretRequest
	GetAppSecretId() *string
}

type GetAppSecretRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the application secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2efd5004-005c-4f05-83c6-5b1dd176****
	AppSecretId *string `json:"AppSecretId,omitempty" xml:"AppSecretId,omitempty"`
}

func (s GetAppSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppSecretRequest) GoString() string {
	return s.String()
}

func (s *GetAppSecretRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetAppSecretRequest) GetAppSecretId() *string {
	return s.AppSecretId
}

func (s *GetAppSecretRequest) SetAppId(v string) *GetAppSecretRequest {
	s.AppId = &v
	return s
}

func (s *GetAppSecretRequest) SetAppSecretId(v string) *GetAppSecretRequest {
	s.AppSecretId = &v
	return s
}

func (s *GetAppSecretRequest) Validate() error {
	return dara.Validate(s)
}
