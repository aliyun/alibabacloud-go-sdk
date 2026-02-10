// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultimodalFineTuneDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *UpdateMultimodalFineTuneDatasetResponseBody
	GetData() interface{}
	SetErrCode(v string) *UpdateMultimodalFineTuneDatasetResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdateMultimodalFineTuneDatasetResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *UpdateMultimodalFineTuneDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMultimodalFineTuneDatasetResponseBody
	GetSuccess() *bool
}

type UpdateMultimodalFineTuneDatasetResponseBody struct {
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

func (s UpdateMultimodalFineTuneDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultimodalFineTuneDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMultimodalFineTuneDatasetResponseBody) GetData() interface{} {
	return s.Data
}

func (s *UpdateMultimodalFineTuneDatasetResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateMultimodalFineTuneDatasetResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdateMultimodalFineTuneDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMultimodalFineTuneDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMultimodalFineTuneDatasetResponseBody) SetData(v interface{}) *UpdateMultimodalFineTuneDatasetResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMultimodalFineTuneDatasetResponseBody) SetErrCode(v string) *UpdateMultimodalFineTuneDatasetResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateMultimodalFineTuneDatasetResponseBody) SetErrMessage(v string) *UpdateMultimodalFineTuneDatasetResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateMultimodalFineTuneDatasetResponseBody) SetRequestId(v string) *UpdateMultimodalFineTuneDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMultimodalFineTuneDatasetResponseBody) SetSuccess(v bool) *UpdateMultimodalFineTuneDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMultimodalFineTuneDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
