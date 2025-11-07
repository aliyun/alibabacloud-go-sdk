// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNisInspectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInspectionTaskId(v string) *UpdateNisInspectionTaskRequest
	GetInspectionTaskId() *string
	SetStatus(v string) *UpdateNisInspectionTaskRequest
	GetStatus() *string
}

type UpdateNisInspectionTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ni-8svmpe0yso2bhzr7fh79
	InspectionTaskId *string `json:"InspectionTaskId,omitempty" xml:"InspectionTaskId,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateNisInspectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNisInspectionTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateNisInspectionTaskRequest) GetInspectionTaskId() *string {
	return s.InspectionTaskId
}

func (s *UpdateNisInspectionTaskRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateNisInspectionTaskRequest) SetInspectionTaskId(v string) *UpdateNisInspectionTaskRequest {
	s.InspectionTaskId = &v
	return s
}

func (s *UpdateNisInspectionTaskRequest) SetStatus(v string) *UpdateNisInspectionTaskRequest {
	s.Status = &v
	return s
}

func (s *UpdateNisInspectionTaskRequest) Validate() error {
	return dara.Validate(s)
}
