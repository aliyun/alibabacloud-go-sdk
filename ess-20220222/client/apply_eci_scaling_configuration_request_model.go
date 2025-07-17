// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyEciScalingConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ApplyEciScalingConfigurationRequest
	GetContent() *string
	SetFormat(v string) *ApplyEciScalingConfigurationRequest
	GetFormat() *string
	SetRegionId(v string) *ApplyEciScalingConfigurationRequest
	GetRegionId() *string
	SetScalingConfigurationId(v string) *ApplyEciScalingConfigurationRequest
	GetScalingConfigurationId() *string
	SetScalingGroupId(v string) *ApplyEciScalingConfigurationRequest
	GetScalingGroupId() *string
}

type ApplyEciScalingConfigurationRequest struct {
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
	//   spec:
	//
	//     replicas: 3
	//
	//     selector:
	//
	//        matchLabels:
	//
	//         app: nginx
	//
	//     template:
	//
	//       metadata:
	//
	//         labels:
	//
	//           app: nginx
	//
	//         annotations:
	//
	//           k8s.aliyun.com/eip-bandwidth: 10
	//
	//           k8s.aliyun.com/eci-with-eip: true
	//
	//         spec:
	//
	//           containers:
	//
	//           - name: nginx
	//
	//             image: nginx:1.14.2
	//
	//             ports:
	//
	//             - containerPort: 80
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Optional. Set the value to YAML.
	//
	// example:
	//
	// YAML
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scaling configuration.
	//
	// If you want the system to update a scaling configuration of the Elastic Container Instance type based on a YAML configuration file, you must specify `ScalingConfigurationId`. If you do not specify `ScalingConfigurationId`, the system creates a new scaling configuration based on the YAML configuration file.
	//
	// example:
	//
	// asc-bp1i65jd06v04vdh****
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1igpak5ft1flyp****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s ApplyEciScalingConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyEciScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ApplyEciScalingConfigurationRequest) GetContent() *string {
	return s.Content
}

func (s *ApplyEciScalingConfigurationRequest) GetFormat() *string {
	return s.Format
}

func (s *ApplyEciScalingConfigurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ApplyEciScalingConfigurationRequest) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *ApplyEciScalingConfigurationRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *ApplyEciScalingConfigurationRequest) SetContent(v string) *ApplyEciScalingConfigurationRequest {
	s.Content = &v
	return s
}

func (s *ApplyEciScalingConfigurationRequest) SetFormat(v string) *ApplyEciScalingConfigurationRequest {
	s.Format = &v
	return s
}

func (s *ApplyEciScalingConfigurationRequest) SetRegionId(v string) *ApplyEciScalingConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyEciScalingConfigurationRequest) SetScalingConfigurationId(v string) *ApplyEciScalingConfigurationRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *ApplyEciScalingConfigurationRequest) SetScalingGroupId(v string) *ApplyEciScalingConfigurationRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ApplyEciScalingConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
