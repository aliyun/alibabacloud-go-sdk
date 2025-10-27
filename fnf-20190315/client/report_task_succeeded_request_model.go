// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportTaskSucceededRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutput(v string) *ReportTaskSucceededRequest
	GetOutput() *string
	SetTaskToken(v string) *ReportTaskSucceededRequest
	GetTaskToken() *string
}

type ReportTaskSucceededRequest struct {
	// The output information of the task whose execution success you want to report.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"key":"value"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The token of the task whose execution you want to report. The task token is passed to the called service, such as Message Service (MNS) or Function Compute. For MNS, the value of this parameter can be obtained from a message. For Function Compute, the value of this parameter can be obtained from an event. For more information, see [Service integration modes](https://help.aliyun.com/document_detail/2592915.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// emptyString
	TaskToken *string `json:"TaskToken,omitempty" xml:"TaskToken,omitempty"`
}

func (s ReportTaskSucceededRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportTaskSucceededRequest) GoString() string {
	return s.String()
}

func (s *ReportTaskSucceededRequest) GetOutput() *string {
	return s.Output
}

func (s *ReportTaskSucceededRequest) GetTaskToken() *string {
	return s.TaskToken
}

func (s *ReportTaskSucceededRequest) SetOutput(v string) *ReportTaskSucceededRequest {
	s.Output = &v
	return s
}

func (s *ReportTaskSucceededRequest) SetTaskToken(v string) *ReportTaskSucceededRequest {
	s.TaskToken = &v
	return s
}

func (s *ReportTaskSucceededRequest) Validate() error {
	return dara.Validate(s)
}
