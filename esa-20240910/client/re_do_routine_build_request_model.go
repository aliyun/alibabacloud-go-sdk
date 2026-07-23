// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReDoRoutineBuildRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoutineBuildId(v int64) *ReDoRoutineBuildRequest
	GetRoutineBuildId() *int64
}

type ReDoRoutineBuildRequest struct {
	// The ID of the ER build task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 156773519472872
	RoutineBuildId *int64 `json:"RoutineBuildId,omitempty" xml:"RoutineBuildId,omitempty"`
}

func (s ReDoRoutineBuildRequest) String() string {
	return dara.Prettify(s)
}

func (s ReDoRoutineBuildRequest) GoString() string {
	return s.String()
}

func (s *ReDoRoutineBuildRequest) GetRoutineBuildId() *int64 {
	return s.RoutineBuildId
}

func (s *ReDoRoutineBuildRequest) SetRoutineBuildId(v int64) *ReDoRoutineBuildRequest {
	s.RoutineBuildId = &v
	return s
}

func (s *ReDoRoutineBuildRequest) Validate() error {
	return dara.Validate(s)
}
