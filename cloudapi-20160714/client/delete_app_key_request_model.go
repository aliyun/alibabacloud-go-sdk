// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAppKeyRequest
	GetAppId() *string
	SetAppKey(v string) *DeleteAppKeyRequest
	GetAppKey() *string
}

type DeleteAppKeyRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 110840611
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The AppKey of the application. The AppKey is used for calling an API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 203708622
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s DeleteAppKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppKeyRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAppKeyRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *DeleteAppKeyRequest) SetAppId(v string) *DeleteAppKeyRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppKeyRequest) SetAppKey(v string) *DeleteAppKeyRequest {
	s.AppKey = &v
	return s
}

func (s *DeleteAppKeyRequest) Validate() error {
	return dara.Validate(s)
}
