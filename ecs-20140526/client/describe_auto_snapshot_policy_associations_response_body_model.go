// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotPolicyAssociationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicyAssociations(v *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations) *DescribeAutoSnapshotPolicyAssociationsResponseBody
	GetAutoSnapshotPolicyAssociations() *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations
	SetNextToken(v string) *DescribeAutoSnapshotPolicyAssociationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeAutoSnapshotPolicyAssociationsResponseBody
	GetRequestId() *string
}

type DescribeAutoSnapshotPolicyAssociationsResponseBody struct {
	AutoSnapshotPolicyAssociations *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations `json:"AutoSnapshotPolicyAssociations,omitempty" xml:"AutoSnapshotPolicyAssociations,omitempty" type:"Struct"`
	NextToken                      *string                                                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                      *string                                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAutoSnapshotPolicyAssociationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBody) GetAutoSnapshotPolicyAssociations() *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations {
	return s.AutoSnapshotPolicyAssociations
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBody) SetAutoSnapshotPolicyAssociations(v *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations) *DescribeAutoSnapshotPolicyAssociationsResponseBody {
	s.AutoSnapshotPolicyAssociations = v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBody) SetNextToken(v string) *DescribeAutoSnapshotPolicyAssociationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBody) SetRequestId(v string) *DescribeAutoSnapshotPolicyAssociationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBody) Validate() error {
	if s.AutoSnapshotPolicyAssociations != nil {
		if err := s.AutoSnapshotPolicyAssociations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations struct {
	AutoSnapshotPolicyAssociation []*DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation `json:"AutoSnapshotPolicyAssociation,omitempty" xml:"AutoSnapshotPolicyAssociation,omitempty" type:"Repeated"`
}

func (s DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations) GetAutoSnapshotPolicyAssociation() []*DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation {
	return s.AutoSnapshotPolicyAssociation
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations) SetAutoSnapshotPolicyAssociation(v []*DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation) *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations {
	s.AutoSnapshotPolicyAssociation = v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations) Validate() error {
	if s.AutoSnapshotPolicyAssociation != nil {
		for _, item := range s.AutoSnapshotPolicyAssociation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation struct {
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	DiskId               *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
}

func (s DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation) SetDiskId(v string) *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation {
	s.DiskId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation) Validate() error {
	return dara.Validate(s)
}
