// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebInstanceRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeWebInstanceRelationsResponseBody
	GetRequestId() *string
	SetWebInstanceRelations(v []*DescribeWebInstanceRelationsResponseBodyWebInstanceRelations) *DescribeWebInstanceRelationsResponseBody
	GetWebInstanceRelations() []*DescribeWebInstanceRelationsResponseBodyWebInstanceRelations
}

type DescribeWebInstanceRelationsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0222382B-5FE5-4FF7-BC9B-97EE31D58818
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the instances to which a website service is added.
	WebInstanceRelations []*DescribeWebInstanceRelationsResponseBodyWebInstanceRelations `json:"WebInstanceRelations,omitempty" xml:"WebInstanceRelations,omitempty" type:"Repeated"`
}

func (s DescribeWebInstanceRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebInstanceRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebInstanceRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebInstanceRelationsResponseBody) GetWebInstanceRelations() []*DescribeWebInstanceRelationsResponseBodyWebInstanceRelations {
	return s.WebInstanceRelations
}

func (s *DescribeWebInstanceRelationsResponseBody) SetRequestId(v string) *DescribeWebInstanceRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebInstanceRelationsResponseBody) SetWebInstanceRelations(v []*DescribeWebInstanceRelationsResponseBodyWebInstanceRelations) *DescribeWebInstanceRelationsResponseBody {
	s.WebInstanceRelations = v
	return s
}

func (s *DescribeWebInstanceRelationsResponseBody) Validate() error {
	if s.WebInstanceRelations != nil {
		for _, item := range s.WebInstanceRelations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebInstanceRelationsResponseBodyWebInstanceRelations struct {
	// The domain name of the website.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The information about the instance to which a website service is added.
	InstanceDetails []*DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails `json:"InstanceDetails,omitempty" xml:"InstanceDetails,omitempty" type:"Repeated"`
}

func (s DescribeWebInstanceRelationsResponseBodyWebInstanceRelations) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebInstanceRelationsResponseBodyWebInstanceRelations) GoString() string {
	return s.String()
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelations) GetDomain() *string {
	return s.Domain
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelations) GetInstanceDetails() []*DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails {
	return s.InstanceDetails
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelations) SetDomain(v string) *DescribeWebInstanceRelationsResponseBodyWebInstanceRelations {
	s.Domain = &v
	return s
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelations) SetInstanceDetails(v []*DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails) *DescribeWebInstanceRelationsResponseBodyWebInstanceRelations {
	s.InstanceDetails = v
	return s
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelations) Validate() error {
	if s.InstanceDetails != nil {
		for _, item := range s.InstanceDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails struct {
	// The IP addresses of the instance.
	EipList []*string `json:"EipList,omitempty" xml:"EipList,omitempty" type:"Repeated"`
	// The function plan of the instance. Valid values:
	//
	// 	- **default**: Standard function plan
	//
	// 	- **enhance**: Enhanced function plan
	//
	// example:
	//
	// enhance
	FunctionVersion *string `json:"FunctionVersion,omitempty" xml:"FunctionVersion,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ddoscoo-cn-0pp163pd****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails) GoString() string {
	return s.String()
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails) GetEipList() []*string {
	return s.EipList
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails) GetFunctionVersion() *string {
	return s.FunctionVersion
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails) SetEipList(v []*string) *DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails {
	s.EipList = v
	return s
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails) SetFunctionVersion(v string) *DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails {
	s.FunctionVersion = &v
	return s
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails) SetInstanceId(v string) *DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails {
	s.InstanceId = &v
	return s
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails) Validate() error {
	return dara.Validate(s)
}
