// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyDtsJobPasswordResponseBody
	GetCode() *string
	SetDynamicMessage(v string) *ModifyDtsJobPasswordResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *ModifyDtsJobPasswordResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyDtsJobPasswordResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyDtsJobPasswordResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyDtsJobPasswordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDtsJobPasswordResponseBody
	GetSuccess() *bool
}

type ModifyDtsJobPasswordResponseBody struct {
	// The error code. This parameter will be removed in the future.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the value of **ErrMessage**.
	//
	// >  If the return value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the return value of **DynamicMessage*	- is **DtsJobId**, the specified value of **DtsJobId*	- in the request is invalid.
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
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8C498360-7892-433C-847A-BA71A850****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDtsJobPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobPasswordResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyDtsJobPasswordResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyDtsJobPasswordResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyDtsJobPasswordResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyDtsJobPasswordResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyDtsJobPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDtsJobPasswordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDtsJobPasswordResponseBody) SetCode(v string) *ModifyDtsJobPasswordResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDtsJobPasswordResponseBody) SetDynamicMessage(v string) *ModifyDtsJobPasswordResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyDtsJobPasswordResponseBody) SetErrCode(v string) *ModifyDtsJobPasswordResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDtsJobPasswordResponseBody) SetErrMessage(v string) *ModifyDtsJobPasswordResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDtsJobPasswordResponseBody) SetHttpStatusCode(v int32) *ModifyDtsJobPasswordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDtsJobPasswordResponseBody) SetRequestId(v string) *ModifyDtsJobPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDtsJobPasswordResponseBody) SetSuccess(v bool) *ModifyDtsJobPasswordResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDtsJobPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
