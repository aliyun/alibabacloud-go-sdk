// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineBuildResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPipeLineRunId(v int64) *CreateRoutineBuildResponseBody
	GetPipeLineRunId() *int64
	SetRequestId(v string) *CreateRoutineBuildResponseBody
	GetRequestId() *string
	SetRoutineBuildId(v int64) *CreateRoutineBuildResponseBody
	GetRoutineBuildId() *int64
}

type CreateRoutineBuildResponseBody struct {
	// The build ID in Apsara Devops.
	//
	// example:
	//
	// 1
	PipeLineRunId *int64 `json:"PipeLineRunId,omitempty" xml:"PipeLineRunId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ER build task ID.
	//
	// example:
	//
	// 159782040838348
	RoutineBuildId *int64 `json:"RoutineBuildId,omitempty" xml:"RoutineBuildId,omitempty"`
}

func (s CreateRoutineBuildResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineBuildResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoutineBuildResponseBody) GetPipeLineRunId() *int64 {
	return s.PipeLineRunId
}

func (s *CreateRoutineBuildResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRoutineBuildResponseBody) GetRoutineBuildId() *int64 {
	return s.RoutineBuildId
}

func (s *CreateRoutineBuildResponseBody) SetPipeLineRunId(v int64) *CreateRoutineBuildResponseBody {
	s.PipeLineRunId = &v
	return s
}

func (s *CreateRoutineBuildResponseBody) SetRequestId(v string) *CreateRoutineBuildResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoutineBuildResponseBody) SetRoutineBuildId(v int64) *CreateRoutineBuildResponseBody {
	s.RoutineBuildId = &v
	return s
}

func (s *CreateRoutineBuildResponseBody) Validate() error {
	return dara.Validate(s)
}
