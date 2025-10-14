// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshImportMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *RefreshImportMetaRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *RefreshImportMetaRequest
	GetRegionId() *string
	SetSlinkTaskId(v string) *RefreshImportMetaRequest
	GetSlinkTaskId() *string
}

type RefreshImportMetaRequest struct {
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
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
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
}

func (s RefreshImportMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshImportMetaRequest) GoString() string {
	return s.String()
}

func (s *RefreshImportMetaRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *RefreshImportMetaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RefreshImportMetaRequest) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *RefreshImportMetaRequest) SetDBInstanceName(v string) *RefreshImportMetaRequest {
	s.DBInstanceName = &v
	return s
}

func (s *RefreshImportMetaRequest) SetRegionId(v string) *RefreshImportMetaRequest {
	s.RegionId = &v
	return s
}

func (s *RefreshImportMetaRequest) SetSlinkTaskId(v string) *RefreshImportMetaRequest {
	s.SlinkTaskId = &v
	return s
}

func (s *RefreshImportMetaRequest) Validate() error {
	return dara.Validate(s)
}
