// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalDatasetEmbeddingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *CreateMultimodalDatasetEmbeddingResponseBody
	GetData() interface{}
	SetErrCode(v string) *CreateMultimodalDatasetEmbeddingResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateMultimodalDatasetEmbeddingResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *CreateMultimodalDatasetEmbeddingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMultimodalDatasetEmbeddingResponseBody
	GetSuccess() *bool
}

type CreateMultimodalDatasetEmbeddingResponseBody struct {
	// example:
	//
	// Embedding ID
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

func (s CreateMultimodalDatasetEmbeddingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalDatasetEmbeddingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMultimodalDatasetEmbeddingResponseBody) GetData() interface{} {
	return s.Data
}

func (s *CreateMultimodalDatasetEmbeddingResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateMultimodalDatasetEmbeddingResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateMultimodalDatasetEmbeddingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMultimodalDatasetEmbeddingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMultimodalDatasetEmbeddingResponseBody) SetData(v interface{}) *CreateMultimodalDatasetEmbeddingResponseBody {
	s.Data = v
	return s
}

func (s *CreateMultimodalDatasetEmbeddingResponseBody) SetErrCode(v string) *CreateMultimodalDatasetEmbeddingResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateMultimodalDatasetEmbeddingResponseBody) SetErrMessage(v string) *CreateMultimodalDatasetEmbeddingResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateMultimodalDatasetEmbeddingResponseBody) SetRequestId(v string) *CreateMultimodalDatasetEmbeddingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMultimodalDatasetEmbeddingResponseBody) SetSuccess(v bool) *CreateMultimodalDatasetEmbeddingResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMultimodalDatasetEmbeddingResponseBody) Validate() error {
	return dara.Validate(s)
}
