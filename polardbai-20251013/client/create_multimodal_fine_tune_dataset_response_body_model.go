// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalFineTuneDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *CreateMultimodalFineTuneDatasetResponseBody
	GetData() interface{}
	SetErrCode(v string) *CreateMultimodalFineTuneDatasetResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateMultimodalFineTuneDatasetResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *CreateMultimodalFineTuneDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMultimodalFineTuneDatasetResponseBody
	GetSuccess() *bool
}

type CreateMultimodalFineTuneDatasetResponseBody struct {
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

func (s CreateMultimodalFineTuneDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalFineTuneDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMultimodalFineTuneDatasetResponseBody) GetData() interface{} {
	return s.Data
}

func (s *CreateMultimodalFineTuneDatasetResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateMultimodalFineTuneDatasetResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateMultimodalFineTuneDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMultimodalFineTuneDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMultimodalFineTuneDatasetResponseBody) SetData(v interface{}) *CreateMultimodalFineTuneDatasetResponseBody {
	s.Data = v
	return s
}

func (s *CreateMultimodalFineTuneDatasetResponseBody) SetErrCode(v string) *CreateMultimodalFineTuneDatasetResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateMultimodalFineTuneDatasetResponseBody) SetErrMessage(v string) *CreateMultimodalFineTuneDatasetResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateMultimodalFineTuneDatasetResponseBody) SetRequestId(v string) *CreateMultimodalFineTuneDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMultimodalFineTuneDatasetResponseBody) SetSuccess(v bool) *CreateMultimodalFineTuneDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMultimodalFineTuneDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
