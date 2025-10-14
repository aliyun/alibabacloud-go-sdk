// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSqlAuditSlsStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CheckSqlAuditSlsStatusRequest
	GetRegionId() *string
}

type CheckSqlAuditSlsStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckSqlAuditSlsStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckSqlAuditSlsStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckSqlAuditSlsStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckSqlAuditSlsStatusRequest) SetRegionId(v string) *CheckSqlAuditSlsStatusRequest {
	s.RegionId = &v
	return s
}

func (s *CheckSqlAuditSlsStatusRequest) Validate() error {
	return dara.Validate(s)
}
