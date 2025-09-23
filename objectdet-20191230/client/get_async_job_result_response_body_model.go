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
	// 87FC80D2-2571-4BBD-BD61-AFF7912C556D
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
	// 35B11E1B-800C-4598-B5AA-577E3BDBD917
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// {\\"inputFile\\":\\"oss://public-vigen-video/guotian.xgt/test_images/test_video\\",\\"width\\":1280,\\"height\\":720,\\"frames\\":[{\\"time\\":6124533574,\\"elements\\":[{\\"type\\":\\"PERSON\\",\\"score\\":0.7812,\\"x\\":289,\\"y\\":271,\\"width\\":100,\\"height\\":156},{\\"type\\":\\"PERSON\\",\\"score\\":0.4377,\\"x\\":917,\\"y\\":267,\\"width\\":34,\\"height\\":51}]}]}]}"}
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
