// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataImportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateDataImportTaskRequest
	GetDBInstanceName() *string
	SetDstDb(v string) *CreateDataImportTaskRequest
	GetDstDb() *string
	SetDstPassword(v string) *CreateDataImportTaskRequest
	GetDstPassword() *string
	SetDstResId(v string) *CreateDataImportTaskRequest
	GetDstResId() *string
	SetDstUserName(v string) *CreateDataImportTaskRequest
	GetDstUserName() *string
	SetRegionId(v string) *CreateDataImportTaskRequest
	GetRegionId() *string
	SetSlinkTaskId(v string) *CreateDataImportTaskRequest
	GetSlinkTaskId() *string
	SetSrcDb(v string) *CreateDataImportTaskRequest
	GetSrcDb() *string
	SetSrcPassword(v string) *CreateDataImportTaskRequest
	GetSrcPassword() *string
	SetSrcResId(v string) *CreateDataImportTaskRequest
	GetSrcResId() *string
	SetSrcUserName(v string) *CreateDataImportTaskRequest
	GetSrcUserName() *string
}

type CreateDataImportTaskRequest struct {
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// transfer_test3
	DstDb       *string `json:"DstDb,omitempty" xml:"DstDb,omitempty"`
	DstPassword *string `json:"DstPassword,omitempty" xml:"DstPassword,omitempty"`
	// example:
	//
	// pxc-shr8****k36vrn
	DstResId    *string `json:"DstResId,omitempty" xml:"DstResId,omitempty"`
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
	// transfer_for_st
	SrcDb       *string `json:"SrcDb,omitempty" xml:"SrcDb,omitempty"`
	SrcPassword *string `json:"SrcPassword,omitempty" xml:"SrcPassword,omitempty"`
	// example:
	//
	// pxc-shrnv****kh87z
	SrcResId    *string `json:"SrcResId,omitempty" xml:"SrcResId,omitempty"`
	SrcUserName *string `json:"SrcUserName,omitempty" xml:"SrcUserName,omitempty"`
}

func (s CreateDataImportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataImportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDataImportTaskRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateDataImportTaskRequest) GetDstDb() *string {
	return s.DstDb
}

func (s *CreateDataImportTaskRequest) GetDstPassword() *string {
	return s.DstPassword
}

func (s *CreateDataImportTaskRequest) GetDstResId() *string {
	return s.DstResId
}

func (s *CreateDataImportTaskRequest) GetDstUserName() *string {
	return s.DstUserName
}

func (s *CreateDataImportTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDataImportTaskRequest) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *CreateDataImportTaskRequest) GetSrcDb() *string {
	return s.SrcDb
}

func (s *CreateDataImportTaskRequest) GetSrcPassword() *string {
	return s.SrcPassword
}

func (s *CreateDataImportTaskRequest) GetSrcResId() *string {
	return s.SrcResId
}

func (s *CreateDataImportTaskRequest) GetSrcUserName() *string {
	return s.SrcUserName
}

func (s *CreateDataImportTaskRequest) SetDBInstanceName(v string) *CreateDataImportTaskRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetDstDb(v string) *CreateDataImportTaskRequest {
	s.DstDb = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetDstPassword(v string) *CreateDataImportTaskRequest {
	s.DstPassword = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetDstResId(v string) *CreateDataImportTaskRequest {
	s.DstResId = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetDstUserName(v string) *CreateDataImportTaskRequest {
	s.DstUserName = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetRegionId(v string) *CreateDataImportTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetSlinkTaskId(v string) *CreateDataImportTaskRequest {
	s.SlinkTaskId = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetSrcDb(v string) *CreateDataImportTaskRequest {
	s.SrcDb = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetSrcPassword(v string) *CreateDataImportTaskRequest {
	s.SrcPassword = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetSrcResId(v string) *CreateDataImportTaskRequest {
	s.SrcResId = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetSrcUserName(v string) *CreateDataImportTaskRequest {
	s.SrcUserName = &v
	return s
}

func (s *CreateDataImportTaskRequest) Validate() error {
	return dara.Validate(s)
}
