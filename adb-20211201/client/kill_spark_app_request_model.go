// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillSparkAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *KillSparkAppRequest
	GetAppId() *string
	SetDBClusterId(v string) *KillSparkAppRequest
	GetDBClusterId() *string
}

type KillSparkAppRequest struct {
	// The ID of the Spark application that you want to terminate.
	//
	// This parameter is required.
	//
	// example:
	//
	// s202204132018hzprec1ac****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// example:
	//
	// amv-bp1ub9grke1****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s KillSparkAppRequest) String() string {
	return dara.Prettify(s)
}

func (s KillSparkAppRequest) GoString() string {
	return s.String()
}

func (s *KillSparkAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *KillSparkAppRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *KillSparkAppRequest) SetAppId(v string) *KillSparkAppRequest {
	s.AppId = &v
	return s
}

func (s *KillSparkAppRequest) SetDBClusterId(v string) *KillSparkAppRequest {
	s.DBClusterId = &v
	return s
}

func (s *KillSparkAppRequest) Validate() error {
	return dara.Validate(s)
}
