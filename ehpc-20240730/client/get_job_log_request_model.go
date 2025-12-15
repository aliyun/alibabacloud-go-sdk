// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetJobLogRequest
	GetClusterId() *string
	SetJobId(v string) *GetJobLogRequest
	GetJobId() *string
	SetLogType(v string) *GetJobLogRequest
	GetLogType() *string
	SetOffset(v string) *GetJobLogRequest
	GetOffset() *string
	SetSize(v string) *GetJobLogRequest
	GetSize() *string
}

type GetJobLogRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-jeJki6****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.manager
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The log type. Valid values:
	//
	// 	- stdout: standard output logs.
	//
	// 	- stderr: error output logs.
	//
	// 	- all: all logs.
	//
	// Default value: all.
	//
	// example:
	//
	// stdout
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The position where logs start to be read.
	//
	// Unit: bytes.
	//
	// Default value: 0.
	//
	// example:
	//
	// 0
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The maximum size of logs that you can read in a single request.
	//
	// Unit: bytes.
	//
	// Default value: 10240.
	//
	// example:
	//
	// 20480
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s GetJobLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobLogRequest) GoString() string {
	return s.String()
}

func (s *GetJobLogRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetJobLogRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetJobLogRequest) GetLogType() *string {
	return s.LogType
}

func (s *GetJobLogRequest) GetOffset() *string {
	return s.Offset
}

func (s *GetJobLogRequest) GetSize() *string {
	return s.Size
}

func (s *GetJobLogRequest) SetClusterId(v string) *GetJobLogRequest {
	s.ClusterId = &v
	return s
}

func (s *GetJobLogRequest) SetJobId(v string) *GetJobLogRequest {
	s.JobId = &v
	return s
}

func (s *GetJobLogRequest) SetLogType(v string) *GetJobLogRequest {
	s.LogType = &v
	return s
}

func (s *GetJobLogRequest) SetOffset(v string) *GetJobLogRequest {
	s.Offset = &v
	return s
}

func (s *GetJobLogRequest) SetSize(v string) *GetJobLogRequest {
	s.Size = &v
	return s
}

func (s *GetJobLogRequest) Validate() error {
	return dara.Validate(s)
}
