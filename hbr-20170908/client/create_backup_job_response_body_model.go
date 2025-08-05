// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateBackupJobResponseBody
	GetCode() *string
	SetJobId(v string) *CreateBackupJobResponseBody
	GetJobId() *string
	SetMessage(v string) *CreateBackupJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBackupJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBackupJobResponseBody
	GetSuccess() *bool
}

type CreateBackupJobResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the backup job.
	//
	// example:
	//
	// job-000csy09q50a2jdcbwbo
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 25F49E7B-7E39-542E-83AD-62E6E7F73786
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBackupJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBackupJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateBackupJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBackupJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBackupJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBackupJobResponseBody) SetCode(v string) *CreateBackupJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBackupJobResponseBody) SetJobId(v string) *CreateBackupJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateBackupJobResponseBody) SetMessage(v string) *CreateBackupJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBackupJobResponseBody) SetRequestId(v string) *CreateBackupJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackupJobResponseBody) SetSuccess(v bool) *CreateBackupJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBackupJobResponseBody) Validate() error {
	return dara.Validate(s)
}
