// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunInventorySynchronizationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInventorySynchronizationJobId(v string) *RunInventorySynchronizationJobResponseBody
	GetInventorySynchronizationJobId() *string
	SetRequestId(v string) *RunInventorySynchronizationJobResponseBody
	GetRequestId() *string
}

type RunInventorySynchronizationJobResponseBody struct {
	// The ID of the synchronization job.
	//
	// example:
	//
	// sync_xxxxxxxxxx
	InventorySynchronizationJobId *string `json:"InventorySynchronizationJobId,omitempty" xml:"InventorySynchronizationJobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunInventorySynchronizationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunInventorySynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *RunInventorySynchronizationJobResponseBody) GetInventorySynchronizationJobId() *string {
	return s.InventorySynchronizationJobId
}

func (s *RunInventorySynchronizationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunInventorySynchronizationJobResponseBody) SetInventorySynchronizationJobId(v string) *RunInventorySynchronizationJobResponseBody {
	s.InventorySynchronizationJobId = &v
	return s
}

func (s *RunInventorySynchronizationJobResponseBody) SetRequestId(v string) *RunInventorySynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunInventorySynchronizationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
