// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetBuildRiskDefineRuleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetBuildRiskDefineRuleConfigResponseBody
	GetCode() *string
	SetData(v *SetBuildRiskDefineRuleConfigResponseBodyData) *SetBuildRiskDefineRuleConfigResponseBody
	GetData() *SetBuildRiskDefineRuleConfigResponseBodyData
	SetMessage(v string) *SetBuildRiskDefineRuleConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetBuildRiskDefineRuleConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetBuildRiskDefineRuleConfigResponseBody
	GetSuccess() *bool
}

type SetBuildRiskDefineRuleConfigResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *SetBuildRiskDefineRuleConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// BA674E4B-00CF-5DEA-8B92-360862FB5133
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

func (s SetBuildRiskDefineRuleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetBuildRiskDefineRuleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetBuildRiskDefineRuleConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetBuildRiskDefineRuleConfigResponseBody) GetData() *SetBuildRiskDefineRuleConfigResponseBodyData {
	return s.Data
}

func (s *SetBuildRiskDefineRuleConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetBuildRiskDefineRuleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetBuildRiskDefineRuleConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetBuildRiskDefineRuleConfigResponseBody) SetCode(v string) *SetBuildRiskDefineRuleConfigResponseBody {
	s.Code = &v
	return s
}

func (s *SetBuildRiskDefineRuleConfigResponseBody) SetData(v *SetBuildRiskDefineRuleConfigResponseBodyData) *SetBuildRiskDefineRuleConfigResponseBody {
	s.Data = v
	return s
}

func (s *SetBuildRiskDefineRuleConfigResponseBody) SetMessage(v string) *SetBuildRiskDefineRuleConfigResponseBody {
	s.Message = &v
	return s
}

func (s *SetBuildRiskDefineRuleConfigResponseBody) SetRequestId(v string) *SetBuildRiskDefineRuleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetBuildRiskDefineRuleConfigResponseBody) SetSuccess(v bool) *SetBuildRiskDefineRuleConfigResponseBody {
	s.Success = &v
	return s
}

func (s *SetBuildRiskDefineRuleConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetBuildRiskDefineRuleConfigResponseBodyData struct {
	// The configuration ID for scanning image build command risks.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SetBuildRiskDefineRuleConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SetBuildRiskDefineRuleConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetBuildRiskDefineRuleConfigResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *SetBuildRiskDefineRuleConfigResponseBodyData) SetId(v int64) *SetBuildRiskDefineRuleConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *SetBuildRiskDefineRuleConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
