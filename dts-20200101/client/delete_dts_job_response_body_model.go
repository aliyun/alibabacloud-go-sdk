// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDtsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *DeleteDtsJobResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteDtsJobResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DeleteDtsJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteDtsJobResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DeleteDtsJobResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteDtsJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDtsJobResponseBody
	GetSuccess() *bool
}

type DeleteDtsJobResponseBody struct {
	// The dynamic error code. This parameter will be deprecated.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the **%s*	- placeholder in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
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

func (s DeleteDtsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDtsJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDtsJobResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteDtsJobResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteDtsJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteDtsJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteDtsJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDtsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDtsJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDtsJobResponseBody) SetDynamicCode(v string) *DeleteDtsJobResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetDynamicMessage(v string) *DeleteDtsJobResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetErrCode(v string) *DeleteDtsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetErrMessage(v string) *DeleteDtsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetHttpStatusCode(v int32) *DeleteDtsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetRequestId(v string) *DeleteDtsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetSuccess(v bool) *DeleteDtsJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDtsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
