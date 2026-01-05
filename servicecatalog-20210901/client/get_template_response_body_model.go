// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTemplateResponseBody
	GetRequestId() *string
	SetTemplateBody(v string) *GetTemplateResponseBody
	GetTemplateBody() *string
}

type GetTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The content of the template.
	//
	// For more information about the template syntax, see [Structure of Terraform templates](https://help.aliyun.com/document_detail/184397.html).
	//
	// example:
	//
	// {
	//
	//   "ROSTemplateFormatVersion": "2015-09-01",
	//
	//   "Transform": "Aliyun::Terraform-v1.1",
	//
	//   "Workspace": {
	//
	//     "main.tf": "variable  \\"name\\" {  default = \\"auto_provisioning_group\\"}"
	//
	//   },
	//
	//   "Outputs": {}
	//
	// }
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateResponseBody) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateBody(v string) *GetTemplateResponseBody {
	s.TemplateBody = &v
	return s
}

func (s *GetTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
