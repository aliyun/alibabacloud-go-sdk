// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachWhitelistTemplateToInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *AttachWhitelistTemplateToInstanceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *AttachWhitelistTemplateToInstanceRequest
	GetRegionId() *string
	SetTemplateId(v string) *AttachWhitelistTemplateToInstanceRequest
	GetTemplateId() *string
}

type AttachWhitelistTemplateToInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-nkhdbf1d,cc-nkhdbf1s
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// RegionId
	//
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

func (s AttachWhitelistTemplateToInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachWhitelistTemplateToInstanceRequest) GoString() string {
	return s.String()
}

func (s *AttachWhitelistTemplateToInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *AttachWhitelistTemplateToInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachWhitelistTemplateToInstanceRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AttachWhitelistTemplateToInstanceRequest) SetDBInstanceId(v string) *AttachWhitelistTemplateToInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceRequest) SetRegionId(v string) *AttachWhitelistTemplateToInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceRequest) SetTemplateId(v string) *AttachWhitelistTemplateToInstanceRequest {
	s.TemplateId = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceRequest) Validate() error {
	return dara.Validate(s)
}
