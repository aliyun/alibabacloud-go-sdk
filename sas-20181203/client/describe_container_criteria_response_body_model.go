// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerCriteriaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCriteriaList(v []*DescribeContainerCriteriaResponseBodyCriteriaList) *DescribeContainerCriteriaResponseBody
	GetCriteriaList() []*DescribeContainerCriteriaResponseBodyCriteriaList
	SetRequestId(v string) *DescribeContainerCriteriaResponseBody
	GetRequestId() *string
}

type DescribeContainerCriteriaResponseBody struct {
	// An array that consists of information about the filter condition.
	CriteriaList []*DescribeContainerCriteriaResponseBodyCriteriaList `json:"CriteriaList,omitempty" xml:"CriteriaList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB3936FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerCriteriaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerCriteriaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerCriteriaResponseBody) GetCriteriaList() []*DescribeContainerCriteriaResponseBodyCriteriaList {
	return s.CriteriaList
}

func (s *DescribeContainerCriteriaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContainerCriteriaResponseBody) SetCriteriaList(v []*DescribeContainerCriteriaResponseBodyCriteriaList) *DescribeContainerCriteriaResponseBody {
	s.CriteriaList = v
	return s
}

func (s *DescribeContainerCriteriaResponseBody) SetRequestId(v string) *DescribeContainerCriteriaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerCriteriaResponseBody) Validate() error {
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

type DescribeContainerCriteriaResponseBodyCriteriaList struct {
	// The name of the search condition. Valid values:
	//
	// 	- **instanceId**: the ID of the container instance.
	//
	// 	- **clusterId**: the cluster ID.
	//
	// 	- **regionId**: the ID of the region in which the container resides.
	//
	// 	- **clusterName**: the name of the cluster.
	//
	// 	- **image**: the name of the image.
	//
	// 	- **imageRepoName**: the name of the image repository.
	//
	// 	- **imageRepoNamespace**: the namespace of the image repository.
	//
	// 	- **imageRepoTag**: the tag of the image repository.
	//
	// 	- **imageDigest**: the image digest.
	//
	// 	- **ClusterType**: the type of the cluster.
	//
	// 	- **hostIp**: the public IP address.
	//
	// 	- **pod**: the pod.
	//
	// 	- **podIp**: the IP address of the pod.
	//
	// 	- **containerId**: the container ID.
	//
	// 	- **vulStatus**: indicates whether vulnerabilities exist in the container.
	//
	// 	- **alarmStatus**: indicates whether alerts are generated for the container.
	//
	// 	- **riskStatus**: indicates whether risks exist in the container.
	//
	// 	- **riskLevel**: the risk level of the container.
	//
	// 	- **containerScope**: the type of the container.
	//
	// example:
	//
	// clusterId
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
	// The values of the search condition. This parameter is returned only if the value of **Type*	- is set to **select**.
	//
	// > If the value of **Type*	- is set to **input**, the return value of this parameter is empty.
	//
	// example:
	//
	// ManagedKubernetes,NotManagedKubernetes,PrivateKubernetes
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeContainerCriteriaResponseBodyCriteriaList) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerCriteriaResponseBodyCriteriaList) GoString() string {
	return s.String()
}

func (s *DescribeContainerCriteriaResponseBodyCriteriaList) GetName() *string {
	return s.Name
}

func (s *DescribeContainerCriteriaResponseBodyCriteriaList) GetType() *string {
	return s.Type
}

func (s *DescribeContainerCriteriaResponseBodyCriteriaList) GetValues() *string {
	return s.Values
}

func (s *DescribeContainerCriteriaResponseBodyCriteriaList) SetName(v string) *DescribeContainerCriteriaResponseBodyCriteriaList {
	s.Name = &v
	return s
}

func (s *DescribeContainerCriteriaResponseBodyCriteriaList) SetType(v string) *DescribeContainerCriteriaResponseBodyCriteriaList {
	s.Type = &v
	return s
}

func (s *DescribeContainerCriteriaResponseBodyCriteriaList) SetValues(v string) *DescribeContainerCriteriaResponseBodyCriteriaList {
	s.Values = &v
	return s
}

func (s *DescribeContainerCriteriaResponseBodyCriteriaList) Validate() error {
	return dara.Validate(s)
}
