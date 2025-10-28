// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTemplateByScratchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GenerateTemplateByScratchResponseBody
	GetRequestId() *string
	SetResourcesToImport(v []*GenerateTemplateByScratchResponseBodyResourcesToImport) *GenerateTemplateByScratchResponseBody
	GetResourcesToImport() []*GenerateTemplateByScratchResponseBodyResourcesToImport
	SetTemplateBody(v string) *GenerateTemplateByScratchResponseBody
	GetTemplateBody() *string
}

type GenerateTemplateByScratchResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D22C1D13-D74E-558C-AF68-1B4C05FA6F1A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources that you want to import into a stack in the resource management scenario.
	//
	// > This parameter is returned only for a resource management scenario.
	ResourcesToImport []*GenerateTemplateByScratchResponseBodyResourcesToImport `json:"ResourcesToImport,omitempty" xml:"ResourcesToImport,omitempty" type:"Repeated"`
	// The template content of the resource scenario.
	//
	// example:
	//
	// {\\"ROSTemplateFormatVersion\\": \\"2015-09-01\\", \\"Resources\\": {\\"ECSVPC_001\\": {\\"DeletionPolicy\\": \\"Retain\\", \\"Type\\": \\"ALIYUN::ECS::VPC\\", \\"Properties\\": {\\"CidrBlock\\": \\"172.16.0.0/12\\", \\"VpcName\\": \\"MyTestVpc\\", \\"EnableIpv6\\": false}}}}
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
}

func (s GenerateTemplateByScratchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateTemplateByScratchResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTemplateByScratchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateTemplateByScratchResponseBody) GetResourcesToImport() []*GenerateTemplateByScratchResponseBodyResourcesToImport {
	return s.ResourcesToImport
}

func (s *GenerateTemplateByScratchResponseBody) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *GenerateTemplateByScratchResponseBody) SetRequestId(v string) *GenerateTemplateByScratchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateTemplateByScratchResponseBody) SetResourcesToImport(v []*GenerateTemplateByScratchResponseBodyResourcesToImport) *GenerateTemplateByScratchResponseBody {
	s.ResourcesToImport = v
	return s
}

func (s *GenerateTemplateByScratchResponseBody) SetTemplateBody(v string) *GenerateTemplateByScratchResponseBody {
	s.TemplateBody = &v
	return s
}

func (s *GenerateTemplateByScratchResponseBody) Validate() error {
	if s.ResourcesToImport != nil {
		for _, item := range s.ResourcesToImport {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GenerateTemplateByScratchResponseBodyResourcesToImport struct {
	// The logical ID of the resource.
	//
	// example:
	//
	// ECSVPC_001
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The key-value mapping between strings. The value is a JSON string that identifies the resource that you want to import into a stack.\\
	//
	// A key is an identifier for a resource, and a value is an assignment of data to the key. For example, VpcId is a key that indicates the ID of a virtual private cloud (VPC), and `vpc-bp1m6fww66xbntjyc****"` is a value that is assigned to VpcId.
	//
	// example:
	//
	// {"VpcId": "vpc-bp1m6fww66xbntjyc****" }
	ResourceIdentifier map[string]interface{} `json:"ResourceIdentifier,omitempty" xml:"ResourceIdentifier,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ALIYUN::ECS::VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GenerateTemplateByScratchResponseBodyResourcesToImport) String() string {
	return dara.Prettify(s)
}

func (s GenerateTemplateByScratchResponseBodyResourcesToImport) GoString() string {
	return s.String()
}

func (s *GenerateTemplateByScratchResponseBodyResourcesToImport) GetLogicalResourceId() *string {
	return s.LogicalResourceId
}

func (s *GenerateTemplateByScratchResponseBodyResourcesToImport) GetResourceIdentifier() map[string]interface{} {
	return s.ResourceIdentifier
}

func (s *GenerateTemplateByScratchResponseBodyResourcesToImport) GetResourceType() *string {
	return s.ResourceType
}

func (s *GenerateTemplateByScratchResponseBodyResourcesToImport) SetLogicalResourceId(v string) *GenerateTemplateByScratchResponseBodyResourcesToImport {
	s.LogicalResourceId = &v
	return s
}

func (s *GenerateTemplateByScratchResponseBodyResourcesToImport) SetResourceIdentifier(v map[string]interface{}) *GenerateTemplateByScratchResponseBodyResourcesToImport {
	s.ResourceIdentifier = v
	return s
}

func (s *GenerateTemplateByScratchResponseBodyResourcesToImport) SetResourceType(v string) *GenerateTemplateByScratchResponseBodyResourcesToImport {
	s.ResourceType = &v
	return s
}

func (s *GenerateTemplateByScratchResponseBodyResourcesToImport) Validate() error {
	return dara.Validate(s)
}
