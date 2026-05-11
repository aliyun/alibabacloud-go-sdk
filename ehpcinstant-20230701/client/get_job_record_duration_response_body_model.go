// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobRecordDurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobRecordDuration(v int64) *GetJobRecordDurationResponseBody
	GetJobRecordDuration() *int64
	SetRequestId(v string) *GetJobRecordDurationResponseBody
	GetRequestId() *string
}

type GetJobRecordDurationResponseBody struct {
	// example:
	//
	// 30
	JobRecordDuration *int64 `json:"JobRecordDuration,omitempty" xml:"JobRecordDuration,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobRecordDurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobRecordDurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobRecordDurationResponseBody) GetJobRecordDuration() *int64 {
	return s.JobRecordDuration
}

func (s *GetJobRecordDurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobRecordDurationResponseBody) SetJobRecordDuration(v int64) *GetJobRecordDurationResponseBody {
	s.JobRecordDuration = &v
	return s
}

func (s *GetJobRecordDurationResponseBody) SetRequestId(v string) *GetJobRecordDurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobRecordDurationResponseBody) Validate() error {
	return dara.Validate(s)
}
