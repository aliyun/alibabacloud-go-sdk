// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmsTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDstName(v string) *UpdateMmsTableRequest
	GetDstName() *string
	SetDstProjectName(v string) *UpdateMmsTableRequest
	GetDstProjectName() *string
	SetDstSchemaName(v string) *UpdateMmsTableRequest
	GetDstSchemaName() *string
	SetStatus(v string) *UpdateMmsTableRequest
	GetStatus() *string
}

type UpdateMmsTableRequest struct {
	// example:
	//
	// dst_table_name
	DstName *string `json:"dstName,omitempty" xml:"dstName,omitempty"`
	// example:
	//
	// dst_project_name
	DstProjectName *string `json:"dstProjectName,omitempty" xml:"dstProjectName,omitempty"`
	// example:
	//
	// default
	DstSchemaName *string `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
	// example:
	//
	// INIT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateMmsTableRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmsTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateMmsTableRequest) GetDstName() *string {
	return s.DstName
}

func (s *UpdateMmsTableRequest) GetDstProjectName() *string {
	return s.DstProjectName
}

func (s *UpdateMmsTableRequest) GetDstSchemaName() *string {
	return s.DstSchemaName
}

func (s *UpdateMmsTableRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateMmsTableRequest) SetDstName(v string) *UpdateMmsTableRequest {
	s.DstName = &v
	return s
}

func (s *UpdateMmsTableRequest) SetDstProjectName(v string) *UpdateMmsTableRequest {
	s.DstProjectName = &v
	return s
}

func (s *UpdateMmsTableRequest) SetDstSchemaName(v string) *UpdateMmsTableRequest {
	s.DstSchemaName = &v
	return s
}

func (s *UpdateMmsTableRequest) SetStatus(v string) *UpdateMmsTableRequest {
	s.Status = &v
	return s
}

func (s *UpdateMmsTableRequest) Validate() error {
	return dara.Validate(s)
}
