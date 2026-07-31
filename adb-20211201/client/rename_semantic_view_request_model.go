// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameSemanticViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *RenameSemanticViewRequest
	GetDBClusterId() *string
	SetNewSchemaName(v string) *RenameSemanticViewRequest
	GetNewSchemaName() *string
	SetNewViewName(v string) *RenameSemanticViewRequest
	GetNewViewName() *string
	SetOldSchemaName(v string) *RenameSemanticViewRequest
	GetOldSchemaName() *string
	SetOldViewName(v string) *RenameSemanticViewRequest
	GetOldViewName() *string
}

type RenameSemanticViewRequest struct {
	// The ID of the AnalyticDB for MySQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1u8c0mgfg58****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The new schema name in which the semantic view resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// adb_sv
	NewSchemaName *string `json:"NewSchemaName,omitempty" xml:"NewSchemaName,omitempty"`
	// The new name of the semantic view.
	//
	// This parameter is required.
	//
	// example:
	//
	// new_sv_name
	NewViewName *string `json:"NewViewName,omitempty" xml:"NewViewName,omitempty"`
	// The original schema name in which the semantic view resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// adb_sv_old
	OldSchemaName *string `json:"OldSchemaName,omitempty" xml:"OldSchemaName,omitempty"`
	// The original name of the semantic view.
	//
	// This parameter is required.
	//
	// example:
	//
	// old_sv_name
	OldViewName *string `json:"OldViewName,omitempty" xml:"OldViewName,omitempty"`
}

func (s RenameSemanticViewRequest) String() string {
	return dara.Prettify(s)
}

func (s RenameSemanticViewRequest) GoString() string {
	return s.String()
}

func (s *RenameSemanticViewRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RenameSemanticViewRequest) GetNewSchemaName() *string {
	return s.NewSchemaName
}

func (s *RenameSemanticViewRequest) GetNewViewName() *string {
	return s.NewViewName
}

func (s *RenameSemanticViewRequest) GetOldSchemaName() *string {
	return s.OldSchemaName
}

func (s *RenameSemanticViewRequest) GetOldViewName() *string {
	return s.OldViewName
}

func (s *RenameSemanticViewRequest) SetDBClusterId(v string) *RenameSemanticViewRequest {
	s.DBClusterId = &v
	return s
}

func (s *RenameSemanticViewRequest) SetNewSchemaName(v string) *RenameSemanticViewRequest {
	s.NewSchemaName = &v
	return s
}

func (s *RenameSemanticViewRequest) SetNewViewName(v string) *RenameSemanticViewRequest {
	s.NewViewName = &v
	return s
}

func (s *RenameSemanticViewRequest) SetOldSchemaName(v string) *RenameSemanticViewRequest {
	s.OldSchemaName = &v
	return s
}

func (s *RenameSemanticViewRequest) SetOldViewName(v string) *RenameSemanticViewRequest {
	s.OldViewName = &v
	return s
}

func (s *RenameSemanticViewRequest) Validate() error {
	return dara.Validate(s)
}
