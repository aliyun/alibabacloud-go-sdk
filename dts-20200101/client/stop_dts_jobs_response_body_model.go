// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDtsJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *StopDtsJobsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *StopDtsJobsResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *StopDtsJobsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *StopDtsJobsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *StopDtsJobsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *StopDtsJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopDtsJobsResponseBody
	GetSuccess() *bool
}

type StopDtsJobsResponseBody struct {
	// The dynamic error code. This parameter will be removed in the future.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the **%s*	- variable in the **ErrMessage*	- parameter.
	//
	// >  If the returned value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the returned value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
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
	// AD823BD3-1BA6-4117-A536-165CB280****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopDtsJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDtsJobsResponseBody) GoString() string {
	return s.String()
}

func (s *StopDtsJobsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *StopDtsJobsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *StopDtsJobsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *StopDtsJobsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *StopDtsJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StopDtsJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDtsJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopDtsJobsResponseBody) SetDynamicCode(v string) *StopDtsJobsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *StopDtsJobsResponseBody) SetDynamicMessage(v string) *StopDtsJobsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *StopDtsJobsResponseBody) SetErrCode(v string) *StopDtsJobsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StopDtsJobsResponseBody) SetErrMessage(v string) *StopDtsJobsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StopDtsJobsResponseBody) SetHttpStatusCode(v int32) *StopDtsJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopDtsJobsResponseBody) SetRequestId(v string) *StopDtsJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDtsJobsResponseBody) SetSuccess(v bool) *StopDtsJobsResponseBody {
	s.Success = &v
	return s
}

func (s *StopDtsJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
