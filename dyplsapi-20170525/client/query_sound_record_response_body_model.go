// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySoundRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuerySoundRecordResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QuerySoundRecordResponseBody
	GetCode() *string
	SetData(v *QuerySoundRecordResponseBodyData) *QuerySoundRecordResponseBody
	GetData() *QuerySoundRecordResponseBodyData
	SetMessage(v string) *QuerySoundRecordResponseBody
	GetMessage() *string
	SetSuccess(v bool) *QuerySoundRecordResponseBody
	GetSuccess() *bool
}

type QuerySoundRecordResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QuerySoundRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 返回信息
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回是否成功 true  表示成功 false表示失败
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySoundRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySoundRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySoundRecordResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuerySoundRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySoundRecordResponseBody) GetData() *QuerySoundRecordResponseBodyData {
	return s.Data
}

func (s *QuerySoundRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySoundRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySoundRecordResponseBody) SetAccessDeniedDetail(v string) *QuerySoundRecordResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySoundRecordResponseBody) SetCode(v string) *QuerySoundRecordResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySoundRecordResponseBody) SetData(v *QuerySoundRecordResponseBodyData) *QuerySoundRecordResponseBody {
	s.Data = v
	return s
}

func (s *QuerySoundRecordResponseBody) SetMessage(v string) *QuerySoundRecordResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySoundRecordResponseBody) SetSuccess(v bool) *QuerySoundRecordResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySoundRecordResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySoundRecordResponseBodyData struct {
	// 通话录音url路径，最大长度1000，有效期1小时
	//
	// example:
	//
	// http://www.oss.com/temepl/a.mp3
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s QuerySoundRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySoundRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySoundRecordResponseBodyData) GetFileUrl() *string {
	return s.FileUrl
}

func (s *QuerySoundRecordResponseBodyData) SetFileUrl(v string) *QuerySoundRecordResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *QuerySoundRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
