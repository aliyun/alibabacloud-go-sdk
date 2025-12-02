// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetADBSparkNecessaryRAMPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetADBSparkNecessaryRAMPermissionsRequest
	GetDBClusterId() *string
}

type GetADBSparkNecessaryRAMPermissionsRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s GetADBSparkNecessaryRAMPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetADBSparkNecessaryRAMPermissionsRequest) GoString() string {
	return s.String()
}

func (s *GetADBSparkNecessaryRAMPermissionsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetADBSparkNecessaryRAMPermissionsRequest) SetDBClusterId(v string) *GetADBSparkNecessaryRAMPermissionsRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetADBSparkNecessaryRAMPermissionsRequest) Validate() error {
	return dara.Validate(s)
}
