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
	// The details of the access denial.
	//
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code. Valid values:
	//
	// - 0000: Success.
	//
	// Other values indicate failure.
	//
	// example:
	//
	// 0000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The call recording query results.
	//
	// example:
	//
	// {"FileUrl":"http://www.oss.com/temepl/a.mp3"}
	Data *QuerySoundRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request failed.
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
	// The URL of the call recording. The maximum length is 1,000 characters. The URL is valid for 1 hour.
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
