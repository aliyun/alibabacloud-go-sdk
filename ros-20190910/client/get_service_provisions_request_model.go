// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceProvisionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParameters(v []*GetServiceProvisionsRequestParameters) *GetServiceProvisionsRequest
	GetParameters() []*GetServiceProvisionsRequestParameters
	SetRegionId(v string) *GetServiceProvisionsRequest
	GetRegionId() *string
	SetServices(v []*GetServiceProvisionsRequestServices) *GetServiceProvisionsRequest
	GetServices() []*GetServiceProvisionsRequestServices
	SetTemplateBody(v string) *GetServiceProvisionsRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *GetServiceProvisionsRequest
	GetTemplateId() *string
	SetTemplateURL(v string) *GetServiceProvisionsRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *GetServiceProvisionsRequest
	GetTemplateVersion() *string
}

type GetServiceProvisionsRequest struct {
	// The parameters.
	Parameters []*GetServiceProvisionsRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The services.
	Services     []*GetServiceProvisionsRequestServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	TemplateBody *string                                `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The template ID. This parameter applies to shared and private templates.
	//
	// You must and can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and Services.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body must be 1 to 524,288 bytes in length. If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// You must and can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and Services.
	//
	// example:
	//
	// oss://ros-template/demo
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. If you do not specify this parameter, the latest version is used.
	//
	// This parameter takes effect only when TemplateId is specified.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetServiceProvisionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvisionsRequest) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsRequest) GetParameters() []*GetServiceProvisionsRequestParameters {
	return s.Parameters
}

func (s *GetServiceProvisionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceProvisionsRequest) GetServices() []*GetServiceProvisionsRequestServices {
	return s.Services
}

func (s *GetServiceProvisionsRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *GetServiceProvisionsRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetServiceProvisionsRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *GetServiceProvisionsRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetServiceProvisionsRequest) SetParameters(v []*GetServiceProvisionsRequestParameters) *GetServiceProvisionsRequest {
	s.Parameters = v
	return s
}

func (s *GetServiceProvisionsRequest) SetRegionId(v string) *GetServiceProvisionsRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetServices(v []*GetServiceProvisionsRequestServices) *GetServiceProvisionsRequest {
	s.Services = v
	return s
}

func (s *GetServiceProvisionsRequest) SetTemplateBody(v string) *GetServiceProvisionsRequest {
	s.TemplateBody = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetTemplateId(v string) *GetServiceProvisionsRequest {
	s.TemplateId = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetTemplateURL(v string) *GetServiceProvisionsRequest {
	s.TemplateURL = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetTemplateVersion(v string) *GetServiceProvisionsRequest {
	s.TemplateVersion = &v
	return s
}

func (s *GetServiceProvisionsRequest) Validate() error {
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetServiceProvisionsRequestParameters struct {
	// The name of the parameter. If you do not specify the name and value of a parameter, Resource Orchestration Service (ROS) uses the default name and value that are specified in the template.
	//
	// > The Parameters parameter is optional. If you specify Parameters, you must specify ParameterKey.
	//
	// This parameter is required.
	//
	// example:
	//
	// Amount
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter.
	//
	// > The Parameters parameter is optional. If you specify Parameters, you must specify ParameterValue.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetServiceProvisionsRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvisionsRequestParameters) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetServiceProvisionsRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GetServiceProvisionsRequestParameters) SetParameterKey(v string) *GetServiceProvisionsRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetServiceProvisionsRequestParameters) SetParameterValue(v string) *GetServiceProvisionsRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetServiceProvisionsRequestParameters) Validate() error {
	return dara.Validate(s)
}

type GetServiceProvisionsRequestServices struct {
	// The name of the service or feature. Valid values:
	//
	// 	- AHAS: Application High Availability Service
	//
	// 	- ARMS: Application Real-Time Monitoring Service (ARMS)
	//
	// 	- ApiGateway: API Gateway
	//
	// 	- BatchCompute: Batch Compute
	//
	// 	- BrainIndustrial: Industrial Brain
	//
	// 	- CloudStorageGateway: Cloud Storage Gateway (CSG)
	//
	// 	- CMS: CloudMonitor
	//
	// 	- CR: Container Registry
	//
	// 	- CS: Container Service for Kubernetes (ACK)
	//
	// 	- DCDN: Dynamic Content Delivery Network (DCDN)
	//
	// 	- DataHub: DataHub
	//
	// 	- DataWorks: DataWorks
	//
	// 	- EDAS: Enterprise Distributed Application Service (EDAS)
	//
	// 	- EHPC: E-HPC
	//
	// 	- EMAS: Enterprise Mobile Application Studio (EMAS)
	//
	// 	- FC: Function Compute
	//
	// 	- FNF: CloudFlow (SWF)
	//
	// 	- MaxCompute: MaxCompute
	//
	// 	- MNS: Message Service (MNS)
	//
	// 	- HBR: Cloud Backup
	//
	// 	- IMM: Intelligent Media Management (IMM)
	//
	// 	- IOT: IoT Platform
	//
	// 	- KMS: Key Management Service (KMS)
	//
	// 	- NAS: File Storage NAS (NAS)
	//
	// 	- NLP: Natural Language Processing (NLP)
	//
	// 	- OSS: Object Storage Service (OSS)
	//
	// 	- OTS: Tablestore
	//
	// 	- PrivateLink: PrivateLink
	//
	// 	- PrivateZone: Alibaba Cloud DNS PrivateZone
	//
	// 	- RocketMQ: ApsaraMQ for RocketMQ
	//
	// 	- SAE: Serverless App Engine (SAE)
	//
	// 	- SLS: Simple Log Service (SLS)
	//
	// 	- TrafficMirror: traffic mirroring
	//
	// 	- VS: Video Surveillance System
	//
	// 	- Xtrace: Managed Service for OpenTelemetry
	//
	// This parameter is required.
	//
	// example:
	//
	// EHPC
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetServiceProvisionsRequestServices) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvisionsRequestServices) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsRequestServices) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceProvisionsRequestServices) SetServiceName(v string) *GetServiceProvisionsRequestServices {
	s.ServiceName = &v
	return s
}

func (s *GetServiceProvisionsRequestServices) Validate() error {
	return dara.Validate(s)
}
