// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOrgByLayerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrgs(v []*DescribeOrgByLayerResponseBodyOrgs) *DescribeOrgByLayerResponseBody
	GetOrgs() []*DescribeOrgByLayerResponseBodyOrgs
	SetRequestId(v string) *DescribeOrgByLayerResponseBody
	GetRequestId() *string
}

type DescribeOrgByLayerResponseBody struct {
	Orgs []*DescribeOrgByLayerResponseBodyOrgs `json:"Orgs,omitempty" xml:"Orgs,omitempty" type:"Repeated"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOrgByLayerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrgByLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOrgByLayerResponseBody) GetOrgs() []*DescribeOrgByLayerResponseBodyOrgs {
	return s.Orgs
}

func (s *DescribeOrgByLayerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOrgByLayerResponseBody) SetOrgs(v []*DescribeOrgByLayerResponseBodyOrgs) *DescribeOrgByLayerResponseBody {
	s.Orgs = v
	return s
}

func (s *DescribeOrgByLayerResponseBody) SetRequestId(v string) *DescribeOrgByLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOrgByLayerResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeOrgByLayerResponseBodyOrgs struct {
	// example:
	//
	// org-1mox****
	OrgId   *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// example:
	//
	// org-ezqr****
	ParentOrgId *string `json:"ParentOrgId,omitempty" xml:"ParentOrgId,omitempty"`
}

func (s DescribeOrgByLayerResponseBodyOrgs) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrgByLayerResponseBodyOrgs) GoString() string {
	return s.String()
}

func (s *DescribeOrgByLayerResponseBodyOrgs) GetOrgId() *string {
	return s.OrgId
}

func (s *DescribeOrgByLayerResponseBodyOrgs) GetOrgName() *string {
	return s.OrgName
}

func (s *DescribeOrgByLayerResponseBodyOrgs) GetParentOrgId() *string {
	return s.ParentOrgId
}

func (s *DescribeOrgByLayerResponseBodyOrgs) SetOrgId(v string) *DescribeOrgByLayerResponseBodyOrgs {
	s.OrgId = &v
	return s
}

func (s *DescribeOrgByLayerResponseBodyOrgs) SetOrgName(v string) *DescribeOrgByLayerResponseBodyOrgs {
	s.OrgName = &v
	return s
}

func (s *DescribeOrgByLayerResponseBodyOrgs) SetParentOrgId(v string) *DescribeOrgByLayerResponseBodyOrgs {
	s.ParentOrgId = &v
	return s
}

func (s *DescribeOrgByLayerResponseBodyOrgs) Validate() error {
	return dara.Validate(s)
}
