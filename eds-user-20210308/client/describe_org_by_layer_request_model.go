// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOrgByLayerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrgName(v string) *DescribeOrgByLayerRequest
	GetOrgName() *string
	SetParentOrgId(v string) *DescribeOrgByLayerRequest
	GetParentOrgId() *string
}

type DescribeOrgByLayerRequest struct {
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// example:
	//
	// org-11fs****
	ParentOrgId *string `json:"ParentOrgId,omitempty" xml:"ParentOrgId,omitempty"`
}

func (s DescribeOrgByLayerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrgByLayerRequest) GoString() string {
	return s.String()
}

func (s *DescribeOrgByLayerRequest) GetOrgName() *string {
	return s.OrgName
}

func (s *DescribeOrgByLayerRequest) GetParentOrgId() *string {
	return s.ParentOrgId
}

func (s *DescribeOrgByLayerRequest) SetOrgName(v string) *DescribeOrgByLayerRequest {
	s.OrgName = &v
	return s
}

func (s *DescribeOrgByLayerRequest) SetParentOrgId(v string) *DescribeOrgByLayerRequest {
	s.ParentOrgId = &v
	return s
}

func (s *DescribeOrgByLayerRequest) Validate() error {
	return dara.Validate(s)
}
