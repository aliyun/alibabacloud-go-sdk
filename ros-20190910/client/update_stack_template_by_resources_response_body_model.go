// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStackTemplateByResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNewTemplateBody(v string) *UpdateStackTemplateByResourcesResponseBody
	GetNewTemplateBody() *string
	SetOldTemplateBody(v string) *UpdateStackTemplateByResourcesResponseBody
	GetOldTemplateBody() *string
	SetRequestId(v string) *UpdateStackTemplateByResourcesResponseBody
	GetRequestId() *string
}

type UpdateStackTemplateByResourcesResponseBody struct {
	// The template content after correction.
	//
	// example:
	//
	// {\\"ROSTemplateFormatVersion\\": \\"2015-09-01\\", \\"Resources\\": {\\"Vpc\\": {\\"Type\\": \\"ALIYUN::ECS::VPC\\", \\"Properties\\": {\\"VpcName\\": \\"test\\", \\"CidrBlock\\": \\"192.168.0.0/16\\", \\"Description\\": \\"test2\\"}}}, \\"Outputs\\": {\\"VpcId\\": {\\"Value\\": {\\"Fn::GetAtt\\": [\\"Vpc\\", \\"VpcId\\"]}}}}
	NewTemplateBody *string `json:"NewTemplateBody,omitempty" xml:"NewTemplateBody,omitempty"`
	// The template content before correction.
	//
	// example:
	//
	// {\\"ROSTemplateFormatVersion\\": \\"2015-09-01\\", \\"Resources\\": {\\"Vpc\\": {\\"Type\\": \\"ALIYUN::ECS::VPC\\", \\"Properties\\": {\\"VpcName\\": \\"test\\", \\"CidrBlock\\": \\"192.168.0.0/16\\", \\"Description\\": \\"test1\\"}}}, \\"Outputs\\": {\\"VpcId\\": {\\"Value\\": {\\"Fn::GetAtt\\": [\\"Vpc\\", \\"VpcId\\"]}}}}
	OldTemplateBody *string `json:"OldTemplateBody,omitempty" xml:"OldTemplateBody,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateStackTemplateByResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackTemplateByResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStackTemplateByResourcesResponseBody) GetNewTemplateBody() *string {
	return s.NewTemplateBody
}

func (s *UpdateStackTemplateByResourcesResponseBody) GetOldTemplateBody() *string {
	return s.OldTemplateBody
}

func (s *UpdateStackTemplateByResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStackTemplateByResourcesResponseBody) SetNewTemplateBody(v string) *UpdateStackTemplateByResourcesResponseBody {
	s.NewTemplateBody = &v
	return s
}

func (s *UpdateStackTemplateByResourcesResponseBody) SetOldTemplateBody(v string) *UpdateStackTemplateByResourcesResponseBody {
	s.OldTemplateBody = &v
	return s
}

func (s *UpdateStackTemplateByResourcesResponseBody) SetRequestId(v string) *UpdateStackTemplateByResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStackTemplateByResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
