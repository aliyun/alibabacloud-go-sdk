// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalSearchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *CreateMultimodalSearchTaskResponseBody
	GetData() interface{}
	SetErrCode(v string) *CreateMultimodalSearchTaskResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateMultimodalSearchTaskResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *CreateMultimodalSearchTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMultimodalSearchTaskResponseBody
	GetSuccess() *bool
}

type CreateMultimodalSearchTaskResponseBody struct {
	// example:
	//
	// embedding ID
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

func (s CreateMultimodalSearchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalSearchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMultimodalSearchTaskResponseBody) GetData() interface{} {
	return s.Data
}

func (s *CreateMultimodalSearchTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateMultimodalSearchTaskResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateMultimodalSearchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMultimodalSearchTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMultimodalSearchTaskResponseBody) SetData(v interface{}) *CreateMultimodalSearchTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateMultimodalSearchTaskResponseBody) SetErrCode(v string) *CreateMultimodalSearchTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateMultimodalSearchTaskResponseBody) SetErrMessage(v string) *CreateMultimodalSearchTaskResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateMultimodalSearchTaskResponseBody) SetRequestId(v string) *CreateMultimodalSearchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMultimodalSearchTaskResponseBody) SetSuccess(v bool) *CreateMultimodalSearchTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMultimodalSearchTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
