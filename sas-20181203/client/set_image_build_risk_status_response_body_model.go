// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetImageBuildRiskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetImageBuildRiskStatusResponseBody
	GetCode() *string
	SetData(v *SetImageBuildRiskStatusResponseBodyData) *SetImageBuildRiskStatusResponseBody
	GetData() *SetImageBuildRiskStatusResponseBodyData
	SetMessage(v string) *SetImageBuildRiskStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetImageBuildRiskStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetImageBuildRiskStatusResponseBody
	GetSuccess() *bool
}

type SetImageBuildRiskStatusResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *SetImageBuildRiskStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 52870893-48A7-5A9E-9E05-6253E5B6****
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

func (s SetImageBuildRiskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetImageBuildRiskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetImageBuildRiskStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetImageBuildRiskStatusResponseBody) GetData() *SetImageBuildRiskStatusResponseBodyData {
	return s.Data
}

func (s *SetImageBuildRiskStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetImageBuildRiskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetImageBuildRiskStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetImageBuildRiskStatusResponseBody) SetCode(v string) *SetImageBuildRiskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *SetImageBuildRiskStatusResponseBody) SetData(v *SetImageBuildRiskStatusResponseBodyData) *SetImageBuildRiskStatusResponseBody {
	s.Data = v
	return s
}

func (s *SetImageBuildRiskStatusResponseBody) SetMessage(v string) *SetImageBuildRiskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *SetImageBuildRiskStatusResponseBody) SetRequestId(v string) *SetImageBuildRiskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetImageBuildRiskStatusResponseBody) SetSuccess(v bool) *SetImageBuildRiskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *SetImageBuildRiskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type SetImageBuildRiskStatusResponseBodyData struct {
	// The ID of the rule.
	//
	// example:
	//
	// 273698***
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SetImageBuildRiskStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SetImageBuildRiskStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetImageBuildRiskStatusResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *SetImageBuildRiskStatusResponseBodyData) SetId(v int64) *SetImageBuildRiskStatusResponseBodyData {
	s.Id = &v
	return s
}

func (s *SetImageBuildRiskStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
