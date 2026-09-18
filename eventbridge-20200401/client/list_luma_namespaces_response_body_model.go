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
	// The response code. A value of Success indicates a successful operation. If the operation fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of namespaces bound to the agent, including entries and pagination information.
	Data *ListLumaNamespacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message. A value of Operation success is returned if the operation succeeds. A specific error description is returned if the operation fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request. Use this ID for troubleshooting and when submitting a ticket.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success.
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
	// The effective page size for this request. If the Limit parameter is not specified, the server default value is used. If the specified value exceeds the upper limit, the value is adjusted to the maximum allowed value.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The list of namespaces bound to the agent.
	//
	// example:
	//
	// [{"Name":"my_namespace"}]
	Namespaces []*Namespace `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The token for the next page. Pass this value as the NextToken parameter in the next request to retrieve the next page. An empty value indicates that no more data is available.
	//
	// example:
	//
	// 10
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of namespaces bound to the agent, regardless of the number of entries returned on the current page.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLumaNamespacesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLumaNamespacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLumaNamespacesResponseBodyData) GetLimit() *int32 {
	return s.Limit
}

func (s *ListLumaNamespacesResponseBodyData) GetNamespaces() []*Namespace {
	return s.Namespaces
}

func (s *ListLumaNamespacesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLumaNamespacesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLumaNamespacesResponseBodyData) SetLimit(v int32) *ListLumaNamespacesResponseBodyData {
	s.Limit = &v
	return s
}

func (s *ListLumaNamespacesResponseBodyData) SetNamespaces(v []*Namespace) *ListLumaNamespacesResponseBodyData {
	s.Namespaces = v
	return s
}

func (s *ListLumaNamespacesResponseBodyData) SetNextToken(v string) *ListLumaNamespacesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListLumaNamespacesResponseBodyData) SetTotalCount(v int32) *ListLumaNamespacesResponseBodyData {
	s.TotalCount = &v
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
