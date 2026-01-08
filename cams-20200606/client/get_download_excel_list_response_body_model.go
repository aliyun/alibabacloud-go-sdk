// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDownloadExcelListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetDownloadExcelListResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetDownloadExcelListResponseBody
	GetCode() *string
	SetData(v string) *GetDownloadExcelListResponseBody
	GetData() *string
	SetMessage(v string) *GetDownloadExcelListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDownloadExcelListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDownloadExcelListResponseBody
	GetSuccess() *bool
}

type GetDownloadExcelListResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 39***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDownloadExcelListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDownloadExcelListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDownloadExcelListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetDownloadExcelListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDownloadExcelListResponseBody) GetData() *string {
	return s.Data
}

func (s *GetDownloadExcelListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDownloadExcelListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDownloadExcelListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDownloadExcelListResponseBody) SetAccessDeniedDetail(v string) *GetDownloadExcelListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetDownloadExcelListResponseBody) SetCode(v string) *GetDownloadExcelListResponseBody {
	s.Code = &v
	return s
}

func (s *GetDownloadExcelListResponseBody) SetData(v string) *GetDownloadExcelListResponseBody {
	s.Data = &v
	return s
}

func (s *GetDownloadExcelListResponseBody) SetMessage(v string) *GetDownloadExcelListResponseBody {
	s.Message = &v
	return s
}

func (s *GetDownloadExcelListResponseBody) SetRequestId(v string) *GetDownloadExcelListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDownloadExcelListResponseBody) SetSuccess(v bool) *GetDownloadExcelListResponseBody {
	s.Success = &v
	return s
}

func (s *GetDownloadExcelListResponseBody) Validate() error {
	return dara.Validate(s)
}
