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
	// A1F44EC4-118D-4A03-B213-F908F36F7DAA
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
	// 7DFDA846-178B-4ADB-B69A-62C641214D81
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// {\\"Outputs\\":[{\\"ImageURL\\":\\"http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-cover/2020-03-25-11/pic_4t7zW6R6SUGn4DLF.mp4_2375_1920_1080_96ce5a96b5b16628cd778c035b68356d_beautified.png?Expires=1585136160&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=MDd7BqbivlLRd16MTKbPFQHV3u****\\",\\"Confidence\\":8.426481079120514},{\\"ImageURL\\":\\"http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-cover/2020-03-25-11/pic_4t7zW6R6SUGn4DLF.mp4_2996_1920_1080_d5df0556bf420242c84fe6f7a45d01e1_beautified.png?Expires=1585136160&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=htaW5G%2BpqPBW%2BEMTe01ckVoGsQ****\\",\\"Confidence\\":6.225726566341124}]}
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
