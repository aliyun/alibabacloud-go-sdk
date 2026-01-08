// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateFlowResponseBody
	GetCode() *string
	SetData(v *CreateFlowResponseBodyData) *CreateFlowResponseBody
	GetData() *CreateFlowResponseBodyData
	SetMessage(v string) *CreateFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateFlowResponseBody
	GetRequestId() *string
}

type CreateFlowResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// If OK is returned, the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateFlowResponseBody) GetData() *CreateFlowResponseBodyData {
	return s.Data
}

func (s *CreateFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFlowResponseBody) SetAccessDeniedDetail(v string) *CreateFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateFlowResponseBody) SetCode(v string) *CreateFlowResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFlowResponseBody) SetData(v *CreateFlowResponseBodyData) *CreateFlowResponseBody {
	s.Data = v
	return s
}

func (s *CreateFlowResponseBody) SetMessage(v string) *CreateFlowResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFlowResponseBody) SetRequestId(v string) *CreateFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateFlowResponseBodyData struct {
	// The categories of the Flow.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The Flow ID.
	//
	// example:
	//
	// 333993838***
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The name of the Flow.
	//
	// example:
	//
	// test1
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s CreateFlowResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateFlowResponseBodyData) GetCategories() []*string {
	return s.Categories
}

func (s *CreateFlowResponseBodyData) GetFlowId() *string {
	return s.FlowId
}

func (s *CreateFlowResponseBodyData) GetFlowName() *string {
	return s.FlowName
}

func (s *CreateFlowResponseBodyData) SetCategories(v []*string) *CreateFlowResponseBodyData {
	s.Categories = v
	return s
}

func (s *CreateFlowResponseBodyData) SetFlowId(v string) *CreateFlowResponseBodyData {
	s.FlowId = &v
	return s
}

func (s *CreateFlowResponseBodyData) SetFlowName(v string) *CreateFlowResponseBodyData {
	s.FlowName = &v
	return s
}

func (s *CreateFlowResponseBodyData) Validate() error {
	return dara.Validate(s)
}
