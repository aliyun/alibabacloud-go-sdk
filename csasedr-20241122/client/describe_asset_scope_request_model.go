// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DescribeAssetScopeRequest
	GetId() *int64
}

type DescribeAssetScopeRequest struct {
	// The ID of the asset scope to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1932
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeAssetScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetScopeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetScopeRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeAssetScopeRequest) SetId(v int64) *DescribeAssetScopeRequest {
	s.Id = &v
	return s
}

func (s *DescribeAssetScopeRequest) Validate() error {
	return dara.Validate(s)
}
