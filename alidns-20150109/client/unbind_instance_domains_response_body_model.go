// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindInstanceDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedCount(v int32) *UnbindInstanceDomainsResponseBody
	GetFailedCount() *int32
	SetRequestId(v string) *UnbindInstanceDomainsResponseBody
	GetRequestId() *string
	SetSuccessCount(v int32) *UnbindInstanceDomainsResponseBody
	GetSuccessCount() *int32
}

type UnbindInstanceDomainsResponseBody struct {
	// The number of domain names that failed to be unbound from the instance.
	//
	// example:
	//
	// 0
	FailedCount *int32 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 123
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of domain names that are unbound from the instance.
	//
	// example:
	//
	// 2
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s UnbindInstanceDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindInstanceDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindInstanceDomainsResponseBody) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *UnbindInstanceDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindInstanceDomainsResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *UnbindInstanceDomainsResponseBody) SetFailedCount(v int32) *UnbindInstanceDomainsResponseBody {
	s.FailedCount = &v
	return s
}

func (s *UnbindInstanceDomainsResponseBody) SetRequestId(v string) *UnbindInstanceDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindInstanceDomainsResponseBody) SetSuccessCount(v int32) *UnbindInstanceDomainsResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *UnbindInstanceDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}
