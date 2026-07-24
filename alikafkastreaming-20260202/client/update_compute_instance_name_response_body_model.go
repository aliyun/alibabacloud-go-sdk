// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeInstanceNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *UpdateComputeInstanceNameResponseBody
	GetCode() *int64
	SetData(v bool) *UpdateComputeInstanceNameResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateComputeInstanceNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateComputeInstanceNameResponseBody
	GetSuccess() *bool
}

type UpdateComputeInstanceNameResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateComputeInstanceNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeInstanceNameResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateComputeInstanceNameResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateComputeInstanceNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateComputeInstanceNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateComputeInstanceNameResponseBody) SetCode(v int64) *UpdateComputeInstanceNameResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateComputeInstanceNameResponseBody) SetData(v bool) *UpdateComputeInstanceNameResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateComputeInstanceNameResponseBody) SetRequestId(v string) *UpdateComputeInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateComputeInstanceNameResponseBody) SetSuccess(v bool) *UpdateComputeInstanceNameResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateComputeInstanceNameResponseBody) Validate() error {
	return dara.Validate(s)
}
