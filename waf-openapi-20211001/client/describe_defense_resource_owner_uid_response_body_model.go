// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceOwnerUidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerInfos(v []*DescribeDefenseResourceOwnerUidResponseBodyOwnerInfos) *DescribeDefenseResourceOwnerUidResponseBody
	GetOwnerInfos() []*DescribeDefenseResourceOwnerUidResponseBodyOwnerInfos
	SetRequestId(v string) *DescribeDefenseResourceOwnerUidResponseBody
	GetRequestId() *string
}

type DescribeDefenseResourceOwnerUidResponseBody struct {
	OwnerInfos []*DescribeDefenseResourceOwnerUidResponseBodyOwnerInfos `json:"OwnerInfos,omitempty" xml:"OwnerInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 7326952B-B83B-5B7C-84FA-77F3E17310A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDefenseResourceOwnerUidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceOwnerUidResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceOwnerUidResponseBody) GetOwnerInfos() []*DescribeDefenseResourceOwnerUidResponseBodyOwnerInfos {
	return s.OwnerInfos
}

func (s *DescribeDefenseResourceOwnerUidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseResourceOwnerUidResponseBody) SetOwnerInfos(v []*DescribeDefenseResourceOwnerUidResponseBodyOwnerInfos) *DescribeDefenseResourceOwnerUidResponseBody {
	s.OwnerInfos = v
	return s
}

func (s *DescribeDefenseResourceOwnerUidResponseBody) SetRequestId(v string) *DescribeDefenseResourceOwnerUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseResourceOwnerUidResponseBody) Validate() error {
	if s.OwnerInfos != nil {
		for _, item := range s.OwnerInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDefenseResourceOwnerUidResponseBodyOwnerInfos struct {
	// example:
	//
	// 125************21
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// example:
	//
	// a.com-waf
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s DescribeDefenseResourceOwnerUidResponseBodyOwnerInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceOwnerUidResponseBodyOwnerInfos) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceOwnerUidResponseBodyOwnerInfos) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *DescribeDefenseResourceOwnerUidResponseBodyOwnerInfos) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeDefenseResourceOwnerUidResponseBodyOwnerInfos) SetOwnerUserId(v string) *DescribeDefenseResourceOwnerUidResponseBodyOwnerInfos {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeDefenseResourceOwnerUidResponseBodyOwnerInfos) SetResourceName(v string) *DescribeDefenseResourceOwnerUidResponseBodyOwnerInfos {
	s.ResourceName = &v
	return s
}

func (s *DescribeDefenseResourceOwnerUidResponseBodyOwnerInfos) Validate() error {
	return dara.Validate(s)
}
