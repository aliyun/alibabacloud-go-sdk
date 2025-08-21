// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartCollectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartCollectorResponseBody
	GetRequestId() *string
	SetResult(v bool) *RestartCollectorResponseBody
	GetResult() *bool
}

type RestartCollectorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 84B4038A-AF38-4BF4-9FAD-EA92A4FFF00A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the shipper is restarted. Valid values:
	//
	// 	- true: The shipper is restarted.
	//
	// 	- false: The shipper fails to be restarted.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RestartCollectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *RestartCollectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartCollectorResponseBody) GetResult() *bool {
	return s.Result
}

func (s *RestartCollectorResponseBody) SetRequestId(v string) *RestartCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartCollectorResponseBody) SetResult(v bool) *RestartCollectorResponseBody {
	s.Result = &v
	return s
}

func (s *RestartCollectorResponseBody) Validate() error {
	return dara.Validate(s)
}
