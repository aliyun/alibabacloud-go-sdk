// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAutoGroupingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableAutoGroupingResponseBody
	GetRequestId() *string
}

type DisableAutoGroupingResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 898FAB24-7509-43EE-A287-086FE4C44394
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAutoGroupingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoGroupingResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAutoGroupingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableAutoGroupingResponseBody) SetRequestId(v string) *DisableAutoGroupingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAutoGroupingResponseBody) Validate() error {
	return dara.Validate(s)
}
