// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadMultimodalSearchTaskResultMetadataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *DownloadMultimodalSearchTaskResultMetadataResponseBody
	GetData() interface{}
	SetErrCode(v string) *DownloadMultimodalSearchTaskResultMetadataResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DownloadMultimodalSearchTaskResultMetadataResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DownloadMultimodalSearchTaskResultMetadataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DownloadMultimodalSearchTaskResultMetadataResponseBody
	GetSuccess() *bool
}

type DownloadMultimodalSearchTaskResultMetadataResponseBody struct {
	// String
	//
	// example:
	//
	// 文件下载Url
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

func (s DownloadMultimodalSearchTaskResultMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadMultimodalSearchTaskResultMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponseBody) SetData(v interface{}) *DownloadMultimodalSearchTaskResultMetadataResponseBody {
	s.Data = v
	return s
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponseBody) SetErrCode(v string) *DownloadMultimodalSearchTaskResultMetadataResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponseBody) SetErrMessage(v string) *DownloadMultimodalSearchTaskResultMetadataResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponseBody) SetRequestId(v string) *DownloadMultimodalSearchTaskResultMetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponseBody) SetSuccess(v bool) *DownloadMultimodalSearchTaskResultMetadataResponseBody {
	s.Success = &v
	return s
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponseBody) Validate() error {
	return dara.Validate(s)
}
