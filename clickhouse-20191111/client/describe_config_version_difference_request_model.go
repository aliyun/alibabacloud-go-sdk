// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigVersionDifferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeId(v string) *DescribeConfigVersionDifferenceRequest
	GetChangeId() *string
	SetDBClusterId(v string) *DescribeConfigVersionDifferenceRequest
	GetDBClusterId() *string
}

type DescribeConfigVersionDifferenceRequest struct {
	// The ID of the change record. You can call the [DescribeConfigHistory](https://help.aliyun.com/document_detail/452209.html) operation to query the ID of the change record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ChangeId *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all the clusters that are deployed in a specific region. The information includes the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1tm8zf130ew****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeConfigVersionDifferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigVersionDifferenceRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigVersionDifferenceRequest) GetChangeId() *string {
	return s.ChangeId
}

func (s *DescribeConfigVersionDifferenceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeConfigVersionDifferenceRequest) SetChangeId(v string) *DescribeConfigVersionDifferenceRequest {
	s.ChangeId = &v
	return s
}

func (s *DescribeConfigVersionDifferenceRequest) SetDBClusterId(v string) *DescribeConfigVersionDifferenceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeConfigVersionDifferenceRequest) Validate() error {
	return dara.Validate(s)
}
