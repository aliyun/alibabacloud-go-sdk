// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCardSmsTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCardSmsTemplateResponseBody
	GetCode() *string
	SetData(v *CreateCardSmsTemplateResponseBodyData) *CreateCardSmsTemplateResponseBody
	GetData() *CreateCardSmsTemplateResponseBodyData
	SetRequestId(v string) *CreateCardSmsTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCardSmsTemplateResponseBody
	GetSuccess() *bool
}

type CreateCardSmsTemplateResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *CreateCardSmsTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
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

func (s CreateCardSmsTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCardSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCardSmsTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCardSmsTemplateResponseBody) GetData() *CreateCardSmsTemplateResponseBodyData {
	return s.Data
}

func (s *CreateCardSmsTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCardSmsTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCardSmsTemplateResponseBody) SetCode(v string) *CreateCardSmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCardSmsTemplateResponseBody) SetData(v *CreateCardSmsTemplateResponseBodyData) *CreateCardSmsTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateCardSmsTemplateResponseBody) SetRequestId(v string) *CreateCardSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCardSmsTemplateResponseBody) SetSuccess(v bool) *CreateCardSmsTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCardSmsTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCardSmsTemplateResponseBodyData struct {
	// The code of the message template.
	//
	// You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm?spm=5176.12818093.categories-n-products.ddysms.3b2816d0xml2NA#/overview).
	//
	// > Make sure that the message template has been approved.
	//
	// example:
	//
	// CARD_SMS_60000****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s CreateCardSmsTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCardSmsTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCardSmsTemplateResponseBodyData) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateCardSmsTemplateResponseBodyData) SetTemplateCode(v string) *CreateCardSmsTemplateResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *CreateCardSmsTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
