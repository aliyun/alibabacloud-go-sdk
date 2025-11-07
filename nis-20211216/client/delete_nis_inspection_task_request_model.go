// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNisInspectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInspectionTaskId(v string) *DeleteNisInspectionTaskRequest
	GetInspectionTaskId() *string
}

type DeleteNisInspectionTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ni-209d****wfirfwc2yl
	InspectionTaskId *string `json:"InspectionTaskId,omitempty" xml:"InspectionTaskId,omitempty"`
}

func (s DeleteNisInspectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNisInspectionTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteNisInspectionTaskRequest) GetInspectionTaskId() *string {
	return s.InspectionTaskId
}

func (s *DeleteNisInspectionTaskRequest) SetInspectionTaskId(v string) *DeleteNisInspectionTaskRequest {
	s.InspectionTaskId = &v
	return s
}

func (s *DeleteNisInspectionTaskRequest) Validate() error {
	return dara.Validate(s)
}
