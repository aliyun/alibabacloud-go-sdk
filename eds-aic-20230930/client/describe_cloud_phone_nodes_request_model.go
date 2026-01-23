// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudPhoneNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *DescribeCloudPhoneNodesRequest
	GetBandwidthPackageId() *string
	SetBizRegionId(v string) *DescribeCloudPhoneNodesRequest
	GetBizRegionId() *string
	SetChargeType(v string) *DescribeCloudPhoneNodesRequest
	GetChargeType() *string
	SetMaxResults(v string) *DescribeCloudPhoneNodesRequest
	GetMaxResults() *string
	SetNextToken(v string) *DescribeCloudPhoneNodesRequest
	GetNextToken() *string
	SetNodeIds(v []*string) *DescribeCloudPhoneNodesRequest
	GetNodeIds() []*string
	SetNodeName(v string) *DescribeCloudPhoneNodesRequest
	GetNodeName() *string
	SetNodeNameList(v []*string) *DescribeCloudPhoneNodesRequest
	GetNodeNameList() []*string
	SetServerType(v string) *DescribeCloudPhoneNodesRequest
	GetServerType() *string
	SetStatus(v string) *DescribeCloudPhoneNodesRequest
	GetStatus() *string
}

type DescribeCloudPhoneNodesRequest struct {
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method. Only the subscription billing method is supported.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If a query doesn\\"t return all results, the response includes a NextToken value for pagination. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uONHqPtDLM2U8s****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The matrix IDs.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The matrix name.
	//
	// example:
	//
	// node_name
	NodeName     *string   `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeNameList []*string `json:"NodeNameList,omitempty" xml:"NodeNameList,omitempty" type:"Repeated"`
	// The matrix specification.
	//
	// Valid values:
	//
	// 	- cpm.gn6.gx1
	//
	// example:
	//
	// cpm.gn6.gx1
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The matrix status.
	//
	// Valid values:
	//
	// 	- FAILED: The matrix failed to be created.
	//
	// 	- RUNNING: The matrix is available.
	//
	// 	- DELETING: The matrix is being deleted.
	//
	// 	- NODE_READY: The matrix is ready, and cloud phone instances are being created.
	//
	// 	- DELETED: The matrix is deleted.
	//
	// 	- CREATING: The matrix is being created.
	//
	// example:
	//
	// CREATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCloudPhoneNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodesRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DescribeCloudPhoneNodesRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *DescribeCloudPhoneNodesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeCloudPhoneNodesRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *DescribeCloudPhoneNodesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudPhoneNodesRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *DescribeCloudPhoneNodesRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeCloudPhoneNodesRequest) GetNodeNameList() []*string {
	return s.NodeNameList
}

func (s *DescribeCloudPhoneNodesRequest) GetServerType() *string {
	return s.ServerType
}

func (s *DescribeCloudPhoneNodesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeCloudPhoneNodesRequest) SetBandwidthPackageId(v string) *DescribeCloudPhoneNodesRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetBizRegionId(v string) *DescribeCloudPhoneNodesRequest {
	s.BizRegionId = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetChargeType(v string) *DescribeCloudPhoneNodesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetMaxResults(v string) *DescribeCloudPhoneNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetNextToken(v string) *DescribeCloudPhoneNodesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetNodeIds(v []*string) *DescribeCloudPhoneNodesRequest {
	s.NodeIds = v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetNodeName(v string) *DescribeCloudPhoneNodesRequest {
	s.NodeName = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetNodeNameList(v []*string) *DescribeCloudPhoneNodesRequest {
	s.NodeNameList = v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetServerType(v string) *DescribeCloudPhoneNodesRequest {
	s.ServerType = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetStatus(v string) *DescribeCloudPhoneNodesRequest {
	s.Status = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) Validate() error {
	return dara.Validate(s)
}
