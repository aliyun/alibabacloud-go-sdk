// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEiamInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ListEiamInstancesResponseBodyInstances) *ListEiamInstancesResponseBody
	GetInstances() []*ListEiamInstancesResponseBodyInstances
	SetRequestId(v string) *ListEiamInstancesResponseBody
	GetRequestId() *string
}

type ListEiamInstancesResponseBody struct {
	// The instance list.
	Instances []*ListEiamInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEiamInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEiamInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEiamInstancesResponseBody) GetInstances() []*ListEiamInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListEiamInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEiamInstancesResponseBody) SetInstances(v []*ListEiamInstancesResponseBodyInstances) *ListEiamInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListEiamInstancesResponseBody) SetRequestId(v string) *ListEiamInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEiamInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEiamInstancesResponseBodyInstances struct {
	// The instance description.
	//
	// example:
	//
	// instance test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The private domain name of the instance Developer API.
	//
	// example:
	//
	// eiam-developerapi-cn.vpc-proxy.aliyuncs.com
	DeveloperAPIPrivateDomain *string `json:"DeveloperAPIPrivateDomain,omitempty" xml:"DeveloperAPIPrivateDomain,omitempty"`
	// The public domain of the instance Developer API.
	//
	// example:
	//
	// eiam-developerapi.cn-hangzhou.aliyuncs.com
	DeveloperAPIPublicDomain *string `json:"DeveloperAPIPublicDomain,omitempty" xml:"DeveloperAPIPublicDomain,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_eypq6ljgyeuwmlw672sulxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance status.
	//
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The instance version.
	//
	// Valid values:
	//
	// 	- EIAM 2.0
	//
	// 	- EIAM 1.0
	//
	// example:
	//
	// EIAM 2.0
	InstanceVersion *string `json:"InstanceVersion,omitempty" xml:"InstanceVersion,omitempty"`
	// The private domain of the instance OpenAPI.
	//
	// example:
	//
	// eiam-cn.vpc-proxy.aliyuncs.com
	OpenAPIPrivateDomain *string `json:"OpenAPIPrivateDomain,omitempty" xml:"OpenAPIPrivateDomain,omitempty"`
	// The public domain of the instance OpenAPI.
	//
	// example:
	//
	// eiam.cn-hangzhou.aliyuncs.com
	OpenAPIPublicDomain *string `json:"OpenAPIPublicDomain,omitempty" xml:"OpenAPIPublicDomain,omitempty"`
	// The single sign-on (SSO) domain  of the instance.
	//
	// example:
	//
	// xxxx.aliyunidaas.com
	SSODomain *string `json:"SSODomain,omitempty" xml:"SSODomain,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 1677810869300
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListEiamInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListEiamInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListEiamInstancesResponseBodyInstances) GetDescription() *string {
	return s.Description
}

func (s *ListEiamInstancesResponseBodyInstances) GetDeveloperAPIPrivateDomain() *string {
	return s.DeveloperAPIPrivateDomain
}

func (s *ListEiamInstancesResponseBodyInstances) GetDeveloperAPIPublicDomain() *string {
	return s.DeveloperAPIPublicDomain
}

func (s *ListEiamInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEiamInstancesResponseBodyInstances) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *ListEiamInstancesResponseBodyInstances) GetInstanceVersion() *string {
	return s.InstanceVersion
}

func (s *ListEiamInstancesResponseBodyInstances) GetOpenAPIPrivateDomain() *string {
	return s.OpenAPIPrivateDomain
}

func (s *ListEiamInstancesResponseBodyInstances) GetOpenAPIPublicDomain() *string {
	return s.OpenAPIPublicDomain
}

func (s *ListEiamInstancesResponseBodyInstances) GetSSODomain() *string {
	return s.SSODomain
}

func (s *ListEiamInstancesResponseBodyInstances) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListEiamInstancesResponseBodyInstances) SetDescription(v string) *ListEiamInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetDeveloperAPIPrivateDomain(v string) *ListEiamInstancesResponseBodyInstances {
	s.DeveloperAPIPrivateDomain = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetDeveloperAPIPublicDomain(v string) *ListEiamInstancesResponseBodyInstances {
	s.DeveloperAPIPublicDomain = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetInstanceId(v string) *ListEiamInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetInstanceStatus(v string) *ListEiamInstancesResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetInstanceVersion(v string) *ListEiamInstancesResponseBodyInstances {
	s.InstanceVersion = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetOpenAPIPrivateDomain(v string) *ListEiamInstancesResponseBodyInstances {
	s.OpenAPIPrivateDomain = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetOpenAPIPublicDomain(v string) *ListEiamInstancesResponseBodyInstances {
	s.OpenAPIPublicDomain = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetSSODomain(v string) *ListEiamInstancesResponseBodyInstances {
	s.SSODomain = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetStartTime(v int64) *ListEiamInstancesResponseBodyInstances {
	s.StartTime = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}
