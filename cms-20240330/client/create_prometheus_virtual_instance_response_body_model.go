// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusVirtualInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstance(v *CreatePrometheusVirtualInstanceResponseBodyInstance) *CreatePrometheusVirtualInstanceResponseBody
	GetInstance() *CreatePrometheusVirtualInstanceResponseBodyInstance
	SetRequestId(v string) *CreatePrometheusVirtualInstanceResponseBody
	GetRequestId() *string
}

type CreatePrometheusVirtualInstanceResponseBody struct {
	Instance *CreatePrometheusVirtualInstanceResponseBodyInstance `json:"instance,omitempty" xml:"instance,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-************
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePrometheusVirtualInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusVirtualInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrometheusVirtualInstanceResponseBody) GetInstance() *CreatePrometheusVirtualInstanceResponseBodyInstance {
	return s.Instance
}

func (s *CreatePrometheusVirtualInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrometheusVirtualInstanceResponseBody) SetInstance(v *CreatePrometheusVirtualInstanceResponseBodyInstance) *CreatePrometheusVirtualInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *CreatePrometheusVirtualInstanceResponseBody) SetRequestId(v string) *CreatePrometheusVirtualInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrometheusVirtualInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreatePrometheusVirtualInstanceResponseBodyInstance struct {
	// example:
	//
	// 1751520976660
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// http://xxxxxxx
	HttpApiUrl *string `json:"httpApiUrl,omitempty" xml:"httpApiUrl,omitempty"`
	// example:
	//
	// rw-e815960b4c9ebc5c3d89790c7e82
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// ack-csi-fuse
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// cn-zhengzhou-jva
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 167212345678
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreatePrometheusVirtualInstanceResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusVirtualInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *CreatePrometheusVirtualInstanceResponseBodyInstance) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreatePrometheusVirtualInstanceResponseBodyInstance) GetHttpApiUrl() *string {
	return s.HttpApiUrl
}

func (s *CreatePrometheusVirtualInstanceResponseBodyInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreatePrometheusVirtualInstanceResponseBodyInstance) GetNamespace() *string {
	return s.Namespace
}

func (s *CreatePrometheusVirtualInstanceResponseBodyInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePrometheusVirtualInstanceResponseBodyInstance) GetUserId() *string {
	return s.UserId
}

func (s *CreatePrometheusVirtualInstanceResponseBodyInstance) SetCreatedAt(v string) *CreatePrometheusVirtualInstanceResponseBodyInstance {
	s.CreatedAt = &v
	return s
}

func (s *CreatePrometheusVirtualInstanceResponseBodyInstance) SetHttpApiUrl(v string) *CreatePrometheusVirtualInstanceResponseBodyInstance {
	s.HttpApiUrl = &v
	return s
}

func (s *CreatePrometheusVirtualInstanceResponseBodyInstance) SetInstanceId(v string) *CreatePrometheusVirtualInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *CreatePrometheusVirtualInstanceResponseBodyInstance) SetNamespace(v string) *CreatePrometheusVirtualInstanceResponseBodyInstance {
	s.Namespace = &v
	return s
}

func (s *CreatePrometheusVirtualInstanceResponseBodyInstance) SetRegionId(v string) *CreatePrometheusVirtualInstanceResponseBodyInstance {
	s.RegionId = &v
	return s
}

func (s *CreatePrometheusVirtualInstanceResponseBodyInstance) SetUserId(v string) *CreatePrometheusVirtualInstanceResponseBodyInstance {
	s.UserId = &v
	return s
}

func (s *CreatePrometheusVirtualInstanceResponseBodyInstance) Validate() error {
	return dara.Validate(s)
}
