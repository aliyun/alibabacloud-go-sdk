// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSuspEventUserSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveSuspEventUserSettingResponseBody
	GetCode() *string
	SetMessage(v string) *SaveSuspEventUserSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveSuspEventUserSettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveSuspEventUserSettingResponseBody
	GetSuccess() *bool
}

type SaveSuspEventUserSettingResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// AE6229A0-BDBE-534C-A3F8-095EBXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveSuspEventUserSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSuspEventUserSettingResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSuspEventUserSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveSuspEventUserSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveSuspEventUserSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSuspEventUserSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveSuspEventUserSettingResponseBody) SetCode(v string) *SaveSuspEventUserSettingResponseBody {
	s.Code = &v
	return s
}

func (s *SaveSuspEventUserSettingResponseBody) SetMessage(v string) *SaveSuspEventUserSettingResponseBody {
	s.Message = &v
	return s
}

func (s *SaveSuspEventUserSettingResponseBody) SetRequestId(v string) *SaveSuspEventUserSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSuspEventUserSettingResponseBody) SetSuccess(v bool) *SaveSuspEventUserSettingResponseBody {
	s.Success = &v
	return s
}

func (s *SaveSuspEventUserSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
