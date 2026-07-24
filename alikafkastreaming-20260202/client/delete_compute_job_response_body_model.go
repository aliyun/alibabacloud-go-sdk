// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeleteComputeJobResponseBody
	GetCode() *int64
	SetData(v bool) *DeleteComputeJobResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteComputeJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteComputeJobResponseBody
	GetSuccess() *bool
}

type DeleteComputeJobResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteComputeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComputeJobResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeleteComputeJobResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteComputeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteComputeJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteComputeJobResponseBody) SetCode(v int64) *DeleteComputeJobResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteComputeJobResponseBody) SetData(v bool) *DeleteComputeJobResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteComputeJobResponseBody) SetRequestId(v string) *DeleteComputeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteComputeJobResponseBody) SetSuccess(v bool) *DeleteComputeJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteComputeJobResponseBody) Validate() error {
	return dara.Validate(s)
}
