// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartComputeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *RestartComputeJobResponseBody
	GetCode() *int64
	SetData(v bool) *RestartComputeJobResponseBody
	GetData() *bool
	SetRequestId(v string) *RestartComputeJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestartComputeJobResponseBody
	GetSuccess() *bool
}

type RestartComputeJobResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestartComputeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartComputeJobResponseBody) GoString() string {
	return s.String()
}

func (s *RestartComputeJobResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *RestartComputeJobResponseBody) GetData() *bool {
	return s.Data
}

func (s *RestartComputeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartComputeJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestartComputeJobResponseBody) SetCode(v int64) *RestartComputeJobResponseBody {
	s.Code = &v
	return s
}

func (s *RestartComputeJobResponseBody) SetData(v bool) *RestartComputeJobResponseBody {
	s.Data = &v
	return s
}

func (s *RestartComputeJobResponseBody) SetRequestId(v string) *RestartComputeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartComputeJobResponseBody) SetSuccess(v bool) *RestartComputeJobResponseBody {
	s.Success = &v
	return s
}

func (s *RestartComputeJobResponseBody) Validate() error {
	return dara.Validate(s)
}
