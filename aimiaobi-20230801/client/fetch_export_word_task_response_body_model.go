// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchExportWordTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FetchExportWordTaskResponseBody
	GetCode() *string
	SetData(v *FetchExportWordTaskResponseBodyData) *FetchExportWordTaskResponseBody
	GetData() *FetchExportWordTaskResponseBodyData
	SetHttpStatusCode(v int32) *FetchExportWordTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *FetchExportWordTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *FetchExportWordTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FetchExportWordTaskResponseBody
	GetSuccess() *bool
}

type FetchExportWordTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *FetchExportWordTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FetchExportWordTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FetchExportWordTaskResponseBody) GoString() string {
	return s.String()
}

func (s *FetchExportWordTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *FetchExportWordTaskResponseBody) GetData() *FetchExportWordTaskResponseBodyData {
	return s.Data
}

func (s *FetchExportWordTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *FetchExportWordTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FetchExportWordTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FetchExportWordTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FetchExportWordTaskResponseBody) SetCode(v string) *FetchExportWordTaskResponseBody {
	s.Code = &v
	return s
}

func (s *FetchExportWordTaskResponseBody) SetData(v *FetchExportWordTaskResponseBodyData) *FetchExportWordTaskResponseBody {
	s.Data = v
	return s
}

func (s *FetchExportWordTaskResponseBody) SetHttpStatusCode(v int32) *FetchExportWordTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FetchExportWordTaskResponseBody) SetMessage(v string) *FetchExportWordTaskResponseBody {
	s.Message = &v
	return s
}

func (s *FetchExportWordTaskResponseBody) SetRequestId(v string) *FetchExportWordTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchExportWordTaskResponseBody) SetSuccess(v bool) *FetchExportWordTaskResponseBody {
	s.Success = &v
	return s
}

func (s *FetchExportWordTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FetchExportWordTaskResponseBodyData struct {
	// example:
	//
	// http://xxx/xxx.xls
	FileUrl   *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	TaskStats *string `json:"TaskStats,omitempty" xml:"TaskStats,omitempty"`
}

func (s FetchExportWordTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FetchExportWordTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *FetchExportWordTaskResponseBodyData) GetFileUrl() *string {
	return s.FileUrl
}

func (s *FetchExportWordTaskResponseBodyData) GetTaskStats() *string {
	return s.TaskStats
}

func (s *FetchExportWordTaskResponseBodyData) SetFileUrl(v string) *FetchExportWordTaskResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *FetchExportWordTaskResponseBodyData) SetTaskStats(v string) *FetchExportWordTaskResponseBodyData {
	s.TaskStats = &v
	return s
}

func (s *FetchExportWordTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
