// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppInstanceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAppInstanceListResponseBody
	GetCode() *int32
	SetInstanceList(v []*DescribeAppInstanceListResponseBodyInstanceList) *DescribeAppInstanceListResponseBody
	GetInstanceList() []*DescribeAppInstanceListResponseBodyInstanceList
	SetMessage(v string) *DescribeAppInstanceListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAppInstanceListResponseBody
	GetRequestId() *string
}

type DescribeAppInstanceListResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The application instances.
	InstanceList []*DescribeAppInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 03FD1520-0FD6-436A-****-265318D7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAppInstanceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppInstanceListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAppInstanceListResponseBody) GetInstanceList() []*DescribeAppInstanceListResponseBodyInstanceList {
	return s.InstanceList
}

func (s *DescribeAppInstanceListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAppInstanceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppInstanceListResponseBody) SetCode(v int32) *DescribeAppInstanceListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAppInstanceListResponseBody) SetInstanceList(v []*DescribeAppInstanceListResponseBodyInstanceList) *DescribeAppInstanceListResponseBody {
	s.InstanceList = v
	return s
}

func (s *DescribeAppInstanceListResponseBody) SetMessage(v string) *DescribeAppInstanceListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAppInstanceListResponseBody) SetRequestId(v string) *DescribeAppInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppInstanceListResponseBody) Validate() error {
	if s.InstanceList != nil {
		for _, item := range s.InstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppInstanceListResponseBodyInstanceList struct {
	// The ID of the application.
	//
	// example:
	//
	// 93fdd228-*****-ed2ae98de18d
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Indicates whether the application was released in canary release mode.
	//
	// 	- `true`: The application was released in canary release mode.
	//
	// 	- `false`: The application was not released in canary release mode
	//
	// example:
	//
	// false
	Canary *bool `json:"Canary,omitempty" xml:"Canary,omitempty"`
	// The ID of the instance group to which the application is deployed.
	//
	// example:
	//
	// 93fdd228-*****-ed2ae98de18d
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the instance group to which the application is deployed.
	//
	// example:
	//
	// _DEFAULT_GROUP
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The labels of the node. The value is a JSON string.
	//
	// example:
	//
	// {"alibabacloud.com/nodepool-id":"np0*5b9377fa907","beta.kubernetes.io/arch":"amd64","beta.kubernetes.io/instance-type":"ecs.*","beta.kubernetes.io/os":"linux","failure-domain.beta.kubernetes.io/region":"cn-hangzhou","failure-domain.beta.kubernetes.io/zone":"cn-hangzhou-b","kubernetes.io/arch":"amd64","kubernetes.io/hostname":"cn-hangzhou*","kubernetes.io/os":"linux","node.kubernetes.io/instance-type":"ecs.*","topology.diskplugin.csi.alibabacloud.com/zone":"cn-hangzhou-b","topology.kubernetes.io/region":"cn-hangzhou","topology.kubernetes.io/zone":"cn-hangzhou-b"}
	NodeLabels *string `json:"NodeLabels,omitempty" xml:"NodeLabels,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// cn-hangzhou.192.168.0.*
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The information about the pod. The value is a JSON string.
	//
	// example:
	//
	// {"metadata":{"name":"oambuild-group-1-*4xthz","generateName":"oambuild-group-*96-","namespace":"default","selfLink":"/api/v1/namespaces/default/pods/oambuild-grou*96-4xthz","uid":"7a23399c-*fe7ff4018","resourceVersion":"969614830","creationTimestamp":"2021-04-06T11:38:46Z","labels":{"ARMSApmAppId":"*","ARMSApmLicenseKey":"*"...
	PodRaw *string `json:"PodRaw,omitempty" xml:"PodRaw,omitempty"`
	// The deployment package version of the node.
	//
	// example:
	//
	// 2021-04-06 19:37:42
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAppInstanceListResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppInstanceListResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) GetCanary() *bool {
	return s.Canary
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) GetNodeLabels() *string {
	return s.NodeLabels
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) GetPodRaw() *string {
	return s.PodRaw
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) GetVersion() *string {
	return s.Version
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) SetAppId(v string) *DescribeAppInstanceListResponseBodyInstanceList {
	s.AppId = &v
	return s
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) SetCanary(v bool) *DescribeAppInstanceListResponseBodyInstanceList {
	s.Canary = &v
	return s
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) SetGroupId(v string) *DescribeAppInstanceListResponseBodyInstanceList {
	s.GroupId = &v
	return s
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) SetGroupName(v string) *DescribeAppInstanceListResponseBodyInstanceList {
	s.GroupName = &v
	return s
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) SetNodeLabels(v string) *DescribeAppInstanceListResponseBodyInstanceList {
	s.NodeLabels = &v
	return s
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) SetNodeName(v string) *DescribeAppInstanceListResponseBodyInstanceList {
	s.NodeName = &v
	return s
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) SetPodRaw(v string) *DescribeAppInstanceListResponseBodyInstanceList {
	s.PodRaw = &v
	return s
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) SetVersion(v string) *DescribeAppInstanceListResponseBodyInstanceList {
	s.Version = &v
	return s
}

func (s *DescribeAppInstanceListResponseBodyInstanceList) Validate() error {
	return dara.Validate(s)
}
