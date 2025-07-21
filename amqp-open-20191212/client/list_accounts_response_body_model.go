// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListAccountsResponseBody
	GetCode() *int32
	SetData(v map[string][]*DataValue) *ListAccountsResponseBody
	GetData() map[string][]*DataValue
	SetMessage(v string) *ListAccountsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAccountsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAccountsResponseBody
	GetSuccess() *bool
}

type ListAccountsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data map[string][]*DataValue `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 549A5A97-FE61-5A23-8126-3A11929C1EC4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListAccountsResponseBody) GetData() map[string][]*DataValue {
	return s.Data
}

func (s *ListAccountsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccountsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAccountsResponseBody) SetCode(v int32) *ListAccountsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAccountsResponseBody) SetData(v map[string][]*DataValue) *ListAccountsResponseBody {
	s.Data = v
	return s
}

func (s *ListAccountsResponseBody) SetMessage(v string) *ListAccountsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAccountsResponseBody) SetRequestId(v string) *ListAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountsResponseBody) SetSuccess(v bool) *ListAccountsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}
