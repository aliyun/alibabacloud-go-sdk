// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomCallTaggingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCustomCallTaggingResponseBody
	GetCode() *string
	SetData(v []*CreateCustomCallTaggingResponseBodyData) *CreateCustomCallTaggingResponseBody
	GetData() []*CreateCustomCallTaggingResponseBodyData
	SetHttpStatusCode(v int32) *CreateCustomCallTaggingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateCustomCallTaggingResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCustomCallTaggingResponseBody
	GetRequestId() *string
}

type CreateCustomCallTaggingResponseBody struct {
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*CreateCustomCallTaggingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomCallTaggingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomCallTaggingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomCallTaggingResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCustomCallTaggingResponseBody) GetData() []*CreateCustomCallTaggingResponseBodyData {
	return s.Data
}

func (s *CreateCustomCallTaggingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateCustomCallTaggingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCustomCallTaggingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomCallTaggingResponseBody) SetCode(v string) *CreateCustomCallTaggingResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCustomCallTaggingResponseBody) SetData(v []*CreateCustomCallTaggingResponseBodyData) *CreateCustomCallTaggingResponseBody {
	s.Data = v
	return s
}

func (s *CreateCustomCallTaggingResponseBody) SetHttpStatusCode(v int32) *CreateCustomCallTaggingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCustomCallTaggingResponseBody) SetMessage(v string) *CreateCustomCallTaggingResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCustomCallTaggingResponseBody) SetRequestId(v string) *CreateCustomCallTaggingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomCallTaggingResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateCustomCallTaggingResponseBodyData struct {
	// example:
	//
	// 1312121****
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// example:
	//
	// CUSTOM_NUMBER_DUPLICATED
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s CreateCustomCallTaggingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomCallTaggingResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCustomCallTaggingResponseBodyData) GetItem() *string {
	return s.Item
}

func (s *CreateCustomCallTaggingResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *CreateCustomCallTaggingResponseBodyData) SetItem(v string) *CreateCustomCallTaggingResponseBodyData {
	s.Item = &v
	return s
}

func (s *CreateCustomCallTaggingResponseBodyData) SetReason(v string) *CreateCustomCallTaggingResponseBodyData {
	s.Reason = &v
	return s
}

func (s *CreateCustomCallTaggingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
