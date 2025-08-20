// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadSampleDataSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *LoadSampleDataSetRequest
	GetDBClusterId() *string
}

type LoadSampleDataSetRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-2ze0z517o1mgp66a
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s LoadSampleDataSetRequest) String() string {
	return dara.Prettify(s)
}

func (s LoadSampleDataSetRequest) GoString() string {
	return s.String()
}

func (s *LoadSampleDataSetRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *LoadSampleDataSetRequest) SetDBClusterId(v string) *LoadSampleDataSetRequest {
	s.DBClusterId = &v
	return s
}

func (s *LoadSampleDataSetRequest) Validate() error {
	return dara.Validate(s)
}
