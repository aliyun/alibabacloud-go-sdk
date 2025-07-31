// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteComputeSourceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteComputeSourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteComputeSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteComputeSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteComputeSourceResponseBody
	GetSuccess() *bool
}

type DeleteComputeSourceResponseBody struct {
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteComputeSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComputeSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteComputeSourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteComputeSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteComputeSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteComputeSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteComputeSourceResponseBody) SetCode(v string) *DeleteComputeSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteComputeSourceResponseBody) SetHttpStatusCode(v int32) *DeleteComputeSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteComputeSourceResponseBody) SetMessage(v string) *DeleteComputeSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteComputeSourceResponseBody) SetRequestId(v string) *DeleteComputeSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteComputeSourceResponseBody) SetSuccess(v bool) *DeleteComputeSourceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteComputeSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
