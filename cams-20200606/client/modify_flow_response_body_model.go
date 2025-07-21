// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ModifyFlowResponseBody
	GetCode() *string
	SetData(v *ModifyFlowResponseBodyData) *ModifyFlowResponseBody
	GetData() *ModifyFlowResponseBodyData
	SetMessage(v string) *ModifyFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyFlowResponseBody
	GetRequestId() *string
}

type ModifyFlowResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// If OK is returned, the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ModifyFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1612C226-E271-4CFE-9F18-4066D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyFlowResponseBody) GetData() *ModifyFlowResponseBodyData {
	return s.Data
}

func (s *ModifyFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFlowResponseBody) SetAccessDeniedDetail(v string) *ModifyFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyFlowResponseBody) SetCode(v string) *ModifyFlowResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyFlowResponseBody) SetData(v *ModifyFlowResponseBodyData) *ModifyFlowResponseBody {
	s.Data = v
	return s
}

func (s *ModifyFlowResponseBody) SetMessage(v string) *ModifyFlowResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyFlowResponseBody) SetRequestId(v string) *ModifyFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFlowResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyFlowResponseBodyData struct {
	// The categories of the Flow.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The Flow ID.
	//
	// example:
	//
	// 3939399****
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The Flow name.
	//
	// example:
	//
	// flow-00203
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s ModifyFlowResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyFlowResponseBodyData) GetCategories() []*string {
	return s.Categories
}

func (s *ModifyFlowResponseBodyData) GetFlowId() *string {
	return s.FlowId
}

func (s *ModifyFlowResponseBodyData) GetFlowName() *string {
	return s.FlowName
}

func (s *ModifyFlowResponseBodyData) SetCategories(v []*string) *ModifyFlowResponseBodyData {
	s.Categories = v
	return s
}

func (s *ModifyFlowResponseBodyData) SetFlowId(v string) *ModifyFlowResponseBodyData {
	s.FlowId = &v
	return s
}

func (s *ModifyFlowResponseBodyData) SetFlowName(v string) *ModifyFlowResponseBodyData {
	s.FlowName = &v
	return s
}

func (s *ModifyFlowResponseBodyData) Validate() error {
	return dara.Validate(s)
}
