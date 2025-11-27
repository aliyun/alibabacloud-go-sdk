// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhitelistTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeWhitelistTemplateResponseBody
	GetCode() *string
	SetData(v *DescribeWhitelistTemplateResponseBodyData) *DescribeWhitelistTemplateResponseBody
	GetData() *DescribeWhitelistTemplateResponseBodyData
	SetHttpStatusCode(v int32) *DescribeWhitelistTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeWhitelistTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeWhitelistTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeWhitelistTemplateResponseBody
	GetSuccess() *bool
}

type DescribeWhitelistTemplateResponseBody struct {
	// The response code returned. Valid values:
	//
	// 	- **200**: success
	//
	// 	- **400**: client error
	//
	// 	- **401**: identity authentication failed
	//
	// 	- **404**: request page not found
	//
	// 	- **500**: server error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribeWhitelistTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned. Valid values:
	//
	// 	- **200**: success
	//
	// 	- **400**: client error
	//
	// 	- **500**: server error
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ED169A3E-1657-4104-82AB-24EA8CD0DB75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeWhitelistTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeWhitelistTemplateResponseBody) GetData() *DescribeWhitelistTemplateResponseBodyData {
	return s.Data
}

func (s *DescribeWhitelistTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeWhitelistTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeWhitelistTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWhitelistTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeWhitelistTemplateResponseBody) SetCode(v string) *DescribeWhitelistTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeWhitelistTemplateResponseBody) SetData(v *DescribeWhitelistTemplateResponseBodyData) *DescribeWhitelistTemplateResponseBody {
	s.Data = v
	return s
}

func (s *DescribeWhitelistTemplateResponseBody) SetHttpStatusCode(v int32) *DescribeWhitelistTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeWhitelistTemplateResponseBody) SetMessage(v string) *DescribeWhitelistTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeWhitelistTemplateResponseBody) SetRequestId(v string) *DescribeWhitelistTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhitelistTemplateResponseBody) SetSuccess(v bool) *DescribeWhitelistTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeWhitelistTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeWhitelistTemplateResponseBodyData struct {
	// The information about the IP whitelist template.
	Template *DescribeWhitelistTemplateResponseBodyDataTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s DescribeWhitelistTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistTemplateResponseBodyData) GetTemplate() *DescribeWhitelistTemplateResponseBodyDataTemplate {
	return s.Template
}

func (s *DescribeWhitelistTemplateResponseBodyData) SetTemplate(v *DescribeWhitelistTemplateResponseBodyDataTemplate) *DescribeWhitelistTemplateResponseBodyData {
	s.Template = v
	return s
}

func (s *DescribeWhitelistTemplateResponseBodyData) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeWhitelistTemplateResponseBodyDataTemplate struct {
	// The primary key of the data table.
	//
	// example:
	//
	// 1013
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IP addresses.
	//
	// example:
	//
	// 10.1.X.X,2.3.X.X
	Ips *string `json:"Ips,omitempty" xml:"Ips,omitempty"`
	// The ID of the whitelist template.
	//
	// example:
	//
	// 424
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the IP whitelist template.
	//
	// example:
	//
	// template_123
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 16****
	UserId *int32 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeWhitelistTemplateResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistTemplateResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistTemplateResponseBodyDataTemplate) GetId() *int32 {
	return s.Id
}

func (s *DescribeWhitelistTemplateResponseBodyDataTemplate) GetIps() *string {
	return s.Ips
}

func (s *DescribeWhitelistTemplateResponseBodyDataTemplate) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *DescribeWhitelistTemplateResponseBodyDataTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeWhitelistTemplateResponseBodyDataTemplate) GetUserId() *int32 {
	return s.UserId
}

func (s *DescribeWhitelistTemplateResponseBodyDataTemplate) SetId(v int32) *DescribeWhitelistTemplateResponseBodyDataTemplate {
	s.Id = &v
	return s
}

func (s *DescribeWhitelistTemplateResponseBodyDataTemplate) SetIps(v string) *DescribeWhitelistTemplateResponseBodyDataTemplate {
	s.Ips = &v
	return s
}

func (s *DescribeWhitelistTemplateResponseBodyDataTemplate) SetTemplateId(v int32) *DescribeWhitelistTemplateResponseBodyDataTemplate {
	s.TemplateId = &v
	return s
}

func (s *DescribeWhitelistTemplateResponseBodyDataTemplate) SetTemplateName(v string) *DescribeWhitelistTemplateResponseBodyDataTemplate {
	s.TemplateName = &v
	return s
}

func (s *DescribeWhitelistTemplateResponseBodyDataTemplate) SetUserId(v int32) *DescribeWhitelistTemplateResponseBodyDataTemplate {
	s.UserId = &v
	return s
}

func (s *DescribeWhitelistTemplateResponseBodyDataTemplate) Validate() error {
	return dara.Validate(s)
}
