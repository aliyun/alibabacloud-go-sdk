// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentAccuracyTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccuracyTestInsId(v string) *DeleteDataAgentAccuracyTestRequest
	GetAccuracyTestInsId() *string
	SetDmsUnit(v string) *DeleteDataAgentAccuracyTestRequest
	GetDmsUnit() *string
	SetRegionId(v string) *DeleteDataAgentAccuracyTestRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *DeleteDataAgentAccuracyTestRequest
	GetWorkspaceId() *string
}

type DeleteDataAgentAccuracyTestRequest struct {
	// The accuracy test instance ID.
	//
	// example:
	//
	// at-106n4rg17gv9fxxxxxxxxxx
	AccuracyTestInsId *string `json:"AccuracyTestInsId,omitempty" xml:"AccuracyTestInsId,omitempty"`
	// The current DMS unit.
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
	// The workspace ID.
	//
	// example:
	//
	// 8wfig6l33n4f4xxxxxxxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDataAgentAccuracyTestRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentAccuracyTestRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentAccuracyTestRequest) GetAccuracyTestInsId() *string {
	return s.AccuracyTestInsId
}

func (s *DeleteDataAgentAccuracyTestRequest) GetDmsUnit() *string {
	return s.DmsUnit
}

func (s *DeleteDataAgentAccuracyTestRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDataAgentAccuracyTestRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteDataAgentAccuracyTestRequest) SetAccuracyTestInsId(v string) *DeleteDataAgentAccuracyTestRequest {
	s.AccuracyTestInsId = &v
	return s
}

func (s *DeleteDataAgentAccuracyTestRequest) SetDmsUnit(v string) *DeleteDataAgentAccuracyTestRequest {
	s.DmsUnit = &v
	return s
}

func (s *DeleteDataAgentAccuracyTestRequest) SetRegionId(v string) *DeleteDataAgentAccuracyTestRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDataAgentAccuracyTestRequest) SetWorkspaceId(v string) *DeleteDataAgentAccuracyTestRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDataAgentAccuracyTestRequest) Validate() error {
	return dara.Validate(s)
}
