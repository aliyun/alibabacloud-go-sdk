// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportTaskNumberDatasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ImportTaskNumberDatasResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ImportTaskNumberDatasResponseBody
	GetCode() *string
	SetData(v int64) *ImportTaskNumberDatasResponseBody
	GetData() *int64
	SetMessage(v string) *ImportTaskNumberDatasResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportTaskNumberDatasResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportTaskNumberDatasResponseBody
	GetSuccess() *bool
}

type ImportTaskNumberDatasResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1223123132123*****
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CFC2F07E-F763-7C48-1A32-6EFFB6EA344E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportTaskNumberDatasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportTaskNumberDatasResponseBody) GoString() string {
	return s.String()
}

func (s *ImportTaskNumberDatasResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ImportTaskNumberDatasResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportTaskNumberDatasResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ImportTaskNumberDatasResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportTaskNumberDatasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportTaskNumberDatasResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportTaskNumberDatasResponseBody) SetAccessDeniedDetail(v string) *ImportTaskNumberDatasResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ImportTaskNumberDatasResponseBody) SetCode(v string) *ImportTaskNumberDatasResponseBody {
	s.Code = &v
	return s
}

func (s *ImportTaskNumberDatasResponseBody) SetData(v int64) *ImportTaskNumberDatasResponseBody {
	s.Data = &v
	return s
}

func (s *ImportTaskNumberDatasResponseBody) SetMessage(v string) *ImportTaskNumberDatasResponseBody {
	s.Message = &v
	return s
}

func (s *ImportTaskNumberDatasResponseBody) SetRequestId(v string) *ImportTaskNumberDatasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportTaskNumberDatasResponseBody) SetSuccess(v bool) *ImportTaskNumberDatasResponseBody {
	s.Success = &v
	return s
}

func (s *ImportTaskNumberDatasResponseBody) Validate() error {
	return dara.Validate(s)
}
