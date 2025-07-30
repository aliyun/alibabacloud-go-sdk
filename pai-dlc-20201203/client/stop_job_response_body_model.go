// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *StopJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *StopJobResponseBody
	GetRequestId() *string
}

type StopJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// dlc-20210126170216-xxxxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-xxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *StopJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopJobResponseBody) SetJobId(v string) *StopJobResponseBody {
	s.JobId = &v
	return s
}

func (s *StopJobResponseBody) SetRequestId(v string) *StopJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopJobResponseBody) Validate() error {
	return dara.Validate(s)
}
