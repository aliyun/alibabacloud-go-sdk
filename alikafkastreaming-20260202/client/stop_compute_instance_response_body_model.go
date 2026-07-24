// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopComputeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *StopComputeInstanceResponseBody
	GetCode() *int64
	SetData(v bool) *StopComputeInstanceResponseBody
	GetData() *bool
	SetRequestId(v string) *StopComputeInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopComputeInstanceResponseBody
	GetSuccess() *bool
}

type StopComputeInstanceResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopComputeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopComputeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopComputeInstanceResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *StopComputeInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *StopComputeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopComputeInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopComputeInstanceResponseBody) SetCode(v int64) *StopComputeInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StopComputeInstanceResponseBody) SetData(v bool) *StopComputeInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *StopComputeInstanceResponseBody) SetRequestId(v string) *StopComputeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopComputeInstanceResponseBody) SetSuccess(v bool) *StopComputeInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *StopComputeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
