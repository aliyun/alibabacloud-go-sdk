// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStructureImportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *CreateStructureImportTaskRequest
	GetConfig() *string
	SetDBInstanceName(v string) *CreateStructureImportTaskRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *CreateStructureImportTaskRequest
	GetRegionId() *string
	SetSlinkTaskId(v string) *CreateStructureImportTaskRequest
	GetSlinkTaskId() *string
}

type CreateStructureImportTaskRequest struct {
	// example:
	//
	// [{\\"schemaName\\":\\"transfer_test\\",\\"tableList\\":[]}]
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
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
}

func (s CreateStructureImportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStructureImportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateStructureImportTaskRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateStructureImportTaskRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateStructureImportTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateStructureImportTaskRequest) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *CreateStructureImportTaskRequest) SetConfig(v string) *CreateStructureImportTaskRequest {
	s.Config = &v
	return s
}

func (s *CreateStructureImportTaskRequest) SetDBInstanceName(v string) *CreateStructureImportTaskRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateStructureImportTaskRequest) SetRegionId(v string) *CreateStructureImportTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStructureImportTaskRequest) SetSlinkTaskId(v string) *CreateStructureImportTaskRequest {
	s.SlinkTaskId = &v
	return s
}

func (s *CreateStructureImportTaskRequest) Validate() error {
	return dara.Validate(s)
}
