// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPptInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPptInfoResponseBody
	GetCode() *string
	SetData(v *GetPptInfoResponseBodyData) *GetPptInfoResponseBody
	GetData() *GetPptInfoResponseBodyData
	SetHttpStatusCode(v int32) *GetPptInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPptInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPptInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPptInfoResponseBody
	GetSuccess() *bool
}

type GetPptInfoResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetPptInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPptInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPptInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPptInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPptInfoResponseBody) GetData() *GetPptInfoResponseBodyData {
	return s.Data
}

func (s *GetPptInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPptInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPptInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPptInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPptInfoResponseBody) SetCode(v string) *GetPptInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetPptInfoResponseBody) SetData(v *GetPptInfoResponseBodyData) *GetPptInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetPptInfoResponseBody) SetHttpStatusCode(v int32) *GetPptInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPptInfoResponseBody) SetMessage(v string) *GetPptInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetPptInfoResponseBody) SetRequestId(v string) *GetPptInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPptInfoResponseBody) SetSuccess(v bool) *GetPptInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetPptInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPptInfoResponseBodyData struct {
	ExportFileLink []*string `json:"ExportFileLink,omitempty" xml:"ExportFileLink,omitempty" type:"Repeated"`
	// example:
	//
	// xxx-xxx-xx
	ExportTaskId *string `json:"ExportTaskId,omitempty" xml:"ExportTaskId,omitempty"`
	// example:
	//
	// 5423431
	PptArtifactId *string `json:"PptArtifactId,omitempty" xml:"PptArtifactId,omitempty"`
	// example:
	//
	// 11231232
	PptProcessId *string `json:"PptProcessId,omitempty" xml:"PptProcessId,omitempty"`
	// example:
	//
	// 关于班会主题的PPT
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// xxx-xxx-xx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetPptInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPptInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPptInfoResponseBodyData) GetExportFileLink() []*string {
	return s.ExportFileLink
}

func (s *GetPptInfoResponseBodyData) GetExportTaskId() *string {
	return s.ExportTaskId
}

func (s *GetPptInfoResponseBodyData) GetPptArtifactId() *string {
	return s.PptArtifactId
}

func (s *GetPptInfoResponseBodyData) GetPptProcessId() *string {
	return s.PptProcessId
}

func (s *GetPptInfoResponseBodyData) GetQuery() *string {
	return s.Query
}

func (s *GetPptInfoResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetPptInfoResponseBodyData) SetExportFileLink(v []*string) *GetPptInfoResponseBodyData {
	s.ExportFileLink = v
	return s
}

func (s *GetPptInfoResponseBodyData) SetExportTaskId(v string) *GetPptInfoResponseBodyData {
	s.ExportTaskId = &v
	return s
}

func (s *GetPptInfoResponseBodyData) SetPptArtifactId(v string) *GetPptInfoResponseBodyData {
	s.PptArtifactId = &v
	return s
}

func (s *GetPptInfoResponseBodyData) SetPptProcessId(v string) *GetPptInfoResponseBodyData {
	s.PptProcessId = &v
	return s
}

func (s *GetPptInfoResponseBodyData) SetQuery(v string) *GetPptInfoResponseBodyData {
	s.Query = &v
	return s
}

func (s *GetPptInfoResponseBodyData) SetTaskId(v string) *GetPptInfoResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetPptInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
