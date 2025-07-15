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
	// The ID of the production studio instance.
	//
	// 	- If you call the [CreateCaster](https://help.aliyun.com/document_detail/69338.html) operation to create a production studio instance, you can obtain the instance ID from the CasterId parameter in the response.
	//
	// 	- If you create a production studio instance in the ApsaraVideo Live console, perform the following operations to obtain the instance ID: Log on to the **ApsaraVideo Live console*	- and click **Production Studios*	- in the left-side navigation pane. Then, view the instance ID on the **Production Studio Management*	- page.
	//
	// >  The value displayed in the Name column for an instance on the Production Studio Management page is the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e91b96****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The ID of the layout.
	//
	// You can specify multiple layout IDs and separate them with commas (,). If you leave this parameter empty, all layouts of the production studio are returned.
	//
	// If you call the [AddStudioLayout](https://help.aliyun.com/document_detail/215388.html) operation to configure a layout for a virtual studio, you can obtain the ID of the layout from the LayoutId parameter in the response.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
