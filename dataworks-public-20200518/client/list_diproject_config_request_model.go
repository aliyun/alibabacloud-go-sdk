// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIProjectConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationType(v string) *ListDIProjectConfigRequest
	GetDestinationType() *string
	SetProjectId(v int64) *ListDIProjectConfigRequest
	GetProjectId() *int64
	SetSourceType(v string) *ListDIProjectConfigRequest
	GetSourceType() *string
}

type ListDIProjectConfigRequest struct {
	// The type of the destinations of the synchronization solutions. This parameter cannot be left empty. Valid values: analyticdb_for_mysql, odps, elasticsearch, holo, mysql, and polardb.
	//
	// This parameter is required.
	//
	// example:
	//
	// elasticsearch
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the sources of the synchronization solutions. Valid values: oracle, mysql, polardb, datahub, drds, and analyticdb_for_mysql. If you do not configure this parameter, DataWorks applies the default global configuration to all sources.
	//
	// example:
	//
	// mysql
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ListDIProjectConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDIProjectConfigRequest) GoString() string {
	return s.String()
}

func (s *ListDIProjectConfigRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *ListDIProjectConfigRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDIProjectConfigRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListDIProjectConfigRequest) SetDestinationType(v string) *ListDIProjectConfigRequest {
	s.DestinationType = &v
	return s
}

func (s *ListDIProjectConfigRequest) SetProjectId(v int64) *ListDIProjectConfigRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDIProjectConfigRequest) SetSourceType(v string) *ListDIProjectConfigRequest {
	s.SourceType = &v
	return s
}

func (s *ListDIProjectConfigRequest) Validate() error {
	return dara.Validate(s)
}
