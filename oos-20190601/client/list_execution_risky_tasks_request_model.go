// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutionRiskyTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListExecutionRiskyTasksRequest
	GetRegionId() *string
	SetTemplateName(v string) *ListExecutionRiskyTasksRequest
	GetTemplateName() *string
}

type ListExecutionRiskyTasksRequest struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// myTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListExecutionRiskyTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionRiskyTasksRequest) GoString() string {
	return s.String()
}

func (s *ListExecutionRiskyTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListExecutionRiskyTasksRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListExecutionRiskyTasksRequest) SetRegionId(v string) *ListExecutionRiskyTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ListExecutionRiskyTasksRequest) SetTemplateName(v string) *ListExecutionRiskyTasksRequest {
	s.TemplateName = &v
	return s
}

func (s *ListExecutionRiskyTasksRequest) Validate() error {
	return dara.Validate(s)
}
