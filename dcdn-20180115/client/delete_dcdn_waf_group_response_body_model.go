// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnWafGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDcdnWafGroupResponseBody
	GetRequestId() *string
}

type DeleteDcdnWafGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2430E05E-1340-5773-B5E1-B743929F46F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnWafGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnWafGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnWafGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDcdnWafGroupResponseBody) SetRequestId(v string) *DeleteDcdnWafGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDcdnWafGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
