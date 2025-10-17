// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsJobDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAPSJobDetail(v *DescribeApsJobDetailResponseBodyAPSJobDetail) *DescribeApsJobDetailResponseBody
	GetAPSJobDetail() *DescribeApsJobDetailResponseBodyAPSJobDetail
	SetRequestId(v string) *DescribeApsJobDetailResponseBody
	GetRequestId() *string
}

type DescribeApsJobDetailResponseBody struct {
	// The queried job.
	//
	// example:
	//
	// -
	APSJobDetail *DescribeApsJobDetailResponseBodyAPSJobDetail `json:"APSJobDetail,omitempty" xml:"APSJobDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******-E606-4A42-BF6D-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApsJobDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApsJobDetailResponseBody) GetAPSJobDetail() *DescribeApsJobDetailResponseBodyAPSJobDetail {
	return s.APSJobDetail
}

func (s *DescribeApsJobDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApsJobDetailResponseBody) SetAPSJobDetail(v *DescribeApsJobDetailResponseBodyAPSJobDetail) *DescribeApsJobDetailResponseBody {
	s.APSJobDetail = v
	return s
}

func (s *DescribeApsJobDetailResponseBody) SetRequestId(v string) *DescribeApsJobDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApsJobDetailResponseBody) Validate() error {
	if s.APSJobDetail != nil {
		if err := s.APSJobDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApsJobDetailResponseBodyAPSJobDetail struct {
	// The objects that are synchronized.
	//
	// example:
	//
	// {}
	DbList *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The ID of the destination cluster.
	//
	// example:
	//
	// amv-******
	DestinationEndpointInstanceID *string `json:"DestinationEndpointInstanceID,omitempty" xml:"DestinationEndpointInstanceID,omitempty"`
	// The region of the destination cluster.
	//
	// example:
	//
	// cn-beijing
	DestinationEndpointRegion *string `json:"DestinationEndpointRegion,omitempty" xml:"DestinationEndpointRegion,omitempty"`
	// The partitions.
	//
	// example:
	//
	// {}
	PartitionList *string `json:"PartitionList,omitempty" xml:"PartitionList,omitempty"`
	// The ID of the source instance.
	//
	// example:
	//
	// pc-*******
	SourceEndpointInstanceID *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	// The region of the source instance.
	//
	// example:
	//
	// cn-beijing
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// The status of the job.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The mode of the destination table.
	//
	// example:
	//
	// 0
	TargetTableMode *string `json:"TargetTableMode,omitempty" xml:"TargetTableMode,omitempty"`
}

func (s DescribeApsJobDetailResponseBodyAPSJobDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsJobDetailResponseBodyAPSJobDetail) GoString() string {
	return s.String()
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) GetDbList() *string {
	return s.DbList
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) GetDestinationEndpointInstanceID() *string {
	return s.DestinationEndpointInstanceID
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) GetDestinationEndpointRegion() *string {
	return s.DestinationEndpointRegion
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) GetPartitionList() *string {
	return s.PartitionList
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) GetSourceEndpointInstanceID() *string {
	return s.SourceEndpointInstanceID
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) GetSourceEndpointRegion() *string {
	return s.SourceEndpointRegion
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) GetStatus() *string {
	return s.Status
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) GetTargetTableMode() *string {
	return s.TargetTableMode
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) SetDbList(v string) *DescribeApsJobDetailResponseBodyAPSJobDetail {
	s.DbList = &v
	return s
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) SetDestinationEndpointInstanceID(v string) *DescribeApsJobDetailResponseBodyAPSJobDetail {
	s.DestinationEndpointInstanceID = &v
	return s
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) SetDestinationEndpointRegion(v string) *DescribeApsJobDetailResponseBodyAPSJobDetail {
	s.DestinationEndpointRegion = &v
	return s
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) SetPartitionList(v string) *DescribeApsJobDetailResponseBodyAPSJobDetail {
	s.PartitionList = &v
	return s
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) SetSourceEndpointInstanceID(v string) *DescribeApsJobDetailResponseBodyAPSJobDetail {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) SetSourceEndpointRegion(v string) *DescribeApsJobDetailResponseBodyAPSJobDetail {
	s.SourceEndpointRegion = &v
	return s
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) SetStatus(v string) *DescribeApsJobDetailResponseBodyAPSJobDetail {
	s.Status = &v
	return s
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) SetTargetTableMode(v string) *DescribeApsJobDetailResponseBodyAPSJobDetail {
	s.TargetTableMode = &v
	return s
}

func (s *DescribeApsJobDetailResponseBodyAPSJobDetail) Validate() error {
	return dara.Validate(s)
}
