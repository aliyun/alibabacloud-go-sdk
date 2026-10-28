// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGadOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *CreateGadOrderResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateGadOrderResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *CreateGadOrderResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateGadOrderResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *CreateGadOrderResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *CreateGadOrderResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateGadOrderResponseBody
	GetSuccess() *string
	SetTaskId(v string) *CreateGadOrderResponseBody
	GetTaskId() *string
}

type CreateGadOrderResponseBody struct {
	// The dynamic error code. This parameter is scheduled to be deprecated.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the **%s*	- placeholder in the error message of the **ErrMessage*	- parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **Type**, it indicates that the value of the **Type*	- request parameter is invalid.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned when the call fails.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned when the call fails.
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
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DB0C7001-4F46-5A4E-BC6B-E608D3F***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID generated for the order.
	//
	// example:
	//
	// dtsGad-m4p109mdv8nt***
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateGadOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGadOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGadOrderResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateGadOrderResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateGadOrderResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateGadOrderResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateGadOrderResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateGadOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGadOrderResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateGadOrderResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateGadOrderResponseBody) SetDynamicCode(v string) *CreateGadOrderResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateGadOrderResponseBody) SetDynamicMessage(v string) *CreateGadOrderResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateGadOrderResponseBody) SetErrCode(v string) *CreateGadOrderResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateGadOrderResponseBody) SetErrMessage(v string) *CreateGadOrderResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateGadOrderResponseBody) SetHttpStatusCode(v string) *CreateGadOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateGadOrderResponseBody) SetRequestId(v string) *CreateGadOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGadOrderResponseBody) SetSuccess(v string) *CreateGadOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGadOrderResponseBody) SetTaskId(v string) *CreateGadOrderResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateGadOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
