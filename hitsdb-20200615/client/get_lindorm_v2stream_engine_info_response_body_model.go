// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2StreamEngineInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetLindormV2StreamEngineInfoResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *GetLindormV2StreamEngineInfoResponseBody
	GetRequestId() *string
	SetResourceGroupList(v []*GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) *GetLindormV2StreamEngineInfoResponseBody
	GetResourceGroupList() []*GetLindormV2StreamEngineInfoResponseBodyResourceGroupList
}

type GetLindormV2StreamEngineInfoResponseBody struct {
	InstanceId        *string                                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId         *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupList []*GetLindormV2StreamEngineInfoResponseBodyResourceGroupList `json:"ResourceGroupList,omitempty" xml:"ResourceGroupList,omitempty" type:"Repeated"`
}

func (s GetLindormV2StreamEngineInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2StreamEngineInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormV2StreamEngineInfoResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormV2StreamEngineInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLindormV2StreamEngineInfoResponseBody) GetResourceGroupList() []*GetLindormV2StreamEngineInfoResponseBodyResourceGroupList {
	return s.ResourceGroupList
}

func (s *GetLindormV2StreamEngineInfoResponseBody) SetInstanceId(v string) *GetLindormV2StreamEngineInfoResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2StreamEngineInfoResponseBody) SetRequestId(v string) *GetLindormV2StreamEngineInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormV2StreamEngineInfoResponseBody) SetResourceGroupList(v []*GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) *GetLindormV2StreamEngineInfoResponseBody {
	s.ResourceGroupList = v
	return s
}

func (s *GetLindormV2StreamEngineInfoResponseBody) Validate() error {
	if s.ResourceGroupList != nil {
		for _, item := range s.ResourceGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLindormV2StreamEngineInfoResponseBodyResourceGroupList struct {
	JmIpList          []*string `json:"JmIpList,omitempty" xml:"JmIpList,omitempty" type:"Repeated"`
	Quantity          *int32    `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ResourceGroupName *string   `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	SgIpList          []*string `json:"SgIpList,omitempty" xml:"SgIpList,omitempty" type:"Repeated"`
	Spec              *string   `json:"Spec,omitempty" xml:"Spec,omitempty"`
	SpecId            *string   `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	Status            *string   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) GoString() string {
	return s.String()
}

func (s *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) GetJmIpList() []*string {
	return s.JmIpList
}

func (s *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) GetQuantity() *int32 {
	return s.Quantity
}

func (s *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) GetSgIpList() []*string {
	return s.SgIpList
}

func (s *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) GetSpec() *string {
	return s.Spec
}

func (s *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) GetSpecId() *string {
	return s.SpecId
}

func (s *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) GetStatus() *string {
	return s.Status
}

func (s *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) SetJmIpList(v []*string) *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList {
	s.JmIpList = v
	return s
}

func (s *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) SetQuantity(v int32) *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList {
	s.Quantity = &v
	return s
}

func (s *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) SetResourceGroupName(v string) *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList {
	s.ResourceGroupName = &v
	return s
}

func (s *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) SetSgIpList(v []*string) *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList {
	s.SgIpList = v
	return s
}

func (s *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) SetSpec(v string) *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList {
	s.Spec = &v
	return s
}

func (s *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) SetSpecId(v string) *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList {
	s.SpecId = &v
	return s
}

func (s *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) SetStatus(v string) *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList {
	s.Status = &v
	return s
}

func (s *GetLindormV2StreamEngineInfoResponseBodyResourceGroupList) Validate() error {
	return dara.Validate(s)
}
