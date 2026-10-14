// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOptimizeRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOptimizeRun(v map[string]interface{}) *CreateOptimizeRunResponseBody
	GetOptimizeRun() map[string]interface{}
	SetOptimizeRunId(v string) *CreateOptimizeRunResponseBody
	GetOptimizeRunId() *string
	SetRequestId(v string) *CreateOptimizeRunResponseBody
	GetRequestId() *string
}

type CreateOptimizeRunResponseBody struct {
	// The details of the created run. You can call the GetOptimizeRun operation to query the subsequent status.
	//
	// example:
	//
	// {"optimizeRunId":"opr-20261014","status":"Pending"}
	OptimizeRun map[string]interface{} `json:"optimizeRun,omitempty" xml:"optimizeRun,omitempty"`
	// The run ID.
	//
	// example:
	//
	// opr-20261014
	OptimizeRunId *string `json:"optimizeRunId,omitempty" xml:"optimizeRunId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateOptimizeRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOptimizeRunResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOptimizeRunResponseBody) GetOptimizeRun() map[string]interface{} {
	return s.OptimizeRun
}

func (s *CreateOptimizeRunResponseBody) GetOptimizeRunId() *string {
	return s.OptimizeRunId
}

func (s *CreateOptimizeRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOptimizeRunResponseBody) SetOptimizeRun(v map[string]interface{}) *CreateOptimizeRunResponseBody {
	s.OptimizeRun = v
	return s
}

func (s *CreateOptimizeRunResponseBody) SetOptimizeRunId(v string) *CreateOptimizeRunResponseBody {
	s.OptimizeRunId = &v
	return s
}

func (s *CreateOptimizeRunResponseBody) SetRequestId(v string) *CreateOptimizeRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOptimizeRunResponseBody) Validate() error {
	return dara.Validate(s)
}
