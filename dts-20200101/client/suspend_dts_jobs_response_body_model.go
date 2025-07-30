// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendDtsJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *SuspendDtsJobsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *SuspendDtsJobsResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *SuspendDtsJobsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *SuspendDtsJobsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *SuspendDtsJobsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *SuspendDtsJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SuspendDtsJobsResponseBody
	GetSuccess() *bool
}

type SuspendDtsJobsResponseBody struct {
	// The dynamic error code. This parameter will be removed in the future.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the **%s*	- variable in the **ErrMessage*	- parameter.
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
	// AD823BD3-1BA6-4117-A536-165CB280****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendDtsJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendDtsJobsResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendDtsJobsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *SuspendDtsJobsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *SuspendDtsJobsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *SuspendDtsJobsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SuspendDtsJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SuspendDtsJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendDtsJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SuspendDtsJobsResponseBody) SetDynamicCode(v string) *SuspendDtsJobsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SuspendDtsJobsResponseBody) SetDynamicMessage(v string) *SuspendDtsJobsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SuspendDtsJobsResponseBody) SetErrCode(v string) *SuspendDtsJobsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SuspendDtsJobsResponseBody) SetErrMessage(v string) *SuspendDtsJobsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SuspendDtsJobsResponseBody) SetHttpStatusCode(v int32) *SuspendDtsJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SuspendDtsJobsResponseBody) SetRequestId(v string) *SuspendDtsJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendDtsJobsResponseBody) SetSuccess(v bool) *SuspendDtsJobsResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendDtsJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
