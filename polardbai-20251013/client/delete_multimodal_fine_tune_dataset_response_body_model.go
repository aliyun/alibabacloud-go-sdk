// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultimodalFineTuneDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *DeleteMultimodalFineTuneDatasetResponseBody
	GetData() interface{}
	SetErrCode(v string) *DeleteMultimodalFineTuneDatasetResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteMultimodalFineTuneDatasetResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DeleteMultimodalFineTuneDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMultimodalFineTuneDatasetResponseBody
	GetSuccess() *bool
}

type DeleteMultimodalFineTuneDatasetResponseBody struct {
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

func (s DeleteMultimodalFineTuneDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultimodalFineTuneDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMultimodalFineTuneDatasetResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DeleteMultimodalFineTuneDatasetResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteMultimodalFineTuneDatasetResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteMultimodalFineTuneDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMultimodalFineTuneDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMultimodalFineTuneDatasetResponseBody) SetData(v interface{}) *DeleteMultimodalFineTuneDatasetResponseBody {
	s.Data = v
	return s
}

func (s *DeleteMultimodalFineTuneDatasetResponseBody) SetErrCode(v string) *DeleteMultimodalFineTuneDatasetResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteMultimodalFineTuneDatasetResponseBody) SetErrMessage(v string) *DeleteMultimodalFineTuneDatasetResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteMultimodalFineTuneDatasetResponseBody) SetRequestId(v string) *DeleteMultimodalFineTuneDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMultimodalFineTuneDatasetResponseBody) SetSuccess(v bool) *DeleteMultimodalFineTuneDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMultimodalFineTuneDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
