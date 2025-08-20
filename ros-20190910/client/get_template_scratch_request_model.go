// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateScratchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetTemplateScratchRequest
	GetRegionId() *string
	SetShowDataOption(v string) *GetTemplateScratchRequest
	GetShowDataOption() *string
	SetTemplateScratchId(v string) *GetTemplateScratchRequest
	GetTemplateScratchId() *string
}

type GetTemplateScratchRequest struct {
	// The region ID of the resource scenario.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The data display option. Valid values:
	//
	// 	- Sources: displays only the data of source nodes. This setting takes effect only when TemplateScratchType is set to ArchitectureDetection.
	//
	// 	- Source: displays only the data of the source node. This setting takes effect only when TemplateScratchType is not set to ArchitectureDetection.
	//
	// 	- Provisions: displays only the data of new nodes. This setting takes effect only when TemplateScratchType is not set to ArchitectureDetection.
	//
	// 	- All: displays all data.
	//
	// For more information about source nodes and new nodes, see [Overview](https://help.aliyun.com/document_detail/352074.html).
	//
	// >  If you do not specify this parameter, the node data is not displayed.
	//
	// example:
	//
	// Source
	ShowDataOption *string `json:"ShowDataOption,omitempty" xml:"ShowDataOption,omitempty"`
	// The ID of the resource scenario.
	//
	// example:
	//
	// ts-7f7a704cf71c49a6****
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
}

func (s GetTemplateScratchRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateScratchRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTemplateScratchRequest) GetShowDataOption() *string {
	return s.ShowDataOption
}

func (s *GetTemplateScratchRequest) GetTemplateScratchId() *string {
	return s.TemplateScratchId
}

func (s *GetTemplateScratchRequest) SetRegionId(v string) *GetTemplateScratchRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateScratchRequest) SetShowDataOption(v string) *GetTemplateScratchRequest {
	s.ShowDataOption = &v
	return s
}

func (s *GetTemplateScratchRequest) SetTemplateScratchId(v string) *GetTemplateScratchRequest {
	s.TemplateScratchId = &v
	return s
}

func (s *GetTemplateScratchRequest) Validate() error {
	return dara.Validate(s)
}
