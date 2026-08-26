// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStudioLayoutsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DescribeStudioLayoutsRequest
	GetCasterId() *string
	SetLayoutId(v string) *DescribeStudioLayoutsRequest
	GetLayoutId() *string
	SetOwnerId(v int64) *DescribeStudioLayoutsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeStudioLayoutsRequest
	GetRegionId() *string
}

type DescribeStudioLayoutsRequest struct {
	// The production studio ID.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the CasterId parameter value returned by the CreateCaster operation.
	//
	// - If you created the production studio in the ApsaraVideo Live console, go to **ApsaraVideo Live console*	- > **Production Studio*	- > **Cloud Production Studio*	- to view the ID.
	//
	// >
	//
	// > - The production studio name in the production studio list on the Cloud Production Studio page is the production studio ID.
	//
	// > - Only virtual studio production studios (NormType=4) are supported. If you pass in a production studio ID of another type, InvalidCaster.NotFound is returned. Call DescribeCasters and filter by NormType=4 to obtain the virtual studio production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e91b96****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The layout ID.
	//
	// Separate multiple layout IDs with commas (,). If this parameter is not specified, all layouts under the production studio are returned.
	//
	// If you added virtual studio layout settings by calling the [AddStudioLayout](https://help.aliyun.com/document_detail/2848062.html) operation, check the LayoutId parameter value returned by the AddStudioLayout operation.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeStudioLayoutsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStudioLayoutsRequest) GoString() string {
	return s.String()
}

func (s *DescribeStudioLayoutsRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeStudioLayoutsRequest) GetLayoutId() *string {
	return s.LayoutId
}

func (s *DescribeStudioLayoutsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeStudioLayoutsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStudioLayoutsRequest) SetCasterId(v string) *DescribeStudioLayoutsRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeStudioLayoutsRequest) SetLayoutId(v string) *DescribeStudioLayoutsRequest {
	s.LayoutId = &v
	return s
}

func (s *DescribeStudioLayoutsRequest) SetOwnerId(v int64) *DescribeStudioLayoutsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStudioLayoutsRequest) SetRegionId(v string) *DescribeStudioLayoutsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStudioLayoutsRequest) Validate() error {
	return dara.Validate(s)
}
