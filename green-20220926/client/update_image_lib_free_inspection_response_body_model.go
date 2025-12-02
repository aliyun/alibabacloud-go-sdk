// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageLibFreeInspectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateImageLibFreeInspectionResponseBody
	GetCode() *int32
	SetData(v bool) *UpdateImageLibFreeInspectionResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpdateImageLibFreeInspectionResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *UpdateImageLibFreeInspectionResponseBody
	GetMsg() *string
	SetRequestId(v string) *UpdateImageLibFreeInspectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateImageLibFreeInspectionResponseBody
	GetSuccess() *bool
}

type UpdateImageLibFreeInspectionResponseBody struct {
	// Error code, consistent with the HTTP status.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return result.
	//
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Further description of the error code.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// ID assigned by the backend to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success indicator.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateImageLibFreeInspectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageLibFreeInspectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageLibFreeInspectionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateImageLibFreeInspectionResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateImageLibFreeInspectionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateImageLibFreeInspectionResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *UpdateImageLibFreeInspectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateImageLibFreeInspectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateImageLibFreeInspectionResponseBody) SetCode(v int32) *UpdateImageLibFreeInspectionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateImageLibFreeInspectionResponseBody) SetData(v bool) *UpdateImageLibFreeInspectionResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateImageLibFreeInspectionResponseBody) SetHttpStatusCode(v int32) *UpdateImageLibFreeInspectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateImageLibFreeInspectionResponseBody) SetMsg(v string) *UpdateImageLibFreeInspectionResponseBody {
	s.Msg = &v
	return s
}

func (s *UpdateImageLibFreeInspectionResponseBody) SetRequestId(v string) *UpdateImageLibFreeInspectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageLibFreeInspectionResponseBody) SetSuccess(v bool) *UpdateImageLibFreeInspectionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateImageLibFreeInspectionResponseBody) Validate() error {
	return dara.Validate(s)
}
