// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateApmResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateApmResponseBody
	GetResult() *bool
}

type UpdateApmResponseBody struct {
	// example:
	//
	// 18061926-CC50-5F9B-9600-034C29F1D5B0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateApmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApmResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApmResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateApmResponseBody) SetRequestId(v string) *UpdateApmResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApmResponseBody) SetResult(v bool) *UpdateApmResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateApmResponseBody) Validate() error {
	return dara.Validate(s)
}
