// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySqlTemplatePositionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifySqlTemplatePositionRequest
	GetDBClusterId() *string
	SetRegionId(v string) *ModifySqlTemplatePositionRequest
	GetRegionId() *string
	SetTargetTemplateGroupId(v int64) *ModifySqlTemplatePositionRequest
	GetTargetTemplateGroupId() *int64
	SetTemplateId(v int64) *ModifySqlTemplatePositionRequest
	GetTemplateId() *int64
}

type ModifySqlTemplatePositionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// amv-uf6wjk5xxxxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
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
	// 1
	TargetTemplateGroupId *int64 `json:"TargetTemplateGroupId,omitempty" xml:"TargetTemplateGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 210208
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifySqlTemplatePositionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySqlTemplatePositionRequest) GoString() string {
	return s.String()
}

func (s *ModifySqlTemplatePositionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifySqlTemplatePositionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySqlTemplatePositionRequest) GetTargetTemplateGroupId() *int64 {
	return s.TargetTemplateGroupId
}

func (s *ModifySqlTemplatePositionRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ModifySqlTemplatePositionRequest) SetDBClusterId(v string) *ModifySqlTemplatePositionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifySqlTemplatePositionRequest) SetRegionId(v string) *ModifySqlTemplatePositionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySqlTemplatePositionRequest) SetTargetTemplateGroupId(v int64) *ModifySqlTemplatePositionRequest {
	s.TargetTemplateGroupId = &v
	return s
}

func (s *ModifySqlTemplatePositionRequest) SetTemplateId(v int64) *ModifySqlTemplatePositionRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifySqlTemplatePositionRequest) Validate() error {
	return dara.Validate(s)
}
