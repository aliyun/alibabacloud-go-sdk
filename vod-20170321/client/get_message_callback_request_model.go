// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetMessageCallbackRequest
	GetAppId() *string
	SetOwnerAccount(v string) *GetMessageCallbackRequest
	GetOwnerAccount() *string
}

type GetMessageCallbackRequest struct {
	// The application ID. If this parameter is not specified, the ID of the system default application is used. The fixed value is **app-1000000**.
	//
	// example:
	//
	// app-1000000
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s GetMessageCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCallbackRequest) GoString() string {
	return s.String()
}

func (s *GetMessageCallbackRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetMessageCallbackRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetMessageCallbackRequest) SetAppId(v string) *GetMessageCallbackRequest {
	s.AppId = &v
	return s
}

func (s *GetMessageCallbackRequest) SetOwnerAccount(v string) *GetMessageCallbackRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetMessageCallbackRequest) Validate() error {
	return dara.Validate(s)
}
