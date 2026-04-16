// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudImportTaskTelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudImportTaskTelResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudImportTaskTelResponseBody
	GetCode() *string
	SetData(v *CloudImportTaskTelResponseBodyData) *CloudImportTaskTelResponseBody
	GetData() *CloudImportTaskTelResponseBodyData
	SetMessage(v string) *CloudImportTaskTelResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudImportTaskTelResponseBody
	GetRequestId() *string
}

type CloudImportTaskTelResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudImportTaskTelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudImportTaskTelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudImportTaskTelResponseBody) GoString() string {
	return s.String()
}

func (s *CloudImportTaskTelResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudImportTaskTelResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudImportTaskTelResponseBody) GetData() *CloudImportTaskTelResponseBodyData {
	return s.Data
}

func (s *CloudImportTaskTelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudImportTaskTelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudImportTaskTelResponseBody) SetAccessDeniedDetail(v string) *CloudImportTaskTelResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudImportTaskTelResponseBody) SetCode(v string) *CloudImportTaskTelResponseBody {
	s.Code = &v
	return s
}

func (s *CloudImportTaskTelResponseBody) SetData(v *CloudImportTaskTelResponseBodyData) *CloudImportTaskTelResponseBody {
	s.Data = v
	return s
}

func (s *CloudImportTaskTelResponseBody) SetMessage(v string) *CloudImportTaskTelResponseBody {
	s.Message = &v
	return s
}

func (s *CloudImportTaskTelResponseBody) SetRequestId(v string) *CloudImportTaskTelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudImportTaskTelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudImportTaskTelResponseBodyData struct {
	// 批次Id
	//
	// example:
	//
	// 0
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// 请求号码总数
	//
	// example:
	//
	// 30
	ImportTotal *string `json:"ImportTotal,omitempty" xml:"ImportTotal,omitempty"`
	// 非法号码数
	//
	// example:
	//
	// 0
	InvalidTotal *string `json:"InvalidTotal,omitempty" xml:"InvalidTotal,omitempty"`
	// 成功导入号码数
	//
	// example:
	//
	// 30
	SuccessTotal *string `json:"SuccessTotal,omitempty" xml:"SuccessTotal,omitempty"`
	// 任务Id
	//
	// example:
	//
	// 13322333
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CloudImportTaskTelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudImportTaskTelResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudImportTaskTelResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *CloudImportTaskTelResponseBodyData) GetImportTotal() *string {
	return s.ImportTotal
}

func (s *CloudImportTaskTelResponseBodyData) GetInvalidTotal() *string {
	return s.InvalidTotal
}

func (s *CloudImportTaskTelResponseBodyData) GetSuccessTotal() *string {
	return s.SuccessTotal
}

func (s *CloudImportTaskTelResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CloudImportTaskTelResponseBodyData) SetFileId(v string) *CloudImportTaskTelResponseBodyData {
	s.FileId = &v
	return s
}

func (s *CloudImportTaskTelResponseBodyData) SetImportTotal(v string) *CloudImportTaskTelResponseBodyData {
	s.ImportTotal = &v
	return s
}

func (s *CloudImportTaskTelResponseBodyData) SetInvalidTotal(v string) *CloudImportTaskTelResponseBodyData {
	s.InvalidTotal = &v
	return s
}

func (s *CloudImportTaskTelResponseBodyData) SetSuccessTotal(v string) *CloudImportTaskTelResponseBodyData {
	s.SuccessTotal = &v
	return s
}

func (s *CloudImportTaskTelResponseBodyData) SetTaskId(v string) *CloudImportTaskTelResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CloudImportTaskTelResponseBodyData) Validate() error {
	return dara.Validate(s)
}
