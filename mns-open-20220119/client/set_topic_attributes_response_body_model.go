// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTopicAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *SetTopicAttributesResponseBody
	GetCode() *int64
	SetData(v *SetTopicAttributesResponseBodyData) *SetTopicAttributesResponseBody
	GetData() *SetTopicAttributesResponseBodyData
	SetMessage(v string) *SetTopicAttributesResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetTopicAttributesResponseBody
	GetRequestId() *string
	SetStatus(v string) *SetTopicAttributesResponseBody
	GetStatus() *string
	SetSuccess(v bool) *SetTopicAttributesResponseBody
	GetSuccess() *bool
}

type SetTopicAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *SetTopicAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetTopicAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetTopicAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *SetTopicAttributesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *SetTopicAttributesResponseBody) GetData() *SetTopicAttributesResponseBodyData {
	return s.Data
}

func (s *SetTopicAttributesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetTopicAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetTopicAttributesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SetTopicAttributesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetTopicAttributesResponseBody) SetCode(v int64) *SetTopicAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *SetTopicAttributesResponseBody) SetData(v *SetTopicAttributesResponseBodyData) *SetTopicAttributesResponseBody {
	s.Data = v
	return s
}

func (s *SetTopicAttributesResponseBody) SetMessage(v string) *SetTopicAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *SetTopicAttributesResponseBody) SetRequestId(v string) *SetTopicAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetTopicAttributesResponseBody) SetStatus(v string) *SetTopicAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *SetTopicAttributesResponseBody) SetSuccess(v bool) *SetTopicAttributesResponseBody {
	s.Success = &v
	return s
}

func (s *SetTopicAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}

type SetTopicAttributesResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetTopicAttributesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SetTopicAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetTopicAttributesResponseBodyData) GetCode() *int64 {
	return s.Code
}

func (s *SetTopicAttributesResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *SetTopicAttributesResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *SetTopicAttributesResponseBodyData) SetCode(v int64) *SetTopicAttributesResponseBodyData {
	s.Code = &v
	return s
}

func (s *SetTopicAttributesResponseBodyData) SetMessage(v string) *SetTopicAttributesResponseBodyData {
	s.Message = &v
	return s
}

func (s *SetTopicAttributesResponseBodyData) SetSuccess(v bool) *SetTopicAttributesResponseBodyData {
	s.Success = &v
	return s
}

func (s *SetTopicAttributesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
