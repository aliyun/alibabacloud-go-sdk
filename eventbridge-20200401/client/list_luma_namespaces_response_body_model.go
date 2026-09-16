// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLumaNamespacesResponseBody
	GetCode() *string
	SetData(v *ListLumaNamespacesResponseBodyData) *ListLumaNamespacesResponseBody
	GetData() *ListLumaNamespacesResponseBodyData
	SetMessage(v string) *ListLumaNamespacesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLumaNamespacesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLumaNamespacesResponseBody
	GetSuccess() *bool
}

type ListLumaNamespacesResponseBody struct {
	// The response code. A value of Success indicates a successful call. A specific error code is returned upon failure.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of namespaces bound to the Agent. All results are returned at once without pagination.
	Data *ListLumaNamespacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned by the operation. The value is Operation success when the call succeeds, or a specific error description when the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, used for troubleshooting and ticket feedback.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLumaNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLumaNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLumaNamespacesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLumaNamespacesResponseBody) GetData() *ListLumaNamespacesResponseBodyData {
	return s.Data
}

func (s *ListLumaNamespacesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLumaNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLumaNamespacesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLumaNamespacesResponseBody) SetCode(v string) *ListLumaNamespacesResponseBody {
	s.Code = &v
	return s
}

func (s *ListLumaNamespacesResponseBody) SetData(v *ListLumaNamespacesResponseBodyData) *ListLumaNamespacesResponseBody {
	s.Data = v
	return s
}

func (s *ListLumaNamespacesResponseBody) SetMessage(v string) *ListLumaNamespacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListLumaNamespacesResponseBody) SetRequestId(v string) *ListLumaNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLumaNamespacesResponseBody) SetSuccess(v bool) *ListLumaNamespacesResponseBody {
	s.Success = &v
	return s
}

func (s *ListLumaNamespacesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLumaNamespacesResponseBodyData struct {
	// The list of namespaces bound to the Agent.
	//
	// example:
	//
	// [{"Name":"my_namespace"}]
	Namespaces []*Namespace `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
}

func (s ListLumaNamespacesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLumaNamespacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLumaNamespacesResponseBodyData) GetNamespaces() []*Namespace {
	return s.Namespaces
}

func (s *ListLumaNamespacesResponseBodyData) SetNamespaces(v []*Namespace) *ListLumaNamespacesResponseBodyData {
	s.Namespaces = v
	return s
}

func (s *ListLumaNamespacesResponseBodyData) Validate() error {
	if s.Namespaces != nil {
		for _, item := range s.Namespaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
