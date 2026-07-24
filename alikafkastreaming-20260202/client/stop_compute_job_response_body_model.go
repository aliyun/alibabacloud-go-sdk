// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopComputeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *StopComputeJobResponseBody
	GetCode() *int64
	SetData(v bool) *StopComputeJobResponseBody
	GetData() *bool
	SetRequestId(v string) *StopComputeJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopComputeJobResponseBody
	GetSuccess() *bool
}

type StopComputeJobResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopComputeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopComputeJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopComputeJobResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *StopComputeJobResponseBody) GetData() *bool {
	return s.Data
}

func (s *StopComputeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopComputeJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopComputeJobResponseBody) SetCode(v int64) *StopComputeJobResponseBody {
	s.Code = &v
	return s
}

func (s *StopComputeJobResponseBody) SetData(v bool) *StopComputeJobResponseBody {
	s.Data = &v
	return s
}

func (s *StopComputeJobResponseBody) SetRequestId(v string) *StopComputeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopComputeJobResponseBody) SetSuccess(v bool) *StopComputeJobResponseBody {
	s.Success = &v
	return s
}

func (s *StopComputeJobResponseBody) Validate() error {
	return dara.Validate(s)
}
