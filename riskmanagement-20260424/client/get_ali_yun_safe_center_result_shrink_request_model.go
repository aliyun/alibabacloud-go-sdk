// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliYunSafeCenterResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateSimilarSecurityEventsQueryTaskRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetCreateSimilarSecurityEventsQueryTaskRequestShrink() *string
	SetDescribeInstancesFullStatusRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetDescribeInstancesFullStatusRequestShrink() *string
	SetDescribeSecurityEventOperationStatusRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetDescribeSecurityEventOperationStatusRequestShrink() *string
	SetDescribeSimilarSecurityEventsRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetDescribeSimilarSecurityEventsRequestShrink() *string
	SetGetAssetDetailByUuidRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetGetAssetDetailByUuidRequestShrink() *string
	SetHandleSecurityEventsRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetHandleSecurityEventsRequestShrink() *string
	SetHandleSimilarSecurityEventsRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetHandleSimilarSecurityEventsRequestShrink() *string
	SetInterfaceCode(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetInterfaceCode() *string
	SetListInstancesRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetListInstancesRequestShrink() *string
	SetRegionId(v string) *GetAliYunSafeCenterResultShrinkRequest
	GetRegionId() *string
}

type GetAliYunSafeCenterResultShrinkRequest struct {
	CreateSimilarSecurityEventsQueryTaskRequestShrink *string `json:"CreateSimilarSecurityEventsQueryTaskRequest,omitempty" xml:"CreateSimilarSecurityEventsQueryTaskRequest,omitempty"`
	DescribeInstancesFullStatusRequestShrink          *string `json:"DescribeInstancesFullStatusRequest,omitempty" xml:"DescribeInstancesFullStatusRequest,omitempty"`
	DescribeSecurityEventOperationStatusRequestShrink *string `json:"DescribeSecurityEventOperationStatusRequest,omitempty" xml:"DescribeSecurityEventOperationStatusRequest,omitempty"`
	DescribeSimilarSecurityEventsRequestShrink        *string `json:"DescribeSimilarSecurityEventsRequest,omitempty" xml:"DescribeSimilarSecurityEventsRequest,omitempty"`
	GetAssetDetailByUuidRequestShrink                 *string `json:"GetAssetDetailByUuidRequest,omitempty" xml:"GetAssetDetailByUuidRequest,omitempty"`
	HandleSecurityEventsRequestShrink                 *string `json:"HandleSecurityEventsRequest,omitempty" xml:"HandleSecurityEventsRequest,omitempty"`
	HandleSimilarSecurityEventsRequestShrink          *string `json:"HandleSimilarSecurityEventsRequest,omitempty" xml:"HandleSimilarSecurityEventsRequest,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ListInstanceStatus
	InterfaceCode              *string `json:"InterfaceCode,omitempty" xml:"InterfaceCode,omitempty"`
	ListInstancesRequestShrink *string `json:"ListInstancesRequest,omitempty" xml:"ListInstancesRequest,omitempty"`
	// example:
	//
	// cn-guangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAliYunSafeCenterResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetCreateSimilarSecurityEventsQueryTaskRequestShrink() *string {
	return s.CreateSimilarSecurityEventsQueryTaskRequestShrink
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetDescribeInstancesFullStatusRequestShrink() *string {
	return s.DescribeInstancesFullStatusRequestShrink
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetDescribeSecurityEventOperationStatusRequestShrink() *string {
	return s.DescribeSecurityEventOperationStatusRequestShrink
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetDescribeSimilarSecurityEventsRequestShrink() *string {
	return s.DescribeSimilarSecurityEventsRequestShrink
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetGetAssetDetailByUuidRequestShrink() *string {
	return s.GetAssetDetailByUuidRequestShrink
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetHandleSecurityEventsRequestShrink() *string {
	return s.HandleSecurityEventsRequestShrink
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetHandleSimilarSecurityEventsRequestShrink() *string {
	return s.HandleSimilarSecurityEventsRequestShrink
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetInterfaceCode() *string {
	return s.InterfaceCode
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetListInstancesRequestShrink() *string {
	return s.ListInstancesRequestShrink
}

func (s *GetAliYunSafeCenterResultShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetCreateSimilarSecurityEventsQueryTaskRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.CreateSimilarSecurityEventsQueryTaskRequestShrink = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetDescribeInstancesFullStatusRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.DescribeInstancesFullStatusRequestShrink = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetDescribeSecurityEventOperationStatusRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.DescribeSecurityEventOperationStatusRequestShrink = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetDescribeSimilarSecurityEventsRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.DescribeSimilarSecurityEventsRequestShrink = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetGetAssetDetailByUuidRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.GetAssetDetailByUuidRequestShrink = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetHandleSecurityEventsRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.HandleSecurityEventsRequestShrink = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetHandleSimilarSecurityEventsRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.HandleSimilarSecurityEventsRequestShrink = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetInterfaceCode(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.InterfaceCode = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetListInstancesRequestShrink(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.ListInstancesRequestShrink = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) SetRegionId(v string) *GetAliYunSafeCenterResultShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetAliYunSafeCenterResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
