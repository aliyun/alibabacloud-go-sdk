// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeleteComputeInstanceResponseBody
	GetCode() *int64
	SetData(v bool) *DeleteComputeInstanceResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteComputeInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteComputeInstanceResponseBody
	GetSuccess() *bool
}

type DeleteComputeInstanceResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteComputeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComputeInstanceResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeleteComputeInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteComputeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteComputeInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteComputeInstanceResponseBody) SetCode(v int64) *DeleteComputeInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteComputeInstanceResponseBody) SetData(v bool) *DeleteComputeInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteComputeInstanceResponseBody) SetRequestId(v string) *DeleteComputeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteComputeInstanceResponseBody) SetSuccess(v bool) *DeleteComputeInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteComputeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
