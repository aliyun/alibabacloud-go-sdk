// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLimitationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListLimitationsResponseBody
	GetRequestId() *string
}

type ListLimitationsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 7CE283B7-50EC-5ABD-9EE5-FC94C38BE37F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListLimitationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLimitationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLimitationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLimitationsResponseBody) SetRequestId(v string) *ListLimitationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLimitationsResponseBody) Validate() error {
	return dara.Validate(s)
}
