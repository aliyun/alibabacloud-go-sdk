// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeJobDraftSqlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *UpdateComputeJobDraftSqlResponseBody
	GetCode() *int64
	SetData(v bool) *UpdateComputeJobDraftSqlResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateComputeJobDraftSqlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateComputeJobDraftSqlResponseBody
	GetSuccess() *bool
}

type UpdateComputeJobDraftSqlResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateComputeJobDraftSqlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeJobDraftSqlResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeJobDraftSqlResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateComputeJobDraftSqlResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateComputeJobDraftSqlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateComputeJobDraftSqlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateComputeJobDraftSqlResponseBody) SetCode(v int64) *UpdateComputeJobDraftSqlResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateComputeJobDraftSqlResponseBody) SetData(v bool) *UpdateComputeJobDraftSqlResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateComputeJobDraftSqlResponseBody) SetRequestId(v string) *UpdateComputeJobDraftSqlResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateComputeJobDraftSqlResponseBody) SetSuccess(v bool) *UpdateComputeJobDraftSqlResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateComputeJobDraftSqlResponseBody) Validate() error {
	return dara.Validate(s)
}
