// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSynchronizationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RunSynchronizationJobResponseBody
	GetRequestId() *string
	SetSynchronizationJobId(v string) *RunSynchronizationJobResponseBody
	GetSynchronizationJobId() *string
}

type RunSynchronizationJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the synchronization job.
	//
	// example:
	//
	// sync_0000347vjovtcf41li0fgsd98gn24q9nj9xxxxx
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s RunSynchronizationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *RunSynchronizationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunSynchronizationJobResponseBody) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *RunSynchronizationJobResponseBody) SetRequestId(v string) *RunSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunSynchronizationJobResponseBody) SetSynchronizationJobId(v string) *RunSynchronizationJobResponseBody {
	s.SynchronizationJobId = &v
	return s
}

func (s *RunSynchronizationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
