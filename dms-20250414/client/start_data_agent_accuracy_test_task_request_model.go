// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDataAgentAccuracyTestTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccuracyTestInsId(v string) *StartDataAgentAccuracyTestTaskRequest
	GetAccuracyTestInsId() *string
	SetCsvFile(v string) *StartDataAgentAccuracyTestTaskRequest
	GetCsvFile() *string
	SetDmsUnit(v string) *StartDataAgentAccuracyTestTaskRequest
	GetDmsUnit() *string
	SetRegionId(v string) *StartDataAgentAccuracyTestTaskRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *StartDataAgentAccuracyTestTaskRequest
	GetWorkspaceId() *string
}

type StartDataAgentAccuracyTestTaskRequest struct {
	// The accuracy test instance ID.
	//
	// example:
	//
	// at-106n4rg17gv9fxxxxxxxxxx
	AccuracyTestInsId *string `json:"AccuracyTestInsId,omitempty" xml:"AccuracyTestInsId,omitempty"`
	// The accuracy test sample.
	//
	// example:
	//
	// 包含问题、答案[、SQL]的文件
	CsvFile *string `json:"CsvFile,omitempty" xml:"CsvFile,omitempty"`
	// The DMS unit used to create the resource.
	//
	// example:
	//
	// cn-hangzhou
	DmsUnit *string `json:"DmsUnit,omitempty" xml:"DmsUnit,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 8wfig6l33n4f4xxxxxxxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s StartDataAgentAccuracyTestTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDataAgentAccuracyTestTaskRequest) GoString() string {
	return s.String()
}

func (s *StartDataAgentAccuracyTestTaskRequest) GetAccuracyTestInsId() *string {
	return s.AccuracyTestInsId
}

func (s *StartDataAgentAccuracyTestTaskRequest) GetCsvFile() *string {
	return s.CsvFile
}

func (s *StartDataAgentAccuracyTestTaskRequest) GetDmsUnit() *string {
	return s.DmsUnit
}

func (s *StartDataAgentAccuracyTestTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartDataAgentAccuracyTestTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *StartDataAgentAccuracyTestTaskRequest) SetAccuracyTestInsId(v string) *StartDataAgentAccuracyTestTaskRequest {
	s.AccuracyTestInsId = &v
	return s
}

func (s *StartDataAgentAccuracyTestTaskRequest) SetCsvFile(v string) *StartDataAgentAccuracyTestTaskRequest {
	s.CsvFile = &v
	return s
}

func (s *StartDataAgentAccuracyTestTaskRequest) SetDmsUnit(v string) *StartDataAgentAccuracyTestTaskRequest {
	s.DmsUnit = &v
	return s
}

func (s *StartDataAgentAccuracyTestTaskRequest) SetRegionId(v string) *StartDataAgentAccuracyTestTaskRequest {
	s.RegionId = &v
	return s
}

func (s *StartDataAgentAccuracyTestTaskRequest) SetWorkspaceId(v string) *StartDataAgentAccuracyTestTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *StartDataAgentAccuracyTestTaskRequest) Validate() error {
	return dara.Validate(s)
}
