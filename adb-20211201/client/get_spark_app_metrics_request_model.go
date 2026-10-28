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
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s202204221525hzca7d8140000003
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// <props="china">The ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to query the IDs of all clusters in the target region.
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
