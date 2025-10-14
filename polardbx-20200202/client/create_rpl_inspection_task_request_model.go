// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRplInspectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDstDb(v string) *CreateRplInspectionTaskRequest
	GetDstDb() *string
	SetDstPassword(v string) *CreateRplInspectionTaskRequest
	GetDstPassword() *string
	SetDstResId(v string) *CreateRplInspectionTaskRequest
	GetDstResId() *string
	SetDstUserName(v string) *CreateRplInspectionTaskRequest
	GetDstUserName() *string
	SetRegionId(v string) *CreateRplInspectionTaskRequest
	GetRegionId() *string
	SetSlinkTaskId(v string) *CreateRplInspectionTaskRequest
	GetSlinkTaskId() *string
	SetSrcPassword(v string) *CreateRplInspectionTaskRequest
	GetSrcPassword() *string
	SetSrcUserName(v string) *CreateRplInspectionTaskRequest
	GetSrcUserName() *string
}

type CreateRplInspectionTaskRequest struct {
	// example:
	//
	// transfer_test3
	DstDb *string `json:"DstDb,omitempty" xml:"DstDb,omitempty"`
	// example:
	//
	// ******
	DstPassword *string `json:"DstPassword,omitempty" xml:"DstPassword,omitempty"`
	// example:
	//
	// pxc-zkrc1****l54rc
	DstResId *string `json:"DstResId,omitempty" xml:"DstResId,omitempty"`
	// example:
	//
	// bbt_item
	DstUserName *string `json:"DstUserName,omitempty" xml:"DstUserName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
	// example:
	//
	// ******
	SrcPassword *string `json:"SrcPassword,omitempty" xml:"SrcPassword,omitempty"`
	// example:
	//
	// bbt_ump
	SrcUserName *string `json:"SrcUserName,omitempty" xml:"SrcUserName,omitempty"`
}

func (s CreateRplInspectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRplInspectionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRplInspectionTaskRequest) GetDstDb() *string {
	return s.DstDb
}

func (s *CreateRplInspectionTaskRequest) GetDstPassword() *string {
	return s.DstPassword
}

func (s *CreateRplInspectionTaskRequest) GetDstResId() *string {
	return s.DstResId
}

func (s *CreateRplInspectionTaskRequest) GetDstUserName() *string {
	return s.DstUserName
}

func (s *CreateRplInspectionTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRplInspectionTaskRequest) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *CreateRplInspectionTaskRequest) GetSrcPassword() *string {
	return s.SrcPassword
}

func (s *CreateRplInspectionTaskRequest) GetSrcUserName() *string {
	return s.SrcUserName
}

func (s *CreateRplInspectionTaskRequest) SetDstDb(v string) *CreateRplInspectionTaskRequest {
	s.DstDb = &v
	return s
}

func (s *CreateRplInspectionTaskRequest) SetDstPassword(v string) *CreateRplInspectionTaskRequest {
	s.DstPassword = &v
	return s
}

func (s *CreateRplInspectionTaskRequest) SetDstResId(v string) *CreateRplInspectionTaskRequest {
	s.DstResId = &v
	return s
}

func (s *CreateRplInspectionTaskRequest) SetDstUserName(v string) *CreateRplInspectionTaskRequest {
	s.DstUserName = &v
	return s
}

func (s *CreateRplInspectionTaskRequest) SetRegionId(v string) *CreateRplInspectionTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRplInspectionTaskRequest) SetSlinkTaskId(v string) *CreateRplInspectionTaskRequest {
	s.SlinkTaskId = &v
	return s
}

func (s *CreateRplInspectionTaskRequest) SetSrcPassword(v string) *CreateRplInspectionTaskRequest {
	s.SrcPassword = &v
	return s
}

func (s *CreateRplInspectionTaskRequest) SetSrcUserName(v string) *CreateRplInspectionTaskRequest {
	s.SrcUserName = &v
	return s
}

func (s *CreateRplInspectionTaskRequest) Validate() error {
	return dara.Validate(s)
}
