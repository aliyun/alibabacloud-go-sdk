// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReopenComputeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ReopenComputeInstanceResponseBody
	GetCode() *int64
	SetData(v bool) *ReopenComputeInstanceResponseBody
	GetData() *bool
	SetRequestId(v string) *ReopenComputeInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReopenComputeInstanceResponseBody
	GetSuccess() *bool
}

type ReopenComputeInstanceResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReopenComputeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReopenComputeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReopenComputeInstanceResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ReopenComputeInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *ReopenComputeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReopenComputeInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReopenComputeInstanceResponseBody) SetCode(v int64) *ReopenComputeInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ReopenComputeInstanceResponseBody) SetData(v bool) *ReopenComputeInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *ReopenComputeInstanceResponseBody) SetRequestId(v string) *ReopenComputeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReopenComputeInstanceResponseBody) SetSuccess(v bool) *ReopenComputeInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ReopenComputeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
