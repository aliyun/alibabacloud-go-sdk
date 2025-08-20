// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetSparkAppInfoRequest
	GetAppId() *string
	SetDBClusterId(v string) *GetSparkAppInfoRequest
	GetDBClusterId() *string
}

type GetSparkAppInfoRequest struct {
	// The application ID.
	//
	// >  You can call the [ListSparkApps](https://help.aliyun.com/document_detail/455888.html) operation to query the Spark application IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// s202205201533hz1209892000****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// example:
	//
	// am-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s GetSparkAppInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSparkAppInfoRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetSparkAppInfoRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetSparkAppInfoRequest) SetAppId(v string) *GetSparkAppInfoRequest {
	s.AppId = &v
	return s
}

func (s *GetSparkAppInfoRequest) SetDBClusterId(v string) *GetSparkAppInfoRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkAppInfoRequest) Validate() error {
	return dara.Validate(s)
}
