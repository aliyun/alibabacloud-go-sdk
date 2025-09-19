// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridProxyLinkedClientListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *DescribeHybridProxyLinkedClientListRequest
	GetClusterName() *string
	SetCurrentPage(v int32) *DescribeHybridProxyLinkedClientListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeHybridProxyLinkedClientListRequest
	GetPageSize() *int32
	SetProxyUuid(v string) *DescribeHybridProxyLinkedClientListRequest
	GetProxyUuid() *string
	SetRemark(v string) *DescribeHybridProxyLinkedClientListRequest
	GetRemark() *string
	SetUuid(v string) *DescribeHybridProxyLinkedClientListRequest
	GetUuid() *string
}

type DescribeHybridProxyLinkedClientListRequest struct {
	// The name of the proxy cluster. You can query the name of the proxy cluster in the Security Center console.
	//
	// example:
	//
	// office-proxy
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The UUID of the proxy node. You can call the DescribeHybridProxyList operation to query the UUID of the proxy node.
	//
	// example:
	//
	// inet-proxy-3bb11fad-37d6-4aee-9c37-b0ad1612XXXX
	ProxyUuid *string `json:"ProxyUuid,omitempty" xml:"ProxyUuid,omitempty"`
	// The description of the proxy cluster.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The UUID of the server on which the Security Center agent is installed. You can query the UUID by querying asset information.
	//
	// example:
	//
	// 80d2f7d6-31a9-4d7f-8ff4-7ecc42f89ca****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeHybridProxyLinkedClientListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyLinkedClientListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyLinkedClientListRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeHybridProxyLinkedClientListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeHybridProxyLinkedClientListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHybridProxyLinkedClientListRequest) GetProxyUuid() *string {
	return s.ProxyUuid
}

func (s *DescribeHybridProxyLinkedClientListRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeHybridProxyLinkedClientListRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeHybridProxyLinkedClientListRequest) SetClusterName(v string) *DescribeHybridProxyLinkedClientListRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListRequest) SetCurrentPage(v int32) *DescribeHybridProxyLinkedClientListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListRequest) SetPageSize(v int32) *DescribeHybridProxyLinkedClientListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListRequest) SetProxyUuid(v string) *DescribeHybridProxyLinkedClientListRequest {
	s.ProxyUuid = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListRequest) SetRemark(v string) *DescribeHybridProxyLinkedClientListRequest {
	s.Remark = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListRequest) SetUuid(v string) *DescribeHybridProxyLinkedClientListRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListRequest) Validate() error {
	return dara.Validate(s)
}
