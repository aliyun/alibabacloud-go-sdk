// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncJobResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAsyncJobResultResponseBodyData) *GetAsyncJobResultResponseBody
	GetData() *GetAsyncJobResultResponseBodyData
	SetRequestId(v string) *GetAsyncJobResultResponseBody
	GetRequestId() *string
}

type GetAsyncJobResultResponseBody struct {
	Data *GetAsyncJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 6BD843FB-E7B9-45B6-ADE1-9AB1FF450148
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAsyncJobResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBody) GetData() *GetAsyncJobResultResponseBodyData {
	return s.Data
}

func (s *GetAsyncJobResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAsyncJobResultResponseBody) SetData(v *GetAsyncJobResultResponseBodyData) *GetAsyncJobResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAsyncJobResultResponseBody) SetRequestId(v string) *GetAsyncJobResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncJobResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAsyncJobResultResponseBodyData struct {
	// example:
	//
	// InvalidParameter
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// paramsIllegal
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 49E2CC28-ED1D-4CC5-854D-7D0AE2B20976
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// {\\"VideoUrl\\":\\"http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/film-summary/49E2CC28-ED1D-4CC5-854D-7D0AE2B20976_rf8mv0.mp4?Expires=1585111875&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=7uE0ppt6uwkdwAFz73PPZN4KG7****\\"}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// PROCESS_SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAsyncJobResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncJobResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAsyncJobResultResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAsyncJobResultResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *GetAsyncJobResultResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *GetAsyncJobResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorCode(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorMessage(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetJobId(v string) *GetAsyncJobResultResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetResult(v string) *GetAsyncJobResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetStatus(v string) *GetAsyncJobResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
