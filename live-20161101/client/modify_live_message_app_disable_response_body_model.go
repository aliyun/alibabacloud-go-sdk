// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageAppDisableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyLiveMessageAppDisableResponseBody
	GetAppId() *string
	SetAppSign(v string) *ModifyLiveMessageAppDisableResponseBody
	GetAppSign() *string
	SetDisable(v bool) *ModifyLiveMessageAppDisableResponseBody
	GetDisable() *bool
	SetRequestId(v string) *ModifyLiveMessageAppDisableResponseBody
	GetRequestId() *string
}

type ModifyLiveMessageAppDisableResponseBody struct {
	// The ID of the interactive messaging application.
	//
	// example:
	//
	// ab6b5740****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The signature of the interactive messaging application. It is required by the interactive messaging SDK.
	//
	// example:
	//
	// H4sIAAAAAAAE/wBwAI//zguHB+lYCilkv7diSkk4GhHQAvMXs5tWyI+I09+uEBiB5sqa28ycJSJFmsd50Mhz8nDrvvqmti+fVaNLC5CMgLvNIy48v1aV9x74LRNFN0+Dxd2Al51xuDNkEIDaEwjqfyxscTXjSr0iQjHu2WgkpQAAAP//AQAA//+yR5XCc****
	AppSign *string `json:"AppSign,omitempty" xml:"AppSign,omitempty"`
	// Indicates whether the interactive messaging application is disabled.
	//
	// example:
	//
	// true
	Disable *bool `json:"Disable,omitempty" xml:"Disable,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6AA1AE11-EA78-1FD4-A966-6BA843073F6D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLiveMessageAppDisableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageAppDisableResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageAppDisableResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *ModifyLiveMessageAppDisableResponseBody) GetAppSign() *string {
	return s.AppSign
}

func (s *ModifyLiveMessageAppDisableResponseBody) GetDisable() *bool {
	return s.Disable
}

func (s *ModifyLiveMessageAppDisableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLiveMessageAppDisableResponseBody) SetAppId(v string) *ModifyLiveMessageAppDisableResponseBody {
	s.AppId = &v
	return s
}

func (s *ModifyLiveMessageAppDisableResponseBody) SetAppSign(v string) *ModifyLiveMessageAppDisableResponseBody {
	s.AppSign = &v
	return s
}

func (s *ModifyLiveMessageAppDisableResponseBody) SetDisable(v bool) *ModifyLiveMessageAppDisableResponseBody {
	s.Disable = &v
	return s
}

func (s *ModifyLiveMessageAppDisableResponseBody) SetRequestId(v string) *ModifyLiveMessageAppDisableResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLiveMessageAppDisableResponseBody) Validate() error {
	return dara.Validate(s)
}
