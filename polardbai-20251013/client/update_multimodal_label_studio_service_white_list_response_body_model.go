// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultimodalLabelStudioServiceWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *UpdateMultimodalLabelStudioServiceWhiteListResponseBody
	GetData() interface{}
	SetErrCode(v string) *UpdateMultimodalLabelStudioServiceWhiteListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdateMultimodalLabelStudioServiceWhiteListResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *UpdateMultimodalLabelStudioServiceWhiteListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMultimodalLabelStudioServiceWhiteListResponseBody
	GetSuccess() *bool
}

type UpdateMultimodalLabelStudioServiceWhiteListResponseBody struct {
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

func (s UpdateMultimodalLabelStudioServiceWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultimodalLabelStudioServiceWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponseBody) GetData() interface{} {
	return s.Data
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponseBody) SetData(v interface{}) *UpdateMultimodalLabelStudioServiceWhiteListResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponseBody) SetErrCode(v string) *UpdateMultimodalLabelStudioServiceWhiteListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponseBody) SetErrMessage(v string) *UpdateMultimodalLabelStudioServiceWhiteListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponseBody) SetRequestId(v string) *UpdateMultimodalLabelStudioServiceWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponseBody) SetSuccess(v bool) *UpdateMultimodalLabelStudioServiceWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
