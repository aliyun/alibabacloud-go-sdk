// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamBatchJobMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateStreamBatchJobMappingResponseBody
	GetCode() *string
	SetData(v *CreateStreamBatchJobMappingResponseBodyData) *CreateStreamBatchJobMappingResponseBody
	GetData() *CreateStreamBatchJobMappingResponseBodyData
	SetHttpStatusCode(v int32) *CreateStreamBatchJobMappingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateStreamBatchJobMappingResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateStreamBatchJobMappingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateStreamBatchJobMappingResponseBody
	GetSuccess() *bool
}

type CreateStreamBatchJobMappingResponseBody struct {
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateStreamBatchJobMappingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateStreamBatchJobMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamBatchJobMappingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStreamBatchJobMappingResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateStreamBatchJobMappingResponseBody) GetData() *CreateStreamBatchJobMappingResponseBodyData {
	return s.Data
}

func (s *CreateStreamBatchJobMappingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateStreamBatchJobMappingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateStreamBatchJobMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStreamBatchJobMappingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateStreamBatchJobMappingResponseBody) SetCode(v string) *CreateStreamBatchJobMappingResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStreamBatchJobMappingResponseBody) SetData(v *CreateStreamBatchJobMappingResponseBodyData) *CreateStreamBatchJobMappingResponseBody {
	s.Data = v
	return s
}

func (s *CreateStreamBatchJobMappingResponseBody) SetHttpStatusCode(v int32) *CreateStreamBatchJobMappingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateStreamBatchJobMappingResponseBody) SetMessage(v string) *CreateStreamBatchJobMappingResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStreamBatchJobMappingResponseBody) SetRequestId(v string) *CreateStreamBatchJobMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStreamBatchJobMappingResponseBody) SetSuccess(v bool) *CreateStreamBatchJobMappingResponseBody {
	s.Success = &v
	return s
}

func (s *CreateStreamBatchJobMappingResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateStreamBatchJobMappingResponseBodyData struct {
	// example:
	//
	// 7083701105376640
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// 123123
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// /dev/streamJob/7083701105376640?env=DEV&projectId=7081229106458752&tenantId=300001420
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateStreamBatchJobMappingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamBatchJobMappingResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateStreamBatchJobMappingResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *CreateStreamBatchJobMappingResponseBodyData) GetJobId() *int64 {
	return s.JobId
}

func (s *CreateStreamBatchJobMappingResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *CreateStreamBatchJobMappingResponseBodyData) SetFileId(v string) *CreateStreamBatchJobMappingResponseBodyData {
	s.FileId = &v
	return s
}

func (s *CreateStreamBatchJobMappingResponseBodyData) SetJobId(v int64) *CreateStreamBatchJobMappingResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateStreamBatchJobMappingResponseBodyData) SetUrl(v string) *CreateStreamBatchJobMappingResponseBodyData {
	s.Url = &v
	return s
}

func (s *CreateStreamBatchJobMappingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
