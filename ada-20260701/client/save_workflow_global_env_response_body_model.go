// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWorkflowGlobalEnvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveWorkflowGlobalEnvResponseBody
	GetRequestId() *string
	SetSize(v int64) *SaveWorkflowGlobalEnvResponseBody
	GetSize() *int64
	SetSuccess(v bool) *SaveWorkflowGlobalEnvResponseBody
	GetSuccess() *bool
}

type SaveWorkflowGlobalEnvResponseBody struct {
	// The request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 01A0677A-49CB-1605-AC3A-26761045B614
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of UTF-8 bytes in the configuration content.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// Indicates whether the request is successful.
	//
	// This parameter is required.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveWorkflowGlobalEnvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveWorkflowGlobalEnvResponseBody) GoString() string {
	return s.String()
}

func (s *SaveWorkflowGlobalEnvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveWorkflowGlobalEnvResponseBody) GetSize() *int64 {
	return s.Size
}

func (s *SaveWorkflowGlobalEnvResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveWorkflowGlobalEnvResponseBody) SetRequestId(v string) *SaveWorkflowGlobalEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveWorkflowGlobalEnvResponseBody) SetSize(v int64) *SaveWorkflowGlobalEnvResponseBody {
	s.Size = &v
	return s
}

func (s *SaveWorkflowGlobalEnvResponseBody) SetSuccess(v bool) *SaveWorkflowGlobalEnvResponseBody {
	s.Success = &v
	return s
}

func (s *SaveWorkflowGlobalEnvResponseBody) Validate() error {
	return dara.Validate(s)
}
