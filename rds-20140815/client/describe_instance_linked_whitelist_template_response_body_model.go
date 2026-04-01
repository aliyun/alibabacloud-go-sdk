// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceLinkedWhitelistTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeInstanceLinkedWhitelistTemplateResponseBody
	GetCode() *string
	SetData(v *DescribeInstanceLinkedWhitelistTemplateResponseBodyData) *DescribeInstanceLinkedWhitelistTemplateResponseBody
	GetData() *DescribeInstanceLinkedWhitelistTemplateResponseBodyData
	SetHttpStatusCode(v int32) *DescribeInstanceLinkedWhitelistTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeInstanceLinkedWhitelistTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeInstanceLinkedWhitelistTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstanceLinkedWhitelistTemplateResponseBody
	GetSuccess() *bool
}

type DescribeInstanceLinkedWhitelistTemplateResponseBody struct {
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
	Data *DescribeInstanceLinkedWhitelistTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1E43AAE0-BEE8-43DA-860D-EAF2AA0724DC
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

func (s DescribeInstanceLinkedWhitelistTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceLinkedWhitelistTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBody) GetData() *DescribeInstanceLinkedWhitelistTemplateResponseBodyData {
	return s.Data
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBody) SetCode(v string) *DescribeInstanceLinkedWhitelistTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBody) SetData(v *DescribeInstanceLinkedWhitelistTemplateResponseBodyData) *DescribeInstanceLinkedWhitelistTemplateResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBody) SetHttpStatusCode(v int32) *DescribeInstanceLinkedWhitelistTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBody) SetMessage(v string) *DescribeInstanceLinkedWhitelistTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBody) SetRequestId(v string) *DescribeInstanceLinkedWhitelistTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBody) SetSuccess(v bool) *DescribeInstanceLinkedWhitelistTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceLinkedWhitelistTemplateResponseBodyData struct {
	// The instance name.
	//
	// example:
	//
	// rm-bp191w771kd3****
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// The information about whitelists that are returned by page.
	Templates []*DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s DescribeInstanceLinkedWhitelistTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceLinkedWhitelistTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBodyData) GetInsName() *string {
	return s.InsName
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBodyData) GetTemplates() []*DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates {
	return s.Templates
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBodyData) SetInsName(v string) *DescribeInstanceLinkedWhitelistTemplateResponseBodyData {
	s.InsName = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBodyData) SetTemplates(v []*DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates) *DescribeInstanceLinkedWhitelistTemplateResponseBodyData {
	s.Templates = v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBodyData) Validate() error {
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates struct {
	// The primary key of the data table.
	//
	// example:
	//
	// 1884
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IP addresses.
	//
	// example:
	//
	// 12.0.X.X,10.2.X.X
	Ips *string `json:"Ips,omitempty" xml:"Ips,omitempty"`
	// The whitelist template ID.
	//
	// example:
	//
	// 412
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The whitelist template name.
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

func (s DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates) GetId() *int32 {
	return s.Id
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates) GetIps() *string {
	return s.Ips
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates) GetUserId() *int32 {
	return s.UserId
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates) SetId(v int32) *DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates {
	s.Id = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates) SetIps(v string) *DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates {
	s.Ips = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates) SetTemplateId(v int32) *DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates {
	s.TemplateId = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates) SetTemplateName(v string) *DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates {
	s.TemplateName = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates) SetUserId(v int32) *DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates {
	s.UserId = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponseBodyDataTemplates) Validate() error {
	return dara.Validate(s)
}
