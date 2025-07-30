// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyDtsJobNameResponseBody
	GetCode() *string
	SetDynamicMessage(v string) *ModifyDtsJobNameResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *ModifyDtsJobNameResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyDtsJobNameResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyDtsJobNameResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyDtsJobNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDtsJobNameResponseBody
	GetSuccess() *bool
}

type ModifyDtsJobNameResponseBody struct {
	// The error code. This parameter is going to be removed in the future.
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
	// The error message returned if the request failed.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code returned.
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
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDtsJobNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyDtsJobNameResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyDtsJobNameResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyDtsJobNameResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyDtsJobNameResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyDtsJobNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDtsJobNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDtsJobNameResponseBody) SetCode(v string) *ModifyDtsJobNameResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDtsJobNameResponseBody) SetDynamicMessage(v string) *ModifyDtsJobNameResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyDtsJobNameResponseBody) SetErrCode(v string) *ModifyDtsJobNameResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDtsJobNameResponseBody) SetErrMessage(v string) *ModifyDtsJobNameResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDtsJobNameResponseBody) SetHttpStatusCode(v int32) *ModifyDtsJobNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDtsJobNameResponseBody) SetRequestId(v string) *ModifyDtsJobNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDtsJobNameResponseBody) SetSuccess(v bool) *ModifyDtsJobNameResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDtsJobNameResponseBody) Validate() error {
	return dara.Validate(s)
}
