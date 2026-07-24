// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *UpdateComputeJobResponseBody
	GetCode() *int64
	SetData(v bool) *UpdateComputeJobResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateComputeJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateComputeJobResponseBody
	GetSuccess() *bool
}

type UpdateComputeJobResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateComputeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeJobResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateComputeJobResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateComputeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateComputeJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateComputeJobResponseBody) SetCode(v int64) *UpdateComputeJobResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateComputeJobResponseBody) SetData(v bool) *UpdateComputeJobResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateComputeJobResponseBody) SetRequestId(v string) *UpdateComputeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateComputeJobResponseBody) SetSuccess(v bool) *UpdateComputeJobResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateComputeJobResponseBody) Validate() error {
	return dara.Validate(s)
}
