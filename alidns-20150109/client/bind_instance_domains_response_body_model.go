// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindInstanceDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedCount(v int32) *BindInstanceDomainsResponseBody
	GetFailedCount() *int32
	SetRequestId(v string) *BindInstanceDomainsResponseBody
	GetRequestId() *string
	SetSuccessCount(v int32) *BindInstanceDomainsResponseBody
	GetSuccessCount() *int32
}

type BindInstanceDomainsResponseBody struct {
	// The number of domain names that failed to be bound to the instance.
	//
	// example:
	//
	// 0
	FailedCount *int32 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of domain names that are bound to the instance.
	//
	// example:
	//
	// 2
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s BindInstanceDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindInstanceDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *BindInstanceDomainsResponseBody) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *BindInstanceDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindInstanceDomainsResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *BindInstanceDomainsResponseBody) SetFailedCount(v int32) *BindInstanceDomainsResponseBody {
	s.FailedCount = &v
	return s
}

func (s *BindInstanceDomainsResponseBody) SetRequestId(v string) *BindInstanceDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindInstanceDomainsResponseBody) SetSuccessCount(v int32) *BindInstanceDomainsResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *BindInstanceDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}
