// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMixStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackConfig(v string) *CreateMixStreamRequest
	GetCallbackConfig() *string
	SetDomainName(v string) *CreateMixStreamRequest
	GetDomainName() *string
	SetInputStreamList(v string) *CreateMixStreamRequest
	GetInputStreamList() *string
	SetLayoutId(v string) *CreateMixStreamRequest
	GetLayoutId() *string
	SetOutputConfig(v string) *CreateMixStreamRequest
	GetOutputConfig() *string
	SetOwnerId(v int64) *CreateMixStreamRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateMixStreamRequest
	GetRegionId() *string
}

type CreateMixStreamRequest struct {
	// The webhook address. This is a JSON array. When an event occurs, the live service sends an HTTP POST request to this address. The request body contains the event details.
	//
	// example:
	//
	// {"CallbackUrl":"http://aliyundoc.com"}
	CallbackConfig *string `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty"`
	// The primary streaming domain.
	//
	// 	Notice:
	//
	// Currently, only domain names in the China (Shanghai) and China (Beijing) regions are supported.
	//
	// </notice>
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The list of input streams. This is a JSON array.
	//
	// For more information, see **InputStreamConfig*	- below.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"LayoutChildId":1,"ResourceType":"live","ResourceValue":"rtmp://example.net/live/f2139ec2b8d6a191068cd****ea9064d?auth_key=1600947017-0-0-0b5645fe35d21a65ab92b394bd4d****","LayoutConfig":{"FillMode":"fit","PositionRefer":"topLeft","FillPositionNormalized":[0,0],"FillSizeNormalized":[1,1]}}]
	InputStreamList *string `json:"InputStreamList,omitempty" xml:"InputStreamList,omitempty"`
	// The layout ID. Valid values:
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
	// - **USERDEFINED*	- (If you do not use a preset layout, set this parameter to **USERDEFINED**.)
	//
	// > For more information, see [Preset stream mix layouts](https://help.aliyun.com/document_detail/199359.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// MixStreamLayout-1-1
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// The output configuration. This is a JSON string.
	//
	// For more information, see **OutputConfig*	- below.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"AppName":"liveApp****","StreamName":"9a78fb3f5c508be0122746f677a3****","MixStreamTemplate":"lp_hd_v","ExpireDuration":"86400"}
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateMixStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMixStreamRequest) GoString() string {
	return s.String()
}

func (s *CreateMixStreamRequest) GetCallbackConfig() *string {
	return s.CallbackConfig
}

func (s *CreateMixStreamRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateMixStreamRequest) GetInputStreamList() *string {
	return s.InputStreamList
}

func (s *CreateMixStreamRequest) GetLayoutId() *string {
	return s.LayoutId
}

func (s *CreateMixStreamRequest) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *CreateMixStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateMixStreamRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMixStreamRequest) SetCallbackConfig(v string) *CreateMixStreamRequest {
	s.CallbackConfig = &v
	return s
}

func (s *CreateMixStreamRequest) SetDomainName(v string) *CreateMixStreamRequest {
	s.DomainName = &v
	return s
}

func (s *CreateMixStreamRequest) SetInputStreamList(v string) *CreateMixStreamRequest {
	s.InputStreamList = &v
	return s
}

func (s *CreateMixStreamRequest) SetLayoutId(v string) *CreateMixStreamRequest {
	s.LayoutId = &v
	return s
}

func (s *CreateMixStreamRequest) SetOutputConfig(v string) *CreateMixStreamRequest {
	s.OutputConfig = &v
	return s
}

func (s *CreateMixStreamRequest) SetOwnerId(v int64) *CreateMixStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMixStreamRequest) SetRegionId(v string) *CreateMixStreamRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMixStreamRequest) Validate() error {
	return dara.Validate(s)
}
