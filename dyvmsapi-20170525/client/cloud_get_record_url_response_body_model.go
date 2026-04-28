// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetRecordUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudGetRecordUrlResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudGetRecordUrlResponseBody
	GetCode() *string
	SetData(v *CloudGetRecordUrlResponseBodyData) *CloudGetRecordUrlResponseBody
	GetData() *CloudGetRecordUrlResponseBodyData
	SetMessage(v string) *CloudGetRecordUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudGetRecordUrlResponseBody
	GetRequestId() *string
}

type CloudGetRecordUrlResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudGetRecordUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6086693B-2250-17CE-A41F-06259AB6DB1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudGetRecordUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudGetRecordUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CloudGetRecordUrlResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudGetRecordUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudGetRecordUrlResponseBody) GetData() *CloudGetRecordUrlResponseBodyData {
	return s.Data
}

func (s *CloudGetRecordUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudGetRecordUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudGetRecordUrlResponseBody) SetAccessDeniedDetail(v string) *CloudGetRecordUrlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudGetRecordUrlResponseBody) SetCode(v string) *CloudGetRecordUrlResponseBody {
	s.Code = &v
	return s
}

func (s *CloudGetRecordUrlResponseBody) SetData(v *CloudGetRecordUrlResponseBodyData) *CloudGetRecordUrlResponseBody {
	s.Data = v
	return s
}

func (s *CloudGetRecordUrlResponseBody) SetMessage(v string) *CloudGetRecordUrlResponseBody {
	s.Message = &v
	return s
}

func (s *CloudGetRecordUrlResponseBody) SetRequestId(v string) *CloudGetRecordUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudGetRecordUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudGetRecordUrlResponseBodyData struct {
	// 试听或下载地址，默认120分钟过期（可配置），过期后，需重新请求接口
	//
	// example:
	//
	// http://voice-1.aliyun.cn/31102016/record/7000000/7000000-2016103111%204742-01087120766-01087120766--record-sip-1-1477885662.222.mp3?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20161114T100003Z&X-Amz-SignedHeaders=host&X-Amz-Expires=119&X-Amz-Credential=AKIAOTCJOBZFUWFI7FSA%2F20161114%2Fcn-north-1%2Fs3%2Faws4_request&X-Amz-Signature=0cda4389ecc743acc71d76a320705afd0c175c5ad2286305675c4dee5189b9c8
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CloudGetRecordUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudGetRecordUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudGetRecordUrlResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *CloudGetRecordUrlResponseBodyData) SetUrl(v string) *CloudGetRecordUrlResponseBodyData {
	s.Url = &v
	return s
}

func (s *CloudGetRecordUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
