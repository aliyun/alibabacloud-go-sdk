// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetQueueAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *SetQueueAttributesResponseBody
	GetCode() *int64
	SetData(v *SetQueueAttributesResponseBodyData) *SetQueueAttributesResponseBody
	GetData() *SetQueueAttributesResponseBodyData
	SetMessage(v string) *SetQueueAttributesResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetQueueAttributesResponseBody
	GetRequestId() *string
	SetStatus(v string) *SetQueueAttributesResponseBody
	GetStatus() *string
	SetSuccess(v bool) *SetQueueAttributesResponseBody
	GetSuccess() *bool
}

type SetQueueAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *SetQueueAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SetQueueAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetQueueAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *SetQueueAttributesResponseBody) GetData() *SetQueueAttributesResponseBodyData {
	return s.Data
}

func (s *SetQueueAttributesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetQueueAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetQueueAttributesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SetQueueAttributesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetQueueAttributesResponseBody) SetCode(v int64) *SetQueueAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *SetQueueAttributesResponseBody) SetData(v *SetQueueAttributesResponseBodyData) *SetQueueAttributesResponseBody {
	s.Data = v
	return s
}

func (s *SetQueueAttributesResponseBody) SetMessage(v string) *SetQueueAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *SetQueueAttributesResponseBody) SetRequestId(v string) *SetQueueAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetQueueAttributesResponseBody) SetStatus(v string) *SetQueueAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *SetQueueAttributesResponseBody) SetSuccess(v bool) *SetQueueAttributesResponseBody {
	s.Success = &v
	return s
}

func (s *SetQueueAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}

type SetQueueAttributesResponseBodyData struct {
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

func (s SetQueueAttributesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SetQueueAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesResponseBodyData) GetCode() *int64 {
	return s.Code
}

func (s *SetQueueAttributesResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *SetQueueAttributesResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *SetQueueAttributesResponseBodyData) SetCode(v int64) *SetQueueAttributesResponseBodyData {
	s.Code = &v
	return s
}

func (s *SetQueueAttributesResponseBodyData) SetMessage(v string) *SetQueueAttributesResponseBodyData {
	s.Message = &v
	return s
}

func (s *SetQueueAttributesResponseBodyData) SetSuccess(v bool) *SetQueueAttributesResponseBodyData {
	s.Success = &v
	return s
}

func (s *SetQueueAttributesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
