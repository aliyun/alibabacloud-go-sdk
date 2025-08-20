// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEssdCacheConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeEssdCacheConfigRequest
	GetDBClusterId() *string
}

type DescribeEssdCacheConfigRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-uf685u1o987hj6rn
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeEssdCacheConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEssdCacheConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeEssdCacheConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeEssdCacheConfigRequest) SetDBClusterId(v string) *DescribeEssdCacheConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeEssdCacheConfigRequest) Validate() error {
	return dara.Validate(s)
}
