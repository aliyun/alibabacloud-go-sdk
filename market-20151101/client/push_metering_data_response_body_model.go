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
	SetSuccess(v bool) *PushMeteringDataResponseBody
	GetSuccess() *bool
}

type PushMeteringDataResponseBody struct {
	// example:
	//
	// 6EF60BEC-0242-43AF-BB20-270359FB54A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *PushMeteringDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PushMeteringDataResponseBody) SetRequestId(v string) *PushMeteringDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetSuccess(v bool) *PushMeteringDataResponseBody {
	s.Success = &v
	return s
}

func (s *PushMeteringDataResponseBody) Validate() error {
	return dara.Validate(s)
}
