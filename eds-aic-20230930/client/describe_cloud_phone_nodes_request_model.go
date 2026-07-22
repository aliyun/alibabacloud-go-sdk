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
	SetSortKey(v string) *DescribeCloudPhoneNodesRequest
	GetSortKey() *string
	SetSortType(v string) *DescribeCloudPhoneNodesRequest
	GetSortType() *string
	SetStatus(v string) *DescribeCloudPhoneNodesRequest
	GetStatus() *string
	SetTags(v []*DescribeCloudPhoneNodesRequestTags) *DescribeCloudPhoneNodesRequest
	GetTags() []*DescribeCloudPhoneNodesRequestTags
}

type DescribeCloudPhoneNodesRequest struct {
	// The instance ID of the bandwidth plan.
	//
	// example:
	//
	// cbwp-bp17psa7fhxqmm*****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing type. Only subscription is supported.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The maximum number of entries per page for a paged query. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. If the results of a query are not completely returned, the returned NextToken is not empty. You can pass the returned NextToken in the next request to continue the query.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uONHqPtDLM2U8s****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of cloud phone matrix IDs.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The name of the cloud phone matrix.
	//
	// example:
	//
	// node_name
	NodeName     *string   `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeNameList []*string `json:"NodeNameList,omitempty" xml:"NodeNameList,omitempty" type:"Repeated"`
	// The specifications of the cloud phone matrix.
	//
	// example:
	//
	// cpm.gx7.10xlarge
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	SortKey    *string `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
	SortType   *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// The status of the cloud phone matrix.
	//
	// example:
	//
	// CREATING
	Status *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*DescribeCloudPhoneNodesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *DescribeCloudPhoneNodesRequest) GetSortKey() *string {
	return s.SortKey
}

func (s *DescribeCloudPhoneNodesRequest) GetSortType() *string {
	return s.SortType
}

func (s *DescribeCloudPhoneNodesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeCloudPhoneNodesRequest) GetTags() []*DescribeCloudPhoneNodesRequestTags {
	return s.Tags
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

func (s *DescribeCloudPhoneNodesRequest) SetSortKey(v string) *DescribeCloudPhoneNodesRequest {
	s.SortKey = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetSortType(v string) *DescribeCloudPhoneNodesRequest {
	s.SortType = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetStatus(v string) *DescribeCloudPhoneNodesRequest {
	s.Status = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetTags(v []*DescribeCloudPhoneNodesRequestTags) *DescribeCloudPhoneNodesRequest {
	s.Tags = v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudPhoneNodesRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCloudPhoneNodesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodesRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodesRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeCloudPhoneNodesRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeCloudPhoneNodesRequestTags) SetKey(v string) *DescribeCloudPhoneNodesRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequestTags) SetValue(v string) *DescribeCloudPhoneNodesRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequestTags) Validate() error {
	return dara.Validate(s)
}
