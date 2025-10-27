// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaStrategyTemplateSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOpaStrategyTemplateSummaryResponseBody
	GetCode() *string
	SetData(v []*GetOpaStrategyTemplateSummaryResponseBodyData) *GetOpaStrategyTemplateSummaryResponseBody
	GetData() []*GetOpaStrategyTemplateSummaryResponseBodyData
	SetMessage(v string) *GetOpaStrategyTemplateSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOpaStrategyTemplateSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOpaStrategyTemplateSummaryResponseBody
	GetSuccess() *bool
}

type GetOpaStrategyTemplateSummaryResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The usage statistics about the templates provided for rules of the at-risk image blocking type.
	Data []*GetOpaStrategyTemplateSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54572138-3390-5774-B71D-799DC8C2161B
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

func (s GetOpaStrategyTemplateSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyTemplateSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyTemplateSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOpaStrategyTemplateSummaryResponseBody) GetData() []*GetOpaStrategyTemplateSummaryResponseBodyData {
	return s.Data
}

func (s *GetOpaStrategyTemplateSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOpaStrategyTemplateSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpaStrategyTemplateSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOpaStrategyTemplateSummaryResponseBody) SetCode(v string) *GetOpaStrategyTemplateSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetOpaStrategyTemplateSummaryResponseBody) SetData(v []*GetOpaStrategyTemplateSummaryResponseBodyData) *GetOpaStrategyTemplateSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetOpaStrategyTemplateSummaryResponseBody) SetMessage(v string) *GetOpaStrategyTemplateSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetOpaStrategyTemplateSummaryResponseBody) SetRequestId(v string) *GetOpaStrategyTemplateSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpaStrategyTemplateSummaryResponseBody) SetSuccess(v bool) *GetOpaStrategyTemplateSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetOpaStrategyTemplateSummaryResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOpaStrategyTemplateSummaryResponseBodyData struct {
	// The number of times that the template is used.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// Custom defense configuration
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// Blank template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetOpaStrategyTemplateSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyTemplateSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyTemplateSummaryResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *GetOpaStrategyTemplateSummaryResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetOpaStrategyTemplateSummaryResponseBodyData) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetOpaStrategyTemplateSummaryResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetOpaStrategyTemplateSummaryResponseBodyData) SetCount(v int32) *GetOpaStrategyTemplateSummaryResponseBodyData {
	s.Count = &v
	return s
}

func (s *GetOpaStrategyTemplateSummaryResponseBodyData) SetDescription(v string) *GetOpaStrategyTemplateSummaryResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetOpaStrategyTemplateSummaryResponseBodyData) SetTemplateId(v int64) *GetOpaStrategyTemplateSummaryResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetOpaStrategyTemplateSummaryResponseBodyData) SetTemplateName(v string) *GetOpaStrategyTemplateSummaryResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *GetOpaStrategyTemplateSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
