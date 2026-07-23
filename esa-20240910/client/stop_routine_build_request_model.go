// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRoutineBuildRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoutineBuildId(v int64) *StopRoutineBuildRequest
	GetRoutineBuildId() *int64
}

type StopRoutineBuildRequest struct {
	// The ID of the ER build task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4195751945250944
	RoutineBuildId *int64 `json:"RoutineBuildId,omitempty" xml:"RoutineBuildId,omitempty"`
}

func (s StopRoutineBuildRequest) String() string {
	return dara.Prettify(s)
}

func (s StopRoutineBuildRequest) GoString() string {
	return s.String()
}

func (s *StopRoutineBuildRequest) GetRoutineBuildId() *int64 {
	return s.RoutineBuildId
}

func (s *StopRoutineBuildRequest) SetRoutineBuildId(v int64) *StopRoutineBuildRequest {
	s.RoutineBuildId = &v
	return s
}

func (s *StopRoutineBuildRequest) Validate() error {
	return dara.Validate(s)
}
