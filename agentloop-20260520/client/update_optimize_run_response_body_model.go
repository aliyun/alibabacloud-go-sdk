// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOptimizeRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOptimizeRunId(v string) *UpdateOptimizeRunResponseBody
	GetOptimizeRunId() *string
	SetRequestId(v string) *UpdateOptimizeRunResponseBody
	GetRequestId() *string
}

type UpdateOptimizeRunResponseBody struct {
	// The ID of the execution instance.
	//
	// example:
	//
	// or-20260812-0001
	OptimizeRunId *string `json:"optimizeRunId,omitempty" xml:"optimizeRunId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateOptimizeRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOptimizeRunResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOptimizeRunResponseBody) GetOptimizeRunId() *string {
	return s.OptimizeRunId
}

func (s *UpdateOptimizeRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOptimizeRunResponseBody) SetOptimizeRunId(v string) *UpdateOptimizeRunResponseBody {
	s.OptimizeRunId = &v
	return s
}

func (s *UpdateOptimizeRunResponseBody) SetRequestId(v string) *UpdateOptimizeRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOptimizeRunResponseBody) Validate() error {
	return dara.Validate(s)
}
