// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunUserPoolSyncJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RunUserPoolSyncJobResponseBody
	GetRequestId() *string
	SetSynchronizationJobId(v string) *RunUserPoolSyncJobResponseBody
	GetSynchronizationJobId() *string
}

type RunUserPoolSyncJobResponseBody struct {
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s RunUserPoolSyncJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunUserPoolSyncJobResponseBody) GoString() string {
	return s.String()
}

func (s *RunUserPoolSyncJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunUserPoolSyncJobResponseBody) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *RunUserPoolSyncJobResponseBody) SetRequestId(v string) *RunUserPoolSyncJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunUserPoolSyncJobResponseBody) SetSynchronizationJobId(v string) *RunUserPoolSyncJobResponseBody {
	s.SynchronizationJobId = &v
	return s
}

func (s *RunUserPoolSyncJobResponseBody) Validate() error {
	return dara.Validate(s)
}
