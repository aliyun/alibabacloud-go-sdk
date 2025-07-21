// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListFlowResponseBody
	GetCode() *string
	SetData(v []*ListFlowResponseBodyData) *ListFlowResponseBody
	GetData() []*ListFlowResponseBodyData
	SetMessage(v string) *ListFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFlowResponseBody
	GetRequestId() *string
}

type ListFlowResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// If OK is returned, the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*ListFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 1612C226-E271-4CFE-9F18-4066D550F91B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFlowResponseBody) GetData() []*ListFlowResponseBodyData {
	return s.Data
}

func (s *ListFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlowResponseBody) SetAccessDeniedDetail(v string) *ListFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListFlowResponseBody) SetCode(v string) *ListFlowResponseBody {
	s.Code = &v
	return s
}

func (s *ListFlowResponseBody) SetData(v []*ListFlowResponseBodyData) *ListFlowResponseBody {
	s.Data = v
	return s
}

func (s *ListFlowResponseBody) SetMessage(v string) *ListFlowResponseBody {
	s.Message = &v
	return s
}

func (s *ListFlowResponseBody) SetRequestId(v string) *ListFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFlowResponseBodyData struct {
	// The categories of the Flows.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The Flow ID.
	//
	// example:
	//
	// 3939393***
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The Flow name.
	//
	// example:
	//
	// flow-02020
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s ListFlowResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFlowResponseBodyData) GetCategories() []*string {
	return s.Categories
}

func (s *ListFlowResponseBodyData) GetFlowId() *string {
	return s.FlowId
}

func (s *ListFlowResponseBodyData) GetFlowName() *string {
	return s.FlowName
}

func (s *ListFlowResponseBodyData) SetCategories(v []*string) *ListFlowResponseBodyData {
	s.Categories = v
	return s
}

func (s *ListFlowResponseBodyData) SetFlowId(v string) *ListFlowResponseBodyData {
	s.FlowId = &v
	return s
}

func (s *ListFlowResponseBodyData) SetFlowName(v string) *ListFlowResponseBodyData {
	s.FlowName = &v
	return s
}

func (s *ListFlowResponseBodyData) Validate() error {
	return dara.Validate(s)
}
