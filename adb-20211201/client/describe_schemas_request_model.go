// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSchemasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSchemasRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeSchemasRequest
	GetRegionId() *string
}

type DescribeSchemasRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1xxxxxxxx47
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSchemasRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSchemasRequest) GoString() string {
	return s.String()
}

func (s *DescribeSchemasRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSchemasRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSchemasRequest) SetDBClusterId(v string) *DescribeSchemasRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSchemasRequest) SetRegionId(v string) *DescribeSchemasRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSchemasRequest) Validate() error {
	return dara.Validate(s)
}
