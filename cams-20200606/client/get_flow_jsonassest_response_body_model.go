// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowJSONAssestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetFlowJSONAssestResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetFlowJSONAssestResponseBody
	GetCode() *string
	SetData(v *GetFlowJSONAssestResponseBodyData) *GetFlowJSONAssestResponseBody
	GetData() *GetFlowJSONAssestResponseBodyData
	SetMessage(v string) *GetFlowJSONAssestResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFlowJSONAssestResponseBody
	GetRequestId() *string
}

type GetFlowJSONAssestResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// If OK is returned, the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetFlowJSONAssestResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error description information.
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

func (s GetFlowJSONAssestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFlowJSONAssestResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowJSONAssestResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetFlowJSONAssestResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFlowJSONAssestResponseBody) GetData() *GetFlowJSONAssestResponseBodyData {
	return s.Data
}

func (s *GetFlowJSONAssestResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFlowJSONAssestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFlowJSONAssestResponseBody) SetAccessDeniedDetail(v string) *GetFlowJSONAssestResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetFlowJSONAssestResponseBody) SetCode(v string) *GetFlowJSONAssestResponseBody {
	s.Code = &v
	return s
}

func (s *GetFlowJSONAssestResponseBody) SetData(v *GetFlowJSONAssestResponseBodyData) *GetFlowJSONAssestResponseBody {
	s.Data = v
	return s
}

func (s *GetFlowJSONAssestResponseBody) SetMessage(v string) *GetFlowJSONAssestResponseBody {
	s.Message = &v
	return s
}

func (s *GetFlowJSONAssestResponseBody) SetRequestId(v string) *GetFlowJSONAssestResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFlowJSONAssestResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFlowJSONAssestResponseBodyData struct {
	// The file path.
	//
	// example:
	//
	// https://url.com/json.json
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The Flow ID.
	//
	// example:
	//
	// flow_id_arms
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s GetFlowJSONAssestResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFlowJSONAssestResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFlowJSONAssestResponseBodyData) GetFilePath() *string {
	return s.FilePath
}

func (s *GetFlowJSONAssestResponseBodyData) GetFlowId() *string {
	return s.FlowId
}

func (s *GetFlowJSONAssestResponseBodyData) SetFilePath(v string) *GetFlowJSONAssestResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *GetFlowJSONAssestResponseBodyData) SetFlowId(v string) *GetFlowJSONAssestResponseBodyData {
	s.FlowId = &v
	return s
}

func (s *GetFlowJSONAssestResponseBodyData) Validate() error {
	return dara.Validate(s)
}
