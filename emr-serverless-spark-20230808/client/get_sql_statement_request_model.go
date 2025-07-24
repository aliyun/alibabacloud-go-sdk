// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlStatementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetSqlStatementRequest
	GetRegionId() *string
}

type GetSqlStatementRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetSqlStatementRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSqlStatementRequest) GoString() string {
	return s.String()
}

func (s *GetSqlStatementRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSqlStatementRequest) SetRegionId(v string) *GetSqlStatementRequest {
	s.RegionId = &v
	return s
}

func (s *GetSqlStatementRequest) Validate() error {
	return dara.Validate(s)
}
