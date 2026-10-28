// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGadInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *DeleteGadInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteGadInstanceResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DeleteGadInstanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteGadInstanceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *DeleteGadInstanceResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *DeleteGadInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteGadInstanceResponseBody
	GetSuccess() *string
}

type DeleteGadInstanceResponseBody struct {
	// The dynamic error code. This parameter will be deprecated soon.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message. It is used to replace **%s*	- in the error message returned in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **Type**, it indicates that the **Type*	- request parameter is invalid.
	//
	// example:
	//
	// can not find env: online
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned when an error occurs during the call.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
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
	// FDC111B1-ACBF-457D-9656-247FDEE9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGadInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGadInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGadInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteGadInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteGadInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteGadInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteGadInstanceResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DeleteGadInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGadInstanceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteGadInstanceResponseBody) SetDynamicCode(v string) *DeleteGadInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteGadInstanceResponseBody) SetDynamicMessage(v string) *DeleteGadInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteGadInstanceResponseBody) SetErrCode(v string) *DeleteGadInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteGadInstanceResponseBody) SetErrMessage(v string) *DeleteGadInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteGadInstanceResponseBody) SetHttpStatusCode(v string) *DeleteGadInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGadInstanceResponseBody) SetRequestId(v string) *DeleteGadInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGadInstanceResponseBody) SetSuccess(v string) *DeleteGadInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGadInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
