// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalEmbeddingModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ListMultimodalEmbeddingModelResponseBody
	GetData() interface{}
	SetErrCode(v string) *ListMultimodalEmbeddingModelResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListMultimodalEmbeddingModelResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ListMultimodalEmbeddingModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMultimodalEmbeddingModelResponseBody
	GetSuccess() *bool
}

type ListMultimodalEmbeddingModelResponseBody struct {
	// example:
	//
	// 模型相关信息
	//
	// 返回示例：
	//
	// {
	//
	//   "Data": {
	//
	//     "Items": [
	//
	//       {
	//
	//         "Model": "MultiModal-Embedding"
	//
	//       }
	//
	//     ]
	//
	//   },
	//
	//   "RequestId": "7A53B3C8-2228-1478-9D12-B4C219369540",
	//
	//   "Success": true
	//
	// }
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

func (s ListMultimodalEmbeddingModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalEmbeddingModelResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultimodalEmbeddingModelResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ListMultimodalEmbeddingModelResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListMultimodalEmbeddingModelResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListMultimodalEmbeddingModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMultimodalEmbeddingModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMultimodalEmbeddingModelResponseBody) SetData(v interface{}) *ListMultimodalEmbeddingModelResponseBody {
	s.Data = v
	return s
}

func (s *ListMultimodalEmbeddingModelResponseBody) SetErrCode(v string) *ListMultimodalEmbeddingModelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListMultimodalEmbeddingModelResponseBody) SetErrMessage(v string) *ListMultimodalEmbeddingModelResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListMultimodalEmbeddingModelResponseBody) SetRequestId(v string) *ListMultimodalEmbeddingModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultimodalEmbeddingModelResponseBody) SetSuccess(v bool) *ListMultimodalEmbeddingModelResponseBody {
	s.Success = &v
	return s
}

func (s *ListMultimodalEmbeddingModelResponseBody) Validate() error {
	return dara.Validate(s)
}
