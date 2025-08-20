// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetSparkAppStateRequest
	GetAppId() *string
	SetDBClusterId(v string) *GetSparkAppStateRequest
	GetDBClusterId() *string
}

type GetSparkAppStateRequest struct {
	// The Spark application ID.
	//
	// >  You can call the [ListSparkApps](https://help.aliyun.com/document_detail/455888.html) operation to query Spark application IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// s202204191546hzpread6a896000****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// example:
	//
	// amv-bp1h405854m****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s GetSparkAppStateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppStateRequest) GoString() string {
	return s.String()
}

func (s *GetSparkAppStateRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetSparkAppStateRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetSparkAppStateRequest) SetAppId(v string) *GetSparkAppStateRequest {
	s.AppId = &v
	return s
}

func (s *GetSparkAppStateRequest) SetDBClusterId(v string) *GetSparkAppStateRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkAppStateRequest) Validate() error {
	return dara.Validate(s)
}
