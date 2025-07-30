// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DeleteJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *DeleteJobResponseBody
	GetRequestId() *string
}

type DeleteJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// dlc*************
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *DeleteJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteJobResponseBody) SetJobId(v string) *DeleteJobResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteJobResponseBody) SetRequestId(v string) *DeleteJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteJobResponseBody) Validate() error {
	return dara.Validate(s)
}
