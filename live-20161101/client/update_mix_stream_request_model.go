// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMixStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *UpdateMixStreamRequest
	GetDomainName() *string
	SetInputStreamList(v string) *UpdateMixStreamRequest
	GetInputStreamList() *string
	SetLayoutId(v string) *UpdateMixStreamRequest
	GetLayoutId() *string
	SetMixStreamId(v string) *UpdateMixStreamRequest
	GetMixStreamId() *string
	SetOwnerId(v int64) *UpdateMixStreamRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateMixStreamRequest
	GetRegionId() *string
}

type UpdateMixStreamRequest struct {
	// The streaming domain.
	//
	// 	Notice:
	//
	// Only domain names in the China (Shanghai) and China (Beijing) regions are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The list of input streams for the mix. This is a JSON array.
	//
	// For more information, see **InputStreamConfig*	- below.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"InputStreamList":[{"LayoutConfig":{"FillSizeNormalized":[0.5,0.5],"FillPositionNormalized":[0,0],"PositionRefer":"topLeft","FillMode":"fit"},"LayoutChildId":1,"ResourceValue":"rtmp://aliyundoc.com/caster/8564a8d1659b4dc69df5f66cf4c9****","ResourceType":"live"},{"LayoutConfig":{"FillSizeNormalized":[0.5,0.5],"FillPositionNormalized":[0.5,0],"PositionRefer":"topLeft","FillMode":"fit"},"LayoutChildId":2,"ResourceValue":"http://developer.aliyundoc.com/3c3c25426cf744fdb90423e76b78a28a/69b1a16e2b1d423d9841bf27a96f134e-0b1cba51f58bb5ad3a854x96a2c735f****.mp4","ResourceType":"url"},{"LayoutConfig":{"FillSizeNormalized":[1,0.5],"FillPositionNormalized":[0,0.5],"PositionRefer":"topLeft","FillMode":"fit"},"LayoutChildId":3,"ResourceValue":"http://aliyundoc.com/c0c6c5446b56432389e91535864938da/ed4adc5263b4474c954b95607a5350ae-fda757b3328438a8cf-4k57f373a0f0****.mp4","ResourceType":"url"}]}
	InputStreamList *string `json:"InputStreamList,omitempty" xml:"InputStreamList,omitempty"`
	// The layout ID. The following values are supported:
	//
	// - **MixStreamLayout-1-1**
	//
	// - **MixStreamLayout-2-1**
	//
	// - **MixStreamLayout-2-2**
	//
	// - **MixStreamLayout-2-3**
	//
	// - **MixStreamLayout-3-1**
	//
	// - **MixStreamLayout-3-2**
	//
	// - **MixStreamLayout-4-1**
	//
	// - **USERDEFINED*	- (If you use a custom layout instead of a preset layout, set this parameter to this value.)
	//
	// example:
	//
	// MixStreamLayout-3-2
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// The ID of the stream mix task. If you created the task by calling the [CreateMixStream](https://help.aliyun.com/document_detail/2848087.html) operation, use the MixStreamId value returned in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5b2a046e-74d7-385e-d2d7-8a5b87e4****
	MixStreamId *string `json:"MixStreamId,omitempty" xml:"MixStreamId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateMixStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMixStreamRequest) GoString() string {
	return s.String()
}

func (s *UpdateMixStreamRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateMixStreamRequest) GetInputStreamList() *string {
	return s.InputStreamList
}

func (s *UpdateMixStreamRequest) GetLayoutId() *string {
	return s.LayoutId
}

func (s *UpdateMixStreamRequest) GetMixStreamId() *string {
	return s.MixStreamId
}

func (s *UpdateMixStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateMixStreamRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateMixStreamRequest) SetDomainName(v string) *UpdateMixStreamRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateMixStreamRequest) SetInputStreamList(v string) *UpdateMixStreamRequest {
	s.InputStreamList = &v
	return s
}

func (s *UpdateMixStreamRequest) SetLayoutId(v string) *UpdateMixStreamRequest {
	s.LayoutId = &v
	return s
}

func (s *UpdateMixStreamRequest) SetMixStreamId(v string) *UpdateMixStreamRequest {
	s.MixStreamId = &v
	return s
}

func (s *UpdateMixStreamRequest) SetOwnerId(v int64) *UpdateMixStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateMixStreamRequest) SetRegionId(v string) *UpdateMixStreamRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateMixStreamRequest) Validate() error {
	return dara.Validate(s)
}
