// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCardSmsTemplateReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCardSmsTemplateReportResponseBody
	GetCode() *string
	SetData(v *QueryCardSmsTemplateReportResponseBodyData) *QueryCardSmsTemplateReportResponseBody
	GetData() *QueryCardSmsTemplateReportResponseBodyData
	SetRequestId(v string) *QueryCardSmsTemplateReportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCardSmsTemplateReportResponseBody
	GetSuccess() *bool
}

type QueryCardSmsTemplateReportResponseBody struct {
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryCardSmsTemplateReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CC89A90C-978F-46AC-B80D-54738371E7CA
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

func (s QueryCardSmsTemplateReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCardSmsTemplateReportResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCardSmsTemplateReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCardSmsTemplateReportResponseBody) GetData() *QueryCardSmsTemplateReportResponseBodyData {
	return s.Data
}

func (s *QueryCardSmsTemplateReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCardSmsTemplateReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCardSmsTemplateReportResponseBody) SetCode(v string) *QueryCardSmsTemplateReportResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCardSmsTemplateReportResponseBody) SetData(v *QueryCardSmsTemplateReportResponseBodyData) *QueryCardSmsTemplateReportResponseBody {
	s.Data = v
	return s
}

func (s *QueryCardSmsTemplateReportResponseBody) SetRequestId(v string) *QueryCardSmsTemplateReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCardSmsTemplateReportResponseBody) SetSuccess(v bool) *QueryCardSmsTemplateReportResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCardSmsTemplateReportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCardSmsTemplateReportResponseBodyData struct {
	// The details of the data returned.
	Model []map[string]interface{} `json:"model,omitempty" xml:"model,omitempty" type:"Repeated"`
}

func (s QueryCardSmsTemplateReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCardSmsTemplateReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCardSmsTemplateReportResponseBodyData) GetModel() []map[string]interface{} {
	return s.Model
}

func (s *QueryCardSmsTemplateReportResponseBodyData) SetModel(v []map[string]interface{}) *QueryCardSmsTemplateReportResponseBodyData {
	s.Model = v
	return s
}

func (s *QueryCardSmsTemplateReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}
