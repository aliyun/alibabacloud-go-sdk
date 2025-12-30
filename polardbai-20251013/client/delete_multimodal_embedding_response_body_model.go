// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultimodalEmbeddingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *DeleteMultimodalEmbeddingResponseBody
	GetData() interface{}
	SetErrCode(v string) *DeleteMultimodalEmbeddingResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteMultimodalEmbeddingResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DeleteMultimodalEmbeddingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMultimodalEmbeddingResponseBody
	GetSuccess() *bool
}

type DeleteMultimodalEmbeddingResponseBody struct {
	// example:
	//
	// embedding ID
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 错误码
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// 错误信息
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 983863E2-****-1D15-A4AE-3468CD75383D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMultimodalEmbeddingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultimodalEmbeddingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMultimodalEmbeddingResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DeleteMultimodalEmbeddingResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteMultimodalEmbeddingResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteMultimodalEmbeddingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMultimodalEmbeddingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMultimodalEmbeddingResponseBody) SetData(v interface{}) *DeleteMultimodalEmbeddingResponseBody {
	s.Data = v
	return s
}

func (s *DeleteMultimodalEmbeddingResponseBody) SetErrCode(v string) *DeleteMultimodalEmbeddingResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteMultimodalEmbeddingResponseBody) SetErrMessage(v string) *DeleteMultimodalEmbeddingResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteMultimodalEmbeddingResponseBody) SetRequestId(v string) *DeleteMultimodalEmbeddingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMultimodalEmbeddingResponseBody) SetSuccess(v bool) *DeleteMultimodalEmbeddingResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMultimodalEmbeddingResponseBody) Validate() error {
	return dara.Validate(s)
}
