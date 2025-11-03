// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateWhiteListResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateWhiteListResponseBody
	GetRequestId() *string
}

type UpdateWhiteListResponseBody struct {
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 576EF709-71CE-500F-95FC-7F7A297D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWhiteListResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWhiteListResponseBody) SetData(v bool) *UpdateWhiteListResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateWhiteListResponseBody) SetRequestId(v string) *UpdateWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
