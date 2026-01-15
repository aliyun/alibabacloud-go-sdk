// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalSearchModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ListMultimodalSearchModelResponseBody
	GetData() interface{}
	SetErrCode(v string) *ListMultimodalSearchModelResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListMultimodalSearchModelResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ListMultimodalSearchModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMultimodalSearchModelResponseBody
	GetSuccess() *bool
}

type ListMultimodalSearchModelResponseBody struct {
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

func (s ListMultimodalSearchModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalSearchModelResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultimodalSearchModelResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ListMultimodalSearchModelResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListMultimodalSearchModelResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListMultimodalSearchModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMultimodalSearchModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMultimodalSearchModelResponseBody) SetData(v interface{}) *ListMultimodalSearchModelResponseBody {
	s.Data = v
	return s
}

func (s *ListMultimodalSearchModelResponseBody) SetErrCode(v string) *ListMultimodalSearchModelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListMultimodalSearchModelResponseBody) SetErrMessage(v string) *ListMultimodalSearchModelResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListMultimodalSearchModelResponseBody) SetRequestId(v string) *ListMultimodalSearchModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultimodalSearchModelResponseBody) SetSuccess(v bool) *ListMultimodalSearchModelResponseBody {
	s.Success = &v
	return s
}

func (s *ListMultimodalSearchModelResponseBody) Validate() error {
	return dara.Validate(s)
}
