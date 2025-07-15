// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportCustomCallTaggingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportCustomCallTaggingResponseBody
	GetCode() *string
	SetData(v []*ImportCustomCallTaggingResponseBodyData) *ImportCustomCallTaggingResponseBody
	GetData() []*ImportCustomCallTaggingResponseBodyData
	SetHttpStatusCode(v int32) *ImportCustomCallTaggingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ImportCustomCallTaggingResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportCustomCallTaggingResponseBody
	GetRequestId() *string
}

type ImportCustomCallTaggingResponseBody struct {
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ImportCustomCallTaggingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s ImportCustomCallTaggingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportCustomCallTaggingResponseBody) GoString() string {
	return s.String()
}

func (s *ImportCustomCallTaggingResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportCustomCallTaggingResponseBody) GetData() []*ImportCustomCallTaggingResponseBodyData {
	return s.Data
}

func (s *ImportCustomCallTaggingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ImportCustomCallTaggingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportCustomCallTaggingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportCustomCallTaggingResponseBody) SetCode(v string) *ImportCustomCallTaggingResponseBody {
	s.Code = &v
	return s
}

func (s *ImportCustomCallTaggingResponseBody) SetData(v []*ImportCustomCallTaggingResponseBodyData) *ImportCustomCallTaggingResponseBody {
	s.Data = v
	return s
}

func (s *ImportCustomCallTaggingResponseBody) SetHttpStatusCode(v int32) *ImportCustomCallTaggingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportCustomCallTaggingResponseBody) SetMessage(v string) *ImportCustomCallTaggingResponseBody {
	s.Message = &v
	return s
}

func (s *ImportCustomCallTaggingResponseBody) SetRequestId(v string) *ImportCustomCallTaggingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportCustomCallTaggingResponseBody) Validate() error {
	return dara.Validate(s)
}

type ImportCustomCallTaggingResponseBodyData struct {
	// example:
	//
	// 1312121****
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// example:
	//
	// CUSTOM_NUMBER_DUPLICATED
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ImportCustomCallTaggingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImportCustomCallTaggingResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportCustomCallTaggingResponseBodyData) GetItem() *string {
	return s.Item
}

func (s *ImportCustomCallTaggingResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *ImportCustomCallTaggingResponseBodyData) SetItem(v string) *ImportCustomCallTaggingResponseBodyData {
	s.Item = &v
	return s
}

func (s *ImportCustomCallTaggingResponseBodyData) SetReason(v string) *ImportCustomCallTaggingResponseBodyData {
	s.Reason = &v
	return s
}

func (s *ImportCustomCallTaggingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
