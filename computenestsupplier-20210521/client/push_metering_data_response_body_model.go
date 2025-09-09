// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushMeteringDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PushMeteringDataResponseBody
	GetRequestId() *string
}

type PushMeteringDataResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 94E89857-B994-44B6-9C4F-DBD200E9XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushMeteringDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushMeteringDataResponseBody) GoString() string {
	return s.String()
}

func (s *PushMeteringDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushMeteringDataResponseBody) SetRequestId(v string) *PushMeteringDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushMeteringDataResponseBody) Validate() error {
	return dara.Validate(s)
}
