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
	// Transcoding job ID. Supports up to 10 IDs, and multiple IDs should be separated by a comma (,). You can obtain this value in the following ways:
	//
	// When initiating a transcoding task through the [SubmitTranscodeJobs](https://help.aliyun.com/document_detail/454920.html) interface, it is the value of the returned parameter JobId.
	//
	// example:
	//
	// 86c1925fba0****,7afb201e7fa****,2cc4997378****
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	// The ID of the transcoding task. You can use one of the following methods to obtain the ID:
	//
	// 	- Obtain the value of TranscodeTaskId from the response to the [SubmitTranscodeJobs](https://help.aliyun.com/document_detail/68570.html) operation.
	//
	// 	- Obtain the value of TranscodeTaskId from the response to the [ListTranscodeTask](https://help.aliyun.com/document_detail/109120.html) operation.
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
