// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDIJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *StopDIJobRequest
	GetDIJobId() *int64
	SetId(v int64) *StopDIJobRequest
	GetId() *int64
	SetInstanceId(v int64) *StopDIJobRequest
	GetInstanceId() *int64
}

type StopDIJobRequest struct {
	// Deprecated
	//
	// This parameter is deprecated and is replaced by the Id parameter.
	//
	// example:
	//
	// 11668
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// 11668
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 1234
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopDIJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDIJobRequest) GoString() string {
	return s.String()
}

func (s *StopDIJobRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *StopDIJobRequest) GetId() *int64 {
	return s.Id
}

func (s *StopDIJobRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *StopDIJobRequest) SetDIJobId(v int64) *StopDIJobRequest {
	s.DIJobId = &v
	return s
}

func (s *StopDIJobRequest) SetId(v int64) *StopDIJobRequest {
	s.Id = &v
	return s
}

func (s *StopDIJobRequest) SetInstanceId(v int64) *StopDIJobRequest {
	s.InstanceId = &v
	return s
}

func (s *StopDIJobRequest) Validate() error {
	return dara.Validate(s)
}
