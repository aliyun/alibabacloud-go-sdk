// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultimodalDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *DeleteMultimodalDatasetResponseBody
	GetData() interface{}
	SetErrCode(v string) *DeleteMultimodalDatasetResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteMultimodalDatasetResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DeleteMultimodalDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMultimodalDatasetResponseBody
	GetSuccess() *bool
}

type DeleteMultimodalDatasetResponseBody struct {
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

func (s DeleteMultimodalDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultimodalDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMultimodalDatasetResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DeleteMultimodalDatasetResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteMultimodalDatasetResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteMultimodalDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMultimodalDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMultimodalDatasetResponseBody) SetData(v interface{}) *DeleteMultimodalDatasetResponseBody {
	s.Data = v
	return s
}

func (s *DeleteMultimodalDatasetResponseBody) SetErrCode(v string) *DeleteMultimodalDatasetResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteMultimodalDatasetResponseBody) SetErrMessage(v string) *DeleteMultimodalDatasetResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteMultimodalDatasetResponseBody) SetRequestId(v string) *DeleteMultimodalDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMultimodalDatasetResponseBody) SetSuccess(v bool) *DeleteMultimodalDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMultimodalDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
