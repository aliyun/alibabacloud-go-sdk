// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertInstanceResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConvertInstanceResourceGroupResponseBody
	GetCode() *string
	SetDynamicMessage(v string) *ConvertInstanceResourceGroupResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *ConvertInstanceResourceGroupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ConvertInstanceResourceGroupResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ConvertInstanceResourceGroupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ConvertInstanceResourceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConvertInstanceResourceGroupResponseBody
	GetSuccess() *bool
}

type ConvertInstanceResourceGroupResponseBody struct {
	// The error code returned by the backend service. The number is incremented.
	//
	// example:
	//
	// 500
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
	//
	// >  If **The Value of Input Parameter %s is not valid*	- is returned and DtsJobId is returned for DynamicMessage, it indicates that the request parameter DtsJobId is invalid.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned when the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned when the request failed.
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
	// The ID of a request.
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

func (s ConvertInstanceResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConvertInstanceResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertInstanceResourceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConvertInstanceResourceGroupResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ConvertInstanceResourceGroupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ConvertInstanceResourceGroupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ConvertInstanceResourceGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ConvertInstanceResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConvertInstanceResourceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConvertInstanceResourceGroupResponseBody) SetCode(v string) *ConvertInstanceResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ConvertInstanceResourceGroupResponseBody) SetDynamicMessage(v string) *ConvertInstanceResourceGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ConvertInstanceResourceGroupResponseBody) SetErrCode(v string) *ConvertInstanceResourceGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConvertInstanceResourceGroupResponseBody) SetErrMessage(v string) *ConvertInstanceResourceGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConvertInstanceResourceGroupResponseBody) SetHttpStatusCode(v int32) *ConvertInstanceResourceGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ConvertInstanceResourceGroupResponseBody) SetRequestId(v string) *ConvertInstanceResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertInstanceResourceGroupResponseBody) SetSuccess(v bool) *ConvertInstanceResourceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ConvertInstanceResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
