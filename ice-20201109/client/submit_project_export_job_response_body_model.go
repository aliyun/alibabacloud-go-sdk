// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitProjectExportJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitProjectExportJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitProjectExportJobResponseBody
	GetRequestId() *string
}

type SubmitProjectExportJobResponseBody struct {
	// The ID of the project export task.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitProjectExportJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitProjectExportJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitProjectExportJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitProjectExportJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitProjectExportJobResponseBody) SetJobId(v string) *SubmitProjectExportJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitProjectExportJobResponseBody) SetRequestId(v string) *SubmitProjectExportJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitProjectExportJobResponseBody) Validate() error {
	return dara.Validate(s)
}
