// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendDtsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *SuspendDtsJobResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *SuspendDtsJobResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *SuspendDtsJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *SuspendDtsJobResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *SuspendDtsJobResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *SuspendDtsJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SuspendDtsJobResponseBody
	GetSuccess() *bool
}

type SuspendDtsJobResponseBody struct {
	// The dynamic error code. This parameter will be deprecated.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message used to replace the **%s*	- variable in the **ErrMessage*	- parameter.
	//
	// > For example, if the **ErrMessage*	- parameter returns **The Value of Input Parameter %s is not valid*	- and the **DynamicMessage*	- parameter returns **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
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
	// The request ID.
	//
	// example:
	//
	// 01B6F25-21E7-4484-99D5-3EF2625C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendDtsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendDtsJobResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendDtsJobResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *SuspendDtsJobResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *SuspendDtsJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *SuspendDtsJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SuspendDtsJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SuspendDtsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendDtsJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SuspendDtsJobResponseBody) SetDynamicCode(v string) *SuspendDtsJobResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SuspendDtsJobResponseBody) SetDynamicMessage(v string) *SuspendDtsJobResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SuspendDtsJobResponseBody) SetErrCode(v string) *SuspendDtsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SuspendDtsJobResponseBody) SetErrMessage(v string) *SuspendDtsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SuspendDtsJobResponseBody) SetHttpStatusCode(v int32) *SuspendDtsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SuspendDtsJobResponseBody) SetRequestId(v string) *SuspendDtsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendDtsJobResponseBody) SetSuccess(v bool) *SuspendDtsJobResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendDtsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
