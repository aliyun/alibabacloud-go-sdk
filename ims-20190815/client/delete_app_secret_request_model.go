// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAppSecretRequest
	GetAppId() *string
	SetAppSecretId(v string) *DeleteAppSecretRequest
	GetAppSecretId() *string
}

type DeleteAppSecretRequest struct {
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

func (s DeleteAppSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppSecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppSecretRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAppSecretRequest) GetAppSecretId() *string {
	return s.AppSecretId
}

func (s *DeleteAppSecretRequest) SetAppId(v string) *DeleteAppSecretRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppSecretRequest) SetAppSecretId(v string) *DeleteAppSecretRequest {
	s.AppSecretId = &v
	return s
}

func (s *DeleteAppSecretRequest) Validate() error {
	return dara.Validate(s)
}
