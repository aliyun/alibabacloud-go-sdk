// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryXpackRelateDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *QueryXpackRelateDBRequest
	GetClusterId() *string
	SetHasSingleNode(v bool) *QueryXpackRelateDBRequest
	GetHasSingleNode() *bool
	SetRelateDbType(v string) *QueryXpackRelateDBRequest
	GetRelateDbType() *string
}

type QueryXpackRelateDBRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ap-bp1qtz9rcbbt3p6ng
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// false
	HasSingleNode *bool `json:"HasSingleNode,omitempty" xml:"HasSingleNode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase
	RelateDbType *string `json:"RelateDbType,omitempty" xml:"RelateDbType,omitempty"`
}

func (s QueryXpackRelateDBRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryXpackRelateDBRequest) GoString() string {
	return s.String()
}

func (s *QueryXpackRelateDBRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *QueryXpackRelateDBRequest) GetHasSingleNode() *bool {
	return s.HasSingleNode
}

func (s *QueryXpackRelateDBRequest) GetRelateDbType() *string {
	return s.RelateDbType
}

func (s *QueryXpackRelateDBRequest) SetClusterId(v string) *QueryXpackRelateDBRequest {
	s.ClusterId = &v
	return s
}

func (s *QueryXpackRelateDBRequest) SetHasSingleNode(v bool) *QueryXpackRelateDBRequest {
	s.HasSingleNode = &v
	return s
}

func (s *QueryXpackRelateDBRequest) SetRelateDbType(v string) *QueryXpackRelateDBRequest {
	s.RelateDbType = &v
	return s
}

func (s *QueryXpackRelateDBRequest) Validate() error {
	return dara.Validate(s)
}
