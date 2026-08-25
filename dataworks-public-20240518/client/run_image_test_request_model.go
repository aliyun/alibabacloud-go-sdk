// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunImageTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCu(v float64) *RunImageTestRequest
	GetCu() *float64
	SetId(v string) *RunImageTestRequest
	GetId() *string
	SetProcessId(v string) *RunImageTestRequest
	GetProcessId() *string
	SetResourceGroupId(v string) *RunImageTestRequest
	GetResourceGroupId() *string
}

type RunImageTestRequest struct {
	// The test compute unit (CU).
	//
	// example:
	//
	// 0.5
	Cu *float64 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom_image_xxxx_xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The image test execution ID, which is used as an idempotence identifier.
	//
	// example:
	//
	// 582d4896-d224-413b-b883-239eeebe0bc5
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The unique identifier of the general-purpose resource group used to run the test task. Only Serverless resource groups are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_xxx_xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s RunImageTestRequest) String() string {
	return dara.Prettify(s)
}

func (s RunImageTestRequest) GoString() string {
	return s.String()
}

func (s *RunImageTestRequest) GetCu() *float64 {
	return s.Cu
}

func (s *RunImageTestRequest) GetId() *string {
	return s.Id
}

func (s *RunImageTestRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *RunImageTestRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RunImageTestRequest) SetCu(v float64) *RunImageTestRequest {
	s.Cu = &v
	return s
}

func (s *RunImageTestRequest) SetId(v string) *RunImageTestRequest {
	s.Id = &v
	return s
}

func (s *RunImageTestRequest) SetProcessId(v string) *RunImageTestRequest {
	s.ProcessId = &v
	return s
}

func (s *RunImageTestRequest) SetResourceGroupId(v string) *RunImageTestRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RunImageTestRequest) Validate() error {
	return dara.Validate(s)
}
