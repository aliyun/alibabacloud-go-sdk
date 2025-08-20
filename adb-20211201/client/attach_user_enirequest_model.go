// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachUserENIRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *AttachUserENIRequest
	GetDBClusterId() *string
}

type AttachUserENIRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL Data Lakehouse Edition clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s AttachUserENIRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachUserENIRequest) GoString() string {
	return s.String()
}

func (s *AttachUserENIRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *AttachUserENIRequest) SetDBClusterId(v string) *AttachUserENIRequest {
	s.DBClusterId = &v
	return s
}

func (s *AttachUserENIRequest) Validate() error {
	return dara.Validate(s)
}
