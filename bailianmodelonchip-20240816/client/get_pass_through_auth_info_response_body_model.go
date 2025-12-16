// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPassThroughAuthInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPassThroughAuthInfoResponseBody
	GetCode() *string
	SetData(v *GetPassThroughAuthInfoResponseBodyData) *GetPassThroughAuthInfoResponseBody
	GetData() *GetPassThroughAuthInfoResponseBodyData
	SetHttpStatusCode(v int32) *GetPassThroughAuthInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPassThroughAuthInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPassThroughAuthInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPassThroughAuthInfoResponseBody
	GetSuccess() *bool
}

type GetPassThroughAuthInfoResponseBody struct {
	// example:
	//
	// success
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetPassThroughAuthInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 7B0FC4BC-9E4B-5AD7-9D35-6559BDC0788E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPassThroughAuthInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPassThroughAuthInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPassThroughAuthInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPassThroughAuthInfoResponseBody) GetData() *GetPassThroughAuthInfoResponseBodyData {
	return s.Data
}

func (s *GetPassThroughAuthInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPassThroughAuthInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPassThroughAuthInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPassThroughAuthInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPassThroughAuthInfoResponseBody) SetCode(v string) *GetPassThroughAuthInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetPassThroughAuthInfoResponseBody) SetData(v *GetPassThroughAuthInfoResponseBodyData) *GetPassThroughAuthInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetPassThroughAuthInfoResponseBody) SetHttpStatusCode(v int32) *GetPassThroughAuthInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPassThroughAuthInfoResponseBody) SetMessage(v string) *GetPassThroughAuthInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetPassThroughAuthInfoResponseBody) SetRequestId(v string) *GetPassThroughAuthInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPassThroughAuthInfoResponseBody) SetSuccess(v bool) *GetPassThroughAuthInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetPassThroughAuthInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPassThroughAuthInfoResponseBodyData struct {
	// example:
	//
	// mm_2b7d37b695fb44faa983e5fbb032
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// 35f9ed99b27a1e8374f6593bf969f7d6
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// example:
	//
	// {"device_info":"246c19a2c721724aa3e52e96a93a730ce9080b3a7680522a0d5210e0dc0a43cd","timestamp":"1765874670867"}
	PassThroughParams map[string]interface{} `json:"passThroughParams,omitempty" xml:"passThroughParams,omitempty"`
	// example:
	//
	// 8e4e0522a986cd641866c2b2453eedfa
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetPassThroughAuthInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPassThroughAuthInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPassThroughAuthInfoResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetPassThroughAuthInfoResponseBodyData) GetDeviceName() *string {
	return s.DeviceName
}

func (s *GetPassThroughAuthInfoResponseBodyData) GetPassThroughParams() map[string]interface{} {
	return s.PassThroughParams
}

func (s *GetPassThroughAuthInfoResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetPassThroughAuthInfoResponseBodyData) SetAppId(v string) *GetPassThroughAuthInfoResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetPassThroughAuthInfoResponseBodyData) SetDeviceName(v string) *GetPassThroughAuthInfoResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *GetPassThroughAuthInfoResponseBodyData) SetPassThroughParams(v map[string]interface{}) *GetPassThroughAuthInfoResponseBodyData {
	s.PassThroughParams = v
	return s
}

func (s *GetPassThroughAuthInfoResponseBodyData) SetTaskId(v string) *GetPassThroughAuthInfoResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetPassThroughAuthInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
