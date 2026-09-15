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
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 062D8E8B-8D47-5DCC-BB12-5A1D93C3A66B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
