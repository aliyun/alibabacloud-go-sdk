// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFabricTopologyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetFabricTopologyResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetFabricTopologyResponseBody
	GetCode() *int32
	SetContent(v *GetFabricTopologyResponseBodyContent) *GetFabricTopologyResponseBody
	GetContent() *GetFabricTopologyResponseBodyContent
	SetMessage(v string) *GetFabricTopologyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFabricTopologyResponseBody
	GetRequestId() *string
}

type GetFabricTopologyResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *GetFabricTopologyResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFabricTopologyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFabricTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *GetFabricTopologyResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetFabricTopologyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetFabricTopologyResponseBody) GetContent() *GetFabricTopologyResponseBodyContent {
	return s.Content
}

func (s *GetFabricTopologyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFabricTopologyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFabricTopologyResponseBody) SetAccessDeniedDetail(v string) *GetFabricTopologyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetFabricTopologyResponseBody) SetCode(v int32) *GetFabricTopologyResponseBody {
	s.Code = &v
	return s
}

func (s *GetFabricTopologyResponseBody) SetContent(v *GetFabricTopologyResponseBodyContent) *GetFabricTopologyResponseBody {
	s.Content = v
	return s
}

func (s *GetFabricTopologyResponseBody) SetMessage(v string) *GetFabricTopologyResponseBody {
	s.Message = &v
	return s
}

func (s *GetFabricTopologyResponseBody) SetRequestId(v string) *GetFabricTopologyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFabricTopologyResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFabricTopologyResponseBodyContent struct {
	// The cluster ID.
	//
	// example:
	//
	// cluster-****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// network interface controller Topology Information
	TopoInfo []*GetFabricTopologyResponseBodyContentTopoInfo `json:"TopoInfo,omitempty" xml:"TopoInfo,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-j6ctp4n75306****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Lingjun CIDR block ID
	//
	// example:
	//
	// vpd-fuli****
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s GetFabricTopologyResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetFabricTopologyResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetFabricTopologyResponseBodyContent) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetFabricTopologyResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetFabricTopologyResponseBodyContent) GetTopoInfo() []*GetFabricTopologyResponseBodyContentTopoInfo {
	return s.TopoInfo
}

func (s *GetFabricTopologyResponseBodyContent) GetVpcId() *string {
	return s.VpcId
}

func (s *GetFabricTopologyResponseBodyContent) GetVpdId() *string {
	return s.VpdId
}

func (s *GetFabricTopologyResponseBodyContent) SetClusterId(v string) *GetFabricTopologyResponseBodyContent {
	s.ClusterId = &v
	return s
}

func (s *GetFabricTopologyResponseBodyContent) SetRegionId(v string) *GetFabricTopologyResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetFabricTopologyResponseBodyContent) SetTopoInfo(v []*GetFabricTopologyResponseBodyContentTopoInfo) *GetFabricTopologyResponseBodyContent {
	s.TopoInfo = v
	return s
}

func (s *GetFabricTopologyResponseBodyContent) SetVpcId(v string) *GetFabricTopologyResponseBodyContent {
	s.VpcId = &v
	return s
}

func (s *GetFabricTopologyResponseBodyContent) SetVpdId(v string) *GetFabricTopologyResponseBodyContent {
	s.VpdId = &v
	return s
}

func (s *GetFabricTopologyResponseBodyContent) Validate() error {
	if s.TopoInfo != nil {
		for _, item := range s.TopoInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFabricTopologyResponseBodyContentTopoInfo struct {
	// The resource name.
	//
	// example:
	//
	// core-1
	LayerName *string `json:"LayerName,omitempty" xml:"LayerName,omitempty"`
	// Hierarchical resource types
	//
	// Valid value:
	//
	// 	- core: core layer.
	//
	// 	- node: Lingjun node.
	//
	// 	- lni: lingjun network interface controller.
	//
	// 	- spine: backbone layer.
	//
	// 	- leaf: access layer
	//
	// example:
	//
	// core
	LayerType *string `json:"LayerType,omitempty" xml:"LayerType,omitempty"`
	// Next Level
	NextLayer []interface{} `json:"NextLayer,omitempty" xml:"NextLayer,omitempty" type:"Repeated"`
}

func (s GetFabricTopologyResponseBodyContentTopoInfo) String() string {
	return dara.Prettify(s)
}

func (s GetFabricTopologyResponseBodyContentTopoInfo) GoString() string {
	return s.String()
}

func (s *GetFabricTopologyResponseBodyContentTopoInfo) GetLayerName() *string {
	return s.LayerName
}

func (s *GetFabricTopologyResponseBodyContentTopoInfo) GetLayerType() *string {
	return s.LayerType
}

func (s *GetFabricTopologyResponseBodyContentTopoInfo) GetNextLayer() []interface{} {
	return s.NextLayer
}

func (s *GetFabricTopologyResponseBodyContentTopoInfo) SetLayerName(v string) *GetFabricTopologyResponseBodyContentTopoInfo {
	s.LayerName = &v
	return s
}

func (s *GetFabricTopologyResponseBodyContentTopoInfo) SetLayerType(v string) *GetFabricTopologyResponseBodyContentTopoInfo {
	s.LayerType = &v
	return s
}

func (s *GetFabricTopologyResponseBodyContentTopoInfo) SetNextLayer(v []interface{}) *GetFabricTopologyResponseBodyContentTopoInfo {
	s.NextLayer = v
	return s
}

func (s *GetFabricTopologyResponseBodyContentTopoInfo) Validate() error {
	return dara.Validate(s)
}
