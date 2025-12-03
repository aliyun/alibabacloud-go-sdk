// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iXpackRelateDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *XpackRelateDBRequest
	GetClusterId() *string
	SetDbClusterIds(v string) *XpackRelateDBRequest
	GetDbClusterIds() *string
	SetRelateDbType(v string) *XpackRelateDBRequest
	GetRelateDbType() *string
}

type XpackRelateDBRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ap-bp1qtz9rcbbt3****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16o0pd52e3****
	DbClusterIds *string `json:"DbClusterIds,omitempty" xml:"DbClusterIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase
	RelateDbType *string `json:"RelateDbType,omitempty" xml:"RelateDbType,omitempty"`
}

func (s XpackRelateDBRequest) String() string {
	return dara.Prettify(s)
}

func (s XpackRelateDBRequest) GoString() string {
	return s.String()
}

func (s *XpackRelateDBRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *XpackRelateDBRequest) GetDbClusterIds() *string {
	return s.DbClusterIds
}

func (s *XpackRelateDBRequest) GetRelateDbType() *string {
	return s.RelateDbType
}

func (s *XpackRelateDBRequest) SetClusterId(v string) *XpackRelateDBRequest {
	s.ClusterId = &v
	return s
}

func (s *XpackRelateDBRequest) SetDbClusterIds(v string) *XpackRelateDBRequest {
	s.DbClusterIds = &v
	return s
}

func (s *XpackRelateDBRequest) SetRelateDbType(v string) *XpackRelateDBRequest {
	s.RelateDbType = &v
	return s
}

func (s *XpackRelateDBRequest) Validate() error {
	return dara.Validate(s)
}
