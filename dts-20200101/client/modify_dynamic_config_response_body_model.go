// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDynamicConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyDynamicConfigResponseBody
	GetCode() *string
	SetDynamicMessage(v string) *ModifyDynamicConfigResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *ModifyDynamicConfigResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyDynamicConfigResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyDynamicConfigResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyDynamicConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDynamicConfigResponseBody
	GetSuccess() *bool
}

type ModifyDynamicConfigResponseBody struct {
	// The error code. This parameter will be removed in the future.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
	//
	// >  If the value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- value is invalid.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
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
	// 8D81829D-1BBD-5CE8-BE75-1CAD5750****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDynamicConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDynamicConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDynamicConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyDynamicConfigResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyDynamicConfigResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyDynamicConfigResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyDynamicConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyDynamicConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDynamicConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDynamicConfigResponseBody) SetCode(v string) *ModifyDynamicConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDynamicConfigResponseBody) SetDynamicMessage(v string) *ModifyDynamicConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyDynamicConfigResponseBody) SetErrCode(v string) *ModifyDynamicConfigResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDynamicConfigResponseBody) SetErrMessage(v string) *ModifyDynamicConfigResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDynamicConfigResponseBody) SetHttpStatusCode(v int32) *ModifyDynamicConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDynamicConfigResponseBody) SetRequestId(v string) *ModifyDynamicConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDynamicConfigResponseBody) SetSuccess(v bool) *ModifyDynamicConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDynamicConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
