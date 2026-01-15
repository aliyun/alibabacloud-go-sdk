// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalEmbeddingModelModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ListMultimodalEmbeddingModelModeResponseBody
	GetData() interface{}
	SetErrCode(v string) *ListMultimodalEmbeddingModelModeResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListMultimodalEmbeddingModelModeResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ListMultimodalEmbeddingModelModeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMultimodalEmbeddingModelModeResponseBody
	GetSuccess() *bool
}

type ListMultimodalEmbeddingModelModeResponseBody struct {
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

func (s ListMultimodalEmbeddingModelModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalEmbeddingModelModeResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultimodalEmbeddingModelModeResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ListMultimodalEmbeddingModelModeResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListMultimodalEmbeddingModelModeResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListMultimodalEmbeddingModelModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMultimodalEmbeddingModelModeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMultimodalEmbeddingModelModeResponseBody) SetData(v interface{}) *ListMultimodalEmbeddingModelModeResponseBody {
	s.Data = v
	return s
}

func (s *ListMultimodalEmbeddingModelModeResponseBody) SetErrCode(v string) *ListMultimodalEmbeddingModelModeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListMultimodalEmbeddingModelModeResponseBody) SetErrMessage(v string) *ListMultimodalEmbeddingModelModeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListMultimodalEmbeddingModelModeResponseBody) SetRequestId(v string) *ListMultimodalEmbeddingModelModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultimodalEmbeddingModelModeResponseBody) SetSuccess(v bool) *ListMultimodalEmbeddingModelModeResponseBody {
	s.Success = &v
	return s
}

func (s *ListMultimodalEmbeddingModelModeResponseBody) Validate() error {
	return dara.Validate(s)
}
