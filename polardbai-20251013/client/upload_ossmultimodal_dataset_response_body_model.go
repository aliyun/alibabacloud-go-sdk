// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadOSSMultimodalDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *UploadOSSMultimodalDatasetResponseBody
	GetData() interface{}
	SetErrCode(v string) *UploadOSSMultimodalDatasetResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UploadOSSMultimodalDatasetResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *UploadOSSMultimodalDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadOSSMultimodalDatasetResponseBody
	GetSuccess() *bool
}

type UploadOSSMultimodalDatasetResponseBody struct {
	// example:
	//
	// 数据集列表，每一项包括数据集id，数据集名，创建时间，修改时间，标签，描述，导入信息，embedding列表相关字段
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

func (s UploadOSSMultimodalDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadOSSMultimodalDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *UploadOSSMultimodalDatasetResponseBody) GetData() interface{} {
	return s.Data
}

func (s *UploadOSSMultimodalDatasetResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UploadOSSMultimodalDatasetResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UploadOSSMultimodalDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadOSSMultimodalDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadOSSMultimodalDatasetResponseBody) SetData(v interface{}) *UploadOSSMultimodalDatasetResponseBody {
	s.Data = v
	return s
}

func (s *UploadOSSMultimodalDatasetResponseBody) SetErrCode(v string) *UploadOSSMultimodalDatasetResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UploadOSSMultimodalDatasetResponseBody) SetErrMessage(v string) *UploadOSSMultimodalDatasetResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UploadOSSMultimodalDatasetResponseBody) SetRequestId(v string) *UploadOSSMultimodalDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadOSSMultimodalDatasetResponseBody) SetSuccess(v bool) *UploadOSSMultimodalDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *UploadOSSMultimodalDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
