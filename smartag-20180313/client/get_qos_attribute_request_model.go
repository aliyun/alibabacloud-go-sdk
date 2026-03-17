// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQosAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQosId(v string) *GetQosAttributeRequest
	GetQosId() *string
	SetRegionId(v string) *GetQosAttributeRequest
	GetRegionId() *string
}

type GetQosAttributeRequest struct {
	// The ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-1iqifund3gcno5****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the region where the QoS policy is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetQosAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQosAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetQosAttributeRequest) GetQosId() *string {
	return s.QosId
}

func (s *GetQosAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetQosAttributeRequest) SetQosId(v string) *GetQosAttributeRequest {
	s.QosId = &v
	return s
}

func (s *GetQosAttributeRequest) SetRegionId(v string) *GetQosAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *GetQosAttributeRequest) Validate() error {
	return dara.Validate(s)
}
