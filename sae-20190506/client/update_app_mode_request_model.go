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
	SetEnableIdle(v bool) *UpdateAppModeRequest
	GetEnableIdle() *bool
}

type UpdateAppModeRequest struct {
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// true
	EnableIdle *bool `json:"EnableIdle,omitempty" xml:"EnableIdle,omitempty"`
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

func (s *UpdateAppModeRequest) GetEnableIdle() *bool {
	return s.EnableIdle
}

func (s *UpdateAppModeRequest) SetAppId(v string) *UpdateAppModeRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAppModeRequest) SetEnableIdle(v bool) *UpdateAppModeRequest {
	s.EnableIdle = &v
	return s
}

func (s *UpdateAppModeRequest) Validate() error {
	return dara.Validate(s)
}
