// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceSharedRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v int64) *CreateDataSourceSharedRuleRequest
	GetDataSourceId() *int64
	SetEnvType(v string) *CreateDataSourceSharedRuleRequest
	GetEnvType() *string
	SetSharedUser(v string) *CreateDataSourceSharedRuleRequest
	GetSharedUser() *string
	SetTargetProjectId(v int64) *CreateDataSourceSharedRuleRequest
	GetTargetProjectId() *int64
}

type CreateDataSourceSharedRuleRequest struct {
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 144544
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Share data sources to the target project environment, including
	//
	// - Dev (Development Environment)
	//
	// - Prod (production environment)
	//
	// This parameter is required.
	//
	// example:
	//
	// Dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The user with which you want to share the data source. If you do not configure this parameter, the data source is shared to an entire workspace.
	//
	// example:
	//
	// 1107550004253538
	SharedUser *string `json:"SharedUser,omitempty" xml:"SharedUser,omitempty"`
	// The ID of the workspace to which you want to share the data source. You cannot share the data source to the workspace with which the data source is associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 106560
	TargetProjectId *int64 `json:"TargetProjectId,omitempty" xml:"TargetProjectId,omitempty"`
}

func (s CreateDataSourceSharedRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceSharedRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceSharedRuleRequest) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *CreateDataSourceSharedRuleRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *CreateDataSourceSharedRuleRequest) GetSharedUser() *string {
	return s.SharedUser
}

func (s *CreateDataSourceSharedRuleRequest) GetTargetProjectId() *int64 {
	return s.TargetProjectId
}

func (s *CreateDataSourceSharedRuleRequest) SetDataSourceId(v int64) *CreateDataSourceSharedRuleRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateDataSourceSharedRuleRequest) SetEnvType(v string) *CreateDataSourceSharedRuleRequest {
	s.EnvType = &v
	return s
}

func (s *CreateDataSourceSharedRuleRequest) SetSharedUser(v string) *CreateDataSourceSharedRuleRequest {
	s.SharedUser = &v
	return s
}

func (s *CreateDataSourceSharedRuleRequest) SetTargetProjectId(v int64) *CreateDataSourceSharedRuleRequest {
	s.TargetProjectId = &v
	return s
}

func (s *CreateDataSourceSharedRuleRequest) Validate() error {
	return dara.Validate(s)
}
