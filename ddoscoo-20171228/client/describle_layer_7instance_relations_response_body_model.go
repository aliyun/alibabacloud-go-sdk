// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribleLayer7InstanceRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLayer7InstanceRelations(v []*DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations) *DescribleLayer7InstanceRelationsResponseBody
	GetLayer7InstanceRelations() []*DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations
	SetRequestId(v string) *DescribleLayer7InstanceRelationsResponseBody
	GetRequestId() *string
}

type DescribleLayer7InstanceRelationsResponseBody struct {
	Layer7InstanceRelations []*DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations `json:"Layer7InstanceRelations,omitempty" xml:"Layer7InstanceRelations,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribleLayer7InstanceRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribleLayer7InstanceRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribleLayer7InstanceRelationsResponseBody) GetLayer7InstanceRelations() []*DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations {
	return s.Layer7InstanceRelations
}

func (s *DescribleLayer7InstanceRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribleLayer7InstanceRelationsResponseBody) SetLayer7InstanceRelations(v []*DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations) *DescribleLayer7InstanceRelationsResponseBody {
	s.Layer7InstanceRelations = v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseBody) SetRequestId(v string) *DescribleLayer7InstanceRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations struct {
	// example:
	//
	// www.aliyun.com
	Domain          *string                                                                               `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceDetails []*DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails `json:"InstanceDetails,omitempty" xml:"InstanceDetails,omitempty" type:"Repeated"`
}

func (s DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations) String() string {
	return dara.Prettify(s)
}

func (s DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations) GoString() string {
	return s.String()
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations) GetDomain() *string {
	return s.Domain
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations) GetInstanceDetails() []*DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails {
	return s.InstanceDetails
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations) SetDomain(v string) *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations {
	s.Domain = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations) SetInstanceDetails(v []*DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations {
	s.InstanceDetails = v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations) Validate() error {
	return dara.Validate(s)
}

type DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails struct {
	EipList []*string `json:"EipList,omitempty" xml:"EipList,omitempty" type:"Repeated"`
	// example:
	//
	// default
	FunctionVersion *string `json:"FunctionVersion,omitempty" xml:"FunctionVersion,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpMode     *string `json:"IpMode,omitempty" xml:"IpMode,omitempty"`
	IpVersion  *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
}

func (s DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) GoString() string {
	return s.String()
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) GetEipList() []*string {
	return s.EipList
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) GetFunctionVersion() *string {
	return s.FunctionVersion
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) GetIpMode() *string {
	return s.IpMode
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) SetEipList(v []*string) *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails {
	s.EipList = v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) SetFunctionVersion(v string) *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails {
	s.FunctionVersion = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) SetInstanceId(v string) *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails {
	s.InstanceId = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) SetIpMode(v string) *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails {
	s.IpMode = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) SetIpVersion(v string) *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails {
	s.IpVersion = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) Validate() error {
	return dara.Validate(s)
}
