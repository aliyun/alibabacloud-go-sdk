// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMigrateRegionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogicalRegionId(v string) *QueryMigrateRegionListRequest
	GetLogicalRegionId() *string
}

type QueryMigrateRegionListRequest struct {
	// The ID of the namespace.
	//
	// example:
	//
	// cn-hangzhou:test
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
}

func (s QueryMigrateRegionListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMigrateRegionListRequest) GoString() string {
	return s.String()
}

func (s *QueryMigrateRegionListRequest) GetLogicalRegionId() *string {
	return s.LogicalRegionId
}

func (s *QueryMigrateRegionListRequest) SetLogicalRegionId(v string) *QueryMigrateRegionListRequest {
	s.LogicalRegionId = &v
	return s
}

func (s *QueryMigrateRegionListRequest) Validate() error {
	return dara.Validate(s)
}
