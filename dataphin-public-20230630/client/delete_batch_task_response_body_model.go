// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBatchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteBatchTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteBatchTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteBatchTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteBatchTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBatchTaskResponseBody
	GetSuccess() *bool
}

type DeleteBatchTaskResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBatchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBatchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBatchTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteBatchTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteBatchTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteBatchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBatchTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBatchTaskResponseBody) SetCode(v string) *DeleteBatchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBatchTaskResponseBody) SetHttpStatusCode(v int32) *DeleteBatchTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteBatchTaskResponseBody) SetMessage(v string) *DeleteBatchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBatchTaskResponseBody) SetRequestId(v string) *DeleteBatchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBatchTaskResponseBody) SetSuccess(v bool) *DeleteBatchTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBatchTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
