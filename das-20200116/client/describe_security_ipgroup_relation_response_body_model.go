// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityIPGroupRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSecurityIPGroupRelationResponseBody
	GetCode() *string
	SetData(v *DescribeSecurityIPGroupRelationResponseBodyData) *DescribeSecurityIPGroupRelationResponseBody
	GetData() *DescribeSecurityIPGroupRelationResponseBodyData
	SetMessage(v string) *DescribeSecurityIPGroupRelationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSecurityIPGroupRelationResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSecurityIPGroupRelationResponseBody
	GetSuccess() *string
}

type DescribeSecurityIPGroupRelationResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// ListResult<InstanceSSL>
	Data *DescribeSecurityIPGroupRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned for the request.
	//
	// > If the request is successful, **Successful*	- is returned. If the request fails, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A6D1C8EE-013C-541F-83EB-B13C8xxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSecurityIPGroupRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPGroupRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSecurityIPGroupRelationResponseBody) GetData() *DescribeSecurityIPGroupRelationResponseBodyData {
	return s.Data
}

func (s *DescribeSecurityIPGroupRelationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSecurityIPGroupRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityIPGroupRelationResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSecurityIPGroupRelationResponseBody) SetCode(v string) *DescribeSecurityIPGroupRelationResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSecurityIPGroupRelationResponseBody) SetData(v *DescribeSecurityIPGroupRelationResponseBodyData) *DescribeSecurityIPGroupRelationResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSecurityIPGroupRelationResponseBody) SetMessage(v string) *DescribeSecurityIPGroupRelationResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSecurityIPGroupRelationResponseBody) SetRequestId(v string) *DescribeSecurityIPGroupRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityIPGroupRelationResponseBody) SetSuccess(v string) *DescribeSecurityIPGroupRelationResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSecurityIPGroupRelationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityIPGroupRelationResponseBodyData struct {
	// The global IP whitelist template information.
	GlobalSecurityIPGroupRel []*DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel `json:"GlobalSecurityIPGroupRel,omitempty" xml:"GlobalSecurityIPGroupRel,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeSecurityIPGroupRelationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPGroupRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPGroupRelationResponseBodyData) GetGlobalSecurityIPGroupRel() []*DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel {
	return s.GlobalSecurityIPGroupRel
}

func (s *DescribeSecurityIPGroupRelationResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSecurityIPGroupRelationResponseBodyData) SetGlobalSecurityIPGroupRel(v []*DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) *DescribeSecurityIPGroupRelationResponseBodyData {
	s.GlobalSecurityIPGroupRel = v
	return s
}

func (s *DescribeSecurityIPGroupRelationResponseBodyData) SetInstanceId(v string) *DescribeSecurityIPGroupRelationResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeSecurityIPGroupRelationResponseBodyData) Validate() error {
	if s.GlobalSecurityIPGroupRel != nil {
		for _, item := range s.GlobalSecurityIPGroupRel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel struct {
	// The IP addresses in the whitelist template.
	//
	// > Separate multiple IP addresses with commas (,). You can add up to 1,000 IP addresses or CIDR blocks across all IP whitelists.
	//
	// example:
	//
	// 192.168.1.28/32
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// The name of the IP whitelist template.
	//
	// example:
	//
	// test2
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	// The ID of the IP whitelist template.
	//
	// example:
	//
	// g-1no2rzybnqcv0xxxxxx
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) GetGIpList() *string {
	return s.GIpList
}

func (s *DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) SetGIpList(v string) *DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel {
	s.GIpList = &v
	return s
}

func (s *DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) SetGlobalIgName(v string) *DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel {
	s.GlobalIgName = &v
	return s
}

func (s *DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) SetGlobalSecurityGroupId(v string) *DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) SetRegionId(v string) *DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) Validate() error {
	return dara.Validate(s)
}
