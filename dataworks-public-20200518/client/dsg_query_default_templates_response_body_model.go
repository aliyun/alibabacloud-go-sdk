// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgQueryDefaultTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DsgQueryDefaultTemplatesResponseBodyData) *DsgQueryDefaultTemplatesResponseBody
	GetData() []*DsgQueryDefaultTemplatesResponseBodyData
	SetErrorCode(v string) *DsgQueryDefaultTemplatesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgQueryDefaultTemplatesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgQueryDefaultTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgQueryDefaultTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgQueryDefaultTemplatesResponseBody
	GetSuccess() *bool
}

type DsgQueryDefaultTemplatesResponseBody struct {
	// The data returned.
	Data []*DsgQueryDefaultTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// 1029030003
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// param error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 102400001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DsgQueryDefaultTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgQueryDefaultTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DsgQueryDefaultTemplatesResponseBody) GetData() []*DsgQueryDefaultTemplatesResponseBodyData {
	return s.Data
}

func (s *DsgQueryDefaultTemplatesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgQueryDefaultTemplatesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgQueryDefaultTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgQueryDefaultTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgQueryDefaultTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgQueryDefaultTemplatesResponseBody) SetData(v []*DsgQueryDefaultTemplatesResponseBodyData) *DsgQueryDefaultTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *DsgQueryDefaultTemplatesResponseBody) SetErrorCode(v string) *DsgQueryDefaultTemplatesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgQueryDefaultTemplatesResponseBody) SetErrorMessage(v string) *DsgQueryDefaultTemplatesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgQueryDefaultTemplatesResponseBody) SetHttpStatusCode(v int32) *DsgQueryDefaultTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgQueryDefaultTemplatesResponseBody) SetRequestId(v string) *DsgQueryDefaultTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgQueryDefaultTemplatesResponseBody) SetSuccess(v bool) *DsgQueryDefaultTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *DsgQueryDefaultTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DsgQueryDefaultTemplatesResponseBodyData struct {
	// The sensitive field type.
	//
	// example:
	//
	// phone
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The supported data masking methods and parameter descriptions.
	DesensPlanTemplate map[string][]*DataDesensPlanTemplateValue `json:"DesensPlanTemplate,omitempty" xml:"DesensPlanTemplate,omitempty"`
}

func (s DsgQueryDefaultTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DsgQueryDefaultTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DsgQueryDefaultTemplatesResponseBodyData) GetDataType() *string {
	return s.DataType
}

func (s *DsgQueryDefaultTemplatesResponseBodyData) GetDesensPlanTemplate() map[string][]*DataDesensPlanTemplateValue {
	return s.DesensPlanTemplate
}

func (s *DsgQueryDefaultTemplatesResponseBodyData) SetDataType(v string) *DsgQueryDefaultTemplatesResponseBodyData {
	s.DataType = &v
	return s
}

func (s *DsgQueryDefaultTemplatesResponseBodyData) SetDesensPlanTemplate(v map[string][]*DataDesensPlanTemplateValue) *DsgQueryDefaultTemplatesResponseBodyData {
	s.DesensPlanTemplate = v
	return s
}

func (s *DsgQueryDefaultTemplatesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
