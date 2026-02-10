// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalFineTuneDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ListMultimodalFineTuneDatasetResponseBody
	GetData() interface{}
	SetErrCode(v string) *ListMultimodalFineTuneDatasetResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListMultimodalFineTuneDatasetResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ListMultimodalFineTuneDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMultimodalFineTuneDatasetResponseBody
	GetSuccess() *bool
}

type ListMultimodalFineTuneDatasetResponseBody struct {
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

func (s ListMultimodalFineTuneDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalFineTuneDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultimodalFineTuneDatasetResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ListMultimodalFineTuneDatasetResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListMultimodalFineTuneDatasetResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListMultimodalFineTuneDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMultimodalFineTuneDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMultimodalFineTuneDatasetResponseBody) SetData(v interface{}) *ListMultimodalFineTuneDatasetResponseBody {
	s.Data = v
	return s
}

func (s *ListMultimodalFineTuneDatasetResponseBody) SetErrCode(v string) *ListMultimodalFineTuneDatasetResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListMultimodalFineTuneDatasetResponseBody) SetErrMessage(v string) *ListMultimodalFineTuneDatasetResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListMultimodalFineTuneDatasetResponseBody) SetRequestId(v string) *ListMultimodalFineTuneDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultimodalFineTuneDatasetResponseBody) SetSuccess(v bool) *ListMultimodalFineTuneDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *ListMultimodalFineTuneDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
