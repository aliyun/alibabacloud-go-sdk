// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventsAttentionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEventsAttentionResponseBody
	GetRequestId() *string
	SetCode(v string) *UpdateEventsAttentionResponseBody
	GetCode() *string
	SetData(v *UpdateEventsAttentionResponseBodyData) *UpdateEventsAttentionResponseBody
	GetData() *UpdateEventsAttentionResponseBodyData
	SetMessage(v string) *UpdateEventsAttentionResponseBody
	GetMessage() *string
}

type UpdateEventsAttentionResponseBody struct {
	// example:
	//
	// 44841312-7227-55C9-AE03-D59729BFAE38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateEventsAttentionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// Instance not belong to current user
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateEventsAttentionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventsAttentionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventsAttentionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEventsAttentionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateEventsAttentionResponseBody) GetData() *UpdateEventsAttentionResponseBodyData {
	return s.Data
}

func (s *UpdateEventsAttentionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateEventsAttentionResponseBody) SetRequestId(v string) *UpdateEventsAttentionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventsAttentionResponseBody) SetCode(v string) *UpdateEventsAttentionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEventsAttentionResponseBody) SetData(v *UpdateEventsAttentionResponseBodyData) *UpdateEventsAttentionResponseBody {
	s.Data = v
	return s
}

func (s *UpdateEventsAttentionResponseBody) SetMessage(v string) *UpdateEventsAttentionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEventsAttentionResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateEventsAttentionResponseBodyData struct {
	// example:
	//
	// 1
	Mode *int32 `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s UpdateEventsAttentionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventsAttentionResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateEventsAttentionResponseBodyData) GetMode() *int32 {
	return s.Mode
}

func (s *UpdateEventsAttentionResponseBodyData) SetMode(v int32) *UpdateEventsAttentionResponseBodyData {
	s.Mode = &v
	return s
}

func (s *UpdateEventsAttentionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
