// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportAgentJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ImportAgentJobsResponseBody
	GetCode() *int32
	SetMessage(v string) *ImportAgentJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportAgentJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportAgentJobsResponseBody
	GetSuccess() *bool
}

type ImportAgentJobsResponseBody struct {
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
	// 2C3E52FF-CBE9-5C0E-8252-37ACFF1F5EFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportAgentJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportAgentJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ImportAgentJobsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ImportAgentJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportAgentJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportAgentJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportAgentJobsResponseBody) SetCode(v int32) *ImportAgentJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ImportAgentJobsResponseBody) SetMessage(v string) *ImportAgentJobsResponseBody {
	s.Message = &v
	return s
}

func (s *ImportAgentJobsResponseBody) SetRequestId(v string) *ImportAgentJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportAgentJobsResponseBody) SetSuccess(v bool) *ImportAgentJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ImportAgentJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
