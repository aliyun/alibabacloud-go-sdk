// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveWhiteListSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveWhiteListSettingResponseBody
	GetCode() *string
	SetMessage(v string) *RemoveWhiteListSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveWhiteListSettingResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *RemoveWhiteListSettingResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *RemoveWhiteListSettingResponseBody
	GetSuccess() *bool
}

type RemoveWhiteListSettingResponseBody struct {
	// Return code: 200 for success, others for failure
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// Whether the response was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveWhiteListSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveWhiteListSettingResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveWhiteListSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveWhiteListSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveWhiteListSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveWhiteListSettingResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *RemoveWhiteListSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveWhiteListSettingResponseBody) SetCode(v string) *RemoveWhiteListSettingResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveWhiteListSettingResponseBody) SetMessage(v string) *RemoveWhiteListSettingResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveWhiteListSettingResponseBody) SetRequestId(v string) *RemoveWhiteListSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveWhiteListSettingResponseBody) SetResultObject(v bool) *RemoveWhiteListSettingResponseBody {
	s.ResultObject = &v
	return s
}

func (s *RemoveWhiteListSettingResponseBody) SetSuccess(v bool) *RemoveWhiteListSettingResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveWhiteListSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
