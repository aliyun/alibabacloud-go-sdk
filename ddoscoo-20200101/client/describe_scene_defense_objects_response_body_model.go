// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSceneDefenseObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetObjects(v []*DescribeSceneDefenseObjectsResponseBodyObjects) *DescribeSceneDefenseObjectsResponseBody
	GetObjects() []*DescribeSceneDefenseObjectsResponseBodyObjects
	SetRequestId(v string) *DescribeSceneDefenseObjectsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSceneDefenseObjectsResponseBody
	GetSuccess() *bool
}

type DescribeSceneDefenseObjectsResponseBody struct {
	// The information about the protected assets.
	Objects []*DescribeSceneDefenseObjectsResponseBodyObjects `json:"Objects,omitempty" xml:"Objects,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// FE07E73F-F19E-4A51-B62F-AC59E3B962D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSceneDefenseObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneDefenseObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefenseObjectsResponseBody) GetObjects() []*DescribeSceneDefenseObjectsResponseBodyObjects {
	return s.Objects
}

func (s *DescribeSceneDefenseObjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSceneDefenseObjectsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSceneDefenseObjectsResponseBody) SetObjects(v []*DescribeSceneDefenseObjectsResponseBodyObjects) *DescribeSceneDefenseObjectsResponseBody {
	s.Objects = v
	return s
}

func (s *DescribeSceneDefenseObjectsResponseBody) SetRequestId(v string) *DescribeSceneDefenseObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSceneDefenseObjectsResponseBody) SetSuccess(v bool) *DescribeSceneDefenseObjectsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSceneDefenseObjectsResponseBody) Validate() error {
	if s.Objects != nil {
		for _, item := range s.Objects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSceneDefenseObjectsResponseBodyObjects struct {
	// The domain name that is protected by the policy.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// 47e07ebd-0ba5-4afc-957b-59d15b90****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The IP address of the Anti-DDoS Pro or Anti-DDoS Premium instance that is protected by the policy.
	//
	// example:
	//
	// 203.XX.XX.119
	Vip *string `json:"Vip,omitempty" xml:"Vip,omitempty"`
}

func (s DescribeSceneDefenseObjectsResponseBodyObjects) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneDefenseObjectsResponseBodyObjects) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefenseObjectsResponseBodyObjects) GetDomain() *string {
	return s.Domain
}

func (s *DescribeSceneDefenseObjectsResponseBodyObjects) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeSceneDefenseObjectsResponseBodyObjects) GetVip() *string {
	return s.Vip
}

func (s *DescribeSceneDefenseObjectsResponseBodyObjects) SetDomain(v string) *DescribeSceneDefenseObjectsResponseBodyObjects {
	s.Domain = &v
	return s
}

func (s *DescribeSceneDefenseObjectsResponseBodyObjects) SetPolicyId(v string) *DescribeSceneDefenseObjectsResponseBodyObjects {
	s.PolicyId = &v
	return s
}

func (s *DescribeSceneDefenseObjectsResponseBodyObjects) SetVip(v string) *DescribeSceneDefenseObjectsResponseBodyObjects {
	s.Vip = &v
	return s
}

func (s *DescribeSceneDefenseObjectsResponseBodyObjects) Validate() error {
	return dara.Validate(s)
}
