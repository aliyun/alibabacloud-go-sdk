// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iXpackRelateDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *XpackRelateDBResponseBody
	GetRequestId() *string
}

type XpackRelateDBResponseBody struct {
	// example:
	//
	// 50373857-C47B-4B64-9332-D0B5280B59EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s XpackRelateDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s XpackRelateDBResponseBody) GoString() string {
	return s.String()
}

func (s *XpackRelateDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *XpackRelateDBResponseBody) SetRequestId(v string) *XpackRelateDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *XpackRelateDBResponseBody) Validate() error {
	return dara.Validate(s)
}
