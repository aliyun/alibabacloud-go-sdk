// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachWhitelistTemplateToInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DetachWhitelistTemplateToInstanceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DetachWhitelistTemplateToInstanceRequest
	GetRegionId() *string
	SetTemplateId(v string) *DetachWhitelistTemplateToInstanceRequest
	GetTemplateId() *string
}

type DetachWhitelistTemplateToInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-nkhdbf1d,cc-nkhdbf1s
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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
	// 98a6d3db05984dca
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DetachWhitelistTemplateToInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachWhitelistTemplateToInstanceRequest) GoString() string {
	return s.String()
}

func (s *DetachWhitelistTemplateToInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DetachWhitelistTemplateToInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachWhitelistTemplateToInstanceRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DetachWhitelistTemplateToInstanceRequest) SetDBInstanceId(v string) *DetachWhitelistTemplateToInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceRequest) SetRegionId(v string) *DetachWhitelistTemplateToInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceRequest) SetTemplateId(v string) *DetachWhitelistTemplateToInstanceRequest {
	s.TemplateId = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceRequest) Validate() error {
	return dara.Validate(s)
}
