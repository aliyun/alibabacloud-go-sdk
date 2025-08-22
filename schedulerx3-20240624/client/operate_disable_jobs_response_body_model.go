// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateDisableJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateDisableJobsResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateDisableJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateDisableJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateDisableJobsResponseBody
	GetSuccess() *bool
}

type OperateDisableJobsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 29ED6209-5DE6-5E1D-89B0-B7B1D823A1BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateDisableJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateDisableJobsResponseBody) GoString() string {
	return s.String()
}

func (s *OperateDisableJobsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateDisableJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateDisableJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateDisableJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateDisableJobsResponseBody) SetCode(v int32) *OperateDisableJobsResponseBody {
	s.Code = &v
	return s
}

func (s *OperateDisableJobsResponseBody) SetMessage(v string) *OperateDisableJobsResponseBody {
	s.Message = &v
	return s
}

func (s *OperateDisableJobsResponseBody) SetRequestId(v string) *OperateDisableJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateDisableJobsResponseBody) SetSuccess(v bool) *OperateDisableJobsResponseBody {
	s.Success = &v
	return s
}

func (s *OperateDisableJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
