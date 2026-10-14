// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOptimizeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOptimizeTaskName(v string) *CreateOptimizeTaskResponseBody
	GetOptimizeTaskName() *string
	SetRequestId(v string) *CreateOptimizeTaskResponseBody
	GetRequestId() *string
}

type CreateOptimizeTaskResponseBody struct {
	// The name of the optimization task.
	//
	// example:
	//
	// efficiency-analysis
	OptimizeTaskName *string `json:"optimizeTaskName,omitempty" xml:"optimizeTaskName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateOptimizeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOptimizeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOptimizeTaskResponseBody) GetOptimizeTaskName() *string {
	return s.OptimizeTaskName
}

func (s *CreateOptimizeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOptimizeTaskResponseBody) SetOptimizeTaskName(v string) *CreateOptimizeTaskResponseBody {
	s.OptimizeTaskName = &v
	return s
}

func (s *CreateOptimizeTaskResponseBody) SetRequestId(v string) *CreateOptimizeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOptimizeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
