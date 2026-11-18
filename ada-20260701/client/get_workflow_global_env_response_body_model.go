// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowGlobalEnvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *GetWorkflowGlobalEnvResponseBody
	GetContent() *string
	SetExists(v bool) *GetWorkflowGlobalEnvResponseBody
	GetExists() *bool
	SetRequestId(v string) *GetWorkflowGlobalEnvResponseBody
	GetRequestId() *string
	SetSize(v int64) *GetWorkflowGlobalEnvResponseBody
	GetSize() *int64
	SetSuccess(v bool) *GetWorkflowGlobalEnvResponseBody
	GetSuccess() *bool
	SetUpdatedAt(v int64) *GetWorkflowGlobalEnvResponseBody
	GetUpdatedAt() *int64
}

type GetWorkflowGlobalEnvResponseBody struct {
	// The configuration content.
	//
	// This parameter is required.
	//
	// example:
	//
	// MODE=demo
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Indicates whether the configuration is saved.
	//
	// This parameter is required.
	Exists *bool `json:"Exists,omitempty" xml:"Exists,omitempty"`
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
	// The timestamp of the last modification in milliseconds.
	//
	// example:
	//
	// 1700000000000
	UpdatedAt *int64 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s GetWorkflowGlobalEnvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowGlobalEnvResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowGlobalEnvResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetWorkflowGlobalEnvResponseBody) GetExists() *bool {
	return s.Exists
}

func (s *GetWorkflowGlobalEnvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkflowGlobalEnvResponseBody) GetSize() *int64 {
	return s.Size
}

func (s *GetWorkflowGlobalEnvResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkflowGlobalEnvResponseBody) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *GetWorkflowGlobalEnvResponseBody) SetContent(v string) *GetWorkflowGlobalEnvResponseBody {
	s.Content = &v
	return s
}

func (s *GetWorkflowGlobalEnvResponseBody) SetExists(v bool) *GetWorkflowGlobalEnvResponseBody {
	s.Exists = &v
	return s
}

func (s *GetWorkflowGlobalEnvResponseBody) SetRequestId(v string) *GetWorkflowGlobalEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowGlobalEnvResponseBody) SetSize(v int64) *GetWorkflowGlobalEnvResponseBody {
	s.Size = &v
	return s
}

func (s *GetWorkflowGlobalEnvResponseBody) SetSuccess(v bool) *GetWorkflowGlobalEnvResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkflowGlobalEnvResponseBody) SetUpdatedAt(v int64) *GetWorkflowGlobalEnvResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *GetWorkflowGlobalEnvResponseBody) Validate() error {
	return dara.Validate(s)
}
