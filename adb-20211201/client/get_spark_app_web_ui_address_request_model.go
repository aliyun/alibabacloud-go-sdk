// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppWebUiAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetSparkAppWebUiAddressRequest
	GetAppId() *string
	SetDBClusterId(v string) *GetSparkAppWebUiAddressRequest
	GetDBClusterId() *string
}

type GetSparkAppWebUiAddressRequest struct {
	// The Spark application ID.
	//
	// >  You can call the [ListSparkApps](https://help.aliyun.com/document_detail/455888.html) operation to query Spark application IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// s202205201533hz1209892000****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition clusters within a region.
	//
	// example:
	//
	// amv-uf6g8w25jacm7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s GetSparkAppWebUiAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppWebUiAddressRequest) GoString() string {
	return s.String()
}

func (s *GetSparkAppWebUiAddressRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetSparkAppWebUiAddressRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetSparkAppWebUiAddressRequest) SetAppId(v string) *GetSparkAppWebUiAddressRequest {
	s.AppId = &v
	return s
}

func (s *GetSparkAppWebUiAddressRequest) SetDBClusterId(v string) *GetSparkAppWebUiAddressRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkAppWebUiAddressRequest) Validate() error {
	return dara.Validate(s)
}
