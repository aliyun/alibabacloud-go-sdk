// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMigrateEcuListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogicalRegionId(v string) *QueryMigrateEcuListRequest
	GetLogicalRegionId() *string
}

type QueryMigrateEcuListRequest struct {
	// The ID of the namespace.
	//
	// 	- The ID of a custom namespace is in the `region ID:namespace identifier` format. Example: `cn-beijing:test`.
	//
	// 	- The ID of the default namespace is in the `region ID` format. Example: `cn-beijing`.
	//
	// example:
	//
	// cn-hangzhou or cn-hangzhou:test
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
}

func (s QueryMigrateEcuListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMigrateEcuListRequest) GoString() string {
	return s.String()
}

func (s *QueryMigrateEcuListRequest) GetLogicalRegionId() *string {
	return s.LogicalRegionId
}

func (s *QueryMigrateEcuListRequest) SetLogicalRegionId(v string) *QueryMigrateEcuListRequest {
	s.LogicalRegionId = &v
	return s
}

func (s *QueryMigrateEcuListRequest) Validate() error {
	return dara.Validate(s)
}
