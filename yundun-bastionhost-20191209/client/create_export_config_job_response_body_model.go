// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExportConfigJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateExportConfigJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateExportConfigJobResponseBody
	GetRequestId() *string
}

type CreateExportConfigJobResponseBody struct {
	// The ID of the configuration backup export task.
	//
	// > You can call the GetExportConfigJob operation to obtain the details of this task. Make sure to record the task ID.
	//
	// example:
	//
	// 1
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The unique identifier that Alibaba Cloud generates for the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExportConfigJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExportConfigJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExportConfigJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateExportConfigJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExportConfigJobResponseBody) SetJobId(v string) *CreateExportConfigJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateExportConfigJobResponseBody) SetRequestId(v string) *CreateExportConfigJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExportConfigJobResponseBody) Validate() error {
	return dara.Validate(s)
}
