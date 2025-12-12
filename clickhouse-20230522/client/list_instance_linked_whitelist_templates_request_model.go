// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceLinkedWhitelistTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListInstanceLinkedWhitelistTemplatesRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *ListInstanceLinkedWhitelistTemplatesRequest
	GetRegionId() *string
}

type ListInstanceLinkedWhitelistTemplatesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-exadfas
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstanceLinkedWhitelistTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceLinkedWhitelistTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceLinkedWhitelistTemplatesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListInstanceLinkedWhitelistTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceLinkedWhitelistTemplatesRequest) SetDBInstanceId(v string) *ListInstanceLinkedWhitelistTemplatesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListInstanceLinkedWhitelistTemplatesRequest) SetRegionId(v string) *ListInstanceLinkedWhitelistTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstanceLinkedWhitelistTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
