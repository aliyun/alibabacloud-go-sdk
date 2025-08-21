// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCollectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCollectorResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteCollectorResponseBody
	GetResult() *bool
}

type DeleteCollectorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the shipper is deleted. Valid values:
	//
	// 	- true: The shipper is deleted.
	//
	// 	- false: The shipper fails to be deleted.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteCollectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCollectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCollectorResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteCollectorResponseBody) SetRequestId(v string) *DeleteCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCollectorResponseBody) SetResult(v bool) *DeleteCollectorResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteCollectorResponseBody) Validate() error {
	return dara.Validate(s)
}
