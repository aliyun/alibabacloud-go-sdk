// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEvaluatorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateEvaluatorResponseBody
	GetName() *string
	SetRequestId(v string) *CreateEvaluatorResponseBody
	GetRequestId() *string
	SetVersion(v string) *CreateEvaluatorResponseBody
	GetVersion() *string
}

type CreateEvaluatorResponseBody struct {
	// The evaluator name.
	//
	// example:
	//
	// trace_task_completion
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The version number that is created.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateEvaluatorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEvaluatorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEvaluatorResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateEvaluatorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEvaluatorResponseBody) GetVersion() *string {
	return s.Version
}

func (s *CreateEvaluatorResponseBody) SetName(v string) *CreateEvaluatorResponseBody {
	s.Name = &v
	return s
}

func (s *CreateEvaluatorResponseBody) SetRequestId(v string) *CreateEvaluatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEvaluatorResponseBody) SetVersion(v string) *CreateEvaluatorResponseBody {
	s.Version = &v
	return s
}

func (s *CreateEvaluatorResponseBody) Validate() error {
	return dara.Validate(s)
}
