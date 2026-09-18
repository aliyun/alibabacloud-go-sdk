// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeBundleId(v int64) *CreateScanRequest
	GetCodeBundleId() *int64
	SetKind(v string) *CreateScanRequest
	GetKind() *string
	SetTaskName(v string) *CreateScanRequest
	GetTaskName() *string
}

type CreateScanRequest struct {
	// The code package ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 111
	CodeBundleId *int64 `json:"codeBundleId,omitempty" xml:"codeBundleId,omitempty"`
	// The type. Valid values:
	//
	// 	- full: full data
	//
	// 	- incremental: incremental
	//
	// This parameter is required.
	//
	// example:
	//
	// full
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The task name.
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s CreateScanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScanRequest) GoString() string {
	return s.String()
}

func (s *CreateScanRequest) GetCodeBundleId() *int64 {
	return s.CodeBundleId
}

func (s *CreateScanRequest) GetKind() *string {
	return s.Kind
}

func (s *CreateScanRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateScanRequest) SetCodeBundleId(v int64) *CreateScanRequest {
	s.CodeBundleId = &v
	return s
}

func (s *CreateScanRequest) SetKind(v string) *CreateScanRequest {
	s.Kind = &v
	return s
}

func (s *CreateScanRequest) SetTaskName(v string) *CreateScanRequest {
	s.TaskName = &v
	return s
}

func (s *CreateScanRequest) Validate() error {
	return dara.Validate(s)
}
