// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOnlineEvalTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteOnlineEvalTaskResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteOnlineEvalTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteOnlineEvalTaskResponseBody
	GetRequestId() *string
}

type DeleteOnlineEvalTaskResponseBody struct {
	// Internal error code. Set only when the response is in error.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response error message. Set only when the response is in error.
	//
	// example:
	//
	// task id is empty
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOnlineEvalTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOnlineEvalTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOnlineEvalTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteOnlineEvalTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteOnlineEvalTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOnlineEvalTaskResponseBody) SetCode(v string) *DeleteOnlineEvalTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteOnlineEvalTaskResponseBody) SetMessage(v string) *DeleteOnlineEvalTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteOnlineEvalTaskResponseBody) SetRequestId(v string) *DeleteOnlineEvalTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOnlineEvalTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
