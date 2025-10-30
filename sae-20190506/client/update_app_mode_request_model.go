// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateAppModeRequest
	GetAppId() *string
	SetAppIds(v string) *UpdateAppModeRequest
	GetAppIds() *string
	SetEnableIdle(v bool) *UpdateAppModeRequest
	GetEnableIdle() *bool
	SetNamespaceId(v string) *UpdateAppModeRequest
	GetNamespaceId() *string
}

type UpdateAppModeRequest struct {
	// The app ID.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppIds *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	// Enable Idle Mode?
	//
	// Enumeration value:
	//
	// 	- true: enables.
	//
	// 	- false: disables.
	//
	// example:
	//
	// true
	EnableIdle  *bool   `json:"EnableIdle,omitempty" xml:"EnableIdle,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s UpdateAppModeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppModeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppModeRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateAppModeRequest) GetAppIds() *string {
	return s.AppIds
}

func (s *UpdateAppModeRequest) GetEnableIdle() *bool {
	return s.EnableIdle
}

func (s *UpdateAppModeRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateAppModeRequest) SetAppId(v string) *UpdateAppModeRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAppModeRequest) SetAppIds(v string) *UpdateAppModeRequest {
	s.AppIds = &v
	return s
}

func (s *UpdateAppModeRequest) SetEnableIdle(v bool) *UpdateAppModeRequest {
	s.EnableIdle = &v
	return s
}

func (s *UpdateAppModeRequest) SetNamespaceId(v string) *UpdateAppModeRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateAppModeRequest) Validate() error {
	return dara.Validate(s)
}
