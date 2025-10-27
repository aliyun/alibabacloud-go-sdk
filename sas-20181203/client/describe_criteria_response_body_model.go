// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCriteriaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCriteriaList(v []*DescribeCriteriaResponseBodyCriteriaList) *DescribeCriteriaResponseBody
	GetCriteriaList() []*DescribeCriteriaResponseBodyCriteriaList
	SetRequestId(v string) *DescribeCriteriaResponseBody
	GetRequestId() *string
}

type DescribeCriteriaResponseBody struct {
	// The information about the search conditions of assets.
	CriteriaList []*DescribeCriteriaResponseBodyCriteriaList `json:"CriteriaList,omitempty" xml:"CriteriaList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8E6DDACF-99AF-5939-AFFD-FCCD3B01E724
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCriteriaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCriteriaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCriteriaResponseBody) GetCriteriaList() []*DescribeCriteriaResponseBodyCriteriaList {
	return s.CriteriaList
}

func (s *DescribeCriteriaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCriteriaResponseBody) SetCriteriaList(v []*DescribeCriteriaResponseBodyCriteriaList) *DescribeCriteriaResponseBody {
	s.CriteriaList = v
	return s
}

func (s *DescribeCriteriaResponseBody) SetRequestId(v string) *DescribeCriteriaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCriteriaResponseBody) Validate() error {
	if s.CriteriaList != nil {
		for _, item := range s.CriteriaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCriteriaResponseBodyCriteriaList struct {
	// The structured attribute values of the assets that match the keyword. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **vendor**: providers.
	//
	// 	- **regionIds**: IDs of supported regions
	//
	// example:
	//
	// [{"vendor":0,"regionIds":{"default":["ap-southeast-1","ap-northeast-2","ap-southeast-3","ap-southeast-5","ap-southeast-7","me-central-1"]}},{"vendor":1,"regionIds":{"default":["outside-of-aliyun"]}}]
	MultiValues *string `json:"MultiValues,omitempty" xml:"MultiValues,omitempty"`
	// The name of the search condition. Valid values:
	//
	// 	- **internetIp**: the public IP address.
	//
	// 	- **intranetIp**: the private IP address.
	//
	// 	- **instanceName**: the name of the instance.
	//
	// 	- **instanceId**: the instance ID.
	//
	// 	- **vpcInstanceId**: The ID of the virtual private cloud (VPC) to which the instance belongs.
	//
	// 	- **osName**: the operating system.
	//
	// 	- **osType**: the operating system type.
	//
	// 	- **hcStatus**: indicates whether baseline risks exist.
	//
	// 	- **vulStatus**: indicates whether vulnerabilities exist.
	//
	// 	- **alarmStatus**: indicates whether security alerts exist.
	//
	// 	- **riskStatus**: indicates whether risks exist.
	//
	// 	- **clientStatus**: indicates the status of the client.
	//
	// 	- **runningStatus**: the running status of the asset.
	//
	// 	- **tagName**: the name of the tag.
	//
	// 	- **groupName**: the name of the server group.
	//
	// 	- **regionId**: the region ID.
	//
	// 	- **importance**: the importance of the asset.
	//
	// 	- **exposedStatus**: indicates whether the server is exposed.
	//
	// 	- **authVersion**: the authorization version.
	//
	// 	- **flag**: the cloud service provider.
	//
	// 	- **ipList**: the IP address list.
	//
	// 	- **uuidList*	- :the UUID.
	//
	// 	- **tagKeyValue**: the ECS tag.
	//
	// 	- **vendorAuthAlias**: the account name.
	//
	// example:
	//
	// internetIp
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the search condition. Valid values:
	//
	// 	- **input**: The search condition needs to be specified.
	//
	// 	- **select**: The search condition is an option that can be selected from the drop-down list.
	//
	// example:
	//
	// input
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The attribute values of the assets that match the keyword.
	//
	// example:
	//
	// 47.96.XX.XX
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeCriteriaResponseBodyCriteriaList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCriteriaResponseBodyCriteriaList) GoString() string {
	return s.String()
}

func (s *DescribeCriteriaResponseBodyCriteriaList) GetMultiValues() *string {
	return s.MultiValues
}

func (s *DescribeCriteriaResponseBodyCriteriaList) GetName() *string {
	return s.Name
}

func (s *DescribeCriteriaResponseBodyCriteriaList) GetType() *string {
	return s.Type
}

func (s *DescribeCriteriaResponseBodyCriteriaList) GetValues() *string {
	return s.Values
}

func (s *DescribeCriteriaResponseBodyCriteriaList) SetMultiValues(v string) *DescribeCriteriaResponseBodyCriteriaList {
	s.MultiValues = &v
	return s
}

func (s *DescribeCriteriaResponseBodyCriteriaList) SetName(v string) *DescribeCriteriaResponseBodyCriteriaList {
	s.Name = &v
	return s
}

func (s *DescribeCriteriaResponseBodyCriteriaList) SetType(v string) *DescribeCriteriaResponseBodyCriteriaList {
	s.Type = &v
	return s
}

func (s *DescribeCriteriaResponseBodyCriteriaList) SetValues(v string) *DescribeCriteriaResponseBodyCriteriaList {
	s.Values = &v
	return s
}

func (s *DescribeCriteriaResponseBodyCriteriaList) Validate() error {
	return dara.Validate(s)
}
