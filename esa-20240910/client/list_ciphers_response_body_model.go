// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCiphersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCiphersGroup(v string) *ListCiphersResponseBody
	GetCiphersGroup() *string
	SetRequestId(v string) *ListCiphersResponseBody
	GetRequestId() *string
	SetResult(v []*string) *ListCiphersResponseBody
	GetResult() []*string
	SetTotalCount(v int64) *ListCiphersResponseBody
	GetTotalCount() *int64
}

type ListCiphersResponseBody struct {
	// Name of the cipher suite group.
	//
	// example:
	//
	// all
	CiphersGroup *string `json:"CiphersGroup,omitempty" xml:"CiphersGroup,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result.
	Result []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Total number of cipher suites.
	//
	// example:
	//
	// 16
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCiphersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCiphersResponseBody) GoString() string {
	return s.String()
}

func (s *ListCiphersResponseBody) GetCiphersGroup() *string {
	return s.CiphersGroup
}

func (s *ListCiphersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCiphersResponseBody) GetResult() []*string {
	return s.Result
}

func (s *ListCiphersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCiphersResponseBody) SetCiphersGroup(v string) *ListCiphersResponseBody {
	s.CiphersGroup = &v
	return s
}

func (s *ListCiphersResponseBody) SetRequestId(v string) *ListCiphersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCiphersResponseBody) SetResult(v []*string) *ListCiphersResponseBody {
	s.Result = v
	return s
}

func (s *ListCiphersResponseBody) SetTotalCount(v int64) *ListCiphersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCiphersResponseBody) Validate() error {
	return dara.Validate(s)
}
