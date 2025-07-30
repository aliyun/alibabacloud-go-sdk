// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDtsJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *DeleteDtsJobsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteDtsJobsResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DeleteDtsJobsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteDtsJobsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DeleteDtsJobsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteDtsJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDtsJobsResponseBody
	GetSuccess() *bool
}

type DeleteDtsJobsResponseBody struct {
	// The dynamic error code. This parameter will be removed soon.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
	//
	// > If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
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
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDtsJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDtsJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDtsJobsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteDtsJobsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteDtsJobsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteDtsJobsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteDtsJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDtsJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDtsJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDtsJobsResponseBody) SetDynamicCode(v string) *DeleteDtsJobsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteDtsJobsResponseBody) SetDynamicMessage(v string) *DeleteDtsJobsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteDtsJobsResponseBody) SetErrCode(v string) *DeleteDtsJobsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteDtsJobsResponseBody) SetErrMessage(v string) *DeleteDtsJobsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteDtsJobsResponseBody) SetHttpStatusCode(v int32) *DeleteDtsJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDtsJobsResponseBody) SetRequestId(v string) *DeleteDtsJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDtsJobsResponseBody) SetSuccess(v bool) *DeleteDtsJobsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDtsJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
