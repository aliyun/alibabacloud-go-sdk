// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadSparkAppMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *PreloadSparkAppMetricsRequest
	GetAppId() *string
	SetDBClusterId(v string) *PreloadSparkAppMetricsRequest
	GetDBClusterId() *string
}

type PreloadSparkAppMetricsRequest struct {
	// The Spark application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s202204221525hzca7d8140000003
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// example:
	//
	// amv-bp1mfe9qmsw1dzyg
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s PreloadSparkAppMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s PreloadSparkAppMetricsRequest) GoString() string {
	return s.String()
}

func (s *PreloadSparkAppMetricsRequest) GetAppId() *string {
	return s.AppId
}

func (s *PreloadSparkAppMetricsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *PreloadSparkAppMetricsRequest) SetAppId(v string) *PreloadSparkAppMetricsRequest {
	s.AppId = &v
	return s
}

func (s *PreloadSparkAppMetricsRequest) SetDBClusterId(v string) *PreloadSparkAppMetricsRequest {
	s.DBClusterId = &v
	return s
}

func (s *PreloadSparkAppMetricsRequest) Validate() error {
	return dara.Validate(s)
}
