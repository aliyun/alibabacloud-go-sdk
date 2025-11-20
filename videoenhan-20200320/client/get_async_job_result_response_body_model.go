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
	// 43A0AEB6-45F4-4138-8E89-E1A5D63200E3
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	// {\\"VideoUrl\\":\\"http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-super-resolution/2020-03-26-02/55%3A59-XFLsuEKpPxleHl4n.mp4?Expires=1585193159&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=OqsxMwAKsFVjPO%2Brgw%2BpXHIm6v****\\"}
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
