// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmsDbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDstName(v string) *UpdateMmsDbRequest
	GetDstName() *string
	SetDstProjectName(v string) *UpdateMmsDbRequest
	GetDstProjectName() *string
	SetStatus(v string) *UpdateMmsDbRequest
	GetStatus() *string
}

type UpdateMmsDbRequest struct {
	// The name of the destination schema in MaxCompute.
	//
	// example:
	//
	// default
	DstName *string `json:"dstName,omitempty" xml:"dstName,omitempty"`
	// The name of the destination MaxCompute project.
	//
	// example:
	//
	// dst_project_name
	DstProjectName *string `json:"dstProjectName,omitempty" xml:"dstProjectName,omitempty"`
	// The migration status.
	//
	// example:
	//
	// INIT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateMmsDbRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmsDbRequest) GoString() string {
	return s.String()
}

func (s *UpdateMmsDbRequest) GetDstName() *string {
	return s.DstName
}

func (s *UpdateMmsDbRequest) GetDstProjectName() *string {
	return s.DstProjectName
}

func (s *UpdateMmsDbRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateMmsDbRequest) SetDstName(v string) *UpdateMmsDbRequest {
	s.DstName = &v
	return s
}

func (s *UpdateMmsDbRequest) SetDstProjectName(v string) *UpdateMmsDbRequest {
	s.DstProjectName = &v
	return s
}

func (s *UpdateMmsDbRequest) SetStatus(v string) *UpdateMmsDbRequest {
	s.Status = &v
	return s
}

func (s *UpdateMmsDbRequest) Validate() error {
	return dara.Validate(s)
}
