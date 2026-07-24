// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReDoRoutineBuildResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPipeLineRunId(v int64) *ReDoRoutineBuildResponseBody
	GetPipeLineRunId() *int64
	SetRequestId(v string) *ReDoRoutineBuildResponseBody
	GetRequestId() *string
	SetRoutineBuildId(v int64) *ReDoRoutineBuildResponseBody
	GetRoutineBuildId() *int64
}

type ReDoRoutineBuildResponseBody struct {
	// The ID of the build task in Yunxiao.
	//
	// example:
	//
	// 70
	PipeLineRunId *int64 `json:"PipeLineRunId,omitempty" xml:"PipeLineRunId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F8AA0364-0FDB-4AD5-AC74-D69FAB8924ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the ER build task.
	//
	// example:
	//
	// 156773519472872
	RoutineBuildId *int64 `json:"RoutineBuildId,omitempty" xml:"RoutineBuildId,omitempty"`
}

func (s ReDoRoutineBuildResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReDoRoutineBuildResponseBody) GoString() string {
	return s.String()
}

func (s *ReDoRoutineBuildResponseBody) GetPipeLineRunId() *int64 {
	return s.PipeLineRunId
}

func (s *ReDoRoutineBuildResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReDoRoutineBuildResponseBody) GetRoutineBuildId() *int64 {
	return s.RoutineBuildId
}

func (s *ReDoRoutineBuildResponseBody) SetPipeLineRunId(v int64) *ReDoRoutineBuildResponseBody {
	s.PipeLineRunId = &v
	return s
}

func (s *ReDoRoutineBuildResponseBody) SetRequestId(v string) *ReDoRoutineBuildResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReDoRoutineBuildResponseBody) SetRoutineBuildId(v int64) *ReDoRoutineBuildResponseBody {
	s.RoutineBuildId = &v
	return s
}

func (s *ReDoRoutineBuildResponseBody) Validate() error {
	return dara.Validate(s)
}
