// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateClientSettingsResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateClientSettingsResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateClientSettingsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateClientSettingsResponseBody
	GetSuccess() *bool
}

type UpdateClientSettingsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateClientSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClientSettingsResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateClientSettingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateClientSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateClientSettingsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateClientSettingsResponseBody) SetCode(v string) *UpdateClientSettingsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateClientSettingsResponseBody) SetMessage(v string) *UpdateClientSettingsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateClientSettingsResponseBody) SetRequestId(v string) *UpdateClientSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClientSettingsResponseBody) SetSuccess(v bool) *UpdateClientSettingsResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateClientSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
