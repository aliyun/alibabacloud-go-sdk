// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertWhiteListSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InsertWhiteListSettingResponseBody
	GetCode() *string
	SetMessage(v string) *InsertWhiteListSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsertWhiteListSettingResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *InsertWhiteListSettingResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *InsertWhiteListSettingResponseBody
	GetSuccess() *bool
}

type InsertWhiteListSettingResponseBody struct {
	// Return code: 200 for success, others for failure.
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
	// Request ID
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// Indicates whether the response was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertWhiteListSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertWhiteListSettingResponseBody) GoString() string {
	return s.String()
}

func (s *InsertWhiteListSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *InsertWhiteListSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsertWhiteListSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertWhiteListSettingResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *InsertWhiteListSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InsertWhiteListSettingResponseBody) SetCode(v string) *InsertWhiteListSettingResponseBody {
	s.Code = &v
	return s
}

func (s *InsertWhiteListSettingResponseBody) SetMessage(v string) *InsertWhiteListSettingResponseBody {
	s.Message = &v
	return s
}

func (s *InsertWhiteListSettingResponseBody) SetRequestId(v string) *InsertWhiteListSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertWhiteListSettingResponseBody) SetResultObject(v bool) *InsertWhiteListSettingResponseBody {
	s.ResultObject = &v
	return s
}

func (s *InsertWhiteListSettingResponseBody) SetSuccess(v bool) *InsertWhiteListSettingResponseBody {
	s.Success = &v
	return s
}

func (s *InsertWhiteListSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
