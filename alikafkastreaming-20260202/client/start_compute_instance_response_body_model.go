// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartComputeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *StartComputeInstanceResponseBody
	GetCode() *int64
	SetData(v bool) *StartComputeInstanceResponseBody
	GetData() *bool
	SetRequestId(v string) *StartComputeInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartComputeInstanceResponseBody
	GetSuccess() *bool
}

type StartComputeInstanceResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartComputeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartComputeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartComputeInstanceResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *StartComputeInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *StartComputeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartComputeInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartComputeInstanceResponseBody) SetCode(v int64) *StartComputeInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StartComputeInstanceResponseBody) SetData(v bool) *StartComputeInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *StartComputeInstanceResponseBody) SetRequestId(v string) *StartComputeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartComputeInstanceResponseBody) SetSuccess(v bool) *StartComputeInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *StartComputeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
