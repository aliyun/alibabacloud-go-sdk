// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetJobLogResponseBody
	GetJobId() *string
	SetRequestId(v string) *GetJobLogResponseBody
	GetRequestId() *string
	SetStderrLog(v string) *GetJobLogResponseBody
	GetStderrLog() *string
	SetStderrLogSize(v string) *GetJobLogResponseBody
	GetStderrLogSize() *string
	SetStdoutLog(v string) *GetJobLogResponseBody
	GetStdoutLog() *string
	SetStdoutLogSize(v string) *GetJobLogResponseBody
	GetStdoutLogSize() *string
	SetSuccess(v string) *GetJobLogResponseBody
	GetSuccess() *string
}

type GetJobLogResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// 1.manager
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B745C159-3155-4B94-95D0-4B73D4D2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error output log that is encoded in Base64.
	//
	// example:
	//
	// aG9zdG5hbWU=
	StderrLog *string `json:"StderrLog,omitempty" xml:"StderrLog,omitempty"`
	// The size of the error output file.
	//
	// example:
	//
	// 0
	StderrLogSize *string `json:"StderrLogSize,omitempty" xml:"StderrLogSize,omitempty"`
	// The standard output log that is encoded in Base64.
	//
	// example:
	//
	// aG9zdG5hbWU=
	StdoutLog *string `json:"StdoutLog,omitempty" xml:"StdoutLog,omitempty"`
	// The size of the standard output file.
	//
	// example:
	//
	// 4096
	StdoutLogSize *string `json:"StdoutLogSize,omitempty" xml:"StdoutLogSize,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobLogResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetJobLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobLogResponseBody) GetStderrLog() *string {
	return s.StderrLog
}

func (s *GetJobLogResponseBody) GetStderrLogSize() *string {
	return s.StderrLogSize
}

func (s *GetJobLogResponseBody) GetStdoutLog() *string {
	return s.StdoutLog
}

func (s *GetJobLogResponseBody) GetStdoutLogSize() *string {
	return s.StdoutLogSize
}

func (s *GetJobLogResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetJobLogResponseBody) SetJobId(v string) *GetJobLogResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobLogResponseBody) SetRequestId(v string) *GetJobLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobLogResponseBody) SetStderrLog(v string) *GetJobLogResponseBody {
	s.StderrLog = &v
	return s
}

func (s *GetJobLogResponseBody) SetStderrLogSize(v string) *GetJobLogResponseBody {
	s.StderrLogSize = &v
	return s
}

func (s *GetJobLogResponseBody) SetStdoutLog(v string) *GetJobLogResponseBody {
	s.StdoutLog = &v
	return s
}

func (s *GetJobLogResponseBody) SetStdoutLogSize(v string) *GetJobLogResponseBody {
	s.StdoutLogSize = &v
	return s
}

func (s *GetJobLogResponseBody) SetSuccess(v string) *GetJobLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetJobLogResponseBody) Validate() error {
	return dara.Validate(s)
}
