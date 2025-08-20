// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceTypeTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResourceTypeTemplateResponseBody
	GetRequestId() *string
	SetTemplateBody(v map[string]interface{}) *GetResourceTypeTemplateResponseBody
	GetTemplateBody() map[string]interface{}
	SetTemplateContent(v string) *GetResourceTypeTemplateResponseBody
	GetTemplateContent() *string
}

type GetResourceTypeTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 87F54B2B-AEF0-4C33-A72A-3F8856A575E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The structure that contains the template body. The template body must be 1 to 51,200 bytes in length. For more information, see [Template syntax](https://help.aliyun.com/document_detail/28857.html).
	//
	// > We recommend that use TemplateContent instead of TemplateBody.
	//
	// example:
	//
	// {"ROSTemplateFormatVersion": "2015-09-01"}
	TemplateBody map[string]interface{} `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The JSON-formatted structure of the template body. For more information, see [Template syntax](https://help.aliyun.com/document_detail/28857.html).
	//
	// example:
	//
	// {
	//
	//       "ROSTemplateFormatVersion": "2015-09-01"
	//
	// }
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
}

func (s GetResourceTypeTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceTypeTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceTypeTemplateResponseBody) GetTemplateBody() map[string]interface{} {
	return s.TemplateBody
}

func (s *GetResourceTypeTemplateResponseBody) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *GetResourceTypeTemplateResponseBody) SetRequestId(v string) *GetResourceTypeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceTypeTemplateResponseBody) SetTemplateBody(v map[string]interface{}) *GetResourceTypeTemplateResponseBody {
	s.TemplateBody = v
	return s
}

func (s *GetResourceTypeTemplateResponseBody) SetTemplateContent(v string) *GetResourceTypeTemplateResponseBody {
	s.TemplateContent = &v
	return s
}

func (s *GetResourceTypeTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
