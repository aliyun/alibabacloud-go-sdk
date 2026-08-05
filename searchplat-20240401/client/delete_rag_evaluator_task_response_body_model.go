// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRagEvaluatorTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteRagEvaluatorTaskResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteRagEvaluatorTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteRagEvaluatorTaskResponseBody
	GetRequestId() *string
}

type DeleteRagEvaluatorTaskResponseBody struct {
	// The status code.
	//
	// example:
	//
	// not found
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message.
	//
	// example:
	//
	// "xx not found"
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0B66A850-506C-56B7-B001-EA09411CCD69
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteRagEvaluatorTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRagEvaluatorTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRagEvaluatorTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRagEvaluatorTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteRagEvaluatorTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRagEvaluatorTaskResponseBody) SetCode(v string) *DeleteRagEvaluatorTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRagEvaluatorTaskResponseBody) SetMessage(v string) *DeleteRagEvaluatorTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRagEvaluatorTaskResponseBody) SetRequestId(v string) *DeleteRagEvaluatorTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRagEvaluatorTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
