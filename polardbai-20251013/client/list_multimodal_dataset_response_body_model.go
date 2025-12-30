// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ListMultimodalDatasetResponseBody
	GetData() interface{}
	SetErrCode(v string) *ListMultimodalDatasetResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListMultimodalDatasetResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ListMultimodalDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMultimodalDatasetResponseBody
	GetSuccess() *bool
}

type ListMultimodalDatasetResponseBody struct {
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

func (s ListMultimodalDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultimodalDatasetResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ListMultimodalDatasetResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListMultimodalDatasetResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListMultimodalDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMultimodalDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMultimodalDatasetResponseBody) SetData(v interface{}) *ListMultimodalDatasetResponseBody {
	s.Data = v
	return s
}

func (s *ListMultimodalDatasetResponseBody) SetErrCode(v string) *ListMultimodalDatasetResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListMultimodalDatasetResponseBody) SetErrMessage(v string) *ListMultimodalDatasetResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListMultimodalDatasetResponseBody) SetRequestId(v string) *ListMultimodalDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultimodalDatasetResponseBody) SetSuccess(v bool) *ListMultimodalDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *ListMultimodalDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
