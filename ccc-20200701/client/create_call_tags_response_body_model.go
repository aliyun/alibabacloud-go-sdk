// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCallTagsResponseBody
	GetCode() *string
	SetData(v []*CreateCallTagsResponseBodyData) *CreateCallTagsResponseBody
	GetData() []*CreateCallTagsResponseBodyData
	SetHttpStatusCode(v int32) *CreateCallTagsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateCallTagsResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCallTagsResponseBody
	GetRequestId() *string
}

type CreateCallTagsResponseBody struct {
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*CreateCallTagsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCallTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCallTagsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCallTagsResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCallTagsResponseBody) GetData() []*CreateCallTagsResponseBodyData {
	return s.Data
}

func (s *CreateCallTagsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateCallTagsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCallTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCallTagsResponseBody) SetCode(v string) *CreateCallTagsResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCallTagsResponseBody) SetData(v []*CreateCallTagsResponseBodyData) *CreateCallTagsResponseBody {
	s.Data = v
	return s
}

func (s *CreateCallTagsResponseBody) SetHttpStatusCode(v int32) *CreateCallTagsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCallTagsResponseBody) SetMessage(v string) *CreateCallTagsResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCallTagsResponseBody) SetRequestId(v string) *CreateCallTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCallTagsResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateCallTagsResponseBodyData struct {
	// example:
	//
	// TagC
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// example:
	//
	// CALL_TAG_NAME_DUPLICATED
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s CreateCallTagsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCallTagsResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCallTagsResponseBodyData) GetItem() *string {
	return s.Item
}

func (s *CreateCallTagsResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *CreateCallTagsResponseBodyData) SetItem(v string) *CreateCallTagsResponseBodyData {
	s.Item = &v
	return s
}

func (s *CreateCallTagsResponseBodyData) SetReason(v string) *CreateCallTagsResponseBodyData {
	s.Reason = &v
	return s
}

func (s *CreateCallTagsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
