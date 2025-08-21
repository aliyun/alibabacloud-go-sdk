// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ListApmRequest
	GetDescription() *string
	SetInstanceId(v string) *ListApmRequest
	GetInstanceId() *string
	SetOutput(v string) *ListApmRequest
	GetOutput() *string
	SetPage(v int64) *ListApmRequest
	GetPage() *int64
	SetSize(v int64) *ListApmRequest
	GetSize() *int64
}

type ListApmRequest struct {
	// example:
	//
	// APMtest
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// apm-cn-i7m2fuae****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// es-cn-i7m2fsfhc001x****
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListApmRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApmRequest) GoString() string {
	return s.String()
}

func (s *ListApmRequest) GetDescription() *string {
	return s.Description
}

func (s *ListApmRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApmRequest) GetOutput() *string {
	return s.Output
}

func (s *ListApmRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListApmRequest) GetSize() *int64 {
	return s.Size
}

func (s *ListApmRequest) SetDescription(v string) *ListApmRequest {
	s.Description = &v
	return s
}

func (s *ListApmRequest) SetInstanceId(v string) *ListApmRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApmRequest) SetOutput(v string) *ListApmRequest {
	s.Output = &v
	return s
}

func (s *ListApmRequest) SetPage(v int64) *ListApmRequest {
	s.Page = &v
	return s
}

func (s *ListApmRequest) SetSize(v int64) *ListApmRequest {
	s.Size = &v
	return s
}

func (s *ListApmRequest) Validate() error {
	return dara.Validate(s)
}
