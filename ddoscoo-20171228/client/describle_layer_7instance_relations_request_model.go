// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribleLayer7InstanceRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainList(v []*string) *DescribleLayer7InstanceRelationsRequest
	GetDomainList() []*string
	SetResourceGroupId(v string) *DescribleLayer7InstanceRelationsRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribleLayer7InstanceRelationsRequest
	GetSourceIp() *string
}

type DescribleLayer7InstanceRelationsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	DomainList []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribleLayer7InstanceRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribleLayer7InstanceRelationsRequest) GoString() string {
	return s.String()
}

func (s *DescribleLayer7InstanceRelationsRequest) GetDomainList() []*string {
	return s.DomainList
}

func (s *DescribleLayer7InstanceRelationsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribleLayer7InstanceRelationsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribleLayer7InstanceRelationsRequest) SetDomainList(v []*string) *DescribleLayer7InstanceRelationsRequest {
	s.DomainList = v
	return s
}

func (s *DescribleLayer7InstanceRelationsRequest) SetResourceGroupId(v string) *DescribleLayer7InstanceRelationsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsRequest) SetSourceIp(v string) *DescribleLayer7InstanceRelationsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsRequest) Validate() error {
	return dara.Validate(s)
}
