// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNumberBindingsOfInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListNumberBindingsOfInstanceResponseBody
	GetCode() *string
	SetData(v []*ListNumberBindingsOfInstanceResponseBodyData) *ListNumberBindingsOfInstanceResponseBody
	GetData() []*ListNumberBindingsOfInstanceResponseBodyData
	SetHttpStatusCode(v int32) *ListNumberBindingsOfInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListNumberBindingsOfInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListNumberBindingsOfInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListNumberBindingsOfInstanceResponseBody
	GetSuccess() *bool
}

type ListNumberBindingsOfInstanceResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of number binding information.
	Data []*ListNumberBindingsOfInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the API operation.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNumberBindingsOfInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNumberBindingsOfInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListNumberBindingsOfInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListNumberBindingsOfInstanceResponseBody) GetData() []*ListNumberBindingsOfInstanceResponseBodyData {
	return s.Data
}

func (s *ListNumberBindingsOfInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListNumberBindingsOfInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNumberBindingsOfInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNumberBindingsOfInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListNumberBindingsOfInstanceResponseBody) SetCode(v string) *ListNumberBindingsOfInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ListNumberBindingsOfInstanceResponseBody) SetData(v []*ListNumberBindingsOfInstanceResponseBodyData) *ListNumberBindingsOfInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ListNumberBindingsOfInstanceResponseBody) SetHttpStatusCode(v int32) *ListNumberBindingsOfInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListNumberBindingsOfInstanceResponseBody) SetMessage(v string) *ListNumberBindingsOfInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ListNumberBindingsOfInstanceResponseBody) SetRequestId(v string) *ListNumberBindingsOfInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNumberBindingsOfInstanceResponseBody) SetSuccess(v bool) *ListNumberBindingsOfInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ListNumberBindingsOfInstanceResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNumberBindingsOfInstanceResponseBodyData struct {
	// Indicates whether the number is bound to the instance specified in the request.
	//
	// example:
	//
	// true
	Bound *bool `json:"Bound,omitempty" xml:"Bound,omitempty"`
	// The bound number.
	//
	// example:
	//
	// 02180542230
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The SIP trunk ID.
	//
	// example:
	//
	// 020ffd7f-e40a-4969-acc3-3e1d80870e3f
	SipTrunkId *string `json:"SipTrunkId,omitempty" xml:"SipTrunkId,omitempty"`
	// The SIP trunk name.
	//
	// example:
	//
	// Sample trunk
	SipTrunkName *string `json:"SipTrunkName,omitempty" xml:"SipTrunkName,omitempty"`
}

func (s ListNumberBindingsOfInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNumberBindingsOfInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNumberBindingsOfInstanceResponseBodyData) GetBound() *bool {
	return s.Bound
}

func (s *ListNumberBindingsOfInstanceResponseBodyData) GetNumber() *string {
	return s.Number
}

func (s *ListNumberBindingsOfInstanceResponseBodyData) GetSipTrunkId() *string {
	return s.SipTrunkId
}

func (s *ListNumberBindingsOfInstanceResponseBodyData) GetSipTrunkName() *string {
	return s.SipTrunkName
}

func (s *ListNumberBindingsOfInstanceResponseBodyData) SetBound(v bool) *ListNumberBindingsOfInstanceResponseBodyData {
	s.Bound = &v
	return s
}

func (s *ListNumberBindingsOfInstanceResponseBodyData) SetNumber(v string) *ListNumberBindingsOfInstanceResponseBodyData {
	s.Number = &v
	return s
}

func (s *ListNumberBindingsOfInstanceResponseBodyData) SetSipTrunkId(v string) *ListNumberBindingsOfInstanceResponseBodyData {
	s.SipTrunkId = &v
	return s
}

func (s *ListNumberBindingsOfInstanceResponseBodyData) SetSipTrunkName(v string) *ListNumberBindingsOfInstanceResponseBodyData {
	s.SipTrunkName = &v
	return s
}

func (s *ListNumberBindingsOfInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
