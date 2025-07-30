// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDtsJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *StartDtsJobsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *StartDtsJobsResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *StartDtsJobsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *StartDtsJobsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *StartDtsJobsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *StartDtsJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartDtsJobsResponseBody
	GetSuccess() *bool
}

type StartDtsJobsResponseBody struct {
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

func (s StartDtsJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDtsJobsResponseBody) GoString() string {
	return s.String()
}

func (s *StartDtsJobsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *StartDtsJobsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *StartDtsJobsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *StartDtsJobsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *StartDtsJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartDtsJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDtsJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartDtsJobsResponseBody) SetDynamicCode(v string) *StartDtsJobsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *StartDtsJobsResponseBody) SetDynamicMessage(v string) *StartDtsJobsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *StartDtsJobsResponseBody) SetErrCode(v string) *StartDtsJobsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartDtsJobsResponseBody) SetErrMessage(v string) *StartDtsJobsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartDtsJobsResponseBody) SetHttpStatusCode(v int32) *StartDtsJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartDtsJobsResponseBody) SetRequestId(v string) *StartDtsJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDtsJobsResponseBody) SetSuccess(v bool) *StartDtsJobsResponseBody {
	s.Success = &v
	return s
}

func (s *StartDtsJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
