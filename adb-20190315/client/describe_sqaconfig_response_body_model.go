// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQAConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSQAConfigResponseBody
	GetDBClusterId() *string
	SetGroupName(v string) *DescribeSQAConfigResponseBody
	GetGroupName() *string
	SetRequestId(v string) *DescribeSQAConfigResponseBody
	GetRequestId() *string
	SetSQAStatus(v string) *DescribeSQAConfigResponseBody
	GetSQAStatus() *string
}

type DescribeSQAConfigResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	//
	// example:
	//
	// am-8vbyw9awuj141haf9
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// test_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBE843D8-964D-5EA3-9D31-822125611B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether short query acceleration (SQA) is enabled.
	//
	// example:
	//
	// off
	SQAStatus *string `json:"SQAStatus,omitempty" xml:"SQAStatus,omitempty"`
}

func (s DescribeSQAConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQAConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQAConfigResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSQAConfigResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeSQAConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSQAConfigResponseBody) GetSQAStatus() *string {
	return s.SQAStatus
}

func (s *DescribeSQAConfigResponseBody) SetDBClusterId(v string) *DescribeSQAConfigResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQAConfigResponseBody) SetGroupName(v string) *DescribeSQAConfigResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeSQAConfigResponseBody) SetRequestId(v string) *DescribeSQAConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQAConfigResponseBody) SetSQAStatus(v string) *DescribeSQAConfigResponseBody {
	s.SQAStatus = &v
	return s
}

func (s *DescribeSQAConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
