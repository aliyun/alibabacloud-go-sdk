// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOnlineEvalTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateOnlineEvalTaskResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateOnlineEvalTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateOnlineEvalTaskResponseBody
	GetRequestId() *string
}

type UpdateOnlineEvalTaskResponseBody struct {
	// The internal error code. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// cannot modify a stopped task
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the POP request
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOnlineEvalTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOnlineEvalTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOnlineEvalTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateOnlineEvalTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateOnlineEvalTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOnlineEvalTaskResponseBody) SetCode(v string) *UpdateOnlineEvalTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateOnlineEvalTaskResponseBody) SetMessage(v string) *UpdateOnlineEvalTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateOnlineEvalTaskResponseBody) SetRequestId(v string) *UpdateOnlineEvalTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOnlineEvalTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
