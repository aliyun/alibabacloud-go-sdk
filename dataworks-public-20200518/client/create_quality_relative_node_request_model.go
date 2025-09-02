// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityRelativeNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvType(v string) *CreateQualityRelativeNodeRequest
	GetEnvType() *string
	SetMatchExpression(v string) *CreateQualityRelativeNodeRequest
	GetMatchExpression() *string
	SetNodeId(v int64) *CreateQualityRelativeNodeRequest
	GetNodeId() *int64
	SetProjectId(v int64) *CreateQualityRelativeNodeRequest
	GetProjectId() *int64
	SetProjectName(v string) *CreateQualityRelativeNodeRequest
	GetProjectName() *string
	SetTableName(v string) *CreateQualityRelativeNodeRequest
	GetTableName() *string
	SetTargetNodeProjectId(v int64) *CreateQualityRelativeNodeRequest
	GetTargetNodeProjectId() *int64
	SetTargetNodeProjectName(v string) *CreateQualityRelativeNodeRequest
	GetTargetNodeProjectName() *string
}

type CreateQualityRelativeNodeRequest struct {
	// The type of the compute engine or data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// ODPS
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The partition filter expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// dt=$[yyyymmdd]
	MatchExpression *string `json:"MatchExpression,omitempty" xml:"MatchExpression,omitempty"`
	// The node ID. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12321
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the compute engine or data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// dual
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the workspace to which the node belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	TargetNodeProjectId *int64 `json:"TargetNodeProjectId,omitempty" xml:"TargetNodeProjectId,omitempty"`
	// The name of the workspace to which the node to be associated with the partition filter expression belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest
	TargetNodeProjectName *string `json:"TargetNodeProjectName,omitempty" xml:"TargetNodeProjectName,omitempty"`
}

func (s CreateQualityRelativeNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityRelativeNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateQualityRelativeNodeRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *CreateQualityRelativeNodeRequest) GetMatchExpression() *string {
	return s.MatchExpression
}

func (s *CreateQualityRelativeNodeRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *CreateQualityRelativeNodeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateQualityRelativeNodeRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateQualityRelativeNodeRequest) GetTableName() *string {
	return s.TableName
}

func (s *CreateQualityRelativeNodeRequest) GetTargetNodeProjectId() *int64 {
	return s.TargetNodeProjectId
}

func (s *CreateQualityRelativeNodeRequest) GetTargetNodeProjectName() *string {
	return s.TargetNodeProjectName
}

func (s *CreateQualityRelativeNodeRequest) SetEnvType(v string) *CreateQualityRelativeNodeRequest {
	s.EnvType = &v
	return s
}

func (s *CreateQualityRelativeNodeRequest) SetMatchExpression(v string) *CreateQualityRelativeNodeRequest {
	s.MatchExpression = &v
	return s
}

func (s *CreateQualityRelativeNodeRequest) SetNodeId(v int64) *CreateQualityRelativeNodeRequest {
	s.NodeId = &v
	return s
}

func (s *CreateQualityRelativeNodeRequest) SetProjectId(v int64) *CreateQualityRelativeNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateQualityRelativeNodeRequest) SetProjectName(v string) *CreateQualityRelativeNodeRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateQualityRelativeNodeRequest) SetTableName(v string) *CreateQualityRelativeNodeRequest {
	s.TableName = &v
	return s
}

func (s *CreateQualityRelativeNodeRequest) SetTargetNodeProjectId(v int64) *CreateQualityRelativeNodeRequest {
	s.TargetNodeProjectId = &v
	return s
}

func (s *CreateQualityRelativeNodeRequest) SetTargetNodeProjectName(v string) *CreateQualityRelativeNodeRequest {
	s.TargetNodeProjectName = &v
	return s
}

func (s *CreateQualityRelativeNodeRequest) Validate() error {
	return dara.Validate(s)
}
