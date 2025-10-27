// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportTaskFailedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCause(v string) *ReportTaskFailedRequest
	GetCause() *string
	SetError(v string) *ReportTaskFailedRequest
	GetError() *string
	SetTaskToken(v string) *ReportTaskFailedRequest
	GetTaskToken() *string
}

type ReportTaskFailedRequest struct {
	// The cause of the failure. The value must be 1 to 4,096 characters in length.
	//
	// example:
	//
	// emptyString
	Cause *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	// The error code for the failed task. The error code must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// nill
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The token of the task whose execution you want to report. The task token is passed to the called service, such as Message Service (MNS) or Function Compute. For MNS, the value of this parameter can be obtained from a message. For Function Compute, the value of this parameter can be obtained from an event. For more information, see [Service integration modes](https://help.aliyun.com/document_detail/2592915.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// emptyString
	TaskToken *string `json:"TaskToken,omitempty" xml:"TaskToken,omitempty"`
}

func (s ReportTaskFailedRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportTaskFailedRequest) GoString() string {
	return s.String()
}

func (s *ReportTaskFailedRequest) GetCause() *string {
	return s.Cause
}

func (s *ReportTaskFailedRequest) GetError() *string {
	return s.Error
}

func (s *ReportTaskFailedRequest) GetTaskToken() *string {
	return s.TaskToken
}

func (s *ReportTaskFailedRequest) SetCause(v string) *ReportTaskFailedRequest {
	s.Cause = &v
	return s
}

func (s *ReportTaskFailedRequest) SetError(v string) *ReportTaskFailedRequest {
	s.Error = &v
	return s
}

func (s *ReportTaskFailedRequest) SetTaskToken(v string) *ReportTaskFailedRequest {
	s.TaskToken = &v
	return s
}

func (s *ReportTaskFailedRequest) Validate() error {
	return dara.Validate(s)
}
