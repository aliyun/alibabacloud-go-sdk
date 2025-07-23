// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSubscriptionAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *SetSubscriptionAttributesResponseBody
	GetCode() *int64
	SetData(v *SetSubscriptionAttributesResponseBodyData) *SetSubscriptionAttributesResponseBody
	GetData() *SetSubscriptionAttributesResponseBodyData
	SetMessage(v string) *SetSubscriptionAttributesResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetSubscriptionAttributesResponseBody
	GetRequestId() *string
	SetStatus(v string) *SetSubscriptionAttributesResponseBody
	GetStatus() *string
	SetSuccess(v bool) *SetSubscriptionAttributesResponseBody
	GetSuccess() *bool
}

type SetSubscriptionAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *SetSubscriptionAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SetSubscriptionAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSubscriptionAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *SetSubscriptionAttributesResponseBody) GetData() *SetSubscriptionAttributesResponseBodyData {
	return s.Data
}

func (s *SetSubscriptionAttributesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetSubscriptionAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSubscriptionAttributesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SetSubscriptionAttributesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetSubscriptionAttributesResponseBody) SetCode(v int64) *SetSubscriptionAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBody) SetData(v *SetSubscriptionAttributesResponseBodyData) *SetSubscriptionAttributesResponseBody {
	s.Data = v
	return s
}

func (s *SetSubscriptionAttributesResponseBody) SetMessage(v string) *SetSubscriptionAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBody) SetRequestId(v string) *SetSubscriptionAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBody) SetStatus(v string) *SetSubscriptionAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBody) SetSuccess(v bool) *SetSubscriptionAttributesResponseBody {
	s.Success = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}

type SetSubscriptionAttributesResponseBodyData struct {
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

func (s SetSubscriptionAttributesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SetSubscriptionAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesResponseBodyData) GetCode() *int64 {
	return s.Code
}

func (s *SetSubscriptionAttributesResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *SetSubscriptionAttributesResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *SetSubscriptionAttributesResponseBodyData) SetCode(v int64) *SetSubscriptionAttributesResponseBodyData {
	s.Code = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBodyData) SetMessage(v string) *SetSubscriptionAttributesResponseBodyData {
	s.Message = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBodyData) SetSuccess(v bool) *SetSubscriptionAttributesResponseBodyData {
	s.Success = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
