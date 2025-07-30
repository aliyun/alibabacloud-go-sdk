// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDtsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *StopDtsJobResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *StopDtsJobResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *StopDtsJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *StopDtsJobResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *StopDtsJobResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *StopDtsJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopDtsJobResponseBody
	GetSuccess() *bool
}

type StopDtsJobResponseBody struct {
	// The dynamic error code. This parameter will be removed in the future.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace **%s*	- in the **ErrMessage*	- parameter.
	//
	// >  If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 01B6F25-21E7-4484-99D5-3EF2625C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopDtsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDtsJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopDtsJobResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *StopDtsJobResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *StopDtsJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *StopDtsJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *StopDtsJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StopDtsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDtsJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopDtsJobResponseBody) SetDynamicCode(v string) *StopDtsJobResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *StopDtsJobResponseBody) SetDynamicMessage(v string) *StopDtsJobResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *StopDtsJobResponseBody) SetErrCode(v string) *StopDtsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StopDtsJobResponseBody) SetErrMessage(v string) *StopDtsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StopDtsJobResponseBody) SetHttpStatusCode(v int32) *StopDtsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopDtsJobResponseBody) SetRequestId(v string) *StopDtsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDtsJobResponseBody) SetSuccess(v bool) *StopDtsJobResponseBody {
	s.Success = &v
	return s
}

func (s *StopDtsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
