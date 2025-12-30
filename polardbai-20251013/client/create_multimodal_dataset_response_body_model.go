// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *CreateMultimodalDatasetResponseBody
	GetData() interface{}
	SetErrCode(v string) *CreateMultimodalDatasetResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateMultimodalDatasetResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *CreateMultimodalDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMultimodalDatasetResponseBody
	GetSuccess() *bool
}

type CreateMultimodalDatasetResponseBody struct {
	// example:
	//
	// 执行成功返回数据集ID，失败返回错误信息
	//
	// "Data":{"DatasetId":"ds-htev80p61324x8qlbn"}
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

func (s CreateMultimodalDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMultimodalDatasetResponseBody) GetData() interface{} {
	return s.Data
}

func (s *CreateMultimodalDatasetResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateMultimodalDatasetResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateMultimodalDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMultimodalDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMultimodalDatasetResponseBody) SetData(v interface{}) *CreateMultimodalDatasetResponseBody {
	s.Data = v
	return s
}

func (s *CreateMultimodalDatasetResponseBody) SetErrCode(v string) *CreateMultimodalDatasetResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateMultimodalDatasetResponseBody) SetErrMessage(v string) *CreateMultimodalDatasetResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateMultimodalDatasetResponseBody) SetRequestId(v string) *CreateMultimodalDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMultimodalDatasetResponseBody) SetSuccess(v bool) *CreateMultimodalDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMultimodalDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
