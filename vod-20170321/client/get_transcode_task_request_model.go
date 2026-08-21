// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscodeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *GetTranscodeTaskRequest
	GetJobIds() *string
	SetTranscodeTaskId(v string) *GetTranscodeTaskRequest
	GetTranscodeTaskId() *string
}

type GetTranscodeTaskRequest struct {
	// The transcoding job IDs. You can specify a maximum of 10 IDs. Separate multiple IDs with commas (,). You can obtain the IDs by using the following method:
	//
	// - Call the [SubmitTranscodeJobs](https://help.aliyun.com/document_detail/68570.html) operation to submit a transcoding task. The value of JobId in the response is the transcoding job ID.
	//
	// example:
	//
	// 86c1925fba0****,7afb201e7fa****,2cc4997378****
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	// The transcoding task ID. You can obtain the ID by using one of the following methods:
	//
	// - Call the [SubmitTranscodeJobs](https://help.aliyun.com/document_detail/68570.html) operation to submit a transcoding task. The value of TranscodeTaskId in the response is the transcoding task ID.
	//
	// - Call the [ListTranscodeTask](https://help.aliyun.com/document_detail/109120.html) operation. The value of TranscodeTaskId in the response is the transcoding task ID.
	//
	// example:
	//
	// b1b65ab107e14*****3dbb900f6c1fe0
	TranscodeTaskId *string `json:"TranscodeTaskId,omitempty" xml:"TranscodeTaskId,omitempty"`
}

func (s GetTranscodeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *GetTranscodeTaskRequest) GetTranscodeTaskId() *string {
	return s.TranscodeTaskId
}

func (s *GetTranscodeTaskRequest) SetJobIds(v string) *GetTranscodeTaskRequest {
	s.JobIds = &v
	return s
}

func (s *GetTranscodeTaskRequest) SetTranscodeTaskId(v string) *GetTranscodeTaskRequest {
	s.TranscodeTaskId = &v
	return s
}

func (s *GetTranscodeTaskRequest) Validate() error {
	return dara.Validate(s)
}
