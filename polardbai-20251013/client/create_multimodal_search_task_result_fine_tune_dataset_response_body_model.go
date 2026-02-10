// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalSearchTaskResultFineTuneDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody
	GetData() interface{}
	SetErrCode(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody
	GetSuccess() *bool
}

type CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody struct {
	// example:
	//
	// JsonObject，包含数据集ID
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

func (s CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody) GetData() interface{} {
	return s.Data
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody) SetData(v interface{}) *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody {
	s.Data = v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody) SetErrCode(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody) SetErrMessage(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody) SetRequestId(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody) SetSuccess(v bool) *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
