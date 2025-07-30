// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDtsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *ResetDtsJobResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ResetDtsJobResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *ResetDtsJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ResetDtsJobResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ResetDtsJobResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ResetDtsJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResetDtsJobResponseBody
	GetSuccess() *bool
}

type ResetDtsJobResponseBody struct {
	// The dynamic error code. This parameter will be removed in the future.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace **%s*	- in the **ErrMessage*	- parameter.
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

func (s ResetDtsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetDtsJobResponseBody) GoString() string {
	return s.String()
}

func (s *ResetDtsJobResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ResetDtsJobResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ResetDtsJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ResetDtsJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ResetDtsJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ResetDtsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetDtsJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResetDtsJobResponseBody) SetDynamicCode(v string) *ResetDtsJobResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ResetDtsJobResponseBody) SetDynamicMessage(v string) *ResetDtsJobResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ResetDtsJobResponseBody) SetErrCode(v string) *ResetDtsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ResetDtsJobResponseBody) SetErrMessage(v string) *ResetDtsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ResetDtsJobResponseBody) SetHttpStatusCode(v int32) *ResetDtsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResetDtsJobResponseBody) SetRequestId(v string) *ResetDtsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetDtsJobResponseBody) SetSuccess(v bool) *ResetDtsJobResponseBody {
	s.Success = &v
	return s
}

func (s *ResetDtsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
