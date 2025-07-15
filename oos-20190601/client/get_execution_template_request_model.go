// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecutionTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionId(v string) *GetExecutionTemplateRequest
	GetExecutionId() *string
	SetRegionId(v string) *GetExecutionTemplateRequest
	GetRegionId() *string
}

type GetExecutionTemplateRequest struct {
	// The ID of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-046490ff88f242
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetExecutionTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExecutionTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetExecutionTemplateRequest) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *GetExecutionTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetExecutionTemplateRequest) SetExecutionId(v string) *GetExecutionTemplateRequest {
	s.ExecutionId = &v
	return s
}

func (s *GetExecutionTemplateRequest) SetRegionId(v string) *GetExecutionTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *GetExecutionTemplateRequest) Validate() error {
	return dara.Validate(s)
}
