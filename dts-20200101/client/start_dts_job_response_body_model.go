// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDtsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *StartDtsJobResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *StartDtsJobResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *StartDtsJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *StartDtsJobResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *StartDtsJobResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *StartDtsJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartDtsJobResponseBody
	GetSuccess() *bool
}

type StartDtsJobResponseBody struct {
	// The dynamic error code. This parameter will be removed in the future.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the **%s*	- variable in the **ErrMessage*	- parameter.
	//
	// >  For example, if the returned value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the returned value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
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
	// 601B6F25-21E7-4484-99D5-3EF2625C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartDtsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDtsJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartDtsJobResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *StartDtsJobResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *StartDtsJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *StartDtsJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *StartDtsJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartDtsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDtsJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartDtsJobResponseBody) SetDynamicCode(v string) *StartDtsJobResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *StartDtsJobResponseBody) SetDynamicMessage(v string) *StartDtsJobResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *StartDtsJobResponseBody) SetErrCode(v string) *StartDtsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartDtsJobResponseBody) SetErrMessage(v string) *StartDtsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartDtsJobResponseBody) SetHttpStatusCode(v int32) *StartDtsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartDtsJobResponseBody) SetRequestId(v string) *StartDtsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDtsJobResponseBody) SetSuccess(v bool) *StartDtsJobResponseBody {
	s.Success = &v
	return s
}

func (s *StartDtsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
