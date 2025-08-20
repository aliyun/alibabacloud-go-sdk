// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetSparkAppMetricsRequest
	GetAppId() *string
	SetDBClusterId(v string) *GetSparkAppMetricsRequest
	GetDBClusterId() *string
}

type GetSparkAppMetricsRequest struct {
	// The ID of the Spark application.
	//
	// This parameter is required.
	//
	// example:
	//
	// s202204221525hzca7d8140000003
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// example:
	//
	// amv-bp1ggnu61d77****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s GetSparkAppMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetSparkAppMetricsRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetSparkAppMetricsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetSparkAppMetricsRequest) SetAppId(v string) *GetSparkAppMetricsRequest {
	s.AppId = &v
	return s
}

func (s *GetSparkAppMetricsRequest) SetDBClusterId(v string) *GetSparkAppMetricsRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkAppMetricsRequest) Validate() error {
	return dara.Validate(s)
}
