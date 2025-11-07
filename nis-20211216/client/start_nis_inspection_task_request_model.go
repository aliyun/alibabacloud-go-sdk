// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartNisInspectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInspectionTaskId(v string) *StartNisInspectionTaskRequest
	GetInspectionTaskId() *string
}

type StartNisInspectionTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ni-8svmpe0yso****r7fh79
	InspectionTaskId *string `json:"InspectionTaskId,omitempty" xml:"InspectionTaskId,omitempty"`
}

func (s StartNisInspectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartNisInspectionTaskRequest) GoString() string {
	return s.String()
}

func (s *StartNisInspectionTaskRequest) GetInspectionTaskId() *string {
	return s.InspectionTaskId
}

func (s *StartNisInspectionTaskRequest) SetInspectionTaskId(v string) *StartNisInspectionTaskRequest {
	s.InspectionTaskId = &v
	return s
}

func (s *StartNisInspectionTaskRequest) Validate() error {
	return dara.Validate(s)
}
