// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateComputeJobResponseBody
	GetCode() *int64
	SetData(v bool) *CreateComputeJobResponseBody
	GetData() *bool
	SetRequestId(v string) *CreateComputeJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateComputeJobResponseBody
	GetSuccess() *bool
}

type CreateComputeJobResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateComputeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComputeJobResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateComputeJobResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateComputeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateComputeJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateComputeJobResponseBody) SetCode(v int64) *CreateComputeJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateComputeJobResponseBody) SetData(v bool) *CreateComputeJobResponseBody {
	s.Data = &v
	return s
}

func (s *CreateComputeJobResponseBody) SetRequestId(v string) *CreateComputeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateComputeJobResponseBody) SetSuccess(v bool) *CreateComputeJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateComputeJobResponseBody) Validate() error {
	return dara.Validate(s)
}
