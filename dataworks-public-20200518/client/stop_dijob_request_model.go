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
}

type StopDIJobRequest struct {
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11668
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
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

func (s *StopDIJobRequest) SetDIJobId(v int64) *StopDIJobRequest {
	s.DIJobId = &v
	return s
}

func (s *StopDIJobRequest) Validate() error {
	return dara.Validate(s)
}
