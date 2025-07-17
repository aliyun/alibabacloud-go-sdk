// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyScalingGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ApplyScalingGroupRequest
	GetContent() *string
	SetFormat(v string) *ApplyScalingGroupRequest
	GetFormat() *string
	SetRegionId(v string) *ApplyScalingGroupRequest
	GetRegionId() *string
}

type ApplyScalingGroupRequest struct {
	// The content of the configuration file.
	//
	// This parameter is required.
	//
	// example:
	//
	// apiVersion: apps/v1
	//
	// kind: Deployment
	//
	// metadata:
	//
	//   name: nginx-deployment
	//
	//   labels:
	//
	//     app: nginx
	//
	// spec:
	//
	//   replicas: 3
	//
	//   selector:
	//
	//     matchLabels:
	//
	//       app: nginx
	//
	//   template:
	//
	//     metadata:
	//
	//       labels:
	//
	//         app: nginx
	//
	//       annotations:
	//
	//         k8s.aliyun.com/eip-bandwidth: 10
	//
	//         k8s.aliyun.com/eci-with-eip: true
	//
	//     spec:
	//
	//       containers:
	//
	//       - name: nginx
	//
	//         image: nginx:1.14.2
	//
	//         ports:
	//
	//         - containerPort: 80
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Optional. The format of the configuration file. Default value: YAML. Set the value to YAML.
	//
	// example:
	//
	// YAML
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ApplyScalingGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *ApplyScalingGroupRequest) GetContent() *string {
	return s.Content
}

func (s *ApplyScalingGroupRequest) GetFormat() *string {
	return s.Format
}

func (s *ApplyScalingGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ApplyScalingGroupRequest) SetContent(v string) *ApplyScalingGroupRequest {
	s.Content = &v
	return s
}

func (s *ApplyScalingGroupRequest) SetFormat(v string) *ApplyScalingGroupRequest {
	s.Format = &v
	return s
}

func (s *ApplyScalingGroupRequest) SetRegionId(v string) *ApplyScalingGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyScalingGroupRequest) Validate() error {
	return dara.Validate(s)
}
