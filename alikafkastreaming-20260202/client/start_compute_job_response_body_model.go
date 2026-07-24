// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartComputeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *StartComputeJobResponseBody
	GetCode() *int64
	SetData(v bool) *StartComputeJobResponseBody
	GetData() *bool
	SetRequestId(v string) *StartComputeJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartComputeJobResponseBody
	GetSuccess() *bool
}

type StartComputeJobResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartComputeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartComputeJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartComputeJobResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *StartComputeJobResponseBody) GetData() *bool {
	return s.Data
}

func (s *StartComputeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartComputeJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartComputeJobResponseBody) SetCode(v int64) *StartComputeJobResponseBody {
	s.Code = &v
	return s
}

func (s *StartComputeJobResponseBody) SetData(v bool) *StartComputeJobResponseBody {
	s.Data = &v
	return s
}

func (s *StartComputeJobResponseBody) SetRequestId(v string) *StartComputeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartComputeJobResponseBody) SetSuccess(v bool) *StartComputeJobResponseBody {
	s.Success = &v
	return s
}

func (s *StartComputeJobResponseBody) Validate() error {
	return dara.Validate(s)
}
