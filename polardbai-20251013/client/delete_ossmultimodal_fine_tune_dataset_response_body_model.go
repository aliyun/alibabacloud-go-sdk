// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOSSMultimodalFineTuneDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *DeleteOSSMultimodalFineTuneDatasetResponseBody
	GetData() interface{}
	SetErrCode(v string) *DeleteOSSMultimodalFineTuneDatasetResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteOSSMultimodalFineTuneDatasetResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DeleteOSSMultimodalFineTuneDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteOSSMultimodalFineTuneDatasetResponseBody
	GetSuccess() *bool
}

type DeleteOSSMultimodalFineTuneDatasetResponseBody struct {
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

func (s DeleteOSSMultimodalFineTuneDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOSSMultimodalFineTuneDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponseBody) SetData(v interface{}) *DeleteOSSMultimodalFineTuneDatasetResponseBody {
	s.Data = v
	return s
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponseBody) SetErrCode(v string) *DeleteOSSMultimodalFineTuneDatasetResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponseBody) SetErrMessage(v string) *DeleteOSSMultimodalFineTuneDatasetResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponseBody) SetRequestId(v string) *DeleteOSSMultimodalFineTuneDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponseBody) SetSuccess(v bool) *DeleteOSSMultimodalFineTuneDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
