// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultimodalDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *UpdateMultimodalDatasetResponseBody
	GetData() interface{}
	SetErrCode(v string) *UpdateMultimodalDatasetResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdateMultimodalDatasetResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *UpdateMultimodalDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMultimodalDatasetResponseBody
	GetSuccess() *bool
}

type UpdateMultimodalDatasetResponseBody struct {
	// example:
	//
	// 数据集ID
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

func (s UpdateMultimodalDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultimodalDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMultimodalDatasetResponseBody) GetData() interface{} {
	return s.Data
}

func (s *UpdateMultimodalDatasetResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateMultimodalDatasetResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdateMultimodalDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMultimodalDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMultimodalDatasetResponseBody) SetData(v interface{}) *UpdateMultimodalDatasetResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMultimodalDatasetResponseBody) SetErrCode(v string) *UpdateMultimodalDatasetResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateMultimodalDatasetResponseBody) SetErrMessage(v string) *UpdateMultimodalDatasetResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateMultimodalDatasetResponseBody) SetRequestId(v string) *UpdateMultimodalDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMultimodalDatasetResponseBody) SetSuccess(v bool) *UpdateMultimodalDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMultimodalDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
