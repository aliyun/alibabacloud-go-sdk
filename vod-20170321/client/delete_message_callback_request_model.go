// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteMessageCallbackRequest
	GetAppId() *string
	SetOwnerAccount(v string) *DeleteMessageCallbackRequest
	GetOwnerAccount() *string
}

type DeleteMessageCallbackRequest struct {
	// The ID of the application. If you do not set this parameter, the default value **app-1000000*	- is used.
	//
	// example:
	//
	// app-1000000
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DeleteMessageCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageCallbackRequest) GoString() string {
	return s.String()
}

func (s *DeleteMessageCallbackRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMessageCallbackRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteMessageCallbackRequest) SetAppId(v string) *DeleteMessageCallbackRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMessageCallbackRequest) SetOwnerAccount(v string) *DeleteMessageCallbackRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteMessageCallbackRequest) Validate() error {
	return dara.Validate(s)
}
