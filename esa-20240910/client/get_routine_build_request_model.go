// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineBuildRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoutineBuildId(v int64) *GetRoutineBuildRequest
	GetRoutineBuildId() *int64
}

type GetRoutineBuildRequest struct {
	// The ID of the ER build task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4354306271619136
	RoutineBuildId *int64 `json:"RoutineBuildId,omitempty" xml:"RoutineBuildId,omitempty"`
}

func (s GetRoutineBuildRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineBuildRequest) GoString() string {
	return s.String()
}

func (s *GetRoutineBuildRequest) GetRoutineBuildId() *int64 {
	return s.RoutineBuildId
}

func (s *GetRoutineBuildRequest) SetRoutineBuildId(v int64) *GetRoutineBuildRequest {
	s.RoutineBuildId = &v
	return s
}

func (s *GetRoutineBuildRequest) Validate() error {
	return dara.Validate(s)
}
