// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalSearchTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ListMultimodalSearchTaskResultResponseBody
	GetData() interface{}
	SetErrCode(v string) *ListMultimodalSearchTaskResultResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListMultimodalSearchTaskResultResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ListMultimodalSearchTaskResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMultimodalSearchTaskResultResponseBody
	GetSuccess() *bool
}

type ListMultimodalSearchTaskResultResponseBody struct {
	// example:
	//
	// {"Data": {
	//
	//   # 分页信息,
	//
	//   "Total": 5,
	//
	//   "PageNumber": 1,
	//
	//   "PageSize": 50,
	//
	//   # 数据列表
	//
	//   "Items": [
	//
	//     # 每个JSONObject为一个视频/图片的元信息
	//
	//     {
	//
	//       # 用户oss图片地址
	//
	//       "OssUri": "",
	//
	//       # 缩略图
	//
	//       "ThumbnailUrl": "",
	//
	//       # 原图，点击查看
	//
	//       "DownloadUrl": "",
	//
	//       # 文件类型（视频/图片）
	//
	//       "FileType": "",
	//
	//       # 文件名
	//
	//       "FileName": "",
	//
	//     }, {}, ..
	//
	//   ]
	//
	// },
	//
	// "Success": true,
	//
	// "RequestId": "***"
	//
	// }
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

func (s ListMultimodalSearchTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalSearchTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultimodalSearchTaskResultResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ListMultimodalSearchTaskResultResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListMultimodalSearchTaskResultResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListMultimodalSearchTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMultimodalSearchTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMultimodalSearchTaskResultResponseBody) SetData(v interface{}) *ListMultimodalSearchTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *ListMultimodalSearchTaskResultResponseBody) SetErrCode(v string) *ListMultimodalSearchTaskResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListMultimodalSearchTaskResultResponseBody) SetErrMessage(v string) *ListMultimodalSearchTaskResultResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListMultimodalSearchTaskResultResponseBody) SetRequestId(v string) *ListMultimodalSearchTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultimodalSearchTaskResultResponseBody) SetSuccess(v bool) *ListMultimodalSearchTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *ListMultimodalSearchTaskResultResponseBody) Validate() error {
	return dara.Validate(s)
}
