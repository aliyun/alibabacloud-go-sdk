// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsVerifyStatisticRecordsExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateSmsVerifyStatisticRecordsExportTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateSmsVerifyStatisticRecordsExportTaskResponseBody
	GetCode() *string
	SetMessage(v string) *CreateSmsVerifyStatisticRecordsExportTaskResponseBody
	GetMessage() *string
	SetModel(v map[string]interface{}) *CreateSmsVerifyStatisticRecordsExportTaskResponseBody
	GetModel() map[string]interface{}
	SetSuccess(v bool) *CreateSmsVerifyStatisticRecordsExportTaskResponseBody
	GetSuccess() *bool
}

type CreateSmsVerifyStatisticRecordsExportTaskResponseBody struct {
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
	// 示例值
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSmsVerifyStatisticRecordsExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsVerifyStatisticRecordsExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponseBody) GetModel() map[string]interface{} {
	return s.Model
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponseBody) SetAccessDeniedDetail(v string) *CreateSmsVerifyStatisticRecordsExportTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponseBody) SetCode(v string) *CreateSmsVerifyStatisticRecordsExportTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponseBody) SetMessage(v string) *CreateSmsVerifyStatisticRecordsExportTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponseBody) SetModel(v map[string]interface{}) *CreateSmsVerifyStatisticRecordsExportTaskResponseBody {
	s.Model = v
	return s
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponseBody) SetSuccess(v bool) *CreateSmsVerifyStatisticRecordsExportTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSmsVerifyStatisticRecordsExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
