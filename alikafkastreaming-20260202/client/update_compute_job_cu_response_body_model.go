// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeJobCuResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *UpdateComputeJobCuResponseBody
	GetCode() *int64
	SetData(v bool) *UpdateComputeJobCuResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateComputeJobCuResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateComputeJobCuResponseBody
	GetSuccess() *bool
}

type UpdateComputeJobCuResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateComputeJobCuResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeJobCuResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeJobCuResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateComputeJobCuResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateComputeJobCuResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateComputeJobCuResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateComputeJobCuResponseBody) SetCode(v int64) *UpdateComputeJobCuResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateComputeJobCuResponseBody) SetData(v bool) *UpdateComputeJobCuResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateComputeJobCuResponseBody) SetRequestId(v string) *UpdateComputeJobCuResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateComputeJobCuResponseBody) SetSuccess(v bool) *UpdateComputeJobCuResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateComputeJobCuResponseBody) Validate() error {
	return dara.Validate(s)
}
