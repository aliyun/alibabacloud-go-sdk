// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaResourceIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMediaResourceIdResponseBody
	GetCode() *string
	SetData(v *GetMediaResourceIdResponseBodyData) *GetMediaResourceIdResponseBody
	GetData() *GetMediaResourceIdResponseBodyData
	SetRequestId(v string) *GetMediaResourceIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMediaResourceIdResponseBody
	GetSuccess() *bool
}

type GetMediaResourceIdResponseBody struct {
	// 请求状态码。
	//
	// - 返回OK代表请求成功。
	//
	// - 其他错误码，请参见[错误码列表](https://help.aliyun.com/document_detail/101346.html)。
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据。
	Data *GetMediaResourceIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 请求ID。
	//
	// example:
	//
	// F07CF237-F6E3-5F77-B91B-F9B7C5DE84AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用接口是否成功。取值：
	//
	// - **true**：调用成功。
	//
	// - **false**：调用失败。
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMediaResourceIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaResourceIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaResourceIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMediaResourceIdResponseBody) GetData() *GetMediaResourceIdResponseBodyData {
	return s.Data
}

func (s *GetMediaResourceIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaResourceIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMediaResourceIdResponseBody) SetCode(v string) *GetMediaResourceIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetMediaResourceIdResponseBody) SetData(v *GetMediaResourceIdResponseBodyData) *GetMediaResourceIdResponseBody {
	s.Data = v
	return s
}

func (s *GetMediaResourceIdResponseBody) SetRequestId(v string) *GetMediaResourceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaResourceIdResponseBody) SetSuccess(v bool) *GetMediaResourceIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetMediaResourceIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaResourceIdResponseBodyData struct {
	// 资源下载地址。
	//
	// example:
	//
	// http://test-example.com/download.jpg
	ResUrlDownload *string `json:"ResUrlDownload,omitempty" xml:"ResUrlDownload,omitempty"`
	// 资源ID。
	//
	// example:
	//
	// 0
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s GetMediaResourceIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMediaResourceIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMediaResourceIdResponseBodyData) GetResUrlDownload() *string {
	return s.ResUrlDownload
}

func (s *GetMediaResourceIdResponseBodyData) GetResourceId() *int64 {
	return s.ResourceId
}

func (s *GetMediaResourceIdResponseBodyData) SetResUrlDownload(v string) *GetMediaResourceIdResponseBodyData {
	s.ResUrlDownload = &v
	return s
}

func (s *GetMediaResourceIdResponseBodyData) SetResourceId(v int64) *GetMediaResourceIdResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *GetMediaResourceIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
