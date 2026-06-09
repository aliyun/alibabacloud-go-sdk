// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListNamespacesResponseBody
	GetCode() *string
	SetData(v *ListNamespacesResponseBodyData) *ListNamespacesResponseBody
	GetData() *ListNamespacesResponseBodyData
	SetMessage(v string) *ListNamespacesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListNamespacesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListNamespacesResponseBody
	GetSuccess() *bool
}

type ListNamespacesResponseBody struct {
	// example:
	//
	// 200
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListNamespacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListNamespacesResponseBody) GetData() *ListNamespacesResponseBodyData {
	return s.Data
}

func (s *ListNamespacesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNamespacesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListNamespacesResponseBody) SetCode(v string) *ListNamespacesResponseBody {
	s.Code = &v
	return s
}

func (s *ListNamespacesResponseBody) SetData(v *ListNamespacesResponseBodyData) *ListNamespacesResponseBody {
	s.Data = v
	return s
}

func (s *ListNamespacesResponseBody) SetMessage(v string) *ListNamespacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListNamespacesResponseBody) SetRequestId(v string) *ListNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNamespacesResponseBody) SetSuccess(v bool) *ListNamespacesResponseBody {
	s.Success = &v
	return s
}

func (s *ListNamespacesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNamespacesResponseBodyData struct {
	Namespaces []*Namespace `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListNamespacesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBodyData) GetNamespaces() []*Namespace {
	return s.Namespaces
}

func (s *ListNamespacesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNamespacesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListNamespacesResponseBodyData) SetNamespaces(v []*Namespace) *ListNamespacesResponseBodyData {
	s.Namespaces = v
	return s
}

func (s *ListNamespacesResponseBodyData) SetNextToken(v string) *ListNamespacesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListNamespacesResponseBodyData) SetTotal(v int32) *ListNamespacesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListNamespacesResponseBodyData) Validate() error {
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
