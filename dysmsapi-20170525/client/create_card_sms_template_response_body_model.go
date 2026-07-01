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
	// The request status code. Valid values:
	//
	// - OK: The request was successful.
	//
	// - For a list of other error codes, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned by the operation.
	Data *CreateCardSmsTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
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
	// The code for the card SMS template. You can view the **Template Code*	- on the **Card SMS*	- > [template management](https://dysms.console.aliyun.com/domestic/card) page in the console.
	//
	// > The card SMS template must be approved before it can be used.
	//
	// example:
	//
	// CARD_SMS_2****
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
