// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateSqlStatementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *TerminateSqlStatementRequest
	GetRegionId() *string
}

type TerminateSqlStatementRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s TerminateSqlStatementRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminateSqlStatementRequest) GoString() string {
	return s.String()
}

func (s *TerminateSqlStatementRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TerminateSqlStatementRequest) SetRegionId(v string) *TerminateSqlStatementRequest {
	s.RegionId = &v
	return s
}

func (s *TerminateSqlStatementRequest) Validate() error {
	return dara.Validate(s)
}
