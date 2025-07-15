// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomizedListHeadersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCustomizedListHeadersResponseBody
	GetRequestId() *string
}

type ModifyCustomizedListHeadersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCustomizedListHeadersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomizedListHeadersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCustomizedListHeadersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCustomizedListHeadersResponseBody) SetRequestId(v string) *ModifyCustomizedListHeadersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCustomizedListHeadersResponseBody) Validate() error {
	return dara.Validate(s)
}
