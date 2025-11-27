// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCardSmsTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCardSmsTemplateResponseBody
	GetCode() *string
	SetData(v *QueryCardSmsTemplateResponseBodyData) *QueryCardSmsTemplateResponseBody
	GetData() *QueryCardSmsTemplateResponseBodyData
	SetRequestId(v string) *QueryCardSmsTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCardSmsTemplateResponseBody
	GetSuccess() *bool
}

type QueryCardSmsTemplateResponseBody struct {
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryCardSmsTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
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

func (s QueryCardSmsTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCardSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCardSmsTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCardSmsTemplateResponseBody) GetData() *QueryCardSmsTemplateResponseBodyData {
	return s.Data
}

func (s *QueryCardSmsTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCardSmsTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCardSmsTemplateResponseBody) SetCode(v string) *QueryCardSmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCardSmsTemplateResponseBody) SetData(v *QueryCardSmsTemplateResponseBodyData) *QueryCardSmsTemplateResponseBody {
	s.Data = v
	return s
}

func (s *QueryCardSmsTemplateResponseBody) SetRequestId(v string) *QueryCardSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCardSmsTemplateResponseBody) SetSuccess(v bool) *QueryCardSmsTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCardSmsTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCardSmsTemplateResponseBodyData struct {
	// The array of objects.
	Templates []map[string]interface{} `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s QueryCardSmsTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCardSmsTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCardSmsTemplateResponseBodyData) GetTemplates() []map[string]interface{} {
	return s.Templates
}

func (s *QueryCardSmsTemplateResponseBodyData) SetTemplates(v []map[string]interface{}) *QueryCardSmsTemplateResponseBodyData {
	s.Templates = v
	return s
}

func (s *QueryCardSmsTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
