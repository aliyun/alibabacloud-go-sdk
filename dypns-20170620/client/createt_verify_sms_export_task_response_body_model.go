// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatetVerifySmsExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreatetVerifySmsExportTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreatetVerifySmsExportTaskResponseBody
	GetCode() *string
	SetData(v *CreatetVerifySmsExportTaskResponseBodyData) *CreatetVerifySmsExportTaskResponseBody
	GetData() *CreatetVerifySmsExportTaskResponseBodyData
	SetMessage(v string) *CreatetVerifySmsExportTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatetVerifySmsExportTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatetVerifySmsExportTaskResponseBody
	GetSuccess() *bool
}

type CreatetVerifySmsExportTaskResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreatetVerifySmsExportTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2E7CA885-8D97-C497-A02C-2D31214D3285
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatetVerifySmsExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatetVerifySmsExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreatetVerifySmsExportTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreatetVerifySmsExportTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatetVerifySmsExportTaskResponseBody) GetData() *CreatetVerifySmsExportTaskResponseBodyData {
	return s.Data
}

func (s *CreatetVerifySmsExportTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatetVerifySmsExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatetVerifySmsExportTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatetVerifySmsExportTaskResponseBody) SetAccessDeniedDetail(v string) *CreatetVerifySmsExportTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreatetVerifySmsExportTaskResponseBody) SetCode(v string) *CreatetVerifySmsExportTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreatetVerifySmsExportTaskResponseBody) SetData(v *CreatetVerifySmsExportTaskResponseBodyData) *CreatetVerifySmsExportTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreatetVerifySmsExportTaskResponseBody) SetMessage(v string) *CreatetVerifySmsExportTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreatetVerifySmsExportTaskResponseBody) SetRequestId(v string) *CreatetVerifySmsExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatetVerifySmsExportTaskResponseBody) SetSuccess(v bool) *CreatetVerifySmsExportTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreatetVerifySmsExportTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatetVerifySmsExportTaskResponseBodyData struct {
	// example:
	//
	// 69
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreatetVerifySmsExportTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatetVerifySmsExportTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatetVerifySmsExportTaskResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreatetVerifySmsExportTaskResponseBodyData) SetTaskId(v int64) *CreatetVerifySmsExportTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreatetVerifySmsExportTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
