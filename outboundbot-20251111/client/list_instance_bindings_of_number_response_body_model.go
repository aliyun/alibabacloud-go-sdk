// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceBindingsOfNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstanceBindingsOfNumberResponseBody
	GetCode() *string
	SetData(v []*ListInstanceBindingsOfNumberResponseBodyData) *ListInstanceBindingsOfNumberResponseBody
	GetData() []*ListInstanceBindingsOfNumberResponseBodyData
	SetHttpStatusCode(v int32) *ListInstanceBindingsOfNumberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListInstanceBindingsOfNumberResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListInstanceBindingsOfNumberResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListInstanceBindingsOfNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstanceBindingsOfNumberResponseBody
	GetSuccess() *bool
}

type ListInstanceBindingsOfNumberResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of instance binding information.
	Data []*ListInstanceBindingsOfNumberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the API.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstanceBindingsOfNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceBindingsOfNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceBindingsOfNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstanceBindingsOfNumberResponseBody) GetData() []*ListInstanceBindingsOfNumberResponseBodyData {
	return s.Data
}

func (s *ListInstanceBindingsOfNumberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstanceBindingsOfNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstanceBindingsOfNumberResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListInstanceBindingsOfNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceBindingsOfNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstanceBindingsOfNumberResponseBody) SetCode(v string) *ListInstanceBindingsOfNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceBindingsOfNumberResponseBody) SetData(v []*ListInstanceBindingsOfNumberResponseBodyData) *ListInstanceBindingsOfNumberResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceBindingsOfNumberResponseBody) SetHttpStatusCode(v int32) *ListInstanceBindingsOfNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstanceBindingsOfNumberResponseBody) SetMessage(v string) *ListInstanceBindingsOfNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceBindingsOfNumberResponseBody) SetParams(v []*string) *ListInstanceBindingsOfNumberResponseBody {
	s.Params = v
	return s
}

func (s *ListInstanceBindingsOfNumberResponseBody) SetRequestId(v string) *ListInstanceBindingsOfNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceBindingsOfNumberResponseBody) SetSuccess(v bool) *ListInstanceBindingsOfNumberResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstanceBindingsOfNumberResponseBody) Validate() error {
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

type ListInstanceBindingsOfNumberResponseBodyData struct {
	// Indicates whether the instance is bound to the number specified in the request.
	//
	// example:
	//
	// false
	Bound *bool `json:"Bound,omitempty" xml:"Bound,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 7f04f92c-ccfc-4f8f-a816-6902023be5c6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// Instance example
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s ListInstanceBindingsOfNumberResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceBindingsOfNumberResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceBindingsOfNumberResponseBodyData) GetBound() *bool {
	return s.Bound
}

func (s *ListInstanceBindingsOfNumberResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceBindingsOfNumberResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstanceBindingsOfNumberResponseBodyData) SetBound(v bool) *ListInstanceBindingsOfNumberResponseBodyData {
	s.Bound = &v
	return s
}

func (s *ListInstanceBindingsOfNumberResponseBodyData) SetInstanceId(v string) *ListInstanceBindingsOfNumberResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceBindingsOfNumberResponseBodyData) SetInstanceName(v string) *ListInstanceBindingsOfNumberResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListInstanceBindingsOfNumberResponseBodyData) Validate() error {
	return dara.Validate(s)
}
